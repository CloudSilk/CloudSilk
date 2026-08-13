package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/system/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddSystemEvent godoc
// @Summary 新增
// @Description 新增
// @Tags 系统事件管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.SystemEventInfo true "Add SystemEvent"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/systemevent/add [post]
func AddSystemEvent(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.SystemEventInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建系统事件请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateSystemEvent(model.PBToSystemEvent(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateSystemEvent godoc
// @Summary 更新
// @Description 更新
// @Tags 系统事件管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.SystemEventInfo true "Update SystemEvent"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/systemevent/update [put]
func UpdateSystemEvent(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.SystemEventInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新系统事件请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateSystemEvent(model.PBToSystemEvent(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QuerySystemEvent godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 系统事件管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Success 200 {object} proto.QuerySystemEventResponse
// @Router /api/mom/system/systemevent/query [get]
func QuerySystemEvent(c *gin.Context) {
	req := &proto.QuerySystemEventRequest{}
	resp := &proto.QuerySystemEventResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QuerySystemEvent(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllSystemEvent godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 系统事件管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllSystemEventResponse
// @Router /api/mom/system/systemevent/all [get]
func GetAllSystemEvent(c *gin.Context) {
	resp := &proto.GetAllSystemEventResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllSystemEvents()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.SystemEventsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetSystemEventDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 系统事件管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetSystemEventDetailResponse
// @Router /api/mom/system/systemevent/detail [get]
func GetSystemEventDetail(c *gin.Context) {
	resp := &proto.GetSystemEventDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetSystemEventByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.SystemEventToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteSystemEvent godoc
// @Summary 删除
// @Description 删除
// @Tags 系统事件管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete SystemEvent"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/systemevent/delete [delete]
func DeleteSystemEvent(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除系统事件请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteSystemEvent(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterSystemEventRouter(r *gin.Engine) {
	g := r.Group("/api/mom/system/systemevent")

	g.POST("add", AddSystemEvent)
	g.PUT("update", UpdateSystemEvent)
	g.GET("query", QuerySystemEvent)
	g.DELETE("delete", DeleteSystemEvent)
	g.GET("all", GetAllSystemEvent)
	g.GET("detail", GetSystemEventDetail)
}
