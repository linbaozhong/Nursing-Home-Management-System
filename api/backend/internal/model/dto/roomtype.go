package dto

// ============ RoomTypeController 请求 ============

// @request
// PageRoomTypeByKeyQuery 分页查询房间类型请求
type PageRoomTypeByKeyQuery struct {
	PageNum      *int    `json:"pageNum"`
	PageSize     *int    `json:"pageSize"`
	RoomTypeName *string `json:"roomTypeName"`
}

// @request
// OperateRoomTypeQuery 操作房间类型请求（新增/编辑）
type OperateRoomTypeQuery struct {
	ID         *int64     `json:"id"`
	Name       *string    `json:"name"`
	MonthPrice *float64   `json:"monthPrice"`
}

// ============ RoomTypeController 响应 ============

// @response
// PageRoomTypeByKeyVO 分页查询房间类型响应
type PageRoomTypeByKeyVO struct {
	Rank
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	MonthPrice float64 `json:"monthPrice"`
}

// @response
// OperateRoomTypeVo 操作房间类型响应（继承 OperateRoomTypeQuery）
type OperateRoomTypeVo struct {
	OperateRoomTypeQuery
}