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

// AddProductPackageMatchRule godoc
// @Summary 新增
// @Description 新增
// @Tags 包装匹配规则管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductPackageMatchRuleInfo true "Add ProductPackageMatchRule"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productpackagematchrule/add [post]
func AddProductPackageMatchRule(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductPackageMatchRuleInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建包装匹配规则请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductPackageMatchRule(model.PBToProductPackageMatchRule(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductPackageMatchRule godoc
// @Summary 更新
// @Description 更新
// @Tags 包装匹配规则管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductPackageMatchRuleInfo true "Update ProductPackageMatchRule"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productpackagematchrule/update [put]
func UpdateProductPackageMatchRule(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductPackageMatchRuleInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新包装匹配规则请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductPackageMatchRule(model.PBToProductPackageMatchRule(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductPackageMatchRule godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 包装匹配规则管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Success 200 {object} proto.QueryProductPackageMatchRuleResponse
// @Router /api/mom/product/productpackagematchrule/query [get]
func QueryProductPackageMatchRule(c *gin.Context) {
	req := &proto.QueryProductPackageMatchRuleRequest{}
	resp := &proto.QueryProductPackageMatchRuleResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductPackageMatchRule(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductPackageMatchRule godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 包装匹配规则管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductPackageMatchRuleResponse
// @Router /api/mom/product/productpackagematchrule/all [get]
func GetAllProductPackageMatchRule(c *gin.Context) {
	resp := &proto.GetAllProductPackageMatchRuleResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductPackageMatchRules()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductPackageMatchRulesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductPackageMatchRuleDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 包装匹配规则管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductPackageMatchRuleDetailResponse
// @Router /api/mom/product/productpackagematchrule/detail [get]
func GetProductPackageMatchRuleDetail(c *gin.Context) {
	resp := &proto.GetProductPackageMatchRuleDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductPackageMatchRuleByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductPackageMatchRuleToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductPackageMatchRule godoc
// @Summary 删除
// @Description 删除
// @Tags 包装匹配规则管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductPackageMatchRule"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productpackagematchrule/delete [delete]
func DeleteProductPackageMatchRule(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除包装匹配规则请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductPackageMatchRule(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductPackageMatchRuleRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productpackagematchrule")

	g.POST("add", AddProductPackageMatchRule)
	g.PUT("update", UpdateProductPackageMatchRule)
	g.GET("query", QueryProductPackageMatchRule)
	g.DELETE("delete", DeleteProductPackageMatchRule)
	g.GET("all", GetAllProductPackageMatchRule)
	g.GET("detail", GetProductPackageMatchRuleDetail)
}
