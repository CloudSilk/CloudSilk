package logic

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProductModel(m *model.ProductModel) (string, error) {
	var count int64
	err := model.DB.DB().Model(m).Where(" material_no  = ? ", m.MaterialNo).Count(&count).Error
	if err != nil {
		return "", err
	}
	if count > 0 {
		return "", errors.New("存在相同产品型号")
	}

	//新增型号配置关联同产品类别下的产品类别特性
	var productCategoryAttributes []*model.ProductCategoryAttribute
	if err := model.DB.DB().Where("`product_category_id` = ?", m.ProductCategoryID).Find(&productCategoryAttributes).Error; err != nil {
		return "", err
	}
	var productModelAttributeValues []*model.ProductModelAttributeValue
	for _, _productCategoryAttribute := range productCategoryAttributes {
		productModelAttributeValues = append(productModelAttributeValues, &model.ProductModelAttributeValue{
			ProductAttributeID: _productCategoryAttribute.ProductAttributeID,
			AssignedValue:      _productCategoryAttribute.DefaultValue,
		})
	}
	m.ProductModelAttributeValues = productModelAttributeValues

	err = model.DB.DB().Create(m).Error

	return m.ID, err
}

func DeleteProductModelBoms(tx *gorm.DB, old, m *model.ProductModel) error {
	var deleteIDs []string
	for _, oldObj := range old.ProductModelBoms {
		flag := false
		for _, newObj := range m.ProductModelBoms {
			if newObj.ID == oldObj.ID {
				flag = true
			}
		}
		if !flag {
			deleteIDs = append(deleteIDs, oldObj.ID)
		}
	}

	if len(deleteIDs) > 0 {
		if err := tx.Delete(&model.ProductModelBom{}, "id in ?", deleteIDs).Error; err != nil {
			return err
		}
	}
	return nil
}

func UpdateProductModel(m *model.ProductModel) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		oldProductModel := &model.ProductModel{}
		if err := tx.Preload("ProductModelAttributeValues").Preload("ProductModelBoms").Preload(clause.Associations).Where("id = ?", m.ID).First(oldProductModel).Error; err != nil {
			return err
		}
		if err := tx.Delete(model.ProductModelAttributeValue{}, "product_model_id = ?", m.ID).Error; err != nil {
			return err
		}

		if err := DeleteProductModelBoms(tx, oldProductModel, m); err != nil {
			return err
		}

		// omits := []string{"created_at"}
		// if m.ProductCategoryID == "" {
		// 	omits = append(omits, "ProductCategoryID")
		// }
		duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(tx, m, true, []string{"created_at"}, "id <> ?  AND  material_no  = ? ", m.ID, m.MaterialNo)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同产品型号")
		}
		return nil
	})
}

func QueryProductModel(req *proto.QueryProductModelRequest, resp *proto.QueryProductModelResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductModel{})
	if req.ProductCategoryID != "" {
		db = db.Where("`product_category_id` = ?", req.ProductCategoryID)
	}
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `material_no` LIKE ? OR `material_description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%", "%"+req.Code+"%")
	}
	if req.IsPrefabricated {
		db = db.Where("`is_prefabricated` = ?", req.IsPrefabricated)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductModel
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductModelsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductModels() (list []*model.ProductModel, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductModelByID(id string) (*model.ProductModel, error) {
	m := &model.ProductModel{}
	err := model.DB.DB().
		Preload("ProductModelAttributeValues").
		Preload("ProductModelAttributeValues.ProductAttribute").
		Preload("ProductModelAttributeValues.ProductAttribute.ProductCategoryAttributes").
		Preload("ProductModelBoms").
		Preload(clause.Associations).Where("id = ?", id).First(m).Error

	for _, productModelAttributeValue := range m.ProductModelAttributeValues {
		for _, productCategoryAttribute := range productModelAttributeValue.ProductAttribute.ProductCategoryAttributes {
			if productCategoryAttribute.ProductCategoryID == m.ProductCategoryID {
				productModelAttributeValue.AllowNullOrBlank = productCategoryAttribute.AllowNullOrBlank
			}
		}
	}
	return m, err
}

func GetProductModelByIDs(ids []string) ([]*model.ProductModel, error) {
	var m []*model.ProductModel
	err := model.DB.DB().Preload("ProductModelAttributeValues").
		Preload("ProductModelAttributeValues.ProductAttribute").
		Preload("ProductModelAttributeValues.ProductAttribute.ProductCategoryAttributes").
		Preload("ProductModelBoms").Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductModel(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductModel{}, "`id` = ?", id).Error
}

func ParamProductModelByID(id string) (err error) {
	var productModel model.ProductModel
	if err := model.DB.DB().Preload("ProductModelAttributeValues").Preload("ProductModelBoms").Preload(clause.Associations).First(&productModel, id).Error; err != nil {
		return err
	}
	//物料描述
	materialDescription := productModel.MaterialDescription

	var productCategory model.ProductCategory
	if err := model.DB.DB().First(&productCategory, productModel.ProductCategoryID).Error; err != nil {
		return err
	}
	//正则特性表达式
	attributeExpression := productCategory.AttributeExpression
	if attributeExpression == "" {
		return fmt.Errorf("缺少此产品类别的特性表达式配置，无法解析")
	}

	//TODO 原正则在go语言报错，需替换字符
	attributeExpression = strings.ReplaceAll(attributeExpression, "(?<", "(?P<")
	pattern := regexp.MustCompile(attributeExpression)

	//正则匹配
	if flag := pattern.MatchString(materialDescription); !flag {
		return fmt.Errorf("此产品的型号物料描述不匹配此产品类别的特性表达式的格式，解析失败")
	}

	matches := pattern.FindStringSubmatch(materialDescription)

	productModelAttributeValues := []*model.ProductModelAttributeValue{}
	for i, code := range pattern.SubexpNames() {
		if strings.TrimSpace(code) != "" {
			var productCategoryAttribute model.ProductCategoryAttribute
			if err := model.DB.DB().Joins("JOIN product_attributes ON product_category_attributes.product_attribute_id=product_attributes.id").
				Where("`product_category_id` = ? and `code` = ?", productModel.ProductCategoryID, code).First(&productCategoryAttribute).Error; err != nil {
				continue
			}

			var productModelAttributeValue model.ProductModelAttributeValue
			if err := model.DB.DB().Where("`product_model_id` = ? and `product_attribute_id` = ?", productModel.ID, productCategoryAttribute.ProductAttributeID).First(&productModelAttributeValue).Error; err == gorm.ErrRecordNotFound {
				productModelAttributeValue = model.ProductModelAttributeValue{
					ProductAttributeID: productCategoryAttribute.ProductAttributeID,
					ProductModelID:     productModel.ID,
					AssignedValue:      productCategoryAttribute.DefaultValue,
				}
			}

			if len(strings.TrimSpace(matches[i])) > 0 {
				productModelAttributeValue.AssignedValue = strings.TrimSpace(matches[i])
			}

			productModelAttributeValues = append(productModelAttributeValues, &productModelAttributeValue)
		}
	}

	productModel.ProductModelAttributeValues = productModelAttributeValues
	if err := UpdateProductModel(&productModel); err != nil {
		return err
	}

	return nil
}
