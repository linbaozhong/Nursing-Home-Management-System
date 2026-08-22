package service

import (
	"api/internal/model/define/table/tblelder"
	"api/internal/model/define/table/tblelderlabel"
	"context"

	"api/internal/constant"
	"api/internal/lib"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tbllabel"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/types"
)

var Intention = (*intentionService)(nil)

type intentionService struct{}

// PageIntentionByKey 分页查询意向客户（elder 表中 check_flag=意向）
func (s *intentionService) PageIntentionByKey(ctx context.Context, in *dto.PageIntentionByKeyReq, out *[]dto.PageIntentionByKeyResp) error {
	q := db.Table(tblelder.TableName).
		Where(tblelder.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblelder.CheckFlag.Eq(types.Int8(constant.CheckIntention)))
	if in.ElderName != nil {
		q = q.And(tblelder.Name.Like(*in.ElderName))
	}
	if in.ElderPhone != nil {
		q = q.And(tblelder.Phone.Like(*in.ElderPhone))
	}
	var list []do.Elder
	e := q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(tblelder.Id, tblelder.Name, tblelder.Phone, tblelder.Sex, tblelder.Age).
		Desc(tblelder.CreateTime).
		Select().Gets(ctx, &list)
	if e != nil {
		return e
	}
	// 批量查询老人标签
	ids := make([]any, 0, len(list))
	for _, el := range list {
		ids = append(ids, el.Id)
	}
	labels, _, e := dao.ElderLabel(db).List(ctx, ace.Where(tblelderlabel.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblelderlabel.ElderId.In(ids...)).Cols(tblelderlabel.ElderId, tblelderlabel.LabelId))
	if e != nil {
		return e
	}
	labelMap := make(map[int64][]int64)
	for _, l := range labels {
		labelMap[int64(l.ElderId)] = append(labelMap[int64(l.ElderId)], int64(l.LabelId))
	}
	allLabelIds := make([]any, 0)
	for _, v := range labelMap {
		for _, lid := range v {
			allLabelIds = append(allLabelIds, types.BigInt(lid))
		}
	}
	labelNameMap := make(map[int64]dto.IntentionLabelResp)
	if len(allLabelIds) > 0 {
		lbs, _, e := dao.Label(db).List(ctx, ace.Where(tbllabel.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tbllabel.Id.In(allLabelIds...).And(tbllabel.DelFlag.Eq(types.Int8(constant.YesNoNo)))).Cols(tbllabel.Id, tbllabel.Name, tbllabel.Color))
		if e != nil {
			return e
		}
		for _, lb := range lbs {
			labelNameMap[int64(lb.Id)] = dto.IntentionLabelResp{Name: lb.Name.String(), Color: lb.Color.String()}
		}
	}
	res := make([]dto.PageIntentionByKeyResp, 0, len(list))
	for _, el := range list {
		vo := dto.PageIntentionByKeyResp{
			ID:    int64(el.Id),
			Name:  el.Name.String(),
			Phone: el.Phone.String(),
			Sex:   el.Sex.String(),
			Age:   int(el.Age),
		}
		for _, lid := range labelMap[int64(el.Id)] {
			if lv, ok := labelNameMap[lid]; ok {
				vo.LabelRespList = append(vo.LabelRespList, lv)
			}
		}
		res = append(res, vo)
	}
	*out = res
	return nil
}

// AddIntention 新增意向客户（写入 elder 表，check_flag=意向）
func (s *intentionService) AddIntention(ctx context.Context, in *dto.AddIntentReq, out *dto.EmptyResp) error {
	exist, has, e := dao.Elder(db).Get(ctx, ace.Where(tblelder.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblelder.IdNum.Eq(types.String(*in.IDNum))).Cols(tblelder.Id, tblelder.CheckFlag))
	if e != nil {
		return e
	}
	if !has {
		bean := new(do.Elder)
		bean.TenantId = types.BigInt(lib.TenantID(ctx))
		bean.Name = types.String(*in.Name)
		bean.IdNum = types.String(*in.IDNum)
		bean.Age = types.Int32(*in.Age)
		bean.Sex = types.String(*in.Sex)
		bean.Phone = types.String(*in.Phone)
		bean.Address = types.String(*in.Address)
		bean.Balance = types.Money(0)
		bean.CheckFlag = types.Int8(constant.CheckIntention)
		_, e = dao.Elder(db).InsertOne(ctx, bean,
			tblelder.TenantId, tblelder.Name, tblelder.IdNum, tblelder.Age, tblelder.Sex, tblelder.Phone,
			tblelder.Address, tblelder.Balance, tblelder.CheckFlag, tblelder.CreateId, tblelder.CreateTime)
		if e != nil {
			return e
		}
		return nil
	}
	// 已存在则更新
	_, e = dao.Elder(db).UpdateById(ctx, exist.Id,
		tblelder.Name.Set(types.String(*in.Name)),
		tblelder.Age.Set(types.Int32(*in.Age)),
		tblelder.Sex.Set(types.String(*in.Sex)),
		tblelder.Phone.Set(types.String(*in.Phone)),
		tblelder.Address.Set(types.String(*in.Address)),
	)
	if e != nil {
		return e
	}
	return nil
}

// GetIntentById 查询意向客户详情（含标签）
func (s *intentionService) GetIntentById(ctx context.Context, in *dto.IDReq, out *dto.OperateIntentionResp) error {
	bean, has, e := dao.Elder(db).GetByID(ctx, types.BigInt(*in.ID),
		tblelder.Id, tblelder.Name, tblelder.IdNum, tblelder.Age, tblelder.Sex, tblelder.Phone, tblelder.Address)
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrDataNotExist
	}
	*out.ID = int64(bean.Id)
	*out.Name = bean.Name.String()
	*out.IDNum = bean.IdNum.String()
	*out.Age = int(bean.Age)
	*out.Sex = bean.Sex.String()
	*out.Phone = bean.Phone.String()
	*out.Address = bean.Address.String()
	// 标签
	labels, _, e := dao.ElderLabel(db).List(ctx, ace.Where(tblelderlabel.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblelderlabel.ElderId.Eq(types.BigInt(*in.ID))).Cols(tblelderlabel.LabelId))
	if e != nil {
		return e
	}
	if len(labels) == 0 {
		return nil
	}
	ids := make([]any, 0, len(labels))
	for _, l := range labels {
		ids = append(ids, l.LabelId)
	}
	lbs, _, e := dao.Label(db).List(ctx, ace.Where(tbllabel.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tbllabel.Id.In(ids...).And(tbllabel.DelFlag.Eq(types.Int8(constant.YesNoNo)))).Cols(tbllabel.Id, tbllabel.Name, tbllabel.Color))
	if e != nil {
		return e
	}
	list := make([]dto.IntentionLabelResp, 0, len(lbs))
	for _, lb := range lbs {
		list = append(list, dto.IntentionLabelResp{Name: lb.Name.String(), Color: lb.Color.String()})
	}
	out.IntentionLabelRespList = list
	return nil
}

// EditIntention 编辑意向客户
func (s *intentionService) EditIntention(ctx context.Context, in *dto.EditIntentReq, out *dto.EmptyResp) error {
	_, has, e := dao.Elder(db).Get(ctx, ace.Where(tblelder.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblelder.Id.Eq(types.BigInt(*in.ID))))
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrDataNotExist
	}
	_, e = dao.Elder(db).UpdateById(ctx, types.BigInt(*in.ID),
		tblelder.Name.Set(types.String(*in.Name)),
		tblelder.IdNum.Set(types.String(*in.IDNum)),
		tblelder.Age.Set(types.Int32(*in.Age)),
		tblelder.Sex.Set(types.String(*in.Sex)),
		tblelder.Phone.Set(types.String(*in.Phone)),
		tblelder.Address.Set(types.String(*in.Address)),
	)
	if e != nil {
		return e
	}
	return nil
}

// DeleteIntention 删除意向客户（逻辑删除）
func (s *intentionService) DeleteIntention(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, has, e := dao.Elder(db).Get(ctx, ace.Where(tblelder.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblelder.Id.Eq(types.BigInt(*in.ID))))
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrDataNotExist
	}
	_, e = dao.Elder(db).DeleteById(ctx, types.BigInt(*in.ID))
	if e != nil {
		return e
	}
	return nil
}
