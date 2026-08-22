package dto

import "time"

// ============ AccidentController 请求 ============

// PageAccidentByKeyReq 分页查询事故登记请求
// @request
type PageAccidentByKeyReq struct {
	PageNum   *int    `json:"page_num" valid:"required"`  // 页码
	PageSize  *int    `json:"page_size" valid:"required"` // 条数
	ElderName *string `json:"elder_name"`                 // 老人姓名
	StaffName *string `json:"staff_name"`                 // 护工姓名
}

// EditAccidentReq 编辑事故登记请求
// @request
type EditAccidentReq struct {
	ID *int64 `json:"id" valid:"required"` // id
	AccidentQuery
}

// AddAccidentReq 新增事故登记请求（继承 EditAccidentReq）
// @request
type AddAccidentReq struct {
	AccidentQuery
	ElderID *int64 `json:"elder_id" valid:"required"` // 老人编号
}

type AccidentQuery struct {
	StaffID     *int64     `json:"staff_id" valid:"required"`    // 护工编号
	OccurDate   *time.Time `json:"occur_date" valid:"required"`  // 发生时间
	Description *string    `json:"description" valid:"required"` // 事故描述
	Picture     *string    `json:"picture" valid:"required"`     // 事故图片
}

// ============ AccidentController 响应 ============

// PageAccidentByKeyResp 分页查询事故登记响应
// @response
type PageAccidentByKeyResp struct {
	ID        int64     `json:"id"`         // id
	ElderName string    `json:"elder_name"` // 老人姓名
	StaffName string    `json:"staff_name"` // 护工姓名
	OccurDate time.Time `json:"occur_date"` // 发生时间
}

// GetAccidentByIDResp 根据编号获取事故登记响应
// @response
type GetAccidentByIDResp struct {
	ID          int64     `json:"id"`          // id
	ElderName   string    `json:"elder_name"`  // 老人姓名
	StaffID     int64     `json:"staff_id"`    // 护工编号
	OccurDate   time.Time `json:"occur_date"`  // 发生时间
	Description string    `json:"description"` // 事故描述
	Picture     string    `json:"picture"`     // 事故图片
}
