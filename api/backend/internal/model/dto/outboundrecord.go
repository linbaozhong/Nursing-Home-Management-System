package dto

import "time"

// ============ OutboundRecordController 请求 ============

// @request
// PageOutboundRecordByKeyQuery 分页查询出库记录请求
type PageOutboundRecordByKeyQuery struct {
	PageNum       *int    `json:"pageNum" valid:"required"` // 页码
	PageSize      *int    `json:"pageSize" valid:"required"` // 条数
	WarehouseName *string `json:"warehouseName"` // 仓库名称
	MaterialName  *string `json:"materialName"` // 物资名称
	StartTime     *string `json:"startTime"` // 开始时间
	EndTime       *string `json:"endTime"` // 结束时间
	Recipient     *string `json:"recipient"` // 领用人
}

// @request
// AddOutboundRecordQuery 新增出库记录请求
type AddOutboundRecordQuery struct {
	RecipientType              *string                     `json:"recipientType" valid:"required"` // 领用人类型
	RecipientID                *int64                      `json:"recipientId" valid:"required"` // 领用人编号
	WarehouseID                *int64                      `json:"warehouseId" valid:"required"` // 仓库编号
	OutboundDate               *string                     `json:"outboundDate" valid:"required"` // 出库时间
	MaterialUse                *string                     `json:"materialUse" valid:"required"` // 物资去向
	StaffID                    *int64                      `json:"staffId" valid:"required"` // 登记人编号
	OutboundMaterialQueryList  []AddOutboundMaterialQuery `json:"outboundMaterialQueryList" valid:"required"` // 出库物资列表
}

// @request
// AddOutboundMaterialQuery 新增出库物资请求（嵌套）
type AddOutboundMaterialQuery struct {
	WarehouseMaterialID *int64 `json:"warehouseMaterialId" valid:"required"` // 入库物资编号
	OutboundNum         *int   `json:"outboundNum" valid:"required"` // 出库数量
}

// @request
// AuditOutboundRecordQuery 审核出库记录请求
type AuditOutboundRecordQuery struct {
	OutboundRecordID *int64  `json:"outboundRecordId" valid:"required"` // 出库记录编号
	AuditResult      *string `json:"auditResult" valid:"required"` // 审核结果
}

// ============ OutboundRecordController 响应 ============

// @response
// PageOutboundRecordByKeyVO 分页查询出库记录响应
type PageOutboundRecordByKeyVO struct {
	Rank
	ID            int64     `json:"id"` // id
	WarehouseName string    `json:"warehouseName"` // 仓库名称
	MaterialName  string    `json:"materialName"` // 物资名称
	OutboundDate  time.Time `json:"outboundDate"` // 出库时间
	MaterialUse   string    `json:"materialUse"` // 物资去向
	Recipient     string    `json:"recipient"` // 领用人
	StaffName     string    `json:"staffName"` // 登记人
	OutboundFlag  string    `json:"outboundFlag"` // 出库状态
}

// @response
// GetOutboundRecordByIDVO 根据编号查询出库记录响应
type GetOutboundRecordByIDVO struct {
	ID                          int64                       `json:"id"` // id
	RecipientType               string                      `json:"recipientType"` // 领用人类型
	Recipient                   string                      `json:"recipient"` // 领用人
	WarehouseName               string                      `json:"warehouseName"` // 仓库名称
	OutboundDate                time.Time                   `json:"outboundDate"` // 出库时间
	MaterialUse                 string                      `json:"materialUse"` // 物资去向
	StaffName                   string                      `json:"staffName"` // 登记人
	OutboundMaterialByIDVOList  []GetOutboundMaterialByIDVO `json:"outboundMaterialByIdVoList"` // 出库物资列表
}

// @response
// GetOutboundMaterialByIDVO 出库物资响应（嵌套，继承 Rank）
type GetOutboundMaterialByIDVO struct {
	Rank
	MaterialName string `json:"materialName"` // 物资名称
	OutboundNum  int    `json:"outboundNum"` // 出库数量
}