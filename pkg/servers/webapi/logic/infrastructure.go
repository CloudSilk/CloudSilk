package logic

import (
	"fmt"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// GetAllProductionLine 获取全部产线信息
func GetAllProductionLine() ([]map[string]interface{}, error) {
	productionLines := []*model.ProductionLine{}
	if err := model.DB.DB().Preload("ProductionStations").
		Preload("ProductionStations.AvailableProductionProcesses").
		Preload("ProductionStations.AvailableProductionProcesses.ProductionProcess").
		Find(&productionLines).Error; err != nil {
		return nil, err
	}

	data := make([]map[string]interface{}, len(productionLines))
	for pli, pl := range productionLines {
		productionStations := make([]map[string]interface{}, len(pl.ProductionStations))
		for psi, ps := range pl.ProductionStations {
			productionProcesses := []map[string]interface{}{}
			for _, pp := range ps.AvailableProductionProcesses {
				if pp.ProductionStationID == ps.ID && pp.ProductionProcess.Enable {
					productionProcesses = append(productionProcesses, map[string]interface{}{
						"id":            pp.ProductionProcess.ID,
						"code":          pp.ProductionProcess.Code,
						"description":   pp.ProductionProcess.Description,
						"processType":   pp.ProductionProcess.ProcessType,
						"vehicleType":   pp.ProductionProcess.VehicleType,
						"enableControl": pp.ProductionProcess.EnableControl,
					})
				}
			}

			productionStation := map[string]interface{}{
				"id":                  ps.ID,
				"code":                ps.Code,
				"description":         ps.Description,
				"stationType":         ps.StationType,
				"accountControl":      ps.AccountControl,
				"materialControl":     ps.MaterialControl,
				"productionProcesses": productionProcesses,
			}

			productionStations[psi] = productionStation
		}

		productionLine := map[string]interface{}{
			"id":                 pl.ID,
			"code":               pl.Code,
			"description":        pl.Description,
			"accountControl":     pl.AccountControl,
			"materialControl":    pl.MaterialControl,
			"productionStations": productionStations,
		}

		data[pli] = productionLine
	}

	return data, nil
}

// RetrieveProductionStation 查询产线工位信息
func RetrieveProductionStation(req *proto.RetrieveProductionStationRequest) ([]map[string]interface{}, error) {
	if req.ProductionLine == "" {
		return nil, fmt.Errorf("ProductionLine不能为空")
	}

	whereMap := map[string]interface{}{"production_line_id": req.ProductionLine}
	if req.StationType != "" {
		whereMap["station_type"] = req.StationType
	}

	productionStations := []*model.ProductionStation{}
	if err := model.DB.DB().Find(&productionStations, whereMap).Error; err != nil {
		return nil, err
	}

	data := []map[string]interface{}{}
	for _, s := range productionStations {
		data = append(data, map[string]interface{}{
			"id":              s.ID,
			"code":            s.Code,
			"description":     s.Description,
			"stationType":     s.StationType,
			"accountControl":  s.AccountControl,
			"materialControl": s.MaterialControl,
			"currentState":    s.CurrentState,
		})
	}

	return data, nil
}

// RetrieveProductAttribute 查询产品特性信息
func RetrieveProductAttribute(req *proto.RetrieveProductAttributeRequest) ([]map[string]interface{}, error) {
	if req.Code == "" {
		return nil, fmt.Errorf("Code不能为空")
	}

	var productAttributes []*model.ProductAttribute
	if err := model.DB.DB().Where(model.ProductAttribute{Code: req.Code, Description: req.Description}).Find(&productAttributes).Error; err != nil {
		return nil, err
	}

	data := make([]map[string]interface{}, len(productAttributes))
	for pai, pa := range productAttributes {
		data[pai] = map[string]interface{}{
			"code":        pa.Code,
			"description": pa.Description,
		}
	}

	return data, nil
}

// RetrieveProductionCrossway 查询产线路口信息
func RetrieveProductionCrossway(req *proto.RetrieveProductionCrosswayRequest) ([]map[string]interface{}, error) {
	if req.ProductionLine == "" {
		return nil, fmt.Errorf("ProductionLine不能为空")
	}

	var productionCrossways []*model.ProductionCrossway
	if err := model.DB.DB().Where(model.ProductionCrossway{ProductionLineID: req.ProductionLine}).Find(&productionCrossways).Error; err != nil {
		return nil, err
	}

	data := make([]map[string]interface{}, len(productionCrossways))
	for i, v := range productionCrossways {
		data[i] = map[string]interface{}{
			"id":          v.ID,
			"code":        v.Code,
			"description": v.Description,
			"defaultTurn": v.DefaultTurn,
			"remark":      v.Remark,
		}
	}

	return data, nil
}
