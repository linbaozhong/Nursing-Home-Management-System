package dto

import "time"

// ============ VisitController 请求 ============

// PageVisitByKeyReq 分页查询来访登记请求
// @request
type PageVisitByKeyReq struct {
	PageNum    *int    `json:"page_num" valid:"required"`  // 页码
	PageSize   *int    `json:"page_size" valid:"required"` // 条数
	ElderName  *string `json:"elder_name"`                 // 老人姓名
	VisitName  *string `json:"visit_name"`                 // 来访者姓名
	VisitPhone *string `json:"visit_phone"`                // 来访者电话
	VisitFlag  *string `json:"visit_flag"`                 // 来访状态
}

// EditVisitReq 编辑来访登记请求
// @request
type EditVisitReq struct {
	ID        *int64     `json:"id"`                          // id
	Name      *string    `json:"name" valid:"required"`       // 来访者姓名
	Phone     *string    `json:"phone" valid:"required"`      // 来访者电话
	Relation  *string    `json:"relation" valid:"required"`   // 与老人关系
	VisitDate *time.Time `json:"visit_date" valid:"required"` // 来访时间
	VisitNum  *int64     `json:"visit_num" valid:"required"`  // 来访者人数
}

// AddVisitReq 新增来访登记请求（继承 EditVisitReq）
// @request
type AddVisitReq struct {
	EditVisitReq
	ElderID *int64 `json:"elder_id" valid:"required"` // 老人编号
}

// RecordLeaveReq 登记离开请求
// @request
type RecordLeaveReq struct {
	ID        *int64     `json:"id" valid:"required"`         // id
	LeaveDate *time.Time `json:"leave_date" valid:"required"` // 离开时间
}

// PageSearchElderByKeyReq 分页搜索老人请求（定义见 elderrecord.go）

// ============ VisitController 响应 ============

// PageVisitByKeyResp 分页查询来访登记响应
// @response
type PageVisitByKeyResp struct {
	ID         int64     `json:"id"`          // id
	ElderName  string    `json:"elder_name"`  // 老人姓名
	VisitName  string    `json:"visit_name"`  // 来访者姓名
	VisitPhone string    `json:"visit_phone"` // 来访者电话
	Relation   string    `json:"relation"`    // 与老人关系
	VisitDate  time.Time `json:"visit_date"`  // 来访时间
	LeaveDate  time.Time `json:"leave_date"`  // 离开时间
	VisitNum   int64     `json:"visit_num"`   // 来访者人数
	VisitFlag  *int8     `json:"visit_flag"`  // 来访状态（0-在院 1-已离开）
}

// GetVisitByIDResp 根据编号获取来访登记响应
// @response
type GetVisitByIDResp struct {
	ID         int64     `json:"id"`          // id
	ElderName  string    `json:"elder_name"`  // 老人姓名
	VisitName  string    `json:"visit_name"`  // 来访者姓名
	VisitPhone string    `json:"visit_phone"` // 来访者电话
	Relation   string    `json:"relation"`    // 与老人关系
	VisitDate  time.Time `json:"visit_date"`  // 来访时间
	VisitNum   int64     `json:"visit_num"`   // 来访者人数
}
