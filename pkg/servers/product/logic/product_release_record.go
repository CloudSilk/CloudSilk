package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductReleaseRecord(m *model.ProductReleaseRecord) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductReleaseRecord(m *model.ProductReleaseRecord) error {
	return model.DB.DB().Omit("created_at", "create_time").Save(m).Error
}

func QueryProductReleaseRecord(req *proto.QueryProductReleaseRecordRequest, resp *proto.QueryProductReleaseRecordResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductReleaseRecord{}).Preload("ProductInfo").Preload("RemanufacturedProductInfo").Preload("RemanufacturedProductInfo.ProductOrder").Preload(clause.Associations)
	if req.ProductSerialNo != "" || req.ProductOrderNo != "" || req.SecurityCode != "" {
		db = db.Joins("JOIN product_infos ON product_release_records.remanufactured_product_info_id = product_infos.id")

		if req.ProductSerialNo != "" {
			db = db.Where("product_infos.product_serial_no LIKE ?", "%"+req.ProductSerialNo+"%")
		}

		if req.ProductOrderNo != "" {
			db = db.Joins("JOIN product_orders ON product_infos.product_order_id = product_orders.id").
				Where("product_orders.product_order_no LIKE ?", "%"+req.ProductOrderNo+"%")
		}
	}
	if req.SecurityCode != "" {
		db.Joins("JOIN product_infos AS PrefabricatedProductInfo ON product_release_records.prefabricated_product_info_id = PrefabricatedProductInfo.id").
			Where("PrefabricatedProductInfo.product_serial_no LIKE ?", "%"+req.SecurityCode+"%")
	}
	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("product_release_records.create_time BETWEEN ? AND ?", req.CreateTime0, req.CreateTime1)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductReleaseRecord
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductReleaseRecordsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductReleaseRecords() (list []*model.ProductReleaseRecord, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductReleaseRecordByID(id string) (*model.ProductReleaseRecord, error) {
	m := &model.ProductReleaseRecord{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductReleaseRecordByIDs(ids []string) ([]*model.ProductReleaseRecord, error) {
	var m []*model.ProductReleaseRecord
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductReleaseRecord(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductReleaseRecord{}, "id=?", id).Error
}
