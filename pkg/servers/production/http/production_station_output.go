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

// AddProductionStationOutput godoc
// @Summary 新增
// @Description 新增
// @Tags 工站产量记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionStationOutputInfo true "Add ProductionStationOutput"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionstationoutput/add [post]
func AddProductionStationOutput(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionStationOutputInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建工站产量记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductionStationOutput(model.PBToProductionStationOutput(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductionStationOutput godoc
// @Summary 更新
// @Description 更新
// @Tags 工站产量记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionStationOutputInfo true "Update ProductionStationOutput"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionstationoutput/update [put]
func UpdateProductionStationOutput(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionStationOutputInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新工站产量记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductionStationOutput(model.PBToProductionStationOutput(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductionStationOutput godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 工站产量记录管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param outputTime0 query string false "产出时间开始"
// @Param outputTime1 query string false "产出时间结束"
// @Param productionLineID query string false "生产产线ID"
// @Param productSerialNo query string false "产品序列号"
// @Param productOrderNo query string false "产品订单号"
// @Success 200 {object} proto.QueryProductionStationOutputResponse
// @Router /api/mom/production/productionstationoutput/query [get]
func QueryProductionStationOutput(c *gin.Context) {
	req := &proto.QueryProductionStationOutputRequest{}
	resp := &proto.QueryProductionStationOutputResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductionStationOutput(req, resp, false)
	if resp.Code == proto.Code_Success {
		r := &usercenter.QueryUserRequest{}
		for _, u := range resp.Data {
			r.Ids = append(r.Ids, u.LoginUserID)
		}
		r.PageSize = int64(len(r.Ids))
		users, err := clients.UserClient.Query(context.Background(), r)
		if err == nil && users.Code == usercenter.Code_Success {
			for _, u := range resp.Data {
				for _, u2 := range users.Data {
					if u.LoginUserID == u2.Id {
						u.LoginUserName = u2.Nickname
					}
				}
			}
		}
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllProductionStationOutput godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 工站产量记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductionStationOutputResponse
// @Router /api/mom/production/productionstationoutput/all [get]
func GetAllProductionStationOutput(c *gin.Context) {
	resp := &proto.GetAllProductionStationOutputResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductionStationOutputs()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductionStationOutputsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductionStationOutputDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 工站产量记录管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductionStationOutputDetailResponse
// @Router /api/mom/production/productionstationoutput/detail [get]
func GetProductionStationOutputDetail(c *gin.Context) {
	resp := &proto.GetProductionStationOutputDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductionStationOutputByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionStationOutputToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductionStationOutput godoc
// @Summary 删除
// @Description 删除
// @Tags 工站产量记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductionStationOutput"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionstationoutput/delete [delete]
func DeleteProductionStationOutput(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除工站产量记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductionStationOutput(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductionStationOutputRouter(r *gin.Engine) {
	g := r.Group("/api/mom/production/productionstationoutput")

	g.POST("add", AddProductionStationOutput)
	g.PUT("update", UpdateProductionStationOutput)
	g.GET("query", QueryProductionStationOutput)
	g.DELETE("delete", DeleteProductionStationOutput)
	g.GET("all", GetAllProductionStationOutput)
	g.GET("detail", GetProductionStationOutputDetail)
}
