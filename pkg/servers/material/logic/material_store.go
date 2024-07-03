package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateMaterialStore(m *model.MaterialStore) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateMaterialStore(m *model.MaterialStore) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Delete(&model.MaterialStoreAvailableLine{}, "`material_store_id` = ?", m.ID).Error; err != nil {
			return err
		}

		if err := tx.Omit("created_at").Save(m).Error; err != nil {
			return err
		}

		return nil
	})
}

func QueryMaterialStore(req *proto.QueryMaterialStoreRequest, resp *proto.QueryMaterialStoreResponse, preload bool) {
	db := model.DB.DB().Model(&model.MaterialStore{})
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MaterialStore
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialStoresToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMaterialStores() (list []*model.MaterialStore, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMaterialStoreByID(id string) (*model.MaterialStore, error) {
	m := &model.MaterialStore{}
	err := model.DB.DB().Preload("ProductionLines").Preload("ProductionLines.ProductionLine").Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetMaterialStoreByIDs(ids []string) ([]*model.MaterialStore, error) {
	var m []*model.MaterialStore
	err := model.DB.DB().Preload("ProductionLines").Preload("ProductionLines.ProductionLine").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMaterialStore(id string) (err error) {
	return model.DB.DB().Delete(&model.MaterialStore{}, "`id` = ?", id).Error
}
