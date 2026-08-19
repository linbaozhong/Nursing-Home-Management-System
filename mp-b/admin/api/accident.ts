/**
 * 事故登记相关接口
 * 对应后端 AccidentController（/accident/...）
 */
import { get, post } from '../utils/request'
import { PageResult } from './elder'

const PAGE_NUM = 1
const PAGE_SIZE = 200

export interface AccidentItem {
	id: number
	elder_name: string
	staff_name: string
	occur_date: string
}

export interface AccidentDetail {
	id: number
	elder_name: string
	staff_id: number
	occur_date: string
	description: string
	picture: string
}

export interface AccidentForm {
	id?: number
	elder_id: number
	staff_id: number
	occur_date: string
	description: string
	picture: string
}

/** 事故列表（全量） */
export function pageAccidentByKey(elderName?: string, staffName?: string): Promise<PageResult<AccidentItem>> {
	return get<PageResult<AccidentItem>>('/accident/pageAccidentByKey', {
		page_num: PAGE_NUM,
		page_size: PAGE_SIZE,
		elder_name: elderName || undefined,
		staff_name: staffName || undefined
	})
}

export function getAccidentById(id: number): Promise<AccidentDetail> {
	return get<AccidentDetail>('/accident/getAccidentById', { id })
}

export function addAccident(form: AccidentForm): Promise<void> {
	return post<void>('/accident/addAccident', form)
}

export function editAccident(form: AccidentForm): Promise<void> {
	return post<void>('/accident/editAccident', form)
}

export function deleteAccident(id: number): Promise<void> {
	return post<void>('/accident/deleteAccident', { id })
}
