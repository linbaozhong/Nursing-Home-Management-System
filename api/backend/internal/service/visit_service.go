package service

import (
	"context"

	"api/internal/model/define/dao"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type visit struct{}

var Visit = &visit{}

// PageVisitByKey 分页查询来访登记（联表 elder、user）
// 对应 Java: VisitServiceImpl.pageVisitByKey -> VisitMapper.listVisitByKey
// SQL: SELECT v.*, e.elder_name, u.name AS visitor_user_name FROM visit v
//
//	LEFT JOIN elder e ON e.id = v.elder_id
//	LEFT JOIN user u ON u.id = v.visitor_user_id
//	WHERE (e.elder_name LIKE %key% OR v.id = key) [可选]
//	ORDER BY v.create_time DESC; 再由 PageUtil 内存分页。
//
// Todo: 1) in.Key 非空 -> (tbl<visit>.Id.Eq(in.Key) OR tbl<elder>.ElderName.Like(in.Key))
//
//	2) DB 分页: Count + List(联表 LeftJoin)
//	3) 组装含老人/访客姓名的 VO 并赋值 out
func (v *visit) PageVisitByKey(ctx context.Context, in *dto.PageVisitByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 实现联表分页查询
	return nil
}

// PageSearchElderByKey 分页搜索老人（供来访选择老人）
// 对应 Java: VisitServiceImpl.pageSearchElderByKey -> elderMapper.listElderByKey
// SQL: SELECT * FROM elder WHERE (elder_name LIKE %key% OR id = key) [可选] AND del_flag=0
// todo: 查询 elder 表并分页, 结果赋值 out(需定义老人分页返回类型)
func (v *visit) PageSearchElderByKey(ctx context.Context, in *dto.PageSearchElderByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 elder 表并分页
	return nil
}

// AddVisit 新增来访登记
// 对应 Java: VisitServiceImpl.addVisit -> visitMapper.insertSelective
// todo: 标准 CRUD - dao.Visit(db).InsertOne 写入 visit 表(含 elderId/visitorName/visitTime 等)
func (v *visit) AddVisit(ctx context.Context, in *dto.AddVisitQuery, out *dto.EmptyResp) error {
	// todo: bean := do.NewVisit(); 填充 in; dao.Visit(db).InsertOne(ctx, bean)
	return nil
}

// GetVisitById 根据编号获取来访登记
// 对应 Java: VisitServiceImpl.getVisitById -> visitMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.Visit(db).GetByID(ctx, types.BigInt(in.ID))
func (v *visit) GetVisitById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.Visit(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// EditVisit 编辑来访登记
// 对应 Java: VisitServiceImpl.editVisit -> visitMapper.updateByPrimaryKeySelective
// todo: 标准 CRUD - 按主键更新 visit 表
func (v *visit) EditVisit(ctx context.Context, in *dto.EditVisitQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<visit>.VisitorName.Value(in.VisitorName),
	}
	_, e := dao.Visit(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// RecordLeave 登记离开
// 对应 Java: VisitServiceImpl.recordLeave -> 更新 visit 离开时间/状态
// todo: 更新 visit 的离开时间字段(UpdateById), 结果赋值 out
func (v *visit) RecordLeave(ctx context.Context, in *dto.RecordLeaveQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: tbl<visit>.LeaveTime.Value(in.LeaveTime),
	}
	_, e := dao.Visit(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteVisit 删除来访登记
// 对应 Java: VisitServiceImpl.deleteVisit -> visitMapper.deleteByPrimaryKey
// todo: 标准 CRUD - dao.Visit(db).DeleteById(ctx, types.BigInt(in.ID))
func (v *visit) DeleteVisit(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.Visit(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}
