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

// AddMaterialReturnCause godoc
// @Summary 新增
// @Description 新增
// @Tags 物料退料原因管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialReturnCauseInfo true "Add MaterialReturnCause"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialreturncause/add [post]
func AddMaterialReturnCause(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialReturnCauseInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建物料退料原因请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateMaterialReturnCause(model.PBToMaterialReturnCause(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateMaterialReturnCause godoc
// @Summary 更新
// @Description 更新
// @Tags 物料退料原因管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialReturnCauseInfo true "Update MaterialReturnCause"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialreturncause/update [put]
func UpdateMaterialReturnCause(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialReturnCauseInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新物料退料原因请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateMaterialReturnCause(model.PBToMaterialReturnCause(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryMaterialReturnCause godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 物料退料原因管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Success 200 {object} proto.QueryMaterialReturnCauseResponse
// @Router /api/mom/material/materialreturncause/query [get]
func QueryMaterialReturnCause(c *gin.Context) {
	req := &proto.QueryMaterialReturnCauseRequest{}
	resp := &proto.QueryMaterialReturnCauseResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryMaterialReturnCause(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllMaterialReturnCause godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 物料退料原因管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllMaterialReturnCauseResponse
// @Router /api/mom/material/materialreturncause/all [get]
func GetAllMaterialReturnCause(c *gin.Context) {
	resp := &proto.GetAllMaterialReturnCauseResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllMaterialReturnCauses()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.MaterialReturnCausesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetMaterialReturnCauseDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 物料退料原因管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetMaterialReturnCauseDetailResponse
// @Router /api/mom/material/materialreturncause/detail [get]
func GetMaterialReturnCauseDetail(c *gin.Context) {
	resp := &proto.GetMaterialReturnCauseDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetMaterialReturnCauseByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialReturnCauseToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteMaterialReturnCause godoc
// @Summary 删除
// @Description 删除
// @Tags 物料退料原因管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete MaterialReturnCause"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialreturncause/delete [delete]
func DeleteMaterialReturnCause(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除物料退料原因请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteMaterialReturnCause(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterMaterialReturnCauseRouter(r *gin.Engine) {
	g := r.Group("/api/mom/material/materialreturncause")

	g.POST("add", AddMaterialReturnCause)
	g.PUT("update", UpdateMaterialReturnCause)
	g.GET("query", QueryMaterialReturnCause)
	g.DELETE("delete", DeleteMaterialReturnCause)
	g.GET("all", GetAllMaterialReturnCause)
	g.GET("detail", GetMaterialReturnCauseDetail)
}
