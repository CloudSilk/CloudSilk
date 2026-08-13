package logic

import (
	"testing"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/types"
	"github.com/CloudSilk/pkg/db"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// setupRouteTestDB 初始化内存 SQLite 并建表
func setupRouteTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	gdb, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("打开内存数据库失败: %v", err)
	}
	model.InitDB(&testDBClient{DBClientInterface: &db.DBClient{}, gdb: gdb}, false)
	model.AutoMigrate()
	return gdb
}

func seedRouteFixture(t *testing.T, gdb *gorm.DB) (*model.ProductInfo, string, string) {
	t.Helper()
	processA := &model.ProductionProcess{ModelID: model.ModelID{ID: "process-a"}, Code: "P10", SortIndex: 1, Enable: true}
	processB := &model.ProductionProcess{ModelID: model.ModelID{ID: "process-b"}, Code: "P20", SortIndex: 2, Enable: true}
	if err := gdb.Create([]*model.ProductionProcess{processA, processB}).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}

	order := &model.ProductOrder{ModelID: model.ModelID{ID: "order-1"}, ProductOrderNo: "WO-001", CurrentState: types.ProductOrderStateProducting}
	productInfo := &model.ProductInfo{ModelID: model.ModelID{ID: "product-1"}, ProductSerialNo: "SN-001", CurrentState: types.ProductStateReleased, ProductOrderID: order.ID}
	if err := gdb.Create([]*model.ProductOrder{order}).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}
	if err := gdb.Create(productInfo).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}

	//工单工艺（动态创建模式的数据源）：两道工序
	orderProcesses := []*model.ProductOrderProcess{
		{ModelID: model.ModelID{ID: "op-1"}, ProductOrderID: order.ID, ProductionProcessID: processA.ID, SortIndex: 1, Enable: true},
		{ModelID: model.ModelID{ID: "op-2"}, ProductOrderID: order.ID, ProductionProcessID: processB.ID, SortIndex: 2, Enable: true},
	}
	if err := gdb.Create(orderProcesses).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}
	return productInfo, processA.ID, processB.ID
}

// 模式一：直接创建模式——产线已为产品预建工艺路线，直接返回待加工路线
func TestResolveRoute_DirectMode(t *testing.T) {
	gdb := setupRouteTestDB(t)
	productInfo, processA, _ := seedRouteFixture(t, gdb)

	precreated := &model.ProductProcessRoute{
		CurrentProcessID: processA,
		CurrentState:     types.ProductProcessRouteStateWaitProcess,
		RouteIndex:       1,
		ProductInfoID:    productInfo.ID,
	}
	if err := gdb.Create(precreated).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}

	route, err := resolveProductProcessRoute(gdb, productInfo, -1, nil, "route_index")
	if err != nil {
		t.Fatalf("解析失败: %v", err)
	}
	if route.ID != precreated.ID {
		t.Fatalf("应返回预建的待加工路线 %s，实际 %s", precreated.ID, route.ID)
	}

	//不应产生新的动态路线
	var count int64
	gdb.Model(&model.ProductProcessRoute{}).Count(&count)
	if count != 1 {
		t.Fatalf("直接创建模式不应新建路线，当前共 %d 条", count)
	}
}

// 模式二：工单工艺动态创建模式——无预建路线时按工单工序顺序生成
func TestResolveRoute_DynamicMode(t *testing.T) {
	gdb := setupRouteTestDB(t)
	productInfo, processA, processB := seedRouteFixture(t, gdb)

	//从头解析：应动态创建首道工序路线
	route, err := resolveProductProcessRoute(gdb, productInfo, -1, nil, "route_index")
	if err != nil {
		t.Fatalf("解析失败: %v", err)
	}
	if route.ID == "" || route.CurrentProcessID != processA || route.RouteIndex != 1 {
		t.Fatalf("应动态创建首道工序路线（%s, 索引1），实际：%+v", processA, route)
	}

	//解析下一道：应动态创建第二道工序路线，并记录上一道工序
	next, err := resolveProductProcessRoute(gdb, productInfo, route.RouteIndex, &route.CurrentProcessID, "work_index")
	if err != nil {
		t.Fatalf("解析失败: %v", err)
	}
	if next.ID == "" || next.CurrentProcessID != processB || next.RouteIndex != 2 {
		t.Fatalf("应动态创建第二道工序路线（%s, 索引2），实际：%+v", processB, next)
	}
	if next.LastProcessID == nil || *next.LastProcessID != processA {
		t.Fatalf("动态路线应记录上一道工序 %s", processA)
	}

	//全部解析完成后返回空路线（非错误）
	last, err := resolveProductProcessRoute(gdb, productInfo, next.RouteIndex, &next.CurrentProcessID, "work_index")
	if err != nil {
		t.Fatalf("解析失败: %v", err)
	}
	if last.ID != "" {
		t.Fatalf("工序耗尽时应返回空路线，实际 %s", last.ID)
	}
}

// testDBClient 仅用于测试的 DBClient（复用分页/查重逻辑，仅替换底层连接）
type testDBClient struct {
	db.DBClientInterface
	gdb *gorm.DB
}

func (c *testDBClient) DB() *gorm.DB { return c.gdb }
func (c *testDBClient) Close()       {}
