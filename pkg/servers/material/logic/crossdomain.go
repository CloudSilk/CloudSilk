package logic

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/model"
)

// 本文件是 material 域对其他业务域暴露的跨域访问出口（规则见 ARCHITECTURE.md）。

// LatestCompletedPickTime 工单最近一次物料齐套时间（最近一张已完成拣货单的完成时间），
// 供 APS 排程的物料齐套约束使用
func LatestCompletedPickTime(orderID string) (time.Time, bool) {
	bill := &model.WMSBillQueue{}
	err := model.DB.DB().Where("`product_order_id` = ? AND `current_state` = ?", orderID, WMSBillStateCompleted).
		Order("last_update_time desc").First(bill).Error
	if err != nil {
		return time.Time{}, false
	}
	return bill.LastUpdateTime, true
}
