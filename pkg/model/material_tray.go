package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 物料托盘
type MaterialTray struct {
	ModelID
	Code             string          `json:"code" gorm:"size:50;uniqueIndex;comment:代号"`
	Description      string          `json:"description" gorm:"size:500;comment:描述"`
	Identifier       string          `json:"identifier" gorm:"size:50;comment:识别码"`
	TrayType         string          `json:"trayType" gorm:"size:50;comment:托盘类型"`
	Enable           bool            `json:"enable" gorm:"comment:是否启用"`
	CurrentState     string          `json:"currentState" gorm:"comment:当前状态"`
	LastUpdateTime   time.Time       `json:"lastUpdateTime" gorm:"autoUpdateTime:nano;comment:状态变更时间"`
	Remark           string          `json:"remark" gorm:"size:500;comment:备注"`
	ProductionLineID string          `json:"productionLineID" gorm:"size:36;comment:隶属产线ID;"`
	ProductionLine   *ProductionLine `json:"productionLine" gorm:"constraint:OnDelete:CASCADE"` //隶属产线
	ProductInfoID    *string         `json:"productInfoID" gorm:"size:36;comment:当前产品ID;"`
	ProductInfo      *ProductInfo    `json:"productInfo" gorm:"constraint:OnDelete:SET NULL"` //当前产品
}

func PBToMaterialTrays(in []*proto.MaterialTrayInfo) []*MaterialTray {
	var result []*MaterialTray
	for _, c := range in {
		result = append(result, PBToMaterialTray(c))
	}
	return result
}

func PBToMaterialTray(in *proto.MaterialTrayInfo) *MaterialTray {
	if in == nil {
		return nil
	}

	var productInfoID *string
	if in.ProductInfoID != "" {
		productInfoID = &in.ProductInfoID
	}

	return &MaterialTray{
		ModelID:      ModelID{ID: in.Id},
		Code:         in.Code,
		Description:  in.Description,
		Identifier:   in.Identifier,
		TrayType:     in.TrayType,
		Enable:       in.Enable,
		CurrentState: in.CurrentState,
		// LastUpdateTime:   utils.ParseTime(in.LastUpdateTime),
		Remark:           in.Remark,
		ProductionLineID: in.ProductionLineID,
		ProductInfoID:    productInfoID,
	}
}

func MaterialTraysToPB(in []*MaterialTray) []*proto.MaterialTrayInfo {
	var list []*proto.MaterialTrayInfo
	for _, f := range in {
		list = append(list, MaterialTrayToPB(f))
	}
	return list
}

func MaterialTrayToPB(in *MaterialTray) *proto.MaterialTrayInfo {
	if in == nil {
		return nil
	}

	var productInfoID string
	if in.ProductInfoID != nil {
		productInfoID = *in.ProductInfoID
	}

	m := &proto.MaterialTrayInfo{
		Id:               in.ID,
		Code:             in.Code,
		Description:      in.Description,
		Identifier:       in.Identifier,
		TrayType:         in.TrayType,
		Enable:           in.Enable,
		CurrentState:     in.CurrentState,
		LastUpdateTime:   utils.FormatTime(in.LastUpdateTime),
		Remark:           in.Remark,
		ProductionLineID: in.ProductionLineID,
		ProductionLine:   ProductionLineToPB(in.ProductionLine),
		ProductInfoID:    productInfoID,
		ProductInfo:      ProductInfoToPB(in.ProductInfo),
	}
	return m
}
