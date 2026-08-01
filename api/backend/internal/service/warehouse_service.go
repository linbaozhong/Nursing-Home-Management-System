package service

import (
	"context"

	"api/internal/model/define/dao"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type warehouse struct{}

var Warehouse = &warehouse{}

// PageWarehouseByKey 分页查询仓库
// 对应 Java: WarehouseServiceImpl.pageWarehouseByKey -> WarehouseMapper.listWarehouseByKey
// SQL: SELECT * FROM warehouse WHERE (warehouse_name LIKE %key%) [可选] ORDER BY create_time DESC
// todo: 仓库分页查询 - dao.Warehouse(db) 条件 + 分页, 结果赋值 out
func (w *warehouse) PageWarehouseByKey(ctx context.Context, in *dto.PageWarehouseByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 warehouse 表并分页
	return nil
}

// ListWarehouseStaff 仓库管理员列表
// 对应 Java: WarehouseServiceImpl.listWarehouseStaff -> userMapper.listWarehouseStaff(角色=仓库管理员)
// SQL: SELECT * FROM user WHERE role = 仓库管理员
// todo: 查询 user(仓库管理员)列表, 结果赋值 out(需定义返回类型)
func (w *warehouse) ListWarehouseStaff(ctx context.Context, in *dto.EmptyReq, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 user 仓库管理员列表
	return nil
}

// AddWarehouse 新增仓库
// 对应 Java: WarehouseServiceImpl.addWarehouse -> warehouseMapper.insertSelective
// todo: 标准 CRUD - dao.Warehouse(db).InsertOne 写入 warehouse 表(含 warehouseName/managerUserId 等)
func (w *warehouse) AddWarehouse(ctx context.Context, in *dto.OperateWarehouseQuery, out *dto.EmptyResp) error {
	// todo: bean := do.NewWarehouse(); 填充 in; dao.Warehouse(db).InsertOne(ctx, bean)
	return nil
}

// GetWarehouseById 根据编号获取仓库
// 对应 Java: WarehouseServiceImpl.getWarehouseById -> warehouseMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.Warehouse(db).GetByID(ctx, types.BigInt(in.ID))
func (w *warehouse) GetWarehouseById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.Warehouse(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// EditWarehouse 编辑仓库
// 对应 Java: WarehouseServiceImpl.editWarehouse -> warehouseMapper.updateByPrimaryKeySelective
// todo: 标准 CRUD - 按主键更新 warehouse 表
func (w *warehouse) EditWarehouse(ctx context.Context, in *dto.OperateWarehouseQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<warehouse>.WarehouseName.Value(in.WarehouseName),
	}
	_, e := dao.Warehouse(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteWarehouse 删除仓库
// 对应 Java: WarehouseServiceImpl.deleteWarehouse -> warehouseMapper.deleteByPrimaryKey
// todo: 标准 CRUD - dao.Warehouse(db).DeleteById(ctx, types.BigInt(in.ID))
func (w *warehouse) DeleteWarehouse(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.Warehouse(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}
