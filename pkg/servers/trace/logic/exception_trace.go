package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateExceptionTrace(m *model.ExceptionTrace) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateExceptionTrace(m *model.ExceptionTrace) error {
	return model.DB.DB().Omit("created_at", "time_reported").Save(m).Error
}

func QueryExceptionTrace(req *proto.QueryExceptionTraceRequest, resp *proto.QueryExceptionTraceResponse, preload bool) {
	db := model.DB.DB().Model(&model.ExceptionTrace{})
	if req.Message != "" {
		db = db.Where("`host` LIKE ? OR `source` LIKE ? OR `message` LIKE ? OR `stack_trace` LIKE ?", "%"+req.Message+"%", "%"+req.Message+"%", "%"+req.Message+"%", "%"+req.Message+"%")
	}

	if req.TimeReported0 != "" && req.TimeReported1 != "" {
		db = db.Where("`time_reported` BETWEEN ? AND ?", req.TimeReported0, req.TimeReported1)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ExceptionTrace
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ExceptionTracesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllExceptionTraces() (list []*model.ExceptionTrace, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetExceptionTraceByID(id string) (*model.ExceptionTrace, error) {
	m := &model.ExceptionTrace{}
	err := model.DB.DB().Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetExceptionTraceByIDs(ids []string) ([]*model.ExceptionTrace, error) {
	var m []*model.ExceptionTrace
	err := model.DB.DB().Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteExceptionTrace(id string) (err error) {
	return model.DB.DB().Delete(&model.ExceptionTrace{}, "`id` = ?", id).Error
}
