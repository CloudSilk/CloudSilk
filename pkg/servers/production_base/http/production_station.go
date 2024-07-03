package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/clients"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/production_base/logic"
	"github.com/CloudSilk/pkg/utils/log"
	usercenter "github.com/CloudSilk/usercenter/proto"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddProductionStation godoc
// @Summary 新增
// @Description 新增
// @Tags 工站管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionStationInfo true "Add ProductionStation"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productionbase/productionstation/add [post]
func AddProductionStation(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionStationInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建工站请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductionStation(model.PBToProductionStation(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductionStation godoc
// @Summary 更新
// @Description 更新
// @Tags 工站管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionStationInfo true "Update ProductionStation"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productionbase/productionstation/update [put]
func UpdateProductionStation(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionStationInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新工站请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductionStation(model.PBToProductionStation(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductionStation godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 工站管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Param stationType query string false "工位类型"
// @Param productionLineID query string false "生产产线ID"
// @Success 200 {object} proto.QueryProductionStationResponse
// @Router /api/mom/productionbase/productionstation/query [get]
func QueryProductionStation(c *gin.Context) {
	req := &proto.QueryProductionStationRequest{}
	resp := &proto.QueryProductionStationResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductionStation(req, resp, false)
	if resp.Code == proto.Code_Success {
		r := &usercenter.QueryUserRequest{}
		for _, u := range resp.Data {
			r.Ids = append(r.Ids, u.CurrentUserID)
		}
		r.PageSize = int64(len(r.Ids))
		users, err := clients.UserClient.Query(context.Background(), r)
		if err == nil && users.Code == usercenter.Code_Success {
			for _, u := range resp.Data {
				for _, u2 := range users.Data {
					if u.CurrentUserID == u2.Id {
						u.CurrentUserName = u2.Nickname
					}
				}
			}
		}
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllProductionStation godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 工站管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductionStationResponse
// @Router /api/mom/productionbase/productionstation/all [get]
func GetAllProductionStation(c *gin.Context) {
	resp := &proto.GetAllProductionStationResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductionStations()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductionStationsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductionStationDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 工站管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductionStationDetailResponse
// @Router /api/mom/productionbase/productionstation/detail [get]
func GetProductionStationDetail(c *gin.Context) {
	resp := &proto.GetProductionStationDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductionStationByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionStationToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductionStation godoc
// @Summary 删除
// @Description 删除
// @Tags 工站管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductionStation"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/productionbase/productionstation/delete [delete]
func DeleteProductionStation(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除工站请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductionStation(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductionStationRouter(r *gin.Engine) {
	g := r.Group("/api/mom/productionbase/productionstation")

	g.POST("add", AddProductionStation)
	g.PUT("update", UpdateProductionStation)
	g.GET("query", QueryProductionStation)
	g.DELETE("delete", DeleteProductionStation)
	g.GET("all", GetAllProductionStation)
	g.GET("detail", GetProductionStationDetail)
}
