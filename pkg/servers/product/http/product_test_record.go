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

// AddProductTestRecord godoc
// @Summary 新增
// @Description 新增
// @Tags 产品测试记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductTestRecordInfo true "Add ProductTestRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/producttestrecord/add [post]
func AddProductTestRecord(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductTestRecordInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建产品测试记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductTestRecord(model.PBToProductTestRecord(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductTestRecord godoc
// @Summary 更新
// @Description 更新
// @Tags 产品测试记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductTestRecordInfo true "Update ProductTestRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/producttestrecord/update [put]
func UpdateProductTestRecord(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductTestRecordInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新产品测试记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductTestRecord(model.PBToProductTestRecord(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductTestRecord godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 产品测试记录管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productOrderNo query string false "生产工单号"
// @Param productSerialNo query string false "产品序列号"
// @Param testStartTime0 query string false "开始测试时间开始"
// @Param testStartTime1 query string false "开始测试时间结束"
// @Param productionLineID query string false "作业产线"
// @Success 200 {object} proto.QueryProductTestRecordResponse
// @Router /api/mom/product/producttestrecord/query [get]
func QueryProductTestRecord(c *gin.Context) {
	req := &proto.QueryProductTestRecordRequest{}
	resp := &proto.QueryProductTestRecordResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductTestRecord(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductTestRecord godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 产品测试记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductTestRecordResponse
// @Router /api/mom/product/producttestrecord/all [get]
func GetAllProductTestRecord(c *gin.Context) {
	resp := &proto.GetAllProductTestRecordResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductTestRecords()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductTestRecordsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductTestRecordDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 产品测试记录管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductTestRecordDetailResponse
// @Router /api/mom/product/producttestrecord/detail [get]
func GetProductTestRecordDetail(c *gin.Context) {
	resp := &proto.GetProductTestRecordDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductTestRecordByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductTestRecordToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductTestRecord godoc
// @Summary 删除
// @Description 删除
// @Tags 产品测试记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductTestRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/producttestrecord/delete [delete]
func DeleteProductTestRecord(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除产品测试记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductTestRecord(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductTestRecordRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/producttestrecord")

	g.POST("add", AddProductTestRecord)
	g.PUT("update", UpdateProductTestRecord)
	g.GET("query", QueryProductTestRecord)
	g.DELETE("delete", DeleteProductTestRecord)
	g.GET("all", GetAllProductTestRecord)
	g.GET("detail", GetProductTestRecordDetail)
}
