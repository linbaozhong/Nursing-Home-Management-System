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

/** 账号密码登录：成功后写入本地登录态 */
export async function login(params: LoginParams): Promise<LoginUser> {
	const data = await post<LoginUser>('/account/login', params)
	setLoginUser(data)
	return data
}

export interface WxLoginResult {
	need_bind: boolean
	token: string
	user: LoginUser | null
	tenants: TenantVO[]
}

/** 微信静默登录（返回 need_bind / token / tenants 用于分流） */
export async function wxLogin(code: string): Promise<WxLoginResult> {
	return post<WxLoginResult>('/account/wxLogin', { code })
}

/** 修改密码 */
export function editPass(id: number, oldPass: string, newPass: string): Promise<void> {
	return post<void>('/account/edit', { id, old_pass: oldPass, new_pass: newPass })
}

/** 退出登录 */
export function logout(): Promise<void> {
	clearAuth()
	return get<void>('/account/logout')
}
