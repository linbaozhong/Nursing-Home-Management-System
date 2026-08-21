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

// ====== MOCK 数据（后台未完成，测试前端用；后台完成后放开各函数内注释并删除本段）======
export function mockPage<T>(list: T[]): PageResult<T> {
	return { page_num: PAGE_NUM, page_size: PAGE_SIZE, pages: 1, size: list.length, total: list.length, list }
}
const mockElderList: ElderItem[] = [
	{ id: 1, bed_name: 'A-101', name: '张建国', id_num: '110101194801011234', age: 76, sex: '男', phone: '13800000001', address: '北京市朝阳区xx路1号', check_flag: '1' },
	{ id: 2, bed_name: 'A-102', name: '王秀兰', id_num: '110101195203015678', age: 70, sex: '女', phone: '13800000002', address: '北京市海淀区xx路2号', check_flag: '1' },
	{ id: 3, bed_name: 'B-201', name: '李德福', id_num: '110101194507129012', age: 81, sex: '男', phone: '13800000003', address: '北京市丰台区xx路3号', check_flag: '0' }
]
const mockEmergencyList: EmergencyContact[] = [
	{ id: 1, elder_id: 1, name: '张小明', phone: '13900000001', relation: '儿子' },
	{ id: 2, elder_id: 1, name: '张小红', phone: '13900000002', relation: '女儿' },
	{ id: 3, elder_id: 2, name: '王大力', phone: '13900000003', relation: '儿子' }
]

/** 长者档案列表（全量） */
export function pageElderByKey(keyword?: string, idNum?: string, sex?: string): Promise<PageResult<ElderItem>> {
	// return get<PageResult<ElderItem>>('/elderRecord/pageElderRecordByKey', {
	// 	page_num: PAGE_NUM,
	// 	page_size: PAGE_SIZE,
	// 	elder_name: keyword || undefined,
	// 	id_num: idNum || undefined,
	// 	elder_sex: sex || undefined
	// })
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	let list = mockElderList
	if (keyword) list = list.filter((e) => e.name.indexOf(keyword) >= 0)
	if (idNum) list = list.filter((e) => e.id_num.indexOf(idNum) >= 0)
	if (sex) list = list.filter((e) => e.sex === sex)
	return Promise.resolve(mockPage(list))
}

/** 长者搜索器（外出/来访/事故/点餐复用） */
export function pageSearchElderByKey(name?: string, phone?: string): Promise<PageResult<ElderItem>> {
	// return get<PageResult<ElderItem>>('/elderRecord/pageSearchElderByKey', {
	// 	page_num: PAGE_NUM,
	// 	page_size: PAGE_SIZE,
	// 	name: name || undefined,
	// 	phone: phone || undefined
	// })
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	let list = mockElderList
	if (name) list = list.filter((e) => e.name.indexOf(name) >= 0)
	if (phone) list = list.filter((e) => e.phone.indexOf(phone) >= 0)
	return Promise.resolve(mockPage(list))
}

export function getElderRecordById(id: number): Promise<any> {
	// return get<any>('/elderRecord/getElderRecordById', { id })
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	const item = mockElderList.find((e) => e.id === id) || mockElderList[0]
	return Promise.resolve({ ...item, nurse_level: '一级护理', check_flag: item.check_flag })
}

export function addElderRecord(form: ElderRecordForm): Promise<void> {
	// return post<void>('/elderRecord/addElderRecord', form)
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve()
}

export function editElderRecord(form: ElderRecordForm): Promise<void> {
	// return post<void>('/elderRecord/editElderRecord', form)
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve()
}

export function deleteElderRecord(id: number): Promise<void> {
	// return post<void>('/elderRecord/deleteElderRecord', { id })
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve()
}

/** 某老人紧急联系人列表 */
export function pageSearchEmergencyContactByKey(elderId: number, key?: string): Promise<PageResult<EmergencyContact>> {
	// return get<PageResult<EmergencyContact>>('/elderRecord/pageSearchEmergencyContactByKey', {
	// 	page_num: PAGE_NUM,
	// 	page_size: PAGE_SIZE,
	// 	elder_id: elderId,
	// 	key: key || undefined
	// })
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	let list = mockEmergencyList.filter((c) => c.elder_id === elderId)
	if (key) list = list.filter((c) => c.name.indexOf(key) >= 0)
	return Promise.resolve(mockPage(list))
}

export function addEmergencyContact(p: { elder_id: number; name: string; phone: string; relation?: string }): Promise<void> {
	// return post<void>('/elderRecord/addEmergencyContact', p)
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve()
}

export function editEmergencyContact(p: { id: number; name?: string; phone?: string; relation?: string }): Promise<void> {
	// return post<void>('/elderRecord/editEmergencyContact', p)
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve()
}

export function deleteEmergencyContact(id: number): Promise<void> {
	// return post<void>('/elderRecord/deleteEmergencyContact', { id })
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve()
}
