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

// AddProductAttribute godoc
// @Summary 新增
// @Description 新增
// @Tags 产品特性管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductAttributeInfo true "Add ProductAttribute"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productbase/productattribute/add [post]
func AddProductAttribute(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductAttributeInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建产品特性请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	for _, _productAttributeIdentifier := range req.ProductAttributeIdentifiers {
		if len(_productAttributeIdentifier.AvailableCategoryIDs) > 0 {
			var AvailableCategorys []*proto.ProductAttributeIdentifierAvailableCategoryInfo
			for _, categoryID := range _productAttributeIdentifier.AvailableCategoryIDs {
				AvailableCategorys = append(AvailableCategorys, &proto.ProductAttributeIdentifierAvailableCategoryInfo{
					ProductCategoryID: categoryID,
				})
			}
			_productAttributeIdentifier.ProductAttributeIdentifierAvailableCategorys = AvailableCategorys
		}
	}

	id, err := logic.CreateProductAttribute(model.PBToProductAttribute(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductAttribute godoc
// @Summary 更新
// @Description 更新
// @Tags 产品特性管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductAttributeInfo true "Update ProductAttribute"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productbase/productattribute/update [put]
func UpdateProductAttribute(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductAttributeInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新产品特性请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	for _, _productAttributeIdentifier := range req.ProductAttributeIdentifiers {
		if len(_productAttributeIdentifier.AvailableCategoryIDs) > 0 {
			var AvailableCategorys []*proto.ProductAttributeIdentifierAvailableCategoryInfo
			for _, categoryID := range _productAttributeIdentifier.AvailableCategoryIDs {
				AvailableCategorys = append(AvailableCategorys, &proto.ProductAttributeIdentifierAvailableCategoryInfo{
					ProductCategoryID: categoryID,
				})
			}
			_productAttributeIdentifier.ProductAttributeIdentifierAvailableCategorys = AvailableCategorys
		}
	}

	err = logic.UpdateProductAttribute(model.PBToProductAttribute(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductAttribute godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 产品特性管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Success 200 {object} proto.QueryProductAttributeResponse
// @Router /api/mom/productbase/productattribute/query [get]
func QueryProductAttribute(c *gin.Context) {
	req := &proto.QueryProductAttributeRequest{}
	resp := &proto.QueryProductAttributeResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductAttribute(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductAttribute godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 产品特性管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductAttributeResponse
// @Router /api/mom/productbase/productattribute/all [get]
func GetAllProductAttribute(c *gin.Context) {
	resp := &proto.GetAllProductAttributeResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductAttributes()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductAttributesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductAttributeDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 产品特性管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductAttributeDetailResponse
// @Router /api/mom/productbase/productattribute/detail [get]
func GetProductAttributeDetail(c *gin.Context) {
	resp := &proto.GetProductAttributeDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductAttributeByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductAttributeToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductAttribute godoc
// @Summary 删除
// @Description 删除
// @Tags 产品特性管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductAttribute"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productbase/productattribute/delete [delete]
func DeleteProductAttribute(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除产品特性请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductAttribute(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductAttributeRouter(r *gin.Engine) {
	g := r.Group("/api/mom/productbase/productattribute")

	g.POST("add", AddProductAttribute)
	g.PUT("update", UpdateProductAttribute)
	g.GET("query", QueryProductAttribute)
	g.DELETE("delete", DeleteProductAttribute)
	g.GET("all", GetAllProductAttribute)
	g.GET("detail", GetProductAttributeDetail)
}
