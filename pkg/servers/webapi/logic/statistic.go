package logic

import (
	"fmt"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
)

// 获取工位效率统计
func GetProductionStationEfficiency(req *proto.GetProductionStationEfficiencyRequest) (map[string]interface{}, error) {
	if req.ProductionStation == "" {
		return nil, fmt.Errorf("ProductionStation不能为空")
	}

	productionStation := &model.ProductionStation{}
	if err := model.DB.DB().First(productionStation, "`code` = ?", req.ProductionStation).Error; err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("无效的工位代号")
	} else if err != nil {
		return nil, err
	}

	outputDate := time.Now()
	if req.OutputDate != "" {
		outputDate = utils.ParseDate(req.OutputDate)
	}

	productionStationEfficiency := &model.ProductionStationEfficiency{}
	if err := model.DB.DB().Where("`production_station_id` = ? AND `output_date` = ?", productionStation.ID, outputDate.Format("2006-01-02")).First(productionStationEfficiency).Error; err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("无法找到此工站的效率数据")
	} else if err != nil {
		return nil, err
	}

	var outputOfThisMonth int32
	var outputOfThisMonths []int32
	if err := model.DB.DB().Model(model.ProductionStationEfficiency{}).Where("`production_station_id` = ? AND EXTRACT(MONTH FROM `output_date`) = ?", productionStation.ID, outputDate.Month()).Select("output_of_this_day").Find(&outputOfThisMonths).Error; err != nil {
		return nil, err
	}
	for _, v := range outputOfThisMonths {
		outputOfThisMonth += v
	}

	var outputOfThisYear int32
	var outputOfThisYears []int32
	if err := model.DB.DB().Model(model.ProductionStationEfficiency{}).Where("`production_station_id` = ? AND EXTRACT(YEAR FROM `output_date`) = ?", productionStation.ID, outputDate.Year()).Select("output_of_this_day").Find(&outputOfThisYears).Error; err != nil {
		return nil, err
	}
	for _, v := range outputOfThisYears {
		outputOfThisYear += v
	}

	return map[string]interface{}{
		"productionStation":         productionStation.Description,
		"outputDate":                productionStationEfficiency.OutputDate,
		"outputOfThisDay":           productionStationEfficiency.OutputOfThisDay,
		"outputOfThisMonth":         outputOfThisMonth,
		"outputOfThisYear":          outputOfThisYear,
		"numberOfWork":              productionStationEfficiency.NumberOfWork,
		"numberOfPass":              productionStationEfficiency.NumberOfPass,
		"numberOfFail":              productionStationEfficiency.NumberOfFail,
		"fPY":                       productionStationEfficiency.FPY,
		"estimateRhythm":            productionStationEfficiency.EstimateRhythm,
		"averageRhythm":             productionStationEfficiency.AverageRhythm,
		"startupDuration":           productionStationEfficiency.StartupDuration,
		"plannedShutdownDuration":   productionStationEfficiency.PlannedShutdownDuration,
		"estimateAvailableDuration": productionStationEfficiency.EstimateAvailableDuration,
		"numberOfBreakdown":         productionStationEfficiency.NumberOfBreakdown,
		"breakdownDuration":         productionStationEfficiency.BreakdownDuration,
		"numberOfShutdown":          productionStationEfficiency.NumberOfShutdown,
		"unplannedShutdownDuration": productionStationEfficiency.UnplannedShutdownDuration,
		"oEF":                       productionStationEfficiency.OEF,
		"actualAvailableDuration":   productionStationEfficiency.ActualAvailableDuration,
		"actualEffectiveDuration":   productionStationEfficiency.ActualEffectiveDuration,
		"oEU":                       productionStationEfficiency.OEU,
		"oEP":                       productionStationEfficiency.OEP,
		"oEE":                       productionStationEfficiency.OEE,
		"lastUpdateTime":            productionStationEfficiency.LastUpdateTime,
		"queryTime":                 utils.FormatTime(time.Now()),
	}, nil
}

// 更新工位效率统计
func UpdateProductionStationEfficiency(req *proto.UpdateProductionStationEfficiencyRequest) error {
	if req.ProductionStation == "" {
		return fmt.Errorf("ProductionStation不能为空")
	}

	productionStation := &model.ProductionStation{}
	if err := model.DB.DB().First(productionStation, "`code` = ?", req.ProductionStation).Error; err == gorm.ErrRecordNotFound {
		return fmt.Errorf("无效的工位代号")
	} else if err != nil {
		return err
	}

	outputDate := time.Now().Format("2006-01-02")
	if req.OutputDate != "" {
		outputDate = req.OutputDate
	}

	productionStationEfficiency := &model.ProductionStationEfficiency{}
	if err := model.DB.DB().Where("`production_station_id` = ? AND `output_date` = ?", productionStation.ID, outputDate).First(productionStationEfficiency).Error; err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if productionStationEfficiency.ID == "" {
		productionStationEfficiency.ProductionStationID = productionStation.ID
		productionStationEfficiency.OutputDate = utils.ParseDate(outputDate)
	}

	productionStationEfficiency.OutputOfThisDay = req.OutputOfThisDay
	productionStationEfficiency.NumberOfWork = req.NumberOfWork
	productionStationEfficiency.NumberOfPass = req.NumberOfPass
	productionStationEfficiency.NumberOfFail = req.NumberOfFail
	productionStationEfficiency.FPY = req.FPY / 100
	productionStationEfficiency.EstimateRhythm = req.EstimateRhythm
	productionStationEfficiency.AverageRhythm = req.AverageRhythm
	productionStationEfficiency.StartupDuration = req.StartupDuration
	productionStationEfficiency.PlannedShutdownDuration = req.PlannedShutdownDuration
	productionStationEfficiency.EstimateAvailableDuration = req.EstimateAvailableDuration
	productionStationEfficiency.NumberOfBreakdown = req.NumberOfBreakdown
	productionStationEfficiency.BreakdownDuration = req.BreakdownDuration
	productionStationEfficiency.NumberOfShutdown = req.NumberOfShutdown
	productionStationEfficiency.UnplannedShutdownDuration = req.UnplannedShutdownDuration
	productionStationEfficiency.OEF = req.OEF / 100
	productionStationEfficiency.ActualAvailableDuration = req.ActualAvailableDuration
	productionStationEfficiency.ActualEffectiveDuration = req.ActualEffectiveDuration
	productionStationEfficiency.OEU = req.OEU / 100
	productionStationEfficiency.OEP = req.OEP / 100
	productionStationEfficiency.OEE = req.OEE / 100

	return model.DB.DB().Save(productionStationEfficiency).Error
}
