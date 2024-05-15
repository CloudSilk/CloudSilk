package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/servers/webapi/logic"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/pkg/utils/middleware"
	"github.com/gin-gonic/gin"
)

// GetAllProductionLine godoc
// @Summary 获取全部产线信息
// @Description 获取全部产线信息
// @Tags WebAPI
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductionLineResponse
// @Router /api/mom/webapi/infrastructure/getAllProductionLine [get]
func GetAllProductionLine(c *gin.Context) {
	resp := &proto.GetAllProductionLineResponse{
		Code: proto.Code_Success,
	}
	data, err := logic.GetAllProductionLine()
	if err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp.Data = data
	c.JSON(http.StatusOK, resp)
}

// RetrieveProductionStation godoc
// @Summary 查询产线工位信息
// @Description 查询产线工位信息
// @Tags WebAPI
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.RetrieveProductionStationRequest true "RetrieveProductionStation"
// @Success 200 {object} proto.GetAllProductionStationResponse
// @Router /api/mom/webapi/infrastructure/retrieveProductionStation [post]
func RetrieveProductionStation(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.RetrieveProductionStationRequest{}
	resp := &proto.GetAllProductionStationResponse{Code: 200}

	if err := c.BindJSON(req); err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,查询产线工位信息请求参数无效:%v", transID, err)
		return
	}

	if err := middleware.Validate.Struct(req); err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	data, err := logic.RetrieveProductionStation(req)
	if err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp.Data = data
	c.JSON(http.StatusOK, resp)
}

// RetrieveProductAttribute godoc
// @Summary 查询产品特性信息
// @Description 查询产品特性信息
// @Tags WebAPI
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.RetrieveProductAttributeRequest true "RetrieveProductAttribute"
// @Success 200 {object} proto.GetAllProductAttributeResponse
// @Router /api/mom/webapi/infrastructure/retrieveProductAttribute [post]
func RetrieveProductAttribute(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.RetrieveProductAttributeRequest{}
	resp := &proto.GetAllProductAttributeResponse{Code: 200}

	if err := c.BindJSON(req); err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,查询产品特性信息请求参数无效:%v", transID, err)
		return
	}

	if err := middleware.Validate.Struct(req); err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	data, err := logic.RetrieveProductAttribute(req)
	if err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp.Data = data
	c.JSON(http.StatusOK, resp)
}

// RetrieveProductionCrossway godoc
// @Summary 查询产线路口信息
// @Description 查询产线路口信息
// @Tags WebAPI
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.RetrieveProductionCrosswayRequest true "RetrieveProductionCrossway"
// @Success 200 {object} proto.GetAllProductionCrosswayResponse
// @Router /api/mom/webapi/infrastructure/retrieveProductionCrossway [post]
func RetrieveProductionCrossway(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.RetrieveProductionCrosswayRequest{}
	resp := &proto.GetAllProductionCrosswayResponse{Code: 200}

	if err := c.BindJSON(req); err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,查询产品特性信息请求参数无效:%v", transID, err)
		return
	}

	if err := middleware.Validate.Struct(req); err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	data, err := logic.RetrieveProductionCrossway(req)
	if err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp.Data = data
	c.JSON(http.StatusOK, resp)
}

func RegisterInfrastructureRouter(r *gin.Engine) {
	g := r.Group("/api/mom/webapi/infrastructure")

	g.GET("getAllProductionLine", GetAllProductionLine)
	g.POST("retrieveProductionStation", RetrieveProductionStation)
	g.POST("retrieveProductAttribute", RetrieveProductAttribute)
	g.POST("retrieveProductionCrossway", RetrieveProductionCrossway)
}
