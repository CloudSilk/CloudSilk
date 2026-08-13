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

// AddProductOrderBom godoc
// @Summary 新增
// @Description 新增
// @Tags 工单BOM管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductOrderBomInfo true "Add ProductOrderBom"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderbom/add [post]
func AddProductOrderBom(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductOrderBomInfo{
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
		log.Warnf(context.Background(), "TransID:%s,新建工单BOM管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductOrderBom(model.PBToProductOrderBom(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductOrderBom godoc
// @Summary 更新
// @Description 更新
// @Tags 工单BOM管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductOrderBomInfo true "Update ProductOrderBom"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderbom/update [put]
func UpdateProductOrderBom(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductOrderBomInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新工单BOM管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductOrderBom(model.PBToProductOrderBom(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductOrderBom godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 工单BOM管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productOrderNo query string false "生产工单号"
// @Param materialNo query string false "物料信息"
// @Success 200 {object} proto.QueryProductOrderBomResponse
// @Router /api/mom/product/productorderbom/query [get]
func QueryProductOrderBom(c *gin.Context) {
	req := &proto.QueryProductOrderBomRequest{}
	resp := &proto.QueryProductOrderBomResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductOrderBom(req, resp, false)
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

// GetAllProductOrderBom godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 工单BOM管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductOrderBomResponse
// @Router /api/mom/product/productorderbom/all [get]
func GetAllProductOrderBom(c *gin.Context) {
	resp := &proto.GetAllProductOrderBomResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductOrderBoms()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductOrderBomsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductOrderBomDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 工单BOM管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductOrderBomDetailResponse
// @Router /api/mom/product/productorderbom/detail [get]
func GetProductOrderBomDetail(c *gin.Context) {
	resp := &proto.GetProductOrderBomDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductOrderBomByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductOrderBomToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductOrderBom godoc
// @Summary 删除
// @Description 删除
// @Tags 工单BOM管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductOrderBom"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderbom/delete [delete]
func DeleteProductOrderBom(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除工单BOM管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductOrderBom(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductOrderBomRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productorderbom")

	g.POST("add", AddProductOrderBom)
	g.PUT("update", UpdateProductOrderBom)
	g.GET("query", QueryProductOrderBom)
	g.DELETE("delete", DeleteProductOrderBom)
	g.GET("all", GetAllProductOrderBom)
	g.GET("detail", GetProductOrderBomDetail)
}
