package logic

import (
	"errors"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProductionFactory(m *model.ProductionFactory) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, " code =? ", m.Code)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同工厂配置")
	}
	return m.ID, nil
}

func UpdateProductionFactory(m *model.ProductionFactory) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		oldProductionFactory := &model.ProductionFactory{}
		if err := tx.First(oldProductionFactory, "id = ?", m.ID).Error; err != nil {
			return err
		}

		if err := tx.Delete(&model.ProductionLine{}, "production_factory_id=?", m.ID).Error; err != nil {
			return err
		}

		duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(tx, m, true, []string{}, "id <> ?  and  code =? ", m.ID, m.Code)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同工厂配置")
		}

		return nil
	})
}

func QueryProductionFactory(req *proto.QueryProductionFactoryRequest, resp *proto.QueryProductionFactoryResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductionFactory{})
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductionFactory
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionFactorysToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductionFactorys() (list []*model.ProductionFactory, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductionFactoryByID(id string) (*model.ProductionFactory, error) {
	m := &model.ProductionFactory{}
	err := model.DB.DB().Preload("ProductionLines.ProductionStations").Preload("ProductionLines.ProductionCrossways").Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductionFactoryByIDs(ids []string) ([]*model.ProductionFactory, error) {
	var m []*model.ProductionFactory
	err := model.DB.DB().Preload("ProductionLines.ProductionStations").Preload("ProductionLines.ProductionCrossways").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductionFactory(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductionFactory{}, "id=?", id).Error
}
