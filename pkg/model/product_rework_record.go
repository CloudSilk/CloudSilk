package model

import (
	"database/sql"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 产品返工记录
type ProductReworkRecord struct {
	ModelID
	ProductionStationID     string             `json:"productionStationID" gorm:"size:36;comment:生产工站ID"`
	ProductionStation       *ProductionStation `json:"productionStation" gorm:"constraint:OnDelete:CASCADE"` //生产工站
	ProductionProcessID     string             `json:"productionProcessID" gorm:"size:36;comment:生产工序ID"`
	ProductInfoID           string             `json:"productInfoID" gorm:"size:36;comment:产品信息ID"`
	ProductInfo             *ProductInfo       `json:"productInfo" gorm:"constraint:OnDelete:CASCADE"` //产品信息
	CreateUserID            string             `json:"createUserID" gorm:"size:36;comment:返工人员ID"`
	CreateTime              time.Time          `json:"createTime" gorm:"autoCreateTime:nano;comment:创建时间"`
	ReworkReason            string             `json:"reworkReason" gorm:"size:1000;comment:返工原因"`
	ProductReworkTypeID     string             `json:"productReworkTypeID" gorm:"size:36;comment:返工类型ID"`
	ProductReworkCauseID    string             `json:"productReworkCauseID" gorm:"size:36;comment:返工原因ID"`
	ProductReworkSolutionID string             `json:"productReworkSolutionID" gorm:"size:36;comment:返工方案ID"`
	ReworkBrief             string             `json:"reworkBrief" gorm:"size:1000;comment:返工简述"`
	ReworkUserID            string             `json:"reworkUserID" gorm:"size:36;comment:返工人员ID"`
	ReworkTime              sql.NullTime       `json:"reworkTime" gorm:"comment:返工时间"`
	CompleteTime            sql.NullTime       `json:"completeTime" gorm:"comment:完成时间"`
	Remark                  string             `json:"remark" gorm:"size:1000;comment:备注"`
}

func PBToProductReworkRecords(in []*proto.ProductReworkRecordInfo) []*ProductReworkRecord {
	var result []*ProductReworkRecord
	for _, c := range in {
		result = append(result, PBToProductReworkRecord(c))
	}
	return result
}

func PBToProductReworkRecord(in *proto.ProductReworkRecordInfo) *ProductReworkRecord {
	if in == nil {
		return nil
	}
	return &ProductReworkRecord{
		ModelID:             ModelID{ID: in.Id},
		ProductionStationID: in.ProductionStationID,
		ProductionProcessID: in.ProductionProcessID,
		ProductInfoID:       in.ProductInfoID,
		CreateUserID:        in.CreateUserID,
		// CreateTime:              utils.ParseTime(in.CreateTime),
		ReworkReason:            in.ReworkReason,
		ProductReworkTypeID:     in.ProductReworkTypeID,
		ProductReworkCauseID:    in.ProductReworkCauseID,
		ProductReworkSolutionID: in.ProductReworkSolutionID,
		ReworkBrief:             in.ReworkBrief,
		ReworkUserID:            in.ReworkUserID,
		ReworkTime:              utils.ParseSqlNullTime(in.ReworkTime),
		CompleteTime:            utils.ParseSqlNullTime(in.CompleteTime),
		Remark:                  in.Remark,
	}
}

func ProductReworkRecordsToPB(in []*ProductReworkRecord) []*proto.ProductReworkRecordInfo {
	var list []*proto.ProductReworkRecordInfo
	for _, f := range in {
		list = append(list, ProductReworkRecordToPB(f))
	}
	return list
}

func ProductReworkRecordToPB(in *ProductReworkRecord) *proto.ProductReworkRecordInfo {
	if in == nil {
		return nil
	}

	m := &proto.ProductReworkRecordInfo{
		Id:                      in.ID,
		ProductionStationID:     in.ProductionStationID,
		ProductionStation:       ProductionStationToPB(in.ProductionStation),
		ProductionProcessID:     in.ProductionProcessID,
		ProductInfoID:           in.ProductInfoID,
		ProductInfo:             ProductInfoToPB(in.ProductInfo),
		CreateUserID:            in.CreateUserID,
		CreateTime:              utils.FormatTime(in.CreateTime),
		ReworkReason:            in.ReworkReason,
		ProductReworkTypeID:     in.ProductReworkTypeID,
		ProductReworkCauseID:    in.ProductReworkCauseID,
		ProductReworkSolutionID: in.ProductReworkSolutionID,
		ReworkBrief:             in.ReworkBrief,
		ReworkUserID:            in.ReworkUserID,
		ReworkTime:              utils.FormatSqlNullTime(in.ReworkTime),
		CompleteTime:            utils.FormatSqlNullTime(in.CompleteTime),
		Remark:                  in.Remark,
	}
	return m
}
