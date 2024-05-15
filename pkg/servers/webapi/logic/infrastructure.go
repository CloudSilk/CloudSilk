package logic

import (
	"context"
	"fmt"

	"github.com/CloudSilk/CloudSilk/pkg/clients"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// GetAllProductionLine 获取全部产线信息
func GetAllProductionLine() ([]*proto.ProductionLineInfo, error) {
	_productionLine, err := clients.ProductionLineClient.GetAll(context.Background(), &proto.GetAllRequest{})
	if err != nil {
		return nil, err
	}
	productionLines := _productionLine.Data

	data := make([]*proto.ProductionLineInfo, len(productionLines))
	for pli, pl := range productionLines {
		productionLine := &proto.ProductionLineInfo{
			Id:              pl.Id,
			Code:            pl.Code,
			Description:     pl.Description,
			AccountControl:  pl.AccountControl,
			MaterialControl: pl.MaterialControl,
		}

		productionStations := make([]*proto.ProductionStationInfo, len(productionLine.ProductionStations))
		for psi, ps := range productionLine.ProductionStations {
			productionStation := &proto.ProductionStationInfo{
				Id:              ps.Id,
				Code:            ps.Code,
				Description:     ps.Description,
				StationType:     ps.StationType,
				AccountControl:  ps.AccountControl,
				MaterialControl: ps.MaterialControl,
			}

			// productionProcesses := []map[string]interface{}{}
			// for _, _productionProcesse := range _productionStation.ProductionProcesses {
			// 	if _productionProcesse.Enable {
			// 		productionProcesse := map[string]interface{}{}
			// 		productionProcesse["id"] = _productionProcesse.Id
			// 		productionProcesse["code"] = _productionProcesse.Code
			// 		productionProcesse["description"] = _productionProcesse.Description
			// 		productionProcesse["processType"] = _productionProcesse.ProcessType
			// 		productionProcesse["vehicleType"] = _productionProcesse.VehicleType
			// 		productionProcesse["enableControl"] = _productionProcesse.EnableControl
			// 		productionProcesses = append(productionProcesses, productionProcesse)
			// 	}
			// }
			// productionStation["productionProcesses"] = productionProcesses
			productionStations[psi] = productionStation
		}
		productionLine.ProductionStations = productionStations

		data[pli] = productionLine
	}

	return data, nil
}

// RetrieveProductionStation 查询产线工位信息
func RetrieveProductionStation(req *proto.RetrieveProductionStationRequest) ([]*proto.ProductionStationInfo, error) {
	if req.ProductionLine == "" {
		return nil, fmt.Errorf("ProductionLine不能为空")
	}

	r := &proto.GetDetailRequest{Id: req.ProductionLine}
	if req.StationType != "" {
		r.Id = req.StationType
	}
	_productionLine, err := clients.ProductionLineClient.GetDetail(context.Background(), r)
	if err != nil {
		return nil, err
	}

	productionLine := _productionLine.Data

	data := make([]*proto.ProductionStationInfo, len(productionLine.ProductionStations))
	for i, s := range productionLine.ProductionStations {
		data[i] = &proto.ProductionStationInfo{
			Id:              s.Id,
			Code:            s.Code,
			Description:     s.Description,
			StationType:     s.StationType,
			AccountControl:  s.AccountControl,
			MaterialControl: s.MaterialControl,
			CurrentState:    s.CurrentState,
		}
	}

	return data, nil
}

// RetrieveProductAttribute 查询产品特性信息
func RetrieveProductAttribute(req *proto.RetrieveProductAttributeRequest) ([]*proto.ProductAttributeInfo, error) {
	if req.Code == "" {
		return nil, fmt.Errorf("Code不能为空")
	}

	_productAttributes, err := clients.ProductAttributeClient.Query(context.Background(), &proto.QueryProductAttributeRequest{
		PageSize:    1000,
		Code:        req.Code,
		Description: req.Description,
	})
	if err != nil {
		return nil, err
	}
	productAttributes := _productAttributes.Data

	data := make([]*proto.ProductAttributeInfo, len(productAttributes))
	for pai, pa := range productAttributes {
		data[pai] = &proto.ProductAttributeInfo{
			Code:        pa.Code,
			Description: pa.Description,
		}
	}

	return data, nil
}

// RetrieveProductionCrossway 查询产线路口信息
func RetrieveProductionCrossway(req *proto.RetrieveProductionCrosswayRequest) ([]*proto.ProductionCrosswayInfo, error) {
	if req.ProductionLine == "" {
		return nil, fmt.Errorf("ProductionLine不能为空")
	}

	_productionCrossways, err := clients.ProductionCrosswayClient.Query(context.Background(), &proto.QueryProductionCrosswayRequest{ProductionLineID: req.GetProductionLine()})
	if err != nil {
		return nil, err
	}
	productionCrossways := _productionCrossways.Data

	data := make([]*proto.ProductionCrosswayInfo, len(productionCrossways))
	for i, v := range productionCrossways {
		data[i] = &proto.ProductionCrosswayInfo{
			Id:          v.Id,
			Code:        v.Code,
			Description: v.Description,
			DefaultTurn: v.DefaultTurn,
			Remark:      v.Remark,
		}
	}

	return data, nil
}
