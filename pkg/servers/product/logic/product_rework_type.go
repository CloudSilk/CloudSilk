package logic

import (
	"errors"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductReworkType(m *model.ProductReworkType) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, " code =? ", m.Code)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同产品返工类型")
	}
	return m.ID, nil
}

func UpdateProductReworkType(m *model.ProductReworkType) error {
	duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(model.DB.DB(), m, false, []string{"created_at"}, "id != ? and  code =? ", m.ID, m.Code)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同产品返工类型")
	}

	return nil
}

func QueryProductReworkType(req *proto.QueryProductReworkTypeRequest, resp *proto.QueryProductReworkTypeResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductReworkType{})
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductReworkType
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductReworkTypesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductReworkTypes() (list []*model.ProductReworkType, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductReworkTypeByID(id string) (*model.ProductReworkType, error) {
	m := &model.ProductReworkType{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductReworkTypeByIDs(ids []string) ([]*model.ProductReworkType, error) {
	var m []*model.ProductReworkType
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductReworkType(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductReworkType{}, "id=?", id).Error
}
