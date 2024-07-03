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

// AddProductWorkRecord godoc
// @Summary 新增
// @Description 新增
// @Tags 产品作业记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductWorkRecordInfo true "Add ProductWorkRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productworkrecord/add [post]
func AddProductWorkRecord(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductWorkRecordInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建产品作业记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductWorkRecord(model.PBToProductWorkRecord(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductWorkRecord godoc
// @Summary 更新
// @Description 更新
// @Tags 产品作业记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductWorkRecordInfo true "Update ProductWorkRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productworkrecord/update [put]
func UpdateProductWorkRecord(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductWorkRecordInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新产品作业记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductWorkRecord(model.PBToProductWorkRecord(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductWorkRecord godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 产品作业记录管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productOrderNo query string false "生产工单号"
// @Param productSerialNo query string false "产品序列号"
// @Param workStartTime0 query string false "开始作业时间开始"
// @Param workStartTime1 query string false "开始作业时间结束"
// @Param productionLineID query string false "作业产线"
// @Success 200 {object} proto.QueryProductWorkRecordResponse
// @Router /api/mom/product/productworkrecord/query [get]
func QueryProductWorkRecord(c *gin.Context) {
	req := &proto.QueryProductWorkRecordRequest{}
	resp := &proto.QueryProductWorkRecordResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductWorkRecord(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductWorkRecord godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 产品作业记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductWorkRecordResponse
// @Router /api/mom/product/productworkrecord/all [get]
func GetAllProductWorkRecord(c *gin.Context) {
	resp := &proto.GetAllProductWorkRecordResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductWorkRecords()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductWorkRecordsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductWorkRecordDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 产品作业记录管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductWorkRecordDetailResponse
// @Router /api/mom/product/productworkrecord/detail [get]
func GetProductWorkRecordDetail(c *gin.Context) {
	resp := &proto.GetProductWorkRecordDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductWorkRecordByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductWorkRecordToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductWorkRecord godoc
// @Summary 删除
// @Description 删除
// @Tags 产品作业记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductWorkRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productworkrecord/delete [delete]
func DeleteProductWorkRecord(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除产品作业记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductWorkRecord(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductWorkRecordRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productworkrecord")

	g.POST("add", AddProductWorkRecord)
	g.PUT("update", UpdateProductWorkRecord)
	g.GET("query", QueryProductWorkRecord)
	g.DELETE("delete", DeleteProductWorkRecord)
	g.GET("all", GetAllProductWorkRecord)
	g.GET("detail", GetProductWorkRecordDetail)
}
