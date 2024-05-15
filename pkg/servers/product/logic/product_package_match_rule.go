package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProductPackageMatchRule(m *model.ProductPackageMatchRule) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductPackageMatchRule(m *model.ProductPackageMatchRule) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.AttributeExpression{}, "rule_id=? AND rule_type=?", m.ID, "ProductPackageMatchRule").Error; err != nil {
			return err
		}

		if err := tx.Save(m).Error; err != nil {
			return err
		}

		return nil
	})
}

func QueryProductPackageMatchRule(req *proto.QueryProductPackageMatchRuleRequest, resp *proto.QueryProductPackageMatchRuleResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductPackageMatchRule{}).Preload("ProductPackage").Preload(clause.Associations)

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "id")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductPackageMatchRule
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductPackageMatchRulesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductPackageMatchRules() (list []*model.ProductPackageMatchRule, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductPackageMatchRuleByID(id string) (*model.ProductPackageMatchRule, error) {
	m := &model.ProductPackageMatchRule{}
	err := model.DB.DB().Preload("AttributeExpressions", func(db *gorm.DB) *gorm.DB {
		return db.Where("rule_type", "ProductPackageMatchRule")
	}).Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductPackageMatchRuleByIDs(ids []string) ([]*model.ProductPackageMatchRule, error) {
	var m []*model.ProductPackageMatchRule
	err := model.DB.DB().Preload("AttributeExpressions", func(db *gorm.DB) *gorm.DB {
		return db.Where("rule_type", "ProductPackageMatchRule")
	}).Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductPackageMatchRule(id string) (err error) {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.ProductPackageMatchRule{}, "id=?", id).Error; err != nil {
			return err
		}
		if err := tx.Delete(&model.AttributeExpression{}, "rule_id=? AND rule_type=?", id, "ProductPackageMatchRule").Error; err != nil {
			return err
		}
		return nil
	})
}
