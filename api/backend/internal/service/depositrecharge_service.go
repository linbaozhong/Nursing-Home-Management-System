package service

import (
	"context"

	"api/internal/model/define/dao"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type depositrecharge struct{}

var DepositRecharge = &depositrecharge{}

// PageDepositRechargeByKey 分页查询定金/充值（联表 elder、user）
// 对应 Java: DepositRechargeServiceImpl.pageDepositRechargeByKey -> DepositRechargeMapper.listDepositRechargeByKey
// SQL: SELECT dr.*, e.elder_name, u.name AS charge_user_name FROM deposit_recharge dr
//
//	LEFT JOIN elder e ON e.id = dr.elder_id
//	LEFT JOIN user u ON u.id = dr.charge_user_id
//	WHERE (e.elder_name LIKE %key% OR dr.id = key) [可选]
//	ORDER BY dr.create_time DESC; 再由 PageUtil 内存分页。
//
// Todo: 1) in.Key 非空 -> (tbl<depositrecharge>.Id.Eq(in.Key) OR tbl<elder>.ElderName.Like(in.Key))
//
//	2) DB 分页: Count + List(联表 LeftJoin)
//	3) 组装含老人/负责人姓名的 VO 并赋值 out
func (d *depositrecharge) PageDepositRechargeByKey(ctx context.Context, in *dto.PageDepositRechargeByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 实现联表分页查询
	return nil
}

// GetDepositRechargeById 根据编号获取定金/充值
// 对应 Java: DepositRechargeServiceImpl.getDepositRechargeById -> depositRechargeMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.DepositRecharge(db).GetByID(ctx, types.BigInt(in.ID))
func (d *depositrecharge) GetDepositRechargeById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.DepositRecharge(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// AddDepositRecharge 新增定金/充值（同时更新老人账户余额）
// 对应 Java: DepositRechargeServiceImpl.addDepositRecharge -> insert deposit_recharge + 更新 elder_account 余额
// todo: 事务: 1) dao.DepositRecharge(db).InsertOne; 2) 更新对应 elder 账户余额(余额+/-金额)
func (d *depositrecharge) AddDepositRecharge(ctx context.Context, in *dto.AddDepositRechargeQuery, out *dto.EmptyResp) error {
	// todo: 写入 deposit_recharge 并更新老人账户余额
	return nil
}

// EditDepositRecharge 编辑定金/充值
// 对应 Java: DepositRechargeServiceImpl.editDepositRecharge -> depositRechargeMapper.updateByPrimaryKeySelective
// todo: 标准 CRUD - 按主键更新 deposit_recharge 表
func (d *depositrecharge) EditDepositRecharge(ctx context.Context, in *dto.EditDepositRechargeQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<depositrecharge>.Remark.Value(in.Remark),
	}
	_, e := dao.DepositRecharge(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteDepositRecharge 删除定金/充值（回滚老人账户余额）
// 对应 Java: DepositRechargeServiceImpl.deleteDepositRecharge -> 删记录 + 回滚 elder_account 余额
// todo: 事务: 1) 查原记录; 2) 回滚账户余额; 3) dao.DepositRecharge(db).DeleteById
func (d *depositrecharge) DeleteDepositRecharge(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.DepositRecharge(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}

// PageSearchElderByKey 分页搜索老人（供定金/充值选择老人）
// 对应 Java: DepositRechargeServiceImpl.pageSearchElderByKey -> elderMapper.listElderByKey
// SQL: SELECT * FROM elder WHERE (elder_name LIKE %key% OR id = key) [可选] AND del_flag=0
// todo: 查询 elder 表并分页, 结果赋值 out(需定义老人分页返回类型)
func (d *depositrecharge) PageSearchElderByKey(ctx context.Context, in *dto.PageSearchElderByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 elder 表并分页
	return nil
}

// PageSearchStaffByKey 分页搜索员工（供定金/充值选择负责人）
// 对应 Java: DepositRechargeServiceImpl.pageSearchStaffByKey -> userMapper.listStaffByKey
// SQL: SELECT * FROM user WHERE (name LIKE %key% OR id = key) [可选] AND role 为员工
// todo: 查询 user(员工)表并分页, 结果赋值 out(需定义员工分页返回类型)
func (d *depositrecharge) PageSearchStaffByKey(ctx context.Context, in *dto.PageSearchStaffByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 user 员工并分页
	return nil
}

// GetElderFeeById 根据老人编号获取费用明细
// 对应 Java: DepositRechargeServiceImpl.getElderFeeById -> 聚合查询老人各项费用/余额
// todo: 查询 elder_account / elder_fee 等, 汇总返回老人费用(需定义返回类型)
func (d *depositrecharge) GetElderFeeById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 聚合查询老人费用
	return nil
}

// AuditElderFee 审核老人费用
// 对应 Java: DepositRechargeServiceImpl.auditElderFee -> 更新费用审核状态
// todo: 更新 elder_fee 审核状态/字段(UpdateById), 结果赋值 out
func (d *depositrecharge) AuditElderFee(ctx context.Context, in *dto.AuditElderFeeQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: tbl<elderfee>.AuditStatus.Value(in.AuditStatus),
	}
	_, e := dao.ElderFee(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}
