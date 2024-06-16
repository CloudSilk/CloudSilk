package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateMaterialInfo(m *model.MaterialInfo) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateMaterialInfo(m *model.MaterialInfo) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryMaterialInfo(req *proto.QueryMaterialInfoRequest, resp *proto.QueryMaterialInfoResponse, preload bool) {
	db := model.DB.DB().Model(&model.MaterialInfo{}).Preload("MaterialCategory")
	if req.Code != "" {
		db = db.Where("`material_no` LIKE ? OR `material_description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MaterialInfo
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialInfosToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMaterialInfos() (list []*model.MaterialInfo, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMaterialInfoByID(id string) (*model.MaterialInfo, error) {
	m := &model.MaterialInfo{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetMaterialInfoByIDs(ids []string) ([]*model.MaterialInfo, error) {
	var m []*model.MaterialInfo
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMaterialInfo(id string) (err error) {
	return model.DB.DB().Delete(&model.MaterialInfo{}, "`id` = ?", id).Error
}
