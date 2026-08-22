package dto

import "time"

// ============ OutwardController 请求 ============

// PageOutwardByKeyReq 分页查询外出登记请求
// @request
type PageOutwardByKeyReq struct {
	PageNum       *int       `json:"page_num" valid:"required"`  // 页码
	PageSize      *int       `json:"page_size" valid:"required"` // 条数
	ElderName     *string    `json:"elder_name"`                 // 老人姓名
	ChaperoneType *string    `json:"chaperone_type"`             // 陪同人类型
	StartTime     *time.Time `json:"start_time"`                 // 开始时间
	EndTime       *time.Time `json:"end_time"`                   // 结束时间
}

// AddOutwardReq 新增外出登记请求
// @request
type AddOutwardReq struct {
	ElderID        *int64     `json:"elder_id" valid:"required"`         // 老人编号
	ChaperoneName  *string    `json:"chaperone_name" valid:"required"`   // 陪同人姓名
	ChaperonePhone *string    `json:"chaperone_phone" valid:"required"`  // 陪同人电话
	ChaperoneType  *string    `json:"chaperone_type" valid:"required"`   // 陪同人类型
	OutwardDate    *time.Time `json:"outward_date" valid:"required"`     // 外出时间
	PlanReturnDate *time.Time `json:"plan_return_date" valid:"required"` // 计划返回时间
}

// DelayReturnReq 延期返回请求
// @request
type DelayReturnReq struct {
	ID             *int64     `json:"id" valid:"required"`               // id
	PlanReturnDate *time.Time `json:"plan_return_date" valid:"required"` // 计划返回时间
}

// RecordReturnReq 登记返回请求
// @request
type RecordReturnReq struct {
	ID             *int64     `json:"id" valid:"required"`               // id
	RealReturnDate *time.Time `json:"real_return_date" valid:"required"` // 实际返回时间
}

// EditOutwardReq 编辑外出登记请求
// @request
type EditOutwardReq struct {
	ID             *int64     `json:"id"`                                // id
	ElderID        *int64     `json:"elder_id" valid:"required"`         // 老人编号
	ChaperoneName  *string    `json:"chaperone_name" valid:"required"`   // 陪同人姓名
	ChaperonePhone *string    `json:"chaperone_phone" valid:"required"`  // 陪同人电话
	ChaperoneType  *string    `json:"chaperone_type" valid:"required"`   // 陪同人类型
	OutwardDate    *time.Time `json:"outward_date" valid:"required"`     // 外出时间
	PlanReturnDate *time.Time `json:"plan_return_date" valid:"required"` // 计划返回时间
}

// ============ OutwardController 响应 ============

// PageOutwardByKeyResp 分页查询外出登记响应
// @response
type PageOutwardByKeyResp struct {
	ID             int64     `json:"id"`               // 外出登记编号
	ElderName      string    `json:"elder_name"`       // 老人姓名
	ChaperoneName  string    `json:"chaperone_name"`   // 陪同人姓名
	ChaperonePhone string    `json:"chaperone_phone"`  // 陪同人电话
	ChaperoneType  string    `json:"chaperone_type"`   // 陪同人类型
	OutwardDate    time.Time `json:"outward_date"`     // 外出时间
	PlanReturnDate time.Time `json:"plan_return_date"` // 计划返回时间
	RealReturnDate time.Time `json:"real_return_date"` // 实际返回时间
}

// GetOutwardByIDResp 根据编号查询外出登记响应
// @response
type GetOutwardByIDResp struct {
	ID             int64     `json:"id"`               // id
	ElderName      string    `json:"elder_name"`       // 老人姓名
	ChaperoneName  string    `json:"chaperone_name"`   // 陪同人姓名
	ChaperonePhone string    `json:"chaperone_phone"`  // 陪同人电话
	ChaperoneType  string    `json:"chaperone_type"`   // 陪同人类型
	OutwardDate    time.Time `json:"outward_date"`     // 外出时间
	PlanReturnDate time.Time `json:"plan_return_date"` // 计划返回时间
}
