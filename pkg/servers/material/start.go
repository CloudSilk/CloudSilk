package material

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/CloudSilk/CloudSilk/pkg/servers/material/http"
	"github.com/CloudSilk/CloudSilk/pkg/servers/material/provider"
	"github.com/gin-gonic/gin"
)

type Server struct{}

func (s *Server) Start(r *gin.Engine) {
	http.RegisterRouter(r)

	config.SetProviderService(&provider.MaterialTrayProvider{})
	config.SetProviderService(&provider.MaterialTrayBindingRecordProvider{})
	config.SetProviderService(&provider.MaterialChannelLayerProvider{})
}
