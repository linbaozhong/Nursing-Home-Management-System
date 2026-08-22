package dto

import (
	"github.com/linbaozhong/gentity/pkg/types"
	"time"
)

// ============ DepositRechargeController 请求 ============

// PageDepositRechargeByKeyReq 分页查询预存充值请求
// @request
type PageDepositRechargeByKeyReq struct {
	PageNum  *int    `json:"page_num" valid:"required"`  // 页码
	PageSize *int    `json:"page_size" valid:"required"` // 条数
	Name     *string `json:"name"`                       // 姓名
	Phone    *string `json:"phone"`                      // 电话
	IDNum    *string `json:"id_num"`                     // 身份证号
}

// RechargeReq 入住老人账户充值请求
// @request
type RechargeReq struct {
	ElderID *int64       `json:"elder_id" valid:"required"` // 老人编号
	Amount  *types.Money `json:"amount" valid:"required"`   // 充值金额
}

// AddDepositRechargeReq 新增预存充值请求
// @request
type AddDepositRechargeReq struct {
	ID           *int64       `json:"id"`                             // id
	ElderID      *int64       `json:"elder_id" valid:"required"`      // 老人编号
	Amount       *types.Money `json:"amount" valid:"required"`        // 充值金额
	RechargeDate *time.Time   `json:"recharge_date" valid:"required"` // 充值日期
	Remark       *string      `json:"remark"`                         // 备注
}

// EditDepositRechargeReq 编辑预存充值请求
// @request
type EditDepositRechargeReq struct {
	ID           *int64       `json:"id"`                             // id
	ElderID      *int64       `json:"elder_id" valid:"required"`      // 老人编号
	Amount       *types.Money `json:"amount" valid:"required"`        // 充值金额
	RechargeDate *time.Time   `json:"recharge_date" valid:"required"` // 充值日期
	Remark       *string      `json:"remark"`                         // 备注
}

// ============ DepositRechargeController 响应 ============

// PageDepositRechargeByKeyResp 分页查询预存充值响应
// @response
type PageDepositRechargeByKeyResp struct {
	ElderID    string  `json:"elder_id"`    // 老人编号
	ElderName  string  `json:"elder_name"`  // 老人姓名
	ElderPhone string  `json:"elder_phone"` // 老人电话
	IDNum      string  `json:"id_num"`      // 身份证号
	BedName    string  `json:"bed_name"`    // 床位名称
	Balance    float64 `json:"balance"`     // 余额
}
