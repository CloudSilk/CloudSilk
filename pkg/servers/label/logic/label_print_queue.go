package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/types"
	"github.com/CloudSilk/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateLabelPrintQueue(m *model.LabelPrintQueue) (string, error) {
	m.TaskNo = uuid.NewString()
	m.CurrentState = types.LabelPrintQueueStateWaitPrint
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateLabelPrintQueue(m *model.LabelPrintQueue) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.LabelPrintQueueParameter{}, "`label_print_queue_id` = ?", m.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(&model.LabelPrintQueueExecution{}, "`label_print_queue_id` = ?", m.ID).Error; err != nil {
			return err
		}

		if err := tx.Omit("created_at", "create_time").Save(m).Error; err != nil {
			return err
		}

		return nil
	})
}

func QueryLabelPrintQueue(req *proto.QueryLabelPrintQueueRequest, resp *proto.QueryLabelPrintQueueResponse, preload bool) {
	db := model.DB.DB().Model(&model.LabelPrintQueue{}).Preload("Printer").Preload("Printer.PrintServer").Preload("RemoteServiceTask").Preload(clause.Associations)
	if req.TaskNo != "" {
		db.Where("task_no LIKE ?", "%"+req.TaskNo+"%")
	}
	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("create_time BETWEEN ? AND ?", req.CreateTime0, req.CreateTime1)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.LabelPrintQueue
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.LabelPrintQueuesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllLabelPrintQueues() (list []*model.LabelPrintQueue, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetLabelPrintQueueByID(id string) (*model.LabelPrintQueue, error) {
	m := &model.LabelPrintQueue{}
	err := model.DB.DB().Preload("LabelParameters").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetLabelPrintQueueByIDs(ids []string) ([]*model.LabelPrintQueue, error) {
	var m []*model.LabelPrintQueue
	err := model.DB.DB().Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteLabelPrintQueue(id string) (err error) {
	return model.DB.DB().Delete(&model.LabelPrintQueue{}, "`id` = ?", id).Error
}
