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

// AddRemoteServiceTask godoc
// @Summary 新增
// @Description 新增
// @Tags 远程任务管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.RemoteServiceTaskInfo true "Add RemoteServiceTask"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/remoteservicetask/add [post]
func AddRemoteServiceTask(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.RemoteServiceTaskInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建远程任务管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateRemoteServiceTask(model.PBToRemoteServiceTask(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateRemoteServiceTask godoc
// @Summary 更新
// @Description 更新
// @Tags 远程任务管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.RemoteServiceTaskInfo true "Update RemoteServiceTask"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/remoteservicetask/update [put]
func UpdateRemoteServiceTask(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.RemoteServiceTaskInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新远程任务管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateRemoteServiceTask(model.PBToRemoteServiceTask(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryRemoteServiceTask godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 远程任务管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Success 200 {object} proto.QueryRemoteServiceTaskResponse
// @Router /api/mom/system/remoteservicetask/query [get]
func QueryRemoteServiceTask(c *gin.Context) {
	req := &proto.QueryRemoteServiceTaskRequest{}
	resp := &proto.QueryRemoteServiceTaskResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryRemoteServiceTask(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllRemoteServiceTask godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 远程任务管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllRemoteServiceTaskResponse
// @Router /api/mom/system/remoteservicetask/all [get]
func GetAllRemoteServiceTask(c *gin.Context) {
	resp := &proto.GetAllRemoteServiceTaskResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllRemoteServiceTasks()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.RemoteServiceTasksToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetRemoteServiceTaskDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 远程任务管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetRemoteServiceTaskDetailResponse
// @Router /api/mom/system/remoteservicetask/detail [get]
func GetRemoteServiceTaskDetail(c *gin.Context) {
	resp := &proto.GetRemoteServiceTaskDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetRemoteServiceTaskByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.RemoteServiceTaskToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteRemoteServiceTask godoc
// @Summary 删除
// @Description 删除
// @Tags 远程任务管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete RemoteServiceTask"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/remoteservicetask/delete [delete]
func DeleteRemoteServiceTask(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除远程任务管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteRemoteServiceTask(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterRemoteServiceTaskRouter(r *gin.Engine) {
	g := r.Group("/api/mom/system/remoteservicetask")

	g.POST("add", AddRemoteServiceTask)
	g.PUT("update", UpdateRemoteServiceTask)
	g.GET("query", QueryRemoteServiceTask)
	g.DELETE("delete", DeleteRemoteServiceTask)
	g.GET("all", GetAllRemoteServiceTask)
	g.GET("detail", GetRemoteServiceTaskDetail)
}
