package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateRemoteServiceTaskQueue(m *model.RemoteServiceTaskQueue) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateRemoteServiceTaskQueue(m *model.RemoteServiceTaskQueue) error {
	return model.DB.DB().Omit("created_at", "create_time").Save(m).Error
}

func QueryRemoteServiceTaskQueue(req *proto.QueryRemoteServiceTaskQueueRequest, resp *proto.QueryRemoteServiceTaskQueueResponse, preload bool) {
	db := model.DB.DB().Model(&model.RemoteServiceTaskQueue{}).Preload("RemoteServiceTask").Preload("RemoteServiceTask.RemoteService").Preload(clause.Associations)
	if req.TaskNo != "" {
		db.Where("`task_no` LIKE ? OR `request_text` LIKE ? OR `response_text` LIKE ?", "%"+req.TaskNo+"%", "%"+req.TaskNo+"%", "%"+req.TaskNo+"%")
	}
	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("`create_time` BETWEEN ? AND ?", req.CreateTime0, req.CreateTime1)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.RemoteServiceTaskQueue
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.RemoteServiceTaskQueuesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllRemoteServiceTaskQueues() (list []*model.RemoteServiceTaskQueue, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetRemoteServiceTaskQueueByID(id string) (*model.RemoteServiceTaskQueue, error) {
	m := &model.RemoteServiceTaskQueue{}
	err := model.DB.DB().Preload("RemoteServiceTaskQueueParameters").Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetRemoteServiceTaskQueueByIDs(ids []string) ([]*model.RemoteServiceTaskQueue, error) {
	var m []*model.RemoteServiceTaskQueue
	err := model.DB.DB().Preload("RemoteServiceTaskQueueParameters").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteRemoteServiceTaskQueue(id string) (err error) {
	return model.DB.DB().Delete(&model.RemoteServiceTaskQueue{}, "id=?", id).Error
}
