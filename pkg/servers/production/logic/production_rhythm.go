package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProductionRhythm(m *model.ProductionRhythm) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductionRhythm(m *model.ProductionRhythm) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.AttributeExpression{}, "rule_id=? AND rule_type=?", m.ID, "ProductionRhythm").Error; err != nil {
			return err
		}

		if err := tx.Omit("created_at").Save(m).Error; err != nil {
			return err
		}

		return nil
	})
}

func QueryProductionRhythm(req *proto.QueryProductionRhythmRequest, resp *proto.QueryProductionRhythmResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductionRhythm{})
	if req.ProductionLineID != "" {
		db = db.Where("`production_line_id` = ?", req.ProductionLineID)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "priority")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductionRhythm
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionRhythmsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductionRhythms() (list []*model.ProductionRhythm, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductionRhythmByID(id string) (*model.ProductionRhythm, error) {
	m := &model.ProductionRhythm{}
	err := model.DB.DB().Preload("AttributeExpressions", func(db *gorm.DB) *gorm.DB {
		return db.Where("rule_type", "ProductionRhythm")
	}).Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductionRhythmByIDs(ids []string) ([]*model.ProductionRhythm, error) {
	var m []*model.ProductionRhythm
	err := model.DB.DB().Preload("AttributeExpressions", func(db *gorm.DB) *gorm.DB {
		return db.Where("rule_type", "ProductionRhythm")
	}).Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductionRhythm(id string) (err error) {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.ProductionRhythm{}, "id=?", id).Error; err != nil {
			return err
		}
		if err := tx.Delete(&model.AttributeExpression{}, "rule_id=? AND rule_type=?", id, "ProductionRhythm").Error; err != nil {
			return err
		}
		return nil
	})
}
