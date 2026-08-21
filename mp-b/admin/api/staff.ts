/**
 * 员工相关接口（护工/送餐人员搜索）
 * 对应后端 ReserveController（/reserve/pageSearchStaffByKey）与 NurseReserveController
 */
import { get } from '../utils/request'
import { PageResult, mockPage } from './elder'

export interface StaffItem {
	id: number
	name: string
	phone: string
}

// ====== MOCK 数据（后台未完成，测试前端用；后台完成后放开函数内注释并删除本段）======
const mockStaffList: StaffItem[] = [
	{ id: 1, name: '赵护士', phone: '13700000001' },
	{ id: 2, name: '钱护士', phone: '13700000002' },
	{ id: 3, name: '孙护工', phone: '13700000003' }
]

/** 员工搜索（page_size=200 全量本地过滤） */
export function pageSearchStaffByKey(name?: string, phone?: string): Promise<PageResult<StaffItem>> {
	// return get<PageResult<StaffItem>>('/reserve/pageSearchStaffByKey', {
	// 	page_num: 1,
	// 	page_size: 200,
	// 	name: name || undefined,
	// 	phone: phone || undefined
	// })
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	let list = mockStaffList
	if (name) list = list.filter((s) => s.name.indexOf(name) >= 0)
	if (phone) list = list.filter((s) => s.phone.indexOf(phone) >= 0)
	return Promise.resolve(mockPage(list))
}
