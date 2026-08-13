package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreatePersonnelQualification(m *model.PersonnelQualification) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdatePersonnelQualification(m *model.PersonnelQualification) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryPersonnelQualification(req *proto.QueryPersonnelQualificationRequest, resp *proto.QueryPersonnelQualificationResponse, preload bool) {
	db := model.DB.DB().Model(&model.PersonnelQualification{})
	if req.ProductionLineID != "" {
		db = db.Joins("JOIN production_processes ON personnel_qualifications.production_process_id=production_processes.id").
			Where("production_processes.production_line_id = ?", req.ProductionLineID)
	}
	if req.CertifiedUserID != "" {
		db = db.Where("`certified_user_id` = ?", req.CertifiedUserID)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.PersonnelQualification
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.PersonnelQualificationsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllPersonnelQualifications() (list []*model.PersonnelQualification, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetPersonnelQualificationByID(id string) (*model.PersonnelQualification, error) {
	m := &model.PersonnelQualification{}
	err := model.DB.DB().Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetPersonnelQualification(req *proto.GetPersonnelQualificationRequest) (*model.PersonnelQualification, error) {
	m := &model.PersonnelQualification{}
	err := model.DB.DB().Preload(clause.Associations).Where("`certified_user_id` = ?", req.CertifiedUserID).First(m).Error
	return m, err
}

func GetPersonnelQualificationByIDs(ids []string) ([]*model.PersonnelQualification, error) {
	var m []*model.PersonnelQualification
	err := model.DB.DB().Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeletePersonnelQualification(id string) (err error) {
	return model.DB.DB().Delete(&model.PersonnelQualification{}, "`id` = ?", id).Error
}
