package do

import (
	"github.com/linbaozhong/gentity/pkg/ace/pool"
	"github.com/linbaozhong/gentity/pkg/types"
)

// Accident
// @tablename accident
type Accident struct {
	pool.Model
	Description   types.String `json:"description,omitempty" db:"'description'"`                 // 事故描述
	HandleResult  types.String `json:"handle_result,omitempty" db:"'handle_result'"`             // 处理结果/整改措施
	Picture       types.String `json:"picture,omitempty" db:"'picture'"`                         // 事故图片
	CreateId      types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`             // 创建人id
	CreateTime    types.Time   `json:"create_time,omitempty" db:"'create_time'"`                 // 创建时间
	ElderId       types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`               // 老人id
	HandleStaffId types.BigInt `json:"handle_staff_id,omitempty" db:"'handle_staff_id' size:20"` // 处理责任人id
	Id            types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`                   // id
	OccurDate     types.Time   `json:"occur_date,omitempty" db:"'occur_date'"`                   // 发生时间
	StaffId       types.BigInt `json:"staff_id,omitempty" db:"'staff_id' size:20"`               // 值班护工id
	TenantId      types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`             // 租户id
	UpdateId      types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`             // 修改人id
	UpdateTime    types.Time   `json:"update_time,omitempty" db:"'update_time'"`                 // 修改时间
	Severity      types.Int8   `json:"severity,omitempty" db:"'severity' size:3"`                // 严重程度：1轻/2中/3重
	State         types.Int8   `json:"state,omitempty" db:"'state' size:3"`                      // 管理状态：-1=删除，0=禁用，1=可用
}

// Active
// @tablename active
type Active struct {
	pool.Model
	ActivePicture types.String `json:"active_picture,omitempty" db:"'active_picture'"` // 活动图片
	Address       types.String `json:"address,omitempty" db:"'address'"`               // 活动地点
	Content       types.String `json:"content,omitempty" db:"'content'"`               // 活动内容
	Name          types.String `json:"name,omitempty" db:"'name'"`                     // 活动名称
	Organizer     types.String `json:"organizer,omitempty" db:"'organizer'"`           // 组织者姓名
	Phone         types.String `json:"phone,omitempty" db:"'phone'"`                   // 组织者电话
	Theme         types.String `json:"theme,omitempty" db:"'theme'"`                   // 活动主题
	ActiveDate    types.Time   `json:"active_date,omitempty" db:"'active_date'"`       // 活动日期
	CreateId      types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`   // 创建人id
	CreateTime    types.Time   `json:"create_time,omitempty" db:"'create_time'"`       // 创建时间
	Id            types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`         // id
	TenantId      types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`   // 租户id
	TypeId        types.BigInt `json:"type_id,omitempty" db:"'type_id' size:20"`       // 活动类别id
	UpdateId      types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`   // 修改人id
	UpdateTime    types.Time   `json:"update_time,omitempty" db:"'update_time'"`       // 修改时间
	State         types.Int8   `json:"state,omitempty" db:"'state' size:3"`            // 管理状态：-1=删除，0=禁用，1=可用
}

// ActiveParticipant
// @tablename active_participant
type ActiveParticipant struct {
	pool.Model
	ActiveId   types.BigInt `json:"active_id,omitempty" db:"'active_id' size:20"` // 活动id
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	ElderId    types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`   // 老人id
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// ActiveType
// @tablename active_type
type ActiveType struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name'"`                   // 活动类型名称
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	State      types.Int8   `json:"state,omitempty" db:"'state' size:3"`          // 管理状态：-1=删除，0=禁用，1=可用
}

// Assessment 评估记录表
// @tablename assessment
type Assessment struct {
	pool.Model
	ResultLevel    types.String `json:"result_level,omitempty" db:"'result_level'"`              // 评估结果（护理等级）
	ScaleData      types.String `json:"scale_data,omitempty" db:"'scale_data'"`                  // 量表原始数据
	AssessmentDate types.Time   `json:"assessment_date,omitempty" db:"'assessment_date'"`        // 评估日期
	CreatedAt      types.Time   `json:"created_at,omitempty" db:"'created_at' <-"`               // 创建时间
	DeletedAt      types.Time   `json:"deleted_at,omitempty" db:"'deleted_at'"`                  // 软删除时间
	ElderId        types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`              // 老人ID
	EvaluatorId    types.BigInt `json:"evaluator_id,omitempty" db:"'evaluator_id' size:20"`      // 评估人ID
	Id             types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`                  // 主键ID
	TenantId       types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`            // 租户ID
	UpdatedAt      types.Time   `json:"updated_at,omitempty" db:"'updated_at' <-"`               // 更新时间
	AssessmentType types.Int8   `json:"assessment_type,omitempty" db:"'assessment_type' size:3"` // 评估类型：1-能力评估，2-健康评估
}

// AuditLog 数据轨迹日志
// @tablename audit_log
type AuditLog struct {
	pool.Model
	Action       types.String `json:"action,omitempty" db:"'action'"`                   // 操作：create/update/delete
	ChangeAfter  types.String `json:"change_after,omitempty" db:"'change_after'"`       // 变更后整行快照(JSON)
	ChangeLabel  types.String `json:"change_label,omitempty" db:"'change_label'"`       // 可读变更摘要(中文字段名)
	Comment      types.String `json:"comment,omitempty" db:"'comment'"`                 // 业务备注
	OperatorName types.String `json:"operator_name,omitempty" db:"'operator_name'"`     // 操作员名称
	Table        types.String `json:"table,omitempty" db:"'table'"`                     // 被操作表名
	CreateTime   types.Time   `json:"create_time,omitempty" db:"'create_time'"`         // 操作时间
	Id           types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`           // 主键id
	OperatorId   types.BigInt `json:"operator_id,omitempty" db:"'operator_id' size:20"` // 操作员id
	RowId        types.BigInt `json:"row_id,omitempty" db:"'row_id' size:20"`           // 被操作行主键id
	TenantId     types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`     // 租户id
}

// Auth
// @tablename auth
type Auth struct {
	pool.Model
	Icon       types.String `json:"icon,omitempty" db:"'icon'"`                   // 权限图标
	Method     types.String `json:"method,omitempty" db:"'method'"`               // 权限请求方式（GET/POST/PUT/DELETE）
	Name       types.String `json:"name,omitempty" db:"'name'"`                   // 权限名称
	Path       types.String `json:"path,omitempty" db:"'path'"`                   // 权限path
	Title      types.String `json:"title,omitempty" db:"'title'"`                 // 权限标题
	Type       types.String `json:"type,omitempty" db:"'type'"`                   // 权限类别（MENU/BTN）
	Url        types.String `json:"url,omitempty" db:"'url'"`                     // 权限url
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	ParentId   types.BigInt `json:"parent_id,omitempty" db:"'parent_id' size:20"` // 父级id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// BaseAttachment
// @tablename base_attachment
type BaseAttachment struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name'"`                   // 上传后文件名
	Path       types.String `json:"path,omitempty" db:"'path'"`                   // 文件绝对路径
	RealName   types.String `json:"real_name,omitempty" db:"'real_name'"`         // 文件真实名称
	Suff       types.String `json:"suff,omitempty" db:"'suff'"`                   // 文件后缀
	Url        types.String `json:"url,omitempty" db:"'url'"`                     // url相对路径
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	Size       types.Money  `json:"size,omitempty" db:"'size' size:19"`           // 文件大小 B
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	State      types.Int8   `json:"state,omitempty" db:"'state' size:3"`          // 删除状态(Y/N)
}

// Bed 床位表
// @tablename bed
type Bed struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name'"`                   // 床位编号
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	Price      types.Money  `json:"price,omitempty" db:"'price' size:19"`         // 床位费（分）
	RoomId     types.BigInt `json:"room_id,omitempty" db:"'room_id' size:20"`     // 房间id
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	TypeId     types.BigInt `json:"type_id,omitempty" db:"'type_id' size:20"`     // 床型编号：1-普通床，2-护理床，3-加床(关联 material_type.id，kind=1)
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time' <-"`  // 修改时间
	State      types.Int8   `json:"state,omitempty" db:"'state' size:3"`          // 管理状态：-1=删除，0=禁用，1=正常
	Status     types.Uint8  `json:"status,omitempty" db:"'status' size:3"`        // 床位状态(空闲/预定/入住/退住审核/维修)
}

// Bill 账单表
// @tablename bill
type Bill struct {
	pool.Model
	BillNo      types.String `json:"bill_no,omitempty" db:"'bill_no'"`                   // 账单编号
	BillPeriod  types.String `json:"bill_period,omitempty" db:"'bill_period'"`           // 账单周期
	CreateId    types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`       // 创建人id
	CreateTime  types.Time   `json:"create_time,omitempty" db:"'create_time'"`           // 创建时间
	DueDate     types.Time   `json:"due_date,omitempty" db:"'due_date'"`                 // 缴费截止日
	ElderId     types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`         // 老人ID
	Id          types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`             // 主键ID
	PaidAmount  types.Money  `json:"paid_amount,omitempty" db:"'paid_amount' size:19"`   // 已付金额（分）
	TenantId    types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`       // 租户ID
	TotalAmount types.Money  `json:"total_amount,omitempty" db:"'total_amount' size:19"` // 总金额（分）
	UpdateId    types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`       // 修改人id
	UpdateTime  types.Time   `json:"update_time,omitempty" db:"'update_time'"`           // 修改时间
	Status      types.Uint8  `json:"status,omitempty" db:"'status' size:3"`              // 状态：0-未支付，1-部分支付，2-已支付，3-逾期
}

// BillItem 账单明细表
// @tablename bill_item
type BillItem struct {
	pool.Model
	Description types.String  `json:"description,omitempty" db:"'description'"`         // 描述
	Amount      types.Money   `json:"amount,omitempty" db:"'amount' size:19"`           // 小计（分）
	BillId      types.BigInt  `json:"bill_id,omitempty" db:"'bill_id' size:20"`         // 账单ID
	CreatedTime types.Time    `json:"created_time,omitempty" db:"'created_time' <-"`    // 创建时间
	FeeItemId   types.BigInt  `json:"fee_item_id,omitempty" db:"'fee_item_id' size:20"` // 费用项ID
	Id          types.BigInt  `json:"id,omitempty" db:"'id' pk auto size:20"`           // 主键ID
	Quantity    types.Float64 `json:"quantity,omitempty" db:"'quantity' size:10|2"`     // 数量
	UnitPrice   types.Money   `json:"unit_price,omitempty" db:"'unit_price' size:19"`   // 单价（分）
}

// Building
// @tablename building
type Building struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name'"`                   // 楼栋名称
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	FloorNum   types.Int32  `json:"floor_num,omitempty" db:"'floor_num' size:10"` // 楼层数量
	State      types.Int8   `json:"state,omitempty" db:"'state' size:3"`          // 管理状态：-1=删除，0=禁用，1=可用
}

// CateringSet
// @tablename catering_set
type CateringSet struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name'"`                       // 餐饮套餐名称
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`     // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`         // 创建时间
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`           // id
	MonthPrice types.Money  `json:"month_price,omitempty" db:"'month_price' size:19"` // 月套餐费用
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`     // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`     // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`         // 修改时间
	State      types.Int8   `json:"state,omitempty" db:"'state' size:3"`              // 管理状态：-1=删除，0=禁用，1=可用
}

// CommunicationRecord
// @tablename communication_record
type CommunicationRecord struct {
	pool.Model
	CommunicationRecord types.String `json:"communication_record,omitempty" db:"'communication_record'"` // 沟通记录
	CreateId            types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`               // 创建人id
	CreateTime          types.Time   `json:"create_time,omitempty" db:"'create_time'"`                   // 创建时间
	ElderId             types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`                 // 老人id
	Id                  types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`                     // id
	RecordDate          types.Time   `json:"record_date,omitempty" db:"'record_date'"`                   // 记录时间
	TenantId            types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`               // 租户id
	UpdateId            types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`               // 修改人id
	UpdateTime          types.Time   `json:"update_time,omitempty" db:"'update_time'"`                   // 修改时间
	State               types.Int8   `json:"state,omitempty" db:"'state' size:3"`                        // 管理状态：-1=删除，0=禁用，1=可用
}

// Consult
// @tablename consult
type Consult struct {
	pool.Model
	ConsultContent types.String `json:"consult_content,omitempty" db:"'consult_content'"` // 咨询内容
	Name           types.String `json:"name,omitempty" db:"'name'"`                       // 咨询人姓名
	Phone          types.String `json:"phone,omitempty" db:"'phone'"`                     // 咨询人电话
	Relation       types.String `json:"relation,omitempty" db:"'relation'"`               // 与老人关系
	ConsultDate    types.Time   `json:"consult_date,omitempty" db:"'consult_date'"`       // 咨询日期
	CreateId       types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`     // 创建人id
	CreateTime     types.Time   `json:"create_time,omitempty" db:"'create_time'"`         // 创建时间
	ElderId        types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`       // 老人id
	Id             types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`           // id
	SourceId       types.BigInt `json:"source_id,omitempty" db:"'source_id' size:20"`     // 来源渠道id
	StaffId        types.BigInt `json:"staff_id,omitempty" db:"'staff_id' size:20"`       // 接待人id
	TenantId       types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`     // 租户id
	UpdateId       types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`     // 修改人id
	UpdateTime     types.Time   `json:"update_time,omitempty" db:"'update_time'"`         // 修改时间
}

// Consume
// @tablename consume
type Consume struct {
	pool.Model
	ConsumeType   types.String `json:"consume_type,omitempty" db:"'consume_type'"`             // 消费类别
	ConsumeAmount types.Money  `json:"consume_amount,omitempty" db:"'consume_amount' size:19"` // 消费金额
	ConsumeDate   types.Time   `json:"consume_date,omitempty" db:"'consume_date'"`             // 消费日期
	SourceType    types.String `json:"source_type,omitempty" db:"'source_type'"`               // 来源：ORDER/NURSE_RESERVE/MANUAL
	OutTradeNo    types.String `json:"out_trade_no,omitempty" db:"'out_trade_no'"`             // 外部交易单号（对账）
	CreateId      types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`           // 创建人id
	CreateTime    types.Time   `json:"create_time,omitempty" db:"'create_time'"`               // 创建时间
	ElderId       types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`             // 老人id
	Id            types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`                 // id
	SourceId      types.BigInt `json:"source_id,omitempty" db:"'source_id' size:20"`           // 来源业务主键id
	TenantId      types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`           // 租户id
	UpdateId      types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`           // 修改人id
	UpdateTime    types.Time   `json:"update_time,omitempty" db:"'update_time'"`               // 修改时间
}

// Contract
// @tablename contract
type Contract struct {
	pool.Model
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	ElderId    types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`   // 老人id
	EndDate    types.Time   `json:"end_date,omitempty" db:"'end_date'"`           // 合同结束日期
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	SignDate   types.Time   `json:"sign_date,omitempty" db:"'sign_date'"`         // 合同签订日期
	StaffId    types.BigInt `json:"staff_id,omitempty" db:"'staff_id' size:20"`   // 销售人员id
	StartDate  types.Time   `json:"start_date,omitempty" db:"'start_date'"`       // 合同开始日期
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// DepositInfo
// @tablename deposit_info
type DepositInfo struct {
	pool.Model
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`     // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`         // 创建时间
	DepositId  types.BigInt `json:"deposit_id,omitempty" db:"'deposit_id' size:20"`   // 药品缴存id
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`           // id
	MedicineId types.BigInt `json:"medicine_id,omitempty" db:"'medicine_id' size:20"` // 药品id
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`     // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`     // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`         // 修改时间
	DepositNum types.Int32  `json:"deposit_num,omitempty" db:"'deposit_num' size:10"` // 缴存数量
	SurplusNum types.Int32  `json:"surplus_num,omitempty" db:"'surplus_num' size:10"` // 剩余数量
	WarnNum    types.Int32  `json:"warn_num,omitempty" db:"'warn_num' size:10"`       // 预警数量
}

// Dishes
// @tablename dishes
type Dishes struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name'"`                   // 菜品名称
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	Price      types.Money  `json:"price,omitempty" db:"'price' size:19"`         // 菜品价格
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	TypeId     types.BigInt `json:"type_id,omitempty" db:"'type_id' size:20"`     // 菜品类别id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	State      types.Int8   `json:"state,omitempty" db:"'state' size:3"`          // 管理状态：-1=删除，0=禁用，1=可用
}

// DishesType
// @tablename dishes_type
type DishesType struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name'"`                   // 菜品类别名称
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	State      types.Int8   `json:"state,omitempty" db:"'state' size:3"`          // 管理状态：-1=删除，0=禁用，1=可用
}

// DrugDeposit
// @tablename drug_deposit
type DrugDeposit struct {
	pool.Model
	Mode       types.String `json:"mode,omitempty" db:"'mode'"`                   // 缴存药品方式
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	ElderId    types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`   // 老人id
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	State      types.Int8   `json:"state,omitempty" db:"'state' size:3"`          // 管理状态：-1=删除，0=禁用，1=可用
	Status     types.Uint8  `json:"status,omitempty" db:"'status' size:3"`        // 缴存状态
}

// Elder
// @tablename elder
type Elder struct {
	pool.Model
	Address        types.String `json:"address,omitempty" db:"'address'"`                           // 地址
	IdNum          types.String `json:"id_num,omitempty" db:"'id_num'"`                             // 身份证号
	Name           types.String `json:"name,omitempty" db:"'name'"`                                 // 老人姓名
	Phone          types.String `json:"phone,omitempty" db:"'phone'"`                               // 老人电话
	Sex            types.String `json:"sex,omitempty" db:"'sex'"`                                   // 性别(男/女)
	Balance        types.Money  `json:"balance,omitempty" db:"'balance' size:19"`                   // 余额
	BedId          types.BigInt `json:"bed_id,omitempty" db:"'bed_id' size:20"`                     // 床位id
	CateringSetId  types.BigInt `json:"catering_set_id,omitempty" db:"'catering_set_id' size:20"`   // 餐饮套餐id
	CreateId       types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`               // 创建人id
	CreateTime     types.Time   `json:"create_time,omitempty" db:"'create_time'"`                   // 创建时间
	Id             types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`                     // id
	NursingGradeId types.BigInt `json:"nursing_grade_id,omitempty" db:"'nursing_grade_id' size:20"` // 护理等级id
	TenantId       types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`               // 租户id
	UpdateId       types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`               // 修改人id
	UpdateTime     types.Time   `json:"update_time,omitempty" db:"'update_time'"`                   // 修改时间
	Age            types.Int32  `json:"age,omitempty" db:"'age' size:10"`                           // 年龄
	Status         types.Uint8  `json:"status,omitempty" db:"'status' size:3"`                      // 入住状态
}

// ElderAccountLedger 老人资金明细台账
// @tablename elder_account_ledger
type ElderAccountLedger struct {
	pool.Model
	BusinessNo   types.String `json:"business_no,omitempty" db:"'business_no'"`             // 业务单号（可空，用于人工对账）
	Remark       types.String `json:"remark,omitempty" db:"'remark'"`                       // 备注
	SourceType   types.String `json:"source_type,omitempty" db:"'source_type'"`             // 来源类型：BILL_INCOME/RECHARGE/FEED/NURSING/REFUND
	Amount       types.Money  `json:"amount,omitempty" db:"'amount' size:19"`               // 变更金额（分），恒为正
	BalanceAfter types.Money  `json:"balance_after,omitempty" db:"'balance_after' size:19"` // 变动后的账户余额（分）
	CreateTime   types.Time   `json:"create_time,omitempty" db:"'create_time'"`             // 创建时间
	ElderId      types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`           // 老人id
	Id           types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`               // id
	OperatorId   types.BigInt `json:"operator_id,omitempty" db:"'operator_id' size:20"`     // 操作员id
	SourceId     types.BigInt `json:"source_id,omitempty" db:"'source_id' size:20"`         // 来源业务表主键id
	TenantId     types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`         // 租户id
	Direction    types.Int8   `json:"direction,omitempty" db:"'direction' size:3"`          // 1=入账(充值/退款入账)，2=出账(消费/扣费)
}

// ElderLabel
// @tablename elder_label
type ElderLabel struct {
	pool.Model
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	ElderId    types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`   // 老人id
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	LabelId    types.BigInt `json:"label_id,omitempty" db:"'label_id' size:20"`   // 标签id
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// EmergencyContact
// @tablename emergency_contact
type EmergencyContact struct {
	pool.Model
	Email      types.String `json:"email,omitempty" db:"'email'"`                  // 紧急联系人邮箱
	Name       types.String `json:"name,omitempty" db:"'name'"`                    // 紧急联系人姓名
	Phone      types.String `json:"phone,omitempty" db:"'phone'"`                  // 紧急联系人电话
	Relation   types.String `json:"relation,omitempty" db:"'relation'"`            // 与老人关系
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`  // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`      // 创建时间
	ElderId    types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`    // 老人id
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`        // id
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`  // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`  // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`      // 修改时间
	IsReceive  types.Int8   `json:"is_receive,omitempty" db:"'is_receive' size:3"` // 是否接收消息（Y/N）
}

// FamilyAccount 家属账号表
// @tablename family_account
type FamilyAccount struct {
	pool.Model
	Openid     types.String `json:"openid,omitempty" db:"'openid'"`           // 微信 openid（充值支付前置）
	Pass       types.String `json:"pass,omitempty" db:"'pass'"`               // 登录密码（md5）
	Phone      types.String `json:"phone,omitempty" db:"'phone'"`             // 家属手机号（登录账号）
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"` // 创建时间
	Id         types.Money  `json:"id,omitempty" db:"'id' pk auto size:19"`   // 主键
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"` // 更新时间
	State      types.Int8   `json:"state,omitempty" db:"'state' size:3"`      // 逻辑删除
}

// FamilyMember
// @tablename family_member
type FamilyMember struct {
	pool.Model
	Address    types.String `json:"address,omitempty" db:"'address'"`              // 地址
	Email      types.String `json:"email,omitempty" db:"'email'"`                  // 家属邮箱
	IdNum      types.String `json:"id_num,omitempty" db:"'id_num'"`                // 身份证号
	Name       types.String `json:"name,omitempty" db:"'name'"`                    // 家属姓名
	Phone      types.String `json:"phone,omitempty" db:"'phone'"`                  // 家属电话
	Relation   types.String `json:"relation,omitempty" db:"'relation'"`            // 与老人关系
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`  // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`      // 创建时间
	ElderId    types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`    // 老人id
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`        // id
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`  // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`  // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`      // 修改时间
	State      types.Int8   `json:"state,omitempty" db:"'state' size:3"`           // 管理状态：-1=删除，0=禁用，1=可用
	IsReceive  types.Int8   `json:"is_receive,omitempty" db:"'is_receive' size:3"` // 是否接收消息（Y/N）
}

// FamilyRecharge 家属充值订单表
// @tablename family_recharge
type FamilyRecharge struct {
	pool.Model
	OrderNo    types.String `json:"order_no,omitempty" db:"'order_no'"`         // 商户订单号
	Phone      types.String `json:"phone,omitempty" db:"'phone'"`               // 家属手机号
	PrepayId   types.String `json:"prepay_id,omitempty" db:"'prepay_id'"`       // 微信预支付 id
	Amount     types.Money  `json:"amount,omitempty" db:"'amount' size:19"`     // 充值金额（单位：分）
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`   // 创建时间
	ElderId    types.Money  `json:"elder_id,omitempty" db:"'elder_id' size:19"` // 充值到哪位老人
	Id         types.Money  `json:"id,omitempty" db:"'id' pk auto size:19"`     // 主键
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`   // 更新时间
	State      types.Int8   `json:"state,omitempty" db:"'state' size:3"`        // 逻辑删除
	Status     types.Uint8  `json:"status,omitempty" db:"'status' size:3"`      // 0-待支付 1-已支付 2-已关闭
}

// FeeItem 费用项配置表
// @tablename fee_item
type FeeItem struct {
	pool.Model
	Name            types.String `json:"name,omitempty" db:"'name'"`                                // 费用项名称
	CreateId        types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`              // 创建人id
	CreateTime      types.Time   `json:"create_time,omitempty" db:"'create_time'"`                  // 创建时间
	DefaultPrice    types.Money  `json:"default_price,omitempty" db:"'default_price' size:19"`      // 默认单价（分）
	Id              types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`                    // 主键ID
	TenantId        types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`              // 租户ID
	UpdateId        types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`              // 修改人id
	UpdateTime      types.Time   `json:"update_time,omitempty" db:"'update_time'"`                  // 修改时间
	BillingCycle    types.Int8   `json:"billing_cycle,omitempty" db:"'billing_cycle' size:3"`       // 周期性费用计费周期：1-月，2-季，3-年
	CalculationType types.Int8   `json:"calculation_type,omitempty" db:"'calculation_type' size:3"` // 计算方式：1-固定金额，2-按天，3-按用量
	FeeType         types.Int8   `json:"fee_type,omitempty" db:"'fee_type' size:3"`                 // 类型：1-一次性，2-周期性，3-按次
	State           types.Int8   `json:"state,omitempty" db:"'state' size:3"`                       // 管理状态
}

// Floor
// @tablename floor
type Floor struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name'"`                       // 楼层名称
	BuildingId types.BigInt `json:"building_id,omitempty" db:"'building_id' size:20"` // 楼栋id
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`     // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`         // 创建时间
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`           // id
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`     // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`     // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`         // 修改时间
	RoomNum    types.Int32  `json:"room_num,omitempty" db:"'room_num' size:10"`       // 房间数量
	State      types.Int8   `json:"state,omitempty" db:"'state' size:3"`              // 管理状态：-1=删除，0=禁用，1=可用
}

// HealthData 生命体征记录表
// @tablename health_data
type HealthData struct {
	pool.Model
	LeftEar                  types.String  `json:"left_ear,omitempty" db:"'left_ear'"`                                             // 左耳
	RightEar                 types.String  `json:"right_ear,omitempty" db:"'right_ear'"`                                           // 右耳
	CreateId                 types.BigInt  `json:"create_id,omitempty" db:"'create_id' size:20"`                                   // 创建人id
	CreateTime               types.Time    `json:"create_time,omitempty" db:"'create_time'"`                                       // 创建时间
	ElderId                  types.BigInt  `json:"elder_id,omitempty" db:"'elder_id' size:20"`                                     // 老人id
	Id                       types.BigInt  `json:"id,omitempty" db:"'id' pk auto size:20"`                                         // id
	LeftEye                  types.Float64 `json:"left_eye,omitempty" db:"'left_eye' size:22"`                                     // 左眼
	RightEye                 types.Float64 `json:"right_eye,omitempty" db:"'right_eye' size:22"`                                   // 右眼
	Temperature              types.Float64 `json:"temperature,omitempty" db:"'temperature' size:22"`                               // 体温
	TenantId                 types.BigInt  `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`                                   // 租户id
	UpdateId                 types.BigInt  `json:"update_id,omitempty" db:"'update_id' size:20"`                                   // 修改人id
	UpdateTime               types.Time    `json:"update_time,omitempty" db:"'update_time'"`                                       // 修改时间
	Weight                   types.Float64 `json:"weight,omitempty" db:"'weight' size:22"`                                         // 体重
	BloodOxygenSaturation    types.Int32   `json:"blood_oxygen_saturation,omitempty" db:"'blood_oxygen_saturation' size:10"`       // 血氧饱和度
	BodyFatPercentage        types.Int32   `json:"body_fat_percentage,omitempty" db:"'body_fat_percentage' size:10"`               // 体脂率
	Cholesterol              types.Int32   `json:"cholesterol,omitempty" db:"'cholesterol' size:10"`                               // 总胆固醇
	DiastolicBloodPressure   types.Int32   `json:"diastolic_blood_pressure,omitempty" db:"'diastolic_blood_pressure' size:10"`     // 舒张血压
	FastingBloodGlucose      types.Int32   `json:"fasting_blood_glucose,omitempty" db:"'fasting_blood_glucose' size:10"`           // 空腹血糖
	HeartRate                types.Int32   `json:"heart_rate,omitempty" db:"'heart_rate' size:10"`                                 // 心率
	Height                   types.Int32   `json:"height,omitempty" db:"'height' size:10"`                                         // 身高
	HipCircumference         types.Int32   `json:"hip_circumference,omitempty" db:"'hip_circumference' size:10"`                   // 臀围
	MoistureContent          types.Int32   `json:"moisture_content,omitempty" db:"'moisture_content' size:10"`                     // 水分率
	MusclePercentage         types.Int32   `json:"muscle_percentage,omitempty" db:"'muscle_percentage' size:10"`                   // 肌肉率
	PostprandialBloodGlucose types.Int32   `json:"postprandial_blood_glucose,omitempty" db:"'postprandial_blood_glucose' size:10"` // 餐后血糖
	SystolicBloodPressure    types.Int32   `json:"systolic_blood_pressure,omitempty" db:"'systolic_blood_pressure' size:10"`       // 收缩血压
	UricAcid                 types.Int32   `json:"uric_acid,omitempty" db:"'uric_acid' size:10"`                                   // 尿酸
	WaistCircumference       types.Int32   `json:"waist_circumference,omitempty" db:"'waist_circumference' size:10"`               // 腰围
}

// HealthInfo 健康信息表
// @tablename health_info
type HealthInfo struct {
	pool.Model
	AllergyDrug    types.String `json:"allergy_drug,omitempty" db:"'allergy_drug'"`       // 过敏药物
	Doctor         types.String `json:"doctor,omitempty" db:"'doctor'"`                   // 主治医师
	Hearing        types.String `json:"hearing,omitempty" db:"'hearing'"`                 // 听力
	Hospital       types.String `json:"hospital,omitempty" db:"'hospital'"`               // 主治医院
	MajorDisease   types.String `json:"major_disease,omitempty" db:"'major_disease'"`     // 主要疾病
	MedicalHistory types.String `json:"medical_history,omitempty" db:"'medical_history'"` // 病史
	Phone          types.String `json:"phone,omitempty" db:"'phone'"`                     // 医院电话
	SelfCare       types.String `json:"self_care,omitempty" db:"'self_care'"`             // 自理情况
	Vision         types.String `json:"vision,omitempty" db:"'vision'"`                   // 视力
	CreateId       types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`     // 创建人id
	CreateTime     types.Time   `json:"create_time,omitempty" db:"'create_time'"`         // 创建时间
	ElderId        types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`       // 老人id
	Id             types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`           // id
	TenantId       types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`     // 租户id
	UpdateId       types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`     // 修改人id
	UpdateTime     types.Time   `json:"update_time,omitempty" db:"'update_time'"`         // 修改时间
}

// Label
// @tablename label
type Label struct {
	pool.Model
	Color      types.String `json:"color,omitempty" db:"'color'"`                 // 标签颜色
	Name       types.String `json:"name,omitempty" db:"'name'"`                   // 标签名称
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	TypeId     types.BigInt `json:"type_id,omitempty" db:"'type_id' size:20"`     // 类别id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	State      types.Int8   `json:"state,omitempty" db:"'state' size:3"`          // 管理状态：-1=删除，0=禁用，1=可用
}

// LabelType
// @tablename label_type
type LabelType struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name'"`                   // 分类名称
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	State      types.Int8   `json:"state,omitempty" db:"'state' size:3"`          // 管理状态：-1=删除，0=禁用，1=可用
}

// Material
// @tablename material
type Material struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name'"`                   // 物资名称
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	Price      types.Money  `json:"price,omitempty" db:"'price' size:19"`         // 物资单价
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	TypeId     types.BigInt `json:"type_id,omitempty" db:"'type_id' size:20"`     // 物资类别id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	State      types.Int8   `json:"state,omitempty" db:"'state' size:3"`          // 管理状态：-1=删除，0=禁用，1=可用
}

// MaterialType
// @tablename material_type
type MaterialType struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name'"`                   // 物资类别名称
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	Kind       types.Int8   `json:"kind,omitempty" db:"'kind' size:3"`            // 分类用途：1=床型，99=设施/其他
	State      types.Int8   `json:"state,omitempty" db:"'state' size:3"`          // 管理状态：-1=删除，0=禁用，1=可用
}

// Medicine
// @tablename medicine
type Medicine struct {
	pool.Model
	DosageForm    types.String `json:"dosage_form,omitempty" db:"'dosage_form'"`     // 药品剂型
	Manufacturer  types.String `json:"manufacturer,omitempty" db:"'manufacturer'"`   // 生产厂家
	Name          types.String `json:"name,omitempty" db:"'name'"`                   // 药品名称
	Specification types.String `json:"specification,omitempty" db:"'specification'"` // 药品规格
	Type          types.String `json:"type,omitempty" db:"'type'"`                   // 药品类别
	CreateId      types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime    types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	Id            types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	TenantId      types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	UpdateId      types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime    types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	State         types.Int8   `json:"state,omitempty" db:"'state' size:3"`          // 管理状态：-1=删除，0=禁用，1=可用
}

// MedicineRecord
// @tablename medicine_record
type MedicineRecord struct {
	pool.Model
	MedicineTime  types.String `json:"medicine_time,omitempty" db:"'medicine_time'"`             // 用药时间（早/中/晚）
	CreateId      types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`             // 创建人id
	CreateTime    types.Time   `json:"create_time,omitempty" db:"'create_time'"`                 // 创建时间
	DepositInfoId types.BigInt `json:"deposit_info_id,omitempty" db:"'deposit_info_id' size:20"` // 药品缴存信息id
	ElderId       types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`               // 老人id
	Id            types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`                   // id
	MedicineDate  types.Time   `json:"medicine_date,omitempty" db:"'medicine_date'"`             // 用药日期
	TenantId      types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`             // 租户id
	UpdateId      types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`             // 修改人id
	UpdateTime    types.Time   `json:"update_time,omitempty" db:"'update_time'"`                 // 修改时间
}

// MedicineSet
// @tablename medicine_set
type MedicineSet struct {
	pool.Model
	MedicineTime  types.String `json:"medicine_time,omitempty" db:"'medicine_time'"`             // 用药时间（餐前/餐后）
	CreateId      types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`             // 创建人id
	CreateTime    types.Time   `json:"create_time,omitempty" db:"'create_time'"`                 // 创建时间
	DepositInfoId types.BigInt `json:"deposit_info_id,omitempty" db:"'deposit_info_id' size:20"` // 药品缴存信息id
	Id            types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`                   // id
	TenantId      types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`             // 租户id
	UpdateId      types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`             // 修改人id
	UpdateTime    types.Time   `json:"update_time,omitempty" db:"'update_time'"`                 // 修改时间
	DayFrequency  types.Int32  `json:"day_frequency,omitempty" db:"'day_frequency' size:10"`     // 天频率
}

// Member 成员关联表
// @tablename member
type Member struct {
	pool.Model
	Permissions types.String `json:"permissions,omitempty" db:"'permissions'"`     // 预留：细粒度权限
	CreateId    types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime  types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	Id          types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	RoleId      types.BigInt `json:"role_id,omitempty" db:"'role_id' size:20"`     // 角色编号(关联role)
	TenantId    types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	UpdateId    types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime  types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	UserId      types.BigInt `json:"user_id,omitempty" db:"'user_id' size:20"`     // 全局用户id
	State       types.Int8   `json:"state,omitempty" db:"'state' size:3"`          // 管理状态：-1=删除，0=禁用，1=可用
	Status      types.Uint8  `json:"status,omitempty" db:"'status' size:3"`        // 状态：0在职 1离职
}

// Nurse
// @tablename nurse
type Nurse struct {
	pool.Model
	Active         types.String `json:"active,omitempty" db:"'active'"`                          // 活动
	Rest           types.String `json:"rest,omitempty" db:"'rest'"`                              // 休息
	TakeMedicine   types.String `json:"take_medicine,omitempty" db:"'take_medicine'"`            // 服药
	CreateId       types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`            // 创建人id
	CreateTime     types.Time   `json:"create_time,omitempty" db:"'create_time'"`                // 创建时间
	ElderId        types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`              // 老人id
	Id             types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`                  // id
	NurseDate      types.Time   `json:"nurse_date,omitempty" db:"'nurse_date'"`                  // 护理时间
	StaffId        types.BigInt `json:"staff_id,omitempty" db:"'staff_id' size:20"`              // 护理人员id
	TenantId       types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`            // 租户id
	UpdateId       types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`            // 修改人id
	UpdateTime     types.Time   `json:"update_time,omitempty" db:"'update_time'"`                // 修改时间
	CompleteStatus types.Uint8  `json:"complete_status,omitempty" db:"'complete_status' size:3"` // 护理完成情况
	DineStatus     types.Uint8  `json:"dine_status,omitempty" db:"'dine_status' size:3"`         // 进餐情况
}

// NurseGrade
// @tablename nurse_grade
type NurseGrade struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name'"`                       // 级别名称
	Type       types.String `json:"type,omitempty" db:"'type'"`                       // 护理类型
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`     // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`         // 创建时间
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`           // id
	MonthPrice types.Money  `json:"month_price,omitempty" db:"'month_price' size:19"` // 月护理费用
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`     // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`     // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`         // 修改时间
	State      types.Int8   `json:"state,omitempty" db:"'state' size:3"`              // 管理状态：-1=删除，0=禁用，1=可用
}

// NurseGroup
// @tablename nurse_group
type NurseGroup struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name'"`                   // 护工小组名称
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	StaffId    types.BigInt `json:"staff_id,omitempty" db:"'staff_id' size:20"`   // 护工小组组长id
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	State      types.Int8   `json:"state,omitempty" db:"'state' size:3"`          // 管理状态：-1=删除，0=禁用，1=可用
}

// NurseGroupMember
// @tablename nurse_group_member
type NurseGroupMember struct {
	pool.Model
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	GroupId    types.BigInt `json:"group_id,omitempty" db:"'group_id' size:20"`   // 护工小组id
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	StaffId    types.BigInt `json:"staff_id,omitempty" db:"'staff_id' size:20"`   // 护工id
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// NurseItem
// @tablename nurse_item
type NurseItem struct {
	pool.Model
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`   // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`       // 创建时间
	GradeId    types.BigInt `json:"grade_id,omitempty" db:"'grade_id' size:20"`     // 护理等级id
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`         // id
	ServiceId  types.BigInt `json:"service_id,omitempty" db:"'service_id' size:20"` // 服务项目id
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`   // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`   // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`       // 修改时间
}

// NurseReserve
// @tablename nurse_reserve
type NurseReserve struct {
	pool.Model
	ChargeMethod types.String `json:"charge_method,omitempty" db:"'charge_method'"`         // 收费方式
	ServiceName  types.String `json:"service_name,omitempty" db:"'service_name'"`           // 服务项目名称
	CreateId     types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`         // 创建人id
	CreateTime   types.Time   `json:"create_time,omitempty" db:"'create_time'"`             // 创建时间
	ElderId      types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`           // 老人id
	Id           types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`               // id
	NurseDate    types.Time   `json:"nurse_date,omitempty" db:"'nurse_date'"`               // 护理时间
	PayAmount    types.Money  `json:"pay_amount,omitempty" db:"'pay_amount' size:19"`       // 支付总额
	ServicePrice types.Money  `json:"service_price,omitempty" db:"'service_price' size:19"` // 服务费用
	StaffId      types.BigInt `json:"staff_id,omitempty" db:"'staff_id' size:20"`           // 服务人id
	TenantId     types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`         // 租户id
	UpdateId     types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`         // 修改人id
	UpdateTime   types.Time   `json:"update_time,omitempty" db:"'update_time'"`             // 修改时间
	Frequency    types.Int32  `json:"frequency,omitempty" db:"'frequency' size:10"`         // 服务次数
	NeedDate     types.Int32  `json:"need_date,omitempty" db:"'need_date' size:10"`         // 所需时间
	Status       types.Uint8  `json:"status,omitempty" db:"'status' size:3"`                // 订单状态
}

// Order
// @tablename order
type Order struct {
	pool.Model
	DineType          types.String `json:"dine_type,omitempty" db:"'dine_type'"`                     // 就餐方式
	OutTradeNo        types.String `json:"out_trade_no,omitempty" db:"'out_trade_no'"`               // 对外交易单号
	CreateId          types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`             // 创建人id
	CreateTime        types.Time   `json:"create_time,omitempty" db:"'create_time'"`                 // 创建时间
	DeliverDishesDate types.Time   `json:"deliver_dishes_date,omitempty" db:"'deliver_dishes_date'"` // 送餐时间
	DineDate          types.Time   `json:"dine_date,omitempty" db:"'dine_date'"`                     // 就餐时间
	ElderId           types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`               // 老人id
	Id                types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`                   // id
	PayAmount         types.Money  `json:"pay_amount,omitempty" db:"'pay_amount' size:19"`           // 支付总额
	StaffId           types.BigInt `json:"staff_id,omitempty" db:"'staff_id' size:20"`               // 送餐人id
	TenantId          types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`             // 租户id
	UpdateId          types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`             // 修改人id
	UpdateTime        types.Time   `json:"update_time,omitempty" db:"'update_time'"`                 // 修改时间
	Status            types.Uint8  `json:"status,omitempty" db:"'status' size:3"`                    // 订单状态
}

// OrderDishes
// @tablename order_dishes
type OrderDishes struct {
	pool.Model
	DishesName   types.String `json:"dishes_name,omitempty" db:"'dishes_name'"`             // 菜品名称
	CreateId     types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`         // 创建人id
	CreateTime   types.Time   `json:"create_time,omitempty" db:"'create_time'"`             // 创建时间
	DishesPrice  types.Money  `json:"dishes_price,omitempty" db:"'dishes_price' size:19"`   // 菜品价格
	Id           types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`               // id
	OrderId      types.BigInt `json:"order_id,omitempty" db:"'order_id' size:20"`           // 订餐id
	ReallyAmount types.Money  `json:"really_amount,omitempty" db:"'really_amount' size:19"` // 实际总额
	TenantId     types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`         // 租户id
	TotalAmount  types.Money  `json:"total_amount,omitempty" db:"'total_amount' size:19"`   // 菜品总额
	UpdateId     types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`         // 修改人id
	UpdateTime   types.Time   `json:"update_time,omitempty" db:"'update_time'"`             // 修改时间
	OrderNum     types.Int32  `json:"order_num,omitempty" db:"'order_num' size:10"`         // 菜品份数
	Status       types.Uint8  `json:"status,omitempty" db:"'status' size:3"`                // 套餐标记
}

// OutboundRecordItem 出库单明细
// @tablename outbound_record_item
type OutboundRecordItem struct {
	pool.Model
	CreateId   types.BigInt  `json:"create_id,omitempty" db:"'create_id' size:20"`     // 创建人id
	CreateTime types.Time    `json:"create_time,omitempty" db:"'create_time'"`         // 创建时间
	Id         types.BigInt  `json:"id,omitempty" db:"'id' pk auto size:20"`           // id
	MaterialId types.BigInt  `json:"material_id,omitempty" db:"'material_id' size:20"` // 物资id
	RecordId   types.BigInt  `json:"record_id,omitempty" db:"'record_id' size:20"`     // 出库单id
	StockId    *types.BigInt `json:"stock_id,omitempty" db:"'stock_id' size:20"`       // 库存台账id(审核通过后回填)
	TenantId   types.BigInt  `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`     // 租户id
	UpdateId   types.BigInt  `json:"update_id,omitempty" db:"'update_id' size:20"`     // 修改人id
	UpdateTime types.Time    `json:"update_time,omitempty" db:"'update_time'"`         // 修改时间
	Qty        types.Int32   `json:"qty,omitempty" db:"'qty' size:10"`                 // 本次出库数量
}

// OutboundRecord 出库记录
// @tablename outbound_record
type OutboundRecord struct {
	pool.Model
	MaterialUse   types.String `json:"material_use,omitempty" db:"'material_use'"`         // 物资去向
	RecipientType types.String `json:"recipient_type,omitempty" db:"'recipient_type'"`     // 领用人类型
	CreateId      types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`       // 创建人id
	CreateTime    types.Time   `json:"create_time,omitempty" db:"'create_time'"`           // 创建时间
	Id            types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`             // id
	OutboundDate  types.Time   `json:"outbound_date,omitempty" db:"'outbound_date'"`       // 出库时间
	RecipientId   types.BigInt `json:"recipient_id,omitempty" db:"'recipient_id' size:20"` // 领用人id
	StaffId       types.BigInt `json:"staff_id,omitempty" db:"'staff_id' size:20"`         // 经办人id
	TenantId      types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`       // 租户id
	UpdateId      types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`       // 修改人id
	UpdateTime    types.Time   `json:"update_time,omitempty" db:"'update_time'"`           // 修改时间
	WarehouseId   types.BigInt `json:"warehouse_id,omitempty" db:"'warehouse_id' size:20"` // 仓库id
	State         types.Int8   `json:"state,omitempty" db:"'state' size:3"`                // 管理状态：-1=删除，0=禁用，1=可用
	Status        types.Uint8  `json:"status,omitempty" db:"'status' size:3"`              // 出库状态
}

// Outward 外出登记
// @tablename outward
type Outward struct {
	pool.Model
	ChaperoneName  types.String `json:"chaperone_name,omitempty" db:"'chaperone_name'"`     // 陪同人姓名
	ChaperonePhone types.String `json:"chaperone_phone,omitempty" db:"'chaperone_phone'"`   // 陪同人电话
	ChaperoneType  types.String `json:"chaperone_type,omitempty" db:"'chaperone_type'"`     // 陪同人类型（家属/护工）
	CreateId       types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`       // 创建人id
	CreateTime     types.Time   `json:"create_time,omitempty" db:"'create_time'"`           // 创建时间
	ElderId        types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`         // 老人id
	Id             types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`             // id
	OutwardDate    types.Time   `json:"outward_date,omitempty" db:"'outward_date'"`         // 外出时间
	PlanReturnDate types.Time   `json:"plan_return_date,omitempty" db:"'plan_return_date'"` // 计划返回时间
	RealReturnDate types.Time   `json:"real_return_date,omitempty" db:"'real_return_date'"` // 实际返回时间
	TenantId       types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`       // 租户id
	UpdateId       types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`       // 修改人id
	UpdateTime     types.Time   `json:"update_time,omitempty" db:"'update_time'"`           // 修改时间
	State          types.Int8   `json:"state,omitempty" db:"'state' size:3"`                // 管理状态：-1=删除，0=禁用，1=可用
}

// Reserve
// @tablename reserve
type Reserve struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name'"`                   // 交款人姓名
	Phone      types.String `json:"phone,omitempty" db:"'phone'"`                 // 交款人电话
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	Deposit    types.Money  `json:"deposit,omitempty" db:"'deposit' size:19"`     // 定金
	DueDate    types.Time   `json:"due_date,omitempty" db:"'due_date'"`           // 预定到期时间
	ElderId    types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`   // 老人id
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	StaffId    types.BigInt `json:"staff_id,omitempty" db:"'staff_id' size:20"`   // 销售人员id
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	Status     types.Uint8  `json:"status,omitempty" db:"'status' size:3"`        // 退款状态（N/Y）
}

// Retreat
// @tablename retreat
type Retreat struct {
	pool.Model
	RetreatCause types.String `json:"retreat_cause,omitempty" db:"'retreat_cause'"` // 退住原因
	RetreatForm  types.String `json:"retreat_form,omitempty" db:"'retreat_form'"`   // 退住形式
	CreateId     types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime   types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	ElderId      types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`   // 老人id
	Id           types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	TenantId     types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	UpdateId     types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime   types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	Evaluate     types.Int32  `json:"evaluate,omitempty" db:"'evaluate' size:10"`   // 对老人评价
}

// RetreatApply
// @tablename retreat_apply
type RetreatApply struct {
	pool.Model
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	ElderId    types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`   // 老人id
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time' <-"`  // 修改时间
	State      types.Int8   `json:"state,omitempty" db:"'state' size:3"`          // 管理状态：-1=删除，0=禁用，1=正常
	Status     types.Uint8  `json:"status,omitempty" db:"'status' size:3"`        // 业务状态(退住申请流程状态)
}

// Role 角色表
// @tablename role
type Role struct {
	pool.Model
	Name        types.String `json:"name,omitempty" db:"'name'"`                   // 角色名称
	Permissions types.String `json:"permissions,omitempty" db:"'permissions'"`     // 权限列表
	CreateId    types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime  types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	Id          types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	TenantId    types.Money  `json:"tenant_id,omitempty" db:"'tenant_id' size:19"` // 租户ID（平台角色tenant_id=0）
	UpdateId    types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime  types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	IsSystem    types.Int8   `json:"is_system,omitempty" db:"'is_system' size:3"`  // 是否系统预置：0-否，1-是
}

// RoleAuth
// @tablename role_auth
type RoleAuth struct {
	pool.Model
	AuthId     types.BigInt `json:"auth_id,omitempty" db:"'auth_id' size:20"`     // 权限id
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	RoleId     types.BigInt `json:"role_id,omitempty" db:"'role_id' size:20"`     // 角色id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// Room
// @tablename room
type Room struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name'"`                   // 房间名称
	BuildId    types.BigInt `json:"build_id,omitempty" db:"'build_id' size:20"`   // 楼栋id
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	FloorId    types.BigInt `json:"floor_id,omitempty" db:"'floor_id' size:20"`   // 楼层id
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	Price      types.Money  `json:"price,omitempty" db:"'price' size:19"`         // 房间基准价（分）
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	TypeId     types.BigInt `json:"type_id,omitempty" db:"'type_id' size:20"`     // 房间类型id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	BedNum     types.Int32  `json:"bed_num,omitempty" db:"'bed_num' size:10"`     // 床位数量
	State      types.Int8   `json:"state,omitempty" db:"'state' size:3"`          // 管理状态：-1=删除，0=禁用，1=可用
	Status     types.Uint8  `json:"status,omitempty" db:"'status' size:3"`        // 状态：0-空闲，1-部分占用，2-全满，3-维修
	Type       types.Int8   `json:"type,omitempty" db:"'type' size:3"`            // 房间类型：1-单人间，2-双人间，3-多人间
}

// RoomMaterial
// @tablename room_material
type RoomMaterial struct {
	pool.Model
	CreateId       types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`               // 创建人id
	CreateTime     types.Time   `json:"create_time,omitempty" db:"'create_time'"`                   // 创建时间
	Id             types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`                     // id
	MaterialTypeId types.BigInt `json:"material_type_id,omitempty" db:"'material_type_id' size:20"` // 设施(物资分类)编号，kind=0
	RoomId         types.BigInt `json:"room_id,omitempty" db:"'room_id' size:20"`                   // 房间id
	TenantId       types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`               // 租户id
	UpdateId       types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`               // 修改人id
	UpdateTime     types.Time   `json:"update_time,omitempty" db:"'update_time'"`                   // 修改时间
	State          types.Int8   `json:"state,omitempty" db:"'state' size:3"`                        // 删除状态(Y/N)
}

// RoomTransfer 转房记录表
// @tablename room_transfer
type RoomTransfer struct {
	pool.Model
	Reason       types.String `json:"reason,omitempty" db:"'reason'"`                   // 原因
	CreateId     types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`     // 创建人id
	CreateTime   types.Time   `json:"create_time,omitempty" db:"'create_time'"`         // 创建时间
	ElderId      types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`       // 老人ID
	FromBedId    types.BigInt `json:"from_bed_id,omitempty" db:"'from_bed_id' size:20"` // 原床位ID
	Id           types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`           // 主键ID
	TenantId     types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`     // 租户ID
	ToBedId      types.BigInt `json:"to_bed_id,omitempty" db:"'to_bed_id' size:20"`     // 新床位ID
	TransferDate types.Time   `json:"transfer_date,omitempty" db:"'transfer_date'"`     // 转房日期
	UpdateId     types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`     // 修改人id
	UpdateTime   types.Time   `json:"update_time,omitempty" db:"'update_time'"`         // 修改时间
}

// RoomType
// @tablename room_type
type RoomType struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name'"`                       // 房间类型名称
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`     // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`         // 创建时间
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`           // id
	MonthPrice types.Money  `json:"month_price,omitempty" db:"'month_price' size:19"` // 月房间费用
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`     // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`     // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`         // 修改时间
	State      types.Int8   `json:"state,omitempty" db:"'state' size:3"`              // 管理状态：-1=删除，0=禁用，1=可用
}

// ServiceItem
// @tablename service_item
type ServiceItem struct {
	pool.Model
	ChargeMethod types.String `json:"charge_method,omitempty" db:"'charge_method'"` // 收费方式
	Name         types.String `json:"name,omitempty" db:"'name'"`                   // 服务名称
	CreateId     types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime   types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	Id           types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	Price        types.Money  `json:"price,omitempty" db:"'price' size:19"`         // 服务费用
	TenantId     types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	TypeId       types.BigInt `json:"type_id,omitempty" db:"'type_id' size:20"`     // 服务项目类别id
	UpdateId     types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime   types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	NeedDate     types.Int32  `json:"need_date,omitempty" db:"'need_date' size:10"` // 所需时间(分)
	State        types.Int8   `json:"state,omitempty" db:"'state' size:3"`          // 管理状态：-1=删除，0=禁用，1=可用
}

// ServiceType
// @tablename service_type
type ServiceType struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name'"`                   // 服务项目名称
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	State      types.Int8   `json:"state,omitempty" db:"'state' size:3"`          // 管理状态：-1=删除，0=禁用，1=可用
}

// SetDishes
// @tablename set_dishes
type SetDishes struct {
	pool.Model
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	DishesId   types.BigInt `json:"dishes_id,omitempty" db:"'dishes_id' size:20"` // 菜品食物id
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	SetId      types.BigInt `json:"set_id,omitempty" db:"'set_id' size:20"`       // 餐饮套餐id
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// Source
// @tablename source
type Source struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name'"`                   // 来源渠道名称
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	State      types.Int8   `json:"state,omitempty" db:"'state' size:3"`          // 管理状态：-1=删除，0=禁用，1=可用
}

// Staff
// @tablename staff
type Staff struct {
	pool.Model
	Address    types.String `json:"address,omitempty" db:"'address'"`             // 地址
	Avatar     types.String `json:"avatar,omitempty" db:"'avatar'"`               // 头像
	Email      types.String `json:"email,omitempty" db:"'email'"`                 // 邮箱
	IdNum      types.String `json:"id_num,omitempty" db:"'id_num'"`               // 身份证号
	Name       types.String `json:"name,omitempty" db:"'name'"`                   // 姓名
	Pass       types.String `json:"pass,omitempty" db:"'pass'"`                   // 密码
	Phone      types.String `json:"phone,omitempty" db:"'phone'"`                 // 电话
	Sex        types.String `json:"sex,omitempty" db:"'sex'"`                     // 性别(男/女)
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	RoleId     types.BigInt `json:"role_id,omitempty" db:"'role_id' size:20"`     // 角色id
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	Age        types.Int32  `json:"age,omitempty" db:"'age' size:10"`             // 年龄
	Status     types.Uint8  `json:"status,omitempty" db:"'status' size:3"`        // 状态：0-离职，1-在职
}

// Tenant 租户表
// @tablename tenant
type Tenant struct {
	pool.Model
	ContactName  types.String `json:"contact_name,omitempty" db:"'contact_name'"`   // 联系人姓名
	ContactPhone types.String `json:"contact_phone,omitempty" db:"'contact_phone'"` // 联系电话
	Logo         types.String `json:"logo,omitempty" db:"'logo'"`                   // 企业logo
	Name         types.String `json:"name,omitempty" db:"'name'"`                   // 企业名称
	Plan         types.String `json:"plan,omitempty" db:"'plan'"`                   // 套餐
	CreateId     types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime   types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	ExpireTime   types.Time   `json:"expire_time,omitempty" db:"'expire_time'"`     // 套餐到期时间
	Id           types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	UpdateId     types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime   types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	State        types.Int8   `json:"state,omitempty" db:"'state' size:3"`          // 管理状态：-1=删除，0=禁用，1=可用
	Status       types.Uint8  `json:"status,omitempty" db:"'status' size:3"`        // 状态：0试用中 1正式 2锁定
}

// User 全局用户表
// @tablename user
type User struct {
	pool.Model
	Avatar     types.String `json:"avatar,omitempty" db:"'avatar'"`           // 头像
	Name       types.String `json:"name,omitempty" db:"'name'"`               // 姓名
	Openid     types.String `json:"openid,omitempty" db:"'openid'"`           // 微信OpenID（兜底匹配）
	Pass       types.String `json:"pass,omitempty" db:"'pass'"`               // 密码(md5)
	Phone      types.String `json:"phone,omitempty" db:"'phone'"`             // 手机号（账号密码登录）
	UnionId    types.String `json:"union_id,omitempty" db:"'union_id'"`       // 微信UnionID（全局唯一）
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"` // 创建时间
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`   // id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"` // 更新时间
	State      types.Int8   `json:"state,omitempty" db:"'state' size:3"`      // 管理状态：-1=删除，0=禁用，1=可用
}

// Visit 访客记录表
// @tablename visit
type Visit struct {
	pool.Model
	IdCard     types.String `json:"id_card,omitempty" db:"'id_card'"`             // 身份证号（加密）
	Name       types.String `json:"name,omitempty" db:"'name'"`                   // 来访者姓名
	Phone      types.String `json:"phone,omitempty" db:"'phone'"`                 // 来访者电话
	Relation   types.String `json:"relation,omitempty" db:"'relation'"`           // 与老人关系
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	ElderId    types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`   // 老人id
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	LeaveTime  types.Time   `json:"leave_time,omitempty" db:"'leave_time'"`       // 离开时间
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	VisitTime  types.Time   `json:"visit_time,omitempty" db:"'visit_time'"`       // 来访时间
	VisitNum   types.Int32  `json:"visit_num,omitempty" db:"'visit_num' size:10"` // 来访者人数
	State      types.Int8   `json:"state,omitempty" db:"'state' size:3"`          // 删除状态
	Status     types.Uint8  `json:"status,omitempty" db:"'status' size:3"`        // 来访状态
}

// VisitPlan
// @tablename visit_plan
type VisitPlan struct {
	pool.Model
	Content      types.String `json:"content,omitempty" db:"'content'"`             // 回访计划内容
	Title        types.String `json:"title,omitempty" db:"'title'"`                 // 回访计划标题
	CompleteDate types.Time   `json:"complete_date,omitempty" db:"'complete_date'"` // 计划完成时间
	CreateId     types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime   types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	ElderId      types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`   // 老人id
	Id           types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	PlanDate     types.Time   `json:"plan_date,omitempty" db:"'plan_date'"`         // 计划回访时间
	TenantId     types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	UpdateId     types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime   types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	State        types.Int8   `json:"state,omitempty" db:"'state' size:3"`          // 管理状态：-1=删除，0=禁用，1=可用
}

// Warehouse 仓库
// @tablename warehouse
type Warehouse struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name'"`                   // 仓库名称
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人id
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // id
	StaffId    types.BigInt `json:"staff_id,omitempty" db:"'staff_id' size:20"`   // 仓库管理员id
	TenantId   types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"` // 租户id
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人id
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	State      types.Int8   `json:"state,omitempty" db:"'state' size:3"`          // 管理状态：-1=删除，0=禁用，1=可用
}

// Stock 库存台账
// @tablename stock
type Stock struct {
	pool.Model
	CreateId    types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`       // 创建人id
	CreateTime  types.Time   `json:"create_time,omitempty" db:"'create_time'"`           // 创建时间
	ExpireDate  types.Time   `json:"expire_date,omitempty" db:"'expire_date'"`           // 有效期
	Id          types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`             // id
	MaterialId  types.BigInt `json:"material_id,omitempty" db:"'material_id' size:20"`   // 物资id
	ProductDate types.Time   `json:"product_date,omitempty" db:"'product_date'"`         // 生产日期
	TenantId    types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`       // 租户id
	UpdateId    types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`       // 修改人id
	UpdateTime  types.Time   `json:"update_time,omitempty" db:"'update_time'"`           // 修改时间
	WarehouseId types.BigInt `json:"warehouse_id,omitempty" db:"'warehouse_id' size:20"` // 仓库id
	Qty         types.Int32  `json:"qty,omitempty" db:"'qty' size:10"`                   // 实时库存(>=0)
	BatchNo     types.String `json:"batch_no,omitempty" db:"'batch_no'"`                 // 批次号(空表示不按批次管理)
}

// WarehouseRecordItem 入库单明细
// @tablename warehouse_record_item
type WarehouseRecordItem struct {
	pool.Model
	CreateId    types.BigInt  `json:"create_id,omitempty" db:"'create_id' size:20"`     // 创建人id
	CreateTime  types.Time    `json:"create_time,omitempty" db:"'create_time'"`         // 创建时间
	ExpireDate  types.Time    `json:"expire_date,omitempty" db:"'expire_date'"`         // 有效期
	Id          types.BigInt  `json:"id,omitempty" db:"'id' pk auto size:20"`           // id
	MaterialId  types.BigInt  `json:"material_id,omitempty" db:"'material_id' size:20"` // 物资id
	ProductDate types.Time    `json:"product_date,omitempty" db:"'product_date'"`       // 生产日期
	RecordId    types.BigInt  `json:"record_id,omitempty" db:"'record_id' size:20"`     // 入库单id
	StockId     *types.BigInt `json:"stock_id,omitempty" db:"'stock_id' size:20"`       // 库存台账id(审核通过后回填)
	TenantId    types.BigInt  `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`     // 租户id
	UpdateId    types.BigInt  `json:"update_id,omitempty" db:"'update_id' size:20"`     // 修改人id
	UpdateTime  types.Time    `json:"update_time,omitempty" db:"'update_time'"`         // 修改时间
	Qty         types.Int32   `json:"qty,omitempty" db:"'qty' size:10"`                 // 本次入库数量
}

// WarehouseRecord 库存记录
// @tablename warehouse_record
type WarehouseRecord struct {
	pool.Model
	Source        types.String `json:"source,omitempty" db:"'source'"`                     // 物资来源
	CreateId      types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`       // 创建人id
	CreateTime    types.Time   `json:"create_time,omitempty" db:"'create_time'"`           // 创建时间
	Id            types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`             // id
	StaffId       types.BigInt `json:"staff_id,omitempty" db:"'staff_id' size:20"`         // 经办人id
	TenantId      types.BigInt `json:"tenant_id,omitempty" db:"'tenant_id' size:20"`       // 租户id
	UpdateId      types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`       // 修改人id
	UpdateTime    types.Time   `json:"update_time,omitempty" db:"'update_time'"`           // 修改时间
	WarehouseDate types.Time   `json:"warehouse_date,omitempty" db:"'warehouse_date'"`     // 入库时间
	WarehouseId   types.BigInt `json:"warehouse_id,omitempty" db:"'warehouse_id' size:20"` // 仓库id
	State         types.Int8   `json:"state,omitempty" db:"'state' size:3"`                // 管理状态：-1=删除，0=禁用，1=可用
	Status        types.Uint8  `json:"status,omitempty" db:"'status' size:3"`              // 入库状态
}
