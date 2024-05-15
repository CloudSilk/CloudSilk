package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

type ProductReleaseRecord struct {
	ModelID
	CreateTime                  time.Time    `json:"createTime" gorm:"autoCreateTime:nano;comment:发料时间"`
	CreateUserID                string       `json:"createUserID" gorm:"size:36;comment:投料人员ID"`
	CurrentState                string       `json:"currentState" gorm:"size:-1;comment:当前状态"`
	LastUpdateTime              time.Time    `json:"lastUpdateTime" gorm:"autoUpdateTime:nano;comment:状态变更时间"`
	Remark                      string       `json:"remark" gorm:"size:-1;comment:备注"`
	PrefabricatedProductInfoID  string       `json:"prefabricatedProductInfoID" gorm:"size:36;comment:预制产品ID"`
	ProductInfo                 *ProductInfo `json:"productInfo" gorm:"foreignKey:PrefabricatedProductInfoID"`
	RemanufacturedProductInfoID string       `json:"remanufacturedProductInfoID" gorm:"size:36;comment:改制产品ID"`
	RemanufacturedProductInfo   *ProductInfo `json:"remanufacturedProductInfo" gorm:"foreignKey:RemanufacturedProductInfoID"`
}

func PBToProductReleaseRecords(in []*proto.ProductReleaseRecordInfo) []*ProductReleaseRecord {
	var result []*ProductReleaseRecord
	for _, c := range in {
		result = append(result, PBToProductReleaseRecord(c))
	}
	return result
}

func PBToProductReleaseRecord(in *proto.ProductReleaseRecordInfo) *ProductReleaseRecord {
	if in == nil {
		return nil
	}
	return &ProductReleaseRecord{
		ModelID: ModelID{ID: in.Id},
		// CreateTime:                  utils.ParseSqlNullTime(in.CreateTime),
		CreateUserID: in.CreateUserID,
		CurrentState: in.CurrentState,
		// LastUpdateTime:              utils.ParseSqlNullTime(in.LastUpdateTime),
		Remark:                      in.Remark,
		PrefabricatedProductInfoID:  in.PrefabricatedProductInfoID,
		RemanufacturedProductInfoID: in.RemanufacturedProductInfoID,
	}
}

func ProductReleaseRecordsToPB(in []*ProductReleaseRecord) []*proto.ProductReleaseRecordInfo {
	var list []*proto.ProductReleaseRecordInfo
	for _, f := range in {
		list = append(list, ProductReleaseRecordToPB(f))
	}
	return list
}

func ProductReleaseRecordToPB(in *ProductReleaseRecord) *proto.ProductReleaseRecordInfo {
	if in == nil {
		return nil
	}
	productOrderNo := ""
	productSerialNo := ""
	if in.RemanufacturedProductInfo != nil {
		productSerialNo = in.RemanufacturedProductInfo.ProductSerialNo
		if in.RemanufacturedProductInfo.ProductOrder != nil {
			productOrderNo = in.RemanufacturedProductInfo.ProductOrder.ProductOrderNo
		}
	}
	securityCode := ""
	if in.ProductInfo != nil {
		securityCode = in.ProductInfo.ProductSerialNo
	}
	m := &proto.ProductReleaseRecordInfo{
		Id:                          in.ID,
		CreateTime:                  utils.FormatTime(in.CreateTime),
		CreateUserID:                in.CreateUserID,
		CurrentState:                in.CurrentState,
		LastUpdateTime:              utils.FormatTime(in.LastUpdateTime),
		Remark:                      in.Remark,
		PrefabricatedProductInfoID:  in.PrefabricatedProductInfoID,
		RemanufacturedProductInfoID: in.RemanufacturedProductInfoID,
		ProductOrderNo:              productOrderNo,
		ProductSerialNo:             productSerialNo,
		SecurityCode:                securityCode,
	}
	return m
}
