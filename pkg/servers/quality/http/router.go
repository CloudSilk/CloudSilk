package http

import (
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
	o.PUT("complete", CompleteQualityInspectionOrder)

	g.POST("spc/calc", CalcSpc)
}
