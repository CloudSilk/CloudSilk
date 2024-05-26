package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateInvocationTrace(m *model.InvocationTrace) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateInvocationTrace(m *model.InvocationTrace) error {
	return model.DB.DB().Omit("created_at", "request_time").Save(m).Error
}

func QueryInvocationTrace(req *proto.QueryInvocationTraceRequest, resp *proto.QueryInvocationTraceResponse, preload bool) {
	db := model.DB.DB().Model(&model.InvocationTrace{})
	if req.RequestTime0 != "" && req.RequestTime1 != "" {
		db.Where("`request_time` BETWEEN ? AND ?", req.RequestTime0, req.RequestTime1)
	}
	if req.ActionName != "" {
		db.Where("`action_name` LIKE ?", "%"+req.ActionName+"%")
	}
	if req.IPAddress != "" {
		db.Where("`ip_address` LIKE ?", "%"+req.IPAddress+"%")
	}
	if req.RequestText != "" {
		db.Where("`request_text` LIKE ? or `response_text` LIKE ?", "%"+req.RequestText+"%", "%"+req.RequestText+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.InvocationTrace
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.InvocationTracesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllInvocationTraces() (list []*model.InvocationTrace, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetInvocationTraceByID(id string) (*model.InvocationTrace, error) {
	m := &model.InvocationTrace{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetInvocationTraceByIDs(ids []string) ([]*model.InvocationTrace, error) {
	var m []*model.InvocationTrace
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteInvocationTrace(id string) (err error) {
	return model.DB.DB().Delete(&model.InvocationTrace{}, "id=?", id).Error
}
