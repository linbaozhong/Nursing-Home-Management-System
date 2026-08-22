package service

import (
	"context"

	"api/internal/constant"
	"api/internal/lib"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblbed"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/define/table/tblnursereserve"
	"api/internal/model/define/table/tblstaff"
	"api/internal/model/do"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/types"
)

var _ = (*nurseReserveService)(nil)

type nurseReserveService struct{}

// nurseReserveJoin 接收护理预定联表（老人姓名、床位名、护理员姓名）查询结果的中间结构体
type nurseReserveJoin struct {
	ID           types.BigInt `json:"id"`
	ElderName    types.String `json:"elder_name"`
	BedName      types.String `json:"bed_name"`
	ServiceName  types.String `json:"service_name"`
	NeedDate     types.Int32  `json:"need_date"`
	ServicePrice types.Money  `json:"service_price"`
	ChargeMethod types.String `json:"charge_method"`
	Frequency    types.Int32  `json:"frequency"`
	PayAmount    types.Money  `json:"pay_amount"`
	NurseDate    types.Time   `json:"nurse_date"`
	OrderFlag    types.Int8   `json:"order_flag"`
}

// PageNurseReserveByKey 分页查询护理预定
func (s *nurseReserveService) PageNurseReserveByKey(ctx context.Context, in *dto.PageNurseReserveByKeyReq, out *[]dto.PageNurseReserveByKeyResp) error {
	if in.PageNum == nil || in.PageSize == nil {
		return constant.ErrParamInvalid
	}
	q := db.Table(tblnursereserve.TableName).
		LeftJoin(tblnursereserve.ElderId, tblelder.Id).
		LeftJoin(tblelder.BedId, tblbed.Id).
		Where(tblnursereserve.TenantId.Eq(types.BigInt(lib.TenantID(ctx))))
	if in.ElderName != nil && *in.ElderName != "" {
		q = q.Where(tblelder.Name.Like(*in.ElderName))
	}
	if in.ServiceName != nil && *in.ServiceName != "" {
		q = q.And(tblnursereserve.ServiceName.Like(*in.ServiceName))
	}
	if in.BedName != nil && *in.BedName != "" {
		q = q.And(tblbed.Name.Like(*in.BedName))
	}
	var joins []nurseReserveJoin
	e := q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblnursereserve.Id,
			tblnursereserve.ServiceName,
			tblnursereserve.NeedDate,
			tblnursereserve.ServicePrice,
			tblnursereserve.ChargeMethod,
			tblnursereserve.Frequency,
			tblnursereserve.PayAmount,
			tblnursereserve.NurseDate,
			tblnursereserve.OrderFlag,
			tblelder.Name.As("elder_name"),
			tblbed.Name.As("bed_name"),
		).
		Desc(tblnursereserve.Id).
		Select().Gets(ctx, &joins)
	if e != nil {
		return e
	}
	res := make([]dto.PageNurseReserveByKeyResp, 0, len(joins))
	for _, j := range joins {
		res = append(res, dto.PageNurseReserveByKeyResp{
			ID:           int64(j.ID),
			ElderName:    j.ElderName.String(),
			BedName:      j.BedName.String(),
			ServiceName:  j.ServiceName.String(),
			NeedDate:     int(j.NeedDate),
			ServicePrice: j.ServicePrice,
			ChargeMethod: j.ChargeMethod.String(),
			Frequency:    int(j.Frequency),
			PayAmount:    j.PayAmount,
			NurseDate:    j.NurseDate.Time,
			OrderFlag:    constant.YesNo(j.OrderFlag).String(),
		})
	}
	*out = res
	return nil
}

// GetNurseReserveByReserveIdAndElderId 按护理预定/老人编号查询护理预定
func (s *nurseReserveService) GetNurseReserveByReserveIdAndElderId(ctx context.Context, in *dto.GetNurseReserveByReserveIdAndElderIdReq, out *dto.GetNurseReserveByReserveIdAndElderIdResp) error {
	if in.ReserveID == nil || in.ElderID == nil {
		return constant.ErrParamInvalid
	}

	nr, has, e := dao.NurseReserve(db).Get(ctx, ace.Where(tblnursereserve.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblnursereserve.Id.Eq(types.BigInt(*in.ReserveID))).
		And(tblnursereserve.ElderId.Eq(types.BigInt(*in.ElderID))))
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrDataNotExist
	}
	out.ID = int64(nr.Id)
	out.ServiceName = nr.ServiceName.String()
	out.NeedDate = int(nr.NeedDate)
	out.ServicePrice = nr.ServicePrice
	out.ChargeMethod = nr.ChargeMethod.String()
	out.Frequency = int(nr.Frequency)
	out.PayAmount = nr.PayAmount
	out.NurseDate = nr.NurseDate.Time
	out.OrderFlag = constant.YesNo(nr.OrderFlag).String()
	// 老人姓名、床位名
	if el, eh, ee := dao.Elder(db).Get(ctx, ace.Where(tblelder.Id.Eq(nr.ElderId))); ee == nil && eh {
		out.ElderName = el.Name.String()
		if el.BedId != 0 {
			if b, bh, be := dao.Bed(db).Get(ctx, ace.Where(tblbed.Id.Eq(el.BedId))); be == nil && bh {
				out.BedName = b.Name.String()
			}
		}
	}
	return nil
}

// AddNurseReserve 新增护理预定
func (s *nurseReserveService) AddNurseReserve(ctx context.Context, in *dto.AddNurseReserveReq, out *dto.EmptyResp) error {
	if in.ElderID == nil || in.ServiceName == nil || in.PayAmount == nil {
		return constant.ErrParamInvalid
	}
	rec := &do.NurseReserve{
		TenantId:     types.BigInt(lib.TenantID(ctx)),
		ElderId:      types.BigInt(*in.ElderID),
		ServiceName:  types.String(*in.ServiceName),
		NeedDate:     types.Int32(int32(*in.NeedDate)),
		ServicePrice: types.Money(*in.ServicePrice),
		ChargeMethod: types.String(*in.ChargeMethod),
		Frequency:    types.Int32(int32(*in.Frequency)),
		PayAmount:    types.Money(*in.PayAmount),
		OrderFlag:    types.Int8(constant.YesNoNo),
		CreateId:     types.BigInt(*in.ElderID),
	}
	if _, e := dao.NurseReserve(db).InsertOne(ctx, rec); e != nil {
		return e
	}
	return nil
}

// EditNurseReserve 编辑护理预定
func (s *nurseReserveService) EditNurseReserve(ctx context.Context, in *dto.EditNurseReserveReq, out *dto.EmptyResp) error {
	if in.ID == nil {
		return constant.ErrParamInvalid
	}
	if _, e := dao.NurseReserve(db).UpdateById(ctx, types.BigInt(*in.ID),
		tblnursereserve.ElderId.Set(types.BigInt(*in.ElderID)),
		tblnursereserve.NeedDate.Set(types.Int32(int32(*in.NeedDate))),
		tblnursereserve.ChargeMethod.Set(types.String(*in.ChargeMethod)),
		tblnursereserve.Frequency.Set(types.Int32(int32(*in.Frequency))),
		tblnursereserve.ServicePrice.Set(types.Money(*in.ServicePrice)),
		tblnursereserve.PayAmount.Set(types.Money(*in.PayAmount)),
	); e != nil {
		return e
	}
	return nil
}

// DeleteNurseReserve 删除护理预定（物理删除）
func (s *nurseReserveService) DeleteNurseReserve(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	if _, e := dao.NurseReserve(db).DeleteById(ctx, types.BigInt(*in.ID)); e != nil {
		return e
	}
	return nil
}

// PageSearchElderByKey 分页查询老人（供护理预定选择）
func (s *nurseReserveService) PageSearchElderByKey(ctx context.Context, in *dto.PageSearchElderByKeyReq, out *[]dto.PageSearchElderByKeyResp) error {
	if in.PageNum == nil || in.PageSize == nil {
		return constant.ErrParamInvalid
	}
	q := db.Table(tblelder.TableName).
		Where(tblelder.TenantId.Eq(types.BigInt(lib.TenantID(ctx))))
	if in.Name != nil && *in.Name != "" {
		q = q.Where(tblelder.Name.Like(*in.Name))
	}
	if in.Phone != nil && *in.Phone != "" {
		q = q.And(tblelder.Phone.Like(*in.Phone))
	}
	var elders []do.Elder
	e := q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(tblelder.Id, tblelder.Name, tblelder.IdNum, tblelder.Sex, tblelder.Phone, tblelder.Address, tblelder.CheckFlag, tblelder.BedId).
		Desc(tblelder.Id).
		Select().Gets(ctx, &elders)
	if e != nil {
		return e
	}
	res := make([]dto.PageSearchElderByKeyResp, 0, len(elders))
	for _, el := range elders {
		bedName := ""
		if el.BedId != 0 {
			if b, bh, be := dao.Bed(db).Get(ctx, ace.Where(tblbed.Id.Eq(el.BedId))); be == nil && bh {
				bedName = b.Name.String()
			}
		}
		res = append(res, dto.PageSearchElderByKeyResp{
			ElderID:   int64(el.Id),
			ElderName: el.Name.String(),
			IDNum:     el.IdNum.String(),
			ElderSex:  el.Sex.String(),
			BedName:   bedName,
		})
	}
	*out = res
	return nil
}

// ListNurseStaff 查询护理员工（护理员列表）
func (s *nurseReserveService) ListNurseStaff(ctx context.Context, in *dto.EmptyReq, out *[]dto.PageSearchStaffByKeyResp) error {
	staffs, has, e := dao.Staff(db).List(ctx,
		ace.Where(tblstaff.TenantId.Eq(types.BigInt(lib.TenantID(ctx)))).
			Cols(
				tblstaff.Id,
				tblstaff.Name,
				tblstaff.Phone,
			).Desc(tblstaff.Id))
	if e != nil {
		return e
	}
	if !has {
		return nil
	}
	res := make([]dto.PageSearchStaffByKeyResp, 0, len(staffs))
	for _, st := range staffs {
		res = append(res, dto.PageSearchStaffByKeyResp{
			ID:    int64(st.Id),
			Name:  st.Name.String(),
			Phone: st.Phone.String(),
		})
	}
	*out = res
	return nil
}
