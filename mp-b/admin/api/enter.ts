/**
 * 入住签约相关接口
 * 对应后端 CheckContractController（/checkContract/...）与 BuildController（build.getBuildTree）
 * 移动端：多步向导（选择老人 → 护理等级/餐饮套餐/床位 → 合同信息 + 紧急联系人）
 */
import { get, post } from '../utils/request'
import { PageResult, mockPage } from './elder'

export interface DropDownItem {
	id: number
	name: string
}

export interface EmergencyContact {
	name: string
	phone: string
	relation: string
}

export interface EnterForm {
	nursing_grade_id: number // 护理等级编号
	catering_set_id: number // 餐饮套餐编号
	bed_id: number // 床位编号
	name: string // 姓名
	id_num: string // 身份证号
	age: number // 年龄
	sex: string // 性别
	phone: string // 电话
	address: string // 地址
	staff_id: number // 营销人员编号
	sign_date: string // 合同签订日期
	start_date: string // 合同开始日期
	end_date: string // 合同结束日期
	emergency_contacts: EmergencyContact[] // 紧急联系人
}

export interface EnterItem {
	id: number
	name: string
	sex: string
	age: number
	phone: string
	address: string
	check_flag: string
}

export interface BedOption {
	id: number
	name: string
	room_name: string
}

// ====== MOCK 数据（后台未完成，测试前端用）======
const mockEnterList: EnterItem[] = [
	{ id: 1, name: '赵德柱', sex: '男', age: 72, phone: '13800000011', address: '北京市朝阳区', check_flag: '入住中' }
]

const mockNurseGradeList: DropDownItem[] = [
	{ id: 1, name: '三级护理' },
	{ id: 2, name: '二级护理' },
	{ id: 3, name: '一级护理' }
]

const mockCateringSetList: DropDownItem[] = [
	{ id: 1, name: '标准A餐' },
	{ id: 2, name: '特需B餐' }
]

const mockBedList: BedOption[] = [
	{ id: 1, name: 'A栋-101-101室-A床', room_name: '101室' },
	{ id: 2, name: 'A栋-101-102室-A床', room_name: '102室' },
	{ id: 3, name: 'B栋-201-101室-A床', room_name: '101室' }
]

/** 入住签约列表（全量） */
export function pageEnterByKey(keyword?: string): Promise<PageResult<EnterItem>> {
	// return get<PageResult<EnterItem>>('/checkContract/pageCheckContractByKey', {
	// 	page_num: 1, page_size: 200, name: keyword || undefined
	// })
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	let list = mockEnterList
	if (keyword) list = list.filter((it) => it.name.indexOf(keyword) >= 0)
	return Promise.resolve(mockPage(list))
}

/** 护理等级下拉（签约向导） */
export function listNurseGrade(): Promise<DropDownItem[]> {
	// return get<DropDownItem[]>('/checkContract/listNurseGrade')
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	return Promise.resolve(mockNurseGradeList)
}

/** 餐饮套餐下拉（签约向导） */
export function listCateringSet(): Promise<DropDownItem[]> {
	// return get<DropDownItem[]>('/checkContract/listCateringSet')
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	return Promise.resolve(mockCateringSetList)
}

/** 可签约床位列表（签约向导） */
export function listBed(): Promise<BedOption[]> {
	// return get<BedOption[]>('/checkContract/getBuildTree') → 前端从楼栋树筛选空床
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	return Promise.resolve(mockBedList)
}

/** 新增入住签约 */
export function addEnter(form: EnterForm): Promise<void> {
	// return post<void>('/checkContract/addCheckContract', form)
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	return Promise.resolve()
}
