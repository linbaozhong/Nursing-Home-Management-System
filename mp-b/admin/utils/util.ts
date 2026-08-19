/**
 * 通用工具：日期格式化 / 金额 / 手机号脱敏
 */

function pad(n: number): string {
	return n < 10 ? '0' + n : '' + n
}

/** Date -> "yyyy-MM-dd HH:mm" */
export function formatDateTime(d: Date): string {
	if (d == null) {
		return ''
	}
	return d.getFullYear() + '-' + pad(d.getMonth() + 1) + '-' + pad(d.getDate()) +
		' ' + pad(d.getHours()) + ':' + pad(d.getMinutes())
}

/** Date -> "MM-dd HH:mm" */
export function formatShortTime(d: Date): string {
	if (d == null) {
		return ''
	}
	return pad(d.getMonth() + 1) + '-' + pad(d.getDate()) +
		' ' + pad(d.getHours()) + ':' + pad(d.getMinutes())
}

/** 后端 time.Time 字符串 -> Date */
export function parseTime(s: string): Date | null {
	if (s == null || s == '') {
		return null
	}
	// 兼容 "2006-01-02 15:04:05" 与 ISO
	const t = new Date(s.replace(' ', 'T'))
	return isNaN(t.getTime()) ? null : t
}

/** 金额分 -> "¥xx.xx"（传入数字，单位元） */
export function formatMoney(v: number | null | undefined): string {
	if (v == null) {
		return '¥0.00'
	}
	return '¥' + v.toFixed(2)
}

/** 手机号脱敏 138****0000 */
export function maskPhone(phone: string): string {
	if (phone == null || phone.length != 11) {
		return phone || ''
	}
	return phone.substring(0, 3) + '****' + phone.substring(7)
}

/** 身份证尾 4 位 */
export function idNumTail(idNum: string): string {
	if (idNum == null || idNum.length < 4) {
		return idNum || ''
	}
	return idNum.substring(idNum.length - 4)
}
