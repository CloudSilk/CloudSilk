package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateMaterialSupplier(m *model.MaterialSupplier) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateMaterialSupplier(m *model.MaterialSupplier) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryMaterialSupplier(req *proto.QueryMaterialSupplierRequest, resp *proto.QueryMaterialSupplierResponse, preload bool) {
	db := model.DB.DB().Model(&model.MaterialSupplier{})

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MaterialSupplier
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialSuppliersToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMaterialSuppliers() (list []*model.MaterialSupplier, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMaterialSupplierByID(id string) (*model.MaterialSupplier, error) {
	m := &model.MaterialSupplier{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetMaterialSupplierByIDs(ids []string) ([]*model.MaterialSupplier, error) {
	var m []*model.MaterialSupplier
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMaterialSupplier(id string) (err error) {
	return model.DB.DB().Delete(&model.MaterialSupplier{}, "id=?", id).Error
}
