package dto

import "github.com/linbaozhong/gentity/pkg/types"

// ============ InventoryController 请求 ============

// @request
// PageInventoryByKeyQuery 分页查询库存请求
type PageInventoryByKeyQuery struct {
	PageNum      *int    `json:"page_num" valid:"required"`  // 页码
	PageSize     *int    `json:"page_size" valid:"required"` // 条数
	WarehouseID  *int64  `json:"warehouse_id"`               // 仓库编号
	MaterialName *string `json:"material_name"`              // 物资名称
}

// ============ InventoryController 响应 ============

// @response
// PageInventoryByKeyVO 分页查询库存响应
type PageInventoryByKeyVO struct {
	WarehouseName string      `json:"warehouse_name"` // 仓库名称
	MaterialID    int64       `json:"material_id"`    // 物资编号
	MaterialName  string      `json:"material_name"`  // 物资名称
	Total         int         `json:"total"`          // 总库存
	WarehouseNum  int         `json:"warehouse_num"`  // 入库数量
	Inventory     int         `json:"inventory"`      // 库存数量
	OutboundNum   int         `json:"outbound_num"`   // 出库数量
	Price         types.Money `json:"price"`          // 物资单价
}

// @request
// PageInventoryRecordByKeyQuery 分页查询库存记录请求
type PageInventoryRecordByKeyQuery struct {
	PageNum    *int    `json:"page_num" valid:"required"`  // 页码
	PageSize   *int    `json:"page_size" valid:"required"` // 条数
	Type       *string `json:"type"`                       // 记录类型(in/out)
	MaterialID *int64  `json:"material_id"`                // 物资编号
	Key        *string `json:"key"`                        // 关键字
}

// @request
// AddInventoryQuery 新增库存请求
type AddInventoryQuery struct {
	ID          *int64       `json:"id"`                            // id
	WarehouseID *int64       `json:"warehouse_id" valid:"required"` // 仓库编号
	MaterialID  *int64       `json:"material_id" valid:"required"`  // 物资编号
	Num         *int         `json:"num" valid:"required"`          // 数量
	Price       *types.Money `json:"price" valid:"required"`        // 单价
	Remark      *string      `json:"remark"`                        // 备注
}

// @request
// EditInventoryQuery 编辑库存请求
type EditInventoryQuery struct {
	ID          *int64       `json:"id"`                            // id
	WarehouseID *int64       `json:"warehouse_id" valid:"required"` // 仓库编号
	MaterialID  *int64       `json:"material_id" valid:"required"`  // 物资编号
	Num         *int         `json:"num" valid:"required"`          // 数量
	Price       *types.Money `json:"price" valid:"required"`        // 单价
	Remark      *string      `json:"remark"`                        // 备注
}

// @request
// AuditInventoryQuery 审核库存请求
type AuditInventoryQuery struct {
	ID          *int64  `json:"id" valid:"required"`           // 库存记录编号
	AuditResult *string `json:"audit_result" valid:"required"` // 审核结果
}
