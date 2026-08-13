package http

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/monitoring/logic"
	"github.com/gin-gonic/gin"
)

// GetOverview godoc
// @Summary 厂级监控总览
// @Description 产线/工站规模、工单进度、告警与工站状态分布
// @Tags 生产监控
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.MonitoringOverviewResponse
// @Router /api/mom/monitoring/overview [get]
func GetOverview(c *gin.Context) {
	resp, err := logic.GetOverview()
	if err != nil {
		resp = &proto.MonitoringOverviewResponse{
			Code:    proto.Code_InternalServerError,
			Message: err.Error(),
		}
	}
	c.JSON(http.StatusOK, resp)
}

// GetLineMonitoring godoc
// @Summary 产线监控
// @Description OEE、周期产量、工单达成率、在制数、未处理告警
// @Tags 生产监控
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param productionLineID query string true "产线ID"
// @Param startTime query string false "开始时间 yyyy-MM-dd HH:mm:ss（默认近24小时）"
// @Param endTime query string false "结束时间 yyyy-MM-dd HH:mm:ss（默认当前）"
// @Success 200 {object} proto.MonitoringLineResponse
// @Router /api/mom/monitoring/line [get]
func GetLineMonitoring(c *gin.Context) {
	resp, err := logic.GetLineMonitoring(c.Query("productionLineID"), c.Query("startTime"), c.Query("endTime"))
	if err != nil {
		resp = &proto.MonitoringLineResponse{
			Code:    proto.Code_BadRequest,
			Message: err.Error(),
		}
	}
	c.JSON(http.StatusOK, resp)
}

// QueryAlarms godoc
// @Summary 告警列表
// @Description 大屏告警滚动列表（按时间倒序）
// @Tags 生产监控
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param currentState query string false "状态（默认全部）"
// @Param limit query int false "数量上限（默认50）"
// @Success 200 {object} proto.QueryMonitoringAlarmResponse
// @Router /api/mom/monitoring/alarms [get]
func QueryAlarms(c *gin.Context) {
	req := &proto.QueryMonitoringAlarmRequest{}
	_ = c.BindQuery(req)
	resp, err := logic.QueryAlarms(req)
	if err != nil {
		resp = &proto.QueryMonitoringAlarmResponse{
			Code:    proto.Code_InternalServerError,
			Message: err.Error(),
		}
	}
	c.JSON(http.StatusOK, resp)
}

// StreamRealtime godoc
// @Summary 实时监控推送（SSE）
// @Description 每5秒推送一次厂级总览与未处理告警数（text/event-stream，EventSource 兼容）
// @Tags 生产监控
// @Produce  text/event-stream
// @Param authorization header string true "jwt token"
// @Router /api/mom/monitoring/stream [get]
func StreamRealtime(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("X-Accel-Buffering", "no")

	flusher, ok := c.Writer.(http.Flusher)
	if !ok {
		c.String(http.StatusNotImplemented, "streaming unsupported")
		return
	}

	clientGone := c.Request.Context().Done()
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	push := func() {
		overview, err := logic.GetOverview()
		if err != nil {
			return
		}
		unhandled := int64(overview.UnhandledAlarmCount)
		payload := fmt.Sprintf(`{"overview":{"producingOrderCount":%d,"totalFinishedQTY":%d,"unhandledAlarmCount":%d,"breakdownStationCount":%d}}`,
			overview.ProducingOrderCount, overview.TotalFinishedQTY, unhandled, overview.BreakdownStationCount)
		if _, err := io.WriteString(c.Writer, "event: overview\ndata: "+payload+"\n\n"); err == nil {
			flusher.Flush()
		}
	}

	//连接建立即推一次
	push()
	for {
		select {
		case <-clientGone:
			return
		case <-ticker.C:
			push()
		}
	}
}

func RegisterRouter(r *gin.Engine) {
	g := r.Group("/api/mom/monitoring")

	g.GET("overview", GetOverview)
	g.GET("line", GetLineMonitoring)
	g.GET("alarms", QueryAlarms)
	g.GET("stream", StreamRealtime)
}
