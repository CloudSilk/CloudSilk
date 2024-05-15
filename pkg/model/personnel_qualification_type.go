package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 人员资质类型
type PersonnelQualificationType struct {
	ModelID
	Code                   string             `json:"code" gorm:"size:50;comment:代号"`
	Description            string             `json:"description" gorm:"size:500;comment:描述"`
	EffectiveDuration      int32              `json:"effectiveDuration" gorm:"comment:有效时长(天)"`
	ExpirationEarlyWarning bool               `json:"expirationEarlyWarning" gorm:"comment:失效预警"`
	ProductModels          []*ProductModel    `json:"productModels" gorm:"-"` //产品型号
	ProductionProcessID    string             `json:"productionProcessID" gorm:"index;size:36;comment:生产工序ID;"`
	ProductionProcess      *ProductionProcess `json:"productionProcess"` //生产工序
	Remark                 string             `json:"remark" gorm:"size:500;comment:备注"`
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
		ProductModels:          PBToProductModels(in.ProductModels),
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
		ProductModels:          ProductModelsToPB(in.ProductModels),
		ProductionProcessID:    in.ProductionProcessID,
		ProductionProcess:      ProductionProcessToPB(in.ProductionProcess),
		Remark:                 in.Remark,
	}
	return m
}
