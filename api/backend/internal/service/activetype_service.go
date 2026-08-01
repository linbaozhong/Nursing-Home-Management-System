package service

import (
	"context"

	"api/internal/model/define/dao"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type activetype struct{}

var ActiveType = &activetype{}

// PageActiveTypeByKey 分页查询活动类型
// 对应 Java: ActiveTypeServiceImpl.pageActiveTypeByKey -> ActiveTypeMapper.listActiveTypeByKey
// SQL: SELECT * FROM active_type WHERE (type_name LIKE %key%) [可选] ORDER BY create_time DESC
// todo: 类型分页查询 - dao.ActiveType(db) 条件 + 分页, 结果赋值 out
func (a *activetype) PageActiveTypeByKey(ctx context.Context, in *dto.PageActiveTypeByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 active_type 表并分页
	return nil
}

// GetActiveTypeById 根据编号获取活动类型
// 对应 Java: ActiveTypeServiceImpl.getActiveTypeById -> activeTypeMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.ActiveType(db).GetByID(ctx, types.BigInt(in.ID))
func (a *activetype) GetActiveTypeById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.ActiveType(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// AddActiveType 新增活动类型
// 对应 Java: ActiveTypeServiceImpl.addActiveType -> activeTypeMapper.insertSelective
// todo: 标准 CRUD - dao.ActiveType(db).InsertOne 写入 active_type 表
func (a *activetype) AddActiveType(ctx context.Context, in *dto.AddActiveTypeQuery, out *dto.EmptyResp) error {
	// todo: bean := do.NewActiveType(); 填充 in; dao.ActiveType(db).InsertOne(ctx, bean)
	return nil
}

// EditActiveType 编辑活动类型
// 对应 Java: ActiveTypeServiceImpl.editActiveType -> activeTypeMapper.updateByPrimaryKeySelective
// todo: 标准 CRUD - 按主键更新 active_type 表
func (a *activetype) EditActiveType(ctx context.Context, in *dto.OperateActiveTypeQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<activetype>.TypeName.Value(in.TypeName),
	}
	_, e := dao.ActiveType(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteActiveType 删除活动类型
// 对应 Java: ActiveTypeServiceImpl.deleteActiveType -> activeTypeMapper.deleteByPrimaryKey
// todo: 标准 CRUD - dao.ActiveType(db).DeleteById(ctx, types.BigInt(in.ID))
func (a *activetype) DeleteActiveType(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.ActiveType(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}
