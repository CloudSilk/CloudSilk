package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreatePersonnelQualificationType(m *model.PersonnelQualificationType) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdatePersonnelQualificationType(m *model.PersonnelQualificationType) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryPersonnelQualificationType(req *proto.QueryPersonnelQualificationTypeRequest, resp *proto.QueryPersonnelQualificationTypeResponse, preload bool) {
	db := model.DB.DB().Model(&model.PersonnelQualificationType{})

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.PersonnelQualificationType
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.PersonnelQualificationTypesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllPersonnelQualificationTypes() (list []*model.PersonnelQualificationType, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetPersonnelQualificationTypeByID(id string) (*model.PersonnelQualificationType, error) {
	m := &model.PersonnelQualificationType{}
	err := model.DB.DB().Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetPersonnelQualificationTypeByIDs(ids []string) ([]*model.PersonnelQualificationType, error) {
	var m []*model.PersonnelQualificationType
	err := model.DB.DB().Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeletePersonnelQualificationType(id string) (err error) {
	return model.DB.DB().Delete(&model.PersonnelQualificationType{}, "`id` = ?", id).Error
}
