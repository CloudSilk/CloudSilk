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

// AddProductOrderPackage godoc
// @Summary 新增
// @Description 新增
// @Tags 工单包装管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductOrderPackageInfo true "Add ProductOrderPackage"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderpackage/add [post]
func AddProductOrderPackage(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductOrderPackageInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建工单包装管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductOrderPackage(model.PBToProductOrderPackage(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductOrderPackage godoc
// @Summary 更新
// @Description 更新
// @Tags 工单包装管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductOrderPackageInfo true "Update ProductOrderPackage"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderpackage/update [put]
func UpdateProductOrderPackage(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductOrderPackageInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新工单包装管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductOrderPackage(model.PBToProductOrderPackage(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductOrderPackage godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 工单包装管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param packageNo query string false "包装箱号"
// @Param palletNo query string false "栈板标识"
// @Param currentState query string false "当前状态"
// @Param createTime0 query string false "创建时间开始"
// @Param createTime1 query string false "创建时间结束"
// @Param productOrderNo query string false "生产工单号"
// @Success 200 {object} proto.QueryProductOrderPackageResponse
// @Router /api/mom/product/productorderpackage/query [get]
func QueryProductOrderPackage(c *gin.Context) {
	req := &proto.QueryProductOrderPackageRequest{}
	resp := &proto.QueryProductOrderPackageResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductOrderPackage(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductOrderPackage godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 工单包装管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductOrderPackageResponse
// @Router /api/mom/product/productorderpackage/all [get]
func GetAllProductOrderPackage(c *gin.Context) {
	resp := &proto.GetAllProductOrderPackageResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductOrderPackages()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductOrderPackagesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductOrderPackageDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 工单包装管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductOrderPackageDetailResponse
// @Router /api/mom/product/productorderpackage/detail [get]
func GetProductOrderPackageDetail(c *gin.Context) {
	resp := &proto.GetProductOrderPackageDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductOrderPackageByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductOrderPackageToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductOrderPackage godoc
// @Summary 删除
// @Description 删除
// @Tags 工单包装管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductOrderPackage"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderpackage/delete [delete]
func DeleteProductOrderPackage(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除工单包装管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductOrderPackage(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductOrderPackageRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productorderpackage")

	g.POST("add", AddProductOrderPackage)
	g.PUT("update", UpdateProductOrderPackage)
	g.GET("query", QueryProductOrderPackage)
	g.DELETE("delete", DeleteProductOrderPackage)
	g.GET("all", GetAllProductOrderPackage)
	g.GET("detail", GetProductOrderPackageDetail)
}
