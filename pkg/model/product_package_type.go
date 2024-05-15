package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

type ProductPackageType struct {
	ModelID
	Code          string       `json:"code" gorm:"index;size:100;comment:类型"`
	Description   string       `json:"description" gorm:"size:1000;comment:描述"`
	Remark        string       `json:"remark" gorm:"size:1000;comment:备注"`
	LabelTypeID   *string      `json:"labelTypeID" gorm:"comment:标签类型ID"`
	LabelType     *LabelType   `json:"labelType"` //标签类型
	SystemEventID *string      `json:"systemEventID" gorm:"comment:系统事件ID"`
	SystemEvent   *SystemEvent `json:"systemEvent"` //系统事件
}

func PBToProductPackageTypes(in []*proto.ProductPackageTypeInfo) []*ProductPackageType {
	var result []*ProductPackageType
	for _, c := range in {
		result = append(result, PBToProductPackageType(c))
	}
	return result
}

func PBToProductPackageType(in *proto.ProductPackageTypeInfo) *ProductPackageType {
	if in == nil {
		return nil
	}
	var labelTypeID, systemEventID *string
	if in.LabelTypeID != "" {
		labelTypeID = &in.LabelTypeID
	}
	if in.SystemEventID != "" {
		systemEventID = &in.SystemEventID
	}
	return &ProductPackageType{
		ModelID:       ModelID{ID: in.Id},
		Code:          in.Code,
		Description:   in.Description,
		Remark:        in.Remark,
		LabelTypeID:   labelTypeID,
		SystemEventID: systemEventID,
	}
}

func ProductPackageTypesToPB(in []*ProductPackageType) []*proto.ProductPackageTypeInfo {
	var list []*proto.ProductPackageTypeInfo
	for _, f := range in {
		list = append(list, ProductPackageTypeToPB(f))
	}
	return list
}

func ProductPackageTypeToPB(in *ProductPackageType) *proto.ProductPackageTypeInfo {
	if in == nil {
		return nil
	}

	var labelTypeID, systemEventID string
	if in.LabelTypeID != nil {
		labelTypeID = *in.LabelTypeID
	}
	if in.SystemEventID != nil {
		systemEventID = *in.SystemEventID
	}

	m := &proto.ProductPackageTypeInfo{
		Id:            in.ID,
		Code:          in.Code,
		Description:   in.Description,
		Remark:        in.Remark,
		LabelTypeID:   labelTypeID,
		LabelType:     LabelTypeToPB(in.LabelType),
		SystemEventID: systemEventID,
		SystemEvent:   SystemEventToPB(in.SystemEvent),
	}
	return m
}
