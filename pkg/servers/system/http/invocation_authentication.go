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

// AddInvocationAuthentication godoc
// @Summary 新增
// @Description 新增
// @Tags 接口认证管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.InvocationAuthenticationInfo true "Add InvocationAuthentication"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/invocationauthentication/add [post]
func AddInvocationAuthentication(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.InvocationAuthenticationInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建接口认证请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateInvocationAuthentication(model.PBToInvocationAuthentication(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateInvocationAuthentication godoc
// @Summary 更新
// @Description 更新
// @Tags 接口认证管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.InvocationAuthenticationInfo true "Update InvocationAuthentication"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/invocationauthentication/update [put]
func UpdateInvocationAuthentication(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.InvocationAuthenticationInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新接口认证请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateInvocationAuthentication(model.PBToInvocationAuthentication(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryInvocationAuthentication godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 接口认证管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param name query string false "名称"
// @Success 200 {object} proto.QueryInvocationAuthenticationResponse
// @Router /api/mom/system/invocationauthentication/query [get]
func QueryInvocationAuthentication(c *gin.Context) {
	req := &proto.QueryInvocationAuthenticationRequest{}
	resp := &proto.QueryInvocationAuthenticationResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryInvocationAuthentication(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllInvocationAuthentication godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 接口认证管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllInvocationAuthenticationResponse
// @Router /api/mom/system/invocationauthentication/all [get]
func GetAllInvocationAuthentication(c *gin.Context) {
	resp := &proto.GetAllInvocationAuthenticationResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllInvocationAuthentications()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.InvocationAuthenticationsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetInvocationAuthenticationDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 接口认证管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetInvocationAuthenticationDetailResponse
// @Router /api/mom/system/invocationauthentication/detail [get]
func GetInvocationAuthenticationDetail(c *gin.Context) {
	resp := &proto.GetInvocationAuthenticationDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetInvocationAuthenticationByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.InvocationAuthenticationToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteInvocationAuthentication godoc
// @Summary 删除
// @Description 删除
// @Tags 接口认证管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete InvocationAuthentication"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/invocationauthentication/delete [delete]
func DeleteInvocationAuthentication(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除接口认证请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteInvocationAuthentication(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterInvocationAuthenticationRouter(r *gin.Engine) {
	g := r.Group("/api/mom/system/invocationauthentication")

	g.POST("add", AddInvocationAuthentication)
	g.PUT("update", UpdateInvocationAuthentication)
	g.GET("query", QueryInvocationAuthentication)
	g.DELETE("delete", DeleteInvocationAuthentication)
	g.GET("all", GetAllInvocationAuthentication)
	g.GET("detail", GetInvocationAuthenticationDetail)
}
