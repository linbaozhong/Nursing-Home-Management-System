package dto

import "time"

// ============ ActiveController 请求 ============

// @request
// PageActiveByKeyQuery 分页查询活动请求
type PageActiveByKeyQuery struct {
	PageNum   *int    `json:"pageNum" valid:"required"` // 页码
	PageSize  *int    `json:"pageSize" valid:"required"` // 条数
	TypeID    *int64  `json:"typeId"` // 活动类型编号
	Name      *string `json:"name"` // 活动名称
	StartTime *string `json:"startTime"` // 开始时间
	EndTime   *string `json:"endTime"` // 结束时间
}

// @request
// OperateActiveQuery 操作活动请求
type OperateActiveQuery struct {
	ID            *int64    `json:"id"` // id
	TypeID        *int64    `json:"typeId" valid:"required"` // 活动类型编号
	Theme         *string   `json:"theme" valid:"required"` // 活动主题
	Name          *string   `json:"name" valid:"required"` // 活动名称
	Content       *string   `json:"content" valid:"required"` // 活动内容
	Address       *string   `json:"address" valid:"required"` // 活动地点
	Organizer     *string   `json:"organizer" valid:"required"` // 组织者姓名
	Phone         *string   `json:"phone" valid:"required"` // 组织者电话
	ActiveDate    *string   `json:"activeDate" valid:"required"` // 活动日期
	ActivePicture *string   `json:"activePicture" valid:"required"` // 活动图片
	ElderIDList   []int64  `json:"elderIdList" valid:"required"` // 参加活动老人编号列表
}

// ============ ActiveController 响应 ============

// @response
// PageActiveByKeyVO 分页查询活动响应
type PageActiveByKeyVO struct {
	Rank
	ID            int64     `json:"id"` // id
	TypeName      string    `json:"typeName"` // 活动分类名称
	Theme         string    `json:"theme"` // 活动主题
	Name          string    `json:"name"` // 活动名称
	Content       string    `json:"content"` // 活动内容
	Address       string    `json:"address"` // 活动地点
	Organizer     string    `json:"organizer"` // 组织者姓名
	Phone         string    `json:"phone"` // 组织者电话
	ActiveDate    string    `json:"activeDate"` // 活动日期
	ActivePicture string    `json:"activePicture"` // 活动图片
}

// @response
// GetActiveByIDVO 根据编号获取活动响应（继承 OperateActiveQuery）
type GetActiveByIDVO struct {
	OperateActiveQuery
	ParticipateElderVOList []ParticipateElderVO `json:"participateElderVoList"` // 参加活动老人列表
}

// @response
// ParticipateElderVO 参加活动老人响应
type ParticipateElderVO struct {
	ID    int64  `json:"id"` // id
	Name  string `json:"name"` // 老人姓名
	Phone string `json:"phone"` // 老人电话
}

var _ = time.Time{}