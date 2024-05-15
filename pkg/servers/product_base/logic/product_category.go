package logic

import (
	"errors"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductCategory(m *model.ProductCategory) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, " code =? ", m.Code)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同产品类别")
	}
	return m.ID, nil
}

func UpdateProductCategory(m *model.ProductCategory) error {
	var count int64
	if model.DB.DB().Model(&model.ProductCategory{}).Where("id <> ? and code = ?", m.ID, m.Code).Limit(1).Count(&count); count > 0 {
		return errors.New("存在相同产品类别")
	}

	return model.DB.DB().Save(m).Error
}

func QueryProductCategory(req *proto.QueryProductCategoryRequest, resp *proto.QueryProductCategoryResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductCategory{}).Preload("ProductBrand").Preload(clause.Associations)
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}
	if req.ProductBrandID != "" {
		db = db.Where("`product_brand_id` = ?", req.ProductBrandID)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "id")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductCategory
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductCategorysToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductCategorys() (list []*model.ProductCategory, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductCategoryByID(id string) (*model.ProductCategory, error) {
	m := &model.ProductCategory{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductCategoryByIDs(ids []string) ([]*model.ProductCategory, error) {
	var m []*model.ProductCategory
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductCategory(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductCategory{}, "id=?", id).Error
}
