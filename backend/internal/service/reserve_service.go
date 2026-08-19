package service

import (
	"context"
	"time"

	"api/internal/constant"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblbed"
	"api/internal/model/define/table/tblbuilding"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/define/table/tblreserve"
	"api/internal/model/define/table/tblstaff"
	"api/internal/model/do"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/types"
)

var _ = (*reserveService)(nil)

type reserveService struct{}

// reserveJoin 接收预定联表（老人姓名、申请人姓名、床位名）查询结果的中间结构体
type reserveJoin struct {
	ID          types.BigInt `json:"id"`
	ElderName   types.String `json:"elder_name"`
	StaffName   types.String `json:"staff_name"`
	Deposit     types.Money  `json:"deposit"`
	DueDate     types.Time   `json:"due_date"`
	ReserveFlag types.Int8   `json:"reserve_flag"`
}

// PageReserveByKey 分页查询预定
func (s *reserveService) PageReserveByKey(ctx context.Context, in *dto.PageReserveByKeyQuery, out *[]dto.PageReserveByKeyVO) error {
	if in.PageNum == nil || in.PageSize == nil {
		return constant.ErrParamInvalid
	}
	q := db.Table(tblreserve.TableName).
		LeftJoin(tblreserve.ElderId, tblelder.Id).
		LeftJoin(tblreserve.StaffId, tblstaff.Id)
	if in.Key != nil && *in.Key != "" {
		q = q.Where(tblelder.Name.Like(*in.Key))
	}
	var joins []reserveJoin
	has, e := q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblreserve.Id,
			tblreserve.ElderId,
			tblreserve.Deposit,
			tblreserve.DueDate,
			tblreserve.ReserveFlag,
			tblelder.Name.As("elder_name"),
			tblstaff.Name.As("staff_name"),
		).
		Desc(tblreserve.Id).
		Select().Gets(ctx, &joins)
	if e != nil {
		return e
	}
	if !has {
		return nil
	}
	res := make([]dto.PageReserveByKeyVO, 0, len(joins))
	for _, j := range joins {
		res = append(res, dto.PageReserveByKeyVO{
			ID:          int64(j.ID),
			ElderName:   j.ElderName.String(),
			StaffName:   j.StaffName.String(),
			Deposit:     j.Deposit.Float64(),
			DueDate:     j.DueDate.Time(),
			ReserveFlag: constant.YesNo(j.ReserveFlag).String(),
		})
	}
	*out = res
	return nil
}

// GetReserveById 查询预定详情
func (s *reserveService) GetReserveById(ctx context.Context, in *dto.IDReq, out *dto.GetReserveByReserveIDAndElderIDVO) error {
	rec := new(do.Reserve)
	has, e := dao.Reserve(db).Get(ctx, ace.Where(tblreserve.Id.Eq(types.BigInt(*in.ID))))
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrDataNotExist
	}
	out.ReserveID = int64(rec.Id)
	out.ElderID = int64(rec.ElderId)
	out.StaffID = int64(rec.StaffId)
	out.DueDate = rec.DueDate.Time()
	out.Deposit = rec.Deposit.Float64()
	out.Remark = rec.Remark.String()
	out.ReserveFlag = constant.YesNo(rec.ReserveFlag).String()
	return nil
}

// AddReserve 新增预定（含老人与床位初始化）
func (s *reserveService) AddReserve(ctx context.Context, in *dto.AddReserveQuery, out *dto.EmptyResp) error {
	if in.StaffID == nil {
		return constant.ErrParamInvalid
	}
	// 老人不存在则创建
	elderId := orInt64(in.ElderID)
	if elderId == 0 {
		elder := &do.Elder{
			Name:      types.String(orEmpty(in.ElderName)),
			Age:       types.Int8(orInt8(in.ElderAge)),
			Sex:       types.String(orEmpty(in.ElderSex)),
			Phone:     types.String(orEmpty(in.ElderPhone)),
			Address:   types.String(orEmpty(in.ElderAddress)),
			IdNum:     types.String(orEmpty(in.IDNum)),
			CheckFlag: types.Int8(constant.CheckIntention),
		}
		if _, e := dao.Elder(db).InsertOne(ctx, elder); e != nil {
			return e
		}
		elderId = int64(elder.Id)
	}
	// 预留床位
	if in.BedID != nil {
		_, _ = dao.Bed(db).UpdateById(ctx, *in.BedID, tblbed.BedFlag.Set(types.Int8(constant.BedReserve)))
		_, _ = dao.Elder(db).UpdateById(ctx, elderId, tblelder.BedId.Set(types.BigInt(*in.BedID)))
	}
	rec := &do.Reserve{
		Name:        types.String(orEmpty(in.ElderName)),
		Phone:       types.String(orEmpty(in.ElderPhone)),
		Id:          types.String(orEmpty(in.IDNum)),
		ElderId:     types.BigInt(elderId),
		StaffId:     types.BigInt(*in.StaffID),
		DueDate:     types.Time{Time: timePtr(in.DueDate)},
		Deposit:     types.Money(orFloat64(in.Deposit)),
		Remark:      types.String(orEmpty(in.Remark)),
		CreateId:    types.BigInt(*in.StaffID),
		ReserveFlag: types.Int8(constant.YesNoNo),
	}
	if _, e := dao.Reserve(db).InsertOne(ctx, rec); e != nil {
		return e
	}
	// 老人状态置为预定
	if _, e = dao.Elder(db).UpdateById(ctx, elderId, tblelder.CheckFlag.Set(types.Int8(constant.CheckReserve))); e != nil {
		return e
	}
	return nil
}

// EditReserve 编辑预定
func (s *reserveService) EditReserve(ctx context.Context, in *dto.EditReserveQuery, out *dto.EmptyResp) error {
	if in.ReserveID == nil {
		return constant.ErrParamInvalid
	}
	upd := ace.NewUpdateBuilder()
	if in.DueDate != nil {
		upd.Set(tblreserve.DueDate.Set(types.Time{Time: timePtr(in.DueDate)}))
	}
	if in.Deposit != nil {
		upd.Set(tblreserve.Deposit.Set(types.Money(*in.Deposit)))
	}
	if in.Remark != nil {
		upd.Set(tblreserve.Remark.Set(types.String(*in.Remark)))
	}
	if _, e := dao.Reserve(db).Update(ctx, ace.Where(tblreserve.Id.Eq(types.BigInt(*in.ReserveID))).Assign(upd)); e != nil {
		return e
	}
	return nil
}

// DeleteReserve 删除预定（物理删除）
func (s *reserveService) DeleteReserve(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	if _, e := dao.Reserve(db).DeleteById(ctx, *in.ID); e != nil {
		return e
	}
	return nil
}

// PageSearchElderByKey 分页查询老人（供预定选择）
func (s *reserveService) PageSearchElderByKey(ctx context.Context, in *dto.PageSearchElderByKeyQuery, out *[]dto.PageSearchElderByKeyVO) error {
	if in.PageNum == nil || in.PageSize == nil {
		return constant.ErrParamInvalid
	}
	q := db.Table(tblelder.TableName)
	if in.Name != nil && *in.Name != "" {
		q = q.Where(tblelder.Name.Like(*in.Name))
	}
	if in.Phone != nil && *in.Phone != "" {
		q = q.And(tblelder.Phone.Like(*in.Phone))
	}
	var elders []do.Elder
	has, e := q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(tblelder.Id, tblelder.Name, tblelder.IdNum, tblelder.Sex, tblelder.Phone, tblelder.Address, tblelder.CheckFlag).
		Desc(tblelder.Id).
		Select().Gets(ctx, &elders)
	if e != nil {
		return e
	}
	if !has {
		return nil
	}
	res := make([]dto.PageSearchElderByKeyVO, 0, len(elders))
	for _, el := range elders {
		res = append(res, dto.PageSearchElderByKeyVO{
			ID:        int64(el.Id),
			Name:      el.Name.String(),
			IDNum:     el.IdNum.String(),
			Sex:       el.Sex.String(),
			Phone:     el.Phone.String(),
			Address:   el.Address.String(),
			CheckFlag: constant.CheckStatus(el.CheckFlag).String(),
		})
	}
	*out = res
	return nil
}

// PageSearchStaffByKey 分页查询员工（供预定选择经办人）
func (s *reserveService) PageSearchStaffByKey(ctx context.Context, in *dto.PageSearchStaffByKeyQuery, out *[]dto.PageSearchStaffByKeyVO) error {
	if in.PageNum == nil || in.PageSize == nil {
		return constant.ErrParamInvalid
	}
	q := db.Table(tblstaff.TableName)
	if in.Name != nil && *in.Name != "" {
		q = q.Where(tblstaff.Name.Like(*in.Name))
	}
	if in.Phone != nil && *in.Phone != "" {
		q = q.And(tblstaff.Phone.Like(*in.Phone))
	}
	var staffs []do.Staff
	has, e := q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(tblstaff.Id, tblstaff.Name, tblstaff.Phone).
		Desc(tblstaff.Id).
		Select().Gets(ctx, &staffs)
	if e != nil {
		return e
	}
	if !has {
		return nil
	}
	res := make([]dto.PageSearchStaffByKeyVO, 0, len(staffs))
	for _, st := range staffs {
		res = append(res, dto.PageSearchStaffByKeyVO{
			ID:    int64(st.Id),
			Name:  st.Name.String(),
			Phone: st.Phone.String(),
		})
	}
	*out = res
	return nil
}

// GetBuildTree 查询楼栋-房间-床位树（供预定选择床位）
func (s *reserveService) GetBuildTree(ctx context.Context, in *dto.IDReq, out *[]dto.BuildingVO) error {
	buildings := make([]do.Building, 0)
	has, e := dao.Building(db).List(ctx, ace.Where(tblbuilding.DelFlag.Eq(types.Int8(constant.YesNoNo))))
	if e != nil {
		return e
	}
	if !has {
		return nil
	}
	res := make([]dto.BuildingVO, 0, len(buildings))
	for _, b := range buildings {
		res = append(res, dto.BuildingVO{
			ID:   int64(b.Id),
			Name: b.Name.String(),
		})
	}
	*out = res
	return nil
}

// GetReserveByReserveIdAndElderId 按预定编号与老人编号查询预定
func (s *reserveService) GetReserveByReserveIdAndElderId(ctx context.Context, in *dto.GetReserveByReserveIDAndElderIDQuery, out *dto.GetReserveByReserveIDAndElderIDVO) error {
	if in.ReserveID == nil || in.ElderID == nil {
		return constant.ErrParamInvalid
	}
	rec := new(do.Reserve)
	has, e := dao.Reserve(db).Get(ctx, ace.Where(tblreserve.Id.Eq(types.BigInt(*in.ReserveID))).
		And(tblreserve.ElderId.Eq(types.BigInt(*in.ElderID))))
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrDataNotExist
	}
	out.ReserveID = int64(rec.Id)
	out.ElderID = int64(rec.ElderId)
	out.StaffID = int64(rec.StaffId)
	out.DueDate = rec.DueDate.Time()
	out.Deposit = rec.Deposit.Float64()
	out.Remark = rec.Remark.String()
	out.ReserveFlag = constant.YesNo(rec.ReserveFlag).String()
	return nil
}

// Refund 退预定（退款、释放床位）
func (s *reserveService) Refund(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	rec := new(do.Reserve)
	has, e := dao.Reserve(db).Get(ctx, ace.Where(tblreserve.Id.Eq(types.BigInt(*in.ID))))
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrDataNotExist
	}
	if _, e = dao.Reserve(db).UpdateById(ctx, *in.ID, tblreserve.ReserveFlag.Set(types.Int8(constant.YesNoYes))); e != nil {
		return e
	}
	// 释放老人床位并回退状态
	if rec.ElderId != 0 {
		el, found, eerr := dao.Elder(db).Get(ctx, ace.Where(tblelder.Id.Eq(rec.ElderId)))
		if eerr == nil && found {
			if el.BedId != 0 {
				_, _ = dao.Bed(db).UpdateById(ctx, int64(el.BedId), tblbed.BedFlag.Set(types.Int8(constant.BedIdle)))
				_, _ = dao.Elder(db).UpdateById(ctx, int64(rec.ElderId), tblelder.BedId.Set(types.BigInt(0)))
			}
			_, _ = dao.Elder(db).UpdateById(ctx, int64(rec.ElderId), tblelder.CheckFlag.Set(types.Int8(constant.CheckIntention)))
		}
	}
	return nil
}

// ReserveExpireJob 预定到期定时任务（将到期未付的预定标记退款）
func (s *reserveService) ReserveExpireJob(ctx context.Context, in *dto.EmptyReq, out *dto.EmptyResp) error {
	now := time.Now()
	_, e := dao.Reserve(db).Update(ctx, ace.Where(tblreserve.ReserveFlag.Eq(types.Int8(constant.YesNoNo))).
		And(tblreserve.DueDate.Lte(types.Time{Time: now})).
		Assign(ace.NewUpdateBuilder().Set(tblreserve.ReserveFlag.Set(types.Int8(constant.YesNoYes)))))
	return e
}

var _ = time.Now
