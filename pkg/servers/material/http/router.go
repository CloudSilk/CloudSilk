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
	RegisterMaterialContainerTypeRouter(r)
	RegisterMaterialContainerRouter(r)
	RegisterMaterialInventoryRouter(r)
	RegisterMaterialShelfBinRouter(r)
	RegisterMaterialShelfRouter(r)
	RegisterMaterialStoreRouter(r)
	RegisterMaterialStoreFeedRuleRouter(r)
	RegisterAGVTaskTypeRouter(r)
	RegisterAGVTaskQueueRouter(r)
	RegisterWMSBillQueueRouter(r)
	RegisterMaterialReturnCauseRouter(r)
	RegisterMaterialReturnRequestFormRouter(r)
	RegisterMaterialReturnSolutionRouter(r)
	RegisterMaterialReturnTypeRouter(r)
	RegisterMaterialChannelLayerRouter(r)
}
