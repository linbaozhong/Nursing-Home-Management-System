package dto

import "time"

// ============ OutboundRecordController 请求 ============

// PageOutboundRecordByKeyReq 分页查询出库记录请求
// @request
type PageOutboundRecordByKeyReq struct {
	PageNum       *int       `json:"page_num" valid:"required"`  // 页码
	PageSize      *int       `json:"page_size" valid:"required"` // 条数
	WarehouseName *string    `json:"warehouse_name"`             // 仓库名称
	MaterialName  *string    `json:"material_name"`              // 物资名称
	StartTime     *time.Time `json:"start_time"`                 // 开始时间
	EndTime       *time.Time `json:"end_time"`                   // 结束时间
	Recipient     *string    `json:"recipient"`                  // 领用人
}

// AddOutboundRecordReq 新增出库记录请求
// @request
type AddOutboundRecordReq struct {
	RecipientType             *string                  `json:"recipient_type" valid:"required"`               // 领用人类型
	RecipientID               *int64                   `json:"recipient_id" valid:"required"`                 // 领用人编号
	WarehouseID               *int64                   `json:"warehouse_id" valid:"required"`                 // 仓库编号
	OutboundDate              *time.Time               `json:"outbound_date" valid:"required"`                // 出库时间
	MaterialUse               *string                  `json:"material_use" valid:"required"`                 // 物资去向
	StaffID                   *int64                   `json:"staff_id" valid:"required"`                     // 登记人编号
	OutboundMaterialQueryList []AddOutboundMaterialReq `json:"outbound_material_query_list" valid:"required"` // 出库物资列表
}

// AddOutboundMaterialReq 新增出库物资请求（嵌套）
// @request
type AddOutboundMaterialReq struct {
	WarehouseMaterialID *int64 `json:"warehouse_material_id" valid:"required"` // 入库物资编号
	OutboundNum         *int   `json:"outbound_num" valid:"required"`          // 出库数量
}

// AuditOutboundRecordReq 审核出库记录请求
// @request
type AuditOutboundRecordReq struct {
	OutboundRecordID *int64  `json:"outbound_record_id" valid:"required"` // 出库记录编号
	AuditResult      *string `json:"audit_result" valid:"required"`       // 审核结果
}

// ============ OutboundRecordController 响应 ============

// PageOutboundRecordByKeyResp 分页查询出库记录响应
// @response
type PageOutboundRecordByKeyResp struct {
	ID            int64     `json:"id"`             // id
	WarehouseName string    `json:"warehouse_name"` // 仓库名称
	MaterialName  string    `json:"material_name"`  // 物资名称
	OutboundDate  time.Time `json:"outbound_date"`  // 出库时间
	MaterialUse   string    `json:"material_use"`   // 物资去向
	Recipient     string    `json:"recipient"`      // 领用人
	StaffName     string    `json:"staff_name"`     // 登记人
	OutboundFlag  string    `json:"outbound_flag"`  // 出库状态
}

// GetOutboundRecordByIDResp 根据编号查询出库记录响应
// @response
type GetOutboundRecordByIDResp struct {
	ID                           int64                         `json:"id"`                              // id
	RecipientType                string                        `json:"recipient_type"`                  // 领用人类型
	Recipient                    string                        `json:"recipient"`                       // 领用人
	WarehouseName                string                        `json:"warehouse_name"`                  // 仓库名称
	OutboundDate                 time.Time                     `json:"outbound_date"`                   // 出库时间
	MaterialUse                  string                        `json:"material_use"`                    // 物资去向
	StaffName                    string                        `json:"staff_name"`                      // 登记人
	OutboundMaterialByIDRespList []GetOutboundMaterialByIDResp `json:"outbound_material_by_id_vo_list"` // 出库物资列表
}

// GetOutboundMaterialByIDResp 出库物资响应（嵌套，继承 Rank）
// @response
type GetOutboundMaterialByIDResp struct {
	MaterialName string `json:"material_name"` // 物资名称
	OutboundNum  int    `json:"outbound_num"`  // 出库数量
}
