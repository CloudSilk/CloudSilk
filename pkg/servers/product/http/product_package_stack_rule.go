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

// AddProductPackageStackRule godoc
// @Summary 新增
// @Description 新增
// @Tags 产品包装码垛规则
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductPackageStackRuleInfo true "Add ProductPackageStackRule"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productpackagestackrule/add [post]
func AddProductPackageStackRule(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductPackageStackRuleInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建产品信息管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductPackageStackRule(model.PBToProductPackageStackRule(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductPackageStackRule godoc
// @Summary 更新
// @Description 更新
// @Tags 产品包装码垛规则
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductPackageStackRuleInfo true "Update ProductPackageStackRule"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productpackagestackrule/update [put]
func UpdateProductPackageStackRule(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductPackageStackRuleInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新产品信息管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductPackageStackRule(model.PBToProductPackageStackRule(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductPackageStackRule godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 产品包装码垛规则
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productSerialNo query string false "产品序列号"
// @Param createTime0 query string false "创建时间开始"
// @Param createTime1 query string false "创建时间结束"
// @Param currentState query string false "当前状态"
// @Param productOrderNo query string false "生产工单号"
// @Success 200 {object} proto.QueryProductPackageStackRuleResponse
// @Router /api/mom/product/productpackagestackrule/query [get]
func QueryProductPackageStackRule(c *gin.Context) {
	req := &proto.QueryProductPackageStackRuleRequest{}
	resp := &proto.QueryProductPackageStackRuleResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductPackageStackRule(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductPackageStackRule godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 产品包装码垛规则
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductPackageStackRuleResponse
// @Router /api/mom/product/productpackagestackrule/all [get]
func GetAllProductPackageStackRule(c *gin.Context) {
	resp := &proto.GetAllProductPackageStackRuleResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductPackageStackRules()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductPackageStackRulesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductPackageStackRuleDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 产品包装码垛规则
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductPackageStackRuleDetailResponse
// @Router /api/mom/product/productpackagestackrule/detail [get]
func GetProductPackageStackRuleDetail(c *gin.Context) {
	resp := &proto.GetProductPackageStackRuleDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductPackageStackRuleByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductPackageStackRuleToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductPackageStackRule godoc
// @Summary 删除
// @Description 删除
// @Tags 产品包装码垛规则
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductPackageStackRule"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productpackagestackrule/delete [delete]
func DeleteProductPackageStackRule(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除产品信息管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductPackageStackRule(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductPackageStackRuleRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productpackagestackrule")

	g.POST("add", AddProductPackageStackRule)
	g.PUT("update", UpdateProductPackageStackRule)
	g.GET("query", QueryProductPackageStackRule)
	g.DELETE("delete", DeleteProductPackageStackRule)
	g.GET("all", GetAllProductPackageStackRule)
	g.GET("detail", GetProductPackageStackRuleDetail)
}
