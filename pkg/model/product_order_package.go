package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 产品工单包装
type ProductOrderPackage struct {
	ModelID
	PackageNo        string    `json:"packageNo" gorm:"size:100;comment:包装箱号"`
	SecurityCode     string    `json:"securityCode" gorm:"size:100;comment:防伪码"`
	SecurityUrl      string    `json:"securityUrl" gorm:"size:1000;comment:防伪链接"`
	PalletNo         string    `json:"palletNo" gorm:"size:100;comment:栈板标识"`
	CreateTime       time.Time `json:"createTime" gorm:"autoCreateTime:nano;comment:创建时间"`
	CurrentQuantity  int32     `json:"currentQuantity" gorm:"comment:当前包装数量"`
	PackageQuantity  int32     `json:"packageQuantity" gorm:"comment:实际包装数量"`
	GrossWeight      float32   `json:"grossWeight" gorm:"comment:实际毛重"`
	NetWeight        float32   `json:"netWeight" gorm:"comment:实际净重"`
	PrintTimes       int32     `json:"printTimes" gorm:"comment:打印次数"`
	CurrentState     string    `json:"currentState" gorm:"size:100;comment:当前状态"`
	LastUpdateTime   time.Time `json:"lastUpdateTime" gorm:"autoUpdateTime:nano;comment:状态变更时间"`
	Remark           string    `json:"remark" gorm:"size:1000;comment:备注"`
	ProductPackageID string    `json:"productPackageID" gorm:"size:36;comment:使用包装ID"`
	ProductOrderID   string    `json:"productOrderID" gorm:"size:36;comment:归属工单ID"`
	ParentID         string    `json:"parentID" gorm:"size:36;comment:上级包装ID"`
}

func PBToProductOrderPackages(in []*proto.ProductOrderPackageInfo) []*ProductOrderPackage {
	var result []*ProductOrderPackage
	for _, c := range in {
		result = append(result, PBToProductOrderPackage(c))
	}
	return result
}

func PBToProductOrderPackage(in *proto.ProductOrderPackageInfo) *ProductOrderPackage {
	if in == nil {
		return nil
	}
	return &ProductOrderPackage{
		ModelID:      ModelID{ID: in.Id},
		PackageNo:    in.PackageNo,
		SecurityCode: in.SecurityCode,
		SecurityUrl:  in.SecurityUrl,
		PalletNo:     in.PalletNo,
		// CreateTime:       utils.ParseTime(in.CreateTime),
		CurrentQuantity: in.CurrentQuantity,
		PackageQuantity: in.PackageQuantity,
		GrossWeight:     in.GrossWeight,
		NetWeight:       in.NetWeight,
		PrintTimes:      in.PrintTimes,
		CurrentState:    in.CurrentState,
		// LastUpdateTime:   utils.ParseTime(in.LastUpdateTime),
		Remark:           in.Remark,
		ProductPackageID: in.ProductPackageID,
		ProductOrderID:   in.ProductOrderID,
		ParentID:         in.ParentID,
	}
}

func ProductOrderPackagesToPB(in []*ProductOrderPackage) []*proto.ProductOrderPackageInfo {
	var list []*proto.ProductOrderPackageInfo
	for _, f := range in {
		list = append(list, ProductOrderPackageToPB(f))
	}
	return list
}

func ProductOrderPackageToPB(in *ProductOrderPackage) *proto.ProductOrderPackageInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductOrderPackageInfo{
		Id:               in.ID,
		PackageNo:        in.PackageNo,
		SecurityCode:     in.SecurityCode,
		SecurityUrl:      in.SecurityUrl,
		PalletNo:         in.PalletNo,
		CreateTime:       utils.FormatTime(in.CreateTime),
		CurrentQuantity:  in.CurrentQuantity,
		PackageQuantity:  in.PackageQuantity,
		GrossWeight:      in.GrossWeight,
		NetWeight:        in.NetWeight,
		PrintTimes:       in.PrintTimes,
		CurrentState:     in.CurrentState,
		LastUpdateTime:   utils.FormatTime(in.LastUpdateTime),
		Remark:           in.Remark,
		ProductPackageID: in.ProductPackageID,
		ProductOrderID:   in.ProductOrderID,
		ParentID:         in.ParentID,
	}
	return m
}
