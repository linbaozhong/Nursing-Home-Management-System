package dto

import "time"

// ============ OutwardController 请求 ============

// @request
// PageOutwardByKeyQuery 分页查询外出登记请求
type PageOutwardByKeyQuery struct {
	PageNum       *int    `json:"pageNum" valid:"required"` // 页码
	PageSize      *int    `json:"pageSize" valid:"required"` // 条数
	ElderName     *string `json:"elderName"` // 老人姓名
	ChaperoneType *string `json:"chaperoneType"` // 陪同人类型
	StartTime     *string `json:"startTime"` // 开始时间
	EndTime       *string `json:"endTime"` // 结束时间
}

// @request
// AddOutwardQuery 新增外出登记请求
type AddOutwardQuery struct {
	ElderID        *int64  `json:"elderId" valid:"required"` // 老人编号
	ChaperoneName  *string `json:"chaperoneName" valid:"required"` // 陪同人姓名
	ChaperonePhone *string `json:"chaperonePhone" valid:"required"` // 陪同人电话
	ChaperoneType  *string `json:"chaperoneType" valid:"required"` // 陪同人类型
	OutwardDate    *string `json:"outwardDate" valid:"required"` // 外出时间
	PlanReturnDate *string `json:"planReturnDate" valid:"required"` // 计划返回时间
}

// @request
// DelayReturnQuery 延期返回请求
type DelayReturnQuery struct {
	ID            *int64  `json:"id" valid:"required"` // id
	PlanReturnDate *string `json:"planReturnDate" valid:"required"` // 计划返回时间
}

// @request
// RecordReturnQuery 登记返回请求
type RecordReturnQuery struct {
	ID            *int64  `json:"id" valid:"required"` // id
	RealReturnDate *string `json:"realReturnDate" valid:"required"` // 实际返回时间
}

// @request
// EditOutwardQuery 编辑外出登记请求
type EditOutwardQuery struct {
	ID             *int64  `json:"id"` // id
	ElderID        *int64  `json:"elderId" valid:"required"` // 老人编号
	ChaperoneName  *string `json:"chaperoneName" valid:"required"` // 陪同人姓名
	ChaperonePhone *string `json:"chaperonePhone" valid:"required"` // 陪同人电话
	ChaperoneType  *string `json:"chaperoneType" valid:"required"` // 陪同人类型
	OutwardDate    *string `json:"outwardDate" valid:"required"` // 外出时间
	PlanReturnDate *string `json:"planReturnDate" valid:"required"` // 计划返回时间
}

// ============ OutwardController 响应 ============

// @response
// PageOutwardByKeyVO 分页查询外出登记响应
type PageOutwardByKeyVO struct {
	Rank
	ID             int64     `json:"id"` // 外出登记编号
	ElderName      string    `json:"elderName"` // 老人姓名
	ChaperoneName  string    `json:"chaperoneName"` // 陪同人姓名
	ChaperonePhone string    `json:"chaperonePhone"` // 陪同人电话
	ChaperoneType  string    `json:"chaperoneType"` // 陪同人类型
	OutwardDate    time.Time `json:"outwardDate"` // 外出时间
	PlanReturnDate time.Time `json:"planReturnDate"` // 计划返回时间
	RealReturnDate time.Time `json:"realReturnDate"` // 实际返回时间
}

// @response
// GetOutwardByIDVO 根据编号查询外出登记响应
type GetOutwardByIDVO struct {
	ID             int64     `json:"id"` // id
	ElderName      string    `json:"elderName"` // 老人姓名
	ChaperoneName  string    `json:"chaperoneName"` // 陪同人姓名
	ChaperonePhone string    `json:"chaperonePhone"` // 陪同人电话
	ChaperoneType  string    `json:"chaperoneType"` // 陪同人类型
	OutwardDate    time.Time `json:"outwardDate"` // 外出时间
	PlanReturnDate time.Time `json:"planReturnDate"` // 计划返回时间
}