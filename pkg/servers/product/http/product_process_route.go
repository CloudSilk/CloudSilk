package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/product/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddProductProcessRoute godoc
// @Summary 新增
// @Description 新增
// @Tags 产品工序路线管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductProcessRouteInfo true "Add ProductProcessRoute"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productprocessroute/add [post]
func AddProductProcessRoute(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductProcessRouteInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建产品工序路线请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductProcessRoute(model.PBToProductProcessRoute(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductProcessRoute godoc
// @Summary 更新
// @Description 更新
// @Tags 产品工序路线管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductProcessRouteInfo true "Update ProductProcessRoute"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productprocessroute/update [put]
func UpdateProductProcessRoute(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductProcessRouteInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新产品工序路线请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductProcessRoute(model.PBToProductProcessRoute(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductProcessRoute godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 产品工序路线管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param currentProcessID query string false "当前工序"
// @Param productSerialNo query string false "序列号"
// @Param productOrderNo query string false "生产工单号"
// @Param createTime0 query string false "创建时间开始"
// @Param createTime1 query string false "创建时间结束"
// @Success 200 {object} proto.QueryProductProcessRouteResponse
// @Router /api/mom/product/productprocessroute/query [get]
func QueryProductProcessRoute(c *gin.Context) {
	req := &proto.QueryProductProcessRouteRequest{}
	resp := &proto.QueryProductProcessRouteResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductProcessRoute(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductProcessRoute godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 产品工序路线管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductProcessRouteResponse
// @Router /api/mom/product/productprocessroute/all [get]
func GetAllProductProcessRoute(c *gin.Context) {
	resp := &proto.GetAllProductProcessRouteResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductProcessRoutes()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductProcessRoutesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductProcessRouteDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 产品工序路线管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductProcessRouteDetailResponse
// @Router /api/mom/product/productprocessroute/detail [get]
func GetProductProcessRouteDetail(c *gin.Context) {
	resp := &proto.GetProductProcessRouteDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductProcessRouteByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductProcessRouteToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductProcessRoute godoc
// @Summary 删除
// @Description 删除
// @Tags 产品工序路线管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductProcessRoute"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productprocessroute/delete [delete]
func DeleteProductProcessRoute(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除产品工序路线请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductProcessRoute(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductProcessRouteRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productprocessroute")

	g.POST("add", AddProductProcessRoute)
	g.PUT("update", UpdateProductProcessRoute)
	g.GET("query", QueryProductProcessRoute)
	g.DELETE("delete", DeleteProductProcessRoute)
	g.GET("all", GetAllProductProcessRoute)
	g.GET("detail", GetProductProcessRouteDetail)
}
