package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProcessStepMatchRule(m *model.ProcessStepMatchRule) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProcessStepMatchRule(m *model.ProcessStepMatchRule) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.AttributeExpression{}, "rule_id=? AND rule_type=?", m.ID, "ProcessStepMatchRule").Error; err != nil {
			return err
		}

		if err := tx.Save(m).Error; err != nil {
			return err
		}

		return nil
	})
}

func QueryProcessStepMatchRule(req *proto.QueryProcessStepMatchRuleRequest, resp *proto.QueryProcessStepMatchRuleResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProcessStepMatchRule{})
	if req.ProductionLineID != "" {
		db = db.Where("`production_line_id`=?", req.ProductionLineID)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "id")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProcessStepMatchRule
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProcessStepMatchRulesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProcessStepMatchRules() (list []*model.ProcessStepMatchRule, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProcessStepMatchRuleByID(id string) (*model.ProcessStepMatchRule, error) {
	m := &model.ProcessStepMatchRule{}
	err := model.DB.DB().Preload("AttributeExpressions", func(db *gorm.DB) *gorm.DB {
		return db.Where("rule_type", "ProcessStepMatchRule")
	}).Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProcessStepMatchRuleByIDs(ids []string) ([]*model.ProcessStepMatchRule, error) {
	var m []*model.ProcessStepMatchRule
	err := model.DB.DB().Preload("AttributeExpressions", func(db *gorm.DB) *gorm.DB {
		return db.Where("rule_type", "ProcessStepMatchRule")
	}).Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProcessStepMatchRule(id string) (err error) {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.ProcessStepMatchRule{}, "id=?", id).Error; err != nil {
			return err
		}
		if err := tx.Delete(&model.AttributeExpression{}, "rule_id=? AND rule_type=?", id, "ProcessStepMatchRule").Error; err != nil {
			return err
		}
		return nil
	})
}
