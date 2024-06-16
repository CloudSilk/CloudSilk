package logic

import (
	"errors"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductInfo(m *model.ProductInfo) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, "`product_serial_no` = ?", m.ProductSerialNo)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同产品信息配置")
	}
	return m.ID, nil
}

func UpdateProductInfo(m *model.ProductInfo) error {
	duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(model.DB.DB(), m, false, []string{"created_at", "create_time"}, "`id` <> ? and  `product_serial_no` = ?", m.ID, m.ProductSerialNo)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同产品信息配置")
	}

	return nil
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

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
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
	err := model.DB.DB().Preload("ProductOrder").
		Preload("ProductOrder.ProductModel").
		Preload("ProductOrder.ProductModel.ProductCategory").
		Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProductInfo(productSerialNo string) (*model.ProductInfo, error) {
	m := &model.ProductInfo{}
	err := model.DB.DB().
		Preload("ProductOrder").
		Preload("ProductOrder.ProductModel").
		Preload("ProductOrder.ProductModel.ProductCategory").
		Preload(clause.Associations).
		Where("`product_serial_no` = ?", productSerialNo).First(m).Error
	return m, err
}

func GetProductInfoByIDs(ids []string) ([]*model.ProductInfo, error) {
	var m []*model.ProductInfo
	err := model.DB.DB().
		Preload("ProductOrder").
		Preload("ProductOrder.ProductModel").
		Preload("ProductOrder.ProductModel.ProductCategory").
		Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductInfo(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductInfo{}, "`id` = ?", id).Error
}
