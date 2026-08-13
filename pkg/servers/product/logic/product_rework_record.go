package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductReworkRecord(m *model.ProductReworkRecord) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductReworkRecord(m *model.ProductReworkRecord) error {
	return model.DB.DB().Omit("created_at", "create_time").Save(m).Error
}

func QueryProductReworkRecord(req *proto.QueryProductReworkRecordRequest, resp *proto.QueryProductReworkRecordResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductReworkRecord{}).Preload("ProductionStation").Preload("ProductionStation.ProductionLine").Preload("ProductInfo").Preload("ProductInfo.ProductOrder").Preload(clause.Associations)
	if req.ProductionLineID != "" {
		db = db.Joins("JOIN production_stations ON product_rework_records.production_Station_id = production_Stations.id").
			Where("production_stations.production_line_id = ?", req.ProductionLineID)
	}
	if req.ProductSerialNo != "" || req.ProductOrderNo != "" {
		db = db.Joins("JOIN product_infos ON product_rework_records.product_info_id = product_infos.id")

		if req.ProductSerialNo != "" {
			db = db.Where("product_infos.product_serial_no LIKE ?", "%"+req.ProductSerialNo+"%")
		}

		if req.ProductOrderNo != "" {
			db = db.Joins("JOIN product_orders ON product_infos.product_order_id = product_orders.id").
				Where("product_orders.product_order_no LIKE ?", "%"+req.ProductOrderNo+"%")
		}
	}
	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("product_rework_records.create_time BETWEEN ? and ?", req.CreateTime0, req.CreateTime1)
	}
	if req.ReworkBrief != "" {
		db.Where("`rework_brief` LIKE ?", "%"+req.ReworkBrief+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductReworkRecord
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductReworkRecordsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductReworkRecords() (list []*model.ProductReworkRecord, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductReworkRecordByID(id string) (*model.ProductReworkRecord, error) {
	m := &model.ProductReworkRecord{}
	err := model.DB.DB().Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProductReworkRecordByIDs(ids []string) ([]*model.ProductReworkRecord, error) {
	var m []*model.ProductReworkRecord
	err := model.DB.DB().Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductReworkRecord(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductReworkRecord{}, "`id` = ?", id).Error
}
