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

// AddProductReleaseRecord godoc
// @Summary 新增
// @Description 新增
// @Tags 投料记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductReleaseRecordInfo true "Add ProductReleaseRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productreleaserecord/add [post]
func AddProductReleaseRecord(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductReleaseRecordInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建投料记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductReleaseRecord(model.PBToProductReleaseRecord(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductReleaseRecord godoc
// @Summary 更新
// @Description 更新
// @Tags 投料记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductReleaseRecordInfo true "Update ProductReleaseRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productreleaserecord/update [put]
func UpdateProductReleaseRecord(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductReleaseRecordInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新投料记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductReleaseRecord(model.PBToProductReleaseRecord(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductReleaseRecord godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 投料记录管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param ProductSerialNo query string false "产品序列号"
// @Param ProductOrderNo query string false "生产工单号"
// @Param SecurityCode query string false "投料信息"
// @Param CreateTime0 query string false "发料时间"
// @Param CreateTime0 query string false "投料信息"
// @Success 200 {object} proto.QueryProductReleaseRecordResponse
// @Router /api/mom/product/productreleaserecord/query [get]
func QueryProductReleaseRecord(c *gin.Context) {
	req := &proto.QueryProductReleaseRecordRequest{}
	resp := &proto.QueryProductReleaseRecordResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductReleaseRecord(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductReleaseRecord godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 投料记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductReleaseRecordResponse
// @Router /api/mom/product/productreleaserecord/all [get]
func GetAllProductReleaseRecord(c *gin.Context) {
	resp := &proto.GetAllProductReleaseRecordResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductReleaseRecords()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductReleaseRecordsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductReleaseRecordDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 投料记录管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductReleaseRecordDetailResponse
// @Router /api/mom/product/productreleaserecord/detail [get]
func GetProductReleaseRecordDetail(c *gin.Context) {
	resp := &proto.GetProductReleaseRecordDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductReleaseRecordByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductReleaseRecordToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductReleaseRecord godoc
// @Summary 删除
// @Description 删除
// @Tags 投料记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductReleaseRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productreleaserecord/delete [delete]
func DeleteProductReleaseRecord(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除投料记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductReleaseRecord(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductReleaseRecordRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productreleaserecord")

	g.POST("add", AddProductReleaseRecord)
	g.PUT("update", UpdateProductReleaseRecord)
	g.GET("query", QueryProductReleaseRecord)
	g.DELETE("delete", DeleteProductReleaseRecord)
	g.GET("all", GetAllProductReleaseRecord)
	g.GET("detail", GetProductReleaseRecordDetail)
}
