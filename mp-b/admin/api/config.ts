/**
 * 基础配置（企业侧，非平台端）统一接口 + mock
 * 对应后端各配置 Controller：/source、/roomType、/serviceProject、/build、/dishes、/cateringSet、/nurseGrade、/label
 * 移动端入口：我的页右上角「配置」图标 → 配置菜单 → 各配置项列表
 * ⚠️ 与「日程业务」分离：配置是基础字典维护，进入配置页即管理字典，不触达日常登记业务。
 */
import { get, post } from '../utils/request'
import { PageResult, mockPage } from './elder'

// ---------- 通用配置项 ----------
export interface ConfigItem {
	id: number
	name: string
	[extra: string]: any
}

// ---------- 来源渠道 ----------
export interface SourceItem { id: number; name: string }
export interface RoomTypeItem { id: number; name: string; month_price: number }
export interface ServiceItem { id: number; type_name: string; service_name: string; charge_method: string; price: number; need_date: number }
export interface DishesItem { id: number; type_name: string; dishes_name: string; price: number }
export interface NurseGradeItem { id: number; name: string; type: string; month_price: number }
export interface BuildingItem { id: number; name: string; floor_num: number }

// ---------- 来源（配置项 key: source） ----------
export function pageSourceByKey(keyword?: string): Promise<PageResult<SourceItem>> {
	// return get<PageResult<SourceItem>>('/source/pageSourceByKey', { page_num: 1, page_size: 200, source_name: keyword || undefined })
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	let list: SourceItem[] = [
		{ id: 1, name: '社区活动' }, { id: 2, name: '转介绍' }, { id: 3, name: '网络推广' }, { id: 4, name: '医院转诊' }
	]
	if (keyword) list = list.filter((it) => it.name.indexOf(keyword) >= 0)
	return Promise.resolve(mockPage(list))
}
export function addSource(name: string): Promise<void> {
	// return post<void>('/source/addSource', { value: name })
	return Promise.resolve()
}
export function editSource(id: number, name: string): Promise<void> {
	// return post<void>('/source/editSource', { id, name })
	return Promise.resolve()
}
export function deleteSource(id: number): Promise<void> {
	// return post<void>('/source/deleteSource', { id })
	return Promise.resolve()
}

// ---------- 房型（配置项 key: roomtype） ----------
export function pageRoomTypeByKey(keyword?: string): Promise<PageResult<RoomTypeItem>> {
	// return get<PageResult<RoomTypeItem>>('/roomType/pageRoomTypeByKey', { page_num: 1, page_size: 200, room_type_name: keyword || undefined })
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	let list: RoomTypeItem[] = [
		{ id: 1, name: '标准单人间', month_price: 3000 },
		{ id: 2, name: '标准双人间', month_price: 2500 },
		{ id: 3, name: '特护单人间', month_price: 4000 }
	]
	if (keyword) list = list.filter((it) => it.name.indexOf(keyword) >= 0)
	return Promise.resolve(mockPage(list))
}
export function addRoomType(name: string, price: number): Promise<void> { return Promise.resolve() }
export function editRoomType(id: number, name: string, price: number): Promise<void> { return Promise.resolve() }
export function deleteRoomType(id: number): Promise<void> { return Promise.resolve() }

// ---------- 服务（配置项 key: service） ----------
export function pageServiceByKey(keyword?: string): Promise<PageResult<ServiceItem>> {
	// return get<PageResult<ServiceItem>>('/serviceProject/pageServiceByKey', { page_num: 1, page_size: 200, service_name: keyword || undefined })
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	let list: ServiceItem[] = [
		{ id: 1, type_name: '基础护理', service_name: '翻身拍背', charge_method: '按次', price: 60, need_date: 15 },
		{ id: 2, type_name: '健康监测', service_name: '血压监测', charge_method: '按次', price: 20, need_date: 5 },
		{ id: 3, type_name: '生活照料', service_name: '洗澡协助', charge_method: '按次', price: 80, need_date: 30 }
	]
	if (keyword) list = list.filter((it) => it.service_name.indexOf(keyword) >= 0 || it.type_name.indexOf(keyword) >= 0)
	return Promise.resolve(mockPage(list))
}
export function addService(typeName: string, serviceName: string, price: number, chargeMethod: string, needDate: number): Promise<void> { return Promise.resolve() }
export function editService(id: number, typeName: string, serviceName: string, price: number, chargeMethod: string, needDate: number): Promise<void> { return Promise.resolve() }
export function deleteService(id: number): Promise<void> { return Promise.resolve() }

// ---------- 菜品（配置项 key: dishes） ----------
export function pageDishesByKey2(keyword?: string): Promise<PageResult<DishesItem>> {
	// return get<PageResult<DishesItem>>('/dishes/pageDishesByKey', { page_num: 1, page_size: 200, dishes_name: keyword || undefined })
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	let list: DishesItem[] = [
		{ id: 1, type_name: '套餐A', dishes_name: '红烧肉套餐', price: 25 },
		{ id: 2, type_name: '套餐B', dishes_name: '清蒸鱼套餐', price: 28 },
		{ id: 3, type_name: '汤品', dishes_name: '紫菜蛋花汤', price: 6 }
	]
	if (keyword) list = list.filter((it) => it.dishes_name.indexOf(keyword) >= 0)
	return Promise.resolve(mockPage(list))
}
export function addDishes(typeName: string, name: string, price: number): Promise<void> { return Promise.resolve() }
export function editDishes(id: number, typeName: string, name: string, price: number): Promise<void> { return Promise.resolve() }
export function deleteDishes(id: number): Promise<void> { return Promise.resolve() }

// ---------- 护理等级（配置项 key: nursegrade） ----------
export function pageNurseGradeByKey(keyword?: string): Promise<PageResult<NurseGradeItem>> {
	// return get<PageResult<NurseGradeItem>>('/nurseGrade/pageNurseGradeByKey', { page_num: 1, page_size: 200, nurse_grade_name: keyword || undefined })
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	let list: NurseGradeItem[] = [
		{ id: 1, name: '三级护理', type: '一般', month_price: 1500 },
		{ id: 2, name: '二级护理', type: '半失能', month_price: 2500 },
		{ id: 3, name: '一级护理', type: '全失能', month_price: 3500 }
	]
	if (keyword) list = list.filter((it) => it.name.indexOf(keyword) >= 0)
	return Promise.resolve(mockPage(list))
}
export function addNurseGrade(name: string, type: string, price: number): Promise<void> { return Promise.resolve() }
export function editNurseGrade(id: number, name: string, type: string, price: number): Promise<void> { return Promise.resolve() }
export function deleteNurseGrade(id: number): Promise<void> { return Promise.resolve() }

// ---------- 楼栋（配置项 key: building） ----------
export function pageBuildingByKey(keyword?: string): Promise<PageResult<BuildingItem>> {
	// return get<PageResult<BuildingItem>>('/build/pageBuildingByKey', { page_num: 1, page_size: 200, key: keyword || undefined })
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	let list: BuildingItem[] = mockBuildingList
	if (keyword) list = list.filter((it) => it.name.indexOf(keyword) >= 0)
	return Promise.resolve(mockPage(list))
}
export function addBuilding(name: string, floorNum: number): Promise<void> { return Promise.resolve() }
export function editBuilding(id: number, name: string, floorNum: number): Promise<void> { return Promise.resolve() }
export function deleteBuilding(id: number): Promise<void> { return Promise.resolve() }

// ---------- 楼层（配置项 key: floor，可增删；保存后自动统计回写楼栋 floor_num） ----------
export interface FloorCfgItem {
	id: number
	building_id: number
	name: string
}
// mock 内存状态：每个楼栋的楼层层级。初始按 building.floor_num 展开。
const mockFloorStore: Map<number, FloorCfgItem[]> = new Map<number, FloorCfgItem[]>()

function ensureFloorStore() {
	if (mockFloorStore.size > 0) return
	mockBuildingList.forEach((b) => {
		const floors: FloorCfgItem[] = []
		for (let i = 1; i <= b.floor_num; i++) {
			floors.push({ id: b.id * 100 + i, building_id: b.id, name: i + '层' })
		}
		mockFloorStore.set(b.id, floors)
	})
}

// 统计并回写楼栋的 floor_num
function recountBuildingFloorNum(buildingId: number) {
	const floors = mockFloorStore.get(buildingId) || []
	const maxN = floors.reduce((m, f) => {
		const n = Number(f.name.replace(/[^0-9]/g, ''))
		return n > m ? n : m
	}, 0)
	const b = mockBuildingList.find((it) => it.id == buildingId)
	if (b) b.floor_num = maxN
}

export function listFloorByBuilding(buildingId: number): Promise<FloorCfgItem[]> {
	// return get<FloorCfgItem[]>('/build/listFloorByBuildingId', { building_id: buildingId })
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	ensureFloorStore()
	return Promise.resolve(mockFloorStore.get(buildingId) || [])
}

export function addFloor(buildingId: number, name: string): Promise<FloorCfgItem[]> {
	// return post<FloorCfgItem[]>('/build/addFloor', { building_id: buildingId, name })
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	ensureFloorStore()
	const list = mockFloorStore.get(buildingId) || []
	const maxId = list.reduce((m, f) => (f.id > m ? f.id : m), buildingId * 100)
	list.push({ id: maxId + 1, building_id: buildingId, name })
	mockFloorStore.set(buildingId, list)
	recountBuildingFloorNum(buildingId)
	return Promise.resolve(list)
}

export function deleteFloor(buildingId: number, floorId: number): Promise<FloorCfgItem[]> {
	// return post<FloorCfgItem[]>('/build/deleteFloor', { building_id: buildingId, id: floorId })
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	ensureFloorStore()
	const list = mockFloorStore.get(buildingId) || []
	mockFloorStore.set(buildingId, list.filter((f) => f.id != floorId))
	recountBuildingFloorNum(buildingId)
	return Promise.resolve(mockFloorStore.get(buildingId) || [])
}

// ---------- 设施（配置项 key: facility） ----------
export const ROOM_FACILITY_OPTIONS = ['餐桌', '轮椅', '空调', '电视', '独立卫生间', '呼叫铃', '衣柜', '阳台', '无障碍扶手']

// ---------- 床型（配置项 key: bedtype，供房间床位选择） ----------
export interface BedTypeItem { id: number; name: string }
export function pageBedTypeByKey(keyword?: string): Promise<PageResult<BedTypeItem>> {
	// return get<PageResult<BedTypeItem>>('/bed/pageBedTypeByKey', { page_num: 1, page_size: 200, key: keyword || undefined })
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	let list: BedTypeItem[] = [
		{ id: 1, name: '标准床' },
		{ id: 2, name: '智能床' },
		{ id: 3, name: '电动护理床' }
	]
	if (keyword) list = list.filter((it) => it.name.indexOf(keyword) >= 0)
	return Promise.resolve(mockPage(list))
}
export function addBedType(name: string): Promise<void> { return Promise.resolve() }
export function editBedType(id: number, name: string): Promise<void> { return Promise.resolve() }
export function deleteBedType(id: number): Promise<void> { return Promise.resolve() }

// ---------- 房间（配置项 key: room） ----------
export interface BedCfgItem {
	id: number
	name: string
	bed_type: string
	bed_flag: string
}
export interface RoomCfgItem {
	id: number
	floor_id: number
	building_id: number
	name: string
	type_id: number
	type_name: string
	bed_list: BedCfgItem[]
	facilities: string
}
export function pageRoomByKey(keyword?: string): Promise<PageResult<RoomCfgItem>> {
	// return get<PageResult<RoomCfgItem>>('/build/pageRoomByKey', { page_num: 1, page_size: 200, key: keyword || undefined })
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	let list = mockRoomList.filter((it) => !keyword || it.name.indexOf(keyword) >= 0)
	return Promise.resolve(mockPage(list))
}
export function addRoom(form: RoomCfgItem): Promise<void> { return Promise.resolve() }
export function editRoom(id: number, form: RoomCfgItem): Promise<void> { return Promise.resolve() }
export function deleteRoom(id: number): Promise<void> { return Promise.resolve() }

// ====== 房间 mock 数据（在底部维持对楼栋 mock 的依赖）======
const mockBuildingList = [
	{ id: 1, name: 'A栋', floor_num: 3 },
	{ id: 2, name: 'B栋', floor_num: 2 }
]
const mockRoomList: RoomCfgItem[] = [
	{
		id: 101, floor_id: 102, building_id: 1, name: 'A201', type_id: 2, type_name: '标准双人间',
		bed_list: [
			{ id: 10101, name: '01', bed_type: '标准床', bed_flag: '空' },
			{ id: 10102, name: '02', bed_type: '智能床', bed_flag: '空' }
		],
		facilities: '餐桌,轮椅'
	},
	{
		id: 102, floor_id: 102, building_id: 1, name: 'A202', type_id: 1, type_name: '标准单人间',
		bed_list: [{ id: 10201, name: '01', bed_type: '标准床', bed_flag: '空' }],
		facilities: '空调,电视,独立卫生间'
	},
	{
		id: 201, floor_id: 101, building_id: 1, name: 'A101', type_id: 2, type_name: '标准双人间',
		bed_list: [
			{ id: 20101, name: '01', bed_type: '标准床', bed_flag: '已入住' },
			{ id: 20102, name: '02', bed_type: '标准床', bed_flag: '空' }
		],
		facilities: '餐桌'
	}
]
