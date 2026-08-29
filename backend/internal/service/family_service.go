package service

import (
	"api/internal/constant"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/define/table/tblfamilyaccount"
	"api/internal/model/define/table/tblfamilymember"
	"api/internal/model/do"
	"api/internal/model/dto"
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/linbaozhong/gentity/pkg/ace"
	"strconv"
	"time"

	"github.com/linbaozhong/gentity/pkg/oauth/mp/wechat"
	"github.com/linbaozhong/gentity/pkg/token"
	"github.com/linbaozhong/gentity/pkg/types"
)

var Family = &family{}

type family struct{}

// fmMd5 家属密码哈希（与员工端一致，简单 md5）
func fmMd5(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

func idNumTailMatch(full, tail string) bool {
	if len(full) < 4 || len(tail) != 4 {
		return false
	}
	return full[len(full)-4:] == tail
}

// SendCode 发送家属注册/绑定验证码
func (s *family) SendCode(ctx context.Context, in *dto.SendCodeReq, out *dto.SendCodeResp) error {
	if in.Phone == nil || !isValidPhone(*in.Phone) {
		return constant.ErrPhoneError
	}
	code := genCode()
	codeCache.Store(constant.LoginRedis+*in.Phone, codeItem{
		code:   code,
		expire: time.Now().Add(5 * time.Minute).Unix(),
	})
	out.Code = code
	return nil
}

// RegisterBind 家属注册并绑定一位老人
func (s *family) RegisterBind(ctx context.Context, in *dto.RegisterBindReq, out *dto.EmptyResp) error {
	if in.Phone == nil || in.Code == nil || in.Password == nil ||
		in.ElderName == nil || in.IdNumTail == nil {
		return constant.ErrParamError
	}
	if e := checkCode(*in.Phone, *in.Code); e != nil {
		return e
	}
	elder, ok, e := dao.Elder(db).Get(ctx, db.Table(tblelder.TableName).
		Where(tblelder.Name.Eq(types.String(*in.ElderName)), tblelder.Status.Eq(2)))
	if e != nil {
		return e
	}
	if !ok || !idNumTailMatch(elder.IdNum.String(), *in.IdNumTail) {
		return constant.ErrFamilyElderMatch
	}
	exists, e := dao.FamilyAccount(db).Exists(ctx, tblfamilyaccount.Phone.Eq(types.String(*in.Phone)))
	if e != nil {
		return e
	}
	if exists {
		return constant.ErrFamilyPhoneUsed
	}
	if _, e = dao.FamilyAccount(db).Insert(ctx,
		tblfamilyaccount.Phone.Set(types.String(*in.Phone)),
		tblfamilyaccount.Pass.Set(types.String(fmMd5(*in.Password))),
	); e != nil {
		return e
	}
	rel := ""
	if in.Relation != nil {
		rel = *in.Relation
	}
	return s.bindElderRecord(ctx, *in.Phone, elder, rel)
}

// BindElder 已注册家属绑定更多老人
func (s *family) BindElder(ctx context.Context, in *dto.BindElderReq, out *dto.EmptyResp) error {
	if in.Phone == nil || in.Code == nil ||
		in.ElderName == nil || in.IdNumTail == nil {
		return constant.ErrParamError
	}
	if e := checkCode(*in.Phone, *in.Code); e != nil {
		return e
	}
	elder, ok, e := dao.Elder(db).Get(ctx, db.Table(tblelder.TableName).
		Where(tblelder.Name.Eq(types.String(*in.ElderName)), tblelder.Status.Eq(2)))
	if e != nil {
		return e
	}
	if !ok || !idNumTailMatch(elder.IdNum.String(), *in.IdNumTail) {
		return constant.ErrFamilyElderMatch
	}
	ok, e = dao.FamilyAccount(db).Exists(ctx, tblfamilyaccount.Phone.Eq(types.String(*in.Phone)))
	if e != nil {
		return e
	}
	if !ok {
		return constant.ErrFamilyNotBind
	}
	rel := ""
	if in.Relation != nil {
		rel = *in.Relation
	}
	return s.bindElderRecord(ctx, *in.Phone, elder, rel)
}

func (s *family) bindElderRecord(ctx context.Context, phone string, elder *do.Elder, relation string) error {
	_, e := dao.FamilyMember(db).InsertOne(ctx, &do.FamilyMember{
		ElderId:   types.BigInt(elder.Id.Int64()),
		Name:      elder.Name,
		Phone:     types.String(phone),
		IdNum:     elder.IdNum,
		Relation:  types.String(relation),
		IsReceive: types.Int8(0),
	}, tblfamilymember.ElderId, tblfamilymember.Name, tblfamilymember.Phone,
		tblfamilymember.IdNum, tblfamilymember.Relation, tblfamilymember.IsReceive)
	return e
}

// Login 家属登录，返回 token 与绑定老人列表
func (s *family) Login(ctx context.Context, in *dto.FamilyLoginReq, out *dto.FamilyLoginResp) error {
	if in.Phone == nil || in.Password == nil {
		return constant.ErrParamError
	}
	acc, ok, e := dao.FamilyAccount(db).Get(ctx, ace.Where(tblfamilyaccount.Phone.Eq(types.String(*in.Phone))))
	if e != nil {
		return e
	}
	if !ok {
		return constant.ErrFamilyNotBind
	}
	if acc.Pass.String() != fmMd5(*in.Password) {
		return types.NewError(400, "密码错误")
	}
	tk, e := token.GenToken(strconv.FormatInt(acc.Id.Int64(), 10), constant.FamilyRole)
	if e != nil {
		return e
	}
	list, e := s.myElders(ctx, *in.Phone)
	if e != nil {
		return e
	}
	out.Token = tk
	out.ElderList = list
	return nil
}

// MyElders 我的老人列表（家属端数据范围）
func (s *family) MyElders(ctx context.Context, in *dto.FamilyMyEldersReq, out *dto.FamilyMyEldersResp) error {
	if in.Phone == nil {
		return constant.ErrParamError
	}
	list, e := s.myElders(ctx, *in.Phone)
	if e != nil {
		return e
	}
	out.List = list
	return nil
}

func (s *family) myElders(ctx context.Context, phone string) ([]*dto.FamilyElderResp, error) {
	beans, _, e := dao.FamilyMember(db).List(ctx, db.Table(tblfamilymember.TableName).
		Where(tblfamilymember.Phone.Eq(types.String(phone)), tblfamilymember.State.NotEq(types.Int8(constant.StateDeleted))))
	if e != nil {
		return nil, e
	}
	out := make([]*dto.FamilyElderResp, 0, len(beans))
	for _, b := range beans {
		out = append(out, &dto.FamilyElderResp{
			ElderID:  types.BigInt(b.ElderId.Int64()),
			Name:     b.Name.String(),
			Relation: b.Relation.String(),
		})
	}
	return out, nil
}

// MyElderIDs 返回家属可访问的老人 id 集合（供各查询接口做 WHERE elder_id IN (...)）
func (s *family) MyElderIDs(ctx context.Context, phone string) ([]types.BigInt, error) {
	list, e := s.myElders(ctx, phone)
	if e != nil {
		return nil, e
	}
	if len(list) == 0 {
		return nil, constant.ErrFamilyNoElder
	}
	ids := make([]types.BigInt, 0, len(list))
	for _, v := range list {
		ids = append(ids, v.ElderID)
	}
	return ids, nil
}

// IsMyElder 校验某 elderId 是否属于该家属（防越权）
func (s *family) IsMyElder(ctx context.Context, phone string, elderID types.BigInt) (bool, error) {
	ids, e := s.MyElderIDs(ctx, phone)
	if e != nil {
		return false, e
	}
	for _, id := range ids {
		if id == elderID {
			return true, nil
		}
	}
	return false, nil
}

// BindOpenid 家属绑定微信 openid（基于 gentity 的 wechat 小程序包做登录验证）
func (s *family) BindOpenid(ctx context.Context, in *dto.BindOpenidReq, out *dto.BindOpenidResp) error {
	if in.Phone == nil || in.Code == nil {
		return constant.ErrParamError
	}
	// TODO: appid/secret 来源接入配置中心后替换占位符
	mini, e := wechat.Programe(wxAppIDPlaceholder, wxSecretPlaceholder)
	if e != nil {
		return e
	}
	res, e := mini.GetAuth().Code2Session(*in.Code)
	if e != nil {
		return e
	}
	if res.ErrCode != 0 {
		return types.NewError(400, res.ErrMsg)
	}
	if _, e = dao.FamilyAccount(db).Update(ctx, ace.Sets(tblfamilyaccount.Openid.Set(types.String(res.OpenID))).ToSlice(), tblfamilyaccount.Phone.Eq(types.String(*in.Phone))); e != nil {
		return e
	}
	out.Openid = res.OpenID
	return nil
}
