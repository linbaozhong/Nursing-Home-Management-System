package service

// clampPage 统一校正分页参数：页数最小 1；每页上限 PageSizeMax，避免移动端全量加载时超界
// 移动端取消分页，默认取 PageSizeMax 条全量后在本地过滤。
const PageSizeMax = 200

func clampPage(pageNum, pageSize *int) {
	if pageNum == nil || *pageNum < 1 {
		*pageNum = 1
	}
	if pageSize == nil || *pageSize < 1 {
		*pageSize = PageSizeMax
	}
	if *pageSize > PageSizeMax {
		*pageSize = PageSizeMax
	}
}
