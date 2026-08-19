/**
 * 外出登记相关接口
 * 对应后端 OutwardController（/outward/...）
 */
import { get, post } from '../utils/request'
import { PageResult } from './elder'

const PAGE_NUM = 1
const PAGE_SIZE = 200

export interface OutwardItem {
	id: number
	elder_name: string
	chaperone_name: string
	chaperone_phone: string
	chaperone_type: string
	outward_date: string
	plan_return_date: string
	real_return_date: string
}

export interface OutwardForm {
	id?: number
	elder_id: number
	chaperone_name: string
	chaperone_phone: string
	chaperone_type: string
	outward_date: string
	plan_return_date: string
}

/** 外出列表（全量） */
export function pageOutwardByKey(elderName?: string, chaperoneType?: string): Promise<PageResult<OutwardItem>> {
	return get<PageResult<OutwardItem>>('/outward/pageOutwardByKey', {
		page_num: PAGE_NUM,
		page_size: PAGE_SIZE,
		elder_name: elderName || undefined,
		chaperone_type: chaperoneType || undefined
	})
}

export function getOutwardById(id: number): Promise<OutwardItem> {
	return get<OutwardItem>('/outward/getOutwardById', { id })
}

export function addOutward(form: OutwardForm): Promise<void> {
	return post<void>('/outward/addOutward', form)
}

export function editOutward(form: OutwardForm): Promise<void> {
	return post<void>('/outward/editOutward', form)
}

export function delayReturn(id: number, planReturnDate: string): Promise<void> {
	return post<void>('/outward/delayReturn', { id, plan_return_date: planReturnDate })
}

export function recordReturn(id: number, realReturnDate: string): Promise<void> {
	return post<void>('/outward/recordReturn', { id, real_return_date: realReturnDate })
}

export function deleteOutward(id: number): Promise<void> {
	return post<void>('/outward/deleteOutward', { id })
}
