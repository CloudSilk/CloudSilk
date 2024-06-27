package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductionEfficiency(m *model.ProductionEfficiency) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductionEfficiency(m *model.ProductionEfficiency) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryProductionEfficiency(req *proto.QueryProductionEfficiencyRequest, resp *proto.QueryProductionEfficiencyResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductionEfficiency{}).Preload("ProductionStation").Preload(clause.Associations)
	if req.ProductionLineID != "" {
		db = db.Joins("JOIN production_stations ON production_station_efficiencys.production_station_id=production_stations.id").
			Where("production_stations.production_line_id = ?", req.ProductionLineID)
	}
	if req.OutputDate0 != "" && req.OutputDate1 != "" {
		db = db.Where("production_station_efficiencys.output_date BETWEEN ? AND ?", req.OutputDate0, req.OutputDate1)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductionEfficiency
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionEfficiencysToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductionEfficiencys() (list []*model.ProductionEfficiency, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductionEfficiencyByID(id string) (*model.ProductionEfficiency, error) {
	m := &model.ProductionEfficiency{}
	err := model.DB.DB().Preload("ProductionStation").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProductionEfficiencyByIDs(ids []string) ([]*model.ProductionEfficiency, error) {
	var m []*model.ProductionEfficiency
	err := model.DB.DB().Preload("ProductionStation").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductionEfficiency(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductionEfficiency{}, "`id` = ?", id).Error
}
