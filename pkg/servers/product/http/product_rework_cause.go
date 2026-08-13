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

// AddProductReworkCause godoc
// @Summary 新增
// @Description 新增
// @Tags 产品返工原因管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductReworkCauseInfo true "Add ProductReworkCause"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productreworkcause/add [post]
func AddProductReworkCause(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductReworkCauseInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建产品返工原因请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductReworkCause(model.PBToProductReworkCause(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductReworkCause godoc
// @Summary 更新
// @Description 更新
// @Tags 产品返工原因管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductReworkCauseInfo true "Update ProductReworkCause"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productreworkcause/update [put]
func UpdateProductReworkCause(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductReworkCauseInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新产品返工原因请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductReworkCause(model.PBToProductReworkCause(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductReworkCause godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 产品返工原因管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Success 200 {object} proto.QueryProductReworkCauseResponse
// @Router /api/mom/product/productreworkcause/query [get]
func QueryProductReworkCause(c *gin.Context) {
	req := &proto.QueryProductReworkCauseRequest{}
	resp := &proto.QueryProductReworkCauseResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductReworkCause(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductReworkCause godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 产品返工原因管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductReworkCauseResponse
// @Router /api/mom/product/productreworkcause/all [get]
func GetAllProductReworkCause(c *gin.Context) {
	resp := &proto.GetAllProductReworkCauseResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductReworkCauses()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductReworkCausesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductReworkCauseDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 产品返工原因管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductReworkCauseDetailResponse
// @Router /api/mom/product/productreworkcause/detail [get]
func GetProductReworkCauseDetail(c *gin.Context) {
	resp := &proto.GetProductReworkCauseDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductReworkCauseByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductReworkCauseToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductReworkCause godoc
// @Summary 删除
// @Description 删除
// @Tags 产品返工原因管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductReworkCause"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productreworkcause/delete [delete]
func DeleteProductReworkCause(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除产品返工原因请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductReworkCause(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductReworkCauseRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productreworkcause")

	g.POST("add", AddProductReworkCause)
	g.PUT("update", UpdateProductReworkCause)
	g.GET("query", QueryProductReworkCause)
	g.DELETE("delete", DeleteProductReworkCause)
	g.GET("all", GetAllProductReworkCause)
	g.GET("detail", GetProductReworkCauseDetail)
}
