package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 工位效率统计
type ProductionStationEfficiency struct {
	ModelID
	OutputDate                time.Time          `gorm:"type:date;comment:产出日期"`
	OutputOfThisDay           int32              `gorm:"comment:当日产量"`
	NumberOfWork              int32              `gorm:"comment:作业次数"`
	NumberOfPass              int32              `gorm:"comment:合格次数"`
	NumberOfFail              int32              `gorm:"comment:失败次数"`
	FPY                       float32            `gorm:"comment:一次通过率"`
	EstimateRhythm            float32            `gorm:"comment:理论节拍"`
	AverageRhythm             float32            `gorm:"comment:平均节拍"`
	StartupDuration           float32            `gorm:"comment:开机时长"`
	PlannedShutdownDuration   float32            `gorm:"comment:计划内停机时长"`
	EstimateAvailableDuration float32            `gorm:"comment:预计稼动时长"`
	NumberOfBreakdown         int32              `gorm:"comment:故障次数"`
	BreakdownDuration         float32            `gorm:"comment:故障时长"`
	NumberOfShutdown          int32              `gorm:"comment:计划外停机次数"`
	UnplannedShutdownDuration float32            `gorm:"comment:计划外停机时长"`
	OEF                       float32            `gorm:"comment:设备故障率"`
	ActualAvailableDuration   float32            `gorm:"comment:实际稼动时长"`
	ActualEffectiveDuration   float32            `gorm:"comment:实际作业时长"`
	OEU                       float32            `gorm:"comment:时间稼动率"`
	OEP                       float32            `gorm:"comment:性能稼动率"`
	OEE                       float32            `gorm:"comment:设备综合效率"`
	LastUpdateTime            time.Time          `gorm:"autoUpdateTime:nano;comment:最后更新时间"`
	ProductionStationID       string             `gorm:"size:36;comment:生产工站ID"`
	ProductionStation         *ProductionStation `gorm:"constraint:OnDelete:CASCADE"`
}

func PBToProductionStationEfficiencys(in []*proto.ProductionStationEfficiencyInfo) []*ProductionStationEfficiency {
	var result []*ProductionStationEfficiency
	for _, c := range in {
		result = append(result, PBToProductionStationEfficiency(c))
	}
	return result
}

func PBToProductionStationEfficiency(in *proto.ProductionStationEfficiencyInfo) *ProductionStationEfficiency {
	if in == nil {
		return nil
	}
	return &ProductionStationEfficiency{
		ModelID:                   ModelID{ID: in.Id},
		OutputOfThisDay:           in.OutputOfThisDay,
		NumberOfWork:              in.NumberOfWork,
		NumberOfPass:              in.NumberOfPass,
		NumberOfFail:              in.NumberOfFail,
		FPY:                       in.FPY,
		EstimateRhythm:            in.EstimateRhythm,
		AverageRhythm:             in.AverageRhythm,
		StartupDuration:           in.StartupDuration,
		PlannedShutdownDuration:   in.PlannedShutdownDuration,
		EstimateAvailableDuration: in.EstimateAvailableDuration,
		NumberOfBreakdown:         in.NumberOfBreakdown,
		BreakdownDuration:         in.BreakdownDuration,
		NumberOfShutdown:          in.NumberOfShutdown,
		UnplannedShutdownDuration: in.UnplannedShutdownDuration,
		OEF:                       in.OEF,
		ActualAvailableDuration:   in.ActualAvailableDuration,
		ActualEffectiveDuration:   in.ActualEffectiveDuration,
		OEU:                       in.OEU,
		OEP:                       in.OEP,
		OEE:                       in.OEE,
		ProductionStationID:       in.ProductionStationID,
	}
}

func ProductionStationEfficiencysToPB(in []*ProductionStationEfficiency) []*proto.ProductionStationEfficiencyInfo {
	var list []*proto.ProductionStationEfficiencyInfo
	for _, f := range in {
		list = append(list, ProductionStationEfficiencyToPB(f))
	}
	return list
}

func ProductionStationEfficiencyToPB(in *ProductionStationEfficiency) *proto.ProductionStationEfficiencyInfo {
	if in == nil {
		return nil
	}

	m := &proto.ProductionStationEfficiencyInfo{
		Id:                        in.ID,
		OutputDate:                utils.FormatTime(in.OutputDate),
		OutputOfThisDay:           in.OutputOfThisDay,
		NumberOfWork:              in.NumberOfWork,
		NumberOfPass:              in.NumberOfPass,
		NumberOfFail:              in.NumberOfFail,
		FPY:                       in.FPY,
		EstimateRhythm:            in.EstimateRhythm,
		AverageRhythm:             in.AverageRhythm,
		StartupDuration:           in.StartupDuration,
		PlannedShutdownDuration:   in.PlannedShutdownDuration,
		EstimateAvailableDuration: in.EstimateAvailableDuration,
		NumberOfBreakdown:         in.NumberOfBreakdown,
		BreakdownDuration:         in.BreakdownDuration,
		NumberOfShutdown:          in.NumberOfShutdown,
		UnplannedShutdownDuration: in.UnplannedShutdownDuration,
		OEF:                       in.OEF,
		ActualAvailableDuration:   in.ActualAvailableDuration,
		ActualEffectiveDuration:   in.ActualEffectiveDuration,
		OEU:                       in.OEU,
		OEP:                       in.OEP,
		OEE:                       in.OEE,
		LastUpdateTime:            utils.FormatTime(in.LastUpdateTime),
		ProductionStationID:       in.ProductionStationID,
		ProductionStation:         ProductionStationToPB(in.ProductionStation),
	}
	return m
}
