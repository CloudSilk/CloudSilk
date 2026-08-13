package logic

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/modbus"
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
		m.Protocol = "modbus-tcp"
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

// TestScadaDevice 测试设备连通性（读保持寄存器0）
func TestScadaDevice(req *proto.TestScadaDeviceRequest) error {
	if req.Address == "" {
		return errors.New("address不能为空")
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
	if !validDataType(m.DataType) {
		return "", fmt.Errorf("无效的数据类型：%s（支持 bool/uint16/int16/float32）", m.DataType)
	}
	return m.ID, model.DB.DB().Create(m).Error
}

func UpdateScadaTag(m *model.ScadaTag) error {
	if !validDataType(m.DataType) {
		return fmt.Errorf("无效的数据类型：%s（支持 bool/uint16/int16/float32）", m.DataType)
	}
	return model.DB.DB().Omit("created_at").Save(m).Error
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

// StartCollector 启动后台采集器：按设备配置的间隔轮询启用设备的启用点位
// 每个设备一个 goroutine，互不阻塞；连接错误回写设备状态
func StartCollector() {
	devices := []*model.ScadaDevice{}
	if err := model.DB.DB().Preload("Tags").Where("`enable` = ?", true).Find(&devices).Error; err != nil {
		return
	}
	var wg sync.WaitGroup
	for _, device := range devices {
		wg.Add(1)
		go func(d *model.ScadaDevice) {
			defer wg.Done()
			interval := time.Duration(d.IntervalSeconds) * time.Second
			if interval < time.Second {
				interval = time.Second
			}
			ticker := time.NewTicker(interval)
			defer ticker.Stop()
			for range ticker.C {
				collectDevice(d)
			}
		}(device)
	}
}

func collectDevice(device *model.ScadaDevice) {
	if len(device.Tags) == 0 {
		return
	}

	client := modbus.NewClient(modbus.NewTCPClientProvider(device.Address))
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
