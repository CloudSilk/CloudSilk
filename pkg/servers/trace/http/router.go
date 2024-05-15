package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	RegisterOperationTraceRouter(r)
	RegisterInvocationTraceRouter(r)
	RegisterExceptionTraceRouter(r)
}
