package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductionStationEfficiency(m *model.ProductionStationEfficiency) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductionStationEfficiency(m *model.ProductionStationEfficiency) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryProductionStationEfficiency(req *proto.QueryProductionStationEfficiencyRequest, resp *proto.QueryProductionStationEfficiencyResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductionStationEfficiency{}).Preload("ProductionStation").Preload(clause.Associations)
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

	var list []*model.ProductionStationEfficiency
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionStationEfficiencysToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductionStationEfficiencys() (list []*model.ProductionStationEfficiency, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductionStationEfficiencyByID(id string) (*model.ProductionStationEfficiency, error) {
	m := &model.ProductionStationEfficiency{}
	err := model.DB.DB().Preload("ProductionStation").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProductionStationEfficiencyByIDs(ids []string) ([]*model.ProductionStationEfficiency, error) {
	var m []*model.ProductionStationEfficiency
	err := model.DB.DB().Preload("ProductionStation").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductionStationEfficiency(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductionStationEfficiency{}, "`id` = ?", id).Error
}
