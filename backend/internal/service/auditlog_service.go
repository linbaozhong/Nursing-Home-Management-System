package service

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"api/internal/lib"
	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblauditlog"

	"github.com/linbaozhong/gentity/pkg/ace"
	"github.com/linbaozhong/gentity/pkg/types"
)

// ==================== 审计日志（audit_log）写入工具 ====================
// 审计动作常量见 constant.AuditCreate / AuditUpdate / AuditDelete。
// 操作人姓名由鉴权中间件写入 ctx，经 lib.UserName(ctx) 读取（免查库）。

// fieldDictCache 字段中文名字典缓存：key = 表名 -> map[字段名]中文名。
// 按需懒加载：用到某表时才从 field_dict 表读入。
var fieldDictCache = map[string]map[string]string{}
var fieldDictMu sync.Mutex

// LoadFieldDict 从 field_dict 表读取某张表的字段中文名到 cache（列：table/field/label）。
func LoadFieldDict(ctx context.Context, table string) {
	fieldDictMu.Lock()
	defer fieldDictMu.Unlock()
	if _, ok := fieldDictCache[table]; ok {
		return // 已加载
	}
	m := map[string]string{}
	rows, e := db.QueryContext(ctx,
		"SELECT field, label FROM field_dict WHERE `table` = ?",
		table)
	if e != nil {
		fieldDictCache[table] = m
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
	fieldDictCache[table] = m
}

// AuditFieldLabel 将字段名（英文列名）转中文；该表字典未加载则懒加载一次。
func AuditFieldLabel(table, col string) string {
	if col == "" {
		return ""
	}
	fieldDictMu.Lock()
	m, ok := fieldDictCache[table]
	fieldDictMu.Unlock()
	if !ok {
		LoadFieldDict(context.Background(), table)
		fieldDictMu.Lock()
		m, _ = fieldDictCache[table]
		fieldDictMu.Unlock()
	}
	if v, ok := m[col]; ok && v != "" {
		return v
	}
	return col
}

// auditLogRow 审计日志写入参数
type auditLogRow struct {
	action      string
	table       string
	rowID       int64
	changeAfter string // 变更后整行快照(JSON 字符串)
	changeLabel string // 可读中文摘要
	comment     string
}

// WriteAuditLog 向 audit_log 写入一条操作日志。
// x 传 ace.Executer：业务在事务内则传 tx（与业务同一事务；推荐），否则传 db。
// after map[string]any 会被序列化为 change_after（key 保持英文列名）。
func WriteAuditLog(ctx context.Context, x ace.Executer, row auditLogRow, after map[string]any) error {
	var afterJSON string
	if after != nil {
		b, e := json.Marshal(after)
		if e != nil {
			afterJSON = ""
		} else {
			afterJSON = string(b)
		}
	}
	now := types.Time{Time: time.Now()}
	_, e := dao.AuditLog(x).Insert(ctx,
		tblauditlog.TenantId.Set(types.BigInt(lib.TenantID(ctx))),
		tblauditlog.OperatorId.Set(types.BigInt(lib.UserID(ctx))),
		tblauditlog.OperatorName.Set(types.String(lib.UserName(ctx))),
		tblauditlog.Table.Set(types.String(row.table)),
		tblauditlog.RowId.Set(types.BigInt(row.rowID)),
		tblauditlog.Action.Set(types.String(row.action)),
		tblauditlog.ChangeAfter.Set(types.String(afterJSON)),
		tblauditlog.ChangeLabel.Set(types.String(row.changeLabel)),
		tblauditlog.Comment.Set(types.String(row.comment)),
		tblauditlog.CreateTime.Set(now),
	)
	return e
}
