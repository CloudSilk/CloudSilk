package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateMaterialTrayBindingRecord(m *model.MaterialTrayBindingRecord) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateMaterialTrayBindingRecord(m *model.MaterialTrayBindingRecord) error {
	return model.DB.DB().Omit("created_at", "create_time").Save(m).Error
}

func QueryMaterialTrayBindingRecord(req *proto.QueryMaterialTrayBindingRecordRequest, resp *proto.QueryMaterialTrayBindingRecordResponse, preload bool) {
	db := model.DB.DB().Model(&model.MaterialTrayBindingRecord{}).Preload("MaterialTray").Preload("ProductInfo").Preload("ProductInfo.ProductOrder")
	if req.ProductionLineID != "" {
		db = db.Joins("JOIN material_trays ON material_tray_binding_records.material_tray_id=material_trays.id").
			Where("material_trays.production_line_id = ?", req.ProductionLineID)
	}
	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("material_tray_binding_records.create_time BETWEEN ? AND ?", req.CreateTime0, req.CreateTime1)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MaterialTrayBindingRecord
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialTrayBindingRecordsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMaterialTrayBindingRecords() (list []*model.MaterialTrayBindingRecord, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMaterialTrayBindingRecordByID(id string) (*model.MaterialTrayBindingRecord, error) {
	m := &model.MaterialTrayBindingRecord{}
	err := model.DB.DB().Preload("MaterialTray").Preload("ProductInfo").Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetMaterialTrayBindingRecord(req *proto.GetMaterialTrayBindingRecordRequest) (*model.MaterialTrayBindingRecord, error) {
	m := &model.MaterialTrayBindingRecord{}
	err := model.DB.DB().
		Preload("MaterialTray").
		Preload("ProductInfo").
		Preload(clause.Associations).
		Where("`product_info_id` = ? AND `material_tray_id` = ?", req.ProductInfoID, req.MaterialTrayID).First(m).Error
	return m, err
}

func GetMaterialTrayBindingRecordByIDs(ids []string) ([]*model.MaterialTrayBindingRecord, error) {
	var m []*model.MaterialTrayBindingRecord
	err := model.DB.DB().Preload("MaterialTray").Preload("ProductInfo").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMaterialTrayBindingRecord(id string) (err error) {
	return model.DB.DB().Delete(&model.MaterialTrayBindingRecord{}, "`id` = ?", id).Error
}
