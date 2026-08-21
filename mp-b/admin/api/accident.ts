/**
 * 事故登记相关接口
 * 对应后端 AccidentController（/accident/...）
 */
import { get, post } from '../utils/request'
import { PageResult, mockPage } from './elder'

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

// ====== MOCK 数据（后台未完成，测试前端用；后台完成后放开各函数内注释并删除本段）======
const mockAccidentList: AccidentItem[] = [
	{ id: 1, elder_name: '李德福', staff_name: '孙护工', occur_date: '2026-08-15 14:00' }
]
const mockAccidentDetail: AccidentDetail = { id: 1, elder_name: '李德福', staff_id: 3, occur_date: '2026-08-15 14:00', description: '散步时轻微摔倒，已及时处理并包扎', picture: '' }

/** 事故列表（全量） */
export function pageAccidentByKey(elderName?: string, staffName?: string): Promise<PageResult<AccidentItem>> {
	// return get<PageResult<AccidentItem>>('/accident/pageAccidentByKey', {
	// 	page_num: PAGE_NUM,
	// 	page_size: PAGE_SIZE,
	// 	elder_name: elderName || undefined,
	// 	staff_name: staffName || undefined
	// })
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	let list = mockAccidentList
	if (elderName) list = list.filter((a) => a.elder_name.indexOf(elderName) >= 0)
	if (staffName) list = list.filter((a) => a.staff_name.indexOf(staffName) >= 0)
	return Promise.resolve(mockPage(list))
}

export function getAccidentById(id: number): Promise<AccidentDetail> {
	// return get<AccidentDetail>('/accident/getAccidentById', { id })
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve({ ...mockAccidentDetail, id })
}

export function addAccident(form: AccidentForm): Promise<void> {
	// return post<void>('/accident/addAccident', form)
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve()
}

export function editAccident(form: AccidentForm): Promise<void> {
	// return post<void>('/accident/editAccident', form)
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve()
}

export function deleteAccident(id: number): Promise<void> {
	// return post<void>('/accident/deleteAccident', { id })
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve()
}
