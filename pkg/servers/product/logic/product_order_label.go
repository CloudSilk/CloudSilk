package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductOrderLabel(m *model.ProductOrderLabel) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductOrderLabel(m *model.ProductOrderLabel) error {
	return model.DB.DB().Omit("created_at", "create_time").Save(m).Error
}

func QueryProductOrderLabel(req *proto.QueryProductOrderLabelRequest, resp *proto.QueryProductOrderLabelResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductOrderLabel{}).Preload("ProductOrder").Preload("LabelType").Preload(clause.Associations)
	if req.ProductOrderNo != "" {
		db = db.Joins("JOIN product_orders ON model.product_order_labels.product_order_id = product_orders.id").
			Where("product_orders.product_order_no LIKE ?", "%"+req.ProductOrderNo+"%")
	}
	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("product_order_labels.create_time BETWEEN ? and ?", req.CreateTime0, req.CreateTime1)
	}
	if req.CurrentState != "" {
		db = db.Where("product_order_labels.current_state = ?", req.CurrentState)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductOrderLabel
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductOrderLabelsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductOrderLabels() (list []*model.ProductOrderLabel, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductOrderLabelByID(id string) (*model.ProductOrderLabel, error) {
	m := &model.ProductOrderLabel{}
	err := model.DB.DB().Preload("ProductOrderLabelParameters").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProductOrderLabelByIDs(ids []string) ([]*model.ProductOrderLabel, error) {
	var m []*model.ProductOrderLabel
	err := model.DB.DB().Preload("ProductOrderLabelParameters").Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductOrderLabel(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductOrderLabel{}, "`id` = ?", id).Error
}
