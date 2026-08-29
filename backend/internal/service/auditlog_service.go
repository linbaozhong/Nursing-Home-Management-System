package service

import (
	"context"
	"encoding/json"
	"github.com/linbaozhong/gentity/pkg/cachego/mmap"
	"sync"
	"time"

	"api/internal/lib"
	"api/internal/model/define/dao"
	"api/internal/model/do"

	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/types"
)

// ==================== 审计日志（audit_log）写入工具 ====================
// 审计动作常量见 constant.AuditCreate / AuditUpdate / AuditDelete。
// 操作人姓名由鉴权中间件写入 ctx，经 lib.UserName(ctx) 读取（免查库）。

// fieldDictCache 字段中文名字典缓存：key = 表名 -> map[字段名]中文名。
// 按需懒加载：用到某表时才从 field_dict 表读入。
var fieldDictCache = mmap.New[map[string]string](mmap.WithExpired(time.Minute * 10))
var fieldDictMu sync.Mutex

// LoadFieldDict 从 field_dict 表读取某张表的字段中文名到 cache（列：table/field/label）。
func LoadFieldDict(ctx context.Context, table string) {
	if _, e := fieldDictCache.Fetch(ctx, table); e == nil {
		return // 已加载
	}
	var m = make(map[string]string)
	rows, e := db.QueryContext(ctx,
		"SELECT field, label FROM field_dict WHERE `table` = ?",
		table)
	if e != nil {
		fieldDictCache.Save(ctx, table, m)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var fn, fl string
		if rows.Scan(&fn, &fl) == nil {
			if fl != "" {
				m[fn] = fl
			}
		}
	}
	fieldDictCache.Save(ctx, table, m)
}

// AuditFieldLabel 将字段名（英文列名）转中文；该表字典未加载则懒加载一次。
func AuditFieldLabel(ctx context.Context, table, col string) string {
	if col == "" {
		return ""
	}

	m, e := fieldDictCache.Fetch(ctx, table)
	if e != nil {
		LoadFieldDict(context.Background(), table)
		fieldDictMu.Lock()
		m, _ = fieldDictCache.Fetch(ctx, table)
		fieldDictMu.Unlock()
	}
	if v, ok := m[col]; ok && v != "" {
		return v
	}
	return col
}

// WriteAuditLog 向 audit_log 写入一条操作日志。
// x 传 ace.Executer：业务在事务内则传 tx（与业务同一事务；推荐），否则传 db。
// after map[string]any 会被序列化为 change_after（key 保持英文列名）。
func WriteAuditLog(ctx context.Context, x ace.Executer,
	table string, rowID int64, action, label, comment string, after map[string]any) error {

	var afterJSON string
	if after != nil {
		if b, e := json.Marshal(after); e == nil {
			afterJSON = string(b)
		}
	}
	row := do.NewAuditLog()
	defer row.Free()

	row.TenantId = types.BigInt(lib.TenantID(ctx))
	row.OperatorId = types.BigInt(lib.UserID(ctx))
	row.OperatorName = types.String(lib.UserName(ctx))
	row.Table = types.String(table)
	row.RowId = types.BigInt(rowID)
	row.Action = types.String(action)
	row.ChangeAfter = types.String(afterJSON)
	row.ChangeLabel = types.String(label)
	row.Comment = types.String(comment)
	row.CreateTime = types.Time{Time: time.Now()}

	_, e := dao.AuditLog(x).InsertOne(ctx, row)
	return e
}
