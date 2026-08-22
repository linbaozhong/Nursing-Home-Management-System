package service

import (
	"context"
	"errors"
	"github.com/linbaozhong/gentity/pkg/ace"
	"strconv"
	"time"

	"api/internal/constant"
	"api/internal/lib"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblconsult"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/define/table/tblsource"
	"api/internal/model/define/table/tblstaff"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/conv"
	"github.com/linbaozhong/gentity/pkg/types"
)

type consult struct{}

var Consult = &consult{}

// PageConsultByKey 分页查询咨询（联表老人、接待人、来源渠道）
// 对应 Java: ConsultServiceImpl.pageConsultByKey -> ConsultMapper.listConsultByKey
func (c *consult) PageConsultByKey(ctx context.Context, in *dto.PageConsultByKeyReq, out *[]dto.PageConsultByKeyResp) error {
	q := db.Table(tblconsult.TableName).
		LeftJoin(tblconsult.ElderId, tblelder.Id).
		LeftJoin(tblconsult.StaffId, tblstaff.Id).
		LeftJoin(tblconsult.SourceId, tblsource.Id).
		Where(tblconsult.TenantId.Eq(types.BigInt(lib.TenantID(ctx))))
	if in.ConsultName != nil && *in.ConsultName != "" {
		q.And(tblconsult.Name.Like(*in.ConsultName))
	}
	if in.ConsultPhone != nil && *in.ConsultPhone != "" {
		q.And(tblconsult.Phone.Like(*in.ConsultPhone))
	}
	if in.ElderName != nil && *in.ElderName != "" {
		q.And(tblelder.Name.Like(*in.ElderName))
	}
	if in.ElderPhone != nil && *in.ElderPhone != "" {
		q.And(tblelder.Phone.Like(*in.ElderPhone))
	}
	if in.StartTime != nil && !in.StartTime.IsZero() {
		q.And(tblconsult.ConsultDate.Gte(*in.StartTime))
	}
	if in.EndTime != nil && !in.EndTime.IsZero() {
		q.And(tblconsult.ConsultDate.Lte(*in.EndTime))
	}
	if in.SourceID != nil {
		q.And(tblconsult.SourceId.Eq(types.BigInt(*in.SourceID)))
	}
	if in.StaffID != nil {
		q.And(tblconsult.StaffId.Eq(types.BigInt(*in.StaffID)))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblconsult.Id.AsName("consult_id"),
			tblconsult.Name.AsName("consult_name"),
			tblconsult.Phone.AsName("consult_phone"),
			tblconsult.Relation,
			tblconsult.ConsultContent,
			tblconsult.ElderId.AsName("elder_id"),
			tblconsult.SourceId,
			tblconsult.StaffId,
			tblconsult.ConsultDate.AsName("consult_date"),
			tblelder.Name.AsName("elder_name"),
			tblelder.Phone.AsName("elder_phone"),
			tblelder.Sex,
			tblelder.Age,
			tblsource.Name.AsName("source_name"),
			tblstaff.Name.AsName("staff_name"),
		).
		Desc(tblconsult.CreateTime).
		Select().
		Gets(ctx, out)
}

// GetConsultByConsultIdAndElderId 根据咨询编号与老人编号获取唯一咨询（用于编辑回显）
// 对应 Java: ConsultServiceImpl.getConsultByConsultIdAndElderId
func (c *consult) GetConsultByConsultIdAndElderId(ctx context.Context, in *dto.GetConsultByConsultIdAndElderIdQuery, out *dto.GetConsultByConsultIDAndElderIDResp) error {
	obj, has, e := dao.Consult(db).Get(ctx, ace.Where(
		tblconsult.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
		tblconsult.Id.Eq(types.BigInt(*in.ConsultID)),
		tblconsult.ElderId.Eq(types.BigInt(*in.ElderID)),
	))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("咨询记录不存在")
	}
	out.ConsultID = conv.Ptr(int64(obj.Id))
	out.ElderID = conv.Ptr(int64(obj.ElderId))
	out.SourceID = conv.Ptr(int64(obj.SourceId))
	out.StaffID = conv.Ptr(int64(obj.StaffId))
	out.ConsultName = conv.Ptr(obj.Name.String())
	out.ConsultPhone = conv.Ptr(obj.Phone.String())
	out.Relation = conv.Ptr(obj.Relation.String())
	out.ConsultDate = &obj.ConsultDate.Time
	out.ConsultContent = conv.Ptr(obj.ConsultContent.String())
	return nil
}

// AddConsult 新增咨询（先校验身份证号，再新增老人与咨询记录）
// 对应 Java: ConsultServiceImpl.addConsult
func (c *consult) AddConsult(ctx context.Context, in *dto.AddConsultReq, out *dto.EmptyResp) error {
	// 校验身份证号是否已存在（排除已删除老人）
	repeat, e := dao.Elder(db).Exists(ctx,
		tblelder.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
		tblelder.IdNum.Eq(*in.IDNum),
		tblelder.CheckFlag.NotEq(types.Int8(constant.CheckFailure)),
	)
	if e != nil {
		return e
	}
	if repeat {
		return errors.New("身份证号已存在")
	}
	// 新增老人（状态：咨询）
	elder := do.NewElder()
	elder.TenantId = types.BigInt(lib.TenantID(ctx))
	elder.Name = types.String(*in.ElderName)
	elder.IdNum = types.String(*in.IDNum)
	elder.Age = types.Int32(parseInt32(*in.Age))
	elder.Sex = types.String(*in.Sex)
	elder.Phone = types.String(*in.ElderPhone)
	elder.Address = types.String(*in.Address)
	elder.Balance = types.Money(0)
	elder.CheckFlag = types.Int8(constant.CheckConsult)
	_, e = dao.Elder(db).InsertOne(ctx, elder)
	if e != nil {
		return e
	}
	// 新增咨询记录
	bean := do.NewConsult()
	bean.TenantId = types.BigInt(lib.TenantID(ctx))
	bean.ElderId = types.BigInt(elder.Id)
	bean.Name = types.String(*in.ConsultName)
	bean.Phone = types.String(*in.ConsultPhone)
	bean.Relation = types.String(*in.Relation)
	bean.ConsultContent = types.String(*in.ConsultContent)
	bean.SourceId = types.BigInt(*in.SourceID)
	bean.StaffId = types.BigInt(*in.StaffID)
	bean.ConsultDate = types.Time{*in.ConsultDate}
	_, e = dao.Consult(db).InsertOne(ctx, bean)
	return e
}

// PageSearchElderByKey 分页搜索老人（供咨询选择老人）
// 对应 Java: ConsultServiceImpl.pageSearchElderByKey -> ElderFunc.listElderByKey
func (c *consult) PageSearchElderByKey(ctx context.Context, in *dto.PageSearchElderByKeyReq, out *[]dto.PageSearchElderByKeyResp) error {
	q := db.Table(tblelder.TableName).
		Where(tblelder.TenantId.Eq(types.BigInt(lib.TenantID(ctx))))
	if in.Name != nil {
		q.And(tblelder.Name.Like(*in.Name))
	}
	if *in.Phone != "" {
		q.Or(tblelder.Phone.Like(*in.Phone))
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

// PageIntentByKey 分页查询意向客户（CheckFlag=INTENTION）
// 对应 Java: ConsultServiceImpl.pageIntentByKey -> intentionMapper.listIntentByKey
func (c *consult) PageIntentByKey(ctx context.Context, in *dto.PageIntentByKeyReq, out *[]dto.PageIntentionByKeyResp) error {
	q := db.Table(tblelder.TableName).
		Where(
			tblelder.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
			tblelder.CheckFlag.Eq(types.Int8(constant.CheckIntention)),
		)
	if in.Key != nil {
		q.AndOr(tblelder.Name.Like(*in.Key), tblelder.Phone.Like(*in.Key))
	}

	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblelder.Id,
			tblelder.Name,
			tblelder.Phone,
			tblelder.Sex,
			tblelder.Age,
		).
		Desc(tblelder.CreateTime).
		Select().
		Gets(ctx, out)
}

// ---- 辅助函数 ----

func parseInt32(s string) int32 {
	if s == "" {
		return 0
	}
	n, _ := strconv.ParseInt(s, 10, 32)
	return int32(n)
}

func timeParse(layout, value string) (time.Time, error) {
	return time.ParseInLocation(layout, value, time.Local)
}

func parseDate(s *string) types.Time {
	if s == nil || *s == "" {
		return types.Time{}
	}
	for _, layout := range []string{"2006-01-02 15:04:05", "2006-01-02"} {
		if t, err := timeParse(layout, *s); err == nil {
			return types.Time{t}
		}
	}
	return types.Time{}
}

func parseDateStart(s *string) time.Time {
	if s == nil || *s == "" {
		return time.Time{}
	}
	if t, err := timeParse("2006-01-02", *s); err == nil {
		return t
	}
	return time.Time{}
}

func parseDateEnd(s *string) time.Time {
	if s == nil || *s == "" {
		return time.Time{}
	}
	if t, err := timeParse("2006-01-02", *s); err == nil {
		// 加一天减一秒，覆盖当天
		return t.Add(24*time.Hour - time.Second)
	}
	return time.Time{}
}
