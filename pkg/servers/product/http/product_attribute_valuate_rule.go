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

// AddProductAttributeValuateRule godoc
// @Summary 新增
// @Description 新增
// @Tags 特性赋值规则
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductAttributeValuateRuleInfo true "Add ProductAttributeValuateRule"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productattributevaluaterule/add [post]
func AddProductAttributeValuateRule(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductAttributeValuateRuleInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建特性赋值规则请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductAttributeValuateRule(model.PBToProductAttributeValuateRule(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductAttributeValuateRule godoc
// @Summary 更新
// @Description 更新
// @Tags 特性赋值规则
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductAttributeValuateRuleInfo true "Update ProductAttributeValuateRule"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productattributevaluaterule/update [put]
func UpdateProductAttributeValuateRule(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductAttributeValuateRuleInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新特性赋值规则请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductAttributeValuateRule(model.PBToProductAttributeValuateRule(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductAttributeValuateRule godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 特性赋值规则
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productCategoryID query string false "产品类别ID"
// @Param productAttributeID query string false "目标特性ID"
// @Success 200 {object} proto.QueryProductAttributeValuateRuleResponse
// @Router /api/mom/product/productattributevaluaterule/query [get]
func QueryProductAttributeValuateRule(c *gin.Context) {
	req := &proto.QueryProductAttributeValuateRuleRequest{}
	resp := &proto.QueryProductAttributeValuateRuleResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductAttributeValuateRule(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductAttributeValuateRule godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 特性赋值规则
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductAttributeValuateRuleResponse
// @Router /api/mom/product/productattributevaluaterule/all [get]
func GetAllProductAttributeValuateRule(c *gin.Context) {
	resp := &proto.GetAllProductAttributeValuateRuleResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductAttributeValuateRules()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductAttributeValuateRulesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductAttributeValuateRuleDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 特性赋值规则
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductAttributeValuateRuleDetailResponse
// @Router /api/mom/product/productattributevaluaterule/detail [get]
func GetProductAttributeValuateRuleDetail(c *gin.Context) {
	resp := &proto.GetProductAttributeValuateRuleDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductAttributeValuateRuleByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductAttributeValuateRuleToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductAttributeValuateRule godoc
// @Summary 删除
// @Description 删除
// @Tags 特性赋值规则
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductAttributeValuateRule"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productattributevaluaterule/delete [delete]
func DeleteProductAttributeValuateRule(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除特性赋值规则请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductAttributeValuateRule(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductAttributeValuateRuleRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productattributevaluaterule")

	g.POST("add", AddProductAttributeValuateRule)
	g.PUT("update", UpdateProductAttributeValuateRule)
	g.GET("query", QueryProductAttributeValuateRule)
	g.DELETE("delete", DeleteProductAttributeValuateRule)
	g.GET("all", GetAllProductAttributeValuateRule)
	g.GET("detail", GetProductAttributeValuateRuleDetail)
}
