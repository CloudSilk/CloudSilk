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

// AddSystemEventTrigger godoc
// @Summary 新增
// @Description 新增
// @Tags 系统事件触发管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.SystemEventTriggerInfo true "Add SystemEventTrigger"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/systemeventtrigger/add [post]
func AddSystemEventTrigger(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.SystemEventTriggerInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建系统事件触发请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateSystemEventTrigger(model.PBToSystemEventTrigger(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateSystemEventTrigger godoc
// @Summary 更新
// @Description 更新
// @Tags 系统事件触发管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.SystemEventTriggerInfo true "Update SystemEventTrigger"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/systemeventtrigger/update [put]
func UpdateSystemEventTrigger(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.SystemEventTriggerInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新系统事件触发请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateSystemEventTrigger(model.PBToSystemEventTrigger(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QuerySystemEventTrigger godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 系统事件触发管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param eventNo query string false "事件编号"
// @Param createTime0 query string false "触发时间开始"
// @Param createTime1 query string false "触发时间结束"
// @Param systemEvent query string false "系统事件"
// @Success 200 {object} proto.QuerySystemEventTriggerResponse
// @Router /api/mom/system/systemeventtrigger/query [get]
func QuerySystemEventTrigger(c *gin.Context) {
	req := &proto.QuerySystemEventTriggerRequest{}
	resp := &proto.QuerySystemEventTriggerResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QuerySystemEventTrigger(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllSystemEventTrigger godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 系统事件触发管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllSystemEventTriggerResponse
// @Router /api/mom/system/systemeventtrigger/all [get]
func GetAllSystemEventTrigger(c *gin.Context) {
	resp := &proto.GetAllSystemEventTriggerResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllSystemEventTriggers()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.SystemEventTriggersToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetSystemEventTriggerDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 系统事件触发管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetSystemEventTriggerDetailResponse
// @Router /api/mom/system/systemeventtrigger/detail [get]
func GetSystemEventTriggerDetail(c *gin.Context) {
	resp := &proto.GetSystemEventTriggerDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetSystemEventTriggerByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.SystemEventTriggerToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteSystemEventTrigger godoc
// @Summary 删除
// @Description 删除
// @Tags 系统事件触发管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete SystemEventTrigger"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/systemeventtrigger/delete [delete]
func DeleteSystemEventTrigger(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除系统事件触发请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteSystemEventTrigger(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterSystemEventTriggerRouter(r *gin.Engine) {
	g := r.Group("/api/mom/system/systemeventtrigger")

	g.POST("add", AddSystemEventTrigger)
	g.PUT("update", UpdateSystemEventTrigger)
	g.GET("query", QuerySystemEventTrigger)
	g.DELETE("delete", DeleteSystemEventTrigger)
	g.GET("all", GetAllSystemEventTrigger)
	g.GET("detail", GetSystemEventTriggerDetail)
}
