package service

import (
	"context"
	"database/sql"
	"strings"

	"api/internal/constant"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblmaterial"
	"api/internal/model/define/table/tblstaff"
	"api/internal/model/define/table/tblwarehouse"
	"api/internal/model/define/table/tblwarehousematerial"
	"api/internal/model/define/table/tblwarehouserecord"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/types"
)

var _ = (*warehouseRecordService)(nil)

type warehouseRecordService struct{}

// warehouseRecordJoin 接收入库记录联表（仓库名/经办人名）查询结果的中间结构体
type warehouseRecordJoin struct {
	ID            types.BigInt `json:"id"`
	WarehouseDate types.Time   `json:"warehouse_date"`
	Source        types.String `json:"source"`
	WarehouseFlag types.Int8   `json:"warehouse_flag"`
	WarehouseName types.String `json:"warehouse_name"`
	StaffName     types.String `json:"staff_name"`
}

// PageWarehouseRecordByKey 分页查询入库记录
func (s *warehouseRecordService) PageWarehouseRecordByKey(ctx context.Context, in *dto.PageWarehouseRecordByKeyQuery, out *[]dto.PageWarehouseRecordByKeyVO) error {
	q := db.Table(tblwarehouserecord.TableName).
		InnerJoin(tblwarehouserecord.WarehouseId, tblwarehouse.Id).
		InnerJoin(tblwarehouserecord.StaffId, tblstaff.Id).
		Where(tblwarehouserecord.DelFlag.Eq(types.Int8(constant.YesNoNo)))
	if in.WarehouseName != nil {
		q = q.And(tblwarehouse.Name.Like(*in.WarehouseName))
	}
	if in.StaffName != nil {
		q = q.And(tblstaff.Name.Like(*in.StaffName))
	}
	if in.StartTime != nil {
		q = q.And(tblwarehouserecord.WarehouseDate.Gte(*in.StartTime))
	}
	if in.EndTime != nil {
		q = q.And(tblwarehouserecord.WarehouseDate.Lte(*in.EndTime))
	}
	var records []warehouseRecordJoin
	e := q.Page(*in.PageNum, *in.PageSize).
		Cols(
			tblwarehouserecord.Id,
			tblwarehouserecord.WarehouseDate,
			tblwarehouserecord.Source,
			tblwarehouserecord.WarehouseFlag,
			tblwarehouse.Name.AsName("warehouse_name"),
			tblstaff.Name.AsName("staff_name"),
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
	wms, _, e := dao.WarehouseMaterial(db).List(ctx, ace.Where(tblwarehousematerial.WarehouseRecordId.In(ids...)).
		Cols(tblwarehousematerial.WarehouseRecordId, tblwarehousematerial.MaterialId))
	if e != nil {
		return e
	}
	materialIds := make([]any, 0)
	for _, wm := range wms {
		materialIds = append(materialIds, wm.MaterialId)
	}
	nameMap := make(map[int64]string)
	if len(materialIds) > 0 {
		ms, _, e := dao.Material(db).List(ctx,
			ace.Where(tblmaterial.Id.In(materialIds...)).
				And(tblmaterial.DelFlag.Eq(types.Int8(constant.YesNoNo))).
				Cols(tblmaterial.Id, tblmaterial.Name))
		if e != nil {
			return e
		}
		for _, m := range ms {
			nameMap[int64(m.Id)] = m.Name.String()
		}
	}
	matMap := make(map[int64][]string)
	for _, wm := range wms {
		if n, ok := nameMap[int64(wm.MaterialId)]; ok {
			matMap[int64(wm.WarehouseRecordId)] = append(matMap[int64(wm.WarehouseRecordId)], n)
		}
	}
	res := make([]dto.PageWarehouseRecordByKeyVO, 0, len(records))
	for _, r := range records {
		names := matMap[int64(r.ID)]
		if in.MaterialName != nil && !strings.Contains(strings.Join(names, ","), *in.MaterialName) {
			continue
		}
		res = append(res, dto.PageWarehouseRecordByKeyVO{
			ID:            int64(r.ID),
			WarehouseName: r.WarehouseName.String(), // 来自 As
			MaterialName:  strings.Join(names, ","),
			WarehouseDate: r.WarehouseDate.Time,
			Source:        r.Source.String(),
			StaffName:     r.StaffName.String(), // 来自 As
			WarehouseFlag: constant.AuditStatus(r.WarehouseFlag).String(),
		})
	}
	*out = res
	return nil
}

// ListWarehouse 仓库下拉
func (s *warehouseRecordService) ListWarehouse(ctx context.Context, in *dto.EmptyReq, out *[]dto.DropDown) error {
	list, _, e := dao.Warehouse(db).List(ctx, ace.Where(tblwarehouse.DelFlag.Eq(types.Int8(constant.YesNoNo))).Cols(tblwarehouse.Id, tblwarehouse.Name))
	if e != nil {
		return e
	}
	res := make([]dto.DropDown, 0, len(list))
	for _, w := range list {
		res = append(res, dto.DropDown{ID: int64(w.Id), Name: w.Name.String()})
	}
	*out = res
	return nil
}

// PageMaterialByKey 分页查询仓库物资
func (s *warehouseRecordService) PageMaterialByKey(ctx context.Context, in *dto.PageWarehouseMaterialByKeyQuery, out *[]dto.PageWarehouseMaterialByKeyVO) error {
	q := db.Table(tblwarehousematerial.TableName).
		InnerJoin(tblwarehousematerial.MaterialId, tblmaterial.Id).
		Where(tblwarehousematerial.WarehouseRecordId.Gt(0))
	if in.WarehouseID != nil {
		// 通过 warehouse_material 的库存记录不分仓库，此处仅按物资名过滤
		_ = in.WarehouseID
	}
	if in.MaterialName != nil {
		q = q.And(tblmaterial.Name.Like(*in.MaterialName))
	}
	return q.Page(*in.PageNum, *in.PageSize).
		Cols(
			tblwarehousematerial.Id.As("id"),
			tblmaterial.Name.As("material_name"),
			tblmaterial.Price.As("price"),
			tblwarehousematerial.WarehouseNum.As("warehouse_num"),
			tblwarehousematerial.Inventory.As("inventory"),
			tblwarehousematerial.ExpireDate.As("expire_date"),
		).
		Desc(tblwarehousematerial.Id).
		Select().Gets(ctx, out)
}

// AddWarehouseRecord 新增入库记录
func (s *warehouseRecordService) AddWarehouseRecord(ctx context.Context, in *dto.AddWarehouseRecordQuery, out *dto.EmptyResp) error {
	record := new(do.WarehouseRecord)
	record.WarehouseId = types.BigInt(*in.WarehouseID)
	record.StaffId = types.BigInt(*in.StaffID)
	record.Source = types.String(*in.Source)
	record.WarehouseDate = types.Time{*in.WarehouseDate}
	record.WarehouseFlag = types.Int8(constant.AuditStay)
	record.DelFlag = types.Int8(constant.YesNoNo)
	if _, e := dao.WarehouseRecord(db).InsertOne(ctx, record,
		tblwarehouserecord.WarehouseId, tblwarehouserecord.StaffId, tblwarehouserecord.Source,
		tblwarehouserecord.WarehouseDate, tblwarehouserecord.WarehouseFlag, tblwarehouserecord.DelFlag,
		tblwarehouserecord.CreateId, tblwarehouserecord.CreateTime); e != nil {
		return e
	}
	// 按物资编号分组汇总数量
	group := make(map[int64]*dto.AddWarehouseMaterialQuery)
	for i := range in.WarehouseMaterialQueryList {
		m := &in.WarehouseMaterialQueryList[i]
		if g, ok := group[*m.MaterialID]; ok {
			*g.WarehouseNum += *m.WarehouseNum
		} else {
			cp := *m
			group[*m.MaterialID] = &cp
		}
	}
	beans := make([]*do.WarehouseMaterial, 0, len(group))
	for mid, m := range group {
		wm := new(do.WarehouseMaterial)
		wm.WarehouseRecordId = record.Id
		wm.MaterialId = types.BigInt(mid)
		wm.ProductDate = types.Time{*m.ProductDate}
		wm.ExpireDate = types.Time{*m.ExpireDate}
		wm.WarehouseNum = types.Int32(*m.WarehouseNum)
		wm.Inventory = types.Int32(*m.WarehouseNum)
		beans = append(beans, wm)
	}
	if _, e := dao.WarehouseMaterial(db).InsertBatch(ctx, beans,
		tblwarehousematerial.WarehouseRecordId, tblwarehousematerial.MaterialId, tblwarehousematerial.ProductDate,
		tblwarehousematerial.ExpireDate, tblwarehousematerial.WarehouseNum, tblwarehousematerial.Inventory); e != nil {
		return e
	}
	return nil
}

// GetWarehouseRecordById 查询入库记录详情
func (s *warehouseRecordService) GetWarehouseRecordById(ctx context.Context, in *dto.IDReq, out *dto.GetWarehouseRecordByIDVO) error {
	e := db.Table(tblwarehouserecord.TableName).
		InnerJoin(tblwarehouserecord.WarehouseId, tblwarehouse.Id).
		InnerJoin(tblwarehouserecord.StaffId, tblstaff.Id).
		Where(tblwarehouserecord.Id.Eq(types.BigInt(*in.ID))).
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

	// 入库物资列表
	e = db.Table(tblwarehousematerial.TableName).
		InnerJoin(tblwarehousematerial.MaterialId, tblmaterial.Id).
		Where(tblwarehousematerial.WarehouseRecordId.Eq(types.BigInt(*in.ID))).
		Cols(
			tblmaterial.Name.As("material_name"),
			tblwarehousematerial.WarehouseNum,
			tblwarehousematerial.ProductDate,
			tblwarehousematerial.ExpireDate,
		).
		Select().
		Gets(ctx, &out.WarehouseMaterialByIDVOList)

	switch e {
	case nil, sql.ErrNoRows:
		return nil
	default:
		return e
	}
}

// AuditWarehouseRecord 审核入库记录
func (s *warehouseRecordService) AuditWarehouseRecord(ctx context.Context, in *dto.AuditWarehouseRecordQuery, out *dto.EmptyResp) error {
	record, has, e := dao.WarehouseRecord(db).
		Get(ctx,
			ace.Where(
				tblwarehouserecord.Id.Eq(types.BigInt(*in.WarehouseRecordID)),
			).
				Cols(
					tblwarehouserecord.Id,
					tblwarehouserecord.WarehouseFlag,
				))
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrDataNotExist
	}
	if record.WarehouseFlag != types.Int8(constant.AuditStay) {
		return constant.ErrAuditRepeat
	}
	flag := constant.AuditNotPass
	if *in.AuditResult == "通过" {
		flag = constant.AuditPass
	}
	_, e = dao.WarehouseRecord(db).UpdateById(ctx, types.BigInt(*in.WarehouseRecordID),
		tblwarehouserecord.WarehouseFlag.Set(types.Int8(flag)),
	)
	if e != nil {
		return e
	}
	return nil
}

// DeleteWarehouseRecord 删除入库记录（逻辑删除）
func (s *warehouseRecordService) DeleteWarehouseRecord(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.WarehouseRecord(db).UpdateById(ctx, types.BigInt(*in.ID),
		tblwarehouserecord.DelFlag.Set(types.Int8(constant.YesNoYes)),
	)
	if e != nil {
		return e
	}
	return nil
}
