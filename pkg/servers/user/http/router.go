package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	RegisterPersonnelQualificationTypeRouter(r)
	RegisterPersonnelQualificationRouter(r)
}
