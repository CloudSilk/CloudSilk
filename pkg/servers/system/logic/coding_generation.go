package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	apipb "github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateCodingGeneration(m *model.CodingGeneration) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateCodingGeneration(m *model.CodingGeneration) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryCodingGeneration(req *apipb.QueryCodingGenerationRequest, resp *apipb.QueryCodingGenerationResponse, preload bool) {
	db := model.DB.DB().Model(&model.CodingGeneration{})

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "id")
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.CodingGeneration
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.CodingGenerationsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllCodingGenerations() (list []*model.CodingGeneration, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetCodingGenerationByID(id string) (*model.CodingGeneration, error) {
	m := &model.CodingGeneration{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetCodingGenerationByIDs(ids []string) ([]*model.CodingGeneration, error) {
	var m []*model.CodingGeneration
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteCodingGeneration(id string) (err error) {
	return model.DB.DB().Delete(&model.CodingGeneration{}, "id=?", id).Error
}
