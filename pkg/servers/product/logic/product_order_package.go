package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductOrderPackage(m *model.ProductOrderPackage) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductOrderPackage(m *model.ProductOrderPackage) error {
	return model.DB.DB().Omit("create_time").Save(m).Error
}

func QueryProductOrderPackage(req *proto.QueryProductOrderPackageRequest, resp *proto.QueryProductOrderPackageResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductOrderPackage{}).Preload("ProductOrder").Preload("ProductPackage").Preload(clause.Associations)
	if req.PackageNo != "" {
		db = db.Where("`package_no` LIKE ?", "%"+req.PackageNo+"%")
	}
	if req.PalletNo != "" {
		db = db.Where("`pallet_no` LIKE ?", "%"+req.PalletNo+"%")
	}
	if req.CurrentState != "" {
		db = db.Where("`current_state` = ?", req.CurrentState)
	}
	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("`create_time` BETWEEN ? AND ?", req.CreateTime0, req.CreateTime1)
	}
	if req.ProductOrderNo != "" {
		db = db.Joins("JOIN product_orders ON product_order_package.product_order_id=product_orders.id").
			Where("product_orders.product_order_no like ?", "%"+req.ProductOrderNo+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "id")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductOrderPackage
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductOrderPackagesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductOrderPackages() (list []*model.ProductOrderPackage, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductOrderPackageByID(id string) (*model.ProductOrderPackage, error) {
	m := &model.ProductOrderPackage{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductOrderPackageByIDs(ids []string) ([]*model.ProductOrderPackage, error) {
	var m []*model.ProductOrderPackage
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductOrderPackage(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductOrderPackage{}, "id=?", id).Error
}
