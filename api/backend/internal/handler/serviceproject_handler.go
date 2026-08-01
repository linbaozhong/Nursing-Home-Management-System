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
	_g.Put("/editServiceType", s.editServiceType)
	_g.Delete("/deleteServiceType", s.deleteServiceType)
	_g.Post("/addService", s.addService)
	_g.Get("/getServiceById", s.getServiceById)
	_g.Put("/editService", s.editService)
	_g.Delete("/deleteService", s.deleteService)
}

func (s *serviceproject) getServiceType(ctx ack.Context) {
	ack.Get(ctx, service.ServiceProject.ListServiceType)
}

func (s *serviceproject) pageServiceByKey(ctx ack.Context) {
	ack.Get(ctx, service.ServiceProject.PageServiceByKey)
}

func (s *serviceproject) addServiceType(ctx ack.Context) {
	ack.Post(ctx, service.ServiceProject.AddServiceType)
}

func (s *serviceproject) getServiceTypeById(ctx ack.Context) {
	ack.Get(ctx, service.ServiceProject.GetServiceTypeById)
}

func (s *serviceproject) editServiceType(ctx ack.Context) {
	ack.Put(ctx, service.ServiceProject.EditServiceType)
}

func (s *serviceproject) deleteServiceType(ctx ack.Context) {
	ack.Delete(ctx, service.ServiceProject.DeleteServiceType)
}

func (s *serviceproject) addService(ctx ack.Context) {
	ack.Post(ctx, service.ServiceProject.AddService)
}

func (s *serviceproject) getServiceById(ctx ack.Context) {
	ack.Get(ctx, service.ServiceProject.GetServiceById)
}

func (s *serviceproject) editService(ctx ack.Context) {
	ack.Put(ctx, service.ServiceProject.EditService)
}

func (s *serviceproject) deleteService(ctx ack.Context) {
	ack.Delete(ctx, service.ServiceProject.DeleteService)
}
