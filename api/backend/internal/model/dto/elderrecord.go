package dto

// ============ ElderRecordController 请求 ============

// @request
// PageElderByKeyQuery 分页查询长者请求
type PageElderByKeyQuery struct {
	PageNum   *int    `json:"page_num" valid:"required"`  // 页码
	PageSize  *int    `json:"page_size" valid:"required"` // 条数
	ElderName *string `json:"elder_name"`                 // 老人姓名
	IDNum     *string `json:"id_num"`                     // 身份证号
	ElderSex  *string `json:"elder_sex"`                  // 老人性别
}

// @request
// PageSearchElderByKeyQuery 分页搜索老人请求
type PageSearchElderByKeyQuery struct {
	PageNum  *int    `json:"page_num" valid:"required"`  // 页码
	PageSize *int    `json:"page_size" valid:"required"` // 条数
	Name     *string `json:"name"`                       // 老人姓名
	Phone    *string `json:"phone"`                      // 老人电话
}

// @request
// EditElderLabelQuery 编辑老人标签请求
type EditElderLabelQuery struct {
	ElderID     *int64  `json:"elder_id" valid:"required"`      // 老人编号
	LabelIDList []int64 `json:"label_id_list" valid:"required"` // 标签编号列表
}

// @request
// AuditElderFeeQuery 审核老人费用详情请求
type AuditElderFeeQuery struct {
	ApplyID     *int64  `json:"apply_id" valid:"required"`     // 申请编号
	ElderID     *int64  `json:"elder_id" valid:"required"`     // 老人编号
	AuditResult *string `json:"audit_result" valid:"required"` // 审核结果
}

// ============ ElderRecordController 响应 ============

// @response
// PageElderByKeyVO 分页查询长者响应
type PageElderByKeyVO struct {
	Rank
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

// GetElderRecordByIDVO 长者档案详情响应（定义见 common.go）

// GetElderFeeByIDVO 老人费用详情响应（定义见 common.go）

// GetElderLabelByIDLabelVO 客户标签响应（定义见 common.go）

// @request
// PageElderRecordByKeyQuery 分页查询长者档案请求
type PageElderRecordByKeyQuery struct {
	PageNum   *int    `json:"page_num" valid:"required"`  // 页码
	PageSize  *int    `json:"page_size" valid:"required"` // 条数
	ElderName *string `json:"elder_name"`                 // 老人姓名
	IDNum     *string `json:"id_num"`                     // 身份证号
	ElderSex  *string `json:"elder_sex"`                  // 老人性别
}

// @request
// AddElderRecordQuery 新增长者档案请求
type AddElderRecordQuery struct {
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

// @request
// EditElderRecordQuery 编辑长者档案请求
type EditElderRecordQuery struct {
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

// @request
// PageSearchEmergencyContactByKeyQuery 分页查询紧急联系人请求
type PageSearchEmergencyContactByKeyQuery struct {
	PageNum  *int    `json:"page_num" valid:"required"`  // 页码
	PageSize *int    `json:"page_size" valid:"required"` // 条数
	ElderID  *int64  `json:"elder_id" valid:"required"`  // 老人编号
	Key      *string `json:"key"`                        // 关键字(联系人姓名/电话), 可选
}

// @request
// PageLabelByKeyQuery 分页查询客户标签请求
type PageLabelByKeyQuery struct {
	PageNum  *int    `json:"page_num" valid:"required"`  // 页码
	PageSize *int    `json:"page_size" valid:"required"` // 条数
	Key      *string `json:"key"`                        // 关键字(标签名称), 可选
}

// ListLabelVO 客户标签分类列表响应（定义见 common.go）
