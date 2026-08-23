package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"

	"api/internal/constant"
	"api/internal/lib"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblauth"
	"api/internal/model/define/table/tblmember"
	"api/internal/model/define/table/tblroleauth"
	"api/internal/model/define/table/tbltenant"
	"api/internal/model/define/table/tbluser"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/oauth/mp/wechat"
	"github.com/linbaozhong/gentity/pkg/token"
	"github.com/linbaozhong/gentity/pkg/types"
)

type tenant struct{}

var Tenant = &tenant{}

// md5Str 简单 md5 哈希
func md5Str(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

// ptrStr 指针字符串取值
func ptrStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// wxCode2Session 微信小程序登录换取 openid/unionid
func wxCode2Session(code string) (openID, unionID string, err error) {
	mini, e := wechat.Programe(wxAppIDPlaceholder, wxSecretPlaceholder)
	if e != nil {
		return "", "", e
	}
	res, e := mini.Code2Session(code)
	if e != nil {
		return "", "", e
	}
	if res.ErrCode != 0 {
		return "", "", types.NewError(400, res.ErrMsg)
	}
	return res.OpenID, res.UnionID, nil
}

// Register 租户自助注册：创建租户 + 创建管理员 User + 创建 Member（租户管理员），立即开通并设试用期
func (t *tenant) Register(ctx context.Context, in *dto.RegisterTenantReq, out *dto.LoginUserResp) error {
	if in.Name == nil || in.ContactName == nil || in.ContactPhone == nil || in.Password == nil {
		return constant.ErrParamError
	}
	if !isValidPhone(*in.ContactPhone) {
		return constant.ErrPhoneError
	}
	// 校验企业名称唯一
	if has, e := dao.Tenant(db).Exists(ctx,
		tbltenant.Name.Eq(types.String(*in.Name)), tbltenant.DelFlag.Eq(0)); e != nil {
		return e
	} else if has {
		return constant.ErrTenantNameRepeat
	}

	pass := md5Str(*in.Password)

	// 查找或创建全局 User
	var userID int64
	user, has, e := dao.User(db).Get(ctx, db.Table(tbluser.TableName).
		Where(tbluser.Phone.Eq(types.String(*in.ContactPhone)), tbluser.DelFlag.Eq(0)))
	if e != nil {
		return e
	}
	unionID := ""
	if has {
		userID = user.Id.Int64()
	} else {
		// 绑定微信
		if in.WxCode != nil && *in.WxCode != "" {
			openID, unionID2, e2 := wxCode2Session(*in.WxCode)
			if e2 != nil {
				return e2
			}
			unionID = unionID2
			// 尝试按 unionID 匹配已有用户（同一微信已注册）
			if u, ok, e3 := dao.User(db).Get(ctx, db.Table(tbluser.TableName).
				Where(tbluser.UnionId.Eq(types.String(unionID)), tbluser.DelFlag.Eq(0))); e3 != nil {
				return e3
			} else if ok {
				user = u
				has = true
				userID = user.Id.Int64()
				_ = openID
			}
		}
		if !has {
			newID, e2 := dao.User(db).Insert(ctx,
				tbluser.Phone.Set(types.String(*in.ContactPhone)),
				tbluser.Pass.Set(types.String(pass)),
				tbluser.Name.Set(types.String(*in.ContactName)),
				tbluser.UnionId.Set(types.String(unionID)),
				tbluser.CreateTime.Set(types.Time{Time: time.Now()}),
				tbluser.UpdateTime.Set(types.Time{Time: time.Now()}),
				tbluser.DelFlag.Set(types.Int8(0)),
			)
			if e2 != nil {
				return e2
			}
			userID = newID
		}
	}

	// 创建租户，立即开通并设置试用期
	now := time.Now()
	trialEnd := now.AddDate(0, 0, constant.TrialDays)
	tenantID, e := dao.Tenant(db).Insert(ctx,
		tbltenant.Name.Set(types.String(*in.Name)),
		tbltenant.Logo.Set(types.String(ptrStr(in.Logo))),
		tbltenant.ContactName.Set(types.String(*in.ContactName)),
		tbltenant.ContactPhone.Set(types.String(*in.ContactPhone)),
		tbltenant.Plan.Set(types.String("trial")),
		tbltenant.Status.Set(types.Int8(constant.TenantStatusTrial)),
		tbltenant.TrialStart.Set(types.Time{Time: now}),
		tbltenant.TrialEnd.Set(types.Time{Time: trialEnd}),
		tbltenant.CreateId.Set(types.BigInt(userID)),
		tbltenant.CreateTime.Set(types.Time{Time: now}),
		tbltenant.UpdateId.Set(types.BigInt(userID)),
		tbltenant.UpdateTime.Set(types.Time{Time: now}),
		tbltenant.DelFlag.Set(types.Int8(0)),
	)
	if e != nil {
		return e
	}

	// 创建成员（租户管理员），role_id=1
	memberID, e := dao.Member(db).Insert(ctx,
		tblmember.UserId.Set(types.BigInt(userID)),
		tblmember.TenantId.Set(types.BigInt(tenantID)),
		tblmember.RoleId.Set(types.BigInt(1)),
		tblmember.Status.Set(types.Int8(constant.MemberStatusActive)),
		tblmember.CreateId.Set(types.BigInt(userID)),
		tblmember.CreateTime.Set(types.Time{Time: now}),
		tblmember.UpdateId.Set(types.BigInt(userID)),
		tblmember.UpdateTime.Set(types.Time{Time: now}),
		tblmember.DelFlag.Set(types.Int8(0)),
	)
	if e != nil {
		return e
	}

	return Account.fillUserByMember(ctx, userID, tenantID, memberID, 1, out)
}

// Open 平台开通/解锁租户（仅平台管理员可调用）
func (t *tenant) Open(ctx context.Context, in *dto.OpenTenantReq, out *dto.EmptyResp) error {
	if in.ID == nil {
		return constant.ErrParamError
	}
	if lib.Role(ctx) != constant.PlatformRole {
		return constant.ErrNotPlatformAdmin
	}
	tenantID := types.BigInt(*in.ID)
	_, has, e := dao.Tenant(db).GetByID(ctx, tenantID,
		tbltenant.Id, tbltenant.Status, tbltenant.TrialStart, tbltenant.TrialEnd)
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrTenantNotExist
	}
	// 解锁并延长试用期（重新计 30 天）
	now := time.Now()
	_, e = dao.Tenant(db).UpdateById(ctx, tenantID,
		tbltenant.Status.Set(types.Int8(constant.TenantStatusTrial)),
		tbltenant.TrialStart.Set(types.Time{Time: now}),
		tbltenant.TrialEnd.Set(types.Time{Time: now.AddDate(0, 0, constant.TrialDays)}),
		tbltenant.UpdateTime.Set(types.Time{Time: now}),
	)
	return e
}

// Lock 试用到期锁定
func (t *tenant) Lock(ctx context.Context, in *dto.OpenTenantReq, out *dto.EmptyResp) error {
	if in.ID == nil {
		return constant.ErrParamError
	}
	tenantID := types.BigInt(*in.ID)
	tn, has, e := dao.Tenant(db).GetByID(ctx, tenantID, tbltenant.Id, tbltenant.Status, tbltenant.TrialEnd)
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrTenantNotExist
	}
	if tn.Status.Int8() == constant.TenantStatusTrial && !tn.TrialEnd.IsZero() &&
		tn.TrialEnd.Time.Before(time.Now()) {
		_, e = dao.Tenant(db).UpdateById(ctx, tenantID,
			tbltenant.Status.Set(types.Int8(constant.TenantStatusLocked)),
			tbltenant.UpdateTime.Set(types.Time{Time: time.Now()}),
		)
		return e
	}
	return nil
}

// MyTenants 当前用户已绑定的企业列表（一人多企）
func (t *tenant) MyTenants(ctx context.Context, in *dto.EmptyReq, out *dto.UserTenantListResp) error {
	userID := lib.UserID(ctx)
	memberList, _, e := dao.Member(db).List(ctx, db.Table(tblmember.TableName).
		Where(tblmember.UserId.Eq(types.BigInt(userID)),
			tblmember.Status.Eq(types.Int8(constant.MemberStatusActive)),
			tblmember.DelFlag.Eq(0)))
	if e != nil {
		return e
	}
	tenants := make([]dto.TenantResp, 0, len(memberList))
	for _, m := range memberList {
		tn, has, e2 := dao.Tenant(db).GetByID(ctx, m.TenantId,
			tbltenant.Id, tbltenant.Name, tbltenant.Logo, tbltenant.ContactName,
			tbltenant.ContactPhone, tbltenant.Plan, tbltenant.Status,
			tbltenant.TrialStart, tbltenant.TrialEnd)
		if e2 != nil {
			return e2
		}
		if !has {
			continue
		}
		tenants = append(tenants, dto.TenantResp{
			ID:           types.BigInt(tn.Id.Int64()),
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
	out.Tenants = tenants
	out.Current = types.BigInt(lib.TenantID(ctx))
	return nil
}

// SwitchTenant 切换当前租户（一人多企时）
func (t *tenant) SwitchTenant(ctx context.Context, in *dto.SwitchTenantReq, out *dto.LoginUserResp) error {
	if in.TenantID == nil {
		return constant.ErrParamError
	}
	userID := lib.UserID(ctx)
	member, has, e := dao.Member(db).Get(ctx, db.Table(tblmember.TableName).
		Where(tblmember.UserId.Eq(types.BigInt(userID)),
			tblmember.TenantId.Eq(types.BigInt(*in.TenantID)),
			tblmember.DelFlag.Eq(0)))
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrMemberNotExist
	}
	tn, has, e := dao.Tenant(db).GetByID(ctx, member.TenantId,
		tbltenant.Id, tbltenant.Status, tbltenant.TrialEnd)
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrTenantNotExist
	}
	if e = checkTenantUsable(tn); e != nil {
		return e
	}
	return Account.fillUserByMember(ctx, userID, member.TenantId.Int64(), member.Id.Int64(), member.RoleId.Int64(), out)
}

// InviteMember 租户管理员邀请成员（根据手机号加入）
func (t *tenant) InviteMember(ctx context.Context, in *dto.InviteMemberReq, out *dto.MemberResp) error {
	if in.TenantID == nil || in.Phone == nil || in.RoleID == nil {
		return constant.ErrParamError
	}
	tenantID := types.BigInt(*in.TenantID)
	if !isTenantAdmin(ctx, tenantID.Int64()) {
		return constant.ErrNotTenantAdmin
	}
	user, has, e := dao.User(db).Get(ctx, db.Table(tbluser.TableName).
		Where(tbluser.Phone.Eq(types.String(*in.Phone)), tbluser.DelFlag.Eq(0)))
	if e != nil {
		return e
	}
	var userID int64
	if has {
		userID = user.Id.Int64()
	} else {
		newID, e2 := dao.User(db).Insert(ctx,
			tbluser.Phone.Set(types.String(*in.Phone)),
			tbluser.CreateTime.Set(types.Time{Time: time.Now()}),
			tbluser.UpdateTime.Set(types.Time{Time: time.Now()}),
			tbluser.DelFlag.Set(types.Int8(0)),
		)
		if e2 != nil {
			return e2
		}
		userID = newID
	}
	if hasMember, e := dao.Member(db).Exists(ctx,
		tblmember.UserId.Eq(types.BigInt(userID)),
		tblmember.TenantId.Eq(tenantID),
		tblmember.DelFlag.Eq(0)); e != nil {
		return e
	} else if hasMember {
		return constant.ErrMemberAlreadyExist
	}
	now := time.Now()
	memberID, e := dao.Member(db).Insert(ctx,
		tblmember.UserId.Set(types.BigInt(userID)),
		tblmember.TenantId.Set(tenantID),
		tblmember.RoleId.Set(types.BigInt(*in.RoleID)),
		tblmember.Status.Set(types.Int8(constant.MemberStatusActive)),
		tblmember.CreateId.Set(types.BigInt(lib.UserID(ctx))),
		tblmember.CreateTime.Set(types.Time{Time: now}),
		tblmember.UpdateId.Set(types.BigInt(lib.UserID(ctx))),
		tblmember.UpdateTime.Set(types.Time{Time: now}),
		tblmember.DelFlag.Set(types.Int8(0)),
	)
	if e != nil {
		return e
	}
	out.ID = types.BigInt(memberID)
	out.UserID = types.BigInt(userID)
	out.Phone = *in.Phone
	out.RoleID = types.BigInt(*in.RoleID)
	out.Status = constant.MemberStatusActive
	return nil
}

// WxLogin 小程序微信静默登录（智能分流）
func (t *tenant) WxLogin(ctx context.Context, in *dto.WxLoginReq, out *dto.WxLoginResp) error {
	if in.Code == nil || *in.Code == "" {
		return constant.ErrParamError
	}
	openID, unionID, e := wxCode2Session(*in.Code)
	if e != nil {
		return e
	}
	// 按 unionID 优先匹配，其次 openid
	var user *do.User
	var has bool
	if unionID != "" {
		user, has, e = dao.User(db).Get(ctx, db.Table(tbluser.TableName).
			Where(tbluser.UnionId.Eq(types.String(unionID)), tbluser.DelFlag.Eq(0)))
		if e != nil {
			return e
		}
	}
	if !has && openID != "" {
		user, has, e = dao.User(db).Get(ctx, db.Table(tbluser.TableName).
			Where(tbluser.Openid.Eq(types.String(openID)), tbluser.DelFlag.Eq(0)))
		if e != nil {
			return e
		}
	}
	// 新用户：返回 need_bind，交由前端引导绑定/创建企业
	if !has {
		out.NeedBind = true
		return nil
	}
	userID := user.Id.Int64()
	// 查询该用户所有在职成员（绑定企业）
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

	// 构建企业列表
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
		if e2 = checkTenantUsable(tn); e2 != nil {
			return e2
		}
		tenants = append(tenants, dto.TenantResp{
			ID:           types.BigInt(tn.Id.Int64()),
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
	out.Tenants = tenants

	// 唯一企业：直接登录，返回 token
	if len(tenants) == 1 {
		m := memberList[0]
		lu := &dto.LoginUserResp{}
		if e = Account.fillUserByMember(ctx, userID, m.TenantId.Int64(), m.Id.Int64(), m.RoleId.Int64(), lu); e != nil {
			return e
		}
		out.Token = lu.Token
		out.User = lu
		out.NeedBind = false
		return nil
	}
	// 多企业：返回企业列表供选择
	out.NeedBind = true
	return nil
}

// ============ 辅助方法 ============

// fillUserByMember 根据 user/tenant/member 组装登录信息并生成 token
func (a *account) fillUserByMember(ctx context.Context, userID, tenantID, memberID, roleID int64, out *dto.LoginUserResp) error {
	user, has, e := dao.User(db).GetByID(ctx, types.BigInt(userID),
		tbluser.Id, tbluser.Name, tbluser.Avator, tbluser.Phone, tbluser.Pass)
	if e != nil {
		return e
	}
	if !has {
		return constant.ErrUserNotExist
	}
	authUrls, e := Account.authUrlsByRole(ctx, roleID)
	if e != nil {
		return e
	}
	tk, e := token.GenToken(
		strconv.FormatInt(userID, 10)+"_"+strconv.FormatInt(tenantID, 10)+"_"+strconv.FormatInt(memberID, 10),
		constant.TenantRole)
	if e != nil {
		return e
	}
	out.ID = types.BigInt(userID)
	out.Name = user.Name.String()
	out.Avator = user.Avator.String()
	out.Phone = user.Phone.String()
	out.Pass = user.Pass.String()
	out.Token = tk
	out.AuthUrlList = authUrls
	return nil
}

// authUrlsByRole 根据角色编号查权限 URL 列表
func (a *account) authUrlsByRole(ctx context.Context, roleID int64) ([]string, error) {
	ras, _, e := dao.RoleAuth(db).List(ctx, db.Table(tblroleauth.TableName).
		Where(tblroleauth.RoleId.Eq(types.BigInt(roleID))))
	if e != nil {
		return nil, e
	}
	authIDs := make([]any, 0, len(ras))
	for _, ra := range ras {
		authIDs = append(authIDs, int64(ra.AuthId))
	}
	if len(authIDs) == 0 {
		return nil, nil
	}
	auths, _, e := dao.Auth(db).List(ctx, db.Table(tblauth.TableName).
		Where(tblauth.Id.In(authIDs...)))
	if e != nil {
		return nil, e
	}
	urls := make([]string, 0, len(auths))
	for _, au := range auths {
		if au.Url.String() != "" {
			urls = append(urls, au.Url.String())
		}
	}
	return urls, nil
}

// checkTenantUsable 校验租户是否可用（未锁定且未过期）
func checkTenantUsable(tn *do.Tenant) error {
	if tn.Status.Int8() == constant.TenantStatusLocked {
		return constant.ErrTenantLocked
	}
	if tn.Status.Int8() == constant.TenantStatusTrial &&
		!tn.TrialEnd.IsZero() && tn.TrialEnd.Time.Before(time.Now()) {
		return constant.ErrTenantExpired
	}
	return nil
}

// isTenantAdmin 判断当前用户是否为指定租户的管理员（role_id=1）
func isTenantAdmin(ctx context.Context, tenantID int64) bool {
	userID := lib.UserID(ctx)
	member, has, e := dao.Member(db).Get(ctx, db.Table(tblmember.TableName).
		Where(tblmember.UserId.Eq(types.BigInt(userID)),
			tblmember.TenantId.Eq(types.BigInt(tenantID)),
			tblmember.DelFlag.Eq(0)))
	if e != nil || !has {
		return false
	}
	return member.RoleId.Int64() == 1
}
