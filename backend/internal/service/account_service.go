package service

import (
	"context"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"api/internal/constant"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblmember"
	"api/internal/model/define/table/tbltenant"
	"api/internal/model/define/table/tbluser"
	"api/internal/model/do"
	"api/internal/model/dto"

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
func (a *account) SendCode(ctx context.Context, in *dto.SendCodeReq, out *dto.LoginUserResp) error {
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

// Login 登录（账号密码，备用登录）
func (a *account) Login(ctx context.Context, in *dto.LoginReq, out *dto.LoginUserResp) error {
	if in.Phone == nil || !isValidPhone(*in.Phone) {
		return constant.ErrPhoneError
	}
	if in.Pass == nil || *in.Pass == "" {
		return constant.ErrPassword
	}
	user, has, e := dao.User(db).Get(ctx, db.Table(tbluser.TableName).Where(tbluser.Phone.Eq(types.String(*in.Phone))))
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrNoRegister
	}
	if *in.Pass != user.Pass.String() {
		return constant.ErrPassword
	}
	return a.loginWithUser(ctx, user, out)
}

// loginWithUser 根据全局用户完成登录分流（单企业直接登录，多企业返回列表）
func (a *account) loginWithUser(ctx context.Context, user *do.User, out *dto.LoginUserResp) error {
	userID := user.Id.Int64()
	memberList, _, e := dao.Member(db).List(ctx, db.Table(tblmember.TableName).
		Where(tblmember.UserId.Eq(types.BigInt(userID)),
			tblmember.Status.Eq(types.Int8(constant.MemberStatusActive)),
			tblmember.DelFlag.Eq(0)))
	if e != nil {
		return e
	}
	if len(memberList) == 0 {
		out.NeedBind = true
		return nil
	}
	tenants := make([]dto.TenantResp, 0, len(memberList))
	for _, m := range memberList {
		tn, hasT, e2 := dao.Tenant(db).GetByID(ctx, m.TenantId,
			tbltenant.Id, tbltenant.Name, tbltenant.Logo, tbltenant.ContactName,
			tbltenant.ContactPhone, tbltenant.Plan, tbltenant.Status,
			tbltenant.TrialStart, tbltenant.TrialEnd)
		if e2 != nil {
			return e2
		}
		if !hasT {
			continue
		}
		tenants = append(tenants, dto.TenantResp{
			ID:           tn.Id.Int64(),
			Name:         tn.Name.String(),
			Logo:         tn.Logo.String(),
			ContactName:  tn.ContactName.String(),
			ContactPhone: tn.ContactPhone.String(),
			Plan:         tn.Plan.String(),
			Status:       tn.Status.Int8(),
			TrialStart:   tn.TrialStart.Time,
			TrialEnd:     tn.TrialEnd.Time,
		})
	}
	if len(tenants) == 0 {
		out.NeedBind = true
		return nil
	}
	out.Tenants = tenants
	if len(tenants) == 1 {
		m := memberList[0]
		return a.fillUserByMember(ctx, userID, m.TenantId.Int64(), m.Id.Int64(), m.RoleId.Int64(), out)
	}
	out.NeedBind = true
	return nil
}

// Forget 忘记密码
func (a *account) Forget(ctx context.Context, in *dto.ForgetReq, out *dto.LoginUserResp) error {
	if in.Account == nil || !isValidPhone(*in.Account) {
		return constant.ErrPhoneError
	}
	if in.Code == nil || *in.Code == "" {
		return constant.ErrNullCode
	}
	if e := checkCode(*in.Account, *in.Code); e != nil {
		return e
	}
	user, has, e := dao.User(db).Get(ctx, db.Table(tbluser.TableName).Where(tbluser.Phone.Eq(types.String(*in.Account))))
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrNoRegister
	}
	_, e = dao.User(db).UpdateById(ctx, user.Id, tbluser.Pass.Set(types.String(md5Str(*in.Pass))))
	if e != nil {
		return e
	}
	user.Pass = types.String(md5Str(*in.Pass))
	return a.loginWithUser(ctx, user, out)
}

// Edit 修改密码
func (a *account) Edit(ctx context.Context, in *dto.EditReq, out *dto.LoginUserResp) error {
	if in.ID == nil {
		return constant.ErrDataNotExist
	}
	user, has, e := dao.User(db).GetByID(ctx, types.BigInt(*in.ID),
		tbluser.Id, tbluser.Name, tbluser.Phone, tbluser.Avator, tbluser.Pass)
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrDataNotExist
	}
	if in.OldPass == nil || md5Str(*in.OldPass) != user.Pass.String() {
		return constant.ErrPassword
	}
	_, e = dao.User(db).UpdateById(ctx, user.Id, tbluser.Pass.Set(types.String(md5Str(*in.NewPass))))
	if e != nil {
		return e
	}
	user.Pass = types.String(md5Str(*in.NewPass))
	return a.loginWithUser(ctx, user, out)
}

// Logout 退出登录
func (a *account) Logout(ctx context.Context, in *dto.LoginReq, out *dto.EmptyResp) error {
	// 无状态 token，直接返回成功
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
