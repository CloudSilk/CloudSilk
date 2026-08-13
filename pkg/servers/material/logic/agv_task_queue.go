package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/types"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateAGVTaskQueue(m *model.AGVTaskQueue) (string, error) {
	m.CurrentState = types.AGVTaskQueueStateWaitDispatch
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateAGVTaskQueue(m *model.AGVTaskQueue) error {
	return model.DB.DB().Omit("created_at", "create_time").Save(m).Error
}

func QueryAGVTaskQueue(req *proto.QueryAGVTaskQueueRequest, resp *proto.QueryAGVTaskQueueResponse, preload bool) {
	db := model.DB.DB().Model(&model.AGVTaskQueue{}).Preload("AGVTaskType").Preload("Departure").Preload("Destination")
	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("`create_time` BETWEEN ? AND ?", req.CreateTime0, req.CreateTime1)
	}
	if req.TaskNo != "" {
		db = db.Where("`task_no` = ?", req.TaskNo)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.AGVTaskQueue
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.AGVTaskQueuesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllAGVTaskQueues() (list []*model.AGVTaskQueue, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetAGVTaskQueueByID(id string) (*model.AGVTaskQueue, error) {
	m := &model.AGVTaskQueue{}
	err := model.DB.DB().Preload("Departure").Preload("Destination").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetAGVTaskQueueByIDs(ids []string) ([]*model.AGVTaskQueue, error) {
	var m []*model.AGVTaskQueue
	err := model.DB.DB().Preload("Departure").Preload("Destination").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteAGVTaskQueue(id string) (err error) {
	return model.DB.DB().Delete(&model.AGVTaskQueue{}, "`id` = ?", id).Error
}
