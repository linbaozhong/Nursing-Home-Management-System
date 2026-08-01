package service

import (
	"context"

	"api/internal/model/define/dao"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type warehouserecord struct{}

var WarehouseRecord = &warehouserecord{}

// PageWarehouseRecordByKey 分页查询入库记录（联表 material、warehouse、user）
// 对应 Java: WarehouseRecordServiceImpl.pageWarehouseRecordByKey -> WarehouseRecordMapper.listWarehouseRecordByKey
// SQL: SELECT wr.*, m.material_name, w.warehouse_name, u.name AS charge_user_name
//
//	FROM warehouse_record wr
//	LEFT JOIN material m ON m.id = wr.material_id
//	LEFT JOIN warehouse w ON w.id = wr.warehouse_id
//	LEFT JOIN user u ON u.id = wr.charge_user_id
//	WHERE (m.material_name LIKE %key% OR wr.id = key) [可选]
//	ORDER BY wr.create_time DESC; 再由 PageUtil 内存分页。
//
// Todo: 1) in.Key 非空 -> (tbl<warehouserecord>.Id.Eq(in.Key) OR tbl<material>.MaterialName.Like(in.Key))
//
//	2) DB 分页: Count + List(联表 LeftJoin)
//	3) 组装含物资/仓库/负责人名的 VO 并赋值 out
func (w *warehouserecord) PageWarehouseRecordByKey(ctx context.Context, in *dto.PageWarehouseRecordByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 实现联表分页查询
	return nil
}

// ListWarehouse 获取仓库列表
// 对应 Java: WarehouseRecordServiceImpl.listWarehouse -> warehouseMapper.selectAll
// SQL: SELECT * FROM warehouse
// todo: 查询 warehouse 列表, 结果赋值 out(需定义返回类型)
func (w *warehouserecord) ListWarehouse(ctx context.Context, in *dto.EmptyReq, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 warehouse 列表
	return nil
}

// PageMaterialByKey 分页查询物资（供入库选择物资）
// 对应 Java: WarehouseRecordServiceImpl.pageMaterialByKey -> materialMapper.listMaterialByKey
// SQL: SELECT m.*, mt.type_name FROM material m LEFT JOIN material_type mt ON mt.id = m.material_type_id
//
//	WHERE (m.material_name LIKE %key% OR mt.type_name LIKE %key%) [可选]
//
// todo: 物资分页查询 - 联表 material_type + 分页, 结果赋值 out
func (w *warehouserecord) PageMaterialByKey(ctx context.Context, in *dto.PageMaterialByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 material 联表 material_type 并分页
	return nil
}

// AddWarehouseRecord 新增入库记录（增加仓库物资库存与 material 库存）
// 对应 Java: WarehouseRecordServiceImpl.addWarehouseRecord -> insert warehouse_record + 增 warehouse_material + 更新 material.stock
// todo: 事务: 1) dao.WarehouseRecord(db).InsertOne; 2) 写/更新 warehouse_material; 3) 更新 material.stock
func (w *warehouserecord) AddWarehouseRecord(ctx context.Context, in *dto.AddWarehouseRecordQuery, out *dto.EmptyResp) error {
	// todo: 写入 warehouse_record 并同步 warehouse_material/material 库存
	return nil
}

// GetWarehouseRecordById 根据编号获取入库记录
// 对应 Java: WarehouseRecordServiceImpl.getWarehouseRecordById -> warehouseRecordMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.WarehouseRecord(db).GetByID(ctx, types.BigInt(in.ID))
func (w *warehouserecord) GetWarehouseRecordById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.WarehouseRecord(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// AuditWarehouseRecord 审核入库记录
// 对应 Java: WarehouseRecordServiceImpl.auditWarehouseRecord -> 更新审核状态(审核通过后实际入库)
// todo: 更新 warehouse_record 审核状态(UpdateById); 审核通过则执行入库库存联动, 结果赋值 out
func (w *warehouserecord) AuditWarehouseRecord(ctx context.Context, in *dto.AuditWarehouseRecordQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: tbl<warehouserecord>.AuditStatus.Value(in.AuditStatus),
	}
	_, e := dao.WarehouseRecord(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteWarehouseRecord 删除入库记录（回冲库存）
// 对应 Java: WarehouseRecordServiceImpl.deleteWarehouseRecord -> 删记录 + 回冲 warehouse_material/material 库存
// todo: 事务: 1) 查原记录; 2) 回冲库存; 3) dao.WarehouseRecord(db).DeleteById
func (w *warehouserecord) DeleteWarehouseRecord(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.WarehouseRecord(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}
