package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateMaterialReturnCause(m *model.MaterialReturnCause) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateMaterialReturnCause(m *model.MaterialReturnCause) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Delete(&model.MaterialReturnCauseAvailableCategory{}, "`material_return_cause_id` = ?", m.ID).Error; err != nil {
			return err
		}
		if err := tx.Unscoped().Delete(&model.MaterialReturnCauseAvailableType{}, "`material_return_cause_id` = ?", m.ID).Error; err != nil {
			return err
		}

		if err := tx.Omit("created_at").Save(m).Error; err != nil {
			return err
		}

		return nil
	})
}

func QueryMaterialReturnCause(req *proto.QueryMaterialReturnCauseRequest, resp *proto.QueryMaterialReturnCauseResponse, preload bool) {
	db := model.DB.DB().Model(&model.MaterialReturnCause{})
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MaterialReturnCause
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialReturnCausesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMaterialReturnCauses() (list []*model.MaterialReturnCause, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMaterialReturnCauseByID(id string) (*model.MaterialReturnCause, error) {
	m := &model.MaterialReturnCause{}
	err := model.DB.DB().
		Preload("MaterialCategories").Preload("MaterialCategories.MaterialCategory").
		Preload("MaterialReturnTypes").Preload("MaterialReturnTypes.MaterialReturnType").
		Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetMaterialReturnCauseByIDs(ids []string) ([]*model.MaterialReturnCause, error) {
	var m []*model.MaterialReturnCause
	err := model.DB.DB().
		Preload("MaterialCategories").Preload("MaterialCategories.MaterialCategory").
		Preload("MaterialReturnTypes").Preload("MaterialReturnTypes.MaterialReturnType").
		Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMaterialReturnCause(id string) (err error) {
	return model.DB.DB().Delete(&model.MaterialReturnCause{}, "`id` = ?", id).Error
}
