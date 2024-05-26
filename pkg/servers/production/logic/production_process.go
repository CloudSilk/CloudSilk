package logic

import (
	"errors"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProductionProcess(m *model.ProductionProcess) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, " code =? ", m.Code)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同生产工序")
	}
	return m.ID, nil
}

func UpdateProductionProcess(m *model.ProductionProcess) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.ProductionProcessAvailableStation{}, "production_process_id=?", m.ID).Error; err != nil {
			return err
		}

		if err := tx.Delete(&model.AttributeExpression{}, "rule_id=? AND rule_type=?", m.ID, "ProductionProcess").Error; err != nil {
			return err
		}

		duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(tx, m, true, []string{"created_at"}, "id <> ?  and  code =? ", m.ID, m.Code)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同生产工序")
		}

		return nil
	})
}

func QueryProductionProcess(req *proto.QueryProductionProcessRequest, resp *proto.QueryProductionProcessResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductionProcess{}).Preload("ProductionLine").Preload(clause.Associations)
	if req.ProductionLineID != "" {
		db = db.Where("`production_line_id` = ?", req.ProductionLineID)
	}
	if req.Code != "" {
		db = db.Where("`code` = ?", req.Code)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "sort_index")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductionProcess
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionProcesssToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductionProcesss() (list []*model.ProductionProcess, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductionProcessByID(id string) (*model.ProductionProcess, error) {
	m := &model.ProductionProcess{}
	err := model.DB.DB().Preload("ProductionLine").Preload("ProductionProcessAvailableStations").Preload("ProductionProcessAvailableStations.ProductionStation").
		Preload("AttributeExpressions", func(db *gorm.DB) *gorm.DB {
			return db.Where("rule_type", "ProductionProcess")
		}).Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductionProcessByIDs(ids []string) ([]*model.ProductionProcess, error) {
	var m []*model.ProductionProcess
	err := model.DB.DB().Preload("ProductionLine").Preload("ProductionProcessAvailableStations").Preload("ProductionProcessAvailableStations.ProductionStation").
		Preload("AttributeExpressions", func(db *gorm.DB) *gorm.DB {
			return db.Where("rule_type", "ProductionProcess")
		}).Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductionProcess(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductionProcess{}, "id=?", id).Error
}
