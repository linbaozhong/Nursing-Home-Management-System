package service

import (
	"context"

	"api/internal/model/define/dao"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type source struct{}

var Source = &source{}

// PageSourceByKey 分页查询来源渠道
// 对应 Java: SourceServiceImpl.pageSourceByKey -> SourceMapper.listSourceByKey
// SQL: SELECT * FROM source WHERE (source_name LIKE %key%) [可选] ORDER BY create_time DESC
// todo: 来源分页查询 - dao.Source(db) 条件 + 分页, 结果赋值 out
func (s *source) PageSourceByKey(ctx context.Context, in *dto.PageSourceByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 source 表并分页
	return nil
}

// AddSource 新增来源渠道
// 对应 Java: SourceServiceImpl.addSource -> sourceMapper.insertSelective
// todo: 标准 CRUD - dao.Source(db).InsertOne 写入 source 表(含 sourceName)
func (s *source) AddSource(ctx context.Context, in *dto.StringReq, out *dto.EmptyResp) error {
	// todo: bean := do.NewSource(); 填充 in; dao.Source(db).InsertOne(ctx, bean)
	return nil
}

// GetSourceById 根据编号获取来源渠道
// 对应 Java: SourceServiceImpl.getSourceById -> sourceMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.Source(db).GetByID(ctx, types.BigInt(in.ID))
func (s *source) GetSourceById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.Source(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// EditSource 编辑来源渠道
// 对应 Java: SourceServiceImpl.editSource -> sourceMapper.updateByPrimaryKeySelective
// todo: 标准 CRUD - 按主键更新 source 表
func (s *source) EditSource(ctx context.Context, in *dto.OperateSourceQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<source>.SourceName.Value(in.SourceName),
	}
	_, e := dao.Source(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteSource 删除来源渠道
// 对应 Java: SourceServiceImpl.deleteSource -> sourceMapper.deleteByPrimaryKey
// todo: 标准 CRUD - dao.Source(db).DeleteById(ctx, types.BigInt(in.ID))
func (s *source) DeleteSource(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.Source(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}
