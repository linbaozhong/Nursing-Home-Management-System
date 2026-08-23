/**
 * 全局 uni API 类型声明
 * uni-app-x 对 .uvue 内联脚本自动注入 uni 全局类型，
 * 但对独立 .ts/.uts 工具文件需要显式声明，否则 TS 报 "Cannot find name 'uni'"。
 * 此处声明为宽泛类型，运行时由 uni-app-x 提供真实实现。
 */
declare const uni: any
