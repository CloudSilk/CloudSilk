package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProductOrderReleaseRule(m *model.ProductOrderReleaseRule) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductOrderReleaseRule(m *model.ProductOrderReleaseRule) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Delete(&model.AttributeExpression{}, "`rule_id` = ? AND `rule_type` = ?", m.ID, "ProductOrderReleaseRule").Error; err != nil {
			return err
		}

		if err := tx.Save(m).Error; err != nil {
			return err
		}

		return nil
	})
}

func QueryProductOrderReleaseRule(req *proto.QueryProductOrderReleaseRuleRequest, resp *proto.QueryProductOrderReleaseRuleResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductOrderReleaseRule{}).Preload("ProductionLine").Preload(clause.Associations)

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductOrderReleaseRule
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductOrderReleaseRulesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductOrderReleaseRules() (list []*model.ProductOrderReleaseRule, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductOrderReleaseRuleByID(id string) (*model.ProductOrderReleaseRule, error) {
	m := &model.ProductOrderReleaseRule{}
	err := model.DB.DB().Preload("AttributeExpressions", func(db *gorm.DB) *gorm.DB {
		return db.Where("rule_type", "ProductOrderReleaseRule")
	}).Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductOrderReleaseRuleByIDs(ids []string) ([]*model.ProductOrderReleaseRule, error) {
	var m []*model.ProductOrderReleaseRule
	err := model.DB.DB().Preload("AttributeExpressions", func(db *gorm.DB) *gorm.DB {
		return db.Where("rule_type", "ProductOrderReleaseRule")
	}).Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductOrderReleaseRule(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductOrderReleaseRule{}, "id=?", id).Error
}
