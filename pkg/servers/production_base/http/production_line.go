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

// AddProductionLine godoc
// @Summary 新增
// @Description 新增
// @Tags 产线管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionLineInfo true "Add ProductionLine"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productionbase/productionline/add [post]
func AddProductionLine(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionLineInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建产线请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductionLine(model.PBToProductionLine(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductionLine godoc
// @Summary 更新
// @Description 更新
// @Tags 产线管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionLineInfo true "Update ProductionLine"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productionbase/productionline/update [put]
func UpdateProductionLine(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionLineInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新产线请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductionLine(model.PBToProductionLine(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductionLine godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 产线管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Param productionFactoryID query string false "生产工厂ID "
// @Success 200 {object} proto.QueryProductionLineResponse
// @Router /api/mom/productionbase/productionline/query [get]
func QueryProductionLine(c *gin.Context) {
	req := &proto.QueryProductionLineRequest{}
	resp := &proto.QueryProductionLineResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductionLine(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductionLine godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 产线管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductionLineResponse
// @Router /api/mom/productionbase/productionline/all [get]
func GetAllProductionLine(c *gin.Context) {
	resp := &proto.GetAllProductionLineResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductionLines()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductionLinesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductionLineDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 产线管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductionLineDetailResponse
// @Router /api/mom/productionbase/productionline/detail [get]
func GetProductionLineDetail(c *gin.Context) {
	resp := &proto.GetProductionLineDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductionLineByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionLineToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductionLine godoc
// @Summary 删除
// @Description 删除
// @Tags 产线管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductionLine"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productionbase/productionline/delete [delete]
func DeleteProductionLine(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除产线请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductionLine(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductionLineRouter(r *gin.Engine) {
	g := r.Group("/api/mom/productionbase/productionline")

	g.POST("add", AddProductionLine)
	g.PUT("update", UpdateProductionLine)
	g.GET("query", QueryProductionLine)
	g.DELETE("delete", DeleteProductionLine)
	g.GET("all", GetAllProductionLine)
	g.GET("detail", GetProductionLineDetail)
}
