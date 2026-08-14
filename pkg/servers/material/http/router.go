package http

import (
	mw "github.com/CloudSilk/CloudSilk/pkg/middleware"
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

	// WMS作业流程：库存变更类操作需授权角色，查询类开放
	w := r.Group("/api/mom/material/wms")
	w.POST("receive", ReceiveMaterial, mw.RequirePermissionParam("wms.operate"))
	w.POST("pickbill", CreatePickBill, mw.RequirePermissionParam("wms.operate"))
	w.PUT("pickbill/complete", CompletePickBill, mw.RequirePermissionParam("wms.operate"))
	w.PUT("pickbill/cancel", CancelPickBill, mw.RequirePermissionParam("wms.operate"))
	w.PUT("stocktake", StocktakeInventory, mw.RequirePermissionParam("wms.operate"))
	w.GET("transaction/query", QueryInventoryTransaction)
	w.GET("alert", GetInventoryAlerts)
}
