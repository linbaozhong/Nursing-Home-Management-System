package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type serviceproject struct{}

func init() {
	ack.RegisterRoute(&serviceproject{})
}

func (s *serviceproject) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/service")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/getServiceType", s.getServiceType)
	_g.Get("/pageServiceByKey", s.pageServiceByKey)
	_g.Post("/addServiceType", s.addServiceType)
	_g.Get("/getServiceTypeById", s.getServiceTypeById)
	_g.Post("/editServiceType", s.editServiceType)
	_g.Post("/deleteServiceType", s.deleteServiceType)
	_g.Post("/addService", s.addService)
	_g.Get("/getServiceById", s.getServiceById)
	_g.Post("/editService", s.editService)
	_g.Post("/deleteService", s.deleteService)
}

// 获取服务分类
// @Summary 获取服务分类
// @Tags 服务项目
// @Accept application/json
// @Produce application/json
// @Param data query dto.NameReq true "NameReq"
// @Success 200 {object} dto.EmptyResp
// @Router /service/getServiceType [get]
func (s *serviceproject) getServiceType(ctx ack.Context) {
	ack.Get(ctx, service.ServiceProject.ListServiceType)
}

// 分页查询服务
// @Summary 分页查询服务
// @Tags 服务项目
// @Accept application/json
// @Produce application/json
// @Param data query dto.PageServiceByKeyQuery true "PageServiceByKeyQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /service/pageServiceByKey [get]
func (s *serviceproject) pageServiceByKey(ctx ack.Context) {
	ack.Get(ctx, service.ServiceProject.PageServiceByKey)
}

// 新增服务分类
// @Summary 新增服务分类
// @Tags 服务项目
// @Accept application/json
// @Produce application/json
// @Param data body dto.OperateServiceTypeQuery true "OperateServiceTypeQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /service/addServiceType [post]
func (s *serviceproject) addServiceType(ctx ack.Context) {
	ack.Post(ctx, service.ServiceProject.AddServiceType)
}

// 获取服务分类
// @Summary 获取服务分类
// @Tags 服务项目
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /service/getServiceTypeById [get]
func (s *serviceproject) getServiceTypeById(ctx ack.Context) {
	ack.Get(ctx, service.ServiceProject.GetServiceTypeById)
}

// 编辑服务分类
// @Summary 编辑服务分类
// @Tags 服务项目
// @Accept application/json
// @Produce application/json
// @Param data body dto.OperateServiceTypeQuery true "OperateServiceTypeQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /service/editServiceType [post]
func (s *serviceproject) editServiceType(ctx ack.Context) {
	ack.Post(ctx, service.ServiceProject.EditServiceType)
}

// 删除服务分类
// @Summary 删除服务分类
// @Tags 服务项目
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /service/deleteServiceType [post]
func (s *serviceproject) deleteServiceType(ctx ack.Context) {
	ack.Post(ctx, service.ServiceProject.DeleteServiceType)
}

// 新增服务
// @Summary 新增服务
// @Tags 服务项目
// @Accept application/json
// @Produce application/json
// @Param data body dto.OperateServiceQuery true "OperateServiceQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /service/addService [post]
func (s *serviceproject) addService(ctx ack.Context) {
	ack.Post(ctx, service.ServiceProject.AddService)
}

// 获取服务
// @Summary 获取服务
// @Tags 服务项目
// @Accept application/json
// @Produce application/json
// @Param data query dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /service/getServiceById [get]
func (s *serviceproject) getServiceById(ctx ack.Context) {
	ack.Get(ctx, service.ServiceProject.GetServiceById)
}

// 编辑服务
// @Summary 编辑服务
// @Tags 服务项目
// @Accept application/json
// @Produce application/json
// @Param data body dto.OperateServiceQuery true "OperateServiceQuery"
// @Success 200 {object} dto.EmptyResp
// @Router /service/editService [post]
func (s *serviceproject) editService(ctx ack.Context) {
	ack.Post(ctx, service.ServiceProject.EditService)
}

// 删除服务
// @Summary 删除服务
// @Tags 服务项目
// @Accept application/json
// @Produce application/json
// @Param data body dto.IDReq true "IDReq"
// @Success 200 {object} dto.EmptyResp
// @Router /service/deleteService [post]
func (s *serviceproject) deleteService(ctx ack.Context) {
	ack.Post(ctx, service.ServiceProject.DeleteService)
}
