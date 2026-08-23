package dto

import (
	"github.com/linbaozhong/gentity/pkg/types"
	"time"
)

// ============ ActiveController 请求 ============

// PageActiveByKeyReq 分页查询活动请求
// @request
type PageActiveByKeyReq struct {
	PageNum   *int       `json:"page_num" valid:"required"`  // 页码
	PageSize  *int       `json:"page_size" valid:"required"` // 条数
	TypeID    *int64     `json:"type_id"`                    // 活动类型编号
	Name      *string    `json:"name"`                       // 活动名称
	StartTime *time.Time `json:"start_time"`                 // 开始时间
	EndTime   *time.Time `json:"end_time"`                   // 结束时间
}

// OperateActiveReq 操作活动请求
// @request
type OperateActiveReq struct {
	ID            *int64     `json:"id"`                              // id
	TypeID        *int64     `json:"type_id" valid:"required"`        // 活动类型编号
	Theme         *string    `json:"theme" valid:"required"`          // 活动主题
	Name          *string    `json:"name" valid:"required"`           // 活动名称
	Content       *string    `json:"content" valid:"required"`        // 活动内容
	Address       *string    `json:"address" valid:"required"`        // 活动地点
	Organizer     *string    `json:"organizer" valid:"required"`      // 组织者姓名
	Phone         *string    `json:"phone" valid:"required"`          // 组织者电话
	ActiveDate    *time.Time `json:"active_date" valid:"required"`    // 活动日期
	ActivePicture *string    `json:"active_picture" valid:"required"` // 活动图片
	ElderIDList   []int64    `json:"elder_id_list" valid:"required"`  // 参加活动老人编号列表
}

// ============ ActiveController 响应 ============

// PageActiveByKeyResp 分页查询活动响应
// @response
type PageActiveByKeyResp struct {
	ID            types.BigInt `json:"id"`             // id
	TypeName      string       `json:"type_name"`      // 活动分类名称
	Theme         string       `json:"theme"`          // 活动主题
	Name          string       `json:"name"`           // 活动名称
	Content       string       `json:"content"`        // 活动内容
	Address       string       `json:"address"`        // 活动地点
	Organizer     string       `json:"organizer"`      // 组织者姓名
	Phone         string       `json:"phone"`          // 组织者电话
	ActiveDate    time.Time    `json:"active_date"`    // 活动日期
	ActivePicture string       `json:"active_picture"` // 活动图片
}

// GetActiveByIDResp 根据编号获取活动响应（继承 OperateActiveReq）
// @response
type GetActiveByIDResp struct {
	OperateActiveReq
	ParticipateElderRespList []ParticipateElderResp `json:"participate_elder_vo_list"` // 参加活动老人列表
}

// ParticipateElderResp 参加活动老人响应
// @response
type ParticipateElderResp struct {
	ID    types.BigInt `json:"id"`    // id
	Name  string       `json:"name"`  // 老人姓名
	Phone string       `json:"phone"` // 老人电话
}

var _ = time.Time{}
