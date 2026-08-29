package service

import (
	"context"
	"errors"
	"time"

	"api/internal/constant"
	"api/internal/lib"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblconsume"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type consume struct{}

var Consume = &consume{}

// PageConsumeByKey 分页查询消费记录（联表 elder 获取老人姓名/身份证）
// 对应 Java: ConsumeServiceImpl.pageConsumeByKey
func (c *consume) PageConsumeByKey(ctx context.Context, in *dto.PageConsumeByKeyReq, out *[]dto.PageConsumeByKeyResp) error {
	q := db.Table(tblconsume.TableName).
		LeftJoin(tblconsume.ElderId, tblelder.Id).
		Where(tblconsume.TenantId.Eq(types.BigInt(lib.TenantID(ctx))))
	if in.ElderName != nil && *in.ElderName != "" {
		q.And(tblelder.Name.Like(*in.ElderName))
	}
	if in.StartTime != nil && !in.StartTime.IsZero() {
		q.And(tblconsume.ConsumeDate.Gte(*in.StartTime))
	}
	if in.EndTime != nil && !in.EndTime.IsZero() {
		q.And(tblconsume.ConsumeDate.Lte(in.EndTime.Add(24*time.Hour - time.Second)))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblconsume.Id,
			tblelder.Name.AsName("elder_name"),
			tblelder.IdNum.AsName("id_num"),
			tblconsume.ConsumeType,
			tblconsume.ConsumeAmount,
			tblconsume.ConsumeDate,
		).
		Desc(tblconsume.ConsumeDate).
		Select().
		Gets(ctx, out)
}

// GetConsumeById 根据编号获取消费记录
func (c *consume) GetConsumeById(ctx context.Context, in *dto.IDReq, out *dto.GetConsumeByIdResp) error {
	obj, has, e := dao.Consume(db).GetByID(ctx, types.BigInt(*in.ID),
		tblconsume.Id,
		tblconsume.ElderId,
		tblconsume.ConsumeType,
		tblconsume.ConsumeAmount,
		tblconsume.ConsumeDate,
		tblconsume.SourceType,
		tblconsume.SourceId,
		tblconsume.OutTradeNo,
	)
	if e != nil {
		return e
	}
	if !has {
		return errors.New("消费记录不存在")
	}
	out.ID = types.BigInt(obj.Id)
	out.ElderID = types.BigInt(obj.ElderId)
	out.ConsumeType = obj.ConsumeType.String()
	out.ConsumeAmount = obj.ConsumeAmount
	out.ConsumeDate = obj.ConsumeDate.Time
	out.SourceType = obj.SourceType.String()
	out.SourceID = types.BigInt(obj.SourceId)
	out.OutTradeNo = obj.OutTradeNo.String()
	return nil
}

// AddConsume 新增消费（扣减老人余额并记账，同事务）
func (c *consume) AddConsume(ctx context.Context, in *dto.AddConsumeReq, out *dto.EmptyResp) error {
	if in.ElderID == nil || in.ConsumeAmount == nil {
		return constant.ErrParamError
	}
	bean := do.NewConsume()
	bean.TenantId = types.BigInt(lib.TenantID(ctx))
	bean.ElderId = types.BigInt(*in.ElderID)
	bean.ConsumeType = types.String(*in.ConsumeType)
	bean.ConsumeAmount = *in.ConsumeAmount
	bean.ConsumeDate = types.Time{*in.ConsumeDate}
	if in.SourceType != nil {
		bean.SourceType = types.String(*in.SourceType)
	}
	if in.SourceID != nil {
		bean.SourceId = types.BigInt(*in.SourceID)
	}
	if in.OutTradeNo != nil {
		bean.OutTradeNo = types.String(*in.OutTradeNo)
	}

	// 记录消费 + 扣减老人余额记账：同一事务，原子、幂等
	_, e := db.Transaction(ctx, func(tx *ace.Tx) (any, error) {
		// 先落消费记录，拿回自增 id 作为记账 source_id
		if _, e := dao.Consume(tx).InsertOne(ctx, bean, tblconsume.ConsumeType, tblconsume.ElderId, tblconsume.ConsumeAmount, tblconsume.ConsumeDate, tblconsume.SourceType, tblconsume.SourceId, tblconsume.OutTradeNo); e != nil {
			return nil, e
		}
		elderID := bean.ElderId.Int64()
		sourceID := bean.Id.Int64()
		amt := bean.ConsumeAmount
		direction := int8(constant.LedgerOutcome)
		sourceType := constant.LedgerSourceMANUAL
		deduct := &dto.ChangeBalanceReq{
			ElderID:    &elderID,
			Direction:  &direction,
			Amount:     &amt,
			SourceType: &sourceType,
			SourceID:   &sourceID,
		}
		if e := AccountLedger.changeBalanceTx(ctx, tx, deduct); e != nil {
			return nil, e
		}
		return nil, nil
	})
	return e
}

// EditConsume 编辑消费
func (c *consume) EditConsume(ctx context.Context, in *dto.EditConsumeReq, out *dto.EmptyResp) error {
	var sets = make([]dialect.Setter, 0, 4)
	sets = append(sets, tblconsume.ElderId.Set(*in.ElderID))
	sets = append(sets, tblconsume.ConsumeType.Set(*in.ConsumeType))
	sets = append(sets, tblconsume.ConsumeAmount.Set(*in.ConsumeAmount))
	sets = append(sets, tblconsume.ConsumeDate.Set(*in.ConsumeDate))
	if in.SourceType != nil {
		sets = append(sets, tblconsume.SourceType.Set(*in.SourceType))
	}
	if in.SourceID != nil {
		sets = append(sets, tblconsume.SourceId.Set(*in.SourceID))
	}
	if in.OutTradeNo != nil {
		sets = append(sets, tblconsume.OutTradeNo.Set(*in.OutTradeNo))
	}
	_, e := dao.Consume(db).UpdateById(ctx, types.BigInt(*in.ID), sets...)
	return e
}

// DeleteConsume 删除消费
func (c *consume) DeleteConsume(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.Consume(db).DeleteById(ctx, types.BigInt(*in.ID))
	return e
}

// parseTimeStr 解析消费日期字符串（支持两种格式）
func parseTimeStr(s *string) types.Time {
	if s == nil || *s == "" {
		return types.Time{}
	}
	for _, layout := range []string{"2006-01-02 15:04:05", "2006-01-02"} {
		if t, err := time.ParseInLocation(layout, *s, time.Local); err == nil {
			return types.Time{t}
		}
	}
	return types.Time{}
}
