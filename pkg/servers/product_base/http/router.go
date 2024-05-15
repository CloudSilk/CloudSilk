package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	RegisterProductAttributeRouter(r)
	RegisterProductBrandRouter(r)
	RegisterProductCategoryAttributeRouter(r)
	RegisterProductCategoryRouter(r)
	RegisterProductModelRouter(r)
	RegisterProductModelBomRouter(r)
}
