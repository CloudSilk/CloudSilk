package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/servers/production/logic"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/pkg/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddProcessStepMatchRule godoc
// @Summary 新增
// @Description 新增
// @Tags 工步匹配规则管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProcessStepMatchRuleInfo true "Add ProcessStepMatchRule"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/processstepmatchrule/add [post]
func AddProcessStepMatchRule(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProcessStepMatchRuleInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建工步匹配规则请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProcessStepMatchRule(model.PBToProcessStepMatchRule(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProcessStepMatchRule godoc
// @Summary 更新
// @Description 更新
// @Tags 工步匹配规则管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProcessStepMatchRuleInfo true "Update ProcessStepMatchRule"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/processstepmatchrule/update [put]
func UpdateProcessStepMatchRule(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProcessStepMatchRuleInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新工步匹配规则请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProcessStepMatchRule(model.PBToProcessStepMatchRule(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProcessStepMatchRule godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 工步匹配规则管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Success 200 {object} proto.QueryProcessStepMatchRuleResponse
// @Router /api/mom/production/processstepmatchrule/query [get]
func QueryProcessStepMatchRule(c *gin.Context) {
	req := &proto.QueryProcessStepMatchRuleRequest{}
	resp := &proto.QueryProcessStepMatchRuleResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProcessStepMatchRule(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProcessStepMatchRule godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 工步匹配规则管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProcessStepMatchRuleResponse
// @Router /api/mom/production/processstepmatchrule/all [get]
func GetAllProcessStepMatchRule(c *gin.Context) {
	resp := &proto.GetAllProcessStepMatchRuleResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProcessStepMatchRules()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProcessStepMatchRulesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProcessStepMatchRuleDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 工步匹配规则管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProcessStepMatchRuleDetailResponse
// @Router /api/mom/production/processstepmatchrule/detail [get]
func GetProcessStepMatchRuleDetail(c *gin.Context) {
	resp := &proto.GetProcessStepMatchRuleDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProcessStepMatchRuleByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProcessStepMatchRuleToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProcessStepMatchRule godoc
// @Summary 删除
// @Description 删除
// @Tags 工步匹配规则管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProcessStepMatchRule"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/processstepmatchrule/delete [delete]
func DeleteProcessStepMatchRule(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除工步匹配规则请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProcessStepMatchRule(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProcessStepMatchRuleRouter(r *gin.Engine) {
	g := r.Group("/api/mom/production/processstepmatchrule")

	g.POST("add", AddProcessStepMatchRule)
	g.PUT("update", UpdateProcessStepMatchRule)
	g.GET("query", QueryProcessStepMatchRule)
	g.DELETE("delete", DeleteProcessStepMatchRule)
	g.GET("all", GetAllProcessStepMatchRule)
	g.GET("detail", GetProcessStepMatchRuleDetail)
}
