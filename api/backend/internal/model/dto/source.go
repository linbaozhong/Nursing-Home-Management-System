package dto

// ============ SourceController 请求 ============

// @request
// PageSourceByKeyQuery 分页查询来源渠道请求
type PageSourceByKeyQuery struct {
	PageNum    *int    `json:"page_num" valid:"required"`  // 页码
	PageSize   *int    `json:"page_size" valid:"required"` // 条数
	SourceName *string `json:"source_name"`                // 来源渠道名称
}

// @request
// OperateSourceQuery 操作来源渠道请求
type OperateSourceQuery struct {
	ID   *int64  `json:"id"`                    // id
	Name *string `json:"name" valid:"required"` // 来源渠道名称
}

// ============ SourceController 响应 ============

// @response
// PageSourceByKeyVO 分页查询来源渠道响应
type PageSourceByKeyVO struct {
	Rank
	ID   int64  `json:"id"`   // id
	Name string `json:"name"` // 来源渠道名称
}
