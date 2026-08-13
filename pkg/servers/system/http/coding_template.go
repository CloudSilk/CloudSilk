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

// AddCodingTemplate godoc
// @Summary 新增
// @Description 新增
// @Tags 编码模版管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.CodingTemplateInfo true "Add CodingTemplate"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/mom/codingtemplate/add [post]
func AddCodingTemplate(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.CodingTemplateInfo{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建编码模版管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateCodingTemplate(model.PBToCodingTemplate(req))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateCodingTemplate godoc
// @Summary 更新
// @Description 更新
// @Tags 编码模版管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.CodingTemplateInfo true "Update CodingTemplate"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/mom/codingtemplate/update [put]
func UpdateCodingTemplate(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.CodingTemplateInfo{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新编码模版管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateCodingTemplate(model.PBToCodingTemplate(req))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryCodingTemplate godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 编码模版管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Success 200 {object} apipb.QueryCodingTemplateResponse
// @Router /api/mom/codingtemplate/query [get]
func QueryCodingTemplate(c *gin.Context) {
	req := &apipb.QueryCodingTemplateRequest{}
	resp := &apipb.QueryCodingTemplateResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryCodingTemplate(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllCodingTemplate godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 编码模版管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetAllCodingTemplateResponse
// @Router /api/mom/codingtemplate/all [get]
func GetAllCodingTemplate(c *gin.Context) {
	resp := &apipb.GetAllCodingTemplateResponse{
		Code: apipb.Code_Success,
	}
	list, err := logic.GetAllCodingTemplates()
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.CodingTemplatesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetCodingTemplateDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 编码模版管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetCodingTemplateDetailResponse
// @Router /api/mom/codingtemplate/detail [get]
func GetCodingTemplateDetail(c *gin.Context) {
	resp := &apipb.GetCodingTemplateDetailResponse{
		Code: apipb.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = apipb.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetCodingTemplateByID(id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.CodingTemplateToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteCodingTemplate godoc
// @Summary 删除
// @Description 删除
// @Tags 编码模版管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body apipb.DelRequest true "Delete CodingTemplate"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/mom/codingtemplate/delete [delete]
func DeleteCodingTemplate(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除编码模版管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteCodingTemplate(req.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterCodingTemplateRouter(r *gin.Engine) {
	g := r.Group("/api/mom/codingtemplate")

	g.POST("add", AddCodingTemplate)
	g.PUT("update", UpdateCodingTemplate)
	g.GET("query", QueryCodingTemplate)
	g.DELETE("delete", DeleteCodingTemplate)
	g.GET("all", GetAllCodingTemplate)
	g.GET("detail", GetCodingTemplateDetail)
}
