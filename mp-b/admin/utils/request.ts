/**
 * 统一请求封装
 * - BaseURL：Go 后端路由全带 /v1 前缀
 * - 自动注入 Authorization: Bearer <token>
 * - 统一处理 Result { code, msg, data }；code!=0 toast；401 跳登录
 */

const BASE_URL = 'https://<域名>/v1'

export interface Result<T> {
	code: number
	msg: string
	data: T
}

export function getToken(): string {
	return uni.getStorageSync('token') as string || ''
}

function realUrl(url: string): string {
	if (url.startsWith('http')) {
		return url
	}
	return BASE_URL + url
}

export function request<T>(options: {
	url: string
	method?: 'GET' | 'POST'
	data?: any
	header?: Record<string, string>
	showLoading?: boolean
}): Promise<T> {
	const {
		url,
		method = 'GET',
		data,
		header = {},
		showLoading = false
	} = options

	if (showLoading) {
		uni.showLoading({ title: '加载中', mask: true })
	}

	return new Promise<T>((resolve, reject) => {
		const token = getToken()
		const mergedHeader: Record<string, string> = {
			'Content-Type': 'application/json',
			...header
		}
		if (token != '') {
			mergedHeader['Authorization'] = 'Bearer ' + token
		}

		uni.request({
			url: realUrl(url),
			method: method as any,
			data: data,
			header: mergedHeader,
			success: (res) => {
				if (showLoading) {
					uni.hideLoading()
				}
				const body = res.data as Result<T>
				if (body == null) {
					uni.showToast({ title: '请求异常', icon: 'none' })
					reject(new Error('empty response'))
					return
				}
				if (body.code == 0) {
					resolve(body.data)
					return
				}
				// 401 未登录
				if (body.code == 401) {
					uni.removeStorageSync('token')
					uni.reLaunch({ url: '/pages/login/login' })
				} else {
					uni.showToast({ title: body.msg || '请求失败', icon: 'none' })
				}
				reject(new Error(body.msg))
			},
			fail: (err) => {
				if (showLoading) {
					uni.hideLoading()
				}
				uni.showToast({ title: '网络异常', icon: 'none' })
				reject(err)
			}
		})
	})
}

export function get<T>(url: string, data?: any, header?: Record<string, string>): Promise<T> {
	return request<T>({ url, method: 'GET', data, header })
}

export function post<T>(url: string, data?: any, header?: Record<string, string>): Promise<T> {
	return request<T>({ url, method: 'POST', data, header })
}
