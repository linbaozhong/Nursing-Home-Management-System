/**
 * 登录信息 / 权限存取
 * 对应后端 dto.LoginUserVO
 */
// 全局 uni API（uni-app-x 运行时提供，此处供 TS 工具文件识别）
declare const uni: any

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

export interface LoginUser {
	id: number
	name: string
	avator: string
	phone: string
	tenant_id: number
	member_id: number
	role_id: number
	role_name: string
	need_bind: boolean
	auth_id_list: number[]
	auth_url_list: string[]
	tenants: TenantVO[]
	token: string
}

const TOKEN_KEY = 'token'
const USER_KEY = 'login_user'
const TENANT_KEY = 'current_tenant'

export function setToken(token: string): void {
	uni.setStorageSync(TOKEN_KEY, token)
}

export function getToken(): string {
	return uni.getStorageSync(TOKEN_KEY) as string || ''
}

export function setLoginUser(user: LoginUser): void {
	setToken(user.token)
	uni.setStorageSync(USER_KEY, user)
	if (user.tenant_id && user.tenant_id > 0) {
		uni.setStorageSync(TENANT_KEY, user.tenant_id)
	}
}

export function getLoginUser(): LoginUser | null {
	const raw = uni.getStorageSync(USER_KEY)
	if (raw == null || raw == '') {
		return null
	}
	return raw as LoginUser
}

/** 当前选择的租户编号（记住上次选择） */
export function getCurrentTenantId(): number {
	return (uni.getStorageSync(TENANT_KEY) as number) || 0
}

export function setCurrentTenantId(id: number): void {
	uni.setStorageSync(TENANT_KEY, id)
}

export function clearAuth(): void {
	uni.removeStorageSync(TOKEN_KEY)
	uni.removeStorageSync(USER_KEY)
	uni.removeStorageSync(TENANT_KEY)
}

/** 是否拥有指定权限 id */
export function hasAuth(authId: number): boolean {
	const u = getLoginUser()
	if (u == null) {
		return false
	}
	return u.auth_id_list.indexOf(authId) >= 0
}

/** 是否已登录 */
export function isLogin(): boolean {
	return getToken() != ''
}

/**
 * 模块权限点（对应后端 auth 表 url 字段，登录返回 auth_url_list）
 * P0 六模块映射：
 *   /people/old        长者档案
 *   /check-in/leave    外出登记
 *   /check-in/visit    来访登记
 *   /check-in/accident 事故登记
 *   /food/order        点餐
 */
export type ModuleKey = 'elder' | 'leave' | 'visit' | 'accident' | 'order'

const MODULE_AUTH_URL: Record<ModuleKey, string> = {
	elder: '/people/old',
	leave: '/check-in/leave',
	visit: '/check-in/visit',
	accident: '/check-in/accident',
	order: '/food/order'
}

/** 判断当前登录用户是否有某模块权限（基于 auth_url_list 精确匹配）。
 * 可传入已取好的 user 以避免重复读取本地存储；不传则内部取一次。 */
export function canAccess(module: ModuleKey, user?: LoginUser | null): boolean {
	const u = user === undefined ? getLoginUser() : user
	if (u == null) {
		return false
	}
	return (u.auth_url_list ?? []).indexOf(MODULE_AUTH_URL[module]) >= 0
}

/** 返回当前用户可访问的模块列表（顺序：档案/外出/来访/事故/点餐）。
 * 可传入已取好的 user 以避免重复读取本地存储；不传则内部取一次。 */
export function accessibleModules(user?: LoginUser | null): ModuleKey[] {
	const u = user === undefined ? getLoginUser() : user // 只读一次本地存储，循环内复用
	const order: ModuleKey[] = ['elder', 'leave', 'visit', 'accident', 'order']
	return order.filter((m) => canAccess(m, u))
}
