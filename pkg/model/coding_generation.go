package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 编码序列记录
type CodingGeneration struct {
	ModelID
	Code             string          `gorm:"size:50;comment:编码"`
	CreateTime       time.Time       `gorm:"autoCreateTime:nano;comment:创建时间"`
	CreateUserID     string          `gorm:"size:36;comment:创建人员ID"`
	CodingTemplateID string          `gorm:"size:36;comment:编码模板ID"`
	CodingTemplate   *CodingTemplate `gorm:"constraint:OnDelete:CASCADE"` //编码模板
}

func PBToCodingGenerations(in []*proto.CodingGenerationInfo) []*CodingGeneration {
	var result []*CodingGeneration
	for _, c := range in {
		result = append(result, PBToCodingGeneration(c))
	}
	return result
}

func PBToCodingGeneration(in *proto.CodingGenerationInfo) *CodingGeneration {
	if in == nil {
		return nil
	}
	return &CodingGeneration{
		ModelID:          ModelID{ID: in.Id},
		Code:             in.Code,
		CreateUserID:     in.CreateUserID,
		CodingTemplateID: in.CodingTemplateID,
	}
}

func CodingGenerationsToPB(in []*CodingGeneration) []*proto.CodingGenerationInfo {
	var list []*proto.CodingGenerationInfo
	for _, f := range in {
		list = append(list, CodingGenerationToPB(f))
	}
	return list
}

func CodingGenerationToPB(in *CodingGeneration) *proto.CodingGenerationInfo {
	if in == nil {
		return nil
	}
	m := &proto.CodingGenerationInfo{
		Id:               in.ID,
		Code:             in.Code,
		CreateTime:       utils.FormatTime(in.CreateTime),
		CreateUserID:     in.CreateUserID,
		CodingTemplateID: in.CodingTemplateID,
		CodingTemplate:   CodingTemplateToPB(in.CodingTemplate),
	}
	return m
}
