package service

import (
	"context"

	"api/internal/model/define/dao"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type outboundrecord struct{}

var OutboundRecord = &outboundrecord{}

// PageOutboundRecordByKey 分页查询出库记录（联表 material、warehouse、elder、user）
// 对应 Java: OutboundRecordServiceImpl.pageOutboundRecordByKey -> OutboundRecordMapper.listOutboundRecordByKey
// SQL: SELECT ob.*, m.material_name, w.warehouse_name, e.elder_name, u.name AS charge_user_name
//
//	FROM outbound_record ob
//	LEFT JOIN material m ON m.id = ob.material_id
//	LEFT JOIN warehouse w ON w.id = ob.warehouse_id
//	LEFT JOIN elder e ON e.id = ob.elder_id
//	LEFT JOIN user u ON u.id = ob.charge_user_id
//	WHERE (m.material_name LIKE %key% OR ob.id = key) [可选]
//	ORDER BY ob.create_time DESC; 再由 PageUtil 内存分页。
//
// Todo: 1) in.Key 非空 -> (tbl<outboundrecord>.Id.Eq(in.Key) OR tbl<material>.MaterialName.Like(in.Key))
//
//	2) DB 分页: Count + List(联表 LeftJoin)
//	3) 组装含物资/仓库/老人/负责人名的 VO 并赋值 out
func (o *outboundrecord) PageOutboundRecordByKey(ctx context.Context, in *dto.PageOutboundRecordByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 实现联表分页查询
	return nil
}

// GetOutboundRecordById 根据编号获取出库记录
// 对应 Java: OutboundRecordServiceImpl.getOutboundRecordById -> outboundRecordMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.OutboundRecord(db).GetByID(ctx, types.BigInt(in.ID))
func (o *outboundrecord) GetOutboundRecordById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.OutboundRecord(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// PageSearchElderByKey 分页搜索老人（供出库选择老人）
// 对应 Java: OutboundRecordServiceImpl.pageSearchElderByKey -> elderMapper.listElderByKey
// SQL: SELECT * FROM elder WHERE (elder_name LIKE %key% OR id = key) [可选] AND del_flag=0
// todo: 查询 elder 表并分页, 结果赋值 out(需定义老人分页返回类型)
func (o *outboundrecord) PageSearchElderByKey(ctx context.Context, in *dto.PageSearchElderByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 elder 表并分页
	return nil
}

// PageWarehouseMaterialByKey 分页查询仓库物资（供出库选择物资）
// 对应 Java: OutboundRecordServiceImpl.pageWarehouseMaterialByKey -> warehouseMaterialMapper.listWarehouseMaterialByKey
// SQL: SELECT wm.*, m.material_name, w.warehouse_name FROM warehouse_material wm
//
//	LEFT JOIN material m ON m.id = wm.material_id
//	LEFT JOIN warehouse w ON w.id = wm.warehouse_id
//	WHERE (m.material_name LIKE %key%) [可选]
//
// todo: 仓库物资分页查询 - 联表 material/warehouse + 分页, 结果赋值 out
func (o *outboundrecord) PageWarehouseMaterialByKey(ctx context.Context, in *dto.PageWarehouseMaterialByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 warehouse_material 联表并分页
	return nil
}

// AddOutboundRecord 新增出库记录（扣减仓库物资库存）
// 对应 Java: OutboundRecordServiceImpl.addOutboundRecord -> insert outbound_record + 扣减 warehouse_material.stock
// todo: 事务: 1) dao.OutboundRecord(db).InsertOne; 2) 扣减对应 warehouse_material 库存
func (o *outboundrecord) AddOutboundRecord(ctx context.Context, in *dto.AddOutboundRecordQuery, out *dto.EmptyResp) error {
	// todo: 写入 outbound_record 并扣减 warehouse_material 库存
	return nil
}

// AuditOutboundRecord 审核出库记录
// 对应 Java: OutboundRecordServiceImpl.auditOutboundRecord -> 更新审核状态
// todo: 更新 outbound_record 审核状态字段(UpdateById), 结果赋值 out
func (o *outboundrecord) AuditOutboundRecord(ctx context.Context, in *dto.AuditOutboundRecordQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: tbl<outboundrecord>.AuditStatus.Value(in.AuditStatus),
	}
	_, e := dao.OutboundRecord(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteOutboundRecord 删除出库记录（回冲仓库物资库存）
// 对应 Java: OutboundRecordServiceImpl.deleteOutboundRecord -> 删记录 + 回冲 warehouse_material.stock
// todo: 事务: 1) 查原记录; 2) 回冲库存; 3) dao.OutboundRecord(db).DeleteById
func (o *outboundrecord) DeleteOutboundRecord(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.OutboundRecord(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}
