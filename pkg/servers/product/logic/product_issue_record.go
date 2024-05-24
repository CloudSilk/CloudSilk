package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductIssueRecord(m *model.ProductIssueRecord) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductIssueRecord(m *model.ProductIssueRecord) error {
	return model.DB.DB().Omit("create_time").Save(m).Error
}

func QueryProductIssueRecord(req *proto.QueryProductIssueRecordRequest, resp *proto.QueryProductIssueRecordResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductIssueRecord{}).Preload("ProductOrderBom").Preload("ProductInfo").Preload("ProductInfo.ProductOrder").Preload("IssuanceProcess").Preload(clause.Associations)
	if req.ProductSerialNo != "" || req.ProductOrderNo != "" {
		db = db.Joins("JOIN product_infos ON product_issue_records.product_info_id = product_infos.id")

		if req.ProductSerialNo != "" {
			db = db.Where("product_infos.product_serial_no LIKE ?", "%"+req.ProductSerialNo+"%")
		}

		if req.ProductOrderNo != "" {
			db = db.Joins("JOIN product_orders ON product_infos.product_order_id = product_orders.id").
				Where("product_orders.product_order_no LIKE ?", "%"+req.ProductOrderNo+"%")
		}
	}
	if req.MaterialDescription != "" {
		db = db.Joins("JOIN product_order_boms ON product_issue_records.product_order_bom_id = product_order_boms.id").
			Where("product_order_boms.material_description LIKE ?", "%"+req.MaterialDescription+"%")
	}
	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("product_issue_records.create_time BETWEEN ? AND ?", req.CreateTime0, req.CreateTime1)
	}
	if req.ProductionProcess != "" {
		db = db.Joins("JOIN production_processes ON product_issue_records.issuance_process_id = production_processes.id").
			Where("production_processes.description LIKE ?", "%"+req.ProductionProcess+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductIssueRecord
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductIssueRecordsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductIssueRecords() (list []*model.ProductIssueRecord, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductIssueRecordByID(id string) (*model.ProductIssueRecord, error) {
	m := &model.ProductIssueRecord{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductIssueRecordByIDs(ids []string) ([]*model.ProductIssueRecord, error) {
	var m []*model.ProductIssueRecord
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductIssueRecord(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductIssueRecord{}, "id=?", id).Error
}
