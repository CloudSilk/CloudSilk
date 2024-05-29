package logic

import (
	"errors"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProductReworkSolution(m *model.ProductReworkSolution) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, " `code`  = ? ", m.Code)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同产品返工方案")
	}
	return m.ID, nil
}

func UpdateProductReworkSolution(m *model.ProductReworkSolution) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.ProductReworkCauseAvailableSolution{}, "`product_rework_solution_id` = ?", m.ID).Error; err != nil {
			return err
		}

		duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(tx, m, true, []string{"created_at"}, "`id` <> ?  and  `code`  = ? ", m.ID, m.Code)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同产品返工方案")
		}

		return nil
	})
}

func QueryProductReworkSolution(req *proto.QueryProductReworkSolutionRequest, resp *proto.QueryProductReworkSolutionResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductReworkSolution{})
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductReworkSolution
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductReworkSolutionsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductReworkSolutions() (list []*model.ProductReworkSolution, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductReworkSolutionByID(id string) (*model.ProductReworkSolution, error) {
	m := &model.ProductReworkSolution{}
	err := model.DB.DB().Preload("ProductReworkCauseAvailableSolutions").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProductReworkSolutionByIDs(ids []string) ([]*model.ProductReworkSolution, error) {
	var m []*model.ProductReworkSolution
	err := model.DB.DB().Preload("ProductReworkCauseAvailableSolutions").Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductReworkSolution(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductReworkSolution{}, "`id` = ?", id).Error
}
