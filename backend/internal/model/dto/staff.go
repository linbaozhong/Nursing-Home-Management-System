package dto

import "github.com/linbaozhong/gentity/pkg/types"

// ============ StaffController 请求 ============

// PageStaffByKeyReq 分页查询员工请求
// @request
type PageStaffByKeyReq struct {
	PageNum  *int    `json:"page_num" valid:"required"`  // 页码
	PageSize *int    `json:"page_size" valid:"required"` // 条数
	Name     *string `json:"name"`                       // 员工姓名
	Phone    *string `json:"phone"`                      // 员工电话
	RoleID   *int64  `json:"role_id"`                    // 角色编号
}

// PageSearchStaffByKeyReq 分页搜索员工请求
// @request
type PageSearchStaffByKeyReq struct {
	PageNum  *int    `json:"page_num" valid:"required"`  // 页码
	PageSize *int    `json:"page_size" valid:"required"` // 条数
	Name     *string `json:"name"`                       // 姓名
	Phone    *string `json:"phone"`                      // 电话
}

// OperateStaffReq 操作员工请求
// @request
type OperateStaffReq struct {
	ID      *int64  `json:"id"`                       // id
	RoleID  *int64  `json:"role_id" valid:"required"` // 角色编号
	Name    *string `json:"name" valid:"required"`    // 员工姓名
	IDNum   *string `json:"id_num" valid:"required"`  // 员工身份证号
	Age     *int    `json:"age" valid:"required"`     // 员工年龄
	Sex     *string `json:"sex" valid:"required"`     // 员工性别
	Phone   *string `json:"phone" valid:"required"`   // 员工电话
	Email   *string `json:"email" valid:"required"`   // 员工邮箱
	Address *string `json:"address" valid:"required"` // 员工地址
	Avator  *string `json:"avator" valid:"required"`  // 员工头像
}

// ============ StaffController 响应 ============

// PageStaffByKeyResp 分页查询员工响应
// @response
type PageStaffByKeyResp struct {
	ID        types.BigInt `json:"id"`         // id
	Name      string       `json:"name"`       // 员工姓名
	RoleName  string       `json:"role_name"`  // 角色名称
	IDNum     string       `json:"id_num"`     // 身份证号
	Sex       string       `json:"sex"`        // 员工性别
	Phone     string       `json:"phone"`      // 员工电话
	Email     string       `json:"email"`      // 员工邮箱
	LeaveFlag string       `json:"leave_flag"` // 离职状态
}

// OperateStaffResp 操作员工响应（继承 OperateStaffReq）
// @response
type OperateStaffResp struct {
	OperateStaffReq
}

// PageSearchStaffByKeyResp 分页搜索员工响应
// @response
type PageSearchStaffByKeyResp struct {
	ID    types.BigInt `json:"id"`    // id
	Name  string       `json:"name"`  // 员工姓名
	Phone string       `json:"phone"` // 员工电话
}
