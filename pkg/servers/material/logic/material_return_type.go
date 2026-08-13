package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateMaterialReturnType(m *model.MaterialReturnType) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateMaterialReturnType(m *model.MaterialReturnType) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryMaterialReturnType(req *proto.QueryMaterialReturnTypeRequest, resp *proto.QueryMaterialReturnTypeResponse, preload bool) {
	db := model.DB.DB().Model(&model.MaterialReturnType{})
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MaterialReturnType
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialReturnTypesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMaterialReturnTypes() (list []*model.MaterialReturnType, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMaterialReturnTypeByID(id string) (*model.MaterialReturnType, error) {
	m := &model.MaterialReturnType{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetMaterialReturnTypeByIDs(ids []string) ([]*model.MaterialReturnType, error) {
	var m []*model.MaterialReturnType
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMaterialReturnType(id string) (err error) {
	return model.DB.DB().Delete(&model.MaterialReturnType{}, "`id` = ?", id).Error
}
