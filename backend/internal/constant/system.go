package constant

// 系统通用常量
const (
	Subject       = "敬老院管理系统"
	Developer     = "@developer "
	EmperorWen    = "Xxx"
	AnonymousUser = "anonymousUser"

	// 正则
	PhoneRegular = "^1[35789]\\d{9}$"

	// REDIS
	LoginRedis = "login:"

	// OTHER
	TotalLimit = int64(10)
	Comma      = ","
)

// RecipientType 接收方类型
type RecipientType uint8

// 接收方类型
const (
	RecipientElder RecipientType = 1 // 老人
	RecipientStaff RecipientType = 2 // 员工
)

func (r RecipientType) String() string {
	switch r {
	case RecipientElder:
		return "老人"
	case RecipientStaff:
		return "员工"
	default:
		return "未知"
	}
}
