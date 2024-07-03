package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/production_base/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddProductionCrossway godoc
// @Summary 新增
// @Description 新增
// @Tags 产线路口管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionCrosswayInfo true "Add ProductionCrossway"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productionbase/productioncrossway/add [post]
func AddProductionCrossway(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionCrosswayInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建产线路口请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductionCrossway(model.PBToProductionCrossway(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductionCrossway godoc
// @Summary 更新
// @Description 更新
// @Tags 产线路口管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionCrosswayInfo true "Update ProductionCrossway"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productionbase/productioncrossway/update [put]
func UpdateProductionCrossway(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionCrosswayInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新产线路口请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductionCrossway(model.PBToProductionCrossway(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductionCrossway godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 产线路口管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Param productionLineID query string false "生产产线ID"
// @Success 200 {object} proto.QueryProductionCrosswayResponse
// @Router /api/mom/productionbase/productioncrossway/query [get]
func QueryProductionCrossway(c *gin.Context) {
	req := &proto.QueryProductionCrosswayRequest{}
	resp := &proto.QueryProductionCrosswayResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductionCrossway(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductionCrossway godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 产线路口管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductionCrosswayResponse
// @Router /api/mom/productionbase/productioncrossway/all [get]
func GetAllProductionCrossway(c *gin.Context) {
	resp := &proto.GetAllProductionCrosswayResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductionCrossways()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductionCrosswaysToPB(list)

	c.JSON(http.StatusOK, resp)
}

// GetProductionCrosswayDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 产线路口管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductionCrosswayDetailResponse
// @Router /api/mom/productionbase/productioncrossway/detail [get]
func GetProductionCrosswayDetail(c *gin.Context) {
	resp := &proto.GetProductionCrosswayDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductionCrosswayByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionCrosswayToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductionCrossway godoc
// @Summary 删除
// @Description 删除
// @Tags 产线路口管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductionCrossway"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productionbase/productioncrossway/delete [delete]
func DeleteProductionCrossway(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.DelRequest{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,删除产线路口请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductionCrossway(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductionCrosswayRouter(r *gin.Engine) {
	g := r.Group("/api/mom/productionbase/productioncrossway")

	g.POST("add", AddProductionCrossway)
	g.PUT("update", UpdateProductionCrossway)
	g.GET("query", QueryProductionCrossway)
	g.DELETE("delete", DeleteProductionCrossway)
	g.GET("all", GetAllProductionCrossway)
	g.GET("detail", GetProductionCrosswayDetail)
}
