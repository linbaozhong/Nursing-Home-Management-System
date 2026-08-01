package service

import (
	"context"
	"time"

	"api/internal/model/define/dao"
	"api/internal/model/define/table/tblcontract"
	"api/internal/model/define/table/tblemergencycontact"
	"api/internal/model/dto"
	"github.com/linbaozhong/gentity/pkg/ace"
)

type contract struct{}

var Contract = &contract{}

// ContractExpireJob 合同到期提醒定时任务
// 对应 Java: ContractServiceImpl.contractExpireJob
// 逻辑：
//  1. 查出 7 天内即将到期的合同(contract 表 end_date 在 [now, now+7天])
//  2. 遍历合同，按 elder_id 查紧急联系人(emergency_contact)邮箱
//  3. 向每个紧急联系人邮箱发送合同到期提醒邮件
//
// 说明：Go 侧暂无 SMTP/邮件发送工具，发邮件部分以 //todo 标注。
func (c *contract) ContractExpireJob(ctx context.Context, in *dto.EmptyReq, out *dto.EmptyResp) error {
	now := time.Now()
	start := parseTime(now.Format("2006-01-02 15:04:05"))
	end := parseTime(now.AddDate(0, 0, 7).Format("2006-01-02 15:04:05"))
	// 1) 查 7 天内即将到期的合同
	contracts, _, e := dao.Contract(db).List(ctx,
		ace.Where(tblcontract.EndDate.Gte(start)),
		ace.Where(tblcontract.EndDate.Lte(end)),
	)
	if e != nil {
		return e
	}
	for _, ct := range contracts {
		// 2) 查该合同老人的紧急联系人邮箱
		contacts, _, e := dao.EmergencyContact(db).List(ctx,
			ace.Where(tblemergencycontact.ElderId.Eq(ct.ElderId)),
		)
		if e != nil {
			return e
		}
		// 3) 发送合同到期提醒邮件
		for _, ec := range contacts {
			// todo: 调用邮件发送工具, 向 ec.Email 发送合同(ct.EndDate)到期提醒
			_ = ec
		}
	}
	return nil
}
