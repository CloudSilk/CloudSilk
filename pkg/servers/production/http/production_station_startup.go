package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/production/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddProductionStationStartup godoc
// @Summary 新增
// @Description 新增
// @Tags 工站开机记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionStationStartupInfo true "Add ProductionStationStartup"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionstationstartup/add [post]
func AddProductionStationStartup(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionStationStartupInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建工站开机记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductionStationStartup(model.PBToProductionStationStartup(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductionStationStartup godoc
// @Summary 更新
// @Description 更新
// @Tags 工站开机记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionStationStartupInfo true "Update ProductionStationStartup"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionstationstartup/update [put]
func UpdateProductionStationStartup(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionStationStartupInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新工站开机记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductionStationStartup(model.PBToProductionStationStartup(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductionStationStartup godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 工站开机记录管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param startupTime0 query string false "开机时间开始"
// @Param startupTime1 query string false "开机时间结束"
// @Param productionLineID query string false "生产产线ID"
// @Success 200 {object} proto.QueryProductionStationStartupResponse
// @Router /api/mom/production/productionstationstartup/query [get]
func QueryProductionStationStartup(c *gin.Context) {
	req := &proto.QueryProductionStationStartupRequest{}
	resp := &proto.QueryProductionStationStartupResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductionStationStartup(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductionStationStartup godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 工站开机记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductionStationStartupResponse
// @Router /api/mom/production/productionstationstartup/all [get]
func GetAllProductionStationStartup(c *gin.Context) {
	resp := &proto.GetAllProductionStationStartupResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductionStationStartups()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductionStationStartupsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductionStationStartupDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 工站开机记录管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductionStationStartupDetailResponse
// @Router /api/mom/production/productionstationstartup/detail [get]
func GetProductionStationStartupDetail(c *gin.Context) {
	resp := &proto.GetProductionStationStartupDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductionStationStartupByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionStationStartupToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductionStationStartup godoc
// @Summary 删除
// @Description 删除
// @Tags 工站开机记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductionStationStartup"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionstationstartup/delete [delete]
func DeleteProductionStationStartup(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除工站开机记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductionStationStartup(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductionStationStartupRouter(r *gin.Engine) {
	g := r.Group("/api/mom/production/productionstationstartup")

	g.POST("add", AddProductionStationStartup)
	g.PUT("update", UpdateProductionStationStartup)
	g.GET("query", QueryProductionStationStartup)
	g.DELETE("delete", DeleteProductionStationStartup)
	g.GET("all", GetAllProductionStationStartup)
	g.GET("detail", GetProductionStationStartupDetail)
}
