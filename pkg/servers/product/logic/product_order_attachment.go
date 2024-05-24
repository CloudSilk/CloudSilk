package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductOrderAttachment(m *model.ProductOrderAttachment) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductOrderAttachment(m *model.ProductOrderAttachment) error {
	return model.DB.DB().Omit("create_time").Save(m).Error
}

func QueryProductOrderAttachment(req *proto.QueryProductOrderAttachmentRequest, resp *proto.QueryProductOrderAttachmentResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductOrderAttachment{})

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductOrderAttachment
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductOrderAttachmentsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductOrderAttachments() (list []*model.ProductOrderAttachment, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductOrderAttachmentByID(id string) (*model.ProductOrderAttachment, error) {
	m := &model.ProductOrderAttachment{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetProductOrderAttachmentByIDs(ids []string) ([]*model.ProductOrderAttachment, error) {
	var m []*model.ProductOrderAttachment
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductOrderAttachment(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductOrderAttachment{}, "id=?", id).Error
}
