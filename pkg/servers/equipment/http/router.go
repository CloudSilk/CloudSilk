package http

import (
	mw "github.com/CloudSilk/CloudSilk/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	g := r.Group("/api/mom/equipment")

	e := g.Group("equipment")
	e.POST("add", AddEquipment)
	e.PUT("update", UpdateEquipment)
	e.GET("query", QueryEquipment)
	e.GET("all", GetAllEquipment)
	e.GET("detail", GetEquipmentDetail)
	e.DELETE("delete", DeleteEquipment)
	e.PUT("changestate", ChangeEquipmentState, mw.RequirePermissionParam("equipment.manage"))
	e.POST("oee", CalcEquipmentOee)

	p := g.Group("maintenplan")
	p.POST("add", AddEquipmentMaintenancePlan)
	p.PUT("update", UpdateEquipmentMaintenancePlan)
	p.GET("query", QueryEquipmentMaintenancePlan)
	p.GET("all", GetAllEquipmentMaintenancePlan)
	p.GET("detail", GetEquipmentMaintenancePlanDetail)
	p.DELETE("delete", DeleteEquipmentMaintenancePlan)
	p.PUT("execute", ExecuteEquipmentMaintenance, mw.RequirePermissionParam("equipment.manage"))

	m := g.Group("maintenrecord")
	m.GET("query", QueryEquipmentMaintenanceRecord)
	m.GET("detail", GetEquipmentMaintenanceRecordDetail)
}
