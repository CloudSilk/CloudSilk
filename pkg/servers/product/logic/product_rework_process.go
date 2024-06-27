package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProductReworkProcess(m *model.ProductReworkProcess) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductReworkProcess(m *model.ProductReworkProcess) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.ProductReworkProcessAvailableStation{}, "`product_rework_process_id` = ?", m.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(&model.ProductReworkProcessAvailableProcess{}, "`product_rework_process_id` = ?", m.ID).Error; err != nil {
			return err
		}

		if err := tx.Omit("created_at").Save(m).Error; err != nil {
			return err
		}

		return nil
	})
}

func QueryProductReworkProcess(req *proto.QueryProductReworkProcessRequest, resp *proto.QueryProductReworkProcessResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductReworkProcess{}).Preload("ProductionLine").Preload(clause.Associations)
	if req.ProductionLineID != "" {
		db = db.Where("`production_line_id` = ?", req.ProductionLineID)
	}
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductReworkProcess
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductReworkProcesssToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductReworkProcesss() (list []*model.ProductReworkProcess, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductReworkProcessByID(id string) (*model.ProductReworkProcess, error) {
	m := &model.ProductReworkProcess{}
	err := model.DB.DB().Preload("ProductionLine").Preload("ProductionStations").Preload("ProductionProcesses").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProductReworkProcessByIDs(ids []string) ([]*model.ProductReworkProcess, error) {
	var m []*model.ProductReworkProcess
	err := model.DB.DB().Preload("ProductionLine").Preload("ProductionStations").Preload("ProductionProcesses").Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductReworkProcess(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductReworkProcess{}, "`id` = ?", id).Error
}
