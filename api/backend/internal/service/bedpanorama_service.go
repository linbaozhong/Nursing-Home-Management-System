package service

import (
	"context"

	"api/internal/model/dto"
)

type bedpanorama struct{}

var BedPanorama = &bedpanorama{}

// GetBedPanorama 获取床位全景（首页看板）
// 对应 Java: BedPanoramaServiceImpl.getBedPanorama -> 多表聚合查询
// 返回 BedPanoramaVO, 含:
//   - 楼宇列表(building) 及其每层楼的房间(room)与床位(bed)信息(房间状态、入住老人 elder 姓名/头像)
//   - 当前楼宇的空床位统计
//   - 异常提醒(accident 待处理数、reserve 待跟进数、retreat_apply 退住申请数等)
//
// todo: 1) 查 building 列表; 2) 按楼宇查 floor/room/bed, 联表 elder 获取入住老人;
//
//	3) 统计各类异常数量(accident/reserve/retreat_apply 等未处理数);
//	4) 组装 BedPanoramaVO(需定义对应返回类型)并赋值 out
func (b *bedpanorama) GetBedPanorama(ctx context.Context, in *dto.EmptyReq, out *dto.EmptyResp) error {
	// todo: 见上方方法注释, 实现多表聚合看板查询
	return nil
}
