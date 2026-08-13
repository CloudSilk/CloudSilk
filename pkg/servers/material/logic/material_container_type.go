package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateMaterialContainerType(m *model.MaterialContainerType) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateMaterialContainerType(m *model.MaterialContainerType) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryMaterialContainerType(req *proto.QueryMaterialContainerTypeRequest, resp *proto.QueryMaterialContainerTypeResponse, preload bool) {
	db := model.DB.DB().Model(&model.MaterialContainerType{})
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MaterialContainerType
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialContainerTypesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMaterialContainerTypes() (list []*model.MaterialContainerType, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMaterialContainerTypeByID(id string) (*model.MaterialContainerType, error) {
	m := &model.MaterialContainerType{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetMaterialContainerTypeByIDs(ids []string) ([]*model.MaterialContainerType, error) {
	var m []*model.MaterialContainerType
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMaterialContainerType(id string) (err error) {
	return model.DB.DB().Delete(&model.MaterialContainerType{}, "`id` = ?", id).Error
}
