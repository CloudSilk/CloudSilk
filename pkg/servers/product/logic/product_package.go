package logic

import (
	"errors"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProductPackage(m *model.ProductPackage) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, " code =? ", m.Code)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同产品包装管理")
	}
	return m.ID, nil
}

func UpdateProductPackage(m *model.ProductPackage) error {
	duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(model.DB.DB(), m, false, []string{}, "id != ? and  code =? ", m.ID, m.Code)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同产品包装管理")
	}

	return nil
}

func QueryProductPackage(req *proto.QueryProductPackageRequest, resp *proto.QueryProductPackageResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductPackage{}).Preload("ProductPackageType").Preload(clause.Associations)
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "code")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductPackage
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductPackagesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductPackages() (list []*model.ProductPackage, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductPackageByID(id string) (*model.ProductPackage, error) {
	m := &model.ProductPackage{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductPackageByIDs(ids []string) ([]*model.ProductPackage, error) {
	var m []*model.ProductPackage
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

// 删除子表
func DeleteProductPackage(id string) (err error) {
	deleteProductPackages := []interface{}{&model.ProductOrderPackage{}, &model.ProductPackageMatchRule{}}

	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		for _, model := range deleteProductPackages {
			if err := tx.Delete(model, "ProductPackageID=?", id).Error; err != nil {
				return err
			}
		}

		return tx.Delete(&model.ProductPackage{}, "id=?", id).Error
	})
}

func UpdateProductPackageAll(m *model.ProductPackage) error {
	duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(model.DB.DB(), m, true, []string{"CreateTime"}, "id != ? and  code =? ", m.ID, m.Code)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同产品包装管理")
	}

	return nil
}
