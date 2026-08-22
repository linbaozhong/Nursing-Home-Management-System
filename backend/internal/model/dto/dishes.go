package dto

import "github.com/linbaozhong/gentity/pkg/types"

// ============ DishesController 请求 ============

// PageDishesByKeyReq 分页查询菜品请求
// @request
type PageDishesByKeyReq struct {
	PageNum    *int    `json:"page_num" valid:"required"`  // 页码
	PageSize   *int    `json:"page_size" valid:"required"` // 条数
	TypeID     *int64  `json:"type_id"`                    // 菜品分类编号
	DishesName *string `json:"dishes_name"`                // 菜品名称
}

// OperateDishesReq 操作菜品请求
// @request
type OperateDishesReq struct {
	ID     *int64       `json:"id"`                       // id
	TypeID *int64       `json:"type_id" valid:"required"` // 菜品分类编号
	Name   *string      `json:"name" valid:"required"`    // 菜品名称
	Price  *types.Money `json:"price" valid:"required"`   // 菜品价格
}

// OperateDishesTypeReq 操作菜品分类请求
// @request
type OperateDishesTypeReq struct {
	ID   *int64  `json:"id"`                    // id
	Name *string `json:"name" valid:"required"` // 菜品分类名称
}

// ============ DishesController 响应 ============

// PageDishesByKeyResp 分页查询菜品响应
// @response
type PageDishesByKeyResp struct {
	ID         int64       `json:"id"`          // id
	TypeName   string      `json:"type_name"`   // 类别名称
	DishesName string      `json:"dishes_name"` // 菜品名称
	Price      types.Money `json:"price"`       // 价格
}

// OperateDishesResp 操作菜品响应（继承 OperateDishesReq）
// @response
type OperateDishesResp struct {
	OperateDishesReq
}

// AddDishesReq 新增菜品请求
// @request
type AddDishesReq struct {
	ID     *int64       `json:"id"`                       // id
	TypeID *int64       `json:"type_id" valid:"required"` // 菜品分类编号
	Name   *string      `json:"name" valid:"required"`    // 菜品名称
	Price  *types.Money `json:"price" valid:"required"`   // 菜品价格
}

// EditDishesReq 编辑菜品请求
// @request
type EditDishesReq struct {
	ID     *int64       `json:"id"`                       // id
	TypeID *int64       `json:"type_id" valid:"required"` // 菜品分类编号
	Name   *string      `json:"name" valid:"required"`    // 菜品名称
	Price  *types.Money `json:"price" valid:"required"`   // 菜品价格
}

// PageDishesTypeByKeyReq 分页查询菜品分类请求
// @request
type PageDishesTypeByKeyReq struct {
	PageNum  *int    `json:"page_num" valid:"required"`  // 页码
	PageSize *int    `json:"page_size" valid:"required"` // 条数
	Name     *string `json:"name"`                       // 菜品分类名称
}

// AddDishesTypeReq 新增菜品分类请求
// @request
type AddDishesTypeReq struct {
	ID   *int64  `json:"id"`                    // id
	Name *string `json:"name" valid:"required"` // 菜品分类名称
}

// EditDishesTypeReq 编辑菜品分类请求
// @request
type EditDishesTypeReq struct {
	ID   *int64  `json:"id"`                    // id
	Name *string `json:"name" valid:"required"` // 菜品分类名称
}
