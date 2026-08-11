package service

import (
	"context"
	"errors"

	"api/internal/constant"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblserviceitem"
	"api/internal/model/define/table/tblservicetype"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/types"
)

// serviceTotalLimit 服务 / 服务类型数量上限（对应 Java Constant.TOTAL_LIMIT）
const serviceTotalLimit = 10

type serviceproject struct{}

var ServiceProject = &serviceproject{}

// ListServiceType 获取服务类型下拉列表（未删除）
// 对应 Java: ServiceProjectServiceImpl.listServiceType -> ServiceTypeFunc.listNotDelServiceType
func (s *serviceproject) ListServiceType(ctx context.Context, in *dto.OperateServiceTypeQuery, out *[]dto.DropDown) error {
	list, _, e := dao.ServiceType(db).List(ctx,
		db.Table(tblservicetype.TableName).
			Where(tblservicetype.DelFlag.Eq(constant.YesNoNo)),
	)
	if e != nil {
		return e
	}
	*out = make([]dto.DropDown, 0, len(list))
	for _, v := range list {
		*out = append(*out, dto.DropDown{
			ID:   int64(v.Id),
			Name: v.Name.String(),
		})
	}
	return nil
}

// PageServiceByKey 分页查询服务项目（关联服务类型名称）
// 对应 Java: ServiceProjectServiceImpl.pageServiceByKey -> ServiceItemFunc.listNotDelServiceItemByKey (ChargeEnum.ALL)
func (s *serviceproject) PageServiceByKey(ctx context.Context, in *dto.PageServiceByKeyQuery, out *[]dto.PageServiceByKeyVO) error {
	q := db.Table(tblserviceitem.TableName).
		LeftJoin(tblserviceitem.TypeId, tblservicetype.Id).
		Where(tblserviceitem.DelFlag.Eq(constant.YesNoNo))
	if in.ServiceName != nil && *in.ServiceName != "" {
		q.And(tblserviceitem.Name.Like(*in.ServiceName))
	}
	if in.TypeName != nil && *in.TypeName != "" {
		q.And(tblservicetype.Name.Like(*in.TypeName))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblserviceitem.Id.AsName("id"),
			tblservicetype.Name.AsName("type_name"),
			tblserviceitem.Name.AsName("service_name"),
			tblserviceitem.ChargeMethod.AsName("charge_method"),
			tblserviceitem.Price.AsName("price"),
			tblserviceitem.NeedDate.AsName("need_date"),
		).
		Desc(tblserviceitem.UpdateTime).
		Select().
		Gets(ctx, out)
}

// GetServiceById 根据编号获取服务项目（编辑回显）
// 对应 Java: ServiceProjectServiceImpl.getServiceById
func (s *serviceproject) GetServiceById(ctx context.Context, in *dto.IDReq, out *dto.OperateServiceQuery) error {
	obj, has, e := dao.ServiceItem(db).GetByID(ctx, types.BigInt(*in.ID))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("服务项目不存在")
	}
	out.ID = int64Ptr(int64(obj.Id))
	out.TypeID = int64Ptr(int64(obj.TypeId))
	out.Name = strPtr(obj.Name.String())
	out.ChargeMethod = strPtr(obj.ChargeMethod.String())
	out.Price = float64Ptr(obj.Price.Float64())
	out.NeedDate = intPtr(int(obj.NeedDate))
	return nil
}

// AddService 新增服务项目（同一类型下名称不重复 + 类型数量上限）
// 对应 Java: ServiceProjectServiceImpl.addService -> ServiceItemFunc.getServiceItemByName / checkServiceTotal
func (s *serviceproject) AddService(ctx context.Context, in *dto.OperateServiceQuery, out *dto.EmptyResp) error {
	repeat, e := dao.ServiceItem(db).Exists(ctx,
		tblserviceitem.TypeId.Eq(types.BigInt(*in.TypeID)),
		tblserviceitem.Name.Eq(*in.Name),
		tblserviceitem.DelFlag.Eq(types.Int8(constant.YesNoNo)),
	)
	if e != nil {
		return e
	}
	if repeat {
		return errors.New("该服务类型下服务名称已存在")
	}
	total, e := dao.ServiceItem(db).Count(ctx,
		tblserviceitem.TypeId.Eq(types.BigInt(*in.TypeID)),
		tblserviceitem.DelFlag.Eq(types.Int8(constant.YesNoNo)),
	)
	if e != nil {
		return e
	}
	if int(total) >= serviceTotalLimit {
		return errors.New("该服务类型下服务项目数量已达上限")
	}
	bean := do.NewServiceItem()
	bean.TypeId = types.BigInt(*in.TypeID)
	bean.Name = types.String(*in.Name)
	bean.ChargeMethod = types.String(*in.ChargeMethod)
	bean.Price = types.Float64(*in.Price)
	bean.NeedDate = types.Int32(*in.NeedDate)
	bean.DelFlag = types.Int8(constant.YesNoNo)
	_, e = dao.ServiceItem(db).InsertOne(ctx, bean)
	return e
}

// EditService 编辑服务项目（同一类型下名称不重复排除自身 + 类型数量上限）
// 对应 Java: ServiceProjectServiceImpl.editService
func (s *serviceproject) EditService(ctx context.Context, in *dto.OperateServiceQuery, out *dto.EmptyResp) error {
	repeat, e := dao.ServiceItem(db).Exists(ctx,
		tblserviceitem.TypeId.Eq(types.BigInt(*in.TypeID)),
		tblserviceitem.Name.Eq(*in.Name),
		tblserviceitem.DelFlag.Eq(types.Int8(constant.YesNoNo)),
		tblserviceitem.Id.NotEq(types.BigInt(*in.ID)),
	)
	if e != nil {
		return e
	}
	if repeat {
		return errors.New("该服务类型下服务名称已存在")
	}
	total, e := dao.ServiceItem(db).Count(ctx,
		tblserviceitem.TypeId.Eq(types.BigInt(*in.TypeID)),
		tblserviceitem.DelFlag.Eq(types.Int8(constant.YesNoNo)),
	)
	if e != nil {
		return e
	}
	if int(total) >= serviceTotalLimit {
		return errors.New("该服务类型下服务项目数量已达上限")
	}
	bean := do.NewServiceItem()
	bean.Id = types.BigInt(*in.ID)
	bean.TypeId = types.BigInt(*in.TypeID)
	bean.Name = types.String(*in.Name)
	bean.ChargeMethod = types.String(*in.ChargeMethod)
	bean.Price = types.Float64(*in.Price)
	bean.NeedDate = types.Int32(*in.NeedDate)
	_, e = dao.ServiceItem(db).UpdateOne(ctx, bean)
	return e
}

// DeleteService 删除服务项目（逻辑删除）
// 对应 Java: ServiceProjectServiceImpl.deleteService -> del_flag = YES
func (s *serviceproject) DeleteService(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	bean := do.NewServiceItem()
	bean.Id = types.BigInt(*in.ID)
	bean.DelFlag = types.Int8(constant.YesNoYes)
	_, e := dao.ServiceItem(db).UpdateOne(ctx, bean)
	return e
}

// AddServiceType 新增服务类型（名称不重复 + 总数上限）
// 对应 Java: ServiceProjectServiceImpl.addServiceType -> ServiceTypeFunc.getServiceTypeByName / checkTypeTotal
func (s *serviceproject) AddServiceType(ctx context.Context, in *dto.OperateServiceTypeQuery, out *dto.EmptyResp) error {
	repeat, e := dao.ServiceType(db).Exists(ctx,
		tblservicetype.Name.Eq(*in.Name),
		tblservicetype.DelFlag.Eq(types.Int8(constant.YesNoNo)),
	)
	if e != nil {
		return e
	}
	if repeat {
		return errors.New("服务类型名称已存在")
	}
	total, e := dao.ServiceType(db).Count(ctx, tblservicetype.DelFlag.Eq(constant.YesNoNo))
	if e != nil {
		return e
	}
	if int(total) >= serviceTotalLimit {
		return errors.New("服务类型数量已达上限")
	}
	bean := do.NewServiceType()
	bean.Name = types.String(*in.Name)
	bean.DelFlag = types.Int8(constant.YesNoNo)
	_, e = dao.ServiceType(db).InsertOne(ctx, bean)
	return e
}

// GetServiceTypeById 根据编号获取服务类型（编辑回显）
// 对应 Java: ServiceProjectServiceImpl.getServiceTypeById
func (s *serviceproject) GetServiceTypeById(ctx context.Context, in *dto.IDReq, out *dto.OperateServiceTypeQuery) error {
	obj, has, e := dao.ServiceType(db).GetByID(ctx, types.BigInt(*in.ID))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("服务类型不存在")
	}
	out.ID = int64Ptr(int64(obj.Id))
	out.Name = strPtr(obj.Name.String())
	return nil
}

// EditServiceType 编辑服务类型（名称不重复排除自身 + 总数上限）
// 对应 Java: ServiceProjectServiceImpl.editServiceType
func (s *serviceproject) EditServiceType(ctx context.Context, in *dto.OperateServiceTypeQuery, out *dto.EmptyResp) error {
	repeat, e := dao.ServiceType(db).Exists(ctx,
		tblservicetype.Name.Eq(*in.Name),
		tblservicetype.DelFlag.Eq(types.Int8(constant.YesNoNo)),
		tblservicetype.Id.NotEq(types.BigInt(*in.ID)),
	)
	if e != nil {
		return e
	}
	if repeat {
		return errors.New("服务类型名称已存在")
	}
	total, e := dao.ServiceType(db).Count(ctx, tblservicetype.DelFlag.Eq(constant.YesNoNo))
	if e != nil {
		return e
	}
	if int(total) >= serviceTotalLimit {
		return errors.New("服务类型数量已达上限")
	}
	bean := do.NewServiceType()
	bean.Id = types.BigInt(*in.ID)
	bean.Name = types.String(*in.Name)
	_, e = dao.ServiceType(db).UpdateOne(ctx, bean)
	return e
}

// DeleteServiceType 删除服务类型（存在子服务则不允许删除）
// 对应 Java: ServiceProjectServiceImpl.deleteServiceType -> ServiceItemFunc.checkServiceItem
func (s *serviceproject) DeleteServiceType(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	hasChild, e := dao.ServiceItem(db).Exists(ctx,
		tblserviceitem.TypeId.Eq(types.BigInt(*in.ID)),
		tblserviceitem.DelFlag.Eq(types.Int8(constant.YesNoNo)),
	)
	if e != nil {
		return e
	}
	if hasChild {
		return errors.New("该服务类型下存在服务项目，无法删除")
	}
	bean := do.NewServiceType()
	bean.Id = types.BigInt(*in.ID)
	bean.DelFlag = types.Int8(constant.YesNoYes)
	_, e = dao.ServiceType(db).UpdateOne(ctx, bean)
	return e
}
