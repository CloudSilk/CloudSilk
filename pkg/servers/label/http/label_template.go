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

// AddLabelTemplate godoc
// @Summary 新增
// @Description 新增
// @Tags 标签模板管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.LabelTemplateInfo true "Add LabelTemplate"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/label/labeltemplate/add [post]
func AddLabelTemplate(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.LabelTemplateInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建标签模板管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateLabelTemplate(model.PBToLabelTemplate(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateLabelTemplate godoc
// @Summary 更新
// @Description 更新
// @Tags 标签模板管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.LabelTemplateInfo true "Update LabelTemplate"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/label/labeltemplate/update [put]
func UpdateLabelTemplate(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.LabelTemplateInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新标签模板管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateLabelTemplate(model.PBToLabelTemplate(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryLabelTemplate godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 标签模板管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "标签模版或描述"
// @Param labelTypeID query string false "标签类型ID"
// @Success 200 {object} proto.QueryLabelTemplateResponse
// @Router /api/mom/label/labeltemplate/query [get]
func QueryLabelTemplate(c *gin.Context) {
	req := &proto.QueryLabelTemplateRequest{}
	resp := &proto.QueryLabelTemplateResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryLabelTemplate(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllLabelTemplate godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 标签模板管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllLabelTemplateResponse
// @Router /api/mom/label/labeltemplate/all [get]
func GetAllLabelTemplate(c *gin.Context) {
	resp := &proto.GetAllLabelTemplateResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllLabelTemplates()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.LabelTemplatesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetLabelTemplateDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 标签模板管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetLabelTemplateDetailResponse
// @Router /api/mom/label/labeltemplate/detail [get]
func GetLabelTemplateDetail(c *gin.Context) {
	resp := &proto.GetLabelTemplateDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetLabelTemplateByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.LabelTemplateToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteLabelTemplate godoc
// @Summary 删除
// @Description 删除
// @Tags 标签模板管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete LabelTemplate"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/label/labeltemplate/delete [delete]
func DeleteLabelTemplate(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除标签模板管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteLabelTemplate(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterLabelTemplateRouter(r *gin.Engine) {
	g := r.Group("/api/mom/label/labeltemplate")

	g.POST("add", AddLabelTemplate)
	g.PUT("update", UpdateLabelTemplate)
	g.GET("query", QueryLabelTemplate)
	g.DELETE("delete", DeleteLabelTemplate)
	g.GET("all", GetAllLabelTemplate)
	g.GET("detail", GetLabelTemplateDetail)
}
