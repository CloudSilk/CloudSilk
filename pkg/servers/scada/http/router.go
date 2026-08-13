package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	g := r.Group("/api/mom/scada")

	d := g.Group("device")
	d.POST("add", AddScadaDevice)
	d.PUT("update", UpdateScadaDevice)
	d.GET("query", QueryScadaDevice)
	d.GET("all", GetAllScadaDevice)
	d.GET("detail", GetScadaDeviceDetail)
	d.DELETE("delete", DeleteScadaDevice)
	d.POST("test", TestScadaDevice)

	t := g.Group("tag")
	t.POST("add", AddScadaTag)
	t.PUT("update", UpdateScadaTag)
	t.GET("query", QueryScadaTag)
	t.GET("all", GetAllScadaTag)
	t.GET("detail", GetScadaTagDetail)
	t.DELETE("delete", DeleteScadaTag)
	t.GET("values", GetScadaTagValues)
	t.GET("history", QueryScadaTagHistory)
}
