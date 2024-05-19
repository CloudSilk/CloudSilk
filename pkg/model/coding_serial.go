package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 编码序列
type CodingSerial struct {
	ModelID
	Prefix           string    `gorm:"size:100;comment:前缀"`
	Seed             int32     `gorm:"comment:种子"`
	CreateTime       time.Time `gorm:"autoCreateTime:nano;comment:创建时间"`
	LastIncreaseTime time.Time `gorm:"comment:最后增长时间"`
}

func PBToCodingSerials(in []*proto.CodingSerialInfo) []*CodingSerial {
	var result []*CodingSerial
	for _, c := range in {
		result = append(result, PBToCodingSerial(c))
	}
	return result
}

func PBToCodingSerial(in *proto.CodingSerialInfo) *CodingSerial {
	if in == nil {
		return nil
	}
	return &CodingSerial{
		ModelID:          ModelID{ID: in.Id},
		Prefix:           in.Prefix,
		Seed:             in.Seed,
		LastIncreaseTime: utils.ParseTime(in.LastIncreaseTime),
	}
}

func CodingSerialsToPB(in []*CodingSerial) []*proto.CodingSerialInfo {
	var list []*proto.CodingSerialInfo
	for _, f := range in {
		list = append(list, CodingSerialToPB(f))
	}
	return list
}

func CodingSerialToPB(in *CodingSerial) *proto.CodingSerialInfo {
	if in == nil {
		return nil
	}
	m := &proto.CodingSerialInfo{
		Id:               in.ID,
		Prefix:           in.Prefix,
		Seed:             in.Seed,
		CreateTime:       utils.FormatTime(in.CreateTime),
		LastIncreaseTime: utils.FormatTime(in.LastIncreaseTime),
	}
	return m
}
