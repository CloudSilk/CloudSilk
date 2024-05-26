package logic

import (
	"errors"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateTaskQueue(m *model.TaskQueue) (string, error) {
	if m.Enable {
		m.CauseOfState = "处理中"
		m.RunningState = "队列启用"
	}
	duplication, err := model.DB.CreateWithCheckDuplication(m, " name=? ", m.Name)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同任务队列")
	}
	return m.ID, nil
}

func UpdateTaskQueue(m *model.TaskQueue) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.TaskQueueParameter{}, "task_queue_id = ?", m.ID).Error; err != nil {
			return err
		}

		duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(tx, m, true, []string{"created_at"}, "id <> ? and name=? ", m.ID, m.Name)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同任务队列")
		}

		return nil
	})
}

func QueryTaskQueue(req *proto.QueryTaskQueueRequest, resp *proto.QueryTaskQueueResponse, preload bool) {
	db := model.DB.DB().Model(&model.TaskQueue{})

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.TaskQueue
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.TaskQueuesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllTaskQueues() (list []*model.TaskQueue, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetTaskQueueByID(id string) (*model.TaskQueue, error) {
	m := &model.TaskQueue{}
	err := model.DB.DB().Preload("TaskQueueParameters").Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetTaskQueueByIDs(ids []string) ([]*model.TaskQueue, error) {
	var m []*model.TaskQueue
	err := model.DB.DB().Preload("TaskQueueParameters").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteTaskQueue(id string) (err error) {
	return model.DB.DB().Delete(&model.TaskQueue{}, "id=?", id).Error
}
