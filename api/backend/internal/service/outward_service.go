package service

import (
	"context"

	"api/internal/model/define/dao"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type outward struct{}

var Outward = &outward{}

// PageOutwardByKey 分页查询外出（联表 elder、user）
// 对应 Java: OutwardServiceImpl.pageOutwardByKey -> OutwardMapper.listOutwardByKey
// SQL: SELECT o.*, e.elder_name, u.name AS accompany_user_name FROM outward o
//
//	LEFT JOIN elder e ON e.id = o.elder_id
//	LEFT JOIN user u ON u.id = o.accompany_user_id
//	WHERE (e.elder_name LIKE %key% OR o.id = key) [可选]
//	ORDER BY o.create_time DESC; 再由 PageUtil 内存分页。
//
// Todo: 1) in.Key 非空 -> (tbl<outward>.Id.Eq(in.Key) OR tbl<elder>.ElderName.Like(in.Key))
//
//	2) DB 分页: Count + List(联表 LeftJoin)
//	3) 组装含老人/陪同人姓名的 VO 并赋值 out
func (o *outward) PageOutwardByKey(ctx context.Context, in *dto.PageOutwardByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 实现联表分页查询
	return nil
}

// GetOutwardById 根据编号获取外出
// 对应 Java: OutwardServiceImpl.getOutwardById -> outwardMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.Outward(db).GetByID(ctx, types.BigInt(in.ID))
func (o *outward) GetOutwardById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.Outward(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// AddOutward 新增外出
// 对应 Java: OutwardServiceImpl.addOutward -> outwardMapper.insertSelective
// todo: 标准 CRUD - dao.Outward(db).InsertOne 写入 outward 表(含 elderId/leaveTime/accompanyUserId 等)
func (o *outward) AddOutward(ctx context.Context, in *dto.AddOutwardQuery, out *dto.EmptyResp) error {
	// todo: bean := do.NewOutward(); 填充 in; dao.Outward(db).InsertOne(ctx, bean)
	return nil
}

// EditOutward 编辑外出
// 对应 Java: OutwardServiceImpl.editOutward -> outwardMapper.updateByPrimaryKeySelective
// todo: 标准 CRUD - 按主键更新 outward 表
func (o *outward) EditOutward(ctx context.Context, in *dto.EditOutwardQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<outward>.LeaveTime.Value(in.LeaveTime),
	}
	_, e := dao.Outward(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteOutward 删除外出
// 对应 Java: OutwardServiceImpl.deleteOutward -> outwardMapper.deleteByPrimaryKey
// todo: 标准 CRUD - dao.Outward(db).DeleteById(ctx, types.BigInt(in.ID))
func (o *outward) DeleteOutward(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.Outward(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}

// PageSearchElderByKey 分页搜索老人（供外出选择老人）
// 对应 Java: OutwardServiceImpl.pageSearchElderByKey -> elderMapper.listElderByKey
// SQL: SELECT * FROM elder WHERE (elder_name LIKE %key% OR id = key) [可选] AND del_flag=0
// todo: 查询 elder 表并分页, 结果赋值 out(需定义老人分页返回类型)
func (o *outward) PageSearchElderByKey(ctx context.Context, in *dto.PageSearchElderByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 elder 表并分页
	return nil
}

// DelayReturn 延期返回
// 对应 Java: OutwardServiceImpl.delayReturn -> 更新 outward 预计返回时间/状态
// todo: 更新 outward 的预计返回时间字段(UpdateById), 结果赋值 out
func (o *outward) DelayReturn(ctx context.Context, in *dto.DelayReturnQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: tbl<outward>.ExpectReturnTime.Value(in.ExpectReturnTime),
	}
	_, e := dao.Outward(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// RecordReturn 登记返回
// 对应 Java: OutwardServiceImpl.recordReturn -> 更新 outward 实际返回时间/状态
// todo: 更新 outward 的实际返回时间字段(UpdateById), 结果赋值 out
func (o *outward) RecordReturn(ctx context.Context, in *dto.RecordReturnQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: tbl<outward>.ActualReturnTime.Value(in.ActualReturnTime),
	}
	_, e := dao.Outward(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}
