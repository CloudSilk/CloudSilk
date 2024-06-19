package logic

import (
	"database/sql"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/tool"
	"github.com/CloudSilk/CloudSilk/pkg/types"
	"github.com/CloudSilk/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// 设定产品信息状态为上线装配
func OnlineProductInfo(req *proto.OnlineProductInfoRequest) (code int32, err error) {
	if req.ProductionLine == "" {
		return 40000, fmt.Errorf("ProductionLine不能为空")
	}
	if req.ProductSerialNo == "" {
		return 40000, fmt.Errorf("ProductSerialNo不能为空")
	}

	if err := model.DB.DB().Transaction(func(tx *gorm.DB) error {
		productionLine := &model.ProductionLine{}
		if err := tx.First(productionLine, "`code` = ?", req.ProductionLine).Error; err == gorm.ErrRecordNotFound {
			code = 10001
			return fmt.Errorf("无效的产线代号")
		} else if err != nil {
			code = 50000
			return err
		}

		productInfo := &model.ProductInfo{}
		if err := tx.Preload("ProductOrder").First(productInfo, "`product_serial_no` = ?", req.ProductSerialNo).Error; err == gorm.ErrRecordNotFound {
			code = 10002
			return fmt.Errorf("读取产品信息失败，请联系管理员处理")
		} else if err != nil {
			code = 50000
			return err
		}

		if productInfo.CurrentState != types.ProductStateReleased {
			code = 10002
			return fmt.Errorf("产品当前状态为%s，无法上线", productInfo.CurrentState)
		}

		productOrder := productInfo.ProductOrder
		if productOrder == nil {
			code = 10003
			return fmt.Errorf("读取生产工单失败，请联系管理员处理")
		}

		if *productOrder.ProductionLineID != productionLine.ID {
			code = 10003
			return fmt.Errorf("此生产工单发放产线与上线产线不匹配")
		}
		if req.ProductOrderNo != "" && productOrder.ProductOrderNo != req.ProductOrderNo {
			code = 10003
			return fmt.Errorf("此产品的隶属工单与当前工单不匹配")
		}

		//TODO: 兼容，部分产线是直接创建产品工艺路线，部分是根据工单工艺动态创建
		productProcessRoute := &model.ProductProcessRoute{}
		if err := tx.Preload("CurrentProcess").Where(model.ProductProcessRoute{ProductInfoID: productInfo.ID, CurrentState: types.ProductProcessRouteStateWaitProcess}).Order("route_index").First(productProcessRoute).Error; err != nil && err != gorm.ErrRecordNotFound {
			code = 50000
			return err
		}

		if productProcessRoute.ID == "" {
			productOrderProcess := &model.ProductOrderProcess{}
			if err := tx.Preload("ProductionProcess").Where(model.ProductOrderProcess{ProductOrderID: productInfo.ProductOrderID, Enable: true}).Order("sort_index").First(productOrderProcess).Error; err == gorm.ErrRecordNotFound {
				code = 10004
				return fmt.Errorf("上线失败，此工单缺少工艺路线")
			} else if err != nil {
				code = 50000
				return err
			}

			productProcessRoute = &model.ProductProcessRoute{
				LastProcessID:    nil,
				CurrentProcessID: productOrderProcess.ProductionProcessID,
				CurrentProcess:   productOrderProcess.ProductionProcess,
				CurrentState:     types.ProductProcessRouteStateWaitProcess,
				RouteIndex:       productOrderProcess.SortIndex,
				ProductInfoID:    productInfo.ID,
			}

			if err := tx.Create(productProcessRoute).Error; err != nil {
				code = 50000
				return err
			}
		}

		productProcessRoute.WorkIndex = 1

		//更新产品信息
		var remainingRoutes int64
		if err := tx.Model(model.ProductOrderProcess{}).Where("`product_order_id` = ? AND `enable` = ? AND `sort_index` > ?", productInfo.ProductOrderID, true, productProcessRoute.WorkIndex).Count(&remainingRoutes).Error; err != nil {
			code = 50000
			return err
		}

		nowTime := time.Now()
		productInfo.ProductionProcessID = &productProcessRoute.CurrentProcessID
		productInfo.RemainingRoutes = int32(remainingRoutes)
		productInfo.EstimateTime = tool.Time2NullTime(nowTime.Add(time.Duration(int32(remainingRoutes)*productInfo.ProductOrder.StandardWorkTime) * time.Second))
		productInfo.StartedTime = tool.Time2NullTime(nowTime)
		if productProcessRoute.CurrentProcess != nil && productProcessRoute.CurrentProcess.ProductState != "" {
			productInfo.CurrentState = productProcessRoute.CurrentProcess.ProductState
		} else {
			productInfo.CurrentState = types.ProductStateAssembling
		}

		//更新工单信息
		if !productOrder.ActualStartTime.Valid {
			productOrder.ActualStartTime = tool.Time2NullTime(nowTime)
			productOrder.StartedQTY = 0
			productOrder.CurrentState = types.ProductOrderStateProducting

			systemEvent := &model.SystemEvent{}
			if err := tx.Preload("SystemEventParameters").First(systemEvent, "`code` = ?", types.SystemEventProductOrderStarted).Error; err != nil && err != gorm.ErrRecordNotFound {
				code = 50000
				return err
			}

			if systemEvent.ID != "" {
				systemEventTriggerParameter := make([]*model.SystemEventTriggerParameter, len(systemEvent.SystemEventParameters))
				for i, systemEventParameter := range systemEvent.SystemEventParameters {
					value := strings.ReplaceAll(systemEventParameter.Value, "{ProductOrderNo}", productOrder.ProductOrderNo)

					systemEventTriggerParameter[i] = &model.SystemEventTriggerParameter{
						DataType:    systemEventParameter.DataType,
						Name:        systemEventParameter.Name,
						Description: systemEventParameter.Description,
						Value:       value,
					}
				}

				if err := tx.Create(&model.SystemEventTrigger{
					SystemEventID:                systemEvent.ID,
					EventNo:                      uuid.NewString(),
					CurrentState:                 types.SystemEventTriggerStateWaitExecute,
					SystemEventTriggerParameters: systemEventTriggerParameter,
				}).Error; err != nil {
					code = 50000
					return err
				}
			}
			productOrder.StartedQTY += 1

			//绑定载具
			if req.TrayNo != "" {
				materialTray := &model.MaterialTray{}
				if err := tx.First(materialTray, "`identifier` = ?", req.TrayNo).Error; err == gorm.ErrRecordNotFound {
					code = 10005
					return fmt.Errorf("无效的物料载具识别码")
				} else if err != nil {
					code = 50000
					return err
				}

				if !materialTray.Enable {
					code = 10005
					return fmt.Errorf("载具已禁用")
				}
				if materialTray.TrayType != types.MaterialTrayTypeMaterialTray {
					code = 10005
					return fmt.Errorf("非法操作，只允许使用物料载具上线")
				}
				if materialTray.ProductionLineID != productionLine.ID {
					code = 10005
					return fmt.Errorf("非法操作，此载具不属于当前产线")
				}
				if materialTray.ProductInfoID != nil && *materialTray.ProductInfoID != productInfo.ID {
					code = 10005
					return fmt.Errorf("非法操作，此载具已绑定其他产品")
				}

				if err := tx.Create(&model.MaterialTrayBindingRecord{
					MaterialTrayID: materialTray.ID,
					ProductInfoID:  productInfo.ID,
					CurrentState:   types.MaterialTrayBindingRecordStateEffected,
				}).Error; err != nil {
					code = 50000
					return err
				}
			}

			//TODO: 触发事件
			systemEvent2 := &model.SystemEvent{}
			if err := tx.Preload("SystemEventParameters").First(systemEvent2, "`code` = ?", types.SystemEventProductInfoOnlined).Error; err != nil && err != gorm.ErrRecordNotFound {
				code = 50000
				return err
			}

			if systemEvent2.ID != "" {
				productOrderAttributes := []*model.ProductOrderAttribute{}
				if err := tx.Find(&productOrderAttributes, "`product_order_id` = ?", productInfo.ProductOrderID).Error; err != nil {
					code = 50000
					return err
				}

				systemEventTriggerParameter := make([]*model.SystemEventTriggerParameter, len(systemEvent2.SystemEventParameters))
				for i, systemEventParameter := range systemEvent2.SystemEventParameters {
					value := systemEventParameter.Value
					value = strings.ReplaceAll(value, "{ProductSerialNo}", productInfo.ProductSerialNo)
					value = strings.ReplaceAll(value, "{ProductOrderNo}", productOrder.ProductOrderNo)
					value = strings.ReplaceAll(value, "{SalesOrderNo}", productOrder.SalesOrderNo)
					value = strings.ReplaceAll(value, "{ItemNo}", productOrder.ItemNo)
					for _, productOrderAttribute := range productOrderAttributes {
						value = strings.ReplaceAll(value, "{"+productOrderAttribute.ProductAttribute.Code+"}", productOrderAttribute.Value)
					}

					systemEventTriggerParameter[i] = &model.SystemEventTriggerParameter{
						DataType:    systemEventParameter.DataType,
						Name:        systemEventParameter.Name,
						Description: systemEventParameter.Description,
						Value:       value,
					}
				}

				if err := tx.Create(&model.SystemEventTrigger{
					SystemEventID:                systemEvent2.ID,
					EventNo:                      uuid.NewString(),
					CurrentState:                 types.SystemEventTriggerStateWaitExecute,
					SystemEventTriggerParameters: systemEventTriggerParameter,
				}).Error; err != nil {
					code = 50000
					return err
				}
			}
		}

		if err := tx.Save(productInfo).Error; err != nil {
			code = 50000
			return err
		}
		if err := tx.Save(productOrder).Error; err != nil {
			code = 50000
			return err
		}

		return nil
	}); err != nil {
		return code, err
	}

	return 20000, nil
}

// Code = 0, 工艺路线正确
// Code = 1, 校验失败
// Code = 2, 返工产品
// Code = 3，工艺路线错误
// Code = 4, 完工产品
// Code = 5, 读取载具信息失败
// 请求入站
func EnterProductionStation(req *proto.EnterProductionStationRequest) (data *proto.EnterProductionStationData, code int32, err error) {
	if req.ProductionStation == "" {
		return nil, 5, fmt.Errorf("ProductionStation不能为空")
	}

	productInfo := &model.ProductInfo{}
	productionProcess := &model.ProductionProcess{}
	var sopLink string
	var toward int32
	var isAccept bool
	if err := model.DB.DB().Transaction(func(tx *gorm.DB) error {
		nowTime := time.Now()
		if req.TrayNo != "" {
			//根据载具号获取物料载具
			materialTray := &model.MaterialTray{}
			if err := tx.Preload("ProductInfo").First(materialTray, "`identifier` = ?", req.TrayNo).Error; err == gorm.ErrRecordNotFound {
				code = 5
				return fmt.Errorf("无效的物料载具识别码")
			} else if err != nil {
				code = 1
				return err
			}

			if materialTray.ProductInfo == nil {
				code = 5
				return fmt.Errorf("载具未绑定任何产品")
			}

			req.ProductSerialNo = materialTray.ProductInfo.ProductSerialNo
		}
		if req.PackageNo != "" {
			//根据包装箱号获取产品包装记录
			productPackageRecord := &model.ProductPackageRecord{}
			if err := tx.Preload("ProductInfo").First(productPackageRecord, "`package_no` = ?", req.PackageNo).Error; err == gorm.ErrRecordNotFound {
				code = 5
				return fmt.Errorf("无效的包装箱号")
			} else if err != nil {
				code = 1
				return err
			}

			if productPackageRecord.ProductInfo == nil {
				code = 5
				return fmt.Errorf("包装箱未绑定任何产品")
			}

			req.ProductSerialNo = productPackageRecord.ProductInfo.ProductSerialNo
		}
		req.ProductSerialNo = strings.Trim(strings.Trim(req.ProductSerialNo, "\000"), "\r")
		if req.ProductSerialNo == "" {
			code = 5
			return fmt.Errorf("ProductSerialNo不能为空")
		}

		//根据工位代号获取产线工站
		productionStation := &model.ProductionStation{}
		if err := tx.Preload("ProductionLine").First(productionStation, "`code` = ?", req.ProductionStation).Error; err == gorm.ErrRecordNotFound {
			code = 1
			return fmt.Errorf("无效的工位代号")
		} else if err != nil {
			code = 1
			return err
		}

		if productionStation.AccountControl && productionStation.CurrentUserID == nil {
			code = 1
			return fmt.Errorf("工位未登录，无法进站")
		}
		if productionStation.CurrentState == types.ProductionStationStateBreakdown {
			code = 1
			return fmt.Errorf("设备故障中，请尽快联系人员处理并恢复设备故障")
		}

		//根据产品序列号获取产品
		if err := tx.Preload("ProductOrder").
			Preload("ProductOrder.ProductModel").
			Preload("ProductOrder.ProductModel.ProductCategory").
			First(productInfo, "`product_serial_no` = ?", req.ProductSerialNo).Error; err == gorm.ErrRecordNotFound {
			code = 1
			return fmt.Errorf("读取产品信息失败")
		} else if err != nil {
			code = 1
			return err
		}

		switch productInfo.CurrentState {
		case types.ProductStateChecking:
			code = 2
			return fmt.Errorf("产品状态错误，此产品状态为检查中")
		case types.ProductStateReworking:
			code = 2
			return fmt.Errorf("产品状态错误，此产品状态为返工中")
		case types.ProductStateCompleted:
			code = 4
			return fmt.Errorf("产品状态错误，此产品状态为已完工")
		}
		if productInfo.ProductionProcessID == nil {
			code = 1
			return fmt.Errorf("此产品未开工")
		}

		//根据id获取产品订单
		productOrder := &model.ProductOrder{}
		if err := tx.First(productOrder, "`id` = ?", productInfo.ProductOrderID).Error; err == gorm.ErrRecordNotFound {
			code = 1
			return fmt.Errorf("读取产品工单失败")
		} else if err != nil {
			code = 1
			return err
		}

		//获取产品节拍
		productRhythmRecord := &model.ProductRhythmRecord{}
		if err := tx.Where("`production_process_id` = ? AND `product_info_id` = ? AND `production_station_id` = ? AND work_end_time IS NULL", productInfo.ProductionProcessID, productionStation.ID, productionStation.ID).First(productRhythmRecord).Error; err != nil && err != gorm.ErrRecordNotFound {
			code = 1
			return err
		}

		if productRhythmRecord.ID == "" {
			//重复进站不重复报工，以第一次进站时间为准
			if err := tx.Create(&model.ProductRhythmRecord{
				WorkUserID:          *productionStation.CurrentUserID,
				ProductionStationID: productionStation.ID,
				ProductInfoID:       productInfo.ID,
				ProductionProcessID: *productInfo.ProductionProcessID,
				StandardWorkTime:    productInfo.ProductOrder.StandardWorkTime,
				WorkStartTime:       tool.Time2NullTime(nowTime),
			}).Error; err != nil {
				code = 1
				return err
			}
		}

		//获取产品工艺路线
		targetStates := []string{types.ProductProcessRouteStateWaitProcess, types.ProductProcessRouteStateProcessing}
		productProcessRoute := &model.ProductProcessRoute{}
		if err := tx.Where("`product_info_id` = ? AND `current_process_id` = ? AND `current_state` in ?", productInfo.ID, productInfo.ProductionProcessID, targetStates).First(productProcessRoute).Error; err == gorm.ErrRecordNotFound {
			code = 1
			return fmt.Errorf("读取产品当前工艺路线错误")
		} else if err != nil {
			code = 1
			return err
		}

		//修改工艺路线状态和执行工位
		productProcessRoute.ProcessUserID = productionStation.CurrentUserID
		productProcessRoute.ProductionStationID = &productionStation.ID
		productProcessRoute.CurrentState = types.ProductProcessRouteStateProcessing
		if err := tx.Save(productProcessRoute).Error; err != nil {
			code = 1
			return err
		}

		//获取生产工艺
		if err := tx.Preload("ProductionProcessAvailableStations").First(productionProcess, "`id` = ?", productInfo.ProductionProcessID).Error; err == gorm.ErrRecordNotFound {
			code = 3
			return fmt.Errorf("读取产品当前工艺错误")
		} else if err != nil {
			code = 1
			return err
		}

		var accept bool
		for _, _productionStation := range productionProcess.ProductionProcessAvailableStations {
			if _productionStation.ProductionStationID == productionStation.ID {
				accept = true
				break
			}
		}

		//工序朝向，大于0表示当前工序在当前工位之后，可以放行；小于0表示当前工序在当前工位之前，禁止放行。
		if !accept {
			stationRoute := &model.ProductionProcess{}
			if err := tx.Joins("JOIN production_process_available_stations ON production_processes.id=production_process_available_stations.production_process_id").
				Where("production_process_available_stations.production_station_id = ?", productionStation.ID).
				Where("production_processes.production_line_id = ?", productOrder.ProductionLineID).Order("sort_index").First(stationRoute).Error; err == gorm.ErrRecordNotFound {
				code = 3
				return fmt.Errorf("工艺路线错误，且当前工位并不在当前产线的工艺路线之内，请联系管理员处理")
			} else if err != nil {
				code = 1
				return err
			}

			toward = productionProcess.SortIndex - stationRoute.SortIndex
			towardStr := "前"
			if toward > 0 {
				towardStr = "后"
			}

			code = 3
			isAccept = true
			return fmt.Errorf("工艺路线错误，且此产品的当前工序为%s(%s)，在当前工位的执行工序之%s", productionProcess.Description, productionProcess.Code, towardStr)
		}

		//检查人员资质
		//检查工序是否启用人员管控
		if productionProcess.EnableControl {
			//获取产品型号
			productModel := &model.ProductModel{}
			if err := tx.First(productModel, "`id` = ?", productOrder.ProductModelID).Error; err == gorm.ErrRecordNotFound {
				code = 1
				return fmt.Errorf("读取产品型号失败")
			} else if err != nil {
				code = 1
				return err
			}

			//获取人员资格
			personnelQualification := &model.PersonnelQualification{}
			if err := tx.Joins("JOIN personnel_qualification_types ON personnel_qualifications.personnel_qualification_type_id=personnel_qualification_types.id").
				Joins("JOIN personnel_qualification_type_available_models ON personnel_qualification_types.id=personnel_qualification_type_available_models.personnel_qualification_type_id").
				Where("personnel_qualification_types.production_process_id = ?", productionProcess.ID).
				Where("personnel_qualification_type_available_models.product_model_id = ?", productModel.ID).
				Where("personnel_qualifications.certified_user_id = ?", productionStation.CurrentUserID).
				First(personnelQualification).Error; err == gorm.ErrRecordNotFound {
				code = 1
				return fmt.Errorf("当前作业人员缺少认证资质，无法开工")
			} else if err != nil {
				code = 1
				return err
			}

			if personnelQualification.ExpirationDate.Time.Before(nowTime) || personnelQualification.ExpirationDate.Time.Equal(nowTime) {
				code = 1
				return fmt.Errorf("当前作业人员的认证资质已过期，无法开工")
			}
		}

		//获取作业手册
		productionProcessSop := &model.ProductionProcessSop{}
		if err := tx.Where("`production_process_id` = ? AND `product_model_id` = ?", productionProcess.ID, productOrder.ProductModelID).First(productionProcessSop).Error; err != nil && err != gorm.ErrRecordNotFound {
			code = 1
			return err
		}
		if productionProcessSop.FileLink != "" {
			sopLink = productionProcessSop.FileLink
		}

		if productionProcess.EnableReport {
			//创建系统事件上报开工
			systemEvent := &model.SystemEvent{}
			if err := tx.Preload("SystemEventParameters").Where("`code` = ? AND `enable` = ?", types.SystemEventProductionProcessStarted, true).First(systemEvent).Error; err != nil && err != gorm.ErrRecordNotFound {
				code = 1
				return err
			}

			if systemEvent.ID != "" {
				systemEventTrigger := &model.SystemEventTrigger{
					SystemEventID: systemEvent.ID,
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

					systemEventTrigger.SystemEventTriggerParameters = append(systemEventTrigger.SystemEventTriggerParameters, &model.SystemEventTriggerParameter{
						DataType:    _systemEventParameter.DataType,
						Name:        _systemEventParameter.Name,
						Description: _systemEventParameter.Description,
						Value:       value,
					})
				}

				if err := tx.Create(systemEventTrigger).Error; err != nil {
					code = 1
					return err
				}
			}
		}

		if productionStation.AllowReport {
			productionStation.CurrentState = types.ProductionStationStateOccupied
			//创建系统事件上报工位开始处于作业状态
			systemEvent := &model.SystemEvent{}
			if err := tx.Preload("SystemEventParameters").Where("`code` = ? AND `enable` = ?", types.SystemEventProductionStationOccupied, true).First(systemEvent).Error; err != nil && err != gorm.ErrRecordNotFound {
				code = 1
				return err
			}

			if systemEvent.ID != "" {
				systemEventTrigger := &model.SystemEventTrigger{
					SystemEventID: systemEvent.ID,
					EventNo:       uuid.NewString(),
					CurrentState:  types.SystemEventTriggerStateWaitExecute,
				}
				for _, _systemEventParameter := range systemEvent.SystemEventParameters {
					value := _systemEventParameter.Value
					value = strings.ReplaceAll(value, "{ProductionStation}", productionStation.Code)
					value = strings.ReplaceAll(value, "{ProductSerialNo}", productInfo.ProductSerialNo)
					value = strings.ReplaceAll(value, "{ProductionProcess}", productionProcess.Code)

					systemEventTrigger.SystemEventTriggerParameters = append(systemEventTrigger.SystemEventTriggerParameters, &model.SystemEventTriggerParameter{
						DataType:    _systemEventParameter.DataType,
						Name:        _systemEventParameter.Name,
						Description: _systemEventParameter.Description,
						Value:       value,
					})
				}

				if err := tx.Create(systemEventTrigger).Error; err != nil {
					code = 1
					return err
				}
			}

			if err := tx.Save(productionStation).Error; err != nil {
				code = 1
				return err
			}
		}

		return nil
	}); err != nil {
		if isAccept {
			return &proto.EnterProductionStationData{
				Toward: toward,
				ProductionProcess: &proto.EnterProductionStationInfo{
					Code:        productionProcess.Code,
					Description: productionProcess.Description,
					Identifier:  productionProcess.Identifier,
				}}, code, err
		}
		return nil, code, err
	}

	return &proto.EnterProductionStationData{
		ProductOrderNo:      productInfo.ProductOrder.ProductOrderNo,
		ProductSerialNo:     productInfo.ProductSerialNo,
		SalesOrderNo:        productInfo.ProductOrder.SalesOrderNo,
		ItemNo:              productInfo.ProductOrder.ItemNo,
		OrderTime:           utils.FormatSqlNullTime(productInfo.ProductOrder.OrderTime),
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
	}, 0, nil
}

// 获取作业步骤和作业参数
func GetProductionProcessStepWithParameter(req *proto.GetProductionProcessStepWithParameterRequest) (map[string]interface{}, error) {
	if req.ProductionStation == "" {
		return nil, fmt.Errorf("ProductionStation不能为空")
	}

	if req.TrayNo != "" {
		materialTray := &model.MaterialTray{}
		if err := model.DB.DB().Preload("ProductInfo").First(materialTray, "`identifier` = ?", req.TrayNo).Error; err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("无效的物料载具识别码")
		} else if err != nil {
			return nil, err
		}

		if materialTray.ProductInfo == nil {
			return nil, fmt.Errorf("载具未绑定任何产品")
		}

		req.ProductSerialNo = materialTray.ProductInfo.ProductSerialNo
	}

	if strings.TrimSpace(req.ProductSerialNo) == "" {
		return nil, fmt.Errorf("ProductSerialNo不能为空")
	}

	productInfo := &model.ProductInfo{}
	if err := model.DB.DB().First(productInfo, "`product_serial_no` = ?", req.ProductSerialNo).Error; err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("无效的产品序列号")
	} else if err != nil {
		return nil, err
	}

	if productInfo.ProductionProcessID == nil {
		return nil, fmt.Errorf("无法读取产品的当前工序")
	}

	//获取生产工艺
	productionProcess := &model.ProductionProcess{}
	if err := model.DB.DB().Preload("ProductionProcessAvailableStations").
		Preload("ProductionProcessAvailableStations.ProductionStation").
		First(productionProcess, "`id` = ?", productInfo.ProductionProcessID).Error; err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("读取产品的当前工序失败")
	} else if err != nil {
		return nil, err
	}

	var processable bool
	for _, v := range productionProcess.ProductionProcessAvailableStations {
		if v.ProductionStation.Code == req.ProductionStation {
			processable = true
			break
		}
	}
	if !processable {
		return nil, fmt.Errorf("非法操作，产品的当前工序不支持在此工位进行")
	}

	productionStation := &model.ProductionStation{}
	if err := model.DB.DB().First(productionStation, "`code` = ?", req.ProductionStation).Error; err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("无效的生产工站")
	} else if err != nil {
		return nil, err
	}

	productOrder := &model.ProductOrder{}
	if err := model.DB.DB().Preload("ProductOrderAttributes").
		Preload("ProductOrderAttributes.ProductAttribute").
		Preload("ProductModel").
		First(productOrder, "`id` = ?", productInfo.ProductOrderID).Error; err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("读取产品工单失败")
	} else if err != nil {
		return nil, err
	}

	productOrderBoms := []map[string]interface{}{}
	materialNoArray := []string{}
	if err := model.DB.DB().Model(model.ProductOrderBom{}).
		Where("`product_order_id` = ? AND `production_process` in ?", productOrder.ID, []string{productionProcess.Identifier, productionProcess.Code}).
		Select("material_no materialNo", "material_description materialDescription", "piece_qty pieceQTY", "require_qty requireQTY", "unit", "enable_control enableControl", "control_type controlType", "warehouse", "production_process productionProcess").
		Find(&productOrderBoms).Error; err != nil {
		return nil, err
	}
	for _, v := range productOrderBoms {
		materialNoArray = append(materialNoArray, v["materialNo"].(string))
	}

	materialChannels := []*model.MaterialChannel{}
	if err := model.DB.DB().Preload("MaterialChannelLayer").
		Joins("JOIN material_channel_layers ON material_channels.material_channel_layer_id=material_channel_layers.id").
		Where("material_channel_layers.production_station_id = ?", productionStation.ID).
		Find(&materialChannels).Error; err != nil {
		return nil, err
	}

	materialChannelGroup := map[int32][]*model.MaterialChannel{}
	for _, v := range materialChannels {
		materialChannelGroup[v.MaterialChannelLayer.LightRegisterAddress] = append(materialChannelGroup[v.MaterialChannelLayer.LightRegisterAddress], v)
	}

	materialChannelLayers := []map[string]interface{}{}
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

	_productionProcessSteps := []*model.ProductionProcessStep{}
	if err := model.DB.DB().Preload("AttributeExpressions").Preload("ProcessStepType").Preload("ProcessStepType.ProcessStepTypeParameters").
		Joins("JOIN available_processes ON production_process_steps.id=available_processes.production_process_step_id").
		Where("available_processes.production_process_id = ? AND production_process_steps.enable = ?", productionProcess.ID, true).
		Order("sort_index").
		Find(&_productionProcessSteps).Error; err != nil {
		return nil, err
	}

	productionProcessSteps := []map[string]interface{}{}
	for _, v := range _productionProcessSteps {
		match := v.InitialValue
		for _, attributeExpression := range v.AttributeExpressions {
			match = false
			for _, productOrderAttribute := range productOrder.ProductOrderAttributes {
				if productOrderAttribute.ProductAttributeID == attributeExpression.ProductAttributeID {
					fmt.Println(productOrderAttribute.Value, attributeExpression.MathOperator, attributeExpression.AttributeValue)
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
			processStepTypeParameters := make([]map[string]interface{}, len(v.ProcessStepType.ProcessStepTypeParameters))
			for i, vv := range v.ProcessStepType.ProcessStepTypeParameters {
				processStepTypeParameters[i] = map[string]interface{}{
					"code":        vv.Code,
					"description": vv.Description,
					"value":       vv.DefaultValue,
				}
			}
			productionProcessSteps = append(productionProcessSteps, map[string]interface{}{
				"id":                  v.ID,
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

	productOrderProcessStep := []*model.ProductOrderProcessStep{}
	if err := model.DB.DB().Preload("ProductOrderProcessStepTypeParameters").Preload("ProductOrderProcessStepTypeParameters.ProcessStepTypeParameter").
		Joins("JOIN product_order_processes ON product_order_process_steps.product_order_process_id=product_order_processes.id").
		Where("product_order_processes.production_process_id=? AND product_order_processes.product_order_id=?", productionProcess.ID, productOrder.ID).
		Find(&productOrderProcessStep).Error; err != nil {
		return nil, err
	}

	productOrderProcess := []map[string]interface{}{}
	for _, v := range productOrderProcessStep {
		parameters := make([]map[string]interface{}, len(v.ProductOrderProcessStepTypeParameters))
		for i, vv := range v.ProductOrderProcessStepTypeParameters {
			parameters[i] = map[string]interface{}{
				"Code":        vv.ProcessStepTypeParameter.Code,
				"description": vv.ProcessStepTypeParameter.Description,
				"value":       vv.Value,
			}
		}
		productOrderProcess = append(productOrderProcess, map[string]interface{}{
			"id":                  v.ID,
			"sortIndex":           v.SortIndex,
			"code":                "",
			"workDescription":     v.WorkDescription,
			"processStepType":     v.ProcessStepType.Code,
			"processStepTypeDesc": v.ProcessStepType.Description,
			"workGraphic":         v.WorkGraphic,
			"remark":              v.Remark,
			"workResult":          "",
			"parameters":          parameters,
		})
	}

	productWorkRecords := []map[string]interface{}{}
	if err := model.DB.DB().Model(model.ProductWorkRecord{}).
		Select("production_process_step_id AS productionProcessStepID", "is_qualified AS isQualified").
		Where("`product_info_id` = ?", productInfo.ID).
		Find(&productWorkRecords).Error; err != nil {
		return nil, err
	}

	for i, v := range productionProcessSteps {
		for _, vv := range productWorkRecords {
			if vv["productionProcessStepID"] == v["id"] {
				workResult := ""
				if vv["isQualified"].(bool) {
					workResult = "PASS"
				}
				productionProcessSteps[i]["workResult"] = workResult
				break
			}
		}
	}

	var materialTray string
	if err := model.DB.DB().Model(model.MaterialTray{}).Where("`product_info_id` = ? AND `tray_type` = ?", productInfo.ID, types.MaterialTrayTypeMaterialTray).Select("identifier").Scan(&materialTray).Error; err != nil {
		return nil, err
	}
	var assembleTray string
	if err := model.DB.DB().Model(model.MaterialTray{}).Where("`product_info_id` = ? AND `tray_type` = ?", productInfo.ID, types.MaterialTrayTypeAssembleTray).Select("identifier").Scan(&assembleTray).Error; err != nil {
		return nil, err
	}

	productOrderAttributes := make([]map[string]interface{}, len(productOrder.ProductOrderAttributes))
	for i, v := range productOrder.ProductOrderAttributes {
		productOrderAttributes[i] = map[string]interface{}{
			"id":               v.ID,
			"code":             v.ProductAttribute.Code,
			"codeDescription":  v.ProductAttribute.Description,
			"value":            v.Value,
			"valueDescription": v.Description,
		}
	}

	return map[string]interface{}{
		"productOrder": map[string]interface{}{
			"id":                  productOrder.ID,
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
			"id":              productInfo.ID,
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
		ProductionProcessID: *productInfo.ProductionProcessID,
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

	if productInfo.ProductionProcessID == nil {
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
func ExitProductionStation(req *proto.ExitProductionStationRequest) error {
	if req.ProductionStation == "" {
		return fmt.Errorf("ProductionStation不能为空")
	}
	nowTime := time.Now()
	if err := model.DB.DB().Transaction(func(tx *gorm.DB) error {
		if req.TrayNo != "" {
			//根据载具号获取物料载具
			materialTray := &model.MaterialTray{}
			if err := tx.Preload("ProductInfo").First(materialTray, "`identifier` = ?", req.TrayNo).Error; err == gorm.ErrRecordNotFound {
				return fmt.Errorf("无效的物料载具识别码")
			} else if err != nil {
				return err
			}

			if materialTray.ProductInfo == nil {
				return fmt.Errorf("载具未绑定任何产品")
			}

			req.ProductSerialNo = materialTray.ProductInfo.ProductSerialNo
		}
		if req.PackageNo != "" {
			//根据包装箱号获取产品包装记录
			productPackageRecord := &model.ProductPackageRecord{}
			if err := tx.Preload("ProductInfo").First(productPackageRecord, "`package_no` = ?", req.PackageNo).Error; err == gorm.ErrRecordNotFound {
				return fmt.Errorf("无效的包装箱号")
			} else if err != nil {
				return err
			}

			if productPackageRecord.ProductInfo == nil {
				return fmt.Errorf("包装箱未绑定任何产品")
			}

			req.ProductSerialNo = productPackageRecord.ProductInfo.ProductSerialNo
		}

		req.ProductSerialNo = strings.Trim(strings.Trim(req.ProductSerialNo, "\000"), "\r")
		if req.ProductSerialNo == "" {
			return fmt.Errorf("ProductSerialNo不能为空")
		}

		productionStation := &model.ProductionStation{}
		if err := tx.Preload("ProductionLine").First(productionStation, "`code` = ?", req.ProductionStation).Error; err == gorm.ErrRecordNotFound {
			return fmt.Errorf("非法的工站代号")
		} else if err != nil {
			return err
		}

		productInfo := &model.ProductInfo{}
		if err := tx.Preload("ProductOrder").First(productInfo, "`product_serial_no` = ?", req.ProductSerialNo).Error; err == gorm.ErrRecordNotFound {
			return fmt.Errorf("读取产品信息失败")
		} else if err != nil {
			return err
		}

		if productInfo.ProductionProcessID == nil {
			return fmt.Errorf("无法读取产品的当前工序")
		}

		//上传节拍
		productRhythmRecord := &model.ProductRhythmRecord{}
		if err := tx.Where("`production_station_id` = ? AND `product_info_id` = ? AND `work_end_time` IS NULL", productionStation.ID, productInfo.ID).First(productRhythmRecord).Error; err == gorm.ErrRecordNotFound {
			return fmt.Errorf("读取工位当前节拍数据失败")
		} else if err != nil {
			return err
		}

		productRhythmRecord.WorkEndTime = tool.Time2NullTime(nowTime)
		productRhythmRecord.WorkTime = int32(tool.NullTime2Time(productRhythmRecord.WorkEndTime).Sub(tool.NullTime2Time(productRhythmRecord.WorkStartTime)).Seconds())
		productRhythmRecord.OverTime = productRhythmRecord.WorkTime - productRhythmRecord.StandardWorkTime
		if productRhythmRecord.OverTime < 0 {
			productRhythmRecord.OverTime = 0
		}
		productRhythmRecord.WaitTime = req.WaitTime

		//修改工艺记录
		lastProductProcessRoute := &model.ProductProcessRoute{}
		if err := tx.Where("`product_info_id` = ? AND `current_process_id` = ? AND `current_state` = ?", productInfo.ID, productInfo.ProductionProcessID, types.ProductProcessRouteStateProcessing).First(lastProductProcessRoute).Error; err == gorm.ErrRecordNotFound {
			return fmt.Errorf("读取产品当前工艺路线失败")
		} else if err != nil {
			return err
		}

		if req.IsRework {
			productRhythmRecord.IsRework = true
			lastProductProcessRoute.CurrentState = types.ProductProcessRouteStateReworking

			if err := tx.Create(&model.ProductReworkRecord{
				ProductionStationID: productionStation.ID,
				ProductInfoID:       productInfo.ID,
				ProductionProcessID: *productInfo.ProductionProcessID,
				ReworkTime:          tool.Time2NullTime(nowTime),
				ReworkReason:        req.ReworkReason,
			}).Error; err != nil {
				return err
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
			nextProductProcessRoute := &model.ProductProcessRoute{}
			if err := tx.Preload("CurrentProcess").Where("`product_info_id` = ? AND `route_index` > ? AND `current_state` = ?", productInfo.ID, lastProductProcessRoute.RouteIndex, types.ProductProcessRouteStateWaitProcess).Order("work_index").First(nextProductProcessRoute).Error; err != nil && err != gorm.ErrRecordNotFound {
				return err
			}

			//兼容，部分产线是直接创建产品工艺路线，部分是根据工单工艺动态创建
			if nextProductProcessRoute.ID == "" {
				productOrderProcess := &model.ProductOrderProcess{}
				if err := tx.Preload("ProductionProcess").Where("`product_order_id` = ? AND `enable` = ? AND `sort_index` > ?", productInfo.ProductOrderID, true, lastProductProcessRoute.RouteIndex).Order("sort_index").First(productOrderProcess).Error; err != nil && err != gorm.ErrRecordNotFound {
					return err
				}

				if productOrderProcess.ID != "" {
					nextProductProcessRoute = &model.ProductProcessRoute{
						LastProcessID:    &lastProductProcessRoute.CurrentProcessID,
						CurrentProcessID: productOrderProcess.ProductionProcessID,
						CurrentProcess:   productOrderProcess.ProductionProcess,
						CurrentState:     types.ProductProcessRouteStateWaitProcess,
						RouteIndex:       productOrderProcess.SortIndex,
						ProductInfoID:    productInfo.ID,
					}
					if err := tx.Create(nextProductProcessRoute).Error; err != nil {
						return err
					}
				}
			}

			if nextProductProcessRoute.ID != "" {
				nextProductProcessRoute.WorkIndex = lastProductProcessRoute.WorkIndex + 1
				if nextProductProcessRoute.CurrentProcess != nil && nextProductProcessRoute.CurrentProcess.ProductState != "" {
					//设定当前工序的产品信息状态
					productInfo.CurrentState = nextProductProcessRoute.CurrentProcess.ProductState
				}
				productInfo.ProductionProcessID = &nextProductProcessRoute.CurrentProcessID

				//计算预计下线时间
				var remainingRoutes int64
				if err := tx.Model(model.ProductOrderProcess{}).Where("`product_order_id` = ? AND `enable` = ? AND `sort_index` > ?", productInfo.ProductOrderID, true, nextProductProcessRoute.RouteIndex).Count(&remainingRoutes).Error; err != nil {
					return err
				}

				productInfo.RemainingRoutes = int32(remainingRoutes)
				if remainingRoutes > 0 {
					productInfo.EstimateTime = tool.Time2NullTime(nowTime.Add(time.Duration(int32(remainingRoutes)*productInfo.ProductOrder.StandardWorkTime) * time.Second))
				}
			} else {
				//没有下一个工序判定为完工
				productInfo.CurrentState = types.ProductStateCompleted
				productInfo.FinishedTime = tool.Time2NullTime(nowTime)
				productInfo.ProductionProcessID = nil

				productOrder := &model.ProductOrder{}
				if err := tx.Preload("ProductInfos").First(productOrder, "`id` = ?", productInfo.ProductOrderID).Error; err != nil && err != gorm.ErrRecordNotFound {
					return err
				}

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
					productOrder.ActualFinishTime = tool.Time2NullTime(nowTime)

					systemEvent := &model.SystemEvent{}
					if err := tx.Preload("SystemEventParameters").First(systemEvent, "`code` = ?", types.SystemEventProductOrderFinished).Error; err != nil && err != gorm.ErrRecordNotFound {
						return err
					}
					if systemEvent.ID != "" {
						systemEventTrigger := &model.SystemEventTrigger{
							SystemEventID: systemEvent.ID,
							EventNo:       uuid.NewString(),
							CurrentState:  types.SystemEventTriggerStateWaitExecute,
						}
						for _, _systemEventParameter := range systemEvent.SystemEventParameters {
							value := _systemEventParameter.Value
							value = strings.ReplaceAll(value, "{ProductOrderNo}", productOrder.ProductOrderNo)

							systemEventTrigger.SystemEventTriggerParameters = append(systemEventTrigger.SystemEventTriggerParameters, &model.SystemEventTriggerParameter{
								DataType:    _systemEventParameter.DataType,
								Name:        _systemEventParameter.Name,
								Description: _systemEventParameter.Description,
								Value:       value,
							})
						}

						if err := tx.Create(systemEventTrigger).Error; err != nil {
							return err
						}
					}
				}

				if err := tx.Save(productOrder).Error; err != nil {
					return err
				}
			}

			//创建产量记录
			productionStationOutput := &model.ProductionStationOutput{}
			if err := tx.Where("`production_station_id` = ? AND `product_info_id` = ?", productionStation.ID, productInfo.ID).First(productionStationOutput).Error; err != nil && err != gorm.ErrRecordNotFound {
				return err
			}
			if productionStationOutput.ID == "" {
				if err := tx.Create(&model.ProductionStationOutput{
					LoginUserID:         *productionStation.CurrentUserID,
					ProductionProcessID: productRhythmRecord.ProductionProcessID,
					ProductInfoID:       productRhythmRecord.ProductInfoID,
					ProductionStationID: productRhythmRecord.ProductionStationID,
				}).Error; err != nil {
					return err
				}
			} else {
				productionStationOutput.LoginUserID = *productionStation.CurrentUserID
				if err := tx.Save(productionStationOutput).Error; err != nil {
					return err
				}
			}

			// 判断是否要解绑载具
			if req.UnbindTray {
				materialTray := &model.MaterialTray{}
				if err := tx.First(materialTray, "`product_info_id` = ?", productInfo.ID).Error; err != nil && err != gorm.ErrRecordNotFound {
					return err
				}
				if materialTray.ID != "" {
					materialTray.ProductInfoID = nil
					materialTray.CurrentState = types.MaterialTrayStateWaitBind
					if err := tx.Save(materialTray).Error; err != nil {
						return err
					}
				}
			}

			productionProcess := &model.ProductionProcess{}
			if err := tx.First(productionProcess, "`id` = ?", productRhythmRecord.ProductionProcessID).Error; err != nil && err != gorm.ErrRecordNotFound {
				return err
			}

			if productionProcess.ID != "" {
				if productionProcess.EnableReport {
					//创建系统事件上报完工
					systemEvent := &model.SystemEvent{}
					if err := tx.Preload("SystemEventParameters").First(systemEvent, "`code` = ?", types.SystemEventProductionProcessFinished).Error; err != nil && err != gorm.ErrRecordNotFound {
						return err
					}
					if systemEvent.ID != "" {
						systemEventTrigger := &model.SystemEventTrigger{
							SystemEventID: systemEvent.ID,
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

							systemEventTrigger.SystemEventTriggerParameters = append(systemEventTrigger.SystemEventTriggerParameters, &model.SystemEventTriggerParameter{
								DataType:    _systemEventParameter.DataType,
								Name:        _systemEventParameter.Name,
								Description: _systemEventParameter.Description,
								Value:       value,
							})
						}

						if err := tx.Create(systemEventTrigger).Error; err != nil {
							return err
						}
					}
				}
			}

			if productionStation.AllowReport {
				productionStation.CurrentState = types.ProductionStationStateStandby
				// 创建系统事件上报工位完成作业
				systemEvent := &model.SystemEvent{}
				if err := tx.Preload("SystemEventParameters").First(systemEvent, "`code` = ?", types.SystemEventProductionStationReleased).Error; err != nil && err != gorm.ErrRecordNotFound {
					return err
				}
				if systemEvent.ID != "" {
					systemEventTrigger := &model.SystemEventTrigger{
						SystemEventID: systemEvent.ID,
						EventNo:       uuid.NewString(),
						CurrentState:  types.SystemEventTriggerStateWaitExecute,
					}
					for _, _systemEventParameter := range systemEvent.SystemEventParameters {
						value := _systemEventParameter.Value
						value = strings.ReplaceAll(value, "{ProductionStation}", productionStation.Code)

						systemEventTrigger.SystemEventTriggerParameters = append(systemEventTrigger.SystemEventTriggerParameters, &model.SystemEventTriggerParameter{
							DataType:    _systemEventParameter.DataType,
							Name:        _systemEventParameter.Name,
							Description: _systemEventParameter.Description,
							Value:       value,
						})
					}

					if err := tx.Create(systemEventTrigger).Error; err != nil {
						return err
					}
				}
			}
		}

		if err := tx.Save(productRhythmRecord).Error; err != nil {
			return err
		}
		if err := tx.Save(lastProductProcessRoute).Error; err != nil {
			return err
		}
		if err := tx.Save(productInfo).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

// 设置失败后续处理
func CheckProductProcessRouteFailure(req *proto.CheckProductProcessRouteFailureRequest) error {
	if req.ProductionStation == "" {
		return fmt.Errorf("ProductionStation不能为空")
	}
	if req.ProductSerialNo == "" {
		return fmt.Errorf("ProductSerialNo不能为空")
	}

	if err := model.DB.DB().Transaction(func(tx *gorm.DB) error {
		nowTime := time.Now()

		productionStation := &model.ProductionStation{}
		if err := tx.First(productionStation, "`code` = ?", req.ProductionStation).Error; err == gorm.ErrRecordNotFound {
			return fmt.Errorf("无效的工站代号")
		} else if err != nil {
			return err
		}

		productInfo := &model.ProductInfo{}
		if err := tx.Preload("ProductOrder").First(productInfo, "`product_serial_no` = ?", req.ProductSerialNo).Error; err == gorm.ErrRecordNotFound {
			return fmt.Errorf("无效的产品序列号")
		} else if err != nil {
			return err
		}

		if productInfo.ProductionProcessID == nil {
			return fmt.Errorf("数据错误，此产品的当前工序为空")
		}

		lastProductProcessRoute := &model.ProductProcessRoute{}
		if err := tx.Preload("CurrentProcess").Where("`current_process_id` = ? AND `product_info_id` = ? AND `current_state` = ?", productInfo.ProductionProcessID, productInfo.ID, types.ProductProcessRouteStateChecking).First(lastProductProcessRoute).Error; err == gorm.ErrRecordNotFound {
			return fmt.Errorf("状态错误，此产品的当前工艺状态不是" + types.ProductProcessRouteStateChecking)
		} else if err != nil {
			return err
		}

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

			if err := tx.Create(&model.ProductReworkRecord{
				ProductionStationID: productionStation.ID,
				ProductInfoID:       productInfo.ID,
				ProductionProcessID: *productInfo.ProductionProcessID,
				ReworkTime:          tool.Time2NullTime(nowTime),
				ReworkReason:        lastProductProcessRoute.Remark,
			}).Error; err != nil {
				return err
			}

			productInfo.CurrentState = types.ProductStateReworking
		case types.ProductionProcessHandleMethodIgnore:
			lastProductProcessRoute.CurrentState = types.ProductProcessRouteStateProcessed

			//切换到下个工艺
			nextProductProcessRoute := &model.ProductProcessRoute{}
			if err := tx.Preload("CurrentProcess").Where("`product_info_id` = ? AND `route_index` > ? AND `current_state` = ?", productInfo.ID, lastProductProcessRoute.RouteIndex, types.ProductProcessRouteStateWaitProcess).Order("route_index").First(nextProductProcessRoute).Error; err != nil && err != gorm.ErrRecordNotFound {
				return err
			}

			if nextProductProcessRoute.ID == "" {
				productOrderProcess := &model.ProductOrderProcess{}
				if err := tx.Preload("ProductionProcess").Where("`product_order_id` = ? AND `enable` = ? AND `sort_index` > ?", productInfo.ProductOrderID, true, lastProductProcessRoute.RouteIndex).Order("sort_index").First(productOrderProcess).Error; err != nil && err != gorm.ErrRecordNotFound {
					return err
				}

				if productOrderProcess.ID != "" {
					nextProductProcessRoute = &model.ProductProcessRoute{
						LastProcessID:    &lastProductProcessRoute.CurrentProcessID,
						CurrentProcessID: productOrderProcess.ProductionProcessID,
						CurrentProcess:   productOrderProcess.ProductionProcess,
						CurrentState:     types.ProductProcessRouteStateWaitProcess,
						RouteIndex:       productOrderProcess.SortIndex,
						ProductInfoID:    productInfo.ID,
					}
					if err := tx.Create(nextProductProcessRoute).Error; err != nil {
						return err
					}
				}
			}

			if nextProductProcessRoute.ID != "" {
				nextProductProcessRoute.WorkIndex = lastProductProcessRoute.WorkIndex + 1
				nextProductProcessRoute.CurrentState = types.ProductProcessRouteStateProcessing

				if nextProductProcessRoute.CurrentProcess.ProductState != "" {
					//设定当前工序的产品信息状态
					productInfo.CurrentState = nextProductProcessRoute.CurrentProcess.ProductState
				}
				productInfo.ProductionProcessID = &nextProductProcessRoute.CurrentProcessID

				//TODO: 计算预计下线时间
				var remainingRoutes int64
				if err := tx.Model(model.ProductProcessRoute{}).Where("`product_info_id` = ? AND `current_state` = ?", productInfo.ID, types.ProductProcessRouteStateWaitProcess).Count(&remainingRoutes).Error; err != nil {
					return err
				}

				productInfo.RemainingRoutes = int32(remainingRoutes)
				if remainingRoutes > 0 {
					productInfo.EstimateTime = tool.Time2NullTime(nowTime.Add(time.Duration(int32(remainingRoutes)*productInfo.ProductOrder.StandardWorkTime) * time.Second))
				}

				if err := tx.Save(nextProductProcessRoute).Error; err != nil {
					return err
				}
			} else {
				productInfo.CurrentState = types.ProductStateCompleted
				productInfo.FinishedTime = tool.Time2NullTime(nowTime)
				productInfo.ProductionProcessID = nil
			}
		default:
			return fmt.Errorf("无效的处理方式")
		}

		if err := tx.Save(productInfo).Error; err != nil {
			return err
		}
		if err := tx.Save(lastProductProcessRoute).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
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
		productReworkRecord.CompleteTime = tool.Time2NullTime(time.Now())

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
			CurrentProcessID: *productInfo.ProductionProcessID,
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
			LastProcessID:    productInfo.ProductionProcessID,
			CurrentProcessID: nextProcess.ID,
			RouteIndex:       nextProcess.SortIndex,
			ProductInfoID:    productInfo.ID,
			WorkIndex:        lastProductProcessRoute.WorkIndex + 1,
		}).Error; err != nil {
			return err
		}
		productReworkRecord.ProductionProcessID = nextProcess.ID
		productInfo.ProductionProcessID = &nextProcess.ID

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
			"completeTime":            utils.FormatSqlNullTime(productReworkRecord.CompleteTime),
			"createTime":              utils.FormatTime(productReworkRecord.CreateTime),
		},
	}, nil
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
