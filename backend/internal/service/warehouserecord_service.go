package service

import (
	"context"
	"database/sql"
	"strings"

	"api/internal/constant"
	"api/internal/lib"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblmaterial"
	"api/internal/model/define/table/tblstaff"
	"api/internal/model/define/table/tblstock"
	"api/internal/model/define/table/tblwarehouse"
	"api/internal/model/define/table/tblwarehouserecord"
	"api/internal/model/define/table/tblwarehouserecorditem"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

var WarehouseRecord = (*warehouseRecordService)(nil)

type warehouseRecordService struct{}

// findingStock 入库接收入库单联表（仓库名/经办人名）结果的中间结构体
type warehouseRecordJoin struct {
	ID            types.BigInt `json:"id"`
	WarehouseDate types.Time   `json:"warehouse_date"`
	Source        types.String `json:"source"`
	WarehouseFlag types.Int8   `json:"warehouse_flag"`
	WarehouseName types.String `json:"warehouse_name"`
	StaffName     types.String `json:"staff_name"`
}

// findOrCreateStock 查找给定仓库+物资+批次的库存台账行，不存在则创建（qty 暂为 0）。
// 返回 stock 行指针。
func findOrCreateStock(ctx context.Context, exec ace.Executer, tenantID, warehouseID, materialID types.BigInt, batchNo string, productDate, expireDate types.Time, createID types.BigInt) (*do.Stock, error) {
	if st, has, e := dao.Stock(exec).Get(ctx, ace.Where(
		tblstock.TenantId.Eq(tenantID),
		tblstock.WarehouseId.Eq(warehouseID),
		tblstock.MaterialId.Eq(materialID),
		tblstock.BatchNo.Eq(batchNo),
	)); e != nil {
		return nil, e
	} else if has {
		return st, nil
	}
	// 不存在则新建
	st := new(do.Stock)
	st.TenantId = tenantID
	st.WarehouseId = warehouseID
	st.MaterialId = materialID
	st.BatchNo = types.String(batchNo)
	st.ProductDate = productDate
	st.ExpireDate = expireDate
	st.Qty = 0
	st.CreateId = createID
	id, e := dao.Stock(exec).Insert(ctx,
		tblstock.TenantId.Set(tenantID),
		tblstock.WarehouseId.Set(warehouseID),
		tblstock.MaterialId.Set(materialID),
		tblstock.BatchNo.Set(batchNo),
		tblstock.ProductDate.Set(productDate),
		tblstock.ExpireDate.Set(expireDate),
		tblstock.Qty.Set(0),
		tblstock.CreateId.Set(createID),
		tblstock.CreateTime.Set(types.Now()),
	)
	if e != nil {
		return nil, e
	}
	st.Id = types.BigInt(id)
	return st, nil
}

// PageWarehouseRecordByKey 分页查询入库记录
func (s *warehouseRecordService) PageWarehouseRecordByKey(ctx context.Context, in *dto.PageWarehouseRecordByKeyReq, out *[]dto.PageWarehouseRecordByKeyResp) error {
	q := db.Table(tblwarehouserecord.TableName).
		InnerJoin(tblwarehouserecord.WarehouseId, tblwarehouse.Id).
		InnerJoin(tblwarehouserecord.StaffId, tblstaff.Id).
		Where(tblwarehouserecord.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblwarehouserecord.State.NotEq(types.Int8(constant.StateDeleted)))
	if in.WarehouseName != nil {
		q = q.And(tblwarehouse.Name.Like(*in.WarehouseName))
	}
	if in.StaffName != nil {
		q = q.And(tblstaff.Name.Like(*in.StaffName))
	}
	if in.StartTime != nil {
		q = q.And(tblwarehouserecord.WarehouseDate.Gte(types.Time{Time: *in.StartTime}))
	}
	if in.EndTime != nil {
		q = q.And(tblwarehouserecord.WarehouseDate.Lte(types.Time{Time: *in.EndTime}))
	}
	var records []warehouseRecordJoin
	e := q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblwarehouserecord.Id,
			tblwarehouserecord.WarehouseDate,
			tblwarehouserecord.Source,
			tblwarehouserecord.Status,
			tblwarehouse.Name.As("warehouse_name"),
			tblstaff.Name.As("staff_name"),
		).
		Desc(tblwarehouserecord.WarehouseDate).
		Select().Gets(ctx, &records)
	if e != nil {
		return e
	}
	// 聚合每笔记录的物资名称
	ids := make([]any, 0, len(records))
	for _, r := range records {
		ids = append(ids, r.ID)
	}
	items, _, e := dao.WarehouseRecordItem(db).List(ctx, ace.Where(tblwarehouserecorditem.RecordId.In(ids...)).
		Cols(tblwarehouserecorditem.RecordId, tblwarehouserecorditem.MaterialId))
	if e != nil {
		return e
	}
	materialIds := make([]any, 0)
	for _, it := range items {
		materialIds = append(materialIds, it.MaterialId)
	}
	nameMap := make(map[int64]string)
	if len(materialIds) > 0 {
		ms, _, e := dao.Material(db).List(ctx,
			ace.Where(tblmaterial.Id.In(materialIds...)).
				And(tblmaterial.State.NotEq(types.Int8(constant.StateDeleted))).
				Cols(tblmaterial.Id, tblmaterial.Name))
		if e != nil {
			return e
		}
		for _, m := range ms {
			nameMap[int64(m.Id)] = m.Name.String()
		}
	}
	matMap := make(map[int64][]string)
	for _, it := range items {
		if n, ok := nameMap[int64(it.MaterialId)]; ok {
			matMap[int64(it.RecordId)] = append(matMap[int64(it.RecordId)], n)
		}
	}
	res := make([]dto.PageWarehouseRecordByKeyResp, 0, len(records))
	for _, r := range records {
		names := matMap[int64(r.ID)]
		if in.MaterialName != nil && !strings.Contains(strings.Join(names, ","), *in.MaterialName) {
			continue
		}
		res = append(res, dto.PageWarehouseRecordByKeyResp{
			ID:            types.BigInt(r.ID),
			WarehouseName: r.WarehouseName.String(),
			MaterialName:  strings.Join(names, ","),
			WarehouseDate: r.WarehouseDate.Time,
			Source:        r.Source.String(),
			StaffName:     r.StaffName.String(),
			WarehouseFlag: constant.AuditStatus(r.WarehouseFlag).String(),
		})
	}
	*out = res
	return nil
}

// ListWarehouse 仓库下拉
func (s *warehouseRecordService) ListWarehouse(ctx context.Context, in *dto.EmptyReq, out *[]dto.DropDown) error {
	list, _, e := dao.Warehouse(db).List(ctx, ace.Where(tblwarehouse.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblwarehouse.State.NotEq(types.Int8(constant.StateDeleted))).Cols(tblwarehouse.Id, tblwarehouse.Name))
	if e != nil {
		return e
	}
	res := make([]dto.DropDown, 0, len(list))
	for _, w := range list {
		res = append(res, dto.DropDown{ID: types.BigInt(w.Id), Name: w.Name.String()})
	}
	*out = res
	return nil
}

// PageMaterialByKey 分页查询库存台账（供出库选择入库批次）
func (s *warehouseRecordService) PageMaterialByKey(ctx context.Context, in *dto.PageWarehouseMaterialByKeyReq, out *[]dto.PageWarehouseMaterialByKeyResp) error {
	q := db.Table(tblstock.TableName).
		InnerJoin(tblstock.MaterialId, tblmaterial.Id).
		Where(tblstock.TenantId.Eq(types.BigInt(lib.TenantID(ctx))))
	if in.WarehouseID != nil {
		q = q.And(tblstock.WarehouseId.Eq(types.BigInt(*in.WarehouseID)))
	}
	if in.MaterialName != nil {
		q = q.And(tblmaterial.Name.Like(*in.MaterialName))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblstock.Id.As("id"),
			tblmaterial.Name.As("material_name"),
			tblmaterial.Price.As("price"),
			tblstock.Qty.As("inventory"),
			tblstock.ExpireDate.As("expire_date"),
		).
		Desc(tblstock.Id).
		Select().Gets(ctx, out)
}

// AddWarehouseRecord 新增入库记录（仅登记，不入库。库存增减在审核通过时生效）
func (s *warehouseRecordService) AddWarehouseRecord(ctx context.Context, in *dto.AddWarehouseRecordReq, out *dto.EmptyResp) error {
	if in.WarehouseID == nil || in.StaffID == nil || len(in.WarehouseMaterialQueryList) == 0 {
		return constant.ErrParamInvalid
	}
	record := new(do.WarehouseRecord)
	record.TenantId = types.BigInt(lib.TenantID(ctx))
	record.WarehouseId = types.BigInt(*in.WarehouseID)
	record.StaffId = types.BigInt(*in.StaffID)
	record.Source = types.String(*in.Source)
	record.WarehouseDate = types.Time{Time: *in.WarehouseDate}
	record.Status = types.Uint8(constant.AuditStay)
	record.State = types.Int8(constant.StateEnabled)
	record.CreateId = types.BigInt(*in.StaffID)

	_, err := db.Transaction(ctx, func(tx *ace.Tx) (any, error) {
		id, e := dao.WarehouseRecord(tx).Insert(ctx,
			tblwarehouserecord.TenantId.Set(record.TenantId),
			tblwarehouserecord.WarehouseId.Set(record.WarehouseId),
			tblwarehouserecord.StaffId.Set(record.StaffId),
			tblwarehouserecord.Source.Set(record.Source),
			tblwarehouserecord.WarehouseDate.Set(record.WarehouseDate),
			tblwarehouserecord.Status.Set(types.Int8(record.Status)),
			tblwarehouserecord.State.Set(types.Int8(record.State)),
			tblwarehouserecord.CreateId.Set(record.CreateId),
			tblwarehouserecord.CreateTime.Set(types.Now()),
		)
		if e != nil {
			return nil, e
		}
		record.Id = types.BigInt(id)
		beans := make([]*do.WarehouseRecordItem, 0, len(in.WarehouseMaterialQueryList))
		for i := range in.WarehouseMaterialQueryList {
			m := &in.WarehouseMaterialQueryList[i]
			it := new(do.WarehouseRecordItem)
			it.TenantId = record.TenantId
			it.RecordId = record.Id
			it.MaterialId = types.BigInt(*m.MaterialID)
			it.Qty = types.Int32(*m.WarehouseNum)
			if m.ProductDate != nil {
				it.ProductDate = types.Time{Time: *m.ProductDate}
			}
			if m.ExpireDate != nil {
				it.ExpireDate = types.Time{Time: *m.ExpireDate}
			}
			it.CreateId = record.CreateId
			// StockId 留空，审核通过时回填
			beans = append(beans, it)
		}
		if _, e := dao.WarehouseRecordItem(tx).InsertBatch(ctx, beans,
			tblwarehouserecorditem.TenantId, tblwarehouserecorditem.RecordId, tblwarehouserecorditem.MaterialId,
			tblwarehouserecorditem.ProductDate, tblwarehouserecorditem.ExpireDate,
			tblwarehouserecorditem.Qty, tblwarehouserecorditem.CreateId, tblwarehouserecorditem.CreateTime,
			tblwarehouserecorditem.UpdateId, tblwarehouserecorditem.UpdateTime); e != nil {
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

// GetWarehouseRecordById 查询入库记录详情
func (s *warehouseRecordService) GetWarehouseRecordById(ctx context.Context, in *dto.IDReq, out *dto.GetWarehouseRecordByIDResp) error {
	e := db.Table(tblwarehouserecord.TableName).
		InnerJoin(tblwarehouserecord.WarehouseId, tblwarehouse.Id).
		InnerJoin(tblwarehouserecord.StaffId, tblstaff.Id).
		Where(tblwarehouserecord.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblwarehouserecord.Id.Eq(types.BigInt(*in.ID))).
		Cols(
			tblwarehouserecord.Id,
			tblwarehouserecord.Source,
			tblwarehouserecord.WarehouseDate,
			tblwarehouse.Name.As("warehouse_name"),
			tblstaff.Name.As("staff_name"),
		).
		Select().
		Get(ctx, out)

	switch e {
	case nil:
	case sql.ErrNoRows:
		return nil
	default:
		return e
	}

	// 入库物资列表（明细自带生产/效期）
	items, _, ie := dao.WarehouseRecordItem(db).List(ctx, ace.Where(
		tblwarehouserecorditem.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
		tblwarehouserecorditem.RecordId.Eq(types.BigInt(*in.ID)),
	))
	if ie != nil {
		return ie
	}
	for _, it := range items {
		nm := ""
		if mm, mh, me := dao.Material(db).Get(ctx, ace.Where(tblmaterial.Id.Eq(it.MaterialId)).Cols(tblmaterial.Name)); me == nil && mh {
			nm = mm.Name.String()
		}
		out.WarehouseMaterialByIDRespList = append(out.WarehouseMaterialByIDRespList, dto.GetWarehouseMaterialByIDResp{
			MaterialName: nm,
			WarehouseNum: int(it.Qty),
			ProductDate:  it.ProductDate.Time,
			ExpireDate:   it.ExpireDate.Time,
		})
	}
	return nil
}

// AuditWarehouseRecord 审核入库记录：只在「通过」时把明细计入库存台账
func (s *warehouseRecordService) AuditWarehouseRecord(ctx context.Context, in *dto.AuditWarehouseRecordReq, out *dto.EmptyResp) error {
	if in.WarehouseRecordID == nil || in.AuditResult == nil {
		return constant.ErrParamInvalid
	}
	record, has, e := dao.WarehouseRecord(db).
		Get(ctx,
			ace.Where(
				tblwarehouserecord.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
				tblwarehouserecord.Id.Eq(types.BigInt(*in.WarehouseRecordID)),
			).
				Cols(
					tblwarehouserecord.Id,
					tblwarehouserecord.Status,
					tblwarehouserecord.WarehouseId,
				))
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrDataNotExist
	}
	if record.Status != types.Uint8(constant.AuditStay) {
		return constant.ErrAuditRepeat
	}

	flag := constant.AuditNotPass
	if *in.AuditResult == "通过" {
		flag = constant.AuditPass
	}

	_, err := db.Transaction(ctx, func(tx *ace.Tx) (any, error) {
		if _, e := dao.WarehouseRecord(tx).UpdateById(ctx, types.BigInt(*in.WarehouseRecordID),
			tblwarehouserecord.Status.Set(types.Int8(flag))); e != nil {
			return nil, e
		}
		// 通过时：逐条明细计入库存台账
		if flag == constant.AuditPass {
			items, _, e := dao.WarehouseRecordItem(tx).List(ctx, ace.Where(
				tblwarehouserecorditem.RecordId.Eq(types.BigInt(*in.WarehouseRecordID)),
			).Cols(tblwarehouserecorditem.Id, tblwarehouserecorditem.MaterialId, tblwarehouserecorditem.ProductDate, tblwarehouserecorditem.ExpireDate, tblwarehouserecorditem.Qty))
			if e != nil {
				return nil, e
			}
			for _, it := range items {
				st, e := findOrCreateStock(ctx, tx, record.TenantId, record.WarehouseId, it.MaterialId, "", it.ProductDate, it.ExpireDate, record.CreateId)
				if e != nil {
					return nil, e
				}
				if _, e := dao.Stock(tx).Update(ctx, []dialect.Setter{tblstock.Qty.Incr(it.Qty)},
					tblstock.Id.Eq(st.Id)); e != nil {
					return nil, e
				}
				if _, e := dao.WarehouseRecordItem(tx).Update(ctx, []dialect.Setter{
					tblwarehouserecorditem.StockId.Set(st.Id),
				}, tblwarehouserecorditem.Id.Eq(it.Id)); e != nil {
					return nil, e
				}
			}
		}
		return nil, nil
	})
	if err != nil {
		return err
	}
	_ = out
	return nil
}

// DeleteWarehouseRecord 删除入库记录（逻辑删除；若已审核通过则回冲库存）
func (s *warehouseRecordService) DeleteWarehouseRecord(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	if in.ID == nil {
		return constant.ErrParamInvalid
	}
	record, has, e := dao.WarehouseRecord(db).
		Get(ctx, ace.Where(
			tblwarehouserecord.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
			tblwarehouserecord.Id.Eq(types.BigInt(*in.ID)),
		).Cols(tblwarehouserecord.Id, tblwarehouserecord.Status, tblwarehouserecord.State))
	if e != nil {
		return e
	}
	if !has || record.State == types.Int8(constant.StateDeleted) {
		return constant.ErrDelRepeat
	}

	_, err := db.Transaction(ctx, func(tx *ace.Tx) (any, error) {
		if _, e := dao.WarehouseRecord(tx).UpdateById(ctx, types.BigInt(*in.ID),
			tblwarehouserecord.State.Set(types.Int8(constant.StateDeleted))); e != nil {
			return nil, e
		}
		// 若已审核通过，回冲已入账库存
		if record.Status == types.Uint8(constant.AuditPass) {
			items, _, e := dao.WarehouseRecordItem(tx).List(ctx, ace.Where(
				tblwarehouserecorditem.RecordId.Eq(types.BigInt(*in.ID)),
			).Cols(tblwarehouserecorditem.StockId, tblwarehouserecorditem.Qty))
			if e != nil {
				return nil, e
			}
			for _, it := range items {
				if it.StockId == nil {
					continue
				}
				if _, e := dao.Stock(tx).Update(ctx, []dialect.Setter{tblstock.Qty.Decr(it.Qty)},
					tblstock.Id.Eq(*it.StockId)); e != nil {
					return nil, e
				}
			}
		}
		return nil, nil
	})
	if err != nil {
		return err
	}
	_ = out
	return nil
}
