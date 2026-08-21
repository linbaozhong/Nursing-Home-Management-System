/**
 * 租户相关接口
 * 对应后端 TenantController（/tenant/...）
 */
import { get, post } from '../utils/request'
import { setLoginUser, LoginUser } from '../utils/auth'

export interface TenantVO {
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
	tenants: TenantVO[]
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

/** 租户自助注册：成功后写入登录态 */
export async function registerTenant(params: RegisterParams): Promise<LoginUser> {
	const data = await post<LoginUser>('/tenant/register', params)
	setLoginUser(data)
	return data
}

/** 我的企业列表 */
export function myTenants(): Promise<UserTenantList> {
	return get<UserTenantList>('/tenant/myTenants')
}

/** 切换当前租户 */
export async function switchTenant(tenantId: number): Promise<LoginUser> {
	const data = await post<LoginUser>('/tenant/switchTenant', { tenant_id: tenantId })
	setLoginUser(data)
	return data
}

/** 加入已有企业 */
export async function joinMember(inviteCode: string): Promise<LoginUser> {
	const data = await post<LoginUser>('/tenant/joinMember', { invite_code: inviteCode })
	setLoginUser(data)
	return data
}

/** 邀请成员 */
export function inviteMember(params: { tenant_id: number; phone: string; role_id: number }): Promise<void> {
	return post<void>('/tenant/inviteMember', params)
}
