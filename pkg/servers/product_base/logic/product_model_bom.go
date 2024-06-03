package logic

import (
	"errors"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProductModelBom(m *model.ProductModelBom) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, " `material_no`  = ? ", m.MaterialNo)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同产品型号Bom")
	}
	return m.ID, nil
}

func UpdateProductModelBom(m *model.ProductModelBom) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		oldProductModelBom := &model.ProductModelBom{}
		if err := tx.Preload(clause.Associations).Where("`id` = ?", m.ID).First(oldProductModelBom).Error; err != nil {
			return err
		}
		// omits := []string{}
		// if m.ProductModelID == "" {
		// 	omits = append(omits, "ProductModelID")
		// }
		duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(tx, m, true, []string{"created_at"}, "`id` <> ?  and  `material_no`  = ? ", m.ID, m.MaterialNo)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同产品型号Bom")
		}

		return nil
	})
}

func QueryProductModelBom(req *proto.QueryProductModelBomRequest, resp *proto.QueryProductModelBomResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductModelBom{})
	if req.ProductCategoryID != "" {
		db = db.Joins("JOIN product_models ON product_model_boms.product_model_id=product_models.id").
			Where("product_models.product_category_id = ?", req.ProductCategoryID)
	}
	if req.ProductModelID != "" {
		db = db.Where("`product_model_id`=?", req.ProductModelID)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductModelBom
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductModelBomsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductModelBoms() (list []*model.ProductModelBom, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductModelBomByID(id string) (*model.ProductModelBom, error) {
	m := &model.ProductModelBom{}
	err := model.DB.DB().Preload("ProductModel").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProductModelBomByIDs(ids []string) ([]*model.ProductModelBom, error) {
	var m []*model.ProductModelBom
	err := model.DB.DB().Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductModelBom(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductModelBom{}, "`id` = ?", id).Error
}
