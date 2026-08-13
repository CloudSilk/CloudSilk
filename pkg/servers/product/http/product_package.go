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

// AddProductPackage godoc
// @Summary 新增
// @Description 新增
// @Tags 产品包装管理管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductPackageInfo true "Add ProductPackage"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productpackage/add [post]
func AddProductPackage(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductPackageInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建产品包装管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductPackage(model.PBToProductPackage(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductPackage godoc
// @Summary 更新
// @Description 更新
// @Tags 产品包装管理管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductPackageInfo true "Update ProductPackage"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productpackage/update [put]
func UpdateProductPackage(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductPackageInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新产品包装管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductPackage(model.PBToProductPackage(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductPackage godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 产品包装管理管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号"
// @Success 200 {object} proto.QueryProductPackageResponse
// @Router /api/mom/product/productpackage/query [get]
func QueryProductPackage(c *gin.Context) {
	req := &proto.QueryProductPackageRequest{}
	resp := &proto.QueryProductPackageResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductPackage(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductPackage godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 产品包装管理管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductPackageResponse
// @Router /api/mom/product/productpackage/all [get]
func GetAllProductPackage(c *gin.Context) {
	resp := &proto.GetAllProductPackageResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductPackages()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductPackagesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductPackageDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 产品包装管理管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductPackageDetailResponse
// @Router /api/mom/product/productpackage/detail [get]
func GetProductPackageDetail(c *gin.Context) {
	resp := &proto.GetProductPackageDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductPackageByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductPackageToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductPackage godoc
// @Summary 删除
// @Description 删除
// @Tags 产品包装管理管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductPackage"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productpackage/delete [delete]
func DeleteProductPackage(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除产品包装管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductPackage(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductPackageRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productpackage")

	g.POST("add", AddProductPackage)
	g.PUT("update", UpdateProductPackage)
	g.GET("query", QueryProductPackage)
	g.DELETE("delete", DeleteProductPackage)
	g.GET("all", GetAllProductPackage)
	g.GET("detail", GetProductPackageDetail)
}
