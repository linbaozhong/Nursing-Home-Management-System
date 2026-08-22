package handler

import (
	"api/internal/lib"
	"api/internal/service"
	"github.com/linbaozhong/gentity/pkg/ack/iris"
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

// 上传Image
// @Summary 上传Image
// @Tags 文件
// @Accept application/json
// @Produce application/json
// @Param data body dto.UploadImageReq true "UploadImageReq"
// @Success 200 {object} dto.EmptyResp
// @Router /file/uploadImage [post]
func (f *file) uploadImage(ctx ack.Context) {
	ack.Post(ctx, service.File.UploadImage)
}

// 上传文件
// @Summary 上传文件
// @Tags 文件
// @Accept application/json
// @Produce application/json
// @Param data body dto.UploadFileReq true "UploadFileReq"
// @Success 200 {object} dto.EmptyResp
// @Router /file/uploadFile [post]
func (f *file) uploadFile(ctx ack.Context) {
	ack.Post(ctx, service.File.UploadFile)
}

// 下载文件
// @Summary 下载文件
// @Tags 文件
// @Accept application/json
// @Produce application/json
// @Param data query dto.DownloadFileReq true "DownloadFileReq"
// @Success 200 {object} dto.EmptyResp
// @Router /file/downloadFile [get]
func (f *file) downloadFile(ctx ack.Context) {
	ack.Get(ctx, service.File.DownloadFile)
}
