package service

import (
	"context"

	"api/internal/model/define/dao"
	"api/internal/model/define/table/tbldishes"
	"api/internal/model/define/table/tbldishestype"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type dishes struct{}

var Dishes = &dishes{}

// PageDishesByKey 分页查询菜品（联表 dishes_type 获取类型名）
// 对应 Java: DishesServiceImpl.pageDishesByKey -> DishesMapper.listDishesByKey
// SQL: SELECT d.*, dt.type_name FROM dishes d
//
//	LEFT JOIN dishes_type dt ON dt.id = d.dishes_type_id
//	WHERE (dt.type_name LIKE %key% OR d.dishes_name LIKE %key%) [可选]
//	ORDER BY d.create_time DESC; 再由 PageUtil 内存分页。
//
// todo: 1) in.Key 非空 -> (tbl<dishestype>.TypeName.Like(in.Key) OR tbl<dishes>.DishesName.Like(in.Key))
//
//	2) DB 分页: Count + List(联表 LeftJoin)
//	3) 组装含类型名的 VO 并赋值 out
func (d *dishes) PageDishesByKey(ctx context.Context, in *dto.PageDishesByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 实现联表分页查询
	return nil
}

// GetDishesById 根据编号获取菜品
// 对应 Java: DishesServiceImpl.getDishesById -> dishesMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.Dishes(db).GetByID(ctx, types.BigInt(in.ID))
func (d *dishes) GetDishesById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.Dishes(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// AddDishes 新增菜品
// 对应 Java: DishesServiceImpl.addDishes -> dishesMapper.insertSelective
// todo: 标准 CRUD - dao.Dishes(db).InsertOne 写入 dishes 表(含 dishesName/dishesTypeId/price 等)
func (d *dishes) AddDishes(ctx context.Context, in *dto.AddDishesQuery, out *dto.EmptyResp) error {
	// todo: bean := do.NewDishes(); 填充 in; dao.Dishes(db).InsertOne(ctx, bean)
	return nil
}

// EditDishes 编辑菜品
// 对应 Java: DishesServiceImpl.editDishes -> dishesMapper.updateByPrimaryKeySelective
// todo: 标准 CRUD - 按主键更新 dishes 表(tbl<dishes>.X.Value(in.X))
func (d *dishes) EditDishes(ctx context.Context, in *dto.EditDishesQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<dishes>.DishesName.Value(in.DishesName),
	}
	_, e := dao.Dishes(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteDishes 删除菜品
// 对应 Java: DishesServiceImpl.deleteDishes -> dishesMapper.deleteByPrimaryKey
// todo: 标准 CRUD - dao.Dishes(db).DeleteById(ctx, types.BigInt(in.ID))
func (d *dishes) DeleteDishes(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.Dishes(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}

// PageDishesTypeByKey 分页查询菜品类型
// 对应 Java: DishesServiceImpl.pageDishesTypeByKey -> DishesTypeMapper.listDishesTypeByKey
// SQL: SELECT * FROM dishes_type WHERE (type_name LIKE %key%) [可选] ORDER BY create_time DESC
// todo: 类型分页查询 - dao.DishesType(db) 条件 + 分页, 结果赋值 out
func (d *dishes) PageDishesTypeByKey(ctx context.Context, in *dto.PageDishesTypeByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 dishes_type 表并分页
	return nil
}

// GetDishesTypeById 根据编号获取菜品类型
// 对应 Java: DishesServiceImpl.getDishesTypeById
// todo: 标准 CRUD - dao.DishesType(db).GetByID(ctx, types.BigInt(in.ID))
func (d *dishes) GetDishesTypeById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.DishesType(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// AddDishesType 新增菜品类型
// 对应 Java: DishesServiceImpl.addDishesType
// todo: 标准 CRUD - dao.DishesType(db).InsertOne 写入 dishes_type 表
func (d *dishes) AddDishesType(ctx context.Context, in *dto.AddDishesTypeQuery, out *dto.EmptyResp) error {
	// todo: bean := do.NewDishesType(); 填充 in; dao.DishesType(db).InsertOne(ctx, bean)
	return nil
}

// EditDishesType 编辑菜品类型
// 对应 Java: DishesServiceImpl.editDishesType
// todo: 标准 CRUD - 按主键更新 dishes_type 表
func (d *dishes) EditDishesType(ctx context.Context, in *dto.EditDishesTypeQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<dishestype>.TypeName.Value(in.TypeName),
	}
	_, e := dao.DishesType(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteDishesType 删除菜品类型
// 对应 Java: DishesServiceImpl.deleteDishesType
// todo: 标准 CRUD - dao.DishesType(db).DeleteById(ctx, types.BigInt(in.ID))
func (d *dishes) DeleteDishesType(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.DishesType(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}
