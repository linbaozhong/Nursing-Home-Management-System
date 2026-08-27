package service

import (
	"context"
	"errors"

	"api/internal/constant"
	"api/internal/lib"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblroomtype"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/types"
)

type roomtype struct{}

var RoomType = &roomtype{}

// PageRoomTypeByKey 分页查询房间类型（按名称模糊过滤 + 未删除）
// 对应 Java: RoomTypeServiceImpl.pageRoomTypeByKey -> RoomTypeFunc.listNotDelRoomType
func (r *roomtype) PageRoomTypeByKey(ctx context.Context, in *dto.PageRoomTypeByKeyReq, out *[]dto.PageRoomTypeByKeyResp) error {
	q := db.Table(tblroomtype.TableName).
		Where(tblroomtype.TenantId.Eq(types.BigInt(lib.TenantID(ctx))), tblroomtype.State.NotEq(types.Int8(constant.StateDeleted)))
	if in.RoomTypeName != nil && *in.RoomTypeName != "" {
		q.And(tblroomtype.Name.Like(*in.RoomTypeName))
	}
	return q.Page(uint(*in.PageNum), uint(*in.PageSize)).
		Cols(
			tblroomtype.Id.AsName("id"),
			tblroomtype.Name.AsName("name"),
			tblroomtype.MonthPrice.AsName("month_price"),
		).
		Desc(tblroomtype.CreateTime).
		Select().
		Gets(ctx, out)
}

// GetRoomTypeById 根据编号获取房间类型（编辑回显）
// 对应 Java: RoomTypeServiceImpl.getRoomTypeById -> OperateRoomTypeResp
func (r *roomtype) GetRoomTypeById(ctx context.Context, in *dto.IDReq, out *dto.OperateRoomTypeResp) error {
	obj, has, e := dao.RoomType(db).GetByID(ctx, types.BigInt(*in.ID))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("房间类型不存在")
	}
	*out.ID = int64(obj.Id)
	*out.Name = obj.Name.String()
	*out.MonthPrice = obj.MonthPrice
	return nil
}

// AddRoomType 新增房间类型（校验名称不重复）
// 对应 Java: RoomTypeServiceImpl.addRoomType -> RoomTypeFunc.getRoomTypeByName
func (r *roomtype) AddRoomType(ctx context.Context, in *dto.OperateRoomTypeReq, out *dto.EmptyResp) error {
	repeat, e := dao.RoomType(db).Exists(ctx,
		tblroomtype.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
		tblroomtype.Name.Eq(*in.Name),
		tblroomtype.State.NotEq(types.Int8(constant.StateDeleted)),
	)
	if e != nil {
		return e
	}
	if repeat {
		return errors.New("房间类型名称已存在")
	}
	bean := do.NewRoomType()
	bean.TenantId = types.BigInt(lib.TenantID(ctx))
	bean.Name = types.String(*in.Name)
	bean.MonthPrice = types.Money(*in.MonthPrice)
	bean.State = types.Int8(constant.StateEnabled)
	_, e = dao.RoomType(db).InsertOne(ctx, bean)
	return e
}

// EditRoomType 编辑房间类型（校验名称不重复排除自身）
// 对应 Java: RoomTypeServiceImpl.editRoomType
func (r *roomtype) EditRoomType(ctx context.Context, in *dto.OperateRoomTypeReq, out *dto.EmptyResp) error {
	repeat, e := dao.RoomType(db).Exists(ctx,
		tblroomtype.TenantId.Eq(types.BigInt(lib.TenantID(ctx))),
		tblroomtype.Name.Eq(*in.Name),
		tblroomtype.State.NotEq(types.Int8(constant.StateDeleted)),
		tblroomtype.Id.NotEq(types.BigInt(*in.ID)),
	)
	if e != nil {
		return e
	}
	if repeat {
		return errors.New("房间类型名称已存在")
	}
	bean := do.NewRoomType()
	bean.Id = types.BigInt(*in.ID)
	bean.Name = types.String(*in.Name)
	bean.MonthPrice = types.Money(*in.MonthPrice)
	_, e = dao.RoomType(db).UpdateOne(ctx, bean)
	return e
}

// DeleteRoomType 删除房间类型（逻辑删除）
// 对应 Java: RoomTypeServiceImpl.deleteRoomType -> del_flag = YES
func (r *roomtype) DeleteRoomType(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	bean := do.NewRoomType()
	bean.Id = types.BigInt(*in.ID)
	bean.State = types.Int8(constant.StateDeleted)
	_, e := dao.RoomType(db).UpdateOne(ctx, bean)
	return e
}
