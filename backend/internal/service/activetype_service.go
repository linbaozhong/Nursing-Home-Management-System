package service

import (
	"api/internal/model/define/table/tblactivetype"
	"context"
	"errors"

	"api/internal/constant"
	"api/internal/lib"
	"api/internal/model/define/dao"
	"api/internal/model/do"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type activetype struct{}

var ActiveType = &activetype{}

// PageActiveTypeByKey 分页查询活动类型
// 对应 Java: ActiveTypeServiceImpl.pageActiveTypeByKey -> ActiveTypeFunc.listNotDelActiveType
func (a *activetype) PageActiveTypeByKey(ctx context.Context, in *dto.PageActiveTypeByKeyReq, out *[]dto.PageActiveTypeByKeyResp) error {
	q := db.Table(tblactivetype.TableName).Where(tblactivetype.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblactivetype.State.NotEq(types.Int8(constant.StateDeleted)))
	if in.ActiveTypeName != nil {
		q.And(tblactivetype.Name.Like(*in.ActiveTypeName))
	}
	e := q.Desc(tblactivetype.Id).
		Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblactivetype.Id,
			tblactivetype.Name,
		).
		Select().
		Gets(ctx, out)
	return e
}

// GetActiveTypeById 根据编号获取活动类型
// 对应 Java: ActiveTypeServiceImpl.getActiveTypeById -> activeTypeMapper.selectById
func (a *activetype) GetActiveTypeById(ctx context.Context, in *dto.IDReq, out *dto.OperateActiveTypeResp) error {
	obj, has, e := dao.ActiveType(db).Get(ctx, ace.Where(tblactivetype.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblactivetype.Id.Eq(types.BigInt(*in.ID))))
	if e != nil {
		return e
	}
	if !has || obj == nil {
		return errors.New("活动类型不存在")
	}
	*out.ID = int64(obj.Id)
	*out.Name = obj.Name.String()
	return nil
}

// AddActiveType 新增活动类型
// 对应 Java: ActiveTypeServiceImpl.addActiveType -> activeTypeMapper.insert
func (a *activetype) AddActiveType(ctx context.Context, in *dto.AddActiveTypeReq, out *dto.EmptyResp) error {
	// 判断活动分类是否已存在
	exist, e := a.getActiveTypeByName(ctx, *in.Name)
	if e != nil {
		return e
	}
	if exist {
		return errors.New("活动分类已存在")
	}
	bean := do.NewActiveType()
	bean.TenantId = types.BigInt(lib.TenantID(ctx))
	bean.Name = types.String(*in.Name)
	bean.State = types.Int8(constant.StateEnabled)
	_, e = dao.ActiveType(db).InsertOne(ctx, bean)
	return e
}

// EditActiveType 编辑活动类型
// 对应 Java: ActiveTypeServiceImpl.editActiveType -> activeTypeMapper.updateById
func (a *activetype) EditActiveType(ctx context.Context, in *dto.OperateActiveTypeReq, out *dto.EmptyResp) error {
	// 判断活动分类是否已存在（排除自身）
	exist, e := a.getActiveTypeByNameExclude(ctx, *in.Name, *in.ID)
	if e != nil {
		return e
	}
	if exist {
		return errors.New("活动分类已存在")
	}
	sets := []dialect.Setter{
		tblactivetype.Name.Set(types.String(*in.Name)),
	}
	_, e = dao.ActiveType(db).UpdateById(ctx, types.BigInt(*in.ID), sets...)
	return e
}

// DeleteActiveType 删除活动类型（逻辑删除）
// 对应 Java: ActiveTypeServiceImpl.deleteActiveType -> activeTypeMapper.updateById(delFlag=Y)
func (a *activetype) DeleteActiveType(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		tblactivetype.State.Set(types.Int8(constant.StateDeleted)),
	}
	_, e := dao.ActiveType(db).UpdateById(ctx, types.BigInt(*in.ID), sets...)
	return e
}

// getActiveTypeByName 根据名称判断是否存在未删除的活动分类
func (a *activetype) getActiveTypeByName(ctx context.Context, name string) (bool, error) {
	return dao.ActiveType(db).Exists(ctx,
		tblactivetype.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
		tblactivetype.Name.Eq(types.String(name)),
		tblactivetype.State.NotEq(types.Int8(constant.StateDeleted)),
	)
}

// getActiveTypeByNameExclude 根据名称判断是否存在未删除的活动分类（排除指定 id）
func (a *activetype) getActiveTypeByNameExclude(ctx context.Context, name string, excludeID int64) (bool, error) {
	return dao.ActiveType(db).Exists(ctx,
		tblactivetype.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
		tblactivetype.Name.Eq(types.String(name)),
		tblactivetype.State.NotEq(types.Int8(constant.StateDeleted)),
		tblactivetype.Id.NotEq(types.BigInt(excludeID)),
	)
}
