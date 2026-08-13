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

// AddProductCategory godoc
// @Summary 新增
// @Description 新增
// @Tags 产品类别管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductCategoryInfo true "Add ProductCategory"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productbase/productcategory/add [post]
func AddProductCategory(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductCategoryInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建产品类别请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductCategory(model.PBToProductCategory(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductCategory godoc
// @Summary 更新
// @Description 更新
// @Tags 产品类别管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductCategoryInfo true "Update ProductCategory"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productbase/productcategory/update [put]
func UpdateProductCategory(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductCategoryInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新产品类别请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductCategory(model.PBToProductCategory(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductCategory godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 产品类别管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Param productBrandID query string false "产品品牌ID"
// @Success 200 {object} proto.QueryProductCategoryResponse
// @Router /api/mom/productbase/productcategory/query [get]
func QueryProductCategory(c *gin.Context) {
	req := &proto.QueryProductCategoryRequest{}
	resp := &proto.QueryProductCategoryResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductCategory(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductCategory godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 产品类别管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductCategoryResponse
// @Router /api/mom/productbase/productcategory/all [get]
func GetAllProductCategory(c *gin.Context) {
	resp := &proto.GetAllProductCategoryResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductCategorys()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductCategorysToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductCategoryDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 产品类别管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductCategoryDetailResponse
// @Router /api/mom/productbase/productcategory/detail [get]
func GetProductCategoryDetail(c *gin.Context) {
	resp := &proto.GetProductCategoryDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductCategoryByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductCategoryToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductCategory godoc
// @Summary 删除
// @Description 删除
// @Tags 产品类别管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductCategory"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productbase/productcategory/delete [delete]
func DeleteProductCategory(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除产品类别请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductCategory(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductCategoryRouter(r *gin.Engine) {
	g := r.Group("/api/mom/productbase/productcategory")

	g.POST("add", AddProductCategory)
	g.PUT("update", UpdateProductCategory)
	g.GET("query", QueryProductCategory)
	g.DELETE("delete", DeleteProductCategory)
	g.GET("all", GetAllProductCategory)
	g.GET("detail", GetProductCategoryDetail)
}
