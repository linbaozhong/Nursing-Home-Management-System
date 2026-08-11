package service

import (
	"context"

	"api/internal/constant"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tbelder"
	"api/internal/model/define/table/tbelderlabel"
	"api/internal/model/define/table/tbllabel"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/types"
)

var _ = (*intentionService)(nil)

type intentionService struct{}

// PageIntentionByKey 分页查询意向客户（elder 表中 check_flag=意向）
func (s *intentionService) PageIntentionByKey(ctx context.Context, in *dto.PageIntentionByKeyQuery, out *[]dto.PageIntentionByKeyVO) error {
	q := db.Table(tblelder.TableName).
		Where(ace.Where(tblelder.CheckFlag.Eq(types.Int8(constant.CheckIntention))))
	if in.ElderName != nil {
		q = q.And(tblelder.Name.Like(*in.ElderName))
	}
	if in.ElderPhone != nil {
		q = q.And(tblelder.Phone.Like(*in.ElderPhone))
	}
	var list []do.Elder
	has, e := q.Page(in.PageNum, in.PageSize).
		Cols(tblelder.Id, tblelder.Name, tblelder.Phone, tblelder.Sex, tblelder.Age).
		OrderBy(tblelder.CreateTime, false).
		Select().Gets(ctx, &list)
	if e != nil {
		return e
	}
	if !has {
		return nil
	}
	// 批量查询老人标签
	ids := make([]any, 0, len(list))
	for _, el := range list {
		ids = append(ids, el.Id)
	}
	labels, _, e := dao.ElderLabel(db).List(ctx, ace.Where(tbelderlabel.ElderId.In(ids...)).Cols(tbelderlabel.ElderId, tbelderlabel.LabelId))
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
	labelNameMap := make(map[int64]dto.IntentionLabelVO)
	if len(allLabelIds) > 0 {
		lbs, _, e := dao.Label(db).List(ctx, ace.Where(tbllabel.Id.In(allLabelIds...).And(tbllabel.DelFlag.Eq(types.Int8(constant.YesNoNo)))).Cols(tbllabel.Id, tbllabel.Name, tbllabel.Color))
		if e != nil {
			return e
		}
		for _, lb := range lbs {
			labelNameMap[int64(lb.Id)] = dto.IntentionLabelVO{Name: lb.Name.String(), Color: lb.Color.String()}
		}
	}
	res := make([]dto.PageIntentionByKeyVO, 0, len(list))
	for _, el := range list {
		vo := dto.PageIntentionByKeyVO{
			ID:    int64(el.Id),
			Name:  el.Name.String(),
			Phone: el.Phone.String(),
			Sex:   el.Sex.String(),
			Age:   int(el.Age),
		}
		for _, lid := range labelMap[int64(el.Id)] {
			if lv, ok := labelNameMap[lid]; ok {
				vo.LabelVOList = append(vo.LabelVOList, lv)
			}
		}
		res = append(res, vo)
	}
	*out = res
	return nil
}

// AddIntention 新增意向客户（写入 elder 表，check_flag=意向）
func (s *intentionService) AddIntention(ctx context.Context, in *dto.AddIntentQuery) (*dto.EmptyResp, error) {
	exist, has, e := dao.Elder(db).Get(ctx, ace.Where(tblelder.IdNum.Eq(types.String(*in.IDNum))).Cols(tblelder.Id, tblelder.CheckFlag))
	if e != nil {
		return nil, e
	}
	if !has {
		bean := new(do.Elder)
		bean.Name = types.String(*in.Name)
		bean.IdNum = types.String(*in.IDNum)
		bean.Age = types.Int32(*in.Age)
		bean.Sex = types.String(*in.Sex)
		bean.Phone = types.String(*in.Phone)
		bean.Address = types.String(*in.Address)
		bean.Balance = types.Float64(0)
		bean.CheckFlag = types.Int8(constant.CheckIntention)
		_, e = dao.Elder(db).InsertOne(ctx, bean,
			tblelder.Name, tblelder.IdNum, tblelder.Age, tblelder.Sex, tblelder.Phone,
			tblelder.Address, tblelder.Balance, tblelder.CheckFlag, tblelder.CreateId, tblelder.CreateTime)
		if e != nil {
			return nil, e
		}
		return new(dto.EmptyResp), nil
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
		return nil, e
	}
	return new(dto.EmptyResp), nil
}

// GetIntentById 查询意向客户详情（含标签）
func (s *intentionService) GetIntentById(ctx context.Context, in *dto.IDReq, out *dto.OperateIntentionVO) error {
	bean, has, e := dao.Elder(db).GetByID(ctx, types.BigInt(*in.ID),
		tblelder.Id, tblelder.Name, tblelder.IdNum, tblelder.Age, tblelder.Sex, tblelder.Phone, tblelder.Address)
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrDataNotExist
	}
	out.ID = int64(bean.Id)
	out.Name = bean.Name.String()
	out.IDNum = bean.IdNum.String()
	out.Age = int(bean.Age)
	out.Sex = bean.Sex.String()
	out.Phone = bean.Phone.String()
	out.Address = bean.Address.String()
	// 标签
	labels, _, e := dao.ElderLabel(db).List(ctx, ace.Where(tbelderlabel.ElderId.Eq(types.BigInt(*in.ID))).Cols(tbelderlabel.LabelId))
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
	lbs, _, e := dao.Label(db).List(ctx, ace.Where(tbllabel.Id.In(ids...).And(tbllabel.DelFlag.Eq(types.Int8(constant.YesNoNo)))).Cols(tbllabel.Id, tbllabel.Name, tbllabel.Color))
	if e != nil {
		return e
	}
	list := make([]dto.IntentionLabelVO, 0, len(lbs))
	for _, lb := range lbs {
		list = append(list, dto.IntentionLabelVO{Name: lb.Name.String(), Color: lb.Color.String()})
	}
	out.IntentionLabelVOList = list
	return nil
}

// EditIntention 编辑意向客户
func (s *intentionService) EditIntention(ctx context.Context, in *dto.EditIntentQuery) (*dto.EmptyResp, error) {
	_, e := dao.Elder(db).UpdateById(ctx, types.BigInt(*in.ID),
		tblelder.Name.Set(types.String(*in.Name)),
		tblelder.IdNum.Set(types.String(*in.IDNum)),
		tblelder.Age.Set(types.Int32(*in.Age)),
		tblelder.Sex.Set(types.String(*in.Sex)),
		tblelder.Phone.Set(types.String(*in.Phone)),
		tblelder.Address.Set(types.String(*in.Address)),
	)
	if e != nil {
		return nil, e
	}
	return new(dto.EmptyResp), nil
}

// DeleteIntention 删除意向客户（逻辑删除）
func (s *intentionService) DeleteIntention(ctx context.Context, in *dto.IDReq) (*dto.EmptyResp, error) {
	_, e := dao.Elder(db).DeleteById(ctx, types.BigInt(*in.ID))
	if e != nil {
		return nil, e
	}
	return new(dto.EmptyResp), nil
}
