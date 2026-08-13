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

// AddProductOrderAttachment godoc
// @Summary 新增
// @Description 新增
// @Tags 工单附件管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductOrderAttachmentInfo true "Add ProductOrderAttachment"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderattachment/add [post]
func AddProductOrderAttachment(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductOrderAttachmentInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建工单附件请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	_, userInfo := ucmiddleware.GetUser(c)
	req.CreateUserID = userInfo.Id

	id, err := logic.CreateProductOrderAttachment(model.PBToProductOrderAttachment(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductOrderAttachment godoc
// @Summary 更新
// @Description 更新
// @Tags 工单附件管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductOrderAttachmentInfo true "Update ProductOrderAttachment"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderattachment/update [put]
func UpdateProductOrderAttachment(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductOrderAttachmentInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新工单附件请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductOrderAttachment(model.PBToProductOrderAttachment(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductOrderAttachment godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 工单附件管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productOrderNo query string false "生产工单号"
// @Param materialNo query string false "物料信息"
// @Success 200 {object} proto.QueryProductOrderAttachmentResponse
// @Router /api/mom/product/productorderattachment/query [get]
func QueryProductOrderAttachment(c *gin.Context) {
	req := &proto.QueryProductOrderAttachmentRequest{}
	resp := &proto.QueryProductOrderAttachmentResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductOrderAttachment(req, resp, false)
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

// GetAllProductOrderAttachment godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 工单附件管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductOrderAttachmentResponse
// @Router /api/mom/product/productorderattachment/all [get]
func GetAllProductOrderAttachment(c *gin.Context) {
	resp := &proto.GetAllProductOrderAttachmentResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductOrderAttachments()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductOrderAttachmentsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductOrderAttachmentDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 工单附件管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductOrderAttachmentDetailResponse
// @Router /api/mom/product/productorderattachment/detail [get]
func GetProductOrderAttachmentDetail(c *gin.Context) {
	resp := &proto.GetProductOrderAttachmentDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductOrderAttachmentByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductOrderAttachmentToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductOrderAttachment godoc
// @Summary 删除
// @Description 删除
// @Tags 工单附件管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductOrderAttachment"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderattachment/delete [delete]
func DeleteProductOrderAttachment(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除工单附件请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductOrderAttachment(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductOrderAttachmentRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productorderattachment")

	g.POST("add", AddProductOrderAttachment)
	g.PUT("update", UpdateProductOrderAttachment)
	g.GET("query", QueryProductOrderAttachment)
	g.DELETE("delete", DeleteProductOrderAttachment)
	g.GET("all", GetAllProductOrderAttachment)
	g.GET("detail", GetProductOrderAttachmentDetail)
}
