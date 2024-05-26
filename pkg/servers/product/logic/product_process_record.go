package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductProcessRecord(m *model.ProductProcessRecord) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductProcessRecord(m *model.ProductProcessRecord) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryProductProcessRecord(req *proto.QueryProductProcessRecordRequest, resp *proto.QueryProductProcessRecordResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductProcessRecord{}).Preload("ProductionStation").Preload("ProductInfo").Preload("ProductInfo.ProductOrder").Preload(clause.Associations)
	if req.ProductSerialNo != "" || req.ProductOrderNo != "" {
		db = db.Joins("JOIN product_infos ON product_process_records.product_info_id = product_infos.id")

		if req.ProductSerialNo != "" {
			db = db.Where("product_infos.product_serial_no LIKE ?", "%"+req.ProductSerialNo+"%")
		}

		if req.ProductOrderNo != "" {
			db = db.Joins("JOIN product_orders ON product_infos.product_order_id = product_orders.id").
				Where("product_orders.product_order_no LIKE ?", "%"+req.ProductOrderNo+"%")
		}
	}
	if req.WorkTime0 != "" && req.WorkTime1 != "" {
		db = db.Where("product_process_records.work_time BETWEEN ? AND ?", req.WorkTime0, req.WorkTime1)
	}
	if req.WorkDescription != "" {
		db = db.Where("product_process_records.work_description LIKE ?", "%"+req.WorkDescription+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductProcessRecord
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductProcessRecordsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductProcessRecords() (list []*model.ProductProcessRecord, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductProcessRecordByID(id string) (*model.ProductProcessRecord, error) {
	m := &model.ProductProcessRecord{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductProcessRecordByIDs(ids []string) ([]*model.ProductProcessRecord, error) {
	var m []*model.ProductProcessRecord
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductProcessRecord(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductProcessRecord{}, "id=?", id).Error
}
