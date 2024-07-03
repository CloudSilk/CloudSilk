package logic

import (
	"fmt"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/types"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateMaterialContainer(m *model.MaterialContainer) (string, error) {
	if err := model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if m.MaterialShelfBinID != nil {
			if err := tx.Model(model.MaterialShelfBin{}).Where("`id` = ?", m.MaterialShelfBinID).Update("current_state", types.MaterialShelfStateFilled).Error; err != nil {
				return err
			}
		}

		m.CurrentState = fmt.Sprintf("%d", types.MaterialShelfStateFilled)
		if err := tx.Create(m).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return "", err
	}

	return m.ID, nil
}

func UpdateMaterialContainer(m *model.MaterialContainer) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if m.MaterialShelfBinID != nil {
			if err := tx.Model(model.MaterialShelfBin{}).Where("`id` = ?", m.MaterialShelfBinID).Update("current_state", types.MaterialShelfStateFilled).Error; err != nil {
				return err
			}
		}

		if err := tx.Omit("created_at").Save(m).Error; err != nil {
			return err
		}

		return nil
	})
}

func QueryMaterialContainer(req *proto.QueryMaterialContainerRequest, resp *proto.QueryMaterialContainerResponse, preload bool) {
	db := model.DB.DB().Model(&model.MaterialContainer{}).Preload("MaterialContainerType").Preload("MaterialShelfBin")
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}
	if req.CurrentState != "" {
		db = db.Where("`current_state` = ?", req.CurrentState)
	}
	if req.MaterialContainerTypeID != "" {
		db = db.Where("`material_container_type_id` = ?", req.MaterialContainerTypeID)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MaterialContainer
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialContainersToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMaterialContainers() (list []*model.MaterialContainer, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMaterialContainerByID(id string) (*model.MaterialContainer, error) {
	m := &model.MaterialContainer{}
	err := model.DB.DB().Preload("MaterialContainerType").Preload("MaterialShelfBin").Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetMaterialContainerByIDs(ids []string) ([]*model.MaterialContainer, error) {
	var m []*model.MaterialContainer
	err := model.DB.DB().Preload("MaterialContainerType").Preload("MaterialShelfBin").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMaterialContainer(id string) (err error) {
	return model.DB.DB().Delete(&model.MaterialContainer{}, "`id` = ?", id).Error
}
