package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductionStationAlarm(m *model.ProductionStationAlarm) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductionStationAlarm(m *model.ProductionStationAlarm) error {
	return model.DB.DB().Omit("create_time").Save(m).Error
}

func QueryProductionStationAlarm(req *proto.QueryProductionStationAlarmRequest, resp *proto.QueryProductionStationAlarmResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductionStationAlarm{}).Preload("ProductionStation").Preload("ProductionStation.ProductionLine").Preload("ProductionProcess").Preload("ProductInfo").Preload("ProductInfo.ProductOrder").Preload(clause.Associations)
	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("production_station_alarms.create_time BETWEEN ? AND ?", req.CreateTime0, req.CreateTime1)
	}
	if req.ProductionLineID != "" {
		db = db.Joins("JOIN production_stations ON production_station_alarms.production_station_id=production_stations.id").
			Where("production_stations.production_line_id = ?", req.ProductionLineID)
	}
	if req.ProductOrderNo != "" || req.ProductSerialNo != "" {
		db.Joins("JOIN product_infos ON production_station_alarms.product_info_id=product_infos.id")

		if req.ProductOrderNo != "" {
			db = db.Joins("JOIN product_orders ON product_infos.product_order_id=product_orders.id").
				Where("product_orders.product_order_no LIKE ?", "%"+req.ProductOrderNo+"%")
		}
		if req.ProductSerialNo != "" {
			db = db.Where("product_infos.product_serial_no LIKE ?", "%"+req.ProductSerialNo+"%")
		}
	}
	if req.AlarmMessage != "" {
		db.Where("production_station_alarms.alarm_message LIKE ?", "%"+req.AlarmMessage+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "id")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductionStationAlarm
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionStationAlarmsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductionStationAlarms() (list []*model.ProductionStationAlarm, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductionStationAlarmByID(id string) (*model.ProductionStationAlarm, error) {
	m := &model.ProductionStationAlarm{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductionStationAlarmByIDs(ids []string) ([]*model.ProductionStationAlarm, error) {
	var m []*model.ProductionStationAlarm
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductionStationAlarm(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductionStationAlarm{}, "id=?", id).Error
}
