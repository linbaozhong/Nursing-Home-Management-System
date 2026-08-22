package service

import (
	"context"
	"path"
	"strings"

	"api/internal/constant"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblbaseattachment"
	"api/internal/model/do"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/types"
)

var File = (*fileService)(nil)

type fileService struct{}

// UploadImage 上传图片，保存附件记录并返回访问信息
func (s *fileService) UploadImage(ctx context.Context, in *dto.UploadImageReq, out *dto.FileInfoResp) error {
	// file 字段约定为完整的访问 URL（或相对路径），取文件名作附件名
	name := ""
	if in.File != nil {
		name = path.Base(*in.File)
		if i := strings.LastIndex(name, "/"); i >= 0 {
			name = name[i+1:]
		}
	}
	att := &do.BaseAttachment{
		Name: types.String(name),
		Url:  types.String(orEmpty(in.File)),
		Path: types.String(orEmpty(in.File)),
		Suff: types.String(suffix(name)),
	}
	if _, e := dao.BaseAttachment(db).InsertOne(ctx, att); e != nil {
		return e
	}
	out.ID = int64(att.Id)
	out.URL = att.Url.String()
	return nil
}

// UploadFile 上传文件，保存附件记录并返回访问信息
func (s *fileService) UploadFile(ctx context.Context, in *dto.UploadFileReq, out *dto.FileInfoResp) error {
	name := ""
	if in.File != nil {
		name = path.Base(*in.File)
		if i := strings.LastIndex(name, "/"); i >= 0 {
			name = name[i+1:]
		}
	}
	att := &do.BaseAttachment{
		Name: types.String(name),
		Url:  types.String(orEmpty(in.File)),
		Path: types.String(orEmpty(in.File)),
		Suff: types.String(suffix(name)),
	}
	if _, e := dao.BaseAttachment(db).InsertOne(ctx, att); e != nil {
		return e
	}
	out.ID = int64(att.Id)
	out.URL = att.Url.String()
	return nil
}

// DownloadFile 根据编号下载文件，返回附件记录
func (s *fileService) DownloadFile(ctx context.Context, in *dto.DownloadFileReq, out *dto.FileInfoResp) error {
	if in.ID == nil {
		return constant.ErrParamInvalid
	}
	att := new(do.BaseAttachment)
	has, e := dao.BaseAttachment(db).Get(ctx, ace.Where(tblbaseattachment.Id.Eq(types.BigInt(*in.ID))).
		And(tblbaseattachment.DelFlag.Eq(types.Int8(constant.YesNoNo))))
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrDataNotExist
	}
	out.ID = int64(att.Id)
	out.URL = att.Url.String()
	return nil
}

func orEmpty(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func suffix(name string) string {
	if i := strings.LastIndex(name, "."); i >= 0 && i < len(name)-1 {
		return name[i+1:]
	}
	return ""
}
