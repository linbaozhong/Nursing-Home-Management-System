package dto

// ============ FileController 响应 ============

// @response
// FileInfoVO 文件上传成功响应
type FileInfoVO struct {
	ID  int64  `json:"id"`  // id
	URL string `json:"url"` // url
}

// @request
// UploadImageQuery 上传图片请求
type UploadImageQuery struct {
	Module *string `json:"module" valid:"required"` // 业务模块标识
	File   *string `json:"file" valid:"required"`   // 文件内容(base64 或路径)
}

// @request
// UploadFileQuery 上传文件请求
type UploadFileQuery struct {
	Module *string `json:"module" valid:"required"` // 业务模块标识
	File   *string `json:"file" valid:"required"`   // 文件内容(base64 或路径)
}

// @request
// DownloadFileQuery 下载文件请求
type DownloadFileQuery struct {
	ID       *int64  `json:"id" valid:"required"` // 文件编号
	Module   *string `json:"module"`              // 业务模块标识
	FileName *string `json:"file_name"`           // 文件名
}
