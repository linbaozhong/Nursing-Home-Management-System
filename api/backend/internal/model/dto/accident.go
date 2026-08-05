package dto

import "time"

// ============ AccidentController 请求 ============

// @request
// PageAccidentByKeyQuery 分页查询事故登记请求
type PageAccidentByKeyQuery struct {
	PageNum   *int    `json:"page_num" valid:"required"`  // 页码
	PageSize  *int    `json:"page_size" valid:"required"` // 条数
	ElderName *string `json:"elder_name"`                 // 老人姓名
	StaffName *string `json:"staff_name"`                 // 护工姓名
}

// @request
// EditAccidentQuery 编辑事故登记请求
type EditAccidentQuery struct {
	ID *int64 `json:"id" valid:"required"` // id
	AccidentQuery
}

// @request
// AddAccidentQuery 新增事故登记请求（继承 EditAccidentQuery）
type AddAccidentQuery struct {
	AccidentQuery
	ElderID *int64 `json:"elder_id" valid:"required"` // 老人编号
}

type AccidentQuery struct {
	StaffID     *int64  `json:"staff_id" valid:"required"`    // 护工编号
	OccurDate   *string `json:"occur_date" valid:"required"`  // 发生时间
	Description *string `json:"description" valid:"required"` // 事故描述
	Picture     *string `json:"picture" valid:"required"`     // 事故图片
}

// ============ AccidentController 响应 ============

// @response
// PageAccidentByKeyVO 分页查询事故登记响应
type PageAccidentByKeyVO struct {
	// Rank
	ID        int64     `json:"id"`         // id
	ElderName string    `json:"elder_name"` // 老人姓名
	StaffName string    `json:"staff_name"` // 护工姓名
	OccurDate time.Time `json:"occur_date"` // 发生时间
}

// @response
// GetAccidentByIDVO 根据编号获取事故登记响应
type GetAccidentByIDVO struct {
	ID          int64     `json:"id"`          // id
	ElderName   string    `json:"elder_name"`  // 老人姓名
	StaffID     int64     `json:"staff_id"`    // 护工编号
	OccurDate   time.Time `json:"occur_date"`  // 发生时间
	Description string    `json:"description"` // 事故描述
	Picture     string    `json:"picture"`     // 事故图片
}
