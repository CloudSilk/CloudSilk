package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductOrderBom(m *model.ProductOrderBom) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductOrderBom(m *model.ProductOrderBom) error {
	return model.DB.DB().Omit("created_at", "create_time").Save(m).Error
}

func QueryProductOrderBom(req *proto.QueryProductOrderBomRequest, resp *proto.QueryProductOrderBomResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductOrderBom{}).Preload("ProductOrder").Preload(clause.Associations)
	if req.ProductOrderNo != "" {
		db = db.Joins("JOIN product_orders ON product_order_boms.product_order_id = product_orders.id").
			Where("product_orders.product_order_no LIKE ?", "%"+req.ProductOrderNo+"%")
	}
	if req.MaterialNo != "" {
		db = db.Where("`material_no` LIKE ? OR `material_description` LIKE ?", "%"+req.MaterialNo+"%", "%"+req.MaterialNo+"%")
	}
	if req.ProductOrderID != "" {
		db = db.Where("`product_order_id` = ?", req.ProductOrderID)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductOrderBom
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductOrderBomsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductOrderBoms() (list []*model.ProductOrderBom, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductOrderBomByID(id string) (*model.ProductOrderBom, error) {
	m := &model.ProductOrderBom{}
	err := model.DB.DB().Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProductOrderBomByIDs(ids []string) ([]*model.ProductOrderBom, error) {
	var m []*model.ProductOrderBom
	err := model.DB.DB().Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductOrderBom(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductOrderBom{}, "`id` = ?", id).Error
}
