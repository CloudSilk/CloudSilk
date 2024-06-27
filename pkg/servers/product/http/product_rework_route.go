package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/product/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/pkg/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddProductReworkRoute godoc
// @Summary 新增
// @Description 新增
// @Tags 返工路线管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductReworkRouteInfo true "Add ProductReworkRoute"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productreworkroute/add [post]
func AddProductReworkRoute(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductReworkRouteInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建返工路线请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductReworkRoute(model.PBToProductReworkRoute(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductReworkRoute godoc
// @Summary 更新
// @Description 更新
// @Tags 返工路线管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductReworkRouteInfo true "Update ProductReworkRoute"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productreworkroute/update [put]
func UpdateProductReworkRoute(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductReworkRouteInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新返工路线请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductReworkRoute(model.PBToProductReworkRoute(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductReworkRoute godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 返工路线管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productionLineID query string false "生产产线ID"
// @Param code query string false "代号或描述"
// @Success 200 {object} proto.QueryProductReworkRouteResponse
// @Router /api/mom/product/productreworkroute/query [get]
func QueryProductReworkRoute(c *gin.Context) {
	req := &proto.QueryProductReworkRouteRequest{}
	resp := &proto.QueryProductReworkRouteResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductReworkRoute(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductReworkRoute godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 返工路线管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductReworkRouteResponse
// @Router /api/mom/product/productreworkroute/all [get]
func GetAllProductReworkRoute(c *gin.Context) {
	resp := &proto.GetAllProductReworkRouteResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductReworkRoutes()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductReworkRoutesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductReworkRouteDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 返工路线管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductReworkRouteDetailResponse
// @Router /api/mom/product/productreworkroute/detail [get]
func GetProductReworkRouteDetail(c *gin.Context) {
	resp := &proto.GetProductReworkRouteDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductReworkRouteByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductReworkRouteToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductReworkRoute godoc
// @Summary 删除
// @Description 删除
// @Tags 返工路线管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductReworkRoute"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productreworkroute/delete [delete]
func DeleteProductReworkRoute(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除返工路线请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductReworkRoute(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductReworkRouteRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productreworkroute")

	g.POST("add", AddProductReworkRoute)
	g.PUT("update", UpdateProductReworkRoute)
	g.GET("query", QueryProductReworkRoute)
	g.DELETE("delete", DeleteProductReworkRoute)
	g.GET("all", GetAllProductReworkRoute)
	g.GET("detail", GetProductReworkRouteDetail)
}
