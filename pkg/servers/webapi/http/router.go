package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	RegisterInfrastructureRouter(r)
	RegisterProductionRouter(r)
	// RegisterMaterialRouter(r)
}
