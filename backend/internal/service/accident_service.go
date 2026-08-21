package service

import (
	"api/internal/constant"
	"api/internal/lib"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblaccident"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/define/table/tblstaff"
	"api/internal/model/do"
	"api/internal/model/dto"
	"context"
	"errors"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type accident struct{}

var Accident = &accident{}

// PageAccidentByKey 分页查询事故记录（联表 elder 获取老人姓名）
func (a *accident) PageAccidentByKey(ctx context.Context, in *dto.PageAccidentByKeyQuery, out *[]dto.PageAccidentByKeyVO) error {
	clampPage(in.PageNum, in.PageSize)
	// 构造查询
	q := db.Table(tblaccident.TableName).
		LeftJoin(tblaccident.ElderId, tblelder.Id).
		LeftJoin(tblaccident.StaffId, tblstaff.Id).
		Where(
			tblaccident.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
			tblaccident.DelFlag.Eq(constant.YesNoNo),
		).
		Cols(
			tblaccident.Id,
			tblelder.Name.AsName("elder_name"),
			tblstaff.Name.AsName("staff_name"),
			tblaccident.OccurDate,
		).Desc(tblaccident.OccurDate)

	if in.ElderName != nil {
		q.And(tblelder.Name.Like(*in.ElderName))
	}
	if in.StaffName != nil {
		q.And(tblstaff.Name.Like(*in.StaffName))
	}
	// 查询事故列表
	e := q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Select().
		Gets(ctx, out)
	if e != nil {
		return e
	}

	return nil
}

// GetAccidentById 根据编号获取事故
func (a *accident) GetAccidentById(ctx context.Context, in *dto.IDReq, out *dto.GetAccidentByIDVO) error {
	return db.Table(tblaccident.TableName).
		LeftJoin(tblaccident.ElderId, tblelder.Id).
		Where(tblaccident.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblaccident.Id.Eq(*in.ID)).
		Cols(
			tblaccident.Id,
			tblelder.Name.AsName("elder_name"),
			tblaccident.StaffId,
			tblaccident.OccurDate,
			tblaccident.Description,
			tblaccident.Picture,
		).
		Select().
		Get(ctx, out)
}

// AddAccident 新增事故
func (a *accident) AddAccident(ctx context.Context, in *dto.AddAccidentQuery, out *dto.EmptyResp) error {
	// 初始化事故登记
	bean := do.NewAccident()
	bean.TenantId = types.BigInt(lib.TenantID(ctx))
	if in.ElderID != nil {
		bean.ElderId = types.BigInt(*in.ElderID)
	}
	if in.StaffID != nil {
		bean.StaffId = types.BigInt(*in.StaffID)
	}
	if in.OccurDate != nil {
		bean.OccurDate = types.Time{*in.OccurDate}
	}
	if in.Description != nil {
		bean.Description = types.String(*in.Description)
	}
	if in.Picture != nil {
		bean.Picture = types.String(*in.Picture)
	}
	bean.DelFlag = types.Int8(constant.YesNoNo)
	// 新增
	_, e := dao.Accident(db).InsertOne(ctx, bean)
	return e
}

// EditAccident 编辑事故
func (a *accident) EditAccident(ctx context.Context, in *dto.EditAccidentQuery, out *dto.EmptyResp) error {
	obj, has, e := dao.Accident(db).GetByID(ctx, types.BigInt(*in.ID),
		tblaccident.StaffId,
		tblaccident.OccurDate,
		tblaccident.Description,
		tblaccident.Picture,
	)
	if e != nil {
		return e
	}
	if !has {
		return errors.New("事故不存在")
	}
	//
	var sets = make([]dialect.Setter, 0, 4)
	if in.StaffID != nil && obj.StaffId.Int64() != *in.StaffID {
		sets = append(sets, tblaccident.StaffId.Set(*in.StaffID))
	}
	if in.OccurDate != nil && !obj.OccurDate.Time.Equal(*in.OccurDate) {
		sets = append(sets, tblaccident.OccurDate.Set(*in.OccurDate))
	}
	if in.Description != nil && obj.Description.String() != *in.Description {
		sets = append(sets, tblaccident.Description.Set(*in.Description))
	}
	if in.Picture != nil && obj.Picture.String() != *in.Picture {
		sets = append(sets, tblaccident.Picture.Set(*in.Picture))
	}
	_, e = dao.Accident(db).UpdateById(ctx, types.BigInt(*in.ID), sets...)
	return e
}

// DeleteAccident 删除事故
func (a *accident) DeleteAccident(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.Accident(db).UpdateById(ctx, types.BigInt(*in.ID),
		tblaccident.DelFlag.Set(constant.YesNoYes),
	)
	return e
}
