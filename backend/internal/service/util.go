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

// ptrI8v 解引用 int8 指针；nil 时返回 0
func ptrI8v(p *int8) int8 {
	if p == nil {
		return 0
	}
	return *p
}

// ptrI64v 解引用 int64 指针；nil 时返回 0
func ptrI64v(p *int64) int64 {
	if p == nil {
		return 0
	}
	return *p
}
