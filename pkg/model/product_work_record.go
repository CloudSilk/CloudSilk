package model

import (
	"database/sql"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

type ProductWorkRecord struct {
	ModelID
	ProductionStationID     string                 `json:"ProductionStationID" gorm:"size:36;comment:生产工站ID"`
	ProductionStation       *ProductionStation     ` gorm:"constraint:OnDelete:CASCADE"`
	ProductionProcessStepID string                 `json:"ProductionProcessStepID" gorm:"size:36;comment:作业步骤ID"`
	ProductionProcessStep   *ProductionProcessStep ` gorm:"constraint:OnDelete:CASCADE"`
	ProductInfoID           string                 `json:"ProductInfoID" gorm:"size:36;comment:产品信息ID"`
	ProductInfo             *ProductInfo           ` gorm:"constraint:OnDelete:CASCADE"`
	WorkStartTime           sql.NullTime           `json:"WorkStartTime" gorm:"comment:开始作业时间"`
	WorkEndTime             sql.NullTime           `json:"WorkEndTime" gorm:"comment:结束作业时间"`
	Duration                int32                  `json:"Duration" gorm:"comment:耗时(秒)"`
	WorkData                string                 `json:"WorkData" gorm:"size:-1;comment:作业数据"`
	IsQualified             bool                   `json:"IsQualified" gorm:"comment:是否合格"`
	WorkUserID              string                 `json:"WorkUserID" gorm:"size:36;omment:作业人员ID"`
	Remark                  string                 `json:"Remark" gorm:"size:1000;comment:备注"`
}

func PBToProductWorkRecords(in []*proto.ProductWorkRecordInfo) []*ProductWorkRecord {
	var result []*ProductWorkRecord
	for _, c := range in {
		result = append(result, PBToProductWorkRecord(c))
	}
	return result
}

func PBToProductWorkRecord(in *proto.ProductWorkRecordInfo) *ProductWorkRecord {
	if in == nil {
		return nil
	}
	return &ProductWorkRecord{
		ModelID:                 ModelID{ID: in.Id},
		ProductionStationID:     in.ProductionStationID,
		ProductionProcessStepID: in.ProductionProcessStepID,
		ProductInfoID:           in.ProductInfoID,
		WorkStartTime:           utils.ParseSqlNullTime(in.WorkStartTime),
		WorkEndTime:             utils.ParseSqlNullTime(in.WorkEndTime),
		Duration:                in.Duration,
		WorkData:                in.WorkData,
		IsQualified:             in.IsQualified,
		WorkUserID:              in.WorkUserID,
		Remark:                  in.Remark,
	}
}

func ProductWorkRecordsToPB(in []*ProductWorkRecord) []*proto.ProductWorkRecordInfo {
	var list []*proto.ProductWorkRecordInfo
	for _, f := range in {
		list = append(list, ProductWorkRecordToPB(f))
	}
	return list
}

func ProductWorkRecordToPB(in *ProductWorkRecord) *proto.ProductWorkRecordInfo {
	if in == nil {
		return nil
	}
	productOrderNo := ""
	productSerialNo := ""
	processStepType := ""
	productionProcessStep := ""
	productionLine := ""
	productionStation := ""
	if in.ProductInfo != nil {
		productSerialNo = in.ProductInfo.ProductSerialNo
		if in.ProductInfo.ProductOrder != nil {
			productOrderNo = in.ProductInfo.ProductOrder.ProductOrderNo
		}
	}
	if in.ProductionProcessStep != nil {
		productionProcessStep = in.ProductionProcessStep.Code
		if in.ProductionProcessStep.ProcessStepType != nil {
			processStepType = in.ProductionProcessStep.ProcessStepType.Description
		}
	}
	if in.ProductionStation != nil {
		productionStation = in.ProductionStation.Description
		if in.ProductionStation.ProductionLine != nil {
			productionLine = in.ProductionStation.ProductionLine.Description
		}
	}
	m := &proto.ProductWorkRecordInfo{
		Id:                      in.ID,
		ProductionStationID:     in.ProductionStationID,
		ProductionProcessStepID: in.ProductionProcessStepID,
		ProductInfoID:           in.ProductInfoID,
		WorkStartTime:           utils.FormatSqlNullTime(in.WorkStartTime),
		WorkEndTime:             utils.FormatSqlNullTime(in.WorkEndTime),
		Duration:                in.Duration,
		WorkData:                in.WorkData,
		IsQualified:             in.IsQualified,
		WorkUserID:              in.WorkUserID,
		Remark:                  in.Remark,
		ProductOrderNo:          productOrderNo,
		ProductSerialNo:         productSerialNo,
		ProcessStepType:         processStepType,
		ProductionProcessStep:   productionProcessStep,
		ProductionLine:          productionLine,
		ProductionStation:       productionStation,
	}
	return m
}
