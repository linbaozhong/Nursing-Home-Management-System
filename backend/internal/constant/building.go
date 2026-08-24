package constant

import "github.com/linbaozhong/gentity/pkg/types"

// 楼栋
var (
	ErrBuildingRepeat = types.NewError(500, "楼栋已存在")
	ErrBuildingOut    = types.NewError(500, "楼栋总数超过限制")
)

// 楼层
var (
	ErrFloorRepeat = types.NewError(500, "楼层已存在")
	ErrFloorOut    = types.NewError(500, "楼层总数超过限制")
)

// 房间类型
var (
	ErrRoomTypeRepeat = types.NewError(500, "房间类型已存在")
)

// 房间
var (
	ErrRoomRepeat = types.NewError(500, "房间已存在")
	ErrRoomOut    = types.NewError(500, "房间总数超过限制")
)

// MaterialTypeKind 物资分类用途
type MaterialTypeKind int8

// 物资分类用途常量
const (
	KindBed      MaterialTypeKind = 1  // 床型
	KindFacility MaterialTypeKind = 99 // 设施/其他
)

// BedStatus 床位状态
type BedStatus int8

// 床位状态
const (
	BedIdle      BedStatus = 0 // 空闲
	BedReserve   BedStatus = 1 // 预定
	BedEnter     BedStatus = 2 // 入住
	BedExitAudit BedStatus = 3 // 退住审核
)

func (b BedStatus) String() string {
	switch b {
	case BedIdle:
		return "空闲"
	case BedReserve:
		return "预定"
	case BedEnter:
		return "入住"
	case BedExitAudit:
		return "退住审核"
	default:
		return "未知"
	}
}

// 床位
var (
	ErrBedNull    = types.NewError(500, "床位不存在")
	ErrBedRepeat  = types.NewError(500, "床位已存在")
	ErrBedOut     = types.NewError(500, "床位总数超过限制")
	ErrBedOccupy  = types.NewError(500, "该床位被占用")
	ErrBedNotIdle = types.NewError(500, "该床位被占用，删除失败")
)

// NodeMark 节点标记
type NodeMark int8

// 节点标记
const (
	MarkBuilding NodeMark = 1 // 楼栋
	MarkFloor    NodeMark = 2 // 楼层
	MarkRoom     NodeMark = 3 // 房间
)

func (n NodeMark) String() string {
	switch n {
	case MarkBuilding:
		return "楼栋"
	case MarkFloor:
		return "楼层"
	case MarkRoom:
		return "房间"
	default:
		return "未知"
	}
}

// 节点
var (
	ErrNodeMarkNotExist = types.NewError(500, "该节点标记不存在")
	ErrNodeBedNotIdle   = types.NewError(500, "该节点有床位被占用，删除失败")
)
