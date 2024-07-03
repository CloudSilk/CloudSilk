package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/clients"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/product/logic"
	"github.com/CloudSilk/pkg/utils/log"
	usercenter "github.com/CloudSilk/usercenter/proto"
	"github.com/CloudSilk/usercenter/utils/middleware"
	ucmiddleware "github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddProductOrder godoc
// @Summary 新增
// @Description 新增
// @Tags 生产工单管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductOrderInfo true "Add ProductOrder"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorder/add [post]
func AddProductOrder(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductOrderInfo{
		CreateUserID: ucmiddleware.GetUserID(c),
	}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建生产工单请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductOrder(model.PBToProductOrder(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductOrder godoc
// @Summary 更新
// @Description 更新
// @Tags 生产工单管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductOrderInfo true "Update ProductOrder"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorder/update [put]
func UpdateProductOrder(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductOrderInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新生产工单请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductOrder(model.PBToProductOrder(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductOrder godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 生产工单管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productOrderNo query string false "生产工单号"
// @Param salesOrderNo query string false "销售单号"
// @Param itemNo query string false "销售项号"
// @Param orderTime query string false "下单时间"
// @Param createTime0 query string false "创建时间开始"
// @Param createTime1 query string false "创建时间结束"
// @Success 200 {object} proto.QueryProductOrderResponse
// @Router /api/mom/product/productorder/query [get]
func QueryProductOrder(c *gin.Context) {
	req := &proto.QueryProductOrderRequest{}
	resp := &proto.QueryProductOrderResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductOrder(req, resp, false)
	if resp.Code == proto.Code_Success {
		r := &usercenter.QueryUserRequest{}
		for _, u := range resp.Data {
			r.Ids = append(r.Ids, u.CreateUserID)
		}
		r.PageSize = int64(len(r.Ids))
		users, err := clients.UserClient.Query(context.Background(), r)
		if err == nil && users.Code == usercenter.Code_Success {
			for _, u := range resp.Data {
				for _, u2 := range users.Data {
					if u.CreateUserID == u2.Id {
						u.CreateUserName = u2.Nickname
					}
				}
			}
		}
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllProductOrder godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 生产工单管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductOrderResponse
// @Router /api/mom/product/productorder/all [get]
func GetAllProductOrder(c *gin.Context) {
	resp := &proto.GetAllProductOrderResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductOrders()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductOrdersToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductOrderDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 生产工单管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductOrderDetailResponse
// @Router /api/mom/product/productorder/detail [get]
func GetProductOrderDetail(c *gin.Context) {
	resp := &proto.GetProductOrderDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductOrderByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductOrderToPB(data)
		user, err := clients.UserClient.GetDetail(context.Background(), &usercenter.GetDetailRequest{Id: resp.Data.CreateUserID})
		if err != nil {
			resp.Code = proto.Code_InternalServerError
			resp.Message = err.Error()
			c.JSON(http.StatusOK, resp)
			return
		}
		resp.Data.CreateUserName = user.Data.Nickname
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductOrder godoc
// @Summary 删除
// @Description 删除
// @Tags 生产工单管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductOrder"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorder/delete [delete]
func DeleteProductOrder(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除生产工单请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductOrder(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// CancelProductOrder godoc
// @Summary 取消
// @Description 取消
// @Tags 生产工单管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.GetByIDsRequest true "Cancel ProductOrder"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorder/cancel [put]
func CancelProductOrder(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.GetByIDsRequest{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,取消生产工单请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.CancelProductOrder(req.Ids)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// SuspendProductOrder godoc
// @Summary 暂缓
// @Description 暂缓
// @Tags 生产工单管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.GetByIDsRequest true "Suspend ProductOrder"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorder/suspend [put]
func SuspendProductOrder(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.GetByIDsRequest{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,暂缓生产工单请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.SuspendProductOrder(req.Ids)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// ResumeProductOrder godoc
// @Summary 恢复
// @Description 恢复
// @Tags 生产工单管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.GetByIDsRequest true "Resume ProductOrder"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorder/resume [put]
func ResumeProductOrder(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.GetByIDsRequest{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,恢复生产工单请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.ResumeProductOrder(req.Ids)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductOrderRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productorder")

	g.POST("add", AddProductOrder)
	g.PUT("update", UpdateProductOrder)
	g.GET("query", QueryProductOrder)
	g.DELETE("delete", DeleteProductOrder)
	g.GET("all", GetAllProductOrder)
	g.GET("detail", GetProductOrderDetail)
	g.PUT("cancel", CancelProductOrder)
	g.PUT("suspend", SuspendProductOrder)
	g.PUT("resume", ResumeProductOrder)
}
