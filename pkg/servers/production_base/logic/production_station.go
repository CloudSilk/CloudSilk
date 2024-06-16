package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductionStation(m *model.ProductionStation) (string, error) {
	err := model.DB.DB().Model(m).Create(m).Error
	return m.ID, err
}

func UpdateProductionStation(m *model.ProductionStation) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryProductionStation(req *proto.QueryProductionStationRequest, resp *proto.QueryProductionStationResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductionStation{}).Preload("ProductionLine").Preload(clause.Associations)
	if req.ProductionLineID != "" {
		db = db.Where("`production_line_id` = ?", req.ProductionLineID)
	}
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}
	if req.StationType != "" {
		db = db.Where("`station_type` = ?", req.StationType)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductionStation
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionStationsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductionStations() (list []*model.ProductionStation, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductionStationByID(id string) (*model.ProductionStation, error) {
	m := &model.ProductionStation{}
	err := model.DB.DB().Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProductionStation(req *proto.GetProductionStationRequest) (*model.ProductionStation, error) {
	m := &model.ProductionStation{}
	err := model.DB.DB().Preload(clause.Associations).Where("`code` = ?", req.Code).First(m).Error
	return m, err
}

func GetProductionStationByIDs(ids []string) ([]*model.ProductionStation, error) {
	var m []*model.ProductionStation
	err := model.DB.DB().Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductionStation(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductionStation{}, "`id` = ?", id).Error
}
