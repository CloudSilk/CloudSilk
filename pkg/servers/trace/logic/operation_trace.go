package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateOperationTrace(m *model.OperationTrace) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateOperationTrace(m *model.OperationTrace) error {
	return model.DB.DB().Omit("operate_time").Save(m).Error
}

func QueryOperationTrace(req *proto.QueryOperationTraceRequest, resp *proto.QueryOperationTraceResponse, preload bool) {
	db := model.DB.DB().Model(&model.OperationTrace{})
	if req.OperateTime0 != "" && req.OperateTime1 != "" {
		db = db.Where("`operate_time` BETWEEN ? and ?", req.OperateTime0, req.OperateTime1)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "id")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.OperationTrace
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.OperationTracesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllOperationTraces() (list []*model.OperationTrace, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetOperationTraceByID(id string) (*model.OperationTrace, error) {
	m := &model.OperationTrace{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetOperationTraceByIDs(ids []string) ([]*model.OperationTrace, error) {
	var m []*model.OperationTrace
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteOperationTrace(id string) (err error) {
	return model.DB.DB().Delete(&model.OperationTrace{}, "id=?", id).Error
}
