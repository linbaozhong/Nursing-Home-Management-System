package constant

// VisitEnum 探访状态
type VisitEnum int8

// 探访状态
const (
	VisitEnumStayLeave    VisitEnum = 0 // 在院（未离开）
	VisitEnumAlreadyLeave VisitEnum = 1 // 已离开
)
