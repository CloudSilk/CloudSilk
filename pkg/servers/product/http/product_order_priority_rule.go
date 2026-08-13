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

// AddProductOrderPriorityRule godoc
// @Summary 新增
// @Description 新增
// @Tags 工单优先规则
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductOrderPriorityRuleInfo true "Add ProductOrderPriorityRule"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderpriorityrule/add [post]
func AddProductOrderPriorityRule(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductOrderPriorityRuleInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建工单优先规则请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductOrderPriorityRule(model.PBToProductOrderPriorityRule(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductOrderPriorityRule godoc
// @Summary 更新
// @Description 更新
// @Tags 工单优先规则
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductOrderPriorityRuleInfo true "Update ProductOrderPriorityRule"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderpriorityrule/update [put]
func UpdateProductOrderPriorityRule(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductOrderPriorityRuleInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新工单优先规则请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductOrderPriorityRule(model.PBToProductOrderPriorityRule(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductOrderPriorityRule godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 工单优先规则
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param priorityLevel query int false "生产优先级"
// @Success 200 {object} proto.QueryProductOrderPriorityRuleResponse
// @Router /api/mom/product/productorderpriorityrule/query [get]
func QueryProductOrderPriorityRule(c *gin.Context) {
	req := &proto.QueryProductOrderPriorityRuleRequest{}
	resp := &proto.QueryProductOrderPriorityRuleResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductOrderPriorityRule(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductOrderPriorityRule godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 工单优先规则
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductOrderPriorityRuleResponse
// @Router /api/mom/product/productorderpriorityrule/all [get]
func GetAllProductOrderPriorityRule(c *gin.Context) {
	resp := &proto.GetAllProductOrderPriorityRuleResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductOrderPriorityRules()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductOrderPriorityRulesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductOrderPriorityRuleDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 工单优先规则
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductOrderPriorityRuleDetailResponse
// @Router /api/mom/product/productorderpriorityrule/detail [get]
func GetProductOrderPriorityRuleDetail(c *gin.Context) {
	resp := &proto.GetProductOrderPriorityRuleDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductOrderPriorityRuleByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductOrderPriorityRuleToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductOrderPriorityRule godoc
// @Summary 删除
// @Description 删除
// @Tags 工单优先规则
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductOrderPriorityRule"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productorderpriorityrule/delete [delete]
func DeleteProductOrderPriorityRule(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除工单优先规则请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductOrderPriorityRule(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductOrderPriorityRuleRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productorderpriorityrule")

	g.POST("add", AddProductOrderPriorityRule)
	g.PUT("update", UpdateProductOrderPriorityRule)
	g.GET("query", QueryProductOrderPriorityRule)
	g.DELETE("delete", DeleteProductOrderPriorityRule)
	g.GET("all", GetAllProductOrderPriorityRule)
	g.GET("detail", GetProductOrderPriorityRuleDetail)
}
