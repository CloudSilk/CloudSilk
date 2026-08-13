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

// AddEquipment godoc
// @Summary 新增
// @Description 新增
// @Tags 设备管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.EquipmentInfo true "Add Equipment"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/equipment/equipment/add [post]
func AddEquipment(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.EquipmentInfo{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新增设备请求参数无效:%v", transID, err)
		return
	}
	id, err := logic.CreateEquipment(model.PBToEquipment(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateEquipment godoc
// @Summary 更新
// @Description 更新
// @Tags 设备管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.EquipmentInfo true "Update Equipment"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/equipment/equipment/update [put]
func UpdateEquipment(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.EquipmentInfo{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新设备请求参数无效:%v", transID, err)
		return
	}
	err = logic.UpdateEquipment(model.PBToEquipment(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryEquipment godoc
// @Summary 查询
// @Description 查询
// @Tags 设备管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.QueryEquipmentResponse
// @Router /api/mom/equipment/equipment/query [get]
func QueryEquipment(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.QueryEquipmentRequest{}
	resp := &proto.QueryEquipmentResponse{}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,查询设备请求参数无效:%v", transID, err)
		return
	}
	logic.QueryEquipment(req, resp, true)
	c.JSON(http.StatusOK, resp)
}

// GetAllEquipment godoc
// @Summary 获取所有
// @Description 获取所有
// @Tags 设备管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllEquipmentResponse
// @Router /api/mom/equipment/equipment/all [get]
func GetAllEquipment(c *gin.Context) {
	resp := &proto.GetAllEquipmentResponse{}
	list, err := logic.GetAllEquipments()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.EquipmentsToPB(list)
		resp.Code = proto.Code_Success
	}
	c.JSON(http.StatusOK, resp)
}

// GetEquipmentDetail godoc
// @Summary 获取详情
// @Description 获取详情
// @Tags 设备管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param id query string true "ID"
// @Success 200 {object} proto.GetEquipmentDetailResponse
// @Router /api/mom/equipment/equipment/detail [get]
func GetEquipmentDetail(c *gin.Context) {
	resp := &proto.GetEquipmentDetailResponse{}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		resp.Message = "id不能为空"
		c.JSON(http.StatusOK, resp)
		return
	}
	m, err := logic.GetEquipmentByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.EquipmentToPB(m)
		resp.Code = proto.Code_Success
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteEquipment godoc
// @Summary 删除
// @Description 删除
// @Tags 设备管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.GetByIDsRequest true "Delete Equipment"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/equipment/equipment/delete [delete]
func DeleteEquipment(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.GetByIDsRequest{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,删除设备请求参数无效:%v", transID, err)
		return
	}
	if len(req.Ids) == 0 || req.Ids[0] == "" {
		resp.Code = proto.Code_BadRequest
		resp.Message = "id不能为空"
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteEquipment(req.Ids[0])
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}
