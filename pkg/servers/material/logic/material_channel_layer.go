package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateMaterialChannelLayer(m *model.MaterialChannelLayer) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateMaterialChannelLayer(m *model.MaterialChannelLayer) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryMaterialChannelLayer(req *proto.QueryMaterialChannelLayerRequest, resp *proto.QueryMaterialChannelLayerResponse, preload bool) {
	db := model.DB.DB().Model(&model.MaterialChannelLayer{})

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MaterialChannelLayer
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialChannelLayersToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMaterialChannelLayers() (list []*model.MaterialChannelLayer, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMaterialChannelLayerByID(id string) (*model.MaterialChannelLayer, error) {
	m := &model.MaterialChannelLayer{}
	err := model.DB.DB().Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetMaterialChannels(req *proto.GetMaterialChannelRequest) ([]*model.MaterialChannel, error) {
	var m []*model.MaterialChannel
	err := model.DB.DB().
		Preload("MaterialChannelLayer").
		Preload("MaterialInfo").
		Joins("JOIN material_channel_layers ON material_channels.material_channel_layer_id=material_channel_layers.id").
		Where("material_channel_layers.production_station_id=?", req.ProductionStationID).Find(&m).Error

	return m, err
}

func GetMaterialChannelLayerByIDs(ids []string) ([]*model.MaterialChannelLayer, error) {
	var m []*model.MaterialChannelLayer
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMaterialChannelLayer(id string) (err error) {
	return model.DB.DB().Delete(&model.MaterialChannelLayer{}, "`id` = ?", id).Error
}
