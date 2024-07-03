package logic

import (
	"fmt"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/types"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateMaterialShelfBin(m *model.MaterialShelfBin) (string, error) {
	m.CurrentState = fmt.Sprintf("%d", types.MaterialShelfStateEmptied)
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateMaterialShelfBin(m *model.MaterialShelfBin) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryMaterialShelfBin(req *proto.QueryMaterialShelfBinRequest, resp *proto.QueryMaterialShelfBinResponse, preload bool) {
	db := model.DB.DB().Model(&model.MaterialShelfBin{}).Preload("MaterialShelf").Preload("MaterialInfo")
	if req.Code != "" {
		db = db.Where("`code` LIKE ?", "%"+req.Code+"%")
	}
	if req.CurrentState != "" {
		db = db.Where("`current_state` = ?", req.CurrentState)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MaterialShelfBin
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialShelfBinsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMaterialShelfBins() (list []*model.MaterialShelfBin, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMaterialShelfBinByID(id string) (*model.MaterialShelfBin, error) {
	m := &model.MaterialShelfBin{}
	err := model.DB.DB().Preload("MaterialContainerType").Preload("MaterialShelf").Preload("MaterialInfo").Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetMaterialShelfBinByIDs(ids []string) ([]*model.MaterialShelfBin, error) {
	var m []*model.MaterialShelfBin
	err := model.DB.DB().Preload("MaterialContainerType").Preload("MaterialShelf").Preload("MaterialInfo").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMaterialShelfBin(id string) (err error) {
	return model.DB.DB().Delete(&model.MaterialShelfBin{}, "`id` = ?", id).Error
}
