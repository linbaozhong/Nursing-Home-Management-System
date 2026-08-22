package dto

import (
	"github.com/linbaozhong/gentity/pkg/types"
	"time"
)

// ============ CheckContractController 请求 ============

// PageCheckContractByKeyReq 分页查询入住签约请求
// @request
type PageCheckContractByKeyReq struct {
	PageNum  *int    `json:"page_num" valid:"required"`  // 页码
	PageSize *int    `json:"page_size" valid:"required"` // 条数
	Name     *string `json:"name"`                       // 姓名
	Sex      *string `json:"sex"`                        // 性别
	IDNum    *string `json:"id_num"`                     // 身份证号
}

// OperateCheckContractReq 操作入住签约请求
// @request
type OperateCheckContractReq struct {
	ID                        *types.BigInt                `json:"id"`                                                    // id
	NursingGradeID            *types.BigInt                `json:"nursing_grade_id" valid:"required"`                     // 护理等级编号
	CateringSetID             *types.BigInt                `json:"catering_set_id" valid:"required"`                      // 餐饮套餐编号
	BedID                     *types.BigInt                `json:"bed_id" valid:"required"`                               // 床位编号
	Name                      *string                      `json:"name" valid:"required"`                                 // 姓名
	IDNum                     *string                      `json:"id_num" valid:"required"`                               // 身份证号
	Age                       *int                         `json:"age" valid:"required"`                                  // 年龄
	Sex                       *string                      `json:"sex" valid:"required"`                                  // 性别
	Phone                     *string                      `json:"phone" valid:"required"`                                // 电话
	Address                   *string                      `json:"address" valid:"required"`                              // 地址
	StaffID                   *types.BigInt                `json:"staff_id" valid:"required"`                             // 营销人员编号
	SignDate                  *time.Time                   `json:"sign_date" valid:"required"`                            // 合同签订日期
	StartDate                 *time.Time                   `json:"start_date" valid:"required"`                           // 合同开始日期
	EndDate                   *time.Time                   `json:"end_date" valid:"required"`                             // 合同结束日期
	EmergencyContactQueryList []OperateEmergencyContactReq `json:"operate_emergency_contact_query_list" valid:"required"` // 紧急联系人
}

// ============ CheckContractController 响应 ============

// PageSearchElderByKeyReq 分页搜索老人请求
// @request
type PageSearchElderByKeyReq struct {
	PageNum  *int    `json:"page_num" valid:"required"`  // 页码
	PageSize *int    `json:"page_size" valid:"required"` // 条数
	Name     *string `json:"name"`                       // 姓名
	Phone    *string `json:"phone"`                      // 联系电话
}

// PageCheckContractByKeyResp 分页查询入住签约响应
// @response
type PageCheckContractByKeyResp struct {
	ID        int64  `json:"id"`         // id
	Name      string `json:"name"`       // 姓名
	IDNum     string `json:"id_num"`     // 身份证号
	Age       int    `json:"age"`        // 年龄
	Sex       string `json:"sex"`        // 性别
	Phone     string `json:"phone"`      // 电话
	Address   string `json:"address"`    // 地址
	CheckFlag string `json:"check_flag"` // 入住状态
}

// PageSearchElderByKeyResp 分页搜索老人响应
// @response
type PageSearchElderByKeyResp struct {
	ID        int64  `json:"id"`         // id
	Name      string `json:"name"`       // 姓名
	IDNum     string `json:"id_num"`     // 身份证号
	Sex       string `json:"sex"`        // 性别
	Phone     string `json:"phone"`      // 电话
	Address   string `json:"address"`    // 地址
	CheckFlag string `json:"check_flag"` // 入住状态
}

// GetCheckContractByIDResp 根据编号获取入住签约响应（继承 OperateCheckContractReq）
// @response
type GetCheckContractByIDResp struct {
	OperateCheckContractReq
	StaffID                        *int64                       `json:"staff_id"`                             // 营销人员编号
	SignDate                       *time.Time                   `json:"sign_date"`                            // 合同签订日期
	StartDate                      *time.Time                   `json:"start_date"`                           // 合同开始日期
	EndDate                        *time.Time                   `json:"end_date"`                             // 合同结束日期
	OperateEmergencyContactReqList []OperateEmergencyContactReq `json:"operate_emergency_contact_query_list"` // 紧急联系人
}
