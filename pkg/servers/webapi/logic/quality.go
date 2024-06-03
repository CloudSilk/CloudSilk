package logic

import (
	"context"
	"fmt"
	"strings"

	"github.com/CloudSilk/CloudSilk/pkg/clients"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/tool"
	"gorm.io/gorm"
)

// 获取测试项接口数据
func GetTestProjectWithParameter(req *proto.GetTestProjectWithParameterRequest) (*proto.GetTestProjectWithParameterResponse, error) {
	response := &proto.GetTestProjectWithParameterResponse{}
	if req.ProductionStation == "" {
		return nil, fmt.Errorf("ProductionStation不能为空")
	}

	if req.TrayNo != "" {
		_material, err := clients.MaterialTrayClient.Get(context.Background(), &proto.GetMaterialTrayRequest{Identifier: req.TrayNo})
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("无效的托盘号")
		} else if err != nil {
			return nil, err
		}

		material := _material.Data
		if material.ProductInfoID == "" {
			return nil, fmt.Errorf("托盘未绑定任何产品")
		}

		req.ProductSerialNo = material.ProductInfo.ProductSerialNo
	}
	req.ProductSerialNo = strings.Trim(strings.Trim(req.ProductSerialNo, "\000"), "\r")
	if req.ProductSerialNo == "" {
		return nil, fmt.Errorf("ProductSerialNo不能为空")
	}

	_productInfo, err := clients.ProductInfoClient.Get(context.Background(), &proto.GetProductInfoRequest{ProductSerialNo: req.ProductSerialNo})
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("读取产品信息失败")
	} else if err != nil {
		return nil, err
	}

	productInfo := _productInfo.Data
	productionLineID := productInfo.ProductOrder.ProductionLineID
	if productionLineID == "" {
		return nil, fmt.Errorf("序列号所在的工单%s发放的产线为空", productInfo.ProductOrder.ProductOrderNo)
	}

	_productionProcess, err := clients.ProductionProcessClient.Query(context.Background(), &proto.QueryProductionProcessRequest{
		PageSize:         1,
		Code:             req.ProductionStation,
		ProductionLineID: productionLineID,
	})
	if err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("未能找到对应的测试工序")
	} else if err != nil {
		return nil, err
	}
	productionProcess := _productionProcess.Data[0]

	// propertyInfoes := reflect.TypeOf(model.ProductOrderAttribute{})
	// var attributeIDProperty, attributeValueProperty *reflect.StructField
	// for i := 0; i < propertyInfoes.NumField(); i++ {
	// 	field := propertyInfoes.Field(i)
	// 	switch field.Name {
	// 	case "ProductAttributeID":
	// 		attributeIDProperty = &field
	// 	case "Value":
	// 		attributeValueProperty = &field
	// 	}
	// }
	_productOrderAttributes, err := clients.ProductOrderAttributeClient.Query(context.Background(), &proto.QueryProductOrderAttributeRequest{
		PageSize:       1000,
		ProductOrderID: productInfo.ProductOrderID,
	})
	if err != nil {
		return nil, err
	}
	productOrderAttributes := _productOrderAttributes.Data

	//查找匹配的测试规则
	_processStepMatchRules, err := clients.ProcessStepParameterClient.Query(context.Background(), &proto.QueryProcessStepParameterRequest{
		PageSize:                   1000,
		SortConfig:                 "priority",
		Enable:                     true,
		HasProductionProcessStepID: true,
		Code:                       productionProcess.Code,
	})
	if err != nil {
		return nil, err
	}

	productionProcessStepIDs := []string{}
	for _, _productOrderAttributes := range productOrderAttributes {
		productAttributeID := _productOrderAttributes.ProductAttributeID
		value := _productOrderAttributes.ProductAttribute.DefaultValue
		for _, _processStepMatchRule := range _processStepMatchRules.Data {
			// initialValue := _processStepMatchRule.InitialValue
			initialValue := true
			// match := initialValue
			for _, _attributeExpression := range _processStepMatchRule.AttributeExpressions {
				if _attributeExpression.ProductAttributeID == productAttributeID && initialValue {
					b, err := tool.MathOperator(value, _attributeExpression.MathOperator, _attributeExpression.AttributeValue)
					if err != nil {
						return nil, err
					}
					if b {
						break
					}
					// productionProcessStepIDs = append(productionProcessStepIDs, _processStepMatchRule.ProductionProcessStepID)
				}

				// attributeIDPropertyType := attributeIDProperty.Type.String()
				// attributeIDCriterionMethod := "="
				// attributeID = _attributeExpression.ProductAttributeID

				// attributeValuePropertyType = attributeValueProperty.Type.String()
				// attributeValueCriterionMethod = _attributeExpression.MathOperator
				// attributeValue := _attributeExpression.AttributeValue
			}
		}
	}
	if len(productionProcessStepIDs) == 0 {
		return nil, fmt.Errorf("无法匹配测试项，请检查测试项匹配规则")
	}

	_productionProcessSteps, err := clients.ProductionProcessStepClient.Query(context.Background(), &proto.QueryProductionProcessStepRequest{
		Ids: productionProcessStepIDs,
	})
	if err != nil {
		return nil, err
	}

	testProjects := make([]*proto.TestProjectInfo, 0)
	inputParameters := []*proto.ParameterInfo{}
	outputParameters := []*proto.ParameterInfo{}
	for _, _productionProcessStep := range _productionProcessSteps.Data {
		for _, v := range _productionProcessStep.ProcessStepType.ProcessStepTypeParameters {
			parameter := &proto.ParameterInfo{
				Code:          v.Code,
				Description:   v.Description,
				StandardValue: v.DefaultValue,
				MaximumValue:  v.MaximumValue,
				MinimumValue:  v.MinimumValue,
				Unit:          v.Unit,
				GroupCode:     _productionProcessStep.ProcessStepType.Code,
			}
			if v.ParameterType {
				inputParameters = append(inputParameters, parameter)
			} else {
				outputParameters = append(outputParameters, parameter)
			}
		}

		testProjects = append(testProjects, &proto.TestProjectInfo{
			Id:          _productionProcessStep.Id,
			Code:        _productionProcessStep.Code,
			Description: _productionProcessStep.Description,
		})
	}
	response.Code = 1
	response.Data = &proto.TestProjectWithParameterInfo{
		TestProjects:     testProjects,
		InputParameters:  inputParameters,
		OutputParameters: outputParameters,
	}

	return response, nil
}
