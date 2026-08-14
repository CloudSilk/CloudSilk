package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/ua"
)

// 支持的采集协议
const (
	ProtocolModbusTCP = "modbus-tcp"
	ProtocolOpcUA     = "opcua"
)

// opcUATimeout 单次 OPC UA 读写超时
const opcUATimeout = 5 * time.Second

// newOpcUAClient 建立 OPC UA 连接（None 安全模式）
func newOpcUAClient(endpoint string) (*opcua.Client, error) {
	c, err := opcua.NewClient(endpoint, opcua.SecurityMode(ua.MessageSecurityModeNone))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), opcUATimeout)
	defer cancel()
	if err := c.Connect(ctx); err != nil {
		return nil, fmt.Errorf("OPC UA 连接失败：%w", err)
	}
	return c, nil
}

// readOpcUaNode 读取单个 OPC UA 节点值（数值以 %g 格式化，其余取字符串表示）
func readOpcUaNode(client *opcua.Client, nodeID string) (string, error) {
	if nodeID == "" {
		return "", fmt.Errorf("OPC UA 点位缺少节点ID")
	}
	id, err := ua.ParseNodeID(nodeID)
	if err != nil {
		return "", fmt.Errorf("节点ID无效（%s）：%w", nodeID, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), opcUATimeout)
	defer cancel()
	resp, err := client.Read(ctx, &ua.ReadRequest{
		NodesToRead: []*ua.ReadValueID{{NodeID: id}},
	})
	if err != nil {
		return "", err
	}
	if len(resp.Results) == 0 {
		return "", fmt.Errorf("节点%s无返回结果", nodeID)
	}
	result := resp.Results[0]
	if result.Status != ua.StatusOK {
		return "", fmt.Errorf("读取节点%s失败：%s", nodeID, result.Status.Error())
	}
	if result.Value == nil || result.Value.Value() == nil {
		return "", nil
	}
	switch v := result.Value.Value().(type) {
	case float64:
		return fmt.Sprintf("%g", v), nil
	case float32:
		return fmt.Sprintf("%g", v), nil
	case int64:
		return fmt.Sprintf("%d", v), nil
	case int32:
		return fmt.Sprintf("%d", v), nil
	case bool:
		return boolText(v), nil
	default:
		return fmt.Sprintf("%v", v), nil
	}
}

// TestOpcUAConnection 测试 OPC UA 服务器连通性（读取 Server_ServerStatus_CurrentTime 井号节点）
func TestOpcUAConnection(endpoint string) error {
	client, err := newOpcUAClient(endpoint)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), opcUATimeout)
	defer cancel()
	defer client.Close(ctx)

	_, err = readOpcUaNode(client, "i=2258")
	if err != nil {
		return fmt.Errorf("OPC UA 服务器可达但读取状态节点失败：%w", err)
	}
	return nil
}
