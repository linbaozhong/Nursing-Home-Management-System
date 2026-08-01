package service

import (
	"context"

	"api/internal/model/define/dao"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type retreatapply struct{}

var RetreatApply = &retreatapply{}

// PageRetreatApplyByKey 分页查询退住申请（联表 elder、user）
// 对应 Java: RetreatApplyServiceImpl.pageRetreatApplyByKey -> RetreatApplyMapper.listRetreatApplyByKey
// SQL: SELECT ra.*, e.elder_name, u.name AS apply_user_name FROM retreat_apply ra
//
//	LEFT JOIN elder e ON e.id = ra.elder_id
//	LEFT JOIN user u ON u.id = ra.apply_user_id
//	WHERE (e.elder_name LIKE %key% OR ra.id = key) [可选]
//	ORDER BY ra.create_time DESC; 再由 PageUtil 内存分页。
//
// Todo: 1) in.Key 非空 -> (tbl<retreatapply>.Id.Eq(in.Key) OR tbl<elder>.ElderName.Like(in.Key))
//
//	2) DB 分页: Count + List(联表 LeftJoin)
//	3) 组装含老人/申请人姓名的 VO 并赋值 out
func (r *retreatapply) PageRetreatApplyByKey(ctx context.Context, in *dto.PageRetreatApplyByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 实现联表分页查询
	return nil
}

// GetRetreatApplyById 根据编号获取退住申请
// 对应 Java: RetreatApplyServiceImpl.getRetreatApplyById -> retreatApplyMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.RetreatApply(db).GetByID(ctx, types.BigInt(in.ID))
func (r *retreatapply) GetRetreatApplyById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.RetreatApply(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// PageSearchElderByKey 分页搜索老人（供退住申请选择老人）
// 对应 Java: RetreatApplyServiceImpl.pageSearchElderByKey -> elderMapper.listElderByKey
// SQL: SELECT * FROM elder WHERE (elder_name LIKE %key% OR id = key) [可选] AND del_flag=0
// todo: 查询 elder 表并分页, 结果赋值 out(需定义老人分页返回类型)
func (r *retreatapply) PageSearchElderByKey(ctx context.Context, in *dto.PageSearchElderByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 elder 表并分页
	return nil
}

// AddRetreatApply 新增退住申请
// 对应 Java: RetreatApplyServiceImpl.addRetreatApply -> retreatApplyMapper.insertSelective
// todo: 标准 CRUD - dao.RetreatApply(db).InsertOne 写入 retreat_apply 表(含 elderId/applyUserId/retreatReason 等)
func (r *retreatapply) AddRetreatApply(ctx context.Context, in *dto.AddRetreatApplyQuery, out *dto.EmptyResp) error {
	// todo: bean := do.NewRetreatApply(); 填充 in; dao.RetreatApply(db).InsertOne(ctx, bean)
	return nil
}

// EditRetreatApply 编辑退住申请
// 对应 Java: RetreatApplyServiceImpl.editRetreatApply -> retreatApplyMapper.updateByPrimaryKeySelective
// todo: 标准 CRUD - 按主键更新 retreat_apply 表
func (r *retreatapply) EditRetreatApply(ctx context.Context, in *dto.EditRetreatApplyQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<retreatapply>.RetreatReason.Value(in.RetreatReason),
	}
	_, e := dao.RetreatApply(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteRetreatApply 删除退住申请
// 对应 Java: RetreatApplyServiceImpl.deleteRetreatApply -> retreatApplyMapper.deleteByPrimaryKey
// todo: 标准 CRUD - dao.RetreatApply(db).DeleteById(ctx, types.BigInt(in.ID))
func (r *retreatapply) DeleteRetreatApply(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.RetreatApply(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}
