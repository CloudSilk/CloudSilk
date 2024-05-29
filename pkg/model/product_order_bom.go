package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
)

// 产品工单BOM
type ProductOrderBom struct {
	ModelID
	ItemNo              string        `json:"itemNo" gorm:"size:100;comment:项目号"`                 //项目号
	MaterialNo          string        `json:"materialNo" gorm:"size:100;comment:物料号"`             //物料号
	MaterialDescription string        `json:"materialDescription" gorm:"size:1000;comment:物料描述"`  //物料描述
	PieceQTY            float32       `json:"pieceQTY" gorm:"comment:单件数量"`                       //单件数量
	RequireQTY          float32       `json:"requireQTY" gorm:"comment:需求数量"`                     //需求数量
	Unit                string        `json:"unit" gorm:"size:20;comment:单位"`                     //单位
	EnableControl       bool          `json:"enableControl" gorm:"comment:是否管控"`                  //是否管控
	ControlType         int32         `json:"controlType" gorm:"comment:管控类型"`                    //管控类型
	ProductionProcess   string        `json:"productionProcess" gorm:"size:20;comment:需求工序"`      //需求工序
	Warehouse           string        `json:"warehouse" gorm:"size:100;comment:物料仓库"`             //物料仓库
	CreateTime          time.Time     `json:"createTime" gorm:"autoCreateTime:nano;comment:创建时间"` //创建时间
	CreateUserID        string        `json:"createUserID" gorm:"size:36;comment:创建人员ID"`         //创建人员ID
	Remark              string        `json:"remark" gorm:"size:1000;comment:备注"`                 //备注
	ProductOrderID      string        `json:"productOrderID" gorm:"size:36;comment:隶属工单ID"`       //隶属工单ID
	ProductOrder        *ProductOrder `json:"productOrder" gorm:"constraint:OnDelete:CASCADE"`    //工单
}

func PBToProductOrderBoms(in []*proto.ProductOrderBomInfo) []*ProductOrderBom {
	var result []*ProductOrderBom
	for _, c := range in {
		result = append(result, PBToProductOrderBom(c))
	}
	return result
}

func PBToProductOrderBom(in *proto.ProductOrderBomInfo) *ProductOrderBom {
	if in == nil {
		return nil
	}
	return &ProductOrderBom{
		ModelID:             ModelID{ID: in.Id},
		ItemNo:              in.ItemNo,
		MaterialNo:          in.MaterialNo,
		MaterialDescription: in.MaterialDescription,
		PieceQTY:            in.PieceQTY,
		RequireQTY:          in.RequireQTY,
		Unit:                in.Unit,
		EnableControl:       in.EnableControl,
		ControlType:         in.ControlType,
		ProductionProcess:   in.ProductionProcess,
		Warehouse:           in.Warehouse,
		// CreateTime:          utils.ParseTime(in.CreateTime),
		CreateUserID:   in.CreateUserID,
		Remark:         in.Remark,
		ProductOrderID: in.ProductOrderID,
	}
}

func ProductOrderBomsToPB(in []*ProductOrderBom) []*proto.ProductOrderBomInfo {
	var list []*proto.ProductOrderBomInfo
	for _, f := range in {
		list = append(list, ProductOrderBomToPB(f))
	}
	return list
}

func ProductOrderBomToPB(in *ProductOrderBom) *proto.ProductOrderBomInfo {
	if in == nil {
		return nil
	}
	productOrderNo := ""
	if in.ProductOrder != nil {
		productOrderNo = in.ProductOrder.ProductOrderNo
	}
	m := &proto.ProductOrderBomInfo{
		Id:                  in.ID,
		ItemNo:              in.ItemNo,
		MaterialNo:          in.MaterialNo,
		MaterialDescription: in.MaterialDescription,
		PieceQTY:            in.PieceQTY,
		RequireQTY:          in.RequireQTY,
		Unit:                in.Unit,
		EnableControl:       in.EnableControl,
		ControlType:         in.ControlType,
		ProductionProcess:   in.ProductionProcess,
		Warehouse:           in.Warehouse,
		CreateTime:          utils.FormatTime(in.CreateTime),
		CreateUserID:        in.CreateUserID,
		Remark:              in.Remark,
		ProductOrderID:      in.ProductOrderID,
		ProductOrderNo:      productOrderNo,
	}
	return m
}
