package dto

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
	ID                        *int64                         `json:"id"`                                                    // id
	NursingGradeID            *int64                         `json:"nursing_grade_id" valid:"required"`                     // 护理等级编号
	CateringSetID             *int64                         `json:"catering_set_id" valid:"required"`                      // 餐饮套餐编号
	BedID                     *int64                         `json:"bed_id" valid:"required"`                               // 床位编号
	Name                      *string                        `json:"name" valid:"required"`                                 // 姓名
	IDNum                     *string                        `json:"id_num" valid:"required"`                               // 身份证号
	Age                       *int                           `json:"age" valid:"required"`                                  // 年龄
	Sex                       *string                        `json:"sex" valid:"required"`                                  // 性别
	Phone                     *string                        `json:"phone" valid:"required"`                                // 电话
	Address                   *string                        `json:"address" valid:"required"`                              // 地址
	StaffID                   *int64                         `json:"staff_id" valid:"required"`                             // 营销人员编号
	SignDate                  *string                        `json:"sign_date" valid:"required"`                            // 合同签订日期
	StartDate                 *string                        `json:"start_date" valid:"required"`                           // 合同开始日期
	EndDate                   *string                        `json:"end_date" valid:"required"`                             // 合同结束日期
	EmergencyContactQueryList []OperateEmergencyContactQuery `json:"operate_emergency_contact_query_list" valid:"required"` // 紧急联系人
}

// ============ CheckContractController 响应 ============

// @response
// PageCheckContractByKeyVO 分页查询入住签约响应
type PageCheckContractByKeyVO struct {
	Rank
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
// GetCheckContractByIDVO 根据编号获取入住签约响应（继承 OperateCheckContractQuery）
type GetCheckContractByIDVO struct {
	OperateCheckContractQuery
}
