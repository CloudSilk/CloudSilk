package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateLabelAdaptationRule(m *model.LabelAdaptationRule) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateLabelAdaptationRule(m *model.LabelAdaptationRule) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.AttributeExpression{}, "`rule_id` = ? AND `rule_type` = ?", m.ID, "LabelAdaptationRule").Error; err != nil {
			return err
		}

		if err := tx.Omit("created_at").Save(m).Error; err != nil {
			return err
		}

		return nil
	})
}

func QueryLabelAdaptationRule(req *proto.QueryLabelAdaptationRuleRequest, resp *proto.QueryLabelAdaptationRuleResponse, preload bool) {
	db := model.DB.DB().Model(&model.LabelAdaptationRule{}).Preload("LabelTemplate").Preload(clause.Associations)

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.LabelAdaptationRule
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.LabelAdaptationRulesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllLabelAdaptationRules() (list []*model.LabelAdaptationRule, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetLabelAdaptationRuleByID(id string) (*model.LabelAdaptationRule, error) {
	m := &model.LabelAdaptationRule{}
	err := model.DB.DB().Preload("AttributeExpressions", func(db *gorm.DB) *gorm.DB {
		return db.Where("rule_type", "LabelAdaptationRule")
	}).Preload("LabelTemplate").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetLabelAdaptationRuleByIDs(ids []string) ([]*model.LabelAdaptationRule, error) {
	var m []*model.LabelAdaptationRule
	err := model.DB.DB().Preload("AttributeExpressions", func(db *gorm.DB) *gorm.DB {
		return db.Where("rule_type", "LabelAdaptationRule")
	}).Preload("LabelTemplate").Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteLabelAdaptationRule(id string) (err error) {
	return model.DB.DB().Delete(&model.LabelAdaptationRule{}, "`id` = ?", id).Error
}
