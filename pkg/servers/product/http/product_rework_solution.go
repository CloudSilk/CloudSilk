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

// AddProductReworkSolution godoc
// @Summary 新增
// @Description 新增
// @Tags 产品返工方案管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductReworkSolutionInfo true "Add ProductReworkSolution"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productreworksolution/add [post]
func AddProductReworkSolution(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductReworkSolutionInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建产品返工方案请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductReworkSolution(model.PBToProductReworkSolution(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductReworkSolution godoc
// @Summary 更新
// @Description 更新
// @Tags 产品返工方案管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductReworkSolutionInfo true "Update ProductReworkSolution"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productreworksolution/update [put]
func UpdateProductReworkSolution(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductReworkSolutionInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新产品返工方案请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductReworkSolution(model.PBToProductReworkSolution(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductReworkSolution godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 产品返工方案管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Success 200 {object} proto.QueryProductReworkSolutionResponse
// @Router /api/mom/product/productreworksolution/query [get]
func QueryProductReworkSolution(c *gin.Context) {
	req := &proto.QueryProductReworkSolutionRequest{}
	resp := &proto.QueryProductReworkSolutionResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductReworkSolution(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductReworkSolution godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 产品返工方案管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductReworkSolutionResponse
// @Router /api/mom/product/productreworksolution/all [get]
func GetAllProductReworkSolution(c *gin.Context) {
	resp := &proto.GetAllProductReworkSolutionResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductReworkSolutions()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductReworkSolutionsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductReworkSolutionDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 产品返工方案管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductReworkSolutionDetailResponse
// @Router /api/mom/product/productreworksolution/detail [get]
func GetProductReworkSolutionDetail(c *gin.Context) {
	resp := &proto.GetProductReworkSolutionDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductReworkSolutionByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductReworkSolutionToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductReworkSolution godoc
// @Summary 删除
// @Description 删除
// @Tags 产品返工方案管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductReworkSolution"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productreworksolution/delete [delete]
func DeleteProductReworkSolution(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除产品返工方案请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductReworkSolution(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductReworkSolutionRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productreworksolution")

	g.POST("add", AddProductReworkSolution)
	g.PUT("update", UpdateProductReworkSolution)
	g.GET("query", QueryProductReworkSolution)
	g.DELETE("delete", DeleteProductReworkSolution)
	g.GET("all", GetAllProductReworkSolution)
	g.GET("detail", GetProductReworkSolutionDetail)
}
