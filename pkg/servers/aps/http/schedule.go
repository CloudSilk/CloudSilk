package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/aps/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// GenerateSchedule godoc
// @Summary 生成排程计划
// @Description 按产线对已发放/已签派工单执行前向贪心排程（优先级→交期），生成计划与明细
// @Tags APS排程
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.GenerateScheduleRequest true "GenerateScheduleRequest"
// @Success 200 {object} proto.GenerateScheduleResponse
// @Router /api/mom/aps/schedule/generate [post]
func GenerateSchedule(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.GenerateScheduleRequest{}
	resp := &proto.GenerateScheduleResponse{}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,生成排程请求参数无效:%v", transID, err)
		return
	}
	data, err := logic.GenerateSchedule(req, middleware.GetUserID(c))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp = data
	}
	c.JSON(http.StatusOK, resp)
}

// ReleaseSchedule godoc
// @Summary 下发排程
// @Description 回写各工单预计开工/完工时间
// @Tags APS排程
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ReleaseScheduleRequest true "ReleaseScheduleRequest"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/aps/schedule/release [put]
func ReleaseSchedule(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ReleaseScheduleRequest{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,下发排程请求参数无效:%v", transID, err)
		return
	}
	if err := logic.ReleaseSchedule(req); err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// VoidSchedule godoc
// @Summary 作废排程计划
// @Description 作废未下发的排程计划
// @Tags APS排程
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param id query string true "排程计划ID"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/aps/schedule/void [put]
func VoidSchedule(c *gin.Context) {
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		resp.Message = "id不能为空"
		c.JSON(http.StatusOK, resp)
		return
	}
	if err := logic.VoidSchedule(id); err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QuerySchedulePlan godoc
// @Summary 查询排程计划
// @Description 查询排程计划（含明细，可直接渲染甘特图）
// @Tags APS排程
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.QueryProductionSchedulePlanResponse
// @Router /api/mom/aps/schedule/query [get]
func QuerySchedulePlan(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.QueryProductionSchedulePlanRequest{}
	resp := &proto.QueryProductionSchedulePlanResponse{}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,查询排程计划请求参数无效:%v", transID, err)
		return
	}
	logic.QueryProductionSchedulePlan(req, resp, true)
	c.JSON(http.StatusOK, resp)
}

// GetSchedulePlanDetail godoc
// @Summary 排程计划详情
// @Description 获取排程计划详情（含明细时段，甘特图数据源）
// @Tags APS排程
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param id query string true "排程计划ID"
// @Success 200 {object} proto.GetProductionSchedulePlanDetailResponse
// @Router /api/mom/aps/schedule/detail [get]
func GetSchedulePlanDetail(c *gin.Context) {
	resp := &proto.GetProductionSchedulePlanDetailResponse{}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		resp.Message = "id不能为空"
		c.JSON(http.StatusOK, resp)
		return
	}
	m, err := logic.GetProductionSchedulePlanByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionSchedulePlanToPB(m)
		resp.Code = proto.Code_Success
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteSchedulePlan godoc
// @Summary 删除排程计划
// @Description 删除未下发的排程计划
// @Tags APS排程
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.GetByIDsRequest true "Delete SchedulePlan"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/aps/schedule/delete [delete]
func DeleteSchedulePlan(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.GetByIDsRequest{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,删除排程计划请求参数无效:%v", transID, err)
		return
	}
	if len(req.Ids) == 0 || req.Ids[0] == "" {
		resp.Code = proto.Code_BadRequest
		resp.Message = "id不能为空"
		c.JSON(http.StatusOK, resp)
		return
	}
	if err := logic.DeleteProductionSchedulePlan(req.Ids[0]); err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterRouter(r *gin.Engine) {
	g := r.Group("/api/mom/aps/schedule")

	g.POST("generate", GenerateSchedule)
	g.PUT("release", ReleaseSchedule)
	g.PUT("void", VoidSchedule)
	g.GET("query", QuerySchedulePlan)
	g.GET("detail", GetSchedulePlanDetail)
	g.DELETE("delete", DeleteSchedulePlan)
}
