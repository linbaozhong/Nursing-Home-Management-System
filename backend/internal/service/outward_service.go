package service

import (
	"context"
	"errors"

	"api/internal/constant"
	"api/internal/lib"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/define/table/tbloutward"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/types"
)

type outward struct{}

var Outward = &outward{}

// PageOutwardByKey 分页查询外出登记（关联老人姓名）
// 对应 Java: OutwardServiceImpl.pageOutwardByKey -> OutwardMapper.listOutwardByKey
func (o *outward) PageOutwardByKey(ctx context.Context, in *dto.PageOutwardByKeyReq, out *[]dto.PageOutwardByKeyResp) error {
	clampPage(in.PageNum, in.PageSize)
	q := db.Table(tbloutward.TableName).
		RightJoin(tbloutward.ElderId, tblelder.Id).
		Where(tbloutward.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tbloutward.DelFlag.Eq(constant.YesNoNo))
	if in.ElderName != nil && *in.ElderName != "" {
		q.And(tblelder.Name.Like(*in.ElderName))
	}
	if in.StartTime != nil {
		q.And(tbloutward.OutwardDate.Gte(types.Time{Time: *in.StartTime}))
	}
	if in.EndTime != nil {
		q.And(tbloutward.OutwardDate.Lte(types.Time{Time: *in.EndTime}))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tbloutward.Id.AsName("id"),
			tblelder.Name.AsName("elder_name"),
			tbloutward.ChaperoneName.AsName("chaperone_name"),
			tbloutward.ChaperonePhone.AsName("chaperone_phone"),
			tbloutward.ChaperoneType.AsName("chaperone_type"),
			tbloutward.OutwardDate.AsName("outward_date"),
			tbloutward.PlanReturnDate.AsName("plan_return_date"),
			tbloutward.RealReturnDate.AsName("real_return_date"),
		).
		Desc(tbloutward.CreateTime).
		Select().
		Gets(ctx, out)
}

// GetOutwardById 根据编号获取外出登记（含老人姓名）
// 对应 Java: OutwardServiceImpl.getOutwardById -> outwardMapper.getOutwardById
func (o *outward) GetOutwardById(ctx context.Context, in *dto.IDReq, out *dto.GetOutwardByIDResp) error {
	obj, has, e := dao.Outward(db).Get(ctx, ace.Where(tbloutward.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tbloutward.Id.Eq(types.BigInt(*in.ID))))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("外出登记不存在")
	}
	out.ID = int64(obj.Id)
	out.ChaperoneName = obj.ChaperoneName.String()
	out.ChaperonePhone = obj.ChaperonePhone.String()
	out.ChaperoneType = obj.ChaperoneType.String()
	out.OutwardDate = obj.OutwardDate.Time
	out.PlanReturnDate = obj.PlanReturnDate.Time
	elder, hasE, e := dao.Elder(db).Get(ctx, ace.Where(tblelder.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblelder.Id.Eq(obj.ElderId)))
	if e != nil {
		return e
	}
	if hasE {
		out.ElderName = elder.Name.String()
	}
	return nil
}

// AddOutward 新增外出登记（校验老人无未删除的外出登记）
// 对应 Java: OutwardServiceImpl.addOutward
func (o *outward) AddOutward(ctx context.Context, in *dto.AddOutwardReq, out *dto.EmptyResp) error {
	_, exist, e := dao.Outward(db).Get(ctx,
		ace.Where(
			tbloutward.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
			tbloutward.ElderId.Eq(types.BigInt(*in.ElderID))).
			And(tbloutward.DelFlag.Eq(constant.YesNoNo)),
	)
	if e != nil {
		return e
	}
	if exist {
		return errors.New("该老人已存在外出登记")
	}
	bean := do.NewOutward()
	bean.TenantId = types.BigInt(lib.TenantID(ctx))
	bean.ElderId = types.BigInt(*in.ElderID)
	bean.ChaperoneName = types.String(*in.ChaperoneName)
	bean.ChaperonePhone = types.String(*in.ChaperonePhone)
	bean.ChaperoneType = types.String(*in.ChaperoneType)
	bean.OutwardDate = types.Time{Time: *in.OutwardDate}
	bean.PlanReturnDate = types.Time{Time: *in.PlanReturnDate}
	bean.DelFlag = types.Int8(constant.YesNoNo)
	_, e = dao.Outward(db).InsertOne(ctx, bean)
	return e
}

// EditOutward 编辑外出登记
// 对应 Java: OutwardServiceImpl.editOutward
func (o *outward) EditOutward(ctx context.Context, in *dto.EditOutwardReq, out *dto.EmptyResp) error {
	_, has, e := dao.Outward(db).Get(ctx, ace.Where(tbloutward.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tbloutward.Id.Eq(types.BigInt(*in.ID))))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("外出登记不存在")
	}
	bean := do.NewOutward()
	bean.Id = types.BigInt(*in.ID)
	bean.ElderId = types.BigInt(*in.ElderID)
	bean.ChaperoneName = types.String(*in.ChaperoneName)
	bean.ChaperonePhone = types.String(*in.ChaperonePhone)
	bean.ChaperoneType = types.String(*in.ChaperoneType)
	bean.OutwardDate = types.Time{Time: *in.OutwardDate}
	bean.PlanReturnDate = types.Time{Time: *in.PlanReturnDate}
	_, e = dao.Outward(db).UpdateOne(ctx, bean,
		tbloutward.ElderId,
		tbloutward.ChaperoneName,
		tbloutward.ChaperonePhone,
		tbloutward.ChaperoneType,
		tbloutward.OutwardDate,
		tbloutward.PlanReturnDate,
	)
	return e
}

// DeleteOutward 删除外出登记（逻辑删除）
// 对应 Java: OutwardServiceImpl.deleteOutward
func (o *outward) DeleteOutward(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.Outward(db).UpdateById(ctx, types.BigInt(*in.ID),
		tbloutward.DelFlag.Set(types.Int8(constant.YesNoYes)),
	)
	return e
}

// PageSearchElderByKey 分页搜索老人（入住中或退住审核、且当前无未删除外出登记）
// 对应 Java: OutwardServiceImpl.pageSearchElderByKey -> commonFunc.listPageElderByKey + 过滤
func (o *outward) PageSearchElderByKey(ctx context.Context, in *dto.PageSearchElderByKeyReq, out *[]dto.PageSearchElderByKeyResp) error {
	// 排除已有未删除外出登记的老人
	outwardElderIds := make([]any, 0)
	records, _, e := dao.Outward(db).List(ctx,
		db.Table(tbloutward.TableName).
			Cols(tbloutward.ElderId).
			Where(tbloutward.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
				tbloutward.DelFlag.Eq(constant.YesNoNo)),
	)
	if e != nil {
		return e
	}
	for _, r := range records {
		outwardElderIds = append(outwardElderIds, r.ElderId)
	}

	q := db.Table(tblelder.TableName).
		Where(tblelder.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
			tblelder.CheckFlag.In(
				types.Int8(constant.CheckEnter),
				types.Int8(constant.CheckExitAudit),
			))
	if in.Name != nil && *in.Name != "" {
		q.And(tblelder.Name.Like(*in.Name))
	}
	if len(outwardElderIds) > 0 {
		q.And(tblelder.Id.NotIn(outwardElderIds...))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblelder.Id.AsName("elder_id"),
			tblelder.Name.AsName("elder_name"),
			tblelder.Sex.AsName("elder_sex"),
			tblelder.IdNum.AsName("id_num"),
		).
		Desc(tblelder.CreateTime).
		Select().
		Gets(ctx, out)
}

// DelayReturn 延期返回（更新计划返回时间）
// 对应 Java: OutwardServiceImpl.delayReturn
func (o *outward) DelayReturn(ctx context.Context, in *dto.DelayReturnReq, out *dto.EmptyResp) error {
	_, has, e := dao.Outward(db).Get(ctx, ace.Where(tbloutward.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tbloutward.Id.Eq(types.BigInt(*in.ID))))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("外出登记不存在")
	}
	_, e = dao.Outward(db).UpdateById(ctx, types.BigInt(*in.ID),
		tbloutward.PlanReturnDate.Set(types.Time{Time: *in.PlanReturnDate}),
	)
	return e
}

// RecordReturn 登记返回（更新实际返回时间）
// 对应 Java: OutwardServiceImpl.recordReturn
func (o *outward) RecordReturn(ctx context.Context, in *dto.RecordReturnReq, out *dto.EmptyResp) error {
	obj, has, e := dao.Outward(db).Get(ctx, ace.Where(tbloutward.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tbloutward.Id.Eq(types.BigInt(*in.ID))))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("外出登记不存在")
	}
	if !obj.RealReturnDate.IsZero() {
		return errors.New("该外出登记已登记返回")
	}
	_, e = dao.Outward(db).UpdateById(ctx, types.BigInt(*in.ID),
		tbloutward.RealReturnDate.Set(types.Time{Time: *in.RealReturnDate}),
	)
	return e
}
