package service

import (
	"api/internal/model/define/table/tblbed"
	"context"

	"api/internal/constant"
	"api/internal/lib"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/define/table/tblmaterial"
	"api/internal/model/define/table/tbloutboundmaterial"
	"api/internal/model/define/table/tbloutboundrecord"
	"api/internal/model/define/table/tblstaff"
	"api/internal/model/define/table/tblwarehouse"
	"api/internal/model/define/table/tblwarehousematerial"
	"api/internal/model/do"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/types"
)

var _ = (*outboundRecordService)(nil)

type outboundRecordService struct{}

// outboundRecordJoin 接收入库记录联表（仓库名/经办人名/领用人名）查询结果的中间结构体
type outboundRecordJoin struct {
	ID            types.BigInt `json:"id"`
	RecipientType types.String `json:"recipient_type"`
	OutboundDate  types.Time   `json:"outbound_date"`
	OutboundFlag  types.Int8   `json:"outbound_flag"`
	WarehouseName types.String `json:"warehouse_name"`
	StaffName     types.String `json:"staff_name"`
	ElderName     types.String `json:"elder_name"`
}

// PageOutboundRecordByKey 分页查询出库记录
func (s *outboundRecordService) PageOutboundRecordByKey(ctx context.Context, in *dto.PageOutboundRecordByKeyQuery, out *[]dto.PageOutboundRecordByKeyVO) error {
	if in.PageNum == nil || in.PageSize == nil {
		return constant.ErrParamInvalid
	}
	q := db.Table(tbloutboundrecord.TableName).
		LeftJoin(tbloutboundrecord.WarehouseId, tblwarehouse.Id).
		LeftJoin(tbloutboundrecord.StaffId, tblstaff.Id).
		LeftJoin(tbloutboundrecord.RecipientId, tblelder.Id).
		Where(tbloutboundrecord.TenantId.Eq(types.BigInt(lib.TenantID(ctx))))
	if in.WarehouseName != nil && *in.WarehouseName != "" {
		q = q.Where(tblwarehouse.Name.Like(*in.WarehouseName))
	}
	if in.Recipient != nil && *in.Recipient != "" {
		q = q.And(tblelder.Name.Like(*in.Recipient))
	}
	if in.StartTime != nil {
		q = q.And(tbloutboundrecord.OutboundDate.Gte(types.Time{Time: *in.StartTime}))
	}
	if in.EndTime != nil {
		q = q.And(tbloutboundrecord.OutboundDate.Lte(types.Time{Time: *in.EndTime}))
	}
	var joins []outboundRecordJoin
	e := q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tbloutboundrecord.Id,
			tbloutboundrecord.RecipientType,
			tbloutboundrecord.OutboundDate,
			tbloutboundrecord.OutboundFlag,
			tblwarehouse.Name.As("warehouse_name"),
			tblstaff.Name.As("staff_name"),
			tblelder.Name.As("elder_name"),
		).
		Desc(tbloutboundrecord.OutboundDate).
		Select().Gets(ctx, &joins)
	if e != nil {
		return e
	}
	res := make([]dto.PageOutboundRecordByKeyVO, 0, len(joins))
	for _, j := range joins {
		res = append(res, dto.PageOutboundRecordByKeyVO{
			ID:            int64(j.ID),
			WarehouseName: j.WarehouseName.String(),
			OutboundDate:  j.OutboundDate.Time(),
			Recipient:     j.ElderName.String(),
			StaffName:     j.StaffName.String(),
			OutboundFlag:  constant.AuditStatus(j.OutboundFlag).String(),
		})
	}
	*out = res
	return nil
}

// GetOutboundRecordById 查询出库记录详情
func (s *outboundRecordService) GetOutboundRecordById(ctx context.Context, in *dto.IDReq, out *dto.GetOutboundRecordByIDVO) error {
	rec, has, e := dao.OutboundRecord(db).Get(ctx, ace.Where(tbloutboundrecord.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tbloutboundrecord.Id.Eq(types.BigInt(*in.ID))))
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrDataNotExist
	}
	out.ID = int64(rec.Id)
	out.RecipientType = rec.RecipientType.String()
	out.MaterialUse = rec.MaterialUse.String()
	out.OutboundDate = rec.OutboundDate.Time
	// 仓库名
	if w, wh, we := dao.Warehouse(db).Get(ctx, ace.Where(tblwarehouse.Id.Eq(rec.WarehouseId))); we == nil && wh {
		out.WarehouseName = w.Name.String()
	}
	// 登记人
	if st, sh, se := dao.Staff(db).Get(ctx, ace.Where(tblstaff.Id.Eq(rec.StaffId))); se == nil && sh {
		out.StaffName = st.Name.String()
	}
	// 出库物资明细
	materials, _, me := dao.OutboundMaterial(db).List(ctx, ace.Where(tbloutboundmaterial.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tbloutboundmaterial.OutboundRecordId.Eq(rec.Id)))
	if me == nil {
		for _, m := range materials {
			mn := ""
			if mm, mh, merr := dao.Material(db).Get(ctx, ace.Where(tblmaterial.Id.Eq(m.MaterialId))); merr == nil && mh {
				mn = mm.Name.String()
			}
			out.OutboundMaterialByIDVOList = append(out.OutboundMaterialByIDVOList, dto.GetOutboundMaterialByIDVO{
				MaterialName: mn,
				OutboundNum:  int(m.OutboundNum),
			})
		}
	}
	return nil
}

// PageSearchElderByKey 分页查询老人（供出库选择领用人）
func (s *outboundRecordService) PageSearchElderByKey(ctx context.Context, in *dto.PageSearchElderByKeyQuery, out *[]dto.PageSearchElderByKeyVO) error {
	if in.PageNum == nil || in.PageSize == nil {
		return constant.ErrParamInvalid
	}
	q := db.Table(tblelder.TableName).
		InnerJoin(tblelder.BedId, tblbed.Id).
		Where(tblelder.TenantId.Eq(types.BigInt(lib.TenantID(ctx)))).
		Cols(
			tblelder.Id,
			tblelder.Name,
			tblelder.IdNum,
			tblelder.Sex,
			tblbed.Name.As("phone"), // 床位名称,借用elder表的phone字段
		)
	if in.Name != nil && *in.Name != "" {
		q = q.Where(tblelder.Name.Like(*in.Name))
	}
	if in.Phone != nil && *in.Phone != "" {
		q = q.And(tblelder.Phone.Like(*in.Phone))
	}
	var elders []do.Elder
	e := q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(tblelder.Id, tblelder.Name, tblelder.IdNum, tblelder.Sex, tblelder.Phone, tblelder.Address, tblelder.CheckFlag).
		Desc(tblelder.Id).
		Select().Gets(ctx, &elders)
	if e != nil {
		return e
	}
	res := make([]dto.PageSearchElderByKeyVO, 0, len(elders))
	for _, el := range elders {
		res = append(res, dto.PageSearchElderByKeyVO{
			ElderID:   int64(el.Id),
			ElderName: el.Name.String(),
			IDNum:     el.IdNum.String(),
			ElderSex:  el.Sex.String(),
			BedName:   el.Phone.String(),
		})
	}
	*out = res
	return nil
}

// PageWarehouseMaterialByKey 分页查询仓库物资（供出库选择）
func (s *outboundRecordService) PageWarehouseMaterialByKey(ctx context.Context, in *dto.PageWarehouseMaterialByKeyQuery, out *[]dto.PageWarehouseMaterialByKeyVO) error {
	if in.PageNum == nil || in.PageSize == nil || in.WarehouseID == nil {
		return constant.ErrParamInvalid
	}
	q := db.Table(tblwarehousematerial.TableName).
		LeftJoin(tblwarehousematerial.MaterialId, tblmaterial.Id).
		Where(tblwarehousematerial.TenantId.Eq(types.BigInt(lib.TenantID(ctx))))
	if in.MaterialName != nil && *in.MaterialName != "" {
		q = q.Where(tblmaterial.Name.Like(*in.MaterialName))
	}
	var list []struct {
		ID           types.BigInt `json:"id"`
		MaterialName types.String `json:"material_name"`
		Price        types.Money  `json:"price"`
		WarehouseNum types.Int    `json:"warehouse_num"`
		Inventory    types.BigInt `json:"inventory"`
		ExpireDate   types.Time   `json:"expire_date"`
	}
	e := q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblwarehousematerial.Id,
			tblwarehousematerial.Price,
			tblwarehousematerial.WarehouseNum,
			tblwarehousematerial.Inventory,
			tblwarehousematerial.ExpireDate,
			tblmaterial.Name.As("material_name"),
		).
		Desc(tblwarehousematerial.Id).
		Select().Gets(ctx, &list)
	if e != nil {
		return e
	}
	res := make([]dto.PageWarehouseMaterialByKeyVO, 0, len(list))
	for _, m := range list {
		res = append(res, dto.PageWarehouseMaterialByKeyVO{
			ID:           int64(m.ID),
			MaterialName: m.MaterialName.String(),
			Price:        m.Price,
			WarehouseNum: int(m.WarehouseNum),
			Inventory:    int(m.Inventory),
			ExpireDate:   m.ExpireDate.Time,
		})
	}
	*out = res
	return nil
}

// AddOutboundRecord 新增出库记录并扣减库存
func (s *outboundRecordService) AddOutboundRecord(ctx context.Context, in *dto.AddOutboundRecordQuery, out *dto.EmptyResp) error {
	if in.WarehouseID == nil || in.StaffID == nil || len(in.OutboundMaterialQueryList) == 0 {
		return constant.ErrParamInvalid
	}
	rec := &do.OutboundRecord{
		TenantId:      types.BigInt(lib.TenantID(ctx)),
		RecipientType: types.String(orEmpty(in.RecipientType)),
		MaterialUse:   types.String(orEmpty(in.MaterialUse)),
		WarehouseId:   types.BigInt(*in.WarehouseID),
		StaffId:       types.BigInt(*in.StaffID),
		RecipientId:   types.BigInt(*in.RecipientID),
		OutboundDate:  types.Time{Time: *in.OutboundDate},
		OutboundFlag:  types.Int8(constant.AuditStay),
		CreateId:      types.BigInt(*in.StaffID),
	}
	if _, e := dao.OutboundRecord(db).InsertOne(ctx, rec); e != nil {
		return e
	}
	for _, m := range in.OutboundMaterialQueryList {
		if m.WarehouseMaterialID != nil && m.OutboundNum != nil {
			// 扣减对应仓库物资库存
			wm, has, e := dao.WarehouseMaterial(db).Get(ctx, ace.Where(tblwarehousematerial.Id.Eq(types.BigInt(*m.WarehouseMaterialID))))
			if e == nil && has {
				newInv := int32(wm.Inventory) - int32(*m.OutboundNum)
				if newInv < 0 {
					newInv = 0
				}
				_, _ = dao.WarehouseMaterial(db).UpdateById(ctx, wm.Id, tblwarehousematerial.Inventory.Set(types.Int32(newInv)))
			}
			om := &do.OutboundMaterial{
				TenantId:            types.BigInt(lib.TenantID(ctx)),
				OutboundRecordId:    rec.Id,
				WarehouseMaterialId: types.BigInt(*m.WarehouseMaterialID),
				MaterialId:          wm.MaterialId,
				OutboundNum:         types.Int32(int32(*m.OutboundNum)),
				CreateId:            rec.CreateId,
			}
			_, _ = dao.OutboundMaterial(db).InsertOne(ctx, om)
		}
	}
	return nil
}

// AuditOutboundRecord 审核出库记录（通过/不通过）
func (s *outboundRecordService) AuditOutboundRecord(ctx context.Context, in *dto.AuditOutboundRecordQuery, out *dto.EmptyResp) error {
	if in.OutboundRecordID == nil || in.AuditResult == nil {
		return constant.ErrParamInvalid
	}
	rec, has, e := dao.OutboundRecord(db).Get(ctx, ace.Where(tbloutboundrecord.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tbloutboundrecord.Id.Eq(types.BigInt(*in.OutboundRecordID))))
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrDataNotExist
	}
	if *in.AuditResult == "不通过" {
		if _, e = dao.OutboundRecord(db).UpdateById(ctx, types.BigInt(*in.OutboundRecordID), tbloutboundrecord.OutboundFlag.Set(types.Int8(constant.AuditNotPass))); e != nil {
			return e
		}
		return nil
	}
	if _, e = dao.OutboundRecord(db).UpdateById(ctx, types.BigInt(*in.OutboundRecordID), tbloutboundrecord.OutboundFlag.Set(types.Int8(constant.AuditPass))); e != nil {
		return e
	}
	return nil
}

// DeleteOutboundRecord 删除出库记录（逻辑删除）
func (s *outboundRecordService) DeleteOutboundRecord(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	if _, e := dao.OutboundRecord(db).UpdateById(ctx, types.BigInt(*in.ID), tbloutboundrecord.DelFlag.Set(types.Int8(constant.YesNoYes))); e != nil {
		return e
	}
	return nil
}
