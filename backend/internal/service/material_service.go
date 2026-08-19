package service

import (
	"context"
	"errors"

	"api/internal/constant"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblmaterial"
	"api/internal/model/define/table/tblmaterialtype"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/types"
)

// materialTotalLimit 物资 / 物资分类数量上限（对应 Java Constant.TOTAL_LIMIT）
const materialTotalLimit = 10

type material struct{}

var Material = &material{}

// PageMaterialByKey 分页查询物资（关联物资类型名称）
// 对应 Java: MaterialServiceImpl.pageMaterialByKey -> MaterialMapper.listMaterialByKey
func (m *material) PageMaterialByKey(ctx context.Context, in *dto.PageMaterialByKeyQuery, out *[]dto.PageMaterialByKeyVO) error {
	q := db.Table(tblmaterial.TableName).
		LeftJoin(tblmaterial.TypeId, tblmaterialtype.Id).
		Where(tblmaterial.DelFlag.Eq(constant.YesNoNo))
	if in.MaterialName != nil && *in.MaterialName != "" {
		q.And(tblmaterial.Name.Like(*in.MaterialName))
	}
	if in.MaterialTypeID != nil {
		q.And(tblmaterial.TypeId.Eq(types.BigInt(*in.MaterialTypeID)))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblmaterial.Id.AsName("id"),
			tblmaterial.Name.AsName("name"),
			tblmaterialtype.Name.AsName("type_name"),
			tblmaterial.Price.AsName("price"),
		).
		Desc(tblmaterial.CreateTime).
		Select().
		Gets(ctx, out)
}

// GetMaterialById 根据编号获取物资（编辑回显）
// 对应 Java: MaterialServiceImpl.getMaterialById
func (m *material) GetMaterialById(ctx context.Context, in *dto.IDReq, out *dto.OperateMaterialVO) error {
	obj, has, e := dao.Material(db).GetByID(ctx, types.BigInt(*in.ID))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("物资不存在")
	}
	*out.ID = int64(obj.Id)
	*out.TypeID = int64(obj.TypeId)
	*out.Name = obj.Name.String()
	*out.Price = obj.Price
	return nil
}

// AddMaterial 新增物资（同一分类下名称不重复 + 分类数量上限）
// 对应 Java: MaterialServiceImpl.addMaterial -> MaterialFunc.getMaterialByName / checkTypeTotal
func (m *material) AddMaterial(ctx context.Context, in *dto.AddMaterialQuery, out *dto.EmptyResp) error {
	repeat, e := dao.Material(db).Exists(ctx,
		tblmaterial.TypeId.Eq(types.BigInt(*in.TypeID)),
		tblmaterial.Name.Eq(*in.Name),
		tblmaterial.DelFlag.Eq(constant.YesNoNo),
	)
	if e != nil {
		return e
	}
	if repeat {
		return errors.New("该分类下物资名称已存在")
	}
	total, e := dao.Material(db).Count(ctx,
		tblmaterial.TypeId.Eq(types.BigInt(*in.TypeID)),
		tblmaterial.DelFlag.Eq(constant.YesNoNo),
	)
	if e != nil {
		return e
	}
	if int(total) >= materialTotalLimit {
		return errors.New("该分类下物资数量已达上限")
	}
	bean := do.NewMaterial()
	bean.TypeId = types.BigInt(*in.TypeID)
	bean.Name = types.String(*in.Name)
	bean.Price = types.Money(*in.Price)
	bean.DelFlag = types.Int8(constant.YesNoNo)
	_, e = dao.Material(db).InsertOne(ctx, bean)
	return e
}

// EditMaterial 编辑物资（同一分类下名称不重复排除自身）
// 对应 Java: MaterialServiceImpl.editMaterial
func (m *material) EditMaterial(ctx context.Context, in *dto.EditMaterialQuery, out *dto.EmptyResp) error {
	repeat, e := dao.Material(db).Exists(ctx,
		tblmaterial.TypeId.Eq(types.BigInt(*in.TypeID)),
		tblmaterial.Name.Eq(*in.Name),
		tblmaterial.DelFlag.Eq(constant.YesNoNo),
		tblmaterial.Id.NotEq(*in.ID),
	)
	if e != nil {
		return e
	}
	if repeat {
		return errors.New("该分类下物资名称已存在")
	}
	bean := do.NewMaterial()
	bean.Id = types.BigInt(*in.ID)
	bean.TypeId = types.BigInt(*in.TypeID)
	bean.Name = types.String(*in.Name)
	bean.Price = types.Money(*in.Price)
	_, e = dao.Material(db).UpdateOne(ctx, bean)
	return e
}

// DeleteMaterial 删除物资（逻辑删除）
// 对应 Java: MaterialServiceImpl.deleteMaterial -> del_flag = YES
func (m *material) DeleteMaterial(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	bean := do.NewMaterial()
	bean.Id = types.BigInt(*in.ID)
	bean.DelFlag = types.Int8(constant.YesNoYes)
	_, e := dao.Material(db).UpdateOne(ctx, bean)
	return e
}

// PageMaterialTypeByKey 分页查询物资分类
// 对应 Java: MaterialServiceImpl.getMaterialType -> MaterialTypeFunc.listNotDelMaerialType
func (m *material) PageMaterialTypeByKey(ctx context.Context, in *dto.PageMaterialTypeByKeyQuery, out *[]dto.PageMaterialTypeVO) error {
	q := db.Table(tblmaterial.TableName).
		Where(tblmaterialtype.DelFlag.Eq(constant.YesNoNo))
	if in.Name != nil && *in.Name != "" {
		q.And(tblmaterialtype.Name.Like(*in.Name))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblmaterialtype.Id.AsName("id"),
			tblmaterialtype.Name.AsName("name"),
		).
		Desc(tblmaterialtype.CreateTime).
		Select().
		Gets(ctx, out)
}

// GetMaterialTypeById 根据编号获取物资分类（编辑回显）
// 对应 Java: MaterialServiceImpl.getMaterialTypeById
func (m *material) GetMaterialTypeById(ctx context.Context, in *dto.IDReq, out *dto.OperateMaterialTypeQuery) error {
	obj, has, e := dao.MaterialType(db).GetByID(ctx, types.BigInt(*in.ID))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("物资分类不存在")
	}
	out.ID = int64Ptr(int64(obj.Id))
	out.Name = strPtr(obj.Name.String())
	return nil
}

// AddMaterialType 新增物资分类（名称不重复 + 总数上限）
// 对应 Java: MaterialServiceImpl.addMaterialType -> MaterialTypeFunc.getMaterialTypeByName / checkTypeTotal
func (m *material) AddMaterialType(ctx context.Context, in *dto.AddMaterialTypeQuery, out *dto.EmptyResp) error {
	repeat, e := dao.MaterialType(db).Exists(ctx,
		tblmaterialtype.Name.Eq(*in.Name),
		tblmaterialtype.DelFlag.Eq(constant.YesNoNo),
	)
	if e != nil {
		return e
	}
	if repeat {
		return errors.New("物资分类名称已存在")
	}
	total, e := dao.MaterialType(db).Count(ctx, tblmaterialtype.DelFlag.Eq(constant.YesNoNo))
	if e != nil {
		return e
	}
	if int(total) >= materialTotalLimit {
		return errors.New("物资分类数量已达上限")
	}
	bean := do.NewMaterialType()
	bean.Name = types.String(*in.Name)
	bean.DelFlag = types.Int8(constant.YesNoNo)
	_, e = dao.MaterialType(db).InsertOne(ctx, bean)
	return e
}

// EditMaterialType 编辑物资分类（名称不重复排除自身）
// 对应 Java: MaterialServiceImpl.editMaterialType
func (m *material) EditMaterialType(ctx context.Context, in *dto.EditMaterialTypeQuery, out *dto.EmptyResp) error {
	repeat, e := dao.MaterialType(db).Exists(ctx,
		tblmaterialtype.Name.Eq(*in.Name),
		tblmaterialtype.DelFlag.Eq(constant.YesNoNo),
		tblmaterialtype.Id.NotEq(types.BigInt(*in.ID)),
	)
	if e != nil {
		return e
	}
	if repeat {
		return errors.New("物资分类名称已存在")
	}
	bean := do.NewMaterialType()
	bean.Id = types.BigInt(*in.ID)
	bean.Name = types.String(*in.Name)
	_, e = dao.MaterialType(db).UpdateOne(ctx, bean)
	return e
}

// DeleteMaterialType 删除物资分类（存在子物资则不允许删除）
// 对应 Java: MaterialServiceImpl.deleteMaterialType -> MaterialFunc.checkMaterialItem
func (m *material) DeleteMaterialType(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	hasChild, e := dao.Material(db).Exists(ctx,
		tblmaterial.TypeId.Eq(types.BigInt(*in.ID)),
		tblmaterial.DelFlag.Eq(constant.YesNoNo),
	)
	if e != nil {
		return e
	}
	if hasChild {
		return errors.New("该物资分类下存在物资，无法删除")
	}
	bean := do.NewMaterialType()
	bean.Id = types.BigInt(*in.ID)
	bean.DelFlag = types.Int8(constant.YesNoYes)
	_, e = dao.MaterialType(db).UpdateOne(ctx, bean)
	return e
}
