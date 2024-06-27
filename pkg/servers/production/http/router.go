package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	RegisterProcessStepParameterRouter(r)
	RegisterProcessStepTypeRouter(r)
	RegisterProductionProcessSopRouter(r)
	RegisterProductionProcessStepRouter(r)
	RegisterProductionProcessRouter(r)
	RegisterProductionRhythmRouter(r)
	RegisterProductionStationAlarmRouter(r)
	RegisterProductionStationBreakdownRouter(r)
	RegisterProductionStationOutputRouter(r)
	RegisterProductionStationSignupRouter(r)
	RegisterProductionStationStartupRouter(r)
	RegisterPersonnelQualificationTypeRouter(r)
	RegisterPersonnelQualificationRouter(r)
}
