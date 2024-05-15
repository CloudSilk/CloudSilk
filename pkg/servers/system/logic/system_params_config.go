package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateSystemParamsConfig(m *model.SystemParamsConfig) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateSystemParamsConfig(m *model.SystemParamsConfig) error {
	return model.DB.DB().Save(m).Error
}

func QuerySystemParamsConfig(req *proto.QuerySystemParamsConfigRequest, resp *proto.QuerySystemParamsConfigResponse, preload bool) {
	db := model.DB.DB().Model(&model.SystemParamsConfig{})
	if req.Key != "" {
		db = db.Where("`key` LIKE ?", "%"+req.Key+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "id")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.SystemParamsConfig
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.SystemParamsConfigsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllSystemParamsConfigs() (list []*model.SystemParamsConfig, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetSystemParamsConfigByID(id string) (*model.SystemParamsConfig, error) {
	m := &model.SystemParamsConfig{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetSystemParamsConfigByIDs(ids []string) ([]*model.SystemParamsConfig, error) {
	var m []*model.SystemParamsConfig
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteSystemParamsConfig(id string) (err error) {
	return model.DB.DB().Delete(&model.SystemParamsConfig{}, "id=?", id).Error
}
