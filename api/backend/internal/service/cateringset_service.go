package service

import (
	"context"

	"api/internal/model/define/dao"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type cateringset struct{}

var CateringSet = &cateringset{}

// PageCateringSetByKey 分页查询套餐（联表 catering_set 套餐主表 + 关联菜品）
// 对应 Java: CateringSetServiceImpl.pageCateringSetByKey -> CateringSetMapper.listCateringSetByKey
// SQL: SELECT cs.* FROM catering_set cs WHERE (cs.set_name LIKE %key%) [可选] ORDER BY create_time DESC
// todo: 套餐分页查询 - dao.CateringSet(db) 条件 + 分页, 结果赋值 out(如需菜品明细再联 catering_set_detail)
func (c *cateringset) PageCateringSetByKey(ctx context.Context, in *dto.PageCateringSetByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 catering_set 表并分页
	return nil
}

// GetCateringSetById 根据编号获取套餐
// 对应 Java: CateringSetServiceImpl.getCateringSetById -> cateringSetMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.CateringSet(db).GetByID(ctx, types.BigInt(in.ID))
func (c *cateringset) GetCateringSetById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.CateringSet(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// AddCateringSet 新增套餐（含套餐菜品明细 catering_set_detail）
// 对应 Java: CateringSetServiceImpl.addCateringSet -> cateringSetMapper.insertSelective + 批量 insert detail
// todo: 事务: 1) dao.CateringSet(db).InsertOne; 2) 遍历 in.DishesList 批量 dao.CateringSetDetail(db).InsertOne
func (c *cateringset) AddCateringSet(ctx context.Context, in *dto.OperateCateringSetQuery, out *dto.EmptyResp) error {
	// todo: 写入 catering_set 主表 + catering_set_detail 明细
	return nil
}

// EditCateringSet 编辑套餐（先删后插明细）
// 对应 Java: CateringSetServiceImpl.editCateringSet -> 更新主表 + 删旧明细 + 批量新增
// todo: 事务: 1) UpdateById; 2) 删 catering_set_detail(set_id); 3) 批量新增明细
func (c *cateringset) EditCateringSet(ctx context.Context, in *dto.OperateCateringSetQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<cateringset>.SetName.Value(in.SetName),
	}
	_, e := dao.CateringSet(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	if e != nil {
		return e
	}
	// todo: 删除并重新写入 catering_set_detail
	return nil
}

// DeleteCateringSet 删除套餐
// 对应 Java: CateringSetServiceImpl.deleteCateringSet -> 删主表(级联删明细)
// todo: 事务: 1) 删 catering_set_detail(set_id); 2) dao.CateringSet(db).DeleteById(ctx, types.BigInt(in.ID))
func (c *cateringset) DeleteCateringSet(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	// todo: 先删明细再删主表
	_, e := dao.CateringSet(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}
