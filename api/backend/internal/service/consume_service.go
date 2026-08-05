package service

import (
	"context"
	"errors"
	"time"

	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblconsume"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type consume struct{}

var Consume = &consume{}

// PageConsumeByKey 分页查询消费记录（联表 elder 获取老人姓名/身份证）
// 对应 Java: ConsumeServiceImpl.pageConsumeByKey
func (c *consume) PageConsumeByKey(ctx context.Context, in *dto.PageConsumeByKeyQuery, out *[]dto.PageConsumeByKeyVO) error {
	q := db.Table(do.ConsumeTableName).
		LeftJoin(tblconsume.ElderId, tblelder.Id).
		Where(tblconsume.DelFlag.Eq(0))
	if in.ElderName != nil && *in.ElderName != "" {
		q.And(tblelder.Name.Like(*in.ElderName))
	}
	if in.StartTime != nil && *in.StartTime != "" {
		if t, err := time.ParseInLocation("2006-01-02", *in.StartTime, time.Local); err == nil {
			q.And(tblconsume.ConsumeDate.Gte(types.Time(t)))
		}
	}
	if in.EndTime != nil && *in.EndTime != "" {
		if t, err := time.ParseInLocation("2006-01-02", *in.EndTime, time.Local); err == nil {
			q.And(tblconsume.ConsumeDate.Lte(types.Time(t.Add(24*time.Hour - time.Second))))
		}
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblconsume.Id,
			tblelder.Name.AsName("elder_name"),
			tblelder.IdNum.AsName("id_num"),
			tblconsume.ConsumeType,
			tblconsume.ConsumeAmount,
			tblconsume.ConsumeDate,
		).
		Desc(tblconsume.ConsumeDate).
		Select().
		Gets(ctx, out)
}

// GetConsumeById 根据编号获取消费记录
func (c *consume) GetConsumeById(ctx context.Context, in *dto.IDReq, out *dto.GetConsumeByIdVO) error {
	obj, has, e := dao.Consume(db).GetByID(ctx, types.BigInt(*in.ID),
		tblconsume.Id,
		tblconsume.ElderId,
		tblconsume.ConsumeType,
		tblconsume.ConsumeAmount,
		tblconsume.ConsumeDate,
		tblconsume.Remark,
	)
	if e != nil {
		return e
	}
	if !has {
		return errors.New("消费记录不存在")
	}
	out.ID = int64(obj.Id)
	out.ElderID = int64(obj.ElderId)
	out.ConsumeType = obj.ConsumeType.String()
	out.ConsumeAmount = obj.ConsumeAmount.Float64()
	out.ConsumeDate = time.Time(obj.ConsumeDate)
	out.Remark = obj.Remark.String()
	return nil
}

// AddConsume 新增消费
func (c *consume) AddConsume(ctx context.Context, in *dto.AddConsumeQuery, out *dto.EmptyResp) error {
	bean := do.NewConsume()
	bean.ElderId = types.BigInt(*in.ElderID)
	bean.ConsumeType = types.String(*in.ConsumeType)
	bean.ConsumeAmount = types.Float64(*in.ConsumeAmount)
	bean.ConsumeDate = types.Time(parseTimeStr(*in.ConsumeDate))
	_, e := dao.Consume(db).InsertOne(ctx, bean)
	return e
}

// EditConsume 编辑消费
func (c *consume) EditConsume(ctx context.Context, in *dto.EditConsumeQuery, out *dto.EmptyResp) error {
	var sets = make([]dialect.Setter, 0, 4)
	sets = append(sets, tblconsume.ElderId.Set(*in.ElderID))
	sets = append(sets, tblconsume.ConsumeType.Set(*in.ConsumeType))
	sets = append(sets, tblconsume.ConsumeAmount.Set(*in.ConsumeAmount))
	sets = append(sets, tblconsume.ConsumeDate.Set(parseTimeStr(*in.ConsumeDate)))
	if in.Remark != nil {
		sets = append(sets, tblconsume.Remark.Set(*in.Remark))
	}
	_, e := dao.Consume(db).UpdateById(ctx, types.BigInt(*in.ID), sets...)
	return e
}

// DeleteConsume 删除消费
func (c *consume) DeleteConsume(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.Consume(db).DeleteById(ctx, types.BigInt(*in.ID))
	return e
}

// parseTimeStr 解析消费日期字符串（支持两种格式）
func parseTimeStr(s *string) types.Time {
	if s == nil || *s == "" {
		return types.Time{}
	}
	for _, layout := range []string{"2006-01-02 15:04:05", "2006-01-02"} {
		if t, err := time.ParseInLocation(layout, *s, time.Local); err == nil {
			return types.Time(t)
		}
	}
	return types.Time{}
}
