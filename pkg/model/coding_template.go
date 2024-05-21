package model

import (
	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 编码模版
type CodingTemplate struct {
	ModelID
	Code              string              `gorm:"index;size:50;comment:代号"`
	IndexType         int32               `gorm:"comment:索引类型"`
	IndexBits         int32               `gorm:"comment:索引位数"`
	Description       string              `gorm:"size:500;comment:描述"`
	Remark            string              `gorm:"size:500;comment:备注"`
	CodingElements    []*CodingElement    `gorm:"constraint:OnDelete:CASCADE"` //comment:编码元素
	CodingGenerations []*CodingGeneration `gorm:"constraint:OnDelete:CASCADE"` //comment:编码序列记录
}

// 编码元素
type CodingElement struct {
	ModelID
	CodingTemplateID    string                `gorm:"index;size:36;comment:编码模板ID"`
	SortIndex           int32                 `gorm:"comment:顺序"`
	Name                string                `gorm:"size:50;comment:名称"`
	ElementType         int32                 `gorm:"comment:元素类型"`
	DataBits            int32                 `gorm:"comment:数据位数"`
	DefaultValue        string                `gorm:"size:100;comment:预设值"`
	PlaceHolder         string                `gorm:"size:1;comment:占位符"`
	Remark              string                `gorm:"size:500;comment:备注"`
	CodingElementValues []*CodingElementValue `gorm:"constraint:OnDelete:CASCADE"` //编码元素值
}

// 编码元素值
type CodingElementValue struct {
	ModelID
	CodingElementID string `gorm:"index;size:36;comment:编码元素ID"`
	Value           string `gorm:"size:10;comment:值"`
	Description     string `gorm:"size:500;comment:描述"`
	// CodingElementValueConvertions []*CodingElementValueConvertion //转换方式
	// CodingElementValueLimitations []*CodingElementValueLimitation //限定规则
}

func PBToCodingTemplates(in []*proto.CodingTemplateInfo) []*CodingTemplate {
	var result []*CodingTemplate
	for _, c := range in {
		result = append(result, PBToCodingTemplate(c))
	}
	return result
}

func PBToCodingTemplate(in *proto.CodingTemplateInfo) *CodingTemplate {
	if in == nil {
		return nil
	}
	return &CodingTemplate{
		ModelID:        ModelID{ID: in.Id},
		Code:           in.Code,
		IndexType:      in.IndexType,
		IndexBits:      in.IndexBits,
		Description:    in.Description,
		Remark:         in.Remark,
		CodingElements: PBToCodingElements(in.CodingElements),
	}
}

func CodingTemplatesToPB(in []*CodingTemplate) []*proto.CodingTemplateInfo {
	var list []*proto.CodingTemplateInfo
	for _, f := range in {
		list = append(list, CodingTemplateToPB(f))
	}
	return list
}

func CodingTemplateToPB(in *CodingTemplate) *proto.CodingTemplateInfo {
	if in == nil {
		return nil
	}
	m := &proto.CodingTemplateInfo{
		Id:             in.ID,
		Code:           in.Code,
		IndexType:      in.IndexType,
		IndexBits:      in.IndexBits,
		Description:    in.Description,
		Remark:         in.Remark,
		CodingElements: CodingElementsToPB(in.CodingElements),
	}
	return m
}

func PBToCodingElements(in []*proto.CodingElementInfo) []*CodingElement {
	var result []*CodingElement
	for _, c := range in {
		result = append(result, PBToCodingElement(c))
	}
	return result
}

func PBToCodingElement(in *proto.CodingElementInfo) *CodingElement {
	if in == nil {
		return nil
	}
	return &CodingElement{
		ModelID:             ModelID{ID: in.Id},
		SortIndex:           in.SortIndex,
		Name:                in.Name,
		ElementType:         in.ElementType,
		DataBits:            in.DataBits,
		DefaultValue:        in.DefaultValue,
		PlaceHolder:         in.PlaceHolder,
		Remark:              in.Remark,
		CodingElementValues: PBToCodingElementValues(in.CodingElementValues),
	}
}

func CodingElementsToPB(in []*CodingElement) []*proto.CodingElementInfo {
	var list []*proto.CodingElementInfo
	for _, f := range in {
		list = append(list, CodingElementToPB(f))
	}
	return list
}

func CodingElementToPB(in *CodingElement) *proto.CodingElementInfo {
	if in == nil {
		return nil
	}
	m := &proto.CodingElementInfo{
		Id:                  in.ID,
		CodingTemplateID:    in.CodingTemplateID,
		SortIndex:           in.SortIndex,
		Name:                in.Name,
		ElementType:         in.ElementType,
		DataBits:            in.DataBits,
		DefaultValue:        in.DefaultValue,
		PlaceHolder:         in.PlaceHolder,
		Remark:              in.Remark,
		CodingElementValues: CodingElementValuesToPB(in.CodingElementValues),
	}
	return m
}

func PBToCodingElementValues(in []*proto.CodingElementValueInfo) []*CodingElementValue {
	var result []*CodingElementValue
	for _, c := range in {
		result = append(result, PBToCodingElementValue(c))
	}
	return result
}

func PBToCodingElementValue(in *proto.CodingElementValueInfo) *CodingElementValue {
	if in == nil {
		return nil
	}
	return &CodingElementValue{
		ModelID:     ModelID{ID: in.Id},
		Value:       in.Value,
		Description: in.Description,
	}
}

func CodingElementValuesToPB(in []*CodingElementValue) []*proto.CodingElementValueInfo {
	var list []*proto.CodingElementValueInfo
	for _, f := range in {
		list = append(list, CodingElementValueToPB(f))
	}
	return list
}

func CodingElementValueToPB(in *CodingElementValue) *proto.CodingElementValueInfo {
	if in == nil {
		return nil
	}
	m := &proto.CodingElementValueInfo{
		Id:              in.ID,
		CodingElementID: in.CodingElementID,
		Value:           in.Value,
		Description:     in.Description,
	}
	return m
}

// 转换方式
// type CodingElementValueConvertion struct {
// 	ModelID
// 	Value                string              `gorm:"size:100;comment:设定值"`
// 	Description          string              `gorm:"size:500;comment:值描述"`
// 	ProductAttributeID   string              `gorm:"size:36;comment:转换特性ID"`
// 	ProductAttribute     *ProductAttribute   `gorm:""` //转换特性
// 	CodingElementValueID string              `gorm:"size:36;comment:编码元素值ID"`
// 	CodingElementValue   *CodingElementValue `gorm:""` //编码元素值关联
// }

// 限定规则
// type CodingElementValueLimitation struct {
// 	ModelID
// 	CodingElementID      string                         `gorm:"size:36;comment:限定元素ID"`
// 	CodingElement        *CodingElement                 `` //限定元素
// 	Value                string                         `gorm:"size:100;comment:限定值"`
// 	CodingElementValueID string                         `gorm:"size:36;comment:归属编码元素值ID"`
// 	CodingElementValue   *CodingElementValue            `` //归属编码元素值
// 	ParentID             string                         `gorm:"size:36;comment:串联限定父ID"`
// 	Parent               *CodingElementValueLimitation  ``         //串联限定父级关联
// 	Children             []CodingElementValueLimitation `gorm:"-"` //串联限定子项集合
// 	TreeLevel            int32                          `gorm:"-"`
// 	IsTreeLeaf           bool                           `gorm:"-"`
// }
