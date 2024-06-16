package model

import (
	"database/sql"
	"time"
)

// 材料退货申请表
type MaterialReturnRequestForm struct {
	ModelID
	FormNo              string                  `gorm:"size:50;comment:申请单号"`
	CreateTime          time.Time               `gorm:"autoCreateTime:nano;comment:创建时间"`
	CreateUserID        string                  `gorm:"size:36;comment:创建人员ID"`
	MaterialSupplierID  *string                 `gorm:"size:36;comment:物料供应商ID"`
	MaterialSupplier    *MaterialSupplier       `gorm:"constraint:OnDelete:SET NULL"`
	MaterialInfoID      string                  `gorm:"size:36;comment:物料信息ID"`
	MaterialInfo        *MaterialInfo           `gorm:"constraint:OnDelete:CASCADE"`
	ProductionStationID *string                 `gorm:"size:36;comment:发现工站ID"`
	ProductionStation   *ProductionStation      `gorm:"constraint:OnDelete:SET NULL"`
	MaterialTraceNo     string                  `gorm:"size:100;comment:物料追溯号"`
	Quantity            float64                 `gorm:"comment:退料数量"`
	ReturnID            string                  `gorm:"size:36;comment:退料来源ID"`
	ReturnSource        string                  `gorm:"size:50;comment:退料来源"`
	ReturnReason        string                  `gorm:"size:500;comment:退料原因"`
	ReturnTypeID        *string                 `gorm:"size:36;comment:退料类型ID"`
	ReturnType          *MaterialReturnType     `gorm:"constraint:OnDelete:SET NULL"`
	ReturnCauseID       *string                 `gorm:"size:36;comment:退料原因ID"`
	ReturnCause         *MaterialReturnCause    `gorm:"constraint:OnDelete:SET NULL"`
	ReturnSolutionID    *string                 `gorm:"size:36;comment:处理方案ID"`
	ReturnSolution      *MaterialReturnSolution `gorm:"constraint:OnDelete:SET NULL"`
	ReturnBrief         string                  `gorm:"size:500;comment:退料简述"`
	CheckTime           sql.NullTime            `gorm:"comment:复核时间"`
	CheckUserID         *string                 `gorm:"size:36;comment:复核人员ID"`
	HandleMethod        string                  `gorm:"size:50;comment:旧件处理"`
	CurrentState        int32                   `gorm:"comment:当前状态"`
	Remark              string                  `gorm:"size:500;comment:备注"`
}

// 退料类型
type MaterialReturnType struct {
	ModelID
	Code                 string                 `gorm:"size:50;comment:代号"`  //代号
	Description          string                 `gorm:"size:500;comment:描述"` //描述
	Remark               string                 `gorm:"size:500;comment:备注"` //备注
	MaterialReturnCauses []*MaterialReturnCause `gorm:"-"`                   //退料原因
}

// 退料原因
type MaterialReturnCause struct {
	ModelID
	Code                    string                    `gorm:"size:50;comment:代号"`  //代号
	Description             string                    `gorm:"size:500;comment:描述"` //描述
	Remark                  string                    `gorm:"size:500;comment:备注"` //备注
	MaterialReturnTypes     []*MaterialReturnType     `gorm:"-"`                   //归属类型
	MaterialReturnSolutions []*MaterialReturnSolution `gorm:"-"`                   //可用方案
	MaterialCategories      []*MaterialCategory       `gorm:"-"`                   //物料类别
}

// 可用方案
type MaterialReturnSolution struct {
	ModelID
	Code                 string                 `gorm:"size:50;comment:代号"`  //代号
	Description          string                 `gorm:"size:500;comment:描述"` //描述
	Remark               string                 `gorm:"size:500;comment:备注"` //备注
	MaterialReturnCauses []*MaterialReturnCause `gorm:"-"`                   //退料原因
}
