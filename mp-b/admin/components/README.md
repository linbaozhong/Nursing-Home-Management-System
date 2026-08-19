# admin/components · 业务组件目录

> 本目录存放员工端小程序（`mp-b/admin`）**业务组件**，与第三方 `uni_modules/` 隔离。
> 组件 **props / 事件 / 状态字段级契约** 以规格文档 §9 为准：
> `mp-b/员工端P0六大模块-字段级规格.md`

## 目录结构（契约基线）

```
admin/components/
├── high-level/
│   ├── list-page/ListPage.uvue          # 列表页高阶模板（6 列表页复用）
│   └── wizard-form/WizardForm.uvue       # 多步向导模板（4 新增向导复用）
├── elder-search/ElderSearch.uvue
├── staff-search/StaffSearch.uvue
├── date-time-picker/DateTimePicker.uvue
├── image-uploader/ImageUploader.uvue
├── status-tag/StatusTag.uvue
├── dishes-picker/DishesPicker.uvue
├── contact-quick-fill/ContactQuickFill.uvue
└── section-card/SectionCard.uvue
```

## 约定

- 组件命名遵循 uni-app-x 规范（多词小驼峰，目录与组件名一致）。
- 页面中通过 `easycom` 自动引入 `components/` 下同名组件。
- 依赖的官方组件（`uni-nav-bar-x`、`uni-number-box-x`、`uni-tab-bar`、`uni-fab-button`、`uni-time-format`）从 `uni_modules/` 直接引用，不重复封装。

## 当前状态

骨架已搭建，部分组件已实现（**官方优先**：官方有基础组件则基于其实现，无则自建）。

| 组件 | 契约章节 | 实现状态 | 实现说明 |
|---|---|---|---|
| ListPage | §9.1 | ⬜ 待实现 | 列表页骨架暂由各页内联，后续抽高阶模板 |
| WizardForm | §9.2 | ✅ 已实现 | 多步向导模板，外出/来访/事故/点餐已复用；`next/prev` 由父组件校验后驱动 |
| ElderSearch | §9.3 | ✅ 已实现 | 自建半屏弹层，`pageSearchElderByKey` 全量本地过滤 |
| StaffSearch | §9.4 | ✅ 已实现 | 自建半屏弹层，`/reserve/pageSearchStaffByKey` |
| DateTimePicker | §9.5 | ✅ 已实现 | 基于原生 `picker`（date+time） |
| ImageUploader | §9.5 | ✅ 已实现 | `uni.chooseMedia` 选图→转 base64→`POST /file/uploadImage`，已接入事故编辑 |
| StatusTag | §9.6 | ✅ 已实现 | 自建文字标签（`uni-badge-view` 仅数字角标，不适用故自建） |
| DishesPicker | §9.7 | ✅ 已实现 | 半屏弹层菜品选择器，已接入点餐向导（`v-model:model-value`） |
| ContactQuickFill | §9.8 | ✅ 已实现 | 紧急联系人快捷填充，已接入外出向导 |
| SectionCard | §9.9 | ⬜ 待实现 | 详情分区块暂由页面内联 |
