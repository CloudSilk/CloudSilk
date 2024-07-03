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

// AddAGVTaskQueue godoc
// @Summary 新增
// @Description 新增
// @Tags AGV任务队列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.AGVTaskQueueInfo true "Add AGVTaskQueue"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/agvtaskqueue/add [post]
func AddAGVTaskQueue(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.AGVTaskQueueInfo{CreateUserID: middleware.GetUserID(c)}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建AGV任务队列请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateAGVTaskQueue(model.PBToAGVTaskQueue(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateAGVTaskQueue godoc
// @Summary 更新
// @Description 更新
// @Tags AGV任务队列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.AGVTaskQueueInfo true "Update AGVTaskQueue"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/agvtaskqueue/update [put]
func UpdateAGVTaskQueue(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.AGVTaskQueueInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新AGV任务队列请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateAGVTaskQueue(model.PBToAGVTaskQueue(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryAGVTaskQueue godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags AGV任务队列管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param createTime0 query string false "创建时间开始"
// @Param createTime1 query string false "创建时间结束"
// @Param taskNo query string false "任务编号"
// @Success 200 {object} proto.QueryAGVTaskQueueResponse
// @Router /api/mom/material/agvtaskqueue/query [get]
func QueryAGVTaskQueue(c *gin.Context) {
	req := &proto.QueryAGVTaskQueueRequest{}
	resp := &proto.QueryAGVTaskQueueResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryAGVTaskQueue(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllAGVTaskQueue godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags AGV任务队列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllAGVTaskQueueResponse
// @Router /api/mom/material/agvtaskqueue/all [get]
func GetAllAGVTaskQueue(c *gin.Context) {
	resp := &proto.GetAllAGVTaskQueueResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllAGVTaskQueues()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.AGVTaskQueuesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetAGVTaskQueueDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags AGV任务队列管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAGVTaskQueueDetailResponse
// @Router /api/mom/material/agvtaskqueue/detail [get]
func GetAGVTaskQueueDetail(c *gin.Context) {
	resp := &proto.GetAGVTaskQueueDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetAGVTaskQueueByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.AGVTaskQueueToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteAGVTaskQueue godoc
// @Summary 删除
// @Description 删除
// @Tags AGV任务队列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete AGVTaskQueue"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/agvtaskqueue/delete [delete]
func DeleteAGVTaskQueue(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除AGV任务队列请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteAGVTaskQueue(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterAGVTaskQueueRouter(r *gin.Engine) {
	g := r.Group("/api/mom/material/agvtaskqueue")

	g.POST("add", AddAGVTaskQueue)
	g.PUT("update", UpdateAGVTaskQueue)
	g.GET("query", QueryAGVTaskQueue)
	g.DELETE("delete", DeleteAGVTaskQueue)
	g.GET("all", GetAllAGVTaskQueue)
	g.GET("detail", GetAGVTaskQueueDetail)
}
