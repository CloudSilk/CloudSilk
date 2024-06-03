package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProductAttributeValuateRule(m *model.ProductAttributeValuateRule) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductAttributeValuateRule(m *model.ProductAttributeValuateRule) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Delete(&model.PropertyExpression{}, "`rule_id` = ? and `rule_type` = ?", m.ID, "ProductAttributeValuateRule").Error; err != nil {
			return err
		}

		if err := tx.Omit("created_at").Save(m).Error; err != nil {
			return err
		}

		return nil
	})
}

func QueryProductAttributeValuateRule(req *proto.QueryProductAttributeValuateRuleRequest, resp *proto.QueryProductAttributeValuateRuleResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductAttributeValuateRule{}).Preload("ProductCategory").Preload("ProductAttribute")
	if req.ProductCategoryID != "" {
		db = db.Where("`product_category_id` = ?", req.ProductCategoryID)
	}

	if req.ProductAttributeID != "" {
		db = db.Where("`product_attribute_id` = ?", req.ProductAttributeID)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductAttributeValuateRule
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductAttributeValuateRulesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductAttributeValuateRules() (list []*model.ProductAttributeValuateRule, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductAttributeValuateRuleByID(id string) (*model.ProductAttributeValuateRule, error) {
	m := &model.ProductAttributeValuateRule{}
	err := model.DB.DB().Preload("PropertyExpressions").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProductAttributeValuateRuleByIDs(ids []string) ([]*model.ProductAttributeValuateRule, error) {
	var m []*model.ProductAttributeValuateRule
	err := model.DB.DB().Preload("PropertyExpressions").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductAttributeValuateRule(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductAttributeValuateRule{}, "`id` = ?", id).Error
}
