package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductReworkRoute(m *model.ProductReworkRoute) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductReworkRoute(m *model.ProductReworkRoute) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryProductReworkRoute(req *proto.QueryProductReworkRouteRequest, resp *proto.QueryProductReworkRouteResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductReworkRoute{}).Preload("ProductionLine").Preload("MaterialCategory").Preload("FollowProcess").Preload(clause.Associations)
	if req.ProductionLineID != "" {
		db = db.Where("`production_line_id` = ?", req.ProductionLineID)
	}
	if req.Code != "" {
		db = db.Joins("JOIN material_categories ON product_rework_routes.material_category_id = material_categories.id").
			Where("material_categories.code LIKE ? OR material_categories.description LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductReworkRoute
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductReworkRoutesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductReworkRoutes() (list []*model.ProductReworkRoute, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductReworkRouteByID(id string) (*model.ProductReworkRoute, error) {
	m := &model.ProductReworkRoute{}
	err := model.DB.DB().Preload("ProductionLine").Preload("MaterialCategory").Preload("FollowProcess").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProductReworkRouteByIDs(ids []string) ([]*model.ProductReworkRoute, error) {
	var m []*model.ProductReworkRoute
	err := model.DB.DB().Preload("ProductionLine").Preload("MaterialCategory").Preload("FollowProcess").Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductReworkRoute(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductReworkRoute{}, "`id` = ?", id).Error
}
