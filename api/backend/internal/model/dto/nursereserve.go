package dto

// ============ NurseReserveController 请求 ============

// @request
// PageNurseReserveByKeyQuery 分页查询护理预定请求
type PageNurseReserveByKeyQuery struct {
	PageNum     *int    `json:"page_num" valid:"required"`  // 页码
	PageSize    *int    `json:"page_size" valid:"required"` // 条数
	ElderName   *string `json:"elder_name"`                 // 老人姓名
	ServiceName *string `json:"service_name"`               // 项目名称
	BedName     *string `json:"bed_name"`                   // 床位名称
}

// @request
// AddNurseReserveQuery 新增护理预定请求
type AddNurseReserveQuery struct {
	ElderID      *int64   `json:"elder_id" valid:"required"`      // 老人编号
	ServiceName  *string  `json:"service_name" valid:"required"`  // 项目名称
	NeedDate     *int     `json:"need_date" valid:"required"`     // 所需时间
	ServicePrice *float64 `json:"service_price" valid:"required"` // 服务费用
	ChargeMethod *string  `json:"charge_method" valid:"required"` // 收费方式
	Frequency    *int     `json:"frequency" valid:"required"`     // 服务次数
	PayAmount    *float64 `json:"pay_amount" valid:"required"`    // 支付总额
}

// @request
// ExecuteNurseReserveQuery 执行护理预定请求
type ExecuteNurseReserveQuery struct {
	ID        *int64  `json:"id"`                          // id
	StaffID   *int64  `json:"staff_id" valid:"required"`   // 服务员工编号
	NurseDate *string `json:"nurse_date" valid:"required"` // 护理时间
}

// ============ NurseReserveController 响应 ============

// @response
// PageNurseReserveByKeyVO 分页查询护理预定响应
type PageNurseReserveByKeyVO struct {
	Rank
	ID           int64   `json:"id"`            // id
	ElderName    string  `json:"elder_name"`    // 老人姓名
	BedName      string  `json:"bed_name"`      // 床位名称
	ServiceName  string  `json:"service_name"`  // 项目名称
	NeedDate     int     `json:"need_date"`     // 所需时间
	ServicePrice float64 `json:"service_price"` // 服务费用
	ChargeMethod string  `json:"charge_method"` // 收费方式
	Frequency    int     `json:"frequency"`     // 服务次数
	PayAmount    float64 `json:"pay_amount"`    // 支付总额
	NurseDate    string  `json:"nurse_date"`    // 护理时间
	OrderFlag    string  `json:"order_flag"`    // 订单状态
}

// @request
// GetNurseReserveByReserveIdAndElderIdQuery 按护理预定/老人编号获取护理预定请求
type GetNurseReserveByReserveIdAndElderIdQuery struct {
	ReserveID *int64 `json:"reserve_id" valid:"required"` // 护理预定编号
	ElderID   *int64 `json:"elder_id" valid:"required"`   // 老人编号
}

// @request
// EditNurseReserveQuery 编辑护理预定请求
type EditNurseReserveQuery struct {
	ID           *int64   `json:"id"`                             // id
	ElderID      *int64   `json:"elder_id" valid:"required"`      // 老人编号
	ReserveDate  *string  `json:"reserve_date" valid:"required"`  // 预定时间
	ServiceID    *int64   `json:"service_id" valid:"required"`    // 服务编号
	NeedDate     *int     `json:"need_date" valid:"required"`     // 所需时间
	ChargeMethod *string  `json:"charge_method" valid:"required"` // 收费方式
	Frequency    *int     `json:"frequency" valid:"required"`     // 服务次数
	ServicePrice *float64 `json:"service_price" valid:"required"` // 服务费用
	PayAmount    *float64 `json:"pay_amount" valid:"required"`    // 支付总额
}
