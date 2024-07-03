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

// AddMaterialChannelLayer godoc
// @Summary 新增
// @Description 新增
// @Tags 料架通道管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialChannelLayerInfo true "Add MaterialChannelLayer"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialchannellayer/add [post]
func AddMaterialChannelLayer(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialChannelLayerInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建料架通道请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateMaterialChannelLayer(model.PBToMaterialChannelLayer(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateMaterialChannelLayer godoc
// @Summary 更新
// @Description 更新
// @Tags 料架通道管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialChannelLayerInfo true "Update MaterialChannelLayer"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialchannellayer/update [put]
func UpdateMaterialChannelLayer(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialChannelLayerInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新料架通道请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateMaterialChannelLayer(model.PBToMaterialChannelLayer(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryMaterialChannelLayer godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 料架通道管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Success 200 {object} proto.QueryMaterialChannelLayerResponse
// @Router /api/mom/material/materialchannellayer/query [get]
func QueryMaterialChannelLayer(c *gin.Context) {
	req := &proto.QueryMaterialChannelLayerRequest{}
	resp := &proto.QueryMaterialChannelLayerResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryMaterialChannelLayer(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllMaterialChannelLayer godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 料架通道管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllMaterialChannelLayerResponse
// @Router /api/mom/material/materialchannellayer/all [get]
func GetAllMaterialChannelLayer(c *gin.Context) {
	resp := &proto.GetAllMaterialChannelLayerResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllMaterialChannelLayers()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.MaterialChannelLayersToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetMaterialChannelLayerDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 料架通道管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetMaterialChannelLayerDetailResponse
// @Router /api/mom/material/materialchannellayer/detail [get]
func GetMaterialChannelLayerDetail(c *gin.Context) {
	resp := &proto.GetMaterialChannelLayerDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetMaterialChannelLayerByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialChannelLayerToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteMaterialChannelLayer godoc
// @Summary 删除
// @Description 删除
// @Tags 料架通道管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete MaterialChannelLayer"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialchannellayer/delete [delete]
func DeleteMaterialChannelLayer(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除料架通道请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteMaterialChannelLayer(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterMaterialChannelLayerRouter(r *gin.Engine) {
	g := r.Group("/api/mom/material/materialchannellayer")

	g.POST("add", AddMaterialChannelLayer)
	g.PUT("update", UpdateMaterialChannelLayer)
	g.GET("query", QueryMaterialChannelLayer)
	g.DELETE("delete", DeleteMaterialChannelLayer)
	g.GET("all", GetAllMaterialChannelLayer)
	g.GET("detail", GetMaterialChannelLayerDetail)
}
