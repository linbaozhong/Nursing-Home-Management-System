/**
 * 外出登记相关接口
 * 对应后端 OutwardController（/outward/...）
 */
import { get, post } from '../utils/request'
import { PageResult, mockPage } from './elder'

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

// ====== MOCK 数据（后台未完成，测试前端用；后台完成后放开各函数内注释并删除本段）======
const mockOutwardList: OutwardItem[] = [
	{ id: 1, elder_name: '张建国', chaperone_name: '张小明', chaperone_phone: '13900000001', chaperone_type: '家属', outward_date: '2026-08-20 09:00', plan_return_date: '2026-08-20 18:00', real_return_date: '' },
	{ id: 2, elder_name: '王秀兰', chaperone_name: '王大力', chaperone_phone: '13900000003', chaperone_type: '家属', outward_date: '2026-08-21 10:00', plan_return_date: '2026-08-22 17:00', real_return_date: '2026-08-21 16:00' }
]

/** 外出列表（全量） */
export function pageOutwardByKey(elderName?: string, chaperoneType?: string): Promise<PageResult<OutwardItem>> {
	// return get<PageResult<OutwardItem>>('/outward/pageOutwardByKey', {
	// 	page_num: PAGE_NUM,
	// 	page_size: PAGE_SIZE,
	// 	elder_name: elderName || undefined,
	// 	chaperone_type: chaperoneType || undefined
	// })
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	let list = mockOutwardList
	if (elderName) list = list.filter((o) => o.elder_name.indexOf(elderName) >= 0)
	if (chaperoneType) list = list.filter((o) => o.chaperone_type === chaperoneType)
	return Promise.resolve(mockPage(list))
}

export function getOutwardById(id: number): Promise<OutwardItem> {
	// return get<OutwardItem>('/outward/getOutwardById', { id })
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve(mockOutwardList.find((o) => o.id === id) || mockOutwardList[0])
}

export function addOutward(form: OutwardForm): Promise<void> {
	// return post<void>('/outward/addOutward', form)
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve()
}

export function editOutward(form: OutwardForm): Promise<void> {
	// return post<void>('/outward/editOutward', form)
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve()
}

export function delayReturn(id: number, planReturnDate: string): Promise<void> {
	// return post<void>('/outward/delayReturn', { id, plan_return_date: planReturnDate })
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve()
}

export function recordReturn(id: number, realReturnDate: string): Promise<void> {
	// return post<void>('/outward/recordReturn', { id, real_return_date: realReturnDate })
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve()
}

export function deleteOutward(id: number): Promise<void> {
	// return post<void>('/outward/deleteOutward', { id })
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve()
}
