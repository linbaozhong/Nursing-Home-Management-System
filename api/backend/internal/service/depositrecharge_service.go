package service

import (
	"context"
	"errors"
	"strconv"

	"api/internal/constant"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblbed"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/types"
)

// depositRechargeRow 联表查询接收结构（elder_id 用 int64 接收，再转 VO 的 string）
type depositRechargeRow struct {
	ElderID    int64   `json:"elder_id"`
	ElderName  string  `json:"elder_name"`
	ElderPhone string  `json:"elder_phone"`
	IDNum      string  `json:"id_num"`
	BedName    string  `json:"bed_name"`
	Balance    float64 `json:"balance"`
}

type depositrecharge struct{}

var DepositRecharge = &depositrecharge{}

// PageDepositRechargeByKey 分页查询预存充值（入住/退住审核老人，联表床位）
// 对应 Java: DepositRechargeServiceImpl.pageDepositRechargeByKey -> ElderMapper.listDepositRechargeByKey
func (d *depositrecharge) PageDepositRechargeByKey(ctx context.Context, in *dto.PageDepositRechargeByKeyQuery, out *[]dto.PageDepositRechargeByKeyVO) error {
	q := db.Table(do.ElderTableName).
		LeftJoin(tblelder.BedId, tblbed.Id).
		Where(
			tblelder.DelFlag.Eq(constant.YesNoNo),
			tblelder.CheckFlag.In(
				types.Int8(constant.CheckEnter),
				types.Int8(constant.CheckExitAudit),
			),
		)
	if in.Name != nil && *in.Name != "" {
		q.And(tblelder.Name.Like(*in.Name))
	}
	if in.Phone != nil && *in.Phone != "" {
		q.And(tblelder.Phone.Like(*in.Phone))
	}
	if in.IDNum != nil && *in.IDNum != "" {
		q.And(tblelder.IdNum.Like(*in.IDNum))
	}
	var rows []depositRechargeRow
	e := q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblelder.Id.AsName("elder_id"),
			tblelder.Name.AsName("elder_name"),
			tblelder.Phone.AsName("elder_phone"),
			tblelder.IdNum.AsName("id_num"),
			tblbed.Name.AsName("bed_name"),
			tblelder.Balance,
		).
		Desc(tblelder.CreateTime).
		Select().
		Gets(ctx, &rows)
	if e != nil {
		return e
	}
	*out = make([]dto.PageDepositRechargeByKeyVO, 0, len(rows))
	for _, r := range rows {
		*out = append(*out, dto.PageDepositRechargeByKeyVO{
			ElderID:    strconv.FormatInt(r.ElderID, 10),
			ElderName:  r.ElderName,
			ElderPhone: r.ElderPhone,
			IDNum:      r.IDNum,
			BedName:    r.BedName,
			Balance:    r.Balance,
		})
	}
	return nil
}

// PageSearchElderByKey 分页搜索老人（入住/退住审核老人，供选择）
// 对应 Java: DepositRechargeServiceImpl.pageSearchElderByKey -> CommonFunc.pageSearchElderByKeyResult
func (d *depositrecharge) PageSearchElderByKey(ctx context.Context, in *dto.PageSearchElderByKeyQuery, out *[]dto.PageSearchElderByKeyVO) error {
	q := db.Table(do.ElderTableName).
		Where(
			tblelder.DelFlag.Eq(constant.YesNoNo),
			tblelder.CheckFlag.In(
				types.Int8(constant.CheckEnter),
				types.Int8(constant.CheckExitAudit),
			),
		)
	if in.Key != nil && *in.Key != "" {
		q.And(tblelder.Name.Like(*in.Key))
		q.Or(tblelder.Phone.Like(*in.Key))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblelder.Id,
			tblelder.Name,
			tblelder.IdNum,
			tblelder.Sex,
			tblelder.Phone,
			tblelder.Address,
			tblelder.CheckFlag,
		).
		Desc(tblelder.CreateTime).
		Select().
		Gets(ctx, out)
}

// 以下方法为脚手架生成但无对应 Java 后端业务/数据表（deposit_recharge、elder_fee 等表在 Go 端未生成），
// 保留签名以维持 handler 路由可编译，返回未实现错误。
func (d *depositrecharge) GetDepositRechargeById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	return errors.New("not implemented")
}
func (d *depositrecharge) AddDepositRecharge(ctx context.Context, in *dto.AddDepositRechargeQuery, out *dto.EmptyResp) error {
	return errors.New("not implemented")
}
func (d *depositrecharge) EditDepositRecharge(ctx context.Context, in *dto.EditDepositRechargeQuery, out *dto.EmptyResp) error {
	return errors.New("not implemented")
}
func (d *depositrecharge) DeleteDepositRecharge(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	return errors.New("not implemented")
}
func (d *depositrecharge) PageSearchStaffByKey(ctx context.Context, in *dto.PageSearchStaffByKeyQuery, out *dto.EmptyResp) error {
	return errors.New("not implemented")
}
func (d *depositrecharge) GetElderFeeById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	return errors.New("not implemented")
}
func (d *depositrecharge) AuditElderFee(ctx context.Context, in *dto.AuditElderFeeQuery, out *dto.EmptyResp) error {
	return errors.New("not implemented")
}
