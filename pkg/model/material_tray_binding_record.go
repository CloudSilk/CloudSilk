package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

type MaterialTrayBindingRecord struct {
	ModelID
	CreateTime     time.Time     `gorm:"autoCreateTime:nano;comment:发料时间"`
	CreateUserID   string        `gorm:"size:36;comment:发料人员ID"`
	CurrentState   string        `gorm:"size:50;comment:当前状态"`
	LastUpdateTime time.Time     `gorm:"autoUpdateTime:nano;comment:状态变更时间"`
	Remark         string        `gorm:"size:500;comment:备注"`
	MaterialTrayID string        `gorm:"size:36;comment:使用托盘ID"`
	MaterialTray   *MaterialTray `gorm:"constraint:OnDelete:CASCADE"`
	ProductInfoID  string        `gorm:"size:36;comment:绑定产品ID"`
	ProductInfo    *ProductInfo  `gorm:"constraint:OnDelete:CASCADE"`
}

func PBToMaterialTrayBindingRecords(in []*proto.MaterialTrayBindingRecordInfo) []*MaterialTrayBindingRecord {
	var result []*MaterialTrayBindingRecord
	for _, c := range in {
		result = append(result, PBToMaterialTrayBindingRecord(c))
	}
	return result
}

func PBToMaterialTrayBindingRecord(in *proto.MaterialTrayBindingRecordInfo) *MaterialTrayBindingRecord {
	if in == nil {
		return nil
	}

	return &MaterialTrayBindingRecord{
		ModelID:        ModelID{ID: in.Id},
		CreateUserID:   in.CreateUserID,
		CurrentState:   in.CurrentState,
		Remark:         in.Remark,
		MaterialTrayID: in.MaterialTrayID,
		ProductInfoID:  in.ProductInfoID,
	}
}

func MaterialTrayBindingRecordsToPB(in []*MaterialTrayBindingRecord) []*proto.MaterialTrayBindingRecordInfo {
	var list []*proto.MaterialTrayBindingRecordInfo
	for _, f := range in {
		list = append(list, MaterialTrayBindingRecordToPB(f))
	}
	return list
}

func MaterialTrayBindingRecordToPB(in *MaterialTrayBindingRecord) *proto.MaterialTrayBindingRecordInfo {
	if in == nil {
		return nil
	}

	m := &proto.MaterialTrayBindingRecordInfo{
		Id:             in.ID,
		CreateTime:     utils.FormatTime(in.CreateTime),
		CreateUserID:   in.CreateUserID,
		CurrentState:   in.CurrentState,
		LastUpdateTime: utils.FormatTime(in.LastUpdateTime),
		Remark:         in.Remark,
		MaterialTrayID: in.MaterialTrayID,
		MaterialTray:   MaterialTrayToPB(in.MaterialTray),
		ProductInfoID:  in.ProductInfoID,
		ProductInfo:    ProductInfoToPB(in.ProductInfo),
	}
	return m
}
