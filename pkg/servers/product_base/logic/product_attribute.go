package logic

import (
	"errors"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProductAttribute(m *model.ProductAttribute) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, " code =? ", m.Code)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同产品特性")
	}
	return m.ID, nil
}

func UpdateProductAttribute(m *model.ProductAttribute) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.ProductAttributeIdentifier{}, "product_attribute_id=?", m.ID).Error; err != nil {
			return err
		}

		duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(tx, m, true, []string{}, "id <> ? and code =? ", m.ID, m.Code)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同产品特性")
		}

		return nil
	})
}

func QueryProductAttribute(req *proto.QueryProductAttributeRequest, resp *proto.QueryProductAttributeResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductAttribute{})
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductAttribute
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductAttributesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductAttributes() (list []*model.ProductAttribute, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductAttributeByID(id string) (*model.ProductAttribute, error) {
	m := &model.ProductAttribute{}
	err := model.DB.DB().Preload("ProductAttributeIdentifiers").Preload("ProductAttributeIdentifiers.ProductAttributeIdentifierAvailableCategorys").Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductAttributeByIDs(ids []string) ([]*model.ProductAttribute, error) {
	var m []*model.ProductAttribute
	err := model.DB.DB().Preload("ProductAttributeIdentifiers").Preload("ProductAttributeIdentifiers.ProductAttributeIdentifierAvailableCategorys").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductAttribute(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductAttribute{}, "id=?", id).Error
}

// func DeleteProductAttribute(id string) (err error) {
// 检查是否有关联的产品类别特性
// err = model.DB.DB().Where("ProductAttributeID = ?", id).First(&model.ProductCategoryAttribute{}).Error
// switch err {
// case nil:
// 	return fmt.Errorf("数据冲突，请先清空关联此特性的产品类别特性")
// case gorm.ErrRecordNotFound:
// 	break
// default:
// 	return err
// }

// productAttribute := []interface{}{&AttributeExpression{}, &ProductAttributeValuateRule{}, &ProductModelAttributeValue{}, &ProductOrderAttribute{}, &ProductAttributeIdentifier{}}

// return model.DB.DB().Transaction(func(tx *gorm.DB) error {
// 	for _, model := range productAttribute {
// 		if err := tx.Delete(model, "ProductAttributeID=?", id).Error; err != nil {
// 			return err
// 		}
// 	}

// 	if err := tx.Delete(&ProductAttribute{}, "id=?", id).Error; err != nil {
// 		return err
// 	}

// 	return nil
// })
// 	return nil
// }
