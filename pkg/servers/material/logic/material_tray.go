package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateMaterialTray(m *model.MaterialTray) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateMaterialTray(m *model.MaterialTray) error {
	return model.DB.DB().Save(m).Error
}

func QueryMaterialTray(req *proto.QueryMaterialTrayRequest, resp *proto.QueryMaterialTrayResponse, preload bool) {
	db := model.DB.DB().Model(&model.MaterialTray{})

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MaterialTray
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialTraysToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMaterialTrays() (list []*model.MaterialTray, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMaterialTrayByID(id string) (*model.MaterialTray, error) {
	m := &model.MaterialTray{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetMaterialTray(req *proto.GetMaterialTrayRequest) (*model.MaterialTray, error) {
	m := &model.MaterialTray{}
	r := map[string]interface{}{}
	if req.Identifier != "" {
		r["identifier"] = req.Identifier
	}
	if req.ProductInfoID != "" {
		r["product_info_id"] = req.ProductInfoID
	}

	err := model.DB.DB().Preload("ProductionLine").Preload("ProductInfo").Preload(clause.Associations).First(m, r).Error
	return m, err
}

func GetMaterialTrayByIDs(ids []string) ([]*model.MaterialTray, error) {
	var m []*model.MaterialTray
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMaterialTray(id string) (err error) {
	return model.DB.DB().Delete(&model.MaterialTray{}, "id=?", id).Error
}
