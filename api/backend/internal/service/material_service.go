package service

import (
	"context"

	"api/internal/model/define/dao"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type material struct{}

var Material = &material{}

// PageMaterialByKey 分页查询物资（联表 material_type 获取类型名）
// 对应 Java: MaterialServiceImpl.pageMaterialByKey -> MaterialMapper.listMaterialByKey
// SQL: SELECT m.*, mt.type_name FROM material m
//
//	LEFT JOIN material_type mt ON mt.id = m.material_type_id
//	WHERE (mt.type_name LIKE %key% OR m.material_name LIKE %key%) [可选]
//	ORDER BY m.create_time DESC; 再由 PageUtil 内存分页。
//
// todo: 1) in.Key 非空 -> (tbl<materialtype>.TypeName.Like(in.Key) OR tbl<material>.MaterialName.Like(in.Key))
//
//	2) DB 分页: Count + List(联表 LeftJoin)
//	3) 组装含类型名的 VO 并赋值 out
func (m *material) PageMaterialByKey(ctx context.Context, in *dto.PageMaterialByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 实现联表分页查询
	return nil
}

// GetMaterialById 根据编号获取物资
// 对应 Java: MaterialServiceImpl.getMaterialById -> materialMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.Material(db).GetByID(ctx, types.BigInt(in.ID))
func (m *material) GetMaterialById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.Material(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// AddMaterial 新增物资
// 对应 Java: MaterialServiceImpl.addMaterial -> materialMapper.insertSelective
// todo: 标准 CRUD - dao.Material(db).InsertOne 写入 material 表(含 materialName/materialTypeId/stock 等)
func (m *material) AddMaterial(ctx context.Context, in *dto.AddMaterialQuery, out *dto.EmptyResp) error {
	// todo: bean := do.NewMaterial(); 填充 in; dao.Material(db).InsertOne(ctx, bean)
	return nil
}

// EditMaterial 编辑物资
// 对应 Java: MaterialServiceImpl.editMaterial -> materialMapper.updateByPrimaryKeySelective
// todo: 标准 CRUD - 按主键更新 material 表(tbl<material>.X.Value(in.X))
func (m *material) EditMaterial(ctx context.Context, in *dto.EditMaterialQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<material>.MaterialName.Value(in.MaterialName),
	}
	_, e := dao.Material(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteMaterial 删除物资
// 对应 Java: MaterialServiceImpl.deleteMaterial -> materialMapper.deleteByPrimaryKey
// todo: 标准 CRUD - dao.Material(db).DeleteById(ctx, types.BigInt(in.ID))
func (m *material) DeleteMaterial(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.Material(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}

// PageMaterialTypeByKey 分页查询物资类型
// 对应 Java: MaterialServiceImpl.pageMaterialTypeByKey -> MaterialTypeMapper.listMaterialTypeByKey
// SQL: SELECT * FROM material_type WHERE (type_name LIKE %key%) [可选] ORDER BY create_time DESC
// todo: 类型分页查询 - dao.MaterialType(db) 条件 + 分页, 结果赋值 out
func (m *material) PageMaterialTypeByKey(ctx context.Context, in *dto.PageMaterialTypeByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 material_type 表并分页
	return nil
}

// GetMaterialTypeById 根据编号获取物资类型
// 对应 Java: MaterialServiceImpl.getMaterialTypeById
// todo: 标准 CRUD - dao.MaterialType(db).GetByID(ctx, types.BigInt(in.ID))
func (m *material) GetMaterialTypeById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.MaterialType(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// AddMaterialType 新增物资类型
// 对应 Java: MaterialServiceImpl.addMaterialType
// todo: 标准 CRUD - dao.MaterialType(db).InsertOne 写入 material_type 表
func (m *material) AddMaterialType(ctx context.Context, in *dto.AddMaterialTypeQuery, out *dto.EmptyResp) error {
	// todo: bean := do.NewMaterialType(); 填充 in; dao.MaterialType(db).InsertOne(ctx, bean)
	return nil
}

// EditMaterialType 编辑物资类型
// 对应 Java: MaterialServiceImpl.editMaterialType
// todo: 标准 CRUD - 按主键更新 material_type 表
func (m *material) EditMaterialType(ctx context.Context, in *dto.EditMaterialTypeQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<materialtype>.TypeName.Value(in.TypeName),
	}
	_, e := dao.MaterialType(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteMaterialType 删除物资类型
// 对应 Java: MaterialServiceImpl.deleteMaterialType
// todo: 标准 CRUD - dao.MaterialType(db).DeleteById(ctx, types.BigInt(in.ID))
func (m *material) DeleteMaterialType(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.MaterialType(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}
