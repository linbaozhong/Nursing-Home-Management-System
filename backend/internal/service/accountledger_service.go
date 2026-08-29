package service

import (
	"context"
	"errors"
	"fmt"

	"api/internal/constant"
	"api/internal/lib"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/define/table/tblelderaccountledger"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type accountLedger struct{}

var AccountLedger = &accountLedger{}

// PageAccountLedgerByKey 分页查询老人资金明细台账（联表 elder 获取老人姓名）
func (l *accountLedger) PageAccountLedgerByKey(ctx context.Context, in *dto.PageAccountLedgerByKeyReq, out *[]dto.PageAccountLedgerByKeyResp) error {
	q := db.Table(tblelderaccountledger.TableName).
		LeftJoin(tblelderaccountledger.ElderId, tblelder.Id).
		Where(tblelderaccountledger.TenantId.Eq(types.BigInt(lib.TenantID(ctx))))
	if in.ElderName != nil && *in.ElderName != "" {
		q.And(tblelder.Name.Like(*in.ElderName))
	}
	if in.Direction != nil && *in.Direction != 0 {
		q.And(tblelderaccountledger.Direction.Eq(types.Int8(*in.Direction)))
	}
	if in.StartTime != nil && !in.StartTime.IsZero() {
		q.And(tblelderaccountledger.CreateTime.Gte(*in.StartTime))
	}
	if in.EndTime != nil && !in.EndTime.IsZero() {
		q.And(tblelderaccountledger.CreateTime.Lte(*in.EndTime))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblelderaccountledger.Id,
			tblelderaccountledger.ElderId,
			tblelder.Name.AsName("elder_name"),
			tblelderaccountledger.Direction,
			tblelderaccountledger.Amount,
			tblelderaccountledger.BalanceAfter,
			tblelderaccountledger.SourceType,
			tblelderaccountledger.BusinessNo,
			tblelderaccountledger.Remark,
			tblelderaccountledger.CreateTime,
		).
		Desc(tblelderaccountledger.CreateTime).
		Select().
		Gets(ctx, out)
}

// GetBalanceSummary 老人资金账户汇总（当前余额 + 累计收支）
func (l *accountLedger) GetBalanceSummary(ctx context.Context, in *dto.IDReq, out *dto.PageAccountLedgerBalanceResp) error {
	elder, has, e := dao.Elder(db).GetByID(ctx, types.BigInt(*in.ID), tblelder.Id, tblelder.Name, tblelder.Balance)
	if e != nil {
		return e
	}
	if !has {
		return errors.New("老人不存在")
	}
	income, e := dao.ElderAccountLedger(db).Sum(ctx,
		[]dialect.Field{tblelderaccountledger.Amount},
		tblelderaccountledger.ElderId.Eq(elder.Id),
		tblelderaccountledger.Direction.Eq(types.Int8(constant.LedgerIncome)),
	)
	if e != nil {
		return e
	}
	outcome, e := dao.ElderAccountLedger(db).Sum(ctx,
		[]dialect.Field{tblelderaccountledger.Amount},
		tblelderaccountledger.ElderId.Eq(elder.Id),
		tblelderaccountledger.Direction.Eq(types.Int8(constant.LedgerOutcome)),
	)
	if e != nil {
		return e
	}
	out.ElderID = types.BigInt(elder.Id)
	out.ElderName = elder.Name.String()
	out.Balance = elder.Balance
	// Sum 返回 map[string]any，key 为列字段
	if v, ok := income[tblelderaccountledger.Amount.Name]; ok {
		out.TotalIncome = types.Money(parseMoneyAny(v))
	}
	if v, ok := outcome[tblelderaccountledger.Amount.Name]; ok {
		out.TotalOutcome = types.Money(parseMoneyAny(v))
	}
	return nil
}

// ChangeBalance 在单个事务内完成"老人余额增减 + 资金明细记账"。
// - 幂等：uk_source(source_type, source_id) 唯一约束，重复记账会失败。
// - 原子：balance 更新与 ledger 插入同事务，要么都成功要么都回滚。
// - 余额不足：direction=出账 且 余额不足以扣款时返回 ErrBalanceDeficiency。
func (l *accountLedger) ChangeBalance(ctx context.Context, in *dto.ChangeBalanceReq, out *dto.EmptyResp) error {
	if e := validateChangeBalance(in); e != nil {
		return e
	}
	_, e := db.Transaction(ctx, func(tx *ace.Tx) (any, error) {
		return nil, l.changeBalanceTx(ctx, tx, in)
	})
	return e
}

// changeBalanceTx 在给定事务内完成老人余额增减 + 资金明细记账（供 ChangeBalance 及
// 其它需要与业务变更同事务的调用方复用，保证扣款/加款与订单、充值业务原子）。
func (l *accountLedger) changeBalanceTx(ctx context.Context, tx *ace.Tx, in *dto.ChangeBalanceReq) error {
	if e := validateChangeBalance(in); e != nil {
		return e
	}
	elderID := types.BigInt(*in.ElderID)
	amount := *in.Amount
	direction := types.Int8(*in.Direction)
	sourceType := ptrStr(in.SourceType)
	if sourceType == "" {
		sourceType = constant.LedgerSourceMANUAL
	}
	sourceID := types.BigInt(ptrI64v(in.SourceID))
	operatorID := types.BigInt(lib.UserID(ctx))
	tenantID := types.BigInt(lib.TenantID(ctx))

	elder, has, e := dao.Elder(tx).Get(ctx, tx.Table(tblelder.TableName).
		Where(tblelder.Id.Eq(elderID), tblelder.TenantId.Eq(tenantID)).
		Cols(tblelder.Id, tblelder.Balance, tblelder.TenantId))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("老人不存在")
	}
	cur := elder.Balance.Int64()
	var newBalance int64
	if direction == types.Int8(constant.LedgerIncome) {
		newBalance = cur + amount.Int64()
	} else {
		newBalance = cur - amount.Int64()
		if newBalance < 0 {
			return constant.ErrBalanceDeficiency
		}
	}
	// 更新老人余额
	if _, e := dao.Elder(tx).UpdateById(ctx, elderID,
		tblelder.Balance.Set(types.Money(newBalance)),
	); e != nil {
		return e
	}
	// 记录资金明细（uk_source 幂等）
	ledger := do.NewElderAccountLedger()
	ledger.TenantId = tenantID
	ledger.ElderId = elderID
	ledger.Direction = direction
	ledger.Amount = amount
	ledger.BalanceAfter = types.Money(newBalance)
	ledger.SourceType = types.String(sourceType)
	ledger.SourceId = sourceID
	ledger.BusinessNo = types.String(ptrStr(in.BusinessNo))
	ledger.Remark = types.String(ptrStr(in.Remark))
	ledger.OperatorId = operatorID
	if _, e := dao.ElderAccountLedger(tx).InsertOne(ctx, ledger,
		tblelderaccountledger.TenantId,
		tblelderaccountledger.ElderId,
		tblelderaccountledger.Direction,
		tblelderaccountledger.Amount,
		tblelderaccountledger.BalanceAfter,
		tblelderaccountledger.SourceType,
		tblelderaccountledger.SourceId,
		tblelderaccountledger.BusinessNo,
		tblelderaccountledger.Remark,
		tblelderaccountledger.OperatorId,
		tblelderaccountledger.CreateTime,
	); e != nil {
		if isDuplicateEntry(e) {
			return constant.ErrLedgerSourceRepeat
		}
		return e
	}
	return nil
}

// validateChangeBalance 校验 ChangeBalanceReq 公共参数
func validateChangeBalance(in *dto.ChangeBalanceReq) error {
	if in.ElderID == nil || in.Amount == nil || in.Direction == nil {
		return constant.ErrParamError
	}
	if *in.Amount == 0 {
		return errors.New("变更金额不能为0")
	}
	if *in.Direction != int8(constant.LedgerIncome) && *in.Direction != int8(constant.LedgerOutcome) {
		return errors.New("资金方向非法")
	}
	return nil
}

// ============ 内部辅助 ============

// parseMoneyAny 将 Sum 接口值安全转为 int64
func parseMoneyAny(v any) int64 {
	switch n := v.(type) {
	case int64:
		return n
	case float64:
		return int64(n)
	case string:
		var r int64
		_, _ = fmt.Sscanf(n, "%d", &r)
		return r
	default:
		return 0
	}
}

// isDuplicateEntry 判断是否为 uk_source 唯一键冲突
func isDuplicateEntry(e error) bool {
	if e == nil {
		return false
	}
	msg := e.Error()
	for _, k := range []string{"Duplicate entry", "10523", "1062"} {
		if readErrContains(msg, k) {
			return true
		}
	}
	return false
}

// readErrContains 简易字符包含判断（避免引入额外依赖）
func readErrContains(s, sub string) bool {
	return len(s) >= len(sub) && (func() bool {
		for i := 0; i+len(sub) <= len(s); i++ {
			if s[i:i+len(sub)] == sub {
				return true
			}
		}
		return false
	})()
}
