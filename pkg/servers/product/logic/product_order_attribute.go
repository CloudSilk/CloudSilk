package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductOrderAttribute(m *model.ProductOrderAttribute) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductOrderAttribute(m *model.ProductOrderAttribute) error {
	return model.DB.DB().Omit("created_at", "create_time").Save(m).Error
}

func QueryProductOrderAttribute(req *proto.QueryProductOrderAttributeRequest, resp *proto.QueryProductOrderAttributeResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductOrderAttribute{}).Preload("ProductAttribute").Preload("ProductOrder").Preload(clause.Associations)
	if req.ProductOrderNo != "" {
		db = db.Joins("JOIN product_orders ON product_order_attributes.product_order_id=product_orders.id").
			Where("product_orders.product_order_no LIKE ?", "%"+req.ProductOrderNo+"%")
	}
	if req.Code != "" {
		db = db.Joins("JOIN product_attributes ON product_order_attributes.product_attribute_id=product_attributes.id").
			Where("product_attributes.code LIKE ? OR product_attributes.description LIKE ? OR product_order_attributes.value LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%", "%"+req.Code+"%")
	}
	if req.ProductOrderID != "" {
		db = db.Where("`product_order_id`=?", req.ProductOrderID)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "`id`")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductOrderAttribute
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductOrderAttributesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductOrderAttributes() (list []*model.ProductOrderAttribute, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductOrderAttributeByID(id string) (*model.ProductOrderAttribute, error) {
	m := &model.ProductOrderAttribute{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductOrderAttributeByIDs(ids []string) ([]*model.ProductOrderAttribute, error) {
	var m []*model.ProductOrderAttribute
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductOrderAttribute(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductOrderAttribute{}, "id=?", id).Error
}
