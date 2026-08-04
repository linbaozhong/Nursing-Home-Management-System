package service

import (
	"api/internal/model/define/table/tblbed"
	"api/internal/model/define/table/tblbuilding"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/define/table/tblfloor"
	"api/internal/model/define/table/tblroom"
	"context"
	"github.com/linbaozhong/gentity/pkg/ace"

	"api/internal/constant"
	"api/internal/model/define/dao"
	"api/internal/model/do"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/types"
)

type bedpanorama struct{}

var BedPanorama = &bedpanorama{}

// ListBuilding 获取楼栋列表
// 对应 Java: BedPanoramaServiceImpl.listBuilding -> buildingFunc.listNotDelBuilding
func (b *bedpanorama) ListBuilding(ctx context.Context, in *dto.EmptyReq, out *[]dto.DropDown) error {
	return db.Table(do.BuildingTableName).
		Cols(tblbuilding.Id, tblbuilding.Name).
		Where(tblbuilding.DelFlag.Eq(constant.YesNoNo)).
		Select().
		Gets(ctx, out)
}

// ListFloorByBuildingId 获取楼层列表
// 对应 Java: BedPanoramaServiceImpl.listFloorByBuildingId -> floorFunc.listNotDelFloorByBuildingId
func (b *bedpanorama) ListFloorByBuildingId(ctx context.Context, in *dto.ListFloorByBuildingIdQuery, out *[]dto.DropDown) error {
	q := db.Table(do.FloorTableName).
		Cols(tblfloor.Id, tblfloor.Name).
		Where(tblfloor.DelFlag.Eq(constant.YesNoNo))
	if in.BuildingID != nil {
		q = q.Where(tblfloor.BuildingId.Eq(types.BigInt(*in.BuildingID)))
	}
	return q.Select().Gets(ctx, out)
}

// ListRoomByKey 获取房间列表（含床位与入住老人）
// 对应 Java: BedPanoramaServiceImpl.listRoomByKey -> roomFunc.generateRoomTree
func (b *bedpanorama) ListRoomByKey(ctx context.Context, in *dto.ListRoomByKeyQuery, out *[]dto.FloorItemVO) error {
	// 1. 获取楼层编号列表
	q := db.Cols(
		tblfloor.Id,
		tblfloor.Name,
	).
		Where(tblfloor.DelFlag.Eq(constant.YesNoNo))
	if in.BuildingID != nil {
		q.And(tblfloor.BuildingId.Eq(*in.BuildingID))
	}
	if in.FloorID != nil {
		q.And(tblfloor.Id.Eq(*in.FloorID))
	}

	floors, _, e := dao.Floor(db).List(ctx, q)
	if e != nil {
		return e
	}
	if len(floors) == 0 {
		return nil
	}
	floorIdSet := make(map[types.BigInt]*do.Floor, len(floors))
	floorIds := make([]any, 0, len(floors))
	for _, f := range floors {
		floorIdSet[f.Id] = f
		floorIds = append(floorIds, f.Id)
	}

	// 2. 根据楼层编号获取未删除的房间
	rooms, _, e := dao.Room(db).List(ctx,
		ace.Where(
			tblroom.DelFlag.Eq(constant.YesNoNo),
			tblroom.FloorId.In(floorIds...),
		))
	if e != nil {
		return e
	}
	if len(rooms) == 0 {
		return nil
	}
	roomIds := make([]any, 0, len(rooms))
	for _, r := range rooms {
		roomIds = append(roomIds, r.Id)
	}

	// 3. 获取未删除的床位
	beds, _, e := dao.Bed(db).List(ctx,
		ace.Where(
			tblbed.DelFlag.Eq(constant.YesNoNo),
			tblbed.RoomId.In(roomIds...),
		))
	if e != nil {
		return e
	}

	// 4. 获取入住老人（按床位编号关联），可按老人姓名过滤
	elderCond := ace.Cols(
		tblelder.BedId,
		tblelder.Name,
		tblelder.Age,
	).Where(tblelder.CheckFlag.Eq(constant.YesNoYes))
	if in.ElderName != nil {
		elderCond.And(tblelder.Name.Like(*in.ElderName))
	}
	elders, _, e := dao.Elder(db).List(ctx, elderCond)
	if e != nil {
		return e
	}
	elderByBed := make(map[types.BigInt]*do.Elder, len(elders))
	for _, ed := range elders {
		elderByBed[ed.BedId] = ed
	}

	// 5. 组装房间-床位树（按楼层分组）
	roomByFloor := make(map[int64][]*do.Room, len(rooms))
	for _, r := range rooms {
		roomByFloor[int64(r.FloorId)] = append(roomByFloor[int64(r.FloorId)], r)
	}
	bedByRoom := make(map[int64][]*do.Bed, len(beds))
	for _, bd := range beds {
		bedByRoom[int64(bd.RoomId)] = append(bedByRoom[int64(bd.RoomId)], bd)
	}

	*out = make([]dto.FloorItemVO, 0, len(floorIdSet))
	for _, floor := range floorIdSet {
		floorVO := dto.FloorItemVO{
			ID:      int64(floor.Id),
			Name:    floor.Name.String(),
			BedList: make([]dto.RoomItemVO, 0),
		}
		bedNum := 0
		for _, room := range roomByFloor[int64(floor.Id)] {
			roomBeds := bedByRoom[int64(room.Id)]
			bedNum += len(roomBeds)
			for _, bed := range roomBeds {
				bedVO := dto.RoomItemVO{
					ID:      int64(bed.Id),
					Name:    bed.Name.String(),
					BedFlag: bed.BedFlag.String(),
				}
				if ed, ok := elderByBed[bed.Id]; ok {
					bedVO.ElderName = ed.Name.String()
					bedVO.Sex = ed.Sex.String()
					bedVO.Age = int(ed.Age)
				}
				floorVO.BedList = append(floorVO.BedList, bedVO)
			}
		}
		floorVO.BedNum = bedNum
		*out = append(*out, floorVO)
	}

	return nil
}
