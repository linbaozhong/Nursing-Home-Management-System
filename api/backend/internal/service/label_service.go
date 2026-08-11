package service

import (
	"context"
	"errors"

	"api/internal/constant"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tbllabel"
	"api/internal/model/define/table/tbllabeltype"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/types"
)

// labelTotalLimit 标签 / 标签类型数量上限（对应 Java Constant.TOTAL_LIMIT）
const labelTotalLimit = 10

type label struct{}

var Label = &label{}

// PageLabelByKey 分页查询标签（关联标签类型名称）
// 对应 Java: LabelServiceImpl.pageLabelByKey -> LabelFunc.listLabelByKey
func (l *label) PageLabelByKey(ctx context.Context, in *dto.PageLabelByKeyQuery, out *[]dto.PageLabelByKeyVO) error {
	q := db.Table(tbllabel.TableName).
		LeftJoin(tbllabel.TypeId, tbllabeltype.Id).
		Where(tbllabel.DelFlag.Eq(constant.YesNoNo))
	if in.Key != nil && *in.Key != "" {
		q.And(tbllabel.Name.Like(*in.Key))
		q.Or(tbllabeltype.Name.Like(*in.Key))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tbllabel.Id.AsName("id"),
			tbllabel.Name.AsName("name"),
			tbllabel.Color.AsName("color"),
			tbllabel.TypeId.AsName("type_id"),
			tbllabeltype.Name.AsName("type_name"),
		).
		Desc(tbllabel.CreateTime).
		Select().
		Gets(ctx, out)
}

// GetLabelById 根据编号获取标签（编辑回显）
// 对应 Java: LabelServiceImpl.getLabelById
func (l *label) GetLabelById(ctx context.Context, in *dto.IDReq, out *dto.OperateLabelVO) error {
	obj, has, e := dao.Label(db).Get(ctx, ace.Where(tbllabel.Id.Eq(types.BigInt(*in.ID))))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("标签不存在")
	}
	out.ID = int64Ptr(int64(obj.Id))
	out.TypeID = int64Ptr(int64(obj.TypeId))
	out.Name = strPtr(obj.Name.String())
	out.Color = strPtr(obj.Color.String())
	return nil
}

// AddLabel 新增标签（校验类型存在、名称不重复、总数不超上限）
// 对应 Java: LabelServiceImpl.addLabel -> LabelFunc.checkLabelLimit / checkLabelType
func (l *label) AddLabel(ctx context.Context, in *dto.AddLabelQuery, out *dto.EmptyResp) error {
	// 校验标签类型存在
	_, has, e := dao.LabelType(db).Get(ctx, ace.Where(tbllabeltype.Id.Eq(types.BigInt(*in.TypeID))))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("标签类型不存在")
	}
	// 校验同一类型下名称不重复
	repeat, e := dao.Label(db).Exists(ctx,
		tbllabel.TypeId.Eq(types.BigInt(*in.TypeID)),
		tbllabel.Name.Eq(*in.Name),
		tbllabel.DelFlag.Eq(types.Int8(constant.YesNoNo)),
	)
	if e != nil {
		return e
	}
	if repeat {
		return errors.New("标签名称已存在")
	}
	// 校验总数不超上限
	total, e := dao.Label(db).Count(ctx, tbllabel.DelFlag.Eq(constant.YesNoNo))
	if e != nil {
		return e
	}
	if int(total) >= labelTotalLimit {
		return errors.New("标签数量已达上限")
	}
	bean := do.NewLabel()
	bean.Name = types.String(*in.Name)
	bean.Color = types.String(*in.Color)
	bean.TypeId = types.BigInt(*in.TypeID)
	_, e = dao.Label(db).InsertOne(ctx, bean)
	return e
}

// EditLabel 修改标签（校验类型存在、名称不重复排除自身、总数不超上限）
// 对应 Java: LabelServiceImpl.updateLabel
func (l *label) EditLabel(ctx context.Context, in *dto.EditLabelQuery, out *dto.EmptyResp) error {
	_, has, e := dao.LabelType(db).Get(ctx, ace.Where(tbllabeltype.Id.Eq(types.BigInt(*in.TypeID))))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("标签类型不存在")
	}
	repeat, e := dao.Label(db).Exists(ctx,
		tbllabel.TypeId.Eq(types.BigInt(*in.TypeID)),
		tbllabel.Name.Eq(*in.Name),
		tbllabel.DelFlag.Eq(types.Int8(constant.YesNoNo)),
		tbllabel.Id.NotEq(types.BigInt(*in.ID)),
	)
	if e != nil {
		return e
	}
	if repeat {
		return errors.New("标签名称已存在")
	}
	total, e := dao.Label(db).Count(ctx, tbllabel.DelFlag.Eq(constant.YesNoNo))
	if e != nil {
		return e
	}
	if int(total) >= labelTotalLimit {
		return errors.New("标签数量已达上限")
	}
	bean := do.NewLabel()
	bean.Id = types.BigInt(*in.ID)
	bean.Name = types.String(*in.Name)
	bean.Color = types.String(*in.Color)
	bean.TypeId = types.BigInt(*in.TypeID)
	_, e = dao.Label(db).UpdateOne(ctx, bean)
	return e
}

// DeleteLabel 删除标签
// 对应 Java: LabelServiceImpl.deleteLabel
func (l *label) DeleteLabel(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.Label(db).Delete(ctx, tbllabel.Id.Eq(types.BigInt(*in.ID)))
	return e
}

// PageLabelType 分页查询标签类型
// 对应 Java: LabelServiceImpl.pageLabelTypeByKey -> LabelTypeFunc.listLabelTypeByKey
func (l *label) PageLabelType(ctx context.Context, in *dto.PageLabelTypeByKeyQuery, out *[]dto.PageLabelTypeVO) error {
	q := db.Table(tbllabeltype.TableName).
		Where(tbllabeltype.DelFlag.Eq(constant.YesNoNo))
	if in.Name != nil && *in.Name != "" {
		q.And(tbllabeltype.Name.Like(*in.Name))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tbllabeltype.Id.AsName("id"),
			tbllabeltype.Name.AsName("name"),
			tbllabeltype.CreateTime.AsName("create_time"),
			tbllabeltype.UpdateTime.AsName("update_time"),
		).
		Desc(tbllabeltype.CreateTime).
		Select().
		Gets(ctx, out)
}

// GetLabelTypeById 根据编号获取标签类型（编辑回显）
// 对应 Java: LabelServiceImpl.getLabelTypeById
func (l *label) GetLabelTypeById(ctx context.Context, in *dto.IDReq, out *dto.OperateLabelTypeVO) error {
	obj, has, e := dao.LabelType(db).Get(ctx, ace.Where(tbllabeltype.Id.Eq(types.BigInt(*in.ID))))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("标签类型不存在")
	}
	out.ID = int64Ptr(int64(obj.Id))
	out.Name = strPtr(obj.Name.String())
	return nil
}

// AddLabelType 新增标签类型（校验名称不重复、类型总数不超上限）
// 对应 Java: LabelServiceImpl.addLabelType -> LabelTypeFunc.checkLabelTypeLimit
func (l *label) AddLabelType(ctx context.Context, in *dto.AddLabelTypeQuery, out *dto.EmptyResp) error {
	repeat, e := dao.LabelType(db).Exists(ctx,
		tbllabeltype.Name.Eq(*in.Name),
		tbllabeltype.DelFlag.Eq(types.Int8(constant.YesNoNo)),
	)
	if e != nil {
		return e
	}
	if repeat {
		return errors.New("标签类型名称已存在")
	}
	total, e := dao.LabelType(db).Count(ctx, tbllabeltype.DelFlag.Eq(constant.YesNoNo))
	if e != nil {
		return e
	}
	if int(total) >= labelTotalLimit {
		return errors.New("标签类型数量已达上限")
	}
	bean := do.NewLabelType()
	bean.Name = types.String(*in.Name)
	_, e = dao.LabelType(db).InsertOne(ctx, bean)
	return e
}

// EditLabelType 修改标签类型（校验名称不重复排除自身、类型总数不超上限）
// 对应 Java: LabelServiceImpl.updateLabelType
func (l *label) EditLabelType(ctx context.Context, in *dto.EditLabelTypeQuery, out *dto.EmptyResp) error {
	repeat, e := dao.LabelType(db).Exists(ctx,
		tbllabeltype.Name.Eq(*in.Name),
		tbllabeltype.DelFlag.Eq(types.Int8(constant.YesNoNo)),
		tbllabeltype.Id.NotEq(types.BigInt(*in.ID)),
	)
	if e != nil {
		return e
	}
	if repeat {
		return errors.New("标签类型名称已存在")
	}
	total, e := dao.LabelType(db).Count(ctx, tbllabeltype.DelFlag.Eq(constant.YesNoNo))
	if e != nil {
		return e
	}
	if int(total) >= labelTotalLimit {
		return errors.New("标签类型数量已达上限")
	}
	bean := do.NewLabelType()
	bean.Id = types.BigInt(*in.ID)
	bean.Name = types.String(*in.Name)
	_, e = dao.LabelType(db).UpdateOne(ctx, bean)
	return e
}

// DeleteLabelType 删除标签类型（存在子标签则不允许删除）
// 对应 Java: LabelServiceImpl.deleteLabelType -> LabelFunc.checkLabelByTypeId
func (l *label) DeleteLabelType(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	hasChild, e := dao.Label(db).Exists(ctx,
		tbllabel.TypeId.Eq(types.BigInt(*in.ID)),
		tbllabel.DelFlag.Eq(types.Int8(constant.YesNoNo)),
	)
	if e != nil {
		return e
	}
	if hasChild {
		return errors.New("该标签类型下存在标签，无法删除")
	}
	_, e = dao.LabelType(db).Delete(ctx, tbllabeltype.Id.Eq(types.BigInt(*in.ID)))
	return e
}
