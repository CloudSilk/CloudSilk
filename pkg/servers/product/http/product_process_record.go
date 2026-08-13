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

// AddProductProcessRecord godoc
// @Summary 新增
// @Description 新增
// @Tags 产品过程记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductProcessRecordInfo true "Add ProductProcessRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productprocessrecord/add [post]
func AddProductProcessRecord(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductProcessRecordInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建产品过程记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductProcessRecord(model.PBToProductProcessRecord(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductProcessRecord godoc
// @Summary 更新
// @Description 更新
// @Tags 产品过程记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductProcessRecordInfo true "Update ProductProcessRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productprocessrecord/update [put]
func UpdateProductProcessRecord(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductProcessRecordInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新产品过程记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductProcessRecord(model.PBToProductProcessRecord(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductProcessRecord godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 产品过程记录管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productSerialNo query string false "产品序列号"
// @Param productOrderNo query string false "生产工单号"
// @Param workTime0 query string false "作业时间"
// @Param workTime1 query string false "作业时间"
// @Param workDescription query string false "作业信息"
// @Success 200 {object} proto.QueryProductProcessRecordResponse
// @Router /api/mom/product/productprocessrecord/query [get]
func QueryProductProcessRecord(c *gin.Context) {
	req := &proto.QueryProductProcessRecordRequest{}
	resp := &proto.QueryProductProcessRecordResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductProcessRecord(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductProcessRecord godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 产品过程记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductProcessRecordResponse
// @Router /api/mom/product/productprocessrecord/all [get]
func GetAllProductProcessRecord(c *gin.Context) {
	resp := &proto.GetAllProductProcessRecordResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductProcessRecords()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductProcessRecordsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductProcessRecordDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 产品过程记录管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductProcessRecordDetailResponse
// @Router /api/mom/product/productprocessrecord/detail [get]
func GetProductProcessRecordDetail(c *gin.Context) {
	resp := &proto.GetProductProcessRecordDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductProcessRecordByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductProcessRecordToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductProcessRecord godoc
// @Summary 删除
// @Description 删除
// @Tags 产品过程记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductProcessRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productprocessrecord/delete [delete]
func DeleteProductProcessRecord(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除产品过程记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductProcessRecord(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductProcessRecordRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productprocessrecord")

	g.POST("add", AddProductProcessRecord)
	g.PUT("update", UpdateProductProcessRecord)
	g.GET("query", QueryProductProcessRecord)
	g.DELETE("delete", DeleteProductProcessRecord)
	g.GET("all", GetAllProductProcessRecord)
	g.GET("detail", GetProductProcessRecordDetail)
}
