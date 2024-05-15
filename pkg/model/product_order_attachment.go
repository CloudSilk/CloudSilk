package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 产品工单附件
type ProductOrderAttachment struct {
	ModelID
	FileName       string        `json:"fileName" gorm:"size:100;comment:文件名"`
	FileType       string        `json:"fileType" gorm:"size:100;comment:文件类型"`
	FileSize       float32       `json:"fileSize" gorm:"comment:文件大小"`
	SavedFileName  string        `json:"savedFileName" gorm:"size:200;comment:保存文件名"`
	FileStatus     string        `json:"fileStatus" gorm:"size:100;comment:文件状态"`
	CreateTime     time.Time     `json:"createTime" gorm:"autoCreateTime:nano;comment:创建时间"`
	CreateUserID   string        `json:"createUserID" gorm:"size:36;comment:创建人员ID"`
	ProductOrderID string        `json:"productOrderID" gorm:"size:36;comment:隶属工单ID"`
	ProductOrder   *ProductOrder `json:"productOrder"`
}

func PBToProductOrderAttachments(in []*proto.ProductOrderAttachmentInfo) []*ProductOrderAttachment {
	var result []*ProductOrderAttachment
	for _, c := range in {
		result = append(result, PBToProductOrderAttachment(c))
	}
	return result
}

func PBToProductOrderAttachment(in *proto.ProductOrderAttachmentInfo) *ProductOrderAttachment {
	if in == nil {
		return nil
	}
	return &ProductOrderAttachment{
		ModelID:       ModelID{ID: in.Id},
		FileName:      in.FileName,
		FileType:      in.FileType,
		FileSize:      in.FileSize,
		SavedFileName: in.SavedFileName,
		FileStatus:    in.FileStatus,
		// CreateTime:     utils.ParseSqlNullTime(in.CreateTime),
		CreateUserID:   in.CreateUserID,
		ProductOrderID: in.ProductOrderID,
	}
}

func ProductOrderAttachmentsToPB(in []*ProductOrderAttachment) []*proto.ProductOrderAttachmentInfo {
	var list []*proto.ProductOrderAttachmentInfo
	for _, f := range in {
		list = append(list, ProductOrderAttachmentToPB(f))
	}
	return list
}

func ProductOrderAttachmentToPB(in *ProductOrderAttachment) *proto.ProductOrderAttachmentInfo {
	if in == nil {
		return nil
	}
	m := &proto.ProductOrderAttachmentInfo{
		Id:             in.ID,
		FileName:       in.FileName,
		FileType:       in.FileType,
		FileSize:       in.FileSize,
		SavedFileName:  in.SavedFileName,
		FileStatus:     in.FileStatus,
		CreateTime:     utils.FormatTime(in.CreateTime),
		CreateUserID:   in.CreateUserID,
		ProductOrderID: in.ProductOrderID,
	}
	return m
}
