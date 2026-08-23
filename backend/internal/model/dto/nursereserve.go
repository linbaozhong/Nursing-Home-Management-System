package dto

import (
	"github.com/linbaozhong/gentity/pkg/types"
	"time"
)

// ============ NurseReserveController 请求 ============

// PageNurseReserveByKeyReq 分页查询护理预定请求
// @request
type PageNurseReserveByKeyReq struct {
	PageNum     *int    `json:"page_num" valid:"required"`  // 页码
	PageSize    *int    `json:"page_size" valid:"required"` // 条数
	ElderName   *string `json:"elder_name"`                 // 老人姓名
	ServiceName *string `json:"service_name"`               // 项目名称
	BedName     *string `json:"bed_name"`                   // 床位名称
}

// AddNurseReserveReq 新增护理预定请求
// @request
type AddNurseReserveReq struct {
	ElderID      *int64       `json:"elder_id" valid:"required"`      // 老人编号
	ServiceName  *string      `json:"service_name" valid:"required"`  // 项目名称
	NeedDate     *int         `json:"need_date" valid:"required"`     // 所需时间
	ServicePrice *types.Money `json:"service_price" valid:"required"` // 服务费用
	ChargeMethod *string      `json:"charge_method" valid:"required"` // 收费方式
	Frequency    *int         `json:"frequency" valid:"required"`     // 服务次数
	PayAmount    *types.Money `json:"pay_amount" valid:"required"`    // 支付总额
}

// ExecuteNurseReserveReq 执行护理预定请求
// @request
type ExecuteNurseReserveReq struct {
	ID        *int64     `json:"id"`                          // id
	StaffID   *int64     `json:"staff_id" valid:"required"`   // 服务员工编号
	NurseDate *time.Time `json:"nurse_date" valid:"required"` // 护理时间
}

// ============ NurseReserveController 响应 ============

// PageNurseReserveByKeyResp 分页查询护理预定响应
// @response
type PageNurseReserveByKeyResp struct {
	ID           types.BigInt `json:"id"`            // id
	ElderName    string       `json:"elder_name"`    // 老人姓名
	BedName      string       `json:"bed_name"`      // 床位名称
	ServiceName  string       `json:"service_name"`  // 项目名称
	NeedDate     int          `json:"need_date"`     // 所需时间
	ServicePrice types.Money  `json:"service_price"` // 服务费用
	ChargeMethod string       `json:"charge_method"` // 收费方式
	Frequency    int          `json:"frequency"`     // 服务次数
	PayAmount    types.Money  `json:"pay_amount"`    // 支付总额
	NurseDate    time.Time    `json:"nurse_date"`    // 护理时间
	OrderFlag    string       `json:"order_flag"`    // 订单状态
}

// GetNurseReserveByReserveIdAndElderIdReq 按护理预定/老人编号获取护理预定请求
// @request
type GetNurseReserveByReserveIdAndElderIdReq struct {
	ReserveID *int64 `json:"reserve_id" valid:"required"` // 护理预定编号
	ElderID   *int64 `json:"elder_id" valid:"required"`   // 老人编号
}

// EditNurseReserveReq 编辑护理预定请求
// @request
type EditNurseReserveReq struct {
	ID           *int64       `json:"id"`                             // id
	ElderID      *int64       `json:"elder_id" valid:"required"`      // 老人编号
	ReserveDate  *time.Time   `json:"reserve_date" valid:"required"`  // 预定时间
	ServiceID    *int64       `json:"service_id" valid:"required"`    // 服务编号
	NeedDate     *int         `json:"need_date" valid:"required"`     // 所需时间
	ChargeMethod *string      `json:"charge_method" valid:"required"` // 收费方式
	Frequency    *int         `json:"frequency" valid:"required"`     // 服务次数
	ServicePrice *types.Money `json:"service_price" valid:"required"` // 服务费用
	PayAmount    *types.Money `json:"pay_amount" valid:"required"`    // 支付总额
}

// GetNurseReserveByReserveIdAndElderIdResp 按护理预定/老人编号获取护理预定响应
// @response
type GetNurseReserveByReserveIdAndElderIdResp struct {
	ID           types.BigInt `json:"id"`            // 护理预定编号
	ElderName    string       `json:"elder_name"`    // 老人姓名
	BedName      string       `json:"bed_name"`      // 床位名称
	ServiceName  string       `json:"service_name"`  // 项目名称
	NeedDate     int          `json:"need_date"`     // 所需时间
	ServicePrice types.Money  `json:"service_price"` // 服务费用
	ChargeMethod string       `json:"charge_method"` // 收费方式
	Frequency    int          `json:"frequency"`     // 服务次数
	PayAmount    types.Money  `json:"pay_amount"`    // 支付总额
	NurseDate    time.Time    `json:"nurse_date"`    // 护理时间
	OrderFlag    string       `json:"order_flag"`    // 订单状态
}
