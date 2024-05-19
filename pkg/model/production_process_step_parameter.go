package model

// 生产工步参数
type ProductionProcessStepParameter struct {
	ModelID
	Code                                 string
	Priority                             int32                                  `gorm:"comment:优先级"`
	Description                          string                                 `gorm:"size:500;comment:描述"`
	Enable                               bool                                   `gorm:"comment:是否启用"`
	Remark                               string                                 `gorm:"size:500;comment:备注"`
	ProductionLineID                     string                                 `gorm:"size:36;comment:产线ID"`
	ProductionLine                       *ProductionLine                        ``                                   //生产产线
	ProductionProcessStepParameterValues []*ProductionProcessStepParameterValue `gorm:"constraint:OnDelete:CASCADE"` //生产工步参数值
	AttributeExpression                  []*AttributeExpression                 `gorm:"polymorphic:Rule;polymorphicValue:ProductionProcessStepParameter"`
}

// 生产工步参数值
type ProductionProcessStepParameterValue struct {
	ModelID
	ProductionProcessStepParameterID string                 `gorm:"index;size:36;comment:生产工步参数ID"`
	ProductionProcessID              string                 `gorm:"size:36;comment:生产工序ID"`
	ProductionProcessStepID          string                 `gorm:"size:36;comment:生产工步ID"`
	ProductionProcessStep            *ProductionProcessStep `` //生产工步
	ProcessStepTypeParameterID       string                 `gorm:"size:36;comment:生产工步类型参数ID"`
	StandardValue                    string                 `gorm:"size:100;comment:标准值"`
	MaximumValue                     string                 `gorm:"size:100;comment:最大值"`
	MinimumValue                     string                 `gorm:"size:100;comment:最小值"`
}
