package model

import (
	"database/sql"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 检验单明细
type QualityInspectionOrderItem struct {
	ModelID
	QualityInspectionOrderID           string                      `json:"qualityInspectionOrderID" gorm:"index;size:36;comment:检验单ID"`
	QualityInspectionOrder             *QualityInspectionOrder     `json:"qualityInspectionOrder" gorm:"constraint:OnDelete:CASCADE"`
	QualityInspectionStandardID        string                      `json:"qualityInspectionStandardID" gorm:"size:36;comment:检验标准ID"`
	QualityInspectionStandard          *QualityInspectionStandard  `json:"qualityInspectionStandard" gorm:"constraint:OnDelete:SET NULL"`
	MeasuredValue                      string                      `json:"measuredValue" gorm:"size:1000;comment:实测值"`
	IsQualified                        bool                        `json:"isQualified" gorm:"comment:是否合格"`
	Remark                             string                      `json:"remark" gorm:"size:1000;comment:备注"`
}

// 检验单
type QualityInspectionOrder struct {
	ModelID
	InspectionOrderNo       string                     `json:"inspectionOrderNo" gorm:"index;size:100;comment:检验单号"`
	QualityInspectionTypeID *string                    `json:"qualityInspectionTypeID" gorm:"size:36;comment:检验类型ID"`
	QualityInspectionType   *QualityInspectionType     `json:"qualityInspectionType" gorm:"constraint:OnDelete:SET NULL"`
	ProductOrderID          *string                    `json:"productOrderID" gorm:"size:36;comment:生产工单ID"`
	ProductOrder            *ProductOrder              `json:"productOrder" gorm:"constraint:OnDelete:SET NULL"`
	ProductInfoID           *string                    `json:"productInfoID" gorm:"size:36;comment:产品ID"`
	ProductInfo             *ProductInfo               `json:"productInfo" gorm:"constraint:OnDelete:SET NULL"`
	LotQTY                  int32                      `json:"lotQTY" gorm:"comment:批次数量"`
	SampleQTY               int32                      `json:"sampleQTY" gorm:"comment:抽样数量"`
	CurrentState            string                     `json:"currentState" gorm:"size:100;comment:当前状态"`
	PlanInspectionTime      sql.NullTime               `json:"planInspectionTime" gorm:"comment:计划检验时间"`
	ActualInspectionTime    sql.NullTime               `json:"actualInspectionTime" gorm:"comment:实际检验时间"`
	InspectionUserID        string                     `json:"inspectionUserID" gorm:"size:36;comment:检验人"`
	Conclusion              string                     `json:"conclusion" gorm:"size:100;comment:整单结论"`
	Disposition             string                     `json:"disposition" gorm:"size:1000;comment:不合格处理说明"`
	ReworkCreated           bool                       `json:"reworkCreated" gorm:"comment:是否已生成返工记录"`
	Remark                  string                     `json:"remark" gorm:"size:1000;comment:备注"`
	CreateTime              time.Time                  `json:"createTime" gorm:"autoCreateTime:nano;comment:创建时间"`
	QualityInspectionOrderItems []*QualityInspectionOrderItem `json:"qualityInspectionOrderItems" gorm:"constraint:OnDelete:CASCADE"`
}

func PBToQualityInspectionOrders(in []*proto.QualityInspectionOrderInfo) []*QualityInspectionOrder {
	var result []*QualityInspectionOrder
	for _, c := range in {
		result = append(result, PBToQualityInspectionOrder(c))
	}
	return result
}

func PBToQualityInspectionOrder(in *proto.QualityInspectionOrderInfo) *QualityInspectionOrder {
	if in == nil {
		return nil
	}
	var qualityInspectionTypeID, productOrderID *string
	if in.QualityInspectionTypeID != "" {
		qualityInspectionTypeID = &in.QualityInspectionTypeID
	}
	if in.ProductOrderID != "" {
		productOrderID = &in.ProductOrderID
	}

	items := []*QualityInspectionOrderItem{}
	for _, item := range in.QualityInspectionOrderItems {
		items = append(items, &QualityInspectionOrderItem{
			ModelID:                     ModelID{ID: item.Id},
			QualityInspectionOrderID:    in.Id,
			QualityInspectionStandardID: item.QualityInspectionStandardID,
			MeasuredValue:               item.MeasuredValue,
			IsQualified:                 item.IsQualified,
			Remark:                      item.Remark,
		})
	}

	return &QualityInspectionOrder{
		ModelID:                 ModelID{ID: in.Id},
		InspectionOrderNo:       in.InspectionOrderNo,
		QualityInspectionTypeID: qualityInspectionTypeID,
		ProductOrderID:          productOrderID,
		LotQTY:                  in.LotQTY,
		SampleQTY:               in.SampleQTY,
		CurrentState:            in.CurrentState,
		InspectionUserID:        in.InspectionUserID,
		Conclusion:              in.Conclusion,
		Disposition:             in.Disposition,
		ReworkCreated:           in.ReworkCreated,
		Remark:                  in.Remark,
		QualityInspectionOrderItems: items,
	}
}

func QualityInspectionOrdersToPB(in []*QualityInspectionOrder) []*proto.QualityInspectionOrderInfo {
	var list []*proto.QualityInspectionOrderInfo
	for _, f := range in {
		list = append(list, QualityInspectionOrderToPB(f))
	}
	return list
}

func QualityInspectionOrderToPB(in *QualityInspectionOrder) *proto.QualityInspectionOrderInfo {
	if in == nil {
		return nil
	}
	m := &proto.QualityInspectionOrderInfo{
		Id:                   in.ID,
		InspectionOrderNo:    in.InspectionOrderNo,
		LotQTY:               in.LotQTY,
		SampleQTY:            in.SampleQTY,
		CurrentState:         in.CurrentState,
		InspectionUserID:     in.InspectionUserID,
		Conclusion:           in.Conclusion,
		Disposition:          in.Disposition,
		ReworkCreated:        in.ReworkCreated,
		Remark:               in.Remark,
	}
	if in.QualityInspectionTypeID != nil {
		m.QualityInspectionTypeID = *in.QualityInspectionTypeID
	}
	if in.QualityInspectionType != nil {
		m.QualityInspectionTypeCode = in.QualityInspectionType.Code
	}
	if in.ProductOrderID != nil {
		m.ProductOrderID = *in.ProductOrderID
	}
	if in.ProductOrder != nil {
		m.ProductOrderNo = in.ProductOrder.ProductOrderNo
	}
	if in.ProductInfoID != nil {
		m.ProductSerialNo = *in.ProductInfoID
	}
	if in.ProductInfo != nil {
		m.ProductSerialNo = in.ProductInfo.ProductSerialNo
	}
	if in.PlanInspectionTime.Valid {
		m.PlanInspectionTime = in.PlanInspectionTime.Time.Format("2006-01-02 15:04:05")
	}
	if in.ActualInspectionTime.Valid {
		m.ActualInspectionTime = in.ActualInspectionTime.Time.Format("2006-01-02 15:04:05")
	}
	for _, item := range in.QualityInspectionOrderItems {
		m.QualityInspectionOrderItems = append(m.QualityInspectionOrderItems, QualityInspectionOrderItemToPB(item))
	}
	return m
}

func QualityInspectionOrderItemToPB(in *QualityInspectionOrderItem) *proto.QualityInspectionOrderItemInfo {
	if in == nil {
		return nil
	}
	m := &proto.QualityInspectionOrderItemInfo{
		Id:                          in.ID,
		QualityInspectionOrderID:    in.QualityInspectionOrderID,
		QualityInspectionStandardID: in.QualityInspectionStandardID,
		MeasuredValue:               in.MeasuredValue,
		IsQualified:                 in.IsQualified,
		Remark:                      in.Remark,
	}
	if in.QualityInspectionStandard != nil {
		m.QualityInspectionStandardCode = in.QualityInspectionStandard.Code
		m.QualityInspectionStandardDescription = in.QualityInspectionStandard.Description
	}
	return m
}
