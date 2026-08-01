package dto

// ============ DepositRechargeController 请求 ============

// @request
// PageDepositRechargeByKeyQuery 分页查询预存充值请求
type PageDepositRechargeByKeyQuery struct {
	PageNum *int    `json:"pageNum" valid:"required"` // 页码
	PageSize *int   `json:"pageSize" valid:"required"` // 条数
	Name    *string `json:"name"` // 姓名
	Phone   *string `json:"phone"` // 电话
	IDNum   *string `json:"idNum"` // 身份证号
}

// @request
// RechargeQuery 入住老人账户充值请求
type RechargeQuery struct {
	ElderID *int64   `json:"elderId" valid:"required"` // 老人编号
	Amount  *float64 `json:"amount" valid:"required"` // 充值金额
}

// @request
// AddDepositRechargeQuery 新增预存充值请求
type AddDepositRechargeQuery struct {
	ID       *int64   `json:"id"` // id
	ElderID  *int64   `json:"elderId" valid:"required"` // 老人编号
	Amount   *float64 `json:"amount" valid:"required"` // 充值金额
	RechargeDate *string `json:"rechargeDate" valid:"required"` // 充值日期
	Remark   *string  `json:"remark"` // 备注
}

// @request
// EditDepositRechargeQuery 编辑预存充值请求
type EditDepositRechargeQuery struct {
	ID       *int64   `json:"id"` // id
	ElderID  *int64   `json:"elderId" valid:"required"` // 老人编号
	Amount   *float64 `json:"amount" valid:"required"` // 充值金额
	RechargeDate *string `json:"rechargeDate" valid:"required"` // 充值日期
	Remark   *string  `json:"remark"` // 备注
}

// ============ DepositRechargeController 响应 ============

// @response
// PageDepositRechargeByKeyVO 分页查询预存充值响应
type PageDepositRechargeByKeyVO struct {
	Rank
	ElderID    string  `json:"elderId"` // 老人编号
	ElderName  string  `json:"elderName"` // 老人姓名
	ElderPhone string  `json:"elderPhone"` // 老人电话
	IDNum      string  `json:"idNum"` // 身份证号
	BedName    string  `json:"bedName"` // 床位名称
	Balance    float64 `json:"balance"` // 余额
}