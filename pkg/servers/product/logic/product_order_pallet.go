package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductOrderPallet(m *model.ProductOrderPallet) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductOrderPallet(m *model.ProductOrderPallet) error {
	return model.DB.DB().Omit("create_time").Save(m).Error
}

func QueryProductOrderPallet(req *proto.QueryProductOrderPalletRequest, resp *proto.QueryProductOrderPalletResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductOrderPallet{}).Preload("ProductPackageStackRule").Preload("ProductOrder").Preload(clause.Associations)
	if req.PalletNo != "" {
		db = db.Where("product_order_pallets.pallet_no LIKE ?", "%"+req.PalletNo+"%")
	}

	if req.PalletSn != "" {
		db = db.Where("product_order_pallets.pallet_sn LIKE ?", "%"+req.PalletSn+"%")
	}

	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("product_order_pallets.create_time BETWEEN ? AND ?", req.CreateTime0, req.CreateTime1)
	}

	if req.CurrentState != "" {
		db = db.Where("product_order_pallets.current_state = ?", req.CurrentState)
	}

	if req.ProductOrderNo != "" {
		db = db.Joins("JOIN product_orders ON model.product_order_pallets.product_order_id=product_orders.id").
			Where("product_orders.product_order_no like ?", "%"+req.ProductOrderNo+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductOrderPallet
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductOrderPalletsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductOrderPallets() (list []*model.ProductOrderPallet, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductOrderPalletByID(id string) (*model.ProductOrderPallet, error) {
	m := &model.ProductOrderPallet{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductOrderPalletByIDs(ids []string) ([]*model.ProductOrderPallet, error) {
	var m []*model.ProductOrderPallet
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductOrderPallet(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductOrderPallet{}, "id=?", id).Error
}
