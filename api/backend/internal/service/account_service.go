package service

import (
	"context"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"api/internal/constant"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblauth"
	"api/internal/model/define/table/tblroleauth"
	"api/internal/model/define/table/tblstaff"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/token"
	"github.com/linbaozhong/gentity/pkg/types"
)

type account struct{}

var Account = &account{}

// codeItem 验证码缓存项
type codeItem struct {
	code   string
	expire int64
}

// codeCache 进程内验证码缓存 (key: phone -> codeItem)
var codeCache sync.Map

// SendCode 发送验证码
func (a *account) SendCode(ctx context.Context, in *dto.SendCodeQuery, out *dto.LoginUserVO) error {
	if in.Phone == nil || !isValidPhone(*in.Phone) {
		return constant.ErrPhoneError
	}
	code := genCode()
	codeCache.Store(constant.LoginRedis+*in.Phone, codeItem{
		code:   code,
		expire: time.Now().Add(5 * time.Minute).Unix(),
	})
	out.Pass = code
	return nil
}

// Login 登录
func (a *account) Login(ctx context.Context, in *dto.LoginQuery, out *dto.LoginUserVO) error {
	if in.Phone == nil || !isValidPhone(*in.Phone) {
		return constant.ErrPhoneError
	}
	if in.Code == nil || *in.Code == "" {
		return constant.ErrNullCode
	}
	if e := checkCode(*in.Phone, *in.Code); e != nil {
		return e
	}
	staff, has, e := dao.Staff(db).Get(ctx, db.Table(tblstaff.TableName).Where(tblstaff.Phone.Eq(types.String(*in.Phone))))
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrNoRegister
	}
	if in.Pass == nil || *in.Pass != staff.Pass.String() {
		return constant.ErrPassword
	}
	return a.fillUser(ctx, staff, out)
}

// Forget 忘记密码
func (a *account) Forget(ctx context.Context, in *dto.ForgetQuery, out *dto.LoginUserVO) error {
	if in.Account == nil || !isValidPhone(*in.Account) {
		return constant.ErrPhoneError
	}
	if in.Code == nil || *in.Code == "" {
		return constant.ErrNullCode
	}
	if e := checkCode(*in.Account, *in.Code); e != nil {
		return e
	}
	staff, has, e := dao.Staff(db).Get(ctx, db.Table(tblstaff.TableName).Where(tblstaff.Phone.Eq(types.String(*in.Account))))
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrNoRegister
	}
	_, e = dao.Staff(db).UpdateById(ctx, staff.Id, tblstaff.Pass.Set(types.String(*in.Pass)))
	if e != nil {
		return e
	}
	staff.Pass = types.String(*in.Pass)
	return a.fillUser(ctx, staff, out)
}

// Edit 修改密码
func (a *account) Edit(ctx context.Context, in *dto.EditQuery, out *dto.LoginUserVO) error {
	if in.ID == nil {
		return constant.ErrDataNotExist
	}
	staff, has, e := dao.Staff(db).GetByID(ctx, types.BigInt(*in.ID),
		tblstaff.Id, tblstaff.RoleId, tblstaff.Name, tblstaff.Phone, tblstaff.Avator, tblstaff.Pass)
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrDataNotExist
	}
	if in.OldPass == nil || *in.OldPass != staff.Pass.String() {
		return constant.ErrPassword
	}
	_, e = dao.Staff(db).UpdateById(ctx, types.BigInt(*in.ID), tblstaff.Pass.Set(types.String(*in.NewPass)))
	if e != nil {
		return e
	}
	staff.Pass = types.String(*in.NewPass)
	return a.fillUser(ctx, staff, out)
}

// Logout 退出登录
func (a *account) Logout(ctx context.Context, in *dto.LoginQuery, out *dto.EmptyResp) error {
	// 无状态 token，直接返回成功
	return nil
}

// fillUser 组装登录用户信息与权限，并生成 token
func (a *account) fillUser(ctx context.Context, staff *do.Staff, out *dto.LoginUserVO) error {
	var ras []do.RoleAuth
	e := db.Table(tblroleauth.TableName).
		Where(tblroleauth.RoleId.Eq(staff.RoleId)).
		Cols(tblroleauth.AuthId).
		Select().
		Gets(ctx, &ras)
	if e != nil {
		return e
	}
	var authIDs []int64
	for _, ra := range ras {
		authIDs = append(authIDs, int64(ra.AuthId))
	}
	var authUrls []string
	if len(authIDs) > 0 {
		var auths []do.Auth
		e = db.Table(tblauth.TableName).
			Where(tblauth.Id.In(int64ToAny(authIDs)...)).
			Cols(tblauth.Id, tblauth.Name, tblauth.Url).
			Select().
			Gets(ctx, &auths)
		if e != nil {
			return e
		}
		for _, au := range auths {
			authUrls = append(authUrls, au.Url.String())
		}
	}
	tk, e := token.GenToken(strconv.FormatInt(int64(staff.Id), 10), "staff")
	if e != nil {
		return e
	}
	out.ID = int64(staff.Id)
	out.Name = staff.Name.String()
	out.Avator = staff.Avator.String()
	out.Phone = staff.Phone.String()
	out.Pass = staff.Pass.String()
	out.AuthIDList = authIDs
	out.AuthUrlList = authUrls
	out.Token = tk
	return nil
}

// isValidPhone 简单手机号格式校验
func isValidPhone(phone string) bool {
	if len(phone) != 11 {
		return false
	}
	for _, c := range phone {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

// checkCode 校验验证码
func checkCode(phone, code string) error {
	v, ok := codeCache.Load(constant.LoginRedis + phone)
	if !ok {
		return constant.ErrCodeExpire
	}
	item := v.(codeItem)
	if item.expire < time.Now().Unix() {
		codeCache.Delete(constant.LoginRedis + phone)
		return constant.ErrCodeExpire
	}
	if item.code != code {
		return constant.ErrCodeError
	}
	codeCache.Delete(constant.LoginRedis + phone)
	return nil
}

// genCode 生成 6 位随机验证码
func genCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.FormatInt(int64(r.Intn(1000000)), 10)
}

// int64ToAny []int64 -> []any
func int64ToAny(in []int64) []any {
	out := make([]any, len(in))
	for i, v := range in {
		out[i] = v
	}
	return out
}
