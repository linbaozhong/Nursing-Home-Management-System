package dto

// ============ BedPanoramaController 请求/响应 ============

// GetBuildingVO 床位全景-楼栋树响应
type GetBuildingVO = BuildingVO

// @request
// ListFloorByBuildingIdQuery 根据楼栋编号获取楼层列表请求
type ListFloorByBuildingIdQuery struct {
	BuildingID *int64 `json:"building_id" form:"building_id"` // 楼栋编号
}

// // @request
// // ListRoomByKeyQuery 获取房间列表请求（对应 Java ListRoomByKeyQuery）
// type ListRoomByKeyQuery struct {
// 	BuildingID *int64  `json:"building_id" form:"building_id"` // 楼栋编号
// 	FloorID    *int64  `json:"floor_id" form:"floor_id"`       // 楼层编号
// 	ElderName  *string `json:"elder_name" form:"elder_name"`   // 老人姓名
// }
