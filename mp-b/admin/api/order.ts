/**
 * 点餐相关接口
 * 对应后端 OrderController（/order/...）与 Dishes（/dishes/...）、NurseReserve（/nurseReserve/...）
 */
import { get, post } from '../utils/request'
import { PageResult } from './elder'

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

/** 点餐列表（全量） */
export function pageOrderByKey(elderName?: string, elderPhone?: string): Promise<PageResult<OrderItem>> {
	return get<PageResult<OrderItem>>('/order/pageOrderByKey', {
		page_num: PAGE_NUM,
		page_size: PAGE_SIZE,
		elder_name: elderName || undefined,
		elder_phone: elderPhone || undefined
	})
}

export function getOrderById(id: number): Promise<any> {
	return get<any>('/order/getOrderById', { id })
}

export function addOrder(p: {
	elder_id: number
	dine_date: string
	dine_type: string
	order_dishes_list: { dishes_id: number; order_num: number }[]
}): Promise<void> {
	return post<void>('/order/addOrder', p)
}

export function sendOrder(id: number, staffId: number, deliverDishesDate: string): Promise<void> {
	return post<void>('/order/sendOrder', { id, staff_id: staffId, deliver_dishes_date: deliverDishesDate })
}

/** 菜品列表（全量） */
export function pageDishesByKey(dishesName?: string): Promise<PageResult<DishesItem>> {
	return get<PageResult<DishesItem>>('/dishes/pageDishesByKey', {
		page_num: PAGE_NUM,
		page_size: PAGE_SIZE,
		dishes_name: dishesName || undefined
	})
}

/** 送餐人员列表（无参全量） */
export function listNurseStaff(): Promise<NurseStaff[]> {
	return get<NurseStaff[]>('/nurseReserve/listNurseStaff')
}
