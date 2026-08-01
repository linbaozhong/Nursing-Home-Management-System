package service

import (
	"context"
	"errors"

	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblbed"
	"api/internal/model/define/table/tblroomtype"
	"api/internal/model/do"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

// 床位状态/删除标记的占位常量，Go 侧暂无对应枚举包，待补充
const (
	bedFlagIdle = "0"  // 对应 Java BedEnum.IDLE
	delFlagYes  = "Y"  // 对应 Java YesNoEnum.YES
	markBuilding = "building"
	markFloor    = "floor"
	markRoom     = "room"
)

type build struct{}

var Build = &build{}

// PageBuildingByKey 分页查询楼宇
// 对应 Java: BuildServiceImpl.pageBuildingByKey -> BuildingMapper.listBuildingByKey
// SQL: SELECT * FROM building WHERE (building_name LIKE %key%) [可选] ORDER BY create_time DESC
// todo: 楼宇分页查询 - dao.Building(db) 条件 + 分页, 结果赋值 out
func (b *build) PageBuildingByKey(ctx context.Context, in *dto.PageBuildingByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 building 表并分页
	return nil
}

// GetBuildingById 根据编号获取楼宇
// 对应 Java: BuildServiceImpl.getBuildingById -> buildingMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.Building(db).GetByID(ctx, types.BigInt(in.ID))
func (b *build) GetBuildingById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.Building(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// AddBuilding 新增楼宇
// 对应 Java: BuildServiceImpl.addBuilding -> buildingMapper.insertSelective
// todo: 标准 CRUD - dao.Building(db).InsertOne 写入 building 表
func (b *build) AddBuilding(ctx context.Context, in *dto.AddBuildingQuery, out *dto.EmptyResp) error {
	// todo: bean := do.NewBuilding(); 填充 in; dao.Building(db).InsertOne(ctx, bean)
	return nil
}

// EditBuilding 编辑楼宇
// 对应 Java: BuildServiceImpl.editBuilding -> buildingMapper.updateByPrimaryKeySelective
// todo: 标准 CRUD - 按主键更新 building 表
func (b *build) EditBuilding(ctx context.Context, in *dto.EditBuildingQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<building>.BuildingName.Value(in.BuildingName),
	}
	_, e := dao.Building(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteBuilding 删除楼宇
// 对应 Java: BuildServiceImpl.deleteBuilding -> buildingMapper.deleteByPrimaryKey
// todo: 标准 CRUD - dao.Building(db).DeleteById(ctx, types.BigInt(in.ID))
func (b *build) DeleteBuilding(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.Building(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}

// PageFloorByKey 分页查询楼层（联表 building 获取楼宇名）
// 对应 Java: BuildServiceImpl.pageFloorByKey -> FloorMapper.listFloorByKey
// SQL: SELECT f.*, b.building_name FROM floor f LEFT JOIN building b ON b.id = f.building_id
//
//	WHERE (b.building_name LIKE %key% OR f.floor_name LIKE %key%) [可选]
//
// todo: 楼层分页查询 - 联表 building + 分页, 结果赋值 out
func (b *build) PageFloorByKey(ctx context.Context, in *dto.PageFloorByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 floor 联表 building 并分页
	return nil
}

// GetFloorByBuildingId 根据楼宇编号获取楼层
// 对应 Java: BuildServiceImpl.getFloorByBuildingId -> floorMapper.selectByBuildingId
// SQL: SELECT * FROM floor WHERE building_id = #{buildingId}
// todo: 标准查询 - dao.Floor(db).List(ace.Where(tbl<floor>.BuildingId.Eq(in.BuildingId)))
func (b *build) GetFloorByBuildingId(ctx context.Context, in *dto.GetFloorByBuildingIdQuery, out *dto.EmptyResp) error {
	// todo: list, e := dao.Floor(db).List(ctx, ace.Where(tbl<floor>.BuildingId.Eq(in.BuildingId)))
	return nil
}

// AddFloor 新增楼层
// 对应 Java: BuildServiceImpl.addFloor -> floorMapper.insertSelective
// todo: 标准 CRUD - dao.Floor(db).InsertOne 写入 floor 表(含 buildingId/floorName 等)
func (b *build) AddFloor(ctx context.Context, in *dto.AddFloorQuery, out *dto.EmptyResp) error {
	// todo: bean := do.NewFloor(); 填充 in; dao.Floor(db).InsertOne(ctx, bean)
	return nil
}

// EditFloor 编辑楼层
// 对应 Java: BuildServiceImpl.editFloor -> floorMapper.updateByPrimaryKeySelective
// todo: 标准 CRUD - 按主键更新 floor 表
func (b *build) EditFloor(ctx context.Context, in *dto.EditFloorQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<floor>.FloorName.Value(in.FloorName),
	}
	_, e := dao.Floor(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteFloor 删除楼层
// 对应 Java: BuildServiceImpl.deleteFloor -> floorMapper.deleteByPrimaryKey
// todo: 标准 CRUD - dao.Floor(db).DeleteById(ctx, types.BigInt(in.ID))
func (b *build) DeleteFloor(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.Floor(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}

// PageRoomByKey 分页查询房间（联表 floor、building 获取楼层名与楼宇名）
// 对应 Java: BuildServiceImpl.pageRoomByKey -> RoomMapper.listRoomByKey
// SQL: SELECT r.*, f.floor_name, b.building_name FROM room r
//
//	LEFT JOIN floor f ON f.id = r.floor_id
//	LEFT JOIN building b ON b.id = f.building_id
//	WHERE (b.building_name LIKE %key% OR f.floor_name LIKE %key% OR r.room_name LIKE %key%) [可选]
//
// todo: 房间分页查询 - 联表 floor/building + 分页, 结果赋值 out
func (b *build) PageRoomByKey(ctx context.Context, in *dto.PageRoomByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 room 联表 floor/building 并分页
	return nil
}

// GetRoomByFloorId 根据楼层编号获取房间
// 对应 Java: BuildServiceImpl.getRoomByFloorId -> roomMapper.selectByFloorId
// SQL: SELECT * FROM room WHERE floor_id = #{floorId}
// todo: 标准查询 - dao.Room(db).List(ace.Where(tbl<room>.FloorId.Eq(in.FloorId)))
func (b *build) GetRoomByFloorId(ctx context.Context, in *dto.GetRoomByFloorIdQuery, out *dto.EmptyResp) error {
	// todo: list, e := dao.Room(db).List(ctx, ace.Where(tbl<room>.FloorId.Eq(in.FloorId)))
	return nil
}

// AddRoom 新增房间
// 对应 Java: BuildServiceImpl.addRoom -> roomMapper.insertSelective
// todo: 标准 CRUD - dao.Room(db).InsertOne 写入 room 表(含 floorId/roomName/roomStatus 等)
func (b *build) AddRoom(ctx context.Context, in *dto.AddRoomQuery, out *dto.EmptyResp) error {
	// todo: bean := do.NewRoom(); 填充 in; dao.Room(db).InsertOne(ctx, bean)
	return nil
}

// EditRoom 编辑房间
// 对应 Java: BuildServiceImpl.editRoom -> roomMapper.updateByPrimaryKeySelective
// todo: 标准 CRUD - 按主键更新 room 表
func (b *build) EditRoom(ctx context.Context, in *dto.EditRoomQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<room>.RoomName.Value(in.RoomName),
	}
	_, e := dao.Room(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteRoom 删除房间
// 对应 Java: BuildServiceImpl.deleteRoom -> roomMapper.deleteByPrimaryKey
// todo: 标准 CRUD - dao.Room(db).DeleteById(ctx, types.BigInt(in.ID))
func (b *build) DeleteRoom(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.Room(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}

// GetFloorById 根据楼层编号获取楼层详情
// 对应 Java: BuildServiceImpl.getFloorById -> floorMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.Floor(db).GetByID(ctx, types.BigInt(in.ID)) 赋值 out
func (b *build) GetFloorById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.Floor(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// PageBedByKey 分页查询床位（按关键词）
// 对应 Java: BuildServiceImpl.pageBedByKey -> BedMapper.listBedByKey
// SQL: SELECT * FROM bed WHERE bed_name LIKE %key% [可选] ORDER BY create_time DESC
// todo: 床位分页查询 - dao.Bed(db) 条件 + 分页, 结果赋值 out
func (b *build) PageBedByKey(ctx context.Context, in *dto.PageBedByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 bed 表并按 bedName 等条件分页
	return nil
}

// AddBed 新增床位
// 对应 Java: BuildServiceImpl.addBed -> bedMapper.insertSelective
// todo: 标准 CRUD - dao.Bed(db).InsertOne 写入 bed 表
func (b *build) AddBed(ctx context.Context, in *dto.OperateBedQuery, out *dto.EmptyResp) error {
	bean := &do.Bed{}
	// todo: 将 in(OperateBedQuery) 字段映射到 bean；BedLimit 对应房间床位总数，需确认是否写入 bed 表
	bean.RoomId = types.BigInt(in.RoomID)
	bean.BedName = in.Name
	bean.BedFlag = bedFlagIdle
	bean.DelFlag = "N"
	_, e := dao.Bed(db).InsertOne(ctx, bean)
	if e != nil {
		return e
	}
	return nil
}

// GetBedById 根据床位编号获取床位详情
// 对应 Java: BuildServiceImpl.getBedById -> bedMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.Bed(db).GetByID(ctx, types.BigInt(in.ID)) 赋值 out
func (b *build) GetBedById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.Bed(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// EditBed 编辑床位
// 对应 Java: BuildServiceImpl.editBed -> bedMapper.updateByPrimaryKeySelective
// todo: 标准 CRUD - 按主键更新 bed 表
func (b *build) EditBed(ctx context.Context, in *dto.OperateBedQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tblbed.BedName.Set(in.Name)
	}
	_, e := dao.Bed(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteBed 删除床位（逻辑删除，需校验床位空闲）
// 对应 Java: BuildServiceImpl.deleteBed
// 逻辑：判断 bed_flag != IDLE 则抛 BED_NOT_IDLE；否则 del_flag 置 YES
func (b *build) DeleteBed(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	// 判断该床位是否被占用
	bed, has, e := dao.Bed(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("床位不存在")
	}
	if bed.BedFlag != bedFlagIdle {
		return errors.New("该床位被占用，删除失败") // 对应 Java ExceptionEnum.BED_NOT_IDLE
	}
	// 逻辑删除：del_flag = YES
	_, e = dao.Bed(db).UpdateById(ctx, types.BigInt(in.ID),
		ace.Where(tblbed.DelFlag.Set(delFlagYes)),
	)
	if e != nil {
		return e
	}
	return nil
}

// GetNoBedTreeAndPageBedByKey 查询空闲床位树 + 分页床位
// 对应 Java: BuildServiceImpl.getNoBedTreeAndPageBedByKey
// 逻辑：组装 楼栋->楼层->房间 树（仅含空闲 idle 床位），并对床位做分页
func (b *build) GetNoBedTreeAndPageBedByKey(ctx context.Context, in *dto.PageBedByKeyQuery, out *dto.EmptyResp) error {
	// todo: 1) 查所有楼栋/楼层/房间（按 tree 条件）；2) 查空闲床位(bed_flag=idle, del_flag=N)并按关键词分页；3) 组装树结构返回 out
	return nil
}

// DeleteNode 删除楼栋/楼层/房间节点（按 mark 级联删除）
// 对应 Java: BuildServiceImpl.deleteNode
// 逻辑：根据 mark(building/floor/room) 调用对应 func 逻辑删除节点，删除前需校验无占用床位(NODE_BED_NOT_IDLE)
func (b *build) DeleteNode(ctx context.Context, in *dto.DeleteNodeQuery, out *dto.EmptyResp) error {
	switch in.Mark {
	case markBuilding:
		// todo: 校验该楼栋下无占用床位 -> dao.Building(db).UpdateById 置 del_flag=YES（逻辑删除）
	case markFloor:
		// todo: 校验该楼层下无占用床位 -> dao.Floor(db).UpdateById 置 del_flag=YES
	case markRoom:
		// todo: 校验该房间下无占用床位 -> dao.Room(db).UpdateById 置 del_flag=YES
	default:
		return errors.New("该节点标记不存在") // 对应 Java ExceptionEnum.NODE_MARK_NOT_EXIST
	}
	return nil
}

// ListRoomType 房间类型下拉列表
// 对应 Java: BuildServiceImpl.listRoomType -> roomTypeFunc.listNotDelRoomType
// 逻辑：查 room_type 表(del_flag=N) 转为 DropDown 列表
func (b *build) ListRoomType(ctx context.Context, in *dto.EmptyReq, out *dto.EmptyResp) error {
	list, _, e := dao.RoomType(db).List(ctx, ace.Where(tblroomtype.DelFlag.Eq("N")))
	if e != nil {
		return e
	}
	// todo: 将 list(room_type) 转为 []dto.DropDown{ID,Name} 并赋值 out
	_ = list
	return nil
}

// GetRoomById 根据房间编号获取房间详情
// 对应 Java: BuildServiceImpl.getRoomById -> roomMapper.selectById
// todo: 标准 CRUD - dao.Room(db).GetByID(ctx, types.BigInt(in.ID)) 赋值 out
func (b *build) GetRoomById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.Room(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}
