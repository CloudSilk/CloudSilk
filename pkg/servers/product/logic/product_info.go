package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductInfo(m *model.ProductInfo) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductInfo(m *model.ProductInfo) error {
	return model.DB.DB().Omit("create_time").Save(m).Error
}

func QueryProductInfo(req *proto.QueryProductInfoRequest, resp *proto.QueryProductInfoResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductInfo{}).Preload("ProductOrder").Preload("ProductOrder.ProductModel").Preload(clause.Associations)
	if req.ProductOrderNo != "" {
		db = db.Joins("JOIN product_orders ON product_infos.product_order_id = product_order.id").
			Where("product_orders.product_order_no LIKE ?", "%"+req.ProductOrderNo+"%")
	}
	if req.ProductSerialNo != "" {
		db = db.Where("`product_serial_no` LIKE ?", "%"+req.ProductSerialNo+"%")
	}
	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("product_infos.create_time BETWEEN ? AND ?", req.CreateTime0, req.CreateTime1)
	}
	if req.CurrentState != "" {
		db = db.Where("`current_state` = ?", req.CurrentState)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "id")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductInfo
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductInfosToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductInfos() (list []*model.ProductInfo, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductInfoByID(id string) (*model.ProductInfo, error) {
	m := &model.ProductInfo{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductInfo(productSerialNo string) (*model.ProductInfo, error) {
	m := &model.ProductInfo{}
	err := model.DB.DB().Preload(clause.Associations).Where("product_serial_no = ?", productSerialNo).First(m).Error
	return m, err
}

func GetProductInfoByIDs(ids []string) ([]*model.ProductInfo, error) {
	var m []*model.ProductInfo
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductInfo(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductInfo{}, "id=?", id).Error
}

// func DeleteProductInfo(id string) (err error) {
// var ProductInfo model.ProductInfo
// if err := model.DB.DB().First(&model.ProductInfo, id).Error; err != nil {
// 	return nil
// }
// validStates := []string{"已取消", "已接单", "已上传", "已确认"}
// if !tool.Contains(validStates, model.ProductInfo.CurrentState) {
// 	return fmt.Errorf("只有工单状态为已取消、已接单、已上传或已确认时，才可以删除，请核对后再试")
// }
// deleteProductInfos := []interface{}{&ProductIssueRecord{}, &ProductPackageRecord{}, &ProductProcessRecord{}, &ProductProcessRoute{}, &ProductReworkRecord{}, &ProductRhythmRecord{}, &ProductTestRecord{}, &ProductWorkRecord{}, &ProductionStationAlarm{}, &ProductionStationOutput{}}

// return model.DB.DB().Transaction(func(tx *gorm.DB) error {
// 	for _, model := range deleteProductInfos {
// 		if err := tx.Delete(model, "ProductInfoID=?", id).Error; err != nil {
// 			return err
// 		}
// 	}

// 	if err := tx.Delete(&model.ProductInfo{}, "id=?", id).Error; err != nil {
// 		return err
// 	}

// 	return nil
// })
// 	return nil
// }
