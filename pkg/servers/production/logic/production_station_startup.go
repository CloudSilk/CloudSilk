package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductionStationStartup(m *model.ProductionStationStartup) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductionStationStartup(m *model.ProductionStationStartup) error {
	return model.DB.DB().Omit("startup_time").Save(m).Error
}

func QueryProductionStationStartup(req *proto.QueryProductionStationStartupRequest, resp *proto.QueryProductionStationStartupResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductionStationStartup{}).Preload("ProductionStation").Preload("ProductionStation.ProductionLine").Preload(clause.Associations)
	if req.ProductionLineID != "" {
		db = db.Joins("JOIN production_stations ON production_station_startups.production_station_id=production_stations.id").
			Where("production_stations.production_line_id = ?", req.ProductionLineID)
	}
	if req.StartupTime0 != "" && req.StartupTime1 != "" {
		db = db.Where("`startup_time` BETWEEN ? and ?", req.StartupTime0, req.StartupTime1)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "id")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductionStationStartup
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionStationStartupsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductionStationStartups() (list []*model.ProductionStationStartup, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductionStationStartupByID(id string) (*model.ProductionStationStartup, error) {
	m := &model.ProductionStationStartup{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductionStationStartupByIDs(ids []string) ([]*model.ProductionStationStartup, error) {
	var m []*model.ProductionStationStartup
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductionStationStartup(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductionStationStartup{}, "id=?", id).Error
}
