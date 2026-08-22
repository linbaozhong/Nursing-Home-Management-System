package service

import (
	"context"
	"errors"

	"api/internal/constant"
	"api/internal/lib"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/define/table/tblvisit"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type visit struct{}

var Visit = &visit{}

// PageVisitByKey 分页查询来访登记（关联老人姓名）
// 对应 Java: VisitServiceImpl.pageVisitByKey -> VisitFunc.listVisitByKey
func (v *visit) PageVisitByKey(ctx context.Context, in *dto.PageVisitByKeyReq, out *[]dto.PageVisitByKeyResp) error {
	clampPage(in.PageNum, in.PageSize)
	q := db.Table(tblvisit.TableName).
		LeftJoin(tblvisit.ElderId, tblelder.Id).
		Where(tblvisit.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblvisit.DelFlag.Eq(constant.YesNoNo))
	if in.ElderName != nil && *in.ElderName != "" {
		q.And(tblelder.Name.Like(*in.ElderName))
	}
	if in.VisitName != nil && *in.VisitName != "" {
		q.And(tblvisit.Name.Like(*in.VisitName))
	}
	if in.VisitPhone != nil && *in.VisitPhone != "" {
		q.And(tblvisit.Phone.Like(*in.VisitPhone))
	}
	if in.VisitFlag != nil && *in.VisitFlag != "" {
		q.And(tblvisit.VisitFlag.Eq(types.Int8(parseVisitFlag(*in.VisitFlag))))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblvisit.Id.AsName("id"),
			tblelder.Name.AsName("elder_name"),
			tblvisit.Name.AsName("visit_name"),
			tblvisit.Phone.AsName("visit_phone"),
			tblvisit.Relation.AsName("relation"),
			tblvisit.VisitDate.AsName("visit_date"),
			tblvisit.LeaveDate.AsName("leave_date"),
			tblvisit.VisitNum.AsName("visit_num"),
			tblvisit.VisitFlag.AsName("visit_flag"),
		).
		Desc(tblvisit.CreateTime).
		Select().
		Gets(ctx, out)
}

// parseVisitFlag 将前端传入的来访状态字符串转为常量值
func parseVisitFlag(s string) int8 {
	if s == "ALREADY_LEAVE" {
		return int8(constant.VisitAlreadyLeave)
	}
	return int8(constant.VisitStayLeave)
}

// AddVisit 新增来访登记（校验老人存在、设置来访状态为在院）
// 对应 Java: VisitServiceImpl.addVisit -> VisitFunc.checkElderByVisit
func (v *visit) AddVisit(ctx context.Context, in *dto.AddVisitReq, out *dto.EmptyResp) error {
	_, has, e := dao.Elder(db).Get(ctx, ace.Where(tblelder.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblelder.Id.Eq(types.BigInt(*in.ElderID))))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("老人不存在")
	}
	bean := do.NewVisit()
	bean.TenantId = types.BigInt(lib.TenantID(ctx))
	bean.ElderId = types.BigInt(*in.ElderID)
	bean.Name = types.String(*in.Name)
	bean.Phone = types.String(*in.Phone)
	bean.Relation = types.String(*in.Relation)
	bean.VisitDate = types.Time{Time: *in.VisitDate}
	bean.VisitNum = types.Int32(int32(*in.VisitNum))
	bean.VisitFlag = types.Int8(constant.VisitStayLeave)
	bean.DelFlag = types.Int8(constant.YesNoNo)
	_, e = dao.Visit(db).InsertOne(ctx, bean)
	return e
}

// GetVisitById 根据编号获取来访登记（编辑回显，关联老人姓名）
// 对应 Java: VisitServiceImpl.getVisitById
func (v *visit) GetVisitById(ctx context.Context, in *dto.IDReq, out *dto.GetVisitByIDResp) error {
	obj, has, e := dao.Visit(db).Get(ctx, ace.Where(tblvisit.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblvisit.Id.Eq(types.BigInt(*in.ID))))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("来访登记不存在")
	}
	out.ID = int64(obj.Id)
	out.VisitName = obj.Name.String()
	out.VisitPhone = obj.Phone.String()
	out.Relation = obj.Relation.String()
	out.VisitDate = obj.VisitDate.Time
	out.VisitNum = int64(obj.VisitNum)
	elder, hasE, e := dao.Elder(db).Get(ctx, ace.Where(tblelder.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblelder.Id.Eq(obj.ElderId)))
	if e != nil {
		return e
	}
	if hasE {
		out.ElderName = elder.Name.String()
	}
	return nil
}

// EditVisit 修改来访登记（不含老人编号、来访状态）
// 对应 Java: VisitServiceImpl.updateVisit
func (v *visit) EditVisit(ctx context.Context, in *dto.EditVisitReq, out *dto.EmptyResp) error {
	_, has, e := dao.Visit(db).Get(ctx, ace.Where(tblvisit.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblvisit.Id.Eq(types.BigInt(*in.ID))))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("来访登记不存在")
	}
	bean := do.NewVisit()
	bean.Id = types.BigInt(*in.ID)
	bean.Name = types.String(*in.Name)
	bean.Phone = types.String(*in.Phone)
	bean.Relation = types.String(*in.Relation)
	bean.VisitDate = types.Time{Time: *in.VisitDate}
	bean.VisitNum = types.Int32(int32(*in.VisitNum))
	_, e = dao.Visit(db).UpdateOne(ctx, bean,
		tblvisit.Name,
		tblvisit.Phone,
		tblvisit.Relation,
		tblvisit.VisitDate,
		tblvisit.VisitNum,
	)
	return e
}

// RecordLeave 登记离开（更新离开时间、置来访状态为已离开）
// 对应 Java: VisitServiceImpl.recordLeave
func (v *visit) RecordLeave(ctx context.Context, in *dto.RecordLeaveReq, out *dto.EmptyResp) error {
	_, has, e := dao.Visit(db).Get(ctx, ace.Where(tblvisit.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblvisit.Id.Eq(types.BigInt(*in.ID))))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("来访登记不存在")
	}
	bean := do.NewVisit()
	bean.Id = types.BigInt(*in.ID)
	bean.LeaveDate = types.Time{Time: *in.LeaveDate}
	bean.VisitFlag = types.Int8(constant.VisitAlreadyLeave)
	_, e = dao.Visit(db).UpdateOne(ctx, bean,
		tblvisit.LeaveDate,
		tblvisit.VisitFlag,
	)
	return e
}

// DeleteVisit 删除来访登记（逻辑删除）
// 对应 Java: VisitServiceImpl.deleteVisit
func (v *visit) DeleteVisit(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		tblvisit.DelFlag.Set(types.Int8(constant.YesNoYes)),
	}
	_, e := dao.Visit(db).UpdateById(ctx, types.BigInt(*in.ID), sets...)
	return e
}
