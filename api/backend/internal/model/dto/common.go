package dto

import "time"

// @response
// Result 统一响应包装（对应 Java Result<T>）
type Result struct {
	Code int         `json:"code"` // 状态码
	Msg  string      `json:"msg"` // 提示信息
	Data interface{} `json:"data"` // 响应数据
}

// @response
// PageResult 分页响应包装（对应 Java PageResult<T>）
type PageResult struct {
	PageNum  int         `json:"pageNum"` // 当前页码
	PageSize int         `json:"pageSize"` // 每页条数
	Pages    int         `json:"pages"` // 总页数
	Size     int         `json:"size"` // 当前页条数
	Total    int64       `json:"total"` // 总记录数
	List     interface{} `json:"list"` // 数据列表
}

// @response
// DropDown 下拉列表响应（对应 Java DropDown）
type DropDown struct {
	ID   int64  `json:"id"` // 编号
	Name string `json:"name"` // 名称
}

// @response
// Rank 序号基类（对应 Java Rank）
type Rank struct {
	Rank int64 `json:"rank"` // 序号
}

// @response
// BuildingVO 楼栋-楼层-房间-床位树响应（对应 Java base.BuildingVo）
type BuildingVO struct {
	ID        int64            `json:"id"` // 楼栋编号
	Name      string           `json:"name"` // 楼栋名称
	FloorNum  int              `json:"floorNum"` // 楼层数量
	FloorList []BuildingItemVO `json:"floorList"` // 楼层列表
}

// @response
// BuildingItemVO 楼层（对应 Java base.BuildingVo.BuildingItem）
type BuildingItemVO struct {
	ID       int64        `json:"id"` // 楼层编号
	Name     string       `json:"name"` // 楼层名称
	RoomNum  int          `json:"roomNum"` // 房间数量
	RoomList []FloorItemVO `json:"roomList"` // 房间列表
}

// @response
// FloorItemVO 房间（对应 Java base.BuildingVo.BuildingItem.FloorItem）
type FloorItemVO struct {
	ID     int64        `json:"id"` // 房间编号
	Name   string       `json:"name"` // 房间名称
	BedNum int          `json:"bedNum"` // 床位数量
	BedList []RoomItemVO `json:"bedList"` // 床位列表
}

// @response
// RoomItemVO 床位（对应 Java base.BuildingVo...RoomItem）
type RoomItemVO struct {
	ID        int64  `json:"id"` // 床位编号
	Name      string `json:"name"` // 床位名称
	BedFlag   string `json:"bedFlag"` // 床位状态
	ElderName string `json:"elderName"` // 所住老人
	Sex       string `json:"sex"` // 老人性别
	Age       int    `json:"age"` // 老人年龄
}

// ============ 跨 Controller 共享请求对象 ============

// @request
// OperateEmergencyContactQuery 紧急联系人请求（被 CheckContract、ElderRecord 引用，对应 Java 内部静态类）
type OperateEmergencyContactQuery struct {
	Name        *string `json:"name" valid:"required"` // 紧急联系人姓名
	Phone       *string `json:"phone" valid:"required"` // 紧急联系人电话
	Email       *string `json:"email" valid:"required"` // 紧急联系人邮箱
	Relation    *string `json:"relation" valid:"required"` // 与老人关系
	ReceiveFlag *string `json:"receiveFlag" valid:"required"` // 是否接收消息
}

// @request
// OperateNurseGradeQuery 操作护理等级请求（被 NurseGrade、ElderRecord 引用）
type OperateNurseGradeQuery struct {
	ID            *int64    `json:"id"` // id
	Name          *string   `json:"name" valid:"required"` // 护理等级名称
	Type          *string   `json:"type" valid:"required"` // 护理类型
	MonthPrice    *float64  `json:"monthPrice" valid:"required"` // 月护理费用
	ServiceIDList []int64  `json:"serviceIdList" valid:"required"` // 护理服务编号列表
}

// @request
// OperateServiceQuery 操作服务请求（被 ServiceProject、NurseGrade 引用）
type OperateServiceQuery struct {
	ID           *int64   `json:"id"` // id
	TypeID       *int64   `json:"typeId" valid:"required"` // 服务类型编号
	Name         *string  `json:"name" valid:"required"` // 服务名称
	ChargeMethod *string  `json:"chargeMethod" valid:"required"` // 收费方式
	Price        *float64 `json:"price" valid:"required"` // 服务价格
	NeedDate     *int     `json:"needDate" valid:"required"` // 所需时间
}

// @response
// SetDishesVO 套餐/菜品嵌套响应（被 CateringSet、ElderRecord 引用）
type SetDishesVO struct {
	ID    int64   `json:"id"` // id
	Name  string  `json:"name"` // 菜品名称
	Price float64 `json:"price"` // 菜品价格
}

// ============ 跨 Controller 共享响应对象 ============

// @response
// GetNurseGradeByIDVO 护理等级详情响应（被 ElderRecord 引用）
type GetNurseGradeByIDVO struct {
	OperateNurseGradeQuery
	NurseGradeServiceVOList []NurseGradeServiceVO `json:"nurseGradeServiceVoList"` // 护理等级服务列表
}

// @response
// NurseGradeServiceVO 护理等级服务响应（嵌套，继承 OperateServiceQuery）
type NurseGradeServiceVO struct {
	OperateServiceQuery
}

// @response
// GetCateringSetByIDVO 餐饮套餐详情响应（被 ElderRecord 引用）
type GetCateringSetByIDVO struct {
	ID              int64        `json:"id"` // id
	Name            string       `json:"name"` // 套餐名称
	MonthPrice      float64      `json:"monthPrice"` // 月套餐费用
	SetDishesVOList []SetDishesVO `json:"setDishesVoList"` // 护理等级服务列表
}

// @response
// GetBedByIDVO 床位详情响应（被 ElderRecord 引用）
type GetBedByIDVO struct {
	BedID      int64   `json:"bedId"` // 床位编号
	BedName    string  `json:"bedName"` // 床位名称
	RoomType   string  `json:"roomType"` // 房间类型
	MonthPrice float64 `json:"monthPrice"` // 月床位费用
}

// @response
// GetElderRecordByIDVO 长者档案详情响应（被 ElderRecord 引用）
type GetElderRecordByIDVO struct {
	Name                          string                      `json:"name"` // 姓名
	IDNum                         string                      `json:"idNum"` // 身份证号
	Age                           int                         `json:"age"` // 年龄
	Sex                           string                      `json:"sex"` // 性别
	Phone                         string                      `json:"phone"` // 电话
	Address                       string                      `json:"address"` // 地址
	ElderEmergencyContactByIDVOList []OperateEmergencyContactQuery `json:"elderEmergencyContactByIdVoList"` // 长者紧急联系人
	ElderNurseGradeByIDVO         *GetNurseGradeByIDVO        `json:"elderNurseGradeByIdVo"` // 长者护理等级
	ElderCateringSetByIDVO        *GetCateringSetByIDVO       `json:"elderCateringSetByIdVo"` // 长者餐饮套餐
	ElderBedByIDVO                *GetBedByIDVO               `json:"elderBedByIdVo"` // 长者床位
}

// @response
// GetElderLabelByIDLabelVO 客户标签响应（被 ElderRecord 引用）
type GetElderLabelByIDLabelVO struct {
	ID    int64  `json:"id"` // 编号
	Name  string `json:"name"` // 名称
	Color string `json:"color"` // 颜色
}

// @response
// ListLabelVO 客户标签分类列表响应（被 ElderRecord、Label 引用）
type ListLabelVO struct {
	ID            int64        `json:"id"` // 标签分类编号
	Name          string       `json:"name"` // 标签分类名称
	LabelItemList []LabelItemVO `json:"labelItemList"` // 标签分类子项列表
}

// @response
// LabelItemVO 标签子项（嵌套）
type LabelItemVO struct {
	ID     int64  `json:"id"` // 标签编号
	TypeID int64  `json:"typeId"` // 标签分类编号
	Name   string `json:"name"` // 标签名称
	Color  string `json:"color"` // 标签颜色
}

// @response
// GetElderFeeByIDVO 老人费用详情响应（被 ElderRecord 引用）
type GetElderFeeByIDVO struct {
	ElderName         string      `json:"elderName"` // 老人姓名
	ContractStartTime time.Time   `json:"contractStartTime"` // 合同开始时间
	ContractEndTime   time.Time   `json:"contractEndTime"` // 合同结束时间
	FeeDetailList     []FeeDetail `json:"feeDetailList"` // 费用详情列表
}

// @response
// FeeDetail 费用详情（嵌套）
type FeeDetail struct {
	FeeDate                     string  `json:"feeDate"` // 消费时间
	ContractInsideNurseFee      float64 `json:"contractInsideNurseFee"` // 合同内护理费用
	ContractInsideDishesFee     float64 `json:"contractInsideDishesFee"` // 合同内餐饮费用
	ContractInsideBedFee        float64 `json:"contractInsideBedFee"` // 合同内床位费用
	OrderNurseFee               float64 `json:"orderNurseFee"` // 预定护理费用
	OrderDishesFee              float64 `json:"orderDishesFee"` // 预定套餐费用
	ContractOutsideNurseFee     float64 `json:"contractOutsideNurseFee"` // 合同外护理费用
	ContractOutsideDishesFee    float64 `json:"contractOutsideDishesFee"` // 合同外餐饮费用
	ContractOutsideBedFee       float64 `json:"contractOutsideBedFee"` // 合同外床位费用
	PayableFee                  float64 `json:"payableFee"` // 应缴费用
	ReturnableFee               float64 `json:"returnableFee"` // 应退费用
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
// IDReq 单编号请求（对应 Java @RequestParam Long xxxId）
type IDReq struct {
	ID *int64 `json:"id"` // 编号
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