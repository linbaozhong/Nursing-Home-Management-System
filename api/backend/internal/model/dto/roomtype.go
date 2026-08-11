package dto

// ============ RoomTypeController 请求 ============

// @request
// PageRoomTypeByKeyQuery 分页查询房间类型请求
type PageRoomTypeByKeyQuery struct {
	PageNum      *int    `json:"page_num"`
	PageSize     *int    `json:"page_size"`
	RoomTypeName *string `json:"room_type_name"`
}

// @request
// OperateRoomTypeQuery 操作房间类型请求（新增/编辑）
type OperateRoomTypeQuery struct {
	ID         *int64   `json:"id"`
	Name       *string  `json:"name"`
	MonthPrice *float64 `json:"month_price"`
}

// ============ RoomTypeController 响应 ============

// @response
// PageRoomTypeByKeyVO 分页查询房间类型响应
type PageRoomTypeByKeyVO struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	MonthPrice float64 `json:"month_price"`
}

// @response
// OperateRoomTypeVo 操作房间类型响应（继承 OperateRoomTypeQuery）
type OperateRoomTypeVo struct {
	OperateRoomTypeQuery
}
