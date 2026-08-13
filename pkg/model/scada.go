package model

import (
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 采集设备（SCADA 网关连接的 PLC/仪表）
type ScadaDevice struct {
	ModelID
	Code             string     `json:"code" gorm:"index;size:100;comment:代号"`
	Name             string     `json:"name" gorm:"size:200;comment:名称"`
	Protocol         string     `json:"protocol" gorm:"size:50;comment:协议"`
	Address          string     `json:"address" gorm:"size:100;comment:地址"`
	SlaveID          int32      `json:"slaveID" gorm:"comment:从站号"`
	IntervalSeconds  int32      `json:"intervalSeconds" gorm:"comment:采集间隔(秒)"`
	ConnectionState  string     `json:"connectionState" gorm:"size:50;comment:连接状态"`
	LastCollectTime  *time.Time `json:"lastCollectTime" gorm:"comment:最近采集时间"`
	LastError        string     `json:"lastError" gorm:"size:500;comment:最近错误"`
	Enable           bool       `json:"enable" gorm:"default:true;comment:是否启用"`
	Remark           string     `json:"remark" gorm:"size:1000;comment:备注"`
	Tags             []*ScadaTag `json:"tags" gorm:"constraint:OnDelete:CASCADE"`
}

// 采集点位
type ScadaTag struct {
	ModelID
	ScadaDeviceID string      `json:"scadaDeviceID" gorm:"index;size:36;comment:设备ID"`
	ScadaDevice   *ScadaDevice `json:"scadaDevice" gorm:"constraint:OnDelete:CASCADE"`
	Name          string      `json:"name" gorm:"size:100;comment:点位名称"`
	Address       int32       `json:"address" gorm:"comment:寄存器地址"`
	FunctionCode  int32       `json:"functionCode" gorm:"comment:功能码"`
	DataType      string      `json:"dataType" gorm:"size:20;comment:数据类型"`
	Scale         float64     `json:"scale" gorm:"comment:缩放系数"`
	Unit          string      `json:"unit" gorm:"size:50;comment:单位"`
	SaveHistory   bool        `json:"saveHistory" gorm:"default:true;comment:是否存历史"`
	Enable        bool        `json:"enable" gorm:"default:true;comment:是否启用"`
	Remark        string      `json:"remark" gorm:"size:1000;comment:备注"`
}

// 点位实时值（每个点位一行）
type ScadaTagValue struct {
	ModelID
	ScadaTagID  string     `json:"scadaTagID" gorm:"index;size:36;comment:点位ID"`
	ScadaTag    *ScadaTag  `json:"scadaTag" gorm:"constraint:OnDelete:CASCADE"`
	Value       string     `json:"value" gorm:"size:100;comment:值"`
	CollectTime *time.Time `json:"collectTime" gorm:"comment:采集时间"`
	Quality     string     `json:"quality" gorm:"size:20;comment:质量"`
}

// 点位历史值
type ScadaTagHistory struct {
	ModelID
	ScadaTagID  string    `json:"scadaTagID" gorm:"index;size:36;comment:点位ID"`
	Value       string    `json:"value" gorm:"size:100;comment:值"`
	CollectTime time.Time `json:"collectTime" gorm:"autoCreateTime:nano;comment:采集时间"`
}

func PBToScadaDevices(in []*proto.ScadaDeviceInfo) []*ScadaDevice {
	var result []*ScadaDevice
	for _, c := range in {
		result = append(result, PBToScadaDevice(c))
	}
	return result
}

func PBToScadaDevice(in *proto.ScadaDeviceInfo) *ScadaDevice {
	if in == nil {
		return nil
	}
	return &ScadaDevice{
		ModelID:         ModelID{ID: in.Id},
		Code:            in.Code,
		Name:            in.Name,
		Protocol:        in.Protocol,
		Address:         in.Address,
		SlaveID:         in.SlaveID,
		IntervalSeconds: in.IntervalSeconds,
		ConnectionState: in.ConnectionState,
		LastError:       in.LastError,
		Enable:          in.Enable,
		Remark:          in.Remark,
	}
}

func ScadaDevicesToPB(in []*ScadaDevice) []*proto.ScadaDeviceInfo {
	var list []*proto.ScadaDeviceInfo
	for _, f := range in {
		list = append(list, ScadaDeviceToPB(f))
	}
	return list
}

func ScadaDeviceToPB(in *ScadaDevice) *proto.ScadaDeviceInfo {
	if in == nil {
		return nil
	}
	m := &proto.ScadaDeviceInfo{
		Id:              in.ID,
		Code:            in.Code,
		Name:            in.Name,
		Protocol:        in.Protocol,
		Address:         in.Address,
		SlaveID:         in.SlaveID,
		IntervalSeconds: in.IntervalSeconds,
		ConnectionState: in.ConnectionState,
		LastError:       in.LastError,
		Enable:          in.Enable,
		Remark:          in.Remark,
	}
	if in.LastCollectTime != nil {
		m.LastCollectTime = in.LastCollectTime.Format("2006-01-02 15:04:05")
	}
	return m
}

func PBToScadaTags(in []*proto.ScadaTagInfo) []*ScadaTag {
	var result []*ScadaTag
	for _, c := range in {
		result = append(result, PBToScadaTag(c))
	}
	return result
}

func PBToScadaTag(in *proto.ScadaTagInfo) *ScadaTag {
	if in == nil {
		return nil
	}
	return &ScadaTag{
		ModelID:      ModelID{ID: in.Id},
		ScadaDeviceID: in.ScadaDeviceID,
		Name:         in.Name,
		Address:      in.Address,
		FunctionCode: in.FunctionCode,
		DataType:     in.DataType,
		Scale:        in.Scale,
		Unit:         in.Unit,
		SaveHistory:  in.SaveHistory,
		Enable:       in.Enable,
		Remark:       in.Remark,
	}
}

func ScadaTagsToPB(in []*ScadaTag) []*proto.ScadaTagInfo {
	var list []*proto.ScadaTagInfo
	for _, f := range in {
		list = append(list, ScadaTagToPB(f))
	}
	return list
}

func ScadaTagToPB(in *ScadaTag) *proto.ScadaTagInfo {
	if in == nil {
		return nil
	}
	m := &proto.ScadaTagInfo{
		Id:           in.ID,
		ScadaDeviceID: in.ScadaDeviceID,
		Name:         in.Name,
		Address:      in.Address,
		FunctionCode:  in.FunctionCode,
		DataType:     in.DataType,
		Scale:        in.Scale,
		Unit:         in.Unit,
		SaveHistory:  in.SaveHistory,
		Enable:       in.Enable,
		Remark:       in.Remark,
	}
	if in.ScadaDevice != nil {
		m.ScadaDeviceCode = in.ScadaDevice.Code
	}
	return m
}

func ScadaTagValuesToPB(in []*ScadaTagValue) []*proto.ScadaTagValueInfo {
	var list []*proto.ScadaTagValueInfo
	for _, f := range in {
		list = append(list, ScadaTagValueToPB(f))
	}
	return list
}

func ScadaTagValueToPB(in *ScadaTagValue) *proto.ScadaTagValueInfo {
	if in == nil {
		return nil
	}
	m := &proto.ScadaTagValueInfo{
		ScadaTagID: in.ScadaTagID,
		Value:      in.Value,
		Quality:    in.Quality,
	}
	if in.CollectTime != nil {
		m.CollectTime = in.CollectTime.Format("2006-01-02 15:04:05")
	}
	if in.ScadaTag != nil {
		m.TagName = in.ScadaTag.Name
		m.Unit = in.ScadaTag.Unit
	}
	return m
}

func ScadaTagHistoriesToPB(in []*ScadaTagHistory) []*proto.ScadaTagHistoryInfo {
	var list []*proto.ScadaTagHistoryInfo
	for _, f := range in {
		list = append(list, &proto.ScadaTagHistoryInfo{
			Id:          f.ID,
			ScadaTagID:  f.ScadaTagID,
			Value:       f.Value,
			CollectTime: f.CollectTime.Format("2006-01-02 15:04:05"),
		})
	}
	return list
}
