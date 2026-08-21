/**
 * 点餐相关接口
 * 对应后端 OrderController（/order/...）与 Dishes（/dishes/...）、NurseReserve（/nurseReserve/...）
 */
import { get, post } from '../utils/request'
import { PageResult, mockPage } from './elder'

const PAGE_NUM = 1
const PAGE_SIZE = 200

export interface OrderItem {
	id: number
	elder_name: string
	elder_phone: string
	dine_date: string
	dine_type: string
	staff_name: string
	deliver_dishes_date: string
	pay_amount: number
	order_flag: string
}

export interface DishesItem {
	id: number
	type_name: string
	dishes_name: string
	price: number
}

export interface NurseStaff {
	id: number
	name: string
	phone: string
}

// ====== MOCK 数据（后台未完成，测试前端用；后台完成后放开各函数内注释并删除本段）======
const mockOrderList: OrderItem[] = [
	{ id: 1, elder_name: '张建国', elder_phone: '13800000001', dine_date: '2026-08-21', dine_type: '午餐', staff_name: '赵护士', deliver_dishes_date: '', pay_amount: 25, order_flag: '0' },
	{ id: 2, elder_name: '王秀兰', elder_phone: '13800000002', dine_date: '2026-08-21', dine_type: '晚餐', staff_name: '钱护士', deliver_dishes_date: '2026-08-21 18:30', pay_amount: 30, order_flag: '1' }
]
const mockDishesList: DishesItem[] = [
	{ id: 1, type_name: '套餐A', dishes_name: '红烧肉套餐', price: 25 },
	{ id: 2, type_name: '套餐B', dishes_name: '清蒸鱼套餐', price: 28 },
	{ id: 3, type_name: '汤品', dishes_name: '紫菜蛋花汤', price: 6 }
]
const mockNurseStaffList: NurseStaff[] = [
	{ id: 1, name: '赵护士', phone: '13700000001' },
	{ id: 2, name: '钱护士', phone: '13700000002' }
]

/** 点餐列表（全量） */
export function pageOrderByKey(elderName?: string, elderPhone?: string): Promise<PageResult<OrderItem>> {
	// return get<PageResult<OrderItem>>('/order/pageOrderByKey', {
	// 	page_num: PAGE_NUM,
	// 	page_size: PAGE_SIZE,
	// 	elder_name: elderName || undefined,
	// 	elder_phone: elderPhone || undefined
	// })
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	let list = mockOrderList
	if (elderName) list = list.filter((o) => o.elder_name.indexOf(elderName) >= 0)
	if (elderPhone) list = list.filter((o) => o.elder_phone.indexOf(elderPhone) >= 0)
	return Promise.resolve(mockPage(list))
}

export function getOrderById(id: number): Promise<any> {
	// return get<any>('/order/getOrderById', { id })
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	const item = mockOrderList.find((o) => o.id === id) || mockOrderList[0]
	return Promise.resolve({ ...item, order_dishes_list: [{ dishes_name: '红烧肉套餐', order_num: 1 }] })
}

export function addOrder(p: {
	elder_id: number
	dine_date: string
	dine_type: string
	order_dishes_list: { dishes_id: number; order_num: number }[]
}): Promise<void> {
	// return post<void>('/order/addOrder', p)
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve()
}

export function sendOrder(id: number, staffId: number, deliverDishesDate: string): Promise<void> {
	// return post<void>('/order/sendOrder', { id, staff_id: staffId, deliver_dishes_date: deliverDishesDate })
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve()
}

/** 菜品列表（全量） */
export function pageDishesByKey(dishesName?: string): Promise<PageResult<DishesItem>> {
	// return get<PageResult<DishesItem>>('/dishes/pageDishesByKey', {
	// 	page_num: PAGE_NUM,
	// 	page_size: PAGE_SIZE,
	// 	dishes_name: dishesName || undefined
	// })
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	let list = mockDishesList
	if (dishesName) list = list.filter((d) => d.dishes_name.indexOf(dishesName) >= 0)
	return Promise.resolve(mockPage(list))
}

/** 送餐人员列表（无参全量） */
export function listNurseStaff(): Promise<NurseStaff[]> {
	// return get<NurseStaff[]>('/nurseReserve/listNurseStaff')
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve(mockNurseStaffList)
}
