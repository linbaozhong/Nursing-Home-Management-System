/**
 * 退住申请相关接口
 * 对应后端 RetreatApplyController（/retreatApply/...）
 * 移动端：列表 + 新增/编辑退住申请
 */
import { get, post } from '../utils/request'
import { PageResult, mockPage } from './elder'

export interface ApplyItem {
	apply_id: number
	elder_id: number
	elder_name: string
	elder_sex: string
	id_num: string
	bed_name: string
	apply_flag: number // 审核状态：0-待审核 1-审核中 2-通过 -1-不通过
}

export interface ApplyForm {
	id?: number
	elder_id: number
	apply_reason: string
	apply_date: string
}

// ====== MOCK 数据（后台未完成，测试前端用）======
const mockApplyList: ApplyItem[] = [
	{ apply_id: 1, elder_id: 1, elder_name: '张建国', elder_sex: '男', id_num: '110101194801011234', bed_name: 'A栋-101-101室-A床', apply_flag: 0 },
	{ apply_id: 2, elder_id: 3, elder_name: '李德福', elder_sex: '男', id_num: '110101194507129012', bed_name: 'B栋-201-101室-A床', apply_flag: 2 }
]

/** 退住申请列表（全量） */
export function pageApplyByKey(keyword?: string): Promise<PageResult<ApplyItem>> {
	// return get<PageResult<ApplyItem>>('/retreatApply/pageRetreatApplyByKey', {
	// 	page_num: 1, page_size: 200, elder_name: keyword || undefined
	// })
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	let list = mockApplyList
	if (keyword) list = list.filter((it) => it.elder_name.indexOf(keyword) >= 0)
	return Promise.resolve(mockPage(list))
}

/** 新增退住申请 */
export function addApply(form: ApplyForm): Promise<void> {
	// return post<void>('/retreatApply/addRetreatApply', form)
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	return Promise.resolve()
}

/** 编辑退住申请 */
export function editApply(form: ApplyForm): Promise<void> {
	// return post<void>('/retreatApply/editRetreatApply', form)
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	return Promise.resolve()
}
