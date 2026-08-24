/**
 * 租户相关接口
 * 对应后端 TenantController（/tenant/...）
 */
import { get, post } from '../utils/request'
import { setLoginUser, LoginUser } from '../utils/auth'

export interface TenantResp {
	id: number
	name: string
	logo: string
	contact_name: string
	contact_phone: string
	plan: string
	status: number
	trial_start: string
	trial_end: string
}

export interface UserTenantList {
	tenants: TenantResp[]
	current: number
}

export interface RegisterParams {
	name: string
	logo?: string
	contact_name: string
	contact_phone: string
	password: string
	wx_code?: string
}

// ====== MOCK 数据（后台未完成，测试前端用；后台完成后放开各函数内注释并删除本段）======
// 模拟“多企业”场景：选择企业页会列出 2 家企业
const mockTenantList: TenantResp[] = [
	{ id: 1, name: '示例养老院', logo: '', contact_name: '张三', contact_phone: '13500000000', plan: '专业版', status: 1, trial_start: '2026-01-01', trial_end: '2027-01-01' },
	{ id: 2, name: '康泰护理中心', logo: '', contact_name: '李四', contact_phone: '13500000001', plan: '标准版', status: 1, trial_start: '2026-02-01', trial_end: '2027-02-01' }
]
const mockLoginUser: LoginUser = {
	id: 1,
	name: '管理员',
	avator: '',
	phone: '13800000000',
	tenant_id: 1,
	member_id: 1,
	role_id: 1,
	role_name: '管理员',
	need_bind: true,
	auth_id_list: [],
	auth_url_list: ['/people/old', '/check-in/leave', '/check-in/visit', '/check-in/accident', '/food/order', '/live/bed', '/live/enter', '/live/apply', '/serve/book', '/tenant/member', '/home/dashboard', '/audit'],
	tenants: mockTenantList,
	token: 'mock-token-001'
}

/** 租户自助注册：成功后写入登录态 */
export async function registerTenant(params: RegisterParams): Promise<LoginUser> {
	// const data = await post<LoginUser>('/tenant/register', params)
	// setLoginUser(data)
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	const mockUser: LoginUser = { ...mockLoginUser, name: params.name, phone: params.contact_phone }
	setLoginUser(mockUser)
	return mockUser
}

/** 我的企业列表 */
export function myTenants(): Promise<UserTenantList> {
	// return get<UserTenantList>('/tenant/myTenants')
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve({ tenants: mockTenantList, current: 1 })
}

/** 切换当前租户：改本地登录态里的当前企业 id */
export async function switchTenant(tenantId: number): Promise<LoginUser> {
	// const data = await post<LoginUser>('/tenant/switchTenant', { tenant_id: tenantId })
	// setLoginUser(data)
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	const mockUser: LoginUser = { ...mockLoginUser, tenant_id: tenantId }
	setLoginUser(mockUser)
	return mockUser
}

/** 加入已有企业 */
export async function joinMember(inviteCode: string): Promise<LoginUser> {
	// const data = await post<LoginUser>('/tenant/joinMember', { invite_code: inviteCode })
	// setLoginUser(data)
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	setLoginUser(mockLoginUser)
	return mockLoginUser
}

/** 邀请成员 */
export function inviteMember(params: { tenant_id: number; phone: string; role_id: number }): Promise<void> {
	// return post<void>('/tenant/inviteMember', params)
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve()
}
