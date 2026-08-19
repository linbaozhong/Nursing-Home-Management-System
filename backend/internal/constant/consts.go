// Package constant 业务常量与错误码定义。
//
// 原单一 consts.go 已按业务域拆分到同包下的多个文件，便于检索：
//   - common.go        : 授权、响应码、通用错误、是否(Y/N)
//   - audit.go         : 审计状态
//   - building.go      : 楼栋/楼层/房间/床位/节点
//   - elder.go         : 老人、咨询状态、标签、员工
//   - finance.go       : 收费方式、预定/退款、预存充值
//   - activity.go      : 活动/来源/消费/回访/密码/上传
//   - service_goods.go : 服务/护理/菜品/套餐/订单/物资/仓库
//   - storage.go       : 出入库/外出/来访/退住
//   - system.go        : 系统通用常量、接收方/来访类型定义
//   - crypto.go        : AES 加密配置
//   - jwt.go           : JWT 配置与过期时间
//   - email.go         : 邮箱服务配置与邮箱正则
//   - file.go          : 文件存储路径常量
//
// 调用方仍统一使用 constant.Xxx，拆分后无需改动。
package constant
