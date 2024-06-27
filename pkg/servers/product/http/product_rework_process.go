package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/product/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/pkg/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddProductReworkProcess godoc
// @Summary 新增
// @Description 新增
// @Tags 返工工序管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductReworkProcessInfo true "Add ProductReworkProcess"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productreworkprocess/add [post]
func AddProductReworkProcess(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductReworkProcessInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建返工工序请求参数无效:%v", transID, err)
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
		var availableStations []*proto.ProductReworkProcessAvailableStationInfo
		for _, productionStationID := range req.AvailableStationIDs {
			availableStations = append(availableStations, &proto.ProductReworkProcessAvailableStationInfo{
				ProductionStationID: productionStationID,
			})
		}
		req.ProductionStations = availableStations
	}
	if len(req.AvailableProcessIDs) > 0 {
		var availableProcesss []*proto.ProductReworkProcessAvailableProcessInfo
		for _, productionProcessID := range req.AvailableProcessIDs {
			availableProcesss = append(availableProcesss, &proto.ProductReworkProcessAvailableProcessInfo{
				ProductionProcessID: productionProcessID,
			})
		}
		req.ProductionProcesses = availableProcesss
	}

	id, err := logic.CreateProductReworkProcess(model.PBToProductReworkProcess(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductReworkProcess godoc
// @Summary 更新
// @Description 更新
// @Tags 返工工序管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductReworkProcessInfo true "Update ProductReworkProcess"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productreworkprocess/update [put]
func UpdateProductReworkProcess(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductReworkProcessInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新返工工序请求参数无效:%v", transID, err)
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
		var availableStations []*proto.ProductReworkProcessAvailableStationInfo
		for _, productionStationID := range req.AvailableStationIDs {
			availableStations = append(availableStations, &proto.ProductReworkProcessAvailableStationInfo{
				ProductionStationID: productionStationID,
			})
		}
		req.ProductionStations = availableStations
	}
	if len(req.AvailableProcessIDs) > 0 {
		var availableProcesss []*proto.ProductReworkProcessAvailableProcessInfo
		for _, productionProcessID := range req.AvailableProcessIDs {
			availableProcesss = append(availableProcesss, &proto.ProductReworkProcessAvailableProcessInfo{
				ProductionProcessID: productionProcessID,
			})
		}
		req.ProductionProcesses = availableProcesss
	}

	err = logic.UpdateProductReworkProcess(model.PBToProductReworkProcess(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductReworkProcess godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 返工工序管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productionLineID query string false "生产产线ID"
// @Param code query string false "代号或描述"
// @Success 200 {object} proto.QueryProductReworkProcessResponse
// @Router /api/mom/product/productreworkprocess/query [get]
func QueryProductReworkProcess(c *gin.Context) {
	req := &proto.QueryProductReworkProcessRequest{}
	resp := &proto.QueryProductReworkProcessResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductReworkProcess(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductReworkProcess godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 返工工序管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductReworkProcessResponse
// @Router /api/mom/product/productreworkprocess/all [get]
func GetAllProductReworkProcess(c *gin.Context) {
	resp := &proto.GetAllProductReworkProcessResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductReworkProcesss()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductReworkProcesssToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductReworkProcessDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 返工工序管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductReworkProcessDetailResponse
// @Router /api/mom/product/productreworkprocess/detail [get]
func GetProductReworkProcessDetail(c *gin.Context) {
	resp := &proto.GetProductReworkProcessDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductReworkProcessByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductReworkProcessToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductReworkProcess godoc
// @Summary 删除
// @Description 删除
// @Tags 返工工序管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductReworkProcess"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productreworkprocess/delete [delete]
func DeleteProductReworkProcess(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除返工工序请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductReworkProcess(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductReworkProcessRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productreworkprocess")

	g.POST("add", AddProductReworkProcess)
	g.PUT("update", UpdateProductReworkProcess)
	g.GET("query", QueryProductReworkProcess)
	g.DELETE("delete", DeleteProductReworkProcess)
	g.GET("all", GetAllProductReworkProcess)
	g.GET("detail", GetProductReworkProcessDetail)
}
