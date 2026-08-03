package dto

import "time"

// ============ WarehouseRecordController 请求 ============

// @request
// PageWarehouseRecordByKeyQuery 分页查询入库记录请求
type PageWarehouseRecordByKeyQuery struct {
	PageNum       *int    `json:"page_num" valid:"required"`  // 页码
	PageSize      *int    `json:"page_size" valid:"required"` // 条数
	WarehouseName *string `json:"warehouse_name"`             // 仓库名称
	MaterialName  *string `json:"material_name"`              // 物资名称
	StartTime     *string `json:"start_time"`                 // 开始时间
	EndTime       *string `json:"end_time"`                   // 结束时间
	StaffName     *string `json:"staff_name"`                 // 经办人
}

// @request
// PageWarehouseMaterialByKeyQuery 分页查询仓库物资请求
type PageWarehouseMaterialByKeyQuery struct {
	PageNum      *int    `json:"page_num" valid:"required"`     // 页码
	PageSize     *int    `json:"page_size" valid:"required"`    // 条数
	WarehouseID  *int64  `json:"warehouse_id" valid:"required"` // 仓库编号
	MaterialName *string `json:"material_name"`                 // 物资名称
}

// @request
// AddWarehouseRecordQuery 新增入库记录请求
type AddWarehouseRecordQuery struct {
	WarehouseID                *int64                      `json:"warehouse_id" valid:"required"`                  // 仓库编号
	StaffID                    *int64                      `json:"staff_id" valid:"required"`                      // 经办人编号
	Source                     *string                     `json:"source" valid:"required"`                        // 物资来源
	WarehouseDate              *string                     `json:"warehouse_date" valid:"required"`                // 入库时间
	WarehouseMaterialQueryList []AddWarehouseMaterialQuery `json:"warehouse_material_query_list" valid:"required"` // 入库物资列表
}

// @request
// AddWarehouseMaterialQuery 新增入库物资请求（嵌套）
type AddWarehouseMaterialQuery struct {
	MaterialID   *int64  `json:"material_id" valid:"required"`   // 物资编号
	WarehouseNum *int    `json:"warehouse_num" valid:"required"` // 入库数量
	ProductDate  *string `json:"product_date" valid:"required"`  // 生产日期
	ExpireDate   *string `json:"expire_date" valid:"required"`   // 有效期
}

// @request
// AuditWarehouseRecordQuery 审核入库记录请求
type AuditWarehouseRecordQuery struct {
	WarehouseRecordID *int64  `json:"warehouse_record_id" valid:"required"` // 入库记录编号
	AuditResult       *string `json:"audit_result" valid:"required"`        // 审核结果
}

// PageMaterialByKeyQuery 分页查询物资请求（定义见 material.go）

// ============ WarehouseRecordController 响应 ============

// @response
// PageWarehouseRecordByKeyVO 入库记录响应
type PageWarehouseRecordByKeyVO struct {
	Rank
	ID            int64     `json:"id"`             // id
	WarehouseName string    `json:"warehouse_name"` // 仓库名称
	MaterialName  string    `json:"material_name"`  // 物资名称
	WarehouseDate time.Time `json:"warehouse_date"` // 入库时间
	Source        string    `json:"source"`         // 物资来源
	StaffName     string    `json:"staff_name"`     // 经办人
	WarehouseFlag string    `json:"warehouse_flag"` // 入库状态
}

// @response
// PageWarehouseMaterialByKeyVO 分页查询仓库物资响应
type PageWarehouseMaterialByKeyVO struct {
	Rank
	ID           int64     `json:"id"`            // id
	MaterialName string    `json:"material_name"` // 物资名称
	Price        float64   `json:"price"`         // 物资单价
	WarehouseNum int       `json:"warehouse_num"` // 入库数量
	Inventory    int       `json:"inventory"`     // 库存数量
	ExpireDate   time.Time `json:"expire_date"`   // 有效期
}

// @response
// GetWarehouseRecordByIDVO 根据编号查询入库记录响应
type GetWarehouseRecordByIDVO struct {
	WarehouseName               string                       `json:"warehouse_name"`                   // 仓库名称
	StaffName                   string                       `json:"staff_name"`                       // 经办人姓名
	Source                      string                       `json:"source"`                           // 物资来源
	WarehouseDate               time.Time                    `json:"warehouse_date"`                   // 入库时间
	WarehouseMaterialByIDVOList []GetWarehouseMaterialByIDVO `json:"warehouse_material_by_id_vo_list"` // 入库物资列表
}

// @response
// GetWarehouseMaterialByIDVO 入库物资响应（嵌套，继承 Rank）
type GetWarehouseMaterialByIDVO struct {
	Rank
	MaterialName string    `json:"material_name"` // 物资名称
	WarehouseNum int       `json:"warehouse_num"` // 入库数量
	ProductDate  time.Time `json:"product_date"`  // 生产日期
	ExpireDate   time.Time `json:"expire_date"`   // 有效期
}
