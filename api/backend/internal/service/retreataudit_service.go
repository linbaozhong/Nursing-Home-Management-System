package service

import (
	"context"

	"api/internal/model/define/dao"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type retreataudit struct{}

var RetreatAudit = &retreataudit{}

// PageRetreatAuditByKey 分页查询退住审核（联表 elder、user、retreat_apply）
// 对应 Java: RetreatAuditServiceImpl.pageRetreatAuditByKey -> RetreatAuditMapper.listRetreatAuditByKey
// SQL: SELECT ra.*, e.elder_name, u.name AS audit_user_name, a.retreat_reason
//
//	FROM retreat_audit ra
//	LEFT JOIN retreat_apply a ON a.id = ra.retreat_apply_id
//	LEFT JOIN elder e ON e.id = ra.elder_id
//	LEFT JOIN user u ON u.id = ra.audit_user_id
//	WHERE (e.elder_name LIKE %key% OR ra.id = key) [可选]
//	ORDER BY ra.create_time DESC; 再由 PageUtil 内存分页。
//
// Todo: 1) in.Key 非空 -> (tbl<retreataudit>.Id.Eq(in.Key) OR tbl<elder>.ElderName.Like(in.Key))
//
//	2) DB 分页: Count + List(联表 LeftJoin)
//	3) 组装含老人/审核人/退住原因的 VO 并赋值 out
func (r *retreataudit) PageRetreatAuditByKey(ctx context.Context, in *dto.PageRetreatAuditByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 实现联表分页查询
	return nil
}

// GetRetreatAuditById 根据编号获取退住审核
// 对应 Java: RetreatAuditServiceImpl.getRetreatAuditById -> retreatAuditMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.RetreatAudit(db).GetByID(ctx, types.BigInt(in.ID))
func (r *retreataudit) GetRetreatAuditById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.RetreatAudit(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// AuditRetreat 审核退住（审核通过时结算费用/释放床位）
// 对应 Java: RetreatAuditServiceImpl.auditRetreat -> 更新 retreat_audit 审核结果 + 结算 + 更新退住申请状态
// todo: 事务: 1) UpdateById(retreat_audit 审核结果); 2) 结算老人费用(更新 elder_account/order_fee); 3) 释放床位(更新 bed/room 状态)
func (r *retreataudit) AuditRetreat(ctx context.Context, in *dto.AuditRetreatQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: tbl<retreataudit>.AuditResult.Value(in.AuditResult),
		//       tbl<retreataudit>.AuditRemark.Value(in.AuditRemark),
	}
	_, e := dao.RetreatAudit(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	if e != nil {
		return e
	}
	// todo: 审核通过 -> 结算费用、释放床位、更新 retreat_apply 状态
	return nil
}

// PageSearchElderByKey 分页搜索老人（供选择老人）
// 对应 Java: RetreatAuditServiceImpl.pageSearchElderByKey -> elderMapper.listElderByKey
// SQL: SELECT * FROM elder WHERE (elder_name LIKE %key% OR id = key) [可选] AND del_flag=0
// todo: 查询 elder 表并分页, 结果赋值 out(需定义老人分页返回类型)
func (r *retreataudit) PageSearchElderByKey(ctx context.Context, in *dto.PageSearchElderByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 elder 表并分页
	return nil
}
