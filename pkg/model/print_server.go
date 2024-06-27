package model

import (
	apipb "github.com/CloudSilk/CloudSilk/pkg/proto"
)

type PrintServer struct {
	ModelID
	Name         string     `json:"name" gorm:"size:100;comment:名称"`
	Identity     string     `json:"identity" gorm:"size:100;comment:身份标识"`
	RunningState string     `json:"runningState" gorm:"size:100;comment:运行状态"`
	Printers     []*Printer `json:"printers" gorm:"constraint:OnDelete:CASCADE"` //打印机
}

type Printer struct {
	ModelID
	PrintServerID string       `json:"printServerID" gorm:"index;size:36;comment:打印服务器ID"`
	PrintServer   *PrintServer `gorm:"constraint:OnDelete:CASCADE"`
	Name          string       `json:"name" gorm:"size:200;comment:名称"`
	Enable        bool         `json:"enable" gorm:"comment:是否启用"`
}

func PBToPrintServers(in []*apipb.PrintServerInfo) []*PrintServer {
	var result []*PrintServer
	for _, c := range in {
		result = append(result, PBToPrintServer(c))
	}
	return result
}

func PBToPrintServer(in *apipb.PrintServerInfo) *PrintServer {
	if in == nil {
		return nil
	}
	return &PrintServer{
		ModelID:      ModelID{ID: in.Id},
		Name:         in.Name,
		Identity:     in.Identity,
		RunningState: in.RunningState,
		Printers:     PBToPrinters(in.Printers),
	}
}

func PBToPrinters(in []*apipb.PrinterInfo) []*Printer {
	var result []*Printer
	for _, c := range in {
		result = append(result, PBToPrinter(c))
	}
	return result
}

func PBToPrinter(in *apipb.PrinterInfo) *Printer {
	if in == nil {
		return nil
	}
	return &Printer{
		ModelID:       ModelID{ID: in.Id},
		Name:          in.Name,
		Enable:        in.Enable,
		PrintServerID: in.PrintServerID,
	}
}

func PrintServersToPB(in []*PrintServer) []*apipb.PrintServerInfo {
	var list []*apipb.PrintServerInfo
	for _, f := range in {
		list = append(list, PrintServerToPB(f))
	}
	return list
}

func PrintServerToPB(in *PrintServer) *apipb.PrintServerInfo {
	if in == nil {
		return nil
	}
	m := &apipb.PrintServerInfo{
		Printers:     PrintersToPB(in.Printers),
		Id:           in.ID,
		Name:         in.Name,
		Identity:     in.Identity,
		RunningState: in.RunningState,
	}
	return m
}

func PrintersToPB(in []*Printer) []*apipb.PrinterInfo {
	var list []*apipb.PrinterInfo
	for _, f := range in {
		list = append(list, PrinterToPB(f))
	}
	return list
}

func PrinterToPB(in *Printer) *apipb.PrinterInfo {
	if in == nil {
		return nil
	}
	m := &apipb.PrinterInfo{
		Id:            in.ID,
		Name:          in.Name,
		Enable:        in.Enable,
		PrintServerID: in.PrintServerID,
		PrintServer:   PrintServerToPB(in.PrintServer),
	}
	return m
}
