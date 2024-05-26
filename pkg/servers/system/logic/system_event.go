package logic

import (
	"errors"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateSystemEvent(m *model.SystemEvent) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, " code=? ", m.Code)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同系统事件")
	}
	return m.ID, nil
}

func UpdateSystemEvent(m *model.SystemEvent) error {
	duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(model.DB.DB(), m, true, []string{"created_at"}, "id <> ? and code=? ", m.ID, m.Code)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同系统事件")
	}

	return nil
}

func QuerySystemEvent(req *proto.QuerySystemEventRequest, resp *proto.QuerySystemEventResponse, preload bool) {
	db := model.DB.DB().Model(&model.SystemEvent{})

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.SystemEvent
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.SystemEventsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllSystemEvents() (list []*model.SystemEvent, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetSystemEventByID(id string) (*model.SystemEvent, error) {
	m := &model.SystemEvent{}
	err := model.DB.DB().Preload("SystemEventParameters").Preload("SystemEventSubscriptions").Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetSystemEvent(req *proto.GetSystemEventRequest) (*model.SystemEvent, error) {
	m := &model.SystemEvent{}
	err := model.DB.DB().Preload("SystemEventParameters").Preload(clause.Associations).First(m, map[string]interface{}{
		"code":   req.Code,
		"enable": req.Enable,
	}).Error
	return m, err
}

func GetSystemEventByIDs(ids []string) ([]*model.SystemEvent, error) {
	var m []*model.SystemEvent
	err := model.DB.DB().Preload("SystemEventParameters").Preload("SystemEventSubscriptions").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteSystemEvent(id string) (err error) {
	return model.DB.DB().Delete(&model.SystemEvent{}, "id=?", id).Error
}

// func DeleteSystemEvent(id string) (err error) {
// // 检查是否有关联的产品类别特性
// err = model.DB.DB().Where("SystemEventID = ?", id).First(&SystemEventTrigger{}).Error
// switch err {
// case nil:
// 	return fmt.Errorf("数据冲突，此事件已被多次触发，如需删除，请先清空触发记录")
// case gorm.ErrRecordNotFound:
// 	break
// default:
// 	return err
// }

// err = model.DB.DB().Where("SystemEventID = ?", id).First(&ProductPackageType{}).Error
// switch err {
// case nil:
// 	return fmt.Errorf("数据冲突，如需删除，请先清空包装类型管理")
// case gorm.ErrRecordNotFound:
// 	break
// default:
// 	return err
// }

// model.SystemEvent := []interface{}{&SystemEventSubscription{}, &SystemEventParameter{}}

// return model.DB.DB().Transaction(func(tx *gorm.DB) error {
// 	for _, model := range model.SystemEvent {
// 		if err := tx.Delete(model, "SystemEventID=?", id).Error; err != nil {
// 			return err
// 		}
// 	}

// 	if err := tx.Delete(&model.SystemEvent{}, "id=?", id).Error; err != nil {
// 		return err
// 	}

// 	return nil
// })

// 	return nil
// }
