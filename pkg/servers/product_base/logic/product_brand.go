package logic

import (
	"errors"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductBrand(m *model.ProductBrand) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, " code =? ", m.Code)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同产品品牌")
	}
	return m.ID, nil
}

func UpdateProductBrand(m *model.ProductBrand) error {
	var count int64
	if model.DB.DB().Model(&model.ProductBrand{}).Where("id <> ? and code = ?", m.ID, m.Code).Limit(1).Count(&count); count > 0 {
		return errors.New("存在相同尺码")
	}

	return model.DB.DB().Save(m).Error
}

func QueryProductBrand(req *proto.QueryProductBrandRequest, resp *proto.QueryProductBrandResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductBrand{})
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `description` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "id")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductBrand
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductBrandsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductBrands() (list []*model.ProductBrand, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductBrandByID(id string) (*model.ProductBrand, error) {
	m := &model.ProductBrand{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductBrandByIDs(ids []string) ([]*model.ProductBrand, error) {
	var m []*model.ProductBrand
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductBrand(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductBrand{}, "id=?", id).Error
}
