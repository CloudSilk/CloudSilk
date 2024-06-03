package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/webapi/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/pkg/utils/middleware"
	"github.com/gin-gonic/gin"
)

// EnterProductionStation godoc
// @Summary 请求进站接口
// @Description 请求进站接口
// @Tags WebAPI
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.EnterProductionStationRequest true "EnterProductionStationRequest"
// @Success 200 {object} proto.EnterProductionStationResponse
// @Router /api/mom/webapi/production/enterProductionStation [post]
func EnterProductionStation(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.EnterProductionStationRequest{}
	resp := &proto.EnterProductionStationResponse{Code: 0}

	var err error
	if err = c.BindJSON(req); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,请求进站接口参数无效:%v", transID, err)
		return
	}

	if err = middleware.Validate.Struct(req); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp, err = logic.EnterProductionStation(req)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	}

	c.JSON(http.StatusOK, resp)
}

// ExitProductionStation godoc
// @Summary 请求出站接口
// @Description 请求出站接口
// @Tags WebAPI
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ExitProductionStationRequest true "ExitProductionStationRequest"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/webapi/production/exitProductionStation [post]
func ExitProductionStation(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ExitProductionStationRequest{}
	resp := &proto.CommonResponse{Code: 200}

	if err := c.BindJSON(req); err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,请求进站接口参数无效:%v", transID, err)
		return
	}

	if err := middleware.Validate.Struct(req); err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp, err := logic.ExitProductionStation(req)
	if err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CreateProductTestRecord godoc
// @Summary 创建产品测试记录
// @Description 创建产品测试记录
// @Tags WebAPI
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.CreateProductTestRecordRequest true "CreateProductTestRecordRequest"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/webapi/production/createProductTestRecord [post]
func CreateProductTestRecord(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.CreateProductTestRecordRequest{}
	resp := &proto.CommonResponse{Code: 200}

	if err := c.BindJSON(req); err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,创建产品测试记录参数无效:%v", transID, err)
		return
	}

	if err := middleware.Validate.Struct(req); err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp, err := logic.CreateProductTestRecord(req)
	if err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CheckProductProcessRouteFailure godoc
// @Summary 设置失败后续处理接口
// @Description 设置失败后续处理接口
// @Tags WebAPI
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.CheckProductProcessRouteFailureRequest true "CheckProductProcessRouteFailureRequest"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/webapi/production/checkProductProcessRouteFailure [post]
func CheckProductProcessRouteFailure(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.CheckProductProcessRouteFailureRequest{}
	resp := &proto.CommonResponse{Code: 200}

	if err := c.BindJSON(req); err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,设置失败后续处理接口参数无效:%v", transID, err)
		return
	}

	if err := middleware.Validate.Struct(req); err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp, err := logic.CheckProductProcessRouteFailure(req)
	if err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func RegisterProductionRouter(r *gin.Engine) {
	g := r.Group("/api/mom/webapi/production")

	g.POST("enterProductionStation", EnterProductionStation)
	g.POST("exitProductionStation", ExitProductionStation)
	g.POST("createProductTestRecord", CreateProductTestRecord)
	g.POST("checkProductProcessRouteFailure", CheckProductProcessRouteFailure)
}
