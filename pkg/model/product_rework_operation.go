package model

import "time"

// 产品返工操作
type ProductReworkOperation struct {
	ModelID
	CreateTime            time.Time            `gorm:"autoCreateTime:nano;comment:创建时间"`
	ProductReworkRecordID string               `gorm:"size:36;comment:返工记录ID"`
	ProductReworkRecord   *ProductReworkRecord `gorm:"constraint:OnDelete:CASCADE"`
	ProductOrderBomID     string               `gorm:"size:36;comment:返工记录ID"`
	ProductOrderBom       *ProductOrderBom     `gorm:"constraint:OnDelete:CASCADE"`
	OperationMode         string               `gorm:"size:50;comment:操作方式"`
}
