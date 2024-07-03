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

// AddProductionProcessStep godoc
// @Summary 新增
// @Description 新增
// @Tags 生产工步管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionProcessStepInfo true "Add ProductionProcessStep"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionprocessstep/add [post]
func AddProductionProcessStep(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionProcessStepInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建生产工步请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	if len(req.AvailableProcessIDs) > 0 {
		var availableProcesses []*proto.AvailableProcessInfo
		for _, productionProcessID := range req.AvailableProcessIDs {
			availableProcesses = append(availableProcesses, &proto.AvailableProcessInfo{
				ProductionProcessID: productionProcessID,
			})
		}
		req.AvailableProcesses = availableProcesses
	}

	id, err := logic.CreateProductionProcessStep(model.PBToProductionProcessStep(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductionProcessStep godoc
// @Summary 更新
// @Description 更新
// @Tags 生产工步管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionProcessStepInfo true "Update ProductionProcessStep"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionprocessstep/update [put]
func UpdateProductionProcessStep(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionProcessStepInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新生产工步请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	if len(req.AvailableProcessIDs) > 0 {
		var availableProcesses []*proto.AvailableProcessInfo
		for _, productionProcessID := range req.AvailableProcessIDs {
			availableProcesses = append(availableProcesses, &proto.AvailableProcessInfo{
				ProductionProcessID: productionProcessID,
			})
		}
		req.AvailableProcesses = availableProcesses
	}

	err = logic.UpdateProductionProcessStep(model.PBToProductionProcessStep(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductionProcessStep godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 生产工步管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productionProcessID query string false "生产工序ID"
// @Success 200 {object} proto.QueryProductionProcessStepResponse
// @Router /api/mom/production/productionprocessstep/query [get]
func QueryProductionProcessStep(c *gin.Context) {
	req := &proto.QueryProductionProcessStepRequest{}
	resp := &proto.QueryProductionProcessStepResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductionProcessStep(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductionProcessStep godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 生产工步管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductionProcessStepResponse
// @Router /api/mom/production/productionprocessstep/all [get]
func GetAllProductionProcessStep(c *gin.Context) {
	resp := &proto.GetAllProductionProcessStepResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductionProcessSteps()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductionProcessStepsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductionProcessStepDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 生产工步管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductionProcessStepDetailResponse
// @Router /api/mom/production/productionprocessstep/detail [get]
func GetProductionProcessStepDetail(c *gin.Context) {
	resp := &proto.GetProductionProcessStepDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductionProcessStepByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionProcessStepToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductionProcessStep godoc
// @Summary 删除
// @Description 删除
// @Tags 生产工步管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductionProcessStep"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionprocessstep/delete [delete]
func DeleteProductionProcessStep(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除生产工步请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductionProcessStep(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductionProcessStepRouter(r *gin.Engine) {
	g := r.Group("/api/mom/production/productionprocessstep")

	g.POST("add", AddProductionProcessStep)
	g.PUT("update", UpdateProductionProcessStep)
	g.GET("query", QueryProductionProcessStep)
	g.DELETE("delete", DeleteProductionProcessStep)
	g.GET("all", GetAllProductionProcessStep)
	g.GET("detail", GetProductionProcessStepDetail)
}
