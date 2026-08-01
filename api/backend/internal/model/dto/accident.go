package dto

import "time"

// ============ AccidentController 请求 ============

// @request
// PageAccidentByKeyQuery 分页查询事故登记请求
type PageAccidentByKeyQuery struct {
	PageNum   *int    `json:"pageNum" valid:"required"` // 页码
	PageSize  *int    `json:"pageSize" valid:"required"` // 条数
	ElderName *string `json:"elderName"` // 老人姓名
	StaffName *string `json:"staffName"` // 护工姓名
}

// @request
// EditAccidentQuery 编辑事故登记请求
type EditAccidentQuery struct {
	ID          *int64  `json:"id"` // id
	StaffID     *int64  `json:"staffId" valid:"required"` // 护工编号
	OccurDate   *string `json:"occurDate" valid:"required"` // 发生时间
	Description *string `json:"description" valid:"required"` // 事故描述
	Picture     *string `json:"picture" valid:"required"` // 事故图片
}

// @request
// AddAccidentQuery 新增事故登记请求（继承 EditAccidentQuery）
type AddAccidentQuery struct {
	EditAccidentQuery
	ElderID *int64 `json:"elderId" valid:"required"` // 老人编号
}

// ============ AccidentController 响应 ============

// @response
// PageAccidentByKeyVO 分页查询事故登记响应
type PageAccidentByKeyVO struct {
	Rank
	ID        int64     `json:"id"` // id
	ElderName string    `json:"elderName"` // 老人姓名
	StaffName string    `json:"staffName"` // 护工姓名
	OccurDate time.Time `json:"occurDate"` // 发生时间
}

// @response
// GetAccidentByIDVO 根据编号获取事故登记响应
type GetAccidentByIDVO struct {
	ID          int64     `json:"id"` // id
	ElderName   string    `json:"elderName"` // 老人姓名
	StaffID     int64     `json:"staffId"` // 护工编号
	OccurDate   time.Time `json:"occurDate"` // 发生时间
	Description string    `json:"description"` // 事故描述
	Picture     string    `json:"picture"` // 事故图片
}