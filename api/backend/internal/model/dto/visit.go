package dto

import "time"

// ============ VisitController 请求 ============

// @request
// PageVisitByKeyQuery 分页查询来访登记请求
type PageVisitByKeyQuery struct {
	PageNum    *int    `json:"page_num" valid:"required"`  // 页码
	PageSize   *int    `json:"page_size" valid:"required"` // 条数
	ElderName  *string `json:"elder_name"`                 // 老人姓名
	VisitName  *string `json:"visit_name"`                 // 来访者姓名
	VisitPhone *string `json:"visit_phone"`                // 来访者电话
	VisitFlag  *string `json:"visit_flag"`                 // 来访状态
}

// @request
// EditVisitQuery 编辑来访登记请求
type EditVisitQuery struct {
	ID        *int64  `json:"id"`                          // id
	Name      *string `json:"name" valid:"required"`       // 来访者姓名
	Phone     *string `json:"phone" valid:"required"`      // 来访者电话
	Relation  *string `json:"relation" valid:"required"`   // 与老人关系
	VisitDate *string `json:"visit_date" valid:"required"` // 来访时间
	VisitNum  *int64  `json:"visit_num" valid:"required"`  // 来访者人数
}

// @request
// AddVisitQuery 新增来访登记请求（继承 EditVisitQuery）
type AddVisitQuery struct {
	EditVisitQuery
	ElderID *int64 `json:"elder_id" valid:"required"` // 老人编号
}

// @request
// RecordLeaveQuery 登记离开请求
type RecordLeaveQuery struct {
	ID        *int64  `json:"id" valid:"required"`         // id
	LeaveDate *string `json:"leave_date" valid:"required"` // 离开时间
}

// PageSearchElderByKeyQuery 分页搜索老人请求（定义见 elderrecord.go）

// ============ VisitController 响应 ============

// @response
// PageVisitByKeyVO 分页查询来访登记响应
type PageVisitByKeyVO struct {
	Rank
	ID         int64     `json:"id"`          // id
	ElderName  string    `json:"elder_name"`  // 老人姓名
	VisitName  string    `json:"visit_name"`  // 来访者姓名
	VisitPhone string    `json:"visit_phone"` // 来访者电话
	Relation   string    `json:"relation"`    // 与老人关系
	VisitDate  time.Time `json:"visit_date"`  // 来访时间
	LeaveDate  time.Time `json:"leave_date"`  // 离开时间
	VisitNum   int64     `json:"visit_num"`   // 来访者人数
	VisitFlag  string    `json:"visit_flag"`  // 来访状态
}

// @response
// GetVisitByIDVO 根据编号获取来访登记响应
type GetVisitByIDVO struct {
	ID         int64     `json:"id"`          // id
	ElderName  string    `json:"elder_name"`  // 老人姓名
	VisitName  string    `json:"visit_name"`  // 来访者姓名
	VisitPhone string    `json:"visit_phone"` // 来访者电话
	Relation   string    `json:"relation"`    // 与老人关系
	VisitDate  time.Time `json:"visit_date"`  // 来访时间
	VisitNum   int64     `json:"visit_num"`   // 来访者人数
}
