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

// AddOperationTrace godoc
// @Summary 新增
// @Description 新增
// @Tags 后台操作日志管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.OperationTraceInfo true "Add OperationTrace"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/trace/operationtrace/add [post]
func AddOperationTrace(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.OperationTraceInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建后台操作日志请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateOperationTrace(model.PBToOperationTrace(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateOperationTrace godoc
// @Summary 更新
// @Description 更新
// @Tags 后台操作日志管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.OperationTraceInfo true "Update OperationTrace"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/trace/operationtrace/update [put]
func UpdateOperationTrace(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.OperationTraceInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新后台操作日志请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateOperationTrace(model.PBToOperationTrace(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryOperationTrace godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 后台操作日志管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param group query string false "分组"
// @Param operateTime0 query string false "操作时间开始"
// @Param operateTime0 query string false "操作时间结束"
// @Success 200 {object} proto.QueryOperationTraceResponse
// @Router /api/mom/trace/operationtrace/query [get]
func QueryOperationTrace(c *gin.Context) {
	req := &proto.QueryOperationTraceRequest{}
	resp := &proto.QueryOperationTraceResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryOperationTrace(req, resp, false)
	if resp.Code == proto.Code_Success {
		r := &usercenter.QueryUserRequest{}
		for _, u := range resp.Data {
			r.Ids = append(r.Ids, u.OperateUserID)
		}
		r.PageSize = int64(len(r.Ids))
		users, err := clients.UserClient.Query(context.Background(), r)
		if err == nil && users.Code == usercenter.Code_Success {
			for _, u := range resp.Data {
				for _, u2 := range users.Data {
					if u.OperateUserID == u2.Id {
						u.OperateUserName = u2.Nickname
					}
				}
			}
		}
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllOperationTrace godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 后台操作日志管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllOperationTraceResponse
// @Router /api/mom/trace/operationtrace/all [get]
func GetAllOperationTrace(c *gin.Context) {
	resp := &proto.GetAllOperationTraceResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllOperationTraces()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.OperationTracesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetOperationTraceDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 后台操作日志管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetOperationTraceDetailResponse
// @Router /api/mom/trace/operationtrace/detail [get]
func GetOperationTraceDetail(c *gin.Context) {
	resp := &proto.GetOperationTraceDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetOperationTraceByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.OperationTraceToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteOperationTrace godoc
// @Summary 删除
// @Description 删除
// @Tags 后台操作日志管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete OperationTrace"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/trace/operationtrace/delete [delete]
func DeleteOperationTrace(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除后台操作日志请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteOperationTrace(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterOperationTraceRouter(r *gin.Engine) {
	g := r.Group("/api/mom/trace/operationtrace")

	g.POST("add", AddOperationTrace)
	g.PUT("update", UpdateOperationTrace)
	g.GET("query", QueryOperationTrace)
	g.DELETE("delete", DeleteOperationTrace)
	g.GET("all", GetAllOperationTrace)
	g.GET("detail", GetOperationTraceDetail)
}
