package logic

import (
	"errors"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateQualityInspectionType(m *model.QualityInspectionType) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, " `code`  = ? ", m.Code)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同检验类型")
	}
	return m.ID, nil
}

func UpdateQualityInspectionType(m *model.QualityInspectionType) error {
	duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(model.DB.DB(), m, false, []string{"created_at"}, "`id` != ? and  `code`  = ? ", m.ID, m.Code)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同检验类型")
	}

	return nil
}

func QueryQualityInspectionType(req *proto.QueryQualityInspectionTypeRequest, resp *proto.QueryQualityInspectionTypeResponse, preload bool) {
	db := model.DB.DB().Model(&model.QualityInspectionType{})
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.QualityInspectionType
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.QualityInspectionTypesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllQualityInspectionTypes() (list []*model.QualityInspectionType, err error) {
	err = model.DB.DB().Find(&list, "`enable` = ?", true).Error
	return
}

func GetQualityInspectionTypeByID(id string) (*model.QualityInspectionType, error) {
	m := &model.QualityInspectionType{}
	err := model.DB.DB().Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func DeleteQualityInspectionType(id string) (err error) {
	var count int64
	if err := model.DB.DB().Model(&model.QualityInspectionStandard{}).Where("`quality_inspection_type_id` = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("此检验类型下存在检验标准，无法删除")
	}
	return model.DB.DB().Delete(&model.QualityInspectionType{}, "`id` = ?", id).Error
}
