package logic

import (
	"database/sql"
	"errors"
	"fmt"
	"runtime/debug"
	"strings"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	system "github.com/CloudSilk/CloudSilk/pkg/servers/system/logic"
	"github.com/CloudSilk/CloudSilk/pkg/tool"
	"github.com/CloudSilk/CloudSilk/pkg/types"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// 上传
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
	var systemConfigValue string
	if err := model.DB.DB().Model(model.SystemParamsConfig{}).Where("`key` = ?", systemConfigKey).Select("value").Scan(&systemConfigValue).Error; err == gorm.ErrRecordNotFound {
		return "", fmt.Errorf("系统参数配置缺少项: %s", systemConfigKey)
	} else if err != nil {
		return "", err
	}

	dateStamp := time.Now().Format("20060102")
	prefix := fmt.Sprintf("%s%s", systemConfigValue, dateStamp)
	productOrderNo, err := system.GenerateSerialNumber(prefix, "生产工单号", prefix, 6, 1)
	if err != nil {
		return "", err
	}
	m.ProductOrderNo = productOrderNo

	//产品工单BOM
	for i, v := range m.ProductOrderBoms {
		v.ItemNo = fmt.Sprintf("%04d", i+1)
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

	if err := model.DB.DB().Model(model.SystemParamsConfig{}).Where("`key` = ?", "ProductInfoPrefix").Select("value").Scan(&systemConfigValue).Error; err == gorm.ErrRecordNotFound {
		return "", fmt.Errorf("系统参数配置缺少项: ProductInfoPrefix")
	} else if err != nil {
		return "", err
	}
	prefix = fmt.Sprintf("%s%s", systemConfigValue, dateStamp)

	//产品信息
	productInfos := make([]*model.ProductInfo, m.OrderQTY)
	for i := 0; i < int(m.OrderQTY); i++ {
		productSerialNo, err := system.GenerateSerialNumber(prefix, "产品序列号", prefix, 6, 1)
		if err != nil {
			return "", err
		}
		productInfos[i] = &model.ProductInfo{
			CreateUserID:    m.CreateUserID,
			CurrentState:    types.ProductStateReceipted,
			ProductSerialNo: productSerialNo,
		}
	}
	m.ProductInfos = productInfos

	m.CurrentState = types.ProductOrderStateUploaded
	// if false {
	// 	m.CurrentState = types.ProductOrderStateReceipted
	// }
	m.OrderTime = sql.NullTime{Time: time.Now(), Valid: true}

	err = model.DB.DB().Transaction(func(tx *gorm.DB) error {
		//上传
		if err := tx.Create(m).Error; err != nil {
			return err
		}

		//工单接单
		if err := ReceiveProductOrder(tx, m.ID); err != nil {
			if err := model.DB.DB().Create(&model.ExceptionTrace{
				Host:         "/api/mom/product/productorder/add",
				Level:        types.EventTypeError,
				ReportUserID: m.CreateUserID,
				Source:       "工单接单",
				Message:      err.Error(),
				StackTrace:   string(debug.Stack()),
			}).Error; err != nil {
				return err
			}
			return err
		}

		//工单核验
		if err := VerifyProductOrder(tx, m.ID); err != nil {
			if err := model.DB.DB().Create(&model.ExceptionTrace{
				Host:         "/api/mom/product/productorder/add",
				Level:        types.EventTypeError,
				ReportUserID: m.CreateUserID,
				Source:       "工单核验",
				Message:      err.Error(),
				StackTrace:   string(debug.Stack()),
			}).Error; err != nil {
				return err
			}
			return err
		}

		//工单发放
		if err := ReleaseProductOrder(tx, m.ID); err != nil {
			if err := model.DB.DB().Create(&model.ExceptionTrace{
				Host:         "/api/mom/product/productorder/add",
				Level:        types.EventTypeError,
				ReportUserID: m.CreateUserID,
				Source:       "工单发放",
				Message:      err.Error(),
				StackTrace:   string(debug.Stack()),
			}).Error; err != nil {
				return err
			}
			return err
		}

		return nil
	})

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
		Preload("ProductOrderAttributes.ProductAttribute").
		Preload("ProductInfos").
		Preload("ProductOrderProcesses", func(db *gorm.DB) *gorm.DB {
			return db.Order("sort_index")
		}).
		Preload("ProductOrderProcesses.ProductionProcess").
		Preload("ProductOrderLabels").
		Preload("ProductOrderLabels.LabelType").
		Preload("ProductOrderPackages").
		Preload("ProductOrderPallets").
		Preload("ProductionLine").
		Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func GetProductOrderByIDs(ids []string) ([]*model.ProductOrder, error) {
	var m []*model.ProductOrder
	err := model.DB.DB().
		Preload("ProductOrderAttachments").
		Preload("ProductOrderBoms").
		Preload("ProductOrderAttributes").
		Preload("ProductOrderAttributes.ProductAttribute").
		Preload("ProductInfos").
		Preload("ProductOrderProcesses").
		Preload("ProductOrderProcesses.ProductionProcess").
		Preload("ProductOrderLabels").
		Preload("ProductOrderLabels.LabelType").
		Preload("ProductOrderPackages").
		Preload("ProductOrderPallets").
		Preload("ProductionLine").
		Preload(clause.Associations).Where("`id` in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteProductOrder(id string) (err error) {
	var currentState string
	if err := model.DB.DB().Model(&model.ProductOrder{}).Where("`id` = ?", id).Select("current_state").Scan(&currentState).Error; err != nil {
		return err
	}

	if currentState != types.ProductOrderStateCancelled && currentState != types.ProductOrderStateUploaded && currentState != types.ProductOrderStateReceipted {
		return fmt.Errorf("只有工单状态为已上传、已取消、已接单时，才可以删除。如需删除已发放的工单，可以尝试撤回发放后删除；如需删除生产中的工单，可以尝试取消生产后删除。")
	}

	return model.DB.DB().Delete(&model.ProductOrder{}, "`id` = ?", id).Error
}

// 接单
func ReceiveProductOrder(tx *gorm.DB, id string) (err error) {
	productOrder := &model.ProductOrder{}
	if err = tx.Preload("ProductOrderAttributes").First(productOrder, "`id` = ?", id).Error; err != nil {
		return
	}

	if productOrder.CurrentState != types.ProductOrderStateUploaded {
		return fmt.Errorf("工单的状态错误，只能接收状态为%s的工单。", types.ProductOrderStateUploaded)
	}

	var productOrderReleaseRules []*model.ProductOrderReleaseRule
	if err = tx.Order("priority").Find(&productOrderReleaseRules, "`enable` = ?", true).Error; err != nil {
		return
	}

	var productOrderReleaseRule *model.ProductOrderReleaseRule
	for _, _productOrderReleaseRule := range productOrderReleaseRules {
		match := _productOrderReleaseRule.InitialValue
		var attributeExpressions []*model.AttributeExpression
		if err = tx.Find(&attributeExpressions, "`rule_id` = ? AND `rule_type` = ?", _productOrderReleaseRule.ID, "ProductOrderReleaseRule").Error; err != nil {
			return
		}
		for _, attributeExpression := range attributeExpressions {
			match = false
			for _, productOrderAttribute := range productOrder.ProductOrderAttributes {
				if productOrderAttribute.ProductAttributeID == attributeExpression.ProductAttributeID {
					if b, err := tool.MathOperator(productOrderAttribute.Value, attributeExpression.MathOperator, attributeExpression.AttributeValue); b {
						match = true
						break
					} else if err != nil {
						return err
					}
				}
			}
			if match {
				break
			}
		}
		if match {
			productOrderReleaseRule = _productOrderReleaseRule
			break
		}
	}
	if productOrderReleaseRule == nil {
		return fmt.Errorf("接单失败，无法匹配发放规则")
	}

	return tx.Model(productOrder).Where("`id` = ?", id).Updates(map[string]interface{}{
		"production_line_id": &productOrderReleaseRule.ProductionLineID,
		"current_state":      types.ProductOrderStateReceipted,
	}).Error
}

// 核验
func VerifyProductOrder(tx *gorm.DB, id string) (err error) {
	productOrder := &model.ProductOrder{}
	if err = tx.First(productOrder, "`id` = ?", id).Error; err != nil {
		return
	}

	if productOrder.CurrentState != types.ProductOrderStateReceipted {
		return fmt.Errorf("工单的状态错误，只能核验状态为%s的工单。", types.ProductOrderStateReceipted)
	}

	if err = tx.Model(productOrder).Where("`id` = ?", id).Update("current_state", types.ProductOrderStateVerified).Error; err != nil {
		// if err = model.DB.DB().Create(model.TaskQueueExecution{
		// 	Success:       false,
		// 	FailureReason: fmt.Sprintf("%v", err),
		// 	DataTrace:     fmt.Sprintf("数据表: ProductOrder, 索引: %s", productOrder.ProductOrderNo),
		// }).Error; err != nil {
		// 	return
		// }
		return
	}

	return
}

// TODO 签派

// 发放
func ReleaseProductOrder(tx *gorm.DB, id string) (err error) {
	productOrder := &model.ProductOrder{}
	if err = tx.Preload("ProductOrderAttributes").First(productOrder, "`id` = ?", id).Error; err != nil {
		return
	}

	if productOrder.CurrentState != types.ProductOrderStateVerified {
		return fmt.Errorf("工单的状态错误，只能发放状态为%s的工单", types.ProductOrderStateVerified)
	}

	var productionRhythms []*model.ProductionRhythm
	if err = tx.Order("priority").Find(&productionRhythms, "`enable` = ? AND `production_line_id`=?", true, productOrder.ProductionLineID).Error; err != nil {
		return
	}

	var productionRhythm *model.ProductionRhythm
	for _, _productionRhythm := range productionRhythms {
		match := _productionRhythm.InitialValue
		var attributeExpressions []*model.AttributeExpression
		if err = tx.Find(&attributeExpressions, "`rule_id` = ? AND `rule_type` = ?", _productionRhythm.ID, "ProductionRhythm").Error; err != nil {
			return
		}
		for _, attributeExpression := range attributeExpressions {
			match = false
			for _, productOrderAttribute := range productOrder.ProductOrderAttributes {
				if productOrderAttribute.ProductAttributeID == attributeExpression.ProductAttributeID {
					if b, err := tool.MathOperator(productOrderAttribute.Value, attributeExpression.MathOperator, attributeExpression.AttributeValue); b {
						match = true
						break
					} else if err != nil {
						return err
					}
				}
			}
			if match {
				break
			}
		}
		if match {
			productionRhythm = _productionRhythm
			break
		}
	}
	if productionRhythm == nil {
		return fmt.Errorf("发放失败，无法匹配生产节拍")
	}
	var _productionProcesses []*model.ProductionProcess
	if err = tx.Order("sort_index").Find(&_productionProcesses, "`enable` = ? AND `production_line_id`=?", true, productOrder.ProductionLineID).Error; err != nil {
		return
	}
	var productionProcesses []*model.ProductionProcess
	for _, _productionProcess := range _productionProcesses {
		match := _productionProcess.InitialValue
		var attributeExpressions []*model.AttributeExpression
		if err = tx.Find(&attributeExpressions, "`rule_id` = ? AND `rule_type` = ?", _productionProcess.ID, "ProductionProcess").Error; err != nil {
			return
		}
		for _, attributeExpression := range attributeExpressions {
			match = false
			for _, productOrderAttribute := range productOrder.ProductOrderAttributes {
				if productOrderAttribute.ProductAttributeID == attributeExpression.ProductAttributeID {
					if b, err := tool.MathOperator(productOrderAttribute.Value, attributeExpression.MathOperator, attributeExpression.AttributeValue); b {
						match = true
						break
					} else if err != nil {
						return err
					}
				}
			}
			if match {
				break
			}
		}
		if match {
			productionProcesses = append(productionProcesses, _productionProcess)
		}
	}
	var ProductOrderProcesses []*model.ProductOrderProcess
	for _, productionProcess := range productionProcesses {
		ProductOrderProcesses = append(ProductOrderProcesses, &model.ProductOrderProcess{
			CreateUserID:        productOrder.CreateUserID,
			SortIndex:           productionProcess.SortIndex,
			ProductionProcessID: productionProcess.ID,
			Enable:              true,
		})
	}
	productOrder.ProductOrderProcesses = ProductOrderProcesses

	productOrder.StandardWorkTime = productionRhythm.StandardTime
	productOrder.ReleaseTime = sql.NullTime{Time: time.Now(), Valid: true}
	productOrder.CurrentState = types.ProductOrderStateReleased

	propertyBrief := ""
	for _, productOrderAttribute := range productOrder.ProductOrderAttributes {
		if productOrderAttribute.Value == "" {
			productOrderAttribute.Value = "NA"
		}
		propertyBrief += productOrderAttribute.Value + "\\"
	}
	productOrder.PropertyBrief = strings.TrimSuffix(propertyBrief, "\\")

	if err = tx.Save(productOrder).Error; err != nil {
		return
	}

	if err = tx.Model(&model.ProductInfo{}).Where("`product_order_id`=? AND `current_state`=?", productOrder.ID, types.ProductStateReceipted).Updates(map[string]interface{}{
		"release_time":  time.Now(),
		"current_state": types.ProductStateReleased,
	}).Error; err != nil {
		return
	}

	return nil
}

// 取消
func CancelProductOrder(ids []string) (err error) {
	var productOrders []*model.ProductOrder
	if err := model.DB.DB().Find(&productOrders, "`id` in (?)", ids).Error; err != nil {
		return err
	}

	productOrderNoArray := []string{}
	for _, v := range productOrders {
		if v.CurrentState != types.ProductOrderStateDispatched && v.CurrentState != types.ProductOrderStateReleased && v.CurrentState != types.ProductOrderStateProducting {
			productOrderNoArray = append(productOrderNoArray, v.ProductOrderNo)
		}
	}
	if len(productOrderNoArray) != 0 {
		return fmt.Errorf("下列工单的状态错误，只能取消状态为%s或%s或%s的工单。%s", types.ProductOrderStateDispatched, types.ProductOrderStateReleased, types.ProductOrderStateProducting, strings.Join(productOrderNoArray, ","))
	}

	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.ProductOrder{}).Where("`id` in (?)", ids).Update("current_state", types.ProductOrderStateCancelled).Error; err != nil {
			return err
		}
		if err := tx.Model(&model.ProductInfo{}).Where("`product_order_id` in (?) AND `current_state` = ?", ids, types.ProductStateReleased).Update("current_state", types.ProductStateCancelled).Error; err != nil {
			return err
		}
		return nil
	})
}

// 暂缓
func SuspendProductOrder(ids []string) (err error) {
	var productOrders []*model.ProductOrder
	if err := model.DB.DB().Find(&productOrders, "`id` in (?)", ids).Error; err != nil {
		return err
	}

	productOrderNoArray := []string{}
	for _, v := range productOrders {
		if v.CurrentState != types.ProductOrderStateReleased && v.CurrentState != types.ProductOrderStateProducting {
			productOrderNoArray = append(productOrderNoArray, v.ProductOrderNo)
		}
	}
	if len(productOrderNoArray) != 0 {
		return fmt.Errorf("下列工单的状态错误，只能暂缓状态为%s或%s的工单。%s", types.ProductOrderStateReleased, types.ProductOrderStateProducting, strings.Join(productOrderNoArray, ","))
	}

	return model.DB.DB().Model(&model.ProductOrder{}).Where("`id` in (?)", ids).Update("current_state", types.ProductOrderStateSuspended).Error
}

// 恢复
func ResumeProductOrder(ids []string) (err error) {
	var productOrders []*model.ProductOrder
	if err := model.DB.DB().Find(&productOrders, "`id` in (?)", ids).Error; err != nil {
		return err
	}

	productOrderNoArray := []string{}
	for _, v := range productOrders {
		if v.CurrentState != types.ProductOrderStateSuspended {
			productOrderNoArray = append(productOrderNoArray, v.ProductOrderNo)
		}
	}
	if len(productOrderNoArray) != 0 {
		return fmt.Errorf("下列工单的状态错误，只能恢复状态为%s的工单。%s", types.ProductOrderStateSuspended, strings.Join(productOrderNoArray, ","))
	}

	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.ProductOrder{}).Where("`id` in (?) AND `started_qty` > 0", ids).Update("current_state", types.ProductOrderStateProducting).Error; err != nil {
			return err
		}
		if err := tx.Model(&model.ProductOrder{}).Where("`id` in (?) AND `started_qty` = 0", ids).Update("current_state", types.ProductOrderStateReleased).Error; err != nil {
			return err
		}
		if err := tx.Model(&model.ProductInfo{}).Where("`product_order_id` in (?) AND current_state=?", ids, types.ProductStateSuspended).Update("current_state", types.ProductStateReleased).Error; err != nil {
			return err
		}
		return nil
	})
}
