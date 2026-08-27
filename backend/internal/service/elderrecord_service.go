package service

import (
	"context"
	"errors"
	"time"

	"api/internal/constant"
	"api/internal/lib"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblbed"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/define/table/tblemergencycontact"
	"api/internal/model/define/table/tbllabel"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/conv"
	"github.com/linbaozhong/gentity/pkg/types"
)

type elderrecord struct{}

var ElderRecord = &elderrecord{}

// PageElderRecordByKey 分页查询长者档案（联表 bed 取床位名称）
// 对应 Java: ElderRecordServiceImpl.pageElderByKey -> ElderMapper.listElderByKey
func (e *elderrecord) PageElderRecordByKey(ctx context.Context, in *dto.PageElderRecordByKeyReq, out *[]dto.PageElderByKeyResp) error {
	clampPage(in.PageNum, in.PageSize)
	q := db.Table(tblelder.TableName).
		LeftJoin(tblelder.BedId, tblbed.Id).
		Where(tblelder.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblelder.Id.Gt(types.BigInt(0)))
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
			tblelder.Status,
		).
		Desc(tblelder.CreateTime).
		Select().
		Gets(ctx, out)
}

// GetElderRecordById 根据编号获取长者档案（含紧急联系人）
// 对应 Java: ElderRecordServiceImpl.getElderRecordById
func (e *elderrecord) GetElderRecordById(ctx context.Context, in *dto.IDReq, out *dto.GetElderRecordByIDResp) error {
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
			tblemergencycontact.Email,
			tblemergencycontact.Relation,
			tblemergencycontact.Status,
		).
		Where(tblemergencycontact.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblemergencycontact.ElderId.Eq(types.BigInt(*in.ID))).
		Select().
		Gets(ctx, &contacts)
	if e2 != nil {
		return e2
	}
	out.ElderEmergencyContactByIDRespList = make([]dto.OperateEmergencyContactReq, 0, len(contacts))
	for _, ct := range contacts {
		out.ElderEmergencyContactByIDRespList = append(out.ElderEmergencyContactByIDRespList, dto.OperateEmergencyContactReq{
			Name:        conv.Ptr(ct.Name.String()),
			Phone:       conv.Ptr(ct.Phone.String()),
			Email:       conv.Ptr(ct.Email.String()),
			Relation:    conv.Ptr(ct.Relation.String()),
			ReceiveFlag: conv.Ptr(ct.Status.Int8()),
		})
	}
	return nil
}

// AddElderRecord 新增长者档案（含紧急联系人）
// 对应 Java 不存在, Go handler 注册, 实现基础 elder 插入 + 紧急联系人
func (e *elderrecord) AddElderRecord(ctx context.Context, in *dto.AddElderRecordReq, out *dto.EmptyResp) error {
	bean := do.NewElder()
	bean.TenantId = types.BigInt(lib.TenantID(ctx))
	bean.Name = types.String(*in.Name)
	bean.IdNum = types.String(*in.IDNum)
	bean.Sex = types.String(*in.Sex)
	bean.Age = types.Int32(int32(*in.Age))
	bean.Phone = types.String(*in.Phone)
	bean.Address = types.String(*in.Address)
	bean.Status = types.Int8(constant.ElderCheckEntered)
	_, e2 := dao.Elder(db).InsertOne(ctx, bean)
	if e2 != nil {
		return e2
	}
	// 紧急联系人通过独立接口 addEmergencyContact 维护，新增老人时不在此批量写入
	return nil
}

// EditElderRecord 编辑长者档案
func (e *elderrecord) EditElderRecord(ctx context.Context, in *dto.EditElderRecordReq, out *dto.EmptyResp) error {
	var sets = make([]dialect.Setter, 0, 7)
	sets = append(sets, tblelder.Name.Set(*in.Name))
	sets = append(sets, tblelder.IdNum.Set(*in.IDNum))
	sets = append(sets, tblelder.Sex.Set(*in.Sex))
	sets = append(sets, tblelder.Age.Set(*in.Age))
	sets = append(sets, tblelder.Phone.Set(*in.Phone))
	sets = append(sets, tblelder.Address.Set(*in.Address))
	if in.NurseLevel != nil {
		sets = append(sets, tblelder.NursingGradeId.Set(*in.NurseLevel))
	}
	if in.CheckFlag != nil {
		sets = append(sets, tblelder.Status.Set(*in.CheckFlag))
	}
	_, e2 := dao.Elder(db).UpdateById(ctx, types.BigInt(*in.ID), sets...)
	return e2
}

// AddEmergencyContact 新增紧急联系人
func (e *elderrecord) AddEmergencyContact(ctx context.Context, in *dto.AddEmergencyContactReq, out *dto.EmptyResp) error {
	bean := do.NewEmergencyContact()
	bean.TenantId = types.BigInt(lib.TenantID(ctx))
	bean.ElderId = types.BigInt(*in.ElderID)
	bean.Name = types.String(*in.Name)
	bean.Phone = types.String(*in.Phone)
	if in.Relation != nil {
		bean.Relation = types.String(*in.Relation)
	}
	bean.CreateId = types.BigInt(ctxUserID(ctx))
	bean.UpdateId = types.BigInt(ctxUserID(ctx))
	_, e2 := dao.EmergencyContact(db).InsertOne(ctx, bean)
	return e2
}

// EditEmergencyContact 编辑紧急联系人
func (e *elderrecord) EditEmergencyContact(ctx context.Context, in *dto.EditEmergencyContactReq, out *dto.EmptyResp) error {
	id := types.BigInt(*in.ID)
	sets := make([]dialect.Setter, 0, 4)
	if in.Name != nil {
		sets = append(sets, tblemergencycontact.Name.Set(types.String(*in.Name)))
	}
	if in.Phone != nil {
		sets = append(sets, tblemergencycontact.Phone.Set(types.String(*in.Phone)))
	}
	if in.Relation != nil {
		sets = append(sets, tblemergencycontact.Relation.Set(types.String(*in.Relation)))
	}
	if len(sets) == 0 {
		return nil
	}
	sets = append(sets, tblemergencycontact.UpdateId.Set(types.BigInt(ctxUserID(ctx))))
	sets = append(sets, tblemergencycontact.UpdateTime.Set(types.Time{Time: time.Now()}))
	_, e2 := dao.EmergencyContact(db).UpdateById(ctx, id, sets...)
	return e2
}

// DeleteEmergencyContact 删除紧急联系人
func (e *elderrecord) DeleteEmergencyContact(ctx context.Context, in *dto.DeleteEmergencyContactReq, out *dto.EmptyResp) error {
	_, e2 := dao.EmergencyContact(db).DeleteById(ctx, types.BigInt(*in.ID))
	return e2
}

// ctxUserID 取当前登录用户编号（兜底 0）
func ctxUserID(ctx context.Context) int64 {
	if c, ok := ctx.(interface{ UserID() int64 }); ok {
		return c.UserID()
	}
	return 0
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
func (e *elderrecord) PageSearchElderByKey(ctx context.Context, in *dto.PageSearchElderByKeyReq, out *[]dto.PageSearchElderByKeyResp) error {
	clampPage(in.PageNum, in.PageSize)
	q := db.Table(tblelder.TableName).
		Where(tblelder.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblelder.Id.Gt(types.BigInt(0)))
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
			tblelder.Status,
		).
		Desc(tblelder.CreateTime).
		Select().
		Gets(ctx, out)
}

// PageSearchEmergencyContactByKey 分页搜索紧急联系人
func (e *elderrecord) PageSearchEmergencyContactByKey(ctx context.Context, in *dto.PageSearchEmergencyContactByKeyReq, out *[]dto.PageSearchEmergencyContactByKeyResp) error {
	clampPage(in.PageNum, in.PageSize)
	q := db.Table(tblemergencycontact.TableName).
		Where(tblemergencycontact.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblemergencycontact.ElderId.Eq(types.BigInt(*in.ElderID)))
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
		).
		Desc(tblemergencycontact.CreateTime).
		Select().
		Gets(ctx, out)
}

// PageLabelByKey 分页查询客户标签
func (e *elderrecord) PageLabelByKey(ctx context.Context, in *dto.PageLabelByKeyReq, out *[]dto.ListLabelResp) error {
	clampPage(in.PageNum, in.PageSize)
	q := db.Table(tbllabel.TableName).
		Where(tbllabel.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tbllabel.State.NotEq(types.Int8(constant.StateDeleted)))
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
	*out = make([]dto.ListLabelResp, 0, len(labels))
	for _, l := range labels {
		*out = append(*out, dto.ListLabelResp{
			ID:   types.BigInt(l.Id),
			Name: l.Name.String(),
		})
	}
	return nil
}

// EditElderLabel 编辑老人标签
// 说明: Go 端 elder 表无 label_id 字段, 该方法保留以便后续扩展
func (e *elderrecord) EditElderLabel(ctx context.Context, in *dto.EditElderLabelReq, out *dto.EmptyResp) error {
	return errors.New("not implemented: elder 表暂无 label_id 字段")
}

// ---- 辅助函数（int64Ptr/strPtr 已在 util.go 定义）----
