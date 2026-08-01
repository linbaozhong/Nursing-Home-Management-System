package service

import (
	"context"

	"api/internal/model/define/dao"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type nursereserve struct{}

var NurseReserve = &nursereserve{}

// PageNurseReserveByKey 分页查询护理预订（联表 elder、user、reserve、nurse_grade）
// 对应 Java: NurseReserveServiceImpl.pageNurseReserveByKey -> NurseReserveMapper.listNurseReserveByKey
// SQL: SELECT nr.*, e.elder_name, u.name AS nurse_user_name, r.reserve_no, ng.grade_name
//
//	FROM nurse_reserve nr
//	LEFT JOIN elder e ON e.id = nr.elder_id
//	LEFT JOIN user u ON u.id = nr.nurse_user_id
//	LEFT JOIN reserve r ON r.id = nr.reserve_id
//	LEFT JOIN nurse_grade ng ON ng.id = nr.nurse_grade_id
//	WHERE (e.elder_name LIKE %key% OR nr.id = key) [可选]
//	ORDER BY nr.create_time DESC; 再由 PageUtil 内存分页。
//
// Todo: 1) in.Key 非空 -> (tbl<nursereserve>.Id.Eq(in.Key) OR tbl<elder>.ElderName.Like(in.Key))
//
//	2) DB 分页: Count + List(联表 LeftJoin)
//	3) 组装含老人/护理员/护理等级名的 VO 并赋值 out
func (n *nursereserve) PageNurseReserveByKey(ctx context.Context, in *dto.PageNurseReserveByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 实现联表分页查询
	return nil
}

// GetNurseReserveByReserveIdAndElderId 根据预订编号与老人编号获取护理预订（编辑回显）
// 对应 Java: NurseReserveServiceImpl.getNurseReserveByReserveIdAndElderId
// SQL: SELECT * FROM nurse_reserve WHERE reserve_id = #{reserveId} AND elder_id = #{elderId}
// todo: 标准查询 - dao.NurseReserve(db).Get(ctx, ace.Where(tbl<nursereserve>.ReserveId.Eq(in.ReserveId), tbl<nursereserve>.ElderId.Eq(in.ElderId)))
func (n *nursereserve) GetNurseReserveByReserveIdAndElderId(ctx context.Context, in *dto.GetNurseReserveByReserveIdAndElderIdQuery, out *dto.EmptyResp) error {
	// todo: obj, has, e := dao.NurseReserve(db).Get(ctx, ace.Where(
	//	tbl<nursereserve>.ReserveId.Eq(in.ReserveId),
	//	tbl<nursereserve>.ElderId.Eq(in.ElderId),
	// ))
	return nil
}

// AddNurseReserve 新增护理预订
// 对应 Java: NurseReserveServiceImpl.addNurseReserve -> nurseReserveMapper.insertSelective
// todo: 标准 CRUD - dao.NurseReserve(db).InsertOne 写入 nurse_reserve 表
func (n *nursereserve) AddNurseReserve(ctx context.Context, in *dto.AddNurseReserveQuery, out *dto.EmptyResp) error {
	// todo: bean := do.NewNurseReserve(); 填充 in; dao.NurseReserve(db).InsertOne(ctx, bean)
	return nil
}

// EditNurseReserve 编辑护理预订
// 对应 Java: NurseReserveServiceImpl.editNurseReserve -> nurseReserveMapper.updateByPrimaryKeySelective
// todo: 标准 CRUD - 按主键更新 nurse_reserve 表
func (n *nursereserve) EditNurseReserve(ctx context.Context, in *dto.EditNurseReserveQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<nursereserve>.NurseGradeId.Value(in.NurseGradeId),
	}
	_, e := dao.NurseReserve(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteNurseReserve 删除护理预订
// 对应 Java: NurseReserveServiceImpl.deleteNurseReserve -> nurseReserveMapper.deleteByPrimaryKey
// todo: 标准 CRUD - dao.NurseReserve(db).DeleteById(ctx, types.BigInt(in.ID))
func (n *nursereserve) DeleteNurseReserve(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.NurseReserve(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}

// PageSearchElderByKey 分页搜索老人（供护理预订选择老人）
// 对应 Java: NurseReserveServiceImpl.pageSearchElderByKey -> elderMapper.listElderByKey
// SQL: SELECT * FROM elder WHERE (elder_name LIKE %key% OR id = key) [可选] AND del_flag=0
// todo: 查询 elder 表并分页, 结果赋值 out(需定义老人分页返回类型)
func (n *nursereserve) PageSearchElderByKey(ctx context.Context, in *dto.PageSearchElderByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 elder 表并分页
	return nil
}

// ListNurseStaff 护理员工列表
// 对应 Java: NurseReserveServiceImpl.listNurseStaff -> userMapper.listNurseStaff(角色=护理员)
// SQL: SELECT * FROM user WHERE role = 护理员
// todo: 查询 user(护理员)列表, 结果赋值 out(需定义返回类型)
func (n *nursereserve) ListNurseStaff(ctx context.Context, in *dto.EmptyReq, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 user 护理员列表
	return nil
}
