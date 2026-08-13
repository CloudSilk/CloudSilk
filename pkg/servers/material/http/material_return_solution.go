package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/material/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddMaterialReturnSolution godoc
// @Summary 新增
// @Description 新增
// @Tags 物料退料方案管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialReturnSolutionInfo true "Add MaterialReturnSolution"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialreturnsolution/add [post]
func AddMaterialReturnSolution(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialReturnSolutionInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建物料退料方案请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateMaterialReturnSolution(model.PBToMaterialReturnSolution(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateMaterialReturnSolution godoc
// @Summary 更新
// @Description 更新
// @Tags 物料退料方案管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialReturnSolutionInfo true "Update MaterialReturnSolution"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialreturnsolution/update [put]
func UpdateMaterialReturnSolution(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialReturnSolutionInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新物料退料方案请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateMaterialReturnSolution(model.PBToMaterialReturnSolution(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryMaterialReturnSolution godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 物料退料方案管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Success 200 {object} proto.QueryMaterialReturnSolutionResponse
// @Router /api/mom/material/materialreturnsolution/query [get]
func QueryMaterialReturnSolution(c *gin.Context) {
	req := &proto.QueryMaterialReturnSolutionRequest{}
	resp := &proto.QueryMaterialReturnSolutionResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryMaterialReturnSolution(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllMaterialReturnSolution godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 物料退料方案管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllMaterialReturnSolutionResponse
// @Router /api/mom/material/materialreturnsolution/all [get]
func GetAllMaterialReturnSolution(c *gin.Context) {
	resp := &proto.GetAllMaterialReturnSolutionResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllMaterialReturnSolutions()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.MaterialReturnSolutionsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetMaterialReturnSolutionDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 物料退料方案管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetMaterialReturnSolutionDetailResponse
// @Router /api/mom/material/materialreturnsolution/detail [get]
func GetMaterialReturnSolutionDetail(c *gin.Context) {
	resp := &proto.GetMaterialReturnSolutionDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetMaterialReturnSolutionByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialReturnSolutionToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteMaterialReturnSolution godoc
// @Summary 删除
// @Description 删除
// @Tags 物料退料方案管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete MaterialReturnSolution"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialreturnsolution/delete [delete]
func DeleteMaterialReturnSolution(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除物料退料方案请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteMaterialReturnSolution(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterMaterialReturnSolutionRouter(r *gin.Engine) {
	g := r.Group("/api/mom/material/materialreturnsolution")

	g.POST("add", AddMaterialReturnSolution)
	g.PUT("update", UpdateMaterialReturnSolution)
	g.GET("query", QueryMaterialReturnSolution)
	g.DELETE("delete", DeleteMaterialReturnSolution)
	g.GET("all", GetAllMaterialReturnSolution)
	g.GET("detail", GetMaterialReturnSolutionDetail)
}
