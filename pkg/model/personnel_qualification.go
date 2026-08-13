package model

import (
	"database/sql"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 人员资质
type PersonnelQualification struct {
	ModelID
	ProductModelID      string            `json:"productModelID" gorm:"size:36;comment:产品型号ID"`
	ProductionProcessID string            `json:"productionProcessID" gorm:"size:36;comment:生产工序ID"`
	ProductionProcess   ProductionProcess `json:"productionProcess" gorm:"size:36;comment:生产工序ID"`
	CertifiedUserID     string            `json:"certifiedUserID" gorm:"index;size:36;comment:认证人员ID"`
	EffectiveDate       sql.NullTime      `json:"effectiveDate" gorm:"comment:生效日期"`
	ExpirationDate      sql.NullTime      `json:"expirationDate" gorm:"comment:失效日期"`
	AuthorizedUserID    string            `json:"authorizedUserID" gorm:"index;size:36;comment:授权人员ID"`
	AuthorizedTime      time.Time         `json:"authorizedTime" gorm:"autoCreateTime:nano;comment:授权时间"`
	Remark              string            `json:"remark" gorm:"size:1000;comment:备注"`
}

func PBToPersonnelQualifications(in []*proto.PersonnelQualificationInfo) []*PersonnelQualification {
	var result []*PersonnelQualification
	for _, c := range in {
		result = append(result, PBToPersonnelQualification(c))
	}
	return result
}

func PBToPersonnelQualification(in *proto.PersonnelQualificationInfo) *PersonnelQualification {
	if in == nil {
		return nil
	}
	return &PersonnelQualification{
		ModelID:             ModelID{ID: in.Id},
		ProductModelID:      in.ProductModelID,
		ProductionProcessID: in.ProductionProcessID,
		CertifiedUserID:     in.CertifiedUserID,
		EffectiveDate:       utils.ParseSqlNullDate(in.EffectiveDate),
		ExpirationDate:      utils.ParseSqlNullDate(in.ExpirationDate),
		AuthorizedUserID:    in.AuthorizedUserID,
		AuthorizedTime:      utils.ParseTime(in.AuthorizedTime),
		Remark:              in.Remark,
	}
}

func PersonnelQualificationsToPB(in []*PersonnelQualification) []*proto.PersonnelQualificationInfo {
	var list []*proto.PersonnelQualificationInfo
	for _, f := range in {
		list = append(list, PersonnelQualificationToPB(f))
	}
	return list
}

func PersonnelQualificationToPB(in *PersonnelQualification) *proto.PersonnelQualificationInfo {
	if in == nil {
		return nil
	}
	m := &proto.PersonnelQualificationInfo{
		Id:                  in.ID,
		ProductModelID:      in.ProductModelID,
		ProductionProcessID: in.ProductionProcessID,
		CertifiedUserID:     in.CertifiedUserID,
		EffectiveDate:       utils.FormatSqlNullDate(in.EffectiveDate),
		ExpirationDate:      utils.FormatSqlNullDate(in.ExpirationDate),
		AuthorizedUserID:    in.AuthorizedUserID,
		AuthorizedTime:      utils.FormatTime(in.AuthorizedTime),
		Remark:              in.Remark,
	}
	return m
}
