package main

import (
	"flag"
	"fmt"
	"os"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/CloudSilk/CloudSilk/docs"
	"github.com/CloudSilk/CloudSilk/pkg/clients"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/servers/label"
	"github.com/CloudSilk/CloudSilk/pkg/servers/material"
	"github.com/CloudSilk/CloudSilk/pkg/servers/product"
	"github.com/CloudSilk/CloudSilk/pkg/servers/product_base"
	"github.com/CloudSilk/CloudSilk/pkg/servers/production"
	"github.com/CloudSilk/CloudSilk/pkg/servers/production_base"
	"github.com/CloudSilk/CloudSilk/pkg/servers/system"
	"github.com/CloudSilk/CloudSilk/pkg/servers/trace"
	"github.com/CloudSilk/CloudSilk/pkg/servers/user"
	"github.com/CloudSilk/CloudSilk/pkg/servers/webapi"
	"github.com/CloudSilk/CloudSilk/pkg/servers/webapi/http"
	"github.com/CloudSilk/curd/gen"
	curdhttp "github.com/CloudSilk/curd/http"
	curdmodel "github.com/CloudSilk/curd/model"
	curdservice "github.com/CloudSilk/curd/service"
	pkgconfig "github.com/CloudSilk/pkg/config"
	"github.com/CloudSilk/pkg/constants"
	"github.com/CloudSilk/pkg/db"
	"github.com/CloudSilk/pkg/db/mysql"
	"github.com/CloudSilk/pkg/db/sqlite"
	"github.com/CloudSilk/pkg/utils"
	uchttp "github.com/CloudSilk/usercenter/http"
	ucmodel "github.com/CloudSilk/usercenter/model"
	"github.com/CloudSilk/usercenter/model/token"
	ucmiddleware "github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type IServer interface {
	Start(*gin.Engine)
}

func StartAll(webPath string, port int, singleDB bool) {
	err := pkgconfig.InitFromFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	if singleDB {
		ok, dbClient := pkgconfig.NewDB("all")
		if !ok {
			panic("未配置数据库")
		}
		model.InitDB(dbClient, pkgconfig.DefaultConfig.Debug)
		curdmodel.InitDB(dbClient, pkgconfig.DefaultConfig.Debug)
		ucmodel.InitDB(dbClient, pkgconfig.DefaultConfig.Debug)
	} else {
		// 单独配置数据库
		ok, usercenterDBClient := pkgconfig.NewDB("Usercenter")
		if !ok {
			panic("未配置Usercenter数据库")
		}
		ok, curdDBClient := pkgconfig.NewDB("Curd")
		if !ok {
			panic("未配置Curd数据库")
		}
		ok, cloudSilkDBClient := pkgconfig.NewDB("CloudSilk")
		if !ok {
			panic("未配置CloudSilk数据库")
		}
		model.InitDB(cloudSilkDBClient, pkgconfig.DefaultConfig.Debug)
		curdmodel.InitDB(curdDBClient, pkgconfig.DefaultConfig.Debug)
		ucmodel.InitDB(usercenterDBClient, pkgconfig.DefaultConfig.Debug)
	}

	curdservice.Init()
	token.InitTokenCache(pkgconfig.DefaultConfig.Token.Key, pkgconfig.DefaultConfig.Token.RedisAddr, pkgconfig.DefaultConfig.Token.RedisName, pkgconfig.DefaultConfig.Token.RedisPwd, pkgconfig.DefaultConfig.Token.Expired)
	constants.SetPlatformTenantID(pkgconfig.DefaultConfig.PlatformTenantID)
	constants.SetSuperAdminRoleID(pkgconfig.DefaultConfig.SuperAdminRoleID)
	constants.SetDefaultRoleID(pkgconfig.DefaultConfig.DefaultRoleID)
	constants.SetEnabelTenant(pkgconfig.DefaultConfig.EnableTenant)
	ucmodel.SetDefaultPwd(pkgconfig.DefaultConfig.DefaultPwd)

	gen.LoadCache()

	r := gin.Default()
	r.Use(utils.Cors())
	http.RegisterAdminRouter(r)
	r.Use(ucmiddleware.AuthRequired)

	uchttp.RegisterAuthRouter(r)
	curdhttp.RegisterRouter(r)

	startMom(r)

	r.Static("/web", webPath)
	r.Run(fmt.Sprintf(":%d", port))
}

func startMom(r *gin.Engine) {
	docs.SwaggerInfo.Title = "mom API"
	docs.SwaggerInfo.Description = "This is a mom server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = ""
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/mom/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	servers := []IServer{
		&production_base.Server{},
		&production.Server{},
		&product_base.Server{},
		&product.Server{},
		&label.Server{},
		&system.Server{},
		&material.Server{},
		&user.Server{},
		&trace.Server{},
		&webapi.Server{},
	}
	for _, server := range servers {
		server.Start(r)
	}
	clients.Init(os.Getenv("SERVICE_MODE"))
}

func StartOne(dbType string, port int) {
	r := gin.Default()
	r.Use(utils.Cors())
	http.RegisterAdminRouter(r)

	if os.Getenv("MOM_DISABLE_AUTH") != "true" {
		ucmiddleware.InitIdentity()
		r.Use(ucmiddleware.AuthRequiredWithRPC)
	}

	startMom(r)
	if err := config.Load(); err != nil {
		panic(err)
	}
	params := config.GetRootConfig().ConfigCenter.Params
	var dbClient db.DBClientInterface
	debug := params["debug"] == "true"
	if dbType == "sqlite" {
		dbClient = sqlite.NewSqlite2("", "", "./mom.s3db", "mom", debug)
	} else {
		dbClient = mysql.NewMysql(params["mysql"], debug)
	}

	model.InitDB(dbClient, debug)
	fmt.Println("started server")

	r.Run(fmt.Sprintf(":%d", port))
}

func main() {
	webPath := flag.String("ui", "./web", "web路径")
	dbType := flag.String("db_type", "mysql", "数据库类型：sqlite和mysql两种")
	serviceMode := flag.String("service_mode", "One", "运行模式：ALL、One")
	singleDB := flag.Bool("single_db", true, "使用同一个数据")
	port := flag.Int("port", 48900, "端口")
	flag.Parse()
	if *serviceMode == "ALL" {
		StartAll(*webPath, *port, *singleDB)
	} else {
		StartOne(*dbType, *port)
	}
}
