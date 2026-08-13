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

// AddRemoteServiceTaskQueue godoc
// @Summary 新增
// @Description 新增
// @Tags 远程任务队列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.RemoteServiceTaskQueueInfo true "Add RemoteServiceTaskQueue"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/remoteservicetaskqueue/add [post]
func AddRemoteServiceTaskQueue(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.RemoteServiceTaskQueueInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建远程任务队列请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateRemoteServiceTaskQueue(model.PBToRemoteServiceTaskQueue(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateRemoteServiceTaskQueue godoc
// @Summary 更新
// @Description 更新
// @Tags 远程任务队列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.RemoteServiceTaskQueueInfo true "Update RemoteServiceTaskQueue"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/remoteservicetaskqueue/update [put]
func UpdateRemoteServiceTaskQueue(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.RemoteServiceTaskQueueInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新远程任务队列请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateRemoteServiceTaskQueue(model.PBToRemoteServiceTaskQueue(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryRemoteServiceTaskQueue godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 远程任务队列管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param taskNo query string false "任务编号或请求内容或响应内容"
// @Param createTime0 query string false "创建时间开始"
// @Param createTime1 query string false "创建时间结束"
// @Success 200 {object} proto.QueryRemoteServiceTaskQueueResponse
// @Router /api/mom/system/remoteservicetaskqueue/query [get]
func QueryRemoteServiceTaskQueue(c *gin.Context) {
	req := &proto.QueryRemoteServiceTaskQueueRequest{}
	resp := &proto.QueryRemoteServiceTaskQueueResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryRemoteServiceTaskQueue(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllRemoteServiceTaskQueue godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 远程任务队列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllRemoteServiceTaskQueueResponse
// @Router /api/mom/system/remoteservicetaskqueue/all [get]
func GetAllRemoteServiceTaskQueue(c *gin.Context) {
	resp := &proto.GetAllRemoteServiceTaskQueueResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllRemoteServiceTaskQueues()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.RemoteServiceTaskQueuesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetRemoteServiceTaskQueueDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 远程任务队列管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetRemoteServiceTaskQueueDetailResponse
// @Router /api/mom/system/remoteservicetaskqueue/detail [get]
func GetRemoteServiceTaskQueueDetail(c *gin.Context) {
	resp := &proto.GetRemoteServiceTaskQueueDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetRemoteServiceTaskQueueByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.RemoteServiceTaskQueueToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteRemoteServiceTaskQueue godoc
// @Summary 删除
// @Description 删除
// @Tags 远程任务队列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete RemoteServiceTaskQueue"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/remoteservicetaskqueue/delete [delete]
func DeleteRemoteServiceTaskQueue(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除远程任务队列请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteRemoteServiceTaskQueue(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterRemoteServiceTaskQueueRouter(r *gin.Engine) {
	g := r.Group("/api/mom/system/remoteservicetaskqueue")

	g.POST("add", AddRemoteServiceTaskQueue)
	g.PUT("update", UpdateRemoteServiceTaskQueue)
	g.GET("query", QueryRemoteServiceTaskQueue)
	g.DELETE("delete", DeleteRemoteServiceTaskQueue)
	g.GET("all", GetAllRemoteServiceTaskQueue)
	g.GET("detail", GetRemoteServiceTaskQueueDetail)
}
