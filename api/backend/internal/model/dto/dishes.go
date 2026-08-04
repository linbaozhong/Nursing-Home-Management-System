package dto

// ============ DishesController 请求 ============

// @request
// PageDishesByKeyQuery 分页查询菜品请求
type PageDishesByKeyQuery struct {
	PageNum    *int    `json:"page_num" valid:"required"`  // 页码
	PageSize   *int    `json:"page_size" valid:"required"` // 条数
	TypeID     *int64  `json:"type_id"`                    // 菜品分类编号
	DishesName *string `json:"dishes_name"`                // 菜品名称
}

// @request
// OperateDishesQuery 操作菜品请求
type OperateDishesQuery struct {
	ID     *int64   `json:"id"`                       // id
	TypeID *int64   `json:"type_id" valid:"required"` // 菜品分类编号
	Name   *string  `json:"name" valid:"required"`    // 菜品名称
	Price  *float64 `json:"price" valid:"required"`   // 菜品价格
}

// @request
// OperateDishesTypeQuery 操作菜品分类请求
type OperateDishesTypeQuery struct {
	ID   *int64  `json:"id"`                    // id
	Name *string `json:"name" valid:"required"` // 菜品分类名称
}

// ============ DishesController 响应 ============

// @response
// PageDishesByKeyVO 分页查询菜品响应
type PageDishesByKeyVO struct {
	ID         int64   `json:"id"`          // id
	TypeName   string  `json:"type_name"`   // 类别名称
	DishesName string  `json:"dishes_name"` // 菜品名称
	Price      float64 `json:"price"`       // 价格
}

// @response
// OperateDishesVO 操作菜品响应（继承 OperateDishesQuery）
type OperateDishesVO struct {
	OperateDishesQuery
}

// @request
// AddDishesQuery 新增菜品请求
type AddDishesQuery struct {
	ID     *int64   `json:"id"`                       // id
	TypeID *int64   `json:"type_id" valid:"required"` // 菜品分类编号
	Name   *string  `json:"name" valid:"required"`    // 菜品名称
	Price  *float64 `json:"price" valid:"required"`   // 菜品价格
}

// @request
// EditDishesQuery 编辑菜品请求
type EditDishesQuery struct {
	ID     *int64   `json:"id"`                       // id
	TypeID *int64   `json:"type_id" valid:"required"` // 菜品分类编号
	Name   *string  `json:"name" valid:"required"`    // 菜品名称
	Price  *float64 `json:"price" valid:"required"`   // 菜品价格
}

// @request
// PageDishesTypeByKeyQuery 分页查询菜品分类请求
type PageDishesTypeByKeyQuery struct {
	PageNum  *int    `json:"page_num" valid:"required"`  // 页码
	PageSize *int    `json:"page_size" valid:"required"` // 条数
	Name     *string `json:"name"`                       // 菜品分类名称
}

// @request
// AddDishesTypeQuery 新增菜品分类请求
type AddDishesTypeQuery struct {
	ID   *int64  `json:"id"`                    // id
	Name *string `json:"name" valid:"required"` // 菜品分类名称
}

// @request
// EditDishesTypeQuery 编辑菜品分类请求
type EditDishesTypeQuery struct {
	ID   *int64  `json:"id"`                    // id
	Name *string `json:"name" valid:"required"` // 菜品分类名称
}
