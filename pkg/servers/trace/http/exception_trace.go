package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/clients"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/trace/logic"
	"github.com/CloudSilk/pkg/utils/log"
	usercenter "github.com/CloudSilk/usercenter/proto"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddExceptionTrace godoc
// @Summary 新增
// @Description 新增
// @Tags 系统异常日志管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ExceptionTraceInfo true "Add ExceptionTrace"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/trace/exceptiontrace/add [post]
func AddExceptionTrace(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ExceptionTraceInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建系统异常日志请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateExceptionTrace(model.PBToExceptionTrace(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateExceptionTrace godoc
// @Summary 更新
// @Description 更新
// @Tags 系统异常日志管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ExceptionTraceInfo true "Update ExceptionTrace"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/trace/exceptiontrace/update [put]
func UpdateExceptionTrace(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ExceptionTraceInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新系统异常日志请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateExceptionTrace(model.PBToExceptionTrace(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryExceptionTrace godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 系统异常日志管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param group query string false "分组"
// @Param message query string false "异常信息"
// @Param timeReported0 query string false "上报时间开始"
// @Param timeReported1 query string false "上报时间结束"
// @Success 200 {object} proto.QueryExceptionTraceResponse
// @Router /api/mom/trace/exceptiontrace/query [get]
func QueryExceptionTrace(c *gin.Context) {
	req := &proto.QueryExceptionTraceRequest{}
	resp := &proto.QueryExceptionTraceResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryExceptionTrace(req, resp, false)
	if resp.Code == proto.Code_Success {
		r := &usercenter.QueryUserRequest{}
		for _, u := range resp.Data {
			r.Ids = append(r.Ids, u.ReportUserID)
		}
		r.PageSize = int64(len(r.Ids))
		users, err := clients.UserClient.Query(context.Background(), r)
		if err == nil && users.Code == usercenter.Code_Success {
			for _, u := range resp.Data {
				for _, u2 := range users.Data {
					if u.ReportUserID == u2.Id {
						u.ReportUserName = u2.Nickname
					}
				}
			}
		}
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllExceptionTrace godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 系统异常日志管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllExceptionTraceResponse
// @Router /api/mom/trace/exceptiontrace/all [get]
func GetAllExceptionTrace(c *gin.Context) {
	resp := &proto.GetAllExceptionTraceResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllExceptionTraces()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ExceptionTracesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetExceptionTraceDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 系统异常日志管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetExceptionTraceDetailResponse
// @Router /api/mom/trace/exceptiontrace/detail [get]
func GetExceptionTraceDetail(c *gin.Context) {
	resp := &proto.GetExceptionTraceDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetExceptionTraceByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ExceptionTraceToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteExceptionTrace godoc
// @Summary 删除
// @Description 删除
// @Tags 系统异常日志管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ExceptionTrace"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/trace/exceptiontrace/delete [delete]
func DeleteExceptionTrace(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除系统异常日志请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteExceptionTrace(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterExceptionTraceRouter(r *gin.Engine) {
	g := r.Group("/api/mom/trace/exceptiontrace")
	g.POST("add", AddExceptionTrace)
	g.PUT("update", UpdateExceptionTrace)
	g.GET("query", QueryExceptionTrace)
	g.DELETE("delete", DeleteExceptionTrace)
	g.GET("all", GetAllExceptionTrace)
	g.GET("detail", GetExceptionTraceDetail)
}
