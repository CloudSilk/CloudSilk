package model

import (
	"database/sql"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 产品测试记录
type ProductTestRecord struct {
	ModelID
	ProductionStationID     string                 `json:"productionStationID" gorm:"size:36;comment:测试工站ID"`
	ProductionStation       *ProductionStation     `json:"productionStation" gorm:"constraint:OnDelete:CASCADE"` //测试工站
	ProductionProcessID     string                 `json:"productionProcessID" gorm:"size:36;comment:生产工序ID"`
	ProductionProcess       *ProductionProcess     `json:"productionProcess" gorm:"constraint:OnDelete:CASCADE"` //生产工序
	ProductInfoID           string                 `json:"productInfoID" gorm:"size:36;comment:产品信息ID"`
	ProductInfo             *ProductInfo           `json:"productInfo" gorm:"constraint:OnDelete:CASCADE"` //产品信息
	ProductionProcessStepID string                 `json:"productionProcessStepID" gorm:"size:36;comment:生产工步ID"`
	ProductionProcessStep   *ProductionProcessStep `json:"productionProcessStep" gorm:"constraint:OnDelete:CASCADE"` //生产工步
	TestStartTime           sql.NullTime           `json:"testStartTime" gorm:"comment:开始测试时间"`
	TestEndTime             sql.NullTime           `json:"testEndTime" gorm:"comment:结束测试时间"`
	Duration                int32                  `json:"duration" gorm:"comment:耗时(秒)"`
	TestData                string                 `json:"testData" gorm:"size:-1;comment:测试数据"`
	IsQualified             bool                   `json:"isQualified" gorm:"comment:是否合格"`
	TestUserID              string                 `json:"testUserID" gorm:"size:36;comment:测试人员ID"`
	CheckUserID             string                 `json:"checkUserID" gorm:"size:36;comment:复核人员ID"`
	Remark                  string                 `json:"remark" gorm:"size:1000;comment:备注"`
	// ProductTestConclusions  []*ProductTestConclusion `json:"productTestConclusions"` //测试结论
}

// 测试结论
// type ProductTestConclusion struct {
// 	ID              string            `json:"id" gorm:"primarykey;size:36"`
// 	ProductTestRecordID string             `json:"productTestRecordID" gorm:"size:36;comment:产品测试记录ID"`
// 	ProductTestRecord   *ProductTestRecord `json:"productTestRecord"`
// }

func PBToProductTestRecords(in []*proto.ProductTestRecordInfo) []*ProductTestRecord {
	var result []*ProductTestRecord
	for _, c := range in {
		result = append(result, PBToProductTestRecord(c))
	}
	return result
}

func PBToProductTestRecord(in *proto.ProductTestRecordInfo) *ProductTestRecord {
	if in == nil {
		return nil
	}
	return &ProductTestRecord{
		ModelID:                 ModelID{ID: in.Id},
		ProductionStationID:     in.ProductionStationID,
		ProductionProcessID:     in.ProductionProcessID,
		ProductInfoID:           in.ProductInfoID,
		ProductionProcessStepID: in.ProductionProcessStepID,
		TestStartTime:           utils.ParseSqlNullTime(in.TestStartTime),
		TestEndTime:             utils.ParseSqlNullTime(in.TestEndTime),
		Duration:                in.Duration,
		TestData:                in.TestData,
		IsQualified:             in.IsQualified,
		TestUserID:              in.TestUserID,
		CheckUserID:             in.CheckUserID,
		Remark:                  in.Remark,
	}
}

func ProductTestRecordsToPB(in []*ProductTestRecord) []*proto.ProductTestRecordInfo {
	var list []*proto.ProductTestRecordInfo
	for _, f := range in {
		list = append(list, ProductTestRecordToPB(f))
	}
	return list
}

func ProductTestRecordToPB(in *ProductTestRecord) *proto.ProductTestRecordInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductTestRecordInfo{
		Id:                      in.ID,
		ProductionStationID:     in.ProductionStationID,
		ProductionStation:       ProductionStationToPB(in.ProductionStation),
		ProductionProcessID:     in.ProductionProcessID,
		ProductionProcess:       ProductionProcessToPB(in.ProductionProcess),
		ProductInfoID:           in.ProductInfoID,
		ProductInfo:             ProductInfoToPB(in.ProductInfo),
		ProductionProcessStepID: in.ProductionProcessStepID,
		ProductionProcessStep:   ProductionProcessStepToPB(in.ProductionProcessStep),
		TestStartTime:           utils.FormatSqlNullTime(in.TestStartTime),
		TestEndTime:             utils.FormatSqlNullTime(in.TestEndTime),
		Duration:                in.Duration,
		TestData:                in.TestData,
		IsQualified:             in.IsQualified,
		TestUserID:              in.TestUserID,
		CheckUserID:             in.CheckUserID,
		Remark:                  in.Remark,
	}
	return m
}
