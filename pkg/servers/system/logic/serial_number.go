package logic

import (
	"fmt"
	"math"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateSerialNumber(m *model.SerialNumber) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateSerialNumber(m *model.SerialNumber) error {
	return model.DB.DB().Omit("created_at", "create_time").Save(m).Error
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
	err := model.DB.DB().Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetSerialNumberByIDs(ids []string) ([]*model.SerialNumber, error) {
	var m []*model.SerialNumber
	err := model.DB.DB().Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteSerialNumber(id string) (err error) {
	return model.DB.DB().Delete(&model.SerialNumber{}, "`id` = ?", id).Error
}

func GenerateSerialNumber(name, description, prefix string, length, increment int32) (string, error) {
	serialNumber := &model.SerialNumber{}
	if err := model.DB.DB().Where("`name` = ?", name).First(serialNumber).Error; err == gorm.ErrRecordNotFound {
		serialNumber.Name = name
		serialNumber.Description = description
		serialNumber.Seed = 0
		serialNumber.Increment = increment
		serialNumber.Prefix = prefix
		serialNumber.Length = length
	} else if err != nil {
		return "", err
	}

	serialNumber.Seed += serialNumber.Increment
	seed := serialNumber.Seed
	maximum := math.Pow10(int(serialNumber.Length))
	if float64(seed) > maximum {
		return "", fmt.Errorf("序列超限错误，种子(%d)超出最大限定值(%f)", seed, maximum)
	}

	if err := model.DB.DB().Save(serialNumber).Error; err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%0*d", serialNumber.Prefix, serialNumber.Length, serialNumber.Seed), nil
}
