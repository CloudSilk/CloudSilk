package http

import (
	"context"
	"net/http"

	model "github.com/CloudSilk/CloudSilk/pkg/model"
	apipb "github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/system/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddCodingGeneration godoc
// @Summary 新增
// @Description 新增
// @Tags 编码序列记录
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.CodingGenerationInfo true "Add CodingGeneration"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/mom/codinggeneration/add [post]
func AddCodingGeneration(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.CodingGenerationInfo{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建编码序列记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateCodingGeneration(model.PBToCodingGeneration(req))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateCodingGeneration godoc
// @Summary 更新
// @Description 更新
// @Tags 编码序列记录
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.CodingGenerationInfo true "Update CodingGeneration"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/mom/codinggeneration/update [put]
func UpdateCodingGeneration(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.CodingGenerationInfo{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新编码序列记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateCodingGeneration(model.PBToCodingGeneration(req))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryCodingGeneration godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 编码序列记录
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Success 200 {object} apipb.QueryCodingGenerationResponse
// @Router /api/mom/codinggeneration/query [get]
func QueryCodingGeneration(c *gin.Context) {
	req := &apipb.QueryCodingGenerationRequest{}
	resp := &apipb.QueryCodingGenerationResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryCodingGeneration(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllCodingGeneration godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 编码序列记录
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetAllCodingGenerationResponse
// @Router /api/mom/codinggeneration/all [get]
func GetAllCodingGeneration(c *gin.Context) {
	resp := &apipb.GetAllCodingGenerationResponse{
		Code: apipb.Code_Success,
	}
	list, err := logic.GetAllCodingGenerations()
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.CodingGenerationsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetCodingGenerationDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 编码序列记录
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetCodingGenerationDetailResponse
// @Router /api/mom/codinggeneration/detail [get]
func GetCodingGenerationDetail(c *gin.Context) {
	resp := &apipb.GetCodingGenerationDetailResponse{
		Code: apipb.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = apipb.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetCodingGenerationByID(id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.CodingGenerationToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteCodingGeneration godoc
// @Summary 删除
// @Description 删除
// @Tags 编码序列记录
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body apipb.DelRequest true "Delete CodingGeneration"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/mom/codinggeneration/delete [delete]
func DeleteCodingGeneration(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.DelRequest{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,删除编码序列记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteCodingGeneration(req.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterCodingGenerationRouter(r *gin.Engine) {
	g := r.Group("/api/mom/codinggeneration")

	g.POST("add", AddCodingGeneration)
	g.PUT("update", UpdateCodingGeneration)
	g.GET("query", QueryCodingGeneration)
	g.DELETE("delete", DeleteCodingGeneration)
	g.GET("all", GetAllCodingGeneration)
	g.GET("detail", GetCodingGenerationDetail)
}
