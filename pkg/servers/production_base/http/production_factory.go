package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/production_base/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddProductionFactory godoc
// @Summary 新增
// @Description 新增
// @Tags 工厂管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionFactoryInfo true "Add ProductionFactory"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productionbase/productionfactory/add [post]
func AddProductionFactory(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionFactoryInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建工厂请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductionFactory(model.PBToProductionFactory(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductionFactory godoc
// @Summary 更新
// @Description 更新
// @Tags 工厂管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionFactoryInfo true "Update ProductionFactory"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productionbase/productionfactory/update [put]
func UpdateProductionFactory(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionFactoryInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新工厂请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductionFactory(model.PBToProductionFactory(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductionFactory godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 工厂管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Success 200 {object} proto.QueryProductionFactoryResponse
// @Router /api/mom/productionbase/productionfactory/query [get]
func QueryProductionFactory(c *gin.Context) {
	req := &proto.QueryProductionFactoryRequest{}
	resp := &proto.QueryProductionFactoryResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductionFactory(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductionFactory godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 工厂管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductionFactoryResponse
// @Router /api/mom/productionbase/productionfactory/all [get]
func GetAllProductionFactory(c *gin.Context) {
	resp := &proto.GetAllProductionFactoryResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductionFactorys()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductionFactorysToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductionFactoryDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 工厂管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductionFactoryDetailResponse
// @Router /api/mom/productionbase/productionfactory/detail [get]
func GetProductionFactoryDetail(c *gin.Context) {
	resp := &proto.GetProductionFactoryDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductionFactoryByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionFactoryToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductionFactory godoc
// @Summary 删除
// @Description 删除
// @Tags 工厂管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductionFactory"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productionbase/productionfactory/delete [delete]
func DeleteProductionFactory(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除工厂请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductionFactory(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductionFactoryRouter(r *gin.Engine) {
	g := r.Group("/api/mom/productionbase/productionfactory")

	g.POST("add", AddProductionFactory)
	g.PUT("update", UpdateProductionFactory)
	g.GET("query", QueryProductionFactory)
	g.DELETE("delete", DeleteProductionFactory)
	g.GET("all", GetAllProductionFactory)
	g.GET("detail", GetProductionFactoryDetail)
}
