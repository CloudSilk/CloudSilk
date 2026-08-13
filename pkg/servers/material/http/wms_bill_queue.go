package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/clients"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/material/logic"
	"github.com/CloudSilk/pkg/utils/log"
	usercenter "github.com/CloudSilk/usercenter/proto"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddWMSBillQueue godoc
// @Summary 新增
// @Description 新增
// @Tags WMS过帐队列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.WMSBillQueueInfo true "Add WMSBillQueue"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/wmsbillqueue/add [post]
func AddWMSBillQueue(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.WMSBillQueueInfo{CreateUserID: middleware.GetUserID(c)}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建WMS过帐队列请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateWMSBillQueue(model.PBToWMSBillQueue(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateWMSBillQueue godoc
// @Summary 更新
// @Description 更新
// @Tags WMS过帐队列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.WMSBillQueueInfo true "Update WMSBillQueue"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/wmsbillqueue/update [put]
func UpdateWMSBillQueue(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.WMSBillQueueInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新WMS过帐队列请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateWMSBillQueue(model.PBToWMSBillQueue(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryWMSBillQueue godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags WMS过帐队列管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param createTime0 query string false "创建时间开始"
// @Param createTime1 query string false "创建时间结束"
// @Param materialStore query string false "物料仓库信息"
// @Success 200 {object} proto.QueryWMSBillQueueResponse
// @Router /api/mom/material/wmsbillqueue/query [get]
func QueryWMSBillQueue(c *gin.Context) {
	req := &proto.QueryWMSBillQueueRequest{}
	resp := &proto.QueryWMSBillQueueResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryWMSBillQueue(req, resp, false)
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

// GetAllWMSBillQueue godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags WMS过帐队列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllWMSBillQueueResponse
// @Router /api/mom/material/wmsbillqueue/all [get]
func GetAllWMSBillQueue(c *gin.Context) {
	resp := &proto.GetAllWMSBillQueueResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllWMSBillQueues()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.WMSBillQueuesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetWMSBillQueueDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags WMS过帐队列管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetWMSBillQueueDetailResponse
// @Router /api/mom/material/wmsbillqueue/detail [get]
func GetWMSBillQueueDetail(c *gin.Context) {
	resp := &proto.GetWMSBillQueueDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetWMSBillQueueByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.WMSBillQueueToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteWMSBillQueue godoc
// @Summary 删除
// @Description 删除
// @Tags WMS过帐队列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete WMSBillQueue"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/wmsbillqueue/delete [delete]
func DeleteWMSBillQueue(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除WMS过帐队列请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteWMSBillQueue(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterWMSBillQueueRouter(r *gin.Engine) {
	g := r.Group("/api/mom/material/wmsbillqueue")

	g.POST("add", AddWMSBillQueue)
	g.PUT("update", UpdateWMSBillQueue)
	g.GET("query", QueryWMSBillQueue)
	g.DELETE("delete", DeleteWMSBillQueue)
	g.GET("all", GetAllWMSBillQueue)
	g.GET("detail", GetWMSBillQueueDetail)
}
