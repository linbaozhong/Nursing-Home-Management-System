package dto

// ============ ReserveController 请求 ============

// @request
// PageReserveByKeyQuery 分页查询预定请求
type PageReserveByKeyQuery struct {
	PageNum    *int    `json:"pageNum" valid:"required"` // 页码
	PageSize   *int    `json:"pageSize" valid:"required"` // 条数
	ElderName  *string `json:"elderName"` // 老人姓名
	PayerPhone *string `json:"payerPhone"` // 交款人电话
}

// @request
// AddReserveQuery 新增预定请求
type AddReserveQuery struct {
	BedID        *int64   `json:"bedId" valid:"required"` // 床位编号
	StaffID      *int64   `json:"staffId" valid:"required"` // 营销人员编号
	PayerName    *string  `json:"payerName" valid:"required"` // 交款人姓名
	PayerPhone   *string  `json:"payerPhone" valid:"required"` // 交款人电话
	DueDate      *string  `json:"dueDate" valid:"required"` // 预定到期时间
	Deposit      *float64 `json:"deposit" valid:"required"` // 预付定金
	ElderName    *string  `json:"elderName" valid:"required"` // 老人姓名
	IDNum        *string  `json:"idNum" valid:"required"` // 老人身份证号
	ElderAge     *int     `json:"elderAge" valid:"required"` // 老人年龄
	ElderSex     *string  `json:"elderSex" valid:"required"` // 老人性别
	ElderPhone   *string  `json:"elderPhone" valid:"required"` // 老人电话
	ElderAddress *string  `json:"elderAddress" valid:"required"` // 老人地址
}

// @request
// GetReserveByReserveIDAndElderIDQuery 根据预定编号和老人编号获取预定信息请求
type GetReserveByReserveIDAndElderIDQuery struct {
	ReserveID *int64 `json:"reserveId" valid:"required"` // 预定编号
	ElderID   *int64 `json:"elderId" valid:"required"` // 老人编号
}

// @request
// EditReserveQuery 编辑预定请求
type EditReserveQuery struct {
	ID            *int64   `json:"id"` // id
	BedID         *int64   `json:"bedId" valid:"required"` // 床位编号
	StaffID       *int64   `json:"staffId" valid:"required"` // 营销人员编号
	PayerName     *string  `json:"payerName" valid:"required"` // 交款人姓名
	PayerPhone    *string  `json:"payerPhone" valid:"required"` // 交款人电话
	DueDate       *string  `json:"dueDate" valid:"required"` // 预定到期时间
	Deposit       *float64 `json:"deposit" valid:"required"` // 预付定金
	ElderName     *string  `json:"elderName" valid:"required"` // 老人姓名
	IDNum         *string  `json:"idNum" valid:"required"` // 老人身份证号
	ElderAge      *int     `json:"elderAge" valid:"required"` // 老人年龄
	ElderSex      *string  `json:"elderSex" valid:"required"` // 老人性别
	ElderPhone    *string  `json:"elderPhone" valid:"required"` // 老人电话
	ElderAddress *string  `json:"elderAddress" valid:"required"` // 老人地址
}

// ============ ReserveController 响应 ============

// @response
// PageReserveByKeyVO 分页查询预定响应
type PageReserveByKeyVO struct {
	Rank
	ReserveID   int64   `json:"reserveId"` // 预定编号
	ElderID     int64   `json:"elderId"` // 老人编号
	StaffName   string  `json:"staffName"` // 销售人员姓名
	BedName     string  `json:"bedName"` // 预定床位名称
	ElderName   string  `json:"elderName"` // 老人姓名
	IDNum       string  `json:"idNum"` // 身份证号
	ElderSex    string  `json:"elderSex"` // 老人性别
	ElderAge    int     `json:"elderAge"` // 老人年龄
	PayerPhone  string  `json:"payerPhone"` // 交款人联系电话
	Deposit     float64 `json:"deposit"` // 定金
	ReserveFlag string  `json:"reserveFlag"` // 退款状态
	CheckFlag   string  `json:"checkFlag"` // 入住状态
}

// @response
// GetReserveByReserveIDAndElderIDVO 根据预定编号和老人编号获取预定信息响应
type GetReserveByReserveIDAndElderIDVO struct {
	ElderName    string  `json:"elderName"` // 老人姓名
	IDNum        string  `json:"idNum"` // 身份证号
	ElderSex     string  `json:"elderSex"` // 老人性别
	ElderAge     int     `json:"elderAge"` // 老人年龄
	ElderPhone   string  `json:"elderPhone"` // 老人联系电话
	ElderAddress string  `json:"elderAddress"` // 老人地址
	BedName      string  `json:"bedName"` // 预定床位名称
	StaffID      int64   `json:"staffId"` // 销售人员编号
	PayerName    string  `json:"payerName"` // 交款人姓名
	PayerPhone   string  `json:"payerPhone"` // 交款人联系电话
	DueDate      string  `json:"dueDate"` // 预定到期时间
	Deposit      float64 `json:"deposit"` // 定金
}