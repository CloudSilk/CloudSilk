package logic

import (
	"errors"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateRemoteServiceTask(m *model.RemoteServiceTask) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, " code=? ", m.Code)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同远程任务管理")
	}
	return m.ID, nil
}

func UpdateRemoteServiceTask(m *model.RemoteServiceTask) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.RemoteServiceTaskParameter{}, "remote_service_task_id = ?", m.ID).Error; err != nil {
			return err
		}

		duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(tx, m, true, []string{}, "id <> ? and code=? ", m.ID, m.Code)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同远程任务管理")
		}

		return nil
	})
}

func QueryRemoteServiceTask(req *proto.QueryRemoteServiceTaskRequest, resp *proto.QueryRemoteServiceTaskResponse, preload bool) {
	db := model.DB.DB().Model(&model.RemoteServiceTask{}).Preload("RemoteService").Preload(clause.Associations)

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.RemoteServiceTask
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.RemoteServiceTasksToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllRemoteServiceTasks() (list []*model.RemoteServiceTask, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetRemoteServiceTaskByID(id string) (*model.RemoteServiceTask, error) {
	m := &model.RemoteServiceTask{}
	err := model.DB.DB().Preload("RemoteServiceTaskParameters").Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetRemoteServiceTaskByIDs(ids []string) ([]*model.RemoteServiceTask, error) {
	var m []*model.RemoteServiceTask
	err := model.DB.DB().Preload("RemoteServiceTaskParameters").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteRemoteServiceTask(id string) (err error) {
	return model.DB.DB().Delete(&model.RemoteServiceTask{}, "id=?", id).Error
}
