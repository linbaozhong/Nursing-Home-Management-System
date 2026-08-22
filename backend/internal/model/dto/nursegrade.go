package dto

import "github.com/linbaozhong/gentity/pkg/types"

// ============ NurseGradeController 请求 ============

// PageNurseGradeByKeyReq 分页查询护理等级请求
// @request
type PageNurseGradeByKeyReq struct {
	PageNum   *int    `json:"page_num" valid:"required"`  // 页码
	PageSize  *int    `json:"page_size" valid:"required"` // 条数
	GradeName *string `json:"grade_name"`                 // 级别名称
	NurseType *string `json:"nurse_type"`                 // 护理类型
}

// OperateNurseGradeReq 操作护理等级请求（定义见 common.go）

// ============ NurseGradeController 响应 ============

// PageNurseGradeByKeyResp 分页查询护理等级响应
// @response
type PageNurseGradeByKeyResp struct {
	ID         int64       `json:"id"`          // id
	Name       string      `json:"name"`        // 护理等级名称
	Type       string      `json:"type"`        // 护理类型
	MonthPrice types.Money `json:"month_price"` // 月护理费用
}

// GetNurseGradeByIDResp 护理等级详情响应（定义见 common.go）

// AddNurseGradeReq 新增护理等级请求
// @request
type AddNurseGradeReq struct {
	ID         *int64       `json:"id"`                           // id
	Name       *string      `json:"name" valid:"required"`        // 护理等级名称
	Type       *string      `json:"type" valid:"required"`        // 护理类型
	MonthPrice *types.Money `json:"month_price" valid:"required"` // 月护理费用
}

// EditNurseGradeReq 编辑护理等级请求
// @request
type EditNurseGradeReq struct {
	ID         *int64       `json:"id"`                           // id
	Name       *string      `json:"name" valid:"required"`        // 护理等级名称
	Type       *string      `json:"type" valid:"required"`        // 护理类型
	MonthPrice *types.Money `json:"month_price" valid:"required"` // 月护理费用
}

// PageNurseByKeyReq 分页查询护理员请求
// @request
type PageNurseByKeyReq struct {
	PageNum   *int    `json:"page_num" valid:"required"`  // 页码
	PageSize  *int    `json:"page_size" valid:"required"` // 条数
	NurseName *string `json:"nurse_name"`                 // 护理员姓名
	Key       *string `json:"key"`                        // 关键字
}

// PageNurseByKeyResp 分页查询护理员响应
// @response
type PageNurseByKeyResp struct {
	ID    int64  `json:"id"`    // id
	Name  string `json:"name"`  // 护理员姓名
	Phone string `json:"phone"` // 电话
}

// GetNurseByIdResp 护理员详情响应
// @response
type GetNurseByIdResp struct {
	ID    int64  `json:"id"`    // id
	Name  string `json:"name"`  // 护理员姓名
	Phone string `json:"phone"` // 电话
}

// AddNurseReq 新增护理员请求
// @request
type AddNurseReq struct {
	ID        *int64  `json:"id"`                          // id
	NurseName *string `json:"nurse_name" valid:"required"` // 护理员姓名
	Phone     *string `json:"phone" valid:"required"`      // 电话
	GradeID   *int64  `json:"grade_id" valid:"required"`   // 护理等级编号
}

// EditNurseReq 编辑护理员请求
// @request
type EditNurseReq struct {
	ID        *int64  `json:"id"`                          // id
	NurseName *string `json:"nurse_name" valid:"required"` // 护理员姓名
	Phone     *string `json:"phone" valid:"required"`      // 电话
	GradeID   *int64  `json:"grade_id" valid:"required"`   // 护理等级编号
}
