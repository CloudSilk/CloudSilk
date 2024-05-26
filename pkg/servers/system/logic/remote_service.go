package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateRemoteService(m *model.RemoteService) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateRemoteService(m *model.RemoteService) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryRemoteService(req *proto.QueryRemoteServiceRequest, resp *proto.QueryRemoteServiceResponse, preload bool) {
	db := model.DB.DB().Model(&model.RemoteService{})
	if req.Name != "" {
		db.Where("`name` LIKE ? OR `address` LIKE ?", "%s"+req.Name+"%", "%s"+req.Name+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.RemoteService
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.RemoteServicesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllRemoteServices() (list []*model.RemoteService, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetRemoteServiceByID(id string) (*model.RemoteService, error) {
	m := &model.RemoteService{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetRemoteServiceByIDs(ids []string) ([]*model.RemoteService, error) {
	var m []*model.RemoteService
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteRemoteService(id string) (err error) {
	return model.DB.DB().Delete(&model.RemoteService{}, "id=?", id).Error
}
