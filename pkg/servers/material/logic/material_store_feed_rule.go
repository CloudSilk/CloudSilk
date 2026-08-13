package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateMaterialStoreFeedRule(m *model.MaterialStoreFeedRule) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateMaterialStoreFeedRule(m *model.MaterialStoreFeedRule) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Delete(&model.AttributeExpression{}, "`rule_id` = ? AND `rule_type` = ?", m.ID, "MaterialStoreFeedRule").Error; err != nil {
			return err
		}

		if err := tx.Omit("created_at").Save(m).Error; err != nil {
			return err
		}

		return nil
	})
}

func QueryMaterialStoreFeedRule(req *proto.QueryMaterialStoreFeedRuleRequest, resp *proto.QueryMaterialStoreFeedRuleResponse, preload bool) {
	db := model.DB.DB().Model(&model.MaterialStoreFeedRule{}).Preload("MaterialStore").Preload("MaterialInfo")
	if req.MaterialInfo != "" {
		db = db.Joins("JOIN material_infos ON material_store_feed_rules.material_info_id=material_infos.id").
			Where("material_infos.material_no LIKE ? OR material_infos.material_description LIKE ?", "%"+req.MaterialInfo+"%", "%"+req.MaterialInfo+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MaterialStoreFeedRule
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialStoreFeedRulesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMaterialStoreFeedRules() (list []*model.MaterialStoreFeedRule, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMaterialStoreFeedRuleByID(id string) (*model.MaterialStoreFeedRule, error) {
	m := &model.MaterialStoreFeedRule{}
	err := model.DB.DB().Preload("MaterialStore").Preload("MaterialInfo").Preload("AttributeExpressions").Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetMaterialStoreFeedRuleByIDs(ids []string) ([]*model.MaterialStoreFeedRule, error) {
	var m []*model.MaterialStoreFeedRule
	err := model.DB.DB().Preload("MaterialStore").Preload("MaterialInfo").Preload("AttributeExpressions").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMaterialStoreFeedRule(id string) (err error) {
	return model.DB.DB().Delete(&model.MaterialStoreFeedRule{}, "`id` = ?", id).Error
}
