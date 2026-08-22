package dto

// ============ BuildController 请求 ============

// PageBedByKeyReq 分页查询床位请求
// @request
type PageBedByKeyReq struct {
	PageNum  *int    `json:"page_num" valid:"required"`  // 页码
	PageSize *int    `json:"page_size" valid:"required"` // 条数
	BuildID  *int64  `json:"build_id"`                   // 楼栋编号
	FloorID  *int64  `json:"floor_id"`                   // 楼层编号
	RoomID   *int64  `json:"room_id"`                    // 房间编号
	BedFlag  *string `json:"bed_flag"`                   // 床位状态
}

// OperateFloorReq 操作楼层请求
// @request
type OperateFloorReq struct {
	ID         *int64  `json:"id"`                           // id
	BuildingID *int64  `json:"building_id" valid:"required"` // 楼栋编号
	Name       *string `json:"name" valid:"required"`        // 楼层名称
	RoomNum    *int    `json:"room_num" valid:"required"`    // 房间数量
	FloorLimit *int    `json:"floor_limit" valid:"required"` // 楼栋楼层总数限制
}

// OperateRoomReq 操作房间请求
// @request
type OperateRoomReq struct {
	ID        *int64  `json:"id"`                          // id
	TypeID    *int64  `json:"type_id" valid:"required"`    // 房间类型编号
	FloorID   *int64  `json:"floor_id" valid:"required"`   // 楼层编号
	Name      *string `json:"name" valid:"required"`       // 房间名称
	BedNum    *int    `json:"bed_num" valid:"required"`    // 床位数量
	RoomLimit *int    `json:"room_limit" valid:"required"` // 楼层房间总数限制
}

// OperateBedReq 操作床位请求
// @request
type OperateBedReq struct {
	ID       *int64  `json:"id"`                         // id
	RoomID   *int64  `json:"room_id" valid:"required"`   // 房间编号
	Name     *string `json:"name" valid:"required"`      // 床位名称
	BedLimit *int    `json:"bed_limit" valid:"required"` // 房间床位总数限制
}

// DeleteNodeReq 删除节点请求
// @request
type DeleteNodeReq struct {
	ID   *int64  `json:"id" valid:"required"`   // id
	Mark *string `json:"mark" valid:"required"` // 节点标识
}

// ListRoomByKeyReq 获取房间列表请求
// @request
type ListRoomByKeyReq struct {
	BuildingID *int64  `json:"building_id"` // 楼栋编号
	FloorID    *int64  `json:"floor_id"`    // 楼层编号
	ElderName  *string `json:"elder_name"`  // 老人姓名
}

// ============ BuildController 响应 ============

// GetBuildingTreeResp 楼栋树响应（对应 BuildingResp）
// @response
type GetBuildingTreeResp = BuildingResp

// PageBuildingByKeyResp 分页查询楼栋响应
// @response
type PageBuildingByKeyResp struct {
	ID       int64  `json:"id"`        // 楼栋编号
	Name     string `json:"name"`      // 楼栋名称
	FloorNum int    `json:"floor_num"` // 楼层数量
}

// PageBuildingByKeyReq 分页查询楼宇请求
// @request
type PageBuildingByKeyReq struct {
	PageNum  *int    `json:"page_num" valid:"required"`  // 页码
	PageSize *int    `json:"page_size" valid:"required"` // 条数
	Key      *string `json:"key"`                        // 楼宇名称关键字
}

// AddBuildingReq 新增楼宇请求
// @request
type AddBuildingReq struct {
	ID           *int64  `json:"id"`                    // id
	Name         *string `json:"name" valid:"required"` // 楼宇名称
	Remark       *string `json:"remark"`                // 备注
	FloorNum     *int    `json:"floor_num"`             // 楼层数
	ElevatorFlag *string `json:"elevator_flag"`         // 有无电梯
}

// EditBuildingReq 编辑楼宇请求
// @request
type EditBuildingReq struct {
	ID           *int64  `json:"id"`                    // id
	Name         *string `json:"name" valid:"required"` // 楼宇名称
	Remark       *string `json:"remark"`                // 备注
	FloorNum     *int    `json:"floor_num"`             // 楼层数
	ElevatorFlag *string `json:"elevator_flag"`         // 有无电梯
}

// ListBuildingReq 楼宇下拉列表请求
// @request
type ListBuildingReq struct {
	Name *string `json:"name"` // 楼宇名称
}

// @request
// @request
// PageFloorByKeyReq 分页查询楼层请求
type PageFloorByKeyReq struct {
	PageNum  *int    `json:"page_num" valid:"required"`  // 页码
	PageSize *int    `json:"page_size" valid:"required"` // 条数
	BuildID  *int64  `json:"build_id"`                   // 楼栋编号
	Key      *string `json:"key"`                        // 楼层名称关键字
}

// AddFloorReq 新增楼层请求
// @request
type AddFloorReq struct {
	ID         *int64  `json:"id"`                           // id
	BuildingID *int64  `json:"building_id" valid:"required"` // 楼栋编号
	Name       *string `json:"name" valid:"required"`        // 楼层名称
	RoomNum    *int    `json:"room_num" valid:"required"`    // 房间数量
	FloorLimit *int    `json:"floor_limit" valid:"required"` // 楼栋楼层总数限制
}

// EditFloorReq 编辑楼层请求
// @request
type EditFloorReq struct {
	ID         *int64  `json:"id"`                           // id
	BuildingID *int64  `json:"building_id" valid:"required"` // 楼栋编号
	Name       *string `json:"name" valid:"required"`        // 楼层名称
	RoomNum    *int    `json:"room_num" valid:"required"`    // 房间数量
	FloorLimit *int    `json:"floor_limit" valid:"required"` // 楼栋楼层总数限制
}

// PageRoomByKeyReq 分页查询房间请求
// @request
type PageRoomByKeyReq struct {
	PageNum  *int    `json:"page_num" valid:"required"`  // 页码
	PageSize *int    `json:"page_size" valid:"required"` // 条数
	BuildID  *int64  `json:"build_id"`                   // 楼栋编号
	FloorID  *int64  `json:"floor_id"`                   // 楼层编号
	RoomFlag *string `json:"room_flag"`                  // 房间状态
	Key      *string `json:"key"`                        // 房间名称关键字
}

// AddRoomReq 新增房间请求
// @request
type AddRoomReq struct {
	ID        *int64  `json:"id"`                          // id
	TypeID    *int64  `json:"type_id" valid:"required"`    // 房间类型编号
	FloorID   *int64  `json:"floor_id" valid:"required"`   // 楼层编号
	Name      *string `json:"name" valid:"required"`       // 房间名称
	BedNum    *int    `json:"bed_num" valid:"required"`    // 床位数量
	RoomLimit *int    `json:"room_limit" valid:"required"` // 楼层房间总数限制
}

// EditRoomReq 编辑房间请求
// @request
type EditRoomReq struct {
	ID        *int64  `json:"id"`                          // id
	TypeID    *int64  `json:"type_id" valid:"required"`    // 房间类型编号
	FloorID   *int64  `json:"floor_id" valid:"required"`   // 楼层编号
	Name      *string `json:"name" valid:"required"`       // 房间名称
	BedNum    *int    `json:"bed_num" valid:"required"`    // 床位数量
	RoomLimit *int    `json:"room_limit" valid:"required"` // 楼层房间总数限制
}

// GetFloorByBuildingIdReq 根据楼栋编号获取楼层列表请求
// @request
type GetFloorByBuildingIdReq struct {
	BuildingID *int64  `json:"building_id" valid:"required"` // 楼栋编号
	Name       *string `json:"name"`                         // 楼层名称
}

// GetRoomByFloorIdReq 根据楼层编号获取房间列表请求
// @request
type GetRoomByFloorIdReq struct {
	FloorID *int64  `json:"floor_id" valid:"required"` // 楼层编号
	Name    *string `json:"name"`                      // 房间名称
}

// ============ BuildController 响应 Resp ============

// OperateBuildingResp 楼栋详情响应（对应 OperateBuildingResp）
// @response
type OperateBuildingResp struct {
	ID       int64  `json:"id"`        // 楼栋编号
	Name     string `json:"name"`      // 楼栋名称
	FloorNum int    `json:"floor_num"` // 楼层数量
}

// PageFloorByKeyResp 分页查询楼层响应（对应 PageFloorByKeyResp）
// @response
type PageFloorByKeyResp struct {
	ID         int64  `json:"id"`          // 楼层编号
	BuildingID int64  `json:"building_id"` // 楼栋编号
	Name       string `json:"name"`        // 楼层名称
	RoomNum    int    `json:"room_num"`    // 房间数量
}

// OperateFloorResp 楼层详情响应（对应 OperateFloorResp）
// @response
type OperateFloorResp struct {
	ID         int64  `json:"id"`          // 楼层编号
	BuildingID int64  `json:"building_id"` // 楼栋编号
	Name       string `json:"name"`        // 楼层名称
	RoomNum    int    `json:"room_num"`    // 房间数量
}

// PageRoomByKeyResp 分页查询房间响应（对应 PageRoomByKeyResp）
// @response
type PageRoomByKeyResp struct {
	ID      int64  `json:"id"`       // 房间编号
	TypeId  int64  `json:"type_id"`  // 房间类型
	FloorId int64  `json:"floor_id"` // 楼层编号
	Name    string `json:"name"`     // 房间名称
	BedNum  int    `json:"bed_num"`  // 床位数量
}

// OperateRoomResp 房间详情响应（对应 OperateRoomResp）
// @response
type OperateRoomResp struct {
	ID      int64  `json:"id"`       // 房间编号
	TypeId  int64  `json:"type_id"`  // 房间类型
	FloorId int64  `json:"floor_id"` // 楼层编号
	Name    string `json:"name"`     // 房间名称
	BedNum  int    `json:"bed_num"`  // 床位数量
}

// RoomByFloorIdResp 根据楼层获取房间列表响应（对应 FloorItem）
// @response
type RoomByFloorIdResp struct {
	ID     int64  `json:"id"`      // 房间编号
	Name   string `json:"name"`    // 房间名称
	BedNum int    `json:"bed_num"` // 床位数量
}

// PageBedByKeyResp 分页查询床位响应（对应 PageBedByKeyResp）
// @response
type PageBedByKeyResp struct {
	ID      int64  `json:"id"`       // 床位编号
	Name    string `json:"name"`     // 床位名称
	BedFlag string `json:"bed_flag"` // 床位状态
}

// OperateBedResp 床位详情响应（对应 OperateBedResp）
// @response
type OperateBedResp struct {
	ID     int64  `json:"id"`      // 床位编号
	RoomId int64  `json:"room_id"` // 房间编号
	Name   string `json:"name"`    // 床位名称
}
