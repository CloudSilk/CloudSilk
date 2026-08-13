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

// AddProductOrderProcess godoc
// @Summary 新增
// @Description 新增
// @Tags 工单工序管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductOrderProcessInfo true "Add ProductOrderProcess"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderprocess/add [post]
func AddProductOrderProcess(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductOrderProcessInfo{
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
		log.Warnf(context.Background(), "TransID:%s,新建工单工序请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductOrderProcess(model.PBToProductOrderProcess(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductOrderProcess godoc
// @Summary 更新
// @Description 更新
// @Tags 工单工序管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductOrderProcessInfo true "Update ProductOrderProcess"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderprocess/update [put]
func UpdateProductOrderProcess(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductOrderProcessInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新工单工序请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductOrderProcess(model.PBToProductOrderProcess(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductOrderProcess godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 工单工序管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productOrderNo query string false "工单号"
// @Param createTime0 query string false "创建时间开始"
// @Param createTime1 query string false "创建时间结束"
// @Success 200 {object} proto.QueryProductOrderProcessResponse
// @Router /api/mom/product/productorderprocess/query [get]
func QueryProductOrderProcess(c *gin.Context) {
	req := &proto.QueryProductOrderProcessRequest{}
	resp := &proto.QueryProductOrderProcessResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductOrderProcess(req, resp, false)
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

// GetAllProductOrderProcess godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 工单工序管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductOrderProcessResponse
// @Router /api/mom/product/productorderprocess/all [get]
func GetAllProductOrderProcess(c *gin.Context) {
	resp := &proto.GetAllProductOrderProcessResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductOrderProcesss()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductOrderProcesssToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductOrderProcessDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 工单工序管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductOrderProcessDetailResponse
// @Router /api/mom/product/productorderprocess/detail [get]
func GetProductOrderProcessDetail(c *gin.Context) {
	resp := &proto.GetProductOrderProcessDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductOrderProcessByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductOrderProcessToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductOrderProcess godoc
// @Summary 删除
// @Description 删除
// @Tags 工单工序管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductOrderProcess"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderprocess/delete [delete]
func DeleteProductOrderProcess(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除工单工序请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductOrderProcess(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductOrderProcessRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productorderprocess")

	g.POST("add", AddProductOrderProcess)
	g.PUT("update", UpdateProductOrderProcess)
	g.GET("query", QueryProductOrderProcess)
	g.DELETE("delete", DeleteProductOrderProcess)
	g.GET("all", GetAllProductOrderProcess)
	g.GET("detail", GetProductOrderProcessDetail)
}
