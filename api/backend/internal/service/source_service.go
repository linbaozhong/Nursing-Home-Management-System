package service

import (
	"api/internal/constant"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblsource"
	"api/internal/model/do"
	"api/internal/model/dto"
	"context"
	"errors"
	"github.com/linbaozhong/gentity/pkg/types"
)

type source struct{}

var Source = &source{}

// PageSourceByKey 分页查询来源渠道
// 对应 Java: SourceServiceImpl.pageSourceByKey -> SourceFunc.listNotDelSource
// 逻辑: 查询未删除(del_flag=N)的来源, 可选按名称模糊匹配, 按创建时间倒序, 分页返回
func (s *source) PageSourceByKey(ctx context.Context, in *dto.PageSourceByKeyQuery, out *[]dto.PageSourceByKeyVO) error {
	q := db.Table(do.SourceTableName).
		Where(
			tblsource.DelFlag.Eq(constant.YesNoNo),
		).
		Desc(tblsource.CreateTime)
	if in.SourceName != nil {
		q.And(tblsource.Name.Like(*in.SourceName))
	}
	e := q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Select().
		Gets(ctx, out)
	return e
}

// AddSource 新增来源渠道
// 对应 Java: SourceServiceImpl.addSource -> SourceFunc.getSourceByName 判重 + sourceMapper.insert
func (s *source) AddSource(ctx context.Context, in *dto.StringReq, out *dto.EmptyResp) error {
	// 判断来源渠道是否已存在
	has, e := dao.Source(db).Exists(ctx,
		tblsource.Name.Eq(*in.Value),
		tblsource.DelFlag.Eq(constant.YesNoNo),
	)
	if e != nil {
		return e
	}
	if has {
		return errors.New("来源渠道已存在")
	}
	// 初始化来源渠道
	bean := do.NewSource()
	defer bean.Free()

	bean.Name = types.String(*in.Value)
	bean.DelFlag = constant.YesNoNo
	// 新增
	_, e = dao.Source(db).InsertOne(ctx, bean)
	return e
}

// GetSourceById 根据编号获取来源渠道
// 对应 Java: SourceServiceImpl.getSourceById -> sourceMapper.selectById
func (s *source) GetSourceById(ctx context.Context, in *dto.IDReq, out *dto.OperateSourceVo) error {
	return db.Table(do.SourceTableName).
		Where(tblsource.Id.Eq(*in.ID)).
		Cols(
			tblsource.Id,
			tblsource.Name).
		Select().Get(ctx, out)
}

// EditSource 编辑来源渠道
// 对应 Java: SourceServiceImpl.editSource -> SourceFunc.getSourceByName 判重(排除自身) + sourceMapper.updateById
func (s *source) EditSource(ctx context.Context, in *dto.OperateSourceQuery, out *dto.EmptyResp) error {
	// 判断来源渠道是否已存在(排除自身)
	has, e := dao.Source(db).Exists(ctx,
		tblsource.Name.Eq(*in.Name),
		tblsource.DelFlag.Eq(constant.YesNoNo),
		tblsource.Id.NotEq(types.BigInt(*in.ID)),
	)
	if e != nil {
		return e
	}
	if has {
		return errors.New("来源渠道已存在")
	}

	// 修改
	_, e = dao.Source(db).UpdateById(ctx, types.BigInt(*in.ID),
		tblsource.Name.Set(*in.Name),
	)
	return e
}

// DeleteSource 删除来源渠道(软删除)
// 对应 Java: SourceServiceImpl.deleteSource -> sourceMapper.updateById(delFlag=YES)
func (s *source) DeleteSource(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.Source(db).UpdateById(ctx, types.BigInt(*in.ID),
		tblsource.DelFlag.Set(constant.YesNoYes),
	)
	return e
}
