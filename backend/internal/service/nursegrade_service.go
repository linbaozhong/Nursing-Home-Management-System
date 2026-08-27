package service

import (
	"api/internal/constant"
	"api/internal/lib"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblnursegrade"
	"api/internal/model/define/table/tblnurseitem"
	"api/internal/model/define/table/tblserviceitem"
	"api/internal/model/define/table/tblstaff"
	"api/internal/model/do"
	"api/internal/model/dto"
	"context"

	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/types"
)

// nurseRoleID 护理员角色编号（对应用户端 role=护理员）
const nurseRoleID = 6

var NurseGrade = (*nurseGradeService)(nil)

type nurseGradeService struct{}

// PageNurseGradeByKey 分页查询护理等级
func (s *nurseGradeService) PageNurseGradeByKey(ctx context.Context, in *dto.PageNurseGradeByKeyReq, out *[]dto.PageNurseGradeByKeyResp) error {
	q := db.Table(tblnursegrade.TableName)
	q = q.Where(tblnursegrade.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblnursegrade.State.NotEq(types.Int8(constant.StateDeleted)))
	if in.GradeName != nil {
		q = q.And(tblnursegrade.Name.Like(*in.GradeName))
	}
	if in.NurseType != nil {
		q = q.And(tblnursegrade.Type.Like(*in.NurseType))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblnursegrade.Id.As("id"),
			tblnursegrade.Name.As("name"),
			tblnursegrade.Type.As("type"),
			tblnursegrade.MonthPrice.As("month_price"),
		).
		Desc(tblnursegrade.CreateTime).
		Select().Gets(ctx, out)
}

// GetNurseGradeById 查询护理等级详情（含护理服务列表）
func (s *nurseGradeService) GetNurseGradeById(ctx context.Context, in *dto.IDReq, out *dto.GetNurseGradeByIDResp) error {
	bean, has, e := dao.NurseGrade(db).GetByID(ctx, types.BigInt(*in.ID), tblnursegrade.Id, tblnursegrade.Name, tblnursegrade.Type, tblnursegrade.MonthPrice, tblnursegrade.State)
	if e != nil {
		return e
	}
	if !has {
		return nil
	}
	*out.ID = int64(bean.Id)
	*out.Name = bean.Name.String()
	*out.Type = bean.Type.String()
	*out.MonthPrice = bean.MonthPrice
	// 关联护理服务
	items, _, e := dao.NurseItem(db).List(ctx, ace.Where(tblnurseitem.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblnurseitem.GradeId.Eq(types.BigInt(*in.ID))).Cols(tblnurseitem.ServiceId))
	if e != nil {
		return e
	}
	if len(items) == 0 {
		return nil
	}
	ids := make([]any, 0, len(items))
	for _, it := range items {
		ids = append(ids, it.ServiceId)
	}
	services, _, e := dao.ServiceItem(db).List(ctx, ace.Where(tblserviceitem.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblserviceitem.Id.In(ids...)).
		And(tblserviceitem.State.NotEq(types.Int8(constant.StateDeleted))).
		Cols(tblserviceitem.Id, tblserviceitem.TypeId, tblserviceitem.Name, tblserviceitem.ChargeMethod, tblserviceitem.Price, tblserviceitem.NeedDate))
	if e != nil {
		return e
	}
	list := make([]dto.NurseGradeServiceResp, 0, len(services))
	for _, sv := range services {
		vo := dto.NurseGradeServiceResp{}
		*vo.ID = int64(sv.Id)
		*vo.TypeID = int64(sv.TypeId)
		*vo.Name = sv.Name.String()
		*vo.ChargeMethod = sv.ChargeMethod.String()
		*vo.Price = sv.Price
		*vo.NeedDate = int(sv.NeedDate)
		list = append(list, vo)
	}
	out.NurseGradeServiceRespList = list
	return nil
}

// AddNurseGrade 新增护理等级
func (s *nurseGradeService) AddNurseGrade(ctx context.Context, in *dto.AddNurseGradeReq, out *dto.EmptyResp) error {
	has, e := dao.NurseGrade(db).Exists(ctx,
		tblnursegrade.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
		tblnursegrade.Name.Eq(types.String(*in.Name)),
		tblnursegrade.Type.Eq(types.String(*in.Type)),
		tblnursegrade.State.NotEq(types.Int8(constant.StateDeleted)),
	)
	if e != nil {
		return e
	}
	if has {
		return constant.ErrNurseGradeRepeat
	}
	bean := new(do.NurseGrade)
	bean.TenantId = types.BigInt(lib.TenantID(ctx))
	bean.Name = types.String(*in.Name)
	bean.Type = types.String(*in.Type)
	bean.MonthPrice = types.Money(*in.MonthPrice)
	bean.State = types.Int8(constant.StateEnabled)
	ok, e := dao.NurseGrade(db).InsertOne(ctx, bean, tblnursegrade.TenantId, tblnursegrade.Name, tblnursegrade.Type, tblnursegrade.MonthPrice, tblnursegrade.State, tblnursegrade.CreateId, tblnursegrade.CreateTime)
	if e != nil {
		return e
	}
	if !ok {
		return constant.ErrNurseGradeRepeat
	}
	// 关联护理服务
	if len(in.ServiceIDList) > 0 {
		beans := make([]*do.NurseItem, 0, len(in.ServiceIDList))
		for _, sid := range in.ServiceIDList {
			ni := new(do.NurseItem)
			ni.TenantId = types.BigInt(lib.TenantID(ctx))
			ni.NurseGradeId = bean.Id
			ni.ServiceItemId = types.BigInt(sid)
			beans = append(beans, ni)
		}
		if _, e = dao.NurseItem(db).InsertBatch(ctx, beans, tblnurseitem.TenantId, tblnurseitem.NurseGradeId, tblnurseitem.ServiceItemId); e != nil {
			return e
		}
	}
	return nil
}

// EditNurseGrade 编辑护理等级
func (s *nurseGradeService) EditNurseGrade(ctx context.Context, in *dto.EditNurseGradeReq, out *dto.EmptyResp) error {
	has, e := dao.NurseGrade(db).Exists(ctx,
		tblnursegrade.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
		tblnursegrade.Name.Eq(types.String(*in.Name)),
		tblnursegrade.Type.Eq(types.String(*in.Type)),
		tblnursegrade.State.NotEq(types.Int8(constant.StateDeleted)),
		tblnursegrade.Id.NotEq(types.BigInt(*in.ID)),
	)
	if e != nil {
		return e
	}
	if has {
		return constant.ErrNurseGradeRepeat
	}
	ok, e := dao.NurseGrade(db).UpdateById(ctx, types.BigInt(*in.ID),
		tblnursegrade.Name.Set(types.String(*in.Name)),
		tblnursegrade.Type.Set(types.String(*in.Type)),
		tblnursegrade.MonthPrice.Set(types.Money(*in.MonthPrice)),
	)
	if e != nil {
		return e
	}
	if !ok {
		return constant.ErrNurseGradeRepeat
	}
	// 重建关联护理服务
	if _, e = dao.NurseItem(db).Delete(ctx,
		tblnurseitem.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
		tblnurseitem.GradeId.Eq(types.BigInt(*in.ID))); e != nil {
		return e
	}
	if len(in.ServiceIDList) > 0 {
		beans := make([]*do.NurseItem, 0, len(in.ServiceIDList))
		for _, sid := range in.ServiceIDList {
			ni := new(do.NurseItem)
			ni.TenantId = types.BigInt(lib.TenantID(ctx))
			ni.NurseGradeId = types.BigInt(*in.ID)
			ni.ServiceItemId = types.BigInt(sid)
			beans = append(beans, ni)
		}
		if _, e = dao.NurseItem(db).InsertBatch(ctx, beans, tblnurseitem.TenantId, tblnurseitem.NurseGradeId, tblnurseitem.ServiceItemId); e != nil {
			return e
		}
	}
	return nil
}

// DeleteNurseGrade 删除护理等级（逻辑删除）
func (s *nurseGradeService) DeleteNurseGrade(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.NurseGrade(db).UpdateById(ctx, types.BigInt(*in.ID),
		tblnursegrade.State.Set(types.Int8(constant.StateDeleted)),
	)
	if e != nil {
		return e
	}
	return nil
}

// PageNurseByKey 分页查询护理员（staff 表中 role=护理员）
func (s *nurseGradeService) PageNurseByKey(ctx context.Context, in *dto.PageNurseByKeyReq, out *[]dto.PageNurseByKeyResp) error {
	q := db.Table(tblstaff.TableName).
		Where(ace.Where(tblstaff.TenantId.Eq(types.BigInt(lib.TenantID(ctx)))).
			And(tblstaff.RoleId.Eq(types.BigInt(nurseRoleID))).
			And(tblstaff.LeaveFlag.Eq(types.Int8(constant.YesNoNo))))
	if in.NurseName != nil {
		q = q.And(tblstaff.Name.Like(*in.NurseName))
	}
	if in.Key != nil {
		q = q.And(tblstaff.Name.Like(*in.Key))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblstaff.Id.As("id"),
			tblstaff.Name.As("name"),
			tblstaff.Phone.As("phone"),
		).
		Desc(tblstaff.CreateTime).
		Select().Gets(ctx, out)
}

// GetNurseById 查询护理员详情
func (s *nurseGradeService) GetNurseById(ctx context.Context, in *dto.IDReq, out *dto.GetNurseByIdResp) error {
	bean, has, e := dao.Staff(db).GetByID(ctx, types.BigInt(*in.ID), tblstaff.Id, tblstaff.Name, tblstaff.Phone)
	if e != nil {
		return e
	}
	if !has {
		return nil
	}
	out.ID = types.BigInt(bean.Id)
	out.Name = bean.Name.String()
	out.Phone = bean.Phone.String()
	return nil
}

// AddNurse 新增护理员
func (s *nurseGradeService) AddNurse(ctx context.Context, in *dto.AddNurseReq, out *dto.EmptyResp) error {
	bean := new(do.Staff)
	bean.TenantId = types.BigInt(lib.TenantID(ctx))
	bean.Name = types.String(*in.NurseName)
	bean.Phone = types.String(*in.Phone)
	bean.RoleId = types.BigInt(nurseRoleID)
	bean.Status = types.Int8(constant.YesNoNo)
	_, e := dao.Staff(db).InsertOne(ctx, bean, tblstaff.TenantId, tblstaff.Name, tblstaff.Phone, tblstaff.RoleId, tblstaff.LeaveFlag, tblstaff.CreateId, tblstaff.CreateTime)
	if e != nil {
		return e
	}
	return nil
}

// EditNurse 编辑护理员
func (s *nurseGradeService) EditNurse(ctx context.Context, in *dto.EditNurseReq, out *dto.EmptyResp) error {
	_, e := dao.Staff(db).UpdateById(ctx, types.BigInt(*in.ID),
		tblstaff.Name.Set(types.String(*in.NurseName)),
		tblstaff.Phone.Set(types.String(*in.Phone)),
	)
	if e != nil {
		return e
	}
	return nil
}

// DeleteNurse 删除护理员
func (s *nurseGradeService) DeleteNurse(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.Staff(db).DeleteById(ctx, types.BigInt(*in.ID))
	if e != nil {
		return e
	}
	return nil
}
