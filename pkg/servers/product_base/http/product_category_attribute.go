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

// AddProductCategoryAttribute godoc
// @Summary 新增
// @Description 新增
// @Tags 产品类别特性管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductCategoryAttributeInfo true "Add ProductCategoryAttribute"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productbase/productcategoryattribute/add [post]
func AddProductCategoryAttribute(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductCategoryAttributeInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建产品类别特性请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductCategoryAttribute(model.PBToProductCategoryAttribute(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductCategoryAttribute godoc
// @Summary 更新
// @Description 更新
// @Tags 产品类别特性管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductCategoryAttributeInfo true "Update ProductCategoryAttribute"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productbase/productcategoryattribute/update [put]
func UpdateProductCategoryAttribute(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductCategoryAttributeInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新产品类别特性请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductCategoryAttribute(model.PBToProductCategoryAttribute(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductCategoryAttribute godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 产品类别特性管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productCategoryID query string false "产品类别ID"
// @Param productAtribute query string false "特性或描述"
// @Param productAttributeID query string false "产品特性ID"
// @Success 200 {object} proto.QueryProductCategoryAttributeResponse
// @Router /api/mom/productbase/productcategoryattribute/query [get]
func QueryProductCategoryAttribute(c *gin.Context) {
	req := &proto.QueryProductCategoryAttributeRequest{}
	resp := &proto.QueryProductCategoryAttributeResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductCategoryAttribute(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductCategoryAttribute godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 产品类别特性管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductCategoryAttributeResponse
// @Router /api/mom/productbase/productcategoryattribute/all [get]
func GetAllProductCategoryAttribute(c *gin.Context) {
	resp := &proto.GetAllProductCategoryAttributeResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductCategoryAttributes()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductCategoryAttributesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductCategoryAttributeDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 产品类别特性管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductCategoryAttributeDetailResponse
// @Router /api/mom/productbase/productcategoryattribute/detail [get]
func GetProductCategoryAttributeDetail(c *gin.Context) {
	resp := &proto.GetProductCategoryAttributeDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductCategoryAttributeByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductCategoryAttributeToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductCategoryAttribute godoc
// @Summary 删除
// @Description 删除
// @Tags 产品类别特性管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductCategoryAttribute"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productbase/productcategoryattribute/delete [delete]
func DeleteProductCategoryAttribute(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除产品类别特性请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductCategoryAttribute(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductCategoryAttributeRouter(r *gin.Engine) {
	g := r.Group("/api/mom/productbase/productcategoryattribute")

	g.POST("add", AddProductCategoryAttribute)
	g.PUT("update", UpdateProductCategoryAttribute)
	g.GET("query", QueryProductCategoryAttribute)
	g.DELETE("delete", DeleteProductCategoryAttribute)
	g.GET("all", GetAllProductCategoryAttribute)
	g.GET("detail", GetProductCategoryAttributeDetail)
}
