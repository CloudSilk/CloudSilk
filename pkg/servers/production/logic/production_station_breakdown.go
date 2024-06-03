package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductionStationBreakdown(m *model.ProductionStationBreakdown) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductionStationBreakdown(m *model.ProductionStationBreakdown) error {
	return model.DB.DB().Omit("created_at", "create_time").Save(m).Error
}

func QueryProductionStationBreakdown(req *proto.QueryProductionStationBreakdownRequest, resp *proto.QueryProductionStationBreakdownResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductionStationBreakdown{}).Preload("ProductionStation").Preload("ProductionStation.ProductionLine").Preload(clause.Associations)
	if req.ProductionLineID != "" {
		db = db.Joins("JOIN production_stations ON production_station_breakdowns.production_station_id=production_stations.id").
			Where("production_stations.production_line_id = ?", req.ProductionLineID)
	}
	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("production_station_breakdowns.create_time BETWEEN ? AND ?", req.CreateTime0, req.CreateTime1)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductionStationBreakdown
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionStationBreakdownsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductionStationBreakdowns() (list []*model.ProductionStationBreakdown, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductionStationBreakdownByID(id string) (*model.ProductionStationBreakdown, error) {
	m := &model.ProductionStationBreakdown{}
	err := model.DB.DB().Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProductionStationBreakdownByIDs(ids []string) ([]*model.ProductionStationBreakdown, error) {
	var m []*model.ProductionStationBreakdown
	err := model.DB.DB().Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductionStationBreakdown(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductionStationBreakdown{}, "`id` = ?", id).Error
}
