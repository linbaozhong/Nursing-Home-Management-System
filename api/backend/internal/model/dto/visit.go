package dto

import "time"

// ============ VisitController 请求 ============

// @request
// PageVisitByKeyQuery 分页查询来访登记请求
type PageVisitByKeyQuery struct {
	PageNum    *int    `json:"pageNum" valid:"required"` // 页码
	PageSize   *int    `json:"pageSize" valid:"required"` // 条数
	ElderName  *string `json:"elderName"` // 老人姓名
	VisitName  *string `json:"visitName"` // 来访者姓名
	VisitPhone *string `json:"visitPhone"` // 来访者电话
	VisitFlag  *string `json:"visitFlag"` // 来访状态
}

// @request
// EditVisitQuery 编辑来访登记请求
type EditVisitQuery struct {
	ID        *int64  `json:"id"` // id
	Name      *string `json:"name" valid:"required"` // 来访者姓名
	Phone     *string `json:"phone" valid:"required"` // 来访者电话
	Relation  *string `json:"relation" valid:"required"` // 与老人关系
	VisitDate *string `json:"visitDate" valid:"required"` // 来访时间
	VisitNum  *int64  `json:"visitNum" valid:"required"` // 来访者人数
}

// @request
// AddVisitQuery 新增来访登记请求（继承 EditVisitQuery）
type AddVisitQuery struct {
	EditVisitQuery
	ElderID *int64 `json:"elderId" valid:"required"` // 老人编号
}

// @request
// RecordLeaveQuery 登记离开请求
type RecordLeaveQuery struct {
	ID        *int64  `json:"id" valid:"required"` // id
	LeaveDate *string `json:"leaveDate" valid:"required"` // 离开时间
}

// PageSearchElderByKeyQuery 分页搜索老人请求（定义见 elderrecord.go）

// ============ VisitController 响应 ============

// @response
// PageVisitByKeyVO 分页查询来访登记响应
type PageVisitByKeyVO struct {
	Rank
	ID         int64     `json:"id"` // id
	ElderName  string    `json:"elderName"` // 老人姓名
	VisitName  string    `json:"visitName"` // 来访者姓名
	VisitPhone string    `json:"visitPhone"` // 来访者电话
	Relation   string    `json:"relation"` // 与老人关系
	VisitDate  time.Time `json:"visitDate"` // 来访时间
	LeaveDate  time.Time `json:"leaveDate"` // 离开时间
	VisitNum   int64     `json:"visitNum"` // 来访者人数
	VisitFlag  string    `json:"visitFlag"` // 来访状态
}

// @response
// GetVisitByIDVO 根据编号获取来访登记响应
type GetVisitByIDVO struct {
	ID         int64     `json:"id"` // id
	ElderName  string    `json:"elderName"` // 老人姓名
	VisitName  string    `json:"visitName"` // 来访者姓名
	VisitPhone string    `json:"visitPhone"` // 来访者电话
	Relation   string    `json:"relation"` // 与老人关系
	VisitDate  time.Time `json:"visitDate"` // 来访时间
	VisitNum   int64     `json:"visitNum"` // 来访者人数
}