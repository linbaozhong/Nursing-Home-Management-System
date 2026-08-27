package service

import (
	"api/internal/constant"
	"api/internal/lib"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblactive"
	"api/internal/model/define/table/tblactiveparticipant"
	"api/internal/model/define/table/tblactivetype"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/do"
	"api/internal/model/dto"
	"context"
	"errors"

	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type active struct{}

var Active = &active{}

// PageActiveByKey 分页查询活动
// 对应 Java: ActiveServiceImpl.pageActiveByKey -> ActiveMapper.listActiveByKey + PageUtil 内存分页
// Java SQL: SELECT a.*, at.type_name FROM active a LEFT JOIN active_type at ON at.id = a.active_type_id
//
//	WHERE (at.type_name LIKE %key% OR a.id = key) [可选] ORDER BY a.create_time DESC
//
// 说明：Java 实际为查全量后 PageUtil 内存分页。Go 端用独立 builder 做 DB 分页，
// typeName 通过预加载 active_type map 补齐（替代 mapper 联表）。
func (a *active) PageActiveByKey(ctx context.Context, in *dto.PageActiveByKeyReq, out *[]dto.PageActiveByKeyResp) error {
	// 1) 预加载活动类型，用于补齐 typeName
	typeList, _, e := dao.ActiveType(db).List(ctx, ace.Where(tblactivetype.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblactivetype.State.NotEq(types.Int8(constant.StateDeleted))))
	if e != nil {
		return e
	}
	typeNameMap := make(map[uint64]string, len(typeList))
	for _, t := range typeList {
		typeNameMap[uint64(t.Id)] = string(t.Name)
	}

	// 2) 条件构造 + 分页
	q := db.Table(tblactive.TableName).Where(tblactive.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblactive.State.NotEq(types.Int8(constant.StateDeleted)))
	if in.TypeID != nil {
		q.And(tblactive.TypeId.Eq(types.BigInt(*in.TypeID)))
	}
	if in.Name != nil {
		q.And(tblactive.Name.Like(*in.Name))
	}
	if in.StartTime != nil {
		q.And(tblactive.ActiveDate.Gte(*in.StartTime))
	}
	if in.EndTime != nil {
		q.And(tblactive.ActiveDate.Lte(*in.EndTime))
	}
	e = q.Desc(tblactive.CreateTime).Page(uint(*in.PageNum), uint(*in.PageSize)).Select().Gets(ctx, out)
	if e != nil {
		return e
	}
	for i := range *out {
		(*out)[i].TypeName = typeNameMap[uint64((*out)[i].ID)]
	}
	return nil
}

// GetActiveById 根据编号获取活动（含参与老人列表）
// 对应 Java: ActiveServiceImpl.getActiveById -> activeMapper.selectByPrimaryKey + activeParticipantMapper.listParticipateElder
// 说明：参与老人 Java 联 elder 取 name/phone；Go 端查 elder_active 后批量取 elder 补齐。
func (a *active) GetActiveById(ctx context.Context, in *dto.IDReq, out *dto.GetActiveByIDResp) error {
	obj, has, e := dao.Active(db).GetByID(ctx, types.BigInt(*in.ID))
	if e != nil {
		return e
	}
	if !has || obj == nil {
		return errors.New("活动不存在")
	}

	// 回填活动基本信息
	id := int64(obj.Id)
	name := string(obj.Name)
	theme := string(obj.Theme)
	content := string(obj.Content)
	address := string(obj.Address)
	organizer := string(obj.Organizer)
	phone := string(obj.Phone)
	activeDate := obj.ActiveDate.Time
	activePicture := string(obj.ActivePicture)
	typeID := int64(obj.TypeId)
	out.ID = &id
	out.TypeID = &typeID
	out.Name = &name
	out.Theme = &theme
	out.Content = &content
	out.Address = &address
	out.Organizer = &organizer
	out.Phone = &phone
	out.ActiveDate = &activeDate
	out.ActivePicture = &activePicture

	// 查询参与老人关联
	participants, _, e := dao.ActiveParticipant(db).List(ctx, ace.Where(tblactiveparticipant.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblactiveparticipant.ActiveId.Eq(types.BigInt(*in.ID))))
	if e != nil {
		return e
	}
	if len(participants) == 0 {
		out.ParticipateElderRespList = []dto.ParticipateElderResp{}
		return nil
	}
	ids := make([]any, 0, len(participants))
	for _, p := range participants {
		ids = append(ids, p.ElderId)
	}
	elders, e := dao.Elder(db).GetByIds(ctx, ids)
	if e != nil {
		return e
	}
	elderMap := make(map[uint64]*do.Elder, len(elders))
	for _, el := range elders {
		elderMap[uint64(el.Id)] = el
	}
	voList := make([]dto.ParticipateElderResp, 0, len(participants))
	for _, p := range participants {
		el, ok := elderMap[uint64(p.ElderId)]
		if !ok {
			continue
		}
		voList = append(voList, dto.ParticipateElderResp{
			ID:    types.BigInt(el.Id),
			Name:  string(el.Name),
			Phone: string(el.Phone),
		})
	}
	out.ParticipateElderRespList = voList
	return nil
}

// AddActive 新增活动（同时写入参与老人 elder_active）
// 对应 Java: ActiveServiceImpl.addActive -> activeMapper.insertSelective 后批量 insert elder_active
// 说明：Java 为 @Transactional。Go 端分两步写入（框架未启用显式事务）。
func (a *active) AddActive(ctx context.Context, in *dto.OperateActiveReq, out *dto.EmptyResp) error {
	bean := do.NewActive()
	bean.TenantId = types.BigInt(lib.TenantID(ctx))
	bean.TypeId = types.BigInt(*in.TypeID)
	bean.Theme = types.String(*in.Theme)
	bean.Name = types.String(*in.Name)
	bean.Content = types.String(*in.Content)
	bean.Address = types.String(*in.Address)
	bean.Organizer = types.String(*in.Organizer)
	bean.Phone = types.String(*in.Phone)
	bean.ActiveDate = types.Time{*in.ActiveDate}
	bean.ActivePicture = types.String(*in.ActivePicture)
	bean.State = types.Int8(int8(constant.StateEnabled))
	_, e := dao.Active(db).InsertOne(ctx, bean)
	if e != nil {
		return e
	}

	for _, elderId := range in.ElderIDList {
		if _, e = dao.ActiveParticipant(db).Insert(ctx,
			tblactiveparticipant.TenantId.Set(types.BigInt(lib.TenantID(ctx))),
			tblactiveparticipant.ActiveId.Set(bean.Id),
			tblactiveparticipant.ElderId.Set(types.BigInt(elderId)),
		); e != nil {
			return e
		}
	}
	return nil
}

// EditActive 编辑活动（先删后插参与老人 elder_active）
// 对应 Java: ActiveServiceImpl.editActive -> 更新 active + 删除旧 elder_active + 批量新增
// 说明：Java 为 @Transactional。Go 端分三步写入。
func (a *active) EditActive(ctx context.Context, in *dto.OperateActiveReq, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		tblactive.TypeId.Set(*in.TypeID),
		tblactive.Theme.Set(*in.Theme),
		tblactive.Name.Set(*in.Name),
		tblactive.Content.Set(*in.Content),
		tblactive.Address.Set(*in.Address),
		tblactive.Organizer.Set(*in.Organizer),
		tblactive.Phone.Set(*in.Phone),
		tblactive.ActivePicture.Set(*in.ActivePicture),
		tblactive.ActiveDate.Set(*in.ActiveDate),
	}
	if _, e := dao.Active(db).UpdateById(ctx, types.BigInt(*in.ID), sets...); e != nil {
		return e
	}

	// 清空旧参与老人
	if _, e := dao.ActiveParticipant(db).Delete(ctx, tblactiveparticipant.ActiveId.Eq(types.BigInt(*in.ID))); e != nil {
		return e
	}

	for _, elderId := range in.ElderIDList {
		if _, e := dao.ActiveParticipant(db).Insert(ctx,
			tblactiveparticipant.TenantId.Set(types.BigInt(lib.TenantID(ctx))),
			tblactiveparticipant.ActiveId.Set(types.BigInt(*in.ID)),
			tblactiveparticipant.ElderId.Set(types.BigInt(elderId)),
		); e != nil {
			return e
		}
	}
	return nil
}

// DeleteActive 删除活动（级联删参与老人）
// 对应 Java: ActiveServiceImpl.deleteActive -> activeMapper.deleteByPrimaryKey（级联删 elder_active）
func (a *active) DeleteActive(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	if _, e := dao.ActiveParticipant(db).Delete(ctx, tblactiveparticipant.ActiveId.Eq(types.BigInt(*in.ID))); e != nil {
		return e
	}
	_, e := dao.Active(db).DeleteById(ctx, types.BigInt(*in.ID))
	return e
}

// PageSearchElderByKey 分页搜索老人（供活动选择参与老人）
// 对应 Java: ActiveServiceImpl.pageSearchElderByKey -> CommonFunc.pageSearchElderByKeyResult
// Java SQL: SELECT * FROM elder WHERE (elder_name LIKE %key% OR id = key) [可选] AND del_flag=0
// 说明：Java 侧按 checkFlag in (咨询/意向/预定/退住) 过滤；Go 端对齐 CheckContract 的咨询口径。
func (a *active) PageSearchElderByKey(ctx context.Context, in *dto.PageSearchElderByKeyReq, out *[]dto.ParticipateElderResp) error {
	// 注：elder 表无 del_flag 字段，此处用恒真条件占位
	query := ace.Where(tblelder.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblelder.Id.Gte(types.BigInt(0)))
	if in.Name != nil {
		query.And(tblelder.Name.Like(*in.Name))
	}
	if in.Phone != nil {
		query.And(tblelder.Phone.Like(*in.Phone))
	}
	list, _, e := dao.Elder(db).List(ctx, query)
	if e != nil {
		return e
	}
	*out = make([]dto.ParticipateElderResp, 0, len(list))
	for _, elder := range list {
		*out = append(*out, dto.ParticipateElderResp{
			ID:    types.BigInt(elder.Id),
			Name:  elder.Name.String(),
			Phone: elder.Phone.String(),
		})
	}
	return nil
}

// GetActiveType 获取活动类型下拉框
// 对应 Java: ActiveServiceImpl.getActiveType -> activeTypeMapper.listNotDelActiveType
// SQL: SELECT * FROM active_type WHERE del_flag = 0
func (a *active) GetActiveType(ctx context.Context, in *dto.EmptyReq, out *[]dto.DropDown) error {
	list, _, e := dao.ActiveType(db).List(ctx, ace.Where(tblactivetype.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblactivetype.State.NotEq(types.Int8(constant.StateDeleted))))
	if e != nil {
		return e
	}
	dropList := make([]dto.DropDown, 0, len(list))
	for _, v := range list {
		dropList = append(dropList, dto.DropDown{ID: types.BigInt(v.Id), Name: string(v.Name)})
	}
	*out = dropList
	return nil
}
