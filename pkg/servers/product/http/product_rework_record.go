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

// AddProductReworkRecord godoc
// @Summary 新增
// @Description 新增
// @Tags 返工记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductReworkRecordInfo true "Add ProductReworkRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productreworkrecord/add [post]
func AddProductReworkRecord(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductReworkRecordInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建返工记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductReworkRecord(model.PBToProductReworkRecord(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductReworkRecord godoc
// @Summary 更新
// @Description 更新
// @Tags 返工记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductReworkRecordInfo true "Update ProductReworkRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productreworkrecord/update [put]
func UpdateProductReworkRecord(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductReworkRecordInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新返工记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductReworkRecord(model.PBToProductReworkRecord(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductReworkRecord godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 返工记录管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productionLineID query string false "生产产线ID"
// @Param productSerialNo query string false "产品序列号"
// @Param productOrderNo query string false "生产工单号"
// @Param createTime0 query string false "创建时间开始"
// @Param createTime1 query string false "创建时间结束"
// @Param reworkBrief query string false "故障信息"
// @Success 200 {object} proto.QueryProductReworkRecordResponse
// @Router /api/mom/product/productreworkrecord/query [get]
func QueryProductReworkRecord(c *gin.Context) {
	req := &proto.QueryProductReworkRecordRequest{}
	resp := &proto.QueryProductReworkRecordResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductReworkRecord(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductReworkRecord godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 返工记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductReworkRecordResponse
// @Router /api/mom/product/productreworkrecord/all [get]
func GetAllProductReworkRecord(c *gin.Context) {
	resp := &proto.GetAllProductReworkRecordResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductReworkRecords()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductReworkRecordsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductReworkRecordDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 返工记录管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductReworkRecordDetailResponse
// @Router /api/mom/product/productreworkrecord/detail [get]
func GetProductReworkRecordDetail(c *gin.Context) {
	resp := &proto.GetProductReworkRecordDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductReworkRecordByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductReworkRecordToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductReworkRecord godoc
// @Summary 删除
// @Description 删除
// @Tags 返工记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductReworkRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productreworkrecord/delete [delete]
func DeleteProductReworkRecord(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除返工记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductReworkRecord(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductReworkRecordRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productreworkrecord")

	g.POST("add", AddProductReworkRecord)
	g.PUT("update", UpdateProductReworkRecord)
	g.GET("query", QueryProductReworkRecord)
	g.DELETE("delete", DeleteProductReworkRecord)
	g.GET("all", GetAllProductReworkRecord)
	g.GET("detail", GetProductReworkRecordDetail)
}
