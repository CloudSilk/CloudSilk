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

// AddProductOrderAttribute godoc
// @Summary 新增
// @Description 新增
// @Tags 工单特性管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductOrderAttributeInfo true "Add ProductOrderAttribute"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderattribute/add [post]
func AddProductOrderAttribute(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductOrderAttributeInfo{
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
		log.Warnf(context.Background(), "TransID:%s,新建工单特性请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductOrderAttribute(model.PBToProductOrderAttribute(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductOrderAttribute godoc
// @Summary 更新
// @Description 更新
// @Tags 工单特性管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductOrderAttributeInfo true "Update ProductOrderAttribute"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderattribute/update [put]
func UpdateProductOrderAttribute(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductOrderAttributeInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新工单特性请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductOrderAttribute(model.PBToProductOrderAttribute(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductOrderAttribute godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 工单特性管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productOrderNo query string false "工单号"
// @Param code query string false "代号或描述"
// @Success 200 {object} proto.QueryProductOrderAttributeResponse
// @Router /api/mom/product/productorderattribute/query [get]
func QueryProductOrderAttribute(c *gin.Context) {
	req := &proto.QueryProductOrderAttributeRequest{}
	resp := &proto.QueryProductOrderAttributeResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductOrderAttribute(req, resp, false)
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

// GetAllProductOrderAttribute godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 工单特性管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductOrderAttributeResponse
// @Router /api/mom/product/productorderattribute/all [get]
func GetAllProductOrderAttribute(c *gin.Context) {
	resp := &proto.GetAllProductOrderAttributeResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductOrderAttributes()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductOrderAttributesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductOrderAttributeDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 工单特性管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductOrderAttributeDetailResponse
// @Router /api/mom/product/productorderattribute/detail [get]
func GetProductOrderAttributeDetail(c *gin.Context) {
	resp := &proto.GetProductOrderAttributeDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductOrderAttributeByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductOrderAttributeToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductOrderAttribute godoc
// @Summary 删除
// @Description 删除
// @Tags 工单特性管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductOrderAttribute"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderattribute/delete [delete]
func DeleteProductOrderAttribute(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除工单特性请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductOrderAttribute(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductOrderAttributeRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productorderattribute")

	g.POST("add", AddProductOrderAttribute)
	g.PUT("update", UpdateProductOrderAttribute)
	g.GET("query", QueryProductOrderAttribute)
	g.DELETE("delete", DeleteProductOrderAttribute)
	g.GET("all", GetAllProductOrderAttribute)
	g.GET("detail", GetProductOrderAttributeDetail)
}
