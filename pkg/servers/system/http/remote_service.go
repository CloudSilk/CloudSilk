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

// AddRemoteService godoc
// @Summary 新增
// @Description 新增
// @Tags 远程服务管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.RemoteServiceInfo true "Add RemoteService"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/remoteservice/add [post]
func AddRemoteService(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.RemoteServiceInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建远程服务管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateRemoteService(model.PBToRemoteService(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateRemoteService godoc
// @Summary 更新
// @Description 更新
// @Tags 远程服务管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.RemoteServiceInfo true "Update RemoteService"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/remoteservice/update [put]
func UpdateRemoteService(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.RemoteServiceInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新远程服务管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateRemoteService(model.PBToRemoteService(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryRemoteService godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 远程服务管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param name query string false "名称"
// @Success 200 {object} proto.QueryRemoteServiceResponse
// @Router /api/mom/system/remoteservice/query [get]
func QueryRemoteService(c *gin.Context) {
	req := &proto.QueryRemoteServiceRequest{}
	resp := &proto.QueryRemoteServiceResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryRemoteService(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllRemoteService godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 远程服务管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllRemoteServiceResponse
// @Router /api/mom/system/remoteservice/all [get]
func GetAllRemoteService(c *gin.Context) {
	resp := &proto.GetAllRemoteServiceResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllRemoteServices()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.RemoteServicesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetRemoteServiceDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 远程服务管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetRemoteServiceDetailResponse
// @Router /api/mom/system/remoteservice/detail [get]
func GetRemoteServiceDetail(c *gin.Context) {
	resp := &proto.GetRemoteServiceDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetRemoteServiceByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.RemoteServiceToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteRemoteService godoc
// @Summary 删除
// @Description 删除
// @Tags 远程服务管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete RemoteService"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/remoteservice/delete [delete]
func DeleteRemoteService(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除远程服务管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteRemoteService(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterRemoteServiceRouter(r *gin.Engine) {
	g := r.Group("/api/mom/system/remoteservice")

	g.POST("add", AddRemoteService)
	g.PUT("update", UpdateRemoteService)
	g.GET("query", QueryRemoteService)
	g.DELETE("delete", DeleteRemoteService)
	g.GET("all", GetAllRemoteService)
	g.GET("detail", GetRemoteServiceDetail)
}
