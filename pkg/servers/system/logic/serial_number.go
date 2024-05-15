package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateSerialNumber(m *model.SerialNumber) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateSerialNumber(m *model.SerialNumber) error {
	return model.DB.DB().Omit("create_time").Save(m).Error
}

func QuerySerialNumber(req *proto.QuerySerialNumberRequest, resp *proto.QuerySerialNumberResponse, preload bool) {
	db := model.DB.DB().Model(&model.SerialNumber{})
	if req.Name != "" {
		db.Where("`name` LIKE ? OR `description` LIKE ? OR `prefix` LIKE ?", "%"+req.Name+"%", "%"+req.Name+"%", "%"+req.Name+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "create_time desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.SerialNumber
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.SerialNumbersToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllSerialNumbers() (list []*model.SerialNumber, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetSerialNumberByID(id string) (*model.SerialNumber, error) {
	m := &model.SerialNumber{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetSerialNumberByIDs(ids []string) ([]*model.SerialNumber, error) {
	var m []*model.SerialNumber
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteSerialNumber(id string) (err error) {
	return model.DB.DB().Delete(&model.SerialNumber{}, "id=?", id).Error
}
