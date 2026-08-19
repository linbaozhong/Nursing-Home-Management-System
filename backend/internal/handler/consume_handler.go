package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
)

type consume struct{}

func init() {
	ack.RegisterRoute(&consume{})
}

func (c *consume) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/consume")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageConsumeByKey", c.pageConsumeByKey)
	_g.Get("/getConsumeById", c.getConsumeById)
	_g.Post("/addConsume", c.addConsume)
	_g.Post("/editConsume", c.editConsume)
	_g.Post("/deleteConsume", c.deleteConsume)
}

// 分页查询消费
// @Summary 分页查询消费
// @Tags 消费
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageConsumeByKeyQuery true "PageConsumeByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /consume/pageConsumeByKey [get]
func (c *consume) pageConsumeByKey(ctx ack.Context) {
	ack.Get(ctx, service.Consume.PageConsumeByKey)
}

// 获取消费
// @Summary 获取消费
// @Tags 消费
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /consume/getConsumeById [get]
func (c *consume) getConsumeById(ctx ack.Context) {
	ack.Get(ctx, service.Consume.GetConsumeById)
}

// 新增消费
// @Summary 新增消费
// @Tags 消费
// @Accept application/json
// @Produce application/json
// @Param data body dto.AddConsumeQuery true "AddConsumeQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /consume/addConsume [post]
func (c *consume) addConsume(ctx ack.Context) {
	ack.Post(ctx, service.Consume.AddConsume)
}

// 编辑消费
// @Summary 编辑消费
// @Tags 消费
// @Accept application/json
// @Produce application/json
// @Param data body dto.EditConsumeQuery true "EditConsumeQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /consume/editConsume [post]
func (c *consume) editConsume(ctx ack.Context) {
	ack.Post(ctx, service.Consume.EditConsume)
}

// 删除消费
// @Summary 删除消费
// @Tags 消费
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /consume/deleteConsume [post]
func (c *consume) deleteConsume(ctx ack.Context) {
	ack.Post(ctx, service.Consume.DeleteConsume)
}
