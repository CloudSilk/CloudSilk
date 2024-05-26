package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateInvocationAuthentication(m *model.InvocationAuthentication) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateInvocationAuthentication(m *model.InvocationAuthentication) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryInvocationAuthentication(req *proto.QueryInvocationAuthenticationRequest, resp *proto.QueryInvocationAuthenticationResponse, preload bool) {
	db := model.DB.DB().Model(&model.InvocationAuthentication{})
	if req.Name != "" {
		db.Where("`name` LIKE ?", "%"+req.Name+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.InvocationAuthentication
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.InvocationAuthenticationsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllInvocationAuthentications() (list []*model.InvocationAuthentication, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetInvocationAuthenticationByID(id string) (*model.InvocationAuthentication, error) {
	m := &model.InvocationAuthentication{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetInvocationAuthenticationByIDs(ids []string) ([]*model.InvocationAuthentication, error) {
	var m []*model.InvocationAuthentication
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteInvocationAuthentication(id string) (err error) {
	return model.DB.DB().Delete(&model.InvocationAuthentication{}, "id=?", id).Error
}
