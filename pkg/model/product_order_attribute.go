package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 产品工单特性
type ProductOrderAttribute struct {
	ModelID
	Value              string            `json:"Value" gorm:"size:200;comment:值"`
	Description        string            `json:"Description" gorm:"size:500;comment:描述"`
	ProductAttributeID string            `json:"ProductAttributeID" gorm:"size:36;comment:产品特性ID"`
	ProductAttribute   *ProductAttribute `json:"ProductAttribute"` //产品特性
	CreateTime         time.Time         `json:"CreateTime" gorm:"autoCreateTime:nano;comment:创建时间"`
	CreateUserID       string            `json:"CreateUserID" gorm:"size:36;comment:创建人员ID"`
	ProductOrderID     string            `json:"ProductOrderID" gorm:"size:36;comment:隶属工单ID"`
	ProductOrder       *ProductOrder     `json:"ProductOrder"` //隶属工单
}

func PBToProductOrderAttributes(in []*proto.ProductOrderAttributeInfo) []*ProductOrderAttribute {
	var result []*ProductOrderAttribute
	for _, c := range in {
		result = append(result, PBToProductOrderAttribute(c))
	}
	return result
}

func PBToProductOrderAttribute(in *proto.ProductOrderAttributeInfo) *ProductOrderAttribute {
	if in == nil {
		return nil
	}
	return &ProductOrderAttribute{
		ModelID:            ModelID{ID: in.Id},
		Value:              in.Value,
		Description:        in.Description,
		ProductAttributeID: in.ProductAttributeID,
		// CreateTime:         utils.ParseTime(in.CreateTime),
		CreateUserID:   in.CreateUserID,
		ProductOrderID: in.ProductOrderID,
	}
}

func ProductOrderAttributesToPB(in []*ProductOrderAttribute) []*proto.ProductOrderAttributeInfo {
	var list []*proto.ProductOrderAttributeInfo
	for _, f := range in {
		list = append(list, ProductOrderAttributeToPB(f))
	}
	return list
}

func ProductOrderAttributeToPB(in *ProductOrderAttribute) *proto.ProductOrderAttributeInfo {
	if in == nil {
		return nil
	}

	m := &proto.ProductOrderAttributeInfo{
		Id:                 in.ID,
		Value:              in.Value,
		Description:        in.Description,
		ProductAttributeID: in.ProductAttributeID,
		CreateTime:         utils.FormatTime(in.CreateTime),
		CreateUserID:       in.CreateUserID,
		ProductOrderID:     in.ProductOrderID,
		ProductAttribute:   ProductAttributeToPB(in.ProductAttribute),
		ProductOrder:       ProductOrderToPB(in.ProductOrder),
	}
	return m
}
