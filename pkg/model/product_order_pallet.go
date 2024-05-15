package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

type ProductOrderPallet struct {
	ModelID
	PalletNo                  string                   `json:"palletNo" gorm:"size:100;comment:栈板编号"`
	PalletSn                  string                   `json:"palletSn" gorm:"size:100;comment:栈板序号"`
	CreateTime                time.Time                `json:"createTime" gorm:"autoCreateTime:nano;comment:创建时间"`
	PackageQuantity           int32                    `json:"packageQuantity" gorm:"comment:码垛箱数"`
	CurrentState              string                   `json:"currentState" gorm:"size:100;comment:当前状态"`
	LastUpdateTime            time.Time                `json:"lastUpdateTime" gorm:"autoUpdateTime:nano;comment:状态变更时间"`
	Remark                    string                   `json:"remark" gorm:"size:1000;comment:备注"`
	ProductPackageStackRuleID string                   `json:"productPackageStackRuleID" gorm:"size:36;comment:码垛规则ID"`
	ProductPackageStackRule   *ProductPackageStackRule `json:"productPackageStackRule"`
	ProductOrderID            string                   `json:"productOrderID" gorm:"size:36;comment:归属工单ID"`
	ProductOrder              *ProductOrder            `json:"productOrder"`
}

func PBToProductOrderPallets(in []*proto.ProductOrderPalletInfo) []*ProductOrderPallet {
	var result []*ProductOrderPallet
	for _, c := range in {
		result = append(result, PBToProductOrderPallet(c))
	}
	return result
}

func PBToProductOrderPallet(in *proto.ProductOrderPalletInfo) *ProductOrderPallet {
	if in == nil {
		return nil
	}
	return &ProductOrderPallet{
		ModelID:  ModelID{ID: in.Id},
		PalletNo: in.PalletNo,
		PalletSn: in.PalletSn,
		// CreateTime:                utils.ParseSqlNullTime(in.CreateTime),
		PackageQuantity: in.PackageQuantity,
		CurrentState:    in.CurrentState,
		// LastUpdateTime:            utils.ParseSqlNullTime(in.LastUpdateTime),
		Remark:                    in.Remark,
		ProductPackageStackRuleID: in.ProductPackageStackRuleID,
		ProductOrderID:            in.ProductOrderID,
	}
}

func ProductOrderPalletsToPB(in []*ProductOrderPallet) []*proto.ProductOrderPalletInfo {
	var list []*proto.ProductOrderPalletInfo
	for _, f := range in {
		list = append(list, ProductOrderPalletToPB(f))
	}
	return list
}

func ProductOrderPalletToPB(in *ProductOrderPallet) *proto.ProductOrderPalletInfo {
	if in == nil {
		return nil
	}

	m := &proto.ProductOrderPalletInfo{
		Id:                        in.ID,
		PalletNo:                  in.PalletNo,
		PalletSn:                  in.PalletSn,
		CreateTime:                utils.FormatTime(in.CreateTime),
		PackageQuantity:           in.PackageQuantity,
		CurrentState:              in.CurrentState,
		LastUpdateTime:            utils.FormatTime(in.LastUpdateTime),
		Remark:                    in.Remark,
		ProductPackageStackRuleID: in.ProductPackageStackRuleID,
		ProductPackageStackRule:   ProductPackageStackRuleToPB(in.ProductPackageStackRule),
		ProductOrderID:            in.ProductOrderID,
		ProductOrder:              ProductOrderToPB(in.ProductOrder),
	}
	return m
}
