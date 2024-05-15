package logic

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/clients"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/types"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Code = 0, 工艺路线正确
// Code = 1, 校验失败
// Code = 2, 返工产品
// Code = 3，工艺路线错误
// Code = 4, 完工产品
// Code = 5, 读取托盘信息失败
func EnterProductionStation(req *proto.EnterProductionStationRequest) (*proto.EnterProductionStationResponse, error) {
	if req.ProductionStation == "" {
		return nil, fmt.Errorf("ProductionStation不能为空")
	}
	response := &proto.EnterProductionStationResponse{}

	nowStr := time.Now().Format("2006-01-02 15:04:05")
	if req.TrayNo != "" {
		//根据托盘号获取物料托盘
		_materialTray, err := clients.MaterialTrayClient.Get(context.Background(), &proto.GetMaterialTrayRequest{Identifier: req.TrayNo})
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("无效的托盘号")
		} else if err != nil {
			return nil, err
		}

		materialTray := _materialTray.Data
		if materialTray.ProductionLineID == "" {
			return &proto.EnterProductionStationResponse{Code: 5}, fmt.Errorf("托盘未绑定任何产品")
		}

		req.ProductSerialNo = materialTray.ProductInfo.ProductSerialNo
	}
	if req.PackageNo != "" {
		//根据包装箱号获取产品包装记录
		_productPackageRecord, err := clients.ProductPackageRecordClient.Get(context.Background(), &proto.GetProductPackageRecordRequest{PackageNo: req.PackageNo})
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("无效的包装箱号")
		} else if err != nil {
			return nil, err
		}

		productPackageRecord := _productPackageRecord.Data
		if productPackageRecord.ProductInfoID == "" {
			return &proto.EnterProductionStationResponse{Code: 5}, fmt.Errorf("包装箱未绑定任何产品")
		}

		req.ProductSerialNo = productPackageRecord.ProductInfo.ProductSerialNo
	}
	req.ProductSerialNo = strings.Trim(strings.Trim(req.ProductSerialNo, "\000"), "\r")
	if req.ProductSerialNo == "" {
		return nil, fmt.Errorf("ProductSerialNo不能为空")
	}

	//根据工位代号获取产线工站
	_productionStation, err := clients.ProductionStationClient.Get(context.Background(), &proto.GetProductionStationRequest{Code: req.ProductionStation})
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("无效的工位代号")
	} else if err != nil {
		return nil, err
	}

	productionStation := _productionStation.Data
	if productionStation.AccountControl && productionStation.CurrentUserID == "" {
		return nil, fmt.Errorf("工位未登录，无法进站")
	}
	if productionStation.CurrentState == types.ProductionStationStateBreakdown {
		return nil, fmt.Errorf("设备故障中，请尽快联系人员处理并恢复设备故障")
	}
	//根据产品序列号获取产品
	_productInfo, err := clients.ProductInfoClient.Get(context.Background(), &proto.GetProductInfoRequest{ProductSerialNo: req.ProductSerialNo})
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("读取产品信息失败")
	} else if err != nil {
		return nil, err
	}

	productInfo := _productInfo.Data
	switch productInfo.CurrentState {
	case types.ProductStateChecking:
		return &proto.EnterProductionStationResponse{Code: 2}, fmt.Errorf("产品状态错误，此产品状态为检查中")
	case types.ProductStateReworking:
		return &proto.EnterProductionStationResponse{Code: 2}, fmt.Errorf("产品状态错误，此产品状态为返工中")
	case types.ProductStateCompleted:
		return &proto.EnterProductionStationResponse{Code: 4}, fmt.Errorf("产品状态错误，此产品状态为已完工")
	}
	if productInfo.ProductionProcessID == "" {
		return nil, fmt.Errorf("此产品未开工")
	}

	//根据id获取产品订单
	_productOrder, err := clients.ProductOrderClient.GetDetail(context.Background(), &proto.GetDetailRequest{Id: productInfo.ProductOrderID})
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("读取产品工单失败")
	} else if err != nil {
		return nil, err
	}

	productOrder := _productOrder.Data
	//获取产品节拍
	if _, err := clients.ProductRhythmRecordClient.Get(context.Background(), &proto.GetProductRhythmRecordRequest{
		ProductionProcessID: productInfo.ProductionProcessID,
		ProductInfoID:       productInfo.Id,
		ProductionStationID: productionStation.Id,
		HasWorkEndTime:      false,
	}); err == gorm.ErrRecordNotFound {
		//TODO: 重复进站不重复报工，以第一次进站时间为准
		if _, err := clients.ProductRhythmRecordClient.Add(context.Background(), &proto.ProductRhythmRecordInfo{
			WorkUserID:          productionStation.CurrentUserID,
			CreateTime:          nowStr,
			ProductionStationID: productionStation.Id,
			ProductInfoID:       productInfo.Id,
			ProductionProcessID: productInfo.ProductionProcessID,
			StandardWorkTime:    productInfo.ProductOrder.StandardWorkTime,
			WorkStartTime:       nowStr,
		}); err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	targetStates := []string{types.ProductProcessRouteStateWaitProcess, types.ProductProcessRouteStateProcessing}
	//获取产品工艺路线
	_productProcessRoute, err := clients.ProductProcessRouteClient.Get(context.Background(), &proto.GetProductProcessRouteRequest{
		ProductInfoID:    productInfo.Id,
		CurrentProcessID: productInfo.ProductionProcessID,
		CurrentStates:    targetStates,
	})
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("读取产品当前工艺路线错误")
	} else if err != nil {
		return nil, err
	}

	//修改工艺路线状态和执行工位
	productProcessRoute := _productProcessRoute.Data
	productProcessRoute.ProcessUserID = productionStation.CurrentUserID
	productProcessRoute.ProductionStationID = productionStation.Id
	productProcessRoute.CurrentState = types.ProductProcessRouteStateProcessing
	productProcessRoute.LastUpdateTime = nowStr

	//获取生产工艺
	_productionProcess, err := clients.ProductionProcessClient.GetDetail(context.Background(), &proto.GetDetailRequest{Id: productInfo.ProductionProcessID})
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("读取产品当前工艺错误")
	} else if err != nil {
		return nil, err
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
		_stationRoute, err := clients.ProductionProcessClient.Query(context.Background(), &proto.QueryProductionProcessRequest{
			PageSize:            1,
			ProductionLineID:    productOrder.ProductionLineID,
			SortConfig:          "sort_index",
			ProductionStationID: productionStation.Id,
		})
		if err == gorm.ErrRecordNotFound {
			return &proto.EnterProductionStationResponse{Code: 3}, fmt.Errorf("工艺路线错误，且当前工位并不在当前产线的工艺路线之内，请联系管理员处理")
		} else if err != nil {
			return nil, err
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
		}}, fmt.Errorf("工艺路线错误，且此产品的当前工序为%s(%s)，在当前工位的执行工序之%s", productionProcess.Description, productionProcess.Code, towardStr)
	}

	//检查人员资质
	//检查工序是否启用人员管控
	if productionProcess.EnableControl {
		//获取产品型号
		_productModel, err := clients.ProductModelClient.GetDetail(context.Background(), &proto.GetDetailRequest{Id: productInfo.ProductionProcessID})
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("读取产品型号失败")
		} else if err != nil {
			return nil, err
		}

		//获取人员资格
		_personnelQualifications, err := clients.PersonnelQualificationClient.Query(context.Background(), &proto.QueryPersonnelQualificationRequest{
			ProductionProcessID: productionProcess.Id,
			ProductModelID:      _productModel.Data.Id,
		})
		if err != nil {
			return nil, err
		}

		if len(_personnelQualifications.Data) > 0 {
			_personnelQualification, err := clients.PersonnelQualificationClient.Get(context.Background(), &proto.GetPersonnelQualificationRequest{CertifiedUserID: productionStation.CurrentUserID})
			if err == gorm.ErrRecordNotFound {
				return nil, fmt.Errorf("当前作业人员缺少认证资质，无法开工")
			} else if err != nil {
				return nil, err
			}

			if _personnelQualification.Data.ExpirationDate <= nowStr {
				return nil, fmt.Errorf("当前作业人员的认证资质已过期，无法开工")
			}
		}
	}

	//获取作业手册
	_productionProcessSop, err := clients.ProductionProcessSopClient.Get(context.Background(), &proto.GetProductionProcessSopRequest{
		ProductionProcessID: productionProcess.Id,
		ProductModelID:      productOrder.ProductModelID,
	})
	switch err {
	case nil:
		if _productionProcessSop.Data.FileLink != "" {
			sopLink = _productionProcessSop.Data.FileLink
		}
	case gorm.ErrRecordNotFound:
	default:
		return nil, err
	}

	if productionProcess.EnableReport {
		//创建系统事件上报开工
		_systemEvent, err := clients.SystemEventClient.Get(context.Background(), &proto.GetSystemEventRequest{Code: types.SystemEventProductionProcessStarted, Enable: true})
		switch err {
		case nil:
			systemEvent := _systemEvent.Data
			systemEventTrigger := &proto.SystemEventTriggerInfo{
				SystemEventID:  systemEvent.Id,
				CreateTime:     nowStr,
				EventNo:        uuid.NewString(),
				CurrentState:   types.SystemEventTriggerStateWaitExecute,
				LastUpdateTime: nowStr,
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

			if _, err := clients.SystemEventTriggerClient.Add(context.Background(), systemEventTrigger); err != nil {
				return nil, err
			}
		case gorm.ErrRecordNotFound:
		default:
			return nil, err
		}
	}

	if productionStation.AllowReport {
		productionStation.CurrentState = types.ProductionStationStateOccupied
		//创建系统事件上报工位开始处于作业状态
		_systemEvent, err := clients.SystemEventClient.Get(context.Background(), &proto.GetSystemEventRequest{Code: types.SystemEventProductionStationOccupied, Enable: true})
		switch err {
		case nil:
			systemEvent := _systemEvent.Data
			systemEventTrigger := &proto.SystemEventTriggerInfo{
				SystemEventID:  systemEvent.Id,
				CreateTime:     nowStr,
				EventNo:        uuid.NewString(),
				CurrentState:   types.SystemEventTriggerStateWaitExecute,
				LastUpdateTime: nowStr,
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

			if _, err := clients.SystemEventTriggerClient.Add(context.Background(), systemEventTrigger); err != nil {
				return nil, err
			}
		case gorm.ErrRecordNotFound:
		default:
			return nil, err
		}
	}

	response.Data = &proto.EnterProductionStationData{
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
	}

	return response, nil
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
		_materialTray, err := clients.MaterialTrayClient.Get(context.Background(), &proto.GetMaterialTrayRequest{Identifier: req.TrayNo})
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("无效的托盘号")
		} else if err != nil {
			return nil, err
		}

		materialTray := _materialTray.Data
		if materialTray.ProductionLineID == "" {
			return nil, fmt.Errorf("托盘未绑定任何产品")
		}

		req.ProductSerialNo = materialTray.ProductInfo.ProductSerialNo
	}
	if req.PackageNo != "" {
		//根据包装箱号获取产品包装记录
		_productPackageRecord, err := clients.ProductPackageRecordClient.Get(context.Background(), &proto.GetProductPackageRecordRequest{PackageNo: req.PackageNo})
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("无效的包装箱号")
		} else if err != nil {
			return nil, err
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

	_productionStation, err := clients.ProductionStationClient.Get(context.Background(), &proto.GetProductionStationRequest{Code: req.ProductionStation})
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("非法的工站代号")
	} else if err != nil {
		return nil, err
	}
	productionStation := _productionStation.Data

	_productInfo, err := clients.ProductInfoClient.Get(context.Background(), &proto.GetProductInfoRequest{ProductSerialNo: req.ProductSerialNo})
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("读取产品信息失败")
	} else if err != nil {
		return nil, err
	}
	productInfo := _productInfo.Data
	if productInfo.ProductionProcessID == "" {
		return nil, fmt.Errorf("无法读取产品的当前工序")
	}

	//上传节拍
	_productRhythmRecord, err := clients.ProductRhythmRecordClient.Get(context.Background(), &proto.GetProductRhythmRecordRequest{ProductionStationID: productionStation.Id, ProductInfoID: productInfo.Id, HasWorkEndTime: false})
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("读取工位当前节拍数据失败")
	} else if err != nil {
		return nil, err
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
	// productRhythmRecord.OverTime = max(productRhythmRecord.WorkTime-productRhythmRecord.StandardWorkTime, 0)
	productRhythmRecord.WaitTime = req.WaitTime

	//修改工艺记录
	_productProcessRoute, err := clients.ProductProcessRouteClient.Get(context.Background(), &proto.GetProductProcessRouteRequest{ProductInfoID: productInfo.Id, CurrentProcessID: productInfo.ProductionProcessID, CurrentStates: []string{types.ProductProcessRouteStateProcessing}})
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("读取产品当前工艺路线失败")
	} else if err != nil {
		return nil, err
	}
	lastProductProcessRoute := _productProcessRoute.Data

	if req.IsRework {
		productRhythmRecord.IsRework = true

		lastProductProcessRoute.CurrentState = types.ProductProcessRouteStateReworking
		lastProductProcessRoute.LastUpdateTime = timeNowStr

		if _, err := clients.ProductReworkRecordClient.Add(context.Background(), &proto.ProductReworkRecordInfo{
			ProductionStationID: productionStation.Id,
			ProductInfoID:       productInfo.Id,
			ProductionProcessID: productInfo.ProductionProcessID,
			ReworkTime:          timeNowStr,
			ReworkReason:        req.ReworkReason,
		}); err != nil {
			return nil, err
		}

		productInfo.CurrentState = types.ProductStateReworking
		productInfo.LastUpdateTime = timeNowStr
	} else if req.IsFail {
		productRhythmRecord.IsRework = true
		lastProductProcessRoute.Remark = req.ReworkReason
		lastProductProcessRoute.CurrentState = types.ProductProcessRouteStateChecking
		lastProductProcessRoute.LastUpdateTime = timeNowStr
		productInfo.CurrentState = types.ProductStateChecking
		productInfo.LastUpdateTime = timeNowStr
	} else {
		lastProductProcessRoute.CurrentState = types.ProductProcessRouteStateProcessed
		lastProductProcessRoute.LastUpdateTime = timeNowStr

		//切换到下个工艺
		_productProcessRoutes, err := clients.ProductProcessRouteClient.Query(context.Background(), &proto.QueryProductProcessRouteRequest{
			PageSize:      1,
			SortConfig:    "sort_index",
			ProductInfoID: productInfo.Id,
			RouteIndex:    lastProductProcessRoute.RouteIndex,
			CurrentState:  types.ProductProcessRouteStateWaitProcess,
		})
		if err != nil {
			return nil, err
		}
		var nextProductProcessRoute *proto.ProductProcessRouteInfo
		if len(_productProcessRoutes.Data) > 0 {
			nextProductProcessRoute = _productProcessRoutes.Data[0]
		}
		//兼容，部分产线是直接创建产品工艺路线，部分是根据工单工艺动态创建
		if nextProductProcessRoute == nil {
			_productOrderProcesses, err := clients.ProductOrderProcessClient.Query(context.Background(), &proto.QueryProductOrderProcessRequest{
				ProductOrderID: productInfo.ProductOrderID,
				Enable:         true,
				SortIndex:      lastProductProcessRoute.RouteIndex,
				SortConfig:     "sort_index",
				PageSize:       1,
			})
			if err != nil {
				return nil, err
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
					LastUpdateTime:   timeNowStr,
					ProductInfoID:    productInfo.Id,
				}
				if _, err := clients.ProductProcessRouteClient.Add(context.Background(), nextProductProcessRoute); err != nil {
					return nil, err
				}
			}
		}

		if nextProductProcessRoute != nil {
			nextProductProcessRoute.WorkIndex = lastProductProcessRoute.WorkIndex + 1
			nextProductProcessRoute.LastUpdateTime = timeNowStr
			if nextProductProcessRoute.CurrentProcess != nil {
				if nextProductProcessRoute.CurrentProcess.ProductState != "" {
					//设定当前工序的产品信息状态
					productInfo.CurrentState = nextProductProcessRoute.CurrentProcess.ProductState
				}
			}
			productInfo.ProductionProcessID = nextProductProcessRoute.CurrentProcessID
			productInfo.LastUpdateTime = timeNowStr

			//计算预计下线时间
			_productOrderProcesses, err := clients.ProductOrderProcessClient.Query(context.Background(), &proto.QueryProductOrderProcessRequest{
				ProductOrderID: productInfo.ProductOrderID,
				Enable:         true,
				SortIndex:      nextProductProcessRoute.RouteIndex,
			})
			if err != nil {
				return nil, err
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
			productInfo.LastUpdateTime = timeNowStr
			productInfo.ProductionProcessID = ""

			_productOrder, err := clients.ProductOrderClient.GetDetail(context.Background(), &proto.GetDetailRequest{Id: productInfo.ProductOrderID})
			if err != nil {
				return nil, err
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
				productOrder.LastUpdateTime = timeNowStr
				productOrder.ActualFinishTime = timeNowStr

				_systemEvent, err := clients.SystemEventClient.Get(context.Background(), &proto.GetSystemEventRequest{Code: types.SystemEventProductOrderFinished})
				switch err {
				case nil:
					systemEvent := _systemEvent.Data
					systemEventTrigger := &proto.SystemEventTriggerInfo{
						SystemEventID:  systemEvent.Id,
						CreateTime:     timeNowStr,
						EventNo:        uuid.NewString(),
						CurrentState:   types.SystemEventTriggerStateWaitExecute,
						LastUpdateTime: timeNowStr,
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

					if _, err := clients.SystemEventTriggerClient.Add(context.Background(), systemEventTrigger); err != nil {
						return nil, err
					}
				case gorm.ErrRecordNotFound:
				default:
					return nil, err
				}
			}
		}

		//创建产量记录
		_ProductionStationOutput, err := clients.ProductionStationOutputClient.Get(context.Background(), &proto.GetProductionStationOutputRequest{
			ProductionStationID: productionStation.Id,
			ProductInfoID:       productInfo.Id,
		})
		productionStationOutput := _ProductionStationOutput.Data
		switch err {
		case nil:
			productionStationOutput.OutputTime = timeNowStr
			productionStationOutput.LoginUserID = productionStation.CurrentUserID
		case gorm.ErrRecordNotFound:
			productionStationOutput = &proto.ProductionStationOutputInfo{
				OutputTime:          timeNowStr,
				LoginUserID:         productionStation.CurrentUserID,
				ProductionProcessID: productRhythmRecord.ProductionProcessID,
				ProductInfoID:       productRhythmRecord.ProductInfoID,
				ProductionStationID: productRhythmRecord.ProductionStationID,
			}
			if _, err := clients.ProductionStationOutputClient.Add(context.Background(), productionStationOutput); err != nil {
				return nil, err
			}
		default:
			return nil, err
		}

		// 判断是否要解绑托盘
		if req.UnbindTray {
			_materialTray, err := clients.MaterialTrayClient.Get(context.Background(), &proto.GetMaterialTrayRequest{ProductInfoID: productInfo.Id})
			materialTray := _materialTray.Data
			switch err {
			case nil:
				materialTray.ProductInfoID = "null"
				materialTray.CurrentState = types.MaterialTrayStateWaitBind
			case gorm.ErrRecordNotFound:
			default:
				return nil, err
			}
		}

		_productionProcess, err := clients.ProductionProcessClient.GetDetail(context.Background(), &proto.GetDetailRequest{Id: productRhythmRecord.ProductionProcessID})
		switch err {
		case nil:
			productionProcess := _productionProcess.Data
			if productionProcess.EnableReport {
				_systemEvent, err := clients.SystemEventClient.Get(context.Background(), &proto.GetSystemEventRequest{Code: types.SystemEventProductionProcessFinished, Enable: true})
				switch err {
				case nil:
					systemEvent := _systemEvent.Data
					systemEventTrigger := &proto.SystemEventTriggerInfo{
						SystemEventID:  systemEvent.Id,
						CreateTime:     timeNowStr,
						EventNo:        uuid.NewString(),
						CurrentState:   types.SystemEventTriggerStateWaitExecute,
						LastUpdateTime: timeNowStr,
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

					if _, err := clients.SystemEventTriggerClient.Add(context.Background(), systemEventTrigger); err != nil {
						return nil, err
					}
				case gorm.ErrRecordNotFound:
				default:
					return nil, err
				}
			}
		case gorm.ErrRecordNotFound:
		default:
			return nil, err
		}

		if productionStation.AllowReport {
			productionStation.CurrentState = types.ProductionStationStateStandby
			// 创建系统事件上报工位完成作业
			_systemEvent, err := clients.SystemEventClient.Get(context.Background(), &proto.GetSystemEventRequest{Code: types.SystemEventProductionStationReleased, Enable: true})
			switch err {
			case nil:
				systemEvent := _systemEvent.Data
				systemEventTrigger := &proto.SystemEventTriggerInfo{
					SystemEventID:  systemEvent.Id,
					CreateTime:     timeNowStr,
					EventNo:        uuid.NewString(),
					CurrentState:   types.SystemEventTriggerStateWaitExecute,
					LastUpdateTime: timeNowStr,
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

				if _, err := clients.SystemEventTriggerClient.Add(context.Background(), systemEventTrigger); err != nil {
					return nil, err
				}
			case gorm.ErrRecordNotFound:
			default:
				return nil, err
			}
		}
	}

	return &proto.CommonResponse{Code: 0}, nil
}

// 创建产品测试记录
func CreateProductTestRecord(req *proto.CreateProductTestRecordRequest) (*proto.CommonResponse, error) {
	if req.TestStartTime == "" {
		return CommonFailureResponse("TestStartTime不能为空")
	}
	if req.TestEndTime == "" {
		return CommonFailureResponse("TestEndTime不能为空")
	}
	if req.TestData == "" {
		return CommonFailureResponse("TestData不能为空")
	}
	if req.ProductionStation == "" {
		return CommonFailureResponse("ProductionStation不能为空")
	}
	if req.ProductSerialNo == "" {
		return CommonFailureResponse("ProductSerialNo不能为空")
	}
	if req.TestProject == "" {
		return CommonFailureResponse("TestProject不能为空")
	}
	timeNow := time.Now()
	timeNowStr := timeNow.Format("2006-01-02 15:04:05")
	req.ProductSerialNo = strings.Trim(strings.Trim(req.ProductSerialNo, "\000"), "\r")
	_productInfo, err := clients.ProductInfoClient.Get(context.Background(), &proto.GetProductInfoRequest{ProductSerialNo: req.ProductSerialNo})
	if err == gorm.ErrRecordNotFound {
		return CommonFailureResponse("无效的产品信息")
	} else if err != nil {
		return nil, err
	}
	productInfo := _productInfo.Data

	_productionProcessStep, err := clients.ProductionProcessStepClient.Get(context.Background(), &proto.GetProductionProcessStepRequest{Code: req.TestProject})
	if err == gorm.ErrRecordNotFound {
		return CommonFailureResponse("无效的测试项目")
	} else if err != nil {
		return nil, err
	}
	productionProcessStep := _productionProcessStep.Data

	if productInfo.ProductionProcessID == "" && productionProcessStep.ProcessControl {
		return CommonFailureResponse("无法获取产品的当前工序")
	}
	_ProductionProcess, err := clients.ProductionProcessClient.GetDetail(context.Background(), &proto.GetDetailRequest{Id: productInfo.ProductionProcessID})
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	productionProcess := _ProductionProcess.Data
	if productionProcessStep.ProcessControl {
		if productionProcess == nil {
			return CommonFailureResponse("读取产品的当前工序失败")
		}

		var hasProductionStation bool
		for _, v := range productionProcess.ProductionProcessAvailableStations {
			if v.ProductionStation.Code == req.ProductionStation {
				hasProductionStation = true
				break
			}
		}
		if productionProcess.EnableControl && !hasProductionStation {
			return CommonFailureResponse("非法操作，产品的当前工序不支持在此工位进行")
		}
	}

	_productionStation, err := clients.ProductionStationClient.Get(context.Background(), &proto.GetProductionStationRequest{Code: req.ProductionStation})
	if err == gorm.ErrRecordNotFound {
		return CommonFailureResponse("无效的产线工位")
	} else if err != nil {
		return nil, err
	}
	productionStation := _productionStation.Data

	testEndTime, err := time.Parse("2006-01-02 15:04:05", req.TestEndTime)
	if err != nil {
		return nil, err
	}
	testStartTime, err := time.Parse("2006-01-02 15:04:05", req.TestStartTime)
	if err != nil {
		return nil, err
	}

	productTestRecord := &proto.ProductTestRecordInfo{
		ProductionProcessStepID: productionProcessStep.Id,
		ProductInfoID:           productInfo.Id,
		ProductionProcessID:     productionProcess.Id,
		ProductionStationID:     productionStation.Id,
		TestUserID:              productionStation.CurrentUserID,
		TestData:                req.TestData,
		TestEndTime:             req.TestEndTime,
		TestStartTime:           req.TestStartTime,
		Duration:                int32(testEndTime.Sub(testStartTime).Seconds()),
		IsQualified:             req.IsQualified,
	}
	if _, err := clients.ProductTestRecordClient.Add(context.Background(), productTestRecord); err != nil {
		return nil, err
	}

	if productionProcess != nil && productionProcess.Identifier != "" {
		//创建系统事件上报测试数据
		_systemEvent, err := clients.SystemEventClient.Get(context.Background(), &proto.GetSystemEventRequest{Code: types.SystemEventProductInfoTested, Enable: true})
		switch err {
		case nil:
			systemEvent := _systemEvent.Data
			systemEventTrigger := &proto.SystemEventTriggerInfo{
				SystemEventID:  systemEvent.Id,
				CreateTime:     timeNowStr,
				EventNo:        uuid.NewString(),
				CurrentState:   types.SystemEventTriggerStateWaitExecute,
				LastUpdateTime: timeNowStr,
			}

			for _, _systemEventParameter := range systemEvent.SystemEventParameters {
				value := _systemEventParameter.Value
				if _systemEventParameter.Name == "TestData" {
					value = productTestRecord.TestData
				}
				value = strings.ReplaceAll(value, "{ProductSerialNo}", productInfo.ProductSerialNo)
				value = strings.ReplaceAll(value, "{ProductionStation}", productionStation.Code)
				value = strings.ReplaceAll(value, "{ProductionProcessIdentifier}", productionProcess.Identifier)
				value = strings.ReplaceAll(value, "{ProductionProcess}", productionProcess.Code)

				systemEventTrigger.SystemEventTriggerParameters = append(systemEventTrigger.SystemEventTriggerParameters, &proto.SystemEventTriggerParameterInfo{
					DataType:    _systemEventParameter.DataType,
					Name:        _systemEventParameter.Name,
					Description: _systemEventParameter.Description,
					Value:       value,
				})
			}

			if _, err := clients.SystemEventTriggerClient.Add(context.Background(), systemEventTrigger); err != nil {
				return nil, err
			}
		case gorm.ErrRecordNotFound:
		default:
			return nil, err
		}
	}

	return &proto.CommonResponse{Code: types.ServiceResponseCodeSuccess, Message: "创建成功"}, nil
}

// 设置失败后续处理
func CheckProductProcessRouteFailure(req *proto.CheckProductProcessRouteFailureRequest) (*proto.CommonResponse, error) {
	if req.ProductionStation == "" {
		return CommonFailureResponse("ProductionStation不能为空")
	}
	if req.ProductSerialNo == "" {
		return CommonFailureResponse("ProductSerialNo不能为空")
	}

	timeNow := time.Now()
	timeNowStr := timeNow.Format("2006-01-02 15:04:05")
	_productionStation, err := clients.ProductionStationClient.Get(context.Background(), &proto.GetProductionStationRequest{Code: req.ProductionStation})
	if err == gorm.ErrRecordNotFound {
		return CommonFailureResponse("无效的工站代号")
	} else if err != nil {
		return nil, err
	}
	productionStation := _productionStation.Data

	_productInfo, err := clients.ProductInfoClient.Get(context.Background(), &proto.GetProductInfoRequest{ProductSerialNo: req.ProductSerialNo})
	if err == gorm.ErrRecordNotFound {
		return CommonFailureResponse("无效的产品序列号")
	} else if err != nil {
		return nil, err
	}
	productInfo := _productInfo.Data

	if productInfo.ProductionProcessID == "" {
		return CommonFailureResponse("数据错误，此产品的当前工序为空")
	}

	_productProcessRoute, err := clients.ProductProcessRouteClient.Get(context.Background(), &proto.GetProductProcessRouteRequest{
		CurrentProcessID: productInfo.ProductionProcessID,
		ProductInfoID:    productInfo.Id,
		CurrentStates:    []string{types.ProductProcessRouteStateChecking},
	})
	if err == gorm.ErrRecordNotFound {
		return CommonFailureResponse("状态错误，此产品的当前工艺状态不是" + types.ProductProcessRouteStateChecking)
	} else if err != nil {
		return nil, err
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
		lastProductProcessRoute.LastUpdateTime = timeNowStr

		productInfo.CurrentState = lastProductProcessRoute.CurrentProcess.ProductState
		if productInfo.CurrentState == "" {
			productInfo.CurrentState = types.ProductStateTesting
		}
		productInfo.LastUpdateTime = timeNowStr
	case types.ProductionProcessHandleMethodRework:
		lastProductProcessRoute.CurrentState = types.ProductProcessRouteStateReworking
		lastProductProcessRoute.LastUpdateTime = timeNowStr

		productReworkRecord := &proto.ProductReworkRecordInfo{
			ProductionStationID: productionStation.Id,
			ProductInfoID:       productInfo.Id,
			ProductionProcessID: productInfo.ProductionProcessID,
			ReworkTime:          timeNowStr,
			ReworkReason:        lastProductProcessRoute.Remark,
		}
		if _, err := clients.ProductReworkRecordClient.Add(context.Background(), productReworkRecord); err != nil {
			return nil, err
		}

		productInfo.CurrentState = types.ProductStateReworking
		productInfo.LastUpdateTime = timeNowStr
	case types.ProductionProcessHandleMethodIgnore:
		lastProductProcessRoute.CurrentState = types.ProductProcessRouteStateProcessed
		lastProductProcessRoute.LastUpdateTime = timeNowStr

		//切换到下个工艺
		_productProcessRoutes, err := clients.ProductProcessRouteClient.Query(context.Background(), &proto.QueryProductProcessRouteRequest{
			PageSize:      1,
			SortConfig:    "route_index",
			ProductInfoID: productInfo.Id,
			RouteIndex:    lastProductProcessRoute.RouteIndex,
			CurrentState:  types.ProductProcessRouteStateWaitProcess,
		})
		if err != nil {
			return nil, err
		}
		var nextProductProcessRoute *proto.ProductProcessRouteInfo
		if len(_productProcessRoutes.Data) > 0 {
			nextProductProcessRoute = _productProcessRoutes.Data[0]
		}
		if nextProductProcessRoute == nil {
			_productOrderProcesses, err := clients.ProductOrderProcessClient.Query(context.Background(), &proto.QueryProductOrderProcessRequest{
				PageSize:       1,
				SortConfig:     "route_index",
				ProductOrderID: productInfo.ProductOrderID,
				Enable:         true,
				SortIndex:      lastProductProcessRoute.RouteIndex,
			})
			if err != nil {
				return nil, err
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
					LastUpdateTime:   timeNowStr,
					ProductInfoID:    productInfo.Id,
				}
				if _, err := clients.ProductProcessRouteClient.Add(context.Background(), nextProductProcessRoute); err != nil {
					return nil, err
				}
			}

			if nextProductProcessRoute != nil {
				nextProductProcessRoute.WorkIndex = lastProductProcessRoute.WorkIndex + 1
				nextProductProcessRoute.CurrentState = types.ProductProcessRouteStateProcessing
				nextProductProcessRoute.LastUpdateTime = timeNowStr

				if nextProductProcessRoute.CurrentProcess.ProductState != "" {
					//设定当前工序的产品信息状态
					productInfo.CurrentState = nextProductProcessRoute.CurrentProcess.ProductState
				}
				productInfo.ProductionProcessID = nextProductProcessRoute.CurrentProcessID
				productInfo.LastUpdateTime = timeNowStr
				//TODO: 计算预计下线时间
				_productProcessRoutes, err := clients.ProductProcessRouteClient.Query(context.Background(), &proto.QueryProductProcessRouteRequest{
					ProductInfoID: productInfo.Id,
					CurrentState:  types.ProductProcessRouteStateWaitProcess,
				})
				if err != nil {
					return nil, err
				}
				remainingRoutes := int32(_productProcessRoutes.Total)
				productInfo.RemainingRoutes = remainingRoutes
				if remainingRoutes > 0 {
					productInfo.EstimateTime = timeNow.Add(time.Duration(remainingRoutes*productInfo.ProductOrder.StandardWorkTime) * time.Second).Format("2006-01-02 15:04:05")
				}
			} else {
				productInfo.CurrentState = types.ProductStateCompleted
				productInfo.FinishedTime = timeNowStr
				productInfo.LastUpdateTime = timeNowStr
				productInfo.ProductionProcessID = ""
			}
		}
	default:
		return CommonFailureResponse("无效的处理方式")
	}

	return &proto.CommonResponse{Code: types.ServiceResponseCodeSuccess, Message: "处理完成"}, nil
}

func CommonFailureResponse(message string) (*proto.CommonResponse, error) {
	return &proto.CommonResponse{Code: types.ServiceResponseCodeFailure, Message: message}, nil
}
