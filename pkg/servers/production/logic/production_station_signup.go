package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductionStationSignup(m *model.ProductionStationSignup) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductionStationSignup(m *model.ProductionStationSignup) error {
	return model.DB.DB().Omit("login_time").Save(m).Error
}

func QueryProductionStationSignup(req *proto.QueryProductionStationSignupRequest, resp *proto.QueryProductionStationSignupResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductionStationSignup{}).Preload("ProductionStation").Preload("ProductionStation.ProductionLine").Preload(clause.Associations)
	if req.ProductionLineID != "" {
		db = db.Joins("JOIN production_stations ON production_station_signups.production_station_id=production_stations.id").
			Where("production_stations.production_line_id = ?", req.ProductionLineID)
	}
	if req.LoginTime0 != "" && req.LoginTime1 != "" {
		db = db.Where("`login_time` BETWEEN ? and ?", req.LoginTime0, req.LoginTime1)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "id")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductionStationSignup
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionStationSignupsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductionStationSignups() (list []*model.ProductionStationSignup, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductionStationSignupByID(id string) (*model.ProductionStationSignup, error) {
	m := &model.ProductionStationSignup{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductionStationSignupByIDs(ids []string) ([]*model.ProductionStationSignup, error) {
	var m []*model.ProductionStationSignup
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductionStationSignup(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductionStationSignup{}, "id=?", id).Error
}

func GetProductionStationSignup(productionStationID, loginUserID string, hasLogoutTime bool) (*model.ProductionStationSignup, error) {
	m := &model.ProductionStationSignup{}

	logoutTime := " and logout_time IS NOT NULL"
	if !hasLogoutTime {
		logoutTime = " and logout_time IS NULL"
	}

	err := model.DB.DB().Preload(clause.Associations).Where("productionStationID = ? and loginUserID=?"+logoutTime,
		productionStationID, loginUserID).First(m).Error
	return m, err
}
