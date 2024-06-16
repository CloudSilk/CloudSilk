package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductWorkRecord(m *model.ProductWorkRecord) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductWorkRecord(m *model.ProductWorkRecord) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryProductWorkRecord(req *proto.QueryProductWorkRecordRequest, resp *proto.QueryProductWorkRecordResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductWorkRecord{}).Preload("ProductionStation").Preload("ProductionStation.ProductionLine").Preload("ProductionProcessStep").Preload("ProductionProcessStep.ProcessStepType").Preload("ProductInfo").Preload("ProductInfo.ProductOrder").Preload(clause.Associations)
	if req.ProductOrderNo != "" || req.ProductSerialNo != "" {
		db.Joins("JOIN product_infos ON product_work_records.product_info_id=product_infos.id")

		if req.ProductOrderNo != "" {
			db = db.Joins("JOIN product_orders ON product_infos.product_order_id=product_orders.id").
				Where("product_orders.product_order_no LIKE ?", "%"+req.ProductOrderNo+"%")
		}
		if req.ProductSerialNo != "" {
			db = db.Where("product_infos.product_serial_no LIKE ?", "%"+req.ProductSerialNo+"%")
		}
	}
	if req.WorkStartTime0 != "" && req.WorkStartTime1 != "" {
		db = db.Where("product_work_records.work_start_time BETWEEN ? AND ?", req.WorkStartTime0, req.WorkStartTime1)
	}
	if req.ProductionLineID != "" {
		db = db.Joins("JOIN production_stations ON product_work_records.production_station_id=production_stations.id").
			Where("production_stations.production_line_id = ?", req.ProductionLineID)
	}
	if req.ProductInfoID != "" {
		db = db.Where("`product_info_id` = ?", req.ProductInfoID)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductWorkRecord
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductWorkRecordsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductWorkRecords() (list []*model.ProductWorkRecord, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductWorkRecordByID(id string) (*model.ProductWorkRecord, error) {
	m := &model.ProductWorkRecord{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductWorkRecordByIDs(ids []string) ([]*model.ProductWorkRecord, error) {
	var m []*model.ProductWorkRecord
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductWorkRecord(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductWorkRecord{}, "id = ?", id).Error
}
