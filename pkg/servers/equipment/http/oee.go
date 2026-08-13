package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/equipment/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// ChangeEquipmentState godoc
// @Summary 设备状态流转
// @Description 设备状态流转（在用/停用/维修中/报废）
// @Tags 设备管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param id query string true "设备ID"
// @Param state query string true "目标状态"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/equipment/equipment/changestate [put]
func ChangeEquipmentState(c *gin.Context) {
	transID := middleware.GetTransID(c)
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	id := c.Query("id")
	state := c.Query("state")
	if id == "" || state == "" {
		resp.Code = proto.Code_BadRequest
		resp.Message = "id与state不能为空"
		c.JSON(http.StatusOK, resp)
		return
	}
	if err := logic.ChangeEquipmentState(id, state); err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		log.Warnf(context.Background(), "TransID:%s,设备状态流转失败:%v", transID, err)
	}
	c.JSON(http.StatusOK, resp)
}

// ExecuteEquipmentMaintenance godoc
// @Summary 执行维保
// @Description 生成维保记录并滚动计划的下次执行日期，异常时联动设备状态
// @Tags 设备维保计划管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ExecuteEquipmentMaintenanceRequest true "ExecuteEquipmentMaintenanceRequest"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/equipment/maintenplan/execute [put]
func ExecuteEquipmentMaintenance(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ExecuteEquipmentMaintenanceRequest{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,执行维保请求参数无效:%v", transID, err)
		return
	}
	if req.EquipmentMaintenancePlanID == "" {
		resp.Code = proto.Code_BadRequest
		resp.Message = "equipmentMaintenancePlanID不能为空"
		c.JSON(http.StatusOK, resp)
		return
	}
	req.ExecutorID = middleware.GetUserID(c)
	if err := logic.ExecuteEquipmentMaintenance(req); err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// CalcEquipmentOee godoc
// @Summary OEE计算
// @Description 按时间范围计算设备/工站的OEE（时间稼动率×性能稼动率×良品率）
// @Tags 设备管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.EquipmentOeeRequest true "EquipmentOeeRequest"
// @Success 200 {object} proto.EquipmentOeeResponse
// @Router /api/mom/equipment/equipment/oee [post]
func CalcEquipmentOee(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.EquipmentOeeRequest{}
	resp := &proto.EquipmentOeeResponse{}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,OEE计算请求参数无效:%v", transID, err)
		return
	}
	data, err := logic.CalcOee(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	c.JSON(http.StatusOK, data)
}
