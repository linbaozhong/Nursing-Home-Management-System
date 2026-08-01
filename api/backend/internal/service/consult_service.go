package service

import (
	"context"

	"api/internal/model/define/dao"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/ace/dialect"
)

type consult struct{}

var Consult = &consult{}

// PageConsultByKey 分页查询咨询记录（联表 elder、user 获取老人姓名与咨询人姓名）
// 对应 Java: ConsultServiceImpl.pageConsultByKey -> ConsultMapper.listConsultByKey
// SQL: SELECT c.*, e.elder_name, u.name AS consult_user_name
//
//	FROM consult c
//	LEFT JOIN elder e ON e.id = c.elder_id
//	LEFT JOIN user u ON u.id = c.consult_user_id
//	WHERE (e.elder_name LIKE %key% OR c.id = key) [可选]
//	ORDER BY c.create_time DESC; 再由 PageUtil 内存分页。
//
// todo: 1) in.Key 非空 -> (tblconsult.Id.Eq(in.Key) OR tbl<elder>.ElderName.Like(in.Key))
//
//	2) DB 分页: Count + List(联表用 LeftJoin)
//	3) 组装含老人姓名/咨询人姓名的 VO 并赋值 out
func (c *consult) PageConsultByKey(ctx context.Context, in *dto.PageConsultByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 实现联表分页查询
	return nil
}

// GetConsultByConsultIdAndElderId 根据咨询编号与老人编号获取唯一咨询（用于编辑回显）
// 对应 Java: ConsultServiceImpl.getConsultByConsultIdAndElderId
// SQL: SELECT * FROM consult WHERE id = #{consultId} AND elder_id = #{elderId}
// todo: 标准查询 - dao.Consult(db).Get(ctx, ace.Where(tblconsult.Id.Eq(in.ConsultId), tblconsult.ElderId.Eq(in.ElderId)))
func (c *consult) GetConsultByConsultIdAndElderId(ctx context.Context, in *dto.GetConsultByConsultIdAndElderIdQuery, out *dto.EmptyResp) error {
	obj, has, e := dao.Consult(db).Get(ctx, ace.Where(
	// todo: tblconsult.Id.Eq(in.ConsultId), tblconsult.ElderId.Eq(in.ElderId),
	))
	if e != nil {
		return e
	}
	_ = has
	_ = obj
	// 注意: ace.Where 内条件需在 todo 中补全, 当前为空条件会查全表
	return nil
}

// AddConsult 新增咨询
// 对应 Java: ConsultServiceImpl.addConsult -> consultMapper.insertSelective 后回填 id
// todo: 标准 CRUD - dao.Consult(db).InsertOne 写入 consult 表(含 elderId/consultUserId/consultContent 等)
func (c *consult) AddConsult(ctx context.Context, in *dto.AddConsultQuery, out *dto.EmptyResp) error {
	// todo: bean := do.NewConsult(); 填充 in; dao.Consult(db).InsertOne(ctx, bean)
	return nil
}

// PageSearchElderByKey 分页搜索老人（供咨询选择老人）
// 对应 Java: ConsultServiceImpl.pageSearchElderByKey -> elderMapper.listElderByKey
// SQL: SELECT * FROM elder WHERE (elder_name LIKE %key% OR id = key) [可选] AND del_flag=0
// todo: 实现老人分页查询 - 复用 elder 表条件, 结果赋值 out(需定义老人分页返回类型)
func (c *consult) PageSearchElderByKey(ctx context.Context, in *dto.PageSearchElderByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 elder 表并分页
	return nil
}

// PageIntentByKey 分页查询意向（咨询模块内查询意向客户）
// 对应 Java: ConsultServiceImpl.pageIntentByKey -> intentionMapper.listIntentByKey
// SQL: SELECT i.*, e.elder_name FROM intention i LEFT JOIN elder e ON e.id = i.elder_id
//
//	WHERE (e.elder_name LIKE %key% OR i.id = key) [可选]
//
// todo: 实现意向分页查询(可复用 intention service 的 PageIntentByKey 逻辑)
func (c *consult) PageIntentByKey(ctx context.Context, in *dto.PageIntentByKeyQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 查询 intention 联表 elder 并分页
	return nil
}

var _ = dialect.Setter(nil)
