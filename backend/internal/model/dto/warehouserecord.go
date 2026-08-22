package dto

import (
	"github.com/linbaozhong/gentity/pkg/types"
	"time"
)

// ============ WarehouseRecordController 请求 ============

// PageWarehouseRecordByKeyReq 分页查询入库记录请求
// @request
type PageWarehouseRecordByKeyReq struct {
	PageNum       *uint      `json:"page_num" valid:"required"`  // 页码
	PageSize      *uint      `json:"page_size" valid:"required"` // 条数
	WarehouseName *string    `json:"warehouse_name"`             // 仓库名称
	MaterialName  *string    `json:"material_name"`              // 物资名称
	StartTime     *time.Time `json:"start_time"`                 // 开始时间
	EndTime       *time.Time `json:"end_time"`                   // 结束时间
	StaffName     *string    `json:"staff_name"`                 // 经办人
}

// PageWarehouseMaterialByKeyReq 分页查询仓库物资请求
// @request
type PageWarehouseMaterialByKeyReq struct {
	PageNum      *uint   `json:"page_num" valid:"required"`     // 页码
	PageSize     *uint   `json:"page_size" valid:"required"`    // 条数
	WarehouseID  *int64  `json:"warehouse_id" valid:"required"` // 仓库编号
	MaterialName *string `json:"material_name"`                 // 物资名称
}

// AddWarehouseRecordReq 新增入库记录请求
// @request
type AddWarehouseRecordReq struct {
	WarehouseID                *int64                    `json:"warehouse_id" valid:"required"`                  // 仓库编号
	StaffID                    *int64                    `json:"staff_id" valid:"required"`                      // 经办人编号
	Source                     *string                   `json:"source" valid:"required"`                        // 物资来源
	WarehouseDate              *time.Time                `json:"warehouse_date" valid:"required"`                // 入库时间
	WarehouseMaterialQueryList []AddWarehouseMaterialReq `json:"warehouse_material_query_list" valid:"required"` // 入库物资列表
}

// AddWarehouseMaterialReq 新增入库物资请求（嵌套）
// @request
type AddWarehouseMaterialReq struct {
	MaterialID   *int64     `json:"material_id" valid:"required"`   // 物资编号
	WarehouseNum *int       `json:"warehouse_num" valid:"required"` // 入库数量
	ProductDate  *time.Time `json:"product_date" valid:"required"`  // 生产日期
	ExpireDate   *time.Time `json:"expire_date" valid:"required"`   // 有效期
}

// AuditWarehouseRecordReq 审核入库记录请求
// @request
type AuditWarehouseRecordReq struct {
	WarehouseRecordID *int64  `json:"warehouse_record_id" valid:"required"` // 入库记录编号
	AuditResult       *string `json:"audit_result" valid:"required"`        // 审核结果
}

// PageMaterialByKeyReq 分页查询物资请求（定义见 material.go）

// ============ WarehouseRecordController 响应 ============

// PageWarehouseRecordByKeyResp 入库记录响应
// @response
type PageWarehouseRecordByKeyResp struct {
	ID            int64     `json:"id"`             // id
	WarehouseName string    `json:"warehouse_name"` // 仓库名称
	MaterialName  string    `json:"material_name"`  // 物资名称
	WarehouseDate time.Time `json:"warehouse_date"` // 入库时间
	Source        string    `json:"source"`         // 物资来源
	StaffName     string    `json:"staff_name"`     // 经办人
	WarehouseFlag string    `json:"warehouse_flag"` // 入库状态
}

// PageWarehouseMaterialByKeyResp 分页查询仓库物资响应
// @response
type PageWarehouseMaterialByKeyResp struct {
	ID           int64       `json:"id"`            // id
	MaterialName string      `json:"material_name"` // 物资名称
	Price        types.Money `json:"price"`         // 物资单价
	WarehouseNum int         `json:"warehouse_num"` // 入库数量
	Inventory    int         `json:"inventory"`     // 库存数量
	ExpireDate   time.Time   `json:"expire_date"`   // 有效期
}

// GetWarehouseRecordByIDResp 根据编号查询入库记录响应
// @response
type GetWarehouseRecordByIDResp struct {
	WarehouseName                 string                         `json:"warehouse_name"`                   // 仓库名称
	StaffName                     string                         `json:"staff_name"`                       // 经办人姓名
	Source                        string                         `json:"source"`                           // 物资来源
	WarehouseDate                 time.Time                      `json:"warehouse_date"`                   // 入库时间
	WarehouseMaterialByIDRespList []GetWarehouseMaterialByIDResp `json:"warehouse_material_by_id_vo_list"` // 入库物资列表
}

// GetWarehouseMaterialByIDResp 入库物资响应（嵌套，继承 Rank）
// @response
type GetWarehouseMaterialByIDResp struct {
	MaterialName string    `json:"material_name"` // 物资名称
	WarehouseNum int       `json:"warehouse_num"` // 入库数量
	ProductDate  time.Time `json:"product_date"`  // 生产日期
	ExpireDate   time.Time `json:"expire_date"`   // 有效期
}
