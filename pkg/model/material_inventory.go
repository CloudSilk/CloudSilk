package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 物料库存
type MaterialInventory struct {
	ModelID
	MaterialInfoID  string         `gorm:"size:36;comment:物料信息ID"`
	MaterialInfo    *MaterialInfo  `gorm:"constraint:OnDelete:CASCADE"` //物料信息
	MaterialStoreID string         `gorm:"size:36;comment:物料仓库ID"`
	MaterialStore   *MaterialStore `gorm:"constraint:OnDelete:CASCADE"` //物料仓库
	StoredQTY       int64          `gorm:"comment:库存数量"`
	IssuedQTY       int64          `gorm:"comment:锁定数量"`
	FeedingQTY      int64          `gorm:"comment:正在补料数量"`
	IssuingQTY      int64          `gorm:"comment:正在备库数量"`
	CreateTime      time.Time      `gorm:"autoCreateTime:nano;comment:创建时间"`
	CreateUserID    string         `gorm:"size:36;comment:创建人员"`
	CurrentState    string         `gorm:"size:50;comment:当前状态"`
	LastUpdateTime  time.Time      `gorm:"autoUpdateTime:nano;comment:状态变更时间"`
	Remark          string         `gorm:"size:500;comment:备注"`
}

func PBToMaterialInventorys(in []*proto.MaterialInventoryInfo) []*MaterialInventory {
	var result []*MaterialInventory
	for _, c := range in {
		result = append(result, PBToMaterialInventory(c))
	}
	return result
}

func PBToMaterialInventory(in *proto.MaterialInventoryInfo) *MaterialInventory {
	if in == nil {
		return nil
	}
	return &MaterialInventory{
		ModelID:         ModelID{ID: in.Id},
		MaterialInfoID:  in.MaterialInfoID,
		MaterialStoreID: in.MaterialStoreID,
		StoredQTY:       in.StoredQTY,
		IssuedQTY:       in.IssuedQTY,
		FeedingQTY:      in.FeedingQTY,
		IssuingQTY:      in.IssuingQTY,
		CreateUserID:    in.CreateUserID,
		CurrentState:    in.CurrentState,
		Remark:          in.Remark,
	}
}

func MaterialInventorysToPB(in []*MaterialInventory) []*proto.MaterialInventoryInfo {
	var list []*proto.MaterialInventoryInfo
	for _, f := range in {
		list = append(list, MaterialInventoryToPB(f))
	}
	return list
}

func MaterialInventoryToPB(in *MaterialInventory) *proto.MaterialInventoryInfo {
	if in == nil {
		return nil
	}
	m := &proto.MaterialInventoryInfo{
		Id:              in.ID,
		MaterialInfoID:  in.MaterialInfoID,
		MaterialInfo:    MaterialInfoToPB(in.MaterialInfo),
		MaterialStoreID: in.MaterialStoreID,
		MaterialStore:   MaterialStoreToPB(in.MaterialStore),
		StoredQTY:       in.StoredQTY,
		IssuedQTY:       in.IssuedQTY,
		FeedingQTY:      in.FeedingQTY,
		IssuingQTY:      in.IssuingQTY,
		CreateTime:      utils.FormatTime(in.CreateTime),
		CreateUserID:    in.CreateUserID,
		CurrentState:    in.CurrentState,
		LastUpdateTime:  utils.FormatTime(in.LastUpdateTime),
		Remark:          in.Remark,
	}
	return m
}
