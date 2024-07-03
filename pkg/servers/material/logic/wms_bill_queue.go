package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateWMSBillQueue(m *model.WMSBillQueue) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateWMSBillQueue(m *model.WMSBillQueue) error {
	return model.DB.DB().Omit("created_at", "create_time").Save(m).Error
}

func QueryWMSBillQueue(req *proto.QueryWMSBillQueueRequest, resp *proto.QueryWMSBillQueueResponse, preload bool) {
	db := model.DB.DB().Model(&model.WMSBillQueue{}).Preload("ProductOrder")
	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("`create_time` BETWEEN ? AND ?", req.CreateTime0, req.CreateTime1)
	}
	if req.MaterialStore != "" {
		db = db.Joins("JOIN material_stores ON wms_bill_queues.material_store_id=material_stores.id").
			Where("material_stores.code LIKE ? OR material_stores.description LIKE ?", "%"+req.MaterialStore+"%", "%"+req.MaterialStore+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.WMSBillQueue
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.WMSBillQueuesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllWMSBillQueues() (list []*model.WMSBillQueue, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetWMSBillQueueByID(id string) (*model.WMSBillQueue, error) {
	m := &model.WMSBillQueue{}
	err := model.DB.DB().Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetWMSBillQueueByIDs(ids []string) ([]*model.WMSBillQueue, error) {
	var m []*model.WMSBillQueue
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteWMSBillQueue(id string) (err error) {
	return model.DB.DB().Delete(&model.WMSBillQueue{}, "`id` = ?", id).Error
}
