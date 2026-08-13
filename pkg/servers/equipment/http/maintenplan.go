package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/equipment/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddEquipmentMaintenancePlan godoc
// @Summary 新增
// @Description 新增
// @Tags 设备维保计划管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.EquipmentMaintenancePlanInfo true "Add EquipmentMaintenancePlan"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/equipment/maintenplan/add [post]
func AddEquipmentMaintenancePlan(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.EquipmentMaintenancePlanInfo{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新增设备维保计划请求参数无效:%v", transID, err)
		return
	}
	id, err := logic.CreateEquipmentMaintenancePlan(model.PBToEquipmentMaintenancePlan(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateEquipmentMaintenancePlan godoc
// @Summary 更新
// @Description 更新
// @Tags 设备维保计划管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.EquipmentMaintenancePlanInfo true "Update EquipmentMaintenancePlan"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/equipment/maintenplan/update [put]
func UpdateEquipmentMaintenancePlan(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.EquipmentMaintenancePlanInfo{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新设备维保计划请求参数无效:%v", transID, err)
		return
	}
	err = logic.UpdateEquipmentMaintenancePlan(model.PBToEquipmentMaintenancePlan(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryEquipmentMaintenancePlan godoc
// @Summary 查询
// @Description 查询
// @Tags 设备维保计划管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.QueryEquipmentMaintenancePlanResponse
// @Router /api/mom/equipment/maintenplan/query [get]
func QueryEquipmentMaintenancePlan(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.QueryEquipmentMaintenancePlanRequest{}
	resp := &proto.QueryEquipmentMaintenancePlanResponse{}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,查询设备维保计划请求参数无效:%v", transID, err)
		return
	}
	logic.QueryEquipmentMaintenancePlan(req, resp, true)
	c.JSON(http.StatusOK, resp)
}

// GetAllEquipmentMaintenancePlan godoc
// @Summary 获取所有
// @Description 获取所有
// @Tags 设备维保计划管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllEquipmentMaintenancePlanResponse
// @Router /api/mom/equipment/maintenplan/all [get]
func GetAllEquipmentMaintenancePlan(c *gin.Context) {
	resp := &proto.GetAllEquipmentMaintenancePlanResponse{}
	list, err := logic.GetAllEquipmentMaintenancePlans()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.EquipmentMaintenancePlansToPB(list)
		resp.Code = proto.Code_Success
	}
	c.JSON(http.StatusOK, resp)
}

// GetEquipmentMaintenancePlanDetail godoc
// @Summary 获取详情
// @Description 获取详情
// @Tags 设备维保计划管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param id query string true "ID"
// @Success 200 {object} proto.GetEquipmentMaintenancePlanDetailResponse
// @Router /api/mom/equipment/maintenplan/detail [get]
func GetEquipmentMaintenancePlanDetail(c *gin.Context) {
	resp := &proto.GetEquipmentMaintenancePlanDetailResponse{}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		resp.Message = "id不能为空"
		c.JSON(http.StatusOK, resp)
		return
	}
	m, err := logic.GetEquipmentMaintenancePlanByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.EquipmentMaintenancePlanToPB(m)
		resp.Code = proto.Code_Success
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteEquipmentMaintenancePlan godoc
// @Summary 删除
// @Description 删除
// @Tags 设备维保计划管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.GetByIDsRequest true "Delete EquipmentMaintenancePlan"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/equipment/maintenplan/delete [delete]
func DeleteEquipmentMaintenancePlan(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.GetByIDsRequest{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,删除设备维保计划请求参数无效:%v", transID, err)
		return
	}
	if len(req.Ids) == 0 || req.Ids[0] == "" {
		resp.Code = proto.Code_BadRequest
		resp.Message = "id不能为空"
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteEquipmentMaintenancePlan(req.Ids[0])
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}
