package service

import (
	"api/internal/model/define/table/tblaccident"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/define/table/tblstaff"
	"api/internal/model/do"
	"context"

	"api/internal/model/define/dao"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type accident struct{}

var Accident = &accident{}

// PageAccidentByKey 分页查询事故记录（联表 elder 获取老人姓名）
// 对应 Java: AccidentServiceImpl.pageAccidentByKey -> AccidentMapper.listAccidentByKey
// SQL: SELECT a.*, e.elder_name FROM accident a LEFT JOIN elder e ON e.id = a.elder_id
//
//	WHERE (e.elder_name LIKE %key% OR a.id = key) [可选]
//	ORDER BY a.create_time DESC; 再由 PageUtil 内存分页。
//
// todo: 1) in.Key 非空 -> (tbl<accident>.Id.Eq(in.Key) OR tbl<elder>.ElderName.Like(in.Key))
//
//  2. DB 分页: Count + List(联表 LeftJoin)
//  3. 组装含老人姓名的 VO 并赋值 out
func (a *accident) PageAccidentByKey(ctx context.Context, in *dto.PageAccidentByKeyQuery, out *dto.PageAccidentByKeyVO) error {
	// todo: 实现联表分页查询
	db.Table(do.AccidentTableName).
		LeftJoin(tblaccident.ElderId, tblelder.Id).
		LeftJoin(tblaccident.StaffId, tblstaff.Id).
		Cols(
			tblaccident.Id,
			tblelder.Name.AsName("elder_name"),
			tblstaff.Name.AsName("staff_name"),
			tblaccident.OccurDate,
		).Select().Slices(ctx)
	return nil
}

// GetAccidentById 根据编号获取事故
// 对应 Java: AccidentServiceImpl.getAccidentById -> accidentMapper.selectByPrimaryKey
// todo: 标准 CRUD - dao.Accident(db).GetByID(ctx, types.BigInt(in.ID))
func (a *accident) GetAccidentById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.Accident(db).GetByID(ctx, types.BigInt(*in.ID))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	return nil
}

// AddAccident 新增事故
// 对应 Java: AccidentServiceImpl.addAccident -> accidentMapper.insertSelective
// todo: 标准 CRUD - dao.Accident(db).InsertOne 写入 accident 表(含 elderId/accidentType/remark 等)
func (a *accident) AddAccident(ctx context.Context, in *dto.AddAccidentQuery, out *dto.EmptyResp) error {
	// todo: bean := do.NewAccident(); 填充 in; dao.Accident(db).InsertOne(ctx, bean)
	return nil
}

// EditAccident 编辑事故
// 对应 Java: AccidentServiceImpl.editAccident -> accidentMapper.updateByPrimaryKeySelective
// todo: 标准 CRUD - 按主键更新 accident 表
func (a *accident) EditAccident(ctx context.Context, in *dto.EditAccidentQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tbl<accident>.Remark.Value(in.Remark),
	}
	_, e := dao.Accident(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteAccident 删除事故
// 对应 Java: AccidentServiceImpl.deleteAccident -> accidentMapper.deleteByPrimaryKey
// todo: 标准 CRUD - dao.Accident(db).DeleteById(ctx, types.BigInt(in.ID))
func (a *accident) DeleteAccident(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.Accident(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}
