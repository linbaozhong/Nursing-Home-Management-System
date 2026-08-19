package service

import (
	"context"

	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblmaterial"
	"api/internal/model/define/table/tblmaterialtype"
	"api/internal/model/define/table/tblwarehouse"
	"api/internal/model/define/table/tblwarehousematerial"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/types"
)

type inventory struct{}

var Inventory = &inventory{}

// PageInventoryByKey 分页查询库存
// 对应 Java: InventoryServiceImpl.pageInventoryByKey
// 说明: Java 的 InventoryServiceImpl 实际只实现此方法, 查询 warehouse_material 表联 material/material_type/warehouse;
//
//	其余 Go 桩方法(增删改/审核/记录)在 Java 中不存在, 故保留占位。
//
// 注: warehouse_material 表无 outbound_num/price 列, outbound_num 暂为 0, price 取自 material.price;
//
//	Java 的 total(按 material_id 跨记录求和 inventory) 由 sumInventoryByMaterial 计算填充。
func (i *inventory) PageInventoryByKey(ctx context.Context, in *dto.PageInventoryByKeyQuery, out *[]dto.PageInventoryByKeyVO) error {
	q := db.Table(tblwarehousematerial.TableName).
		LeftJoin(tblwarehousematerial.MaterialId, tblmaterial.Id).
		LeftJoin(tblmaterial.TypeId, tblmaterialtype.Id).
		LeftJoin(tblwarehousematerial.WarehouseRecordId, tblwarehouse.Id)
	if in.WarehouseID != nil {
		q.And(tblwarehousematerial.WarehouseRecordId.Eq(types.BigInt(*in.WarehouseID)))
	}
	if in.MaterialName != nil && *in.MaterialName != "" {
		q.And(tblmaterial.Name.Like(*in.MaterialName))
	}
	if in.Key != nil && *in.Key != "" {
		q.And(tblmaterial.Name.Like(*in.Key))
	}
	if err := q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblwarehouse.Name.AsName("warehouse_name"),
			tblmaterial.Id.AsName("material_id"),
			tblmaterial.Name.AsName("material_name"),
			tblwarehousematerial.WarehouseNum.AsName("warehouse_num"),
			tblwarehousematerial.Inventory.AsName("inventory"),
			tblmaterial.Price.AsName("price"),
		).
		Desc(tblwarehousematerial.CreateTime).
		Select().
		Gets(ctx, out); err != nil {
		return err
	}
	// 填充 total: 同一 material_id 在所有仓库中的 inventory 之和
	for idx := range *out {
		total, err := i.sumInventoryByMaterial(ctx, (*out)[idx].MaterialID)
		if err != nil {
			return err
		}
		(*out)[idx].Total = total
	}
	return nil
}

// sumInventoryByMaterial 按 material_id 汇总所有仓库的库存数量
func (i *inventory) sumInventoryByMaterial(ctx context.Context, materialID int64) (int, error) {
	list, _, e := dao.WarehouseMaterial(db).List(ctx,
		db.Table(tblwarehousematerial.TableName).
			Cols(tblwarehousematerial.Inventory).
			Where(tblwarehousematerial.MaterialId.Eq(types.BigInt(materialID))),
	)
	if e != nil {
		return 0, e
	}
	total := 0
	for _, v := range list {
		total += int(v.Inventory)
	}
	return total, nil
}

// GetInventoryById Java InventoryServiceImpl 未实现, 保留占位
func (i *inventory) GetInventoryById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	return nil
}

// AddInventory Java InventoryServiceImpl 未实现, 保留占位
func (i *inventory) AddInventory(ctx context.Context, in *dto.OperateInventoryQuery, out *dto.EmptyResp) error {
	return nil
}

// EditInventory Java InventoryServiceImpl 未实现, 保留占位
func (i *inventory) EditInventory(ctx context.Context, in *dto.OperateInventoryQuery, out *dto.EmptyResp) error {
	return nil
}

// AuditInventory Java InventoryServiceImpl 未实现, 保留占位
func (i *inventory) AuditInventory(ctx context.Context, in *dto.OperateInventoryQuery, out *dto.EmptyResp) error {
	return nil
}

// PageInventoryRecordByKey Java InventoryServiceImpl 未实现, 保留占位
func (i *inventory) PageInventoryRecordByKey(ctx context.Context, in *dto.PageInventoryRecordByKeyQuery, out *[]dto.PageInventoryRecordByKeyVO) error {
	return nil
}

// GetInventoryRecordById Java InventoryServiceImpl 未实现, 保留占位
func (i *inventory) GetInventoryRecordById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	return nil
}
