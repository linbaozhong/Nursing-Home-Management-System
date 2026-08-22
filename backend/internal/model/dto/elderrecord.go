package dto

// ============ ElderRecordController 请求 ============

// PageElderByKeyReq 分页查询长者请求
// @request
type PageElderByKeyReq struct {
	PageNum   *int    `json:"page_num" valid:"required"`  // 页码
	PageSize  *int    `json:"page_size" valid:"required"` // 条数
	ElderName *string `json:"elder_name"`                 // 老人姓名
	IDNum     *string `json:"id_num"`                     // 身份证号
	ElderSex  *string `json:"elder_sex"`                  // 老人性别
}

// // @request
// // PageSearchElderByKeyReq 分页搜索老人请求
// type PageSearchElderByKeyReq struct {
// 	PageNum  *int    `json:"page_num" valid:"required"`  // 页码
// 	PageSize *int    `json:"page_size" valid:"required"` // 条数
// 	Name     *string `json:"name"`                       // 老人姓名
// 	Phone    *string `json:"phone"`                      // 老人电话
// }

// EditElderLabelReq 编辑老人标签请求
// @request
type EditElderLabelReq struct {
	ElderID     *int64  `json:"elder_id" valid:"required"`      // 老人编号
	LabelIDList []int64 `json:"label_id_list" valid:"required"` // 标签编号列表
}

// AuditElderFeeReq 审核老人费用详情请求
// @request
type AuditElderFeeReq struct {
	ApplyID     *int64  `json:"apply_id" valid:"required"`     // 申请编号
	ElderID     *int64  `json:"elder_id" valid:"required"`     // 老人编号
	AuditResult *string `json:"audit_result" valid:"required"` // 审核结果
}

// ============ ElderRecordController 响应 ============

// PageElderByKeyResp 分页查询长者响应
// @response
type PageElderByKeyResp struct {
	ID        int64  `json:"id"`         // id
	BedName   string `json:"bed_name"`   // 床位名称
	Name      string `json:"name"`       // 姓名
	IDNum     string `json:"id_num"`     // 身份证号
	Age       int    `json:"age"`        // 年龄
	Sex       string `json:"sex"`        // 性别
	Phone     string `json:"phone"`      // 电话
	Address   string `json:"address"`    // 地址
	CheckFlag string `json:"check_flag"` // 入住状态
}

// GetElderRecordByIDResp 长者档案详情响应（定义见 common.go）

// GetElderFeeByIDResp 老人费用详情响应（定义见 common.go）

// GetElderLabelByIDLabelResp 客户标签响应（定义见 common.go）

// PageElderRecordByKeyReq 分页查询长者档案请求
// @request
type PageElderRecordByKeyReq struct {
	PageNum   *int    `json:"page_num" valid:"required"`  // 页码
	PageSize  *int    `json:"page_size" valid:"required"` // 条数
	ElderName *string `json:"elder_name"`                 // 老人姓名
	IDNum     *string `json:"id_num"`                     // 身份证号
	ElderSex  *string `json:"elder_sex"`                  // 老人性别
}

// AddElderRecordReq 新增长者档案请求
// @request
type AddElderRecordReq struct {
	ID         *int64  `json:"id"`                       // id
	Name       *string `json:"name" valid:"required"`    // 姓名
	IDNum      *string `json:"id_num" valid:"required"`  // 身份证号
	Sex        *string `json:"sex" valid:"required"`     // 性别
	Age        *int    `json:"age" valid:"required"`     // 年龄
	Phone      *string `json:"phone" valid:"required"`   // 电话
	Address    *string `json:"address" valid:"required"` // 地址
	NurseLevel *string `json:"nurse_level"`              // 护理等级
	CheckFlag  *string `json:"check_flag"`               // 入住状态
}

// EditElderRecordReq 编辑长者档案请求
// @request
type EditElderRecordReq struct {
	ID         *int64  `json:"id"`                       // id
	Name       *string `json:"name" valid:"required"`    // 姓名
	IDNum      *string `json:"id_num" valid:"required"`  // 身份证号
	Sex        *string `json:"sex" valid:"required"`     // 性别
	Age        *int    `json:"age" valid:"required"`     // 年龄
	Phone      *string `json:"phone" valid:"required"`   // 电话
	Address    *string `json:"address" valid:"required"` // 地址
	NurseLevel *string `json:"nurse_level"`              // 护理等级
	CheckFlag  *string `json:"check_flag"`               // 入住状态
}

// PageSearchEmergencyContactByKeyReq 分页查询紧急联系人请求
// @request
type PageSearchEmergencyContactByKeyReq struct {
	PageNum  *int    `json:"page_num" valid:"required"`  // 页码
	PageSize *int    `json:"page_size" valid:"required"` // 条数
	ElderID  *int64  `json:"elder_id" valid:"required"`  // 老人编号
	Key      *string `json:"key"`                        // 关键字(联系人姓名/电话), 可选
}

// PageLabelByKeyReq 分页查询客户标签请求
// @request
type PageLabelByKeyReq struct {
	PageNum  *int    `json:"page_num" valid:"required"`  // 页码
	PageSize *int    `json:"page_size" valid:"required"` // 条数
	Key      *string `json:"key"`                        // 关键字(标签名称), 可选
}

// AddEmergencyContactReq 新增紧急联系人请求
// @request
type AddEmergencyContactReq struct {
	ElderID  *int64  `json:"elder_id" valid:"required"` // 老人编号
	Name     *string `json:"name" valid:"required"`     // 联系人姓名
	Phone    *string `json:"phone" valid:"required"`    // 联系人电话
	Relation *string `json:"relation"`                  // 与老人关系
}

// EditEmergencyContactReq 编辑紧急联系人请求
// @request
type EditEmergencyContactReq struct {
	ID       *int64  `json:"id" valid:"required"` // 紧急联系人编号
	ElderID  *int64  `json:"elder_id"`            // 老人编号
	Name     *string `json:"name"`                // 联系人姓名
	Phone    *string `json:"phone"`               // 联系人电话
	Relation *string `json:"relation"`            // 与老人关系
}

// DeleteEmergencyContactReq 删除紧急联系人请求
// @request
type DeleteEmergencyContactReq struct {
	ID      *int64 `json:"id" valid:"required"` // 紧急联系人编号
	ElderID *int64 `json:"elder_id"`            // 老人编号(可选, 仅作校验)
}

// PageSearchEmergencyContactByKeyResp 分页查询紧急联系人响应
// @response
type PageSearchEmergencyContactByKeyResp struct {
	ID       int64  `json:"id"`       // 紧急联系人编号
	ElderId  int64  `json:"elder_id"` // 老人编号
	Name     string `json:"name"`     // 联系人姓名
	Phone    string `json:"phone"`    // 联系人电话
	Relation string `json:"relation"` // 与老人关系
	Remark   string `json:"remark"`   // 备注（DB 当前无该列，保留字段以兼容 gen 序列化）
}

// ListLabelResp 客户标签分类列表响应（定义见 common.go）
