package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	apipb "github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateCodingTemplate(m *model.CodingTemplate) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateCodingTemplate(m *model.CodingTemplate) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryCodingTemplate(req *apipb.QueryCodingTemplateRequest, resp *apipb.QueryCodingTemplateResponse, preload bool) {
	db := model.DB.DB().Model(&model.CodingTemplate{})

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.CodingTemplate
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.CodingTemplatesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllCodingTemplates() (list []*model.CodingTemplate, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetCodingTemplateByID(id string) (*model.CodingTemplate, error) {
	m := &model.CodingTemplate{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetCodingTemplateByIDs(ids []string) ([]*model.CodingTemplate, error) {
	var m []*model.CodingTemplate
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteCodingTemplate(id string) (err error) {
	return model.DB.DB().Delete(&model.CodingTemplate{}, "id=?", id).Error
}
