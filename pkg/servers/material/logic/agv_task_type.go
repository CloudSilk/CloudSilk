package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateAGVTaskType(m *model.AGVTaskType) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateAGVTaskType(m *model.AGVTaskType) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryAGVTaskType(req *proto.QueryAGVTaskTypeRequest, resp *proto.QueryAGVTaskTypeResponse, preload bool) {
	db := model.DB.DB().Model(&model.AGVTaskType{}).Preload("MaterialContainerType")
	if req.ShelfType != 0 {
		db = db.Where("`shelf_type` = ?", req.ShelfType)
	}
	if req.SceneType != 0 {
		db = db.Where("`scene_type` = ?", req.SceneType)
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

	var list []*model.AGVTaskType
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.AGVTaskTypesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllAGVTaskTypes() (list []*model.AGVTaskType, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetAGVTaskTypeByID(id string) (*model.AGVTaskType, error) {
	m := &model.AGVTaskType{}
	err := model.DB.DB().Preload("MaterialContainerType").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetAGVTaskTypeByIDs(ids []string) ([]*model.AGVTaskType, error) {
	var m []*model.AGVTaskType
	err := model.DB.DB().Preload("MaterialContainerType").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteAGVTaskType(id string) (err error) {
	return model.DB.DB().Delete(&model.AGVTaskType{}, "`id` = ?", id).Error
}
