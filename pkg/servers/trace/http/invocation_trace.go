package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/trace/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddInvocationTrace godoc
// @Summary 新增
// @Description 新增
// @Tags 接口调用日志管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.InvocationTraceInfo true "Add InvocationTrace"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/trace/invocationtrace/add [post]
func AddInvocationTrace(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.InvocationTraceInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建接口调用日志请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateInvocationTrace(model.PBToInvocationTrace(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateInvocationTrace godoc
// @Summary 更新
// @Description 更新
// @Tags 接口调用日志管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.InvocationTraceInfo true "Update InvocationTrace"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/trace/invocationtrace/update [put]
func UpdateInvocationTrace(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.InvocationTraceInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新接口调用日志请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateInvocationTrace(model.PBToInvocationTrace(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryInvocationTrace godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 接口调用日志管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param group query string false "分组"
// @Param requestTime0 query string false "请求时间开始"
// @Param requestTime1 query string false "请求时间结束"
// @Param actionName query string false "路由"
// @Param iPAddress query string false "IP地址"
// @Param requestText query string false "请求文本或响应文本"
// @Success 200 {object} proto.QueryInvocationTraceResponse
// @Router /api/mom/trace/invocationtrace/query [get]
func QueryInvocationTrace(c *gin.Context) {
	req := &proto.QueryInvocationTraceRequest{}
	resp := &proto.QueryInvocationTraceResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryInvocationTrace(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllInvocationTrace godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 接口调用日志管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllInvocationTraceResponse
// @Router /api/mom/trace/invocationtrace/all [get]
func GetAllInvocationTrace(c *gin.Context) {
	resp := &proto.GetAllInvocationTraceResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllInvocationTraces()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.InvocationTracesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetInvocationTraceDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 接口调用日志管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetInvocationTraceDetailResponse
// @Router /api/mom/trace/invocationtrace/detail [get]
func GetInvocationTraceDetail(c *gin.Context) {
	resp := &proto.GetInvocationTraceDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetInvocationTraceByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.InvocationTraceToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteInvocationTrace godoc
// @Summary 删除
// @Description 删除
// @Tags 接口调用日志管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete InvocationTrace"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/trace/invocationtrace/delete [delete]
func DeleteInvocationTrace(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除接口调用日志请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteInvocationTrace(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterInvocationTraceRouter(r *gin.Engine) {
	g := r.Group("/api/mom/trace/invocationtrace")

	g.POST("add", AddInvocationTrace)
	g.PUT("update", UpdateInvocationTrace)
	g.GET("query", QueryInvocationTrace)
	g.DELETE("delete", DeleteInvocationTrace)
	g.GET("all", GetAllInvocationTrace)
	g.GET("detail", GetInvocationTraceDetail)
}
