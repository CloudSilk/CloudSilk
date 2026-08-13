package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/product/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	ucmiddleware "github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddProductOrderProcessStep godoc
// @Summary 新增
// @Description 新增
// @Tags 工步管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductOrderProcessStepInfo true "Add ProductOrderProcessStep"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderprocessstep/add [post]
func AddProductOrderProcessStep(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductOrderProcessStepInfo{
		CreateUserID: ucmiddleware.GetUserID(c),
	}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建工步请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductOrderProcessStep(model.PBToProductOrderProcessStep(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductOrderProcessStep godoc
// @Summary 更新
// @Description 更新
// @Tags 工步管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductOrderProcessStepInfo true "Update ProductOrderProcessStep"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderprocessstep/update [put]
func UpdateProductOrderProcessStep(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductOrderProcessStepInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新工步请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductOrderProcessStep(model.PBToProductOrderProcessStep(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductOrderProcessStep godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 工步管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productOrderProcessID query string false "工单工序ID"
// @Success 200 {object} proto.QueryProductOrderProcessStepResponse
// @Router /api/mom/product/productorderprocessstep/query [get]
func QueryProductOrderProcessStep(c *gin.Context) {
	req := &proto.QueryProductOrderProcessStepRequest{}
	resp := &proto.QueryProductOrderProcessStepResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductOrderProcessStep(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductOrderProcessStep godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 工步管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductOrderProcessStepResponse
// @Router /api/mom/product/productorderprocessstep/all [get]
func GetAllProductOrderProcessStep(c *gin.Context) {
	resp := &proto.GetAllProductOrderProcessStepResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductOrderProcessSteps()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductOrderProcessStepsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductOrderProcessStepDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 工步管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductOrderProcessStepDetailResponse
// @Router /api/mom/product/productorderprocessstep/detail [get]
func GetProductOrderProcessStepDetail(c *gin.Context) {
	resp := &proto.GetProductOrderProcessStepDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductOrderProcessStepByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductOrderProcessStepToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductOrderProcessStep godoc
// @Summary 删除
// @Description 删除
// @Tags 工步管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductOrderProcessStep"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderprocessstep/delete [delete]
func DeleteProductOrderProcessStep(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除工步请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductOrderProcessStep(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductOrderProcessStepRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productorderprocessstep")

	g.POST("add", AddProductOrderProcessStep)
	g.PUT("update", UpdateProductOrderProcessStep)
	g.GET("query", QueryProductOrderProcessStep)
	g.DELETE("delete", DeleteProductOrderProcessStep)
	g.GET("all", GetAllProductOrderProcessStep)
	g.GET("detail", GetProductOrderProcessStepDetail)
}
