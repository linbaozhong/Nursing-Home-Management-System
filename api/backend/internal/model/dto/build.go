package dto

// ============ BuildController 请求 ============

// @request
// PageBedByKeyQuery 分页查询床位请求
type PageBedByKeyQuery struct {
	PageNum  *int    `json:"pageNum" valid:"required"` // 页码
	PageSize *int    `json:"pageSize" valid:"required"` // 条数
	BuildID  *int64  `json:"buildId"` // 楼栋编号
	FloorID  *int64  `json:"floorId"` // 楼层编号
	RoomID   *int64  `json:"roomId"` // 房间编号
	BedFlag  *string `json:"bedFlag"` // 床位状态
}

// @request
// OperateFloorQuery 操作楼层请求
type OperateFloorQuery struct {
	ID         *int64  `json:"id"` // id
	BuildingID *int64  `json:"buildingId" valid:"required"` // 楼栋编号
	Name       *string `json:"name" valid:"required"` // 楼层名称
	RoomNum    *int    `json:"roomNum" valid:"required"` // 房间数量
	FloorLimit *int    `json:"floorLimit" valid:"required"` // 楼栋楼层总数限制
}

// @request
// OperateRoomQuery 操作房间请求
type OperateRoomQuery struct {
	ID        *int64  `json:"id"` // id
	TypeID    *int64  `json:"typeId" valid:"required"` // 房间类型编号
	FloorID   *int64  `json:"floorId" valid:"required"` // 楼层编号
	Name      *string `json:"name" valid:"required"` // 房间名称
	BedNum    *int    `json:"bedNum" valid:"required"` // 床位数量
	RoomLimit *int    `json:"roomLimit" valid:"required"` // 楼层房间总数限制
}

// @request
// OperateBedQuery 操作床位请求
type OperateBedQuery struct {
	ID       *int64  `json:"id"` // id
	RoomID   *int64  `json:"roomId" valid:"required"` // 房间编号
	Name     *string `json:"name" valid:"required"` // 床位名称
	BedLimit *int    `json:"bedLimit" valid:"required"` // 房间床位总数限制
}

// @request
// DeleteNodeQuery 删除节点请求
type DeleteNodeQuery struct {
	ID   *int64  `json:"id" valid:"required"` // id
	Mark *string `json:"mark" valid:"required"` // 节点标识
}

// @request
// ListRoomByKeyQuery 获取房间列表请求
type ListRoomByKeyQuery struct {
	BuildingID *int64  `json:"buildingId"` // 楼栋编号
	FloorID    *int64  `json:"floorId"` // 楼层编号
	ElderName  *string `json:"elderName"` // 老人姓名
}

// ============ BuildController 响应 ============

// GetBuildingTreeVO 楼栋树响应（对应 BuildingVO）
type GetBuildingTreeVO = BuildingVO

// @request
// PageBuildingByKeyQuery 分页查询楼宇请求
type PageBuildingByKeyQuery struct {
	PageNum  *int    `json:"pageNum" valid:"required"` // 页码
	PageSize *int    `json:"pageSize" valid:"required"` // 条数
	Key      *string `json:"key"` // 楼宇名称关键字
}

// @request
// AddBuildingQuery 新增楼宇请求
type AddBuildingQuery struct {
	ID            *int64  `json:"id"` // id
	Name          *string `json:"name" valid:"required"` // 楼宇名称
	Remark        *string `json:"remark"` // 备注
	FloorNum      *int    `json:"floorNum"` // 楼层数
	ElevatorFlag  *string `json:"elevatorFlag"` // 有无电梯
}

// @request
// EditBuildingQuery 编辑楼宇请求
type EditBuildingQuery struct {
	ID            *int64  `json:"id"` // id
	Name          *string `json:"name" valid:"required"` // 楼宇名称
	Remark        *string `json:"remark"` // 备注
	FloorNum      *int    `json:"floorNum"` // 楼层数
	ElevatorFlag  *string `json:"elevatorFlag"` // 有无电梯
}

// @request
// ListBuildingQuery 楼宇下拉列表请求
type ListBuildingQuery struct {
	Name *string `json:"name"` // 楼宇名称
}

// @request
// @request
// PageFloorByKeyQuery 分页查询楼层请求
type PageFloorByKeyQuery struct {
	PageNum  *int    `json:"pageNum" valid:"required"` // 页码
	PageSize *int    `json:"pageSize" valid:"required"` // 条数
	BuildID  *int64  `json:"buildId"` // 楼栋编号
	Key      *string `json:"key"` // 楼层名称关键字
}

// @request
// AddFloorQuery 新增楼层请求
type AddFloorQuery struct {
	ID          *int64  `json:"id"` // id
	BuildingID  *int64  `json:"buildingId" valid:"required"` // 楼栋编号
	Name        *string `json:"name" valid:"required"` // 楼层名称
	RoomNum     *int    `json:"roomNum" valid:"required"` // 房间数量
	FloorLimit  *int    `json:"floorLimit" valid:"required"` // 楼栋楼层总数限制
}

// @request
// EditFloorQuery 编辑楼层请求
type EditFloorQuery struct {
	ID          *int64  `json:"id"` // id
	BuildingID  *int64  `json:"buildingId" valid:"required"` // 楼栋编号
	Name        *string `json:"name" valid:"required"` // 楼层名称
	RoomNum     *int    `json:"roomNum" valid:"required"` // 房间数量
	FloorLimit  *int    `json:"floorLimit" valid:"required"` // 楼栋楼层总数限制
}

// @request
// PageRoomByKeyQuery 分页查询房间请求
type PageRoomByKeyQuery struct {
	PageNum   *int    `json:"pageNum" valid:"required"` // 页码
	PageSize  *int    `json:"pageSize" valid:"required"` // 条数
	BuildID   *int64  `json:"buildId"` // 楼栋编号
	FloorID   *int64  `json:"floorId"` // 楼层编号
	RoomFlag  *string `json:"roomFlag"` // 房间状态
	Key       *string `json:"key"` // 房间名称关键字
}

// @request
// AddRoomQuery 新增房间请求
type AddRoomQuery struct {
	ID        *int64  `json:"id"` // id
	TypeID    *int64  `json:"typeId" valid:"required"` // 房间类型编号
	FloorID   *int64  `json:"floorId" valid:"required"` // 楼层编号
	Name      *string `json:"name" valid:"required"` // 房间名称
	BedNum    *int    `json:"bedNum" valid:"required"` // 床位数量
	RoomLimit *int    `json:"roomLimit" valid:"required"` // 楼层房间总数限制
}

// @request
// EditRoomQuery 编辑房间请求
type EditRoomQuery struct {
	ID        *int64  `json:"id"` // id
	TypeID    *int64  `json:"typeId" valid:"required"` // 房间类型编号
	FloorID   *int64  `json:"floorId" valid:"required"` // 楼层编号
	Name      *string `json:"name" valid:"required"` // 房间名称
	BedNum    *int    `json:"bedNum" valid:"required"` // 床位数量
	RoomLimit *int    `json:"roomLimit" valid:"required"` // 楼层房间总数限制
}

// @request
// GetFloorByBuildingIdQuery 根据楼栋编号获取楼层列表请求
type GetFloorByBuildingIdQuery struct {
	BuildingID *int64  `json:"buildingId" valid:"required"` // 楼栋编号
	Name       *string `json:"name"` // 楼层名称
}

// @request
// GetRoomByFloorIdQuery 根据楼层编号获取房间列表请求
type GetRoomByFloorIdQuery struct {
	FloorID *int64  `json:"floorId" valid:"required"` // 楼层编号
	Name    *string `json:"name"` // 房间名称
}