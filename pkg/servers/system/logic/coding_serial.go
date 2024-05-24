package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	apipb "github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateCodingSerial(m *model.CodingSerial) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateCodingSerial(m *model.CodingSerial) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryCodingSerial(req *apipb.QueryCodingSerialRequest, resp *apipb.QueryCodingSerialResponse, preload bool) {
	db := model.DB.DB().Model(&model.CodingSerial{})

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.CodingSerial
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.CodingSerialsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllCodingSerials() (list []*model.CodingSerial, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetCodingSerialByID(id string) (*model.CodingSerial, error) {
	m := &model.CodingSerial{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetCodingSerialByIDs(ids []string) ([]*model.CodingSerial, error) {
	var m []*model.CodingSerial
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteCodingSerial(id string) (err error) {
	return model.DB.DB().Delete(&model.CodingSerial{}, "id=?", id).Error
}
