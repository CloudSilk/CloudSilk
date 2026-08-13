package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductProcessRoute(m *model.ProductProcessRoute) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductProcessRoute(m *model.ProductProcessRoute) error {
	return model.DB.DB().Omit("created_at", "create_time").Save(m).Error
}

func QueryProductProcessRoute(req *proto.QueryProductProcessRouteRequest, resp *proto.QueryProductProcessRouteResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductProcessRoute{}).Preload("ProductionStation").Preload("ProductInfo").Preload("ProductInfo.ProductOrder").Preload("CurrentProcess").Preload(clause.Associations)
	if req.CurrentProcessID != "" {
		db = db.Where("`current_process_id` = ?", req.CurrentProcessID)
	}
	if req.ProductSerialNo != "" {
		db = db.Joins("JOIN product_infos ON product_process_routes.product_info_id = product_info.id").
			Where("product_infos.product_serial_no = ?", req.ProductSerialNo)
	}
	if req.ProductOrderNo != "" {
		db = db.Joins("JOIN product_infos AS Info1 ON product_process_routes.product_info_id = Info1.ID").
			Joins("JOIN product_orders ON Info1.product_order_id = product_orders.id").
			Where("product_orders.product_order_no = ?", req.ProductOrderNo)
	}
	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("product_process_routes.create_time BETWEEN ? AND ?", req.CreateTime0, req.CreateTime1)
	}

	if req.ProductInfoID != "" {
		db = db.Where("`product_info_id` = ?", req.ProductInfoID)
	}
	if req.RouteIndex > 0 {
		db = db.Where("`route_index` = ?", req.RouteIndex)
	}
	if req.CurrentState != "" {
		db = db.Where("`current_state` = ?", req.CurrentState)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "`created_at` desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductProcessRoute
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductProcessRoutesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductProcessRoutes() (list []*model.ProductProcessRoute, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductProcessRouteByID(id string) (*model.ProductProcessRoute, error) {
	m := &model.ProductProcessRoute{}
	err := model.DB.DB().Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProductProcessRoute(req *proto.GetProductProcessRouteRequest) (*model.ProductProcessRoute, error) {
	m := &model.ProductProcessRoute{}
	err := model.DB.DB().
		Preload("LastProcess").
		Preload("CurrentProcess").
		Preload("ProductionStation").
		Preload("ProductInfo").
		First(m, map[string]interface{}{
			"product_info_id":    req.ProductInfoID,
			"current_process_id": req.CurrentProcessID,
			"current_state":      req.CurrentStates,
		}).Error

	return m, err
}

func GetProductProcessRouteByIDs(ids []string) ([]*model.ProductProcessRoute, error) {
	var m []*model.ProductProcessRoute
	err := model.DB.DB().Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductProcessRoute(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductProcessRoute{}, "`id` = ?", id).Error
}
