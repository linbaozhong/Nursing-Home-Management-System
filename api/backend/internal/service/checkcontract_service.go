package service

import (
	"context"
	"errors"
	"time"

	"api/internal/constant"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblbed"
	"api/internal/model/define/table/tblbuilding"
	"api/internal/model/define/table/tblcateringset"
	"api/internal/model/define/table/tblcommunicationrecord"
	"api/internal/model/define/table/tblcontract"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/define/table/tblemergencycontact"
	"api/internal/model/define/table/tblfloor"
	"api/internal/model/define/table/tblnursegrade"
	"api/internal/model/define/table/tblroom"
	"api/internal/model/define/table/tblroomtype"
	"api/internal/model/define/table/tblvisitplan"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type checkcontract struct {
}

var CheckContract = &checkcontract{}

// PageCheckContractByKey 分页查询入住签约老人（入住中 + 退住审核）
// 对应 Java: CheckContractFunc.listCheckContract
func (c *checkcontract) PageCheckContractByKey(ctx context.Context, in *dto.PageCheckContractByKeyQuery, out *[]dto.PageCheckContractByKeyVO) error {
	q := db.Table(do.ElderTableName).
		Where(
			tblelder.DelFlag.Eq(constant.YesNoNo),
			tblelder.CheckFlag.In(
				types.Int8(constant.CheckEnter),
				types.Int8(constant.CheckExitAudit),
			),
		)
	if in.Name != nil && *in.Name != "" {
		q.And(tblelder.Name.Like(*in.Name))
	}
	if in.Phone != nil && *in.Phone != "" {
		q.And(tblelder.Phone.Like(*in.Phone))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblelder.Id,
			tblelder.Name,
			tblelder.IdNum,
			tblelder.Age,
			tblelder.Sex,
			tblelder.Phone,
			tblelder.Address,
			tblelder.CheckFlag,
		).
		Desc(tblelder.CreateTime).
		Select().
		Gets(ctx, out)
}

// PageSearchElderByKey 分页查询搜索老人（咨询/意向/预定/退住）
// 对应 Java: CheckContractFunc.listSearchElder
func (c *checkcontract) PageSearchElderByKey(ctx context.Context, in *dto.PageSearchElderByKeyQuery, out *[]dto.PageSearchElderByKeyVO) error {
	q := db.Table(do.ElderTableName).
		Where(
			tblelder.DelFlag.Eq(constant.YesNoNo),
			tblelder.CheckFlag.In(
				types.Int8(constant.CheckConsult),
				types.Int8(constant.CheckIntention),
				types.Int8(constant.CheckReserve),
				types.Int8(constant.CheckExit),
			),
		)
	if in.Name != nil && *in.Name != "" {
		q.And(tblelder.Name.Like(*in.Name))
	}
	if in.Phone != nil && *in.Phone != "" {
		q.And(tblelder.Phone.Like(*in.Phone))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblelder.Id,
			tblelder.Name,
			tblelder.IdNum,
			tblelder.Sex,
			tblelder.Phone,
			tblelder.Address,
			tblelder.CheckFlag,
		).
		Desc(tblelder.CreateTime).
		Select().
		Gets(ctx, out)
}

// ListNurseGrade 护理等级下拉
func (c *checkcontract) ListNurseGrade(ctx context.Context, in *dto.ListNurseGradeQuery, out *[]dto.DropDown) error {
	return db.Table(do.NurseGradeTableName).
		Cols(tblnursegrade.Id, tblnursegrade.Name).
		Where(tblnursegrade.DelFlag.Eq(constant.YesNoNo)).
		Select().
		Gets(ctx, out)
}

// ListCateringSet 餐饮套餐下拉
func (c *checkcontract) ListCateringSet(ctx context.Context, in *dto.ListCateringSetQuery, out *[]dto.DropDown) error {
	return db.Table(do.CateringSetTableName).
		Cols(tblcateringset.Id, tblcateringset.Name).
		Where(tblcateringset.DelFlag.Eq(constant.YesNoNo)).
		Select().
		Gets(ctx, out)
}

// GetBuildTree 楼宇三层下拉树（楼栋 -> 楼层 -> 房间）
func (c *checkcontract) GetBuildTree(ctx context.Context, in *dto.GetBuildTreeQuery, out *[]dto.BuildingVO) error {
	// 楼栋
	var buildings []do.Building
	e := db.Table(do.BuildingTableName).
		Cols(tblbuilding.Id, tblbuilding.Name).
		Where(tblbuilding.DelFlag.Eq(constant.YesNoNo)).
		Select().
		Gets(ctx, &buildings)
	if e != nil {
		return e
	}
	if len(buildings) == 0 {
		*out = make([]dto.BuildingVO, 0)
		return nil
	}
	// 楼层
	buildingIds := make([]any, 0, len(buildings))
	for _, b := range buildings {
		buildingIds = append(buildingIds, b.Id)
	}
	var floors []do.Floor
	e = db.Table(do.FloorTableName).
		Cols(tblfloor.Id, tblfloor.Name, tblfloor.BuildingId).
		Where(
			tblfloor.DelFlag.Eq(constant.YesNoNo),
			tblfloor.BuildingId.In(buildingIds...),
		).
		Select().
		Gets(ctx, &floors)
	if e != nil {
		return e
	}
	floorIds := make([]any, 0, len(floors))
	for _, f := range floors {
		floorIds = append(floorIds, f.Id)
	}
	// 房间
	var rooms []do.Room
	if len(floorIds) > 0 {
		e = db.Table(do.RoomTableName).
			Cols(tblroom.Id, tblroom.Name, tblroom.FloorId).
			Where(
				tblroom.DelFlag.Eq(constant.YesNoNo),
				tblroom.FloorId.In(floorIds...),
			).
			Select().
			Gets(ctx, &rooms)
		if e != nil {
			return e
		}
	}
	// 组装
	roomMap := make(map[int64][]dto.RoomItemVO, len(floors))
	for _, r := range rooms {
		roomMap[int64(r.FloorId)] = append(roomMap[int64(r.FloorId)], dto.RoomItemVO{
			ID:   int64(r.Id),
			Name: r.Name.String(),
		})
	}
	floorMap := make(map[int64][]dto.FloorItemVO, len(buildings))
	for _, f := range floors {
		floorMap[int64(f.BuildingId)] = append(floorMap[int64(f.BuildingId)], dto.FloorItemVO{
			ID:    int64(f.Id),
			Name:  f.Name.String(),
			Rooms: roomMap[int64(f.Id)],
		})
	}
	tree := make([]dto.BuildingVO, 0, len(buildings))
	for _, b := range buildings {
		tree = append(tree, dto.BuildingVO{
			ID:     int64(b.Id),
			Name:   b.Name.String(),
			Floors: floorMap[int64(b.Id)],
		})
	}
	*out = tree
	return nil
}

// GetBedById 根据编号获取床位（含房间类型信息）
// 对应 Java: BedFunc.getBedById
func (c *checkcontract) GetBedById(ctx context.Context, in *dto.IDReq, out *dto.GetBedByIDVO) error {
	var bed do.Bed
	has, e := db.Table(do.BedTableName).
		LeftJoin(tblbed.RoomId, tblroom.Id).
		LeftJoin(tblroom.TypeId, tblroomtype.Id).
		Where(tblbed.Id.Eq(types.BigInt(*in.ID))).
		Cols(
			tblbed.Id,
			tblbed.Name,
			tblbed.BedFlag,
			tblbed.Remark,
			tblroom.Name.As("room_name"),
			tblroomtype.Name.As("room_type_name"),
			tblroomtype.MonthPrice.As("month_price"),
		).
		Select().
		Get(ctx, &bed)
	if e != nil {
		return e
	}
	if !has {
		return errors.New("床位不存在")
	}
	out.ID = int64(bed.Id)
	out.Name = bed.Name.String()
	out.BedFlag = bed.BedFlag.String()
	out.Remark = bed.Remark.String()
	return nil
}

// checkNamePhoneExist 校验客户姓名+电话是否已存在（除指定老人外）
func (c *checkcontract) checkNamePhoneExist(ctx context.Context, name, phone string, excludeID *int64) (bool, error) {
	cond := []any{
		tblelder.Name.Eq(name),
		tblelder.Phone.Eq(phone),
	}
	if excludeID != nil {
		cond = append(cond, tblelder.Id.Neq(types.BigInt(*excludeID)))
	}
	return dao.Elder(db).Exists(ctx, cond...)
}

// AddCheckContract 新增入住签约
// 对应 Java: CheckContractServiceImpl.addCheckContract
func (c *checkcontract) AddCheckContract(ctx context.Context, in *dto.OperateCheckContractQuery, out *dto.EmptyResp) error {
	// 校验床位存在且为空闲
	bed, has, e := dao.Bed(db).GetByID(ctx, types.BigInt(*in.BedID), tblbed.Id, tblbed.BedFlag, tblbed.Name)
	if e != nil {
		return e
	}
	if !has {
		return errors.New("床位不存在")
	}
	if bed.BedFlag != types.String(constant.BedIdle.String()) {
		return errors.New("床位已被占用")
	}
	// 校验客户姓名+电话是否已存在
	exist, e := c.checkNamePhoneExist(ctx, *in.Name, *in.Phone, nil)
	if e != nil {
		return e
	}
	if exist {
		return errors.New("该客户已存在，请检查姓名和电话")
	}
	// 新增合同
	contract := do.NewContract()
	contract.ElderId = types.BigInt(*in.ID)
	contract.BedId = types.BigInt(*in.BedID)
	contract.StaffId = types.BigInt(*in.StaffID)
	contract.NurseGradeId = types.BigInt(*in.NurseGradeID)
	contract.CateringSetId = types.BigInt(*in.CateringSetID)
	contract.SignDate = types.Time(timeParse(*in.SignDate))
	contract.StartDate = types.Time(timeParse(*in.StartDate))
	contract.EndDate = types.Time(timeParse(*in.EndDate))
	contract.Remark = types.String(*in.Remark)
	_, e = dao.Contract(db).InsertOne(ctx, contract)
	if e != nil {
		return e
	}
	// 修改床位状态为占用
	_, e = dao.Bed(db).UpdateById(ctx, types.BigInt(*in.BedID),
		tblbed.BedFlag.Set(constant.BedEnter.String()),
	)
	if e != nil {
		return e
	}
	// 批量添加紧急联系人
	if len(in.OperateEmergencyContactQueryList) > 0 {
		list := make([]*do.EmergencyContact, 0, len(in.OperateEmergencyContactQueryList))
		for _, ec := range in.OperateEmergencyContactQueryList {
			bean := do.NewEmergencyContact()
			bean.ElderId = types.BigInt(*in.ID)
			bean.Name = types.String(*ec.Name)
			bean.Phone = types.String(*ec.Phone)
			bean.Relation = types.String(*ec.Relation)
			bean.Remark = types.String(*ec.Remark)
			list = append(list, bean)
		}
		_, e = dao.EmergencyContact(db).InsertBatch(ctx, list...)
		if e != nil {
			return e
		}
	}
	// 修改客户状态为入住
	_, e = dao.Elder(db).UpdateById(ctx, types.BigInt(*in.ID),
		tblelder.CheckFlag.Set(constant.CheckEnter),
	)
	return e
}

// GetCheckContractById 根据编号获取入住签约（含合同与紧急联系人）
// 对应 Java: CheckContractServiceImpl.getCheckContractById
func (c *checkcontract) GetCheckContractById(ctx context.Context, in *dto.IDReq, out *dto.GetCheckContractByIDVO) error {
	// 客户
	elder, has, e := dao.Elder(db).GetByID(ctx, types.BigInt(*in.ID),
		tblelder.Id, tblelder.Name, tblelder.IdNum, tblelder.Sex, tblelder.Age,
		tblelder.Phone, tblelder.Address, tblelder.CheckFlag,
	)
	if e != nil {
		return e
	}
	if !has {
		return errors.New("客户不存在")
	}
	out.ID = int64(elder.Id)
	out.Name = elder.Name.String()
	out.IDNum = elder.IdNum.String()
	out.Sex = elder.Sex.String()
	out.Age = int(elder.Age)
	out.Phone = elder.Phone.String()
	out.Address = elder.Address.String()
	out.CheckFlag = string(elder.CheckFlag)

	// 合同
	contract, has, e := dao.Contract(db).Get(ctx, tblcontract.ElderId.Eq(types.BigInt(*in.ID)))
	if e != nil {
		return e
	}
	if has {
		out.StaffID = int64Ptr(int64(contract.StaffId))
		out.SignDate = strPtr(timeFormat(contract.SignDate))
		out.StartDate = strPtr(timeFormat(contract.StartDate))
		out.EndDate = strPtr(timeFormat(contract.EndDate))
		out.NurseGradeID = int64Ptr(int64(contract.NurseGradeId))
		out.CateringSetID = int64Ptr(int64(contract.CateringSetId))
		out.BedID = int64Ptr(int64(contract.BedId))
		out.Remark = strPtr(contract.Remark.String())
	}

	// 紧急联系人
	var contacts []do.EmergencyContact
	e = db.Table(do.EmergencyContactTableName).
		Cols(
			tblemergencycontact.Id,
			tblemergencycontact.Name,
			tblemergencycontact.Phone,
			tblemergencycontact.Relation,
			tblemergencycontact.Remark,
		).
		Where(tblemergencycontact.ElderId.Eq(types.BigInt(*in.ID))).
		Select().
		Gets(ctx, &contacts)
	if e != nil {
		return e
	}
	out.OperateEmergencyContactQueryList = make([]dto.OperateEmergencyContactQuery, 0, len(contacts))
	for _, ct := range contacts {
		out.OperateEmergencyContactQueryList = append(out.OperateEmergencyContactQueryList, dto.OperateEmergencyContactQuery{
			ID:       int64Ptr(int64(ct.Id)),
			Name:     strPtr(ct.Name.String()),
			Phone:    strPtr(ct.Phone.String()),
			Relation: strPtr(ct.Relation.String()),
			Remark:   strPtr(ct.Remark.String()),
		})
	}
	return nil
}

// EditCheckContract 编辑入住签约
// 对应 Java: CheckContractServiceImpl.editCheckContract
func (c *checkcontract) EditCheckContract(ctx context.Context, in *dto.OperateCheckContractQuery, out *dto.EmptyResp) error {
	// 校验客户姓名+电话是否已存在（排除自身）
	exist, e := c.checkNamePhoneExist(ctx, *in.Name, *in.Phone, in.ID)
	if e != nil {
		return e
	}
	if exist {
		return errors.New("该客户已存在，请检查姓名和电话")
	}
	// 修改老人信息
	var elderSets = make([]dialect.Setter, 0, 6)
	elderSets = append(elderSets, tblelder.Name.Set(*in.Name))
	elderSets = append(elderSets, tblelder.IdNum.Set(*in.IDNum))
	elderSets = append(elderSets, tblelder.Sex.Set(*in.Sex))
	elderSets = append(elderSets, tblelder.Age.Set(*in.Age))
	elderSets = append(elderSets, tblelder.Phone.Set(*in.Phone))
	elderSets = append(elderSets, tblelder.Address.Set(*in.Address))
	_, e = dao.Elder(db).UpdateById(ctx, types.BigInt(*in.ID), elderSets...)
	if e != nil {
		return e
	}
	// 修改合同
	var contractSets = make([]dialect.Setter, 0, 8)
	contractSets = append(contractSets, tblcontract.BedId.Set(*in.BedID))
	contractSets = append(contractSets, tblcontract.StaffId.Set(*in.StaffID))
	contractSets = append(contractSets, tblcontract.NurseGradeId.Set(*in.NurseGradeID))
	contractSets = append(contractSets, tblcontract.CateringSetId.Set(*in.CateringSetID))
	contractSets = append(contractSets, tblcontract.SignDate.Set(timeParse(*in.SignDate)))
	contractSets = append(contractSets, tblcontract.StartDate.Set(timeParse(*in.StartDate)))
	contractSets = append(contractSets, tblcontract.EndDate.Set(timeParse(*in.EndDate)))
	contractSets = append(contractSets, tblcontract.Remark.Set(*in.Remark))
	_, e = dao.Contract(db).Update(ctx,
		contractSets...,
		tblcontract.ElderId.Eq(types.BigInt(*in.ID)),
	)
	if e != nil {
		return e
	}
	// 删除原有紧急联系人并重新添加
	_, e = dao.EmergencyContact(db).Delete(ctx, tblemergencycontact.ElderId.Eq(types.BigInt(*in.ID)))
	if e != nil {
		return e
	}
	if len(in.OperateEmergencyContactQueryList) > 0 {
		list := make([]*do.EmergencyContact, 0, len(in.OperateEmergencyContactQueryList))
		for _, ec := range in.OperateEmergencyContactQueryList {
			bean := do.NewEmergencyContact()
			bean.ElderId = types.BigInt(*in.ID)
			bean.Name = types.String(*ec.Name)
			bean.Phone = types.String(*ec.Phone)
			bean.Relation = types.String(*ec.Relation)
			bean.Remark = types.String(*ec.Remark)
			list = append(list, bean)
		}
		_, e = dao.EmergencyContact(db).InsertBatch(ctx, list...)
		if e != nil {
			return e
		}
	}
	return nil
}

// DeleteCheckContract 删除入住签约（退住审核流程）
// 对应 Java: CheckContractServiceImpl.deleteCheckContract -> ElderFunc.checkEnter
func (c *checkcontract) DeleteCheckContract(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	// 仅入住状态可操作
	elder, has, e := dao.Elder(db).GetByID(ctx, types.BigInt(*in.ID), tblelder.Id, tblelder.CheckFlag)
	if e != nil {
		return e
	}
	if !has || elder.CheckFlag != types.Int8(constant.CheckEnter) {
		return errors.New("客户不在入住状态，无法删除")
	}
	// 修改客户状态为退住审核
	_, e = dao.Elder(db).UpdateById(ctx, types.BigInt(*in.ID),
		tblelder.CheckFlag.Set(constant.CheckExitAudit),
	)
	if e != nil {
		return e
	}
	// 释放床位
	contract, has, e := dao.Contract(db).Get(ctx, tblcontract.ElderId.Eq(types.BigInt(*in.ID)), tblcontract.BedId)
	if e != nil {
		return e
	}
	if has {
		_, e = dao.Bed(db).UpdateById(ctx, contract.BedId,
			tblbed.BedFlag.Set(constant.BedIdle.String()),
		)
		if e != nil {
			return e
		}
		// 删除合同
		_, e = dao.Contract(db).Delete(ctx, tblcontract.ElderId.Eq(types.BigInt(*in.ID)))
		if e != nil {
			return e
		}
	}
	return nil
}

// PageVisitPlanByKey 分页查询回访计划
func (c *checkcontract) PageVisitPlanByKey(ctx context.Context, in *dto.PageVisitPlanQuery, out *[]dto.PageVisitPlanVO) error {
	q := db.Table(do.VisitPlanTableName).
		Where(tblvisitplan.DelFlag.Eq(constant.YesNoNo))
	if in.ElderID != nil {
		q.And(tblvisitplan.ElderId.Eq(types.BigInt(*in.ElderID)))
	}
	if in.CompleteFlag != nil {
		if *in.CompleteFlag {
			q.And(tblvisitplan.CompleteDate.IsNotNull())
		} else {
			q.And(tblvisitplan.CompleteDate.IsNull())
		}
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblvisitplan.Id,
			tblvisitplan.ElderId,
			tblvisitplan.Title,
			tblvisitplan.PlanDate,
			tblvisitplan.CompleteDate,
			tblvisitplan.Content,
		).
		Desc(tblvisitplan.CreateTime).
		Select().
		Gets(ctx, out)
}

// AddVisitPlan 新增回访计划
func (c *checkcontract) AddVisitPlan(ctx context.Context, in *dto.AddVisitPlanQuery, out *dto.EmptyResp) error {
	bean := do.NewVisitPlan()
	bean.ElderId = types.BigInt(*in.ElderID)
	bean.Title = types.String(*in.Title)
	bean.PlanDate = types.Time(timeParse(*in.PlanDate))
	_, e := dao.VisitPlan(db).InsertOne(ctx, bean)
	return e
}

// CompleteVisitPlan 完成回访计划
func (c *checkcontract) CompleteVisitPlan(ctx context.Context, in *dto.CompleteVisitPlanQuery, out *dto.EmptyResp) error {
	_, e := dao.VisitPlan(db).UpdateById(ctx, types.BigInt(*in.ID),
		tblvisitplan.Content.Set(*in.Content),
		tblvisitplan.CompleteDate.Set(timeParse(*in.CompleteDate)),
	)
	return e
}

// DeleteVisitPlan 删除回访计划
func (c *checkcontract) DeleteVisitPlan(ctx context.Context, in *dto.DeleteVisitPlanQuery, out *dto.EmptyResp) error {
	_, e := dao.VisitPlan(db).DeleteById(ctx, types.BigInt(*in.ID))
	return e
}

// PageCommunicationRecordByKey 分页查询沟通记录
func (c *checkcontract) PageCommunicationRecordByKey(ctx context.Context, in *dto.PageCommunicationRecordQuery, out *[]dto.PageCommunicationRecordVO) error {
	q := db.Table(do.CommunicationRecordTableName).
		Where(tblcommunicationrecord.DelFlag.Eq(constant.YesNoNo))
	if in.ElderID != nil {
		q.And(tblcommunicationrecord.ElderId.Eq(types.BigInt(*in.ElderID)))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblcommunicationrecord.Id,
			tblcommunicationrecord.RecordDate,
			tblcommunicationrecord.CommunicationRecord,
		).
		Desc(tblcommunicationrecord.RecordDate).
		Select().
		Gets(ctx, out)
}

// AddCommunicationRecord 新增沟通记录
func (c *checkcontract) AddCommunicationRecord(ctx context.Context, in *dto.AddCommunicationRecordQuery, out *dto.EmptyResp) error {
	bean := do.NewCommunicationRecord()
	bean.ElderId = types.BigInt(*in.ElderID)
	bean.RecordDate = types.Time(timeParse(*in.RecordDate))
	bean.CommunicationRecord = types.String(*in.CommunicationRecord)
	_, e := dao.CommunicationRecord(db).InsertOne(ctx, bean)
	return e
}

// EditCommunicationRecord 编辑沟通记录
func (c *checkcontract) EditCommunicationRecord(ctx context.Context, in *dto.EditCommunicationRecordQuery, out *dto.EmptyResp) error {
	var sets = make([]dialect.Setter, 0, 2)
	sets = append(sets, tblcommunicationrecord.RecordDate.Set(timeParse(*in.RecordDate)))
	sets = append(sets, tblcommunicationrecord.CommunicationRecord.Set(*in.CommunicationRecord))
	_, e := dao.CommunicationRecord(db).UpdateById(ctx, types.BigInt(*in.ID), sets...)
	return e
}

// DeleteCommunicationRecord 删除沟通记录
func (c *checkcontract) DeleteCommunicationRecord(ctx context.Context, in *dto.DeleteCommunicationRecordQuery, out *dto.EmptyResp) error {
	_, e := dao.CommunicationRecord(db).DeleteById(ctx, types.BigInt(*in.ID))
	return e
}

// ---- 辅助函数 ----

func int64Ptr(v int64) *int64 { return &v }
func strPtr(v string) *string { return &v }

func timeParse(s *string) types.Time {
	if s == nil || *s == "" {
		return types.Time{}
	}
	for _, layout := range []string{"2006-01-02 15:04:05", "2006-01-02"} {
		if t, err := parseTime(layout, *s); err == nil {
			return types.Time(t)
		}
	}
	return types.Time{}
}

func timeFormat(t types.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.String()
}

func parseTime(layout, value string) (time.Time, error) {
	return time.ParseInLocation(layout, value, time.Local)
}
