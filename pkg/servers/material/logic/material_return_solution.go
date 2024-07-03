package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateMaterialReturnSolution(m *model.MaterialReturnSolution) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateMaterialReturnSolution(m *model.MaterialReturnSolution) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Delete(&model.MaterialReturnSolutionAvailableCause{}, "`material_return_solution_id` = ?", m.ID).Error; err != nil {
			return err
		}

		if err := tx.Omit("created_at").Save(m).Error; err != nil {
			return err
		}

		return nil
	})
}

func QueryMaterialReturnSolution(req *proto.QueryMaterialReturnSolutionRequest, resp *proto.QueryMaterialReturnSolutionResponse, preload bool) {
	db := model.DB.DB().Model(&model.MaterialReturnSolution{})
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MaterialReturnSolution
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialReturnSolutionsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMaterialReturnSolutions() (list []*model.MaterialReturnSolution, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMaterialReturnSolutionByID(id string) (*model.MaterialReturnSolution, error) {
	m := &model.MaterialReturnSolution{}
	err := model.DB.DB().Preload("MaterialReturnCauses").Preload("MaterialReturnCauses.MaterialReturnCause").Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetMaterialReturnSolutionByIDs(ids []string) ([]*model.MaterialReturnSolution, error) {
	var m []*model.MaterialReturnSolution
	err := model.DB.DB().Preload("MaterialReturnCauses").Preload("MaterialReturnCauses.MaterialReturnCause").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMaterialReturnSolution(id string) (err error) {
	return model.DB.DB().Delete(&model.MaterialReturnSolution{}, "`id` = ?", id).Error
}
