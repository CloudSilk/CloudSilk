package logic

import (
	"errors"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProcessStepType(m *model.ProcessStepType) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, " `code`  = ? ", m.Code)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同工步类型配置")
	}
	return m.ID, nil
}

func UpdateProcessStepType(m *model.ProcessStepType) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		oldProcessStepType := &model.ProcessStepType{}
		if err := tx.First(oldProcessStepType, "`id` = ?", m.ID).Error; err != nil {
			return err
		}

		duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(tx, m, true, []string{"created_at"}, "`id` <> ?  and  `code`  = ? ", m.ID, m.Code)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同工步类型配置")
		}

		return nil
	})
}

func QueryProcessStepType(req *proto.QueryProcessStepTypeRequest, resp *proto.QueryProcessStepTypeResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProcessStepType{})
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProcessStepType
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProcessStepTypesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProcessStepTypes() (list []*model.ProcessStepType, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProcessStepTypeByID(id string) (*model.ProcessStepType, error) {
	m := &model.ProcessStepType{}
	err := model.DB.DB().Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProcessStepTypeByIDs(ids []string) ([]*model.ProcessStepType, error) {
	var m []*model.ProcessStepType
	err := model.DB.DB().Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProcessStepType(id string) (err error) {
	return model.DB.DB().Delete(&model.ProcessStepType{}, "`id` = ?", id).Error
}
