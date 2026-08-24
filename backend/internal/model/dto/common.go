package dto

import (
	"github.com/linbaozhong/gentity/pkg/types"
	"time"
)

// Result 统一响应包装（对应 Java Result<T>）
// @response
type Result struct {
	Code int         `json:"code"` // 状态码
	Msg  string      `json:"msg"`  // 提示信息
	Data interface{} `json:"data"` // 响应数据
}

// PageResult 分页响应包装（对应 Java PageResult<T>）
// @response
type PageResult struct {
	PageNum  int         `json:"page_num"`  // 当前页码
	PageSize int         `json:"page_size"` // 每页条数
	Pages    int         `json:"pages"`     // 总页数
	Size     int         `json:"size"`      // 当前页条数
	Total    int64       `json:"total"`     // 总记录数
	List     interface{} `json:"list"`      // 数据列表
}

// DropDown 下拉列表响应（对应 Java DropDown）
// @response
type DropDown struct {
	ID   types.BigInt `json:"id"`   // 编号
	Name string       `json:"name"` // 名称
}

// // @response
// // Rank 序号基类（对应 Java Rank）
// type Rank struct {
// 	Rank int64 `json:"rank"` // 序号
// }

// BuildingResp 楼栋-楼层-房间-床位树响应（对应 Java base.BuildingVo）
// @response
type BuildingResp struct {
	ID        types.BigInt       `json:"id"`         // 楼栋编号
	Name      string             `json:"name"`       // 楼栋名称
	FloorNum  int                `json:"floor_num"`  // 楼层数量
	FloorList []BuildingItemResp `json:"floor_list"` // 楼层列表
}

// BuildingItemResp 楼层（对应 Java base.BuildingVo.BuildingItem）
// @response
type BuildingItemResp struct {
	ID       types.BigInt    `json:"id"`        // 楼层编号
	Name     string          `json:"name"`      // 楼层名称
	RoomNum  int             `json:"room_num"`  // 房间数量
	RoomList []FloorItemResp `json:"room_list"` // 房间列表
}

// FloorItemResp 房间（对应 Java base.BuildingVo.BuildingItem.FloorItem）
// @response
type FloorItemResp struct {
	ID      types.BigInt   `json:"id"`       // 房间编号
	Name    string         `json:"name"`     // 房间名称
	BedNum  int            `json:"bed_num"`  // 床位数量
	BedList []RoomItemResp `json:"bed_list"` // 床位列表
}

// RoomItemResp 床位（对应 Java base.BuildingVo...RoomItem）
// @response
type RoomItemResp struct {
	ID        types.BigInt `json:"id"`         // 床位编号
	Name      string       `json:"name"`       // 床位名称
	BedFlag   string       `json:"bed_flag"`   // 床位状态
	ElderName string       `json:"elder_name"` // 所住老人
	Sex       string       `json:"sex"`        // 老人性别
	Age       int          `json:"age"`        // 老人年龄
}

// ============ 跨 Controller 共享请求对象 ============

// OperateEmergencyContactReq 紧急联系人请求（被 CheckContract、ElderRecord 引用，对应 Java 内部静态类）
// @request
type OperateEmergencyContactReq struct {
	Name        *string `json:"name" valid:"required"`         // 紧急联系人姓名
	Phone       *string `json:"phone" valid:"required"`        // 紧急联系人电话
	Email       *string `json:"email" valid:"required"`        // 紧急联系人邮箱
	Relation    *string `json:"relation" valid:"required"`     // 与老人关系
	ReceiveFlag *int8   `json:"receive_flag" valid:"required"` // 是否接收消息
}

// OperateNurseGradeReq 操作护理等级请求（被 NurseGrade、ElderRecord 引用）
// @request
type OperateNurseGradeReq struct {
	ID            *int64       `json:"id"`                               // id
	Name          *string      `json:"name" valid:"required"`            // 护理等级名称
	Type          *string      `json:"type" valid:"required"`            // 护理类型
	MonthPrice    *types.Money `json:"month_price" valid:"required"`     // 月护理费用
	ServiceIDList []int64      `json:"service_id_list" valid:"required"` // 护理服务编号列表
}

// OperateServiceReq 操作服务请求（被 ServiceProject、NurseGrade 引用）
// @request
type OperateServiceReq struct {
	ID           *int64       `json:"id"`                             // id
	TypeID       *int64       `json:"type_id" valid:"required"`       // 服务类型编号
	Name         *string      `json:"name" valid:"required"`          // 服务名称
	ChargeMethod *string      `json:"charge_method" valid:"required"` // 收费方式
	Price        *types.Money `json:"price" valid:"required"`         // 服务价格
	NeedDate     *int         `json:"need_date" valid:"required"`     // 所需时间
}

// // @response
// // SetDishesResp 套餐/菜品嵌套响应（被 CateringSet、ElderRecord 引用）
// type SetDishesResp struct {
// 	ID    int64   `json:"id"`    // id
// 	Name  string  `json:"name"`  // 菜品名称
// 	Price types.Money `json:"price"` // 菜品价格
// }

// ============ 跨 Controller 共享响应对象 ============

// GetNurseGradeByIDResp 护理等级详情响应（被 ElderRecord 引用）
// @response
type GetNurseGradeByIDResp struct {
	OperateNurseGradeReq
	NurseGradeServiceRespList []NurseGradeServiceResp `json:"nurse_grade_service_resp_list"` // 护理等级服务列表
}

// NurseGradeServiceResp 护理等级服务响应（嵌套，继承 OperateServiceReq）
// @response
type NurseGradeServiceResp struct {
	OperateServiceReq
}

// // @response
// // GetCateringSetByIDResp 餐饮套餐详情响应（被 ElderRecord 引用）
// type GetCateringSetByIDResp struct {
// 	ID              int64         `json:"id"`                 // id
// 	Name            string        `json:"name"`               // 套餐名称
// 	MonthPrice      types.Money       `json:"month_price"`        // 月套餐费用
// 	SetDishesRespList []SetDishesResp `json:"set_dishes_resp_list"` // 护理等级服务列表
// }

// GetBedByIDResp 床位详情响应（被 ElderRecord 引用）
// @response
type GetBedByIDResp struct {
	BedID      types.BigInt `json:"bed_id"`      // 床位编号
	BedName    string       `json:"bed_name"`    // 床位名称
	RoomType   string       `json:"room_type"`   // 房间类型
	MonthPrice types.Money  `json:"month_price"` // 月床位费用
}

// GetElderRecordByIDResp 长者档案详情响应（被 ElderRecord 引用）
// @response
type GetElderRecordByIDResp struct {
	Name                              string                       `json:"name"`                                    // 姓名
	IDNum                             string                       `json:"id_num"`                                  // 身份证号
	Age                               int                          `json:"age"`                                     // 年龄
	Sex                               string                       `json:"sex"`                                     // 性别
	Phone                             string                       `json:"phone"`                                   // 电话
	Address                           string                       `json:"address"`                                 // 地址
	ElderEmergencyContactByIDRespList []OperateEmergencyContactReq `json:"elder_emergency_contact_by_id_resp_list"` // 长者紧急联系人
	ElderNurseGradeByIDResp           *GetNurseGradeByIDResp       `json:"elder_nurse_grade_by_id_resp"`            // 长者护理等级
	ElderCateringSetByIDResp          *GetCateringSetByIDResp      `json:"elder_catering_set_by_id_resp"`           // 长者餐饮套餐
	ElderBedByIDResp                  *GetBedByIDResp              `json:"elder_bed_by_id_resp"`                    // 长者床位
}

// GetElderLabelByIDLabelResp 客户标签响应（被 ElderRecord 引用）
// @response
type GetElderLabelByIDLabelResp struct {
	ID    types.BigInt `json:"id"`    // 编号
	Name  string       `json:"name"`  // 名称
	Color string       `json:"color"` // 颜色
}

// ListLabelResp 客户标签分类列表响应（被 ElderRecord、Label 引用）
// @response
type ListLabelResp struct {
	ID            types.BigInt    `json:"id"`              // 标签分类编号
	Name          string          `json:"name"`            // 标签分类名称
	LabelItemList []LabelItemResp `json:"label_item_list"` // 标签分类子项列表
}

// LabelItemResp 标签子项（嵌套）
// @response
type LabelItemResp struct {
	ID     types.BigInt `json:"id"`      // 标签编号
	TypeID types.BigInt `json:"type_id"` // 标签分类编号
	Name   string       `json:"name"`    // 标签名称
	Color  string       `json:"color"`   // 标签颜色
}

// GetElderFeeByIDResp 老人费用详情响应（被 ElderRecord 引用）
// @response
type GetElderFeeByIDResp struct {
	ElderName         string      `json:"elder_name"`          // 老人姓名
	ContractStartTime time.Time   `json:"contract_start_time"` // 合同开始时间
	ContractEndTime   time.Time   `json:"contract_end_time"`   // 合同结束时间
	FeeDetailList     []FeeDetail `json:"fee_detail_list"`     // 费用详情列表
}

// FeeDetail 费用详情（嵌套）
// @response
type FeeDetail struct {
	FeeDate                  string      `json:"fee_date"`                    // 消费时间
	ContractInsideNurseFee   types.Money `json:"contract_inside_nurse_fee"`   // 合同内护理费用
	ContractInsideDishesFee  types.Money `json:"contract_inside_dishes_fee"`  // 合同内餐饮费用
	ContractInsideBedFee     types.Money `json:"contract_inside_bed_fee"`     // 合同内床位费用
	OrderNurseFee            types.Money `json:"order_nurse_fee"`             // 预定护理费用
	OrderDishesFee           types.Money `json:"order_dishes_fee"`            // 预定套餐费用
	ContractOutsideNurseFee  types.Money `json:"contract_outside_nurse_fee"`  // 合同外护理费用
	ContractOutsideDishesFee types.Money `json:"contract_outside_dishes_fee"` // 合同外餐饮费用
	ContractOutsideBedFee    types.Money `json:"contract_outside_bed_fee"`    // 合同外床位费用
	PayableFee               types.Money `json:"payable_fee"`                 // 应缴费用
	ReturnableFee            types.Money `json:"returnable_fee"`              // 应退费用
}

// Date 仅用于标识时间类型字段（go 中使用 time.Time）
var _ = time.Time{}

// EmptyReq 空请求（无参接口复用）
// @request
type EmptyReq struct{}

// EmptyResp 空响应（无返回内容接口复用）
// @response
type EmptyResp struct{}

// IDReq 单编号请求
// @request
type IDReq struct {
	ID *int64 `json:"id" valid:"required"` // 编号
}

// NameReq 单名称请求（对应 Java @RequestParam String name）
// @request
type NameReq struct {
	Name *string `json:"name"` // 名称
}

// StringReq 单字符串请求
// @request
type StringReq struct {
	Value *string `json:"value"` // 值
}
