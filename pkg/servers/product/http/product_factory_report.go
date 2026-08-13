package http

import (
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/product/logic"
	"github.com/gin-gonic/gin"
)

// QueryProductFactoryReport godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 产品出厂报告管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productOrderNo query string false "生产工单号"
// @Param productSerialNo query string false "产品序列号"
// @Param finishedTime0 query string false "下线时间开始"
// @Param finishedTime1 query string false "下线时间结束"
// @Success 200 {object} proto.QueryProductInfoResponse
// @Router /api/mom/product/productfactoryreport/query [get]
func QueryProductFactoryReport(c *gin.Context) {
	req := &proto.QueryProductInfoRequest{}
	resp := &proto.QueryProductInfoResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductFactoryReport(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

func RegisterProductFactoryReportRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productfactoryreport")

	g.GET("query", QueryProductFactoryReport)
}
