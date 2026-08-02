package constant

import "github.com/linbaozhong/gentity/pkg/types"

const (
	Authorization = "Authorization"
)

// 错误声明
var (
	ErrAuthorizationNotFound = types.NewError(400, "访问令牌缺失")
)

// 审计状态
const (
	AuditPass      = "通过"
	AuditNoPass    = "不通过"
	AuditStayAudit = "待审核"
	AuditNotPass   = "未通过"
	AuditHavePass  = "已通过"
)

// 床位状态
const (
	BedIdle      = "空闲"
	BedReserve   = "预定"
	BedEnter     = "入住"
	BedExitAudit = "退住审核"
)

// 收费方式
const (
	ChargeOnce  = "按次"
	ChargeMonth = "按月"
	ChargeAll   = "所有"
)

// 咨询/客户状态
const (
	CheckConsult     = "咨询中"
	CheckIntention   = "意向跟进"
	CheckReserve     = "预定"
	CheckEnter       = "入住"
	CheckExitAudit   = "退住审核"
	CheckExit        = "已退住"
	CheckFailure     = "已失效"
)

// 响应码
var (
	CodeSuccess = types.NewError(200, "操作成功")
	CodeError   = types.NewError(500, "操作失败")
)

// 消费类型
const (
	ConsumeNurse  = "护理"
	ConsumeDishes = "点餐"
)

// 异常错误码
var (
	// common
	ErrSysError            = types.NewError(500, "系统繁忙")
	ErrNotExist            = types.NewError(404, "资源不存在")
	ErrDataNotExist        = types.NewError(404, "数据不存在")
	ErrMethodError         = types.NewError(500, "请求方式错误")
	ErrParamError          = types.NewError(500, "请求参数异常")
	ErrAccessError         = types.NewError(500, "非法访问")
	ErrTokenError          = types.NewError(500, "令牌无效")
	ErrCertificationError  = types.NewError(500, "认证失败")
	ErrAuthError           = types.NewError(500, "权限不足")
	// 忘记密码
	ErrAccountFormatError  = types.NewError(500, "账号格式有误")
	ErrAccountNotRegister  = types.NewError(500, "账号未注册")
	ErrPassSame            = types.NewError(500, "新密码与原密码相同")
	ErrRepeatSendCode      = types.NewError(500, "请勿重复发送验证码")
	ErrCodeExpire          = types.NewError(500, "验证码过期")
	ErrCodeError           = types.NewError(500, "验证码错误")
	// 修改密码
	ErrOldPassError        = types.NewError(500, "原密码错误")
	// 上传文件
	ErrUploadSuffixError   = types.NewError(500, "不支持的文件后缀")
	// 老人
	ErrElderAlreadyDelete  = types.NewError(500, "该老人已被删除")
	ErrIdNumRepeat         = types.NewError(500, "老人身份证号已存在")
	ErrBalanceDeficiency   = types.NewError(500, "老人余额不足")
	// 标签
	ErrLabelTypeRepeat     = types.NewError(500, "标签分类已存在")
	ErrLabelTypeOut        = types.NewError(500, "标签分类总数超过限制")
	ErrLabelNotNull        = types.NewError(500, "该标签分类存在标签，删除失败")
	ErrLabelRepeat         = types.NewError(500, "标签已存在")
	ErrLabelOut            = types.NewError(500, "该分类标签总数超过限制")
	// 来源渠道
	ErrSourceRepeat        = types.NewError(500, "来源渠道已存在")
	// 活动分类
	ErrActiveTypeRepeat    = types.NewError(500, "活动分类已存在")
	// 活动
	ErrActiveRepeat        = types.NewError(500, "活动名称已存在")
	// 节点
	ErrNodeMarkNotExist    = types.NewError(500, "该节点标记不存在")
	ErrNodeBedNotIdle      = types.NewError(500, "该节点有床位被占用，删除失败")
	// 楼栋
	ErrBuildingRepeat      = types.NewError(500, "楼栋已存在")
	ErrBuildingOut         = types.NewError(500, "楼栋总数超过限制")
	// 楼层
	ErrFloorRepeat         = types.NewError(500, "楼层已存在")
	ErrFloorOut            = types.NewError(500, "楼层总数超过限制")
	// 房间类型
	ErrRoomTypeRepeat      = types.NewError(500, "房间类型已存在")
	// 房间
	ErrRoomRepeat          = types.NewError(500, "房间已存在")
	ErrRoomOut             = types.NewError(500, "房间总数超过限制")
	// 床位
	ErrBedNull             = types.NewError(500, "床位不存在")
	ErrBedRepeat           = types.NewError(500, "床位已存在")
	ErrBedOut              = types.NewError(500, "床位总数超过限制")
	ErrBedOccupy           = types.NewError(500, "该床位被占用")
	ErrBedNotIdle          = types.NewError(500, "该床位被占用，删除失败")
	// 意向客户
	ErrAlreadyIntention    = types.NewError(500, "该老人已是意向客户")
	ErrAlreadyReserve      = types.NewError(500, "该老人已预定")
	ErrAlreadyEnter        = types.NewError(500, "该老人已入住")
	// 回访计划
	ErrVisitPlanRepeat     = types.NewError(500, "回访计划标题已存在")
	// 预定
	ErrRefundRepeat        = types.NewError(500, "请勿重复退款")
	ErrDueDateExpire       = types.NewError(500, "预定已过期")
	// 预存充值
	ErrNotEnter            = types.NewError(500, "该老人非入住状态，不予充值")
	// 员工
	ErrPhoneRepeat         = types.NewError(500, "该手机号已存在")
	ErrEmailRepeat         = types.NewError(500, "该邮箱已存在")
	// 服务
	ErrServiceTypeRepeat   = types.NewError(500, "该服务类型已存在")
	ErrServiceTypeOut      = types.NewError(500, "服务类型总数超过限制")
	ErrServiceNotNull      = types.NewError(500, "该服务类型存在服务，删除失败")
	ErrServiceRepeat       = types.NewError(500, "该服务已存在")
	ErrServiceOut          = types.NewError(500, "该类型服务总数超过限制")
	// 护理等级
	ErrNurseGradeRepeat    = types.NewError(500, "该护理等级已存在")
	ErrNurseGradeSelected  = types.NewError(500, "该护理等级已被选择，删除失败")
	// 菜品
	ErrDishesTypeRepeat    = types.NewError(500, "该菜品分类已存在")
	ErrDishesTypeOut       = types.NewError(500, "菜品分类总数超过限制")
	ErrDishesNotNull       = types.NewError(500, "该菜品分类存在菜品，删除失败")
	ErrDishesRepeat        = types.NewError(500, "该菜品已存在")
	// 套餐
	ErrSetRepeat           = types.NewError(500, "该套餐已存在")
	ErrSetSelected         = types.NewError(500, "该套餐已被选择，删除失败")
	// 订单
	ErrOrderSuccess        = types.NewError(500, "该订单已完成")
	// 物资
	ErrMaterialTypeRepeat  = types.NewError(500, "该物资分类已存在")
	ErrMaterialTypeOut     = types.NewError(500, "物资分类总数超过限制")
	ErrMaterialNotNull     = types.NewError(500, "该物资分类存在物资，删除失败")
	ErrMaterialRepeat      = types.NewError(500, "该物资已存在")
	// 仓库
	ErrWarehouseRepeat     = types.NewError(500, "该仓库已存在")
	ErrWarehouseNotNull    = types.NewError(500, "该仓库存有物资，删除失败")
	// 出入库管理
	ErrAuditResultError    = types.NewError(500, "审核结果非法")
	ErrAuditRepeat         = types.NewError(500, "请勿重复审核")
	ErrOutboundNumError    = types.NewError(500, "出库物资数量超过库存量")
	ErrDelRepeat           = types.NewError(500, "请勿重复删除")
	// 外出登记
	ErrOutwardRepeat       = types.NewError(500, "请勿重复登记")
	ErrRecordAlreadyDelete = types.NewError(500, "该登记已被删除")
	ErrAlreadyReturn       = types.NewError(500, "请勿重复登记返回")
	// 来访登记
	ErrAlreadyLeave        = types.NewError(500, "来访人员已登记离开")
	// 退住申请
	ErrElderNotEnter       = types.NewError(500, "老人暂未入住")
	ErrApplyRepeat         = types.NewError(500, "请勿重复申请")
	// 退住审核
	ErrElderNotExitAudit   = types.NewError(500, "老人暂未申请退住")
)

// 节点标记
const (
	MarkBuilding = "楼栋"
	MarkFloor    = "楼层"
	MarkRoom     = "房间"
)

// 接收方类型
const (
	RecipientElder = "老人"
	RecipientStaff = "员工"
)

// 来访状态
const (
	VisitStayLeave   = "待离开"
	VisitAlreadyLeave = "已离开"
)

// 是否
const (
	YesNoYes = "是"
	YesNoNo  = "否"
)

// 系统常量
const (
	Subject      = "敬老院管理系统"
	Developer    = "@developer "
	EmperorWen   = "Xxx"
	AnonymousUser = "anonymousUser"

	// AES
	Iv      = "1234567890123456"
	AesKey  = "1234567890654321"
	AesType = "AES/CBC/PKCS5Padding"

	// JWT
	TokenSubject = Subject
	TokenIssure  = "Xxx"
	TokenSecret  = "pX2~tQ4*nP6_gJ0%sY8/iY6.kC3|oE2$nT3,"

	// EMAIL
	MailHost = "smtp.qq.com"
	Mail     = "xxxxxxxxxx@qq.com"
	Pass     = "irzvcvfqkdaecigi"

	// 过期时间(毫秒)
	ExpireTime = int64(1000 * 60 * 60 * 24)

	// 正则
	EmailRegular = "[a-zA-Z0-9]+@[a-zA-Z0-9]+\\.[a-zA-Z0-9]+"
	PhoneRegular = "^1[35789]\\d{9}$"

	// REDIS
	LoginRedis = "login:"

	// FILE
	StrDownload    = "download"
	StrUploadImg   = "upload/img"
	StrUploadFile  = "upload/file"
	StrUploadVideo = "upload/video"

	// OTHER
	TotalLimit = int64(10)
	Comma      = ","
)
