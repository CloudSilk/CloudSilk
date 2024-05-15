package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateTaskQueueExecution(m *model.TaskQueueExecution) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateTaskQueueExecution(m *model.TaskQueueExecution) error {
	return model.DB.DB().Omit("create_time").Save(m).Error
}

func QueryTaskQueueExecution(req *proto.QueryTaskQueueExecutionRequest, resp *proto.QueryTaskQueueExecutionResponse, preload bool) {
	db := model.DB.DB().Model(&model.TaskQueueExecution{}).Preload("TaskQueue").Preload(clause.Associations)
	if req.TaskQueueID != "" {
		db = db.Where("`task_queue_id` = ?", req.TaskQueueID)
	}
	if req.DataTrace != "" {
		db = db.Where("`data_trace` LIKE ? OR `failure_reason` LIKE ?", "%"+req.DataTrace+"%", "%"+req.DataTrace+"%")
	}
	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("`create_time` BETWEEN ? and ?", req.CreateTime0, req.CreateTime1)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "id")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.TaskQueueExecution
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.TaskQueueExecutionsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllTaskQueueExecutions() (list []*model.TaskQueueExecution, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetTaskQueueExecutionByID(id string) (*model.TaskQueueExecution, error) {
	m := &model.TaskQueueExecution{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetTaskQueueExecutionByIDs(ids []string) ([]*model.TaskQueueExecution, error) {
	var m []*model.TaskQueueExecution
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteTaskQueueExecution(id string) (err error) {
	return model.DB.DB().Delete(&model.TaskQueueExecution{}, "id=?", id).Error
}
