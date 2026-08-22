package service

import (
	"context"
	"errors"

	"api/internal/constant"
	"api/internal/lib"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblcateringset"
	"api/internal/model/define/table/tbldishes"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/define/table/tblsetdishes"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type cateringset struct{}

var CateringSet = &cateringset{}

// PageCateringSetByKey 分页查询套餐
// 对应 Java: CateringSetServiceImpl.pageCateringSetByKey -> CateringSetFunc.listNotDelCateringSet
func (c *cateringset) PageCateringSetByKey(ctx context.Context, in *dto.PageCateringSetByKeyReq, out *[]dto.PageCateringSetByKeyResp) error {
	q := db.Table(tblcateringset.TableName).
		Where(tblcateringset.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblcateringset.DelFlag.Eq(constant.YesNoNo))
	if in.SetName != nil {
		q.And(tblcateringset.Name.Like(*in.SetName))
	}
	if in.DishesName != nil {
		q.And(tbldishes.Name.Like(*in.DishesName))
	}

	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblcateringset.Id,
			tblcateringset.Name,
			tblcateringset.MonthPrice,
		).
		Desc(tblcateringset.CreateTime).
		Select().
		Gets(ctx, out)
}

// GetCateringSetById 根据编号获取套餐（含菜品明细）
// 对应 Java: CateringSetServiceImpl.getCateringSetById -> CateringSetFunc.getCateringSetById
func (c *cateringset) GetCateringSetById(ctx context.Context, in *dto.IDReq, out *dto.GetCateringSetByIDResp) error {
	// 1. 查询套餐主表
	set, has, e := dao.CateringSet(db).GetByID(ctx, types.BigInt(*in.ID),
		tblcateringset.Id,
		tblcateringset.Name,
		tblcateringset.MonthPrice,
	)
	if e != nil {
		return e
	}
	if !has {
		return errors.New("套餐不存在")
	}
	out.ID = int64(set.Id)
	out.Name = set.Name.String()
	out.MonthPrice = set.MonthPrice

	// 2. 联表查询菜品明细（set_dishes LEFT JOIN dishes）
	var dishes []do.Dishes
	e = db.Table(tblsetdishes.TableName).
		LeftJoin(tblsetdishes.DishesId, tbldishes.Id).
		Where(
			tblsetdishes.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
			tblsetdishes.SetId.Eq(types.BigInt(*in.ID)),
			tbldishes.DelFlag.Eq(constant.YesNoNo),
		).
		Cols(
			tbldishes.Id,
			tbldishes.Name,
			tbldishes.Price,
		).
		Select().
		Gets(ctx, &dishes)
	if e != nil {
		return e
	}
	out.SetDishes = make([]dto.SetDishesResp, 0, len(dishes))
	for _, d := range dishes {
		out.SetDishes = append(out.SetDishes, dto.SetDishesResp{
			ID:    int64(d.Id),
			Name:  d.Name.String(),
			Price: d.Price,
		})
	}
	return nil
}

// checkNameRepeat 校验套餐名称是否重复（排除指定编号）
func (c *cateringset) checkNameRepeat(ctx context.Context, name string, excludeID *int64) (bool, error) {
	cond := []dialect.Condition{
		tblcateringset.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
		tblcateringset.Name.Eq(name),
		tblcateringset.DelFlag.Eq(constant.YesNoNo),
	}
	if excludeID != nil {
		cond = append(cond, tblcateringset.Id.NotEq(types.BigInt(*excludeID)))
	}
	return dao.CateringSet(db).Exists(ctx, cond...)
}

// saveBatchSetDishes 批量写入套餐菜品明细（编辑时先删后插）
func (c *cateringset) saveBatchSetDishes(ctx context.Context, setID int64, dishesIDList []int64, editFlag bool) error {
	if editFlag {
		_, e := dao.SetDishes(db).Delete(ctx,
			tblsetdishes.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
			tblsetdishes.SetId.Eq(types.BigInt(setID)))
		if e != nil {
			return e
		}
	}
	// 去重
	seen := make(map[int64]struct{}, len(dishesIDList))
	list := make([]*do.SetDishes, 0, len(dishesIDList))
	for _, dishesID := range dishesIDList {
		if _, ok := seen[dishesID]; ok {
			continue
		}
		seen[dishesID] = struct{}{}
		bean := do.NewSetDishes()
		bean.TenantId = types.BigInt(lib.TenantID(ctx))
		bean.SetId = types.BigInt(setID)
		bean.DishesId = types.BigInt(dishesID)
		list = append(list, bean)
	}
	if len(list) == 0 {
		return nil
	}
	_, e := dao.SetDishes(db).InsertBatch(ctx, list)
	return e
}

// AddCateringSet 新增套餐（含菜品明细）
// 对应 Java: CateringSetServiceImpl.addCateringSet
func (c *cateringset) AddCateringSet(ctx context.Context, in *dto.OperateCateringSetReq, out *dto.EmptyResp) error {
	// 校验套餐名称是否已存在
	repeat, e := c.checkNameRepeat(ctx, *in.Name, nil)
	if e != nil {
		return e
	}
	if repeat {
		return errors.New("套餐名称已存在")
	}
	// 新增主表
	bean := do.NewCateringSet()
	bean.TenantId = types.BigInt(lib.TenantID(ctx))
	bean.Name = types.String(*in.Name)
	bean.MonthPrice = types.Money(*in.MonthPrice)
	bean.DelFlag = types.Int8(constant.YesNoNo)
	_, e = dao.CateringSet(db).InsertOne(ctx, bean)
	if e != nil {
		return e
	}
	// 批量插入套餐菜品（自增 ID 已写入 bean.Id）
	return c.saveBatchSetDishes(ctx, int64(bean.Id), in.DishesIDList, false)
}

// EditCateringSet 编辑套餐（先删后插明细）
// 对应 Java: CateringSetServiceImpl.editCateringSet
func (c *cateringset) EditCateringSet(ctx context.Context, in *dto.OperateCateringSetReq, out *dto.EmptyResp) error {
	// 校验套餐名称是否已存在（排除自身）
	repeat, e := c.checkNameRepeat(ctx, *in.Name, in.ID)
	if e != nil {
		return e
	}
	if repeat {
		return errors.New("套餐名称已存在")
	}
	// 更新主表
	var sets = make([]dialect.Setter, 0, 2)
	sets = append(sets, tblcateringset.Name.Set(*in.Name))
	sets = append(sets, tblcateringset.MonthPrice.Set(*in.MonthPrice))
	_, e = dao.CateringSet(db).UpdateById(ctx, types.BigInt(*in.ID), sets...)
	if e != nil {
		return e
	}
	// 先删后插明细
	return c.saveBatchSetDishes(ctx, *in.ID, in.DishesIDList, true)
}

// DeleteCateringSet 删除套餐（逻辑删除）
// 对应 Java: CateringSetServiceImpl.deleteCateringSet -> ElderFunc.checkCatering
func (c *cateringset) DeleteCateringSet(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	// 判断是否有入住/退住审核中的老人选择该餐饮套餐
	used, e := dao.Elder(db).Exists(ctx,
		tblelder.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
		tblelder.CateringSetId.Eq(types.BigInt(*in.ID)),
		tblelder.CheckFlag.In(
			types.Int8(constant.CheckEnter),
			types.Int8(constant.CheckExitAudit),
		),
	)
	if e != nil {
		return e
	}
	if used {
		return errors.New("该套餐已有老人选择，无法删除")
	}
	// 逻辑删除
	_, e = dao.CateringSet(db).UpdateById(ctx, types.BigInt(*in.ID),
		tblcateringset.DelFlag.Set(constant.YesNoYes),
	)
	return e
}
