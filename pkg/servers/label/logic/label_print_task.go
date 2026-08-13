package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateLabelPrintTask(m *model.LabelPrintTask) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateLabelPrintTask(m *model.LabelPrintTask) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryLabelPrintTask(req *proto.QueryLabelPrintTaskRequest, resp *proto.QueryLabelPrintTaskResponse, preload bool) {
	db := model.DB.DB().Model(&model.LabelPrintTask{}).Preload("ProductionLine").Preload("ProductCategory").Preload("LabelType").Preload("Printer").Preload("RemoteServiceTask").Preload(clause.Associations)
	if req.ProductionLineID != "" {
		db = db.Where("`production_line_id` = ?", req.ProductionLineID)
	}
	if req.Code != "" {
		db = db.Where("code LIKE ? OR description LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.LabelPrintTask
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.LabelPrintTasksToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllLabelPrintTasks() (list []*model.LabelPrintTask, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetLabelPrintTaskByID(id string) (*model.LabelPrintTask, error) {
	m := &model.LabelPrintTask{}
	err := model.DB.DB().Preload("ProductionLine").Preload("ProductCategory").Preload("LabelType").Preload("Printer").Preload("RemoteServiceTask").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetLabelPrintTaskByIDs(ids []string) ([]*model.LabelPrintTask, error) {
	var m []*model.LabelPrintTask
	err := model.DB.DB().Preload("ProductionLine").Preload("ProductCategory").Preload("LabelType").Preload("Printer").Preload("RemoteServiceTask").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteLabelPrintTask(id string) (err error) {
	return model.DB.DB().Delete(&model.LabelPrintTask{}, "`id` = ?", id).Error
}
