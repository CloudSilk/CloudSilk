package model

import (
	"database/sql"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
)

// 设备台账
type Equipment struct {
	ModelID
	Code                string             `json:"code" gorm:"index;size:100;comment:代号"`
	Name                string             `json:"name" gorm:"size:200;comment:名称"`
	Model               string             `json:"model" gorm:"size:200;comment:型号规格"`
	Manufacturer        string             `json:"manufacturer" gorm:"size:200;comment:生产厂家"`
	SerialNo            string             `json:"serialNo" gorm:"size:100;comment:出厂编号"`
	Location            string             `json:"location" gorm:"size:200;comment:安装位置"`
	Team                string             `json:"team" gorm:"size:100;comment:责任班组"`
	CurrentState        string             `json:"currentState" gorm:"size:100;comment:设备状态"`
	CommissionDate      sql.NullTime       `json:"commissionDate" gorm:"comment:投运日期"`
	ProductionStationID *string            `json:"productionStationID" gorm:"size:36;comment:关联生产工站ID"`
	ProductionStation   *ProductionStation `json:"productionStation" gorm:"constraint:OnDelete:SET NULL"`
	Enable              bool               `json:"enable" gorm:"default:true;comment:是否启用"`
	Remark              string             `json:"remark" gorm:"size:1000;comment:备注"`
}

func PBToEquipments(in []*proto.EquipmentInfo) []*Equipment {
	var result []*Equipment
	for _, c := range in {
		result = append(result, PBToEquipment(c))
	}
	return result
}

func PBToEquipment(in *proto.EquipmentInfo) *Equipment {
	if in == nil {
		return nil
	}
	var productionStationID *string
	if in.ProductionStationID != "" {
		productionStationID = &in.ProductionStationID
	}
	e := &Equipment{
		ModelID:             ModelID{ID: in.Id},
		Code:                in.Code,
		Name:                in.Name,
		Model:               in.Model,
		Manufacturer:        in.Manufacturer,
		SerialNo:            in.SerialNo,
		Location:            in.Location,
		Team:                in.Team,
		CurrentState:        in.CurrentState,
		ProductionStationID: productionStationID,
		Enable:              in.Enable,
		Remark:              in.Remark,
	}
	if in.CommissionDate != "" {
		if t, err := time.Parse("2006-01-02", in.CommissionDate); err == nil {
			e.CommissionDate = sql.NullTime{Time: t, Valid: true}
		}
	}
	return e
}

func EquipmentsToPB(in []*Equipment) []*proto.EquipmentInfo {
	var list []*proto.EquipmentInfo
	for _, f := range in {
		list = append(list, EquipmentToPB(f))
	}
	return list
}

func EquipmentToPB(in *Equipment) *proto.EquipmentInfo {
	if in == nil {
		return nil
	}
	m := &proto.EquipmentInfo{
		Id:       in.ID,
		Code:     in.Code,
		Name:     in.Name,
		Model:    in.Model,
		Manufacturer: in.Manufacturer,
		SerialNo: in.SerialNo,
		Location: in.Location,
		Team:     in.Team,
		CurrentState: in.CurrentState,
		Enable:   in.Enable,
		Remark:   in.Remark,
	}
	if in.CommissionDate.Valid {
		m.CommissionDate = in.CommissionDate.Time.Format("2006-01-02")
	}
	if in.ProductionStationID != nil {
		m.ProductionStationID = *in.ProductionStationID
	}
	if in.ProductionStation != nil {
		m.ProductionStationCode = in.ProductionStation.Code
	}
	return m
}

// 设备维保计划
type EquipmentMaintenancePlan struct {
	ModelID
	EquipmentID      *string     `json:"equipmentID" gorm:"size:36;comment:设备ID"`
	Equipment        *Equipment  `json:"equipment" gorm:"constraint:OnDelete:CASCADE"`
	MaintenanceType  string      `json:"maintenanceType" gorm:"size:100;comment:维保类型"`
	CycleDays        int32       `json:"cycleDays" gorm:"comment:周期(天)"`
	LastExecution    sql.NullTime `json:"lastExecution" gorm:"comment:上次执行日期"`
	NextExecution    sql.NullTime `json:"nextExecution" gorm:"comment:下次执行日期"`
	RemindDays       int32       `json:"remindDays" gorm:"comment:提前提醒天数"`
	Enable           bool        `json:"enable" gorm:"default:true;comment:是否启用"`
	Content          string      `json:"content" gorm:"size:2000;comment:维保内容"`
	Remark           string      `json:"remark" gorm:"size:1000;comment:备注"`
}

func PBToEquipmentMaintenancePlans(in []*proto.EquipmentMaintenancePlanInfo) []*EquipmentMaintenancePlan {
	var result []*EquipmentMaintenancePlan
	for _, c := range in {
		result = append(result, PBToEquipmentMaintenancePlan(c))
	}
	return result
}

func PBToEquipmentMaintenancePlan(in *proto.EquipmentMaintenancePlanInfo) *EquipmentMaintenancePlan {
	if in == nil {
		return nil
	}
	var equipmentID *string
	if in.EquipmentID != "" {
		equipmentID = &in.EquipmentID
	}
	p := &EquipmentMaintenancePlan{
		ModelID:          ModelID{ID: in.Id},
		EquipmentID:      equipmentID,
		MaintenanceType:  in.MaintenanceType,
		CycleDays:        in.CycleDays,
		RemindDays:       in.RemindDays,
		Enable:           in.Enable,
		Content:          in.Content,
		Remark:           in.Remark,
	}
	if in.LastExecutionDate != "" {
		if t, err := time.Parse("2006-01-02", in.LastExecutionDate); err == nil {
			p.LastExecution = sql.NullTime{Time: t, Valid: true}
		}
	}
	if in.NextExecutionDate != "" {
		if t, err := time.Parse("2006-01-02", in.NextExecutionDate); err == nil {
			p.NextExecution = sql.NullTime{Time: t, Valid: true}
		}
	}
	return p
}

func EquipmentMaintenancePlansToPB(in []*EquipmentMaintenancePlan) []*proto.EquipmentMaintenancePlanInfo {
	var list []*proto.EquipmentMaintenancePlanInfo
	for _, f := range in {
		list = append(list, EquipmentMaintenancePlanToPB(f))
	}
	return list
}

func EquipmentMaintenancePlanToPB(in *EquipmentMaintenancePlan) *proto.EquipmentMaintenancePlanInfo {
	if in == nil {
		return nil
	}
	m := &proto.EquipmentMaintenancePlanInfo{
		Id:              in.ID,
		MaintenanceType: in.MaintenanceType,
		CycleDays:       in.CycleDays,
		RemindDays:      in.RemindDays,
		Enable:          in.Enable,
		Content:         in.Content,
		Remark:          in.Remark,
	}
	if in.EquipmentID != nil {
		m.EquipmentID = *in.EquipmentID
	}
	if in.Equipment != nil {
		m.EquipmentCode = in.Equipment.Code
	}
	if in.LastExecution.Valid {
		m.LastExecutionDate = in.LastExecution.Time.Format("2006-01-02")
	}
	if in.NextExecution.Valid {
		m.NextExecutionDate = in.NextExecution.Time.Format("2006-01-02")
	}
	return m
}

// 设备维保记录
type EquipmentMaintenanceRecord struct {
	ModelID
	EquipmentMaintenancePlanID *string                     `json:"equipmentMaintenancePlanID" gorm:"size:36;comment:维保计划ID"`
	EquipmentMaintenancePlan   *EquipmentMaintenancePlan   `json:"equipmentMaintenancePlan" gorm:"constraint:OnDelete:SET NULL"`
	EquipmentID                *string                     `json:"equipmentID" gorm:"size:36;comment:设备ID"`
	Equipment                  *Equipment                  `json:"equipment" gorm:"constraint:OnDelete:CASCADE"`
	MaintenanceType            string                      `json:"maintenanceType" gorm:"size:100;comment:维保类型"`
	ExecutionDate              time.Time                   `json:"executionDate" gorm:"autoCreateTime:nano;comment:执行日期"`
	ExecutorID                 string                      `json:"executorID" gorm:"size:36;comment:执行人"`
	Content                    string                      `json:"content" gorm:"size:2000;comment:执行内容"`
	DurationMinutes            int32                       `json:"durationMinutes" gorm:"comment:耗时(分钟)"`
	Result                     string                      `json:"result" gorm:"size:100;comment:结果"`
	AbnormalDescription        string                      `json:"abnormalDescription" gorm:"size:2000;comment:异常说明"`
	Remark                     string                      `json:"remark" gorm:"size:1000;comment:备注"`
}

func PBToEquipmentMaintenanceRecords(in []*proto.EquipmentMaintenanceRecordInfo) []*EquipmentMaintenanceRecord {
	var result []*EquipmentMaintenanceRecord
	for _, c := range in {
		result = append(result, PBToEquipmentMaintenanceRecord(c))
	}
	return result
}

func PBToEquipmentMaintenanceRecord(in *proto.EquipmentMaintenanceRecordInfo) *EquipmentMaintenanceRecord {
	if in == nil {
		return nil
	}
	var equipmentID, planID *string
	if in.EquipmentID != "" {
		equipmentID = &in.EquipmentID
	}
	if in.EquipmentMaintenancePlanID != "" {
		planID = &in.EquipmentMaintenancePlanID
	}
	return &EquipmentMaintenanceRecord{
		ModelID:                    ModelID{ID: in.Id},
		EquipmentMaintenancePlanID: planID,
		EquipmentID:                equipmentID,
		MaintenanceType:            in.MaintenanceType,
		ExecutorID:                 in.ExecutorID,
		Content:                    in.Content,
		DurationMinutes:            in.DurationMinutes,
		Result:                     in.Result,
		AbnormalDescription:        in.AbnormalDescription,
		Remark:                     in.Remark,
	}
}

func EquipmentMaintenanceRecordsToPB(in []*EquipmentMaintenanceRecord) []*proto.EquipmentMaintenanceRecordInfo {
	var list []*proto.EquipmentMaintenanceRecordInfo
	for _, f := range in {
		list = append(list, EquipmentMaintenanceRecordToPB(f))
	}
	return list
}

func EquipmentMaintenanceRecordToPB(in *EquipmentMaintenanceRecord) *proto.EquipmentMaintenanceRecordInfo {
	if in == nil {
		return nil
	}
	m := &proto.EquipmentMaintenanceRecordInfo{
		Id:                  in.ID,
		MaintenanceType:     in.MaintenanceType,
		ExecutionDate:       in.ExecutionDate.Format("2006-01-02 15:04:05"),
		ExecutorID:          in.ExecutorID,
		Content:             in.Content,
		DurationMinutes:     in.DurationMinutes,
		Result:              in.Result,
		AbnormalDescription: in.AbnormalDescription,
		Remark:              in.Remark,
	}
	if in.EquipmentID != nil {
		m.EquipmentID = *in.EquipmentID
	}
	if in.Equipment != nil {
		m.EquipmentCode = in.Equipment.Code
	}
	if in.EquipmentMaintenancePlanID != nil {
		m.EquipmentMaintenancePlanID = *in.EquipmentMaintenancePlanID
	}
	return m
}
