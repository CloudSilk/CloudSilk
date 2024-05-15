package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 序列号
type SerialNumber struct {
	ModelID
	Name             string    `json:"name" gorm:"size:200;comment:名称"`
	Description      string    `json:"description" gorm:"size:1000;comment:描述"`
	Prefix           string    `json:"prefix" gorm:"size:200;comment:前缀"`
	Length           int32     `json:"length" gorm:"comment:长度"`
	Seed             int32     `json:"seed" gorm:"comment:种子"`
	Increment        int32     `json:"increment" gorm:"comment:增量"`
	CreateTime       time.Time `json:"createTime" gorm:"autoCreateTime;comment:创建时间"`
	LastIncreaseTime time.Time `json:"lastIncreaseTime" gorm:"autoCreateTime;comment:最后增长时间"`
}

func PBToSerialNumbers(in []*proto.SerialNumberInfo) []*SerialNumber {
	var result []*SerialNumber
	for _, c := range in {
		result = append(result, PBToSerialNumber(c))
	}
	return result
}

func PBToSerialNumber(in *proto.SerialNumberInfo) *SerialNumber {
	if in == nil {
		return nil
	}
	return &SerialNumber{
		ModelID:     ModelID{ID: in.Id},
		Name:        in.Name,
		Description: in.Description,
		Prefix:      in.Prefix,
		Length:      in.Length,
		Seed:        in.Seed,
		Increment:   in.Increment,
		// CreateTime:       utils.ParseTime(in.CreateTime),
		// LastIncreaseTime: utils.ParseTime(in.LastIncreaseTime),
	}
}

func SerialNumbersToPB(in []*SerialNumber) []*proto.SerialNumberInfo {
	var list []*proto.SerialNumberInfo
	for _, f := range in {
		list = append(list, SerialNumberToPB(f))
	}
	return list
}

func SerialNumberToPB(in *SerialNumber) *proto.SerialNumberInfo {
	if in == nil {
		return nil
	}
	m := &proto.SerialNumberInfo{
		Id:               in.ID,
		Name:             in.Name,
		Description:      in.Description,
		Prefix:           in.Prefix,
		Length:           in.Length,
		Seed:             in.Seed,
		Increment:        in.Increment,
		CreateTime:       utils.FormatTime(in.CreateTime),
		LastIncreaseTime: utils.FormatTime(in.LastIncreaseTime),
	}
	return m
}
