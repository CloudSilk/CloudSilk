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

// AddTaskQueueExecution godoc
// @Summary 新增
// @Description 新增
// @Tags 任务队列执行管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.TaskQueueExecutionInfo true "Add TaskQueueExecution"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/taskqueueexecution/add [post]
func AddTaskQueueExecution(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.TaskQueueExecutionInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建任务队列执行请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateTaskQueueExecution(model.PBToTaskQueueExecution(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateTaskQueueExecution godoc
// @Summary 更新
// @Description 更新
// @Tags 任务队列执行管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.TaskQueueExecutionInfo true "Update TaskQueueExecution"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/taskqueueexecution/update [put]
func UpdateTaskQueueExecution(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.TaskQueueExecutionInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新任务队列执行请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateTaskQueueExecution(model.PBToTaskQueueExecution(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryTaskQueueExecution godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 任务队列执行管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param createTime0 query string false "创建时间开始"
// @Param createTime1 query string false "创建时间结束"
// @Param taskQueueID query string false "任务队列ID"
// @Param dataTrace query string false "数据跟踪或失败原因"
// @Success 200 {object} proto.QueryTaskQueueExecutionResponse
// @Router /api/mom/system/taskqueueexecution/query [get]
func QueryTaskQueueExecution(c *gin.Context) {
	req := &proto.QueryTaskQueueExecutionRequest{}
	resp := &proto.QueryTaskQueueExecutionResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryTaskQueueExecution(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllTaskQueueExecution godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 任务队列执行管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllTaskQueueExecutionResponse
// @Router /api/mom/system/taskqueueexecution/all [get]
func GetAllTaskQueueExecution(c *gin.Context) {
	resp := &proto.GetAllTaskQueueExecutionResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllTaskQueueExecutions()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.TaskQueueExecutionsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetTaskQueueExecutionDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 任务队列执行管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetTaskQueueExecutionDetailResponse
// @Router /api/mom/system/taskqueueexecution/detail [get]
func GetTaskQueueExecutionDetail(c *gin.Context) {
	resp := &proto.GetTaskQueueExecutionDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetTaskQueueExecutionByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.TaskQueueExecutionToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteTaskQueueExecution godoc
// @Summary 删除
// @Description 删除
// @Tags 任务队列执行管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete TaskQueueExecution"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/taskqueueexecution/delete [delete]
func DeleteTaskQueueExecution(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除任务队列执行请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteTaskQueueExecution(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterTaskQueueExecutionRouter(r *gin.Engine) {
	g := r.Group("/api/mom/system/taskqueueexecution")

	g.POST("add", AddTaskQueueExecution)
	g.PUT("update", UpdateTaskQueueExecution)
	g.GET("query", QueryTaskQueueExecution)
	g.DELETE("delete", DeleteTaskQueueExecution)
	g.GET("all", GetAllTaskQueueExecution)
	g.GET("detail", GetTaskQueueExecutionDetail)
}
