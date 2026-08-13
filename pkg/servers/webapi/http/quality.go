package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/webapi/logic"
	"github.com/CloudSilk/pkg/utils/log"
	usercenter "github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// GetTestProjectWithParameter godoc
// @Summary 获取测试项接口数据
// @Description 根据测试工位/托盘号/产品序列号匹配测试项及其输入输出参数
// @Tags 测试项管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.GetTestProjectWithParameterRequest true "GetTestProjectWithParameterRequest"
// @Success 200 {object} proto.GetTestProjectWithParameterResponse
// @Router /api/mom/webapi/quality/gettestprojectwithparameter [post]
func GetTestProjectWithParameter(c *gin.Context) {
	transID := usercenter.GetTransID(c)
	req := &proto.GetTestProjectWithParameterRequest{}
	resp := &proto.GetTestProjectWithParameterResponse{}

	if err := c.BindJSON(req); err != nil {
		resp.Code = 1
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,获取测试项接口数据请求参数无效:%v", transID, err)
		return
	}

	data, err := logic.GetTestProjectWithParameter(req)
	if err != nil {
		resp.Code = 1
		resp.Message = err.Error()
	} else {
		resp = data
	}

	c.JSON(http.StatusOK, resp)
}

func RegisterQualityRouter(r *gin.Engine) {
	g := r.Group("/api/mom/webapi/quality")

	g.POST("gettestprojectwithparameter", GetTestProjectWithParameter)
}
