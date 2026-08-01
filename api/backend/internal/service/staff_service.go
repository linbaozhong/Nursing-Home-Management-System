package service

import (
	"context"

	"api/internal/model/define/dao"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type staff struct{}

var Staff = &staff{}

// GetRole 获取角色列表（字典/常量）
// 对应 Java: StaffServiceImpl.getRole -> 返回员工角色枚举(管理员/护理员/... )
// todo: 返回角色枚举, 结果赋值 out(需定义返回类型)
func (s *staff) GetRole(ctx context.Context, in *dto.EmptyReq, out *dto.EmptyResp) error {
	// todo: 返回角色枚举列表
	return nil
}

// PageStaffByKey 分页查询员工（联表 role/dept）
// 对应 Java: StaffServiceImpl.pageStaffByKey -> UserMapper.listStaffByKey
// SQL: SELECT u.*, r.role_name FROM user u LEFT JOIN role r ON r.id = u.role_id
//
//	WHERE (u.name LIKE %key% OR u.id = key) [可选] AND u.role != 管理员
//	ORDER BY u.create_time DESC; 再由 PageUtil 内存分页。
//
// todo: 1) in.Key 非空 -> (tbluser.Id.Eq(in.Key) OR tbluser.Name.Like(in.Key))
//
//	2) DB 分页: Count + List(联表 role)
//	3) 组装含角色名的 VO 并赋值 out
func (s *staff) PageStaffByKey(ctx context.Context, in *dto.PageStaffByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 实现联表分页查询
	return nil
}

// AddStaff 新增员工（写 user 表）
// 对应 Java: StaffServiceImpl.addStaff -> userMapper.insertSelective
// todo: 标准 CRUD - dao.User(db).InsertOne 写入 user(含 name/roleId/phone 等)
func (s *staff) AddStaff(ctx context.Context, in *dto.OperateStaffQuery, out *dto.EmptyResp) error {
	// todo: bean := do.NewUser(); 填充 in; dao.User(db).InsertOne(ctx, bean)
	return nil
}

// GetStaffById 根据编号获取员工
// 对应 Java: StaffServiceImpl.getStaffById -> userMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.User(db).GetByID(ctx, types.BigInt(in.ID))
func (s *staff) GetStaffById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.User(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// EditStaff 编辑员工
// 对应 Java: StaffServiceImpl.editStaff -> userMapper.updateByPrimaryKeySelective
// todo: 标准 CRUD - 按主键更新 user 表
func (s *staff) EditStaff(ctx context.Context, in *dto.OperateStaffQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbluser.Name.Value(in.Name),
	}
	_, e := dao.User(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// LeaveStaff 员工离职（逻辑删除/置状态）
// 对应 Java: StaffServiceImpl.leaveStaff -> 更新 user 状态为离职(del_flag=1)
// todo: 更新 user 状态/删除标志(UpdateById), 结果赋值 out
func (s *staff) LeaveStaff(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: tbluser.DelFlag.Value(1) 或 tbluser.Status.Value(离职),
	}
	_, e := dao.User(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}
