package service

import (
	"context"

	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblactivetype"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type active struct{}

var Active = &active{}

// PageActiveByKey 分页查询活动（联表 active_type、elder_active、elder、user）
// 对应 Java: ActiveServiceImpl.pageActiveByKey -> ActiveMapper.listActiveByKey
// SQL: SELECT a.*, at.type_name, ea.elder_id, e.elder_name, u.name AS charge_user_name
//
//	FROM active a
//	LEFT JOIN active_type at ON at.id = a.active_type_id
//	LEFT JOIN elder_active ea ON ea.active_id = a.id
//	LEFT JOIN elder e ON e.id = ea.elder_id
//	LEFT JOIN user u ON u.id = a.charge_user_id
//	WHERE (at.type_name LIKE %key% OR a.id = key) [可选]
//	ORDER BY a.create_time DESC; 再由 PageUtil 内存分页。
//
// todo: 1) in.Key 非空 -> (tbl<activetype>.TypeName.Like(in.Key) OR tbl<active>.Id.Eq(in.Key))
//
//	2) DB 分页: Count + List(联表 LeftJoin active_type/elder_active/elder/user)
//	3) 组装含类型名/参与老人/负责人姓名的 VO 并赋值 out
func (a *active) PageActiveByKey(ctx context.Context, in *dto.PageActiveByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 实现联表分页查询
	return nil
}

// GetActiveById 根据编号获取活动
// 对应 Java: ActiveServiceImpl.getActiveById -> activeMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.Active(db).GetByID(ctx, types.BigInt(in.ID)); 另查 elder_active 获取参与老人
func (a *active) GetActiveById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.Active(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	// todo: 另 dao.ElderActive(db).List(ctx, ...) 获取参与老人列表
	return nil
}

// AddActive 新增活动（同时写入参与老人 elder_active）
// 对应 Java: ActiveServiceImpl.addActive -> activeMapper.insertSelective 后批量 insert elder_active
// todo: 事务: 1) dao.Active(db).InsertOne(active); 2) 遍历 in.ElderIdList 批量 dao.ElderActive(db).InsertOne
func (a *active) AddActive(ctx context.Context, in *dto.OperateActiveQuery, out *dto.EmptyResp) error {
	// todo: bean := do.NewActive(); 填充 in; dao.Active(db).InsertOne(ctx, bean)
	// todo: 遍历 in.ElderIdList 写入 elder_active 表
	return nil
}

// EditActive 编辑活动（先删后插参与老人 elder_active）
// 对应 Java: ActiveServiceImpl.editActive -> 更新 active + 删除旧 elder_active + 批量新增
// todo: 事务: 1) dao.Active(db).UpdateById; 2) 删除 elder_active(active_id); 3) 批量新增 elder_active
func (a *active) EditActive(ctx context.Context, in *dto.OperateActiveQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<active>.ActiveName.Value(in.ActiveName),
	}
	_, e := dao.Active(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	if e != nil {
		return e
	}
	// todo: 删除并重新写入 elder_active
	return nil
}

// DeleteActive 删除活动
// 对应 Java: ActiveServiceImpl.deleteActive -> activeMapper.deleteByPrimaryKey(级联删 elder_active)
// todo: 事务: 1) 删 elder_active(active_id); 2) dao.Active(db).DeleteById(ctx, types.BigInt(in.ID))
func (a *active) DeleteActive(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	// todo: 先删 elder_active, 再删 active
	_, e := dao.Active(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}

// PageSearchElderByKey 分页搜索老人（供活动选择参与老人）
// 对应 Java: ActiveServiceImpl.pageSearchElderByKey -> elderMapper.listElderByKey
// SQL: SELECT * FROM elder WHERE (elder_name LIKE %key% OR id = key) [可选] AND del_flag=0
// todo: 实现老人分页查询 - 复用 elder 表条件, 结果赋值 out(需定义老人分页返回类型)
func (a *active) PageSearchElderByKey(ctx context.Context, in *dto.PageSearchElderByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 elder 表并分页
	return nil
}

// GetActiveType 获取活动类型下拉框
// 对应 Java: ActiveServiceImpl.getActiveType -> activeTypeMapper.listNotDelActiveType(null)
// SQL: SELECT * FROM active_type WHERE del_flag = 0
// todo: 查询未删除的活动类型, 组装 []dto.DropDown{ID, Name} 赋值 out
func (a *active) GetActiveType(ctx context.Context, in *dto.EmptyReq, out *dto.EmptyResp) error {
	list, _, e := dao.ActiveType(db).List(ctx, ace.Where(tblactivetype.DelFlag.Eq("N")))
	if e != nil {
		return e
	}
	dropList := make([]dto.DropDown, 0, len(list))
	for _, v := range list {
		dropList = append(dropList, dto.DropDown{ID: int64(v.Id), Name: string(v.Name)})
	}
	// todo: 结果赋值 out
	return nil
}
