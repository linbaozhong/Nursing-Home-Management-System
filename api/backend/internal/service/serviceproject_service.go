package service

import (
	"context"

	"api/internal/model/define/dao"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type serviceproject struct{}

var ServiceProject = &serviceproject{}

// ListServiceType 获取服务类型列表
// 对应 Java: ServiceProjectServiceImpl.listServiceType -> serviceTypeMapper.selectAll(按 name 过滤)
// SQL: SELECT * FROM service_type WHERE (type_name LIKE %name%) [可选]
// todo: 查询 service_type 列表, 结果赋值 out(需定义返回类型)
func (s *serviceproject) ListServiceType(ctx context.Context, in *dto.NameReq, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 service_type 列表
	return nil
}

// PageServiceByKey 分页查询服务（联表 service_type 获取类型名）
// 对应 Java: ServiceProjectServiceImpl.pageServiceByKey -> ServiceProjectMapper.listServiceByKey
// SQL: SELECT sp.*, st.type_name FROM service_project sp
//
//	LEFT JOIN service_type st ON st.id = sp.service_type_id
//	WHERE (st.type_name LIKE %key% OR sp.service_name LIKE %key%) [可选]
//	ORDER BY sp.create_time DESC; 再由 PageUtil 内存分页。
//
// todo: 1) in.Key 非空 -> (tbl<servicetype>.TypeName.Like(in.Key) OR tbl<serviceproject>.ServiceName.Like(in.Key))
//
//	2) DB 分页: Count + List(联表 LeftJoin service_type)
//	3) 组装含类型名的 VO 并赋值 out
func (s *serviceproject) PageServiceByKey(ctx context.Context, in *dto.PageServiceByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 实现联表分页查询
	return nil
}

// AddServiceType 新增服务类型
// 对应 Java: ServiceProjectServiceImpl.addServiceType -> serviceTypeMapper.insertSelective
// todo: 标准 CRUD - dao.ServiceType(db).InsertOne 写入 service_type 表
func (s *serviceproject) AddServiceType(ctx context.Context, in *dto.OperateServiceTypeQuery, out *dto.EmptyResp) error {
	// todo: bean := do.NewServiceType(); 填充 in; dao.ServiceType(db).InsertOne(ctx, bean)
	return nil
}

// GetServiceTypeById 根据编号获取服务类型
// 对应 Java: ServiceProjectServiceImpl.getServiceTypeById -> serviceTypeMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.ServiceType(db).GetByID(ctx, types.BigInt(in.ID))
func (s *serviceproject) GetServiceTypeById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.ServiceType(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// EditServiceType 编辑服务类型
// 对应 Java: ServiceProjectServiceImpl.editServiceType -> serviceTypeMapper.updateByPrimaryKeySelective
// todo: 标准 CRUD - 按主键更新 service_type 表
func (s *serviceproject) EditServiceType(ctx context.Context, in *dto.OperateServiceTypeQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<servicetype>.TypeName.Value(in.TypeName),
	}
	_, e := dao.ServiceType(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteServiceType 删除服务类型
// 对应 Java: ServiceProjectServiceImpl.deleteServiceType -> serviceTypeMapper.deleteByPrimaryKey
// todo: 标准 CRUD - dao.ServiceType(db).DeleteById(ctx, types.BigInt(in.ID))
func (s *serviceproject) DeleteServiceType(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.ServiceType(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}

// AddService 新增服务
// 对应 Java: ServiceProjectServiceImpl.addService -> serviceProjectMapper.insertSelective
// todo: 标准 CRUD - dao.ServiceProject(db).InsertOne 写入 service_project 表(含 serviceName/serviceTypeId/price 等)
func (s *serviceproject) AddService(ctx context.Context, in *dto.OperateServiceQuery, out *dto.EmptyResp) error {
	// todo: bean := do.NewServiceProject(); 填充 in; dao.ServiceProject(db).InsertOne(ctx, bean)
	return nil
}

// GetServiceById 根据编号获取服务
// 对应 Java: ServiceProjectServiceImpl.getServiceById -> serviceProjectMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.ServiceProject(db).GetByID(ctx, types.BigInt(in.ID))
func (s *serviceproject) GetServiceById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.ServiceProject(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// EditService 编辑服务
// 对应 Java: ServiceProjectServiceImpl.editService -> serviceProjectMapper.updateByPrimaryKeySelective
// todo: 标准 CRUD - 按主键更新 service_project 表
func (s *serviceproject) EditService(ctx context.Context, in *dto.OperateServiceQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<serviceproject>.ServiceName.Value(in.ServiceName),
	}
	_, e := dao.ServiceProject(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteService 删除服务
// 对应 Java: ServiceProjectServiceImpl.deleteService -> serviceProjectMapper.deleteByPrimaryKey
// todo: 标准 CRUD - dao.ServiceProject(db).DeleteById(ctx, types.BigInt(in.ID))
func (s *serviceproject) DeleteService(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.ServiceProject(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}
