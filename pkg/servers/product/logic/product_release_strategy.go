package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProductReleaseStrategy(m *model.ProductReleaseStrategy) (string, error) {
	err := model.DB.DB().Model(m).Create(m).Error
	return m.ID, err
}

func UpdateProductReleaseStrategy(m *model.ProductReleaseStrategy) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.ProductReleaseStrategyComparableAttribute{}, "product_release_strategy_id=?", m.ID).Error; err != nil {
			return err
		}

		return tx.Omit("created_at").Save(m).Error
	})
}

func QueryProductReleaseStrategy(req *proto.QueryProductReleaseStrategyRequest, resp *proto.QueryProductReleaseStrategyResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductReleaseStrategy{}).Preload("ProductCategory").Preload(clause.Associations)
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}
	if req.ProductionLineID != "" {
		db = db.Where("`production_line_id` = ?", req.ProductionLineID)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductReleaseStrategy
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductReleaseStrategysToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductReleaseStrategys() (list []*model.ProductReleaseStrategy, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductReleaseStrategyByID(id string) (*model.ProductReleaseStrategy, error) {
	m := &model.ProductReleaseStrategy{}
	err := model.DB.DB().Preload("ProductReleaseStrategyComparableAttributes").Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductReleaseStrategyByIDs(ids []string) ([]*model.ProductReleaseStrategy, error) {
	var m []*model.ProductReleaseStrategy
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductReleaseStrategy(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductReleaseStrategy{}, "id=?", id).Error
}
