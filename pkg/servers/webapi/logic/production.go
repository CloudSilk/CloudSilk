package logic

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/clients"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/tool"
	"github.com/CloudSilk/CloudSilk/pkg/types"
	modelcode "github.com/CloudSilk/pkg/model"
	"github.com/CloudSilk/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// 设定产品信息状态为上线装配
func OnlineProductInfo(req *proto.OnlineProductInfoRequest) *proto.CommonResponse {
	if req.ProductionLine == "" {
		return &proto.CommonResponse{Code: 40000, Message: "ProductionLine不能为空"}
	}
	if req.ProductSerialNo == "" {
		return &proto.CommonResponse{Code: 40000, Message: "ProductSerialNo不能为空"}
	}

	_productionLine, _ := clients.ProductionLineClient.Get(context.Background(), &proto.GetProductionLineRequest{Code: req.ProductionLine})
	if _productionLine.Message == gorm.ErrRecordNotFound.Error() {
		return &proto.CommonResponse{Code: 10001, Message: "无效的产线代号"}
	}
	if _productionLine.Code != modelcode.Success {
		return &proto.CommonResponse{Code: 50000, Message: _productionLine.Message}
	}
	productionLine := _productionLine.Data

	_productInfo, _ := clients.ProductInfoClient.Get(context.Background(), &proto.GetProductInfoRequest{ProductSerialNo: req.ProductSerialNo})
	if _productInfo.Message == gorm.ErrRecordNotFound.Error() {
		return &proto.CommonResponse{Code: 10002, Message: "读取产品信息失败，请联系管理员处理"}
	}
	if _productInfo.Code != modelcode.Success {
		return &proto.CommonResponse{Code: 50000, Message: _productInfo.Message}
	}
	productInfo := _productInfo.Data

	_productOrder, _ := clients.ProductOrderClient.GetDetail(context.Background(), &proto.GetDetailRequest{Id: productInfo.ProductOrderID})
	if _productOrder.Message == gorm.ErrRecordNotFound.Error() {
		return &proto.CommonResponse{Code: 10003, Message: "此生产工单发放产线与上线产线不匹配"}
	}
	if _productOrder.Code != modelcode.Success {
		return &proto.CommonResponse{Code: 50000, Message: _productOrder.Message}
	}
	productOrder := _productOrder.Data

	if req.ProductOrderNo != "" && productOrder.ProductOrderNo != req.ProductOrderNo {
		return &proto.CommonResponse{Code: 10003, Message: "此产品的隶属工单与当前工单不匹配"}
	}

	//TODO: 兼容，部分产线是直接创建产品工艺路线，部分是根据工单工艺动态创建
	_productProcessRoutes, _ := clients.ProductProcessRouteClient.Query(context.Background(), &proto.QueryProductProcessRouteRequest{
		PageSize:      1,
		SortConfig:    `{"route_index": "asc"}`,
		ProductInfoID: productInfo.Id,
		CurrentState:  types.ProductProcessRouteStateWaitProcess,
	})
	if _productProcessRoutes.Code != modelcode.Success {
		return &proto.CommonResponse{Code: 50000, Message: _productProcessRoutes.Message}
	}
	productProcessRoutes := _productProcessRoutes.Data

	var productProcessRoute *proto.ProductProcessRouteInfo
	if len(productProcessRoutes) == 0 {
		_productOrderProcesses, _ := clients.ProductOrderProcessClient.Query(context.Background(), &proto.QueryProductOrderProcessRequest{
			PageSize:       1,
			SortConfig:     `{"sort_index": "asc"}`,
			ProductOrderID: productInfo.ProductOrderID,
			Enable:         true,
		})
		if _productOrderProcesses.Code != modelcode.Success {
			return &proto.CommonResponse{Code: 50000, Message: _productOrderProcesses.Message}
		}
		if len(_productOrderProcesses.Data) == 0 {
			return &proto.CommonResponse{Code: 10004, Message: "上线失败，此工单缺少工艺路线"}
		}
		productOrderProcess := _productOrderProcesses.Data[0]

		productProcessRoute = &proto.ProductProcessRouteInfo{
			CurrentProcessID: productOrderProcess.ProductionProcessID,
			CurrentState:     types.ProductProcessRouteStateWaitProcess,
			RouteIndex:       productOrderProcess.SortIndex,
			ProductInfoID:    productInfo.Id,
		}

		if _productProcessRoute, _ := clients.ProductProcessRouteClient.Add(context.Background(), productProcessRoute); _productProcessRoute.Code != modelcode.Success {
			return &proto.CommonResponse{Code: 50000, Message: _productProcessRoute.Message}
		}
	} else {
		productProcessRoute = productProcessRoutes[0]
	}

	productProcessRoute.WorkIndex = 1

	//TODO: 更新产品信息
	_productOrderProcesses, _ := clients.ProductOrderProcessClient.Query(context.Background(), &proto.QueryProductOrderProcessRequest{
		ProductOrderID: productInfo.ProductOrderID,
		Enable:         true,
		SortIndex:      productProcessRoute.WorkIndex,
	})
	now := time.Now()
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	remainingRoutes := int32(len(_productOrderProcesses.Data))
	estimateTime := now.Add(time.Duration(remainingRoutes*productInfo.ProductOrder.StandardWorkTime) * time.Second).Format("2006-01-02 15:04:05")
	productInfo.ProductionProcessID = productProcessRoute.CurrentProcessID
	productInfo.RemainingRoutes = remainingRoutes
	productInfo.EstimateTime = estimateTime
	if productProcessRoute.CurrentProcess != nil && productProcessRoute.CurrentProcess.ProductState != "" {
		productInfo.CurrentState = productProcessRoute.CurrentProcess.ProductState
	} else {
		productInfo.CurrentState = types.ProductStateAssembling
	}

	productInfo.StartedTime = nowStr
	if _productInfo, _ := clients.ProductInfoClient.Update(context.Background(), productInfo); _productInfo.Code != modelcode.Success {
		return &proto.CommonResponse{Code: 50000, Message: _productInfo.Message}
	}

	//TODO: 更新工单信息
	if productOrder.ActualStartTime == "" {
		productOrder.ActualStartTime = nowStr
		productOrder.StartedQTY = 0
		productOrder.CurrentState = types.ProductOrderStateProducting
		_systemEvent, _ := clients.SystemEventClient.Get(context.Background(), &proto.GetSystemEventRequest{Code: types.SystemEventProductOrderStarted})
		if _systemEvent.Data != nil {
			systemEvent := _systemEvent.Data
			_systemEventTrigger, _ := clients.SystemEventTriggerClient.Add(context.Background(), &proto.SystemEventTriggerInfo{
				SystemEventID: systemEvent.Id,
				EventNo:       uuid.NewString(),
				CurrentState:  types.SystemEventTriggerStateWaitExecute,
			})
			if _systemEventTrigger.Code != modelcode.Success {
				return &proto.CommonResponse{Code: 50000, Message: _systemEventTrigger.Message}
			}
			systemEventTriggerID := _systemEventTrigger.Message

			for _, systemEventParameter := range systemEvent.SystemEventParameters {
				value := strings.ReplaceAll(systemEventParameter.Value, "{ProductOrderNo}", productOrder.ProductOrderNo)
				if _systemEventTriggerParameter, _ := clients.SystemEventTriggerParameterClient.Add(context.Background(), &proto.SystemEventTriggerParameterInfo{
					SystemEventTriggerID: systemEventTriggerID,
					DataType:             systemEventParameter.DataType,
					Name:                 systemEventParameter.Name,
					Description:          systemEventParameter.Description,
					Value:                value,
				}); _systemEventTriggerParameter.Code != modelcode.Success {
					return &proto.CommonResponse{Code: 50000, Message: _systemEventTriggerParameter.Message}
				}
			}
		}
		productOrder.StartedQTY += 1

		if _productOrder, _ := clients.ProductOrderClient.Update(context.Background(), productOrder); _productOrder.Code != modelcode.Success {
			return &proto.CommonResponse{Code: 50000, Message: _productOrder.Message}
		}

		//TODO: 绑定托盘
		if req.TrayNo != "" {
			_materialTray, _ := clients.MaterialTrayClient.Get(context.Background(), &proto.GetMaterialTrayRequest{Identifier: req.TrayNo})
			if _materialTray.Message == gorm.ErrRecordNotFound.Error() {
				return &proto.CommonResponse{Code: 10005, Message: "无效的物料托盘识别码"}
			}
			if _materialTray.Code != modelcode.Success {
				return &proto.CommonResponse{Code: 50000, Message: _materialTray.Message}
			}
			materialTray := _materialTray.Data

			if !materialTray.Enable {
				return &proto.CommonResponse{Code: 10005, Message: "托盘已禁用"}
			}
			if materialTray.TrayType != types.MaterialTrayTypeMaterialTray {
				return &proto.CommonResponse{Code: 10005, Message: "非法操作，只允许使用物料托盘上线"}
			}
			if materialTray.ProductionLineID != productionLine.Id {
				return &proto.CommonResponse{Code: 10005, Message: "非法操作，此托盘不属于当前产线"}
			}
			if materialTray.ProductInfoID != "" && materialTray.ProductInfoID != productInfo.Id {
				return &proto.CommonResponse{Code: 10005, Message: "非法操作，此托盘已绑定其他产品"}
			}

			if _materialTrayBindingRecord, _ := clients.MaterialTrayBindingRecordClient.Add(context.Background(), &proto.MaterialTrayBindingRecordInfo{
				MaterialTrayID: materialTray.Id,
				ProductInfoID:  productInfo.Id,
				CreateTime:     nowStr,
				CurrentState:   types.MaterialTrayBindingRecordStateEffected,
			}); _materialTrayBindingRecord.Code != modelcode.Success {
				return &proto.CommonResponse{Code: 50000, Message: _materialTrayBindingRecord.Message}
			}
		}

		//TODO: 触发事件
		_systemEvent, _ = clients.SystemEventClient.Get(context.Background(), &proto.GetSystemEventRequest{Code: types.SystemEventProductInfoOnlined})
		if _systemEvent.Code != modelcode.Success && _systemEvent.Message != gorm.ErrRecordNotFound.Error() {
			return &proto.CommonResponse{Code: 50000, Message: _systemEvent.Message}
		}
		if _systemEvent.Data != nil {
			systemEvent := _systemEvent.Data
			_systemEventTrigger, _ := clients.SystemEventTriggerClient.Add(context.Background(), &proto.SystemEventTriggerInfo{
				SystemEventID: systemEvent.Id,
				EventNo:       uuid.NewString(),
				CurrentState:  types.SystemEventTriggerStateWaitExecute,
			})
			if _systemEventTrigger.Code != modelcode.Success {
				return &proto.CommonResponse{Code: 50000, Message: _systemEventTrigger.Message}
			}
			systemEventTriggerID := _systemEventTrigger.Message

			_productOrderAttributes, _ := clients.ProductOrderAttributeClient.Query(context.Background(), &proto.QueryProductOrderAttributeRequest{ProductOrderID: productInfo.ProductOrderID})
			if _productOrderAttributes.Code != modelcode.Success {
				return &proto.CommonResponse{Code: 50000, Message: _systemEventTrigger.Message}
			}
			for _, systemEventParameter := range systemEvent.SystemEventParameters {
				value := systemEventParameter.Value
				value = strings.ReplaceAll(value, "{ProductSerialNo}", productInfo.ProductSerialNo)
				value = strings.ReplaceAll(value, "{ProductOrderNo}", productOrder.ProductOrderNo)
				value = strings.ReplaceAll(value, "{SalesOrderNo}", productOrder.SalesOrderNo)
				value = strings.ReplaceAll(value, "{ItemNo}", productOrder.ItemNo)
				for _, productOrderAttribute := range _productOrderAttributes.Data {
					value = strings.ReplaceAll(value, "{"+productOrderAttribute.ProductAttribute.Code+"}", productOrderAttribute.Value)
				}

				if _systemEventTriggerParameter, _ := clients.SystemEventTriggerParameterClient.Add(context.Background(), &proto.SystemEventTriggerParameterInfo{
					SystemEventTriggerID: systemEventTriggerID,
					DataType:             systemEventParameter.DataType,
					Name:                 systemEventParameter.Name,
					Description:          systemEventParameter.Description,
					Value:                value,
				}); _systemEventTriggerParameter.Code != modelcode.Success {
					return &proto.CommonResponse{Code: 50000, Message: _systemEventTriggerParameter.Message}
				}
			}
		}
	}

	return &proto.CommonResponse{Code: 20000, Message: "上线成功"}
}

// Code = 0, 工艺路线正确
// Code = 1, 校验失败
// Code = 2, 返工产品
// Code = 3，工艺路线错误
// Code = 4, 完工产品
// Code = 5, 读取托盘信息失败
// 请求入站
func EnterProductionStation(req *proto.EnterProductionStationRequest) *proto.EnterProductionStationResponse {
	if req.ProductionStation == "" {
		return &proto.EnterProductionStationResponse{Code: 5, Message: "ProductionStation不能为空"}
	}

	nowStr := time.Now().Format("2006-01-02 15:04:05")
	if req.TrayNo != "" {
		//根据托盘号获取物料托盘
		_materialTray, _ := clients.MaterialTrayClient.Get(context.Background(), &proto.GetMaterialTrayRequest{Identifier: req.TrayNo})
		if _materialTray.Message == gorm.ErrRecordNotFound.Error() {
			return &proto.EnterProductionStationResponse{Code: 5, Message: "无效的托盘识别码"}
		}
		if _materialTray.Code != modelcode.Success {
			return &proto.EnterProductionStationResponse{Code: 1, Message: _materialTray.Message}
		}

		materialTray := _materialTray.Data
		if materialTray.ProductionLineID == "" {
			return &proto.EnterProductionStationResponse{Code: 5, Message: "托盘未绑定任何产品"}
		}

		req.ProductSerialNo = materialTray.ProductInfo.ProductSerialNo
	}
	if req.PackageNo != "" {
		//根据包装箱号获取产品包装记录
		_productPackageRecord, _ := clients.ProductPackageRecordClient.Get(context.Background(), &proto.GetProductPackageRecordRequest{PackageNo: req.PackageNo})
		if _productPackageRecord.Message == gorm.ErrRecordNotFound.Error() {
			return &proto.EnterProductionStationResponse{Code: 5, Message: "无效的包装箱号"}
		}
		if _productPackageRecord.Code != modelcode.Success {
			return &proto.EnterProductionStationResponse{Code: 1, Message: _productPackageRecord.Message}
		}

		productPackageRecord := _productPackageRecord.Data
		if productPackageRecord.ProductInfoID == "" {
			return &proto.EnterProductionStationResponse{Code: 5, Message: "包装箱未绑定任何产品"}
		}

		req.ProductSerialNo = productPackageRecord.ProductInfo.ProductSerialNo
	}
	req.ProductSerialNo = strings.Trim(strings.Trim(req.ProductSerialNo, "\000"), "\r")
	if req.ProductSerialNo == "" {
		return &proto.EnterProductionStationResponse{Code: 5, Message: "ProductSerialNo不能为空"}
	}

	//根据工位代号获取产线工站
	_productionStation, _ := clients.ProductionStationClient.Get(context.Background(), &proto.GetProductionStationRequest{Code: req.ProductionStation})
	if _productionStation.Message == gorm.ErrRecordNotFound.Error() {
		return &proto.EnterProductionStationResponse{Code: 1, Message: "无效的工位代号"}
	}
	if _productionStation.Code != modelcode.Success {
		return &proto.EnterProductionStationResponse{Code: 1, Message: _productionStation.Message}
	}

	productionStation := _productionStation.Data
	if productionStation.AccountControl && productionStation.CurrentUserID == "" {
		return &proto.EnterProductionStationResponse{Code: 1, Message: "工位未登录，无法进站"}
	}
	if productionStation.CurrentState == types.ProductionStationStateBreakdown {
		return &proto.EnterProductionStationResponse{Code: 1, Message: "设备故障中，请尽快联系人员处理并恢复设备故障"}
	}
	//根据产品序列号获取产品
	_productInfo, _ := clients.ProductInfoClient.Get(context.Background(), &proto.GetProductInfoRequest{ProductSerialNo: req.ProductSerialNo})
	if _productInfo.Message == gorm.ErrRecordNotFound.Error() {
		return &proto.EnterProductionStationResponse{Code: 1, Message: "读取产品信息失败"}
	}
	if _productInfo.Code != modelcode.Success {
		return &proto.EnterProductionStationResponse{Code: 1, Message: _productInfo.Message}
	}

	productInfo := _productInfo.Data
	switch productInfo.CurrentState {
	case types.ProductStateChecking:
		return &proto.EnterProductionStationResponse{Code: 2, Message: "产品状态错误，此产品状态为检查中"}
	case types.ProductStateReworking:
		return &proto.EnterProductionStationResponse{Code: 2, Message: "产品状态错误，此产品状态为返工中"}
	case types.ProductStateCompleted:
		return &proto.EnterProductionStationResponse{Code: 4, Message: "产品状态错误，此产品状态为已完工"}
	}
	if productInfo.ProductionProcessID == "" {
		return &proto.EnterProductionStationResponse{Code: 1, Message: "此产品未开工"}
	}

	//根据id获取产品订单
	_productOrder, _ := clients.ProductOrderClient.GetDetail(context.Background(), &proto.GetDetailRequest{Id: productInfo.ProductOrderID})
	if _productOrder.Message == gorm.ErrRecordNotFound.Error() {
		return &proto.EnterProductionStationResponse{Code: 1, Message: "读取产品工单失败"}
	}
	if _productOrder.Code != modelcode.Success {
		return &proto.EnterProductionStationResponse{Code: 1, Message: _productOrder.Message}
	}
	productOrder := _productOrder.Data

	//获取产品节拍
	if _productRhythmRecord, _ := clients.ProductRhythmRecordClient.Get(context.Background(), &proto.GetProductRhythmRecordRequest{
		ProductionProcessID: productInfo.ProductionProcessID,
		ProductInfoID:       productInfo.Id,
		ProductionStationID: productionStation.Id,
		HasWorkEndTime:      false,
	}); _productRhythmRecord.Message == gorm.ErrRecordNotFound.Error() {
		//TODO: 重复进站不重复报工，以第一次进站时间为准
		if _productRhythmRecord, _ := clients.ProductRhythmRecordClient.Add(context.Background(), &proto.ProductRhythmRecordInfo{
			WorkUserID:          productionStation.CurrentUserID,
			ProductionStationID: productionStation.Id,
			ProductInfoID:       productInfo.Id,
			ProductionProcessID: productInfo.ProductionProcessID,
			StandardWorkTime:    productInfo.ProductOrder.StandardWorkTime,
			WorkStartTime:       nowStr,
		}); _productRhythmRecord.Code != modelcode.Success {
			return &proto.EnterProductionStationResponse{Code: 1, Message: _productRhythmRecord.Message}
		}
	} else if _productRhythmRecord.Code != modelcode.Success {
		return &proto.EnterProductionStationResponse{Code: 1, Message: _productRhythmRecord.Message}
	}

	targetStates := []string{types.ProductProcessRouteStateWaitProcess, types.ProductProcessRouteStateProcessing}
	//获取产品工艺路线
	_productProcessRoute, _ := clients.ProductProcessRouteClient.Get(context.Background(), &proto.GetProductProcessRouteRequest{
		ProductInfoID:    productInfo.Id,
		CurrentProcessID: productInfo.ProductionProcessID,
		CurrentStates:    targetStates,
	})
	if _productProcessRoute.Message == gorm.ErrRecordNotFound.Error() {
		return &proto.EnterProductionStationResponse{Code: 3, Message: "读取产品当前工艺路线错误"}
	}
	if _productProcessRoute.Code != modelcode.Success {
		return &proto.EnterProductionStationResponse{Code: 1, Message: _productProcessRoute.Message}
	}

	//修改工艺路线状态和执行工位
	productProcessRoute := _productProcessRoute.Data
	productProcessRoute.ProcessUserID = productionStation.CurrentUserID
	productProcessRoute.ProductionStationID = productionStation.Id
	productProcessRoute.CurrentState = types.ProductProcessRouteStateProcessing
	if _productProcessRoute, _ := clients.ProductProcessRouteClient.Update(context.Background(), productProcessRoute); _productProcessRoute.Code != modelcode.Success {
		return &proto.EnterProductionStationResponse{Code: 1, Message: _productProcessRoute.Message}
	}

	//获取生产工艺
	_productionProcess, _ := clients.ProductionProcessClient.GetDetail(context.Background(), &proto.GetDetailRequest{Id: productInfo.ProductionProcessID})
	if _productionProcess.Message == gorm.ErrRecordNotFound.Error() {
		return &proto.EnterProductionStationResponse{Code: 3, Message: "读取产品当前工艺错误"}
	}
	if _productionProcess.Code != modelcode.Success {
		return &proto.EnterProductionStationResponse{Code: 1, Message: _productionProcess.Message}
	}

	productionProcess := _productionProcess.Data
	var accept bool
	for _, _productionStation := range productionProcess.ProductionProcessAvailableStations {
		if _productionStation.ProductionStationID == productionStation.Id {
			accept = true
			break
		}
	}
	//工序朝向，大于0表示当前工序在当前工位之后，可以放行；小于0表示当前工序在当前工位之前，禁止放行。
	var sopLink string
	if !accept {
		_stationRoute, _ := clients.ProductionProcessClient.Query(context.Background(), &proto.QueryProductionProcessRequest{
			PageSize:            1,
			ProductionLineID:    productOrder.ProductionLineID,
			SortConfig:          `{"sort_index": "asc"}`,
			ProductionStationID: productionStation.Id,
		})
		if _stationRoute.Message == gorm.ErrRecordNotFound.Error() {
			return &proto.EnterProductionStationResponse{Code: 3, Message: "工艺路线错误，且当前工位并不在当前产线的工艺路线之内，请联系管理员处理"}
		}
		if _stationRoute.Code != modelcode.Success {
			return &proto.EnterProductionStationResponse{Code: 1, Message: _stationRoute.Message}
		}

		stationRoute := _stationRoute.Data[0]
		toward := productionProcess.SortIndex - stationRoute.SortIndex
		towardStr := "前"
		if toward > 0 {
			towardStr = "后"
		}
		return &proto.EnterProductionStationResponse{Code: 3, Data: &proto.EnterProductionStationData{
			Toward: toward,
			ProductionProcess: &proto.EnterProductionStationInfo{
				Code:        productionProcess.Code,
				Description: productionProcess.Description,
				Identifier:  productionProcess.Identifier,
			},
		}, Message: fmt.Sprintf("工艺路线错误，且此产品的当前工序为%s(%s)，在当前工位的执行工序之%s", productionProcess.Description, productionProcess.Code, towardStr)}
	}

	//检查人员资质
	//检查工序是否启用人员管控
	if productionProcess.EnableControl {
		//获取产品型号
		_productModel, _ := clients.ProductModelClient.GetDetail(context.Background(), &proto.GetDetailRequest{Id: productOrder.ProductModelID})
		if _productModel.Message == gorm.ErrRecordNotFound.Error() {
			return &proto.EnterProductionStationResponse{Code: 1, Message: "读取产品型号失败"}
		}
		if _productModel.Code != modelcode.Success {
			return &proto.EnterProductionStationResponse{Code: 1, Message: _productModel.Message}
		}

		//获取人员资格
		_personnelQualifications, _ := clients.PersonnelQualificationClient.Query(context.Background(), &proto.QueryPersonnelQualificationRequest{
			ProductionProcessID: productionProcess.Id,
			ProductModelID:      _productModel.Data.Id,
		})
		if _personnelQualifications.Code != modelcode.Success {
			return &proto.EnterProductionStationResponse{Code: 1, Message: _personnelQualifications.Message}
		}

		if len(_personnelQualifications.Data) > 0 {
			_personnelQualification, _ := clients.PersonnelQualificationClient.Get(context.Background(), &proto.GetPersonnelQualificationRequest{CertifiedUserID: productionStation.CurrentUserID})
			if _personnelQualification.Message == gorm.ErrRecordNotFound.Error() {
				return &proto.EnterProductionStationResponse{Code: 1, Message: "当前作业人员缺少认证资质，无法开工"}
			}
			if _personnelQualification.Code != modelcode.Success {
				return &proto.EnterProductionStationResponse{Code: 1, Message: _personnelQualification.Message}
			}
			if _personnelQualification.Data.ExpirationDate <= nowStr {
				return &proto.EnterProductionStationResponse{Code: 1, Message: "当前作业人员的认证资质已过期，无法开工"}
			}
		}
	}

	//获取作业手册
	_productionProcessSop, _ := clients.ProductionProcessSopClient.Get(context.Background(), &proto.GetProductionProcessSopRequest{
		ProductionProcessID: productionProcess.Id,
		ProductModelID:      productOrder.ProductModelID,
	})
	if _productionProcessSop.Code != modelcode.Success && _productionProcessSop.Message != gorm.ErrRecordNotFound.Error() {
		return &proto.EnterProductionStationResponse{Code: 1, Message: _productionProcessSop.Message}
	}

	if _productionProcessSop.Data != nil {
		if _productionProcessSop.Data.FileLink != "" {
			sopLink = _productionProcessSop.Data.FileLink
		}
	}

	if productionProcess.EnableReport {
		//创建系统事件上报开工
		_systemEvent, _ := clients.SystemEventClient.Get(context.Background(), &proto.GetSystemEventRequest{Code: types.SystemEventProductionProcessStarted, Enable: true})
		if _systemEvent.Code != modelcode.Success && _systemEvent.Message != gorm.ErrRecordNotFound.Error() {
			return &proto.EnterProductionStationResponse{Code: 1, Message: _systemEvent.Message}
		}
		if _systemEvent.Data != nil {
			systemEvent := _systemEvent.Data
			systemEventTrigger := &proto.SystemEventTriggerInfo{
				SystemEventID: systemEvent.Id,
				CreateTime:    nowStr,
				EventNo:       uuid.NewString(),
				CurrentState:  types.SystemEventTriggerStateWaitExecute,
			}
			for _, _systemEventParameter := range systemEvent.SystemEventParameters {
				value := _systemEventParameter.Value
				value = strings.ReplaceAll(value, "{ProductionProcess.Identifier}", productionProcess.Identifier)
				value = strings.ReplaceAll(value, "{ProductionLine.Identifier}", productionStation.ProductionLine.Identifier)
				value = strings.ReplaceAll(value, "{ProductionStation}", productionStation.Code)
				value = strings.ReplaceAll(value, "{ProductSerialNo}", productInfo.ProductSerialNo)
				value = strings.ReplaceAll(value, "{ProductionProcess}", productionProcess.Code)

				systemEventTrigger.SystemEventTriggerParameters = append(systemEventTrigger.SystemEventTriggerParameters, &proto.SystemEventTriggerParameterInfo{
					DataType:    _systemEventParameter.DataType,
					Name:        _systemEventParameter.Name,
					Description: _systemEventParameter.Description,
					Value:       value,
				})
			}

			if _systemEventTrigger, _ := clients.SystemEventTriggerClient.Add(context.Background(), systemEventTrigger); _systemEventTrigger.Code != modelcode.Success {
				return &proto.EnterProductionStationResponse{Code: 1, Message: _systemEventTrigger.Message}
			}
		}
	}

	if productionStation.AllowReport {
		productionStation.CurrentState = types.ProductionStationStateOccupied
		//创建系统事件上报工位开始处于作业状态
		_systemEvent, _ := clients.SystemEventClient.Get(context.Background(), &proto.GetSystemEventRequest{Code: types.SystemEventProductionStationOccupied, Enable: true})
		if _systemEvent.Code != modelcode.Success && _systemEvent.Message != gorm.ErrRecordNotFound.Error() {
			return &proto.EnterProductionStationResponse{Code: 1, Message: _systemEvent.Message}
		}
		if _systemEvent.Data != nil {
			systemEvent := _systemEvent.Data
			systemEventTrigger := &proto.SystemEventTriggerInfo{
				SystemEventID: systemEvent.Id,
				CreateTime:    nowStr,
				EventNo:       uuid.NewString(),
				CurrentState:  types.SystemEventTriggerStateWaitExecute,
			}
			for _, _systemEventParameter := range systemEvent.SystemEventParameters {
				value := _systemEventParameter.Value

				value = strings.ReplaceAll(value, "{ProductionStation}", productionStation.Code)
				value = strings.ReplaceAll(value, "{ProductSerialNo}", productInfo.ProductSerialNo)
				value = strings.ReplaceAll(value, "{ProductionProcess}", productionProcess.Code)

				systemEventTrigger.SystemEventTriggerParameters = append(systemEventTrigger.SystemEventTriggerParameters, &proto.SystemEventTriggerParameterInfo{
					DataType:    _systemEventParameter.DataType,
					Name:        _systemEventParameter.Name,
					Description: _systemEventParameter.Description,
					Value:       value,
				})
			}

			if _systemEventTrigger, _ := clients.SystemEventTriggerClient.Add(context.Background(), systemEventTrigger); _systemEventTrigger.Code != modelcode.Success {
				return &proto.EnterProductionStationResponse{Code: 1, Message: _systemEventTrigger.Message}
			}
		}
		if _productionStation, _ := clients.ProductionStationClient.Update(context.Background(), productionStation); _productionStation.Code != modelcode.Success {
			return &proto.EnterProductionStationResponse{Code: 1, Message: _productionStation.Message}
		}
	}

	return &proto.EnterProductionStationResponse{
		Code: 0,
		Data: &proto.EnterProductionStationData{
			ProductOrderNo:      productInfo.ProductOrder.ProductOrderNo,
			ProductSerialNo:     productInfo.ProductSerialNo,
			SalesOrderNo:        productInfo.ProductOrder.SalesOrderNo,
			ItemNo:              productInfo.ProductOrder.ItemNo,
			OrderTime:           productInfo.ProductOrder.OrderTime,
			ProductCategory:     productInfo.ProductOrder.ProductModel.ProductCategory.Code,
			ProductModel:        productInfo.ProductOrder.ProductModel.Code,
			MaterialNo:          productInfo.ProductOrder.ProductModel.MaterialNo,
			MaterialDescription: productInfo.ProductOrder.ProductModel.MaterialDescription,
			CurrentState:        productInfo.CurrentState,
			PropertyBrief:       productInfo.ProductOrder.PropertyBrief,
			Remark:              productInfo.ProductOrder.Remark,
			ProductionProcess: &proto.EnterProductionStationInfo{
				Code:        productionProcess.Code,
				Description: productionProcess.Description,
				SopLink:     sopLink,
			},
		},
	}
}

// 获取工站的当前生产工单，工序展示信息
// func GetProductionStationExhibition(productionStationCode string) (*proto.GetProductionStationExhibitionData, error) {
// 	if productionStationCode == "" {
// 		return nil, fmt.Errorf("ProductionStation不能为空")
// 	}
// 	_productionStation, _ := clients.ProductionStationClient.Get(context.Background(), &proto.GetProductionStationRequest{Code: productionStationCode})
// 	if _productionStation.Message == gorm.ErrRecordNotFound.Error() {
// 		return nil, fmt.Errorf("无效的工站代号")
// 	}
// 	if _productionStation.Code != modelcode.Success {
// 		return nil, fmt.Errorf(_productionStation.Message)
// 	}
// 	productionStation := _productionStation.Data
// 	if productionStation.ProductionLineID == "" {
// 		return nil, fmt.Errorf("读取工站所属产线失败")
// 	}
// 	if productionStation.ProductionLine.ProductModelID == "" {
// 		return nil, fmt.Errorf("读取失败，工站所属的产线未设置当前产品型号")
// 	}
// 	_productionProcesses, _ := clients.ProductionProcessClient.Query(context.Background(), &proto.QueryProductionProcessRequest{
// 		PageSize:            1,
// 		ProductionLineID:    productionStation.ProductionLineID,
// 		ProductionStationID: productionStation.Id,
// 	})
// 	if _productionProcesses.Code != modelcode.Success {
// 		return nil, fmt.Errorf(_productionProcesses.Message)
// 	}
// 	if len(_productionProcesses.Data) == 0 {
// 		return nil, fmt.Errorf("读取工位的归属工序失败")
// 	}
// 	productionProcess := _productionProcesses.Data[0]
// 	_productProcessRoutes, _ := clients.ProductProcessRouteClient.Query(context.Background(), &proto.QueryProductProcessRouteRequest{
// 		PageSize:         1000,
// 		SortConfig:       `{"create_time": "descend"}`,
// 		CurrentProcessID: productionProcess.Id,
// 	})
// 	if _productProcessRoutes.Code != modelcode.Success {
// 		return nil, fmt.Errorf(_productProcessRoutes.Message)
// 	}
// 	var productProcessRoute *proto.ProductProcessRouteInfo
// 	for _, v := range _productProcessRoutes.Data {
// 		if v.ProductInfo.ProductOrder.ProductModelID == productionStation.ProductionLine.ProductModelID && (v.CurrentState == types.ProductProcessRouteStateWaitProcess || v.CurrentState == types.ProductProcessRouteStateProcessing || v.CurrentState == types.ProductProcessRouteStateReworking || v.CurrentState == types.ProductProcessRouteStateChecking) {
// 			productProcessRoute = v
// 		}
// 	}
// 	if productProcessRoute == nil {
// 		return nil, fmt.Errorf("读取工序最后的工艺路线失败")
// 	}
// 	_productOrder, _ := clients.ProductOrderClient.GetDetail(context.Background(), &proto.GetDetailRequest{Id: productProcessRoute.ProductInfo.ProductOrderID})
// 	if _productOrder.Message == gorm.ErrRecordNotFound.Error() {
// 		return nil, fmt.Errorf("读取当前工单数据失败")
// 	}
// 	if _productOrder.Code != modelcode.Success {
// 		return nil, fmt.Errorf(_productOrder.Message)
// 	}
// 	productOrder := _productOrder.Data
// 	_productionProcessSop, _ := clients.ProductionProcessSopClient.Get(context.Background(), &proto.GetProductionProcessSopRequest{
// 		ProductionProcessID: productionProcess.Id,
// 		ProductModelID:      productOrder.ProductModelID,
// 	})
// 	if _productionProcessSop.Code != modelcode.Success && _productionProcessSop.Message != gorm.ErrRecordNotFound.Error() {
// 		return nil, fmt.Errorf(_productionProcessSop.Message)
// 	}
// 	var productionProcessSop *proto.ProductionProcessSopInfo
// 	if _productionProcessSop.Data != nil && _productionProcessSop.Data.ProductionProcess != nil && _productionProcessSop.Data.ProductionProcess.ProductionLineID == productionStation.ProductionLineID {
// 		productionProcessSop = _productionProcessSop.Data
// 	}
// 	productProcessRoutesTemp := []*proto.ProductProcessRouteInfo{}
// 	for _, v := range _productProcessRoutes.Data {
// 		if v.ProductInfo != nil && v.ProductInfo.ProductOrderID == productOrder.Id {
// 			productProcessRoutesTemp = append(productProcessRoutesTemp, v)
// 		}
// 	}
// 	productProcessRoutesMap := make(map[string]*proto.ProductProcessRouteInfo)
// 	for _, route := range productProcessRoutesTemp {
// 		if _, ok := productProcessRoutesMap[route.ProductInfoID]; !ok {
// 			productProcessRoutesMap[route.ProductInfoID] = route
// 		} else if route.CreateTime > productProcessRoutesMap[route.ProductInfoID].CreateTime {
// 			productProcessRoutesMap[route.ProductInfoID] = route
// 		}
// 	}
// 	var countOfProcessing, countOfProcessed int32
// 	for _, v := range productProcessRoutesMap {
// 		if v.CurrentState == types.ProductProcessRouteStateWaitProcess || v.CurrentState == types.ProductProcessRouteStateProcessing || v.CurrentState == types.ProductProcessRouteStateReworking || v.CurrentState == types.ProductProcessRouteStateChecking {
// 			countOfProcessing++
// 		}
// 		if v.CurrentState == types.ProductProcessRouteStateProcessed {
// 			countOfProcessed++
// 		}
// 	}
// 	_productOrderBoms, _ := clients.ProductOrderBomClient.Query(context.Background(), &proto.QueryProductOrderBomRequest{PageSize: 1000, ProductOrderID: productOrder.Id})
// 	if _productOrderBoms.Code != modelcode.Success {
// 		return nil, fmt.Errorf(_productOrderBoms.Message)
// 	}
// 	productOrderBoms := []*proto.ProductOrderBomInfo{}
// 	for _, v := range _productOrderBoms.Data {
// 		if v.ProductionProcess == "" || v.ProductionProcess == productionProcess.Code || v.ProductionProcess == productionProcess.Identifier {
// 			productOrderBoms = append(productOrderBoms, &proto.ProductOrderBomInfo{
// 				ItemNo:              v.ItemNo,
// 				MaterialNo:          v.MaterialNo,
// 				MaterialDescription: v.MaterialDescription,
// 				PieceQTY:            v.PieceQTY,
// 				RequireQTY:          v.RequireQTY,
// 				Unit:                v.Unit,
// 				Remark:              v.Remark,
// 			})
// 		}
// 	}
// 	return &proto.GetProductionStationExhibitionData{
// 		ProductOrderNo:      productOrder.ProductOrderNo,
// 		SalesOrderNo:        productOrder.SalesOrderNo,
// 		ItemNo:              productOrder.ItemNo,
// 		OrderTime:           productOrder.OrderTime,
// 		OrderQTY:            productOrder.OrderQTY,
// 		ProductCategory:     productOrder.ProductModel.ProductCategory.Code,
// 		ProductModel:        productOrder.ProductModel.Code,
// 		MaterialNo:          productOrder.ProductModel.MaterialNo,
// 		MaterialDescription: productOrder.ProductModel.MaterialDescription,
// 		CurrentState:        productOrder.CurrentState,
// 		PropertyBrief:       productOrder.PropertyBrief,
// 		StartedQTY:          productOrder.StartedQTY,
// 		FinishedQTY:         productOrder.FinishedQTY,
// 		Remark:              productOrder.Remark,
// 		ProductOrderBoms:    productOrderBoms,
// 		ProductionProcess: &proto.ProductionProcessInfo{
// 			Code:              productionProcess.Code,
// 			Description:       productionProcess.Description,
// 			SopLink:           productionProcessSop.FileLink,
// 			CountOfProcessing: countOfProcessing,
// 			CountOfProcessed:  countOfProcessed,
// 		},
// 	}, nil
// }

// 获取作业步骤和作业参数
func GetProductionProcessStepWithParameter(req *proto.GetProductionProcessStepWithParameterRequest) (map[string]interface{}, error) {
	if req.ProductionStation == "" {
		return nil, fmt.Errorf("ProductionStation不能为空")
	}

	if req.TrayNo != "" {
		_materialTray, _ := clients.MaterialTrayClient.Get(context.Background(), &proto.GetMaterialTrayRequest{Identifier: req.TrayNo})
		if _materialTray.Message == gorm.ErrRecordNotFound.Error() {
			return nil, fmt.Errorf("无效的托盘号")
		} else if _materialTray.Code != modelcode.Success {
			return nil, fmt.Errorf(_materialTray.Message)
		}

		if _materialTray.Data.ProductInfoID == "" {
			return nil, fmt.Errorf("托盘未绑定任何产品")
		}
		req.ProductSerialNo = _materialTray.Data.ProductInfo.ProductSerialNo
	}

	if strings.TrimSpace(req.ProductSerialNo) == "" {
		return nil, fmt.Errorf("ProductSerialNo不能为空")
	}

	_productInfo, _ := clients.ProductInfoClient.Get(context.Background(), &proto.GetProductInfoRequest{ProductSerialNo: req.ProductSerialNo})
	if _productInfo.Message == gorm.ErrRecordNotFound.Error() {
		return nil, fmt.Errorf("读取产品信息失败")
	} else if _productInfo.Code != modelcode.Success {
		return nil, fmt.Errorf(_productInfo.Message)
	}
	productInfo := _productInfo.Data
	if productInfo.ProductionProcessID == "" {
		return nil, fmt.Errorf("无法读取产品的当前工序")
	}

	//获取生产工艺
	_productionProcess, _ := clients.ProductionProcessClient.GetDetail(context.Background(), &proto.GetDetailRequest{Id: productInfo.ProductionProcessID})
	if _productionProcess.Message == gorm.ErrRecordNotFound.Error() {
		return nil, fmt.Errorf("读取产品的当前工序失败")
	} else if _productionProcess.Code != modelcode.Success {
		return nil, fmt.Errorf(_productionProcess.Message)
	}
	productionProcess := _productionProcess.Data

	var productionStation *proto.ProductionStationInfo
	for _, v := range productionProcess.ProductionProcessAvailableStations {
		if v.ProductionStation.Code == req.ProductionStation {
			productionStation = v.ProductionStation
			break
		}
	}
	if productionStation == nil {
		return nil, fmt.Errorf("非法操作，产品的当前工序不支持在此工位进行")
	}

	_productOrder, _ := clients.ProductOrderClient.GetDetail(context.Background(), &proto.GetDetailRequest{Id: productInfo.ProductOrderID})
	if _productOrder.Message == gorm.ErrRecordNotFound.Error() {
		return nil, fmt.Errorf("读取产品工单失败")
	} else if _productOrder.Code != modelcode.Success {
		return nil, fmt.Errorf(_productOrder.Message)
	}
	productOrder := _productOrder.Data

	productOrderBoms := []*proto.ProductOrderBomInfo{}
	materialNoArray := []string{}
	for _, v := range productOrder.ProductOrderBoms {
		if v.ProductionProcess == productionProcess.Identifier || v.ProductionProcess == productionProcess.Code {
			materialNoArray = append(materialNoArray, v.MaterialNo)
			productOrderBoms = append(productOrderBoms, &proto.ProductOrderBomInfo{
				MaterialNo:          v.MaterialNo,
				MaterialDescription: v.MaterialDescription,
				PieceQTY:            v.PieceQTY,
				RequireQTY:          v.RequireQTY,
				Unit:                v.Unit,
				EnableControl:       v.EnableControl,
				ControlType:         v.ControlType,
				Warehouse:           v.Warehouse,
				ProductionProcess:   v.ProductionProcess,
			})
		}
	}

	_materialChannels, _ := clients.MaterialChannelLayerClient.GetMaterialChannels(context.Background(), &proto.GetMaterialChannelRequest{ProductionStationID: productionStation.Id})
	if _materialChannels.Code != modelcode.Success {
		return nil, fmt.Errorf(_materialChannels.Message)
	}

	materialChannelLayers := []map[string]interface{}{}
	materialChannelGroup := map[int32][]*proto.MaterialChannelInfo{}
	for _, v := range _materialChannels.Data {
		materialChannelGroup[v.MaterialChannelLayer.LightRegisterAddress] = append(materialChannelGroup[v.MaterialChannelLayer.LightRegisterAddress], v)
	}
	for k, _materialChannels := range materialChannelGroup {
		var lightRegisterValue int32
		for _, v := range _materialChannels {
			for _, materialNo := range materialNoArray {
				if materialNo == v.MaterialInfo.MaterialNo {
					lightRegisterValue += int32(math.Pow(2, float64(v.SortIndex)-1))
					break
				}
			}
		}
		materialChannelLayers = append(materialChannelLayers, map[string]interface{}{
			"lightRegisterAddress": k,
			"lightRegisterValue":   lightRegisterValue,
		})
	}

	_productionProcessStep, _ := clients.ProductionProcessStepClient.Query(context.Background(), &proto.QueryProductionProcessStepRequest{
		SortConfig:          `{"sort_index": ""}`,
		Enable:              true,
		ProductionProcessID: productionProcess.Id,
	})
	if _productionProcessStep.Code != modelcode.Success {
		return nil, fmt.Errorf(_productionProcessStep.Message)
	}

	//TODO
	productionProcessSteps := []map[string]interface{}{}
	for _, v := range _productionProcessStep.Data {
		match := v.InitialValue
		for _, attributeExpression := range v.AttributeExpressions {
			match = false
			for _, productOrderAttribute := range productOrder.ProductOrderAttributes {
				if productOrderAttribute.ProductAttributeID == attributeExpression.ProductAttributeID {
					if b, err := tool.MathOperator(productOrderAttribute.Value, attributeExpression.MathOperator, attributeExpression.AttributeValue); b {
						match = true
						break
					} else if err != nil {
						return nil, err
					}
				}
			}
			if match {
				break
			}
		}
		if match {
			processStepTypeParameters := []map[string]interface{}{}
			for _, v := range v.ProcessStepType.ProcessStepTypeParameters {
				processStepTypeParameters = append(processStepTypeParameters, map[string]interface{}{
					"code":        v.Code,
					"description": v.Description,
					"value":       v.DefaultValue,
				})
			}
			productionProcessSteps = append(productionProcessSteps, map[string]interface{}{
				"id":                  v.Id,
				"sortIndex":           v.SortIndex,
				"code":                v.Code,
				"description":         v.Description,
				"processStepType":     v.ProcessStepType.Code,
				"processStepTypeDesc": v.ProcessStepType.Description,
				"remark":              v.Remark,
				"parameters":          processStepTypeParameters,
				"workResult":          "",
			})
			break
		}
	}

	_productOrderProcess, _ := clients.ProductOrderProcessStepClient.Query(context.Background(), &proto.QueryProductOrderProcessStepRequest{
		PageSize:            1000,
		ProductionProcessID: productionProcess.Id,
		ProductOrderID:      productOrder.Id,
	})
	if _productOrderProcess.Code != modelcode.Success {
		return nil, fmt.Errorf(_productOrderProcess.Message)
	}
	productOrderProcess := []map[string]interface{}{}
	for _, v := range _productOrderProcess.Data {
		productOrderProcessStepTypeParameters := []map[string]interface{}{}
		for _, vv := range v.ProductOrderProcessStepTypeParameters {
			productOrderProcessStepTypeParameters = append(productOrderProcessStepTypeParameters, map[string]interface{}{
				// "Code": vv.Code,
				// "Code": vv.Code,
				"value": vv.Value,
			})
		}
		productOrderProcess = append(productOrderProcess, map[string]interface{}{
			"id":                                    v.Id,
			"sortIndex":                             v.SortIndex,
			"code":                                  "",
			"workDescription":                       v.WorkDescription,
			"processStepType":                       v.ProcessStepType.Code,
			"processStepTypeDesc":                   v.ProcessStepType.Description,
			"workGraphic":                           v.WorkGraphic,
			"remark":                                v.Remark,
			"workResult":                            "",
			"productOrderProcessStepTypeParameters": productOrderProcessStepTypeParameters,
		})
	}

	_productWorkRecord, _ := clients.ProductWorkRecordClient.Query(context.Background(), &proto.QueryProductWorkRecordRequest{ProductInfoID: productInfo.Id})
	if _productWorkRecord.Code != modelcode.Success {
		return nil, fmt.Errorf(_productWorkRecord.Message)
	}
	productWorkRecords := []map[string]interface{}{}
	for _, v := range _productWorkRecord.Data {
		productWorkRecords = append(productWorkRecords, map[string]interface{}{
			"productionProcessStepID": v.ProductionProcessStepID,
			"isQualified":             v.IsQualified,
		})
	}

	for i, v := range productionProcessSteps {
		for _, vv := range productWorkRecords {
			if vv["productionProcessStepID"] == v["id"] {
				workResult := ""
				if vv["isQualified"] == "true" {
					workResult = "PASS"
				}
				productionProcessSteps[i]["workResult"] = workResult
				break
			}
		}
	}

	var materialTray, assembleTray string
	_materialTray, _ := clients.MaterialTrayClient.Get(context.Background(), &proto.GetMaterialTrayRequest{ProductInfoID: productInfo.Id, TrayType: types.MaterialTrayTypeMaterialTray})
	if _materialTray.Code != modelcode.Success && _materialTray.Message != gorm.ErrRecordNotFound.Error() {
		return nil, fmt.Errorf(_materialTray.Message)
	}
	if _materialTray.Data != nil {
		materialTray = _materialTray.Data.Identifier
	}
	_assembleTray, _ := clients.MaterialTrayClient.Get(context.Background(), &proto.GetMaterialTrayRequest{ProductInfoID: productInfo.Id, TrayType: types.MaterialTrayTypeAssembleTray})
	if _assembleTray.Code != modelcode.Success && _assembleTray.Message != gorm.ErrRecordNotFound.Error() {
		return nil, fmt.Errorf(_assembleTray.Message)
	}
	if _assembleTray.Data != nil {
		assembleTray = _assembleTray.Data.Identifier
	}

	var productOrderAttributes []map[string]interface{}
	for _, v := range productOrder.ProductOrderAttributes {
		productOrderAttributes = append(productOrderAttributes, map[string]interface{}{
			"ID":               v.Id,
			"Code":             v.ProductAttribute.Code,
			"CodeDescription":  v.ProductAttribute.Description,
			"Value":            v.Value,
			"ValueDescription": v.Description,
		})
	}

	return map[string]interface{}{
		"productOrder": map[string]interface{}{
			"id":                  productOrder.Id,
			"productOrderNo":      productOrder.ProductOrderNo,
			"productModel":        productOrder.ProductModel.Code,
			"materialNo":          productOrder.ProductModel.MaterialNo,
			"materialDescription": productOrder.ProductModel.MaterialDescription,
			"selectedOptions":     productOrder.SelectedOptions,
			"propertyBrief":       productOrder.PropertyBrief,
			"standardWorkTime":    productOrder.StandardWorkTime,
			"remark":              productOrder.Remark,
			"orderQTY":            productOrder.OrderQTY,
		},
		"productInfo": map[string]interface{}{
			"id":              productInfo.Id,
			"productSerialNo": productInfo.ProductSerialNo,
			"materialTray":    materialTray,
			"assembleTray":    assembleTray,
		},
		"productOrderBoms":       productOrderBoms,
		"materialChannelLayers":  materialChannelLayers,
		"productionProcessSteps": append(productionProcessSteps, productOrderProcess...),
		"productOrderAttributes": productOrderAttributes,
	}, nil
}

// 创建产品过程记录
func CreateProductProcessRecord(req *proto.CreateProductProcessRecordRequest) error {
	if req.ProductionStation == "" {
		return fmt.Errorf("ProductionStation不能为空")
	}
	if req.ProductSerialNo == "" {
		return fmt.Errorf("ProductSerialNo不能为空")
	}
	if req.ProcessStepType == "" {
		return fmt.Errorf("ProcessStepType不能为空")
	}
	if req.WorkDescription == "" {
		return fmt.Errorf("WorkDescription不能为空")
	}
	if req.WorkData == "" {
		return fmt.Errorf("WorkData不能为空")
	}
	if req.WorkResult == "" {
		return fmt.Errorf("WorkResult不能为空")
	}

	productionStation := &model.ProductionStation{}
	if err := model.DB.DB().Where(&model.ProductionStation{Code: req.ProductionStation}).First(productionStation).Error; err == gorm.ErrRecordNotFound {
		return fmt.Errorf("无效的工站代号")
	} else if err != nil {
		return err
	}

	productInfo := &model.ProductInfo{}
	if err := model.DB.DB().Where(&model.ProductInfo{ProductSerialNo: req.ProductSerialNo}).First(productInfo).Error; err == gorm.ErrRecordNotFound {
		return fmt.Errorf("无效的产品序列号")
	} else if err != nil {
		return err
	}

	if err := model.DB.DB().Create(&model.ProductProcessRecord{
		ProductInfoID:       productInfo.ID,
		ProductionProcessID: productInfo.ProductionProcessID,
		ProductionStationID: productionStation.ID,
		ProcessStepType:     req.ProcessStepType,
		WorkDescription:     req.WorkDescription,
		WorkData:            req.WorkData,
		WorkResult:          req.WorkResult,
		WorkTime:            sql.NullTime{Time: time.Now(), Valid: true},
	}).Error; err != nil {
		return err
	}

	return nil
}

// 创建产品作业记录
func CreateProductWorkRecord(req *proto.CreateProductWorkRecordRequest) error {
	if req.WorkStartTime == "" {
		return fmt.Errorf("WorkStartTime不能为空")
	}
	if req.WorkEndTime == "" {
		return fmt.Errorf("WorkEndTime不能为空")
	}
	// if req.IsQualified == "" {
	// 	return fmt.Errorf("IsQualified不能为空")
	// }
	if req.WorkData == "" {
		return fmt.Errorf("WorkData不能为空")
	}
	if req.ProductionStation == "" {
		return fmt.Errorf("ProductionStation不能为空")
	}
	if req.ProductSerialNo == "" {
		return fmt.Errorf("ProductSerialNo不能为空")
	}
	if req.ProductionProcessStep == "" {
		return fmt.Errorf("ProductionProcessStep不能为空")
	}

	productInfo := &model.ProductInfo{}
	if err := model.DB.DB().Where(&model.ProductInfo{ProductSerialNo: req.ProductSerialNo}).First(productInfo).Error; err == gorm.ErrRecordNotFound {
		return fmt.Errorf("无效的产品信息")
	} else if err != nil {
		return err
	}

	if productInfo.ProductionProcessID == "" {
		return fmt.Errorf("无法获取产品的当前工序")
	}

	productionProcess := &model.ProductionProcess{}
	if err := model.DB.DB().
		Preload("ProductionProcessAvailableStations").
		Preload("ProductionProcessAvailableStations.ProductionStation").
		First(productionProcess, "`id` = ?", productInfo.ProductionProcessID).Error; err == gorm.ErrRecordNotFound {
		return fmt.Errorf("读取产品的当前工序失败")
	} else if err != nil {
		return err
	}

	var processable bool
	for _, v := range productionProcess.ProductionProcessAvailableStations {
		if v.ProductionStation.Code == req.ProductionStation {
			processable = true
			break
		}
	}
	if !processable {
		return fmt.Errorf("非法操作，产品的当前工序不支持在此工位进行")
	}

	productionProcessStep := &model.ProductionProcessStep{}
	if err := model.DB.DB().
		Preload("AvailableProcesses").
		Where(&model.ProductionProcessStep{Code: req.ProductionProcessStep}).First(productionProcessStep).Error; err == gorm.ErrRecordNotFound {
		return fmt.Errorf("无效的作业步骤")
	} else if err != nil {
		return err
	}

	var workable bool
	for _, v := range productionProcessStep.AvailableProcesses {
		if v.ProductionProcessID == productionProcess.ID {
			workable = true
			break
		}
	}
	if !workable {
		return fmt.Errorf("非法操作，此测试项不支持在此工位进行")
	}

	productionStation := &model.ProductionStation{}
	if err := model.DB.DB().Where(&model.ProductionStation{Code: req.ProductionStation}).First(productionStation).Error; err == gorm.ErrRecordNotFound {
		return fmt.Errorf("无效的产线工位")
	} else if err != nil {
		return err
	}

	workEndTime := utils.ParseTime(req.WorkEndTime)
	workStartTime := utils.ParseTime(req.WorkStartTime)
	if err := model.DB.DB().Create(&model.ProductWorkRecord{
		ProductionProcessStepID: productionProcessStep.ID,
		ProductInfoID:           productInfo.ID,
		ProductionStationID:     productionStation.ID,
		WorkUserID:              *productionStation.CurrentUserID,
		WorkData:                req.WorkData,
		WorkEndTime:             sql.NullTime{Time: workEndTime, Valid: true},
		WorkStartTime:           sql.NullTime{Time: workStartTime, Valid: true},
		Duration:                int32(workEndTime.Sub(workStartTime).Seconds()),
		IsQualified:             req.IsQualified,
	}).Error; err != nil {
		return err
	}

	return nil
}

// 请求出站
func ExitProductionStation(req *proto.ExitProductionStationRequest) (*proto.CommonResponse, error) {
	if req.ProductionStation == "" {
		return nil, fmt.Errorf("ProductionStation不能为空")
	}
	timeNow := time.Now()
	timeNowStr := timeNow.Format("2006-01-02 15:04:05")
	if req.TrayNo != "" {
		//根据托盘号获取物料托盘
		_materialTray, _ := clients.MaterialTrayClient.Get(context.Background(), &proto.GetMaterialTrayRequest{Identifier: req.TrayNo})
		if _materialTray.Message == gorm.ErrRecordNotFound.Error() {
			return nil, fmt.Errorf("无效的托盘号")
		}
		if _materialTray.Code != modelcode.Success {
			return nil, fmt.Errorf(_materialTray.Message)
		}

		materialTray := _materialTray.Data
		if materialTray.ProductionLineID == "" {
			return nil, fmt.Errorf("托盘未绑定任何产品")
		}

		req.ProductSerialNo = materialTray.ProductInfo.ProductSerialNo
	}
	if req.PackageNo != "" {
		//根据包装箱号获取产品包装记录
		_productPackageRecord, _ := clients.ProductPackageRecordClient.Get(context.Background(), &proto.GetProductPackageRecordRequest{PackageNo: req.PackageNo})
		if _productPackageRecord.Message == gorm.ErrRecordNotFound.Error() {
			return nil, fmt.Errorf("无效的包装箱号")
		}
		if _productPackageRecord.Code != modelcode.Success {
			return nil, fmt.Errorf(_productPackageRecord.Message)
		}

		productPackageRecord := _productPackageRecord.Data
		if productPackageRecord.ProductInfoID == "" {
			return &proto.CommonResponse{Code: 5, Message: "包装箱未绑定任何产品"}, nil
		}

		req.ProductSerialNo = productPackageRecord.ProductInfo.ProductSerialNo
	}
	req.ProductSerialNo = strings.Trim(strings.Trim(req.ProductSerialNo, "\000"), "\r")
	if req.ProductSerialNo == "" {
		return nil, fmt.Errorf("ProductSerialNo不能为空")
	}

	_productionStation, _ := clients.ProductionStationClient.Get(context.Background(), &proto.GetProductionStationRequest{Code: req.ProductionStation})
	if _productionStation.Message == gorm.ErrRecordNotFound.Error() {
		return nil, fmt.Errorf("非法的工站代号")
	}
	if _productionStation.Code != modelcode.Success {
		return nil, fmt.Errorf(_productionStation.Message)
	}
	productionStation := _productionStation.Data

	_productInfo, _ := clients.ProductInfoClient.Get(context.Background(), &proto.GetProductInfoRequest{ProductSerialNo: req.ProductSerialNo})
	if _productInfo.Message == gorm.ErrRecordNotFound.Error() {
		return nil, fmt.Errorf("读取产品信息失败")
	}
	if _productInfo.Code != modelcode.Success {
		return nil, fmt.Errorf(_productInfo.Message)
	}
	productInfo := _productInfo.Data
	if productInfo.ProductionProcessID == "" {
		return nil, fmt.Errorf("无法读取产品的当前工序")
	}

	//上传节拍
	_productRhythmRecord, _ := clients.ProductRhythmRecordClient.Get(context.Background(), &proto.GetProductRhythmRecordRequest{ProductionStationID: productionStation.Id, ProductInfoID: productInfo.Id, HasWorkEndTime: false})
	if _productRhythmRecord.Message == gorm.ErrRecordNotFound.Error() {
		return nil, fmt.Errorf("读取工位当前节拍数据失败")
	}
	if _productRhythmRecord.Code != modelcode.Success {
		return nil, fmt.Errorf(_productRhythmRecord.Message)
	}
	productRhythmRecord := _productRhythmRecord.Data

	productRhythmRecord.WorkEndTime = timeNowStr
	workStartTime, err := time.Parse("2006-01-02 15:04:05", productRhythmRecord.WorkStartTime)
	if err != nil {
		return nil, err
	}
	workEndTime, err := time.Parse("2006-01-02 15:04:05", productRhythmRecord.WorkEndTime)
	if err != nil {
		return nil, err
	}
	productRhythmRecord.WorkTime = int32(workEndTime.Sub(workStartTime).Seconds())
	productRhythmRecord.OverTime = productRhythmRecord.WorkTime - productRhythmRecord.StandardWorkTime
	if productRhythmRecord.OverTime < 0 {
		productRhythmRecord.OverTime = 0
	}
	productRhythmRecord.WaitTime = req.WaitTime

	//修改工艺记录
	_productProcessRoute, _ := clients.ProductProcessRouteClient.Get(context.Background(), &proto.GetProductProcessRouteRequest{ProductInfoID: productInfo.Id, CurrentProcessID: productInfo.ProductionProcessID, CurrentStates: []string{types.ProductProcessRouteStateProcessing}})
	if _productProcessRoute.Message == gorm.ErrRecordNotFound.Error() {
		return nil, fmt.Errorf("读取产品当前工艺路线失败")
	}
	if _productProcessRoute.Code != modelcode.Success {
		return nil, fmt.Errorf(_productProcessRoute.Message)
	}
	lastProductProcessRoute := _productProcessRoute.Data

	if req.IsRework {
		productRhythmRecord.IsRework = true

		lastProductProcessRoute.CurrentState = types.ProductProcessRouteStateReworking

		if _productReworkRecord, _ := clients.ProductReworkRecordClient.Add(context.Background(), &proto.ProductReworkRecordInfo{
			ProductionStationID: productionStation.Id,
			ProductInfoID:       productInfo.Id,
			ProductionProcessID: productInfo.ProductionProcessID,
			ReworkTime:          timeNowStr,
			ReworkReason:        req.ReworkReason,
		}); _productReworkRecord.Code != modelcode.Success {
			return nil, fmt.Errorf(_productReworkRecord.Message)
		}

		productInfo.CurrentState = types.ProductStateReworking
	} else if req.IsFail {
		productRhythmRecord.IsRework = true
		lastProductProcessRoute.Remark = req.ReworkReason
		lastProductProcessRoute.CurrentState = types.ProductProcessRouteStateChecking
		productInfo.CurrentState = types.ProductStateChecking
	} else {
		lastProductProcessRoute.CurrentState = types.ProductProcessRouteStateProcessed

		//切换到下个工艺
		_productProcessRoutes, _ := clients.ProductProcessRouteClient.Query(context.Background(), &proto.QueryProductProcessRouteRequest{
			PageSize:      1,
			SortConfig:    `{"work_index": "asc"}`,
			ProductInfoID: productInfo.Id,
			RouteIndex:    lastProductProcessRoute.RouteIndex,
			CurrentState:  types.ProductProcessRouteStateWaitProcess,
		})
		if _productProcessRoutes.Code != modelcode.Success {
			return nil, fmt.Errorf(_productProcessRoutes.Message)
		}
		var nextProductProcessRoute *proto.ProductProcessRouteInfo
		if len(_productProcessRoutes.Data) > 0 {
			nextProductProcessRoute = _productProcessRoutes.Data[0]
		}
		//兼容，部分产线是直接创建产品工艺路线，部分是根据工单工艺动态创建
		if nextProductProcessRoute == nil {
			_productOrderProcesses, _ := clients.ProductOrderProcessClient.Query(context.Background(), &proto.QueryProductOrderProcessRequest{
				ProductOrderID: productInfo.ProductOrderID,
				Enable:         true,
				SortIndex:      lastProductProcessRoute.RouteIndex,
				SortConfig:     `{"sort_index": "asc"}`,
				PageSize:       1,
			})
			if _productOrderProcesses.Code != modelcode.Success {
				return nil, fmt.Errorf(_productOrderProcesses.Message)
			}
			if len(_productOrderProcesses.Data) > 0 {
				productOrderProcess := _productOrderProcesses.Data[0]
				nextProductProcessRoute = &proto.ProductProcessRouteInfo{
					LastProcessID:    lastProductProcessRoute.CurrentProcessID,
					CurrentProcessID: productOrderProcess.ProductionProcessID,
					CurrentProcess:   productOrderProcess.ProductionProcess,
					CreateTime:       timeNowStr,
					CurrentState:     types.ProductProcessRouteStateWaitProcess,
					RouteIndex:       productOrderProcess.SortIndex,
					ProductInfoID:    productInfo.Id,
				}
				if _productProcessRoute, _ := clients.ProductProcessRouteClient.Add(context.Background(), nextProductProcessRoute); _productProcessRoute.Code != modelcode.Success {
					return nil, fmt.Errorf(_productProcessRoute.Message)
				}
			}
		}

		if nextProductProcessRoute != nil {
			nextProductProcessRoute.WorkIndex = lastProductProcessRoute.WorkIndex + 1
			if nextProductProcessRoute.CurrentProcess != nil {
				if nextProductProcessRoute.CurrentProcess.ProductState != "" {
					//设定当前工序的产品信息状态
					productInfo.CurrentState = nextProductProcessRoute.CurrentProcess.ProductState
				}
			}
			productInfo.ProductionProcessID = nextProductProcessRoute.CurrentProcessID

			//计算预计下线时间
			_productOrderProcesses, _ := clients.ProductOrderProcessClient.Query(context.Background(), &proto.QueryProductOrderProcessRequest{
				ProductOrderID: productInfo.ProductOrderID,
				Enable:         true,
				SortIndex:      nextProductProcessRoute.RouteIndex,
			})
			if _productOrderProcesses.Code != modelcode.Success {
				return nil, fmt.Errorf(_productOrderProcesses.Message)
			}
			remainingRoutes := int32(_productOrderProcesses.Total)
			productInfo.RemainingRoutes = remainingRoutes
			if remainingRoutes > 0 {
				productInfo.EstimateTime = timeNow.Add(time.Duration(remainingRoutes*productInfo.ProductOrder.StandardWorkTime) * time.Second).Format("2006-01-02 15:04:05")
			}
		} else {
			//没有下一个工序判定为完工
			productInfo.CurrentState = types.ProductStateCompleted
			productInfo.FinishedTime = timeNowStr
			productInfo.ProductionProcessID = ""

			_productOrder, _ := clients.ProductOrderClient.GetDetail(context.Background(), &proto.GetDetailRequest{Id: productInfo.ProductOrderID})
			if _productOrder.Code != modelcode.Success {
				return nil, fmt.Errorf(_productOrder.Message)
			}
			productOrder := _productOrder.Data
			var finishedCount int32
			for _, _productInfo := range productOrder.ProductInfos {
				if _productInfo.CurrentState == types.ProductStateCompleted {
					finishedCount++
				}
			}
			productOrder.FinishedQTY = finishedCount
			totalCount := int32(len(productOrder.ProductInfos))
			if totalCount == finishedCount {
				productOrder.CurrentState = types.ProductOrderStateCompleted
				productOrder.ActualFinishTime = timeNowStr

				_systemEvent, _ := clients.SystemEventClient.Get(context.Background(), &proto.GetSystemEventRequest{Code: types.SystemEventProductOrderFinished})
				if _systemEvent.Code != modelcode.Success && _systemEvent.Message != gorm.ErrRecordNotFound.Error() {
					return nil, fmt.Errorf(_systemEvent.Message)
				}

				if _systemEvent.Data != nil {
					systemEvent := _systemEvent.Data
					systemEventTrigger := &proto.SystemEventTriggerInfo{
						SystemEventID: systemEvent.Id,
						CreateTime:    timeNowStr,
						EventNo:       uuid.NewString(),
						CurrentState:  types.SystemEventTriggerStateWaitExecute,
					}
					for _, _systemEventParameter := range systemEvent.SystemEventParameters {
						value := strings.ReplaceAll(_systemEventParameter.Value, "{ProductOrderNo}", productOrder.ProductOrderNo)
						systemEventTrigger.SystemEventTriggerParameters = append(systemEventTrigger.SystemEventTriggerParameters, &proto.SystemEventTriggerParameterInfo{
							DataType:    _systemEventParameter.DataType,
							Name:        _systemEventParameter.Name,
							Description: _systemEventParameter.Description,
							Value:       value,
						})
					}

					if _systemEventTrigger, _ := clients.SystemEventTriggerClient.Add(context.Background(), systemEventTrigger); _systemEventTrigger.Code != modelcode.Success {
						return nil, fmt.Errorf(_systemEventTrigger.Message)
					}
				}
			}
			if _productOrder, _ := clients.ProductOrderClient.Update(context.Background(), productOrder); _productOrder.Code != modelcode.Success {
				return nil, fmt.Errorf(_productOrder.Message)
			}
		}

		//创建产量记录
		_ProductionStationOutput, _ := clients.ProductionStationOutputClient.Get(context.Background(), &proto.GetProductionStationOutputRequest{
			ProductionStationID: productionStation.Id,
			ProductInfoID:       productInfo.Id,
		})
		if _ProductionStationOutput.Code != modelcode.Success && _ProductionStationOutput.Message != gorm.ErrRecordNotFound.Error() {
			return nil, fmt.Errorf(_ProductionStationOutput.Message)
		}

		productionStationOutput := _ProductionStationOutput.Data
		if productionStationOutput != nil {
			productionStationOutput.OutputTime = timeNowStr
			productionStationOutput.LoginUserID = productionStation.CurrentUserID
			if _productionStationOutput, _ := clients.ProductionStationOutputClient.Update(context.Background(), productionStationOutput); _productionStationOutput.Code != modelcode.Success {
				return nil, fmt.Errorf(_productionStationOutput.Message)
			}
		} else {
			productionStationOutput = &proto.ProductionStationOutputInfo{
				OutputTime:          timeNowStr,
				LoginUserID:         productionStation.CurrentUserID,
				ProductionProcessID: productRhythmRecord.ProductionProcessID,
				ProductInfoID:       productRhythmRecord.ProductInfoID,
				ProductionStationID: productRhythmRecord.ProductionStationID,
			}
			if _productionStationOutput, _ := clients.ProductionStationOutputClient.Add(context.Background(), productionStationOutput); _productionStationOutput.Code != modelcode.Success {
				return nil, fmt.Errorf(_productionStationOutput.Message)
			}
		}

		// 判断是否要解绑托盘
		if req.UnbindTray {
			_materialTray, _ := clients.MaterialTrayClient.Get(context.Background(), &proto.GetMaterialTrayRequest{ProductInfoID: productInfo.Id})
			if _materialTray.Code != modelcode.Success && _materialTray.Message != gorm.ErrRecordNotFound.Error() {
				return nil, fmt.Errorf(_materialTray.Message)
			}
			if _materialTray.Data != nil {
				materialTray := _materialTray.Data
				materialTray.ProductInfoID = ""
				materialTray.CurrentState = types.MaterialTrayStateWaitBind
				if _materialTray, _ := clients.MaterialTrayClient.Update(context.Background(), materialTray); _materialTray.Code != modelcode.Success {
					return nil, fmt.Errorf(_materialTray.Message)
				}
			}
		}

		_productionProcess, _ := clients.ProductionProcessClient.GetDetail(context.Background(), &proto.GetDetailRequest{Id: productRhythmRecord.ProductionProcessID})
		if _productionProcess.Code != modelcode.Success && _productionProcess.Message != gorm.ErrRecordNotFound.Error() {
			return nil, fmt.Errorf(_productionProcess.Message)
		}
		if _productionProcess.Data != nil {
			productionProcess := _productionProcess.Data
			if productionProcess.EnableReport {
				_systemEvent, _ := clients.SystemEventClient.Get(context.Background(), &proto.GetSystemEventRequest{Code: types.SystemEventProductionProcessFinished, Enable: true})
				if _systemEvent.Code != modelcode.Success && _systemEvent.Message != gorm.ErrRecordNotFound.Error() {
					return nil, fmt.Errorf(_systemEvent.Message)
				}
				if _systemEvent.Data != nil {
					systemEvent := _systemEvent.Data
					systemEventTrigger := &proto.SystemEventTriggerInfo{
						SystemEventID: systemEvent.Id,
						CreateTime:    timeNowStr,
						EventNo:       uuid.NewString(),
						CurrentState:  types.SystemEventTriggerStateWaitExecute,
					}

					for _, _systemEventParameter := range systemEvent.SystemEventParameters {
						value := _systemEventParameter.Value
						value = strings.ReplaceAll(value, "{ProductionProcess.Identifier}", productionProcess.Identifier)
						value = strings.ReplaceAll(value, "{ProductionLine.Identifier}", productionStation.ProductionLine.Identifier)
						value = strings.ReplaceAll(value, "{ProductionStation}", productionStation.Code)
						value = strings.ReplaceAll(value, "{ProductSerialNo}", productInfo.ProductSerialNo)
						value = strings.ReplaceAll(value, "{ProductionProcess}", productionProcess.Code)

						systemEventTrigger.SystemEventTriggerParameters = append(systemEventTrigger.SystemEventTriggerParameters, &proto.SystemEventTriggerParameterInfo{
							DataType:    _systemEventParameter.DataType,
							Name:        _systemEventParameter.Name,
							Description: _systemEventParameter.Description,
							Value:       value,
						})
					}

					if _systemEventTrigger, _ := clients.SystemEventTriggerClient.Add(context.Background(), systemEventTrigger); _systemEventTrigger.Code != modelcode.Success {
						return nil, fmt.Errorf(_systemEventTrigger.Message)
					}
				}
			}
		}

		if productionStation.AllowReport {
			productionStation.CurrentState = types.ProductionStationStateStandby
			// 创建系统事件上报工位完成作业
			_systemEvent, _ := clients.SystemEventClient.Get(context.Background(), &proto.GetSystemEventRequest{Code: types.SystemEventProductionStationReleased, Enable: true})
			if _systemEvent.Code != modelcode.Success && _systemEvent.Message != gorm.ErrRecordNotFound.Error() {
				return nil, fmt.Errorf(_systemEvent.Message)
			}
			if _systemEvent.Data != nil {
				systemEvent := _systemEvent.Data
				systemEventTrigger := &proto.SystemEventTriggerInfo{
					SystemEventID: systemEvent.Id,
					CreateTime:    timeNowStr,
					EventNo:       uuid.NewString(),
					CurrentState:  types.SystemEventTriggerStateWaitExecute,
				}

				for _, _systemEventParameter := range systemEvent.SystemEventParameters {
					value := strings.ReplaceAll(_systemEventParameter.Value, "{ProductionStation}", productionStation.Code)

					systemEventTrigger.SystemEventTriggerParameters = append(systemEventTrigger.SystemEventTriggerParameters, &proto.SystemEventTriggerParameterInfo{
						DataType:    _systemEventParameter.DataType,
						Name:        _systemEventParameter.Name,
						Description: _systemEventParameter.Description,
						Value:       value,
					})
				}

				if _systemEventTrigger, _ := clients.SystemEventTriggerClient.Add(context.Background(), systemEventTrigger); _systemEventTrigger.Code != modelcode.Success {
					return nil, fmt.Errorf(_systemEventTrigger.Message)
				}
			}
		}
	}

	if _productRhythmRecord, _ := clients.ProductRhythmRecordClient.Update(context.Background(), productRhythmRecord); _productRhythmRecord.Code != modelcode.Success {
		return nil, fmt.Errorf(_productRhythmRecord.Message)
	}
	if _productProcessRoute, _ := clients.ProductProcessRouteClient.Update(context.Background(), lastProductProcessRoute); _productProcessRoute.Code != modelcode.Success {
		return nil, fmt.Errorf(_productProcessRoute.Message)
	}
	if _productInfo, _ := clients.ProductInfoClient.Update(context.Background(), productInfo); _productInfo.Code != modelcode.Success {
		return nil, fmt.Errorf(_productInfo.Message)
	}

	return &proto.CommonResponse{Code: 0}, nil
}

// 设置失败后续处理
func CheckProductProcessRouteFailure(req *proto.CheckProductProcessRouteFailureRequest) error {
	if req.ProductionStation == "" {
		return fmt.Errorf("ProductionStation不能为空")
	}
	if req.ProductSerialNo == "" {
		return fmt.Errorf("ProductSerialNo不能为空")
	}

	timeNow := time.Now()
	timeNowStr := timeNow.Format("2006-01-02 15:04:05")
	_productionStation, _ := clients.ProductionStationClient.Get(context.Background(), &proto.GetProductionStationRequest{Code: req.ProductionStation})
	if _productionStation.Message == gorm.ErrRecordNotFound.Error() {
		return fmt.Errorf("无效的工站代号")
	} else if _productionStation.Code != modelcode.Success {
		return fmt.Errorf(_productionStation.Message)
	}
	productionStation := _productionStation.Data

	_productInfo, _ := clients.ProductInfoClient.Get(context.Background(), &proto.GetProductInfoRequest{ProductSerialNo: req.ProductSerialNo})
	if _productInfo.Message == gorm.ErrRecordNotFound.Error() {
		return fmt.Errorf("无效的产品序列号")
	} else if _productInfo.Code != modelcode.Success {
		return fmt.Errorf(_productInfo.Message)
	}
	productInfo := _productInfo.Data

	if productInfo.ProductionProcessID == "" {
		return fmt.Errorf("数据错误，此产品的当前工序为空")
	}

	_productProcessRoute, _ := clients.ProductProcessRouteClient.Get(context.Background(), &proto.GetProductProcessRouteRequest{
		CurrentProcessID: productInfo.ProductionProcessID,
		ProductInfoID:    productInfo.Id,
		CurrentStates:    []string{types.ProductProcessRouteStateChecking},
	})
	if _productProcessRoute.Message == gorm.ErrRecordNotFound.Error() {
		return fmt.Errorf("状态错误，此产品的当前工艺状态不是" + types.ProductProcessRouteStateChecking)
	} else if _productProcessRoute.Code != modelcode.Success {
		return fmt.Errorf(_productProcessRoute.Message)
	}
	lastProductProcessRoute := _productProcessRoute.Data

	handleMethod := req.HandleMethod
	if handleMethod == 0 {
		handleMethod = types.ProductionProcessHandleMethodRetry
	}
	switch handleMethod {
	case types.ProductionProcessHandleMethodRetry:
		lastProductProcessRoute.Remark = ""
		lastProductProcessRoute.CurrentState = types.ProductProcessRouteStateProcessing

		productInfo.CurrentState = lastProductProcessRoute.CurrentProcess.ProductState
		if productInfo.CurrentState == "" {
			productInfo.CurrentState = types.ProductStateTesting
		}
	case types.ProductionProcessHandleMethodRework:
		lastProductProcessRoute.CurrentState = types.ProductProcessRouteStateReworking

		productReworkRecord := &proto.ProductReworkRecordInfo{
			ProductionStationID: productionStation.Id,
			ProductInfoID:       productInfo.Id,
			ProductionProcessID: productInfo.ProductionProcessID,
			ReworkTime:          timeNowStr,
			ReworkReason:        lastProductProcessRoute.Remark,
		}
		if resp, _ := clients.ProductReworkRecordClient.Add(context.Background(), productReworkRecord); resp.Code != modelcode.Success {
			return fmt.Errorf(resp.Message)
		}

		productInfo.CurrentState = types.ProductStateReworking
	case types.ProductionProcessHandleMethodIgnore:
		lastProductProcessRoute.CurrentState = types.ProductProcessRouteStateProcessed

		//切换到下个工艺
		_productProcessRoutes, _ := clients.ProductProcessRouteClient.Query(context.Background(), &proto.QueryProductProcessRouteRequest{
			PageSize:      1,
			SortConfig:    `{"route_index": "asc"}`,
			ProductInfoID: productInfo.Id,
			RouteIndex:    lastProductProcessRoute.RouteIndex,
			CurrentState:  types.ProductProcessRouteStateWaitProcess,
		})
		if _productProcessRoutes.Code != modelcode.Success {
			return fmt.Errorf(_productProcessRoutes.Message)
		}
		var nextProductProcessRoute *proto.ProductProcessRouteInfo
		if len(_productProcessRoutes.Data) > 0 {
			nextProductProcessRoute = _productProcessRoutes.Data[0]
		}
		if nextProductProcessRoute == nil {
			_productOrderProcesses, _ := clients.ProductOrderProcessClient.Query(context.Background(), &proto.QueryProductOrderProcessRequest{
				PageSize:       1,
				SortConfig:     `{"route_index": "asc"}`,
				ProductOrderID: productInfo.ProductOrderID,
				Enable:         true,
				SortIndex:      lastProductProcessRoute.RouteIndex,
			})
			if _productOrderProcesses.Code != modelcode.Success {
				return fmt.Errorf(_productOrderProcesses.Message)
			}
			if len(_productOrderProcesses.Data) == 0 {
				productOrderProcess := _productOrderProcesses.Data[0]
				nextProductProcessRoute = &proto.ProductProcessRouteInfo{
					LastProcessID:    lastProductProcessRoute.CurrentProcessID,
					CurrentProcessID: productOrderProcess.ProductionProcessID,
					CurrentProcess:   productOrderProcess.ProductionProcess,
					CreateTime:       timeNowStr,
					CurrentState:     types.ProductProcessRouteStateWaitProcess,
					RouteIndex:       productOrderProcess.SortIndex,
					ProductInfoID:    productInfo.Id,
				}
				if resp, _ := clients.ProductProcessRouteClient.Add(context.Background(), nextProductProcessRoute); resp.Code != modelcode.Success {
					return fmt.Errorf(resp.Message)
				}
			}

			if nextProductProcessRoute != nil {
				nextProductProcessRoute.WorkIndex = lastProductProcessRoute.WorkIndex + 1
				nextProductProcessRoute.CurrentState = types.ProductProcessRouteStateProcessing

				if nextProductProcessRoute.CurrentProcess.ProductState != "" {
					//设定当前工序的产品信息状态
					productInfo.CurrentState = nextProductProcessRoute.CurrentProcess.ProductState
				}
				productInfo.ProductionProcessID = nextProductProcessRoute.CurrentProcessID
				//TODO: 计算预计下线时间
				_productProcessRoutes, _ := clients.ProductProcessRouteClient.Query(context.Background(), &proto.QueryProductProcessRouteRequest{
					ProductInfoID: productInfo.Id,
					CurrentState:  types.ProductProcessRouteStateWaitProcess,
				})
				if _productProcessRoutes.Code != modelcode.Success {
					return fmt.Errorf(_productProcessRoutes.Message)
				}
				remainingRoutes := int32(_productProcessRoutes.Total)
				productInfo.RemainingRoutes = remainingRoutes
				if remainingRoutes > 0 {
					productInfo.EstimateTime = timeNow.Add(time.Duration(remainingRoutes*productInfo.ProductOrder.StandardWorkTime) * time.Second).Format("2006-01-02 15:04:05")
				}
				if resp, _ := clients.ProductProcessRouteClient.Update(context.Background(), nextProductProcessRoute); resp.Code != modelcode.Success {
					return fmt.Errorf(resp.Message)
				}
			} else {
				productInfo.CurrentState = types.ProductStateCompleted
				productInfo.FinishedTime = timeNowStr
				productInfo.ProductionProcessID = ""
			}
		}
	default:
		return fmt.Errorf("无效的处理方式")
	}

	if resp, _ := clients.ProductInfoClient.Update(context.Background(), productInfo); resp.Code != modelcode.Success {
		return fmt.Errorf(resp.Message)
	}
	if resp, _ := clients.ProductProcessRouteClient.Update(context.Background(), lastProductProcessRoute); resp.Code != modelcode.Success {
		return fmt.Errorf(resp.Message)
	}

	return nil
}

// 获取产品返工记录
func RetrieveProductReworkRecord(req *proto.RetrieveProductReworkRecordRequest) ([]map[string]interface{}, error) {
	if req.ProductionLine == "" {
		return nil, fmt.Errorf("ProductionLine不能为空")
	}

	var productReworkRecords []*model.ProductReworkRecord
	db := model.DB.DB().Preload("ProductInfo").Preload("ProductInfo.ProductOrder").Preload("ProductionStation").
		Joins("JOIN product_infos ON product_rework_records.product_info_id=product_infos.id").
		Joins("JOIN production_stations ON product_rework_records.production_station_id=production_stations.id").
		Joins("JOIN production_lines ON production_stations.production_line_id=production_lines.id").
		Where("production_lines.code = ?", req.ProductionLine)

	if req.ProductSerialNo != "" {
		db = db.Where("product_infos.product_serial_no LIKE ?", "%"+req.ProductSerialNo+"%")
	}
	if req.StartDate != "" {
		db = db.Where("product_rework_records.rework_time >= ?", req.StartDate)
	}
	if req.FinishDate != "" {
		db = db.Where("product_rework_records.rework_time <= ?", req.FinishDate)
	}
	if req.IsCompleted {
		db = db.Where("product_rework_records.complete_time IS NULL")
	}
	if err := db.Find(&productReworkRecords).Error; err != nil {
		return nil, err
	}
	data := make([]map[string]interface{}, len(productReworkRecords))
	for i, v := range productReworkRecords {
		data[i] = map[string]interface{}{
			"id":                v.ID,
			"reworkReason":      v.ReworkReason,
			"reworkTime":        utils.FormatSqlNullTime(v.ReworkTime),
			"completeTime":      utils.FormatSqlNullTime(v.CompleteTime),
			"productOrderNo":    v.ProductInfo.ProductOrder.ProductOrderNo,
			"productSerialNo":   v.ProductInfo.ProductSerialNo,
			"productionStation": v.ProductionStation.Description,
		}
	}

	return data, nil
}

// 更新产品返工记录
func UpdateProductReworkRecord(req *proto.UpdateProductReworkRecordRequest) (map[string]interface{}, error) {
	if req.ProductReworkRecordID == "" {
		return nil, fmt.Errorf("ProductReworkRecordID不能为空")
	}
	if req.ProductReworkCauseID == "" {
		return nil, fmt.Errorf("ProductReworkCauseID不能为空")
	}
	if req.ProductReworkSolutionID == "" {
		return nil, fmt.Errorf("ProductReworkSolutionID不能为空")
	}
	if req.ProductReworkTypeID == "" {
		return nil, fmt.Errorf("ProductReworkTypeID不能为空")
	}
	if req.ProductReworkUserID == "" {
		return nil, fmt.Errorf("ProductReworkUserID不能为空")
	}

	productReworkRecord := &model.ProductReworkRecord{}
	nextProcess := &model.ProductionProcess{}
	if err := model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if err := tx.First(productReworkRecord, "`id` = ?", req.ProductReworkRecordID).Error; err == gorm.ErrRecordNotFound {
			return fmt.Errorf("无效的返工记录ID")
		} else if err != nil {
			return err
		}

		productReworkRecord.ProductReworkCauseID = req.ProductReworkCauseID
		productReworkRecord.ProductReworkSolutionID = req.ProductReworkSolutionID
		productReworkRecord.ProductReworkTypeID = req.ProductReworkTypeID
		productReworkRecord.ReworkUserID = req.ProductReworkUserID
		productReworkRecord.ReworkBrief = req.ReworkBrief
		productReworkRecord.Remark = req.Remark
		productReworkRecord.CompleteTime = sql.NullTime{Time: time.Now(), Valid: true}

		// user, _ := clients.UserClient.GetDetail(context.Background(), &usercenter.GetDetailRequest{Id: req.ProductReworkUserID})

		productInfo := &model.ProductInfo{}
		if err := tx.Preload("ProductOrder").First(productInfo, "`id` = ?", productReworkRecord.ProductInfoID).Error; err == gorm.ErrRecordNotFound {
			return fmt.Errorf("读取产品信息失败")
		} else if err != nil {
			return err
		}

		systemEvent := &model.SystemEvent{}
		if err := tx.Preload("SystemEventParameters").First(systemEvent, map[string]interface{}{
			"code":   types.SystemEventProductInfoReworked,
			"enable": true,
		}).Error; err != nil && err != gorm.ErrRecordNotFound {
			return err
		}

		if systemEvent.ID != "" {
			var reworkCause string
			if err := tx.Model(model.ProductReworkCause{}).Select("`id` = ?", req.ProductReworkCauseID).Where("code").Scan(&reworkCause).Error; err != nil {
				return err
			}

			productionStation := &model.ProductionStation{}
			if err := tx.First(productionStation, "`id` = ?", productReworkRecord.ProductionStationID).Error; err == gorm.ErrRecordNotFound {
				return fmt.Errorf("读取生产工站失败")
			} else if err != nil {
				return err
			}

			productionProcess := &model.ProductionProcess{}
			if err := tx.First(productionProcess, "`id` = ?", productReworkRecord.ProductionProcessID).Error; err != nil && err != gorm.ErrRecordNotFound {
				return fmt.Errorf("读取生产工序失败")
			} else if err != nil {
				return err
			}

			systemEventTrigger := &model.SystemEventTrigger{
				SystemEventID: systemEvent.ID,
				EventNo:       uuid.NewString(),
				CurrentState:  types.SystemEventTriggerStateWaitExecute,
			}
			for _, systemEventParameter := range systemEvent.SystemEventParameters {
				value := systemEventParameter.Value

				value = strings.ReplaceAll(value, "{ProductSerialNo}", productInfo.ProductSerialNo)
				value = strings.ReplaceAll(value, "{ProductionStation}", productionStation.Code)
				value = strings.ReplaceAll(value, "{ProductionProcess.Identifier}", productionProcess.Identifier)
				value = strings.ReplaceAll(value, "{ProductionProcess}", productionProcess.Code)
				value = strings.ReplaceAll(value, "{ProductReworkCause}", reworkCause)

				systemEventTrigger.SystemEventTriggerParameters = append(systemEventTrigger.SystemEventTriggerParameters, &model.SystemEventTriggerParameter{
					DataType:    systemEventParameter.DataType,
					Name:        systemEventParameter.Name,
					Description: systemEventParameter.Description,
					Value:       value,
				})
			}
			if err := tx.Save(systemEventTrigger).Error; err != nil {
				return err
			}
		}
		bomIdMap := map[string]struct{}{}
		if len(req.MaterialRecords) > 0 {
			for _, record := range req.MaterialRecords {
				productOrderBom := &model.ProductOrderBom{}
				if err := tx.Where("`material_no` = ? AND `product_order_id` = ?", record.MaterialNo, productInfo.ProductOrderID).First(productOrderBom).Error; err == gorm.ErrRecordNotFound {
					return fmt.Errorf("未能通过ID找到对应的工单BOM信息")
				} else if err != nil {
					return err
				}

				if _, ok := bomIdMap[productOrderBom.ID]; !ok {
					bomIdMap[productOrderBom.ID] = struct{}{}
				}
				targetOperationModes := []string{types.ProductReworkOperationModeRework, types.ProductReworkOperationModeScrap}
				if record.ProductIssueRecordID != "" {
					if tool.Contains(record.OperationMode, targetOperationModes) && record.NewMaterialTraceNo == "" {
						return fmt.Errorf("新的物料追溯号不允许为空！")
					}

					if tool.Contains(record.OperationMode, targetOperationModes) {
						productIssueRecord := model.ProductIssueRecord{}
						if err := tx.First(productIssueRecord, "`id` = ?", record.ProductIssueRecordID).Error; err == gorm.ErrRecordNotFound {
							return fmt.Errorf("读取发料记录(%s)失败", record.ProductIssueRecordID)
						} else if err != nil {
							return err
						}

						//报废旧发料记录
						productIssueRecord.CurrentState = types.ProductIssueRecordStateScrapped

						//创建新发料记录
						if err := tx.Save(&model.ProductIssueRecord{
							IssuanceProcessID: productIssueRecord.IssuanceProcessID,
							CurrentState:      types.ProductIssueRecordStateBatched,
							MaterialTraceNo:   record.NewMaterialTraceNo,
							ProductInfoID:     productReworkRecord.ProductInfoID,
							ProductOrderBomID: productOrderBom.ID,
							CreateUserID:      req.ProductReworkUserID,
						}).Error; err != nil {
							return err
						}
					}

					//创建物料操作记录
					if err := tx.Save(&model.ProductReworkOperation{
						OperationMode:         record.OperationMode,
						ProductOrderBomID:     productOrderBom.ID,
						ProductReworkRecordID: productReworkRecord.ID,
					}).Error; err != nil {
						return err
					}

					if tool.Contains(record.OperationMode, targetOperationModes) {
						materialInfo := &model.MaterialInfo{}
						if err := tx.First(materialInfo, "`material_no` = ?", record.MaterialNo).Error; err == gorm.ErrRecordNotFound {
							return fmt.Errorf("物料号%s对应的物料信息不存在，请在后台维护！", record.MaterialNo)
						} else if err != nil {
							return err
						}

						if err := tx.Save(&model.MaterialReturnRequestForm{
							FormNo:          uuid.NewString(),
							CreateUserID:    req.ProductReworkUserID,
							HandleMethod:    record.OperationMode,
							MaterialTraceNo: record.OldMaterialTraceNo,
							ReturnSource:    "返工",
							ReturnID:        req.ProductReworkRecordID,
							MaterialInfoID:  materialInfo.ID,
						}).Error; err != nil {
							return err
						}
					}
				}
			}
		}

		lastProductProcessRoute := &model.ProductProcessRoute{}
		if err := tx.Preload("CurrentProcess").Where(&model.ProductProcessRoute{
			ProductInfoID:    productInfo.ID,
			CurrentProcessID: productInfo.ProductionProcessID,
			CurrentState:     types.ProductProcessRouteStateReworking,
		}).First(lastProductProcessRoute).Error; err == gorm.ErrRecordNotFound {
			return fmt.Errorf("读取产品当前工艺路线失败")
		} else if err != nil {
			return err
		}

		//上一步工艺完成
		lastProductProcessRoute.CurrentState = types.ProductProcessRouteStateReworked

		if lastProductProcessRoute.CurrentProcess == nil {
			return fmt.Errorf("读取上一步工艺失败")
		}

		nextRouteID := lastProductProcessRoute.CurrentProcess.ID

		if len(bomIdMap) > 0 {
			bomIds := make([]string, 0, len(bomIdMap))
			for k := range bomIdMap {
				bomIds = append(bomIds, k)
			}
			materialNos := []string{}
			if err := tx.Model(model.ProductOrderBom{}).Where("`id` in ?", bomIds).Select("material_no").Scan(&materialNos).Error; err != nil {
				return err
			}

			materialCategories := []string{}
			if err := tx.Model(model.MaterialInfo{}).Where("`material_no` in ?", materialNos).Select("material_category_id").Scan(&materialCategories).Error; err != nil {
				return err
			}

			var followProcessID string
			if err := tx.Model(model.ProductReworkRoute{}).Where("`material_category_id` in ?", materialCategories).Select("follow_process_id").Order("sort_index").Limit(1).Scan(&followProcessID).Error; err != nil {
				return err
			}
			if followProcessID != "" {
				nextRouteID = followProcessID
			}
		}

		if err := tx.First(nextProcess, "`id` = ?", nextRouteID).Error; err == gorm.ErrRecordNotFound {
			return fmt.Errorf("读取下一步工艺失败")
		} else if err != nil {
			return err
		}

		if err := tx.Create(&model.ProductProcessRoute{
			CurrentState:     types.ProductProcessRouteStateWaitProcess,
			LastProcessID:    &productInfo.ProductionProcessID,
			CurrentProcessID: nextProcess.ID,
			RouteIndex:       nextProcess.SortIndex,
			ProductInfoID:    productInfo.ID,
			WorkIndex:        lastProductProcessRoute.WorkIndex + 1,
		}).Error; err != nil {
			return err
		}
		productReworkRecord.ProductionProcessID = nextProcess.ID
		productInfo.ProductionProcessID = nextProcess.ID

		//从ProductionProcess表中获取当前工艺的产品状态
		currentProductionProcess := &model.ProductionProcess{}
		if err := tx.First(currentProductionProcess, "`id` = ?", productInfo.ProductionProcessID).Error; err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		if currentProductionProcess.ProductState == "" {
			return fmt.Errorf("工序%s的产品状态字段为空，请在后台工艺管理-生产工序管理中维护！", currentProductionProcess.Code)
		}
		productInfo.CurrentState = currentProductionProcess.ProductState

		if err := tx.Save(productReworkRecord).Error; err != nil {
			return err
		}
		if err := tx.Save(productInfo).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"code":    types.ServiceResponseCodeSuccess,
		"message": fmt.Sprintf("返工处理完成，请将产品放行到下个工序: %s", nextProcess.Description),
		"data": map[string]interface{}{
			"id":                      productReworkRecord.ID,
			"productReworkCauseID":    productReworkRecord.ProductReworkCauseID,
			"productReworkSolutionID": productReworkRecord.ProductReworkSolutionID,
			"completeTime":            productReworkRecord.CompleteTime,
			"createTime":              productReworkRecord.CreateTime,
		},
	}, nil
}

// 创建产品测试记录
// func CreateProductTestRecord(req *proto.CreateProductTestRecordRequest) error {
// 	if req.TestStartTime == "" {
// 		return fmt.Errorf("TestStartTime不能为空")
// 	}
// 	if req.TestEndTime == "" {
// 		return fmt.Errorf("TestEndTime不能为空")
// 	}
// 	if req.TestData == "" {
// 		return fmt.Errorf("TestData不能为空")
// 	}
// 	if req.ProductionStation == "" {
// 		return fmt.Errorf("ProductionStation不能为空")
// 	}
// 	if req.ProductSerialNo == "" {
// 		return fmt.Errorf("ProductSerialNo不能为空")
// 	}
// 	if req.TestProject == "" {
// 		return fmt.Errorf("TestProject不能为空")
// 	}
// 	timeNow := time.Now()
// 	timeNowStr := timeNow.Format("2006-01-02 15:04:05")
// 	req.ProductSerialNo = strings.Trim(strings.Trim(req.ProductSerialNo, "\000"), "\r")
// 	_productInfo, err := clients.ProductInfoClient.Get(context.Background(), &proto.GetProductInfoRequest{ProductSerialNo: req.ProductSerialNo})
// 	if err == gorm.ErrRecordNotFound {
// 		return fmt.Errorf("无效的产品信息")
// 	} else if err != nil {
// 		return nil, err
// 	}
// 	productInfo := _productInfo.Data
// 	_productionProcessStep, err := clients.ProductionProcessStepClient.Get(context.Background(), &proto.GetProductionProcessStepRequest{Code: req.TestProject})
// 	if err == gorm.ErrRecordNotFound {
// 		return fmt.Errorf("无效的测试项目")
// 	} else if err != nil {
// 		return nil, err
// 	}
// 	productionProcessStep := _productionProcessStep.Data
// 	if productInfo.ProductionProcessID == "" && productionProcessStep.ProcessControl {
// 		return fmt.Errorf("无法获取产品的当前工序")
// 	}
// 	_ProductionProcess, err := clients.ProductionProcessClient.GetDetail(context.Background(), &proto.GetDetailRequest{Id: productInfo.ProductionProcessID})
// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		return nil, err
// 	}
// 	productionProcess := _ProductionProcess.Data
// 	if productionProcessStep.ProcessControl {
// 		if productionProcess == nil {
// 			return fmt.Errorf("读取产品的当前工序失败")
// 		}
// 		var hasProductionStation bool
// 		for _, v := range productionProcess.ProductionProcessAvailableStations {
// 			if v.ProductionStation.Code == req.ProductionStation {
// 				hasProductionStation = true
// 				break
// 			}
// 		}
// 		if productionProcess.EnableControl && !hasProductionStation {
// 			return fmt.Errorf("非法操作，产品的当前工序不支持在此工位进行")
// 		}
// 	}
// 	_productionStation, err := clients.ProductionStationClient.Get(context.Background(), &proto.GetProductionStationRequest{Code: req.ProductionStation})
// 	if err == gorm.ErrRecordNotFound {
// 		return fmt.Errorf("无效的产线工位")
// 	} else if err != nil {
// 		return nil, err
// 	}
// 	productionStation := _productionStation.Data
// 	testEndTime, err := time.Parse("2006-01-02 15:04:05", req.TestEndTime)
// 	if err != nil {
// 		return nil, err
// 	}
// 	testStartTime, err := time.Parse("2006-01-02 15:04:05", req.TestStartTime)
// 	if err != nil {
// 		return nil, err
// 	}
// 	productTestRecord := &proto.ProductTestRecordInfo{
// 		ProductionProcessStepID: productionProcessStep.Id,
// 		ProductInfoID:           productInfo.Id,
// 		ProductionProcessID:     productionProcess.Id,
// 		ProductionStationID:     productionStation.Id,
// 		TestUserID:              productionStation.CurrentUserID,
// 		TestData:                req.TestData,
// 		TestEndTime:             req.TestEndTime,
// 		TestStartTime:           req.TestStartTime,
// 		Duration:                int32(testEndTime.Sub(testStartTime).Seconds()),
// 		IsQualified:             req.IsQualified,
// 	}
// 	if _, err := clients.ProductTestRecordClient.Add(context.Background(), productTestRecord); err != nil {
// 		return nil, err
// 	}
// 	if productionProcess != nil && productionProcess.Identifier != "" {
// 		//创建系统事件上报测试数据
// 		_systemEvent, err := clients.SystemEventClient.Get(context.Background(), &proto.GetSystemEventRequest{Code: types.SystemEventProductInfoTested, Enable: true})
// 		switch err {
// 		case nil:
// 			systemEvent := _systemEvent.Data
// 			systemEventTrigger := &proto.SystemEventTriggerInfo{
// 				SystemEventID: systemEvent.Id,
// 				CreateTime:    timeNowStr,
// 				EventNo:       uuid.NewString(),
// 				CurrentState:  types.SystemEventTriggerStateWaitExecute,
// 			}
// 			for _, _systemEventParameter := range systemEvent.SystemEventParameters {
// 				value := _systemEventParameter.Value
// 				if _systemEventParameter.Name == "TestData" {
// 					value = productTestRecord.TestData
// 				}
// 				value = strings.ReplaceAll(value, "{ProductSerialNo}", productInfo.ProductSerialNo)
// 				value = strings.ReplaceAll(value, "{ProductionStation}", productionStation.Code)
// 				value = strings.ReplaceAll(value, "{ProductionProcessIdentifier}", productionProcess.Identifier)
// 				value = strings.ReplaceAll(value, "{ProductionProcess}", productionProcess.Code)
// 				systemEventTrigger.SystemEventTriggerParameters = append(systemEventTrigger.SystemEventTriggerParameters, &proto.SystemEventTriggerParameterInfo{
// 					DataType:    _systemEventParameter.DataType,
// 					Name:        _systemEventParameter.Name,
// 					Description: _systemEventParameter.Description,
// 					Value:       value,
// 				})
// 			}
// 			if _, err := clients.SystemEventTriggerClient.Add(context.Background(), systemEventTrigger); err != nil {
// 				return nil, err
// 			}
// 		case gorm.ErrRecordNotFound:
// 		default:
// 			return nil, err
// 		}
// 	}
// 	return &proto.CommonResponse{Code: types.ServiceResponseCodeSuccess, Message: "创建成功"}, nil
// }
