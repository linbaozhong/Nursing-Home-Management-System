package service

import (
	"api/internal/constant"
	"api/internal/lib"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblaccident"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/define/table/tblstaff"
	"api/internal/model/dto"
	"context"
	"errors"
	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
	"strings"
)

type accident struct{}

var Accident = &accident{}

// PageAccidentByKey 分页查询事故记录（联表 elder 获取老人姓名）
func (a *accident) PageAccidentByKey(ctx context.Context, in *dto.PageAccidentByKeyReq, out *[]dto.PageAccidentByKeyResp) error {
	clampPage(in.PageNum, in.PageSize)
	// 构造查询
	q := db.Table(tblaccident.TableName).
		LeftJoin(tblaccident.ElderId, tblelder.Id).
		LeftJoin(tblaccident.StaffId, tblstaff.Id).
		Where(
			tblaccident.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
			tblaccident.State.NotEq(types.Int8(constant.StateDeleted)),
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
func (a *accident) GetAccidentById(ctx context.Context, in *dto.IDReq, out *dto.GetAccidentByIDResp) error {
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
			tblaccident.Severity,
			tblaccident.HandleResult,
			tblaccident.HandleStaffId,
		).
		Select().
		Get(ctx, out)
}

// AddAccident 新增事故
func (a *accident) AddAccident(ctx context.Context, in *dto.AddAccidentReq, out *dto.EmptyResp) error {
	tenantID := types.BigInt(lib.TenantID(ctx))
	if in.ElderID == nil || in.StaffID == nil || in.OccurDate == nil || in.Description == nil {
		return constant.ErrParamError
	}

	// 事务：新增事故 + 写审计日志（同一事务，保证原子）
	_, e := db.Transaction(ctx, func(tx *ace.Tx) (any, error) {
		sets := []dialect.Setter{
			tblaccident.TenantId.Set(tenantID),
			tblaccident.ElderId.Set(types.BigInt(*in.ElderID)),
			tblaccident.StaffId.Set(types.BigInt(*in.StaffID)),
			tblaccident.OccurDate.Set(types.Time{*in.OccurDate}),
			tblaccident.Description.Set(types.String(*in.Description)),
			tblaccident.Picture.Set(types.String(ptrStr(in.Picture))),
			tblaccident.State.Set(types.Int8(constant.StateEnabled)),
		}
		if in.Severity != nil {
			sets = append(sets, tblaccident.Severity.Set(types.Int8(*in.Severity)))
		}
		if in.HandleResult != nil {
			sets = append(sets, tblaccident.HandleResult.Set(types.String(*in.HandleResult)))
		}
		if in.HandleStaffID != nil {
			sets = append(sets, tblaccident.HandleStaffId.Set(types.BigInt(*in.HandleStaffID)))
		}
		id, e := dao.Accident(tx).Insert(ctx, sets...)
		if e != nil {
			return 0, e
		}
		after := map[string]any{
			"id":              id,
			"elder_id":        *in.ElderID,
			"staff_id":        *in.StaffID,
			"occur_date":      in.OccurDate.Format("2006-01-02 15:04:05"),
			"description":     *in.Description,
			"picture":         ptrStr(in.Picture),
			"severity":        ptrI8v(in.Severity),
			"handle_result":   ptrStr(in.HandleResult),
			"handle_staff_id": ptrI64v(in.HandleStaffID),
		}
		e = WriteAuditLog(ctx, tx, tblaccident.TableName, id, constant.AuditCreate,
			"新增事故，发生时间："+in.OccurDate.Format("2006-01-02"), "", after)
		if e != nil {
			return 0, e
		}
		return id, nil
	})
	return e
}

// EditAccident 编辑事故
func (a *accident) EditAccident(ctx context.Context, in *dto.EditAccidentReq, out *dto.EmptyResp) error {
	obj, has, e := dao.Accident(db).GetByID(ctx, types.BigInt(*in.ID),
		tblaccident.StaffId,
		tblaccident.OccurDate,
		tblaccident.Description,
		tblaccident.Picture,
		tblaccident.Severity,
		tblaccident.HandleResult,
		tblaccident.HandleStaffId,
	)
	if e != nil {
		return e
	}
	if !has {
		return errors.New("事故不存在")
	}
	//
	var sets = make([]dialect.Setter, 0, 7)
	var parts = make([]string, 0, 7)
	after := map[string]any{
		"id":              int64(*in.ID),
		"staff_id":        obj.StaffId.Int64(),
		"occur_date":      obj.OccurDate.Time.Format("2006-01-02 15:04:05"),
		"description":     obj.Description.String(),
		"picture":         obj.Picture.String(),
		"severity":        obj.Severity.Int8(),
		"handle_result":   obj.HandleResult.String(),
		"handle_staff_id": obj.HandleStaffId.Int64(),
	}
	if in.StaffID != nil && obj.StaffId.Int64() != *in.StaffID {
		sets = append(sets, tblaccident.StaffId.Set(*in.StaffID))
		after["staff_id"] = *in.StaffID
		parts = append(parts, AuditFieldLabel(ctx, "accident", "staff_id")+"已修改")
	}
	if in.OccurDate != nil && !obj.OccurDate.Time.Equal(*in.OccurDate) {
		sets = append(sets, tblaccident.OccurDate.Set(*in.OccurDate))
		after["occur_date"] = in.OccurDate.Format("2006-01-02 15:04:05")
		parts = append(parts, AuditFieldLabel(ctx, "accident", "occur_date")+"已修改")
	}
	if in.Description != nil && obj.Description.String() != *in.Description {
		sets = append(sets, tblaccident.Description.Set(*in.Description))
		after["description"] = *in.Description
		parts = append(parts, AuditFieldLabel(ctx, "accident", "description")+"已修改")
	}
	if in.Picture != nil && obj.Picture.String() != *in.Picture {
		sets = append(sets, tblaccident.Picture.Set(*in.Picture))
		after["picture"] = *in.Picture
		parts = append(parts, AuditFieldLabel(ctx, "accident", "picture")+"已修改")
	}
	if in.Severity != nil && obj.Severity.Int8() != *in.Severity {
		sets = append(sets, tblaccident.Severity.Set(*in.Severity))
		after["severity"] = *in.Severity
		parts = append(parts, AuditFieldLabel(ctx, "accident", "severity")+"已修改")
	}
	if in.HandleResult != nil && obj.HandleResult.String() != *in.HandleResult {
		sets = append(sets, tblaccident.HandleResult.Set(*in.HandleResult))
		after["handle_result"] = *in.HandleResult
		parts = append(parts, AuditFieldLabel(ctx, "accident", "handle_result")+"已修改")
	}
	if in.HandleStaffID != nil && obj.HandleStaffId.Int64() != *in.HandleStaffID {
		sets = append(sets, tblaccident.HandleStaffId.Set(*in.HandleStaffID))
		after["handle_staff_id"] = *in.HandleStaffID
		parts = append(parts, AuditFieldLabel(ctx, "accident", "handle_staff_id")+"已修改")
	}
	if len(sets) == 0 {
		return nil
	}

	// 事务：编辑事故 + 写审计日志（同一事务，保证原子）
	_, e = db.Transaction(ctx, func(tx *ace.Tx) (any, error) {
		if _, e := dao.Accident(tx).UpdateById(ctx, types.BigInt(*in.ID), sets...); e != nil {
			return 0, e
		}
		label := "更新事故：" + strings.Join(parts, "；")
		if e := WriteAuditLog(ctx, tx, tblaccident.TableName, *in.ID, constant.AuditUpdate, label, "", after); e != nil {
			return 0, e
		}
		return nil, nil
	})
	return e
}

// DeleteAccident 删除事故
func (a *accident) DeleteAccident(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	// 事务：逻辑删除事故 + 写审计日志（同一事务，保证原子）
	_, e := db.Transaction(ctx, func(tx *ace.Tx) (any, error) {
		if _, e := dao.Accident(tx).UpdateById(ctx, types.BigInt(*in.ID),
			tblaccident.State.Set(constant.StateDeleted),
		); e != nil {
			return 0, e
		}
		if e := WriteAuditLog(ctx, tx, tblaccident.TableName, *in.ID, constant.AuditDelete, "删除事故", "", map[string]any{"id": int64(*in.ID)}); e != nil {
			return 0, e
		}
		return nil, nil
	})
	return e
}
