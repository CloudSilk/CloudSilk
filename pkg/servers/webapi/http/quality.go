package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/servers/webapi/logic"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/pkg/utils/middleware"
	"github.com/gin-gonic/gin"
)

// GetTestProjectWithParameter godoc
// @Summary 获取测试项接口数据接口
// @Description 获取测试项接口数据接口
// @Tags WebAPI
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.GetTestProjectWithParameterRequest true "GetTestProjectWithParameterRequest"
// @Success 200 {object} proto.GetTestProjectWithParameterResponse
// @Router /api/mom/webapi/quality/getTestProjectWithParameter [post]
func GetTestProjectWithParameter(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.GetTestProjectWithParameterRequest{}
	resp := &proto.GetTestProjectWithParameterResponse{Code: 200}

	if err := c.BindJSON(req); err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,获取测试项接口数据参数无效:%v", transID, err)
		return
	}

	if err := middleware.Validate.Struct(req); err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	resp, err := logic.GetTestProjectWithParameter(req)
	if err != nil {
		resp.Code = 400
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func RegisterQualityRouter(r *gin.Engine) {
	g := r.Group("/api/mom/webapi/quality")

	g.POST("getTestProjectWithParameter", GetTestProjectWithParameter)
}
