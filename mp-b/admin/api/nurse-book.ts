/**
 * 护理预定相关接口
 * 对应后端 NurseReserveController（/nurseReserve/...）
 * 移动端：列表 + 新增/编辑护理预定
 */
import { get, post } from '../utils/request'
import { PageResult, mockPage } from './elder'

export interface NurseBookItem {
	id: number
	elder_name: string
	bed_name: string
	service_name: string
	need_date: number
	service_price: number
	charge_method: string
	frequency: number
	pay_amount: number
	nurse_date: string
	order_flag: string
}

export interface NurseBookForm {
	id?: number
	elder_id: number
	service_name: string
	service_price: number
	charge_method: string
	need_date: number
	frequency: number
	pay_amount: number
}

export interface NurseStaff {
	id: number
	name: string
	phone: string
}

// ====== MOCK 数据（后台未完成，测试前端用）======
const mockNurseBookList: NurseBookItem[] = [
	{ id: 1, elder_name: '张建国', bed_name: 'A栋-101-101室-A床', service_name: '翻身拍背', need_date: 15, service_price: 60, charge_method: '按次', frequency: 10, pay_amount: 600, nurse_date: '', order_flag: '待执行' },
	{ id: 2, elder_name: '王秀兰', bed_name: 'A栋-101-102室-A床', service_name: '血压监测', need_date: 30, service_price: 20, charge_method: '按次', frequency: 30, pay_amount: 600, nurse_date: '2026-08-21 09:00', order_flag: '已执行' }
]

const mockNurseStaffList: NurseStaff[] = [
	{ id: 1, name: '赵护士', phone: '13700000001' },
	{ id: 2, name: '钱护士', phone: '13700000002' }
]

/** 护理预定列表（全量） */
export function pageNurseBookByKey(keyword?: string): Promise<PageResult<NurseBookItem>> {
	// return get<PageResult<NurseBookItem>>('/nurseReserve/pageNurseReserveByKey', {
	// 	page_num: 1, page_size: 200, elder_name: keyword || undefined
	// })
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	let list = mockNurseBookList
	if (keyword) list = list.filter((it) => it.elder_name.indexOf(keyword) >= 0 || it.service_name.indexOf(keyword) >= 0)
	return Promise.resolve(mockPage(list))
}

/** 新增护理预定 */
export function addNurseBook(form: NurseBookForm): Promise<void> {
	// return post<void>('/nurseReserve/addNurseReserve', form)
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	return Promise.resolve()
}

/** 编辑护理预定 */
export function editNurseBook(form: NurseBookForm): Promise<void> {
	// return post<void>('/nurseReserve/editNurseReserve', form)
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	return Promise.resolve()
}

/** 删除护理预定 */
export function deleteNurseBook(id: number): Promise<void> {
	// return post<void>('/nurseReserve/deleteNurseReserve', { id })
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	return Promise.resolve()
}

/** 执行护理预定（选择护理人员） */
export function executeNurseBook(id: number, staffId: number, nurseDate: string): Promise<void> {
	// return post<void>('/nurseReserve/executeNurseReserve', { id, staff_id: staffId, nurse_date: nurseDate })
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	return Promise.resolve()
}

/** 护理人员列表（执行护理时选择） */
export function listNurseStaff(): Promise<NurseStaff[]> {
	// return get<NurseStaff[]>('/nurseReserve/listNurseStaff')
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	return Promise.resolve(mockNurseStaffList)
}
