/**
 * 长者档案相关接口
 * 对应后端 ElderRecordController（/elderRecord/...）
 * 移动端约定：列表全量拉取（page_num=1&page_size=200），本地过滤
 */
import { get, post } from '../utils/request'

const PAGE_NUM = 1
const PAGE_SIZE = 200

/** 分页响应壳 */
export interface PageResult<T> {
	page_num: number
	page_size: number
	pages: number
	size: number
	total: number
	list: T[]
}

export interface ElderItem {
	id: number
	bed_name: string
	name: string
	id_num: string
	age: number
	sex: string
	phone: string
	address: string
	check_flag: string
}

export interface EmergencyContact {
	id: number
	elder_id: number
	name: string
	phone: string
	relation: string
}

export interface ElderRecordForm {
	id?: number
	name: string
	id_num: string
	sex: string
	age: number
	phone: string
	address: string
	nurse_level?: string
	check_flag?: string
}

/** 长者档案列表（全量） */
export function pageElderByKey(keyword?: string, idNum?: string, sex?: string): Promise<PageResult<ElderItem>> {
	return get<PageResult<ElderItem>>('/elderRecord/pageElderRecordByKey', {
		page_num: PAGE_NUM,
		page_size: PAGE_SIZE,
		elder_name: keyword || undefined,
		id_num: idNum || undefined,
		elder_sex: sex || undefined
	})
}

/** 长者搜索器（外出/来访/事故/点餐复用） */
export function pageSearchElderByKey(name?: string, phone?: string): Promise<PageResult<ElderItem>> {
	return get<PageResult<ElderItem>>('/elderRecord/pageSearchElderByKey', {
		page_num: PAGE_NUM,
		page_size: PAGE_SIZE,
		name: name || undefined,
		phone: phone || undefined
	})
}

export function getElderRecordById(id: number): Promise<any> {
	return get<any>('/elderRecord/getElderRecordById', { id })
}

export function addElderRecord(form: ElderRecordForm): Promise<void> {
	return post<void>('/elderRecord/addElderRecord', form)
}

export function editElderRecord(form: ElderRecordForm): Promise<void> {
	return post<void>('/elderRecord/editElderRecord', form)
}

export function deleteElderRecord(id: number): Promise<void> {
	return post<void>('/elderRecord/deleteElderRecord', { id })
}

/** 某老人紧急联系人列表 */
export function pageSearchEmergencyContactByKey(elderId: number, key?: string): Promise<PageResult<EmergencyContact>> {
	return get<PageResult<EmergencyContact>>('/elderRecord/pageSearchEmergencyContactByKey', {
		page_num: PAGE_NUM,
		page_size: PAGE_SIZE,
		elder_id: elderId,
		key: key || undefined
	})
}

export function addEmergencyContact(p: { elder_id: number; name: string; phone: string; relation?: string }): Promise<void> {
	return post<void>('/elderRecord/addEmergencyContact', p)
}

export function editEmergencyContact(p: { id: number; name?: string; phone?: string; relation?: string }): Promise<void> {
	return post<void>('/elderRecord/editEmergencyContact', p)
}

export function deleteEmergencyContact(id: number): Promise<void> {
	return post<void>('/elderRecord/deleteEmergencyContact', { id })
}
