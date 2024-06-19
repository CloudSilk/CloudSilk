package production

import (
	"github.com/CloudSilk/CloudSilk/pkg/servers/production/http"
	"github.com/gin-gonic/gin"
)

type Server struct{}

func (s *Server) Start(r *gin.Engine) {
	http.RegisterRouter(r)
}
