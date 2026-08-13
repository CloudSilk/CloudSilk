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

// AddProductPackageType godoc
// @Summary 新增
// @Description 新增
// @Tags 包装类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductPackageTypeInfo true "Add ProductPackageType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productpackagetype/add [post]
func AddProductPackageType(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductPackageTypeInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建包装类型请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductPackageType(model.PBToProductPackageType(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductPackageType godoc
// @Summary 更新
// @Description 更新
// @Tags 包装类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductPackageTypeInfo true "Update ProductPackageType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productpackagetype/update [put]
func UpdateProductPackageType(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductPackageTypeInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新包装类型请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductPackageType(model.PBToProductPackageType(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductPackageType godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 包装类型管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "类型或描述"
// @Success 200 {object} proto.QueryProductPackageTypeResponse
// @Router /api/mom/product/productpackagetype/query [get]
func QueryProductPackageType(c *gin.Context) {
	req := &proto.QueryProductPackageTypeRequest{}
	resp := &proto.QueryProductPackageTypeResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductPackageType(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductPackageType godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 包装类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductPackageTypeResponse
// @Router /api/mom/product/productpackagetype/all [get]
func GetAllProductPackageType(c *gin.Context) {
	resp := &proto.GetAllProductPackageTypeResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductPackageTypes()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductPackageTypesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductPackageTypeDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 包装类型管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductPackageTypeDetailResponse
// @Router /api/mom/product/productpackagetype/detail [get]
func GetProductPackageTypeDetail(c *gin.Context) {
	resp := &proto.GetProductPackageTypeDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductPackageTypeByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductPackageTypeToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductPackageType godoc
// @Summary 删除
// @Description 删除
// @Tags 包装类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductPackageType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productpackagetype/delete [delete]
func DeleteProductPackageType(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除包装类型请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductPackageType(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductPackageTypeRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productpackagetype")

	g.POST("add", AddProductPackageType)
	g.PUT("update", UpdateProductPackageType)
	g.GET("query", QueryProductPackageType)
	g.DELETE("delete", DeleteProductPackageType)
	g.GET("all", GetAllProductPackageType)
	g.GET("detail", GetProductPackageTypeDetail)
}
