package constant

// VisitStatus 来访状态
type VisitStatus uint8

// 来访状态
const (
	VisitStayLeave    VisitStatus = 0 // 待离开
	VisitAlreadyLeave VisitStatus = 1 // 已离开
)

func (v VisitStatus) String() string {
	switch v {
	case VisitStayLeave:
		return "待离开"
	case VisitAlreadyLeave:
		return "已离开"
	default:
		return "未知"
	}
}
