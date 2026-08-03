package dto

// ============ ActiveTypeController 请求 ============

// @request
// PageActiveTypeByKeyQuery 分页查询活动分类请求
type PageActiveTypeByKeyQuery struct {
	PageNum        *int    `json:"page_num" valid:"required"`  // 页码
	PageSize       *int    `json:"page_size" valid:"required"` // 条数
	ActiveTypeName *string `json:"active_type_name"`           // 活动分类名称
}

// @request
// OperateActiveTypeQuery 操作活动分类请求
type OperateActiveTypeQuery struct {
	ID   *int64  `json:"id"`                    // id
	Name *string `json:"name" valid:"required"` // 活动分类名称
}

// @request
// AddActiveTypeQuery 新增活动分类请求
type AddActiveTypeQuery struct {
	ID   *int64  `json:"id"`                    // id
	Name *string `json:"name" valid:"required"` // 活动分类名称
}

// ============ ActiveTypeController 响应 ============

// @response
// PageActiveTypeByKeyVO 分页查询活动分类响应
type PageActiveTypeByKeyVO struct {
	Rank
	ID   int64  `json:"id"`   // id
	Name string `json:"name"` // 活动分类名称
}

// @response
// OperateActiveTypeVO 操作活动分类响应（继承 OperateActiveTypeQuery）
type OperateActiveTypeVO struct {
	OperateActiveTypeQuery
}
