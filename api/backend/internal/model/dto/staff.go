package dto

// ============ StaffController 请求 ============

// @request
// PageStaffByKeyQuery 分页查询员工请求
type PageStaffByKeyQuery struct {
	PageNum  *int    `json:"pageNum" valid:"required"` // 页码
	PageSize *int    `json:"pageSize" valid:"required"` // 条数
	Name     *string `json:"name"` // 员工姓名
	Phone    *string `json:"phone"` // 员工电话
	RoleID   *int64  `json:"roleId"` // 角色编号
}

// @request
// PageSearchStaffByKeyQuery 分页搜索员工请求
type PageSearchStaffByKeyQuery struct {
	PageNum *int    `json:"pageNum" valid:"required"` // 页码
	PageSize *int   `json:"pageSize" valid:"required"` // 条数
	Name    *string `json:"name"` // 姓名
	Phone   *string `json:"phone"` // 电话
}

// @request
// OperateStaffQuery 操作员工请求
type OperateStaffQuery struct {
	ID      *int64  `json:"id"` // id
	RoleID  *int64  `json:"roleId" valid:"required"` // 角色编号
	Name    *string `json:"name" valid:"required"` // 员工姓名
	IDNum   *string `json:"idNum" valid:"required"` // 员工身份证号
	Age     *int    `json:"age" valid:"required"` // 员工年龄
	Sex     *string `json:"sex" valid:"required"` // 员工性别
	Phone   *string `json:"phone" valid:"required"` // 员工电话
	Email   *string `json:"email" valid:"required"` // 员工邮箱
	Address *string `json:"address" valid:"required"` // 员工地址
	Avator  *string `json:"avator" valid:"required"` // 员工头像
}

// ============ StaffController 响应 ============

// @response
// PageStaffByKeyVO 分页查询员工响应
type PageStaffByKeyVO struct {
	Rank
	ID       int64  `json:"id"` // id
	Name     string `json:"name"` // 员工姓名
	RoleName string `json:"roleName"` // 角色名称
	IDNum    string `json:"idNum"` // 身份证号
	Sex      string `json:"sex"` // 员工性别
	Phone    string `json:"phone"` // 员工电话
	Email    string `json:"email"` // 员工邮箱
	LeaveFlag string `json:"leaveFlag"` // 离职状态
}

// @response
// OperateStaffVO 操作员工响应（继承 OperateStaffQuery）
type OperateStaffVO struct {
	OperateStaffQuery
}