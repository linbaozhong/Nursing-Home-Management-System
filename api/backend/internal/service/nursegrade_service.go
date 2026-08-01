package service

import (
	"context"

	"api/internal/model/define/dao"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type nursegrade struct{}

var NurseGrade = &nursegrade{}

// PageNurseGradeByKey 分页查询护理等级
// 对应 Java: NurseGradeServiceImpl.pageNurseGradeByKey -> NurseGradeMapper.listNurseGradeByKey
// SQL: SELECT * FROM nurse_grade WHERE (grade_name LIKE %key%) [可选] ORDER BY create_time DESC
// todo: 护理等级分页查询 - dao.NurseGrade(db) 条件 + 分页, 结果赋值 out
func (n *nursegrade) PageNurseGradeByKey(ctx context.Context, in *dto.PageNurseGradeByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 nurse_grade 表并分页
	return nil
}

// GetNurseGradeById 根据编号获取护理等级
// 对应 Java: NurseGradeServiceImpl.getNurseGradeById -> nurseGradeMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.NurseGrade(db).GetByID(ctx, types.BigInt(in.ID))
func (n *nursegrade) GetNurseGradeById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.NurseGrade(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// AddNurseGrade 新增护理等级
// 对应 Java: NurseGradeServiceImpl.addNurseGrade -> nurseGradeMapper.insertSelective
// todo: 标准 CRUD - dao.NurseGrade(db).InsertOne 写入 nurse_grade 表
func (n *nursegrade) AddNurseGrade(ctx context.Context, in *dto.AddNurseGradeQuery, out *dto.EmptyResp) error {
	// todo: bean := do.NewNurseGrade(); 填充 in; dao.NurseGrade(db).InsertOne(ctx, bean)
	return nil
}

// EditNurseGrade 编辑护理等级
// 对应 Java: NurseGradeServiceImpl.editNurseGrade -> nurseGradeMapper.updateByPrimaryKeySelective
// todo: 标准 CRUD - 按主键更新 nurse_grade 表
func (n *nursegrade) EditNurseGrade(ctx context.Context, in *dto.EditNurseGradeQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<nursegrade>.GradeName.Value(in.GradeName),
	}
	_, e := dao.NurseGrade(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteNurseGrade 删除护理等级
// 对应 Java: NurseGradeServiceImpl.deleteNurseGrade -> nurseGradeMapper.deleteByPrimaryKey
// todo: 标准 CRUD - dao.NurseGrade(db).DeleteById(ctx, types.BigInt(in.ID))
func (n *nursegrade) DeleteNurseGrade(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.NurseGrade(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}

// PageNurseByKey 分页查询护理员（user 表中角色为护理员）
// 对应 Java: NurseGradeServiceImpl.pageNurseByKey -> UserMapper.listNurseByKey
// SQL: SELECT * FROM user WHERE (name LIKE %key% OR id = key) [可选] AND role = 护理员
// todo: 查询 user(护理员)表并分页, 结果赋值 out(需定义护理员分页返回类型)
func (n *nursegrade) PageNurseByKey(ctx context.Context, in *dto.PageNurseByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 user 护理员并分页
	return nil
}

// GetNurseById 根据编号获取护理员
// 对应 Java: NurseGradeServiceImpl.getNurseById -> userMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.User(db).GetByID(ctx, types.BigInt(in.ID))
func (n *nursegrade) GetNurseById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.User(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// AddNurse 新增护理员（写 user 表）
// 对应 Java: NurseGradeServiceImpl.addNurse -> userMapper.insertSelective
// todo: 标准 CRUD - dao.User(db).InsertOne 写入 user(角色=护理员)
func (n *nursegrade) AddNurse(ctx context.Context, in *dto.AddNurseQuery, out *dto.EmptyResp) error {
	// todo: bean := do.NewUser(); 填充 in(角色=护理员); dao.User(db).InsertOne(ctx, bean)
	return nil
}

// EditNurse 编辑护理员
// 对应 Java: NurseGradeServiceImpl.editNurse -> userMapper.updateByPrimaryKeySelective
// todo: 标准 CRUD - 按主键更新 user 表
func (n *nursegrade) EditNurse(ctx context.Context, in *dto.EditNurseQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbluser.Name.Value(in.Name),
	}
	_, e := dao.User(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteNurse 删除护理员
// 对应 Java: NurseGradeServiceImpl.deleteNurse -> userMapper.deleteByPrimaryKey
// todo: 标准 CRUD - dao.User(db).DeleteById(ctx, types.BigInt(in.ID))
func (n *nursegrade) DeleteNurse(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.User(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}
