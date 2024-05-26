package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProductPackageStackRule(m *model.ProductPackageStackRule) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductPackageStackRule(m *model.ProductPackageStackRule) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.AttributeExpression{}, "rule_id=? AND rule_type=?", m.ID, "ProductPackageStackRule").Error; err != nil {
			return err
		}

		if err := tx.Omit("created_at").Save(m).Error; err != nil {
			return err
		}

		return nil
	})
}

func QueryProductPackageStackRule(req *proto.QueryProductPackageStackRuleRequest, resp *proto.QueryProductPackageStackRuleResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductPackageStackRule{})

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductPackageStackRule
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductPackageStackRulesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductPackageStackRules() (list []*model.ProductPackageStackRule, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductPackageStackRuleByID(id string) (*model.ProductPackageStackRule, error) {
	m := &model.ProductPackageStackRule{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductPackageStackRule(productSerialNo string) (*model.ProductPackageStackRule, error) {
	m := &model.ProductPackageStackRule{}
	err := model.DB.DB().Preload(clause.Associations).Where("product_serial_no = ?", productSerialNo).First(m).Error
	return m, err
}

func GetProductPackageStackRuleByIDs(ids []string) ([]*model.ProductPackageStackRule, error) {
	var m []*model.ProductPackageStackRule
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductPackageStackRule(id string) (err error) {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.ProductPackageStackRule{}, "id=?", id).Error; err != nil {
			return err
		}
		if err := tx.Delete(&model.AttributeExpression{}, "rule_id=? AND rule_type=?", id, "ProductPackageStackRule").Error; err != nil {
			return err
		}
		return nil
	})
}

// func DeleteProductPackageStackRule(id string) (err error) {
// var ProductPackageStackRule model.ProductPackageStackRule
// if err := model.DB.DB().First(&model.ProductPackageStackRule, id).Error; err != nil {
// 	return nil
// }
// validStates := []string{"已取消", "已接单", "已上传", "已确认"}
// if !tool.Contains(validStates, model.ProductPackageStackRule.CurrentState) {
// 	return fmt.Errorf("只有工单状态为已取消、已接单、已上传或已确认时，才可以删除，请核对后再试")
// }
// deleteProductPackageStackRules := []interface{}{&ProductIssueRecord{}, &ProductPackageRecord{}, &ProductProcessRecord{}, &ProductProcessRoute{}, &ProductReworkRecord{}, &ProductRhythmRecord{}, &ProductTestRecord{}, &ProductWorkRecord{}, &ProductionStationAlarm{}, &ProductionStationOutput{}}

// return model.DB.DB().Transaction(func(tx *gorm.DB) error {
// 	for _, model := range deleteProductPackageStackRules {
// 		if err := tx.Delete(model, "ProductPackageStackRuleID=?", id).Error; err != nil {
// 			return err
// 		}
// 	}

// 	if err := tx.Delete(&model.ProductPackageStackRule{}, "id=?", id).Error; err != nil {
// 		return err
// 	}

// 	return nil
// })
// 	return nil
// }
