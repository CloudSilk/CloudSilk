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

// AddProductBrand godoc
// @Summary 新增
// @Description 新增
// @Tags 产品品牌管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductBrandInfo true "Add ProductBrand"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productbase/productbrand/add [post]
func AddProductBrand(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductBrandInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建产品品牌请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductBrand(model.PBToProductBrand(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductBrand godoc
// @Summary 更新
// @Description 更新
// @Tags 产品品牌管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductBrandInfo true "Update ProductBrand"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productbase/productbrand/update [put]
func UpdateProductBrand(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductBrandInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新产品品牌请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductBrand(model.PBToProductBrand(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductBrand godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 产品品牌管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Success 200 {object} proto.QueryProductBrandResponse
// @Router /api/mom/productbase/productbrand/query [get]
func QueryProductBrand(c *gin.Context) {
	req := &proto.QueryProductBrandRequest{}
	resp := &proto.QueryProductBrandResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductBrand(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductBrand godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 产品品牌管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductBrandResponse
// @Router /api/mom/productbase/productbrand/all [get]
func GetAllProductBrand(c *gin.Context) {
	resp := &proto.GetAllProductBrandResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductBrands()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductBrandsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductBrandDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 产品品牌管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductBrandDetailResponse
// @Router /api/mom/productbase/productbrand/detail [get]
func GetProductBrandDetail(c *gin.Context) {
	resp := &proto.GetProductBrandDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductBrandByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductBrandToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductBrand godoc
// @Summary 删除
// @Description 删除
// @Tags 产品品牌管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductBrand"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productbase/productbrand/delete [delete]
func DeleteProductBrand(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除产品品牌请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductBrand(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductBrandRouter(r *gin.Engine) {
	g := r.Group("/api/mom/productbase/productbrand")

	g.POST("add", AddProductBrand)
	g.PUT("update", UpdateProductBrand)
	g.GET("query", QueryProductBrand)
	g.DELETE("delete", DeleteProductBrand)
	g.GET("all", GetAllProductBrand)
	g.GET("detail", GetProductBrandDetail)
}
