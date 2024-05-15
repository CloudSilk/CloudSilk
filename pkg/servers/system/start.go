package system

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/CloudSilk/CloudSilk/pkg/servers/system/provider"
	"github.com/CloudSilk/CloudSilk/pkg/servers/system/http"
	"github.com/gin-gonic/gin"
)

type Server struct{}

func (s *Server) Start(r *gin.Engine) {
	http.RegisterRouter(r)

	config.SetProviderService(&provider.SystemEventProvider{})
	config.SetProviderService(&provider.SystemEventTriggerProvider{})
	config.SetProviderService(&provider.SystemEventTriggerParameterProvider{})
}
