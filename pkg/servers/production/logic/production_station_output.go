package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductionStationOutput(m *model.ProductionStationOutput) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductionStationOutput(m *model.ProductionStationOutput) error {
	return model.DB.DB().Omit("output_time").Save(m).Error
}

func QueryProductionStationOutput(req *proto.QueryProductionStationOutputRequest, resp *proto.QueryProductionStationOutputResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductionStationOutput{}).Preload("ProductionStation").Preload("ProductionStation.ProductionLine").Preload("ProductionProcess").Preload("ProductInfo").Preload("ProductInfo.ProductOrder").Preload(clause.Associations)
	if req.ProductionLineID != "" {
		db = db.Joins("JOIN production_processes ON production_station_outputs.production_processe_id=production_processes.id").
			Where("production_processes.production_line_id = ?", req.ProductionLineID)
	}
	if req.OutputTime0 != "" && req.OutputTime1 != "" {
		db.Where("`output_time` BETWEEN ? and ?", req.OutputTime0, req.OutputTime1)
	}
	if req.ProductSerialNo != "" || req.ProductOrderNo != "" {
		db = db.Joins("JOIN product_infos ON production_station_outputs.product_info_id = product_infos.id")
		if req.ProductSerialNo != "" {
			db = db.Where("product_infos.product_serial_no LIKE ?", "%"+req.ProductSerialNo+"%")
		}
		if req.ProductOrderNo != "" {
			db = db.Joins("JOIN product_orders ON product_infos.product_order_id = product_orders.id").
				Where("product_orders.product_order_no LIKE ?", "%"+req.ProductOrderNo+"%")
		}
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "id")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductionStationOutput
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionStationOutputsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductionStationOutputs() (list []*model.ProductionStationOutput, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductionStationOutputByID(id string) (*model.ProductionStationOutput, error) {
	m := &model.ProductionStationOutput{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductionStationOutput(req *proto.GetProductionStationOutputRequest) (*model.ProductionStationOutput, error) {
	m := &model.ProductionStationOutput{}
	err := model.DB.DB().Preload("ProductionStation").Preload("ProductionProcess").Preload("ProductInfo").Preload(clause.Associations).First(m, map[string]interface{}{
		"production_station_id": req.ProductionStationID,
		"product_info_id":       req.ProductInfoID,
	}).Error
	return m, err
}

func GetProductionStationOutputByIDs(ids []string) ([]*model.ProductionStationOutput, error) {
	var m []*model.ProductionStationOutput
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductionStationOutput(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductionStationOutput{}, "id=?", id).Error
}
