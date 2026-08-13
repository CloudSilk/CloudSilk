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

// AddProductionStationEfficiency godoc
// @Summary 新增
// @Description 新增
// @Tags 工位效率统计
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionStationEfficiencyInfo true "Add ProductionStationEfficiency"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/efficiency/productionstationefficiency/add [post]
func AddProductionStationEfficiency(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionStationEfficiencyInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建工位效率统计请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductionStationEfficiency(model.PBToProductionStationEfficiency(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductionStationEfficiency godoc
// @Summary 更新
// @Description 更新
// @Tags 工位效率统计
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionStationEfficiencyInfo true "Update ProductionStationEfficiency"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/efficiency/productionstationefficiency/update [put]
func UpdateProductionStationEfficiency(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionStationEfficiencyInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新工位效率统计请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductionStationEfficiency(model.PBToProductionStationEfficiency(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductionStationEfficiency godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 工位效率统计
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
// @Success 200 {object} proto.QueryProductionStationEfficiencyResponse
// @Router /api/mom/efficiency/productionstationefficiency/query [get]
func QueryProductionStationEfficiency(c *gin.Context) {
	req := &proto.QueryProductionStationEfficiencyRequest{}
	resp := &proto.QueryProductionStationEfficiencyResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductionStationEfficiency(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductionStationEfficiency godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 工位效率统计
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductionStationEfficiencyResponse
// @Router /api/mom/efficiency/productionstationefficiency/all [get]
func GetAllProductionStationEfficiency(c *gin.Context) {
	resp := &proto.GetAllProductionStationEfficiencyResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductionStationEfficiencys()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductionStationEfficiencysToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductionStationEfficiencyDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 工位效率统计
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductionStationEfficiencyDetailResponse
// @Router /api/mom/efficiency/productionstationefficiency/detail [get]
func GetProductionStationEfficiencyDetail(c *gin.Context) {
	resp := &proto.GetProductionStationEfficiencyDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductionStationEfficiencyByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionStationEfficiencyToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductionStationEfficiency godoc
// @Summary 删除
// @Description 删除
// @Tags 工位效率统计
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductionStationEfficiency"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/efficiency/productionstationefficiency/delete [delete]
func DeleteProductionStationEfficiency(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除工位效率统计请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductionStationEfficiency(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductionStationEfficiencyRouter(r *gin.Engine) {
	g := r.Group("/api/mom/efficiency/productionstationefficiency")

	g.POST("add", AddProductionStationEfficiency)
	g.PUT("update", UpdateProductionStationEfficiency)
	g.GET("query", QueryProductionStationEfficiency)
	g.DELETE("delete", DeleteProductionStationEfficiency)
	g.GET("all", GetAllProductionStationEfficiency)
	g.GET("detail", GetProductionStationEfficiencyDetail)
}
