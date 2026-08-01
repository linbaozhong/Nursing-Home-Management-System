package handler

import (
	"github.com/linbaozhong/gentity/pkg/ack/iris"
	"api/internal/lib"
	"api/internal/service"
)

type file struct{}

func init() {
	ack.RegisterRoute(&file{})
}

func (f *file) RegisterRoute(group ack.Party) {
	_g := ack.NewParty(group, "/file")
	_g.Use(lib.AuthMiddleware())

	_g.Post("/uploadImage", f.uploadImage)
	_g.Post("/uploadFile", f.uploadFile)
	_g.Get("/downloadFile", f.downloadFile)
}

func (f *file) uploadImage(ctx ack.Context) {
	ack.Post(ctx, service.File.UploadImage)
}

func (f *file) uploadFile(ctx ack.Context) {
	ack.Post(ctx, service.File.UploadFile)
}

func (f *file) downloadFile(ctx ack.Context) {
	ack.Get(ctx, service.File.DownloadFile)
}
