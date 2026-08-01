package service

import (
	"context"

	"api/internal/model/define/dao"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type roomtype struct{}

var RoomType = &roomtype{}

// PageRoomTypeByKey 分页查询房间类型
// 对应 Java: RoomTypeServiceImpl.pageRoomTypeByKey -> RoomTypeMapper.listRoomTypeByKey
// SQL: SELECT * FROM room_type WHERE (type_name LIKE %key%) [可选] ORDER BY create_time DESC
// todo: 房间类型分页查询 - dao.RoomType(db) 条件 + 分页, 结果赋值 out
func (r *roomtype) PageRoomTypeByKey(ctx context.Context, in *dto.PageRoomTypeByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 room_type 表并分页
	return nil
}

// GetRoomTypeById 根据编号获取房间类型
// 对应 Java: RoomTypeServiceImpl.getRoomTypeById -> roomTypeMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.RoomType(db).GetByID(ctx, types.BigInt(in.ID))
func (r *roomtype) GetRoomTypeById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.RoomType(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// AddRoomType 新增房间类型
// 对应 Java: RoomTypeServiceImpl.addRoomType -> roomTypeMapper.insertSelective
// todo: 标准 CRUD - dao.RoomType(db).InsertOne 写入 room_type 表(含 typeName/price 等)
func (r *roomtype) AddRoomType(ctx context.Context, in *dto.OperateRoomTypeQuery, out *dto.EmptyResp) error {
	// todo: bean := do.NewRoomType(); 填充 in; dao.RoomType(db).InsertOne(ctx, bean)
	return nil
}

// EditRoomType 编辑房间类型
// 对应 Java: RoomTypeServiceImpl.editRoomType -> roomTypeMapper.updateByPrimaryKeySelective
// todo: 标准 CRUD - 按主键更新 room_type 表
func (r *roomtype) EditRoomType(ctx context.Context, in *dto.OperateRoomTypeQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<roomtype>.TypeName.Value(in.TypeName),
	}
	_, e := dao.RoomType(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteRoomType 删除房间类型
// 对应 Java: RoomTypeServiceImpl.deleteRoomType -> roomTypeMapper.deleteByPrimaryKey
// todo: 标准 CRUD - dao.RoomType(db).DeleteById(ctx, types.BigInt(in.ID))
func (r *roomtype) DeleteRoomType(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.RoomType(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}
