package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 产品工单工步
type ProductOrderProcessStep struct {
	ModelID
	SortIndex                             int32                                   `json:"sortIndex" gorm:"comment:排序"`
	WorkDescription                       string                                  `json:"workDescription" gorm:"comment:作业描述"`
	WorkGraphic                           string                                  `json:"workGraphic" gorm:"comment:作业图示"`
	CreateTime                            time.Time                               `json:"createTime" gorm:"autoCreateTime:nano;comment:创建时间"`
	CreateUserID                          string                                  `json:"createUserID" gorm:"size:36;comment:创建人员ID"`
	Remark                                string                                  `json:"remark" gorm:"comment:备注"`
	ProcessStepTypeID                     string                                  `json:"processStepTypeID" gorm:"size:36;comment:工序类型ID"`
	ProcessStepType                       *ProcessStepType                        `json:"processStepType" gorm:"constraint:OnDelete:CASCADE"` //工序类型
	ProductOrderProcessID                 string                                  `json:"productOrderProcessID" gorm:"index;size:36;comment:工单工序ID"`
	ProductOrderProcess                   *ProductOrderProcess                    `json:"productOrderProcess" gorm:"constraint:OnDelete:CASCADE"`
	ProductOrderProcessStepAttachments    []*ProductOrderProcessStepAttachment    `json:"productOrderProcessStepAttachments" gorm:"constraint:OnDelete:CASCADE;"`
	ProductOrderProcessStepTypeParameters []*ProductOrderProcessStepTypeParameter `json:"productOrderProcessStepTypeParameters" gorm:"constraint:OnDelete:CASCADE;"`
}

// 产品工单工步附件
type ProductOrderProcessStepAttachment struct {
	ModelID
	ProductOrderProcessStepID string    `json:"productOrderProcessStepID" gorm:"index;size:36;comment:产品工单工步ID"`
	FileName                  string    `json:"fileName" gorm:"size:50;comment:文件名"`
	FileType                  string    `json:"fileType" gorm:"size:50;comment:文件类型"`
	FileSize                  float32   `json:"fileSize" gorm:"comment:文件大小"`
	SavedFileName             string    `json:"savedFileName" gorm:"size:100;comment:保存文件名"`
	FileStatus                string    `json:"fileStatus" gorm:"size:50;comment:文件状态"`
	CreateTime                time.Time `json:"createTime" gorm:"autoCreateTime:nano;comment:创建时间"`
	CreateUserID              string    `json:"createUserID" gorm:"size:36;comment:创建人员ID"`
}

// 产品工单工步参数
type ProductOrderProcessStepTypeParameter struct {
	ModelID
	ProductOrderProcessStepID  string                    `json:"productOrderProcessStepID" gorm:"index;size:36;comment:产品工单工步ID"`
	Value                      string                    `json:"value" gorm:"size:50;comment:参数值"`
	Remark                     string                    `json:"remark" gorm:"size:500;comment:备注"`
	ProcessStepTypeParameterID string                    `json:"processStepTypeParameterID" gorm:"size:36;comment:工步类型参数ID"`
	ProcessStepTypeParameter   *ProcessStepTypeParameter `gorm:"constraint:OnDelete:CASCADE"`
}

func PBToProductOrderProcessSteps(in []*proto.ProductOrderProcessStepInfo) []*ProductOrderProcessStep {
	var result []*ProductOrderProcessStep
	for _, c := range in {
		result = append(result, PBToProductOrderProcessStep(c))
	}
	return result
}

func PBToProductOrderProcessStep(in *proto.ProductOrderProcessStepInfo) *ProductOrderProcessStep {
	if in == nil {
		return nil
	}

	return &ProductOrderProcessStep{
		ModelID:         ModelID{ID: in.Id},
		SortIndex:       in.SortIndex,
		WorkDescription: in.WorkDescription,
		WorkGraphic:     in.WorkGraphic,
		// CreateTime:                            utils.ParseTime(in.CreateTime),
		CreateUserID:                          in.CreateUserID,
		Remark:                                in.Remark,
		ProcessStepTypeID:                     in.ProcessStepTypeID,
		ProductOrderProcessID:                 in.ProductOrderProcessID,
		ProductOrderProcessStepAttachments:    PBToProductOrderProcessStepAttachments(in.ProductOrderProcessStepAttachments),
		ProductOrderProcessStepTypeParameters: PBToProductOrderProcessStepTypeParameters(in.ProductOrderProcessStepTypeParameters),
	}
}

func ProductOrderProcessStepsToPB(in []*ProductOrderProcessStep) []*proto.ProductOrderProcessStepInfo {
	var list []*proto.ProductOrderProcessStepInfo
	for _, f := range in {
		list = append(list, ProductOrderProcessStepToPB(f))
	}
	return list
}

func ProductOrderProcessStepToPB(in *ProductOrderProcessStep) *proto.ProductOrderProcessStepInfo {
	if in == nil {
		return nil
	}

	m := &proto.ProductOrderProcessStepInfo{
		Id:                                    in.ID,
		SortIndex:                             in.SortIndex,
		WorkDescription:                       in.WorkDescription,
		WorkGraphic:                           in.WorkGraphic,
		CreateTime:                            utils.FormatTime(in.CreateTime),
		CreateUserID:                          in.CreateUserID,
		Remark:                                in.Remark,
		ProcessStepTypeID:                     in.ProcessStepTypeID,
		ProcessStepType:                       ProcessStepTypeToPB(in.ProcessStepType),
		ProductOrderProcessID:                 in.ProductOrderProcessID,
		ProductOrderProcessStepAttachments:    ProductOrderProcessStepAttachmentsToPB(in.ProductOrderProcessStepAttachments),
		ProductOrderProcessStepTypeParameters: ProductOrderProcessStepTypeParametersToPB(in.ProductOrderProcessStepTypeParameters),
	}
	return m
}

func PBToProductOrderProcessStepAttachments(in []*proto.ProductOrderProcessStepAttachmentInfo) []*ProductOrderProcessStepAttachment {
	var result []*ProductOrderProcessStepAttachment
	for _, c := range in {
		result = append(result, PBToProductOrderProcessStepAttachment(c))
	}
	return result
}

func PBToProductOrderProcessStepAttachment(in *proto.ProductOrderProcessStepAttachmentInfo) *ProductOrderProcessStepAttachment {
	if in == nil {
		return nil
	}

	return &ProductOrderProcessStepAttachment{
		ModelID:       ModelID{ID: in.Id},
		FileName:      in.FileName,
		FileType:      in.FileType,
		FileSize:      in.FileSize,
		SavedFileName: in.SavedFileName,
		FileStatus:    in.FileStatus,
		// CreateTime:    utils.ParseTime(in.CreateTime),
		CreateUserID: in.CreateUserID,
		// ProductOrderProcessStepID: in.ProductOrderProcessStepID,
	}
}

func ProductOrderProcessStepAttachmentsToPB(in []*ProductOrderProcessStepAttachment) []*proto.ProductOrderProcessStepAttachmentInfo {
	var list []*proto.ProductOrderProcessStepAttachmentInfo
	for _, f := range in {
		list = append(list, ProductOrderProcessStepAttachmentToPB(f))
	}
	return list
}

func ProductOrderProcessStepAttachmentToPB(in *ProductOrderProcessStepAttachment) *proto.ProductOrderProcessStepAttachmentInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductOrderProcessStepAttachmentInfo{
		Id:                        in.ID,
		FileName:                  in.FileName,
		FileType:                  in.FileType,
		FileSize:                  in.FileSize,
		SavedFileName:             in.SavedFileName,
		FileStatus:                in.FileStatus,
		CreateTime:                utils.FormatTime(in.CreateTime),
		CreateUserID:              in.CreateUserID,
		ProductOrderProcessStepID: in.ProductOrderProcessStepID,
	}
	return m
}

func PBToProductOrderProcessStepTypeParameters(in []*proto.ProductOrderProcessStepTypeParameterInfo) []*ProductOrderProcessStepTypeParameter {
	var result []*ProductOrderProcessStepTypeParameter
	for _, c := range in {
		result = append(result, PBToProductOrderProcessStepTypeParameter(c))
	}
	return result
}

func PBToProductOrderProcessStepTypeParameter(in *proto.ProductOrderProcessStepTypeParameterInfo) *ProductOrderProcessStepTypeParameter {
	if in == nil {
		return nil
	}
	return &ProductOrderProcessStepTypeParameter{
		ModelID: ModelID{ID: in.Id},
		Value:   in.Value,
		Remark:  in.Remark,
		// ProductOrderProcessStepID:  in.ProductOrderProcessStepID,
		ProcessStepTypeParameterID: in.ProcessStepTypeParameterID,
	}
}

func ProductOrderProcessStepTypeParametersToPB(in []*ProductOrderProcessStepTypeParameter) []*proto.ProductOrderProcessStepTypeParameterInfo {
	var list []*proto.ProductOrderProcessStepTypeParameterInfo
	for _, f := range in {
		list = append(list, ProductOrderProcessStepTypeParameterToPB(f))
	}
	return list
}

func ProductOrderProcessStepTypeParameterToPB(in *ProductOrderProcessStepTypeParameter) *proto.ProductOrderProcessStepTypeParameterInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductOrderProcessStepTypeParameterInfo{
		Id:                         in.ID,
		Value:                      in.Value,
		Remark:                     in.Remark,
		ProductOrderProcessStepID:  in.ProductOrderProcessStepID,
		ProcessStepTypeParameterID: in.ProcessStepTypeParameterID,
	}
	return m
}
