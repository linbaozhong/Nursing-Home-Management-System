package service

import (
	"api/internal/constant"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblelder"
	"api/internal/model/define/table/tblfamilyrecharge"
	"api/internal/model/dto"
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/linbaozhong/gentity/pkg/ace/dialect"
	"github.com/linbaozhong/gentity/pkg/types"
)

var FamilyRechargeSvc = &familyRecharge{}

type familyRecharge struct{}

// ============ 微信支付配置（占位符，接入配置中心/环境变量后替换） ============
const (
	wxAppIDPlaceholder     = "REPLACE_WX_APPID"      // 小程序 appid
	wxSecretPlaceholder    = "REPLACE_WX_SECRET"     // 小程序 secret
	wxMchIDPlaceholder     = "REPLACE_WX_MCH_ID"     // 商户号
	wxAPIKeyPlaceholder    = "REPLACE_WX_API_KEY"    // 商户 API 密钥
	wxAPIv3KeyPlaceholder  = "REPLACE_WX_APIV3_KEY"  // APIv3 密钥（回调解密）
	wxNotifyURLPlaceholder = "REPLACE_WX_NOTIFY_URL" // 支付回调地址
)

// RechargeUnifiedOrder 家属充值统一下单
func (s *familyRecharge) RechargeUnifiedOrder(ctx context.Context, in *dto.RechargeUnifiedOrderQuery, out *dto.RechargeUnifiedOrderVO) error {
	if in.Phone == nil || in.ElderID == nil || in.Amount == nil {
		return constant.ErrParamError
	}
	phone := *in.Phone
	elderID := *in.ElderID
	amount := *in.Amount

	// 校验老人归属
	ok, e := Family.IsMyElder(ctx, phone, elderID)
	if e != nil {
		return e
	}
	if !ok {
		return constant.ErrFamilyNoElder
	}
	// 取家属 openid（登录验证已绑定）
	acc, ok, e := dao.FamilyAccount(db).GetByPhone(ctx, phone)
	if e != nil {
		return e
	}
	if !ok || acc.Openid.String() == "" {
		return types.NewError(400, "请先在微信中绑定账号后充值")
	}
	if amount <= 0 {
		return types.NewError(400, "充值金额不合法")
	}

	orderNo := genRechargeOrderNo()
	amountFen := amount * 100

	// 落库订单（状态 0=待支付）
	if _, e = dao.FamilyRecharge(db).Insert(ctx,
		tblfamilyrecharge.OrderNo.Set(types.String(orderNo)),
		tblfamilyrecharge.Phone.Set(types.String(phone)),
		tblfamilyrecharge.ElderId.Set(types.BigInt(elderID)),
		tblfamilyrecharge.Amount.Set(types.Int64(amountFen)),
		tblfamilyrecharge.Status.Set(types.Int8(0)),
	); e != nil {
		return e
	}

	// 调用微信支付下单（B 方案：引入 silenceper/wechat/pay，配置占位符）
	payParams, e := placeWxOrder(placeWxOrderIn{
		Description: "敬老院家属预存充值",
		OutTradeNo:  orderNo,
		NotifyURL:   wxNotifyURLPlaceholder,
		Amount:      amountFen,
		Openid:      acc.Openid.String(),
	})
	if e != nil {
		return e
	}
	// 回填 prepay_id
	_, _ = dao.FamilyRecharge(db).Update(ctx,
		[]dialect.Setter{tblfamilyrecharge.PrepayId.Set(types.String(payParams.Package))},
		tblfamilyrecharge.OrderNo.Eq(types.String(orderNo)),
	)

	out.AppId = wxAppIDPlaceholder
	out.TimeStamp = strconv.FormatInt(time.Now().Unix(), 10)
	out.NonceStr = payParams.NonceStr
	out.Package = payParams.Package
	out.SignType = "RSA"
	out.PaySign = payParams.PaySign
	out.OrderNo = orderNo
	return nil
}

// PayNotify 微信支付结果回调
func (s *familyRecharge) PayNotify(ctx context.Context, in *dto.WechatPayNotifyQuery, out *dto.EmptyResp) error {
	// 解密回调（B 方案：用 APIv3 密钥解密 resource）
	plain, e := decryptWxNotify(wxAPIv3KeyPlaceholder, in.Resource)
	if e != nil {
		return e
	}
	var tx struct {
		OutTradeNo string `json:"out_trade_no"`
		TradeState string `json:"trade_state"`
	}
	if e = json.Unmarshal([]byte(plain), &tx); e != nil {
		return e
	}
	if tx.TradeState != "SUCCESS" {
		return nil
	}
	order, ok, e := dao.FamilyRecharge(db).GetByOrderNo(ctx, tx.OutTradeNo)
	if e != nil {
		return e
	}
	if !ok {
		return errors.New("订单不存在")
	}
	if order.Status.Int8() == 1 {
		return nil // 已处理
	}
	// 标记已支付
	if _, e = dao.FamilyRecharge(db).UpdateStatus(ctx, tx.OutTradeNo, 1); e != nil {
		return e
	}
	// 给老人余额加款（amount 单位分 -> 元），先读后写保证增量正确
	elder, ok, e := dao.Elder(db).GetByID(ctx, order.ElderId, tblelder.Id, tblelder.Balance)
	if e != nil {
		return e
	}
	if !ok {
		return errors.New("老人不存在")
	}
	yuan := order.Amount.Int64() / 100
	cur, _ := strconv.ParseFloat(elder.Balance.String(), 64)
	newBalance := int64(cur) + yuan
	_, e = dao.Elder(db).UpdateById(ctx, order.ElderId, tblelder.Balance.Set(types.Money(strconv.FormatInt(newBalance, 10))))
	return e
}

// ============ 微信支付占位实现（B 方案：接入 silenceper/wechat/pay 后替换） ============

type placeWxOrderIn struct {
	Description string
	OutTradeNo  string
	NotifyURL   string
	Amount      int64
	Openid      string
}

type placeWxOrderOut struct {
	Package  string
	NonceStr string
	PaySign  string
}

// placeWxOrder 调用微信支付 JSAPI 下单。
// TODO: 接入 github.com/silenceper/wechat/v2/pay，用商户配置（wxMchIDPlaceholder 等）真实下单，
// 返回 prepay_id 并按小程序支付签名规则生成 PaySign。当前为占位实现。
func placeWxOrder(in placeWxOrderIn) (*placeWxOrderOut, error) {
	// 占位：直接返回模拟参数，保证流程可跑；接入真实支付后删除此占位。
	return &placeWxOrderOut{
		Package:  "prepay_id=PLACEHOLDER_" + in.OutTradeNo,
		NonceStr: "nonce_placeholder",
		PaySign:  "sign_placeholder",
	}, nil
}

// decryptWxNotify 解密微信支付回调报文。
// TODO: 接入真实 AES-GCM 解密（APIv3 密钥）。当前为占位实现。
func decryptWxNotify(apiV3Key string, resource dto.WechatPayResource) (string, error) {
	if resource.Ciphertext == nil {
		return "", errors.New("回调解密失败")
	}
	return *resource.Ciphertext, nil
}

// genRechargeOrderNo 生成充值订单号
func genRechargeOrderNo() string {
	return "FR" + time.Now().Format("20060102150405") + strconv.Itoa(int(time.Now().UnixNano()%1000))
}
