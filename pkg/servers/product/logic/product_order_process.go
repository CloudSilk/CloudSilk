package logic

import (
	"errors"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateProductOrderProcess(m *model.ProductOrderProcess) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, "`product_order_id` = ? and `production_process_id`  = ?", m.ProductOrderID, m.ProductionProcessID)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同工单工序配置")
	}
	return m.ID, nil
}

func UpdateProductOrderProcess(m *model.ProductOrderProcess) error {
	duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(model.DB.DB(), m, false, []string{"created_at", "create_time"}, "`id` <> ? and  `product_order_id` = ? and `production_process_id`  = ?", m.ID, m.ProductOrderID, m.ProductionProcessID)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同工单工序配置")
	}

	return nil
}

func QueryProductOrderProcess(req *proto.QueryProductOrderProcessRequest, resp *proto.QueryProductOrderProcessResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductOrderProcess{}).Preload("ProductionProcess").Preload("ProductOrder").Preload(clause.Associations)
	if req.ProductOrderNo != "" {
		db.Joins("JOIN product_orders ON product_order_processes.product_order_id=product_orders.id").
			Where("product_orders.product_order_no LIKE ?", "%"+req.ProductOrderNo+"%")
	}
	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("product_order_processes.create_time BETWEEN ? and ?", req.CreateTime0, req.CreateTime1)
	}
	if req.ProductOrderID != "" {
		db = db.Where("`product_order_id` = ?", req.ProductOrderID)
	}
	if req.Enable {
		db = db.Where("`enable` = ?", req.Enable)
	}
	if req.SortIndex > 0 {
		db = db.Where("`sort_index`>?", req.SortIndex)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductOrderProcess
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductOrderProcesssToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductOrderProcesss() (list []*model.ProductOrderProcess, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductOrderProcessByID(id string) (*model.ProductOrderProcess, error) {
	m := &model.ProductOrderProcess{}
	err := model.DB.DB().Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProductOrderProcessByIDs(ids []string) ([]*model.ProductOrderProcess, error) {
	var m []*model.ProductOrderProcess
	err := model.DB.DB().Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductOrderProcess(id string) (err error) {
	return model.DB.DB().Delete(&model.ProductOrderProcess{}, "`id` = ?", id).Error
}
