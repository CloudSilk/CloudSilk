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

// AddProcessStepType godoc
// @Summary 新增
// @Description 新增
// @Tags 生产工步类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProcessStepTypeInfo true "Add ProcessStepType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/processsteptype/add [post]
func AddProcessStepType(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProcessStepTypeInfo{}
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

	id, err := logic.CreateProcessStepType(model.PBToProcessStepType(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProcessStepType godoc
// @Summary 更新
// @Description 更新
// @Tags 生产工步类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProcessStepTypeInfo true "Update ProcessStepType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/processsteptype/update [put]
func UpdateProcessStepType(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProcessStepTypeInfo{}
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
	err = logic.UpdateProcessStepType(model.PBToProcessStepType(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProcessStepType godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 生产工步类型管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Success 200 {object} proto.QueryProcessStepTypeResponse
// @Router /api/mom/production/processsteptype/query [get]
func QueryProcessStepType(c *gin.Context) {
	req := &proto.QueryProcessStepTypeRequest{}
	resp := &proto.QueryProcessStepTypeResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProcessStepType(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProcessStepType godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 生产工步类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProcessStepTypeResponse
// @Router /api/mom/production/processsteptype/all [get]
func GetAllProcessStepType(c *gin.Context) {
	resp := &proto.GetAllProcessStepTypeResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProcessStepTypes()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProcessStepTypesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProcessStepTypeDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 生产工步类型管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProcessStepTypeDetailResponse
// @Router /api/mom/production/processsteptype/detail [get]
func GetProcessStepTypeDetail(c *gin.Context) {
	resp := &proto.GetProcessStepTypeDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProcessStepTypeByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProcessStepTypeToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProcessStepType godoc
// @Summary 删除
// @Description 删除
// @Tags 生产工步类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProcessStepType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/processsteptype/delete [delete]
func DeleteProcessStepType(c *gin.Context) {
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
	err = logic.DeleteProcessStepType(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProcessStepTypeRouter(r *gin.Engine) {
	g := r.Group("/api/mom/production/processsteptype")

	g.POST("add", AddProcessStepType)
	g.PUT("update", UpdateProcessStepType)
	g.GET("query", QueryProcessStepType)
	g.DELETE("delete", DeleteProcessStepType)
	g.GET("all", GetAllProcessStepType)
	g.GET("detail", GetProcessStepTypeDetail)
}
