package logic

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// SCADA 采集值 MQTT 转发：
// 配置系统参数（code=scada，见下方 key 常量）后，采集成功的点位值
// 按主题 {prefix}/{deviceCode}/{tagName} 发布 JSON，
// 供 SmartFlow（RuleGo）等规则引擎订阅联动；未配置 broker 时自动禁用。

// 系统参数 key（code=scada）
const (
	scadaParamBroker      = "mqtt.broker"      // 形如 tcp://127.0.0.1:1883，空=禁用
	scadaParamUsername    = "mqtt.username"
	scadaParamPassword    = "mqtt.password"
	scadaParamTopicPrefix = "mqtt.topicPrefix" // 默认 cloudsilk/scada
)

// mqttEvent 一次点位值发布事件
type mqttEvent struct {
	topic   string
	payload string
}

var (
	mqttOnce      sync.Once
	mqttCh        chan mqttEvent
	mqttClient    mqtt.Client
	mqttEnabled   bool
	mqttMu        sync.RWMutex
	mqttCfgLoaded time.Time
)

const (
	mqttQueueSize  = 1024 // 缓冲上限：超过则丢弃（采集不因转发阻塞）
	mqttCfgRefresh = time.Minute
)

// mqttConfig 读取转发配置（带短缓存）
func mqttConfig() (broker, username, password, prefix string) {
	mqttMu.RLock()
	if mqttEnabled || time.Since(mqttCfgLoaded) < mqttCfgRefresh {
		defer mqttMu.RUnlock()
		return currentMqttCfg()
	}
	mqttMu.RUnlock()

	broker = getScadaParam(scadaParamBroker)
	username = getScadaParam(scadaParamUsername)
	password = getScadaParam(scadaParamPassword)
	prefix = getScadaParam(scadaParamTopicPrefix)
	if prefix == "" {
		prefix = "cloudsilk/scada"
	}

	mqttMu.Lock()
	defer mqttMu.Unlock()
	mqttCfgLoaded = time.Now()
	// broker 从有到无时不回退禁用（需重启），从无到有即时生效
	if broker != "" && !mqttEnabled {
		startForwarderLocked(broker, username, password)
	}
	if broker != "" {
		setCurrentMqttCfgLocked(broker, username, password, prefix)
	}
	return broker, username, password, prefix
}

var (
	curBroker, curUser, curPass, curPrefix string
)

func currentMqttCfg() (string, string, string, string) {
	return curBroker, curUser, curPass, curPrefix
}
func setCurrentMqttCfgLocked(broker, username, password, prefix string) {
	curBroker, curUser, curPass, curPrefix = broker, username, password, prefix
}

// startForwarderLocked 初始化客户端与发布协程（仅一次）
func startForwarderLocked(broker, username, password string) {
	opts := mqtt.NewClientOptions().
		AddBroker(broker).
		SetClientID(fmt.Sprintf("cloudsilk-scada-%d", time.Now().UnixNano())).
		SetAutoReconnect(true).
		SetMaxReconnectInterval(30 * time.Second)
	if username != "" {
		opts.SetUsername(username)
		opts.SetPassword(password)
	}
	client := mqtt.NewClient(opts)
	// 连接失败不阻塞采集：AutoReconnect 兜底
	if token := client.Connect(); token.WaitTimeout(3 * time.Second) && token.Error() != nil {
		// 保留客户端依赖自动重连，转发协程仍会尝试发布
	}

	mqttCh = make(chan mqttEvent, mqttQueueSize)
	mqttClient = client
	mqttEnabled = true
	go func() {
		for ev := range mqttCh {
			if !mqttClient.IsConnectionOpen() {
				continue // 断线期间丢弃，避免堆积
			}
			mqttClient.Publish(ev.topic, 0, false, ev.payload)
		}
	}()
}

// getScadaParam 读取系统参数（code=scada）
func getScadaParam(key string) string {
	m := &model.SystemParamsConfig{}
	if err := model.DB.DB().Where("`code` = ? AND `key` = ?", "scada", key).First(m).Error; err != nil {
		return ""
	}
	return strings.TrimSpace(m.Value)
}

// PublishTagValue 采集值发布入口（非阻塞；未启用/队列满时直接返回）
func PublishTagValue(deviceCode, tagName, value, quality string, collectTime time.Time) {
	broker, _, _, prefix := mqttConfig()
	if broker == "" || deviceCode == "" || tagName == "" {
		return
	}
	topic := fmt.Sprintf("%s/%s/%s", prefix, deviceCode, tagName)
	payload := fmt.Sprintf(`{"value":%q,"quality":%q,"collectTime":%q}`,
		value, quality, collectTime.Format("2006-01-02 15:04:05"))

	mqttMu.RLock()
	defer mqttMu.RUnlock()
	if !mqttEnabled || mqttCh == nil {
		return
	}
	select {
	case mqttCh <- mqttEvent{topic: topic, payload: payload}:
	default: // 队列满丢弃，采集优先
	}
}
