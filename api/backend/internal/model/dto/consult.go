package dto

import "time"

// ============ ConsultController 请求 ============

// @request
// PageConsultByKeyQuery 分页查询咨询请求
type PageConsultByKeyQuery struct {
	PageNum      *int    `json:"pageNum" valid:"required"` // 页码
	PageSize     *int    `json:"pageSize" valid:"required"` // 条数
	ConsultName  *string `json:"consultName"` // 咨询人姓名
	ConsultPhone *string `json:"consultPhone"` // 咨询人电话
	ElderName    *string `json:"elderName"` // 老人姓名
	ElderPhone   *string `json:"elderPhone"` // 老人电话
	StartTime    *string `json:"startTime"` // 开始时间
	EndTime      *string `json:"endTime"` // 结束时间
	SourceID     *int64  `json:"sourceId"` // 来源渠道编号
	StaffID      *int64  `json:"staffId"` // 接待人编号
}

// @request
// OperateConsultQuery 操作咨询请求
type OperateConsultQuery struct {
	ConsultID       *int64  `json:"consultId"` // 咨询人编号
	ElderID         *int64  `json:"elderId"` // 老人编号
	SourceID        *int64  `json:"sourceId" valid:"required"` // 来源渠道编号
	StaffID         *int64  `json:"staffId" valid:"required"` // 接待人编号
	ConsultName     *string `json:"consultName" valid:"required"` // 咨询人姓名
	ConsultPhone    *string `json:"consultPhone" valid:"required"` // 咨询人电话
	Relation        *string `json:"relation" valid:"required"` // 与老人关系
	ConsultDate     *string `json:"consultDate" valid:"required"` // 咨询日期
	ConsultContent  *string `json:"consultContent" valid:"required"` // 咨询内容
	ElderName       *string `json:"elderName" valid:"required"` // 老人姓名
	IDNum           *string `json:"idNum" valid:"required"` // 身份证号
	Age             *string `json:"age" valid:"required"` // 年龄
	Sex             *string `json:"sex" valid:"required"` // 性别(男/女)
	ElderPhone      *string `json:"elderPhone" valid:"required"` // 老人电话
	Address         *string `json:"address" valid:"required"` // 地址
}

// @request
// AddCommunicationRecordQuery 新增沟通记录请求
type AddCommunicationRecordQuery struct {
	ElderID             *int64  `json:"elderId" valid:"required"` // 老人编号
	CommunicationRecord *string `json:"communicationRecord" valid:"required"` // 沟通记录
	RecordDate          *string `json:"recordDate" valid:"required"` // 记录时间
}

// @request
// EditCommunicationRecordQuery 编辑沟通记录请求
type EditCommunicationRecordQuery struct {
	ID                  *int64  `json:"id" valid:"required"` // 沟通记录编号
	CommunicationRecord *string `json:"communicationRecord" valid:"required"` // 沟通记录
	RecordDate          *string `json:"recordDate" valid:"required"` // 记录时间
}

// @request
// PageCommunicationRecordQuery 分页查询沟通记录请求
type PageCommunicationRecordQuery struct {
	PageNum  *int    `json:"pageNum" valid:"required"` // 页码
	PageSize *int    `json:"pageSize" valid:"required"` // 条数
	ElderID  *int64  `json:"elderId" valid:"required"` // 老人编号
	Key      *string `json:"key"` // 关键词(沟通内容), 可选
}

// @request
// GetConsultByConsultIDAndElderIDQuery 按咨询人/老人编号获取咨询信息请求
type GetConsultByConsultIDAndElderIDQuery struct {
	ConsultID *int64 `json:"consultId" valid:"required"` // 咨询人编号
	ElderID   *int64 `json:"elderId" valid:"required"` // 老人编号
}

// GetConsultByConsultIdAndElderIdQuery 按咨询人/老人编号获取咨询信息请求（别名）
type GetConsultByConsultIdAndElderIdQuery = GetConsultByConsultIDAndElderIDQuery

// @request
// AddConsultQuery 新增咨询请求
type AddConsultQuery struct {
	ConsultID       *int64  `json:"consultId"` // 咨询人编号
	ElderID         *int64  `json:"elderId"` // 老人编号
	SourceID        *int64  `json:"sourceId" valid:"required"` // 来源渠道编号
	StaffID         *int64  `json:"staffId" valid:"required"` // 接待人编号
	ConsultName     *string `json:"consultName" valid:"required"` // 咨询人姓名
	ConsultPhone    *string `json:"consultPhone" valid:"required"` // 咨询人电话
	Relation        *string `json:"relation" valid:"required"` // 与老人关系
	ConsultDate     *string `json:"consultDate" valid:"required"` // 咨询日期
	ConsultContent  *string `json:"consultContent" valid:"required"` // 咨询内容
	ElderName       *string `json:"elderName" valid:"required"` // 老人姓名
	IDNum           *string `json:"idNum" valid:"required"` // 身份证号
	Age             *string `json:"age" valid:"required"` // 年龄
	Sex             *string `json:"sex" valid:"required"` // 性别(男/女)
	ElderPhone      *string `json:"elderPhone" valid:"required"` // 老人电话
	Address         *string `json:"address" valid:"required"` // 地址
}

// @request
// PageIntentByKeyQuery 分页查询意向客户请求
type PageIntentByKeyQuery struct {
	PageNum  *int    `json:"pageNum" valid:"required"` // 页码
	PageSize *int    `json:"pageSize" valid:"required"` // 条数
	Key      *string `json:"key"` // 关键字(姓名/电话), 可选
}

// ============ ConsultController 响应 ============

// @response
// PageConsultByKeyVO 分页查询咨询响应
type PageConsultByKeyVO struct {
	Rank
	ConsultID    int64     `json:"consultId"` // 咨询人编号
	ElderID      int64     `json:"elderId"` // 老人编号
	ConsultName  string    `json:"consultName"` // 咨询人姓名
	ConsultPhone string    `json:"consultPhone"` // 咨询人电话
	ElderName    string    `json:"elderName"` // 老人姓名
	ElderPhone   string    `json:"elderPhone"` // 老人电话
	Sex          string    `json:"sex"` // 性别
	Age          int       `json:"age"` // 年龄
	ConsultDate  time.Time `json:"consultDate"` // 咨询日期
	SourceName   string    `json:"sourceName"` // 来源渠道
	StaffName    string    `json:"staffName"` // 接待人姓名
}

// @response
// PageCommunicationRecordVO 分页查询沟通记录响应
type PageCommunicationRecordVO struct {
	Rank
	ID                 int64     `json:"id"` // id
	RecordDate         time.Time `json:"recordDate"` // 记录时间
	CommunicationRecord string   `json:"communicationRecord"` // 沟通记录
}

// @response
// GetConsultByConsultIDAndElderIDVO 按咨询人/老人编号获取咨询信息响应（继承 OperateConsultQuery）
type GetConsultByConsultIDAndElderIDVO struct {
	OperateConsultQuery
}