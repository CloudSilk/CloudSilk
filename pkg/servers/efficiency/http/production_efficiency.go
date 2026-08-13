package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/efficiency/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddProductionEfficiency godoc
// @Summary 新增
// @Description 新增
// @Tags 生产效率统计
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionEfficiencyInfo true "Add ProductionEfficiency"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/efficiency/productionefficiency/add [post]
func AddProductionEfficiency(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionEfficiencyInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建生产效率统计请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductionEfficiency(model.PBToProductionEfficiency(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductionEfficiency godoc
// @Summary 更新
// @Description 更新
// @Tags 生产效率统计
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionEfficiencyInfo true "Update ProductionEfficiency"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/efficiency/productionefficiency/update [put]
func UpdateProductionEfficiency(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionEfficiencyInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新生产效率统计请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductionEfficiency(model.PBToProductionEfficiency(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductionEfficiency godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 生产效率统计
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param outputDate0 query string false "创建时间开始"
// @Param outputDate1 query string false "创建时间结束"
// @Param productionLineID query string false "生产产线ID"
// @Success 200 {object} proto.QueryProductionEfficiencyResponse
// @Router /api/mom/efficiency/productionefficiency/query [get]
func QueryProductionEfficiency(c *gin.Context) {
	req := &proto.QueryProductionEfficiencyRequest{}
	resp := &proto.QueryProductionEfficiencyResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductionEfficiency(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductionEfficiency godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 生产效率统计
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductionEfficiencyResponse
// @Router /api/mom/efficiency/productionefficiency/all [get]
func GetAllProductionEfficiency(c *gin.Context) {
	resp := &proto.GetAllProductionEfficiencyResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductionEfficiencys()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductionEfficiencysToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductionEfficiencyDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 生产效率统计
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductionEfficiencyDetailResponse
// @Router /api/mom/efficiency/productionefficiency/detail [get]
func GetProductionEfficiencyDetail(c *gin.Context) {
	resp := &proto.GetProductionEfficiencyDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductionEfficiencyByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionEfficiencyToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductionEfficiency godoc
// @Summary 删除
// @Description 删除
// @Tags 生产效率统计
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductionEfficiency"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/efficiency/productionefficiency/delete [delete]
func DeleteProductionEfficiency(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除生产效率统计请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductionEfficiency(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductionEfficiencyRouter(r *gin.Engine) {
	g := r.Group("/api/mom/efficiency/productionefficiency")

	g.POST("add", AddProductionEfficiency)
	g.PUT("update", UpdateProductionEfficiency)
	g.GET("query", QueryProductionEfficiency)
	g.DELETE("delete", DeleteProductionEfficiency)
	g.GET("all", GetAllProductionEfficiency)
	g.GET("detail", GetProductionEfficiencyDetail)
}
