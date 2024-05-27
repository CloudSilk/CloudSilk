package logic

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	system "github.com/CloudSilk/CloudSilk/pkg/servers/system/logic"
	"github.com/CloudSilk/CloudSilk/pkg/types"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProductOrder(m *model.ProductOrder) (string, error) {
	var count int64
	if err := model.DB.DB().Model(m).Where("`receipt_note_no` = ?", m.ReceiptNoteNo).Count(&count).Error; err != nil {
		return "", err
	}
	if count > 0 {
		return "", errors.New("存在相同入库单号")
	}

	systemConfigKey := types.SystemConfigKeyPrefabricateProductOrderPrefix
	if m.OrderType == types.ProductOrderTypeRoutine {
		systemConfigKey = types.SystemConfigKeyRoutineProductOrderPrefix
	}

	var systemConfig model.SystemParamsConfig
	if err := model.DB.DB().First(&systemConfig, "`key` = ?", systemConfigKey).Error; err == gorm.ErrRecordNotFound {
		return "", fmt.Errorf("缺少系统配置项: %s", systemConfigKey)
	} else if err != nil {
		return "", err
	}

	dateStamp := time.Now().Format("20060102")
	prefix := fmt.Sprintf("%s%s", systemConfig.Value, dateStamp)
	productOrderNo, err := system.GenerateSerialNumber(prefix, dateStamp, prefix, 6, 1)
	if err != nil {
		return "", err
	}
	m.ProductOrderNo = productOrderNo

	//产品工单BOM
	for i, v := range m.ProductOrderBoms {
		v.ItemNo = fmt.Sprintf("%04d", i)
		v.RequireQTY = float32(m.OrderQTY) * v.PieceQTY
		v.CreateUserID = m.CreateUserID
	}

	//产品工单特性
	if len(m.ProductOrderAttributes) > 0 {
		var productModel model.ProductModel
		if err := model.DB.DB().First(&productModel, "`id` = ?", m.ProductModelID).Error; err != nil {
			return "", err
		}
		for _, v := range m.ProductOrderAttributes {
			var productCategoryAttribute model.ProductCategoryAttribute
			if err := model.DB.DB().Preload("ProductAttribute").First(&productCategoryAttribute, "`product_category_id` = ? AND `product_attribute_id` = ?", productModel.ProductCategoryID, v.ProductAttributeID).Error; err != nil {
				return "", err
			}
			if !productCategoryAttribute.AllowNullOrBlank && v.Value == "" {
				return "", fmt.Errorf("产品特性:%s的值不允许为空", productCategoryAttribute.ProductAttribute.Description)
			}

			v.CreateUserID = m.CreateUserID
		}
	}

	//产品工单标签
	for _, v := range m.ProductOrderLabels {
		currentState := types.ProductOrderLabelStateWaitCheck
		if !v.DoubleCheck {
			currentState = types.ProductOrderLabelStateChecked
		}
		v.CurrentState = currentState
		v.CreateUserID = m.CreateUserID
	}

	//产品工单附件
	for _, v := range m.ProductOrderAttachments {
		v.CreateUserID = m.CreateUserID
	}

	m.CurrentState = types.ProductOrderStateReceipted
	m.OrderTime = utils.ParseSqlNullTime(time.Now().Format("2006-01-02 15:04:05"))

	err = model.DB.DB().Create(m).Error

	return m.ID, err
}

func UpdateProductOrder(m *model.ProductOrder) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.ProductOrderBom{}, "`product_order_id` = ?", m.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(&model.ProductInfo{}, "`product_order_id` = ?", m.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(&model.ProductOrderAttribute{}, "`product_order_id` = ?", m.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(&model.ProductOrderProcess{}, "`product_order_id` = ?", m.ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(&model.ProductOrderPackage{}, "`product_order_id` = ?", m.ID).Error; err != nil {
			return err
		}

		duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(tx, m, true, []string{"created_at", "create_time"}, "`id` <> ? and  `receipt_note_no` = ?", m.ID, m.ReceiptNoteNo)
		if err != nil {
			return err
		}
		if duplication {
			return errors.New("存在相同产品工单")
		}

		return nil
	})
}

func QueryProductOrder(req *proto.QueryProductOrderRequest, resp *proto.QueryProductOrderResponse, preload bool) {
	db := model.DB.DB().Model(&model.ProductOrder{})
	if req.ProductOrderNo != "" {
		db = db.Where("`product_order_no` LIKE ?", "%"+req.ProductOrderNo+"%")
	}

	if req.SalesOrderNo != "" {
		db = db.Where("`sales_order_no` LIKE ?", "%"+req.SalesOrderNo+"%")
	}

	if req.ItemNo != "" {
		db = db.Where("`item_no` LIKE ?", "%"+req.ItemNo+"%")
	}

	if req.CreateTime0 != "" && req.CreateTime1 != "" {
		db = db.Where("`create_time` BETWEEN ? and ?", req.CreateTime0, req.CreateTime1)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ProductOrder
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductOrdersToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllProductOrders() (list []*model.ProductOrder, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetProductOrderByID(id string) (*model.ProductOrder, error) {
	m := &model.ProductOrder{}
	err := model.DB.DB().
		Preload("ProductOrderAttachments").
		Preload("ProductOrderBoms").
		Preload("ProductOrderAttributes").
		Preload("ProductInfos").
		Preload("ProductOrderProcesses").
		Preload("ProductOrderLabels").
		Preload("ProductOrderPackages").
		Preload("ProductOrderPallets").
		Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProductOrderByIDs(ids []string) ([]*model.ProductOrder, error) {
	var m []*model.ProductOrder
	err := model.DB.DB().
		Preload("ProductOrderAttachments").
		Preload("ProductOrderBoms").
		Preload("ProductOrderAttributes").
		Preload("ProductInfos").
		Preload("ProductOrderProcesses").
		Preload("ProductOrderLabels").
		Preload("ProductOrderPackages").
		Preload("ProductOrderPallets").
		Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductOrder(id string) (err error) {
	var currentState string
	if err := model.DB.DB().Model(&model.ProductOrder{}).Where("`id` = ?", id).Select("current_state").Scan(&currentState).Error; err != nil {
		return err
	}

	states := []string{types.ProductOrderStateCancelled, types.ProductOrderStateUploaded, types.ProductOrderStateReceipted}
	for _, v := range states {
		if v == currentState {
			return fmt.Errorf("只有工单状态为已上传、已取消时，才可以删除。如需删除已发放的工单，可以尝试撤回发放后删除；如需删除生产中的工单，可以尝试取消生产后删除。")
		}
	}

	return model.DB.DB().Delete(&model.ProductOrder{}, "`id` = ?", id).Error
}

func ReleaseProductOrder(ids []string) (err error) {
	var productOrders []*model.ProductOrder
	if err := model.DB.DB().Find(&productOrders, "`id` in (?)", ids).Error; err != nil {
		return err
	}

	productOrderNoArray := []string{}
	for _, v := range productOrders {
		if v.CurrentState == types.ProductOrderStateDispatched {
			productOrderNoArray = append(productOrderNoArray, v.ProductOrderNo)
		}
	}
	if len(productOrderNoArray) != 0 {
		return fmt.Errorf("下列工单的状态错误，只能发放状态为%s的工单。%s", types.ProductOrderStateDispatched, strings.Join(productOrderNoArray, ","))
	}

	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.ProductOrder{}).Where("`id` in (?)", ids).Update("current_state", types.ProductStateReleased).Error; err != nil {
			return err
		}
		if err := tx.Model(&model.ProductInfo{}).Where("`product_order_id` in (?)", ids).Update("current_state", types.ProductStateReleased).Error; err != nil {
			return err
		}
		return nil
	})
}
