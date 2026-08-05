package dto

// ============ NurseGradeController 请求 ============

// @request
// PageNurseGradeByKeyQuery 分页查询护理等级请求
type PageNurseGradeByKeyQuery struct {
	PageNum   *int    `json:"page_num" valid:"required"`  // 页码
	PageSize  *int    `json:"page_size" valid:"required"` // 条数
	GradeName *string `json:"grade_name"`                 // 级别名称
	NurseType *string `json:"nurse_type"`                 // 护理类型
}

// OperateNurseGradeQuery 操作护理等级请求（定义见 common.go）

// ============ NurseGradeController 响应 ============

// @response
// PageNurseGradeByKeyVO 分页查询护理等级响应
type PageNurseGradeByKeyVO struct {
	Rank
	ID         int64   `json:"id"`          // id
	Name       string  `json:"name"`        // 护理等级名称
	Type       string  `json:"type"`        // 护理类型
	MonthPrice float64 `json:"month_price"` // 月护理费用
}

// GetNurseGradeByIDVO 护理等级详情响应（定义见 common.go）

// @request
// AddNurseGradeQuery 新增护理等级请求
type AddNurseGradeQuery struct {
	ID         *int64   `json:"id"`                           // id
	Name       *string  `json:"name" valid:"required"`        // 护理等级名称
	Type       *string  `json:"type" valid:"required"`        // 护理类型
	MonthPrice *float64 `json:"month_price" valid:"required"` // 月护理费用
}

// @request
// EditNurseGradeQuery 编辑护理等级请求
type EditNurseGradeQuery struct {
	ID         *int64   `json:"id"`                           // id
	Name       *string  `json:"name" valid:"required"`        // 护理等级名称
	Type       *string  `json:"type" valid:"required"`        // 护理类型
	MonthPrice *float64 `json:"month_price" valid:"required"` // 月护理费用
}

// @request
// PageNurseByKeyQuery 分页查询护理员请求
type PageNurseByKeyQuery struct {
	PageNum   *int    `json:"page_num" valid:"required"`  // 页码
	PageSize  *int    `json:"page_size" valid:"required"` // 条数
	NurseName *string `json:"nurse_name"`                 // 护理员姓名
	Key       *string `json:"key"`                        // 关键字
}

// @request
// AddNurseQuery 新增护理员请求
type AddNurseQuery struct {
	ID        *int64  `json:"id"`                          // id
	NurseName *string `json:"nurse_name" valid:"required"` // 护理员姓名
	Phone     *string `json:"phone" valid:"required"`      // 电话
	GradeID   *int64  `json:"grade_id" valid:"required"`   // 护理等级编号
}

// @request
// EditNurseQuery 编辑护理员请求
type EditNurseQuery struct {
	ID        *int64  `json:"id"`                          // id
	NurseName *string `json:"nurse_name" valid:"required"` // 护理员姓名
	Phone     *string `json:"phone" valid:"required"`      // 电话
	GradeID   *int64  `json:"grade_id" valid:"required"`   // 护理等级编号
}
