package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductRhythmRecord(m *model.ProductRhythmRecord) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductRhythmRecord(m *model.ProductRhythmRecord) error {
	return model.DB.DB().Omit("created_at", "create_time").Save(m).Error
}

func QueryProductRhythmRecord(req *proto.QueryProductRhythmRecordRequest, resp *proto.QueryProductRhythmRecordResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductRhythmRecord{}).Preload("ProductionStation").Preload("ProductionProcess").Preload("ProductInfo").Preload("ProductInfo.ProductOrder").Preload(clause.Associations)
	if req.ProductOrderNo != "" || req.ProductSerialNo != "" {
		db.Joins("JOIN product_infos ON product_rhythm_records.product_info_id=product_infos.id")

		if req.ProductOrderNo != "" {
			db = db.Joins("JOIN product_orders ON product_infos.product_order_id=product_orders.id").
				Where("product_orders.product_order_no LIKE ?", "%"+req.ProductOrderNo+"%")
		}
		if req.ProductSerialNo != "" {
			db = db.Where("product_infos.product_serial_no LIKE ?", "%"+req.ProductSerialNo+"%")
		}
	}
	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("product_rhythm_records.create_time BETWEEN ? and ?", req.CreateTime0, req.CreateTime1)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductRhythmRecord
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductRhythmRecordsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductRhythmRecords() (list []*model.ProductRhythmRecord, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductRhythmRecordByID(id string) (*model.ProductRhythmRecord, error) {
	m := &model.ProductRhythmRecord{}
	err := model.DB.DB().Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProductRhythmRecord(req *proto.GetProductRhythmRecordRequest) (*model.ProductRhythmRecord, error) {
	m := &model.ProductRhythmRecord{}

	workEndTime := "work_end_time IS NOT NULL"
	if !req.HasWorkEndTime {
		workEndTime = "work_end_time IS NULL"
	}

	whereMap := map[string]interface{}{}
	if req.ProductionProcessID != "" {
		whereMap["production_process_id"] = req.ProductionProcessID
	}
	if req.ProductInfoID != "" {
		whereMap["product_info_id"] = req.ProductInfoID
	}
	if req.ProductionStationID != "" {
		whereMap["production_station_id"] = req.ProductionStationID
	}

	err := model.DB.DB().Preload(clause.Associations).First(m, whereMap).Where(workEndTime).Error
	return m, err
}

func GetProductRhythmRecordByIDs(ids []string) ([]*model.ProductRhythmRecord, error) {
	var m []*model.ProductRhythmRecord
	err := model.DB.DB().Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductRhythmRecord(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductRhythmRecord{}, "`id` = ?", id).Error
}
