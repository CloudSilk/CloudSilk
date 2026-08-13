package monitoring

import (
	"github.com/CloudSilk/CloudSilk/pkg/servers/monitoring/http"
	"github.com/gin-gonic/gin"
)

type Server struct{}

func (s *Server) Start(r *gin.Engine) {
	http.RegisterRouter(r)
}
