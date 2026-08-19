/**
 * 员工相关接口（护工/送餐人员搜索）
 * 对应后端 ReserveController（/reserve/pageSearchStaffByKey）与 NurseReserveController
 */
import { get } from '../utils/request'
import { PageResult } from './elder'

export interface StaffItem {
	id: number
	name: string
	phone: string
}

/** 员工搜索（page_size=200 全量本地过滤） */
export function pageSearchStaffByKey(name?: string, phone?: string): Promise<PageResult<StaffItem>> {
	return get<PageResult<StaffItem>>('/reserve/pageSearchStaffByKey', {
		page_num: 1,
		page_size: 200,
		name: name || undefined,
		phone: phone || undefined
	})
}
