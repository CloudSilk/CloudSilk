package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateMaterialCategory(m *model.MaterialCategory) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateMaterialCategory(m *model.MaterialCategory) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryMaterialCategory(req *proto.QueryMaterialCategoryRequest, resp *proto.QueryMaterialCategoryResponse, preload bool) {
	db := model.DB.DB().Model(&model.MaterialCategory{})
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MaterialCategory
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialCategorysToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMaterialCategorys() (list []*model.MaterialCategory, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMaterialCategoryByID(id string) (*model.MaterialCategory, error) {
	m := &model.MaterialCategory{}
	err := model.DB.DB().Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetMaterialCategoryByIDs(ids []string) ([]*model.MaterialCategory, error) {
	var m []*model.MaterialCategory
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMaterialCategory(id string) (err error) {
	return model.DB.DB().Delete(&model.MaterialCategory{}, "`id` = ?", id).Error
}
