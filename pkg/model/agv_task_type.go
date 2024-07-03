package model

import "github.com/CloudSilk/CloudSilk/pkg/proto"

type AGVTaskType struct {
	ModelID
	Code                    string                 `gorm:"size:50;comment:代号"`
	Description             string                 `gorm:"size:500;comment:描述"`
	ShelfType               int32                  `gorm:"comment:货架类型"`
	SceneType               int32                  `gorm:"comment:场景类型"`
	PositionCode            string                 `gorm:"size:50;comment:任务点位"`
	Remark                  string                 `gorm:"size:500;comment:备注"`
	MaterialContainerTypeID *string                `gorm:"size:36;comment:容器类型ID"`
	MaterialContainerType   *MaterialContainerType `gorm:"constraint:OnDelete:SET NULL"` //容器类型
}

func PBToAGVTaskTypes(in []*proto.AGVTaskTypeInfo) []*AGVTaskType {
	var result []*AGVTaskType
	for _, c := range in {
		result = append(result, PBToAGVTaskType(c))
	}
	return result
}

func PBToAGVTaskType(in *proto.AGVTaskTypeInfo) *AGVTaskType {
	if in == nil {
		return nil
	}

	var materialContainerTypeID *string
	if in.MaterialContainerTypeID != "" {
		materialContainerTypeID = &in.MaterialContainerTypeID
	}

	return &AGVTaskType{
		ModelID:                 ModelID{ID: in.Id},
		Code:                    in.Code,
		Description:             in.Description,
		ShelfType:               in.ShelfType,
		SceneType:               in.SceneType,
		PositionCode:            in.PositionCode,
		Remark:                  in.Remark,
		MaterialContainerTypeID: materialContainerTypeID,
	}
}

func AGVTaskTypesToPB(in []*AGVTaskType) []*proto.AGVTaskTypeInfo {
	var list []*proto.AGVTaskTypeInfo
	for _, f := range in {
		list = append(list, AGVTaskTypeToPB(f))
	}
	return list
}

func AGVTaskTypeToPB(in *AGVTaskType) *proto.AGVTaskTypeInfo {
	if in == nil {
		return nil
	}

	var materialContainerTypeID string
	if in.MaterialContainerTypeID != nil {
		materialContainerTypeID = *in.MaterialContainerTypeID
	}

	m := &proto.AGVTaskTypeInfo{
		Id:                      in.ID,
		Code:                    in.Code,
		Description:             in.Description,
		ShelfType:               in.ShelfType,
		SceneType:               in.SceneType,
		PositionCode:            in.PositionCode,
		Remark:                  in.Remark,
		MaterialContainerTypeID: materialContainerTypeID,
		MaterialContainerType:   MaterialContainerTypeToPB(in.MaterialContainerType),
	}
	return m
}
