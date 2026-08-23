/**
 * 床位全景相关接口
 * 对应后端 BedPanoramaController（/bedPanorama/...）
 * 移动端：层级下钻 楼栋 → 楼层 → 房间 → 床位
 */
import { get } from '../utils/request'

export interface DropDownItem {
	id: number
	name: string
}

export interface BedItem {
	id: number
	name: string
	bed_flag: string // 床位状态：空/已占/预定...
	elder_name: string // 所住老人（空则空闲）
}

export interface RoomItem {
	id: number
	name: string
	bed_list: BedItem[]
}

export interface FloorItem {
	id: number
	name: string
	room_list: RoomItem[]
}

// ====== MOCK 数据（后台未完成，测试前端用）======
const mockBuildingList: DropDownItem[] = [
	{ id: 1, name: 'A栋' },
	{ id: 2, name: 'B栋' }
]

const mockFloorList = new Map<number, DropDownItem[]>()
mockFloorList.set(1, [{ id: 101, name: '1层' }, { id: 102, name: '2层' }])
mockFloorList.set(2, [{ id: 201, name: '1层' }])

const mockRoomList = new Map<number, FloorItem[]>()
mockRoomList.set(101, [
	{
		id: 1011, name: '101室',
		bed_list: [
			{ id: 10111, name: 'A床', bed_flag: '已入住', elder_name: '张建国' },
			{ id: 10112, name: 'B床', bed_flag: '空', elder_name: '' }
		]
	},
	{
		id: 1012, name: '102室',
		bed_list: [
			{ id: 10121, name: 'A床', bed_flag: '已入住', elder_name: '王秀兰' }
		]
	}
])
mockRoomList.set(102, [
	{
		id: 1021, name: '201室',
		bed_list: [
			{ id: 10211, name: 'A床', bed_flag: '已入住', elder_name: '李德福' },
			{ id: 10212, name: 'B床', bed_flag: '空', elder_name: '' }
		]
	}
])
mockRoomList.set(201, [
	{
		id: 2011, name: '101室',
		bed_list: [
			{ id: 20111, name: 'A床', bed_flag: '空', elder_name: '' }
		]
	}
])

/** 楼栋列表 */
export function listBuilding(): Promise<DropDownItem[]> {
	// return get<DropDownItem[]>('/bedPanorama/listBuilding')
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	return Promise.resolve(mockBuildingList)
}

/** 按楼栋编号获取楼层列表 */
export function listFloorByBuildingId(buildingId: number): Promise<DropDownItem[]> {
	// return get<DropDownItem[]>('/bedPanorama/listFloorByBuildingId', { building_id: buildingId })
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	return Promise.resolve(mockFloorList.get(buildingId) || [])
}

/** 按楼层编号获取房间列表（含床位与入住老人） */
export function listRoomByFloorId(floorId: number): Promise<FloorItem[]> {
	// return get<FloorItem[]>('/bedPanorama/listRoomByKey', { floor_id: floorId })
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	return Promise.resolve(mockRoomList.get(floorId) || [])
}
