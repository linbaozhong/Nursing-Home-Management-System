package service

import (
	"context"
	"errors"

	"api/internal/constant"
	"api/internal/lib"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tbldishes"
	"api/internal/model/define/table/tbldishestype"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type dishes struct{}

var Dishes = &dishes{}

// PageDishesByKey 分页查询菜品（联表菜品类别）
// 对应 Java: DishesServiceImpl.pageDishesByKey -> DishesFunc.listDishes
func (d *dishes) PageDishesByKey(ctx context.Context, in *dto.PageDishesByKeyReq, out *[]dto.PageDishesByKeyResp) error {
	q := db.Table(tbldishes.TableName).
		LeftJoin(tbldishes.TypeId, tbldishestype.Id).
		Where(tbldishes.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tbldishes.State.NotEq(types.Int8(constant.StateDeleted)))
	if in.TypeID != nil {
		q.And(tbldishes.TypeId.Eq(types.BigInt(*in.TypeID)))
	}
	if in.DishesName != nil && *in.DishesName != "" {
		q.And(tbldishes.Name.Like(*in.DishesName))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tbldishes.Id.AsName("id"),
			tbldishestype.Name.AsName("type_name"),
			tbldishes.Name.AsName("dishes_name"),
			tbldishes.Price.AsName("price"),
		).
		Desc(tbldishes.CreateTime).
		Select().
		Gets(ctx, out)
}

// GetDishesById 根据编号获取菜品
func (d *dishes) GetDishesById(ctx context.Context, in *dto.IDReq, out *dto.OperateDishesResp) error {
	obj, has, e := dao.Dishes(db).GetByID(ctx, types.BigInt(*in.ID),
		tbldishes.Id,
		tbldishes.TypeId,
		tbldishes.Name,
		tbldishes.Price,
	)
	if e != nil {
		return e
	}
	if !has {
		return errors.New("菜品不存在")
	}
	*out.ID = int64(obj.Id)
	*out.TypeID = int64(obj.TypeId)
	*out.Name = obj.Name.String()
	*out.Price = obj.Price
	return nil
}

// AddDishes 新增菜品（校验名称+类别唯一）
func (d *dishes) AddDishes(ctx context.Context, in *dto.OperateDishesReq, out *dto.EmptyResp) error {
	repeat, e := dao.Dishes(db).Exists(ctx,
		tbldishes.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
		tbldishes.Name.Eq(*in.Name),
		tbldishes.TypeId.Eq(types.BigInt(*in.TypeID)),
		tbldishes.State.NotEq(types.Int8(constant.StateDeleted)),
	)
	if e != nil {
		return e
	}
	if repeat {
		return errors.New("菜品已存在")
	}
	bean := do.NewDishes()
	bean.TenantId = types.BigInt(lib.TenantID(ctx))
	bean.TypeId = types.BigInt(*in.TypeID)
	bean.Name = types.String(*in.Name)
	bean.Price = types.Money(*in.Price)
	_, e = dao.Dishes(db).InsertOne(ctx, bean)
	return e
}

// EditDishes 编辑菜品（校验同名不同编号）
func (d *dishes) EditDishes(ctx context.Context, in *dto.OperateDishesReq, out *dto.EmptyResp) error {
	repeat, e := dao.Dishes(db).Exists(ctx,
		tbldishes.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
		tbldishes.Name.Eq(*in.Name),
		tbldishes.TypeId.Eq(types.BigInt(*in.TypeID)),
		tbldishes.State.NotEq(types.Int8(constant.StateDeleted)),
		tbldishes.Id.NotEq(types.BigInt(*in.ID)),
	)
	if e != nil {
		return e
	}
	if repeat {
		return errors.New("菜品已存在")
	}
	var sets = make([]dialect.Setter, 0, 3)
	sets = append(sets, tbldishes.TypeId.Set(*in.TypeID))
	sets = append(sets, tbldishes.Name.Set(*in.Name))
	sets = append(sets, tbldishes.Price.Set(*in.Price))
	_, e = dao.Dishes(db).UpdateById(ctx, types.BigInt(*in.ID), sets...)
	return e
}

// DeleteDishes 删除菜品（逻辑删除）
func (d *dishes) DeleteDishes(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.Dishes(db).UpdateById(ctx, types.BigInt(*in.ID),
		tbldishes.State.Set(constant.StateDeleted),
	)
	return e
}

// PageDishesTypeByKey 分页查询菜品类别
// 对应 Java: DishesServiceImpl.pageDishesTypeByKey -> DishesTypeFunc.listDishesType
func (d *dishes) PageDishesTypeByKey(ctx context.Context, in *dto.PageDishesTypeByKeyReq, out *[]dto.DropDown) error {
	q := db.Table(tbldishestype.TableName).
		Where(tbldishestype.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tbldishestype.State.NotEq(types.Int8(constant.StateDeleted)))
	if in.Name != nil && *in.Name != "" {
		q.And(tbldishestype.Name.Like(*in.Name))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(tbldishestype.Id, tbldishestype.Name).
		Desc(tbldishestype.CreateTime).
		Select().
		Gets(ctx, out)
}

// GetDishesTypeById 根据编号获取菜品类别
func (d *dishes) GetDishesTypeById(ctx context.Context, in *dto.IDReq, out *dto.DropDown) error {
	obj, has, e := dao.DishesType(db).GetByID(ctx, types.BigInt(*in.ID),
		tbldishestype.Id,
		tbldishestype.Name,
	)
	if e != nil {
		return e
	}
	if !has {
		return errors.New("菜品类别不存在")
	}
	out.ID = types.BigInt(obj.Id)
	out.Name = obj.Name.String()
	return nil
}

// AddDishesType 新增菜品类别（校验名称唯一 + 类别总数上限）
func (d *dishes) AddDishesType(ctx context.Context, in *dto.OperateDishesTypeReq, out *dto.EmptyResp) error {
	total, e := dao.DishesType(db).Count(ctx,
		tbldishestype.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
		tbldishestype.State.NotEq(types.Int8(constant.StateDeleted)),
	)
	if e != nil {
		return e
	}
	if total >= 10 {
		return errors.New("菜品类别已经达到上限")
	}
	repeat, e := dao.DishesType(db).Exists(ctx,
		tbldishestype.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
		tbldishestype.Name.Eq(*in.Name),
		tbldishestype.State.NotEq(types.Int8(constant.StateDeleted)),
	)
	if e != nil {
		return e
	}
	if repeat {
		return errors.New("菜品类别已存在")
	}
	bean := do.NewDishesType()
	bean.TenantId = types.BigInt(lib.TenantID(ctx))
	bean.Name = types.String(*in.Name)
	_, e = dao.DishesType(db).InsertOne(ctx, bean)
	return e
}

// EditDishesType 编辑菜品类别（校验同名不同编号）
func (d *dishes) EditDishesType(ctx context.Context, in *dto.OperateDishesTypeReq, out *dto.EmptyResp) error {
	repeat, e := dao.DishesType(db).Exists(ctx,
		tbldishestype.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
		tbldishestype.Name.Eq(*in.Name),
		tbldishestype.State.NotEq(types.Int8(constant.StateDeleted)),
		tbldishestype.Id.NotEq(types.BigInt(*in.ID)),
	)
	if e != nil {
		return e
	}
	if repeat {
		return errors.New("菜品类别已存在")
	}
	_, e = dao.DishesType(db).UpdateById(ctx, types.BigInt(*in.ID),
		tbldishestype.Name.Set(*in.Name),
	)
	return e
}

// DeleteDishesType 删除菜品类别（逻辑删除，需无子菜品）
func (d *dishes) DeleteDishesType(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	hasChild, e := dao.Dishes(db).Exists(ctx,
		tbldishes.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
		tbldishes.TypeId.Eq(types.BigInt(*in.ID)),
		tbldishes.State.NotEq(types.Int8(constant.StateDeleted)),
	)
	if e != nil {
		return e
	}
	if hasChild {
		return errors.New("该类别存在菜品，无法删除")
	}
	_, e = dao.DishesType(db).UpdateById(ctx, types.BigInt(*in.ID),
		tbldishestype.State.Set(constant.StateDeleted),
	)
	return e
}
