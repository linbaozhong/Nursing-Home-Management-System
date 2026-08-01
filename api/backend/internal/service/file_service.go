package service

import (
	"context"

	"api/internal/model/dto"
)

type file struct{}

var File = &file{}

// UploadImage 上传图片
// 对应 Java: FileServiceImpl 或通用文件上传(通常存本地/对象存储, 返回访问 URL)
// todo: 接收上传的图片文件, 保存到存储(本地目录或 OSS/MinIO), 返回文件 URL 并赋值 out(需定义返回类型)
func (f *file) UploadImage(ctx context.Context, in *dto.UploadImageQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 实现图片上传
	return nil
}

// UploadFile 上传文件
// 对应 Java: 通用文件上传, 返回访问 URL
// todo: 接收上传文件, 保存到存储, 返回文件 URL 并赋值 out(需定义返回类型)
func (f *file) UploadFile(ctx context.Context, in *dto.UploadFileQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 实现文件上传
	return nil
}

// DownloadFile 下载文件
// 对应 Java: 根据文件 ID/路径从存储读取并流式返回
// todo: 按 in 中的文件标识从存储读取, 写入响应流(本方法直接写 HTTP 响应, 不走 out)
func (f *file) DownloadFile(ctx context.Context, in *dto.DownloadFileQuery, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 实现文件下载(通常需要 ctx 中的 responseWriter)
	return nil
}
