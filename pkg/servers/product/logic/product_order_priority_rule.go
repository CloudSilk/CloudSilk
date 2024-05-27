package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProductOrderPriorityRule(m *model.ProductOrderPriorityRule) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductOrderPriorityRule(m *model.ProductOrderPriorityRule) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.AttributeExpression{}, "`rule_id` = ? AND `rule_type` = ?", m.ID, "ProductOrderPriorityRule").Error; err != nil {
			return err
		}

		if err := tx.Omit("created_at").Save(m).Error; err != nil {
			return err
		}

		return nil
	})
}

func QueryProductOrderPriorityRule(req *proto.QueryProductOrderPriorityRuleRequest, resp *proto.QueryProductOrderPriorityRuleResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductOrderPriorityRule{})
	if req.PriorityLevel > 0 {
		db = db.Where("`priority_level` = ?", req.PriorityLevel)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductOrderPriorityRule
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductOrderPriorityRulesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductOrderPriorityRules() (list []*model.ProductOrderPriorityRule, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductOrderPriorityRuleByID(id string) (*model.ProductOrderPriorityRule, error) {
	m := &model.ProductOrderPriorityRule{}
	err := model.DB.DB().Preload("AttributeExpressions", func(db *gorm.DB) *gorm.DB {
		return db.Where("rule_type", "ProductOrderPriorityRule")
	}).Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProductOrderPriorityRuleByIDs(ids []string) ([]*model.ProductOrderPriorityRule, error) {
	var m []*model.ProductOrderPriorityRule
	err := model.DB.DB().Preload("AttributeExpressions", func(db *gorm.DB) *gorm.DB {
		return db.Where("rule_type", "ProductOrderPriorityRule")
	}).Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductOrderPriorityRule(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductOrderPriorityRule{}, "`id` = ?", id).Error
}
