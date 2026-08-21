package service

import (
	"context"

	"api/internal/constant"
	"api/internal/lib"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/define/table/tblretreatapply"
	"api/internal/model/define/table/tblstaff"
	"api/internal/model/do"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/types"
)

var _ = (*retreatAuditService)(nil)

type retreatAuditService struct{}

// retreatAuditJoin 接收入住审核联表（老人姓名、申请人姓名）查询结果的中间结构体
type retreatAuditJoin struct {
	ID        types.BigInt `json:"id"`
	ElderName types.String `json:"elder_name"`
	ApplyFlag types.Int8   `json:"apply_flag"`
	ApplyName types.String `json:"apply_name"`
}

// PageRetreatAuditByKey 分页查询退住审核
func (s *retreatAuditService) PageRetreatAuditByKey(ctx context.Context, in *dto.PageRetreatAuditByKeyQuery, out *[]dto.PageRetreatAuditByKeyVO) error {
	if in.PageNum == nil || in.PageSize == nil {
		return constant.ErrParamInvalid
	}
	q := db.Table(tblretreatapply.TableName).
		LeftJoin(tblretreatapply.ElderId, tblelder.Id).
		LeftJoin(tblretreatapply.CreateId, tblstaff.Id).
		Where(tblretreatapply.TenantId.Eq(types.BigInt(lib.TenantID(ctx))))
	if in.Key != nil && *in.Key != "" {
		q = q.Where(tblelder.Name.Like(*in.Key))
	}
	var joins []retreatAuditJoin
	has, e := q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblretreatapply.Id,
			tblretreatapply.ApplyFlag,
			tblelder.Name.As("elder_name"),
			tblstaff.Name.As("apply_name"),
		).
		Desc(tblretreatapply.Id).
		Select().Gets(ctx, &joins)
	if e != nil {
		return e
	}
	if !has {
		return nil
	}
	res := make([]dto.PageRetreatAuditByKeyVO, 0, len(joins))
	for _, j := range joins {
		res = append(res, dto.PageRetreatAuditByKeyVO{
			ID:        int64(j.ID),
			ElderName: j.ElderName.String(),
			ApplyFlag: constant.AuditStatus(j.ApplyFlag).String(),
			ApplyName: j.ApplyName.String(),
		})
	}
	*out = res
	return nil
}

// GetRetreatAuditById 查询退住审核详情（含老人信息）
func (s *retreatAuditService) GetRetreatAuditById(ctx context.Context, in *dto.IDReq, out *dto.GetRetreatAuditByIDVO) error {
	apply := new(do.RetreatApply)
	has, e := dao.RetreatApply(db).Get(ctx, ace.Where(tblretreatapply.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblretreatapply.Id.Eq(types.BigInt(*in.ID))))
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrDataNotExist
	}
	out.ID = int64(apply.Id)
	out.ElderID = int64(apply.ElderId)
	out.ApplyFlag = constant.AuditStatus(apply.ApplyFlag).String()
	// 关联老人姓名
	elder := new(do.Elder)
	if eh, ee := dao.Elder(db).Get(ctx, ace.Where(tblelder.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblelder.Id.Eq(apply.ElderId))); ee == nil && eh {
		out.ElderName = elder.Name.String()
	}
	return nil
}

// AuditRetreat 审核退住申请（通过/不通过）
func (s *retreatAuditService) AuditRetreat(ctx context.Context, in *dto.AuditRetreatQuery) error {
	if in.ID == nil || in.AuditResult == nil {
		return constant.ErrParamInvalid
	}
	apply := new(do.RetreatApply)
	has, e := dao.RetreatApply(db).Get(ctx, ace.Where(tblretreatapply.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblretreatapply.Id.Eq(types.BigInt(*in.ID))))
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrDataNotExist
	}
	if *in.AuditResult == "不通过" {
		if _, e = dao.RetreatApply(db).UpdateById(ctx, *in.ID, tblretreatapply.ApplyFlag.Set(types.Int8(constant.AuditNotPass))); e != nil {
			return e
		}
		return nil
	}
	// 审核通过
	if _, e = dao.RetreatApply(db).UpdateById(ctx, *in.ID, tblretreatapply.ApplyFlag.Set(types.Int8(constant.AuditPass))); e != nil {
		return e
	}
	// 同步老人状态为退住审核
	if _, e = dao.Elder(db).UpdateById(ctx, int64(apply.ElderId), tblelder.CheckFlag.Set(types.Int8(constant.CheckExitAudit))); e != nil {
		return e
	}
	return nil
}
