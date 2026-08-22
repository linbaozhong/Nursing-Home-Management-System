package dto

import "github.com/linbaozhong/gentity/pkg/types"

// ============ RoomTypeController 请求 ============

// PageRoomTypeByKeyReq 分页查询房间类型请求
// @request
type PageRoomTypeByKeyReq struct {
	PageNum      *int    `json:"page_num"`
	PageSize     *int    `json:"page_size"`
	RoomTypeName *string `json:"room_type_name"`
}

// OperateRoomTypeReq 操作房间类型请求（新增/编辑）
// @request
type OperateRoomTypeReq struct {
	ID         *int64       `json:"id"`
	Name       *string      `json:"name"`
	MonthPrice *types.Money `json:"month_price"`
}

// ============ RoomTypeController 响应 ============

// PageRoomTypeByKeyResp 分页查询房间类型响应
// @response
type PageRoomTypeByKeyResp struct {
	ID         int64       `json:"id"`
	Name       string      `json:"name"`
	MonthPrice types.Money `json:"month_price"`
}

// OperateRoomTypeResp 操作房间类型响应（继承 OperateRoomTypeReq）
// @response
type OperateRoomTypeResp struct {
	OperateRoomTypeReq
}
