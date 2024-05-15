package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateLabelType(m *model.LabelType) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateLabelType(m *model.LabelType) error {
	return model.DB.DB().Save(m).Error
}

func QueryLabelType(req *proto.QueryLabelTypeRequest, resp *proto.QueryLabelTypeResponse, preload bool) {
	db := model.DB.DB().Model(&model.LabelType{})
	if req.Code != "" {
		db = db.Where("code LIKE ? OR description LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "id")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.LabelType
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.LabelTypesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllLabelTypes() (list []*model.LabelType, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetLabelTypeByID(id string) (*model.LabelType, error) {
	m := &model.LabelType{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetLabelTypeByIDs(ids []string) ([]*model.LabelType, error) {
	var m []*model.LabelType
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteLabelType(id string) (err error) {
	return model.DB.DB().Delete(&model.LabelType{}, "id=?", id).Error
}
