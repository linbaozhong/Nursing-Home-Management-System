package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type source struct{}

func init() {
	ack.RegisterRoute(&source{})
}

func (s *source) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/source")
	_g.Use(lib.AuthMiddleware())

	_g.Get("/pageSourceByKey", s.pageSourceByKey)
	_g.Post("/addSource", s.addSource)
	_g.Get("/getSourceById", s.getSourceById)
	_g.Put("/editSource", s.editSource)
	_g.Delete("/deleteSource", s.deleteSource)
}

func (s *source) pageSourceByKey(ctx ack.Context) {
	ack.Get(ctx, service.Source.PageSourceByKey)
}

func (s *source) addSource(ctx ack.Context) {
	ack.Post(ctx, service.Source.AddSource)
}

func (s *source) getSourceById(ctx ack.Context) {
	ack.Get(ctx, service.Source.GetSourceById)
}

func (s *source) editSource(ctx ack.Context) {
	ack.Put(ctx, service.Source.EditSource)
}

func (s *source) deleteSource(ctx ack.Context) {
	ack.Delete(ctx, service.Source.DeleteSource)
}
