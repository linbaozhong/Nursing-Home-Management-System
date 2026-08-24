# gentity 重新生成清单（床位床型 / 房间设施 / 楼层回写）

本清单对应本次后端改动。**必须先执行数据库变更，再用 gentity 重新生成**，
否则 service 里引用的新字段/新表（`BedTypeId`、`Kind`、`RoomMaterial`）编译不过。

## 一、数据库变更

1. 在数据库中执行 `backend/scripts/0002_bed_room_material.sql`（或手动执行下述 3 条变更）：
   - `material_type` 表新增列 `kind tinyint NOT NULL DEFAULT 99`（1=床型，99=设施/其他）
   - `bed` 表新增列 `bed_type_id bigint DEFAULT NULL`（关联 material_type.id，kind=1）
   - 新建表 `room_material`（room_id + material_type_id，见 SQL）

2. 如需从零重建库，使用已同步修改 `kind`/`bed_type_id`/`room_material` 的
   `backend/internal/model/db_gerocomium.sql`。

## 二、gentity 重新生成

让 gentity 重新扫描上述表结构，覆盖/新增以下 `.gen.go`（全部 `DO NOT EDIT`，勿手改）：

### 变更（已有表，需重生成）
- `internal/model/define/table/tblmaterialtype/gentity_model_tbl_MaterialType.gen.go`  → 新增 `Kind` 字段
- `internal/model/do/gentity_model_do_MaterialType.gen.go`                                → 新增 `Kind`
- `internal/model/define/table/tblbed/gentity_model_tbl_Bed.gen.go`                      → 新增 `BedTypeId`
- `internal/model/do/gentity_model_do_Bed.gen.go`                                        → 新增 `BedTypeId`

### 新增（room_material 全新表，全链路生成）
- `internal/model/define/table/tblroommaterial/gentity_model_tbl_RoomMaterial.gen.go`  （表字段：tenant_id, room_id, material_type_id, del_flag ...）
- `internal/model/do/gentity_model_do_RoomMaterial.gen.go`
- `internal/model/define/dao/gentity_model_dao_RoomMaterial.gen.go`

### 公共生成文件（若 gentity 一并刷新）
- `internal/model/dto/gentity_dto.gen.go`（序列化/校验，会因新增列而变化——**请重新生成**，勿手改）

## 三、生成后预期可用的符号（service 已引用）

- `dao.RoomMaterial(db)`、`tblroommaterial.RoomId / .MaterialTypeId / .TenantId / .DelFlag`
- `do.RoomMaterial`（含 RoomId / MaterialTypeId 等）
- `tblbed.BedTypeId`、`do.Bed.BedTypeId`
- `tblmaterialtype.Kind`、`do.MaterialType.Kind`

## 四、验证

生成并 `go build ./...` 通过后，接口即可工作：
- `/material/pageMaterialTypeByKey?kind=1` 取床型分类项；`?kind=99` 取设施分类项
- `/build/addRoom`、`/build/editRoom` 支持 `beds[].bed_type_id` 与 `facility_ids`
- `/build/getRoomById` 返回 `beds`（含床型名）与 `facilities`
- `/build/addFloor|deleteFloor` 后自动回写楼栋 `floor_num`
