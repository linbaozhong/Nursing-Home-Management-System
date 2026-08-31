package service

import (
	"context"

	"api/internal/constant"
	"api/internal/lib"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblbed"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/define/table/tblmaterial"
	"api/internal/model/define/table/tbloutboundrecord"
	"api/internal/model/define/table/tbloutboundrecorditem"
	"api/internal/model/define/table/tblstaff"
	"api/internal/model/define/table/tblstock"
	"api/internal/model/define/table/tblwarehouse"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

var OutboundRecord = (*outboundRecordService)(nil)

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
func (s *outboundRecordService) PageOutboundRecordByKey(ctx context.Context, in *dto.PageOutboundRecordByKeyReq, out *[]dto.PageOutboundRecordByKeyResp) error {
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
			tbloutboundrecord.Status,
			tblwarehouse.Name.As("warehouse_name"),
			tblstaff.Name.As("staff_name"),
			tblelder.Name.As("elder_name"),
		).
		Desc(tbloutboundrecord.OutboundDate).
		Select().Gets(ctx, &joins)
	if e != nil {
		return e
	}
	res := make([]dto.PageOutboundRecordByKeyResp, 0, len(joins))
	for _, j := range joins {
		res = append(res, dto.PageOutboundRecordByKeyResp{
			ID:            types.BigInt(j.ID),
			WarehouseName: j.WarehouseName.String(),
			OutboundDate:  j.OutboundDate.Time,
			Recipient:     j.ElderName.String(),
			StaffName:     j.StaffName.String(),
			OutboundFlag:  constant.AuditStatus(j.OutboundFlag).String(),
		})
	}
	*out = res
	return nil
}

// GetOutboundRecordById 查询出库记录详情
func (s *outboundRecordService) GetOutboundRecordById(ctx context.Context, in *dto.IDReq, out *dto.GetOutboundRecordByIDResp) error {
	rec, has, e := dao.OutboundRecord(db).Get(ctx, ace.Where(tbloutboundrecord.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tbloutboundrecord.Id.Eq(types.BigInt(*in.ID))))
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrDataNotExist
	}
	out.ID = types.BigInt(rec.Id)
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
	items, _, me := dao.OutboundRecordItem(db).List(ctx, ace.Where(tbloutboundrecorditem.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tbloutboundrecorditem.RecordId.Eq(rec.Id)))
	if me == nil {
		for _, it := range items {
			mn := ""
			if mm, mh, merr := dao.Material(db).Get(ctx, ace.Where(tblmaterial.Id.Eq(it.MaterialId)).Cols(tblmaterial.Name)); merr == nil && mh {
				mn = mm.Name.String()
			}
			out.OutboundMaterialByIDRespList = append(out.OutboundMaterialByIDRespList, dto.GetOutboundMaterialByIDResp{
				MaterialName: mn,
				OutboundNum:  int(it.Qty),
			})
		}
	}
	return nil
}

// PageSearchElderByKey 分页查询老人（供出库选择领用人）
func (s *outboundRecordService) PageSearchElderByKey(ctx context.Context, in *dto.PageSearchElderByKeyReq, out *[]dto.PageSearchElderByKeyResp) error {
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
		Cols(tblelder.Id, tblelder.Name, tblelder.IdNum, tblelder.Sex, tblelder.Phone, tblelder.Address, tblelder.Status).
		Desc(tblelder.Id).
		Select().Gets(ctx, &elders)
	if e != nil {
		return e
	}
	res := make([]dto.PageSearchElderByKeyResp, 0, len(elders))
	for i := range elders {
		el := &elders[i]
		res = append(res, dto.PageSearchElderByKeyResp{
			ID:        types.BigInt(el.Id),
			Name:      el.Name.String(),
			IDNum:     el.IdNum.String(),
			Sex:       el.Sex.String(),
			Phone:     el.Phone.String(),
			Address:   el.Address.String(),
			CheckFlag: el.Status.String(),
		})
	}
	*out = res
	return nil
}

// PageWarehouseMaterialByKey 分页查询库存台账（供出库选择，可按仓库过滤）
func (s *outboundRecordService) PageWarehouseMaterialByKey(ctx context.Context, in *dto.PageWarehouseMaterialByKeyReq, out *[]dto.PageWarehouseMaterialByKeyResp) error {
	if in.PageNum == nil || in.PageSize == nil {
		return constant.ErrParamInvalid
	}
	q := db.Table(tblstock.TableName).
		LeftJoin(tblstock.MaterialId, tblmaterial.Id).
		Where(tblstock.TenantId.Eq(types.BigInt(lib.TenantID(ctx))))
	if in.WarehouseID != nil {
		q = q.And(tblstock.WarehouseId.Eq(types.BigInt(*in.WarehouseID)))
	}
	if in.MaterialName != nil && *in.MaterialName != "" {
		q = q.And(tblmaterial.Name.Like(*in.MaterialName))
	}
	var list []struct {
		ID           types.BigInt `json:"id"`
		MaterialName types.String `json:"material_name"`
		Price        types.Money  `json:"price"`
		Inventory    types.BigInt `json:"inventory"`
		ExpireDate   types.Time   `json:"expire_date"`
	}
	e := q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblstock.Id,
			tblstock.Qty.As("inventory"),
			tblstock.ExpireDate,
			tblmaterial.Name.As("material_name"),
			tblmaterial.Price.As("price"),
		).
		Desc(tblstock.Id).
		Select().Gets(ctx, &list)
	if e != nil {
		return e
	}
	res := make([]dto.PageWarehouseMaterialByKeyResp, 0, len(list))
	for _, m := range list {
		res = append(res, dto.PageWarehouseMaterialByKeyResp{
			ID:           types.BigInt(m.ID),
			MaterialName: m.MaterialName.String(),
			Price:        m.Price,
			WarehouseNum: 0,
			Inventory:    int(m.Inventory),
			ExpireDate:   m.ExpireDate.Time,
		})
	}
	*out = res
	return nil
}

// AddOutboundRecord 新增出库记录（仅登记，不扣库存。库存扣减在审核通过时生效）
func (s *outboundRecordService) AddOutboundRecord(ctx context.Context, in *dto.AddOutboundRecordReq, out *dto.EmptyResp) error {
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
		Status:        types.Uint8(constant.AuditStay),
		State:         types.Int8(constant.StateEnabled),
		CreateId:      types.BigInt(*in.StaffID),
	}
	_, err := db.Transaction(ctx, func(tx *ace.Tx) (any, error) {
		id, e := dao.OutboundRecord(tx).Insert(ctx,
			tbloutboundrecord.TenantId.Set(rec.TenantId),
			tbloutboundrecord.RecipientType.Set(rec.RecipientType),
			tbloutboundrecord.MaterialUse.Set(rec.MaterialUse),
			tbloutboundrecord.WarehouseId.Set(rec.WarehouseId),
			tbloutboundrecord.StaffId.Set(rec.StaffId),
			tbloutboundrecord.RecipientId.Set(rec.RecipientId),
			tbloutboundrecord.OutboundDate.Set(rec.OutboundDate),
			tbloutboundrecord.Status.Set(types.Int8(rec.Status)),
			tbloutboundrecord.State.Set(types.Int8(rec.State)),
			tbloutboundrecord.CreateId.Set(rec.CreateId),
			tbloutboundrecord.CreateTime.Set(types.Now()),
		)
		if e != nil {
			return nil, e
		}
		rec.Id = types.BigInt(id)
		beans := make([]*do.OutboundRecordItem, 0, len(in.OutboundMaterialQueryList))
		for i := range in.OutboundMaterialQueryList {
			m := &in.OutboundMaterialQueryList[i]
			if m.StockID == nil || m.OutboundNum == nil {
				return nil, constant.ErrParamInvalid
			}
			// 校验所选库存行属于该仓库，并取物资id
			st, has, e := dao.Stock(tx).Get(ctx, ace.Where(
				tblstock.TenantId.Eq(rec.TenantId),
				tblstock.Id.Eq(types.BigInt(*m.StockID)),
			).Cols(tblstock.Id, tblstock.WarehouseId, tblstock.MaterialId))
			if e != nil {
				return nil, e
			}
			if !has {
				return nil, constant.ErrDataNotExist
			}
			if st.WarehouseId != rec.WarehouseId {
				return nil, constant.ErrParamInvalid
			}
			it := new(do.OutboundRecordItem)
			it.TenantId = rec.TenantId
			it.RecordId = rec.Id
			it.StockId = &st.Id
			it.MaterialId = st.MaterialId
			it.Qty = types.Int32(*m.OutboundNum)
			it.CreateId = rec.CreateId
			beans = append(beans, it)
		}
		if _, e := dao.OutboundRecordItem(tx).InsertBatch(ctx, beans,
			tbloutboundrecorditem.TenantId, tbloutboundrecorditem.RecordId, tbloutboundrecorditem.StockId, tbloutboundrecorditem.MaterialId,
			tbloutboundrecorditem.Qty, tbloutboundrecorditem.CreateId, tbloutboundrecorditem.CreateTime,
			tbloutboundrecorditem.UpdateId, tbloutboundrecorditem.UpdateTime); e != nil {
			return nil, e
		}
		return nil, nil
	})
	if err != nil {
		return err
	}
	_ = out
	return nil
}

// AuditOutboundRecord 审核出库记录：通过时扣减库存（并发安全，禁止超卖）
func (s *outboundRecordService) AuditOutboundRecord(ctx context.Context, in *dto.AuditOutboundRecordReq, out *dto.EmptyResp) error {
	if in.OutboundRecordID == nil || in.AuditResult == nil {
		return constant.ErrParamInvalid
	}
	rec, has, e := dao.OutboundRecord(db).Get(ctx,
		ace.Where(
			tbloutboundrecord.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
			tbloutboundrecord.Id.Eq(types.BigInt(*in.OutboundRecordID)),
		).Cols(tbloutboundrecord.Id, tbloutboundrecord.Status, tbloutboundrecord.WarehouseId, tbloutboundrecord.TenantId))
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrDataNotExist
	}
	if rec.Status != types.Uint8(constant.AuditStay) {
		return constant.ErrAuditRepeat
	}

	flag := constant.AuditNotPass
	if *in.AuditResult == "通过" {
		flag = constant.AuditPass
	}

	_, tErr := db.Transaction(ctx, func(tx *ace.Tx) (any, error) {
		if _, e := dao.OutboundRecord(tx).UpdateById(ctx, types.BigInt(*in.OutboundRecordID),
			tbloutboundrecord.Status.Set(types.Int8(flag))); e != nil {
			return nil, e
		}
		if flag != constant.AuditPass {
			return nil, nil
		}
		items, _, e := dao.OutboundRecordItem(tx).List(ctx, ace.Where(
			tbloutboundrecorditem.RecordId.Eq(types.BigInt(*in.OutboundRecordID)),
		).Cols(tbloutboundrecorditem.Id, tbloutboundrecorditem.StockId, tbloutboundrecorditem.Qty))
		if e != nil {
			return nil, e
		}
		for _, it := range items {
			if it.StockId == nil {
				return nil, constant.ErrParamInvalid
			}
			// 原子扣减：qty = qty - n WHERE id=? AND tenant_id=? AND qty>=n
			res, e := tx.ExecContext(ctx,
				"UPDATE `stock` SET `qty` = `qty` - ? WHERE `id` = ? AND `tenant_id` = ? AND `qty` >= ?",
				int32(it.Qty), int64(*it.StockId), int64(rec.TenantId), int32(it.Qty))
			if e != nil {
				return nil, e
			}
			rows, _ := res.RowsAffected()
			if rows == 0 {
				return nil, constant.ErrOutboundNumError
			}
		}
		return nil, nil
	})
	if tErr != nil {
		return tErr
	}
	_ = out
	return nil
}

// DeleteOutboundRecord 删除出库记录（逻辑删除；若已审核通过则回补库存）
func (s *outboundRecordService) DeleteOutboundRecord(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	if in.ID == nil {
		return constant.ErrParamInvalid
	}
	rec, has, e := dao.OutboundRecord(db).Get(ctx, ace.Where(
		tbloutboundrecord.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
		tbloutboundrecord.Id.Eq(types.BigInt(*in.ID)),
	).Cols(tbloutboundrecord.Id, tbloutboundrecord.Status, tbloutboundrecord.State))
	if e != nil {
		return e
	}
	if !has || rec.State == types.Int8(constant.StateDeleted) {
		return constant.ErrDelRepeat
	}
	_, tErr := db.Transaction(ctx, func(tx *ace.Tx) (any, error) {
		if _, e := dao.OutboundRecord(tx).UpdateById(ctx, types.BigInt(*in.ID),
			tbloutboundrecord.State.Set(types.Int8(constant.StateDeleted))); e != nil {
			return nil, e
		}
		// 若已审核通过，回补库存
		if rec.Status == types.Uint8(constant.AuditPass) {
			items, _, e := dao.OutboundRecordItem(tx).List(ctx, ace.Where(
				tbloutboundrecorditem.RecordId.Eq(types.BigInt(*in.ID)),
			).Cols(tbloutboundrecorditem.StockId, tbloutboundrecorditem.Qty))
			if e != nil {
				return nil, e
			}
			for _, it := range items {
				if it.StockId == nil {
					continue
				}
				if _, e := dao.Stock(tx).Update(ctx, []dialect.Setter{tblstock.Qty.Incr(it.Qty)},
					tblstock.Id.Eq(*it.StockId)); e != nil {
					return nil, e
				}
			}
		}
		return nil, nil
	})
	if tErr != nil {
		return tErr
	}
	_ = out
	return nil
}
