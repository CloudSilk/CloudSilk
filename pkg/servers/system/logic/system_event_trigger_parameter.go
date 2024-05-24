package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateSystemEventTriggerParameter(m *model.SystemEventTriggerParameter) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateSystemEventTriggerParameter(m *model.SystemEventTriggerParameter) error {
	return model.DB.DB().Save(m).Error
}

func QuerySystemEventTriggerParameter(req *proto.QuerySystemEventTriggerParameterRequest, resp *proto.QuerySystemEventTriggerParameterResponse, preload bool) {
	db := model.DB.DB().Model(&model.SystemEventTriggerParameter{}).Preload("SystemEvent").Preload(clause.Associations)

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.SystemEventTriggerParameter
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.SystemEventTriggerParametersToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllSystemEventTriggerParameters() (list []*model.SystemEventTriggerParameter, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetSystemEventTriggerParameterByID(id string) (*model.SystemEventTriggerParameter, error) {
	m := &model.SystemEventTriggerParameter{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetSystemEventTriggerParameterByIDs(ids []string) ([]*model.SystemEventTriggerParameter, error) {
	var m []*model.SystemEventTriggerParameter
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteSystemEventTriggerParameter(id string) (err error) {
	return model.DB.DB().Delete(&model.SystemEventTriggerParameter{}, "id=?", id).Error
}
