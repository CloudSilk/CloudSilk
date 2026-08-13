package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/product_base/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddProductModel godoc
// @Summary 新增
// @Description 新增
// @Tags 产品型号管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductModelInfo true "Add ProductModel"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productbase/productmodel/add [post]
func AddProductModel(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductModelInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建产品型号请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductModel(model.PBToProductModel(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductModel godoc
// @Summary 更新
// @Description 更新
// @Tags 产品型号管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductModelInfo true "Update ProductModel"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productbase/productmodel/update [put]
func UpdateProductModel(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductModelInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新产品型号请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductModel(model.PBToProductModel(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductModel godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 产品型号管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "型号"
// @Param productCategoryID query string false "产品类别ID"
// @Param IsPrefabricated query bool false "是否预制"
// @Success 200 {object} proto.QueryProductModelResponse
// @Router /api/mom/productbase/productmodel/query [get]
func QueryProductModel(c *gin.Context) {
	req := &proto.QueryProductModelRequest{}
	resp := &proto.QueryProductModelResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductModel(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductModel godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 产品型号管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductModelResponse
// @Router /api/mom/productbase/productmodel/all [get]
func GetAllProductModel(c *gin.Context) {
	resp := &proto.GetAllProductModelResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductModels()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductModelsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductModelDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 产品型号管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductModelDetailResponse
// @Router /api/mom/productbase/productmodel/detail [get]
func GetProductModelDetail(c *gin.Context) {
	resp := &proto.GetProductModelDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductModelByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductModelToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductModel godoc
// @Summary 删除
// @Description 删除
// @Tags 产品型号管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductModel"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productbase/productmodel/delete [delete]
func DeleteProductModel(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除产品型号请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductModel(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// ParseProductModel godoc
// @Summary 解析
// @Description 解析
// @Tags 产品型号管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Parse ProductModel"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productbase/productmodel/parse [post]
func ParseProductModel(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,解析产品型号请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.ParamProductModelByID(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductModelRouter(r *gin.Engine) {
	g := r.Group("/api/mom/productbase/productmodel")

	g.POST("add", AddProductModel)
	g.PUT("update", UpdateProductModel)
	g.GET("query", QueryProductModel)
	g.DELETE("delete", DeleteProductModel)
	g.GET("all", GetAllProductModel)
	g.GET("detail", GetProductModelDetail)
	g.POST("parse", ParseProductModel)
}
