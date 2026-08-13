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

// AddProductIssueRecord godoc
// @Summary 新增
// @Description 新增
// @Tags 发料记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductIssueRecordInfo true "Add ProductIssueRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productissuerecord/add [post]
func AddProductIssueRecord(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductIssueRecordInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建发料记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductIssueRecord(model.PBToProductIssueRecord(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductIssueRecord godoc
// @Summary 更新
// @Description 更新
// @Tags 发料记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductIssueRecordInfo true "Update ProductIssueRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productissuerecord/update [put]
func UpdateProductIssueRecord(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductIssueRecordInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新发料记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductIssueRecord(model.PBToProductIssueRecord(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductIssueRecord godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 发料记录管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productSerialNo query string false "产品序列号"
// @Param productOrderNo query string false "生产工单号"
// @Param materialDescription query string false "物料信息"
// @Param createTime0 query string false "发料时间"
// @Param createTime1 query string false "发料时间"
// @Param productionProcess query string false "发料信息"
// @Success 200 {object} proto.QueryProductIssueRecordResponse
// @Router /api/mom/product/productissuerecord/query [get]
func QueryProductIssueRecord(c *gin.Context) {
	req := &proto.QueryProductIssueRecordRequest{}
	resp := &proto.QueryProductIssueRecordResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductIssueRecord(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductIssueRecord godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 发料记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductIssueRecordResponse
// @Router /api/mom/product/productissuerecord/all [get]
func GetAllProductIssueRecord(c *gin.Context) {
	resp := &proto.GetAllProductIssueRecordResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductIssueRecords()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductIssueRecordsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductIssueRecordDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 发料记录管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductIssueRecordDetailResponse
// @Router /api/mom/product/productissuerecord/detail [get]
func GetProductIssueRecordDetail(c *gin.Context) {
	resp := &proto.GetProductIssueRecordDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductIssueRecordByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductIssueRecordToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductIssueRecord godoc
// @Summary 删除
// @Description 删除
// @Tags 发料记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductIssueRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productissuerecord/delete [delete]
func DeleteProductIssueRecord(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除发料记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductIssueRecord(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductIssueRecordRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productissuerecord")

	g.POST("add", AddProductIssueRecord)
	g.PUT("update", UpdateProductIssueRecord)
	g.GET("query", QueryProductIssueRecord)
	g.DELETE("delete", DeleteProductIssueRecord)
	g.GET("all", GetAllProductIssueRecord)
	g.GET("detail", GetProductIssueRecordDetail)
}
