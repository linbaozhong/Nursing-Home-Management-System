package dto

import (
	"github.com/linbaozhong/gentity/pkg/types"
	"time"
)

// ============ ReserveController 请求 ============

// PageReserveByKeyReq 分页查询预定请求
// @request
type PageReserveByKeyReq struct {
	PageNum    *int    `json:"page_num" valid:"required"`  // 页码
	PageSize   *int    `json:"page_size" valid:"required"` // 条数
	ElderName  *string `json:"elder_name"`                 // 老人姓名
	PayerPhone *string `json:"payer_phone"`                // 交款人电话
}

// AddReserveReq 新增预定请求
// @request
type AddReserveReq struct {
	BedID        *int64       `json:"bed_id" valid:"required"`        // 床位编号
	StaffID      *int64       `json:"staff_id" valid:"required"`      // 营销人员编号
	PayerName    *string      `json:"payer_name" valid:"required"`    // 交款人姓名
	PayerPhone   *string      `json:"payer_phone" valid:"required"`   // 交款人电话
	DueDate      *time.Time   `json:"due_date" valid:"required"`      // 预定到期时间
	Deposit      *types.Money `json:"deposit" valid:"required"`       // 预付定金
	ElderName    *string      `json:"elder_name" valid:"required"`    // 老人姓名
	IDNum        *string      `json:"id_num" valid:"required"`        // 老人身份证号
	ElderAge     *int         `json:"elder_age" valid:"required"`     // 老人年龄
	ElderSex     *string      `json:"elder_sex" valid:"required"`     // 老人性别
	ElderPhone   *string      `json:"elder_phone" valid:"required"`   // 老人电话
	ElderAddress *string      `json:"elder_address" valid:"required"` // 老人地址
}

// GetReserveByReserveIDAndElderIDReq 根据预定编号和老人编号获取预定信息请求
// @request
type GetReserveByReserveIDAndElderIDReq struct {
	ReserveID *int64 `json:"reserve_id" valid:"required"` // 预定编号
	ElderID   *int64 `json:"elder_id" valid:"required"`   // 老人编号
}

// EditReserveReq 编辑预定请求
// @request
type EditReserveReq struct {
	ID           *int64       `json:"id"`                             // id
	BedID        *int64       `json:"bed_id" valid:"required"`        // 床位编号
	StaffID      *int64       `json:"staff_id" valid:"required"`      // 营销人员编号
	PayerName    *string      `json:"payer_name" valid:"required"`    // 交款人姓名
	PayerPhone   *string      `json:"payer_phone" valid:"required"`   // 交款人电话
	DueDate      *time.Time   `json:"due_date" valid:"required"`      // 预定到期时间
	Deposit      *types.Money `json:"deposit" valid:"required"`       // 预付定金
	ElderName    *string      `json:"elder_name" valid:"required"`    // 老人姓名
	IDNum        *string      `json:"id_num" valid:"required"`        // 老人身份证号
	ElderAge     *int         `json:"elder_age" valid:"required"`     // 老人年龄
	ElderSex     *string      `json:"elder_sex" valid:"required"`     // 老人性别
	ElderPhone   *string      `json:"elder_phone" valid:"required"`   // 老人电话
	ElderAddress *string      `json:"elder_address" valid:"required"` // 老人地址
}

// ============ ReserveController 响应 ============

// PageReserveByKeyResp 分页查询预定响应
// @response
type PageReserveByKeyResp struct {
	ReserveID   types.BigInt `json:"reserve_id"`   // 预定编号
	ElderID     types.BigInt `json:"elder_id"`     // 老人编号
	StaffName   string       `json:"staff_name"`   // 销售人员姓名
	BedName     string       `json:"bed_name"`     // 预定床位名称
	ElderName   string       `json:"elder_name"`   // 老人姓名
	IDNum       string       `json:"id_num"`       // 身份证号
	ElderSex    string       `json:"elder_sex"`    // 老人性别
	ElderAge    int          `json:"elder_age"`    // 老人年龄
	PayerPhone  string       `json:"payer_phone"`  // 交款人联系电话
	Deposit     types.Money  `json:"deposit"`      // 定金
	ReserveFlag string       `json:"reserve_flag"` // 退款状态
	CheckFlag   string       `json:"check_flag"`   // 入住状态
	DueDate     time.Time    `json:"due_date"`     // 预定到期时间

}

// GetReserveByReserveIDAndElderIDResp 根据预定编号和老人编号获取预定信息响应
// @response
type GetReserveByReserveIDAndElderIDResp struct {
	ElderName    string       `json:"elder_name"`    // 老人姓名
	IDNum        string       `json:"id_num"`        // 身份证号
	ElderSex     string       `json:"elder_sex"`     // 老人性别
	ElderAge     int          `json:"elder_age"`     // 老人年龄
	ElderPhone   string       `json:"elder_phone"`   // 老人联系电话
	ElderAddress string       `json:"elder_address"` // 老人地址
	BedName      string       `json:"bed_name"`      // 预定床位名称
	StaffID      types.BigInt `json:"staff_id"`      // 销售人员编号
	PayerName    string       `json:"payer_name"`    // 交款人姓名
	PayerPhone   string       `json:"payer_phone"`   // 交款人联系电话
	DueDate      time.Time    `json:"due_date"`      // 预定到期时间
	Deposit      types.Money  `json:"deposit"`       // 定金
}
