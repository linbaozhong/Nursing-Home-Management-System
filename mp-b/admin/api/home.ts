/**
 * 首页驾驶舱相关接口
 * 对应后端 HomeController（/home/...）
 * 移动端：纵向卡片流展示各指标
 */
import { get } from '../utils/request'

export interface TodayOverview {
	today_add_consult_num: number // 今日新增咨询
	today_add_reserve_num: number // 今日新增预定
	today_add_contract_num: number // 今日新增合同
	today_contract_expire_num: number // 合同到期提醒
}

export interface AvailableBed {
	idle_room_num: number // 空闲房间数
	idle_bed_num: number // 空闲床位数
	exit_audit_num: number // 退住审核数
}

export interface TodaySaleFollow {
	today_return_visit_num: number // 今日待回访数
	today_returned_visit_num: number // 今日已回访数
	stay_returned_visit_num: number // 滞留回访数
}

export interface SaleRank {
	rank: number
	name: string
	consult_num: number
	contract_num: number
}

export interface MonthPerformanceRank {
	consult_client_num: number
	consult_client_float_rate: number
	sign_contract_num: number
	sign_contract_float_rate: number
	consult_conversion_rate: number
	consult_conversion_float_rate: number
	sale_rank_list: SaleRank[]
}

export interface ClientSource {
	source_name: string
	consult_num: number
}

export interface BusinessTrend {
	month: string
	consult_num: number
	contract_num: number
}

// ====== MOCK 数据（后台未完成，测试前端用）======
export function todayOverview(): Promise<TodayOverview> {
	// return get<TodayOverview>('/home/todayOverview')
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	return Promise.resolve({
		today_add_consult_num: 5,
		today_add_reserve_num: 3,
		today_add_contract_num: 2,
		today_contract_expire_num: 1
	})
}

export function availableBed(): Promise<AvailableBed> {
	// return get<AvailableBed>('/home/availableBed')
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	return Promise.resolve({ idle_room_num: 12, idle_bed_num: 8, exit_audit_num: 2 })
}

export function todaySaleFollow(): Promise<TodaySaleFollow> {
	// return get<TodaySaleFollow>('/home/todaySaleFollow')
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	return Promise.resolve({ today_return_visit_num: 6, today_returned_visit_num: 4, stay_returned_visit_num: 2 })
}

export function monthPerformanceRank(): Promise<MonthPerformanceRank> {
	// return get<MonthPerformanceRank>('/home/monthPerformanceRank')
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	return Promise.resolve({
		consult_client_num: 30,
		consult_client_float_rate: 0.15,
		sign_contract_num: 18,
		sign_contract_float_rate: 0.1,
		consult_conversion_rate: 0.6,
		consult_conversion_float_rate: 0.05,
		sale_rank_list: [
			{ rank: 1, name: '张三', consult_num: 12, contract_num: 8 },
			{ rank: 2, name: '李四', consult_num: 10, contract_num: 6 },
			{ rank: 3, name: '王五', consult_num: 8, contract_num: 4 }
		]
	})
}

export function clientSource(): Promise<ClientSource[]> {
	// return get<ClientSource[]>('/home/clientSource', { start_time, end_time })
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	return Promise.resolve([
		{ source_name: '社区活动', consult_num: 12 },
		{ source_name: '转介绍', consult_num: 9 },
		{ source_name: '网络推广', consult_num: 5 },
		{ source_name: '医院转诊', consult_num: 4 }
	])
}

export function businessTrend(): Promise<BusinessTrend[]> {
	// return get<BusinessTrend[]>('/home/businessTrend')
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	return Promise.resolve([
		{ month: '2026-07', consult_num: 24, contract_num: 14 },
		{ month: '2026-08', consult_num: 30, contract_num: 18 },
		{ month: '2026-09', consult_num: 27, contract_num: 16 }
	])
}
