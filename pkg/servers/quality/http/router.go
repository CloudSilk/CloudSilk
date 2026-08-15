package http

import (
	mw "github.com/CloudSilk/CloudSilk/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	g := r.Group("/api/mom/quality")

	t := g.Group("qualityinspectiontype")
	t.POST("add", AddQualityInspectionType)
	t.PUT("update", UpdateQualityInspectionType)
	t.GET("query", QueryQualityInspectionType)
	t.GET("all", GetAllQualityInspectionType)
	t.GET("detail", GetQualityInspectionTypeDetail)
	t.DELETE("delete", DeleteQualityInspectionType)

	s := g.Group("qualityinspectionstandard")
	s.POST("add", AddQualityInspectionStandard)
	s.PUT("update", UpdateQualityInspectionStandard)
	s.GET("query", QueryQualityInspectionStandard)
	s.GET("all", GetAllQualityInspectionStandard)
	s.GET("detail", GetQualityInspectionStandardDetail)
	s.DELETE("delete", DeleteQualityInspectionStandard)

	o := g.Group("qualityinspectionorder")
	o.POST("add", AddQualityInspectionOrder)
	o.PUT("update", UpdateQualityInspectionOrder)
	o.GET("query", QueryQualityInspectionOrder)
	o.GET("detail", GetQualityInspectionOrderDetail)
	o.DELETE("delete", DeleteQualityInspectionOrder)
	o.PUT("start", StartQualityInspectionOrder)
	// 判定/让步/返工联动属于质量裁决动作，需授权角色
	o.PUT("complete", CompleteQualityInspectionOrder, mw.RequirePermissionParam("quality.judge"))

	g.POST("spc/calc", CalcSpc)
}
