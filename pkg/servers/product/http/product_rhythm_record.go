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

// AddProductRhythmRecord godoc
// @Summary 新增
// @Description 新增
// @Tags 产品节拍记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductRhythmRecordInfo true "Add ProductRhythmRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productrhythmrecord/add [post]
func AddProductRhythmRecord(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductRhythmRecordInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建产品节拍记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductRhythmRecord(model.PBToProductRhythmRecord(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductRhythmRecord godoc
// @Summary 更新
// @Description 更新
// @Tags 产品节拍记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductRhythmRecordInfo true "Update ProductRhythmRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productrhythmrecord/update [put]
func UpdateProductRhythmRecord(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductRhythmRecordInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新产品节拍记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductRhythmRecord(model.PBToProductRhythmRecord(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductRhythmRecord godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 产品节拍记录管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productOrderNo query string false "生产工单号"
// @Param productSerialNo query string false "产品序列号"
// @Param createTime0 query string false "创建时间开始"
// @Param createTime1 query string false "创建时间结束"
// @Success 200 {object} proto.QueryProductRhythmRecordResponse
// @Router /api/mom/product/productrhythmrecord/query [get]
func QueryProductRhythmRecord(c *gin.Context) {
	req := &proto.QueryProductRhythmRecordRequest{}
	resp := &proto.QueryProductRhythmRecordResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductRhythmRecord(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductRhythmRecord godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 产品节拍记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductRhythmRecordResponse
// @Router /api/mom/product/productrhythmrecord/all [get]
func GetAllProductRhythmRecord(c *gin.Context) {
	resp := &proto.GetAllProductRhythmRecordResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductRhythmRecords()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductRhythmRecordsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductRhythmRecordDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 产品节拍记录管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductRhythmRecordDetailResponse
// @Router /api/mom/product/productrhythmrecord/detail [get]
func GetProductRhythmRecordDetail(c *gin.Context) {
	resp := &proto.GetProductRhythmRecordDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductRhythmRecordByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductRhythmRecordToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductRhythmRecord godoc
// @Summary 删除
// @Description 删除
// @Tags 产品节拍记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductRhythmRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productrhythmrecord/delete [delete]
func DeleteProductRhythmRecord(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除产品节拍记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductRhythmRecord(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductRhythmRecordRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productrhythmrecord")

	g.POST("add", AddProductRhythmRecord)
	g.PUT("update", UpdateProductRhythmRecord)
	g.GET("query", QueryProductRhythmRecord)
	g.DELETE("delete", DeleteProductRhythmRecord)
	g.GET("all", GetAllProductRhythmRecord)
	g.GET("detail", GetProductRhythmRecordDetail)
}
