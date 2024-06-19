package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 人员资质类型
type PersonnelQualificationType struct {
	ModelID
	Code                   string                                      `json:"code" gorm:"size:50;comment:代号"`
	Description            string                                      `json:"description" gorm:"size:500;comment:描述"`
	EffectiveDuration      int32                                       `json:"effectiveDuration" gorm:"comment:有效时长(天)"`
	ExpirationEarlyWarning bool                                        `json:"expirationEarlyWarning" gorm:"comment:失效预警"`
	ProductModels          []*PersonnelQualificationTypeAvailableModel `json:"productModels" gorm:"constraint:OnDelete:CASCADE"` //产品型号
	ProductionProcessID    string                                      `json:"productionProcessID" gorm:"index;size:36;comment:生产工序ID;"`
	ProductionProcess      *ProductionProcess                          `json:"productionProcess" gorm:"constraint:OnDelete:CASCADE"` //生产工序
	Remark                 string                                      `json:"remark" gorm:"size:500;comment:备注"`
}

// 支持产品型号
type PersonnelQualificationTypeAvailableModel struct {
	ModelID
	PersonnelQualificationTypeID string `gorm:"index;size:36;comment:人员资质类型ID"`
	ProductModelID               string `gorm:"size:36;comment:产品型号ID"`
}

func PBToPersonnelQualificationTypes(in []*proto.PersonnelQualificationTypeInfo) []*PersonnelQualificationType {
	var result []*PersonnelQualificationType
	for _, c := range in {
		result = append(result, PBToPersonnelQualificationType(c))
	}
	return result
}

func PBToPersonnelQualificationType(in *proto.PersonnelQualificationTypeInfo) *PersonnelQualificationType {
	if in == nil {
		return nil
	}
	return &PersonnelQualificationType{
		ModelID:                ModelID{ID: in.Id},
		Code:                   in.Code,
		Description:            in.Description,
		EffectiveDuration:      in.EffectiveDuration,
		ExpirationEarlyWarning: in.ExpirationEarlyWarning,
		ProductModels:          PBToPersonnelQualificationTypeAvailableModels(in.ProductModels),
		ProductionProcessID:    in.ProductionProcessID,
		Remark:                 in.Remark,
	}
}

func PersonnelQualificationTypesToPB(in []*PersonnelQualificationType) []*proto.PersonnelQualificationTypeInfo {
	var list []*proto.PersonnelQualificationTypeInfo
	for _, f := range in {
		list = append(list, PersonnelQualificationTypeToPB(f))
	}
	return list
}

func PersonnelQualificationTypeToPB(in *PersonnelQualificationType) *proto.PersonnelQualificationTypeInfo {
	if in == nil {
		return nil
	}
	m := &proto.PersonnelQualificationTypeInfo{
		Id:                     in.ID,
		Code:                   in.Code,
		Description:            in.Description,
		EffectiveDuration:      in.EffectiveDuration,
		ExpirationEarlyWarning: in.ExpirationEarlyWarning,
		ProductModels:          PersonnelQualificationTypeAvailableModelsToPB(in.ProductModels),
		ProductionProcessID:    in.ProductionProcessID,
		ProductionProcess:      ProductionProcessToPB(in.ProductionProcess),
		Remark:                 in.Remark,
	}
	return m
}

func PBToPersonnelQualificationTypeAvailableModels(in []*proto.PersonnelQualificationTypeAvailableModelInfo) []*PersonnelQualificationTypeAvailableModel {
	var result []*PersonnelQualificationTypeAvailableModel
	for _, c := range in {
		result = append(result, PBToPersonnelQualificationTypeAvailableModel(c))
	}
	return result
}

func PBToPersonnelQualificationTypeAvailableModel(in *proto.PersonnelQualificationTypeAvailableModelInfo) *PersonnelQualificationTypeAvailableModel {
	if in == nil {
		return nil
	}
	return &PersonnelQualificationTypeAvailableModel{
		ModelID:                      ModelID{ID: in.Id},
		PersonnelQualificationTypeID: in.PersonnelQualificationTypeID,
		ProductModelID:               in.ProductModelID,
	}
}

func PersonnelQualificationTypeAvailableModelsToPB(in []*PersonnelQualificationTypeAvailableModel) []*proto.PersonnelQualificationTypeAvailableModelInfo {
	var list []*proto.PersonnelQualificationTypeAvailableModelInfo
	for _, f := range in {
		list = append(list, PersonnelQualificationTypeAvailableModelToPB(f))
	}
	return list
}

func PersonnelQualificationTypeAvailableModelToPB(in *PersonnelQualificationTypeAvailableModel) *proto.PersonnelQualificationTypeAvailableModelInfo {
	if in == nil {
		return nil
	}
	m := &proto.PersonnelQualificationTypeAvailableModelInfo{
		Id:                           in.ID,
		PersonnelQualificationTypeID: in.PersonnelQualificationTypeID,
		ProductModelID:               in.ProductModelID,
	}
	return m
}
