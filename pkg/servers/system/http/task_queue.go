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

// AddTaskQueue godoc
// @Summary 新增
// @Description 新增
// @Tags 任务队列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.TaskQueueInfo true "Add TaskQueue"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/taskqueue/add [post]
func AddTaskQueue(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.TaskQueueInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建任务队列请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateTaskQueue(model.PBToTaskQueue(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateTaskQueue godoc
// @Summary 更新
// @Description 更新
// @Tags 任务队列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.TaskQueueInfo true "Update TaskQueue"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/taskqueue/update [put]
func UpdateTaskQueue(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.TaskQueueInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新任务队列请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateTaskQueue(model.PBToTaskQueue(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryTaskQueue godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 任务队列管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Success 200 {object} proto.QueryTaskQueueResponse
// @Router /api/mom/system/taskqueue/query [get]
func QueryTaskQueue(c *gin.Context) {
	req := &proto.QueryTaskQueueRequest{}
	resp := &proto.QueryTaskQueueResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryTaskQueue(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllTaskQueue godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 任务队列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllTaskQueueResponse
// @Router /api/mom/system/taskqueue/all [get]
func GetAllTaskQueue(c *gin.Context) {
	resp := &proto.GetAllTaskQueueResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllTaskQueues()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.TaskQueuesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetTaskQueueDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 任务队列管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetTaskQueueDetailResponse
// @Router /api/mom/system/taskqueue/detail [get]
func GetTaskQueueDetail(c *gin.Context) {
	resp := &proto.GetTaskQueueDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetTaskQueueByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.TaskQueueToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteTaskQueue godoc
// @Summary 删除
// @Description 删除
// @Tags 任务队列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete TaskQueue"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/taskqueue/delete [delete]
func DeleteTaskQueue(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除任务队列请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteTaskQueue(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterTaskQueueRouter(r *gin.Engine) {
	g := r.Group("/api/mom/system/taskqueue")

	g.POST("add", AddTaskQueue)
	g.PUT("update", UpdateTaskQueue)
	g.GET("query", QueryTaskQueue)
	g.DELETE("delete", DeleteTaskQueue)
	g.GET("all", GetAllTaskQueue)
	g.GET("detail", GetTaskQueueDetail)
}
