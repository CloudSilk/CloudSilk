package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 产品包装记录
type ProductPackageRecord struct {
	ModelID
	CreateTime            time.Time            `json:"createTime" gorm:"autoCreateTime:nano;comment:装箱时间"`
	CreateUserID          string               `json:"createUserID" gorm:"size:36;comment:装箱人员ID"`
	ProductOrderPackageID string               `json:"productOrderPackageID" gorm:"size:36;comment:使用包装ID"`
	ProductOrderPackage   *ProductOrderPackage `json:"productOrderPackage"` //包装
	ProductInfoID         string               `json:"productInfoID" gorm:"size:36;comment:绑定产品ID"`
	ProductInfo           *ProductInfo         `json:"productInfo"` //产品
	Remark                string               `json:"remark" gorm:"size:1000;comment:备注"`
}

func PBToProductPackageRecords(in []*proto.ProductPackageRecordInfo) []*ProductPackageRecord {
	var result []*ProductPackageRecord
	for _, c := range in {
		result = append(result, PBToProductPackageRecord(c))
	}
	return result
}

func PBToProductPackageRecord(in *proto.ProductPackageRecordInfo) *ProductPackageRecord {
	if in == nil {
		return nil
	}
	return &ProductPackageRecord{
		ModelID: ModelID{ID: in.Id},
		// CreateTime:            utils.ParseTime(in.CreateTime),
		CreateUserID:          in.CreateUserID,
		ProductOrderPackageID: in.ProductOrderPackageID,
		ProductInfoID:         in.ProductInfoID,
		Remark:                in.Remark,
	}
}

func ProductPackageRecordsToPB(in []*ProductPackageRecord) []*proto.ProductPackageRecordInfo {
	var list []*proto.ProductPackageRecordInfo
	for _, f := range in {
		list = append(list, ProductPackageRecordToPB(f))
	}
	return list
}

func ProductPackageRecordToPB(in *ProductPackageRecord) *proto.ProductPackageRecordInfo {
	if in == nil {
		return nil
	}

	m := &proto.ProductPackageRecordInfo{
		Id:                    in.ID,
		CreateTime:            utils.FormatTime(in.CreateTime),
		CreateUserID:          in.CreateUserID,
		ProductOrderPackageID: in.ProductOrderPackageID,
		ProductOrderPackage:   ProductOrderPackageToPB(in.ProductOrderPackage),
		ProductInfoID:         in.ProductInfoID,
		ProductInfo:           ProductInfoToPB(in.ProductInfo),
		Remark:                in.Remark,
	}
	return m
}
