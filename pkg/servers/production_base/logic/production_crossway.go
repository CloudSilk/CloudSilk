package logic

import (
	"errors"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProductionCrossway(m *model.ProductionCrossway) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, " code  = ? ", m.Code)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同产线路口配置")
	}
	return m.ID, nil
}

func UpdateProductionCrossway(m *model.ProductionCrossway) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.ProductionCrosswayLeftTurnStation{}, "`production_crossway_id` = ?", m.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(&model.ProductionCrosswayRightTurnStation{}, "`production_crossway_id` = ?", m.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(&model.ProductionCrosswayStraightStation{}, "`production_crossway_id` = ?", m.ID).Error; err != nil {
			return err
		}
		duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(tx, m, true, []string{"created_at"}, "id <> ? and code  = ? ", m.ID, m.Code)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同产线路口配置")
		}

		return nil
	})
}

func QueryProductionCrossway(req *proto.QueryProductionCrosswayRequest, resp *proto.QueryProductionCrosswayResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductionCrossway{})
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

	var list []*model.ProductionCrossway
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionCrosswaysToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductionCrossways() (list []*model.ProductionCrossway, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductionCrosswayByID(id string) (*model.ProductionCrossway, error) {
	m := &model.ProductionCrossway{}
	err := model.DB.DB().Preload("ProductionCrosswayLeftTurnStations").Preload("ProductionCrosswayRightTurnStations").Preload("ProductionCrosswayStraightStations").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProductionCrosswayByIDs(ids []string) ([]*model.ProductionCrossway, error) {
	var m []*model.ProductionCrossway
	err := model.DB.DB().Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductionCrossway(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductionCrossway{}, "`id` = ?", id).Error
}
