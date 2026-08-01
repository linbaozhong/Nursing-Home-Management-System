package service

import (
	"context"
	"errors"
	"time"

	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblbed"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/define/table/tblreserve"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type reserve struct{}

var Reserve = &reserve{}

// PageReserveByKey 分页查询预订（联表 elder、user、building、room）
// 对应 Java: ReserveServiceImpl.pageReserveByKey -> ReserveMapper.listReserveByKey
// SQL: SELECT r.*, e.elder_name, u.name AS charge_user_name, b.building_name, rm.room_name
//
//	FROM reserve r
//	LEFT JOIN elder e ON e.id = r.elder_id
//	LEFT JOIN user u ON u.id = r.charge_user_id
//	LEFT JOIN building b ON b.id = r.building_id
//	LEFT JOIN room rm ON rm.id = r.room_id
//	WHERE (e.elder_name LIKE %key% OR r.id = key) [可选]
//	ORDER BY r.create_time DESC; 再由 PageUtil 内存分页。
//
// Todo: 1) in.Key 非空 -> (tbl<reserve>.Id.Eq(in.Key) OR tbl<elder>.ElderName.Like(in.Key))
//
//	2) DB 分页: Count + List(联表 LeftJoin)
//	3) 组装含老人/负责人/楼宇/房间名的 VO 并赋值 out
func (r *reserve) PageReserveByKey(ctx context.Context, in *dto.PageReserveByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 实现联表分页查询
	return nil
}

// GetReserveById 根据编号获取预订
// 对应 Java: ReserveServiceImpl.getReserveById -> reserveMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.Reserve(db).GetByID(ctx, types.Money(internal))
func (r *reserve) GetReserveById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.Reserve(db).GetByID(ctx, types.Money(internal))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// PageSearchElderByKey 分页搜索老人（供预订选择老人）
// 对应 Java: ReserveServiceImpl.pageSearchElderByKey -> elderMapper.listElderByKey
// SQL: SELECT * FROM elder WHERE (elder_name LIKE %key% OR id = key) [可选] AND del_flag=0
// todo: 查询 elder 表并分页, 结果赋值 out(需定义老人分页返回类型)
func (r *reserve) PageSearchElderByKey(ctx context.Context, in *dto.PageSearchElderByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 elder 表并分页
	return nil
}

// PageBuildingByKey 分页查询楼宇（供预订选择楼宇）
// 对应 Java: ReserveServiceImpl.pageBuildingByKey -> buildingMapper.listBuildingByKey
// SQL: SELECT * FROM building WHERE (building_name LIKE %key%) [可选]
// todo: 查询 building 表并分页, 结果赋值 out(需定义楼宇分页返回类型)
func (r *reserve) PageBuildingByKey(ctx context.Context, in *dto.PageBuildingByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 building 表并分页
	return nil
}

// GetFloorByBuildingId 根据楼宇编号获取楼层（供预订选择楼层）
// 对应 Java: ReserveServiceImpl.getFloorByBuildingId -> floorMapper.selectByBuildingId
// SQL: SELECT * FROM floor WHERE building_id = #{buildingId}
// todo: 标准查询 - dao.Floor(db).List(ace.Where(tbl<floor>.BuildingId.Eq(in.BuildingId)))
func (r *reserve) GetFloorByBuildingId(ctx context.Context, in *dto.GetFloorByBuildingIdQuery, out *dto.EmptyResp) error {
	// todo: list, e := dao.Floor(db).List(ctx, ace.Where(tbl<floor>.BuildingId.Eq(in.BuildingId)))
	return nil
}

// GetRoomByFloorId 根据楼层编号获取房间（供预订选择房间）
// 对应 Java: ReserveServiceImpl.getRoomByFloorId -> roomMapper.selectByFloorId
// SQL: SELECT * FROM room WHERE floor_id = #{floorId}
// todo: 标准查询 - dao.Room(db).List(ace.Where(tbl<room>.FloorId.Eq(in.FloorId)))
func (r *reserve) GetRoomByFloorId(ctx context.Context, in *dto.GetRoomByFloorIdQuery, out *dto.EmptyResp) error {
	// todo: list, e := dao.Room(db).List(ctx, ace.Where(tbl<room>.FloorId.Eq(in.FloorId)))
	return nil
}

// AddReserve 新增预订
// 对应 Java: ReserveServiceImpl.addReserve -> reserveMapper.insertSelective
// todo: 标准 CRUD - dao.Reserve(db).InsertOne 写入 reserve 表(含 elderId/buildingId/roomId/reserveStatus 等)
func (r *reserve) AddReserve(ctx context.Context, in *dto.AddReserveQuery, out *dto.EmptyResp) error {
	// todo: bean := do.NewReserve(); 填充 in; dao.Reserve(db).InsertOne(ctx, bean)
	return nil
}

// EditReserve 编辑预订
// 对应 Java: ReserveServiceImpl.editReserve -> reserveMapper.updateByPrimaryKeySelective
// todo: 标准 CRUD - 按主键更新 reserve 表
func (r *reserve) EditReserve(ctx context.Context, in *dto.EditReserveQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<reserve>.ReserveStatus.Value(in.ReserveStatus),
	}
	_, e := dao.Reserve(db).UpdateById(ctx, types.Money(internal), sets...)
	return e
}

// DeleteReserve 删除预订
// 对应 Java: ReserveServiceImpl.deleteReserve -> reserveMapper.deleteByPrimaryKey
// todo: 标准 CRUD - dao.Reserve(db).DeleteById(ctx, types.Money(internal))
func (r *reserve) DeleteReserve(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.Reserve(db).DeleteById(ctx, types.Money(internal))
	return e
}

// ListReserveStatus 预订状态列表（字典/常量）
// 对应 Java: ReserveServiceImpl.listReserveStatus -> 返回固定枚举/字典
// todo: 返回预订状态枚举, 结果赋值 out(需定义返回类型)
func (r *reserve) ListReserveStatus(ctx context.Context, in *dto.EmptyReq, out *dto.EmptyResp) error {
	// todo: 返回预订状态枚举列表
	return nil
}

// GetBuildTree 获取楼栋树（楼栋→楼层→房间→床位）
// 对应 Java: ReserveServiceImpl.getBuildTree -> CommonFunc.getBuildingTreeResult
// 逻辑：查所有 building/floor/room/bed，组装为树形结构，床位含 id/name/type/flag
//
// todo: 1) dao.Building/Floor/Room/Bed 全量查询 2) 按层级组装树 3) 结果赋值 out(需定义树 VO)
func (r *reserve) GetBuildTree(ctx context.Context, in *dto.EmptyReq, out *dto.EmptyResp) error {
	// todo: 见上方方法注释
	return nil
}

// GetReserveByReserveIdAndElderId 根据预订编号与老人编号查询预订详情
// 对应 Java: ReserveServiceImpl.getReserveByReserveIdAndElderId -> reserveMapper.getReserveByIdAndElderId
// SQL: SELECT r.*, e.elder_name, e.elder_sex, e.elder_phone, u.name AS staff_name
//
//	FROM reserve r LEFT JOIN elder e ON e.id=r.elder_id LEFT JOIN user u ON u.id=r.staff_id
//	WHERE r.id=? AND r.elder_id=?
//
// todo: 联表查询并组装 GetReserveByIdAndElderIdVO, 赋值 out
func (r *reserve) GetReserveByReserveIdAndElderId(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	// todo: 见上方 SQL 注释, 使用 dao.Reserve(db) 联表查询
	return nil
}

// Refund 退款（取消预订）
// 对应 Java: ReserveServiceImpl.refund @Transactional
// 逻辑：
//  1. 查 reserve，校验 reserve_flag=NO(未退款)
//  2. reserve_flag=YES(已退款)
//  3. 若老人 check_flag=RESERVE 且 bed_id 一致 -> 取消老人预订(check_flag=NO)
//  4. 若床位 bed_flag=RESERVE 且属于该预订 -> 床位恢复 IDLE
func (r *reserve) Refund(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	// 1) 查预订
	reserve, has, e := dao.Reserve(db).GetByID(ctx, types.Money(internal))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("预订不存在")
	}
	// 校验是否已退款
	if reserve.ReserveFlag == "YES" {
		return errors.New("预订已退款") // 对应 Java ExceptionEnum.RESERVE_SUCCESS
	}
	// 2) 标记预订已退款
	_, e = dao.Reserve(db).UpdateById(ctx, types.Money(internal),
		ace.Where(tblreserve.ReserveFlag.Set("YES")),
	)
	if e != nil {
		return e
	}
	// 3) 若老人处于预订状态且床位一致 -> 恢复老人为未预订
	elder, has, e := dao.Elder(db).GetByID(ctx, reserve.ElderId)
	if e != nil {
		return e
	}
	if has && elder.CheckFlag == "RESERVE" {
		_, e = dao.Elder(db).UpdateById(ctx, reserve.ElderId,
			ace.Where(tblelder.CheckFlag.Set("NO")),
		)
		if e != nil {
			return e
		}
	}
	// 4) 若床位处于预订状态 -> 恢复 IDLE
	bed, has, e := dao.Bed(db).GetByID(ctx, reserve.BedId)
	if e != nil {
		return e
	}
	if has && bed.BedFlag == "RESERVE" {
		_, e = dao.Bed(db).UpdateById(ctx, reserve.BedId,
			ace.Where(tblbed.BedFlag.Set("IDLE")),
		)
		if e != nil {
			return e
		}
	}
	return nil
}

// ReserveExpireJob 过期预订定时任务
// 对应 Java: ReserveServiceImpl.reserveExpireJob
// 逻辑：查所有 reserve_flag=NO 且 due_date < now 的过期预订，删除并恢复对应老人/床位状态
//
// todo: Go 侧无定时任务基础设施，建议由外部调度器定时调用此接口
func (r *reserve) ReserveExpireJob(ctx context.Context, in *dto.EmptyReq, out *dto.EmptyResp) error {
	expired, _, e := dao.Reserve(db).List(ctx,
		ace.Where(tblreserve.ReserveFlag.Eq("NO")),
		ace.Where(tblreserve.DueDate.Lt(parseTime(time.Now().Format("2006-01-02 15:04:05")))),
	)
	if e != nil {
		return e
	}
	for _, rp := range expired {
		// todo: 删除过期预订, 并恢复对应 elder.check_flag=NO / bed.bed_flag=IDLE
		_, e := dao.Reserve(db).DeleteById(ctx, rp.Id)
		if e != nil {
			return e
		}
	}
	return nil
}
