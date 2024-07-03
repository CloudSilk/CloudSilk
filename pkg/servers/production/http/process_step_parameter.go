package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/production/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddProcessStepParameter godoc
// @Summary 新增
// @Description 新增
// @Tags 工步参数管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProcessStepParameterInfo true "Add ProcessStepParameter"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/processstepparameter/add [post]
func AddProcessStepParameter(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProcessStepParameterInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建生产工步类型请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProcessStepParameter(model.PBToProcessStepParameter(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProcessStepParameter godoc
// @Summary 更新
// @Description 更新
// @Tags 工步参数管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProcessStepParameterInfo true "Update ProcessStepParameter"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/processstepparameter/update [put]
func UpdateProcessStepParameter(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProcessStepParameterInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新生产工步类型请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProcessStepParameter(model.PBToProcessStepParameter(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProcessStepParameter godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 工步参数管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Param productionLineID query string false "产线ID"
// @Success 200 {object} proto.QueryProcessStepParameterResponse
// @Router /api/mom/production/processstepparameter/query [get]
func QueryProcessStepParameter(c *gin.Context) {
	req := &proto.QueryProcessStepParameterRequest{}
	resp := &proto.QueryProcessStepParameterResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProcessStepParameter(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProcessStepParameter godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 工步参数管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProcessStepParameterResponse
// @Router /api/mom/production/processstepparameter/all [get]
func GetAllProcessStepParameter(c *gin.Context) {
	resp := &proto.GetAllProcessStepParameterResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProcessStepParameters()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProcessStepParametersToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProcessStepParameterDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 工步参数管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProcessStepParameterDetailResponse
// @Router /api/mom/production/processstepparameter/detail [get]
func GetProcessStepParameterDetail(c *gin.Context) {
	resp := &proto.GetProcessStepParameterDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProcessStepParameterByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProcessStepParameterToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProcessStepParameter godoc
// @Summary 删除
// @Description 删除
// @Tags 工步参数管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProcessStepParameter"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/processstepparameter/delete [delete]
func DeleteProcessStepParameter(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除生产工步类型请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProcessStepParameter(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// GetAllProcessStepParameterByProductionLineID godoc
// @Summary 查询产线下所有工步参数
// @Description 查询产线下所有工步参数
// @Tags 工步参数管理
// @Accept  json
// @Produce  json
// @Param productionLineID query string true "生产产线ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductionLineDetailResponse
// @Router /api/mom/production/processstepparameter/all/productionlineid [get]
func GetAllProcessStepParameterByProductionLineID(c *gin.Context) {
	resp := &proto.GetProductionLineDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("productionLineID")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetAllProcessStepParameterByProductionLineID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionLineToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProcessStepParameterRouter(r *gin.Engine) {
	g := r.Group("/api/mom/production/processstepparameter")

	g.POST("add", AddProcessStepParameter)
	g.PUT("update", UpdateProcessStepParameter)
	g.GET("query", QueryProcessStepParameter)
	g.DELETE("delete", DeleteProcessStepParameter)
	g.GET("all", GetAllProcessStepParameter)
	g.GET("detail", GetProcessStepParameterDetail)
	g.GET("all/productionlineid", GetAllProcessStepParameterByProductionLineID)
}
