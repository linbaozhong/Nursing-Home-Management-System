package constant

// AuditStatus 审计状态
type AuditStatus uint8

// 审计状态
const (
	AuditStay    AuditStatus = 0 // 待审核
	Auditing     AuditStatus = 1 // 审核中
	AuditPass    AuditStatus = 2 // 通过
	AuditNotPass AuditStatus = 3 // 不通过
)

func (a AuditStatus) String() string {
	switch a {
	case AuditNotPass:
		return "不通过"
	case AuditStay:
		return "待审核"
	case Auditing:
		return "审核中"
	case AuditPass:
		return "通过"
	default:
		return "未知"
	}
}
