package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/scada/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddScadaDevice godoc
// @Summary 新增
// @Description 新增
// @Tags 采集设备
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ScadaDeviceInfo true "Add ScadaDevice"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/scada/device/add [post]
func AddScadaDevice(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ScadaDeviceInfo{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新增采集设备请求参数无效:%v", transID, err)
		return
	}
	id, err := logic.CreateScadaDevice(model.PBToScadaDevice(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateScadaDevice godoc
// @Summary 更新
// @Description 更新
// @Tags 采集设备
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ScadaDeviceInfo true "Update ScadaDevice"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/scada/device/update [put]
func UpdateScadaDevice(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ScadaDeviceInfo{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新采集设备请求参数无效:%v", transID, err)
		return
	}
	err = logic.UpdateScadaDevice(model.PBToScadaDevice(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryScadaDevice godoc
// @Summary 查询
// @Description 查询
// @Tags 采集设备
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.QueryScadaDeviceResponse
// @Router /api/mom/scada/device/query [get]
func QueryScadaDevice(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.QueryScadaDeviceRequest{}
	resp := &proto.QueryScadaDeviceResponse{}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,查询采集设备请求参数无效:%v", transID, err)
		return
	}
	logic.QueryScadaDevice(req, resp, true)
	c.JSON(http.StatusOK, resp)
}

// GetAllScadaDevice godoc
// @Summary 获取所有
// @Description 获取所有
// @Tags 采集设备
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllScadaDeviceResponse
// @Router /api/mom/scada/device/all [get]
func GetAllScadaDevice(c *gin.Context) {
	resp := &proto.GetAllScadaDeviceResponse{}
	list, err := logic.GetAllScadaDevices()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ScadaDevicesToPB(list)
		resp.Code = proto.Code_Success
	}
	c.JSON(http.StatusOK, resp)
}

// GetScadaDeviceDetail godoc
// @Summary 获取详情
// @Description 获取详情
// @Tags 采集设备
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param id query string true "ID"
// @Success 200 {object} proto.GetScadaDeviceDetailResponse
// @Router /api/mom/scada/device/detail [get]
func GetScadaDeviceDetail(c *gin.Context) {
	resp := &proto.GetScadaDeviceDetailResponse{}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		resp.Message = "id不能为空"
		c.JSON(http.StatusOK, resp)
		return
	}
	m, err := logic.GetScadaDeviceByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ScadaDeviceToPB(m)
		resp.Code = proto.Code_Success
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteScadaDevice godoc
// @Summary 删除
// @Description 删除
// @Tags 采集设备
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.GetByIDsRequest true "Delete ScadaDevice"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/scada/device/delete [delete]
func DeleteScadaDevice(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.GetByIDsRequest{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,删除采集设备请求参数无效:%v", transID, err)
		return
	}
	if len(req.Ids) == 0 || req.Ids[0] == "" {
		resp.Code = proto.Code_BadRequest
		resp.Message = "id不能为空"
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteScadaDevice(req.Ids[0])
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}
