package logic

import (
	"database/sql"
	"errors"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/types"
	"gorm.io/gorm"
)

// 本文件是 product 域对其他业务域暴露的跨域访问出口。
// 规则见 ARCHITECTURE.md：其他模块不得直接读写 product 域的数据表，
// 一律经由这里导出的函数（保留在事务内执行以组合跨域流程）。

// GetProductOrderForPick 读取工单（含BOM），供 WMS 生成拣货单使用
func GetProductOrderForPick(tx *gorm.DB, id string) (*model.ProductOrder, error) {
	order := &model.ProductOrder{}
	if err := tx.Preload("ProductOrderBoms").First(order, "`id` = ?", id).Error; err != nil {
		return nil, errors.New("读取生产工单失败")
	}
	return order, nil
}

// GetProductInfoByIDTx 事务内读取产品信息（质量/物料域定位产品用）；
// 与既有 GetProductInfoByID（非事务、带聚合预加载）互补
func GetProductInfoByIDTx(tx *gorm.DB, id string) (*model.ProductInfo, error) {
	pi := &model.ProductInfo{}
	if err := tx.First(pi, "`id` = ?", id).Error; err != nil {
		return nil, errors.New("读取产品信息失败")
	}
	return pi, nil
}

// AddProductOrderIssued 累加工单发料数量并记录发料时间（WMS 完成拣货联动）
func AddProductOrderIssued(tx *gorm.DB, orderID string, qty int64, at time.Time) error {
	return tx.Model(&model.ProductOrder{}).Where("`id` = ?", orderID).
		Updates(map[string]interface{}{
			"issued_qty":      gorm.Expr("issued_qty + ?", qty),
			"last_issue_time": at,
		}).Error
}

// MarkProductChecking 将产品状态置为检查中（质量判定不合格联动）
func MarkProductChecking(tx *gorm.DB, productID string) error {
	return tx.Model(&model.ProductInfo{}).Where("`id` = ?", productID).
		Update("current_state", types.ProductStateChecking).Error
}

// CreateProductReworkRecordTx 事务内创建返工记录（质量不合格联动返工）
func CreateProductReworkRecordTx(tx *gorm.DB, m *model.ProductReworkRecord) error {
	return tx.Create(m).Error
}

// GetSchedulableOrders 读取产线下可排程工单（已发放/已签派/生产中），
// 供 APS 生成排程使用；ids 非空时限定工单范围
func GetSchedulableOrders(lineID string, ids []string) ([]*model.ProductOrder, error) {
	query := model.DB.DB().Preload("ProductModel").
		Where("`production_line_id` = ? AND `current_state` in ?", lineID,
			[]string{types.ProductOrderStateReleased, types.ProductOrderStateDispatched, types.ProductOrderStateProducting})
	if len(ids) > 0 {
		query = query.Where("`id` in (?)", ids)
	}
	var orders []*model.ProductOrder
	if err := query.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

// NullTime 辅助：跨域返回时间用（避免各域自行引 database/sql 拼装）
func NullTime(t time.Time) sql.NullTime {
	return sql.NullTime{Time: t, Valid: true}
}
