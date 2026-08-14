// Package testutil 提供跨模块集成测试共享的内存数据库初始化。
package testutil

import (
	"fmt"
	"sync/atomic"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/pkg/db"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// testDBClient 复用分页/查重逻辑，仅替换底层连接
type testDBClient struct {
	db.DBClientInterface
	gdb *gorm.DB
}

func (c *testDBClient) DB() *gorm.DB { return c.gdb }
func (c *testDBClient) Close()       {}

// SetupTestDB 初始化内存 SQLite、替换全局 model.DB 并建全量表。
// 返回的 gorm.DB 供造数与断言使用。
var dbSeq atomic.Int64

// SetupTestDB 初始化相互隔离的内存 SQLite 并建全量表。
// 说明：裸 ":memory:" 与连接池不兼容（每个新连接是独立空库，且事务内嵌套查询
// 会取第二个连接），因此用"唯一命名的共享缓存内存库"——同库多连接可见同一份数据，
// 每次调用独立命名实现用例间隔离。
func SetupTestDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("file:cloudsilk_test_%d?mode=memory&cache=shared&_pragma=busy_timeout(5000)", dbSeq.Add(1))
	gdb, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	model.InitDB(&testDBClient{DBClientInterface: &db.DBClient{}, gdb: gdb}, false)
	model.AutoMigrate()
	return gdb, nil
}
