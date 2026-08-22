package dto

import (
	"github.com/linbaozhong/gentity/pkg/validator"
)

// 消费 gentity_dto.gen.go 中导入的 validator 包（该包级 import 需被引用）
var _ = validator.IsMobile

// ============ FamilyController 请求 ============

// FamilySendCodeReq 发送家属注册/绑定验证码请求
// @request
type FamilySendCodeReq struct {
	Phone *string `json:"phone" valid:"required"` // 家属手机号
}

// RegisterBindReq 家属注册并绑定老人请求
// @request
type RegisterBindReq struct {
	Phone     *string `json:"phone" valid:"required"`       // 家属手机号
	Code      *string `json:"code" valid:"required"`        // 短信验证码
	Password  *string `json:"password" valid:"required"`    // 登录密码
	ElderName *string `json:"elder_name" valid:"required"`  // 老人姓名
	IdNumTail *string `json:"id_num_tail" valid:"required"` // 身份证后4位
	Relation  *string `json:"relation"`                     // 与老人关系
}

// FamilyLoginReq 家属登录请求
// @request
type FamilyLoginReq struct {
	Phone    *string `json:"phone" valid:"required"`    // 家属手机号
	Password *string `json:"password" valid:"required"` // 登录密码
}

// BindElderReq 已注册家属绑定更多老人请求
// @request
type BindElderReq struct {
	Phone     *string `json:"phone" valid:"required"`       // 当前家属手机号
	Code      *string `json:"code" valid:"required"`        // 短信验证码
	ElderName *string `json:"elder_name" valid:"required"`  // 老人姓名
	IdNumTail *string `json:"id_num_tail" valid:"required"` // 身份证后4位
	Relation  *string `json:"relation"`                     // 与老人关系
}

// FamilyMyEldersReq 我的老人请求
// @request
type FamilyMyEldersReq struct {
	Phone *string `json:"phone" valid:"required"` // 当前家属手机号
}

// ============ FamilyController 响应 ============

// FamilyElderResp 家属绑定老人响应
// @response
type FamilyElderResp struct {
	ElderID  int64  `json:"elder_id"` // 老人编号
	Name     string `json:"name"`     // 老人姓名
	Relation string `json:"relation"` // 与老人关系
}

// FamilyLoginResp 家属登录响应
// @response
type FamilyLoginResp struct {
	Token     string             `json:"token"`      // 登录令牌
	ElderList []*FamilyElderResp `json:"elder_list"` // 绑定老人列表
}

// FamilyMyEldersResp 我的老人列表响应
// @response
type FamilyMyEldersResp struct {
	List []*FamilyElderResp `json:"list"` // 绑定老人列表
}

// ============ FamilyRechargeController 请求 ============

// RechargeUnifiedOrderReq 家属充值统一下单请求
// @request
type RechargeUnifiedOrderReq struct {
	Phone   *string `json:"phone" valid:"required"`         // 当前家属手机号
	ElderID *int64  `json:"elder_id" valid:"required"`      // 充值到哪位老人账户
	Amount  *int64  `json:"amount" valid:"required,min(0)"` // 金额（元）
}

// BindOpenidReq 家属绑定微信openid请求
// @request
type BindOpenidReq struct {
	Phone *string `json:"phone" valid:"required"` // 当前家属手机号
	Code  *string `json:"code" valid:"required"`  // wx.login 返回的 code
}

// ============ FamilyRechargeController 响应 ============

// RechargeUnifiedOrderResp 家属充值统一下单响应
// @response
type RechargeUnifiedOrderResp struct {
	AppId     string `json:"app_id"`     // 微信 AppID
	TimeStamp string `json:"time_stamp"` // 时间戳
	NonceStr  string `json:"nonce_str"`  // 随机串
	Package   string `json:"package"`    // 预支付交易会话标识
	SignType  string `json:"sign_type"`  // 签名类型
	PaySign   string `json:"pay_sign"`   // 签名
	OrderNo   string `json:"order_no"`   // 商户订单号
}

// BindOpenidResp 绑定微信openid响应
// @response
type BindOpenidResp struct {
	Openid string `json:"openid"` // 微信 openid
}

// ============ 微信支付回调（struct 由本文件定义，方法由 gentity_dto.gen.go 生成） ============

// WechatPayResource 微信支付回调加密报文
// @request
type WechatPayResource struct {
	Algorithm      *string `json:"algorithm"`       // 加密算法
	Ciphertext     *string `json:"ciphertext"`      // 密文
	AssociatedData *string `json:"associated_data"` // 附加数据
	Nonce          *string `json:"nonce"`           // 随机串
	OriginalType   *string `json:"original_type"`   // 原始类型
}

// WechatPayNotifyReq 微信支付结果通知
// @request
type WechatPayNotifyReq struct {
	Id           *string            `json:"id"`            // 通知 id
	CreateTime   *string            `json:"create_time"`   // 通知创建时间
	ResourceType *string            `json:"resource_type"` // 资源类型
	EventType    *string            `json:"event_type"`    // 事件类型
	Summary      *string            `json:"summary"`       // 摘要
	Resource     *WechatPayResource `json:"resource"`      // 加密资源
}
