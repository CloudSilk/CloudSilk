package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	RegisterLabelTemplateRouter(r)
	RegisterLabelTypeRouter(r)
}
