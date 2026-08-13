package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/clients"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/production/logic"
	"github.com/CloudSilk/pkg/utils/log"
	usercenter "github.com/CloudSilk/usercenter/proto"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddProductionStationBreakdown godoc
// @Summary 新增
// @Description 新增
// @Tags 工站故障记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionStationBreakdownInfo true "Add ProductionStationBreakdown"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionstationbreakdown/add [post]
func AddProductionStationBreakdown(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionStationBreakdownInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建工站故障记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductionStationBreakdown(model.PBToProductionStationBreakdown(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductionStationBreakdown godoc
// @Summary 更新
// @Description 更新
// @Tags 工站故障记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionStationBreakdownInfo true "Update ProductionStationBreakdown"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionstationbreakdown/update [put]
func UpdateProductionStationBreakdown(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionStationBreakdownInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新工站故障记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductionStationBreakdown(model.PBToProductionStationBreakdown(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductionStationBreakdown godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 工站故障记录管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param createTime0 query string false "创建时间开始"
// @Param createTime1 query string false "创建时间结束"
// @Param productionLineID query string false "生产产线ID"
// @Success 200 {object} proto.QueryProductionStationBreakdownResponse
// @Router /api/mom/production/productionstationbreakdown/query [get]
func QueryProductionStationBreakdown(c *gin.Context) {
	req := &proto.QueryProductionStationBreakdownRequest{}
	resp := &proto.QueryProductionStationBreakdownResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductionStationBreakdown(req, resp, false)
	if resp.Code == proto.Code_Success {
		r := &usercenter.QueryUserRequest{}
		for _, u := range resp.Data {
			r.Ids = append(r.Ids, u.CreateUserID)
		}
		r.PageSize = int64(len(r.Ids))
		users, err := clients.UserClient.Query(context.Background(), r)
		if err == nil && users.Code == usercenter.Code_Success {
			for _, u := range resp.Data {
				for _, u2 := range users.Data {
					if u.CreateUserID == u2.Id {
						u.CreateUserName = u2.Nickname
					}
				}
			}
		}
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllProductionStationBreakdown godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 工站故障记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductionStationBreakdownResponse
// @Router /api/mom/production/productionstationbreakdown/all [get]
func GetAllProductionStationBreakdown(c *gin.Context) {
	resp := &proto.GetAllProductionStationBreakdownResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductionStationBreakdowns()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductionStationBreakdownsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductionStationBreakdownDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 工站故障记录管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductionStationBreakdownDetailResponse
// @Router /api/mom/production/productionstationbreakdown/detail [get]
func GetProductionStationBreakdownDetail(c *gin.Context) {
	resp := &proto.GetProductionStationBreakdownDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductionStationBreakdownByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionStationBreakdownToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductionStationBreakdown godoc
// @Summary 删除
// @Description 删除
// @Tags 工站故障记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductionStationBreakdown"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionstationbreakdown/delete [delete]
func DeleteProductionStationBreakdown(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除工站故障记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductionStationBreakdown(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductionStationBreakdownRouter(r *gin.Engine) {
	g := r.Group("/api/mom/production/productionstationbreakdown")

	g.POST("add", AddProductionStationBreakdown)
	g.PUT("update", UpdateProductionStationBreakdown)
	g.GET("query", QueryProductionStationBreakdown)
	g.DELETE("delete", DeleteProductionStationBreakdown)
	g.GET("all", GetAllProductionStationBreakdown)
	g.GET("detail", GetProductionStationBreakdownDetail)
}
