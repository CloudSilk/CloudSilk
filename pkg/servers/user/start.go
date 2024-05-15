package user

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/CloudSilk/CloudSilk/pkg/servers/user/provider"
	"github.com/CloudSilk/CloudSilk/pkg/servers/user/http"
	"github.com/gin-gonic/gin"
)

type Server struct{}

func (s *Server) Start(r *gin.Engine) {
	http.RegisterRouter(r)

	config.SetProviderService(&provider.PersonnelQualificationProvider{})
}
