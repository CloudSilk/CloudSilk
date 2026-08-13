package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/webapi/logic"
	"github.com/CloudSilk/CloudSilk/pkg/types"
	"github.com/CloudSilk/pkg/utils/log"
	ucmiddleware "github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// GetProductionStationEfficiency godoc
// @Summary 获取工位效率统计
// @Description 获取工位效率统计
// @Tags WebAPI
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.GetProductionStationEfficiencyRequest true "GetProductionStationEfficiencyRequest"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/webapi/statistic/getproductionstationefficiency [post]
func GetProductionStationEfficiency(c *gin.Context) {
	transID := ucmiddleware.GetTransID(c)
	req := &proto.GetProductionStationEfficiencyRequest{}
	resp := map[string]interface{}{"code": types.ServiceResponseCodeSuccess, "message": "", "data": nil}

	var err error
	if err = c.BindJSON(req); err != nil {
		resp["code"] = types.ServiceResponseCodeFailure
		resp["message"] = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,请求获取工位效率统计接口参数无效:%v", transID, err)
		return
	}

	if err = ucmiddleware.Validate.Struct(req); err != nil {
		resp["code"] = types.ServiceResponseCodeFailure
		resp["message"] = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	data, err := logic.GetProductionStationEfficiency(req)
	if err != nil {
		resp["code"] = types.ServiceResponseCodeFailure
		resp["message"] = err.Error()
	} else {
		resp["data"] = data
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateProductionStationEfficiency godoc
// @Summary 更新工位效率统计
// @Description 更新工位效率统计
// @Tags WebAPI
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.UpdateProductionStationEfficiencyRequest true "UpdateProductionStationEfficiencyRequest"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/webapi/statistic/updateproductionstationefficiency [post]
func UpdateProductionStationEfficiency(c *gin.Context) {
	transID := ucmiddleware.GetTransID(c)
	req := &proto.UpdateProductionStationEfficiencyRequest{}
	resp := map[string]interface{}{"code": types.ServiceResponseCodeSuccess, "message": "", "data": nil}

	var err error
	if err = c.BindJSON(req); err != nil {
		resp["code"] = types.ServiceResponseCodeFailure
		resp["message"] = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,请求获取工位效率统计接口参数无效:%v", transID, err)
		return
	}

	if err = ucmiddleware.Validate.Struct(req); err != nil {
		resp["code"] = types.ServiceResponseCodeFailure
		resp["message"] = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	if err := logic.UpdateProductionStationEfficiency(req); err != nil {
		resp["code"] = types.ServiceResponseCodeFailure
		resp["message"] = err.Error()
	}

	c.JSON(http.StatusOK, resp)
}

func RegisterStatisticRouter(r *gin.Engine) {
	g := r.Group("/api/mom/webapi/statistic")

	g.POST("getproductionstationefficiency", GetProductionStationEfficiency)
	g.POST("updateproductionstationefficiency", UpdateProductionStationEfficiency)
}
