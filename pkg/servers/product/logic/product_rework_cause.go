package logic

import (
	"errors"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProductReworkCause(m *model.ProductReworkCause) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, " code =? ", m.Code)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同产品返工理由")
	}
	return m.ID, nil
}

func UpdateProductReworkCause(m *model.ProductReworkCause) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.ProductReworkTypePossibleCause{}, "product_rework_cause_id=?", m.ID).Error; err != nil {
			return err
		}

		duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(tx, m, true, []string{}, "id <> ?  and  code =? ", m.ID, m.Code)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同产品返工理由")
		}

		return nil
	})
}

func QueryProductReworkCause(req *proto.QueryProductReworkCauseRequest, resp *proto.QueryProductReworkCauseResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductReworkCause{})
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductReworkCause
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductReworkCausesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductReworkCauses() (list []*model.ProductReworkCause, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductReworkCauseByID(id string) (*model.ProductReworkCause, error) {
	m := &model.ProductReworkCause{}
	err := model.DB.DB().Preload("ProductReworkTypePossibleCauses").Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductReworkCauseByIDs(ids []string) ([]*model.ProductReworkCause, error) {
	var m []*model.ProductReworkCause
	err := model.DB.DB().Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductReworkCause(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductReworkCause{}, "id=?", id).Error
}
