package model

import (
	"database/sql"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 产品节拍记录
type ProductRhythmRecord struct {
	ModelID
	CreateTime          time.Time          `json:"createTime" gorm:"autoCreateTime:nano;comment:创建时间"`
	StandardWorkTime    int32              `json:"standardWorkTime" gorm:"comment:标准时长(秒)"`
	WorkUserID          string             `json:"workUserID" gorm:"comment:作业人员ID"`
	WorkStartTime       sql.NullTime       `json:"workStartTime" gorm:"comment:开始作业时间"`
	WaitTime            int32              `json:"waitTime" gorm:"comment:等待时长(秒)"`
	WorkTime            int32              `json:"workTime" gorm:"comment:作业时长(秒)"`
	OverTime            int32              `json:"overTime" gorm:"comment:超时时长(秒)"`
	WorkEndTime         sql.NullTime       `json:"workEndTime" gorm:"comment:结束作业时间"`
	IsRework            bool               `json:"isRework" gorm:"comment:是否返工"`
	Remark              string             `json:"remark" gorm:"size:1000;comment:备注"`
	ProductionStationID string             `json:"productionStationID" gorm:"comment:生产工站ID"`
	ProductionStation   *ProductionStation `json:"productionStation"` //生产工站
	ProductionProcessID string             `json:"productionProcessID" gorm:"comment:生产工序ID"`
	ProductInfoID       string             `json:"productInfoID" gorm:"comment:产品信息ID"`
	ProductInfo         *ProductInfo
}

func PBToProductRhythmRecords(in []*proto.ProductRhythmRecordInfo) []*ProductRhythmRecord {
	var result []*ProductRhythmRecord
	for _, c := range in {
		result = append(result, PBToProductRhythmRecord(c))
	}
	return result
}

func PBToProductRhythmRecord(in *proto.ProductRhythmRecordInfo) *ProductRhythmRecord {
	if in == nil {
		return nil
	}
	return &ProductRhythmRecord{
		ModelID: ModelID{ID: in.Id},
		// CreateTime:          utils.ParseTime(in.CreateTime),
		StandardWorkTime:    in.StandardWorkTime,
		WorkUserID:          in.WorkUserID,
		WorkStartTime:       utils.ParseSqlNullTime(in.WorkStartTime),
		WaitTime:            in.WaitTime,
		WorkTime:            in.WorkTime,
		OverTime:            in.OverTime,
		WorkEndTime:         utils.ParseSqlNullTime(in.WorkEndTime),
		IsRework:            in.IsRework,
		Remark:              in.Remark,
		ProductionStationID: in.ProductionStationID,
		ProductionProcessID: in.ProductionProcessID,
		ProductInfoID:       in.ProductInfoID,
	}
}

func ProductRhythmRecordsToPB(in []*ProductRhythmRecord) []*proto.ProductRhythmRecordInfo {
	var list []*proto.ProductRhythmRecordInfo
	for _, f := range in {
		list = append(list, ProductRhythmRecordToPB(f))
	}
	return list
}

func ProductRhythmRecordToPB(in *ProductRhythmRecord) *proto.ProductRhythmRecordInfo {
	if in == nil {
		return nil
	}
	productOrderNo := ""
	productSerialNo := ""
	productionStationDescription := ""
	productionProcessDescription := ""
	if in.ProductionStation != nil {
		productionStationDescription = in.ProductionStation.Description
	}
	// if in.ProductionProcess != nil {
	// 	productionProcessDescription = in.ProductionProcess.Description
	// }
	// if in.ProductInfo != nil {
	// 	productSerialNo = in.ProductInfo.ProductSerialNo
	// 	if in.ProductInfo.ProductOrder != nil {
	// 		productOrderNo = in.ProductInfo.ProductOrder.ProductOrderNo
	// 	}
	// }
	m := &proto.ProductRhythmRecordInfo{
		Id:                           in.ID,
		CreateTime:                   utils.FormatTime(in.CreateTime),
		StandardWorkTime:             in.StandardWorkTime,
		WorkUserID:                   in.WorkUserID,
		WorkStartTime:                utils.FormatSqlNullTime(in.WorkStartTime),
		WaitTime:                     in.WaitTime,
		WorkTime:                     in.WorkTime,
		OverTime:                     in.OverTime,
		WorkEndTime:                  utils.FormatSqlNullTime(in.WorkEndTime),
		IsRework:                     in.IsRework,
		Remark:                       in.Remark,
		ProductionStationID:          in.ProductionStationID,
		ProductionProcessID:          in.ProductionProcessID,
		ProductInfoID:                in.ProductInfoID,
		ProductOrderNo:               productOrderNo,
		ProductSerialNo:              productSerialNo,
		ProductionStationDescription: productionStationDescription,
		ProductionProcessDescription: productionProcessDescription,
	}
	return m
}
