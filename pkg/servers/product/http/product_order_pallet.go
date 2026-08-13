package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/product/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddProductOrderPallet godoc
// @Summary 新增
// @Description 新增
// @Tags 工单栈板管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductOrderPalletInfo true "Add ProductOrderPallet"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderpallet/add [post]
func AddProductOrderPallet(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductOrderPalletInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建工单栈板管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductOrderPallet(model.PBToProductOrderPallet(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductOrderPallet godoc
// @Summary 更新
// @Description 更新
// @Tags 工单栈板管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductOrderPalletInfo true "Update ProductOrderPallet"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderpallet/update [put]
func UpdateProductOrderPallet(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductOrderPalletInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新工单栈板管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductOrderPallet(model.PBToProductOrderPallet(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductOrderPallet godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 工单栈板管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productOrderNo query string false "生产工单号"
// @Param currentState query string false "当前状态"
// @Param createTime0 query string false "创建时间开始"
// @Param createTime1 query string false "创建时间结束"
// @Success 200 {object} proto.QueryProductOrderPalletResponse
// @Router /api/mom/product/productorderpallet/query [get]
func QueryProductOrderPallet(c *gin.Context) {
	req := &proto.QueryProductOrderPalletRequest{}
	resp := &proto.QueryProductOrderPalletResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductOrderPallet(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductOrderPallet godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 工单栈板管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductOrderPalletResponse
// @Router /api/mom/product/productorderpallet/all [get]
func GetAllProductOrderPallet(c *gin.Context) {
	resp := &proto.GetAllProductOrderPalletResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductOrderPallets()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductOrderPalletsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductOrderPalletDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 工单栈板管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductOrderPalletDetailResponse
// @Router /api/mom/product/productorderpallet/detail [get]
func GetProductOrderPalletDetail(c *gin.Context) {
	resp := &proto.GetProductOrderPalletDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductOrderPalletByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductOrderPalletToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductOrderPallet godoc
// @Summary 删除
// @Description 删除
// @Tags 工单栈板管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductOrderPallet"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderpallet/delete [delete]
func DeleteProductOrderPallet(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.DelRequest{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,删除工单栈板管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductOrderPallet(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductOrderPalletRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productorderpallet")

	g.POST("add", AddProductOrderPallet)
	g.PUT("update", UpdateProductOrderPallet)
	g.GET("query", QueryProductOrderPallet)
	g.DELETE("delete", DeleteProductOrderPallet)
	g.GET("all", GetAllProductOrderPallet)
	g.GET("detail", GetProductOrderPalletDetail)
}
