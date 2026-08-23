/**
 * 退住审核相关接口
 * 对应后端 RetreatAuditController（/retreatAudit/...）
 * 移动端：待审核列表 + 通过/驳回审核
 */
import { get, post } from '../utils/request'
import { PageResult, mockPage } from './elder'

export interface AuditItem {
	id: number
	elder_name: string
	apply_flag: string // 中文：待审核/审核中/通过/不通过
	apply_name: string // 申请人姓名
}

export interface AuditDetail {
	id: number
	elder_id: number
	elder_name: string
	apply_flag: string
}

// ====== MOCK 数据（后台未完成，测试前端用）======
const mockAuditList: AuditItem[] = [
	{ id: 1, elder_name: '张建国', apply_flag: '待审核', apply_name: '张三' },
	{ id: 2, elder_name: '李德福', apply_flag: '审核中', apply_name: '张三' }
]

/** 退住审核列表（全量） */
export function pageAuditByKey(keyword?: string): Promise<PageResult<AuditItem>> {
	// return get<PageResult<AuditItem>>('/retreatAudit/pageRetreatAuditByKey', {
	// 	page_num: 1, page_size: 200, elder_name: keyword || undefined
	// })
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	let list = mockAuditList
	if (keyword) list = list.filter((it) => it.elder_name.indexOf(keyword) >= 0)
	return Promise.resolve(mockPage(list))
}

/** 审核退住：通过（audit_result="通过"）/ 驳回（audit_result="不通过"） */
export function auditRetreat(id: number, pass: boolean, remark: string): Promise<void> {
	// return post<void>('/retreatAudit/auditRetreat', {
	// 	id,
	// 	audit_result: pass ? '通过' : '不通过',
	// 	audit_remark: remark
	// })
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	return Promise.resolve()
}
