package do

import (
	"github.com/linbaozhong/gentity/pkg/ace/pool"
	"github.com/linbaozhong/gentity/pkg/types"
)

// ActiveParticipant
// @tablename active_participant
type ActiveParticipant struct {
	pool.Model
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // 编号
	ActiveId   types.BigInt `json:"active_id,omitempty" db:"'active_id' size:20"` // 活动编号
	ElderId    types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`   // 老人编号
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// Consume
// @tablename consume
type Consume struct {
	pool.Model
	ConsumeType   types.String  `json:"consume_type,omitempty" db:"'consume_type' size:10"`       // 消费类别
	Id            types.BigInt  `json:"id,omitempty" db:"'id' pk auto size:20"`                   // 编号
	ElderId       types.BigInt  `json:"elder_id,omitempty" db:"'elder_id' size:20"`               // 老人编号
	ConsumeAmount types.Float64 `json:"consume_amount,omitempty" db:"'consume_amount' size:10|2"` // 消费金额
	ConsumeDate   types.Time    `json:"consume_date,omitempty" db:"'consume_date'"`               // 消费日期
	CreateId      types.BigInt  `json:"create_id,omitempty" db:"'create_id' size:20"`             // 创建人编号
	CreateTime    types.Time    `json:"create_time,omitempty" db:"'create_time'"`                 // 创建时间
	UpdateId      types.BigInt  `json:"update_id,omitempty" db:"'update_id' size:20"`             // 修改人编号
	UpdateTime    types.Time    `json:"update_time,omitempty" db:"'update_time'"`                 // 修改时间
}

// Contract
// @tablename contract
type Contract struct {
	pool.Model
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // 编号
	ElderId    types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`   // 老人编号
	StaffId    types.BigInt `json:"staff_id,omitempty" db:"'staff_id' size:20"`   // 销售人员编号
	SignDate   types.Time   `json:"sign_date,omitempty" db:"'sign_date'"`         // 合同签订日期
	StartDate  types.Time   `json:"start_date,omitempty" db:"'start_date'"`       // 合同开始日期
	EndDate    types.Time   `json:"end_date,omitempty" db:"'end_date'"`           // 合同结束日期
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// RoleAuth
// @tablename role_auth
type RoleAuth struct {
	pool.Model
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // 编号
	RoleId     types.BigInt `json:"role_id,omitempty" db:"'role_id' size:20"`     // 角色编号
	AuthId     types.BigInt `json:"auth_id,omitempty" db:"'auth_id' size:20"`     // 权限编号
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// User
// @tablename user
type User struct {
	pool.Model
	Id       types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`     // 编号
	Name     types.String `json:"name,omitempty" db:"'name' size:40"`         // 姓名
	Email    types.String `json:"email,omitempty" db:"'email' size:40"`       // 邮箱
	Mobile   types.String `json:"mobile,omitempty" db:"'mobile' size:20"`     // 手机号
	Gender   types.String `json:"gender,omitempty" db:"'gender' size:5"`      // 性别
	Birthday types.String `json:"birthday,omitempty" db:"'birthday' size:20"` // 生日
	Creator  types.BigInt `json:"creator,omitempty" db:"'creator' size:20"`   // 创建人
	Status   types.Int8   `json:"status,omitempty" db:"'status' size:4"`      // 状态
	State    types.Int8   `json:"state,omitempty" db:"'state' size:4"`        // 启用状态
	Ctime    types.Time   `json:"ctime,omitempty" db:"'ctime'"`               // 创建时间
	Utime    types.Time   `json:"utime,omitempty" db:"'utime'"`               // 修改时间
}

// SetDishes
// @tablename set_dishes
type SetDishes struct {
	pool.Model
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // 编号
	SetId      types.BigInt `json:"set_id,omitempty" db:"'set_id' size:20"`       // 餐饮套餐编号
	DishesId   types.BigInt `json:"dishes_id,omitempty" db:"'dishes_id' size:20"` // 菜品食物编号
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// Bed
// @tablename bed
type Bed struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name' size:40"`           // 床位名称
	BedFlag    types.String `json:"bed_flag,omitempty" db:"'bed_flag' size:5"`    // 床位状态(空闲/预定/入住/退住审核)
	DelFlag    types.String `json:"del_flag,omitempty" db:"'del_flag' size:2"`    // 删除状态（Y/N）
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // 编号
	RoomId     types.BigInt `json:"room_id,omitempty" db:"'room_id' size:20"`     // 房间编号
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// Consult
// @tablename consult
type Consult struct {
	pool.Model
	Name           types.String `json:"name,omitempty" db:"'name' size:10"`                        // 咨询人姓名
	Phone          types.String `json:"phone,omitempty" db:"'phone' size:11"`                      // 咨询人电话
	Relation       types.String `json:"relation,omitempty" db:"'relation' size:5"`                 // 与老人关系
	ConsultContent types.String `json:"consult_content,omitempty" db:"'consult_content' size:255"` // 咨询内容
	Id             types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`                    // 编号
	ElderId        types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`                // 老人编号
	SourceId       types.BigInt `json:"source_id,omitempty" db:"'source_id' size:20"`              // 来源渠道编号
	StaffId        types.BigInt `json:"staff_id,omitempty" db:"'staff_id' size:20"`                // 接待人编号
	ConsultDate    types.Time   `json:"consult_date,omitempty" db:"'consult_date'"`                // 咨询日期
	CreateId       types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`              // 创建人编号
	CreateTime     types.Time   `json:"create_time,omitempty" db:"'create_time'"`                  // 创建时间
	UpdateId       types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`              // 修改人编号
	UpdateTime     types.Time   `json:"update_time,omitempty" db:"'update_time'"`                  // 修改时间
}

// HealthInfo
// @tablename health_info
type HealthInfo struct {
	pool.Model
	SelfCare       types.String `json:"self_care,omitempty" db:"'self_care' size:5"`               // 自理情况
	Vision         types.String `json:"vision,omitempty" db:"'vision' size:5"`                     // 视力
	Hearing        types.String `json:"hearing,omitempty" db:"'hearing' size:10"`                  // 听力
	Hospital       types.String `json:"hospital,omitempty" db:"'hospital' size:50"`                // 主治医院
	Doctor         types.String `json:"doctor,omitempty" db:"'doctor' size:10"`                    // 主治医师
	Phone          types.String `json:"phone,omitempty" db:"'phone' size:11"`                      // 医院电话
	AllergyDrug    types.String `json:"allergy_drug,omitempty" db:"'allergy_drug' size:255"`       // 过敏药物
	MedicalHistory types.String `json:"medical_history,omitempty" db:"'medical_history' size:255"` // 病史
	MajorDisease   types.String `json:"major_disease,omitempty" db:"'major_disease' size:255"`     // 主要疾病
	Id             types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`                    // 编号
	ElderId        types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`                // 老人编号
	CreateId       types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`              // 创建人编号
	CreateTime     types.Time   `json:"create_time,omitempty" db:"'create_time'"`                  // 创建时间
	UpdateId       types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`              // 修改人编号
	UpdateTime     types.Time   `json:"update_time,omitempty" db:"'update_time'"`                  // 修改时间
}

// MaterialType
// @tablename material_type
type MaterialType struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name' size:10"`           // 物资类别名称
	DelFlag    types.String `json:"del_flag,omitempty" db:"'del_flag' size:2"`    // 删除状态（Y/N）
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // 编号
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// RetreatApply
// @tablename retreat_apply
type RetreatApply struct {
	pool.Model
	ApplyFlag  types.String `json:"apply_flag,omitempty" db:"'apply_flag' size:5"` // 申请状态
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`        // 编号
	ElderId    types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`    // 老人编号
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`  // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`      // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`  // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`      // 修改时间
}

// DepositInfo
// @tablename deposit_info
type DepositInfo struct {
	pool.Model
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`           // 编号
	DepositId  types.BigInt `json:"deposit_id,omitempty" db:"'deposit_id' size:20"`   // 药品缴存编号
	MedicineId types.BigInt `json:"medicine_id,omitempty" db:"'medicine_id' size:20"` // 药品编号
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`     // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`         // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`     // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`         // 修改时间
	DepositNum types.Int32  `json:"deposit_num,omitempty" db:"'deposit_num' size:11"` // 缴存数量
	SurplusNum types.Int32  `json:"surplus_num,omitempty" db:"'surplus_num' size:11"` // 剩余数量
	WarnNum    types.Int32  `json:"warn_num,omitempty" db:"'warn_num' size:11"`       // 预警数量
}

// OutboundMaterial
// @tablename outbound_material
type OutboundMaterial struct {
	pool.Model
	Id                  types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`                               // 编号
	OutboundRecordId    types.BigInt `json:"outbound_record_id,omitempty" db:"'outbound_record_id' size:20"`       // 出库登记编号
	WarehouseMaterialId types.BigInt `json:"warehouse_material_id,omitempty" db:"'warehouse_material_id' size:20"` // 入库物资编号
	MaterialId          types.BigInt `json:"material_id,omitempty" db:"'material_id' size:20"`                     // 物资编号
	CreateId            types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`                         // 创建人编号
	CreateTime          types.Time   `json:"create_time,omitempty" db:"'create_time'"`                             // 创建时间
	UpdateId            types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`                         // 修改人编号
	UpdateTime          types.Time   `json:"update_time,omitempty" db:"'update_time'"`                             // 修改时间
	OutboundNum         types.Int32  `json:"outbound_num,omitempty" db:"'outbound_num' size:11"`                   // 出库数量
}

// Visit
// @tablename visit
type Visit struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name' size:10"`            // 来访者姓名
	Phone      types.String `json:"phone,omitempty" db:"'phone' size:11"`          // 来访者电话
	Relation   types.String `json:"relation,omitempty" db:"'relation' size:5"`     // 与老人关系
	VisitFlag  types.String `json:"visit_flag,omitempty" db:"'visit_flag' size:5"` // 来访状态
	DelFlag    types.String `json:"del_flag,omitempty" db:"'del_flag' size:2"`     // 删除状态
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`        // 编号
	ElderId    types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`    // 老人编号
	VisitDate  types.Time   `json:"visit_date,omitempty" db:"'visit_date'"`        // 来访时间
	LeaveDate  types.Time   `json:"leave_date,omitempty" db:"'leave_date'"`        // 离开时间
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`  // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`      // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`  // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`      // 修改时间
	VisitNum   types.Int32  `json:"visit_num,omitempty" db:"'visit_num' size:11"`  // 来访者人数
}

// WarehouseMaterial
// @tablename warehouse_material
type WarehouseMaterial struct {
	pool.Model
	Id                types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`                           // 编号
	WarehouseRecordId types.BigInt `json:"warehouse_record_id,omitempty" db:"'warehouse_record_id' size:20"` // 入库登记编号
	MaterialId        types.BigInt `json:"material_id,omitempty" db:"'material_id' size:20"`                 // 物资编号
	ProductDate       types.Time   `json:"product_date,omitempty" db:"'product_date'"`                       // 生产日期
	ExpireDate        types.Time   `json:"expire_date,omitempty" db:"'expire_date'"`                         // 有效期
	CreateId          types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`                     // 创建人编号
	CreateTime        types.Time   `json:"create_time,omitempty" db:"'create_time'"`                         // 创建时间
	UpdateId          types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`                     // 修改人编号
	UpdateTime        types.Time   `json:"update_time,omitempty" db:"'update_time'"`                         // 修改时间
	WarehouseNum      types.Int32  `json:"warehouse_num,omitempty" db:"'warehouse_num' size:11"`             // 入库数量
	Inventory         types.Int32  `json:"inventory,omitempty" db:"'inventory' size:11"`                     // 库存量
}

// LabelType
// @tablename label_type
type LabelType struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name' size:10"`           // 分类名称
	DelFlag    types.String `json:"del_flag,omitempty" db:"'del_flag' size:2"`    // 删除状态（Y/N）
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // 编号
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// NurseGrade
// @tablename nurse_grade
type NurseGrade struct {
	pool.Model
	Name       types.String  `json:"name,omitempty" db:"'name' size:10"`                 // 级别名称
	Type       types.String  `json:"type,omitempty" db:"'type' size:5"`                  // 护理类型
	DelFlag    types.String  `json:"del_flag,omitempty" db:"'del_flag' size:2"`          // 删除状态（Y/N）
	Id         types.BigInt  `json:"id,omitempty" db:"'id' pk auto size:20"`             // 编号
	MonthPrice types.Float64 `json:"month_price,omitempty" db:"'month_price' size:10|2"` // 月护理费用
	CreateId   types.BigInt  `json:"create_id,omitempty" db:"'create_id' size:20"`       // 创建人编号
	CreateTime types.Time    `json:"create_time,omitempty" db:"'create_time'"`           // 创建时间
	UpdateId   types.BigInt  `json:"update_id,omitempty" db:"'update_id' size:20"`       // 修改人编号
	UpdateTime types.Time    `json:"update_time,omitempty" db:"'update_time'"`           // 修改时间
}

// ActiveType
// @tablename active_type
type ActiveType struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name' size:10"`           // 活动类型名称
	DelFlag    types.String `json:"del_flag,omitempty" db:"'del_flag' size:2"`    // 删除状态（Y/N）
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // 编号
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// CommunicationRecord
// @tablename communication_record
type CommunicationRecord struct {
	pool.Model
	CommunicationRecord types.String `json:"communication_record,omitempty" db:"'communication_record' size:255"` // 沟通记录
	DelFlag             types.String `json:"del_flag,omitempty" db:"'del_flag' size:2"`                           // 删除状态（Y/N）
	Id                  types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`                              // 编号
	ElderId             types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`                          // 老人编号
	RecordDate          types.Time   `json:"record_date,omitempty" db:"'record_date'"`                            // 记录时间
	CreateId            types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`                        // 创建人编号
	CreateTime          types.Time   `json:"create_time,omitempty" db:"'create_time'"`                            // 创建时间
	UpdateId            types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`                        // 修改人编号
	UpdateTime          types.Time   `json:"update_time,omitempty" db:"'update_time'"`                            // 修改时间
}

// DishesType
// @tablename dishes_type
type DishesType struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name' size:10"`           // 菜品类别名称
	DelFlag    types.String `json:"del_flag,omitempty" db:"'del_flag' size:2"`    // 删除状态（Y/N）
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // 编号
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// HealthData
// @tablename health_data
type HealthData struct {
	pool.Model
	LeftEar                  types.String  `json:"left_ear,omitempty" db:"'left_ear' size:5"`                                      // 左耳
	RightEar                 types.String  `json:"right_ear,omitempty" db:"'right_ear' size:5"`                                    // 右耳
	Id                       types.BigInt  `json:"id,omitempty" db:"'id' pk auto size:20"`                                         // 编号
	ElderId                  types.BigInt  `json:"elder_id,omitempty" db:"'elder_id' size:20"`                                     // 老人编号
	Weight                   types.Float64 `json:"weight,omitempty" db:"'weight'"`                                                 // 体重
	Temperature              types.Float64 `json:"temperature,omitempty" db:"'temperature'"`                                       // 体温
	LeftEye                  types.Float64 `json:"left_eye,omitempty" db:"'left_eye'"`                                             // 左眼
	RightEye                 types.Float64 `json:"right_eye,omitempty" db:"'right_eye'"`                                           // 右眼
	CreateId                 types.BigInt  `json:"create_id,omitempty" db:"'create_id' size:20"`                                   // 创建人编号
	CreateTime               types.Time    `json:"create_time,omitempty" db:"'create_time'"`                                       // 创建时间
	UpdateId                 types.BigInt  `json:"update_id,omitempty" db:"'update_id' size:20"`                                   // 修改人编号
	UpdateTime               types.Time    `json:"update_time,omitempty" db:"'update_time'"`                                       // 修改时间
	Height                   types.Int32   `json:"height,omitempty" db:"'height' size:11"`                                         // 身高
	HeartRate                types.Int32   `json:"heart_rate,omitempty" db:"'heart_rate' size:11"`                                 // 心率
	SystolicBloodPressure    types.Int32   `json:"systolic_blood_pressure,omitempty" db:"'systolic_blood_pressure' size:11"`       // 收缩血压
	DiastolicBloodPressure   types.Int32   `json:"diastolic_blood_pressure,omitempty" db:"'diastolic_blood_pressure' size:11"`     // 舒张血压
	FastingBloodGlucose      types.Int32   `json:"fasting_blood_glucose,omitempty" db:"'fasting_blood_glucose' size:11"`           // 空腹血糖
	PostprandialBloodGlucose types.Int32   `json:"postprandial_blood_glucose,omitempty" db:"'postprandial_blood_glucose' size:11"` // 餐后血糖
	BloodOxygenSaturation    types.Int32   `json:"blood_oxygen_saturation,omitempty" db:"'blood_oxygen_saturation' size:11"`       // 血氧饱和度
	Cholesterol              types.Int32   `json:"cholesterol,omitempty" db:"'cholesterol' size:11"`                               // 总胆固醇
	UricAcid                 types.Int32   `json:"uric_acid,omitempty" db:"'uric_acid' size:11"`                                   // 尿酸
	MusclePercentage         types.Int32   `json:"muscle_percentage,omitempty" db:"'muscle_percentage' size:11"`                   // 肌肉率
	BodyFatPercentage        types.Int32   `json:"body_fat_percentage,omitempty" db:"'body_fat_percentage' size:11"`               // 体脂率
	WaistCircumference       types.Int32   `json:"waist_circumference,omitempty" db:"'waist_circumference' size:11"`               // 腰围
	HipCircumference         types.Int32   `json:"hip_circumference,omitempty" db:"'hip_circumference' size:11"`                   // 臀围
	MoistureContent          types.Int32   `json:"moisture_content,omitempty" db:"'moisture_content' size:11"`                     // 水分率
}

// Outward
// @tablename outward
type Outward struct {
	pool.Model
	ChaperoneName  types.String `json:"chaperone_name,omitempty" db:"'chaperone_name' size:10"`   // 陪同人姓名
	ChaperonePhone types.String `json:"chaperone_phone,omitempty" db:"'chaperone_phone' size:11"` // 陪同人电话
	ChaperoneType  types.String `json:"chaperone_type,omitempty" db:"'chaperone_type' size:5"`    // 陪同人类型（家属/护工）
	DelFlag        types.String `json:"del_flag,omitempty" db:"'del_flag' size:2"`                // 删除状态（Y/N）
	Id             types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`                   // 编号
	ElderId        types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`               // 老人编号
	OutwardDate    types.Time   `json:"outward_date,omitempty" db:"'outward_date'"`               // 外出时间
	PlanReturnDate types.Time   `json:"plan_return_date,omitempty" db:"'plan_return_date'"`       // 计划返回时间
	RealReturnDate types.Time   `json:"real_return_date,omitempty" db:"'real_return_date'"`       // 实际返回时间
	CreateId       types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`             // 创建人编号
	CreateTime     types.Time   `json:"create_time,omitempty" db:"'create_time'"`                 // 创建时间
	UpdateId       types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`             // 修改人编号
	UpdateTime     types.Time   `json:"update_time,omitempty" db:"'update_time'"`                 // 修改时间
}

// Reserve
// @tablename reserve
type Reserve struct {
	pool.Model
	Name        types.String  `json:"name,omitempty" db:"'name' size:10"`                // 交款人姓名
	Phone       types.String  `json:"phone,omitempty" db:"'phone' size:11"`              // 交款人电话
	ReserveFlag types.String  `json:"reserve_flag,omitempty" db:"'reserve_flag' size:2"` // 退款状态（N/Y）
	Id          types.BigInt  `json:"id,omitempty" db:"'id' pk auto size:20"`            // 编号
	ElderId     types.BigInt  `json:"elder_id,omitempty" db:"'elder_id' size:20"`        // 老人编号
	StaffId     types.BigInt  `json:"staff_id,omitempty" db:"'staff_id' size:20"`        // 销售人员编号
	DueDate     types.Time    `json:"due_date,omitempty" db:"'due_date'"`                // 预定到期时间
	Deposit     types.Float64 `json:"deposit,omitempty" db:"'deposit' size:10|2"`        // 定金
	CreateId    types.BigInt  `json:"create_id,omitempty" db:"'create_id' size:20"`      // 创建人编号
	CreateTime  types.Time    `json:"create_time,omitempty" db:"'create_time'"`          // 创建时间
	UpdateId    types.BigInt  `json:"update_id,omitempty" db:"'update_id' size:20"`      // 修改人编号
	UpdateTime  types.Time    `json:"update_time,omitempty" db:"'update_time'"`          // 修改时间
}

// VisitPlan
// @tablename visit_plan
type VisitPlan struct {
	pool.Model
	Title        types.String `json:"title,omitempty" db:"'title' size:25"`         // 回访计划标题
	Content      types.String `json:"content,omitempty" db:"'content' size:255"`    // 回访计划内容
	DelFlag      types.String `json:"del_flag,omitempty" db:"'del_flag' size:2"`    // 删除状态（Y/N）
	Id           types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // 编号
	ElderId      types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`   // 老人编号
	PlanDate     types.Time   `json:"plan_date,omitempty" db:"'plan_date'"`         // 计划回访时间
	CompleteDate types.Time   `json:"complete_date,omitempty" db:"'complete_date'"` // 计划完成时间
	CreateId     types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人编号
	CreateTime   types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	UpdateId     types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人编号
	UpdateTime   types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// ElderLabel
// @tablename elder_label
type ElderLabel struct {
	pool.Model
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // 编号
	ElderId    types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`   // 老人编号
	LabelId    types.BigInt `json:"label_id,omitempty" db:"'label_id' size:20"`   // 标签编号
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// Label
// @tablename label
type Label struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name' size:10"`           // 标签名称
	Color      types.String `json:"color,omitempty" db:"'color' size:15"`         // 标签颜色
	DelFlag    types.String `json:"del_flag,omitempty" db:"'del_flag' size:2"`    // 删除状态（Y/N）
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // 编号
	TypeId     types.BigInt `json:"type_id,omitempty" db:"'type_id' size:20"`     // 类别编号
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// Dishes
// @tablename dishes
type Dishes struct {
	pool.Model
	Name       types.String  `json:"name,omitempty" db:"'name' size:15"`           // 菜品名称
	DelFlag    types.String  `json:"del_flag,omitempty" db:"'del_flag' size:2"`    // 删除状态（Y/N）
	Id         types.BigInt  `json:"id,omitempty" db:"'id' pk auto size:20"`       // 编号
	TypeId     types.BigInt  `json:"type_id,omitempty" db:"'type_id' size:20"`     // 菜品类别编号
	Price      types.Float64 `json:"price,omitempty" db:"'price' size:10|2"`       // 菜品价格
	CreateId   types.BigInt  `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人编号
	CreateTime types.Time    `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	UpdateId   types.BigInt  `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人编号
	UpdateTime types.Time    `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// Material
// @tablename material
type Material struct {
	pool.Model
	Name       types.String  `json:"name,omitempty" db:"'name' size:15"`           // 物资名称
	DelFlag    types.String  `json:"del_flag,omitempty" db:"'del_flag' size:2"`    // 删除状态（Y/N）
	Id         types.BigInt  `json:"id,omitempty" db:"'id' pk auto size:20"`       // 编号
	TypeId     types.BigInt  `json:"type_id,omitempty" db:"'type_id' size:20"`     // 物资类别编号
	Price      types.Float64 `json:"price,omitempty" db:"'price' size:10|2"`       // 物资单价
	CreateId   types.BigInt  `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人编号
	CreateTime types.Time    `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	UpdateId   types.BigInt  `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人编号
	UpdateTime types.Time    `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// RoomType
// @tablename room_type
type RoomType struct {
	pool.Model
	Name       types.String  `json:"name,omitempty" db:"'name' size:10"`                 // 房间类型名称
	DelFlag    types.String  `json:"del_flag,omitempty" db:"'del_flag' size:2"`          // 删除状态（Y/N）
	Id         types.BigInt  `json:"id,omitempty" db:"'id' pk auto size:20"`             // 编号
	MonthPrice types.Float64 `json:"month_price,omitempty" db:"'month_price' size:10|2"` // 月房间费用
	CreateId   types.BigInt  `json:"create_id,omitempty" db:"'create_id' size:20"`       // 创建人编号
	CreateTime types.Time    `json:"create_time,omitempty" db:"'create_time'"`           // 创建时间
	UpdateId   types.BigInt  `json:"update_id,omitempty" db:"'update_id' size:20"`       // 修改人编号
	UpdateTime types.Time    `json:"update_time,omitempty" db:"'update_time'"`           // 修改时间
}

// Auth
// @tablename auth
type Auth struct {
	pool.Model
	Title      types.String `json:"title,omitempty" db:"'title' size:10"`         // 权限标题
	Name       types.String `json:"name,omitempty" db:"'name' size:20"`           // 权限名称
	Path       types.String `json:"path,omitempty" db:"'path' size:20"`           // 权限path
	Icon       types.String `json:"icon,omitempty" db:"'icon' size:10"`           // 权限图标
	Url        types.String `json:"url,omitempty" db:"'url' size:50"`             // 权限url
	Type       types.String `json:"type,omitempty" db:"'type' size:5"`            // 权限类别（MENU/BTN）
	Method     types.String `json:"method,omitempty" db:"'method' size:6"`        // 权限请求方式（GET/POST/PUT/DELETE）
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // 编号
	ParentId   types.BigInt `json:"parent_id,omitempty" db:"'parent_id' size:20"` // 父级编号
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// ServiceType
// @tablename service_type
type ServiceType struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name' size:10"`           // 服务项目名称
	DelFlag    types.String `json:"del_flag,omitempty" db:"'del_flag' size:2"`    // 删除状态（Y/N）
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // 编号
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// NurseGroup
// @tablename nurse_group
type NurseGroup struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name' size:10"`           // 护工小组名称
	DelFlag    types.String `json:"del_flag,omitempty" db:"'del_flag' size:2"`    // 删除状态（Y/N）
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // 编号
	StaffId    types.BigInt `json:"staff_id,omitempty" db:"'staff_id' size:20"`   // 护工小组组长编号
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// Retreat
// @tablename retreat
type Retreat struct {
	pool.Model
	RetreatForm  types.String `json:"retreat_form,omitempty" db:"'retreat_form' size:5"`     // 退住形式
	RetreatCause types.String `json:"retreat_cause,omitempty" db:"'retreat_cause' size:255"` // 退住原因
	Id           types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`                // 编号
	ElderId      types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`            // 老人编号
	CreateId     types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`          // 创建人编号
	CreateTime   types.Time   `json:"create_time,omitempty" db:"'create_time'"`              // 创建时间
	UpdateId     types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`          // 修改人编号
	UpdateTime   types.Time   `json:"update_time,omitempty" db:"'update_time'"`              // 修改时间
	Evaluate     types.Int32  `json:"evaluate,omitempty" db:"'evaluate' size:11"`            // 对老人评价
}

// Role
// @tablename role
type Role struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name' size:10"`           // 角色名称
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // 编号
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// MedicineRecord
// @tablename medicine_record
type MedicineRecord struct {
	pool.Model
	MedicineTime  types.String `json:"medicine_time,omitempty" db:"'medicine_time' size:5"`      // 用药时间（早/中/晚）
	Id            types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`                   // 编号
	ElderId       types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`               // 老人编号
	DepositInfoId types.BigInt `json:"deposit_info_id,omitempty" db:"'deposit_info_id' size:20"` // 药品缴存信息编号
	MedicineDate  types.Time   `json:"medicine_date,omitempty" db:"'medicine_date'"`             // 用药日期
	CreateId      types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`             // 创建人编号
	CreateTime    types.Time   `json:"create_time,omitempty" db:"'create_time'"`                 // 创建时间
	UpdateId      types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`             // 修改人编号
	UpdateTime    types.Time   `json:"update_time,omitempty" db:"'update_time'"`                 // 修改时间
}

// NurseItem
// @tablename nurse_item
type NurseItem struct {
	pool.Model
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`         // 编号
	GradeId    types.BigInt `json:"grade_id,omitempty" db:"'grade_id' size:20"`     // 护理等级编号
	ServiceId  types.BigInt `json:"service_id,omitempty" db:"'service_id' size:20"` // 服务项目编号
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`   // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`       // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`   // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`       // 修改时间
}

// NurseReserve
// @tablename nurse_reserve
type NurseReserve struct {
	pool.Model
	ServiceName  types.String  `json:"service_name,omitempty" db:"'service_name' size:10"`     // 服务项目名称
	ChargeMethod types.String  `json:"charge_method,omitempty" db:"'charge_method' size:3"`    // 收费方式
	OrderFlag    types.String  `json:"order_flag,omitempty" db:"'order_flag' size:2"`          // 订单状态
	Id           types.BigInt  `json:"id,omitempty" db:"'id' pk auto size:20"`                 // 编号
	ElderId      types.BigInt  `json:"elder_id,omitempty" db:"'elder_id' size:20"`             // 老人编号
	StaffId      types.BigInt  `json:"staff_id,omitempty" db:"'staff_id' size:20"`             // 服务人编号
	ServicePrice types.Float64 `json:"service_price,omitempty" db:"'service_price' size:10|2"` // 服务费用
	PayAmount    types.Float64 `json:"pay_amount,omitempty" db:"'pay_amount' size:10|2"`       // 支付总额
	NurseDate    types.Time    `json:"nurse_date,omitempty" db:"'nurse_date'"`                 // 护理时间
	CreateId     types.BigInt  `json:"create_id,omitempty" db:"'create_id' size:20"`           // 创建人编号
	CreateTime   types.Time    `json:"create_time,omitempty" db:"'create_time'"`               // 创建时间
	UpdateId     types.BigInt  `json:"update_id,omitempty" db:"'update_id' size:20"`           // 修改人编号
	UpdateTime   types.Time    `json:"update_time,omitempty" db:"'update_time'"`               // 修改时间
	NeedDate     types.Int32   `json:"need_date,omitempty" db:"'need_date' size:11"`           // 所需时间
	Frequency    types.Int32   `json:"frequency,omitempty" db:"'frequency' size:11"`           // 服务次数
}

// WarehouseRecord
// @tablename warehouse_record
type WarehouseRecord struct {
	pool.Model
	Source        types.String `json:"source,omitempty" db:"'source' size:5"`                 // 物资来源
	WarehouseFlag types.String `json:"warehouse_flag,omitempty" db:"'warehouse_flag' size:5"` // 入库状态
	DelFlag       types.String `json:"del_flag,omitempty" db:"'del_flag' size:2"`             // 删除状态（Y/N）
	Id            types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`                // 编号
	WarehouseId   types.BigInt `json:"warehouse_id,omitempty" db:"'warehouse_id' size:20"`    // 仓库编号
	StaffId       types.BigInt `json:"staff_id,omitempty" db:"'staff_id' size:20"`            // 经办人编号
	WarehouseDate types.Time   `json:"warehouse_date,omitempty" db:"'warehouse_date'"`        // 入库时间
	CreateId      types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`          // 创建人编号
	CreateTime    types.Time   `json:"create_time,omitempty" db:"'create_time'"`              // 创建时间
	UpdateId      types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`          // 修改人编号
	UpdateTime    types.Time   `json:"update_time,omitempty" db:"'update_time'"`              // 修改时间
}

// Building
// @tablename building
type Building struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name' size:10"`           // 楼栋名称
	DelFlag    types.String `json:"del_flag,omitempty" db:"'del_flag' size:2"`    // 删除状态（Y/N）
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // 编号
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	FloorNum   types.Int32  `json:"floor_num,omitempty" db:"'floor_num' size:11"` // 楼层数量
}

// OrderDishes
// @tablename order_dishes
type OrderDishes struct {
	pool.Model
	DishesName   types.String  `json:"dishes_name,omitempty" db:"'dishes_name' size:15"`       // 菜品名称
	SetFlag      types.String  `json:"set_flag,omitempty" db:"'set_flag' size:2"`              // 套餐标记
	Id           types.BigInt  `json:"id,omitempty" db:"'id' pk auto size:20"`                 // 编号
	OrderId      types.BigInt  `json:"order_id,omitempty" db:"'order_id' size:20"`             // 订餐编号
	DishesPrice  types.Float64 `json:"dishes_price,omitempty" db:"'dishes_price' size:10|2"`   // 菜品价格
	TotalAmount  types.Float64 `json:"total_amount,omitempty" db:"'total_amount' size:10|2"`   // 菜品总额
	ReallyAmount types.Float64 `json:"really_amount,omitempty" db:"'really_amount' size:10|2"` // 实际总额
	CreateId     types.BigInt  `json:"create_id,omitempty" db:"'create_id' size:20"`           // 创建人编号
	CreateTime   types.Time    `json:"create_time,omitempty" db:"'create_time'"`               // 创建时间
	UpdateId     types.BigInt  `json:"update_id,omitempty" db:"'update_id' size:20"`           // 修改人编号
	UpdateTime   types.Time    `json:"update_time,omitempty" db:"'update_time'"`               // 修改时间
	OrderNum     types.Int32   `json:"order_num,omitempty" db:"'order_num' size:11"`           // 菜品份数
}

// Active
// @tablename active
type Active struct {
	pool.Model
	Theme         types.String `json:"theme,omitempty" db:"'theme' size:10"`                    // 活动主题
	Name          types.String `json:"name,omitempty" db:"'name' size:25"`                      // 活动名称
	Content       types.String `json:"content,omitempty" db:"'content' size:255"`               // 活动内容
	Address       types.String `json:"address,omitempty" db:"'address' size:50"`                // 活动地点
	Organizer     types.String `json:"organizer,omitempty" db:"'organizer' size:10"`            // 组织者姓名
	Phone         types.String `json:"phone,omitempty" db:"'phone' size:11"`                    // 组织者电话
	ActivePicture types.String `json:"active_picture,omitempty" db:"'active_picture' size:255"` // 活动图片
	DelFlag       types.String `json:"del_flag,omitempty" db:"'del_flag' size:2"`               // 删除状态（Y/N）
	Id            types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`                  // 编号
	TypeId        types.BigInt `json:"type_id,omitempty" db:"'type_id' size:20"`                // 活动类别编号
	ActiveDate    types.Time   `json:"active_date,omitempty" db:"'active_date'"`                // 活动日期
	CreateId      types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`            // 创建人编号
	CreateTime    types.Time   `json:"create_time,omitempty" db:"'create_time'"`                // 创建时间
	UpdateId      types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`            // 修改人编号
	UpdateTime    types.Time   `json:"update_time,omitempty" db:"'update_time'"`                // 修改时间
}

// CateringSet
// @tablename catering_set
type CateringSet struct {
	pool.Model
	Name       types.String  `json:"name,omitempty" db:"'name' size:10"`                 // 餐饮套餐名称
	DelFlag    types.String  `json:"del_flag,omitempty" db:"'del_flag' size:2"`          // 删除状态（Y/N）
	Id         types.BigInt  `json:"id,omitempty" db:"'id' pk auto size:20"`             // 编号
	MonthPrice types.Float64 `json:"month_price,omitempty" db:"'month_price' size:10|2"` // 月套餐费用
	CreateId   types.BigInt  `json:"create_id,omitempty" db:"'create_id' size:20"`       // 创建人编号
	CreateTime types.Time    `json:"create_time,omitempty" db:"'create_time'"`           // 创建时间
	UpdateId   types.BigInt  `json:"update_id,omitempty" db:"'update_id' size:20"`       // 修改人编号
	UpdateTime types.Time    `json:"update_time,omitempty" db:"'update_time'"`           // 修改时间
}

// DrugDeposit
// @tablename drug_deposit
type DrugDeposit struct {
	pool.Model
	Mode        types.String `json:"mode,omitempty" db:"'mode' size:5"`                 // 缴存药品方式
	DepositFlag types.String `json:"deposit_flag,omitempty" db:"'deposit_flag' size:5"` // 缴存状态
	DelFlag     types.String `json:"del_flag,omitempty" db:"'del_flag' size:2"`         // 删除状态（Y/N）
	Id          types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`            // 编号
	ElderId     types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`        // 老人编号
	CreateId    types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`      // 创建人编号
	CreateTime  types.Time   `json:"create_time,omitempty" db:"'create_time'"`          // 创建时间
	UpdateId    types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`      // 修改人编号
	UpdateTime  types.Time   `json:"update_time,omitempty" db:"'update_time'"`          // 修改时间
}

// FamilyMember
// @tablename family_member
type FamilyMember struct {
	pool.Model
	Name        types.String `json:"name,omitempty" db:"'name' size:10"`                // 家属姓名
	IdNum       types.String `json:"id_num,omitempty" db:"'id_num' size:18"`            // 身份证号
	Phone       types.String `json:"phone,omitempty" db:"'phone' size:11"`              // 家属电话
	Email       types.String `json:"email,omitempty" db:"'email' size:30"`              // 家属邮箱
	Address     types.String `json:"address,omitempty" db:"'address' size:50"`          // 地址
	Relation    types.String `json:"relation,omitempty" db:"'relation' size:5"`         // 与老人关系
	ReceiveFlag types.String `json:"receive_flag,omitempty" db:"'receive_flag' size:2"` // 是否接收消息（Y/N）
	DelFlag     types.String `json:"del_flag,omitempty" db:"'del_flag' size:2"`         // 删除状态（Y/N）
	Id          types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`            // 编号
	ElderId     types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`        // 老人编号
	CreateId    types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`      // 创建人编号
	CreateTime  types.Time   `json:"create_time,omitempty" db:"'create_time'"`          // 创建时间
	UpdateId    types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`      // 修改人编号
	UpdateTime  types.Time   `json:"update_time,omitempty" db:"'update_time'"`          // 修改时间
}

// Medicine
// @tablename medicine
type Medicine struct {
	pool.Model
	Name          types.String `json:"name,omitempty" db:"'name' size:15"`                   // 药品名称
	Type          types.String `json:"type,omitempty" db:"'type' size:5"`                    // 药品类别
	Specification types.String `json:"specification,omitempty" db:"'specification' size:10"` // 药品规格
	DosageForm    types.String `json:"dosage_form,omitempty" db:"'dosage_form' size:5"`      // 药品剂型
	Manufacturer  types.String `json:"manufacturer,omitempty" db:"'manufacturer' size:25"`   // 生产厂家
	DelFlag       types.String `json:"del_flag,omitempty" db:"'del_flag' size:2"`            // 删除状态（Y/N）
	Id            types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`               // 编号
	CreateId      types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`         // 创建人编号
	CreateTime    types.Time   `json:"create_time,omitempty" db:"'create_time'"`             // 创建时间
	UpdateId      types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`         // 修改人编号
	UpdateTime    types.Time   `json:"update_time,omitempty" db:"'update_time'"`             // 修改时间
}

// NurseGroupMember
// @tablename nurse_group_member
type NurseGroupMember struct {
	pool.Model
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // 编号
	GroupId    types.BigInt `json:"group_id,omitempty" db:"'group_id' size:20"`   // 护工小组编号
	StaffId    types.BigInt `json:"staff_id,omitempty" db:"'staff_id' size:20"`   // 护工编号
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// Source
// @tablename source
type Source struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name' size:10"`           // 来源渠道名称
	DelFlag    types.String `json:"del_flag,omitempty" db:"'del_flag' size:2"`    // 删除状态（Y/N）
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // 编号
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// Warehouse
// @tablename warehouse
type Warehouse struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name' size:10"`           // 仓库名称
	DelFlag    types.String `json:"del_flag,omitempty" db:"'del_flag' size:2"`    // 删除状态（Y/N）
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // 编号
	StaffId    types.BigInt `json:"staff_id,omitempty" db:"'staff_id' size:20"`   // 仓库管理员编号
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// MedicineSet
// @tablename medicine_set
type MedicineSet struct {
	pool.Model
	MedicineTime  types.String `json:"medicine_time,omitempty" db:"'medicine_time' size:5"`      // 用药时间（餐前/餐后）
	Id            types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`                   // 编号
	DepositInfoId types.BigInt `json:"deposit_info_id,omitempty" db:"'deposit_info_id' size:20"` // 药品缴存信息编号
	CreateId      types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`             // 创建人编号
	CreateTime    types.Time   `json:"create_time,omitempty" db:"'create_time'"`                 // 创建时间
	UpdateId      types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`             // 修改人编号
	UpdateTime    types.Time   `json:"update_time,omitempty" db:"'update_time'"`                 // 修改时间
	DayFrequency  types.Int32  `json:"day_frequency,omitempty" db:"'day_frequency' size:11"`     // 天频率
}

// Accident
// @tablename accident
type Accident struct {
	pool.Model
	Description types.String `json:"description,omitempty" db:"'description' size:255"` // 事故描述
	Picture     types.String `json:"picture,omitempty" db:"'picture' size:255"`         // 事故图片
	DelFlag     types.String `json:"del_flag,omitempty" db:"'del_flag' size:2"`         // 删除状态（Y/N）
	Id          types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`            // 编号
	ElderId     types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`        // 老人编号
	StaffId     types.BigInt `json:"staff_id,omitempty" db:"'staff_id' size:20"`        // 值班护工编号
	OccurDate   types.Time   `json:"occur_date,omitempty" db:"'occur_date'"`            // 发生时间
	CreateId    types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`      // 创建人编号
	CreateTime  types.Time   `json:"create_time,omitempty" db:"'create_time'"`          // 创建时间
	UpdateId    types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`      // 修改人编号
	UpdateTime  types.Time   `json:"update_time,omitempty" db:"'update_time'"`          // 修改时间
}

// EmergencyContact
// @tablename emergency_contact
type EmergencyContact struct {
	pool.Model
	Name        types.String `json:"name,omitempty" db:"'name' size:10"`                // 紧急联系人姓名
	Phone       types.String `json:"phone,omitempty" db:"'phone' size:11"`              // 紧急联系人电话
	Email       types.String `json:"email,omitempty" db:"'email' size:30"`              // 紧急联系人邮箱
	Relation    types.String `json:"relation,omitempty" db:"'relation' size:5"`         // 与老人关系
	ReceiveFlag types.String `json:"receive_flag,omitempty" db:"'receive_flag' size:2"` // 是否接收消息（Y/N）
	Id          types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`            // 编号
	ElderId     types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`        // 老人编号
	CreateId    types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`      // 创建人编号
	CreateTime  types.Time   `json:"create_time,omitempty" db:"'create_time'"`          // 创建时间
	UpdateId    types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`      // 修改人编号
	UpdateTime  types.Time   `json:"update_time,omitempty" db:"'update_time'"`          // 修改时间
}

// Nurse
// @tablename nurse
type Nurse struct {
	pool.Model
	CompleteFlag types.String `json:"complete_flag,omitempty" db:"'complete_flag' size:5"` // 护理完成情况
	DineFlag     types.String `json:"dine_flag,omitempty" db:"'dine_flag' size:5"`         // 进餐情况
	Rest         types.String `json:"rest,omitempty" db:"'rest' size:5"`                   // 休息
	TakeMedicine types.String `json:"take_medicine,omitempty" db:"'take_medicine' size:5"` // 服药
	Active       types.String `json:"active,omitempty" db:"'active' size:5"`               // 活动
	Id           types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`              // 编号
	ElderId      types.BigInt `json:"elder_id,omitempty" db:"'elder_id' size:20"`          // 老人编号
	StaffId      types.BigInt `json:"staff_id,omitempty" db:"'staff_id' size:20"`          // 护理人员编号
	NurseDate    types.Time   `json:"nurse_date,omitempty" db:"'nurse_date'"`              // 护理时间
	CreateId     types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`        // 创建人编号
	CreateTime   types.Time   `json:"create_time,omitempty" db:"'create_time'"`            // 创建时间
	UpdateId     types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`        // 修改人编号
	UpdateTime   types.Time   `json:"update_time,omitempty" db:"'update_time'"`            // 修改时间
}

// RoleCopy1
// @tablename role_copy1
type RoleCopy1 struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name' size:10"`           // 角色名称
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // 编号
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
}

// BaseAttachment
// @tablename base_attachment
type BaseAttachment struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name' size:225"`           // 上传后文件名
	RealName   types.String `json:"real_name,omitempty" db:"'real_name' size:225"` // 文件真实名称
	Path       types.String `json:"path,omitempty" db:"'path' size:225"`           // 文件绝对路径
	Url        types.String `json:"url,omitempty" db:"'url' size:225"`             // url相对路径
	Suff       types.String `json:"suff,omitempty" db:"'suff' size:225"`           // 文件后缀
	DelFlag    types.String `json:"del_flag,omitempty" db:"'del_flag' size:2"`     // 删除状态(Y/N)
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`        // 编号
	Size       types.BigInt `json:"size,omitempty" db:"'size' size:20"`            // 文件大小 B
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`  // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`      // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`  // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`      // 修改时间
}

// Order
// @tablename order
type Order struct {
	pool.Model
	DineType          types.String  `json:"dine_type,omitempty" db:"'dine_type' size:5"`              // 就餐方式
	OrderFlag         types.String  `json:"order_flag,omitempty" db:"'order_flag' size:2"`            // 订单状态
	Id                types.BigInt  `json:"id,omitempty" db:"'id' pk auto size:20"`                   // 编号
	ElderId           types.BigInt  `json:"elder_id,omitempty" db:"'elder_id' size:20"`               // 老人编号
	StaffId           types.BigInt  `json:"staff_id,omitempty" db:"'staff_id' size:20"`               // 送餐人编号
	DeliverDishesDate types.Time    `json:"deliver_dishes_date,omitempty" db:"'deliver_dishes_date'"` // 送餐时间
	DineDate          types.Time    `json:"dine_date,omitempty" db:"'dine_date'"`                     // 就餐时间
	PayAmount         types.Float64 `json:"pay_amount,omitempty" db:"'pay_amount' size:10|2"`         // 支付总额
	CreateId          types.BigInt  `json:"create_id,omitempty" db:"'create_id' size:20"`             // 创建人编号
	CreateTime        types.Time    `json:"create_time,omitempty" db:"'create_time'"`                 // 创建时间
	UpdateId          types.BigInt  `json:"update_id,omitempty" db:"'update_id' size:20"`             // 修改人编号
	UpdateTime        types.Time    `json:"update_time,omitempty" db:"'update_time'"`                 // 修改时间
}

// Room
// @tablename room
type Room struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name' size:30"`           // 房间名称
	DelFlag    types.String `json:"del_flag,omitempty" db:"'del_flag' size:2"`    // 删除状态（Y/N）
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`       // 编号
	TypeId     types.BigInt `json:"type_id,omitempty" db:"'type_id' size:20"`     // 房间类型编号
	FloorId    types.BigInt `json:"floor_id,omitempty" db:"'floor_id' size:20"`   // 楼层编号
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"` // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`     // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"` // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`     // 修改时间
	BedNum     types.Int32  `json:"bed_num,omitempty" db:"'bed_num' size:11"`     // 床位数量
}

// Staff
// @tablename staff
type Staff struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name' size:10"`            // 姓名
	IdNum      types.String `json:"id_num,omitempty" db:"'id_num' size:18"`        // 身份证号
	Sex        types.String `json:"sex,omitempty" db:"'sex' size:2"`               // 性别(男/女)
	Phone      types.String `json:"phone,omitempty" db:"'phone' size:11"`          // 电话
	Email      types.String `json:"email,omitempty" db:"'email' size:30"`          // 邮箱
	Pass       types.String `json:"pass,omitempty" db:"'pass' size:255"`           // 密码
	Avator     types.String `json:"avator,omitempty" db:"'avator' size:255"`       // 头像
	Address    types.String `json:"address,omitempty" db:"'address' size:50"`      // 地址
	LeaveFlag  types.String `json:"leave_flag,omitempty" db:"'leave_flag' size:2"` // 离职状态（Y/N）
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`        // 编号
	RoleId     types.BigInt `json:"role_id,omitempty" db:"'role_id' size:20"`      // 角色编号
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`  // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`      // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`  // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`      // 修改时间
	Age        types.Int32  `json:"age,omitempty" db:"'age' size:11"`              // 年龄
}

// Elder
// @tablename elder
type Elder struct {
	pool.Model
	Name           types.String  `json:"name,omitempty" db:"'name' size:10"`                         // 老人姓名
	IdNum          types.String  `json:"id_num,omitempty" db:"'id_num' size:18"`                     // 身份证号
	Sex            types.String  `json:"sex,omitempty" db:"'sex' size:2"`                            // 性别(男/女)
	Phone          types.String  `json:"phone,omitempty" db:"'phone' size:11"`                       // 老人电话
	Address        types.String  `json:"address,omitempty" db:"'address' size:50"`                   // 地址
	CheckFlag      types.String  `json:"check_flag,omitempty" db:"'check_flag' size:11"`             // 入住状态
	Id             types.BigInt  `json:"id,omitempty" db:"'id' pk auto size:20"`                     // 编号
	NursingGradeId types.BigInt  `json:"nursing_grade_id,omitempty" db:"'nursing_grade_id' size:20"` // 护理等级编号
	CateringSetId  types.BigInt  `json:"catering_set_id,omitempty" db:"'catering_set_id' size:20"`   // 餐饮套餐编号
	BedId          types.BigInt  `json:"bed_id,omitempty" db:"'bed_id' size:20"`                     // 床位编号
	Balance        types.Float64 `json:"balance,omitempty" db:"'balance' size:10|2"`                 // 余额
	CreateId       types.BigInt  `json:"create_id,omitempty" db:"'create_id' size:20"`               // 创建人编号
	CreateTime     types.Time    `json:"create_time,omitempty" db:"'create_time'"`                   // 创建时间
	UpdateId       types.BigInt  `json:"update_id,omitempty" db:"'update_id' size:20"`               // 修改人编号
	UpdateTime     types.Time    `json:"update_time,omitempty" db:"'update_time'"`                   // 修改时间
	Age            types.Int32   `json:"age,omitempty" db:"'age' size:11"`                           // 年龄
}

// Floor
// @tablename floor
type Floor struct {
	pool.Model
	Name       types.String `json:"name,omitempty" db:"'name' size:20"`               // 楼层名称
	DelFlag    types.String `json:"del_flag,omitempty" db:"'del_flag' size:2"`        // 删除状态（Y/N）
	Id         types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`           // 编号
	BuildingId types.BigInt `json:"building_id,omitempty" db:"'building_id' size:20"` // 楼栋编号
	CreateId   types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`     // 创建人编号
	CreateTime types.Time   `json:"create_time,omitempty" db:"'create_time'"`         // 创建时间
	UpdateId   types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`     // 修改人编号
	UpdateTime types.Time   `json:"update_time,omitempty" db:"'update_time'"`         // 修改时间
	RoomNum    types.Int32  `json:"room_num,omitempty" db:"'room_num' size:11"`       // 房间数量
}

// OutboundRecord
// @tablename outbound_record
type OutboundRecord struct {
	pool.Model
	RecipientType types.String `json:"recipient_type,omitempty" db:"'recipient_type' size:5"` // 领用人类型
	MaterialUse   types.String `json:"material_use,omitempty" db:"'material_use' size:5"`     // 物资去向
	OutboundFlag  types.String `json:"outbound_flag,omitempty" db:"'outbound_flag' size:5"`   // 出库状态
	DelFlag       types.String `json:"del_flag,omitempty" db:"'del_flag' size:2"`             // 删除状态（Y/N）
	Id            types.BigInt `json:"id,omitempty" db:"'id' pk auto size:20"`                // 编号
	WarehouseId   types.BigInt `json:"warehouse_id,omitempty" db:"'warehouse_id' size:20"`    // 仓库编号
	StaffId       types.BigInt `json:"staff_id,omitempty" db:"'staff_id' size:20"`            // 经办人编号
	RecipientId   types.BigInt `json:"recipient_id,omitempty" db:"'recipient_id' size:20"`    // 领用人编号
	OutboundDate  types.Time   `json:"outbound_date,omitempty" db:"'outbound_date'"`          // 出库时间
	CreateId      types.BigInt `json:"create_id,omitempty" db:"'create_id' size:20"`          // 创建人编号
	CreateTime    types.Time   `json:"create_time,omitempty" db:"'create_time'"`              // 创建时间
	UpdateId      types.BigInt `json:"update_id,omitempty" db:"'update_id' size:20"`          // 修改人编号
	UpdateTime    types.Time   `json:"update_time,omitempty" db:"'update_time'"`              // 修改时间
}

// ServiceItem
// @tablename service_item
type ServiceItem struct {
	pool.Model
	Name         types.String  `json:"name,omitempty" db:"'name' size:10"`                  // 服务名称
	ChargeMethod types.String  `json:"charge_method,omitempty" db:"'charge_method' size:3"` // 收费方式
	DelFlag      types.String  `json:"del_flag,omitempty" db:"'del_flag' size:2"`           // 删除状态（Y/N）
	Id           types.BigInt  `json:"id,omitempty" db:"'id' pk auto size:20"`              // 编号
	TypeId       types.BigInt  `json:"type_id,omitempty" db:"'type_id' size:20"`            // 服务项目类别编号
	Price        types.Float64 `json:"price,omitempty" db:"'price' size:10|2"`              // 服务费用
	CreateId     types.BigInt  `json:"create_id,omitempty" db:"'create_id' size:20"`        // 创建人编号
	CreateTime   types.Time    `json:"create_time,omitempty" db:"'create_time'"`            // 创建时间
	UpdateId     types.BigInt  `json:"update_id,omitempty" db:"'update_id' size:20"`        // 修改人编号
	UpdateTime   types.Time    `json:"update_time,omitempty" db:"'update_time'"`            // 修改时间
	NeedDate     types.Int32   `json:"need_date,omitempty" db:"'need_date' size:11"`        // 所需时间(分)
}
