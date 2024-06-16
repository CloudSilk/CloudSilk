package product

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/CloudSilk/CloudSilk/pkg/servers/product/http"
	"github.com/CloudSilk/CloudSilk/pkg/servers/product/provider"
	"github.com/gin-gonic/gin"
)

type Server struct{}

func (s *Server) Start(r *gin.Engine) {
	http.RegisterRouter(r)

	config.SetProviderService(&provider.ProductInfoProvider{})
	config.SetProviderService(&provider.ProductOrderAttributeProvider{})
	config.SetProviderService(&provider.ProductOrderBomProvider{})
	config.SetProviderService(&provider.ProductOrderProcessStepProvider{})
	config.SetProviderService(&provider.ProductOrderProcessProvider{})
	config.SetProviderService(&provider.ProductOrderProvider{})
	config.SetProviderService(&provider.ProductPackageRecordProvider{})
	config.SetProviderService(&provider.ProductProcessRouteProvider{})
	config.SetProviderService(&provider.ProductReworkRecordProvider{})
	config.SetProviderService(&provider.ProductRhythmRecordProvider{})
	config.SetProviderService(&provider.ProductTestRecordProvider{})
	config.SetProviderService(&provider.ProductWorkRecordProvider{})
}
