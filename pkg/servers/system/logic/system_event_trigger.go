package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateSystemEventTrigger(m *model.SystemEventTrigger) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateSystemEventTrigger(m *model.SystemEventTrigger) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QuerySystemEventTrigger(req *proto.QuerySystemEventTriggerRequest, resp *proto.QuerySystemEventTriggerResponse, preload bool) {
	db := model.DB.DB().Model(&model.SystemEventTrigger{}).Preload("SystemEvent").Preload(clause.Associations)
	if req.EventNo != "" {
		db = db.Where("`event_no` = ?", req.EventNo)
	}
	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("`create_time` BETWEEN ? AND ?", req.CreateTime0, req.CreateTime1)
	}
	if req.SystemEvent != "" {
		db = db.Joins("JOIN system_events ON system_event_triggers.system_event_id = system_events.id").
			Where("system_events.description = ?", req.SystemEvent)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.SystemEventTrigger
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.SystemEventTriggersToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllSystemEventTriggers() (list []*model.SystemEventTrigger, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetSystemEventTriggerByID(id string) (*model.SystemEventTrigger, error) {
	m := &model.SystemEventTrigger{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetSystemEventTriggerByIDs(ids []string) ([]*model.SystemEventTrigger, error) {
	var m []*model.SystemEventTrigger
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteSystemEventTrigger(id string) (err error) {
	return model.DB.DB().Delete(&model.SystemEventTrigger{}, "id=?", id).Error
}
