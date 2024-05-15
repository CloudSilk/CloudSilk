package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	RegisterProductionFactoryRouter(r)
	RegisterProductionLineRouter(r)
	RegisterProductionStationRouter(r)
	RegisterProductionCrosswayRouter(r)
}
