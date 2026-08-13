package logic

import (
	"fmt"
	"strings"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/tool"
	"gorm.io/gorm"
)

// GetTestProjectWithParameter 获取测试项接口数据
// 根据测试工位代号（或托盘号/产品序列号）定位产品及其工单，
// 匹配测试工序的作业步骤（按工单特性表达式筛选），
// 汇总测试项目及其输入/输出参数，供测试设备下发使用
func GetTestProjectWithParameter(req *proto.GetTestProjectWithParameterRequest) (*proto.GetTestProjectWithParameterResponse, error) {
	response := &proto.GetTestProjectWithParameterResponse{}
	if req.ProductionStation == "" {
		return nil, fmt.Errorf("ProductionStation不能为空")
	}

	if req.TrayNo != "" {
		materialTray := &model.MaterialTray{}
		if err := model.DB.DB().Preload("ProductInfo").First(materialTray, "`identifier` = ?", req.TrayNo).Error; err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("无效的托盘号")
		} else if err != nil {
			return nil, err
		}

		if materialTray.ProductInfo == nil {
			return nil, fmt.Errorf("托盘未绑定任何产品")
		}
		req.ProductSerialNo = materialTray.ProductInfo.ProductSerialNo
	}
	req.ProductSerialNo = strings.Trim(strings.Trim(req.ProductSerialNo, "\000"), "\r")
	if req.ProductSerialNo == "" {
		return nil, fmt.Errorf("ProductSerialNo不能为空")
	}

	productInfo := &model.ProductInfo{}
	if err := model.DB.DB().Preload("ProductOrder").First(productInfo, "`product_serial_no` = ?", req.ProductSerialNo).Error; err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("读取产品信息失败")
	} else if err != nil {
		return nil, err
	}

	productionLineID := productInfo.ProductOrder.ProductionLineID
	if productionLineID == nil || *productionLineID == "" {
		return nil, fmt.Errorf("序列号所在的工单%s发放的产线为空", productInfo.ProductOrder.ProductOrderNo)
	}

	//测试工序以测试工位代号作为工序代号
	productionProcess := &model.ProductionProcess{}
	if err := model.DB.DB().First(productionProcess, "`code` = ? AND `production_line_id` = ?", req.ProductionStation, productionLineID).Error; err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("未能找到对应的测试工序")
	} else if err != nil {
		return nil, err
	}

	productOrder := &model.ProductOrder{}
	if err := model.DB.DB().Preload("ProductOrderAttributes").
		Preload("ProductOrderAttributes.ProductAttribute").
		First(productOrder, "`id` = ?", productInfo.ProductOrderID).Error; err == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("读取产品工单失败")
	} else if err != nil {
		return nil, err
	}

	//查找测试工序下启用的作业步骤（含步骤类型与参数定义）
	_productionProcessSteps := []*model.ProductionProcessStep{}
	if err := model.DB.DB().Preload("AttributeExpressions").Preload("ProcessStepType").Preload("ProcessStepType.ProcessStepTypeParameters").
		Joins("JOIN available_processes ON production_process_steps.id=available_processes.production_process_step_id").
		Where("available_processes.production_process_id = ? AND production_process_steps.enable = ?", productionProcess.ID, true).
		Order("sort_index").
		Find(&_productionProcessSteps).Error; err != nil {
		return nil, err
	}

	//按工单特性表达式匹配测试步骤
	matchedProcessSteps := []*model.ProductionProcessStep{}
	for _, v := range _productionProcessSteps {
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
			if !match {
				break
			}
		}
		if match {
			matchedProcessSteps = append(matchedProcessSteps, v)
		}
	}
	if len(matchedProcessSteps) == 0 {
		return nil, fmt.Errorf("无法匹配测试项，请检查测试项匹配规则")
	}

	testProjects := make([]*proto.TestProjectInfo, 0)
	inputParameters := []*proto.ParameterInfo{}
	outputParameters := []*proto.ParameterInfo{}
	for _, v := range matchedProcessSteps {
		for _, p := range v.ProcessStepType.ProcessStepTypeParameters {
			parameter := &proto.ParameterInfo{
				ProjectCode:   v.Code,
				Code:          p.Code,
				Description:   p.Description,
				StandardValue: p.DefaultValue,
				MinimumValue:  p.MinimumValue,
				MaximumValue:  p.MaximumValue,
				Unit:          p.Unit,
			}
			if p.ParameterType {
				inputParameters = append(inputParameters, parameter)
			} else {
				outputParameters = append(outputParameters, parameter)
			}
		}

		testProjects = append(testProjects, &proto.TestProjectInfo{
			Id:          v.ID,
			Code:        v.Code,
			Description: v.Description,
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
