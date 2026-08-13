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

// AddMaterialStoreFeedRule godoc
// @Summary 新增
// @Description 新增
// @Tags 物料仓库补料规则
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialStoreFeedRuleInfo true "Add MaterialStoreFeedRule"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialstorefeedrule/add [post]
func AddMaterialStoreFeedRule(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialStoreFeedRuleInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建物料仓库补料规则请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateMaterialStoreFeedRule(model.PBToMaterialStoreFeedRule(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateMaterialStoreFeedRule godoc
// @Summary 更新
// @Description 更新
// @Tags 物料仓库补料规则
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialStoreFeedRuleInfo true "Update MaterialStoreFeedRule"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialstorefeedrule/update [put]
func UpdateMaterialStoreFeedRule(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialStoreFeedRuleInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新物料仓库补料规则请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateMaterialStoreFeedRule(model.PBToMaterialStoreFeedRule(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryMaterialStoreFeedRule godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 物料仓库补料规则
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param materialInfo query string false "物料号或描述"
// @Success 200 {object} proto.QueryMaterialStoreFeedRuleResponse
// @Router /api/mom/material/materialstorefeedrule/query [get]
func QueryMaterialStoreFeedRule(c *gin.Context) {
	req := &proto.QueryMaterialStoreFeedRuleRequest{}
	resp := &proto.QueryMaterialStoreFeedRuleResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryMaterialStoreFeedRule(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllMaterialStoreFeedRule godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 物料仓库补料规则
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllMaterialStoreFeedRuleResponse
// @Router /api/mom/material/materialstorefeedrule/all [get]
func GetAllMaterialStoreFeedRule(c *gin.Context) {
	resp := &proto.GetAllMaterialStoreFeedRuleResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllMaterialStoreFeedRules()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.MaterialStoreFeedRulesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetMaterialStoreFeedRuleDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 物料仓库补料规则
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetMaterialStoreFeedRuleDetailResponse
// @Router /api/mom/material/materialstorefeedrule/detail [get]
func GetMaterialStoreFeedRuleDetail(c *gin.Context) {
	resp := &proto.GetMaterialStoreFeedRuleDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetMaterialStoreFeedRuleByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialStoreFeedRuleToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteMaterialStoreFeedRule godoc
// @Summary 删除
// @Description 删除
// @Tags 物料仓库补料规则
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete MaterialStoreFeedRule"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialstorefeedrule/delete [delete]
func DeleteMaterialStoreFeedRule(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除物料仓库补料规则请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteMaterialStoreFeedRule(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterMaterialStoreFeedRuleRouter(r *gin.Engine) {
	g := r.Group("/api/mom/material/materialstorefeedrule")

	g.POST("add", AddMaterialStoreFeedRule)
	g.PUT("update", UpdateMaterialStoreFeedRule)
	g.GET("query", QueryMaterialStoreFeedRule)
	g.DELETE("delete", DeleteMaterialStoreFeedRule)
	g.GET("all", GetAllMaterialStoreFeedRule)
	g.GET("detail", GetMaterialStoreFeedRuleDetail)
}
