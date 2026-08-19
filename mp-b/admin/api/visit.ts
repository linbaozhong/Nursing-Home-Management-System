/**
 * 来访登记相关接口
 * 对应后端 VisitController（/visit/...）
 */
import { get, post } from '../utils/request'
import { PageResult } from './elder'

const PAGE_NUM = 1
const PAGE_SIZE = 200

export interface VisitItem {
	id: number
	elder_name: string
	visit_name: string
	visit_phone: string
	relation: string
	visit_date: string
	leave_date: string
	visit_num: number
	visit_flag: number
}

export interface VisitForm {
	id?: number
	elder_id: number
	name: string
	phone: string
	relation: string
	visit_date: string
	visit_num: number
}

/** 来访列表（全量） */
export function pageVisitByKey(elderName?: string, visitFlag?: number): Promise<PageResult<VisitItem>> {
	return get<PageResult<VisitItem>>('/visit/pageVisitByKey', {
		page_num: PAGE_NUM,
		page_size: PAGE_SIZE,
		elder_name: elderName || undefined,
		visit_flag: visitFlag != null ? '' + visitFlag : undefined
	})
}

export function getVisitById(id: number): Promise<VisitItem> {
	return get<VisitItem>('/visit/getVisitById', { id })
}

export function addVisit(form: VisitForm): Promise<void> {
	return post<void>('/visit/addVisit', form)
}

export function editVisit(form: VisitForm): Promise<void> {
	return post<void>('/visit/editVisit', form)
}

export function recordLeave(id: number, leaveDate: string): Promise<void> {
	return post<void>('/visit/recordLeave', { id, leave_date: leaveDate })
}

export function deleteVisit(id: number): Promise<void> {
	return post<void>('/visit/deleteVisit', { id })
}
