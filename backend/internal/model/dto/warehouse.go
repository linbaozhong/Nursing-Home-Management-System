package dto

// ============ WarehouseController 请求 ============

// PageWarehouseByKeyReq 分页查询仓库请求
// @request
type PageWarehouseByKeyReq struct {
	PageNum       *int    `json:"page_num" valid:"required"`  // 页码
	PageSize      *int    `json:"page_size" valid:"required"` // 条数
	WarehouseName *string `json:"warehouse_name"`             // 仓库名称
}

// OperateWarehouseReq 操作仓库请求
// @request
type OperateWarehouseReq struct {
	ID      *int64  `json:"id"`                        // id
	StaffID *int64  `json:"staff_id" valid:"required"` // 仓库管理员编号
	Name    *string `json:"name" valid:"required"`     // 仓库名称
}

// ============ WarehouseController 响应 ============

// PageWarehouseByKeyResp 分页查询仓库响应
// @response
type PageWarehouseByKeyResp struct {
	ID        int64  `json:"id"`         // id
	Name      string `json:"name"`       // 仓库名称
	StaffName string `json:"staff_name"` // 仓库管理员
}

// OperateWarehouseResp 操作仓库响应（继承 OperateWarehouseReq）
// @response
type OperateWarehouseResp struct {
	OperateWarehouseReq
}
