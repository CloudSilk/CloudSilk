package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductPackageRecord(m *model.ProductPackageRecord) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductPackageRecord(m *model.ProductPackageRecord) error {
	return model.DB.DB().Save(m).Error
}

func QueryProductPackageRecord(req *proto.QueryProductPackageRecordRequest, resp *proto.QueryProductPackageRecordResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductPackageRecord{}).Preload("ProductInfo").Preload("ProductInfo.ProductOrder").Preload("ProductInfo.ProductOrder.ProductModel").Preload("ProductOrderPackage").Preload("ProductOrderPackage.ProductPackage").Preload(clause.Associations)
	if req.ProductionLineID != "" {
		db = db.Joins("JOIN product_infos ON product_package_records.product_info_id = product_infos.id").
			Joins("JOIN product_orders ON product_infos.product_order_id = product_orders.id").
			Where("product_orders.production_line_id = ?", req.ProductionLineID)
	}
	if req.PackageNo != "" {
		db = db.Joins("JOIN product_order_packages ON product_package_records.product_order_package_id = product_order_packages.id").
			Where("product_order_packages.package_no LIKE ?", "%"+req.PackageNo+"%")
	}
	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("product_package_records.create_time BETWEEN ? and ?", req.CreateTime0, req.CreateTime1)
	}

	if req.ProductSerialNo != "" {
		db = db.Joins("JOIN product_infos ON product_package_records.product_info_id = product_infos.id").
			Joins("JOIN product_orders ON product_infos.product_order_id = product_orders.id").
			Where("product_infos.product_serial_no LIKE ? or product_orders.product_order_no LIKE ? or product_orders.sales_order_no LIKE ?", "%"+req.ProductSerialNo+"%", "%"+req.ProductSerialNo+"%", "%"+req.ProductSerialNo+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductPackageRecord
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductPackageRecordsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductPackageRecords() (list []*model.ProductPackageRecord, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductPackageRecordByID(id string) (*model.ProductPackageRecord, error) {
	m := &model.ProductPackageRecord{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductPackageRecordByPackageNo(packageNo string) (*model.ProductPackageRecord, error) {
	m := &model.ProductPackageRecord{}
	err := model.DB.DB().Preload("ProductOrderPackage").Preload(clause.Associations).Where("package_no = ?", packageNo).First(m).Error
	return m, err
}

func GetProductPackageRecordByIDs(ids []string) ([]*model.ProductPackageRecord, error) {
	var m []*model.ProductPackageRecord
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductPackageRecord(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductPackageRecord{}, "id=?", id).Error
}
