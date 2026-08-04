package dto

import "time"

// @response
// Result 统一响应包装（对应 Java Result<T>）
type Result struct {
	Code int         `json:"code"` // 状态码
	Msg  string      `json:"msg"`  // 提示信息
	Data interface{} `json:"data"` // 响应数据
}

// @response
// PageResult 分页响应包装（对应 Java PageResult<T>）
type PageResult struct {
	PageNum  int         `json:"page_num"`  // 当前页码
	PageSize int         `json:"page_size"` // 每页条数
	Pages    int         `json:"pages"`     // 总页数
	Size     int         `json:"size"`      // 当前页条数
	Total    int64       `json:"total"`     // 总记录数
	List     interface{} `json:"list"`      // 数据列表
}

// @response
// DropDown 下拉列表响应（对应 Java DropDown）
type DropDown struct {
	ID   int64  `json:"id"`   // 编号
	Name string `json:"name"` // 名称
}

// // @response
// // Rank 序号基类（对应 Java Rank）
// type Rank struct {
// 	Rank int64 `json:"rank"` // 序号
// }

// @response
// BuildingVO 楼栋-楼层-房间-床位树响应（对应 Java base.BuildingVo）
type BuildingVO struct {
	ID        int64            `json:"id"`         // 楼栋编号
	Name      string           `json:"name"`       // 楼栋名称
	FloorNum  int              `json:"floor_num"`  // 楼层数量
	FloorList []BuildingItemVO `json:"floor_list"` // 楼层列表
}

// @response
// BuildingItemVO 楼层（对应 Java base.BuildingVo.BuildingItem）
type BuildingItemVO struct {
	ID       int64         `json:"id"`        // 楼层编号
	Name     string        `json:"name"`      // 楼层名称
	RoomNum  int           `json:"room_num"`  // 房间数量
	RoomList []FloorItemVO `json:"room_list"` // 房间列表
}

// @response
// FloorItemVO 房间（对应 Java base.BuildingVo.BuildingItem.FloorItem）
type FloorItemVO struct {
	ID      int64        `json:"id"`       // 房间编号
	Name    string       `json:"name"`     // 房间名称
	BedNum  int          `json:"bed_num"`  // 床位数量
	BedList []RoomItemVO `json:"bed_list"` // 床位列表
}

// @response
// RoomItemVO 床位（对应 Java base.BuildingVo...RoomItem）
type RoomItemVO struct {
	ID        int64  `json:"id"`         // 床位编号
	Name      string `json:"name"`       // 床位名称
	BedFlag   string `json:"bed_flag"`   // 床位状态
	ElderName string `json:"elder_name"` // 所住老人
	Sex       string `json:"sex"`        // 老人性别
	Age       int    `json:"age"`        // 老人年龄
}

// ============ 跨 Controller 共享请求对象 ============

// @request
// OperateEmergencyContactQuery 紧急联系人请求（被 CheckContract、ElderRecord 引用，对应 Java 内部静态类）
type OperateEmergencyContactQuery struct {
	Name        *string `json:"name" valid:"required"`         // 紧急联系人姓名
	Phone       *string `json:"phone" valid:"required"`        // 紧急联系人电话
	Email       *string `json:"email" valid:"required"`        // 紧急联系人邮箱
	Relation    *string `json:"relation" valid:"required"`     // 与老人关系
	ReceiveFlag *string `json:"receive_flag" valid:"required"` // 是否接收消息
}

// @request
// OperateNurseGradeQuery 操作护理等级请求（被 NurseGrade、ElderRecord 引用）
type OperateNurseGradeQuery struct {
	ID            *int64   `json:"id"`                               // id
	Name          *string  `json:"name" valid:"required"`            // 护理等级名称
	Type          *string  `json:"type" valid:"required"`            // 护理类型
	MonthPrice    *float64 `json:"month_price" valid:"required"`     // 月护理费用
	ServiceIDList []int64  `json:"service_id_list" valid:"required"` // 护理服务编号列表
}

// @request
// OperateServiceQuery 操作服务请求（被 ServiceProject、NurseGrade 引用）
type OperateServiceQuery struct {
	ID           *int64   `json:"id"`                             // id
	TypeID       *int64   `json:"type_id" valid:"required"`       // 服务类型编号
	Name         *string  `json:"name" valid:"required"`          // 服务名称
	ChargeMethod *string  `json:"charge_method" valid:"required"` // 收费方式
	Price        *float64 `json:"price" valid:"required"`         // 服务价格
	NeedDate     *int     `json:"need_date" valid:"required"`     // 所需时间
}

// // @response
// // SetDishesVO 套餐/菜品嵌套响应（被 CateringSet、ElderRecord 引用）
// type SetDishesVO struct {
// 	ID    int64   `json:"id"`    // id
// 	Name  string  `json:"name"`  // 菜品名称
// 	Price float64 `json:"price"` // 菜品价格
// }

// ============ 跨 Controller 共享响应对象 ============

// @response
// GetNurseGradeByIDVO 护理等级详情响应（被 ElderRecord 引用）
type GetNurseGradeByIDVO struct {
	OperateNurseGradeQuery
	NurseGradeServiceVOList []NurseGradeServiceVO `json:"nurse_grade_service_vo_list"` // 护理等级服务列表
}

// @response
// NurseGradeServiceVO 护理等级服务响应（嵌套，继承 OperateServiceQuery）
type NurseGradeServiceVO struct {
	OperateServiceQuery
}

// // @response
// // GetCateringSetByIDVO 餐饮套餐详情响应（被 ElderRecord 引用）
// type GetCateringSetByIDVO struct {
// 	ID              int64         `json:"id"`                 // id
// 	Name            string        `json:"name"`               // 套餐名称
// 	MonthPrice      float64       `json:"month_price"`        // 月套餐费用
// 	SetDishesVOList []SetDishesVO `json:"set_dishes_vo_list"` // 护理等级服务列表
// }

// @response
// GetBedByIDVO 床位详情响应（被 ElderRecord 引用）
type GetBedByIDVO struct {
	BedID      int64   `json:"bed_id"`      // 床位编号
	BedName    string  `json:"bed_name"`    // 床位名称
	RoomType   string  `json:"room_type"`   // 房间类型
	MonthPrice float64 `json:"month_price"` // 月床位费用
}

// @response
// GetElderRecordByIDVO 长者档案详情响应（被 ElderRecord 引用）
type GetElderRecordByIDVO struct {
	Name                            string                         `json:"name"`                                  // 姓名
	IDNum                           string                         `json:"id_num"`                                // 身份证号
	Age                             int                            `json:"age"`                                   // 年龄
	Sex                             string                         `json:"sex"`                                   // 性别
	Phone                           string                         `json:"phone"`                                 // 电话
	Address                         string                         `json:"address"`                               // 地址
	ElderEmergencyContactByIDVOList []OperateEmergencyContactQuery `json:"elder_emergency_contact_by_id_vo_list"` // 长者紧急联系人
	ElderNurseGradeByIDVO           *GetNurseGradeByIDVO           `json:"elder_nurse_grade_by_id_vo"`            // 长者护理等级
	ElderCateringSetByIDVO          *GetCateringSetByIDVO          `json:"elder_catering_set_by_id_vo"`           // 长者餐饮套餐
	ElderBedByIDVO                  *GetBedByIDVO                  `json:"elder_bed_by_id_vo"`                    // 长者床位
}

// @response
// GetElderLabelByIDLabelVO 客户标签响应（被 ElderRecord 引用）
type GetElderLabelByIDLabelVO struct {
	ID    int64  `json:"id"`    // 编号
	Name  string `json:"name"`  // 名称
	Color string `json:"color"` // 颜色
}

// @response
// ListLabelVO 客户标签分类列表响应（被 ElderRecord、Label 引用）
type ListLabelVO struct {
	ID            int64         `json:"id"`              // 标签分类编号
	Name          string        `json:"name"`            // 标签分类名称
	LabelItemList []LabelItemVO `json:"label_item_list"` // 标签分类子项列表
}

// @response
// LabelItemVO 标签子项（嵌套）
type LabelItemVO struct {
	ID     int64  `json:"id"`      // 标签编号
	TypeID int64  `json:"type_id"` // 标签分类编号
	Name   string `json:"name"`    // 标签名称
	Color  string `json:"color"`   // 标签颜色
}

// @response
// GetElderFeeByIDVO 老人费用详情响应（被 ElderRecord 引用）
type GetElderFeeByIDVO struct {
	ElderName         string      `json:"elder_name"`          // 老人姓名
	ContractStartTime time.Time   `json:"contract_start_time"` // 合同开始时间
	ContractEndTime   time.Time   `json:"contract_end_time"`   // 合同结束时间
	FeeDetailList     []FeeDetail `json:"fee_detail_list"`     // 费用详情列表
}

// @response
// FeeDetail 费用详情（嵌套）
type FeeDetail struct {
	FeeDate                  string  `json:"fee_date"`                    // 消费时间
	ContractInsideNurseFee   float64 `json:"contract_inside_nurse_fee"`   // 合同内护理费用
	ContractInsideDishesFee  float64 `json:"contract_inside_dishes_fee"`  // 合同内餐饮费用
	ContractInsideBedFee     float64 `json:"contract_inside_bed_fee"`     // 合同内床位费用
	OrderNurseFee            float64 `json:"order_nurse_fee"`             // 预定护理费用
	OrderDishesFee           float64 `json:"order_dishes_fee"`            // 预定套餐费用
	ContractOutsideNurseFee  float64 `json:"contract_outside_nurse_fee"`  // 合同外护理费用
	ContractOutsideDishesFee float64 `json:"contract_outside_dishes_fee"` // 合同外餐饮费用
	ContractOutsideBedFee    float64 `json:"contract_outside_bed_fee"`    // 合同外床位费用
	PayableFee               float64 `json:"payable_fee"`                 // 应缴费用
	ReturnableFee            float64 `json:"returnable_fee"`              // 应退费用
}

// Date 仅用于标识时间类型字段（go 中使用 time.Time）
var _ = time.Time{}

// @request
// EmptyReq 空请求（无参接口复用）
type EmptyReq struct{}

// @response
// EmptyResp 空响应（无返回内容接口复用）
type EmptyResp struct{}

// @request
// IDReq 单编号请求
type IDReq struct {
	ID *int64 `json:"id" valid:"required"` // 编号
}

// @request
// NameReq 单名称请求（对应 Java @RequestParam String name）
type NameReq struct {
	Name *string `json:"name"` // 名称
}

// @request
// StringReq 单字符串请求
type StringReq struct {
	Value *string `json:"value"` // 值
}
