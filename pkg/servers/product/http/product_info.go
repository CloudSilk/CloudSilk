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

// AddProductInfo godoc
// @Summary 新增
// @Description 新增
// @Tags 产品信息管理管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductInfoInfo true "Add ProductInfo"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productinfo/add [post]
func AddProductInfo(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductInfoInfo{
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
		log.Warnf(context.Background(), "TransID:%s,新建产品信息管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductInfo(model.PBToProductInfo(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductInfo godoc
// @Summary 更新
// @Description 更新
// @Tags 产品信息管理管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductInfoInfo true "Update ProductInfo"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productinfo/update [put]
func UpdateProductInfo(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductInfoInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新产品信息管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductInfo(model.PBToProductInfo(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductInfo godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 产品信息管理管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productSerialNo query string false "产品序列号"
// @Param createTime0 query string false "创建时间开始"
// @Param createTime1 query string false "创建时间结束"
// @Param currentState query string false "当前状态"
// @Param productOrderNo query string false "生产工单号"
// @Success 200 {object} proto.QueryProductInfoResponse
// @Router /api/mom/product/productinfo/query [get]
func QueryProductInfo(c *gin.Context) {
	req := &proto.QueryProductInfoRequest{}
	resp := &proto.QueryProductInfoResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductInfo(req, resp, false)
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

// GetAllProductInfo godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 产品信息管理管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductInfoResponse
// @Router /api/mom/product/productinfo/all [get]
func GetAllProductInfo(c *gin.Context) {
	resp := &proto.GetAllProductInfoResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductInfos()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductInfosToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductInfoDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 产品信息管理管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductInfoDetailResponse
// @Router /api/mom/product/productinfo/detail [get]
func GetProductInfoDetail(c *gin.Context) {
	resp := &proto.GetProductInfoDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductInfoByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductInfoToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductInfo godoc
// @Summary 删除
// @Description 删除
// @Tags 产品信息管理管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductInfo"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productinfo/delete [delete]
func DeleteProductInfo(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除产品信息管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductInfo(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductInfoRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productinfo")

	g.POST("add", AddProductInfo)
	g.PUT("update", UpdateProductInfo)
	g.GET("query", QueryProductInfo)
	g.DELETE("delete", DeleteProductInfo)
	g.GET("all", GetAllProductInfo)
	g.GET("detail", GetProductInfoDetail)
}
