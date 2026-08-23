package dto

import "github.com/linbaozhong/gentity/pkg/types"

// ============ MaterialController 请求 ============

// PageMaterialByKeyReq 分页查询物资请求
// @request
type PageMaterialByKeyReq struct {
	PageNum        *int    `json:"page_num" valid:"required"`  // 页码
	PageSize       *int    `json:"page_size" valid:"required"` // 条数
	MaterialTypeID *int64  `json:"material_type_id"`           // 物资分类编号
	MaterialName   *string `json:"material_name"`              // 物资名称
}

// OperateMaterialReq 操作物资请求
// @request
type OperateMaterialReq struct {
	ID     *int64       `json:"id"`                       // id
	TypeID *int64       `json:"type_id" valid:"required"` // typeId
	Name   *string      `json:"name" valid:"required"`    // 物资名称
	Price  *types.Money `json:"price" valid:"required"`   // 物资单价
}

// OperateMaterialTypeReq 操作物资分类请求
// @request
type OperateMaterialTypeReq struct {
	ID   *int64  `json:"id"`                    // id
	Name *string `json:"name" valid:"required"` // 物资分类名称
}

// AddMaterialReq 新增物资请求
// @request
type AddMaterialReq struct {
	ID     *int64       `json:"id"`                       // id
	TypeID *int64       `json:"type_id" valid:"required"` // 物资分类编号
	Name   *string      `json:"name" valid:"required"`    // 物资名称
	Price  *types.Money `json:"price" valid:"required"`   // 物资单价
	Stock  *int         `json:"stock" valid:"required"`   // 库存数量
}

// EditMaterialReq 编辑物资请求
// @request
type EditMaterialReq struct {
	ID     *int64       `json:"id"`                       // id
	TypeID *int64       `json:"type_id" valid:"required"` // 物资分类编号
	Name   *string      `json:"name" valid:"required"`    // 物资名称
	Price  *types.Money `json:"price" valid:"required"`   // 物资单价
	Stock  *int         `json:"stock" valid:"required"`   // 库存数量
}

// AddMaterialTypeReq 新增物资分类请求
// @request
type AddMaterialTypeReq struct {
	ID   *int64  `json:"id"`                    // id
	Name *string `json:"name" valid:"required"` // 物资分类名称
}

// EditMaterialTypeReq 编辑物资分类请求
// @request
type EditMaterialTypeReq struct {
	ID   *int64  `json:"id"`                    // id
	Name *string `json:"name" valid:"required"` // 物资分类名称
}

// PageMaterialTypeByKeyReq 分页查询物资分类请求
// @request
type PageMaterialTypeByKeyReq struct {
	PageNum  *int    `json:"page_num" valid:"required"`  // 页码
	PageSize *int    `json:"page_size" valid:"required"` // 条数
	Name     *string `json:"name"`                       // 物资分类名称
}

// ============ MaterialController 响应 ============

// PageMaterialByKeyResp 分页查询物资响应
// @response
type PageMaterialByKeyResp struct {
	ID       types.BigInt `json:"id"`        // id
	Name     string       `json:"name"`      // 物资名称
	TypeName string       `json:"type_name"` // 物资分类
	Price    types.Money  `json:"price"`     // 物资单价
}

// OperateMaterialResp 操作物资响应（继承 OperateMaterialReq）
// @response
type OperateMaterialResp struct {
	OperateMaterialReq
}

// PageMaterialTypeResp 分页查询物资分类响应
// @response
type PageMaterialTypeResp struct {
	ID   types.BigInt `json:"id"`   // 物资分类编号
	Name string       `json:"name"` // 物资分类名称
}
