package service

import (
	"context"
	"errors"

	"api/internal/constant"
	"api/internal/lib"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblbed"
	"api/internal/model/define/table/tblbuilding"
	"api/internal/model/define/table/tblfloor"
	"api/internal/model/define/table/tblroom"
	"api/internal/model/define/table/tblroomtype"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

const (
	// delFlagYes   = "Y" // 对应 Java YesNoEnum.YES
	markBuilding = "building"
	markFloor    = "floor"
	markRoom     = "room"
)

type build struct{}

var Build = &build{}

// ==================== 公共查询辅助（对应 Java BuildingFunc/FloorFunc/RoomFunc/BedFunc） ====================

// getBuildingByName 按名称查未删除楼宇（用于防重名），无则返回 nil
func (b *build) getBuildingByName(ctx context.Context, name string) (*do.Building, bool, error) {
	return dao.Building(db).Get(ctx, ace.Where(tblbuilding.Name.Eq(name), tblbuilding.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblbuilding.DelFlag.Eq(constant.YesNoNo)))
}

// getFloorByName 按楼宇+名称查未删除楼层（用于防重名），无则返回 nil
func (b *build) getFloorByName(ctx context.Context, buildingID int64, name string) (*do.Floor, bool, error) {
	return dao.Floor(db).Get(ctx, ace.Where(
		tblfloor.BuildingId.Eq(buildingID),
		tblfloor.Name.Eq(name),
		tblfloor.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
		tblfloor.DelFlag.Eq(constant.YesNoNo),
	))
}

// getRoomByName 按楼层+名称查未删除房间（用于防重名），无则返回 nil
func (b *build) getRoomByName(ctx context.Context, floorID int64, name string) (*do.Room, bool, error) {
	return dao.Room(db).Get(ctx, ace.Where(
		tblroom.FloorId.Eq(floorID),
		tblroom.Name.Eq(name),
		tblroom.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
		tblroom.DelFlag.Eq(constant.YesNoNo),
	))
}

// getBedByName 按房间+名称查未删除床位（用于防重名），无则返回 nil
func (b *build) getBedByName(ctx context.Context, roomID int64, name string) (*do.Bed, bool, error) {
	return dao.Bed(db).Get(ctx, ace.Where(
		tblbed.RoomId.Eq(roomID),
		tblbed.Name.Eq(name),
		tblbed.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
		tblbed.DelFlag.Eq(constant.YesNoNo),
	))
}

// hasOccupiedBed 校验某节点下是否存在被占用的床位（bed_flag != idle 且未删除）
func (b *build) hasOccupiedBed(ctx context.Context, scope string, id int64) (bool, error) {
	var beds []*do.Bed
	var e error
	switch scope {
	case markBuilding:
		floors, _, e := dao.Floor(db).List(ctx, ace.Where(tblfloor.BuildingId.Eq(id), tblfloor.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblfloor.DelFlag.Eq(constant.YesNoNo)))
		if e != nil {
			return false, e
		}
		if len(floors) == 0 {
			return false, nil
		}
		floorIds := make([]any, 0, len(floors))
		for _, f := range floors {
			floorIds = append(floorIds, f.Id)
		}
		rooms, e := dao.Room(db).GetByIds(ctx, floorIds)
		if e != nil {
			return false, e
		}
		if len(rooms) == 0 {
			return false, nil
		}
		roomIds := make([]any, 0, len(rooms))
		for _, r := range rooms {
			roomIds = append(roomIds, r.Id)
		}
		beds, e = dao.Bed(db).GetByIds(ctx, roomIds)
	case markFloor:
		rooms, _, e := dao.Room(db).List(ctx, ace.Where(tblroom.FloorId.Eq(id), tblroom.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblroom.DelFlag.Eq(constant.YesNoNo)))
		if e != nil {
			return false, e
		}
		if len(rooms) == 0 {
			return false, nil
		}
		roomIds := make([]any, 0, len(rooms))
		for _, r := range rooms {
			roomIds = append(roomIds, r.Id)
		}
		beds, e = dao.Bed(db).GetByIds(ctx, roomIds)
	case markRoom:
		beds, _, e = dao.Bed(db).List(ctx, ace.Where(tblbed.RoomId.Eq(id), tblbed.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblbed.DelFlag.Eq(constant.YesNoNo)))
	default:
		return false, nil
	}
	if e != nil {
		return false, e
	}
	for _, bed := range beds {
		if bed.BedFlag != types.Int8(constant.BedIdle) {
			return true, nil
		}
	}
	return false, nil
}

// buildTree 组装 楼栋→楼层→房间→床位 树（对应 Java generateBuildingSimpleTree）
func (b *build) buildTree(buildings []*do.Building, floors []*do.Floor, rooms []*do.Room, beds []*do.Bed) []dto.BuildingResp {
	roomMap := make(map[uint64][]*do.Bed)
	for _, bed := range beds {
		rid := uint64(bed.RoomId)
		roomMap[rid] = append(roomMap[rid], bed)
	}
	floorMap := make(map[uint64][]*do.Room)
	for _, room := range rooms {
		fid := uint64(room.FloorId)
		floorMap[fid] = append(floorMap[fid], room)
	}
	buildingMap := make(map[uint64][]*do.Floor)
	for _, floor := range floors {
		bid := uint64(floor.BuildingId)
		buildingMap[bid] = append(buildingMap[bid], floor)
	}

	tree := make([]dto.BuildingResp, 0, len(buildings))
	for _, bd := range buildings {
		vo := dto.BuildingResp{
			ID:       types.BigInt(bd.Id),
			Name:     string(bd.Name),
			FloorNum: int(bd.FloorNum),
		}
		for _, fl := range buildingMap[uint64(bd.Id)] {
			fi := dto.BuildingItemResp{
				ID:      types.BigInt(fl.Id),
				Name:    string(fl.Name),
				RoomNum: int(fl.RoomNum),
			}
			for _, rm := range floorMap[uint64(fl.Id)] {
				ri := dto.FloorItemResp{
					ID:     types.BigInt(rm.Id),
					Name:   string(rm.Name),
					BedNum: int(rm.BedNum),
				}
				for _, bd2 := range roomMap[uint64(rm.Id)] {
					ri.BedList = append(ri.BedList, dto.RoomItemResp{
						ID:      types.BigInt(bd2.Id),
						Name:    string(bd2.Name),
						BedFlag: constant.YesNo(bd2.BedFlag).String(),
					})
				}
				fi.RoomList = append(fi.RoomList, ri)
			}
			vo.FloorList = append(vo.FloorList, fi)
		}
		tree = append(tree, vo)
	}
	return tree
}

// ==================== 楼宇 ====================

// PageBuildingByKey 分页查询楼宇
// 对应 Java: BuildServiceImpl.pageBuildingByKey -> BuildingMapper.listBuildingByKey
func (b *build) PageBuildingByKey(ctx context.Context, in *dto.PageBuildingByKeyReq, out *[]dto.PageBuildingByKeyResp) error {
	q := db.Table(tblbuilding.TableName).
		Where(tblbuilding.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblbuilding.DelFlag.Eq(constant.YesNoNo)).
		Desc(tblbuilding.Id)
	if in.Key != nil {
		q.And(tblbuilding.Name.Like(*in.Key))
	}
	e := q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblbuilding.Id,
			tblbuilding.Name,
			tblbuilding.FloorNum,
		).
		Select().
		Gets(ctx, out)
	return e
}

// GetBuildingById 根据编号获取楼宇
// 对应 Java: BuildServiceImpl.getBuildingById -> buildingMapper.selectByPrimaryKey
func (b *build) GetBuildingById(ctx context.Context, in *dto.IDReq, out *dto.OperateBuildingResp) error {
	obj, has, e := dao.Building(db).GetByID(ctx, types.BigInt(*in.ID))
	if e != nil {
		return e
	}
	if !has || obj == nil {
		return errors.New("楼宇不存在")
	}
	out.ID = types.BigInt(obj.Id)
	out.Name = obj.Name.String()
	out.FloorNum = int(obj.FloorNum)
	return nil
}

// AddBuilding 新增楼宇
// 对应 Java: BuildServiceImpl.addBuilding -> 防重名 + insertSelective
func (b *build) AddBuilding(ctx context.Context, in *dto.AddBuildingReq, out *dto.EmptyResp) error {
	if in.Name != nil {
		if exist, _, e := b.getBuildingByName(ctx, *in.Name); e != nil {
			return e
		} else if exist != nil {
			return errors.New("楼宇名称已存在")
		}
	}
	bean := do.NewBuilding()
	bean.TenantId = types.BigInt(lib.TenantID(ctx))
	if in.Name != nil {
		bean.Name = types.String(*in.Name)
	}
	if in.FloorNum != nil {
		bean.FloorNum = types.Int32(int32(*in.FloorNum))
	}
	bean.DelFlag = types.Int8(constant.YesNoNo)
	_, e := dao.Building(db).InsertOne(ctx, bean)
	return e
}

// EditBuilding 编辑楼宇
// 对应 Java: BuildServiceImpl.editBuilding -> 防重名 + updateByPrimaryKeySelective
func (b *build) EditBuilding(ctx context.Context, in *dto.EditBuildingReq, out *dto.EmptyResp) error {
	if in.Name != nil {
		if exist, _, e := b.getBuildingByName(ctx, *in.Name); e != nil {
			return e
		} else if exist != nil && int64(exist.Id) != *in.ID {
			return errors.New("楼宇名称已存在")
		}
	}
	sets := []dialect.Setter{}
	if in.Name != nil {
		sets = append(sets, tblbuilding.Name.Set(types.String(*in.Name)))
	}
	if in.FloorNum != nil {
		sets = append(sets, tblbuilding.FloorNum.Set(types.Int32(int32(*in.FloorNum))))
	}
	if len(sets) > 0 {
		if _, e := dao.Building(db).UpdateById(ctx, types.BigInt(*in.ID), sets...); e != nil {
			return e
		}
	}
	return nil
}

// DeleteBuilding 删除楼宇（逻辑删除）
// 对应 Java: BuildServiceImpl.deleteNode -> buildingFunc.deleteBuildingNode
func (b *build) DeleteBuilding(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	if occupied, e := b.hasOccupiedBed(ctx, markBuilding, *in.ID); e != nil {
		return e
	} else if occupied {
		return errors.New("该楼栋下存在占用床位，无法删除")
	}
	_, e := dao.Building(db).UpdateById(ctx, types.BigInt(*in.ID), tblbuilding.DelFlag.Set(constant.YesNoYes))
	return e
}

// ==================== 楼层 ====================

// PageFloorByKey 分页查询楼层
// 对应 Java: BuildServiceImpl.pageFloorByKey -> FloorMapper.listFloorByKey
func (b *build) PageFloorByKey(ctx context.Context, in *dto.PageFloorByKeyReq, out *[]dto.PageFloorByKeyVO) error {
	q := db.Table(tblfloor.TableName).Where(tblfloor.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblfloor.DelFlag.Eq(constant.YesNoNo))
	if in.BuildID != nil {
		q.And(tblfloor.BuildingId.Eq(types.BigInt(*in.BuildID)))
	}
	if in.Key != nil {
		q.And(tblfloor.Name.Like(*in.Key))
	}
	e := q.Desc(tblfloor.Id).
		Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblfloor.Id,
			tblfloor.BuildingId,
			tblfloor.Name,
			tblfloor.RoomNum,
		).
		Select().
		Gets(ctx, out)
	return e
}

// GetFloorByBuildingId 根据楼宇编号获取楼层
// 对应 Java: BuildServiceImpl.getFloorByBuildingId -> floorMapper.selectByBuildingId
func (b *build) GetFloorByBuildingId(ctx context.Context, in *dto.GetFloorByBuildingIdReq, out *[]dto.DropDown) error {
	q := ace.Where(tblfloor.BuildingId.Eq(types.BigInt(*in.BuildingID)), tblfloor.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblfloor.DelFlag.Eq(constant.YesNoNo))
	if in.Name != nil {
		q = q.And(tblfloor.Name.Like(*in.Name))
	}
	list, _, e := dao.Floor(db).List(ctx, q)
	if e != nil {
		return e
	}
	*out = make([]dto.DropDown, 0, len(list))
	for _, floor := range list {
		*out = append(*out, dto.DropDown{
			ID:   types.BigInt(floor.Id),
			Name: floor.Name.String(),
		})
	}
	return nil
}

// GetFloorById 根据楼层编号获取楼层详情
// 对应 Java: BuildServiceImpl.getFloorById -> floorMapper.selectByPrimaryKey
func (b *build) GetFloorById(ctx context.Context, in *dto.IDReq, out *dto.OperateFloorVO) error {
	obj, has, e := dao.Floor(db).GetByID(ctx, types.BigInt(*in.ID))
	if e != nil {
		return e
	}
	if !has || obj == nil {
		return errors.New("楼层不存在")
	}
	out.ID = types.BigInt(obj.Id)
	out.BuildingID = types.BigInt(obj.BuildingId)
	out.Name = obj.Name.String()
	out.RoomNum = int(obj.RoomNum)
	return nil
}

// AddFloor 新增楼层
// 对应 Java: BuildServiceImpl.addFloor -> 防重名 + insertSelective
func (b *build) AddFloor(ctx context.Context, in *dto.AddFloorReq, out *dto.EmptyResp) error {
	if in.Name != nil {
		if exist, _, e := b.getFloorByName(ctx, *in.BuildingID, *in.Name); e != nil {
			return e
		} else if exist != nil {
			return errors.New("楼层名称已存在")
		}
	}
	bean := do.NewFloor()
	bean.TenantId = types.BigInt(lib.TenantID(ctx))
	bean.BuildingId = types.BigInt(*in.BuildingID)
	if in.Name != nil {
		bean.Name = types.String(*in.Name)
	}
	if in.RoomNum != nil {
		bean.RoomNum = types.Int32(int32(*in.RoomNum))
	}
	bean.DelFlag = types.Int8(constant.YesNoNo)
	_, e := dao.Floor(db).InsertOne(ctx, bean)
	return e
}

// EditFloor 编辑楼层
// 对应 Java: BuildServiceImpl.editFloor -> 防重名 + updateByPrimaryKeySelective
func (b *build) EditFloor(ctx context.Context, in *dto.EditFloorReq, out *dto.EmptyResp) error {
	if in.Name != nil {
		if exist, _, e := b.getFloorByName(ctx, *in.BuildingID, *in.Name); e != nil {
			return e
		} else if exist != nil && int64(exist.Id) != *in.ID {
			return errors.New("楼层名称已存在")
		}
	}
	sets := []dialect.Setter{
		tblfloor.BuildingId.Set(*in.BuildingID),
	}
	if in.Name != nil {
		sets = append(sets, tblfloor.Name.Set(*in.Name))
	}
	if in.RoomNum != nil {
		sets = append(sets, tblfloor.RoomNum.Set(types.Int32(int32(*in.RoomNum))))
	}
	_, e := dao.Floor(db).UpdateById(ctx, types.BigInt(*in.ID), sets...)
	return e
}

// DeleteFloor 删除楼层（逻辑删除）
// 对应 Java: BuildServiceImpl.deleteNode -> floorFunc.deleteFloorNode
func (b *build) DeleteFloor(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	if occupied, e := b.hasOccupiedBed(ctx, markFloor, *in.ID); e != nil {
		return e
	} else if occupied {
		return errors.New("该楼层下存在占用床位，无法删除")
	}
	_, e := dao.Floor(db).UpdateById(ctx, types.BigInt(*in.ID), tblfloor.DelFlag.Set(constant.YesNoYes))
	return e
}

// ==================== 房间 ====================

// PageRoomByKey 分页查询房间
// 对应 Java: BuildServiceImpl.pageRoomByKey -> RoomMapper.listRoomByKey
func (b *build) PageRoomByKey(ctx context.Context, in *dto.PageRoomByKeyReq, out *[]dto.PageRoomByKeyVO) error {
	q := db.Table(tblroom.TableName).Where(tblroom.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblroom.DelFlag.Eq(constant.YesNoNo))
	if in.BuildID != nil {
		floors, _, e := dao.Floor(db).List(ctx, ace.Where(tblfloor.BuildingId.Eq(types.BigInt(*in.BuildID)), tblfloor.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblfloor.DelFlag.Eq(constant.YesNoNo)))
		if e != nil {
			return e
		}
		if len(floors) == 0 {
			return nil
		}
		floorIds := make([]any, 0, len(floors))
		for _, f := range floors {
			floorIds = append(floorIds, f.Id)
		}
		q.And(tblroom.FloorId.In(floorIds...))
	}
	if in.FloorID != nil {
		q.And(tblroom.FloorId.Eq(types.BigInt(*in.FloorID)))
	}
	// 注：Go DO.Room 无 room_flag 列，Java 的 room_flag 过滤暂不支持
	if in.Key != nil {
		q.And(tblroom.Name.Like(*in.Key))
	}
	e := q.Desc(tblroom.Id).
		Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblroom.Id,
			tblroom.TypeId,
			tblroom.FloorId,
			tblroom.Name,
			tblroom.BedNum,
		).
		Select().
		Gets(ctx, out)
	return e
}

// GetRoomByFloorId 根据楼层编号获取房间
// 对应 Java: BuildServiceImpl.getRoomByFloorId -> roomMapper.selectByFloorId
func (b *build) GetRoomByFloorId(ctx context.Context, in *dto.GetRoomByFloorIdReq, out *[]dto.RoomByFloorIdVO) error {
	q := ace.Where(tblroom.FloorId.Eq(types.BigInt(*in.FloorID)), tblroom.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblroom.DelFlag.Eq(constant.YesNoNo))
	if in.Name != nil {
		q = q.And(tblroom.Name.Like(*in.Name))
	}
	list, _, e := dao.Room(db).List(ctx, q)
	if e != nil {
		return e
	}
	*out = make([]dto.RoomByFloorIdVO, 0, len(list))
	for _, room := range list {
		*out = append(*out, dto.RoomByFloorIdVO{
			ID:     types.BigInt(room.Id),
			Name:   room.Name.String(),
			BedNum: int(room.BedNum),
		})
	}
	return nil
}

// GetRoomById 根据房间编号获取房间详情
// 对应 Java: BuildServiceImpl.getRoomById -> roomMapper.selectById
func (b *build) GetRoomById(ctx context.Context, in *dto.IDReq, out *dto.OperateRoomVO) error {
	obj, has, e := dao.Room(db).GetByID(ctx, types.BigInt(*in.ID))
	if e != nil {
		return e
	}
	if !has || obj == nil {
		return errors.New("房间不存在")
	}
	out.ID = types.BigInt(obj.Id)
	out.TypeId = types.BigInt(obj.TypeId)
	out.FloorId = types.BigInt(obj.FloorId)
	out.Name = obj.Name.String()
	out.BedNum = int(obj.BedNum)
	return nil
}

// AddRoom 新增房间
// 对应 Java: BuildServiceImpl.addRoom -> 防重名 + insertSelective
func (b *build) AddRoom(ctx context.Context, in *dto.AddRoomReq, out *dto.EmptyResp) error {
	if in.Name != nil {
		if exist, _, e := b.getRoomByName(ctx, *in.FloorID, *in.Name); e != nil {
			return e
		} else if exist != nil {
			return errors.New("房间名称已存在")
		}
	}
	bean := do.NewRoom()
	bean.TenantId = types.BigInt(lib.TenantID(ctx))
	bean.TypeId = types.BigInt(*in.TypeID)
	bean.FloorId = types.BigInt(*in.FloorID)
	if in.Name != nil {
		bean.Name = types.String(*in.Name)
	}
	if in.BedNum != nil {
		bean.BedNum = types.Int32(int32(*in.BedNum))
	}
	bean.DelFlag = types.Int8(constant.YesNoNo)
	_, e := dao.Room(db).InsertOne(ctx, bean)
	return e
}

// EditRoom 编辑房间
// 对应 Java: BuildServiceImpl.editRoom -> 防重名 + updateByPrimaryKeySelective
func (b *build) EditRoom(ctx context.Context, in *dto.EditRoomReq, out *dto.EmptyResp) error {
	if in.Name != nil {
		if exist, _, e := b.getRoomByName(ctx, *in.FloorID, *in.Name); e != nil {
			return e
		} else if exist != nil && int64(exist.Id) != *in.ID {
			return errors.New("房间名称已存在")
		}
	}
	sets := []dialect.Setter{
		tblroom.TypeId.Set(*in.TypeID),
		tblroom.FloorId.Set(*in.FloorID),
	}
	if in.Name != nil {
		sets = append(sets, tblroom.Name.Set(types.String(*in.Name)))
	}
	if in.BedNum != nil {
		sets = append(sets, tblroom.BedNum.Set(types.Int32(int32(*in.BedNum))))
	}
	_, e := dao.Room(db).UpdateById(ctx, types.BigInt(*in.ID), sets...)
	return e
}

// DeleteRoom 删除房间（逻辑删除）
// 对应 Java: BuildServiceImpl.deleteNode -> roomFunc.deleteRoomNode
func (b *build) DeleteRoom(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	var e error
	var occupied bool
	occupied, e = b.hasOccupiedBed(ctx, markRoom, *in.ID)
	if e != nil {
		return e
	}
	if occupied {
		return errors.New("该房间下存在占用床位，无法删除")
	}
	_, e = dao.Room(db).UpdateById(ctx, types.BigInt(*in.ID), tblroom.DelFlag.Set(constant.YesNoYes))
	return e
}

// ==================== 床位 ====================

// PageBedByKey 分页查询床位（按关键词/楼栋/楼层/房间/状态）
// 对应 Java: BuildServiceImpl.pageBedByKey -> bedFunc.filterBedByKey + PageUtil 内存分页
func (b *build) PageBedByKey(ctx context.Context, in *dto.PageBedByKeyReq, out *[]dto.PageBedByKeyVO) error {
	beds, _, e := dao.Bed(db).List(ctx, ace.Where(tblbed.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblbed.DelFlag.Eq(constant.YesNoNo)))
	if e != nil {
		return e
	}
	rooms, _, e := dao.Room(db).List(ctx, ace.Where(tblroom.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblroom.DelFlag.Eq(constant.YesNoNo)))
	if e != nil {
		return e
	}
	roomMap := make(map[int64]*do.Room, len(rooms))
	for _, r := range rooms {
		roomMap[int64(r.Id)] = r
	}
	type wrapped struct {
		bed *do.Bed
	}
	filtered := make([]wrapped, 0, len(beds))
	for _, bed := range beds {
		if in.BedFlag != nil && bed.BedFlag.String() != *in.BedFlag {
			continue
		}
		room, ok := roomMap[int64(bed.RoomId)]
		if !ok {
			continue
		}
		if in.RoomID != nil && int64(bed.RoomId) != *in.RoomID {
			continue
		}
		if in.FloorID != nil && int64(room.FloorId) != *in.FloorID {
			continue
		}
		if in.BuildID != nil {
			floor, _, e := dao.Floor(db).GetByID(ctx, room.FloorId)
			if e != nil {
				return e
			}
			if floor == nil || int64(floor.BuildingId) != *in.BuildID {
				continue
			}
		}
		filtered = append(filtered, wrapped{bed: bed})
	}
	// 内存分页（参考 Java PageUtil）
	total := len(filtered)
	pageNum := int(*in.PageNum)
	pageSize := int(*in.PageSize)
	start := (pageNum - 1) * pageSize
	if start < 0 {
		start = 0
	}
	end := start + pageSize
	if start > total {
		start = total
	}
	if end > total {
		end = total
	}
	pageItems := filtered[start:end]
	*out = make([]dto.PageBedByKeyVO, 0, len(pageItems))
	for _, item := range pageItems {
		*out = append(*out, dto.PageBedByKeyVO{
			ID:      types.BigInt(item.bed.Id),
			Name:    item.bed.Name.String(),
			BedFlag: item.bed.BedFlag.String(),
		})
	}
	return nil
}

// AddBed 新增床位
// 对应 Java: BuildServiceImpl.addBed -> 防重名 + insertSelective（bed_flag=IDLE, del_flag=N）
func (b *build) AddBed(ctx context.Context, in *dto.OperateBedReq, out *dto.EmptyResp) error {
	if in.Name != nil {
		if exist, _, e := b.getBedByName(ctx, *in.RoomID, *in.Name); e != nil {
			return e
		} else if exist != nil {
			return errors.New("床位名称已存在")
		}
	}
	bean := do.NewBed()
	bean.TenantId = types.BigInt(lib.TenantID(ctx))
	bean.RoomId = types.BigInt(*in.RoomID)
	if in.Name != nil {
		bean.Name = types.String(*in.Name)
	}
	bean.BedFlag = types.Int8(constant.BedIdle)
	bean.DelFlag = types.Int8(constant.YesNoNo)
	_, e := dao.Bed(db).InsertOne(ctx, bean)
	return e
}

// GetBedById 根据床位编号获取床位详情
// 对应 Java: BuildServiceImpl.getBedById -> bedMapper.selectByPrimaryKey
func (b *build) GetBedById(ctx context.Context, in *dto.IDReq, out *dto.OperateBedVO) error {
	obj, has, e := dao.Bed(db).GetByID(ctx, types.BigInt(*in.ID))
	if e != nil {
		return e
	}
	if !has || obj == nil {
		return errors.New("床位不存在")
	}
	out.ID = types.BigInt(obj.Id)
	out.RoomId = types.BigInt(obj.RoomId)
	out.Name = obj.Name.String()
	return nil
}

// EditBed 编辑床位
// 对应 Java: BuildServiceImpl.editBed -> 防重名 + updateByPrimaryKeySelective
func (b *build) EditBed(ctx context.Context, in *dto.OperateBedReq, out *dto.EmptyResp) error {
	if in.Name != nil {
		if exist, _, e := b.getBedByName(ctx, *in.RoomID, *in.Name); e != nil {
			return e
		} else if exist != nil && int64(exist.Id) != *in.ID {
			return errors.New("床位名称已存在")
		}
	}
	sets := []dialect.Setter{
		tblbed.RoomId.Set(*in.RoomID),
	}
	if in.Name != nil {
		sets = append(sets, tblbed.Name.Set(types.String(*in.Name)))
	}
	_, e := dao.Bed(db).UpdateById(ctx, types.BigInt(*in.ID), sets...)
	return e
}

// DeleteBed 删除床位（逻辑删除，需校验床位空闲）
// 对应 Java: BuildServiceImpl.deleteBed -> 判断 bed_flag != IDLE 抛 BED_NOT_IDLE；否则 del_flag=YES
func (b *build) DeleteBed(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	bed, has, e := dao.Bed(db).GetByID(ctx, types.BigInt(*in.ID))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("床位不存在")
	}
	if bed.BedFlag != types.Int8(constant.BedIdle) {
		return errors.New("该床位被占用，删除失败")
	}
	_, e = dao.Bed(db).UpdateById(ctx, types.BigInt(*in.ID), tblbed.DelFlag.Set(constant.YesNoYes))
	return e
}

// GetNoBedTreeAndPageBedByKey 查询空闲床位树 + 分页床位
// 对应 Java: BuildServiceImpl.getNoBedTreeAndPageBedByKey -> generateBuildingSimpleTree
func (b *build) GetNoBedTreeAndPageBedByKey(ctx context.Context, in *dto.PageBedByKeyReq, out *[]dto.BuildingResp) error {
	buildings, _, e := dao.Building(db).List(ctx, ace.Where(tblbuilding.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblbuilding.DelFlag.Eq(constant.YesNoNo)))
	if e != nil {
		return e
	}
	floors, _, e := dao.Floor(db).List(ctx, ace.Where(tblfloor.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblfloor.DelFlag.Eq(constant.YesNoNo)))
	if e != nil {
		return e
	}
	rooms, _, e := dao.Room(db).List(ctx, ace.Where(tblroom.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblroom.DelFlag.Eq(constant.YesNoNo)))
	if e != nil {
		return e
	}
	beds, _, e := dao.Bed(db).List(ctx, ace.Where(tblbed.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblbed.DelFlag.Eq(constant.YesNoNo)))
	if e != nil {
		return e
	}
	tree := b.buildTree(buildings, floors, rooms, beds)
	*out = tree
	return nil
}

// DeleteNode 删除楼栋/楼层/房间节点（按 mark 级联删除）
// 对应 Java: BuildServiceImpl.deleteNode -> 按 mark 调用 building/floor/room func 逻辑删除
func (b *build) DeleteNode(ctx context.Context, in *dto.DeleteNodeReq, out *dto.EmptyResp) error {
	if in.Mark == nil {
		return errors.New("节点标记不能为空")
	}
	switch *in.Mark {
	case markBuilding:
		return b.DeleteBuilding(ctx, &dto.IDReq{ID: in.ID}, out)
	case markFloor:
		return b.DeleteFloor(ctx, &dto.IDReq{ID: in.ID}, out)
	case markRoom:
		return b.DeleteRoom(ctx, &dto.IDReq{ID: in.ID}, out)
	default:
		return errors.New("该节点标记不存在")
	}
}

// ListRoomType 房间类型下拉列表
// 对应 Java: BuildServiceImpl.listRoomType -> roomTypeFunc.listNotDelRoomType
func (b *build) ListRoomType(ctx context.Context, in *dto.EmptyReq, out *[]dto.DropDown) error {
	list, _, e := dao.RoomType(db).List(ctx, ace.Where(tblroomtype.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblroomtype.DelFlag.Eq(constant.YesNoNo)))
	if e != nil {
		return e
	}
	dropList := make([]dto.DropDown, 0, len(list))
	for _, v := range list {
		dropList = append(dropList, dto.DropDown{ID: types.BigInt(v.Id), Name: string(v.Name)})
	}
	*out = dropList
	return nil
}
