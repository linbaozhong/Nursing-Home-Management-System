/**
 * 来访登记相关接口
 * 对应后端 VisitController（/visit/...）
 */
import { get, post } from '../utils/request'
import { PageResult, mockPage } from './elder'

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

// ====== MOCK 数据（后台未完成，测试前端用；后台完成后放开各函数内注释并删除本段）======
const mockVisitList: VisitItem[] = [
	{ id: 1, elder_name: '张建国', visit_name: '张小明', visit_phone: '13900000001', relation: '儿子', visit_date: '2026-08-20 10:00', leave_date: '2026-08-20 11:00', visit_num: 1, visit_flag: 1 },
	{ id: 2, elder_name: '李德福', visit_name: '李芳', visit_phone: '13900000009', relation: '女儿', visit_date: '2026-08-21 15:00', leave_date: '', visit_num: 2, visit_flag: 0 }
]

/** 来访列表（全量） */
export function pageVisitByKey(elderName?: string, visitFlag?: number): Promise<PageResult<VisitItem>> {
	// return get<PageResult<VisitItem>>('/visit/pageVisitByKey', {
	// 	page_num: PAGE_NUM,
	// 	page_size: PAGE_SIZE,
	// 	elder_name: elderName || undefined,
	// 	visit_flag: visitFlag != null ? '' + visitFlag : undefined
	// })
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	let list = mockVisitList
	if (elderName) list = list.filter((v) => v.elder_name.indexOf(elderName) >= 0)
	if (visitFlag != null) list = list.filter((v) => v.visit_flag === visitFlag)
	return Promise.resolve(mockPage(list))
}

export function getVisitById(id: number): Promise<VisitItem> {
	// return get<VisitItem>('/visit/getVisitById', { id })
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve(mockVisitList.find((v) => v.id === id) || mockVisitList[0])
}

export function addVisit(form: VisitForm): Promise<void> {
	// return post<void>('/visit/addVisit', form)
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve()
}

export function editVisit(form: VisitForm): Promise<void> {
	// return post<void>('/visit/editVisit', form)
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve()
}

export function recordLeave(id: number, leaveDate: string): Promise<void> {
	// return post<void>('/visit/recordLeave', { id, leave_date: leaveDate })
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve()
}

export function deleteVisit(id: number): Promise<void> {
	// return post<void>('/visit/deleteVisit', { id })
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve()
}
