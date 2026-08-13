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

// AddProductionStationAlarm godoc
// @Summary 新增
// @Description 新增
// @Tags 工站报警记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionStationAlarmInfo true "Add ProductionStationAlarm"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionstationalarm/add [post]
func AddProductionStationAlarm(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionStationAlarmInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建工站报警记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductionStationAlarm(model.PBToProductionStationAlarm(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductionStationAlarm godoc
// @Summary 更新
// @Description 更新
// @Tags 工站报警记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionStationAlarmInfo true "Update ProductionStationAlarm"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionstationalarm/update [put]
func UpdateProductionStationAlarm(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionStationAlarmInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新工站报警记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductionStationAlarm(model.PBToProductionStationAlarm(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductionStationAlarm godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 工站报警记录管理
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
// @Param productOrderNo query int false "生产工单号"
// @Param productSerialNo query int false "产品序列号"
// @Param alarmMessage query int false "报警信息"
// @Success 200 {object} proto.QueryProductionStationAlarmResponse
// @Router /api/mom/production/productionstationalarm/query [get]
func QueryProductionStationAlarm(c *gin.Context) {
	req := &proto.QueryProductionStationAlarmRequest{}
	resp := &proto.QueryProductionStationAlarmResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductionStationAlarm(req, resp, false)
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

// GetAllProductionStationAlarm godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 工站报警记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductionStationAlarmResponse
// @Router /api/mom/production/productionstationalarm/all [get]
func GetAllProductionStationAlarm(c *gin.Context) {
	resp := &proto.GetAllProductionStationAlarmResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductionStationAlarms()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductionStationAlarmsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductionStationAlarmDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 工站报警记录管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductionStationAlarmDetailResponse
// @Router /api/mom/production/productionstationalarm/detail [get]
func GetProductionStationAlarmDetail(c *gin.Context) {
	resp := &proto.GetProductionStationAlarmDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductionStationAlarmByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionStationAlarmToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductionStationAlarm godoc
// @Summary 删除
// @Description 删除
// @Tags 工站报警记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductionStationAlarm"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionstationalarm/delete [delete]
func DeleteProductionStationAlarm(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除工站报警记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductionStationAlarm(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductionStationAlarmRouter(r *gin.Engine) {
	g := r.Group("/api/mom/production/productionstationalarm")

	g.POST("add", AddProductionStationAlarm)
	g.PUT("update", UpdateProductionStationAlarm)
	g.GET("query", QueryProductionStationAlarm)
	g.DELETE("delete", DeleteProductionStationAlarm)
	g.GET("all", GetAllProductionStationAlarm)
	g.GET("detail", GetProductionStationAlarmDetail)
}
