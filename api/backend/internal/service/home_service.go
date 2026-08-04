package service

import (
	"context"
	"math"
	"sort"
	"time"

	"api/internal/constant"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblbed"
	"api/internal/model/define/table/tblconsult"
	"api/internal/model/define/table/tblcontract"
	"api/internal/model/define/table/tblreserve"
	"api/internal/model/define/table/tblroom"
	"api/internal/model/define/table/tblsource"
	"api/internal/model/define/table/tblstaff"
	"api/internal/model/define/table/tblvisitplan"
	"api/internal/model/do"
	"api/internal/model/dto"

	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/types"
)

type home struct{}

var Home = &home{}

// saleRankRoleId 销售管理员角色编号
// 对应 Java: StaffFunc.listSaleRank -> listStaffByRoleId(2L)
const saleRankRoleId = 2

// dayRange 返回指定日期 [00:00:00, 23:59:59] 的时间区间
// 对应 Java: DateUtilWen.getDayStartTime / getDayEndTime
func dayRange(t time.Time) (start, end time.Time) {
	y, m, d := t.Date()
	loc := t.Location()
	start = time.Date(y, m, d, 0, 0, 0, 0, loc)
	end = time.Date(y, m, d, 23, 59, 59, 0, loc)
	return
}

// monthRange 返回指定日期所在月的 [月初 00:00:00, 月末 23:59:59]
// 对应 Java: DateUtilWen.getMonthFirstDay / getMonthLastDay
func monthRange(t time.Time) (start, end time.Time) {
	y, m, _ := t.Date()
	loc := t.Location()
	start = time.Date(y, m, 1, 0, 0, 0, 0, loc)
	end = start.AddDate(0, 1, 0).Add(-time.Second)
	return
}

// yearRange 返回指定日期所在年的 [年初 00:00:00, 年末 23:59:59]
// 对应 Java: DateUtilWen.getYearFirstDay / getYearLastDay
func yearRange(t time.Time) (start, end time.Time) {
	loc := t.Location()
	start = time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, loc)
	end = start.AddDate(1, 0, 0).Add(-time.Second)
	return
}

// ratio 计算比值，分母为 0 时返回 0，保留两位小数
// 对应 Java: CommonUtil.getRatio
func ratio(molecule, denominator float64) float64 {
	if denominator == 0 {
		return 0
	}
	return round2(molecule / denominator)
}

// growthRate 计算增长率，上期为 0 时返回 0，保留两位小数
// 对应 Java: CommonUtil.getGrowthRate
func growthRate(lastNum, thisNum float64) float64 {
	if lastNum == 0 {
		return 0
	}
	return round2((thisNum - lastNum) / lastNum)
}

// round2 保留两位小数
func round2(f float64) float64 {
	return math.Round(f*100) / 100
}

// TodayOverview 今日概览
// 对应 Java: HomeServiceImpl.todayOverview
// 返回: 今日新增咨询、今日新增预定、今日新增合同、今日到期合同
func (h *home) TodayOverview(ctx context.Context, in *dto.EmptyReq, out *dto.TodayOverviewVO) error {
	start, end := dayRange(time.Now())

	// 今日新增咨询
	consultNum, e := dao.Consult(db).Count(ctx,
		tblconsult.CreateTime.Gte(types.Time{start}),
		tblconsult.CreateTime.Lte(types.Time{end}),
	)
	if e != nil {
		return e
	}
	// 今日新增预定
	reserveNum, e := dao.Reserve(db).Count(ctx,
		tblreserve.CreateTime.Gte(types.Time{start}),
		tblreserve.CreateTime.Lte(types.Time{end}),
	)
	if e != nil {
		return e
	}
	// 今日新增合同
	contractNum, e := dao.Contract(db).Count(ctx,
		tblcontract.CreateTime.Gte(types.Time{start}),
		tblcontract.CreateTime.Lte(types.Time{end}),
	)
	if e != nil {
		return e
	}
	// 今日到期合同
	expireNum, e := dao.Contract(db).Count(ctx,
		tblcontract.EndDate.Gte(types.Time{start}),
		tblcontract.EndDate.Lte(types.Time{end}),
	)
	if e != nil {
		return e
	}

	out.TodayAddConsultNum = consultNum
	out.TodayAddReserveNum = reserveNum
	out.TodayAddContractNum = contractNum
	out.TodayContractExpireNum = expireNum
	return nil
}

// AvailableBed 可售床位
// 对应 Java: HomeServiceImpl.availableBed
// 返回: 空闲房间数（房间内有床位且全部空闲）、空闲床位数、已登记退床数
func (h *home) AvailableBed(ctx context.Context, in *dto.EmptyReq, out *dto.AvailableBedVO) error {
	// 获取所有未被删除床位
	bedList, _, e := dao.Bed(db).List(ctx, ace.Where(tblbed.DelFlag.Eq(constant.YesNoNo)))
	if e != nil {
		return e
	}

	// 按房间分组统计：床位总数 与 空闲床位数
	type bedStat struct {
		total int64
		idle  int64
	}
	statByRoom := make(map[types.BigInt]*bedStat, len(bedList))
	var idleBedNum, exitAuditNum int64
	for _, bed := range bedList {
		st, ok := statByRoom[bed.RoomId]
		if !ok {
			st = &bedStat{}
			statByRoom[bed.RoomId] = st
		}
		st.total++
		switch int8(bed.BedFlag) {
		case int8(constant.BedIdle):
			st.idle++
			idleBedNum++
		case int8(constant.BedExitAudit):
			exitAuditNum++
		}
	}

	// 获取所有未被删除房间, 统计空闲房间（房间内有床位且都是空闲状态）
	roomList, _, e := dao.Room(db).List(ctx, ace.Where(tblroom.DelFlag.Eq(constant.YesNoNo)))
	if e != nil {
		return e
	}
	var idleRoomNum int64
	for _, room := range roomList {
		if st, ok := statByRoom[room.Id]; ok && st.total > 0 && st.total == st.idle {
			idleRoomNum++
		}
	}

	out.IdleRoomNum = idleRoomNum
	out.IdleBedNum = idleBedNum
	out.ExitAuditNum = exitAuditNum
	return nil
}

// TodaySaleFollow 今日销售跟进
// 对应 Java: HomeServiceImpl.todaySaleFollow
// 返回: 今日应回访数、今日已回访数、待回访数
func (h *home) TodaySaleFollow(ctx context.Context, in *dto.EmptyReq, out *dto.TodaySaleFollowVO) error {
	start, end := dayRange(time.Now())

	// 获取所有未被删除回访计划
	planList, _, e := dao.VisitPlan(db).List(ctx, ace.Where(tblvisitplan.DelFlag.Eq(constant.YesNoNo)))
	if e != nil {
		return e
	}

	var returnNum, returnedNum, stayNum int64
	for _, plan := range planList {
		planDate := plan.PlanDate.Time
		completeDate := plan.CompleteDate.Time
		// 今日应回访: 计划回访时间落在今日区间内
		if planDate.After(start) && planDate.Before(end) {
			returnNum++
		}
		// 今日已回访: 完成时间落在今日区间内
		if !completeDate.IsZero() && completeDate.After(start) && completeDate.Before(end) {
			returnedNum++
		}
		// 待回访: 回访内容为空 或 未完成
		if string(plan.Content) == "" || completeDate.IsZero() {
			stayNum++
		}
	}

	out.TodayReturnVisitNum = returnNum
	out.TodayReturnedVisitNum = returnedNum
	out.StayReturnedVisitNum = stayNum
	return nil
}

// MonthPerformanceRank 本月业绩排行
// 对应 Java: HomeServiceImpl.monthPerformanceRank
// 返回: 本月咨询客户数/浮动率、签约合同数/浮动率、咨询转化率/浮动率、销售排行列表
func (h *home) MonthPerformanceRank(ctx context.Context, in *dto.EmptyReq, out *dto.MonthPerformanceRankVO) error {
	now := time.Now()
	lastMonthStart, lastMonthEnd := monthRange(now.AddDate(0, -1, 0))
	thisMonthStart, thisMonthEnd := monthRange(now)

	// 上月咨询客户记录
	lastMonthConsultList, _, e := dao.Consult(db).List(ctx, ace.Where(
		tblconsult.CreateTime.Gte(types.Time{lastMonthStart}),
		tblconsult.CreateTime.Lte(types.Time{lastMonthEnd}),
	))
	if e != nil {
		return e
	}
	// 本月咨询客户记录
	thisMonthConsultList, _, e := dao.Consult(db).List(ctx, ace.Where(
		tblconsult.CreateTime.Gte(types.Time{thisMonthStart}),
		tblconsult.CreateTime.Lte(types.Time{thisMonthEnd}),
	))
	if e != nil {
		return e
	}
	// 上月签订合同记录
	lastMonthContractList, _, e := dao.Contract(db).List(ctx, ace.Where(
		tblcontract.CreateTime.Gte(types.Time{lastMonthStart}),
		tblcontract.CreateTime.Lte(types.Time{lastMonthEnd}),
	))
	if e != nil {
		return e
	}
	// 本月签订合同记录
	thisMonthContractList, _, e := dao.Contract(db).List(ctx, ace.Where(
		tblcontract.CreateTime.Gte(types.Time{thisMonthStart}),
		tblcontract.CreateTime.Lte(types.Time{thisMonthEnd}),
	))
	if e != nil {
		return e
	}

	lastConsultNum := float64(len(lastMonthConsultList))
	thisConsultNum := float64(len(thisMonthConsultList))
	lastContractNum := float64(len(lastMonthContractList))
	thisContractNum := float64(len(thisMonthContractList))

	// 上月和本月合同与咨询比值（转换率）
	lastRatio := ratio(lastContractNum, lastConsultNum)
	thisRatio := ratio(thisContractNum, thisConsultNum)

	// 销售排行
	saleRankList, e := h.listSaleRank(ctx, thisMonthConsultList, thisMonthContractList)
	if e != nil {
		return e
	}

	out.ConsultClientNum = int64(len(thisMonthConsultList))
	out.ConsultClientFloatRate = growthRate(lastConsultNum, thisConsultNum)
	out.SignContractNum = int64(len(thisMonthContractList))
	out.SignContractFloatRate = growthRate(lastContractNum, thisContractNum)
	out.ConsultConversionRate = thisRatio
	out.ConsultConversionFloatRate = growthRate(lastRatio, thisRatio)
	out.SaleRankList = saleRankList
	return nil
}

// listSaleRank 统计销售管理员的咨询数/合同数并排行
// 对应 Java: StaffFunc.listSaleRank
func (h *home) listSaleRank(ctx context.Context, consultList []*do.Consult, contractList []*do.Contract) ([]dto.SaleRank, error) {
	// 根据员工编号对咨询分组计数
	consultNumByStaff := make(map[types.BigInt]int64, len(consultList))
	for _, consult := range consultList {
		consultNumByStaff[consult.StaffId]++
	}
	// 根据员工编号对合同分组计数
	contractNumByStaff := make(map[types.BigInt]int64, len(contractList))
	for _, contract := range contractList {
		contractNumByStaff[contract.StaffId]++
	}

	// 获取销售管理员列表
	staffList, _, e := dao.Staff(db).List(ctx, ace.Where(tblstaff.RoleId.Eq(saleRankRoleId)))
	if e != nil {
		return nil, e
	}

	saleRankList := make([]dto.SaleRank, 0, len(staffList))
	for _, staff := range staffList {
		saleRankList = append(saleRankList, dto.SaleRank{
			Name:        string(staff.Name),
			ConsultNum:  consultNumByStaff[staff.Id],
			ContractNum: contractNumByStaff[staff.Id],
		})
	}

	// 按合同数倒序排序后编排名次
	sort.SliceStable(saleRankList, func(i, j int) bool {
		return saleRankList[i].ContractNum > saleRankList[j].ContractNum
	})
	for i := range saleRankList {
		saleRankList[i].Rank = int64(i + 1)
	}
	return saleRankList, nil
}

// ClientSource 客户来源渠道统计
// 对应 Java: HomeServiceImpl.clientSource
// 返回: 各来源渠道的咨询人数
func (h *home) ClientSource(ctx context.Context, in *dto.ClientSourceQuery, out *[]dto.ClientSourceVO) error {
	// 获取开始/结束时间
	// 未传时默认取今天
	start, end := dayRange(time.Now())
	if t := parseDateStart(in.StartTime); !t.IsZero() {
		start, _ = dayRange(t)
	}
	if t := parseDateStart(in.EndTime); !t.IsZero() {
		_, end = dayRange(t)
	}

	// 获取来源渠道列表
	sourceList, _, e := dao.Source(db).List(ctx, ace.Where(tblsource.DelFlag.Eq(constant.YesNoNo)))
	if e != nil {
		return e
	}
	// 获取时间段内咨询人列表
	consultList, _, e := dao.Consult(db).List(ctx, ace.Where(
		tblconsult.CreateTime.Gte(types.Time{start}),
		tblconsult.CreateTime.Lte(types.Time{end}),
	))
	if e != nil {
		return e
	}

	// 根据来源渠道编号对咨询分组计数
	consultNumBySource := make(map[types.BigInt]int64, len(consultList))
	for _, consult := range consultList {
		consultNumBySource[consult.SourceId]++
	}

	list := make([]dto.ClientSourceVO, 0, len(sourceList))
	for _, source := range sourceList {
		list = append(list, dto.ClientSourceVO{
			SourceName: string(source.Name),
			ConsultNum: consultNumBySource[source.Id],
		})
	}
	*out = list
	return nil
}

// BusinessTrend 今年业务趋势
// 对应 Java: HomeServiceImpl.businessTrend
// 返回: 今年 12 个月每月的咨询数与合同数
func (h *home) BusinessTrend(ctx context.Context, in *dto.EmptyReq, out *[]dto.BusinessTrendVO) error {
	yearStart, yearEnd := yearRange(time.Now())

	// 获取今年所有咨询客户记录
	consultList, _, e := dao.Consult(db).List(ctx, ace.Where(
		tblconsult.CreateTime.Gte(types.Time{yearStart}),
		tblconsult.CreateTime.Lte(types.Time{yearEnd}),
	))
	if e != nil {
		return e
	}
	// 获取今年所有签订合同记录
	contractList, _, e := dao.Contract(db).List(ctx, ace.Where(
		tblcontract.CreateTime.Gte(types.Time{yearStart}),
		tblcontract.CreateTime.Lte(types.Time{yearEnd}),
	))
	if e != nil {
		return e
	}

	// 按月归集
	consultNumByMonth := make(map[string]int64, 12)
	for _, consult := range consultList {
		consultNumByMonth[consult.CreateTime.Format("2006-01")]++
	}
	contractNumByMonth := make(map[string]int64, 12)
	for _, contract := range contractList {
		contractNumByMonth[contract.CreateTime.Format("2006-01")]++
	}

	list := make([]dto.BusinessTrendVO, 0, 12)
	for i := 0; i < 12; i++ {
		month := yearStart.AddDate(0, i, 0).Format("2006-01")
		list = append(list, dto.BusinessTrendVO{
			Month:       month,
			ConsultNum:  consultNumByMonth[month],
			ContractNum: contractNumByMonth[month],
		})
	}
	*out = list
	return nil
}
