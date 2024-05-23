package logic

import (
	"errors"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProcessStepParameter(m *model.ProcessStepParameter) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, " code =? ", m.Code)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同生产工步参数")
	}
	return m.ID, nil
}

func UpdateProcessStepParameter(m *model.ProcessStepParameter) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.ProcessStepParameterValue{}, "process_step_parameter_id=?", m.ID).Error; err != nil {
			return err
		}

		if err := tx.Delete(&model.AttributeExpression{}, "rule_id = ? AND rule_type = ?", m.ID, "ProcessStepParameter").Error; err != nil {
			return err
		}

		duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(tx, m, true, []string{}, "id <> ?  and  code =? ", m.ID, m.Code)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同生产工步参数")
		}

		return nil
	})
}

func QueryProcessStepParameter(req *proto.QueryProcessStepParameterRequest, resp *proto.QueryProcessStepParameterResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProcessStepParameter{})

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "id")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProcessStepParameter
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProcessStepParametersToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProcessStepParameters() (list []*model.ProcessStepParameter, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProcessStepParameterByID(id string) (*model.ProcessStepParameter, error) {
	m := &model.ProcessStepParameter{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProcessStepParameterByIDs(ids []string) ([]*model.ProcessStepParameter, error) {
	var m []*model.ProcessStepParameter
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProcessStepParameter(id string) (err error) {
	return model.DB.DB().Delete(&model.ProcessStepParameter{}, "id=?", id).Error
}

func GetAllProcessStepParameterByProductionLineID(productionLineID string) (*model.ProductionLine, error) {
	m := &model.ProductionLine{}
	err := model.DB.DB().
		Preload("ProductionProcesses").
		Preload("ProductionProcesses.ProductionProcessSteps").
		Preload("ProductionProcesses.ProductionProcessSteps.ProductionProcessStep").
		Preload("ProductionProcesses.ProductionProcessSteps.ProductionProcessStep.ProcessStepType").
		Preload("ProductionProcesses.ProductionProcessSteps.ProductionProcessStep.ProcessStepType.ProcessStepTypeParameters").
		Preload("ProcessStepParameters").
		Preload("ProcessStepParameters.ProcessStepParameterValues").
		Where("id = ?", productionLineID).First(m).Error

	for _, processStepParameter := range m.ProcessStepParameters {
		for _, processStepParameterValue := range processStepParameter.ProcessStepParameterValues {
			for _, process := range m.ProductionProcesses {
				if process.ID == processStepParameterValue.ProductionProcessID {
					for _, processStep := range process.ProductionProcessSteps {
						if processStep.ProductionProcessStepID == processStepParameterValue.ProductionProcessStepID {
							for _, processStepTypeParameter := range processStep.ProductionProcessStep.ProcessStepType.ProcessStepTypeParameters {
								if processStepTypeParameter.ID == processStepParameterValue.ProcessStepTypeParameterID {
									processStepTypeParameter.StandardValue = processStepParameterValue.StandardValue
									processStepTypeParameter.MinimumValue = processStepParameterValue.MinimumValue
									processStepTypeParameter.MaximumValue = processStepParameterValue.MaximumValue
								}
							}
						}
					}
				}
			}
		}
	}
	return m, err
}
