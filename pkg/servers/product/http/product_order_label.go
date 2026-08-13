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

// AddProductOrderLabel godoc
// @Summary 新增
// @Description 新增
// @Tags 工单标签管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductOrderLabelInfo true "Add ProductOrderLabel"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderlabel/add [post]
func AddProductOrderLabel(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductOrderLabelInfo{
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
		log.Warnf(context.Background(), "TransID:%s,新建工单标签管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductOrderLabel(model.PBToProductOrderLabel(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductOrderLabel godoc
// @Summary 更新
// @Description 更新
// @Tags 工单标签管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductOrderLabelInfo true "Update ProductOrderLabel"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderlabel/update [put]
func UpdateProductOrderLabel(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductOrderLabelInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新工单标签管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductOrderLabel(model.PBToProductOrderLabel(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductOrderLabel godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 工单标签管理
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
// @Success 200 {object} proto.QueryProductOrderLabelResponse
// @Router /api/mom/product/productorderlabel/query [get]
func QueryProductOrderLabel(c *gin.Context) {
	req := &proto.QueryProductOrderLabelRequest{}
	resp := &proto.QueryProductOrderLabelResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductOrderLabel(req, resp, false)
	if resp.Code == proto.Code_Success {
		r := &usercenter.QueryUserRequest{}
		for _, u := range resp.Data {
			r.Ids = append(r.Ids, u.CreateUserID, u.CheckUserID)
		}

		r.PageSize = int64(len(r.Ids))
		users, err := clients.UserClient.Query(context.Background(), r)
		if err == nil && users.Code == usercenter.Code_Success {
			for _, u := range resp.Data {
				for _, u2 := range users.Data {
					if u.CreateUserID == u2.Id {
						u.CreateUserName = u2.Nickname
					}
					if u.CheckUserID == u2.Id {
						u.CheckUserName = u2.Nickname
					}
				}
			}
		}
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllProductOrderLabel godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 工单标签管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductOrderLabelResponse
// @Router /api/mom/product/productorderlabel/all [get]
func GetAllProductOrderLabel(c *gin.Context) {
	resp := &proto.GetAllProductOrderLabelResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductOrderLabels()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductOrderLabelsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductOrderLabelDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 工单标签管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductOrderLabelDetailResponse
// @Router /api/mom/product/productorderlabel/detail [get]
func GetProductOrderLabelDetail(c *gin.Context) {
	resp := &proto.GetProductOrderLabelDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductOrderLabelByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductOrderLabelToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductOrderLabel godoc
// @Summary 删除
// @Description 删除
// @Tags 工单标签管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductOrderLabel"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderlabel/delete [delete]
func DeleteProductOrderLabel(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除工单标签管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductOrderLabel(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductOrderLabelRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productorderlabel")

	g.POST("add", AddProductOrderLabel)
	g.PUT("update", UpdateProductOrderLabel)
	g.GET("query", QueryProductOrderLabel)
	g.DELETE("delete", DeleteProductOrderLabel)
	g.GET("all", GetAllProductOrderLabel)
	g.GET("detail", GetProductOrderLabelDetail)
}
