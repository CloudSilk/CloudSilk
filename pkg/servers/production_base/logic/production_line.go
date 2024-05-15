package logic

import (
	"errors"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProductionLine(m *model.ProductionLine) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, " code =? ", m.Code)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同产线")
	}
	return m.ID, nil
}

func UpdateProductionLine(m *model.ProductionLine) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.ProductionStation{}, "production_line_id=?", m.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(&model.ProductionCrossway{}, "production_line_id=?", m.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(&model.ProductionLineSupportableCategory{}, "production_line_id=?", m.ID).Error; err != nil {
			return err
		}

		duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(tx, m, true, []string{}, "id <> ?  and  code =? ", m.ID, m.Code)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同产线")
		}

		return nil
	})
}

func QueryProductionLine(req *proto.QueryProductionLineRequest, resp *proto.QueryProductionLineResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductionLine{})
	if req.ProductionFactoryID != "" {
		db = db.Where("`production_factory_id` = ?", req.ProductionFactoryID)
	}
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "id")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductionLine
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionLinesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductionLines() (list []*model.ProductionLine, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductionLineByID(id string) (*model.ProductionLine, error) {
	m := &model.ProductionLine{}
	err := model.DB.DB().
		Preload("ProductionStations").
		Preload("ProductionCrossways").
		Preload("ProductionLineSupportableCategorys").
		Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductionLineByIDs(ids []string) ([]*model.ProductionLine, error) {
	var m []*model.ProductionLine
	err := model.DB.DB().
		Preload("ProductionStations").
		Preload("ProductionCrossways").
		Preload("ProductionLineSupportableCategorys").
		Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductionLine(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductionLine{}, "id=?", id).Error
}

// func DeleteProductionLine(id string) (err error) {

// productionLine := []interface{}{&ProductionStation{}, &ProductionCrossway{}, &ProductionProcess{}, &ProductOrderReleaseRule{},
// 	&ProductOrder{}, &ProductReleaseStrategy{}, &ProductionCrosswayRightTurnStation{}, &ProductionLineSupportableCategory{}, &ProductionRhythm{}}

// if err := model.DB.DB().Transaction(func(tx *gorm.DB) error {
// 	for _, model := range productionLine {
// 		if err := tx.Delete(model, "ProductionLineID=?", id).Error; err != nil {
// 			return err
// 		}
// 	}

// 	if err := tx.Delete(&ProductionLine{}, "id=?", id).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }); err != nil {
// 	return fmt.Errorf("数据冲突，先删除关联数据")
// }

// 	return nil
// }
