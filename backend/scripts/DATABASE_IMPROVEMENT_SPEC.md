# 养老院管理系统数据库结构 · 改造与验收方案

> 适用范围：关系型数据库 `snow_gerocomium`（MySQL 8.0）
> 关联文档：`Dump20260827.sql`（当前结构基线）、`field_dict_insert.sql`（字段中文名字典）
> 作者视角：高级产品专家 / 系统架构师
> 严重度：🔴 P0（影响正确性/资金/安全） · 🟠 P1（高维护成本/易错） · 🟡 P2（演进瓶颈）

---

## 目录

1. [总体结论](#1-总体结论)
2. [P0-1 账务一致性专项](#p0-1-账务一致性专项)
3. [P0-2 枚举与状态字典治理](#p0-2-枚举与状态字典治理)
4. [P0-3 敏感数据（身份证/手机号）处理](#p0-3-敏感数据身份证手机号处理)
5. [P1-1 DDL 改进建议清单](#p1-1-ddl-改进建议清单)
6. [P1-2 命名规范](#p1-2-命名规范)
7. [P1-3 索引与外键策略](#p1-3-索引与外键策略)
8. [P2-1 附件模型抽象](#p2-1-附件模型抽象)
9. [P2-2 Schema 迁移机制](#p2-2-schema-迁移机制)
10. [改造优先级与验收清单](#10-改造优先级与验收清单)

---

## 1. 总体结论

- **定性**：这是一套「业务功能完整度高、数据治理与账务正确性偏弱」的一代系统。
- **亮点**：多租户模型收敛干净；`audit_log` 审计设计优秀；金额按「分」以整数存储；区分软删除(`state`)与业务状态(`status`)；关键单据做了快照冗余。
- **最大风险**：`elder.balance` 无对账约束、枚举取值散落且语义错乱、敏感字段加密策略不一致、关联表缺索引、无 schema 迁移机制。

> 结论一句话：**功能面覆盖了养老院几乎所有核心痛点，能支撑产品落地；但要承载真实预存资金、医养责任与多租户增长，账务、字典、索引、迁移四件事必须在数据量起来之前补完，否则系统会随增长快速劣化。**

---

## 2. 🔴 P0-1 账务一致性专项

### 现状问题

| 点 | 说明 |
|---|---|
| 账户余额无约束 | `elder.balance` 可被 `consume`、`family_recharge`、`bill` 多次扣减，表间无任何一致性保障 |
| 流水无溯源 | `consume` 只有单表，被点餐/护理/杂费多来源写入，**无 `source_type + source_id` 溯源** |
| 无对账机制 | 账不平后无手段定位差异 |
| 并发安全 | 余额增减若无事务+行锁，高并发下丢失更新 |

### 目标模型（推荐：总账 + 明细台账 + 对账）

```
                         ┌─────────────────────────┐
   elder.balance (科目余额) │                         │
                         └──────────┬──────────────┘
                                    │ 受事务约束，恒等:
                              balance = 期初 + Σ(贷方) - Σ(借方)
                                    │
                    ┌───────────────▼───────────────┐
                    │  elder_account_ledger (明细台账) │
                    │  ─ 每笔变动一行, append-only    │
                    │  ─ 唯一业务流水号 ledger_no      │
                    └───────────────┬───────────────┘
        ┌──────────────────────────┼──────────────────────────┐
        │ source_type + source_id  │ 引用来源单据              │
   ┌────▼─────┐   ┌───────────┐   ┌──────────────┐   ┌───────▼─────┐
   │ billing  │   │ recharge  │   │ consume(点餐) │   │  退款/退住   │
   │(账单支付) │   │(家属充值)  │   │ 护理/杂费     │   │  (负数)      │
   └──────────┘   └───────────┘   └──────────────┘   └─────────────┘
```

### 新增/改造表

```sql
-- ── 明细台账（新增，核心）───────────────────────────────
CREATE TABLE `elder_account_ledger` (
  `id`            bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id`     bigint unsigned NOT NULL DEFAULT '1',
  `elder_id`      bigint unsigned NOT NULL,
  `direction`     tinyint NOT NULL COMMENT '1=入账(充值/退款入账)，2=出账(消费/扣费)',
  `amount`        bigint NOT NULL COMMENT '变更金额（分），恒为正',
  `balance_after` bigint NOT NULL COMMENT '变动后的账户余额（分）',
  `source_type`   varchar(32) NOT NULL COMMENT '来源类型：BILL_INCOME/RECHARGE/FEED/NURSING/REFUND...',
  `source_id`     bigint unsigned NOT NULL COMMENT '来源业务表主键id',
  `business_no`   varchar(64) NOT NULL COMMENT '业务单号（可空，用于人工对账）',
  `remark`        varchar(255) NOT NULL DEFAULT '',
  `operator_id`   bigint unsigned NOT NULL,
  `create_time`   datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_source` (`source_type`,`source_id`),   -- 防重复记账（幂等）
  KEY `idx_elder_time` (`elder_id`,`create_time`),
  CONSTRAINT `fk_led_elder` FOREIGN KEY (`elder_id`) REFERENCES `elder` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='老人资金明细台账';
```

> 核心保证：任何改变余额的操作，**必须在同一个数据库事务里**同时写 `elder_account_ledger` 并 `UPDATE elder.balance = balance ± amount`，并给 `elder` 加行锁（`SELECT ... FOR UPDATE`）。`uk_source` 唯一键天然防重复入账（幂等）。

### 一致性校验 SQL（对账脚本，定期跑）

```sql
-- 校验1：余额 = 期初 + 入账合计 - 出账合计
SELECT e.id, e.balance,
       e.balance - COALESCE((
         SELECT SUM(CASE WHEN l.direction=1 THEN l.amount ELSE -l.amount END)
         FROM elder_account_ledger l WHERE l.elder_id = e.id
       ), 0) AS diff
FROM elder e
HAVING ABS(diff) > 0;

-- 校验2：每笔台账的 balance_after 链路自洽
-- （在应用层/批处理中顺序复查，或定期全量重算比对）
```

### `consume` 表的最小改造

```sql
ALTER TABLE `consume`
  ADD COLUMN `source_type` varchar(32) NOT NULL DEFAULT 'MANUAL' COMMENT '来源：ORDER/NURSE_RESERVE/MANUAL', 
  ADD COLUMN `source_id`   bigint unsigned DEFAULT NULL COMMENT '来源业务主键id',
  ADD COLUMN `out_trade_no` varchar(64) DEFAULT NULL COMMENT '外部交易单号（对账）',
  ADD KEY `idx_elder_time` (`elder_id`,`consume_date`),
  ADD KEY `idx_source` (`source_type`,`source_id`);
```

---

## 3. 🔴 P0-2 枚举与状态字典治理

### 现状问题
数十张表的 `status`/`state`/`type`/`fee_type`/`calculation_type` 取值含义各表不同（`bed.status` 的 0/1/2 与 `bill.status` 的 0/1/2/3 完全不同），全靠注释说明，前端/客户端只能硬编码，极易读错、写错。

### 目标模型（两级字典）

**方案A：统一 `enum_dict` 表（推荐给变化多的业务状态）**

```sql
CREATE TABLE `enum_dict` (
  `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
  `group_code` varchar(50) NOT NULL COMMENT '枚举组，如 BED_STATUS/BILL_STATUS/FEE_TYPE',
  `value`      int NOT NULL COMMENT '存储值',
  `label`      varchar(50) NOT NULL COMMENT '中文名',
  `color`      varchar(20) DEFAULT '' COMMENT '前端标签色',
  `sort`       int NOT NULL DEFAULT '0',
  `state`      tinyint NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_group_value` (`group_code`,`value`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='业务状态/枚举字典';
```

**方案B：代码常量枚举（给稳定状态，强烈推荐，与A互补）**
在 Go 层用 `const` 定义每个状态语义，作为**唯一事实来源**，DB 注释只做给人看的说明，前端下拉由后端 `/dict/{group}` 接口下发（读 `enum_dict`）。**禁止前端硬编码数字含义。**

### 配套：规范化 `field_dict`

- 清洗掉「一个 label 里塞整段注释」的做法，只保留简短中文名（如 `管理状态`），详细取值说明放 `enum_dict`。
- `field_dict` 增加 `field_group`/`data_type` 等维度以支持动态表单/明细展示（可选演进）。

---

## 4. 🔴 P0-3 敏感数据（身份证/手机号）处理

### 现状问题
`staff.id_num`、`elder.id_num`、`family_member.id_num` 明文 varchar；`visit.id_card` 注释「加密」但实为明文可存，策略不一且无标记。

### 方案

1. **统一脱敏/加密策略**：身份证、手机号等敏感字段采用 AES-256 加密存储 + 展示时按需脱敏（如 `511901********3448`）。
2. **约定字段后缀**：加密字段名统一用 `*_enc`（如 `id_num_enc`），或至少在 `field_dict` 里用 `sensitive=1` 标记，避免散落语义。
3. **禁止手机号明文入库**（用于营销触达需注意合规：同意书 + 最小化）。
4. 存量数据迁移脚本：写一次性加密回填，随后 `RENAME` 覆盖旧列并 `DROP` 明文。

---

## 5. 🟠 P1-1 DDL 改进建议清单

以下按业务表逐条列出「最小必要」改造，避免大改造成风险。

### 5.1 `accident`（事故登记，医养责任关键）
```sql
ALTER TABLE `accident`
  ADD COLUMN `severity`     tinyint DEFAULT NULL COMMENT '严重程度：1轻/2中/3重',
  ADD COLUMN `handle_result` varchar(255) DEFAULT NULL COMMENT '处理结果/整改措施',
  ADD COLUMN `handle_staff_id` bigint unsigned DEFAULT NULL COMMENT '处理责任人id',
  ADD KEY `idx_elder`  (`elder_id`),
  ADD KEY `idx_staff`  (`staff_id`),
  ADD KEY `idx_occur`  (`occur_date`);
```
> 事故是医养机构核心责任风险，必须有「发生-严重度-处理-责任关人」闭环，且按老人/时间可排查。

### 5.2 `order` / `order_dishes` / `consume`（餐饮账务）
```sql
ALTER TABLE `order`
  ADD COLUMN `out_trade_no` varchar(64) DEFAULT NULL COMMENT '对外交易单号',
  ADD KEY `idx_elder_time` (`elder_id`,`create_time`);

ALTER TABLE `order_dishes`
  ADD KEY `idx_order` (`order_id`);
```

### 5.3 `nurse_reserve` / `nurse` / `nurse_grade`
```sql
ALTER TABLE `nurse_reserve`
  ADD KEY `idx_elder_time` (`elder_id`,`create_time`),
  ADD KEY `idx_status` (`status`);
```

### 5.4 通用审计补齐（可选，若未启用 audit_log）
对资金/责任类写操作，确保持久化落到 `audit_log`；建议 `change_after` 同时存英文 key，展示时查 `field_dict`（避免历史 change_label 过期）。

### 5.5 `visit.id_card` 处理
```sql
-- 若已启用加密：改为 *_enc 列并回填密文；未启用则至少强制非空/格式校验
```

---

## 6. 🟠 P1-2 命名规范

统一以下规则（存量逐步对齐，新表强制执行）：

| 规则 | 说明 | 反例 ➜ 正例 |
|---|---|---|
| 表名单数 + snake_case | 表名统一单数 | `outward` 语义含糊，建议业务语义化 |
| 关联表 + 业务表 | 关联表 `xxx_yyy` | — |
| 时间戳 `*_time` vs `*_date` | 精确时刻用 `*_time`，仅日期用 `*_date` | `health_data.create_time`（实为取样时间）➜ `record_time` |
| 外键字段后缀 `*_id` | 指向哪张表一目了然 | `warehouse_record.source`(文本) 与 `consult.source_id`(外键) 混用 ➜ 前者改名 `source_name` |
| 布尔/标志用 `is_*`/`flag` | `state` 仅表逻辑删除；业务标志不再复用 | — |
| 头像 avatar | 修正拼写 `avator` ➜ `avatar` |
| 避免列表 / 表同名字段 | `communication_record.communication_record` ➜ `content` |
| 敏感字段标记 | 加密字段统一 `*_enc` 或字典标 sensitive |

> 原则：**一个词一个语义，一个语义一个词**。将 `source`、`status` 这类重灾区做全局消歧映射。

---

## 7. 🟠 P1-3 索引与外键策略

> ✅ **状态：已完成**（由项目侧实现）
> 索引与主外键已按以下策略补齐。如后续需复核已落库的 DDL，可对照下方清单执行 `SHOW INDEX FROM <表>;` 校验。

### 经验法则
- 每张表都应有：`PRIMARY KEY` + 至少一个可支撑查询的二级索引。
- 高频按「租户+业务维度+时间」查询的表，建立联合索引 `(tenant_id, xxx_id, create_time)`。
- 关系表(如 `role_auth`、`nurse_item`、`order_dishes`) 给外键列加索引。
- **业务关系加外键(FK)或至少在应用层强约束 + 全部加索引**。对资金/责任表（`accident`、`bill`、`order`、`consume`）建议加 FK。

### 建议统一增加的二级索引清单

| 表 | 建议索引 |
|---|---|
| `accident` | (tenant_id, elder_id, occur_date) |
| `consume` | (tenant_id, elder_id, consume_date) |
| `order` | (tenant_id, elder_id, create_time) |
| `order_dishes` | (order_id) |
| `nurse_reserve` | (tenant_id, elder_id, create_time), (status) |
| `visit` / `outward` / `communication_record` | (tenant_id, elder_id, *_date) |
| `audit_log` | (table, row_id) 已有；补充 (create_time) |
| `bill_item` | (bill_id) 已有 |

---

## 8. 🟡 P2-1 附件模型抽象

把所有 `picture/avator/logo/*_picture` 字符串收敛到 `base_attachment`，建立多对多关联：

```sql
CREATE TABLE `biz_attachment` (
  `id`            bigint unsigned NOT NULL AUTO_INCREMENT,
  `tenant_id`     bigint unsigned NOT NULL DEFAULT '1',
  `biz_type`      varchar(32) NOT NULL COMMENT '业务类型：ELDER_AVATAR/ACCIDENT_PIC/ACTIVE_PIC/STAFF_AVATAR',
  `biz_id`        bigint unsigned NOT NULL COMMENT '业务主键id',
  `attachment_id` bigint unsigned NOT NULL COMMENT 'base_attachment.id',
  `sort`          int NOT NULL DEFAULT '0',
  `create_time`   datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_biz` (`biz_type`,`biz_id`),
  CONSTRAINT `fk_bizatt_att` FOREIGN KEY (`attachment_id`) REFERENCES `base_attachment` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='业务附件关联';
```

> 后续切 OSS/对象存储只需迁移 `base_attachment` 的 path/url，业务表零改动。

---

## 9. 🟡 P2-2 Schema 迁移机制

项目为 Go 技术栈，建议引入版本化迁移，替代「人工改表 + 依赖 dump」：

- **方案推荐**：`golang-migrate/migrate` 或 `pressly/goose`，目录 `backend/migrations/`。
- 约定：`0001_init_schema.up.sql / .down.sql` 按日期递增。
- 好处：生产/测试环境结构一致、可回滚、可审计；`Dump20260827.sql` 退化为初始化种子/基线，不再作为唯一结构来源。

---

## 10. 改造优先级与验收清单

### 排序（按「风险/成本」比）

| 优先级 | 事项 | 预估工作量 | 上线前后置 |
|---|---|---|---|
| **P0-A** | 账务专项（ledger + 事务 + 幂等 + 对账脚本） | 高 | 上线前必须 |
| **P0-B** | 枚举字典（enum_dict + 接口下发 + 去硬编码） | 中 | 上线前必须 |
| **P0-C** | 敏感字段加密策略统一 | 中 | 上线前必须 |
| **P1-A** | 关键索引与 FK 补齐（accident/order/consume/nurse_reserve...） | 低 | ✅ 已完成 |
| **P1-B** | 命名消歧（source/status/id_num 等重灾区） | 中 | 随重构逐步 |
| **P2-A** | 附件抽象 `biz_attachment` | 低 | 切对象存储前 |
| **P2-B** | 引入 `goose/golang-migrate` | 低 | 多环境前 |

### 验收清单（可执行验收项）

- [ ] 对账脚本连续运行 3 天余额差异 = 0
- [ ] `uk_source` 幂等：同一 source_type+source_id 重复写账不产生第二条
- [ ] 枚举差异：`bed.status`、`bill.status` 等均能从 `/dict/{group}` 取到，前端无硬编码数字
- [ ] 身份证/手机号在 DB 中为密文，接口返回为脱敏
- [x] `accident`、`consume`、`order`、`nurse_reserve` 按 elder_id + time 查询走索引（EXPLAIN 无全表扫描）
- [ ] 命名重灾区（source 的双语义等）已消歧，schema 文档与字典同步
- [ ] 新表全部纳入 goose/golang-migrate，`up`/`down` 可往返

---

*本文档为架构级建议，具体改动请结合业务验收与回滚预案实施。*
