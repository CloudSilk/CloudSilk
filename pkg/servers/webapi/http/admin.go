package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/webapi/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// TryLogin godoc
// @Summary 登录
// @Description 登录
// @Tags WebAPI
// @Accept  json
// @Produce  json
// @Param account body proto.LoginRequest true "TryLogin"
// @Success 200 {object} proto.ServiceResponse
// @Router /api/mom/webapi/admin/login [post]
func TryLogin(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.LoginRequest{}
	resp := &proto.ServiceResponse{Code: 200}

	if err := c.BindJSON(req); err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,登录请求参数无效:%v", transID, err)
		return
	}

	if err := middleware.Validate.Struct(req); err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	logic.TryLogin(req, resp)
	c.JSON(http.StatusOK, resp)
}

// TryLogout godoc
// @Summary 注销
// @Description 注销
// @Tags WebAPI
// @Accept  json
// @Produce  json
// @Param account body proto.LogoutRequest true "TryLogout"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/webapi/admin/logout [post]
func TryLogout(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.LogoutRequest{}
	resp := &proto.CommonResponse{Code: 200}

	if err := c.BindJSON(req); err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,注销请求参数无效:%v", transID, err)
		return
	}

	if err := middleware.Validate.Struct(req); err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	logic.TryLogout(req, resp)
	c.JSON(http.StatusOK, resp)
}

func RegisterAdminRouter(r *gin.Engine) {
	g := r.Group("/api/mom/webapi/admin")

	g.POST("login", TryLogin)
	g.POST("logout", TryLogout)
}
