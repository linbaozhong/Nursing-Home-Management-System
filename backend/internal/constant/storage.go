package constant

import "github.com/linbaozhong/gentity/pkg/types"

// 出入库管理
var (
	ErrAuditResultError = types.NewError(500, "审核结果非法")
	ErrAuditRepeat      = types.NewError(500, "请勿重复审核")
	ErrOutboundNumError = types.NewError(500, "出库物资数量超过库存量")
	ErrDelRepeat        = types.NewError(500, "请勿重复删除")
)

// 外出登记
var (
	ErrOutwardRepeat       = types.NewError(500, "请勿重复登记")
	ErrRecordAlreadyDelete = types.NewError(500, "该登记已被删除")
	ErrAlreadyReturn       = types.NewError(500, "请勿重复登记返回")
)

// 来访登记
var (
	ErrAlreadyLeave = types.NewError(500, "来访人员已登记离开")
)

// 退住申请
var (
	ErrApplyRepeat = types.NewError(500, "请勿重复申请")
)
