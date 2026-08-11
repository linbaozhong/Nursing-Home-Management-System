package service

import (
	"context"
	"errors"

	"api/internal/constant"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblbed"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/define/table/tblemergencycontact"
	"api/internal/model/define/table/tbllabel"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type elderrecord struct{}

var ElderRecord = &elderrecord{}

// PageElderRecordByKey 分页查询长者档案（联表 bed 取床位名称）
// 对应 Java: ElderRecordServiceImpl.pageElderByKey -> ElderMapper.listElderByKey
func (e *elderrecord) PageElderRecordByKey(ctx context.Context, in *dto.PageElderRecordByKeyQuery, out *[]dto.PageElderByKeyVO) error {
	q := db.Table(tblelder.TableName).
		LeftJoin(tblelder.BedId, tblbed.Id).
		Where(tblelder.Id.Gt(types.BigInt(0)))
	if in.ElderName != nil && *in.ElderName != "" {
		q.And(tblelder.Name.Like(*in.ElderName))
	}
	if in.IDNum != nil && *in.IDNum != "" {
		q.And(tblelder.IdNum.Like(*in.IDNum))
	}
	if in.ElderSex != nil && *in.ElderSex != "" {
		q.And(tblelder.Sex.Eq(*in.ElderSex))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblelder.Id,
			tblbed.Name.AsName("bed_name"),
			tblelder.Name,
			tblelder.IdNum,
			tblelder.Age,
			tblelder.Sex,
			tblelder.Phone,
			tblelder.Address,
			tblelder.CheckFlag,
		).
		Desc(tblelder.CreateTime).
		Select().
		Gets(ctx, out)
}

// GetElderRecordById 根据编号获取长者档案（含紧急联系人）
// 对应 Java: ElderRecordServiceImpl.getElderRecordById
func (e *elderrecord) GetElderRecordById(ctx context.Context, in *dto.IDReq, out *dto.GetElderRecordByIDVO) error {
	elder, has, e2 := dao.Elder(db).GetByID(ctx, types.BigInt(*in.ID),
		tblelder.Id, tblelder.Name, tblelder.IdNum, tblelder.Age, tblelder.Sex,
		tblelder.Phone, tblelder.Address,
	)
	if e2 != nil {
		return e2
	}
	if !has {
		return errors.New("老人不存在")
	}
	out.Name = elder.Name.String()
	out.IDNum = elder.IdNum.String()
	out.Age = int(elder.Age)
	out.Sex = elder.Sex.String()
	out.Phone = elder.Phone.String()
	out.Address = elder.Address.String()

	// 紧急联系人
	var contacts []do.EmergencyContact
	e2 = db.Table(tblemergencycontact.TableName).
		Cols(
			tblemergencycontact.Id,
			tblemergencycontact.Name,
			tblemergencycontact.Phone,
			tblemergencycontact.Relation,
			tblemergencycontact.Remark,
		).
		Where(tblemergencycontact.ElderId.Eq(types.BigInt(*in.ID))).
		Select().
		Gets(ctx, &contacts)
	if e2 != nil {
		return e2
	}
	out.ElderEmergencyContactByIDVOList = make([]dto.OperateEmergencyContactQuery, 0, len(contacts))
	for _, ct := range contacts {
		out.ElderEmergencyContactByIDVOList = append(out.ElderEmergencyContactByIDVOList, dto.OperateEmergencyContactQuery{
			ID:       int64Ptr(int64(ct.Id)),
			Name:     strPtr(ct.Name.String()),
			Phone:    strPtr(ct.Phone.String()),
			Relation: strPtr(ct.Relation.String()),
			Remark:   strPtr(ct.Remark.String()),
		})
	}
	return nil
}

// AddElderRecord 新增长者档案（含紧急联系人）
// 对应 Java 不存在, Go handler 注册, 实现基础 elder 插入 + 紧急联系人
func (e *elderrecord) AddElderRecord(ctx context.Context, in *dto.AddElderRecordQuery, out *dto.EmptyResp) error {
	bean := do.NewElder()
	bean.Name = types.String(*in.Name)
	bean.IdNum = types.String(*in.IDNum)
	bean.Sex = types.String(*in.Sex)
	bean.Age = types.Int32(int32(*in.Age))
	bean.Phone = types.String(*in.Phone)
	bean.Address = types.String(*in.Address)
	bean.CheckFlag = types.Int8(constant.CheckConsult)
	_, e2 := dao.Elder(db).InsertOne(ctx, bean)
	if e2 != nil {
		return e2
	}
	if len(in.EmergencyContactQueryList) > 0 {
		list := make([]*do.EmergencyContact, 0, len(in.EmergencyContactQueryList))
		for _, ec := range in.EmergencyContactQueryList {
			c := do.NewEmergencyContact()
			c.ElderId = types.BigInt(bean.Id)
			c.Name = types.String(*ec.Name)
			c.Phone = types.String(*ec.Phone)
			c.Relation = types.String(*ec.Relation)
			c.Remark = types.String(*ec.Remark)
			list = append(list, c)
		}
		_, e2 = dao.EmergencyContact(db).InsertBatch(ctx, list...)
	}
	return e2
}

// EditElderRecord 编辑长者档案
func (e *elderrecord) EditElderRecord(ctx context.Context, in *dto.EditElderRecordQuery, out *dto.EmptyResp) error {
	var sets = make([]dialect.Setter, 0, 7)
	sets = append(sets, tblelder.Name.Set(*in.Name))
	sets = append(sets, tblelder.IdNum.Set(*in.IDNum))
	sets = append(sets, tblelder.Sex.Set(*in.Sex))
	sets = append(sets, tblelder.Age.Set(*in.Age))
	sets = append(sets, tblelder.Phone.Set(*in.Phone))
	sets = append(sets, tblelder.Address.Set(*in.Address))
	if in.NurseLevel != nil {
		sets = append(sets, tblelder.NurseGradeId.Set(*in.NurseLevel))
	}
	if in.CheckFlag != nil {
		sets = append(sets, tblelder.CheckFlag.Set(*in.CheckFlag))
	}
	_, e2 := dao.Elder(db).UpdateById(ctx, types.BigInt(*in.ID), sets...)
	return e2
}

// DeleteElderRecord 删除长者档案（级联删紧急联系人 + 逻辑删除）
func (e *elderrecord) DeleteElderRecord(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e2 := dao.EmergencyContact(db).Delete(ctx, tblemergencycontact.ElderId.Eq(types.BigInt(*in.ID)))
	if e2 != nil {
		return e2
	}
	_, e2 = dao.Elder(db).DeleteById(ctx, types.BigInt(*in.ID))
	return e2
}

// PageSearchElderByKey 分页搜索老人
// 对应 Java: ElderRecordServiceImpl.pageSearchElderByKey -> ElderMapper.listElderByKey
func (e *elderrecord) PageSearchElderByKey(ctx context.Context, in *dto.PageSearchElderByKeyQuery, out *[]dto.PageSearchElderByKeyVO) error {
	q := db.Table(tblelder.TableName).
		Where(tblelder.Id.Gt(types.BigInt(0)))
	if in.Name != nil && *in.Name != "" {
		q.And(tblelder.Name.Like(*in.Name))
	}
	if in.Phone != nil && *in.Phone != "" {
		q.And(tblelder.Phone.Like(*in.Phone))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblelder.Id,
			tblelder.Name,
			tblelder.IdNum,
			tblelder.Sex,
			tblelder.Phone,
			tblelder.Address,
			tblelder.CheckFlag,
		).
		Desc(tblelder.CreateTime).
		Select().
		Gets(ctx, out)
}

// PageSearchEmergencyContactByKey 分页搜索紧急联系人
func (e *elderrecord) PageSearchEmergencyContactByKey(ctx context.Context, in *dto.PageSearchEmergencyContactByKeyQuery, out *[]dto.PageSearchEmergencyContactByKeyVO) error {
	q := db.Table(tblemergencycontact.TableName).
		Where(tblemergencycontact.ElderId.Eq(types.BigInt(*in.ElderID)))
	if in.Key != nil && *in.Key != "" {
		q.And(tblemergencycontact.Name.Like(*in.Key))
		q.Or(tblemergencycontact.Phone.Like(*in.Key))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblemergencycontact.Id,
			tblemergencycontact.ElderId,
			tblemergencycontact.Name,
			tblemergencycontact.Phone,
			tblemergencycontact.Relation,
			tblemergencycontact.Remark,
		).
		Desc(tblemergencycontact.CreateTime).
		Select().
		Gets(ctx, out)
}

// PageLabelByKey 分页查询客户标签
func (e *elderrecord) PageLabelByKey(ctx context.Context, in *dto.PageLabelByKeyQuery, out *[]dto.ListLabelVO) error {
	q := db.Table(tbllabel.TableName).
		Where(tbllabel.DelFlag.Eq(types.Int8(constant.YesNoNo)))
	if in.Key != nil && *in.Key != "" {
		q.And(tbllabel.Name.Like(*in.Key))
	}
	var labels []do.Label
	e2 := q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(tbllabel.Id, tbllabel.Name, tbllabel.Color).
		Desc(tbllabel.CreateTime).
		Select().
		Gets(ctx, &labels)
	if e2 != nil {
		return e2
	}
	*out = make([]dto.ListLabelVO, 0, len(labels))
	for _, l := range labels {
		*out = append(*out, dto.ListLabelVO{
			ID:   int64(l.Id),
			Name: l.Name.String(),
		})
	}
	return nil
}

// EditElderLabel 编辑老人标签
// 说明: Go 端 elder 表无 label_id 字段, 该方法保留以便后续扩展
func (e *elderrecord) EditElderLabel(ctx context.Context, in *dto.EditElderLabelQuery, out *dto.EmptyResp) error {
	return errors.New("not implemented: elder 表暂无 label_id 字段")
}

// ---- 辅助函数（int64Ptr/strPtr 已在 util.go 定义）----
