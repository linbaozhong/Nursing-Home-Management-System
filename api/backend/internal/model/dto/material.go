package dto

// ============ MaterialController 请求 ============

// @request
// PageMaterialByKeyQuery 分页查询物资请求
type PageMaterialByKeyQuery struct {
	PageNum        *int    `json:"page_num" valid:"required"`  // 页码
	PageSize       *int    `json:"page_size" valid:"required"` // 条数
	MaterialTypeID *int64  `json:"material_type_id"`           // 物资分类编号
	MaterialName   *string `json:"material_name"`              // 物资名称
}

// @request
// OperateMaterialQuery 操作物资请求
type OperateMaterialQuery struct {
	ID     *int64   `json:"id"`                       // id
	TypeID *int64   `json:"type_id" valid:"required"` // typeId
	Name   *string  `json:"name" valid:"required"`    // 物资名称
	Price  *float64 `json:"price" valid:"required"`   // 物资单价
}

// @request
// OperateMaterialTypeQuery 操作物资分类请求
type OperateMaterialTypeQuery struct {
	ID   *int64  `json:"id"`                    // id
	Name *string `json:"name" valid:"required"` // 物资分类名称
}

// @request
// AddMaterialQuery 新增物资请求
type AddMaterialQuery struct {
	ID     *int64   `json:"id"`                       // id
	TypeID *int64   `json:"type_id" valid:"required"` // 物资分类编号
	Name   *string  `json:"name" valid:"required"`    // 物资名称
	Price  *float64 `json:"price" valid:"required"`   // 物资单价
	Stock  *int     `json:"stock" valid:"required"`   // 库存数量
}

// @request
// EditMaterialQuery 编辑物资请求
type EditMaterialQuery struct {
	ID     *int64   `json:"id"`                       // id
	TypeID *int64   `json:"type_id" valid:"required"` // 物资分类编号
	Name   *string  `json:"name" valid:"required"`    // 物资名称
	Price  *float64 `json:"price" valid:"required"`   // 物资单价
	Stock  *int     `json:"stock" valid:"required"`   // 库存数量
}

// @request
// AddMaterialTypeQuery 新增物资分类请求
type AddMaterialTypeQuery struct {
	ID   *int64  `json:"id"`                    // id
	Name *string `json:"name" valid:"required"` // 物资分类名称
}

// @request
// EditMaterialTypeQuery 编辑物资分类请求
type EditMaterialTypeQuery struct {
	ID   *int64  `json:"id"`                    // id
	Name *string `json:"name" valid:"required"` // 物资分类名称
}

// @request
// PageMaterialTypeByKeyQuery 分页查询物资分类请求
type PageMaterialTypeByKeyQuery struct {
	PageNum  *int    `json:"page_num" valid:"required"`  // 页码
	PageSize *int    `json:"page_size" valid:"required"` // 条数
	Name     *string `json:"name"`                       // 物资分类名称
}

// ============ MaterialController 响应 ============

// @response
// PageMaterialByKeyVO 分页查询物资响应
type PageMaterialByKeyVO struct {
	ID       int64   `json:"id"`        // id
	Name     string  `json:"name"`      // 物资名称
	TypeName string  `json:"type_name"` // 物资分类
	Price    float64 `json:"price"`     // 物资单价
}

// @response
// OperateMaterialVO 操作物资响应（继承 OperateMaterialQuery）
type OperateMaterialVO struct {
	OperateMaterialQuery
}

// @response
// PageMaterialTypeVO 分页查询物资分类响应
type PageMaterialTypeVO struct {
	ID   int64  `json:"id"`   // 物资分类编号
	Name string `json:"name"` // 物资分类名称
}
