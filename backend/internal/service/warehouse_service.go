package service

import (
	"context"
	"errors"

	"api/internal/constant"
	"api/internal/lib"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblwarehouse"
	"api/internal/model/define/table/tblwarehousematerial"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/conv"
	"github.com/linbaozhong/gentity/pkg/types"
)

type warehouse struct{}

var Warehouse = &warehouse{}

// PageWarehouseByKey 分页查询仓库
// 对应 Java: WarehouseServiceImpl.pageWarehouseByKey -> WarehouseMapper.listWarehouseByKey
// 注: staff_name 需联表 user 表(仓库管理员名), 因 user 表尚未在 Go 侧生成, 暂未联表, staff_name 留空。
func (w *warehouse) PageWarehouseByKey(ctx context.Context, in *dto.PageWarehouseByKeyReq, out *[]dto.PageWarehouseByKeyResp) error {
	q := db.Table(tblwarehouse.TableName).
		Where(tblwarehouse.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblwarehouse.DelFlag.Eq(constant.YesNoNo))
	if in.WarehouseName != nil && *in.WarehouseName != "" {
		q.And(tblwarehouse.Name.Like(*in.WarehouseName))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblwarehouse.Id.AsName("id"),
			tblwarehouse.Name.AsName("name"),
		).
		Desc(tblwarehouse.CreateTime).
		Select().
		Gets(ctx, out)
}

// ListWarehouseStaff 仓库管理员列表
// 对应 Java: WarehouseServiceImpl.listWarehouseStaff -> StaffFunc.listStaffByRoleId(6)
// 注: 需查询 user 表(role = 仓库管理员), 因 user 表尚未在 Go 侧生成, 暂留待 account/user 体系完成后补充。
func (w *warehouse) ListWarehouseStaff(ctx context.Context, in *dto.EmptyReq, out *dto.EmptyResp) error {
	// todo: 查询 user 表 role = 仓库管理员(角色编号 6), 返回人员下拉列表
	return nil
}

// AddWarehouse 新增仓库（校验名称不重复）
// 对应 Java: WarehouseServiceImpl.addWarehouse -> WarehouseFunc.getWarehouseByName
func (w *warehouse) AddWarehouse(ctx context.Context, in *dto.OperateWarehouseReq, out *dto.EmptyResp) error {
	repeat, e := dao.Warehouse(db).Exists(ctx,
		tblwarehouse.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
		tblwarehouse.Name.Eq(*in.Name),
		tblwarehouse.DelFlag.Eq(types.Int8(constant.YesNoNo)),
	)
	if e != nil {
		return e
	}
	if repeat {
		return errors.New("仓库名称已存在")
	}
	bean := do.NewWarehouse()
	bean.TenantId = types.BigInt(lib.TenantID(ctx))
	bean.StaffId = types.BigInt(*in.StaffID)
	bean.Name = types.String(*in.Name)
	bean.DelFlag = types.Int8(constant.YesNoNo)
	_, e = dao.Warehouse(db).InsertOne(ctx, bean)
	return e
}

// GetWarehouseById 根据编号获取仓库（编辑回显）
// 对应 Java: WarehouseServiceImpl.getWarehouseById
func (w *warehouse) GetWarehouseById(ctx context.Context, in *dto.IDReq, out *dto.OperateWarehouseResp) error {
	obj, has, e := dao.Warehouse(db).Get(ctx, ace.Where(tblwarehouse.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblwarehouse.Id.Eq(types.BigInt(*in.ID))))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("仓库不存在")
	}
	out.ID = conv.Ptr(int64(obj.Id))
	out.StaffID = conv.Ptr(int64(obj.StaffId))
	out.Name = conv.Ptr(obj.Name.String())
	return nil
}

// EditWarehouse 编辑仓库（校验名称不重复排除自身）
// 对应 Java: WarehouseServiceImpl.editWarehouse
func (w *warehouse) EditWarehouse(ctx context.Context, in *dto.OperateWarehouseReq, out *dto.EmptyResp) error {
	repeat, e := dao.Warehouse(db).Exists(ctx,
		tblwarehouse.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
		tblwarehouse.Name.Eq(*in.Name),
		tblwarehouse.DelFlag.Eq(types.Int8(constant.YesNoNo)),
		tblwarehouse.Id.NotEq(types.BigInt(*in.ID)),
	)
	if e != nil {
		return e
	}
	if repeat {
		return errors.New("仓库名称已存在")
	}
	bean := do.NewWarehouse()
	bean.Id = types.BigInt(*in.ID)
	bean.StaffId = types.BigInt(*in.StaffID)
	bean.Name = types.String(*in.Name)
	_, e = dao.Warehouse(db).UpdateOne(ctx, bean)
	return e
}

// DeleteWarehouse 删除仓库（存在关联物资记录则不允许删除）
// 对应 Java: WarehouseServiceImpl.deleteWarehouse -> warehouseMaterialMapper.sumWarehouseMaterialNumByWarehouseId
func (w *warehouse) DeleteWarehouse(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	hasMaterial, e := dao.WarehouseMaterial(db).Exists(ctx,
		tblwarehousematerial.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
		tblwarehousematerial.WarehouseRecordId.Eq(types.BigInt(*in.ID)),
	)
	if e != nil {
		return e
	}
	if hasMaterial {
		return errors.New("该仓库下存在物资，无法删除")
	}
	bean := do.NewWarehouse()
	bean.Id = types.BigInt(*in.ID)
	bean.DelFlag = types.Int8(constant.YesNoYes)
	_, e = dao.Warehouse(db).UpdateOne(ctx, bean)
	return e
}
