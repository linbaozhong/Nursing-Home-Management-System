package service

import (
	"context"

	"api/internal/model/define/dao"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type elderrecord struct{}

var ElderRecord = &elderrecord{}

// PageElderRecordByKey 分页查询老人档案（联表 label、user 等）
// 对应 Java: ElderRecordServiceImpl.pageElderRecordByKey -> ElderRecordMapper.listElderRecordByKey
// SQL: SELECT er.*, lbl.label_name, u.name AS charge_user_name FROM elder_record er
//
//	LEFT JOIN label lbl ON lbl.id = er.label_id
//	LEFT JOIN user u ON u.id = er.charge_user_id
//	WHERE (er.elder_name LIKE %key% OR er.id = key) [可选]
//	ORDER BY er.create_time DESC; 再由 PageUtil 内存分页。
//
// Todo: 1) in.Key 非空 -> (tbl<elderrecord>.Id.Eq(in.Key) OR tbl<elderrecord>.ElderName.Like(in.Key))
//
//	2) DB 分页: Count + List(联表 LeftJoin)
//	3) 组装含标签/负责人姓名的 VO 并赋值 out
func (e *elderrecord) PageElderRecordByKey(ctx context.Context, in *dto.PageElderRecordByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 实现联表分页查询
	return nil
}

// GetElderRecordById 根据编号获取老人档案
// 对应 Java: ElderRecordServiceImpl.getElderRecordById -> elderRecordMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.ElderRecord(db).GetByID(ctx, types.BigInt(in.ID)); 另查 emergency_contact
func (e *elderrecord) GetElderRecordById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e2 := dao.ElderRecord(db).GetByID(ctx, types.BigInt(in.ID))
	if e2 != nil {
		return e2
	}
	_ = has
	_ = obj
	// todo: 另查 emergency_contact(老人编号) 获取紧急联系人
	return nil
}

// AddElderRecord 新增老人档案（含紧急联系人）
// 对应 Java: ElderRecordServiceImpl.addElderRecord -> insert elder_record + insert emergency_contact
// todo: 事务: 1) dao.ElderRecord(db).InsertOne; 2) dao.EmergencyContact(db).InsertOne(紧急联系人)
func (e *elderrecord) AddElderRecord(ctx context.Context, in *dto.AddElderRecordQuery, out *dto.EmptyResp) error {
	// todo: 写入 elder_record + emergency_contact
	return nil
}

// EditElderRecord 编辑老人档案
// 对应 Java: ElderRecordServiceImpl.editElderRecord -> updateByPrimaryKeySelective + 更新紧急联系人
// todo: 事务: 1) UpdateById; 2) 更新 emergency_contact
func (e *elderrecord) EditElderRecord(ctx context.Context, in *dto.EditElderRecordQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<elderrecord>.ElderName.Value(in.ElderName),
	}
	_, e2 := dao.ElderRecord(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	if e2 != nil {
		return e2
	}
	// todo: 更新 emergency_contact
	return nil
}

// DeleteElderRecord 删除老人档案（级联删紧急联系人）
// 对应 Java: ElderRecordServiceImpl.deleteElderRecord -> 删记录(级联删 emergency_contact)
// todo: 事务: 1) 删 emergency_contact(elder_id); 2) dao.ElderRecord(db).DeleteById
func (e *elderrecord) DeleteElderRecord(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	// todo: 先删紧急联系人再删档案
	_, e2 := dao.ElderRecord(db).DeleteById(ctx, types.BigInt(in.ID))
	return e2
}

// PageSearchElderByKey 分页搜索老人（供档案选择/关联）
// 对应 Java: ElderRecordServiceImpl.pageSearchElderByKey -> elderMapper.listElderByKey
// SQL: SELECT * FROM elder WHERE (elder_name LIKE %key% OR id = key) [可选] AND del_flag=0
// todo: 查询 elder 表并分页, 结果赋值 out(需定义老人分页返回类型)
func (e *elderrecord) PageSearchElderByKey(ctx context.Context, in *dto.PageSearchElderByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 elder 表并分页
	return nil
}

// PageSearchEmergencyContactByKey 分页搜索紧急联系人
// 对应 Java: ElderRecordServiceImpl.pageSearchEmergencyContactByKey -> emergencyContactMapper.listEmergencyContactByKey
// SQL: SELECT * FROM emergency_contact WHERE (name LIKE %key% OR id = key) [可选]
// todo: 查询 emergency_contact 表并分页, 结果赋值 out(需定义返回类型)
func (e *elderrecord) PageSearchEmergencyContactByKey(ctx context.Context, in *dto.PageSearchEmergencyContactByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 emergency_contact 表并分页
	return nil
}

// PageLabelByKey 分页查询标签
// 对应 Java: ElderRecordServiceImpl.pageLabelByKey -> labelMapper.listLabelByKey
// SQL: SELECT * FROM label WHERE (label_name LIKE %key%) [可选] ORDER BY create_time DESC
// todo: 标签分页查询 - dao.Label(db) 条件 + 分页, 结果赋值 out
func (e *elderrecord) PageLabelByKey(ctx context.Context, in *dto.PageLabelByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 label 表并分页
	return nil
}

// EditElderLabel 编辑老人标签
// 对应 Java: ElderRecordServiceImpl.editElderLabel -> 更新 elder_record.label_id
// todo: 更新 elder_record 的 label_id 字段(UpdateById)
func (e *elderrecord) EditElderLabel(ctx context.Context, in *dto.EditElderLabelQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: tbl<elderrecord>.LabelId.Value(in.LabelId),
	}
	_, e2 := dao.ElderRecord(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e2
}
