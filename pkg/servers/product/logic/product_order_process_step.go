package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProductOrderProcessStep(m *model.ProductOrderProcessStep) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateProductOrderProcessStep(m *model.ProductOrderProcessStep) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		// 删除子表
		if err := tx.Delete(&model.ProductOrderProcessStepAttachment{}, "`product_order_process_step_id` = ?", m.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(&model.ProductOrderProcessStepTypeParameter{}, "`product_order_process_step_id` = ?", m.ID).Error; err != nil {
			return err
		}

		if err := tx.Omit("created_at").Save(m).Error; err != nil {
			return err
		}

		return nil
	})
}

func QueryProductOrderProcessStep(req *proto.QueryProductOrderProcessStepRequest, resp *proto.QueryProductOrderProcessStepResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductOrderProcessStep{}).Preload("ProcessStepType").Preload("ProductOrderProcessStepTypeParameters").Preload(clause.Associations)
	if req.ProductOrderProcessID != "" {
		db = db.Where("`product_order_process_id` = ?", req.ProductOrderProcessID)
	}
	if req.ProductionProcessID != "" || req.ProductOrderID != "" {
		db = db.Joins("JOIN product_order_processes ON product_order_process_steps.product_order_process_id=product_order_processes.id")
		if req.ProductionProcessID != "" {
			db = db.Where("product_order_processes.production_process_id=?", req.ProductionProcessID)
		}
		if req.ProductOrderID != "" {
			db = db.Where("product_order_processes.product_order_id=?", req.ProductOrderID)
		}
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductOrderProcessStep
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductOrderProcessStepsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductOrderProcessSteps() (list []*model.ProductOrderProcessStep, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductOrderProcessStepByID(id string) (*model.ProductOrderProcessStep, error) {
	m := &model.ProductOrderProcessStep{}
	err := model.DB.DB().Preload("ProductOrderProcessStepAttachments").Preload("ProductOrderProcessStepParameters").Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProductOrderProcessStepByIDs(ids []string) ([]*model.ProductOrderProcessStep, error) {
	var m []*model.ProductOrderProcessStep
	err := model.DB.DB().Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductOrderProcessStep(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductOrderProcessStep{}, "`id` = ?", id).Error
}
