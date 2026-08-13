package scada

import (
	"sync"

	"github.com/CloudSilk/CloudSilk/pkg/servers/scada/http"
	"github.com/CloudSilk/CloudSilk/pkg/servers/scada/logic"
	"github.com/gin-gonic/gin"
)

type Server struct{}

var collectorOnce sync.Once

func (s *Server) Start(r *gin.Engine) {
	http.RegisterRouter(r)

	//后台采集器仅启动一次：轮询启用设备的启用点位
	collectorOnce.Do(func() {
		go logic.StartCollector()
	})
}
