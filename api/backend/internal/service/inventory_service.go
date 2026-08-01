package service

import (
	"context"

	"api/internal/model/define/dao"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type inventory struct{}

var Inventory = &inventory{}

// PageInventoryByKey 分页查询库存（联表 material 获取物资名/类型）
// 对应 Java: InventoryServiceImpl.pageInventoryByKey -> InventoryMapper.listInventoryByKey
// SQL: SELECT inv.*, m.material_name, mt.type_name FROM inventory inv
//
//	LEFT JOIN material m ON m.id = inv.material_id
//	LEFT JOIN material_type mt ON mt.id = m.material_type_id
//	WHERE (m.material_name LIKE %key% OR inv.id = key) [可选]
//	ORDER BY inv.create_time DESC; 再由 PageUtil 内存分页。
//
// Todo: 1) in.Key 非空 -> (tbl<inventory>.Id.Eq(in.Key) OR tbl<material>.MaterialName.Like(in.Key))
//
//	2) DB 分页: Count + List(联表 LeftJoin material/material_type)
//	3) 组装含物资名/类型名的 VO 并赋值 out
func (i *inventory) PageInventoryByKey(ctx context.Context, in *dto.PageInventoryByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 实现联表分页查询
	return nil
}

// GetInventoryById 根据编号获取库存
// 对应 Java: InventoryServiceImpl.getInventoryById -> inventoryMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.Inventory(db).GetByID(ctx, types.BigInt(in.ID))
func (i *inventory) GetInventoryById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.Inventory(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// PageInventoryRecordByKey 分页查询库存记录（出库/入库流水）
// 对应 Java: InventoryServiceImpl.pageInventoryRecordByKey -> InventoryRecordMapper.listInventoryRecordByKey
// SQL: SELECT ir.*, m.material_name FROM inventory_record ir
//
//	LEFT JOIN material m ON m.id = ir.material_id
//	WHERE (m.material_name LIKE %key% OR ir.id = key) [可选]
//
// todo: 库存记录分页查询 - 联表 material + 分页, 结果赋值 out
func (i *inventory) PageInventoryRecordByKey(ctx context.Context, in *dto.PageInventoryRecordByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 inventory_record 联表 material 并分页
	return nil
}

// GetInventoryRecordById 根据编号获取库存记录
// 对应 Java: InventoryServiceImpl.getInventoryRecordById -> inventoryRecordMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.InventoryRecord(db).GetByID(ctx, types.BigInt(in.ID))
func (i *inventory) GetInventoryRecordById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.InventoryRecord(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// AddInventory 新增库存（同时写库存记录并调整 material 库存）
// 对应 Java: InventoryServiceImpl.addInventory -> insert inventory + insert inventory_record + 更新 material.stock
// todo: 事务: 1) dao.Inventory(db).InsertOne; 2) dao.InventoryRecord(db).InsertOne; 3) 更新 material 库存
func (i *inventory) AddInventory(ctx context.Context, in *dto.AddInventoryQuery, out *dto.EmptyResp) error {
	// todo: 写入 inventory + inventory_record 并同步 material.stock
	return nil
}

// EditInventory 编辑库存
// 对应 Java: InventoryServiceImpl.editInventory -> inventoryMapper.updateByPrimaryKeySelective
// todo: 标准 CRUD - 按主键更新 inventory 表
func (i *inventory) EditInventory(ctx context.Context, in *dto.EditInventoryQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<inventory>.Stock.Value(in.Stock),
	}
	_, e := dao.Inventory(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// AuditInventory 审核库存
// 对应 Java: InventoryServiceImpl.auditInventory -> 更新 inventory 审核状态
// todo: 更新 inventory 审核状态字段(UpdateById), 结果赋值 out
func (i *inventory) AuditInventory(ctx context.Context, in *dto.AuditInventoryQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: tbl<inventory>.AuditStatus.Value(in.AuditStatus),
	}
	_, e := dao.Inventory(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}
