package model

// 产品返工路线
type ProductReworkRoute struct {
	ModelID
	ProductionLineID   string             `gorm:"size:36;comment:生产产线ID"`
	ProductionLine     *ProductionLine    `gorm:"constraint:OnDelete:CASCADE"`
	MaterialCategoryID string             `gorm:"size:36;comment:物料类别ID"`
	MaterialCategory   *MaterialCategory  `gorm:"constraint:OnDelete:CASCADE"`
	FollowProcessID    string             `gorm:"size:36;comment:后续工序ID"`
	FollowProcess      *ProductionProcess `gorm:"constraint:OnDelete:CASCADE"`
}
