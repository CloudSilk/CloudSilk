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

// GetScadaTagValues godoc
// @Summary 点位实时值
// @Description 全部（或指定设备的）点位实时值
// @Tags 采集点位
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.QueryScadaTagValueResponse
// @Router /api/mom/scada/tag/values [get]
func GetScadaTagValues(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.QueryScadaTagValueRequest{}
	resp := &proto.QueryScadaTagValueResponse{}
	if err := c.BindQuery(req); err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,查询点位实时值请求参数无效:%v", transID, err)
		return
	}
	list, err := logic.GetScadaTagValues(req)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ScadaTagValuesToPB(list)
		resp.Code = proto.Code_Success
	}
	c.JSON(http.StatusOK, resp)
}

// QueryScadaTagHistory godoc
// @Summary 点位历史值
// @Description 按点位与时间区间查询历史值
// @Tags 采集点位
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.QueryScadaTagHistoryResponse
// @Router /api/mom/scada/tag/history [get]
func QueryScadaTagHistory(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.QueryScadaTagHistoryRequest{}
	resp := &proto.QueryScadaTagHistoryResponse{}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,查询点位历史值请求参数无效:%v", transID, err)
		return
	}
	logic.QueryScadaTagHistory(req, resp)
	c.JSON(http.StatusOK, resp)
}

// TestScadaDevice godoc
// @Summary 测试设备连接
// @Description 用给定地址与从站号测试 Modbus TCP 连通性
// @Tags 采集设备
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.TestScadaDeviceRequest true "TestScadaDeviceRequest"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/scada/device/test [post]
func TestScadaDevice(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.TestScadaDeviceRequest{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,测试设备连接请求参数无效:%v", transID, err)
		return
	}
	if err := logic.TestScadaDevice(req); err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}
