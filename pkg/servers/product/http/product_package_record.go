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
	"github.com/gin-gonic/gin"
)

// AddProductPackageRecord godoc
// @Summary 新增
// @Description 新增
// @Tags 产品包装记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductPackageRecordInfo true "Add ProductPackageRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productpackagerecord/add [post]
func AddProductPackageRecord(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductPackageRecordInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建产品包装记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductPackageRecord(model.PBToProductPackageRecord(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductPackageRecord godoc
// @Summary 更新
// @Description 更新
// @Tags 产品包装记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductPackageRecordInfo true "Update ProductPackageRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productpackagerecord/update [put]
func UpdateProductPackageRecord(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductPackageRecordInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新产品包装记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductPackageRecord(model.PBToProductPackageRecord(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductPackageRecord godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 产品包装记录管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param packageNo query string false "包装箱号"
// @Param productionLineID query string false "生产产线ID"
// @Param createTime0 query string false "装箱时间开始"
// @Param createTime1 query string false "装箱时间结束"
// @Param productSerialNo query string false "产品信息"
// @Success 200 {object} proto.QueryProductPackageRecordResponse
// @Router /api/mom/product/productpackagerecord/query [get]
func QueryProductPackageRecord(c *gin.Context) {
	req := &proto.QueryProductPackageRecordRequest{}
	resp := &proto.QueryProductPackageRecordResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductPackageRecord(req, resp, false)
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

// GetAllProductPackageRecord godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 产品包装记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductPackageRecordResponse
// @Router /api/mom/product/productpackagerecord/all [get]
func GetAllProductPackageRecord(c *gin.Context) {
	resp := &proto.GetAllProductPackageRecordResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductPackageRecords()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductPackageRecordsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductPackageRecordDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 产品包装记录管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductPackageRecordDetailResponse
// @Router /api/mom/product/productpackagerecord/detail [get]
func GetProductPackageRecordDetail(c *gin.Context) {
	resp := &proto.GetProductPackageRecordDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductPackageRecordByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductPackageRecordToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductPackageRecord godoc
// @Summary 删除
// @Description 删除
// @Tags 产品包装记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductPackageRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productpackagerecord/delete [delete]
func DeleteProductPackageRecord(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除产品包装记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductPackageRecord(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductPackageRecordRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productpackagerecord")

	g.POST("add", AddProductPackageRecord)
	g.PUT("update", UpdateProductPackageRecord)
	g.GET("query", QueryProductPackageRecord)
	g.DELETE("delete", DeleteProductPackageRecord)
	g.GET("all", GetAllProductPackageRecord)
	g.GET("detail", GetProductPackageRecordDetail)
}
