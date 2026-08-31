package service

import (
	"context"

	"api/internal/lib"
	"api/internal/model/define/table/tblmaterial"
	"api/internal/model/define/table/tblmaterialtype"
	"api/internal/model/define/table/tblstock"
	"api/internal/model/define/table/tblwarehouse"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/types"
)

type inventory struct{}

var Inventory = &inventory{}

// PageInventoryByKey 分页查询库存（基于 stock 库存台账，按仓库过滤，total 按物资跨仓库汇总）
func (i *inventory) PageInventoryByKey(ctx context.Context, in *dto.PageInventoryByKeyReq, out *[]dto.PageInventoryByKeyResp) error {
	q := db.Table(tblstock.TableName).
		InnerJoin(tblstock.WarehouseId, tblwarehouse.Id).
		InnerJoin(tblstock.MaterialId, tblmaterial.Id).
		LeftJoin(tblmaterial.TypeId, tblmaterialtype.Id).
		Where(tblstock.TenantId.Eq(types.BigInt(lib.TenantID(ctx))))
	if in.WarehouseID != nil {
		q = q.And(tblstock.WarehouseId.Eq(types.BigInt(*in.WarehouseID)))
	}
	if in.MaterialName != nil && *in.MaterialName != "" {
		q = q.And(tblmaterial.Name.Like(*in.MaterialName))
	}

	if err := q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblwarehouse.Name.As("warehouse_name"),
			tblmaterial.Id.As("material_id"),
			tblmaterial.Name.As("material_name"),
			tblstock.Qty.As("inventory"),
			tblmaterial.Price.As("price"),
		).
		Desc(tblstock.Id).
		Select().
		Gets(ctx, out); err != nil {
		return err
	}
	// 填充 total：同一 material_id 在所有仓库中的库存之和
	// 注：inventory 里同物资可能有多个批次行，跨仓库的 total 按 material_id 汇总
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
func (i *inventory) sumInventoryByMaterial(ctx context.Context, materialID types.BigInt) (int, error) {
	rows, e := db.Table(tblstock.TableName).
		Cols(tblstock.Qty).
		Where(tblstock.MaterialId.Eq(materialID)).
		Select().
		Query(ctx)
	if e != nil {
		return 0, e
	}
	defer rows.Close()
	total := 0
	for rows.Next() {
		var q int
		if e := rows.Scan(&q); e != nil {
			return 0, e
		}
		total += q
	}
	return total, rows.Err()
}

// GetInventoryById Java InventoryServiceImpl 未实现, 保留占位
func (i *inventory) GetInventoryById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	return nil
}

// AddInventory Java InventoryServiceImpl 未实现, 保留占位
func (i *inventory) AddInventory(ctx context.Context, in *dto.AddInventoryReq, out *dto.EmptyResp) error {
	return nil
}

// EditInventory Java InventoryServiceImpl 未实现, 保留占位
func (i *inventory) EditInventory(ctx context.Context, in *dto.EditInventoryReq, out *dto.EmptyResp) error {
	return nil
}

// AuditInventory Java InventoryServiceImpl 未实现, 保留占位
func (i *inventory) AuditInventory(ctx context.Context, in *dto.AuditInventoryReq, out *dto.EmptyResp) error {
	return nil
}

// PageInventoryRecordByKey Java InventoryServiceImpl 未实现, 保留占位
func (i *inventory) PageInventoryRecordByKey(ctx context.Context, in *dto.PageInventoryRecordByKeyReq, out *[]dto.PageInventoryByKeyResp) error {
	return nil
}

// GetInventoryRecordById Java InventoryServiceImpl 未实现, 保留占位
func (i *inventory) GetInventoryRecordById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	return nil
}
