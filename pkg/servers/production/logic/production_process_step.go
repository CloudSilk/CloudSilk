package logic

import (
	"errors"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProductionProcessStep(m *model.ProductionProcessStep) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, " `code`  = ? ", m.Code)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同工步类型")
	}
	return m.ID, nil
}

func UpdateProductionProcessStep(m *model.ProductionProcessStep) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.AttributeExpression{}, "`rule_id` = ? AND `rule_type` = ?", m.ID, "ProductionProcessStep").Error; err != nil {
			return err
		}

		if err := tx.Delete(&model.AvailableProcess{}, "`production_process_step_id` = ?", m.ID).Error; err != nil {
			return err
		}

		duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(tx, m, true, []string{"created_at"}, "`id` <> ? and `code`  = ? ", m.ID, m.Code)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同工步类型")
		}

		return nil
	})
}

func QueryProductionProcessStep(req *proto.QueryProductionProcessStepRequest, resp *proto.QueryProductionProcessStepResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductionProcessStep{}).Preload("ProcessStepType").Preload("AttributeExpressions").Preload(clause.Associations)
	if req.ProductionProcessID != "" {
		db = db.Joins("JOIN available_processes ON production_process_steps.id=available_processes.production_process_step_id").
			Where("available_processes.production_process_id = ?", req.ProductionProcessID)
	}
	if len(req.Ids) > 0 {
		db = db.Where("`id` in ?", req.Ids)
	}
	if req.Enable {
		db = db.Where("`enable` = ?", req.Enable)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductionProcessStep
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionProcessStepsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductionProcessSteps() (list []*model.ProductionProcessStep, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductionProcessStepByID(id string) (*model.ProductionProcessStep, error) {
	m := &model.ProductionProcessStep{}
	err := model.DB.DB().Preload("AttributeExpressions", func(db *gorm.DB) *gorm.DB {
		return db.Where("rule_type", "ProductionProcessStep")
	}).Preload("AvailableProcesses").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProductionProcessStep(req *proto.GetProductionProcessStepRequest) (*model.ProductionProcessStep, error) {
	m := &model.ProductionProcessStep{}
	err := model.DB.DB().Preload("AttributeExpressions", func(db *gorm.DB) *gorm.DB {
		return db.Where("rule_type", "ProductionProcessStep")
	}).Preload("AvailableProcesses").Preload(clause.Associations).First(m, map[string]interface{}{
		"production_process_id": req.ProductionProcessID,
		"product_model_id":      req.ProductModelID,
		"code":                  req.Code,
	}).Error
	return m, err
}

func GetProductionProcessStepByIDs(ids []string) ([]*model.ProductionProcessStep, error) {
	var m []*model.ProductionProcessStep
	err := model.DB.DB().Preload("AttributeExpressions", func(db *gorm.DB) *gorm.DB {
		return db.Where("rule_type", "ProductionProcessStep")
	}).Preload("AvailableProcesses").Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductionProcessStep(id string) (err error) {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.ProductionProcessStep{}, "`id` = ?", id).Error; err != nil {
			return err
		}
		if err := tx.Delete(&model.AttributeExpression{}, "`rule_id` = ? AND `rule_type` = ?", id, "ProductionProcessStep").Error; err != nil {
			return err
		}
		return nil
	})
}
