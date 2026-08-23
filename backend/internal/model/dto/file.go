package dto

import "github.com/linbaozhong/gentity/pkg/types"

// ============ FileController 响应 ============

// FileInfoResp 文件上传成功响应
// @response
type FileInfoResp struct {
	ID  types.BigInt `json:"id"`  // id
	URL string       `json:"url"` // url
}

// UploadImageReq 上传图片请求
// @request
type UploadImageReq struct {
	Module *string `json:"module" valid:"required"` // 业务模块标识
	File   *string `json:"file" valid:"required"`   // 文件内容(base64 或路径)
}

// UploadFileReq 上传文件请求
// @request
type UploadFileReq struct {
	Module *string `json:"module" valid:"required"` // 业务模块标识
	File   *string `json:"file" valid:"required"`   // 文件内容(base64 或路径)
}

// DownloadFileReq 下载文件请求
// @request
type DownloadFileReq struct {
	ID       *int64  `json:"id" valid:"required"` // 文件编号
	Module   *string `json:"module"`              // 业务模块标识
	FileName *string `json:"file_name"`           // 文件名
}
