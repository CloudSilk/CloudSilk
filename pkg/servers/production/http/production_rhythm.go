package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/production/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddProductionRhythm godoc
// @Summary 新增
// @Description 新增
// @Tags 生产节拍管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionRhythmInfo true "Add ProductionRhythm"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionrhythm/add [post]
func AddProductionRhythm(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionRhythmInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建生产节拍请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductionRhythm(model.PBToProductionRhythm(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductionRhythm godoc
// @Summary 更新
// @Description 更新
// @Tags 生产节拍管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionRhythmInfo true "Update ProductionRhythm"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionrhythm/update [put]
func UpdateProductionRhythm(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionRhythmInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新生产节拍请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductionRhythm(model.PBToProductionRhythm(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductionRhythm godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 生产节拍管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productionLineID query string false "生产产线ID"
// @Success 200 {object} proto.QueryProductionRhythmResponse
// @Router /api/mom/production/productionrhythm/query [get]
func QueryProductionRhythm(c *gin.Context) {
	req := &proto.QueryProductionRhythmRequest{}
	resp := &proto.QueryProductionRhythmResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductionRhythm(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductionRhythm godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 生产节拍管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductionRhythmResponse
// @Router /api/mom/production/productionrhythm/all [get]
func GetAllProductionRhythm(c *gin.Context) {
	resp := &proto.GetAllProductionRhythmResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductionRhythms()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductionRhythmsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductionRhythmDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 生产节拍管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductionRhythmDetailResponse
// @Router /api/mom/production/productionrhythm/detail [get]
func GetProductionRhythmDetail(c *gin.Context) {
	resp := &proto.GetProductionRhythmDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductionRhythmByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionRhythmToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductionRhythm godoc
// @Summary 删除
// @Description 删除
// @Tags 生产节拍管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductionRhythm"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionrhythm/delete [delete]
func DeleteProductionRhythm(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除生产节拍请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductionRhythm(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductionRhythmRouter(r *gin.Engine) {
	g := r.Group("/api/mom/production/productionrhythm")

	g.POST("add", AddProductionRhythm)
	g.PUT("update", UpdateProductionRhythm)
	g.GET("query", QueryProductionRhythm)
	g.DELETE("delete", DeleteProductionRhythm)
	g.GET("all", GetAllProductionRhythm)
	g.GET("detail", GetProductionRhythmDetail)
}
