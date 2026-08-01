package service

import (
	"context"

	"api/internal/model/define/dao"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type label struct{}

var Label = &label{}

// PageLabelByKey 分页查询标签
// 对应 Java: LabelServiceImpl.pageLabelByKey -> LabelMapper.listLabelByKey
// SQL: SELECT l.*, lt.type_name FROM label l LEFT JOIN label_type lt ON lt.id = l.label_type_id
//
//	WHERE (lt.type_name LIKE %key% OR l.label_name LIKE %key%) [可选]
//	ORDER BY l.create_time DESC; 再由 PageUtil 内存分页。
//
// todo: 1) in.Key 非空 -> (tbl<labeltype>.TypeName.Like(in.Key) OR tbl<label>.LabelName.Like(in.Key))
//
//	2) DB 分页: Count + List(联表 LeftJoin label_type)
//	3) 组装含类型名的 VO 并赋值 out
func (l *label) PageLabelByKey(ctx context.Context, in *dto.PageLabelByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 实现联表分页查询
	return nil
}

// GetLabelById 根据编号获取标签
// 对应 Java: LabelServiceImpl.getLabelById -> labelMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.Label(db).GetByID(ctx, types.BigInt(in.ID))
func (l *label) GetLabelById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.Label(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// AddLabel 新增标签
// 对应 Java: LabelServiceImpl.addLabel -> labelMapper.insertSelective
// todo: 标准 CRUD - dao.Label(db).InsertOne 写入 label 表(含 labelName/labelTypeId 等)
func (l *label) AddLabel(ctx context.Context, in *dto.AddLabelQuery, out *dto.EmptyResp) error {
	// todo: bean := do.NewLabel(); 填充 in; dao.Label(db).InsertOne(ctx, bean)
	return nil
}

// EditLabel 编辑标签
// 对应 Java: LabelServiceImpl.editLabel -> labelMapper.updateByPrimaryKeySelective
// todo: 标准 CRUD - 按主键更新 label 表
func (l *label) EditLabel(ctx context.Context, in *dto.EditLabelQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<label>.LabelName.Value(in.LabelName),
	}
	_, e := dao.Label(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteLabel 删除标签
// 对应 Java: LabelServiceImpl.deleteLabel -> labelMapper.deleteByPrimaryKey
// todo: 标准 CRUD - dao.Label(db).DeleteById(ctx, types.BigInt(in.ID))
func (l *label) DeleteLabel(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.Label(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}

// PageLabelTypeByKey 分页查询标签类型
// 对应 Java: LabelServiceImpl.pageLabelTypeByKey -> LabelTypeMapper.listLabelTypeByKey
// SQL: SELECT * FROM label_type WHERE (type_name LIKE %key%) [可选] ORDER BY create_time DESC
// todo: 类型分页查询 - dao.LabelType(db) 条件 + 分页, 结果赋值 out
func (l *label) PageLabelTypeByKey(ctx context.Context, in *dto.PageLabelTypeByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 label_type 表并分页
	return nil
}

// GetLabelTypeById 根据编号获取标签类型
// 对应 Java: LabelServiceImpl.getLabelTypeById -> labelTypeMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.LabelType(db).GetByID(ctx, types.BigInt(in.ID))
func (l *label) GetLabelTypeById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.LabelType(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// AddLabelType 新增标签类型
// 对应 Java: LabelServiceImpl.addLabelType -> labelTypeMapper.insertSelective
// todo: 标准 CRUD - dao.LabelType(db).InsertOne 写入 label_type 表
func (l *label) AddLabelType(ctx context.Context, in *dto.AddLabelTypeQuery, out *dto.EmptyResp) error {
	// todo: bean := do.NewLabelType(); 填充 in; dao.LabelType(db).InsertOne(ctx, bean)
	return nil
}

// EditLabelType 编辑标签类型
// 对应 Java: LabelServiceImpl.editLabelType -> labelTypeMapper.updateByPrimaryKeySelective
// todo: 标准 CRUD - 按主键更新 label_type 表
func (l *label) EditLabelType(ctx context.Context, in *dto.EditLabelTypeQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<labeltype>.TypeName.Value(in.TypeName),
	}
	_, e := dao.LabelType(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteLabelType 删除标签类型
// 对应 Java: LabelServiceImpl.deleteLabelType -> labelTypeMapper.deleteByPrimaryKey
// todo: 标准 CRUD - dao.LabelType(db).DeleteById(ctx, types.BigInt(in.ID))
func (l *label) DeleteLabelType(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.LabelType(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}
