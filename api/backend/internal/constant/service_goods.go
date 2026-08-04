package constant

import "github.com/linbaozhong/gentity/pkg/types"

// 服务
var (
	ErrServiceTypeRepeat = types.NewError(500, "该服务类型已存在")
	ErrServiceTypeOut    = types.NewError(500, "服务类型总数超过限制")
	ErrServiceNotNull    = types.NewError(500, "该服务类型存在服务，删除失败")
	ErrServiceRepeat     = types.NewError(500, "该服务已存在")
	ErrServiceOut        = types.NewError(500, "该类型服务总数超过限制")
)

// 护理等级
var (
	ErrNurseGradeRepeat   = types.NewError(500, "该护理等级已存在")
	ErrNurseGradeSelected = types.NewError(500, "该护理等级已被选择，删除失败")
)

// 菜品
var (
	ErrDishesTypeRepeat = types.NewError(500, "该菜品分类已存在")
	ErrDishesTypeOut    = types.NewError(500, "菜品分类总数超过限制")
	ErrDishesNotNull    = types.NewError(500, "该菜品分类存在菜品，删除失败")
	ErrDishesRepeat     = types.NewError(500, "该菜品已存在")
)

// 套餐
var (
	ErrSetRepeat   = types.NewError(500, "该套餐已存在")
	ErrSetSelected = types.NewError(500, "该套餐已被选择，删除失败")
)

// 订单
var (
	ErrOrderSuccess = types.NewError(500, "该订单已完成")
)

// 物资
var (
	ErrMaterialTypeRepeat = types.NewError(500, "该物资分类已存在")
	ErrMaterialTypeOut    = types.NewError(500, "物资分类总数超过限制")
	ErrMaterialNotNull    = types.NewError(500, "该物资分类存在物资，删除失败")
	ErrMaterialRepeat     = types.NewError(500, "该物资已存在")
)

// 仓库
var (
	ErrWarehouseRepeat  = types.NewError(500, "该仓库已存在")
	ErrWarehouseNotNull = types.NewError(500, "该仓库存有物资，删除失败")
)
