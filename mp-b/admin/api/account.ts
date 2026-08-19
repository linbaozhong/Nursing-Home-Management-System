/**
 * 账号相关接口
 * 对应后端 AccountController（/account/...）
 */
import { get, post } from '../utils/request'
import { setLoginUser, clearAuth, LoginUser } from '../utils/auth'

export interface LoginParams {
	phone: string
	pass: string
}

/** 登录：成功后写入本地登录态 */
export async function login(params: LoginParams): Promise<void> {
	const data = await post<LoginUser>('/account/login', params)
	setLoginUser(data)
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
