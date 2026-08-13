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

// AddProductModelBom godoc
// @Summary 新增
// @Description 新增
// @Tags 产品型号Bom管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductModelBomInfo true "Add ProductModelBom"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productbase/productmodelbom/add [post]
func AddProductModelBom(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductModelBomInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建产品型号Bom请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductModelBom(model.PBToProductModelBom(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductModelBom godoc
// @Summary 更新
// @Description 更新
// @Tags 产品型号Bom管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductModelBomInfo true "Update ProductModelBom"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productbase/productmodelbom/update [put]
func UpdateProductModelBom(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductModelBomInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新产品型号Bom请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductModelBom(model.PBToProductModelBom(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductModelBom godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 产品型号Bom管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productCategoryID query string false "产品类别ID"
// @Param productModelID query string false "产品型号ID"
// @Success 200 {object} proto.QueryProductModelBomResponse
// @Router /api/mom/productbase/productmodelbom/query [get]
func QueryProductModelBom(c *gin.Context) {
	req := &proto.QueryProductModelBomRequest{}
	resp := &proto.QueryProductModelBomResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductModelBom(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductModelBom godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 产品型号Bom管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductModelBomResponse
// @Router /api/mom/productbase/productmodelbom/all [get]
func GetAllProductModelBom(c *gin.Context) {
	resp := &proto.GetAllProductModelBomResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductModelBoms()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductModelBomsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductModelBomDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 产品型号Bom管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductModelBomDetailResponse
// @Router /api/mom/productbase/productmodelbom/detail [get]
func GetProductModelBomDetail(c *gin.Context) {
	resp := &proto.GetProductModelBomDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductModelBomByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductModelBomToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductModelBom godoc
// @Summary 删除
// @Description 删除
// @Tags 产品型号Bom管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductModelBom"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productbase/productmodelbom/delete [delete]
func DeleteProductModelBom(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除产品型号Bom请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductModelBom(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductModelBomRouter(r *gin.Engine) {
	g := r.Group("/api/mom/productbase/productmodelbom")

	g.POST("add", AddProductModelBom)
	g.PUT("update", UpdateProductModelBom)
	g.GET("query", QueryProductModelBom)
	g.DELETE("delete", DeleteProductModelBom)
	g.GET("all", GetAllProductModelBom)
	g.GET("detail", GetProductModelBomDetail)
}
