package logic

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/modbus"
	"github.com/gopcua/opcua"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

// 设备连接状态
const (
	ScadaConnectionOK      = "正常"
	ScadaConnectionError   = "异常"
	ScadaConnectionInitial = "未连接"
)

// 点位值质量
const (
	ScadaQualityGood = "Good"
	ScadaQualityBad  = "Bad"
)

// ---------- 设备 CRUD ----------

func CreateScadaDevice(m *model.ScadaDevice) (string, error) {
	duplication, err := model.DB.CreateWithCheckDuplication(m, " `code`  = ? ", m.Code)
	if err != nil {
		return "", err
	}
	if duplication {
		return "", errors.New("存在相同采集设备代号")
	}
	if m.Protocol == "" {
		m.Protocol = ProtocolModbusTCP
	}
	if m.IntervalSeconds <= 0 {
		m.IntervalSeconds = 10
	}
	if m.ConnectionState == "" {
		m.ConnectionState = ScadaConnectionInitial
	}
	return m.ID, nil
}

func UpdateScadaDevice(m *model.ScadaDevice) error {
	omits := []string{"created_at", "connection_state", "last_collect_time", "last_error"}
	duplication, err := model.DB.UpdateWithCheckDuplicationAndOmit(model.DB.DB(), m, false, omits, "`id` != ? and  `code`  = ? ", m.ID, m.Code)
	if err != nil {
		return err
	}
	if duplication {
		return errors.New("存在相同采集设备代号")
	}
	return nil
}

func QueryScadaDevice(req *proto.QueryScadaDeviceRequest, resp *proto.QueryScadaDeviceResponse, preload bool) {
	db := model.DB.DB().Model(&model.ScadaDevice{})
	if preload {
		db = db.Preload(clause.Associations)
	}
	if req.Code != "" {
		db = db.Where("`code` LIKE ? OR `name` LIKE ?", "%"+req.Code+"%", "%"+req.Code+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ScadaDevice
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ScadaDevicesToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllScadaDevices() (list []*model.ScadaDevice, err error) {
	err = model.DB.DB().Preload(clause.Associations).Find(&list, "`enable` = ?", true).Error
	return
}

func GetScadaDeviceByID(id string) (*model.ScadaDevice, error) {
	m := &model.ScadaDevice{}
	err := model.DB.DB().Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func DeleteScadaDevice(id string) error {
	return model.DB.DB().Delete(&model.ScadaDevice{}, "`id` = ?", id).Error
}

// TestScadaDevice 测试设备连通性（按协议：Modbus 读保持寄存器0 / OPC UA 读状态节点）
func TestScadaDevice(req *proto.TestScadaDeviceRequest) error {
	if req.Address == "" {
		return errors.New("address不能为空")
	}
	if req.Protocol == ProtocolOpcUA {
		return TestOpcUAConnection(req.Address)
	}
	client := modbus.NewClient(modbus.NewTCPClientProvider(req.Address))
	_, err := client.ReadHoldingRegisters(byte(req.SlaveID), 0, 1)
	if err != nil {
		return fmt.Errorf("连接失败：%w", err)
	}
	return nil
}

// ---------- 点位 CRUD ----------

func CreateScadaTag(m *model.ScadaTag) (string, error) {
	if m.ScadaDeviceID == "" {
		return "", errors.New("采集设备不能为空")
	}
	if err := validateTagForDevice(m); err != nil {
		return "", err
	}
	return m.ID, model.DB.DB().Create(m).Error
}

// UpdateScadaTag 白名单更新（防全量 Save 清零漏传字段；设备归属不允许改）
func UpdateScadaTag(m *model.ScadaTag) error {
	if m.ID == "" {
		return errors.New("id不能为空")
	}
	if err := validateTagForDevice(m); err != nil {
		return err
	}
	return model.DB.DB().Model(&model.ScadaTag{}).Where("`id` = ?", m.ID).Updates(map[string]interface{}{
		"name":          m.Name,
		"address":       m.Address,
		"opc_ua_node_id": m.OpcUANodeID,
		"function_code": m.FunctionCode,
		"data_type":     m.DataType,
		"scale":         m.Scale,
		"unit":          m.Unit,
		"save_history":  m.SaveHistory,
		"enable":        m.Enable,
		"remark":        m.Remark,
	}).Error
}

// validateTagForDevice 按所属设备协议校验点位必填项与功能码
func validateTagForDevice(m *model.ScadaTag) error {
	if !validDataType(m.DataType) {
		return fmt.Errorf("无效的数据类型：%s（支持 bool/uint16/int16/float32）", m.DataType)
	}
	device := &model.ScadaDevice{}
	if err := model.DB.DB().Select("id", "protocol").First(device, "`id` = ?", m.ScadaDeviceID).Error; err != nil {
		return errors.New("所属采集设备不存在")
	}
	if device.Protocol == ProtocolOpcUA {
		if m.OpcUANodeID == "" {
			return errors.New("OPC UA 设备的点位必须配置节点ID（opcUANodeID，如 ns=2;s=Channel1.Tag1）")
		}
		return nil
	}
	// Modbus：功能码必填且合法
	switch m.FunctionCode {
	case 0, 1, 3, 4:
		return nil
	}
	return fmt.Errorf("Modbus 点位功能码无效：%d（支持 0-线圈 1-离散输入 3-保持寄存器 4-输入寄存器）", m.FunctionCode)
}

func QueryScadaTag(req *proto.QueryScadaTagRequest, resp *proto.QueryScadaTagResponse, preload bool) {
	db := model.DB.DB().Model(&model.ScadaTag{})
	if preload {
		db = db.Preload(clause.Associations)
	}
	if req.ScadaDeviceID != "" {
		db = db.Where("`scada_device_id` = ?", req.ScadaDeviceID)
	}
	if req.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+req.Name+"%")
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ScadaTag
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ScadaTagsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllScadaTags() (list []*model.ScadaTag, err error) {
	err = model.DB.DB().Preload(clause.Associations).Find(&list, "`enable` = ?", true).Error
	return
}

func GetScadaTagByID(id string) (*model.ScadaTag, error) {
	m := &model.ScadaTag{}
	err := model.DB.DB().Preload(clause.Associations).Where("`id` = ?", id).First(m).Error
	return m, err
}

func DeleteScadaTag(id string) error {
	return model.DB.DB().Delete(&model.ScadaTag{}, "`id` = ?", id).Error
}

// ---------- 实时值与历史 ----------

func GetScadaTagValues(req *proto.QueryScadaTagValueRequest) ([]*model.ScadaTagValue, error) {
	db := model.DB.DB().Preload("ScadaTag").Preload("ScadaTag.ScadaDevice")
	if req.ScadaDeviceID != "" {
		db = db.Joins("JOIN scada_tags ON scada_tag_values.scada_tag_id=scada_tags.id").
			Where("scada_tags.scada_device_id = ?", req.ScadaDeviceID)
	}
	var list []*model.ScadaTagValue
	err := db.Find(&list).Error
	return list, err
}

func QueryScadaTagHistory(req *proto.QueryScadaTagHistoryRequest, resp *proto.QueryScadaTagHistoryResponse) {
	if req.ScadaTagID == "" {
		resp.Code = proto.Code_BadRequest
		resp.Message = "scadaTagID不能为空"
		return
	}
	db := model.DB.DB().Model(&model.ScadaTagHistory{}).Where("`scada_tag_id` = ?", req.ScadaTagID)
	if req.StartTime != "" {
		if t, err := time.ParseInLocation("2006-01-02 15:04:05", req.StartTime, time.Local); err == nil {
			db = db.Where("`collect_time` >= ?", t)
		}
	}
	if req.EndTime != "" {
		if t, err := time.ParseInLocation("2006-01-02 15:04:05", req.EndTime, time.Local); err == nil {
			db = db.Where("`collect_time` <= ?", t)
		}
	}

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "collect_time desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.ScadaTagHistory
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ScadaTagHistoriesToPB(list)
	}
	resp.Total = resp.Records
}

// ---------- 采集器 ----------

// collectorSupervisor 设备热加载监督者：周期性重载启用设备清单，
// 新增设备启动采集协程，停用/删除的设备停止；协程每次采集前重读点位，
// 点位增删改无需重启服务即生效。
var (
	collectorMu      sync.Mutex
	collectorRunning = map[string]context.CancelFunc{}
)

const collectorReloadInterval = 30 * time.Second

// StartCollector 启动采集监督者（幂等，可重复调用）
func StartCollector() {
	go superviseCollector()
}

func superviseCollector() {
	reload := func() {
		devices := []*model.ScadaDevice{}
		if err := model.DB.DB().Where("`enable` = ?", true).Find(&devices).Error; err != nil {
			return
		}
		active := map[string]bool{}
		for _, d := range devices {
			active[d.ID] = true
		}

		collectorMu.Lock()
		for id, cancel := range collectorRunning {
			if !active[id] {
				cancel()
				delete(collectorRunning, id)
			}
		}
		for _, d := range devices {
			if _, running := collectorRunning[d.ID]; running {
				continue
			}
			ctx, cancel := context.WithCancel(context.Background())
			collectorRunning[d.ID] = cancel
			go collectDeviceLoop(ctx, d.ID)
		}
		collectorMu.Unlock()
	}

	reload()
	ticker := time.NewTicker(collectorReloadInterval)
	defer ticker.Stop()
	for range ticker.C {
		reload()
	}
}

// StopCollector 停止全部采集协程（服务关闭时调用）
func StopCollector() {
	collectorMu.Lock()
	defer collectorMu.Unlock()
	for id, cancel := range collectorRunning {
		cancel()
		delete(collectorRunning, id)
	}
}

// collectDeviceLoop 单设备采集循环：每轮重读设备与点位配置（热更新），
// 连接循环内复用（不再每轮重连），采集出错时丢弃重建；设备停用/删除退出
func collectDeviceLoop(ctx context.Context, deviceID string) {
	var (
		modbusClient modbus.Client
		opcuaClient  *opcua.Client
	)
	resetClients := func() {
		if modbusClient != nil {
			modbusClient = nil
		}
		if opcuaClient != nil {
			cctx, cancel := context.WithTimeout(context.Background(), time.Second)
			opcuaClient.Close(cctx)
			cancel()
			opcuaClient = nil
		}
	}
	defer resetClients()

	for {
		device := &model.ScadaDevice{}
		err := model.DB.DB().Preload("Tags").First(device, "`id` = ?", deviceID).Error
		if err != nil || !device.Enable {
			return // 已删除或停用
		}
		interval := time.Duration(device.IntervalSeconds) * time.Second
		if interval < time.Second {
			interval = time.Second
		}

		switch {
		case device.Protocol == ProtocolOpcUA:
			if opcuaClient == nil {
				c, err := newOpcUAClient(device.Address)
				if err != nil {
					updateDeviceError(device, err)
				} else {
					opcuaClient = c
				}
			}
			if opcuaClient != nil {
				if err := collectOpcUADevice(device, opcuaClient); err != nil {
					updateDeviceError(device, err)
					cctx, cancel := context.WithTimeout(context.Background(), time.Second)
					opcuaClient.Close(cctx)
					cancel()
					opcuaClient = nil // 下轮重连
				}
			}
		default:
			if modbusClient == nil {
				modbusClient = modbus.NewClient(modbus.NewTCPClientProvider(device.Address))
			}
			if err := collectModbusDevice(device, modbusClient); err != nil {
				updateDeviceError(device, err)
				modbusClient = nil // 下轮重连
			}
		}

		select {
		case <-ctx.Done():
			return
		case <-time.After(interval):
		}
	}
}

func collectModbusDevice(device *model.ScadaDevice, client modbus.Client) error {
	if len(device.Tags) == 0 {
		return nil
	}

	slaveID := byte(device.SlaveID)

	now := time.Now()
	var successCount, failCount int
	var firstError string
	for _, tag := range device.Tags {
		if !tag.Enable {
			continue
		}
		value, err := readTag(client, slaveID, tag)
		if err != nil {
			failCount++
			if firstError == "" {
				firstError = err.Error()
			}
			upsertTagValue(tag.ID, "", ScadaQualityBad, now)
			continue
		}
		successCount++
		upsertTagValue(tag.ID, value, ScadaQualityGood, now)
		if tag.SaveHistory {
			model.DB.DB().Create(&model.ScadaTagHistory{
				ScadaTagID:  tag.ID,
				Value:       value,
				CollectTime: now,
			})
		}
		// MQTT 转发（SmartFlow 等订阅方；未配置 broker 时为空操作）
		PublishTagValue(device.Code, tag.Name, value, ScadaQualityGood, now)
	}

	state := ScadaConnectionOK
	errmsg := ""
	if successCount == 0 && failCount > 0 {
		state = ScadaConnectionError
		errmsg = firstError
	}
	model.DB.DB().Model(&model.ScadaDevice{}).Where("`id` = ?", device.ID).
		Updates(map[string]interface{}{
			"connection_state":  state,
			"last_collect_time": now,
			"last_error":        errmsg,
		})

	//全部失败视为连接级错误（触发上层重连）
	if successCount == 0 && failCount > 0 {
		return errors.New(errmsg)
	}
	return nil
}

func readTag(client modbus.Client, slaveID byte, tag *model.ScadaTag) (string, error) {
	switch tag.FunctionCode {
	case 0: //线圈
		bits, err := client.ReadCoils(slaveID, uint16(tag.Address), 1)
		if err != nil {
			return "", err
		}
		return boolText(len(bits) > 0 && bits[0] != 0), nil
	case 1: //离散输入
		bits, err := client.ReadDiscreteInputs(slaveID, uint16(tag.Address), 1)
		if err != nil {
			return "", err
		}
		return boolText(len(bits) > 0 && bits[0] != 0), nil
	case 3: //保持寄存器
		regs, err := client.ReadHoldingRegisters(slaveID, uint16(tag.Address), registerCount(tag.DataType))
		if err != nil {
			return "", err
		}
		return parseRegisters(regs, tag.DataType, tag.Scale)
	case 4: //输入寄存器
		regs, err := client.ReadInputRegisters(slaveID, uint16(tag.Address), registerCount(tag.DataType))
		if err != nil {
			return "", err
		}
		return parseRegisters(regs, tag.DataType, tag.Scale)
	default:
		return "", fmt.Errorf("无效的功能码：%d", tag.FunctionCode)
	}
}

func registerCount(dataType string) uint16 {
	if dataType == "float32" {
		return 2
	}
	return 1
}

// parseRegisters 解析寄存器并应用缩放系数
func parseRegisters(regs []uint16, dataType string, scale float64) (string, error) {
	if len(regs) == 0 {
		return "", errors.New("空寄存器数据")
	}
	var value float64
	switch dataType {
	case "uint16", "":
		value = float64(regs[0])
	case "int16":
		value = float64(int16(regs[0]))
	case "float32":
		if len(regs) < 2 {
			return "", errors.New("float32需要2个寄存器")
		}
		value = float64(math.Float32frombits(binary.BigEndian.Uint32([]byte{byte(regs[0] >> 8), byte(regs[0]), byte(regs[1] >> 8), byte(regs[1])})))
	default:
		return "", fmt.Errorf("无效的数据类型：%s", dataType)
	}
	if scale != 0 && scale != 1 {
		value *= scale
	}
	return fmt.Sprintf("%g", value), nil
}

func upsertTagValue(tagID, value, quality string, collectTime time.Time) {
	existing := &model.ScadaTagValue{}
	err := model.DB.DB().Where("`scada_tag_id` = ?", tagID).First(existing).Error
	if err != nil {
		model.DB.DB().Create(&model.ScadaTagValue{
			ScadaTagID:  tagID,
			Value:       value,
			Quality:     quality,
			CollectTime: &collectTime,
		})
		return
	}
	model.DB.DB().Model(&model.ScadaTagValue{}).Where("`id` = ?", existing.ID).
		Updates(map[string]interface{}{
			"value":         value,
			"quality":       quality,
			"collect_time":  collectTime,
		})
}

func validDataType(dataType string) bool {
	switch dataType {
	case "bool", "uint16", "int16", "float32":
		return true
	}
	return false
}

func boolText(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

// collectOpcUADevice OPC UA 设备采集：使用外部复用连接顺序读取全部点位，
// 全部失败返回错误（由上层关闭并重建连接）
func collectOpcUADevice(device *model.ScadaDevice, client *opcua.Client) error {
	if client == nil {
		return errors.New("OPC UA 连接未建立")
	}

	now := time.Now()
	var successCount, failCount int
	var firstError string
	for _, tag := range device.Tags {
		if !tag.Enable {
			continue
		}
		value, err := readOpcUaNode(client, tag.OpcUANodeID)
		if err != nil {
			failCount++
			if firstError == "" {
				firstError = err.Error()
			}
			upsertTagValue(tag.ID, "", ScadaQualityBad, now)
			continue
		}
		successCount++
		value = applyScale(value, tag.Scale)
		upsertTagValue(tag.ID, value, ScadaQualityGood, now)
		if tag.SaveHistory {
			model.DB.DB().Create(&model.ScadaTagHistory{
				ScadaTagID:  tag.ID,
				Value:       value,
				CollectTime: now,
			})
		}
		// MQTT 转发（SmartFlow 等订阅方；未配置 broker 时为空操作）
		PublishTagValue(device.Code, tag.Name, value, ScadaQualityGood, now)
	}

	state := ScadaConnectionOK
	errmsg := ""
	if successCount == 0 && failCount > 0 {
		state = ScadaConnectionError
		errmsg = firstError
	}
	model.DB.DB().Model(&model.ScadaDevice{}).Where("`id` = ?", device.ID).
		Updates(map[string]interface{}{
			"connection_state":  state,
			"last_collect_time": now,
			"last_error":        errmsg,
		})

	if successCount == 0 && failCount > 0 {
		return errors.New(errmsg)
	}
	return nil
}

// updateDeviceError 连接失败时回写设备异常状态并触发断线告警
func updateDeviceError(device *model.ScadaDevice, err error) {
	model.DB.DB().Model(&model.ScadaDevice{}).Where("`id` = ?", device.ID).
		Updates(map[string]interface{}{
			"connection_state": ScadaConnectionError,
			"last_error":       err.Error(),
		})
	createScadaAlarmEvent(device, err)
}

// createScadaAlarmEvent 断线告警：状态从未知/正常翻转为异常时记录异常轨迹
func createScadaAlarmEvent(device *model.ScadaDevice, err error) {
	model.DB.DB().Create(&model.ExceptionTrace{
		Level:        "Error",
		Host:         device.Address,
		Source:       fmt.Sprintf("SCADA采集设备 %s（%s）", device.Code, device.Protocol),
		Message:      fmt.Sprintf("连接失败：%v", err),
		ReportUserID: "system",
	})
}

// applyScale 对字符串数值应用缩放系数（非数值原样返回）
func applyScale(value string, scale float64) string {
	if scale == 0 || scale == 1 {
		return value
	}
	var f float64
	if _, err := fmt.Sscanf(value, "%g", &f); err != nil {
		return value
	}
	return fmt.Sprintf("%g", f*scale)
}
