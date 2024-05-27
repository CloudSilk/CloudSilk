package model

import (
	"database/sql"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 工单标签
type ProductOrderLabel struct {
	ModelID
	ProductOrderID              string                        `json:"productOrderID" gorm:"size:36;comment:归属工单ID"`
	ProductOrder                *ProductOrder                 `json:"productOrder" gorm:"constraint:OnDelete:CASCADE"`
	LabelTypeID                 string                        `json:"labelTypeID" gorm:"size:36;comment:标签类型ID"`
	LabelType                   *LabelType                    `json:"labelType" gorm:"constraint:OnDelete:CASCADE"`
	FilePath                    string                        `json:"filePath" gorm:"size:2000;comment:模板文件"`
	ReferTimes                  int32                         `json:"referTimes" gorm:"comment:引用次数"`
	DoubleCheck                 bool                          `json:"doubleCheck" gorm:"comment:需要复核"`
	CreateTime                  time.Time                     `json:"createTime" gorm:"autoCreateTime:nano;comment:创建时间"`
	CreateUserID                string                        `json:"createUserID" gorm:"size:36;comment:创建人员ID"`
	CheckTime                   sql.NullTime                  `json:"checkTime" gorm:"comment:复核时间"`
	CheckUserID                 string                        `json:"checkUserID" gorm:"size:36;comment:复核人员ID"`
	CurrentState                string                        `json:"currentState" gorm:"size:-1;comment:当前状态"`
	LastUpdateTime              sql.NullTime                  `json:"lastUpdateTime" gorm:"autoUpdateTime:nano;comment:状态变更时间"`
	ProductOrderLabelParameters []*ProductOrderLabelParameter `json:"productOrderLabelParameters" gorm:"constraint:OnDelete:CASCADE"` //标签参数
}

// 工单标签参数
type ProductOrderLabelParameter struct {
	ModelID
	ProductOrderLabelID string `json:"productOrderLabelID" gorm:"index;size:36;comment:工单标签ID"`
	Name                string `json:"name" gorm:"comment:名称"`
	FixedValue          string `json:"fixedValue" gorm:"comment:设定值"`
	Remark              string `json:"remark" gorm:"comment:备注"`
}

func PBToProductOrderLabels(in []*proto.ProductOrderLabelInfo) []*ProductOrderLabel {
	var result []*ProductOrderLabel
	for _, c := range in {
		result = append(result, PBToProductOrderLabel(c))
	}
	return result
}

func PBToProductOrderLabel(in *proto.ProductOrderLabelInfo) *ProductOrderLabel {
	if in == nil {
		return nil
	}
	return &ProductOrderLabel{
		ModelID:        ModelID{ID: in.Id},
		ProductOrderID: in.ProductOrderID,
		LabelTypeID:    in.LabelTypeID,
		FilePath:       in.FilePath,
		ReferTimes:     in.ReferTimes,
		DoubleCheck:    in.DoubleCheck,
		// CreateTime:                  utils.ParseTime(in.CreateTime),
		CreateUserID: in.CreateUserID,
		CheckTime:    utils.ParseSqlNullTime(in.CheckTime),
		CheckUserID:  in.CheckUserID,
		CurrentState: in.CurrentState,
		// LastUpdateTime:              utils.ParseSqlNullTime(in.LastUpdateTime),
		ProductOrderLabelParameters: PBToProductOrderLabelParameters(in.ProductOrderLabelParameters),
	}
}

func ProductOrderLabelsToPB(in []*ProductOrderLabel) []*proto.ProductOrderLabelInfo {
	var list []*proto.ProductOrderLabelInfo
	for _, f := range in {
		list = append(list, ProductOrderLabelToPB(f))
	}
	return list
}

func ProductOrderLabelToPB(in *ProductOrderLabel) *proto.ProductOrderLabelInfo {
	if in == nil {
		return nil
	}

	m := &proto.ProductOrderLabelInfo{
		Id:                          in.ID,
		ProductOrderID:              in.ProductOrderID,
		ProductOrder:                ProductOrderToPB(in.ProductOrder),
		LabelTypeID:                 in.LabelTypeID,
		LabelType:                   LabelTypeToPB(in.LabelType),
		FilePath:                    in.FilePath,
		ReferTimes:                  in.ReferTimes,
		DoubleCheck:                 in.DoubleCheck,
		CreateTime:                  utils.FormatTime(in.CreateTime),
		CreateUserID:                in.CreateUserID,
		CheckTime:                   utils.FormatSqlNullTime(in.CheckTime),
		CheckUserID:                 in.CheckUserID,
		CurrentState:                in.CurrentState,
		LastUpdateTime:              utils.FormatSqlNullTime(in.LastUpdateTime),
		ProductOrderLabelParameters: ProductOrderLabelParametersToPB(in.ProductOrderLabelParameters),
	}
	return m
}

func PBToProductOrderLabelParameters(in []*proto.ProductOrderLabelParameterInfo) []*ProductOrderLabelParameter {
	var result []*ProductOrderLabelParameter
	for _, c := range in {
		result = append(result, PBToProductOrderLabelParameter(c))
	}
	return result
}

func PBToProductOrderLabelParameter(in *proto.ProductOrderLabelParameterInfo) *ProductOrderLabelParameter {
	if in == nil {
		return nil
	}
	return &ProductOrderLabelParameter{
		ModelID:             ModelID{ID: in.Id},
		Name:                in.Name,
		FixedValue:          in.FixedValue,
		Remark:              in.Remark,
		ProductOrderLabelID: in.ProductOrderLabelID,
	}
}

func ProductOrderLabelParametersToPB(in []*ProductOrderLabelParameter) []*proto.ProductOrderLabelParameterInfo {
	var list []*proto.ProductOrderLabelParameterInfo
	for _, f := range in {
		list = append(list, ProductOrderLabelParameterToPB(f))
	}
	return list
}

func ProductOrderLabelParameterToPB(in *ProductOrderLabelParameter) *proto.ProductOrderLabelParameterInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductOrderLabelParameterInfo{
		Id:                  in.ID,
		Name:                in.Name,
		FixedValue:          in.FixedValue,
		Remark:              in.Remark,
		ProductOrderLabelID: in.ProductOrderLabelID,
	}
	return m
}
