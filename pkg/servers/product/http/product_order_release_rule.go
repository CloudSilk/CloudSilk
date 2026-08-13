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

// AddProductOrderReleaseRule godoc
// @Summary 新增
// @Description 新增
// @Tags 工单发放规则
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductOrderReleaseRuleInfo true "Add ProductOrderReleaseRule"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderreleaserule/add [post]
func AddProductOrderReleaseRule(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductOrderReleaseRuleInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建工单发放规则请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductOrderReleaseRule(model.PBToProductOrderReleaseRule(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductOrderReleaseRule godoc
// @Summary 更新
// @Description 更新
// @Tags 工单发放规则
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductOrderReleaseRuleInfo true "Update ProductOrderReleaseRule"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderreleaserule/update [put]
func UpdateProductOrderReleaseRule(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductOrderReleaseRuleInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新工单发放规则请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductOrderReleaseRule(model.PBToProductOrderReleaseRule(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductOrderReleaseRule godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 工单发放规则
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productCategoryID query string false "产品类别ID"
// @Param productAttributeID query string false "目标特性ID"
// @Success 200 {object} proto.QueryProductOrderReleaseRuleResponse
// @Router /api/mom/product/productorderreleaserule/query [get]
func QueryProductOrderReleaseRule(c *gin.Context) {
	req := &proto.QueryProductOrderReleaseRuleRequest{}
	resp := &proto.QueryProductOrderReleaseRuleResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductOrderReleaseRule(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductOrderReleaseRule godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 工单发放规则
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductOrderReleaseRuleResponse
// @Router /api/mom/product/productorderreleaserule/all [get]
func GetAllProductOrderReleaseRule(c *gin.Context) {
	resp := &proto.GetAllProductOrderReleaseRuleResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductOrderReleaseRules()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductOrderReleaseRulesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductOrderReleaseRuleDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 工单发放规则
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductOrderReleaseRuleDetailResponse
// @Router /api/mom/product/productorderreleaserule/detail [get]
func GetProductOrderReleaseRuleDetail(c *gin.Context) {
	resp := &proto.GetProductOrderReleaseRuleDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductOrderReleaseRuleByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductOrderReleaseRuleToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductOrderReleaseRule godoc
// @Summary 删除
// @Description 删除
// @Tags 工单发放规则
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductOrderReleaseRule"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderreleaserule/delete [delete]
func DeleteProductOrderReleaseRule(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除工单发放规则请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductOrderReleaseRule(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductOrderReleaseRuleRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productorderreleaserule")

	g.POST("add", AddProductOrderReleaseRule)
	g.PUT("update", UpdateProductOrderReleaseRule)
	g.GET("query", QueryProductOrderReleaseRule)
	g.DELETE("delete", DeleteProductOrderReleaseRule)
	g.GET("all", GetAllProductOrderReleaseRule)
	g.GET("detail", GetProductOrderReleaseRuleDetail)
}
