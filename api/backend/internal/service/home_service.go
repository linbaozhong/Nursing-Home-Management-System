package service

import (
	"context"
	"time"

	"api/internal/model/define/dao"
	"api/internal/model/do"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/types"
)

type home struct{}

var Home = &home{}

// dayRange 返回今天 [00:00:00, 23:59:59] 的时间区间
func dayRange() (start, end types.Time) {
	now := time.Now()
	y, m, d := now.Date()
	loc := now.Location()
	start = types.Time(time.Date(y, m, d, 0, 0, 0, 0, loc))
	end = types.Time(time.Date(y, m, d, 23, 59, 59, 0, loc))
	return
}

// TodayOverview 今日概览
// 对应 Java: HomeServiceImpl.todayOverview
// 返回: 今日咨询数(consult 今日新增)、今日预订数(reserve 今日新增)、今日合同数(contract 今日新增)、
//
//	今日退住申请(retreat_apply 今日新增)、在住老人(elder 在住)、可用床位(bed 空闲)、销售中套餐数、员工数
//
// todo: 各计数用 dao.X(db).Count(ctx, ace.Where(时间区间/状态)); 组装 TodayOverviewVO, 赋值 out
func (h *home) TodayOverview(ctx context.Context, in *dto.EmptyReq, out *dto.EmptyResp) error {
	s, e := dayRange()
	_ = s
	_ = e
	// todo: 见上方注释, 完成各维度计数并组装 TodayOverviewVO
	return nil
}

// AvailableBed 可用床位
// 对应 Java: HomeServiceImpl.availableBed
// 返回: 楼栋列表(含每层空闲床位数)、楼层总数、房间总数、床位总数、空闲床位、占用床位
//
// todo: 1) 统计 floor/room/bed 总数 2) bed 按 bed_flag/occupy 统计空闲与占用 3) 组装 AvailableBedVO, 赋值 out
func (h *home) AvailableBed(ctx context.Context, in *dto.EmptyReq, out *dto.EmptyResp) error {
	// todo: 见上方注释
	return nil
}

// TodaySaleFollow 今日销售跟进（待跟进的咨询/预订）
// 对应 Java: HomeServiceImpl.todaySaleFollow
// 返回: 今日销售跟进列表(consult 今日新增)、预订跟进(reserve 今日新增)、合计
//
// todo: 查 consult/reserve 今日新增列表, 组装 TodaySaleFollowVO, 赋值 out
func (h *home) TodaySaleFollow(ctx context.Context, in *dto.EmptyReq, out *dto.EmptyResp) error {
	// todo: 见上方注释
	return nil
}

// MonthPerformanceRank 本月业绩排名
// 对应 Java: HomeServiceImpl.monthPerformanceRank
// 返回: 员工销售业绩排行(按 contract 本月新增金额聚合 staff_id)
//
// todo: SELECT staff_id, SUM(contract_amount) FROM contract WHERE create_time 本月 GROUP BY staff_id ORDER BY 金额 DESC; 组装 MonthPerformanceRankVO(含 SaleRank 列表), 赋值 out
func (h *home) MonthPerformanceRank(ctx context.Context, in *dto.EmptyReq, out *dto.EmptyResp) error {
	// todo: 见上方注释
	return nil
}

// ClientSource 客户来源分析（饼图）
// 对应 Java: HomeServiceImpl.clientSource -> sourceMapper.listSourceByRank
// 返回: 各来源来源数(rank)及来源名称, 用于饼图
//
// todo: 查 source 表及来源统计, 组装 ClientSourceVO(含 SourceRank 列表), 赋值 out
func (h *home) ClientSource(ctx context.Context, in *dto.EmptyReq, out *dto.EmptyResp) error {
	// todo: 见上方注释
	return nil
}

// BusinessTrend 业务趋势（折线图）
// 对应 Java: HomeServiceImpl.businessTrend -> consultMapper.listDateByMonth
// 返回: 近6月每月的咨询数/预订数/合同数趋势, 用于折线图
//
// todo: 按月份统计 consult/reserve/contract 数量, 组装 BusinessTrendVO(含 Trend 列表), 赋值 out
func (h *home) BusinessTrend(ctx context.Context, in *dto.EmptyReq, out *dto.EmptyResp) error {
	// todo: 见上方注释
	return nil
}

// 确保 do 包被引用(避免未使用), 实际统计查询会用到 do.X 接收结果
var _ = do.Consult{}
