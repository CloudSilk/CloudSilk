package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	RegisterMaterialCategoryRouter(r)
	RegisterMaterialInfoRouter(r)
	RegisterMaterialSupplierRouter(r)
	RegisterMaterialTrayRouter(r)
	RegisterMaterialTrayBindingRecordRouter(r)
}
