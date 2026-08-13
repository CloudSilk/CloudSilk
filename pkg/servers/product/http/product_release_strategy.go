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

// AddProductReleaseStrategy godoc
// @Summary 新增
// @Description 新增
// @Tags 产品投料策略管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductReleaseStrategyInfo true "Add ProductReleaseStrategy"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productreleasestrategy/add [post]
func AddProductReleaseStrategy(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductReleaseStrategyInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建产品投料策略请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductReleaseStrategy(model.PBToProductReleaseStrategy(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductReleaseStrategy godoc
// @Summary 更新
// @Description 更新
// @Tags 产品投料策略管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductReleaseStrategyInfo true "Update ProductReleaseStrategy"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productreleasestrategy/update [put]
func UpdateProductReleaseStrategy(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductReleaseStrategyInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新产品投料策略请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductReleaseStrategy(model.PBToProductReleaseStrategy(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductReleaseStrategy godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 产品投料策略管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productionLineID query string false "产线ID"
// @Param code query string false "类别或描述"
// @Success 200 {object} proto.QueryProductReleaseStrategyResponse
// @Router /api/mom/product/productreleasestrategy/query [get]
func QueryProductReleaseStrategy(c *gin.Context) {
	req := &proto.QueryProductReleaseStrategyRequest{}
	resp := &proto.QueryProductReleaseStrategyResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductReleaseStrategy(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductReleaseStrategy godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 产品投料策略管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductReleaseStrategyResponse
// @Router /api/mom/product/productreleasestrategy/all [get]
func GetAllProductReleaseStrategy(c *gin.Context) {
	resp := &proto.GetAllProductReleaseStrategyResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductReleaseStrategys()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductReleaseStrategysToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductReleaseStrategyDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 产品投料策略管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductReleaseStrategyDetailResponse
// @Router /api/mom/product/productreleasestrategy/detail [get]
func GetProductReleaseStrategyDetail(c *gin.Context) {
	resp := &proto.GetProductReleaseStrategyDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductReleaseStrategyByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductReleaseStrategyToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductReleaseStrategy godoc
// @Summary 删除
// @Description 删除
// @Tags 产品投料策略管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductReleaseStrategy"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productreleasestrategy/delete [delete]
func DeleteProductReleaseStrategy(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除产品投料策略请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductReleaseStrategy(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductReleaseStrategyRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productreleasestrategy")

	g.POST("add", AddProductReleaseStrategy)
	g.PUT("update", UpdateProductReleaseStrategy)
	g.GET("query", QueryProductReleaseStrategy)
	g.DELETE("delete", DeleteProductReleaseStrategy)
	g.GET("all", GetAllProductReleaseStrategy)
	g.GET("detail", GetProductReleaseStrategyDetail)
}
