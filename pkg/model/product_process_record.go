package model

import (
	"database/sql"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

type ProductProcessRecord struct {
	ModelID
	ProcessStepType     string             `json:"processStepType" gorm:"size:100;comment:作业类型"`
	WorkDescription     string             `json:"workDescription" gorm:"size:1000;comment:作业描述"`
	WorkData            string             `json:"workData" gorm:"size:1000;comment:作业数据"`
	WorkResult          string             `json:"workResult" gorm:"size:100;comment:作业结果"`
	WorkTime            sql.NullTime       `json:"workTime" gorm:"comment:作业时间"`
	WorkUserID          string             `json:"workUserID" gorm:"size:36;comment:作业人员ID"`
	Remark              string             `json:"remark" gorm:"size:1000;comment:备注"`
	ProductionStationID string             `json:"productionStationID" gorm:"size:36;comment:生产工站ID"`
	ProductionStation   *ProductionStation `gorm:"constraint:OnDelete:CASCADE"`
	ProductionProcessID string             `json:"productionProcessID" gorm:"size:36;comment:生产工序ID"`
	ProductInfoID       string             `json:"productInfoID" gorm:"size:36;comment:产品信息ID"`
	ProductInfo         *ProductInfo       `gorm:"constraint:OnDelete:CASCADE"`
}

func PBToProductProcessRecords(in []*proto.ProductProcessRecordInfo) []*ProductProcessRecord {
	var result []*ProductProcessRecord
	for _, c := range in {
		result = append(result, PBToProductProcessRecord(c))
	}
	return result
}

func PBToProductProcessRecord(in *proto.ProductProcessRecordInfo) *ProductProcessRecord {
	if in == nil {
		return nil
	}
	return &ProductProcessRecord{
		ModelID:             ModelID{ID: in.Id},
		ProcessStepType:     in.ProcessStepType,
		WorkDescription:     in.WorkDescription,
		WorkData:            in.WorkData,
		WorkResult:          in.WorkResult,
		WorkTime:            utils.ParseSqlNullTime(in.WorkTime),
		WorkUserID:          in.WorkUserID,
		Remark:              in.Remark,
		ProductionStationID: in.ProductionStationID,
		ProductionProcessID: in.ProductionProcessID,
		ProductInfoID:       in.ProductInfoID,
	}
}

func ProductProcessRecordsToPB(in []*ProductProcessRecord) []*proto.ProductProcessRecordInfo {
	var list []*proto.ProductProcessRecordInfo
	for _, f := range in {
		list = append(list, ProductProcessRecordToPB(f))
	}
	return list
}

func ProductProcessRecordToPB(in *ProductProcessRecord) *proto.ProductProcessRecordInfo {
	if in == nil {
		return nil
	}
	productSerialNo := ""
	productionStation := ""
	productOrderNo := ""
	if in.ProductInfo != nil {
		productSerialNo = in.ProductInfo.ProductSerialNo
		if in.ProductInfo.ProductOrder != nil {
			productOrderNo = in.ProductInfo.ProductOrder.ProductOrderNo
		}
	}
	if in.ProductionStation != nil {
		productionStation = in.ProductionStation.Description
	}
	m := &proto.ProductProcessRecordInfo{
		Id:                  in.ID,
		ProcessStepType:     in.ProcessStepType,
		WorkDescription:     in.WorkDescription,
		WorkData:            in.WorkData,
		WorkResult:          in.WorkResult,
		WorkTime:            utils.FormatSqlNullTime(in.WorkTime),
		WorkUserID:          in.WorkUserID,
		Remark:              in.Remark,
		ProductionStationID: in.ProductionStationID,
		ProductionProcessID: in.ProductionProcessID,
		ProductInfoID:       in.ProductInfoID,
		ProductSerialNo:     productSerialNo,
		ProductionStation:   productionStation,
		ProductOrderNo:      productOrderNo,
	}
	return m
}
