package dto

import (
	"github.com/linbaozhong/gentity/pkg/types"
	"time"
)

// ============ ConsultController 请求 ============

// PageConsultByKeyReq 分页查询咨询请求
// @request
type PageConsultByKeyReq struct {
	PageNum      *int       `json:"page_num" valid:"required"`  // 页码
	PageSize     *int       `json:"page_size" valid:"required"` // 条数
	ConsultName  *string    `json:"consult_name"`               // 咨询人姓名
	ConsultPhone *string    `json:"consult_phone"`              // 咨询人电话
	ElderName    *string    `json:"elder_name"`                 // 老人姓名
	ElderPhone   *string    `json:"elder_phone"`                // 老人电话
	StartTime    *time.Time `json:"start_time"`                 // 开始时间
	EndTime      *time.Time `json:"end_time"`                   // 结束时间
	SourceID     *int64     `json:"source_id"`                  // 来源渠道编号
	StaffID      *int64     `json:"staff_id"`                   // 接待人编号
}

// OperateConsultReq 操作咨询请求
// @request
type OperateConsultReq struct {
	ConsultID      *int64     `json:"consult_id"`                       // 咨询人编号
	ElderID        *int64     `json:"elder_id"`                         // 老人编号
	SourceID       *int64     `json:"source_id" valid:"required"`       // 来源渠道编号
	StaffID        *int64     `json:"staff_id" valid:"required"`        // 接待人编号
	ConsultName    *string    `json:"consult_name" valid:"required"`    // 咨询人姓名
	ConsultPhone   *string    `json:"consult_phone" valid:"required"`   // 咨询人电话
	Relation       *string    `json:"relation" valid:"required"`        // 与老人关系
	ConsultDate    *time.Time `json:"consult_date" valid:"required"`    // 咨询日期
	ConsultContent *string    `json:"consult_content" valid:"required"` // 咨询内容
	ElderName      *string    `json:"elder_name" valid:"required"`      // 老人姓名
	IDNum          *string    `json:"id_num" valid:"required"`          // 身份证号
	Age            *string    `json:"age" valid:"required"`             // 年龄
	Sex            *string    `json:"sex" valid:"required"`             // 性别(男/女)
	ElderPhone     *string    `json:"elder_phone" valid:"required"`     // 老人电话
	Address        *string    `json:"address" valid:"required"`         // 地址
}

// AddCommunicationRecordReq 新增沟通记录请求
// @request
type AddCommunicationRecordReq struct {
	ElderID             *int64     `json:"elder_id" valid:"required"`             // 老人编号
	CommunicationRecord *string    `json:"communication_record" valid:"required"` // 沟通记录
	RecordDate          *time.Time `json:"record_date" valid:"required"`          // 记录时间
}

// EditCommunicationRecordReq 编辑沟通记录请求
// @request
type EditCommunicationRecordReq struct {
	ID                  *int64     `json:"id" valid:"required"`                   // 沟通记录编号
	CommunicationRecord *string    `json:"communication_record" valid:"required"` // 沟通记录
	RecordDate          *time.Time `json:"record_date" valid:"required"`          // 记录时间
}

// PageCommunicationRecordReq 分页查询沟通记录请求
// @request
type PageCommunicationRecordReq struct {
	PageNum  *int    `json:"page_num" valid:"required"`  // 页码
	PageSize *int    `json:"page_size" valid:"required"` // 条数
	ElderID  *int64  `json:"elder_id" valid:"required"`  // 老人编号
	Key      *string `json:"key"`                        // 关键词(沟通内容), 可选
}

// GetConsultByConsultIDAndElderIDReq 按咨询人/老人编号获取咨询信息请求
// @request
type GetConsultByConsultIDAndElderIDReq struct {
	ConsultID *int64 `json:"consult_id" valid:"required"` // 咨询人编号
	ElderID   *int64 `json:"elder_id" valid:"required"`   // 老人编号
}

// GetConsultByConsultIdAndElderIdQuery 按咨询人/老人编号获取咨询信息请求（别名）
type GetConsultByConsultIdAndElderIdQuery = GetConsultByConsultIDAndElderIDReq

// AddConsultReq 新增咨询请求
// @request
type AddConsultReq struct {
	ConsultID      *int64     `json:"consult_id"`                       // 咨询人编号
	ElderID        *int64     `json:"elder_id"`                         // 老人编号
	SourceID       *int64     `json:"source_id" valid:"required"`       // 来源渠道编号
	StaffID        *int64     `json:"staff_id" valid:"required"`        // 接待人编号
	ConsultName    *string    `json:"consult_name" valid:"required"`    // 咨询人姓名
	ConsultPhone   *string    `json:"consult_phone" valid:"required"`   // 咨询人电话
	Relation       *string    `json:"relation" valid:"required"`        // 与老人关系
	ConsultDate    *time.Time `json:"consult_date" valid:"required"`    // 咨询日期
	ConsultContent *string    `json:"consult_content" valid:"required"` // 咨询内容
	ElderName      *string    `json:"elder_name" valid:"required"`      // 老人姓名
	IDNum          *string    `json:"id_num" valid:"required"`          // 身份证号
	Age            *string    `json:"age" valid:"required"`             // 年龄
	Sex            *string    `json:"sex" valid:"required"`             // 性别(男/女)
	ElderPhone     *string    `json:"elder_phone" valid:"required"`     // 老人电话
	Address        *string    `json:"address" valid:"required"`         // 地址
}

// PageIntentByKeyReq 分页查询意向客户请求
// @request
type PageIntentByKeyReq struct {
	PageNum  *int    `json:"page_num" valid:"required"`  // 页码
	PageSize *int    `json:"page_size" valid:"required"` // 条数
	Key      *string `json:"key"`                        // 关键字(姓名/电话), 可选
}

// ============ ConsultController 响应 ============

// PageConsultByKeyResp 分页查询咨询响应
// @response
type PageConsultByKeyResp struct {
	ConsultID    types.BigInt `json:"consult_id"`    // 咨询人编号
	ElderID      types.BigInt `json:"elder_id"`      // 老人编号
	ConsultName  string       `json:"consult_name"`  // 咨询人姓名
	ConsultPhone string       `json:"consult_phone"` // 咨询人电话
	ElderName    string       `json:"elder_name"`    // 老人姓名
	ElderPhone   string       `json:"elder_phone"`   // 老人电话
	Sex          string       `json:"sex"`           // 性别
	Age          int          `json:"age"`           // 年龄
	ConsultDate  time.Time    `json:"consult_date"`  // 咨询日期
	SourceName   string       `json:"source_name"`   // 来源渠道
	StaffName    string       `json:"staff_name"`    // 接待人姓名
}

// PageCommunicationRecordResp 分页查询沟通记录响应
// @response
type PageCommunicationRecordResp struct {
	ID                  types.BigInt `json:"id"`                   // id
	RecordDate          time.Time    `json:"record_date"`          // 记录时间
	CommunicationRecord string       `json:"communication_record"` // 沟通记录
}

// GetConsultByConsultIDAndElderIDResp 按咨询人/老人编号获取咨询信息响应（继承 OperateConsultReq）
// @response
type GetConsultByConsultIDAndElderIDResp struct {
	OperateConsultReq
}
