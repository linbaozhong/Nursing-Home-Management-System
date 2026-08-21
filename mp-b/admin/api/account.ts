/**
 * 账号相关接口
 * 对应后端 AccountController（/account/...）
 */
import { get, post } from '../utils/request'
import { setLoginUser, clearAuth, LoginUser, TenantVO } from '../utils/auth'

export interface LoginParams {
	phone: string
	pass: string
}

// ====== MOCK 数据（后台未完成，测试前端用；后台完成后放开各函数内注释并删除本段）======
const mockTenantList: TenantVO[] = [
	{ id: 1, name: '示例养老院', logo: '', contact_name: '张三', contact_phone: '13500000000', plan: '专业版', status: 1, trial_start: '2026-01-01', trial_end: '2027-01-01' }
]
const mockLoginUser: LoginUser = {
	id: 1,
	name: '管理员',
	avator: '',
	phone: '13800000000',
	tenant_id: 1,
	member_id: 1,
	role_id: 1,
	need_bind: false,
	auth_id_list: [],
	auth_url_list: ['/people/old', '/check-in/leave', '/check-in/visit', '/check-in/accident', '/food/order'],
	tenants: mockTenantList,
	token: 'mock-token-001'
}

/** 账号密码登录：成功后写入本地登录态 */
export async function login(params: LoginParams): Promise<LoginUser> {
	// const data = await post<LoginUser>('/account/login', params)
	// setLoginUser(data)
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	const mockUser: LoginUser = { ...mockLoginUser, phone: params.phone }
	setLoginUser(mockUser)
	return mockUser
}

export interface WxLoginResult {
	need_bind: boolean
	token: string
	user: LoginUser | null
	tenants: TenantVO[]
}

/** 微信静默登录（返回 need_bind / token / tenants 用于分流） */
export async function wxLogin(code: string): Promise<WxLoginResult> {
	// return post<WxLoginResult>('/account/wxLogin', { code })
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve({ need_bind: false, token: 'mock-token-001', user: mockLoginUser, tenants: mockTenantList })
}

/** 修改密码 */
export function editPass(id: number, oldPass: string, newPass: string): Promise<void> {
	// return post<void>('/account/edit', { id, old_pass: oldPass, new_pass: newPass })
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve()
}

/** 退出登录 */
export function logout(): Promise<void> {
	clearAuth()
	// return get<void>('/account/logout')
	// ====== MOCK（后台完成并放开上方注释后，可删除此段）======
	return Promise.resolve()
}
