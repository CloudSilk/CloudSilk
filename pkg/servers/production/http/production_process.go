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

// AddProductionProcess godoc
// @Summary 新增
// @Description 新增
// @Tags 生产工序管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionProcessInfo true "Add ProductionProcess"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionprocess/add [post]
func AddProductionProcess(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionProcessInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建生产工序请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	if len(req.AvailableStationIDs) > 0 {
		var availableStations []*proto.ProductionProcessAvailableStationInfo
		for _, productionStationID := range req.AvailableStationIDs {
			availableStations = append(availableStations, &proto.ProductionProcessAvailableStationInfo{
				ProductionStationID: productionStationID,
			})
		}
		req.ProductionProcessAvailableStations = availableStations
	}

	id, err := logic.CreateProductionProcess(model.PBToProductionProcess(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductionProcess godoc
// @Summary 更新
// @Description 更新
// @Tags 生产工序管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionProcessInfo true "Update ProductionProcess"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionprocess/update [put]
func UpdateProductionProcess(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionProcessInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新生产工序请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	if len(req.AvailableStationIDs) > 0 {
		var availableStations []*proto.ProductionProcessAvailableStationInfo
		for _, productionStationID := range req.AvailableStationIDs {
			availableStations = append(availableStations, &proto.ProductionProcessAvailableStationInfo{
				ProductionStationID: productionStationID,
			})
		}
		req.ProductionProcessAvailableStations = availableStations
	}

	err = logic.UpdateProductionProcess(model.PBToProductionProcess(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductionProcess godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 生产工序管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productionLineID query string false "生产产线ID"
// @Success 200 {object} proto.QueryProductionProcessResponse
// @Router /api/mom/production/productionprocess/query [get]
func QueryProductionProcess(c *gin.Context) {
	req := &proto.QueryProductionProcessRequest{}
	resp := &proto.QueryProductionProcessResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductionProcess(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductionProcess godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 生产工序管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductionProcessResponse
// @Router /api/mom/production/productionprocess/all [get]
func GetAllProductionProcess(c *gin.Context) {
	resp := &proto.GetAllProductionProcessResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductionProcesss()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductionProcesssToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductionProcessDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 生产工序管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductionProcessDetailResponse
// @Router /api/mom/production/productionprocess/detail [get]
func GetProductionProcessDetail(c *gin.Context) {
	resp := &proto.GetProductionProcessDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductionProcessByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionProcessToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductionProcess godoc
// @Summary 删除
// @Description 删除
// @Tags 生产工序管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductionProcess"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionprocess/delete [delete]
func DeleteProductionProcess(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除生产工序请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductionProcess(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductionProcessRouter(r *gin.Engine) {
	g := r.Group("/api/mom/production/productionprocess")

	g.POST("add", AddProductionProcess)
	g.PUT("update", UpdateProductionProcess)
	g.GET("query", QueryProductionProcess)
	g.DELETE("delete", DeleteProductionProcess)
	g.GET("all", GetAllProductionProcess)
	g.GET("detail", GetProductionProcessDetail)
}
