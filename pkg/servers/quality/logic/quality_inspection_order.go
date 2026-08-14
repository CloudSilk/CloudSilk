package logic

import (
	"errors"
	"fmt"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	product "github.com/CloudSilk/CloudSilk/pkg/servers/product/logic"
	system "github.com/CloudSilk/CloudSilk/pkg/servers/system/logic"
	"github.com/CloudSilk/CloudSilk/pkg/tool"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	// 检验单状态
	QualityInspectionStateWait     = "待检验"
	QualityInspectionStateDoing    = "检验中"
	QualityInspectionStateFinished = "已完成"
	// 检验结论
	QualityConclusionQualified   = "合格"
	QualityConclusionUnqualified = "不合格"
	QualityConclusionConcession  = "让步接收"
	// 默认比较方式（规格上下限之间的范围内判定）
	CriterionWithinRange = "范围内"
)

// CreateQualityInspectionOrder 创建检验单
// 按检验类型下的启用检验标准生成明细，并按 GB/T 2828 一般检验水平II 简化规则确定抽样数量
func CreateQualityInspectionOrder(m *model.QualityInspectionOrder) (string, error) {
	if m.QualityInspectionTypeID == nil || *m.QualityInspectionTypeID == "" {
		return "", errors.New("检验类型不能为空")
	}
	if m.LotQTY <= 0 {
		return "", errors.New("批次数量必须大于0")
	}

	//按检验类型加载启用的检验标准
	standards := []*model.QualityInspectionStandard{}
	if err := model.DB.DB().Where("`quality_inspection_type_id` = ? AND `enable` = ?", *m.QualityInspectionTypeID, true).Find(&standards).Error; err != nil {
		return "", err
	}
	if len(standards) == 0 {
		return "", errors.New("此检验类型下没有启用的检验标准，无法生成检验单")
	}

	//关联对象（可选）：按ID绑定产品（product 域出口）
	if m.ProductInfoID != nil && *m.ProductInfoID != "" {
		productInfo, err := product.GetProductInfoByIDTx(model.DB.DB(), *m.ProductInfoID)
		if err != nil {
			return "", errors.New("无效的产品ID")
		}
		if productInfo.ProductOrderID != "" {
			m.ProductOrderID = &productInfo.ProductOrderID
		}
	}

	//抽样数量：取各标准抽样量的最大值
	sampleQTY := m.LotQTY
	for _, standard := range standards {
		if n := SampleSizeByLot(m.LotQTY, standard.SampleLevel); n < sampleQTY {
			sampleQTY = n
		}
	}
	m.SampleQTY = sampleQTY
	m.CurrentState = QualityInspectionStateWait

	now := time.Now()
	prefix := fmt.Sprintf("QI%s", now.Format("20060102"))
	inspectionOrderNo, err := system.GenerateSerialNumber("检验单号", "质量检验单流水号", prefix, 4, 1)
	if err != nil {
		return "", err
	}
	m.InspectionOrderNo = inspectionOrderNo

	for _, standard := range standards {
		m.QualityInspectionOrderItems = append(m.QualityInspectionOrderItems, &model.QualityInspectionOrderItem{
			QualityInspectionStandardID: standard.ID,
		})
	}

	return m.ID, model.DB.DB().Create(m).Error
}

func UpdateQualityInspectionOrder(m *model.QualityInspectionOrder) error {
	if m.CurrentState != QualityInspectionStateWait {
		return errors.New("只有待检验状态的检验单可以修改")
	}
	return model.DB.DB().Omit("created_at", "quality_inspection_order_items").Save(m).Error
}

func QueryQualityInspectionOrder(req *proto.QueryQualityInspectionOrderRequest, resp *proto.QueryQualityInspectionOrderResponse, preload bool) {
	db := model.DB.DB().Model(&model.QualityInspectionOrder{})
	if preload {
		db = db.Preload("QualityInspectionType").Preload("ProductOrder").Preload("ProductInfo").
			Preload("QualityInspectionOrderItems").Preload("QualityInspectionOrderItems.QualityInspectionStandard")
	}
	if req.Code != "" {
		db = db.Where("`inspection_order_no` LIKE ? OR `product_serial_no` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}
	if req.CurrentState != "" {
		db = db.Where("`current_state` = ?", req.CurrentState)
	}
	if req.QualityInspectionTypeID != "" {
		db = db.Where("`quality_inspection_type_id` = ?", req.QualityInspectionTypeID)
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.QualityInspectionOrder
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.QualityInspectionOrdersToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllQualityInspectionOrders() (list []*model.QualityInspectionOrder, err error) {
	err = model.DB.DB().Preload(clause.Associations).Find(&list).Error
	return
}

func GetQualityInspectionOrderByID(id string) (*model.QualityInspectionOrder, error) {
	m := &model.QualityInspectionOrder{}
	err := model.DB.DB().Preload("QualityInspectionType").Preload("ProductOrder").Preload("ProductInfo").
		Preload("QualityInspectionOrderItems").Preload("QualityInspectionOrderItems.QualityInspectionStandard").
		Where("`id` = ?", id).First(m).Error
	return m, err
}

func DeleteQualityInspectionOrder(id string) (err error) {
	order := &model.QualityInspectionOrder{}
	if err := model.DB.DB().First(order, "`id` = ?", id).Error; err != nil {
		return err
	}
	if order.CurrentState == QualityInspectionStateFinished {
		return errors.New("已完成的检验单不能删除")
	}
	return model.DB.DB().Delete(&model.QualityInspectionOrder{}, "`id` = ?", id).Error
}

// CompleteQualityInspectionOrder 完成检验：录入实测值 → 按标准判定单项 → 按 AQL 判定整单 → 可选生成返工记录
func CompleteQualityInspectionOrder(req *proto.CompleteQualityInspectionOrderRequest) error {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		order := &model.QualityInspectionOrder{}
		if err := tx.Preload("QualityInspectionOrderItems").Preload("QualityInspectionOrderItems.QualityInspectionStandard").
			First(order, "`id` = ?", req.Id).Error; err != nil {
			return errors.New("读取检验单失败")
		}
		if order.CurrentState == QualityInspectionStateFinished {
			return errors.New("检验单已完成，不能重复提交")
		}

		measured := map[string]string{}
		for _, item := range req.Items {
			measured[item.QualityInspectionStandardID] = item.MeasuredValue
		}

		var defectives int
		for _, item := range order.QualityInspectionOrderItems {
			value, ok := measured[item.QualityInspectionStandardID]
			if !ok {
				return fmt.Errorf("检验标准%s缺少检测结果", item.QualityInspectionStandardID)
			}
			item.MeasuredValue = value
			qualified, err := judgeMeasuredValue(item.QualityInspectionStandard, value)
			if err != nil {
				return err
			}
			item.IsQualified = qualified
			if !qualified {
				defectives++
			}
			if err := tx.Save(item).Error; err != nil {
				return err
			}
		}

		//整单判定：以首个标准的AQL作为整单接收质量限（简化）
		var aql float64
		if len(order.QualityInspectionOrderItems) > 0 && order.QualityInspectionOrderItems[0].QualityInspectionStandard != nil {
			aql = order.QualityInspectionOrderItems[0].QualityInspectionStandard.Aql
		}
		accepted := defectives <= AcceptNumber(order.SampleQTY, aql)

		now := time.Now()
		order.ActualInspectionTime.Valid = true
		order.ActualInspectionTime.Time = now
		order.InspectionUserID = req.InspectionUserID
		order.Disposition = req.Disposition
		order.CurrentState = QualityInspectionStateFinished
		switch {
		case accepted:
			order.Conclusion = QualityConclusionQualified
		case req.Concession:
			//让步接收：判定不合格但特采放行
			order.Conclusion = QualityConclusionConcession
		default:
			order.Conclusion = QualityConclusionUnqualified
		}

		//不合格且勾选生成返工：写入返工记录（复用现有返工模块）
		if order.Conclusion == QualityConclusionUnqualified && req.CreateRework {
			rework := &model.ProductReworkRecord{
				ProductInfoID: deref(order.ProductInfoID),
				ReworkTime:    tool.Time2NullTime(now),
				ReworkReason:  fmt.Sprintf("检验单%s判定不合格：%s", order.InspectionOrderNo, req.Disposition),
			}
			if err := product.CreateProductReworkRecordTx(tx, rework); err != nil {
				return err
			}
			order.ReworkCreated = true
		}

		//不合格时联动产品状态为检查中（product 域出口；让步接收不改产品状态）
		if order.Conclusion == QualityConclusionUnqualified && order.ProductInfoID != nil && *order.ProductInfoID != "" {
			if err := product.MarkProductChecking(tx, *order.ProductInfoID); err != nil {
				return err
			}
		}

		if err := tx.Omit("created_at", "quality_inspection_order_items").Save(order).Error; err != nil {
			return err
		}
		return tx.Create(&model.OperationTrace{
			OperateUserID:  req.InspectionUserID,
			ControllerName: "质量管理",
			ActionName:     "完成检验",
			RequestContent: fmt.Sprintf("检验单:%s,结论:%s,不合格数:%d,让步:%v", order.InspectionOrderNo, order.Conclusion, defectives, req.Concession),
		}).Error
	})
}

// judgeMeasuredValue 按检验标准判定实测值
func judgeMeasuredValue(standard *model.QualityInspectionStandard, value string) (bool, error) {
	if standard == nil {
		return false, errors.New("检验标准不存在")
	}
	switch standard.CriterionMethod {
	case CriterionWithinRange:
		if standard.LowerLimit == "" || standard.UpperLimit == "" {
			return false, fmt.Errorf("检验标准%s配置了范围内判定但缺少规格限", standard.Code)
		}
		okLower, err := tool.MathOperator(value, "大于等于", standard.LowerLimit)
		if err != nil {
			return false, err
		}
		okUpper, err := tool.MathOperator(value, "小于等于", standard.UpperLimit)
		if err != nil {
			return false, err
		}
		return okLower && okUpper, nil
	case "":
		//未配置比较方式时按标准值相等判定
		return tool.MathOperator(value, "等于", standard.StandardValue)
	default:
		return tool.MathOperator(value, standard.CriterionMethod, standard.StandardValue)
	}
}

func deref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
