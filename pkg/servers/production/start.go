package production

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/CloudSilk/CloudSilk/pkg/servers/production/http"
	"github.com/CloudSilk/CloudSilk/pkg/servers/production/provider"
	"github.com/gin-gonic/gin"
)

type Server struct{}

func (s *Server) Start(r *gin.Engine) {
	http.RegisterRouter(r)

	config.SetProviderService(&provider.ProductionProcessProvider{})
	config.SetProviderService(&provider.ProductionProcessSopProvider{})
	config.SetProviderService(&provider.ProductionStationOutputProvider{})
	config.SetProviderService(&provider.ProductionProcessStepProvider{})
	config.SetProviderService(&provider.ProductionStationSignupProvider{})
	config.SetProviderService(&provider.ProcessStepParameterProvider{})
	config.SetProviderService(&provider.PersonnelQualificationProvider{})
}
