/**
 * 成员管理相关接口
 * 对应后端 TenantController（/tenant/...）
 * 移动端：租户管理员查看企业成员列表 + 邀请新成员
 */
import { post } from '../utils/request'
// import { getLoginUser } from '../utils/auth' // 接入真实后端时放开

export interface MemberItem {
	id: number
	user_id: number
	name: string
	phone: string
	role_id: number
	role_name: string // 前端展示用（mock 提供；真实后端需补充角色名映射）
	status: number
	auth_urls: string[]
}

export interface RoleOption {
	id: number
	name: string
}

// ====== MOCK 数据（后台未完成，测试前端用）======
const mockRoleList: RoleOption[] = [
	{ id: 1, name: '管理员' },
	{ id: 2, name: '护士' },
	{ id: 3, name: '营销' }
]

const mockMemberList: MemberItem[] = [
	{ id: 1, user_id: 1, name: '管理员', phone: '13800000000', role_id: 1, role_name: '管理员', status: 0, auth_urls: [] },
	{ id: 2, user_id: 2, name: '赵护士', phone: '13700000001', role_id: 2, role_name: '护士', status: 0, auth_urls: [] },
	{ id: 3, user_id: 3, name: '孙护工', phone: '13700000003', role_id: 2, role_name: '护士', status: 1, auth_urls: [] }
]

/** 角色下拉（邀请成员时选择） */
export function listRoles(): Promise<RoleOption[]> {
	// return get<RoleOption[]>('/tenant/listRole')
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	return Promise.resolve(mockRoleList)
}

/**
 * 当前企业成员列表
 * ⚠️ 后端暂无 listMembers 接口，此处 mock；需新增后端接口后放开注释
 */
export function listMembers(): Promise<MemberItem[]> {
	// return get<MemberItem[]>('/tenant/listMember', { tenant_id: getLoginUser()?.tenant_id })
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	return Promise.resolve(mockMemberList.filter((m) => m.status == 0))
}

/** 邀请成员加入当前企业 */
export function inviteMember(phone: string, roleId: number): Promise<void> {
	// return post<void>('/tenant/inviteMember', {
	// 	tenant_id: getLoginUser()?.tenant_id,
	// 	phone,
	// 	role_id: roleId
	// })
	// ====== MOCK（接入真实后端时放开上方注释并删掉本段）======
	return Promise.resolve()
}
