package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateLabelTemplate(m *model.LabelTemplate) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateLabelTemplate(m *model.LabelTemplate) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.LabelParameter{}, "label_template_id=?", m.ID).Error; err != nil {
			return err
		}

		if err := tx.Omit("created_at").Save(m).Error; err != nil {
			return err
		}

		return nil
	})

}

func QueryLabelTemplate(req *proto.QueryLabelTemplateRequest, resp *proto.QueryLabelTemplateResponse, preload bool) {
	db := model.DB.DB().Model(&model.LabelTemplate{}).Preload("LabelType").Preload(clause.Associations)
	if req.Code != "" {
		db.Where("code LIKE ? OR description LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}
	if req.LabelTypeID != "" {
		db.Where("label_type_id = ?", req.LabelTypeID)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.LabelTemplate
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.LabelTemplatesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllLabelTemplates() (list []*model.LabelTemplate, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetLabelTemplateByID(id string) (*model.LabelTemplate, error) {
	m := &model.LabelTemplate{}
	err := model.DB.DB().Preload("LabelParameters").Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetLabelTemplateByIDs(ids []string) ([]*model.LabelTemplate, error) {
	var m []*model.LabelTemplate
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteLabelTemplate(id string) (err error) {
	return model.DB.DB().Delete(&model.LabelTemplate{}, "id=?", id).Error
}
