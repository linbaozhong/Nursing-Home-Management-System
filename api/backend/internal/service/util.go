package service

import (
	"time"

	"github.com/linbaozhong/gentity/pkg/types"
)

// 公共辅助函数：service 包内多个文件共享，避免重复定义

func int64Ptr(v int64) *int64 { return &v }
func strPtr(v string) *string { return &v }

func parseTime(layout, value string) (time.Time, error) {
	return time.ParseInLocation(layout, value, time.Local)
}

func timeFormat(t types.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.String()
}
