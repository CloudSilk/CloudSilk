package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/label/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddLabelAdaptationRule godoc
// @Summary 新增
// @Description 新增
// @Tags 标签适配规则管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.LabelAdaptationRuleInfo true "Add LabelAdaptationRule"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/label/labeladaptationrule/add [post]
func AddLabelAdaptationRule(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.LabelAdaptationRuleInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建标签适配规则管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateLabelAdaptationRule(model.PBToLabelAdaptationRule(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateLabelAdaptationRule godoc
// @Summary 更新
// @Description 更新
// @Tags 标签适配规则管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.LabelAdaptationRuleInfo true "Update LabelAdaptationRule"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/label/labeladaptationrule/update [put]
func UpdateLabelAdaptationRule(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.LabelAdaptationRuleInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新标签适配规则管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateLabelAdaptationRule(model.PBToLabelAdaptationRule(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryLabelAdaptationRule godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 标签适配规则管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Success 200 {object} proto.QueryLabelAdaptationRuleResponse
// @Router /api/mom/label/labeladaptationrule/query [get]
func QueryLabelAdaptationRule(c *gin.Context) {
	req := &proto.QueryLabelAdaptationRuleRequest{}
	resp := &proto.QueryLabelAdaptationRuleResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryLabelAdaptationRule(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllLabelAdaptationRule godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 标签适配规则管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllLabelAdaptationRuleResponse
// @Router /api/mom/label/labeladaptationrule/all [get]
func GetAllLabelAdaptationRule(c *gin.Context) {
	resp := &proto.GetAllLabelAdaptationRuleResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllLabelAdaptationRules()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.LabelAdaptationRulesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetLabelAdaptationRuleDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 标签适配规则管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetLabelAdaptationRuleDetailResponse
// @Router /api/mom/label/labeladaptationrule/detail [get]
func GetLabelAdaptationRuleDetail(c *gin.Context) {
	resp := &proto.GetLabelAdaptationRuleDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetLabelAdaptationRuleByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.LabelAdaptationRuleToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteLabelAdaptationRule godoc
// @Summary 删除
// @Description 删除
// @Tags 标签适配规则管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete LabelAdaptationRule"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/label/labeladaptationrule/delete [delete]
func DeleteLabelAdaptationRule(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除标签适配规则管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteLabelAdaptationRule(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterLabelAdaptationRuleRouter(r *gin.Engine) {
	g := r.Group("/api/mom/label/labeladaptationrule")

	g.POST("add", AddLabelAdaptationRule)
	g.PUT("update", UpdateLabelAdaptationRule)
	g.GET("query", QueryLabelAdaptationRule)
	g.DELETE("delete", DeleteLabelAdaptationRule)
	g.GET("all", GetAllLabelAdaptationRule)
	g.GET("detail", GetLabelAdaptationRuleDetail)
}
