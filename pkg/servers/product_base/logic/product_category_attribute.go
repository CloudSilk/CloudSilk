package logic

import (
	"errors"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProductCategoryAttribute(m *model.ProductCategoryAttribute) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, " default_value =? ", m.DefaultValue)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同产品类别特性")
	}
	return m.ID, nil
}

// 删除子表
func DeleteProductCategoryAttributeValues(tx *gorm.DB, old, m *model.ProductCategoryAttribute) error {
	var deleteIDs []string
	for _, oldObj := range old.ProductCategoryAttributeValue {
		flag := false
		for _, newObj := range m.ProductCategoryAttributeValue {
			if newObj.ID == oldObj.ID {
				flag = true
			}
		}
		if !flag {
			deleteIDs = append(deleteIDs, oldObj.ID)
		}
	}

	if len(deleteIDs) > 0 {
		if err := tx.Delete(&model.ProductCategoryAttributeValue{}, "id in ?", deleteIDs).Error; err != nil {
			return err
		}
	}
	return nil
}

func UpdateProductCategoryAttribute(m *model.ProductCategoryAttribute) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		oldProductCategoryAttribute := &model.ProductCategoryAttribute{}
		if err := tx.Preload("ProductCategoryAttributeValue").Preload(clause.Associations).Where("id = ?", m.ID).First(oldProductCategoryAttribute).Error; err != nil {
			return err
		}

		if err := DeleteProductCategoryAttributeValues(tx, oldProductCategoryAttribute, m); err != nil {
			return err
		}

		duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(tx, m, true, []string{}, "id <> ?  and  default_value=? ", m.ID, m.DefaultValue)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同产品类别特性")
		}

		return nil
	})
}

func QueryProductCategoryAttribute(req *proto.QueryProductCategoryAttributeRequest, resp *proto.QueryProductCategoryAttributeResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductCategoryAttribute{}).Preload("ProductCategoryAttributeValue").Preload("ProductAttribute").Preload(clause.Associations)
	if req.ProductCategoryID != "" {
		db = db.Where("`product_category_id` = ?", req.ProductCategoryID)
	}
	if req.ProductAttributeID != "" {
		db = db.Where("`product_attribute_id` = ?", req.ProductAttributeID)
	}
	if req.ProductAtribute != "" {
		db = db.Joins("JOIN product_attributes ON product_category_attributes.product_attribute_id=product_attributes.id").
			Where("product_attribute.code LIKE ? OR product_attribute.description LIKE ?", "%"+req.ProductAtribute+"%", "%"+req.ProductAtribute+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "id")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductCategoryAttribute
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductCategoryAttributesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductCategoryAttributes() (list []*model.ProductCategoryAttribute, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductCategoryAttributeByID(id string) (*model.ProductCategoryAttribute, error) {
	m := &model.ProductCategoryAttribute{}
	err := model.DB.DB().Preload("ProductCategoryAttributeValue").Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductCategoryAttributeByIDs(ids []string) ([]*model.ProductCategoryAttribute, error) {
	var m []*model.ProductCategoryAttribute
	err := model.DB.DB().Preload("ProductCategoryAttributeValue").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductCategoryAttribute(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductCategoryAttribute{}, "id=?", id).Error
}

// func DeleteProductCategoryAttribute(id string) (err error) {
// 	models := []interface{}{&model.ProductCategoryAttributeValue{}}

// 	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
// 		for _, model := range models {
// 			if err := tx.Delete(model, "ProductCategoryAttributeID=?", id).Error; err != nil {
// 				return err
// 			}
// 		}

// 		if err := tx.Delete(&model.ProductCategoryAttribute{}, "id=?", id).Error; err != nil {
// 			return err
// 		}

// 		return nil
// 	})
// }
