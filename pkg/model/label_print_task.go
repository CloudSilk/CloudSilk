package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 标签打印任务
type LabelPrintTask struct {
	ModelID
	Code                string             `gorm:"size:50;comment:代号"`
	Description         string             `gorm:"size:500;comment:描述"`
	Enable              bool               `gorm:"comment:是否启用"`
	PrintCopies         int32              `gorm:"comment:打印份数"`
	ProductionLineID    *string            `gorm:"size:36;comment:生产产线ID"`
	ProductionLine      *ProductionLine    `gorm:"constraint:OnDelete:SET NULL"` //生产产线
	ProductCategoryID   *string            `gorm:"size:36;comment:产品类别ID"`
	ProductCategory     *ProductCategory   `gorm:"constraint:OnDelete:SET NULL"` //产品类别
	LabelTypeID         string             `gorm:"size:36;comment:标签类型ID"`
	LabelType           *LabelType         `gorm:"constraint:OnDelete:CASCADE"` //标签类型
	PrinterID           string             `gorm:"size:36;comment:打印机ID"`
	Printer             *Printer           `gorm:"constraint:OnDelete:CASCADE"` //打印机
	RemoteServiceTaskID *string            `gorm:"size:36;comment:远程任务ID"`
	RemoteServiceTask   *RemoteServiceTask `gorm:"constraint:OnDelete:SET NULL"` //远程任务
	Remark              string             `gorm:"size:500;comment:备注"`
}

func PBToLabelPrintTasks(in []*proto.LabelPrintTaskInfo) []*LabelPrintTask {
	var result []*LabelPrintTask
	for _, c := range in {
		result = append(result, PBToLabelPrintTask(c))
	}
	return result
}

func PBToLabelPrintTask(in *proto.LabelPrintTaskInfo) *LabelPrintTask {
	if in == nil {
		return nil
	}

	var productionLineID, productCategoryID, remoteServiceTaskID *string
	if in.ProductionLineID != "" {
		productionLineID = &in.ProductionLineID
	}
	if in.ProductCategoryID != "" {
		productCategoryID = &in.ProductCategoryID
	}
	if in.RemoteServiceTaskID != "" {
		remoteServiceTaskID = &in.RemoteServiceTaskID
	}

	return &LabelPrintTask{
		ModelID:             ModelID{ID: in.Id},
		Code:                in.Code,
		Description:         in.Description,
		Enable:              in.Enable,
		PrintCopies:         in.PrintCopies,
		ProductionLineID:    productionLineID,
		ProductCategoryID:   productCategoryID,
		LabelTypeID:         in.LabelTypeID,
		PrinterID:           in.PrinterID,
		RemoteServiceTaskID: remoteServiceTaskID,
		Remark:              in.Remark,
	}
}

func LabelPrintTasksToPB(in []*LabelPrintTask) []*proto.LabelPrintTaskInfo {
	var list []*proto.LabelPrintTaskInfo
	for _, f := range in {
		list = append(list, LabelPrintTaskToPB(f))
	}
	return list
}

func LabelPrintTaskToPB(in *LabelPrintTask) *proto.LabelPrintTaskInfo {
	if in == nil {
		return nil
	}

	var productionLineID, productCategoryID, remoteServiceTaskID string
	if in.ProductionLineID != nil {
		productionLineID = *in.ProductionLineID
	}
	if in.ProductCategoryID != nil {
		productCategoryID = *in.ProductCategoryID
	}
	if in.RemoteServiceTaskID != nil {
		remoteServiceTaskID = *in.RemoteServiceTaskID
	}

	m := &proto.LabelPrintTaskInfo{
		Id:                  in.ID,
		Code:                in.Code,
		Description:         in.Description,
		Enable:              in.Enable,
		PrintCopies:         in.PrintCopies,
		ProductionLineID:    productionLineID,
		ProductionLine:      ProductionLineToPB(in.ProductionLine),
		ProductCategoryID:   productCategoryID,
		ProductCategory:     ProductCategoryToPB(in.ProductCategory),
		LabelTypeID:         in.LabelTypeID,
		LabelType:           LabelTypeToPB(in.LabelType),
		PrinterID:           in.PrinterID,
		Printer:             PrinterToPB(in.Printer),
		RemoteServiceTaskID: remoteServiceTaskID,
		RemoteServiceTask:   RemoteServiceTaskToPB(in.RemoteServiceTask),
		Remark:              in.Remark,
	}
	return m
}
