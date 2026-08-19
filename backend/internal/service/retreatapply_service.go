package service

import (
	"context"
	"errors"

	"api/internal/constant"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblbed"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/define/table/tblretreatapply"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/conv"
	"github.com/linbaozhong/gentity/pkg/types"
)

type retreatApply struct{}

var RetreatApply = &retreatApply{}

// PageRetreatApplyByKey 分页查询退住申请（关联老人、床位）
// 对应 Java: RetreatApplyServiceImpl.pageRetreatApplyByKey -> RetreatApplyFunc.listRetreatApplyByKey
func (r *retreatApply) PageRetreatApplyByKey(ctx context.Context, in *dto.PageRetreatApplyByKeyQuery, out *[]dto.PageRetreatByKeyVO) error {
	q := db.Table(tblretreatapply.TableName).
		LeftJoin(tblretreatapply.ElderId, tblelder.Id).
		LeftJoin(tblelder.BedId, tblbed.Id)
	if in.ElderName != nil && *in.ElderName != "" {
		q.And(tblelder.Name.Like(*in.ElderName))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblretreatapply.Id.AsName("apply_id"),
			tblretreatapply.ElderId.AsName("elder_id"),
			tblelder.Name.AsName("elder_name"),
			tblelder.Sex.AsName("elder_sex"),
			tblelder.IdNum.AsName("id_num"),
			tblbed.Name.AsName("bed_name"),
			tblretreatapply.ApplyFlag.AsName("apply_flag"),
		).
		Desc(tblretreatapply.CreateTime).
		Select().
		Gets(ctx, out)
}

// PageSearchElderByKey 分页搜索老人（入住中、且不存在未审核退住申请的老人）
// 对应 Java: RetreatApplyServiceImpl.pageSearchElderByKey -> RetreatApplyFunc.listElderByKey
func (r *retreatApply) PageSearchElderByKey(ctx context.Context, in *dto.PageSearchElderByKeyQuery, out *[]dto.PageSearchElderByKeyVO) error {
	// 排除已存在退住申请（待审核/审核中/通过）的老人
	applyingIds := make([]any, 0)
	applies, _, e := dao.RetreatApply(db).List(ctx,
		db.Table(tblretreatapply.TableName).
			Cols(tblretreatapply.ElderId).
			Where(tblretreatapply.ApplyFlag.In(
				types.Int8(constant.AuditStay),
				types.Int8(constant.Auditing),
				types.Int8(constant.AuditPass),
			)),
	)
	if e != nil {
		return e
	}
	for _, a := range applies {
		applyingIds = append(applyingIds, a.ElderId)
	}

	q := db.Table(tblelder.TableName).
		LeftJoin(tblelder.BedId, tblbed.Id).
		Where(tblelder.CheckFlag.Eq(types.Int8(constant.CheckEnter)))
	if in.Name != nil && *in.Name != "" {
		q.And(tblelder.Name.Like(*in.Name))
	}
	if len(applyingIds) > 0 {
		q.And(tblelder.Id.NotIn(applyingIds...))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblelder.Id.AsName("elder_id"),
			tblelder.Name.AsName("elder_name"),
			tblelder.Sex.AsName("elder_sex"),
			tblelder.IdNum.AsName("id_num"),
			tblbed.Name.AsName("bed_name"),
		).
		Desc(tblelder.CreateTime).
		Select().
		Gets(ctx, out)
}

// AddRetreatApply 新增退住申请（校验老人入住中、置退住审核状态、床位退住审核）
// 对应 Java: RetreatApplyServiceImpl.addRetreatApply -> RetreatApplyFunc.checkElderByRetreatApply
func (r *retreatApply) AddRetreatApply(ctx context.Context, in *dto.AddRetreatApplyQuery, out *dto.EmptyResp) error {
	elder, has, e := dao.Elder(db).Get(ctx, ace.Where(tblelder.Id.Eq(types.BigInt(*in.ElderID))))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("老人不存在")
	}
	if elder.CheckFlag != types.Int8(constant.CheckEnter) {
		return errors.New("老人不是入住状态，无法申请退住")
	}
	// 已存在未审核的退住申请则拦截
	_, exist, e := dao.RetreatApply(db).Get(ctx,
		ace.Where(
			tblretreatapply.ElderId.Eq(types.BigInt(*in.ElderID))).
			And(tblretreatapply.ApplyFlag.In(
				types.Int8(constant.AuditStay),
				types.Int8(constant.Auditing),
				types.Int8(constant.AuditPass),
			)),
	)
	if e != nil {
		return e
	}
	if exist {
		return errors.New("该老人已存在退住申请")
	}

	if _, e = dao.RetreatApply(db).InsertOne(ctx, &do.RetreatApply{
		ElderId:   types.BigInt(*in.ElderID),
		ApplyFlag: types.Int8(constant.AuditStay),
	}); e != nil {
		return e
	}
	// 老人置退住审核状态
	if _, e = dao.Elder(db).UpdateById(ctx, types.BigInt(*in.ElderID),
		tblelder.CheckFlag.Set(types.Int8(constant.CheckExitAudit)),
	); e != nil {
		return e
	}
	// 床位置退住审核
	if int64(elder.BedId) > 0 {
		if _, e = dao.Bed(db).UpdateById(ctx, elder.BedId,
			tblbed.BedFlag.Set(types.Int8(constant.BedExitAudit)),
		); e != nil {
			return e
		}
	}
	return nil
}

// GetRetreatApplyById 根据编号获取退住申请详情
func (r *retreatApply) GetRetreatApplyById(ctx context.Context, in *dto.IDReq, out *dto.PageRetreatByKeyVO) error {
	obj, has, e := dao.RetreatApply(db).Get(ctx, ace.Where(tblretreatapply.Id.Eq(types.BigInt(*in.ID))))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("退住申请不存在")
	}
	out.ApplyID = int64(obj.Id)
	out.ElderID = int64(obj.ElderId)
	out.ApplyFlag = conv.Ptr(int8(obj.ApplyFlag))
	elder, hasE, e := dao.Elder(db).Get(ctx, ace.Where(tblelder.Id.Eq(obj.ElderId)))
	if e != nil {
		return e
	}
	if hasE {
		out.ElderName = elder.Name.String()
		out.ElderSex = elder.Sex.String()
		out.IDNum = elder.IdNum.String()
		if bed, hasB, e := dao.Bed(db).Get(ctx, ace.Where(tblbed.Id.Eq(elder.BedId))); e == nil && hasB {
			out.BedName = bed.Name.String()
		}
	}
	return nil
}

// EditRetreatApply 修改退住申请（重新提交待审核）
func (r *retreatApply) EditRetreatApply(ctx context.Context, in *dto.EditRetreatApplyQuery, out *dto.EmptyResp) error {
	_, has, e := dao.RetreatApply(db).Get(ctx, ace.Where(tblretreatapply.Id.Eq(types.BigInt(*in.ID))))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("退住申请不存在")
	}
	_, e = dao.RetreatApply(db).UpdateById(ctx, types.BigInt(*in.ID),
		tblretreatapply.ApplyFlag.Set(types.Int8(constant.AuditStay)),
	)
	return e
}

// DeleteRetreatApply 删除退住申请（物理删除，并还原老人/床位状态）
func (r *retreatApply) DeleteRetreatApply(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.RetreatApply(db).Get(ctx, ace.Where(tblretreatapply.Id.Eq(types.BigInt(*in.ID))))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("退住申请不存在")
	}
	elder, hasE, e := dao.Elder(db).Get(ctx, ace.Where(tblelder.Id.Eq(obj.ElderId)))
	if e != nil {
		return e
	}
	if _, e = dao.RetreatApply(db).DeleteById(ctx, types.BigInt(*in.ID)); e != nil {
		return e
	}
	// 还原老人为入住状态
	if _, e = dao.Elder(db).UpdateById(ctx, obj.ElderId,
		tblelder.CheckFlag.Set(types.Int8(constant.CheckEnter)),
	); e != nil {
		return e
	}
	// 还原床位为入住状态
	if hasE && int64(elder.BedId) > 0 {
		if _, e = dao.Bed(db).UpdateById(ctx, elder.BedId,
			tblbed.BedFlag.Set(types.Int8(constant.BedEnter)),
		); e != nil {
			return e
		}
	}
	return nil
}
