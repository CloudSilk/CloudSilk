package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductionProcessSop(m *model.ProductionProcessSop) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductionProcessSop(m *model.ProductionProcessSop) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryProductionProcessSop(req *proto.QueryProductionProcessSopRequest, resp *proto.QueryProductionProcessSopResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductionProcessSop{}).Preload("ProductModel").Preload(clause.Associations)
	if req.ProductionLineID != "" {
		db = db.Joins("JOIN production_processes ON production_process_sops.production_process_id=production_processes.id").
			Where("production_processes.production_line_id = ?", req.ProductionLineID)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductionProcessSop
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionProcessSopsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductionProcessSops() (list []*model.ProductionProcessSop, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductionProcessSopByID(id string) (*model.ProductionProcessSop, error) {
	m := &model.ProductionProcessSop{}
	err := model.DB.DB().Preload("ProductionProcess").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProductionProcessSop(req *proto.GetProductionProcessSopRequest) (*model.ProductionProcessSop, error) {
	m := &model.ProductionProcessSop{}
	err := model.DB.DB().Preload("ProductionProcess").First(m, map[string]interface{}{
		"production_process_id": req.ProductionProcessID,
		"product_model_id":      req.ProductModelID,
	}).Error
	return m, err
}

func GetProductionProcessSopByIDs(ids []string) ([]*model.ProductionProcessSop, error) {
	var m []*model.ProductionProcessSop
	err := model.DB.DB().Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductionProcessSop(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductionProcessSop{}, "`id` = ?", id).Error
}
