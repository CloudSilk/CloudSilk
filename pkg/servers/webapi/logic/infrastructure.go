package logic

import (
	"context"
	"fmt"

	"github.com/CloudSilk/CloudSilk/pkg/clients"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	modelcode "github.com/CloudSilk/pkg/model"
	"gorm.io/gorm"
)

// GetAllProductionLine 获取全部产线信息
func GetAllProductionLine() ([]*proto.ProductionLineInfo, error) {
	_productionLine, _ := clients.ProductionLineClient.GetAll(context.Background(), &proto.GetAllRequest{})
	if _productionLine.Code != modelcode.Success {
		return nil, fmt.Errorf(_productionLine.Message)
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

		productionStations := make([]*proto.ProductionStationInfo, len(pl.ProductionStations))
		for psi, ps := range pl.ProductionStations {
			productionStation := &proto.ProductionStationInfo{
				Id:              ps.Id,
				Code:            ps.Code,
				Description:     ps.Description,
				StationType:     ps.StationType,
				AccountControl:  ps.AccountControl,
				MaterialControl: ps.MaterialControl,
			}

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

	_productionLine, _ := clients.ProductionLineClient.GetDetail(context.Background(), &proto.GetDetailRequest{Id: req.ProductionLine})
	if _productionLine.Message == gorm.ErrRecordNotFound.Error() {
		return nil, fmt.Errorf("ProductionLine不存在")
	}
	if _productionLine.Code != modelcode.Success {
		return nil, fmt.Errorf(_productionLine.Message)
	}
	productionLine := _productionLine.Data

	data := []*proto.ProductionStationInfo{}
	for _, s := range productionLine.ProductionStations {
		if req.StationType != "" {
			if s.StationType != req.StationType {
				continue
			}
		}
		data = append(data, &proto.ProductionStationInfo{
			Id:              s.Id,
			Code:            s.Code,
			Description:     s.Description,
			StationType:     s.StationType,
			AccountControl:  s.AccountControl,
			MaterialControl: s.MaterialControl,
			CurrentState:    s.CurrentState,
		})
	}

	return data, nil
}

// RetrieveProductAttribute 查询产品特性信息
func RetrieveProductAttribute(req *proto.RetrieveProductAttributeRequest) ([]*proto.ProductAttributeInfo, error) {
	if req.Code == "" {
		return nil, fmt.Errorf("Code不能为空")
	}

	_productAttributes, _ := clients.ProductAttributeClient.Query(context.Background(), &proto.QueryProductAttributeRequest{
		PageSize:    1000,
		Code:        req.Code,
		Description: req.Description,
	})
	if _productAttributes.Code != modelcode.Success {
		return nil, fmt.Errorf(_productAttributes.Message)
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

	_productionCrossways, _ := clients.ProductionCrosswayClient.Query(context.Background(), &proto.QueryProductionCrosswayRequest{ProductionLineID: req.ProductionLine})
	if _productionCrossways.Code != modelcode.Success {
		return nil, fmt.Errorf(_productionCrossways.Message)
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
