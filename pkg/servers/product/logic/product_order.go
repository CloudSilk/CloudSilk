package logic

import (
	"errors"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProductOrder(m *model.ProductOrder) (string, error) {
	// systemConfigKey := systemConfigKeys.PrefabricateProductOrderPrefix
	// if m.OrderType == productOrderTypes.Routine {
	// 	systemConfigKey = systemConfigKeys.RoutineProductOrderPrefix
	// }

	// model.DB.DB().Find(&model.SystemConfigs{}, "", "")

	duplication, err := model.DB.CreateWithCheckDuplication(m, "receipt_note_no=?", m.ReceiptNoteNo)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同产品工单")
	}
	return m.ID, nil
}

func UpdateProductOrder(m *model.ProductOrder) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.ProductOrderBom{}, "`product_order_id` = ?", m.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(&model.ProductInfo{}, "`product_order_id` = ?", m.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(&model.ProductOrderAttribute{}, "`product_order_id` = ?", m.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(&model.ProductOrderProcess{}, "`product_order_id` = ?", m.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(&model.ProductOrderPackage{}, "`product_order_id` = ?", m.ID).Error; err != nil {
			return err
		}

		duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(tx, m, false, []string{"create_time"}, "id <> ? and  receipt_note_no=?", m.ID, m.ReceiptNoteNo)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同产品工单")
		}

		return nil
	})
}

func QueryProductOrder(req *proto.QueryProductOrderRequest, resp *proto.QueryProductOrderResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductOrder{})
	if req.ProductOrderNo != "" {
		db = db.Where("`product_order_no` LIKE ?", "%"+req.ProductOrderNo+"%")
	}

	if req.SalesOrderNo != "" {
		db = db.Where("`sales_order_no` LIKE ?", "%"+req.SalesOrderNo+"%")
	}

	if req.ItemNo != "" {
		db = db.Where("`item_no` LIKE ?", "%"+req.ItemNo+"%")
	}

	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("`create_time` BETWEEN ? and ?", req.CreateTime0, req.CreateTime1)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "id")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductOrder
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductOrdersToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductOrders() (list []*model.ProductOrder, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductOrderByID(id string) (*model.ProductOrder, error) {
	m := &model.ProductOrder{}
	err := model.DB.DB().
		Preload("ProductOrderAttachments").
		Preload("ProductOrderBoms").
		Preload("ProductOrderAttributes").
		Preload("ProductInfos").
		Preload("ProductOrderProcesses").
		Preload("ProductOrderLabels").
		Preload("ProductOrderPackages").
		Preload("ProductOrderPallets").
		Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductOrderByIDs(ids []string) ([]*model.ProductOrder, error) {
	var m []*model.ProductOrder
	err := model.DB.DB().
		Preload("ProductOrderAttachments").
		Preload("ProductOrderBoms").
		Preload("ProductOrderAttributes").
		Preload("ProductInfos").
		Preload("ProductOrderProcesses").
		Preload("ProductOrderLabels").
		Preload("ProductOrderPackages").
		Preload("ProductOrderPallets").
		Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductOrder(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductOrder{}, "id=?", id).Error
}
