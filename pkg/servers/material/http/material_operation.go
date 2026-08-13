package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/material/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// ReceiveMaterial godoc
// @Summary 收货入库
// @Description 收货入库：增加账面库存并写入事务流水
// @Tags WMS作业
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ReceiveMaterialRequest true "ReceiveMaterialRequest"
// @Success 200 {object} proto.ReceiveMaterialResponse
// @Router /api/mom/material/wms/receive [post]
func ReceiveMaterial(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ReceiveMaterialRequest{}
	resp := &proto.ReceiveMaterialResponse{}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,收货入库请求参数无效:%v", transID, err)
		return
	}
	data, err := logic.ReceiveMaterial(req, middleware.GetUserID(c))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp = data
	}
	c.JSON(http.StatusOK, resp)
}

// CreatePickBill godoc
// @Summary 按工单生成拣货单
// @Description 按工单BOM逐行锁定库存生成拣货单（缺料行跳过并提示）
// @Tags WMS作业
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.CreatePickBillRequest true "CreatePickBillRequest"
// @Success 200 {object} proto.CreatePickBillResponse
// @Router /api/mom/material/wms/pickbill [post]
func CreatePickBill(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.CreatePickBillRequest{}
	resp := &proto.CreatePickBillResponse{}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,生成拣货单请求参数无效:%v", transID, err)
		return
	}
	data, err := logic.CreatePickBillFromOrder(req, middleware.GetUserID(c))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp = data
	}
	c.JSON(http.StatusOK, resp)
}

// CompletePickBill godoc
// @Summary 完成拣货
// @Description 扣减账面库存、解除锁定并联动工单发料数量
// @Tags WMS作业
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.CompletePickBillRequest true "CompletePickBillRequest"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/wms/pickbill/complete [put]
func CompletePickBill(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.CompletePickBillRequest{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,完成拣货请求参数无效:%v", transID, err)
		return
	}
	if err := logic.CompletePickBill(req, middleware.GetUserID(c)); err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// CancelPickBill godoc
// @Summary 取消拣货单
// @Description 解除库存锁定
// @Tags WMS作业
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.CancelPickBillRequest true "CancelPickBillRequest"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/wms/pickbill/cancel [put]
func CancelPickBill(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.CancelPickBillRequest{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,取消拣货单请求参数无效:%v", transID, err)
		return
	}
	if err := logic.CancelPickBill(req, middleware.GetUserID(c)); err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// StocktakeInventory godoc
// @Summary 库存盘点
// @Description 按实盘数量调整账面库存并记录差异流水
// @Tags WMS作业
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.StocktakeInventoryRequest true "StocktakeInventoryRequest"
// @Success 200 {object} proto.StocktakeInventoryResponse
// @Router /api/mom/material/wms/stocktake [put]
func StocktakeInventory(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.StocktakeInventoryRequest{}
	resp := &proto.StocktakeInventoryResponse{}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,库存盘点请求参数无效:%v", transID, err)
		return
	}
	data, err := logic.StocktakeInventory(req, middleware.GetUserID(c))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp = data
	}
	c.JSON(http.StatusOK, resp)
}

// QueryInventoryTransaction godoc
// @Summary 查询库存事务流水
// @Description 查询库存事务流水（入库/出库/锁定/解锁/盘点调整）
// @Tags WMS作业
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.QueryMaterialInventoryTransactionResponse
// @Router /api/mom/material/wms/transaction/query [get]
func QueryInventoryTransaction(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.QueryMaterialInventoryTransactionRequest{}
	resp := &proto.QueryMaterialInventoryTransactionResponse{}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,查询库存流水请求参数无效:%v", transID, err)
		return
	}
	logic.QueryMaterialInventoryTransaction(req, resp, true)
	c.JSON(http.StatusOK, resp)
}

// GetInventoryAlerts godoc
// @Summary 库存预警
// @Description 可用量低于补料规则最低库存量或已为负的记录
// @Tags WMS作业
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.MaterialInventoryAlertResponse
// @Router /api/mom/material/wms/alert [get]
func GetInventoryAlerts(c *gin.Context) {
	resp, err := logic.GetInventoryAlerts()
	if err != nil {
		resp = &proto.MaterialInventoryAlertResponse{
			Code:    proto.Code_InternalServerError,
			Message: err.Error(),
		}
	}
	c.JSON(http.StatusOK, resp)
}
