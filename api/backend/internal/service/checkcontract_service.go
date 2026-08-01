package service

import (
	"context"
	"errors"

	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblbed"
	"api/internal/model/define/table/tblbuilding"
	"api/internal/model/define/table/tblcateringset"
	"api/internal/model/define/table/tblcommunicationrecord"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/define/table/tblfloor"
	"api/internal/model/define/table/tblnursegrade"
	"api/internal/model/define/table/tblroom"
	"api/internal/model/define/table/tblvisitplan"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

// 入住办理服务
// 说明：入住办理本质是对 elder 入住信息的维护（关联 bed、contract、emergency_contact），
// 系统中不存在独立的 check_contract 表，所有操作以 elder 为主表。
type CheckContractFunc struct{}

var CheckContract = CheckContractFunc{}

// PageCheckContractByKey 分页查询入住办理列表
// 对应 Java CheckContractServiceImpl.pageCheckContractByKey（基于 elder 表，checkFlag 为入住/退住审核）
func (CheckContractFunc) PageCheckContractByKey(ctx context.Context, in *dto.PageCheckContractByKeyQuery, out *dto.EmptyResp) error {
	query := ace.Where(tblelder.CheckFlag.Eq(checkEnumEnter))
	if in.Name != "" {
		query = query.And(tblelder.Name.Like("%" + in.Name + "%"))
	}
	if in.Sex != "" {
		query = query.And(tblelder.Sex.Eq(in.Sex))
	}
	if in.IDNum != "" {
		query = query.And(tblelder.IdNum.Eq(in.IDNum))
	}
	// todo: 关联 bed、contract 分页查询（Java 通过 elderFunc 关联构建 VO）
	list, _, err := dao.Elder(db).List(ctx, query)
	if err != nil {
		return err
	}
	_ = list
	return nil
}

// PageSearchElderByKey 分页搜索老人
// 对应 Java CheckContractServiceImpl.pageSearchElderByKey（基于 elder 表，checkFlag 为咨询/意向/预定/退住）
func (CheckContractFunc) PageSearchElderByKey(ctx context.Context, in *dto.PageSearchElderByKeyQuery, out *dto.EmptyResp) error {
	query := ace.Where(tblelder.CheckFlag.Eq(checkEnumConsult))
	if in.Name != "" {
		query = query.And(tblelder.Name.Like("%" + in.Name + "%"))
	}
	if in.Phone != "" {
		query = query.And(tblelder.Phone.Like("%" + in.Phone + "%"))
	}
	// todo: 分页查询 elder（Java: checkFlag in CONSULT/INTENTION/RESERVE/EXIT）
	list, _, err := dao.Elder(db).List(ctx, query)
	if err != nil {
		return err
	}
	_ = list
	return nil
}

// AddCheckContract 新增入住办理
// 对应 Java CheckContractServiceImpl.addCheckContract（入参 OperateCheckContractQuery，操作 elder）
func (CheckContractFunc) AddCheckContract(ctx context.Context, in *dto.OperateCheckContractQuery, out *dto.EmptyResp) error {
	bean := &do.Elder{
		Name:           types.String(in.Name),
		IdNum:          types.String(in.IDNum),
		Sex:            types.String(in.Sex),
		Phone:          types.String(in.Phone),
		Address:        types.String(in.Address),
		NursingGradeId: types.BigInt(in.NursingGradeID),
		CateringSetId:  types.BigInt(in.CateringSetID),
		BedId:          types.BigInt(in.BedID),
		Age:            types.Int32(in.Age),
		CheckFlag:      checkEnumEnter, // 初始入住
	}
	// todo: 同步维护 bed（占用）、contract（生成）、emergency_contact（紧急联系人）
	if _, err := dao.Elder(db).InsertOne(ctx, bean); err != nil {
		return err
	}
	return nil
}

// GetCheckContractById 根据编号查询入住办理信息
// 对应 Java CheckContractServiceImpl.getCheckContractById（elderFunc.getElderById）
func (CheckContractFunc) GetCheckContractById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	elder, has, err := dao.Elder(db).GetByID(ctx, types.BigInt(in.ID))
	if err != nil {
		return err
	}
	if !has || elder == nil {
		return errors.New("长者不存在")
	}
	// todo: 组装 bed/contract/emergency_contact 关联信息返回
	return nil
}

// EditCheckContract 编辑入住办理
// 对应 Java CheckContractServiceImpl.editCheckContract（入参 OperateCheckContractQuery，操作 elder）
func (CheckContractFunc) EditCheckContract(ctx context.Context, in *dto.OperateCheckContractQuery, out *dto.EmptyResp) error {
	if in.ID == 0 {
		return errors.New("编号不能为空")
	}
	sets := []dialect.Setter{
		tblelder.Name.Set(in.Name),
		tblelder.IdNum.Set(in.IDNum),
		tblelder.Sex.Set(in.Sex),
		tblelder.Phone.Set(in.Phone),
		tblelder.Address.Set(in.Address),
		tblelder.NursingGradeId.Set(types.BigInt(in.NursingGradeID)),
		tblelder.CateringSetId.Set(types.BigInt(in.CateringSetID)),
		tblelder.BedId.Set(types.BigInt(in.BedID)),
		tblelder.Age.Set(types.Int32(in.Age)),
	}
	// todo: 同步维护 bed/contract/emergency_contact 关联信息
	if _, err := dao.Elder(db).UpdateById(ctx, types.BigInt(in.ID), sets...); err != nil {
		return err
	}
	return nil
}

// DeleteCheckContract 删除入住办理
// 对应 Java CheckContractServiceImpl.deleteCheckContract（入参 elder id，软删 elder）
func (CheckContractFunc) DeleteCheckContract(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	if in.ID == 0 {
		return errors.New("编号不能为空")
	}
	// todo: 解除 bed 占用、作废关联 contract
	if _, err := dao.Elder(db).DeleteById(ctx, types.BigInt(in.ID)); err != nil {
		return err
	}
	return nil
}

// ListNurseGrade 获取护理等级下拉列表
// 对应 Java CheckContractServiceImpl.listNurseGrade
func (CheckContractFunc) ListNurseGrade(ctx context.Context, in *dto.EmptyReq, out *dto.EmptyResp) error {
	list, _, err := dao.NurseGrade(db).List(ctx, ace.Where(tblnursegrade.DelFlag.Eq(YesNoEnum_NO)))
	if err != nil {
		return err
	}
	result := make([]dto.DropDown, 0, len(list))
	for _, v := range list {
		result = append(result, dto.DropDown{ID: int64(v.Id), Name: string(v.Name)})
	}
	// todo: 结果赋值 out
	return nil
}

// ListCateringSet 获取餐饮套餐下拉列表
// 对应 Java CheckContractServiceImpl.listCateringSet
func (CheckContractFunc) ListCateringSet(ctx context.Context, in *dto.EmptyReq, out *dto.EmptyResp) error {
	list, _, err := dao.CateringSet(db).List(ctx, ace.Where(tblcateringset.DelFlag.Eq(YesNoEnum_NO)))
	if err != nil {
		return err
	}
	result := make([]dto.DropDown, 0, len(list))
	for _, v := range list {
		result = append(result, dto.DropDown{ID: int64(v.Id), Name: string(v.Name)})
	}
	// todo: 结果赋值 out
	return nil
}

// GetBuildTree 获取楼宇树（楼栋-楼层-房间-床位）
// 对应 Java CheckContractServiceImpl.getBuildTree（commonFunc.getBuildingTreeResult）
func (CheckContractFunc) GetBuildTree(ctx context.Context, in *dto.EmptyReq, out *dto.EmptyResp) error {
	buildings, _, err := dao.Building(db).List(ctx, ace.Where(tblbuilding.DelFlag.Eq(YesNoEnum_NO)))
	if err != nil {
		return err
	}
	floors, _, err := dao.Floor(db).List(ctx, ace.Where(tblfloor.DelFlag.Eq(YesNoEnum_NO)))
	if err != nil {
		return err
	}
	rooms, _, err := dao.Room(db).List(ctx, ace.Where(tblroom.DelFlag.Eq(YesNoEnum_NO)))
	if err != nil {
		return err
	}
	beds, _, err := dao.Bed(db).List(ctx, ace.Where(tblbed.DelFlag.Eq(YesNoEnum_NO)))
	if err != nil {
		return err
	}
	// 组装楼宇树
	result := make([]dto.BuildingVO, 0, len(buildings))
	floorByBuilding := make(map[int64][]*do.Floor)
	for _, f := range floors {
		floorByBuilding[int64(f.BuildingId)] = append(floorByBuilding[int64(f.BuildingId)], f)
	}
	roomByFloor := make(map[int64][]*do.Room)
	for _, r := range rooms {
		roomByFloor[int64(r.FloorId)] = append(roomByFloor[int64(r.FloorId)], r)
	}
	bedByRoom := make(map[int64][]*do.Bed)
	for _, b := range beds {
		bedByRoom[int64(b.RoomId)] = append(bedByRoom[int64(b.RoomId)], b)
	}
	for _, b := range buildings {
		vo := dto.BuildingVO{ID: int64(b.Id), Name: string(b.Name), FloorNum: int(b.FloorNum)}
		for _, f := range floorByBuilding[int64(b.Id)] {
			fvo := dto.BuildingItemVO{ID: int64(f.Id), Name: string(f.Name), RoomNum: int(f.RoomNum)}
			for _, r := range roomByFloor[int64(f.Id)] {
				rvo := dto.FloorItemVO{ID: int64(r.Id), Name: string(r.Name), BedNum: int(r.BedNum)}
				for _, bed := range bedByRoom[int64(r.Id)] {
					rvo.BedList = append(rvo.BedList, dto.RoomItemVO{
						ID:      int64(bed.Id),
						Name:    string(bed.Name),
						BedFlag: string(bed.BedFlag),
					})
				}
				fvo.RoomList = append(fvo.RoomList, rvo)
			}
			vo.FloorList = append(vo.FloorList, fvo)
		}
		result = append(result, vo)
	}
	// todo: 结果赋值 out
	return nil
}

// GetBedById 根据编号获取床位信息
// 对应 Java CheckContractServiceImpl.getBedById（bedMapper.getBedById）
func (CheckContractFunc) GetBedById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	bed, has, err := dao.Bed(db).GetByID(ctx, types.BigInt(in.ID))
	if err != nil {
		return err
	}
	if !has || bed == nil {
		return errors.New("床位不存在")
	}
	// todo: 关联 room 获取 roomType、monthPrice，组装 GetBedByIDVO 返回
	return nil
}

// ============ 以下方法沿用既有 VisitPlan / CommunicationRecord（Java 中归属 Intention，Go 端合并保留）============

// PageVisitPlan 分页查询回访计划
// 对应 Java CheckContractServiceImpl.pageVisitPlan
func (CheckContractFunc) PageVisitPlan(ctx context.Context, in *dto.PageVisitPlanQuery, out *dto.EmptyResp) error {
	query := ace.Where(tblvisitplan.DelFlag.Eq(YesNoEnum_NO)).
		And(tblvisitplan.ElderId.Eq(types.BigInt(in.ElderID)))
	// todo: 支持 completeFlag 过滤（completeDate 为空表示未完成）
	list, _, err := dao.VisitPlan(db).List(ctx, query)
	if err != nil {
		return err
	}
	_ = list
	return nil
}

// AddVisitPlan 新增回访计划
// 对应 Java CheckContractServiceImpl.addVisitPlan
func (CheckContractFunc) AddVisitPlan(ctx context.Context, in *dto.AddVisitPlanQuery, out *dto.EmptyResp) error {
	bean := &do.VisitPlan{
		ElderId:  types.BigInt(in.ElderID),
		Title:    types.String(in.Title),
		PlanDate: parseTime(in.PlanDate),
		DelFlag:  YesNoEnum_NO,
	}
	if _, err := dao.VisitPlan(db).InsertOne(ctx, bean); err != nil {
		return err
	}
	return nil
}

// ExecuteVisitPlan 执行回访计划
// 对应 Java CheckContractServiceImpl.executeVisitPlan
func (CheckContractFunc) ExecuteVisitPlan(ctx context.Context, in *dto.CompleteVisitPlanQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		tblvisitplan.Content.Set(in.Content),
		tblvisitplan.CompleteDate.Set(parseTime(in.CompleteDate)),
	}
	if _, err := dao.VisitPlan(db).UpdateById(ctx, types.BigInt(in.ID), sets...); err != nil {
		return err
	}
	return nil
}

// DeleteVisitPlan 删除回访计划
// 对应 Java CheckContractServiceImpl.deleteVisitPlan
func (CheckContractFunc) DeleteVisitPlan(ctx context.Context, in *dto.DeleteVisitPlanQuery, out *dto.EmptyResp) error {
	if _, err := dao.VisitPlan(db).DeleteById(ctx, types.BigInt(in.ID)); err != nil {
		return err
	}
	return nil
}

// PageCommunicationRecord 分页查询沟通记录
// 对应 Java CheckContractServiceImpl.pageCommunicationRecord
func (CheckContractFunc) PageCommunicationRecord(ctx context.Context, in *dto.PageCommunicationRecordQuery, out *dto.EmptyResp) error {
	query := ace.Where(tblcommunicationrecord.DelFlag.Eq(YesNoEnum_NO)).
		And(tblcommunicationrecord.ElderId.Eq(types.BigInt(in.ElderID)))
	// todo: 支持关键词过滤
	list, _, err := dao.CommunicationRecord(db).List(ctx, query)
	if err != nil {
		return err
	}
	_ = list
	return nil
}

// AddCommunicationRecord 新增沟通记录
// 对应 Java addCommunicationRecord
func (CheckContractFunc) AddCommunicationRecord(ctx context.Context, in *dto.AddCommunicationRecordQuery, out *dto.EmptyResp) error {
	bean := &do.CommunicationRecord{
		ElderId:             types.BigInt(in.ElderID),
		CommunicationRecord: types.String(in.CommunicationRecord),
		RecordDate:          parseTime(in.RecordDate),
		DelFlag:             YesNoEnum_NO,
	}
	if _, err := dao.CommunicationRecord(db).InsertOne(ctx, bean); err != nil {
		return err
	}
	return nil
}

// EditCommunicationRecord 编辑沟通记录
// 对应 Java editCommunicationRecord
func (CheckContractFunc) EditCommunicationRecord(ctx context.Context, in *dto.EditCommunicationRecordQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		tblcommunicationrecord.CommunicationRecord.Set(in.CommunicationRecord),
		tblcommunicationrecord.RecordDate.Set(parseTime(in.RecordDate)),
	}
	if _, err := dao.CommunicationRecord(db).UpdateById(ctx, types.BigInt(in.ID), sets...); err != nil {
		return err
	}
	return nil
}

// DeleteCommunicationRecord 删除沟通记录
// 对应 Java deleteCommunicationRecord
func (CheckContractFunc) DeleteCommunicationRecord(ctx context.Context, in *dto.DeleteCommunicationRecordQuery, out *dto.EmptyResp) error {
	if _, err := dao.CommunicationRecord(db).DeleteById(ctx, types.BigInt(in.ID)); err != nil {
		return err
	}
	return nil
}

// 占位常量（Java 枚举在 Go 无对应定义，暂以字符串字面量占位）
const (
	YesNoEnum_NO  = "N"
	YesNoEnum_YES = "Y"
	// Java CheckEnum 状态（中文），Go 端占位
	checkEnumEnter    = "入住"
	checkEnumExitAudit = "退住审核"
	checkEnumConsult  = "咨询"
)
