package service

import (
	"context"
	"errors"

	"api/internal/model/define/dao"
	"api/internal/model/do"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type intention struct{}

var Intention = &intention{}

// PageIntentByKey 分页查询意向客户
// 对应 Java: IntentionServiceImpl.pageIntentByKey
// SQL: SELECT * FROM elder WHERE check_flag = 'INTENTION' [AND name LIKE %elderName%] [AND phone LIKE %elderPhone%] [AND id IN (按 labelId 筛选)]
// todo: 条件 + 分页, 结果赋值 out
func (i *intention) PageIntentByKey(ctx context.Context, in *dto.PageIntentionByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 elder(check_flag=INTENTION) 并分页
	return nil
}

// GetIntentById 根据编号获取意向客户
// 对应 Java: IntentionServiceImpl.getIntentionById
// SQL: SELECT * FROM elder WHERE id=? AND check_flag='INTENTION'
// 注意: 意向客户在数据库中复用 elder 表(check_flag='INTENTION'), 无独立 Intention 表
func (i *intention) GetIntentById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	tbl := dao.Elder(db)
	exist, has, e := tbl.GetByID(ctx, types.BigInt(*in.ID))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("老人不存在")
	}
	if exist.CheckFlag == "FAILURE" {
		return errors.New("老人已删除")
	}
	return nil
}

// AddIntent 新增意向客户
// 对应 Java: IntentionServiceImpl.addIntention
// 注意: 无 idNum 时新增 elder(check_flag='INTENTION'); 有 idNum 且为 FAILURE 时复用该记录
// 注意: 意向客户在数据库中复用 elder 表, 无独立 Intention 表
func (i *intention) AddIntent(ctx context.Context, in *dto.AddIntentQuery, out *dto.EmptyResp) error {
	tbl := dao.Elder(db)
	if in.IDNum == nil {
		return errors.New("身份证号不能为空")
	}
	// 查询相同身份证号的记录(排除已逻辑删除的)
	list, _, e := tbl.List(ctx, ace.Where(tblelder.IdNum.Eq(*in.IDNum)))
	if e != nil {
		return e
	}
	for _, v := range list {
		if v.CheckFlag != "FAILURE" {
			return errors.New("身份证号重复")
		}
	}
	bean := &do.Elder{
		Name:      types.String(*in.Name),
		IdNum:     types.String(*in.IDNum),
		Age:       types.Int32(int32(*in.Age)),
		Sex:       types.String(*in.Sex),
		Phone:     types.String(*in.Phone),
		Address:   types.String(*in.Address),
		CheckFlag: types.String("INTENTION"),
		Balance:   types.Float64(0),
	}
	_, e = tbl.InsertOne(ctx, bean)
	if e != nil {
		return e
	}
	return nil
}

// EditIntent 编辑意向客户
// 对应 Java: IntentionServiceImpl.editIntention -> 更新 elder(check_flag='INTENTION')
// 注意: 意向客户在数据库中复用 elder 表, 无独立 Intention 表
func (i *intention) EditIntent(ctx context.Context, in *dto.EditIntentQuery, out *dto.EmptyResp) error {
	tbl := dao.Elder(db)
	if in.ID == nil || *in.ID == 0 {
		return errors.New("老人编号不能为空")
	}
	exist, has, e := tbl.GetByID(ctx, types.BigInt(*in.ID))
	if e != nil {
		return e
	}
	if !has {
		return errors.New("老人不存在")
	}
	if exist.CheckFlag == "FAILURE" {
		return errors.New("老人已删除")
	}
	sets := []dialect.Setter{}
	if in.Name != nil {
		sets = append(sets, tblelder.Name.Set(*in.Name))
	}
	if in.IDNum != nil {
		sets = append(sets, tblelder.IdNum.Set(*in.IDNum))
	}
	if in.Age != nil {
		sets = append(sets, tblelder.Age.Set(int32(*in.Age)))
	}
	if in.Sex != nil {
		sets = append(sets, tblelder.Sex.Set(*in.Sex))
	}
	if in.Phone != nil {
		sets = append(sets, tblelder.Phone.Set(*in.Phone))
	}
	if in.Address != nil {
		sets = append(sets, tblelder.Address.Set(*in.Address))
	}
	sets = append(sets, tblelder.CheckFlag.Set("INTENTION"))
	_, e = tbl.UpdateById(ctx, types.BigInt(*in.ID), sets...)
	if e != nil {
		return e
	}
	return nil
}

// ListLabel 标签下拉列表
// 对应 Java: IntentionServiceImpl.listLabel -> labelFunc.listNotDelLabel
// todo: 查询 label(未删除)列表, 结果赋值 out(需定义返回类型)
func (i *intention) ListLabel(ctx context.Context, in *dto.ListLabelQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 label 下拉列表
	return nil
}

// GetElderLabelById 根据老人编号获取其标签
// 对应 Java: IntentionServiceImpl.getElderLabelById -> labelMapper.listElderLabelById
// todo: 联 elder_label + label 拼 labelVoList, 结果赋值 out
func (i *intention) GetElderLabelById(ctx context.Context, in *dto.GetElderLabelByIdQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询老人标签列表
	return nil
}

// ============ 以下方法(VisitPlan / CommunicationRecord)沿用 CheckContract 已实现的逻辑 ============
// Java 端这些方法归属 Intention, Go 端统一在 CheckContract 实现, 此处委托以避免重复与编译错误。

// PageVisitPlan 分页查询回访计划（按是否完成区分）
// 对应 Java: IntentionServiceImpl / CheckContractServiceImpl
func (i *intention) PageVisitPlan(ctx context.Context, in *dto.PageVisitPlanQuery, out *dto.EmptyResp) error {
	return CheckContract.PageVisitPlan(ctx, in, out)
}

// AddVisitPlan 新增回访计划
// 对应 Java: IntentionServiceImpl.addVisitPlan
func (i *intention) AddVisitPlan(ctx context.Context, in *dto.AddVisitPlanQuery, out *dto.EmptyResp) error {
	return CheckContract.AddVisitPlan(ctx, in, out)
}

// PageCommunicationRecord 分页查询沟通记录
// 对应 Java: IntentionServiceImpl.pageCommunicationRecord
func (i *intention) PageCommunicationRecord(ctx context.Context, in *dto.PageCommunicationRecordQuery, out *dto.EmptyResp) error {
	return CheckContract.PageCommunicationRecord(ctx, in, out)
}

// AddCommunicationRecord 新增沟通记录
// 对应 Java: IntentionServiceImpl.addCommunicationRecord
func (i *intention) AddCommunicationRecord(ctx context.Context, in *dto.AddCommunicationRecordQuery, out *dto.EmptyResp) error {
	return CheckContract.AddCommunicationRecord(ctx, in, out)
}

// EditCommunicationRecord 编辑沟通记录
// 对应 Java: IntentionServiceImpl.editCommunicationRecord
func (i *intention) EditCommunicationRecord(ctx context.Context, in *dto.EditCommunicationRecordQuery, out *dto.EmptyResp) error {
	return CheckContract.EditCommunicationRecord(ctx, in, out)
}

// DeleteCommunicationRecord 删除沟通记录（逻辑删除）
// 对应 Java: IntentionServiceImpl.deleteCommunicationRecord
func (i *intention) DeleteCommunicationRecord(ctx context.Context, in *dto.DeleteCommunicationRecordQuery, out *dto.EmptyResp) error {
	return CheckContract.DeleteCommunicationRecord(ctx, in, out)
}
