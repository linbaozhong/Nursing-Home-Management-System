package dto

import (
	"github.com/linbaozhong/gentity/pkg/types"
	"time"
)

// ============ CheckContractController 请求 ============

// @request
// PageCheckContractByKeyQuery 分页查询入住签约请求
type PageCheckContractByKeyQuery struct {
	PageNum  *int    `json:"page_num" valid:"required"`  // 页码
	PageSize *int    `json:"page_size" valid:"required"` // 条数
	Name     *string `json:"name"`                       // 姓名
	Sex      *string `json:"sex"`                        // 性别
	IDNum    *string `json:"id_num"`                     // 身份证号
}

// @request
// OperateCheckContractQuery 操作入住签约请求
type OperateCheckContractQuery struct {
	ID                        *types.BigInt                  `json:"id"`                                                    // id
	NursingGradeID            *types.BigInt                  `json:"nursing_grade_id" valid:"required"`                     // 护理等级编号
	CateringSetID             *types.BigInt                  `json:"catering_set_id" valid:"required"`                      // 餐饮套餐编号
	BedID                     *types.BigInt                  `json:"bed_id" valid:"required"`                               // 床位编号
	Name                      *string                        `json:"name" valid:"required"`                                 // 姓名
	IDNum                     *string                        `json:"id_num" valid:"required"`                               // 身份证号
	Age                       *int                           `json:"age" valid:"required"`                                  // 年龄
	Sex                       *string                        `json:"sex" valid:"required"`                                  // 性别
	Phone                     *string                        `json:"phone" valid:"required"`                                // 电话
	Address                   *string                        `json:"address" valid:"required"`                              // 地址
	StaffID                   *types.BigInt                  `json:"staff_id" valid:"required"`                             // 营销人员编号
	SignDate                  *time.Time                     `json:"sign_date" valid:"required"`                            // 合同签订日期
	StartDate                 *time.Time                     `json:"start_date" valid:"required"`                           // 合同开始日期
	EndDate                   *time.Time                     `json:"end_date" valid:"required"`                             // 合同结束日期
	EmergencyContactQueryList []OperateEmergencyContactQuery `json:"operate_emergency_contact_query_list" valid:"required"` // 紧急联系人
}

// ============ CheckContractController 响应 ============

// @request
// PageSearchElderByKeyQuery 分页搜索老人请求
type PageSearchElderByKeyQuery struct {
	PageNum  *int    `json:"page_num" valid:"required"`  // 页码
	PageSize *int    `json:"page_size" valid:"required"` // 条数
	Name     *string `json:"name"`                       // 姓名
	Phone    *string `json:"phone"`                      // 联系电话
}

// @response
// PageCheckContractByKeyVO 分页查询入住签约响应
type PageCheckContractByKeyVO struct {
	ID        int64  `json:"id"`         // id
	Name      string `json:"name"`       // 姓名
	IDNum     string `json:"id_num"`     // 身份证号
	Age       int    `json:"age"`        // 年龄
	Sex       string `json:"sex"`        // 性别
	Phone     string `json:"phone"`      // 电话
	Address   string `json:"address"`    // 地址
	CheckFlag string `json:"check_flag"` // 入住状态
}

// @response
// PageSearchElderByKeyVO 分页搜索老人响应
type PageSearchElderByKeyVO struct {
	ID        int64  `json:"id"`         // id
	Name      string `json:"name"`       // 姓名
	IDNum     string `json:"id_num"`     // 身份证号
	Sex       string `json:"sex"`        // 性别
	Phone     string `json:"phone"`      // 电话
	Address   string `json:"address"`    // 地址
	CheckFlag string `json:"check_flag"` // 入住状态
}

// @response
// GetCheckContractByIDVO 根据编号获取入住签约响应（继承 OperateCheckContractQuery）
type GetCheckContractByIDVO struct {
	OperateCheckContractQuery
	StaffID                          *int64                         `json:"staff_id"`                             // 营销人员编号
	SignDate                         *time.Time                     `json:"sign_date"`                            // 合同签订日期
	StartDate                        *time.Time                     `json:"start_date"`                           // 合同开始日期
	EndDate                          *time.Time                     `json:"end_date"`                             // 合同结束日期
	OperateEmergencyContactQueryList []OperateEmergencyContactQuery `json:"operate_emergency_contact_query_list"` // 紧急联系人
}
