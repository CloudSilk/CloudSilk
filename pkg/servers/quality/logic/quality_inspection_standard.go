package logic

import (
	"errors"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateQualityInspectionStandard(m *model.QualityInspectionStandard) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, " `code`  = ? ", m.Code)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同检验标准")
	}
	return m.ID, nil
}

func UpdateQualityInspectionStandard(m *model.QualityInspectionStandard) error {
	// 空值不覆盖已有外键，避免更新时误清检验类型关联
	omits := []string{"created_at"}
	if m.QualityInspectionTypeID == nil || *m.QualityInspectionTypeID == "" {
		omits = append(omits, "QualityInspectionTypeID")
	}
	duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(model.DB.DB(), m, false, omits, "`id` != ? and  `code`  = ? ", m.ID, m.Code)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同检验标准")
	}

	return nil
}

func QueryQualityInspectionStandard(req *proto.QueryQualityInspectionStandardRequest, resp *proto.QueryQualityInspectionStandardResponse, preload bool) {
	db := model.DB.DB().Model(&model.QualityInspectionStandard{})
	if preload {
		db = db.Preload(clause.Associations)
	}
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}
	if req.QualityInspectionTypeID != "" {
		db = db.Where("`quality_inspection_type_id` = ?", req.QualityInspectionTypeID)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.QualityInspectionStandard
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.QualityInspectionStandardsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllQualityInspectionStandards() (list []*model.QualityInspectionStandard, err error) {
	err = model.DB.DB().Preload(clause.Associations).Find(&list, "`enable` = ?", true).Error
	return
}

func GetQualityInspectionStandardByID(id string) (*model.QualityInspectionStandard, error) {
	m := &model.QualityInspectionStandard{}
	err := model.DB.DB().Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func DeleteQualityInspectionStandard(id string) (err error) {
	var count int64
	if err := model.DB.DB().Model(&model.QualityInspectionOrderItem{}).Where("`quality_inspection_standard_id` = ?", id).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("此检验标准已被检验单引用，无法删除")
	}
	return model.DB.DB().Delete(&model.QualityInspectionStandard{}, "`id` = ?", id).Error
}
