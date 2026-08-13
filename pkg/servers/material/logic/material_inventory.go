package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateMaterialInventory(m *model.MaterialInventory) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateMaterialInventory(m *model.MaterialInventory) error {
	return model.DB.DB().Omit("created_at", "create_time").Save(m).Error
}

func QueryMaterialInventory(req *proto.QueryMaterialInventoryRequest, resp *proto.QueryMaterialInventoryResponse, preload bool) {
	db := model.DB.DB().Model(&model.MaterialInventory{}).Preload("MaterialInfo").Preload("MaterialStore")
	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("material_inventories.create_time BETWEEN ? AND ?", req.CreateTime0, req.CreateTime1)
	}
	if req.MaterialInfo != "" {
		db = db.Joins("JOIN material_infos ON material_inventories.material_info_id=material_infos.id").
			Where("material_infos.material_no LIKE ? OR material_infos.material_description LIKE ?", "%"+req.MaterialInfo+"%", "%"+req.MaterialInfo+"%")
	}
	if req.MaterialStore != "" {
		db = db.Joins("JOIN material_stores ON material_inventories.material_store_id=material_stores.id").
			Where("material_stores.code LIKE ? OR material_infos.description LIKE ?", "%"+req.MaterialStore+"%", "%"+req.MaterialStore+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MaterialInventory
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialInventorysToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMaterialInventorys() (list []*model.MaterialInventory, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMaterialInventoryByID(id string) (*model.MaterialInventory, error) {
	m := &model.MaterialInventory{}
	err := model.DB.DB().Preload("MaterialInfo").Preload("MaterialStore").Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetMaterialInventoryByIDs(ids []string) ([]*model.MaterialInventory, error) {
	var m []*model.MaterialInventory
	err := model.DB.DB().Preload("MaterialInfo").Preload("MaterialStore").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMaterialInventory(id string) (err error) {
	return model.DB.DB().Delete(&model.MaterialInventory{}, "`id` = ?", id).Error
}
