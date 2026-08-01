package service

import (
	"context"

	"api/internal/model/define/dao"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

type consume struct{}

var Consume = &consume{}

// PageConsumeByKey 分页查询消费记录（联表 elder 获取老人姓名等）
// 对应 Java: ConsumeServiceImpl.pageConsumeByKey -> ConsumeMapper.listConsumeByKey
// SQL: SELECT c.*, e.elder_name FROM consume c
//
//	LEFT JOIN elder e ON e.id = c.elder_id
//	WHERE (c.consume_type = #{consumeType}) [可选]
//	  AND (e.elder_name LIKE %key% OR c.id = key) [可选]
//	ORDER BY c.create_time DESC; 再由 PageUtil 内存分页。
//
// todo: 1) in.ConsumeType 非空 -> tblconsume.ConsumeType.Eq(in.ConsumeType)
//
//	2) in.Key 非空 -> (tblconsume.Id.Eq(in.Key) OR tbl<elder>.ElderName.Like(in.Key))
//	3) DB 分页: 先 dao.Consume(db).Count(cond...) 再 List(Limit/Offset)
//	4) 组装含老人姓名的 VO 并赋值 out(需定义分页返回类型, 当前为 EmptyResp)
func (c *consume) PageConsumeByKey(ctx context.Context, in *dto.PageConsumeByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 实现联表分页查询
	return nil
}

// GetConsumeById 根据编号获取消费记录
// 对应 Java: 无单独接口(page 结果已携带), 此处为前端补充的详情接口。
// todo: 实现标准 CRUD - 按主键查询 consume 表
func (c *consume) GetConsumeById(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	obj, has, e := dao.Consume(db).GetByID(ctx, types.BigInt(in.ID))
	if e != nil {
		return e
	}
	if !has {
		// todo: 记录不存在, 视业务返回错误或空
		return nil
	}
	_ = obj
	return nil
}

// AddConsume 新增消费
// 对应 Java: 无(前端补充的 CRUD)。入参 AddConsumeQuery 含 elderId/consumeType/consumeAmount/consumeDate。
// todo: 实现标准 CRUD - 将 in 字段写入 consume 表(用 dao.Consume(db).InsertOne)
func (c *consume) AddConsume(ctx context.Context, in *dto.AddConsumeQuery, out *dto.EmptyResp) error {
	// todo: bean := do.NewConsume(); 填充 in 字段; dao.Consume(db).InsertOne(ctx, bean)
	return nil
}

// EditConsume 编辑消费
// 对应 Java: 无(前端补充的 CRUD)。
// todo: 实现标准 CRUD - 按主键更新 consume 表字段(tblconsume.X.Value(in.X))
func (c *consume) EditConsume(ctx context.Context, in *dto.EditConsumeQuery, out *dto.EmptyResp) error {
	sets := []dialect.Setter{
		// todo: 例 tblconsume.ConsumeAmount.Value(in.ConsumeAmount),
		//       tblconsume.ConsumeDate.Value(in.ConsumeDate),
	}
	_, e := dao.Consume(db).UpdateById(ctx, types.BigInt(in.ID), sets...)
	return e
}

// DeleteConsume 删除消费
// 对应 Java: 无(前端补充的 CRUD)。
// todo: 实现标准 CRUD - 按主键删除 consume 记录
func (c *consume) DeleteConsume(ctx context.Context, in *dto.IDReq, out *dto.EmptyResp) error {
	_, e := dao.Consume(db).DeleteById(ctx, types.BigInt(in.ID))
	return e
}
