package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/clients"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/label/logic"
	"github.com/CloudSilk/pkg/utils/log"
	usercenter "github.com/CloudSilk/usercenter/proto"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddLabelPrintQueue godoc
// @Summary 新增
// @Description 新增
// @Tags 标签打印队列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.LabelPrintQueueInfo true "Add LabelPrintQueue"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/label/labelprintqueue/add [post]
func AddLabelPrintQueue(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.LabelPrintQueueInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建标签打印队列管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateLabelPrintQueue(model.PBToLabelPrintQueue(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateLabelPrintQueue godoc
// @Summary 更新
// @Description 更新
// @Tags 标签打印队列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.LabelPrintQueueInfo true "Update LabelPrintQueue"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/label/labelprintqueue/update [put]
func UpdateLabelPrintQueue(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.LabelPrintQueueInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新标签打印队列管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateLabelPrintQueue(model.PBToLabelPrintQueue(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryLabelPrintQueue godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 标签打印队列管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param taskNo query string false "任务编号"
// @Param createTime0 query string false "创建时间开始"
// @Param createTime1 query string false "创建时间结束"
// @Success 200 {object} proto.QueryLabelPrintQueueResponse
// @Router /api/mom/label/labelprintqueue/query [get]
func QueryLabelPrintQueue(c *gin.Context) {
	req := &proto.QueryLabelPrintQueueRequest{}
	resp := &proto.QueryLabelPrintQueueResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryLabelPrintQueue(req, resp, false)
	if resp.Code == proto.Code_Success {
		r := &usercenter.QueryUserRequest{}
		for _, u := range resp.Data {
			r.Ids = append(r.Ids, u.CreateUserID)
		}
		r.PageSize = int64(len(r.Ids))
		users, err := clients.UserClient.Query(context.Background(), r)
		if err == nil && users.Code == usercenter.Code_Success {
			for _, u := range resp.Data {
				for _, u2 := range users.Data {
					if u.CreateUserID == u2.Id {
						u.CreateUserName = u2.Nickname
					}
				}
			}
		}
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllLabelPrintQueue godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 标签打印队列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllLabelPrintQueueResponse
// @Router /api/mom/label/labelprintqueue/all [get]
func GetAllLabelPrintQueue(c *gin.Context) {
	resp := &proto.GetAllLabelPrintQueueResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllLabelPrintQueues()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.LabelPrintQueuesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetLabelPrintQueueDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 标签打印队列管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetLabelPrintQueueDetailResponse
// @Router /api/mom/label/labelprintqueue/detail [get]
func GetLabelPrintQueueDetail(c *gin.Context) {
	resp := &proto.GetLabelPrintQueueDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetLabelPrintQueueByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.LabelPrintQueueToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteLabelPrintQueue godoc
// @Summary 删除
// @Description 删除
// @Tags 标签打印队列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete LabelPrintQueue"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/label/labelprintqueue/delete [delete]
func DeleteLabelPrintQueue(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除标签打印队列管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteLabelPrintQueue(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterLabelPrintQueueRouter(r *gin.Engine) {
	g := r.Group("/api/mom/label/labelprintqueue")

	g.POST("add", AddLabelPrintQueue)
	g.PUT("update", UpdateLabelPrintQueue)
	g.GET("query", QueryLabelPrintQueue)
	g.DELETE("delete", DeleteLabelPrintQueue)
	g.GET("all", GetAllLabelPrintQueue)
	g.GET("detail", GetLabelPrintQueueDetail)
}
