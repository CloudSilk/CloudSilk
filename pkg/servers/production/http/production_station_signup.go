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

// AddProductionStationSignup godoc
// @Summary 新增
// @Description 新增
// @Tags 工站登录记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionStationSignupInfo true "Add ProductionStationSignup"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionstationsignup/add [post]
func AddProductionStationSignup(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionStationSignupInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建工站登录记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductionStationSignup(model.PBToProductionStationSignup(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductionStationSignup godoc
// @Summary 更新
// @Description 更新
// @Tags 工站登录记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionStationSignupInfo true "Update ProductionStationSignup"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionstationsignup/update [put]
func UpdateProductionStationSignup(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionStationSignupInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新工站登录记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductionStationSignup(model.PBToProductionStationSignup(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductionStationSignup godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 工站登录记录管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param loginTime0 query string false "登入时间开始"
// @Param loginTime1 query string false "登入时间结束"
// @Param productionLineID query string false "生产产线ID"
// @Success 200 {object} proto.QueryProductionStationSignupResponse
// @Router /api/mom/production/productionstationsignup/query [get]
func QueryProductionStationSignup(c *gin.Context) {
	req := &proto.QueryProductionStationSignupRequest{}
	resp := &proto.QueryProductionStationSignupResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductionStationSignup(req, resp, false)
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

// GetAllProductionStationSignup godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 工站登录记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductionStationSignupResponse
// @Router /api/mom/production/productionstationsignup/all [get]
func GetAllProductionStationSignup(c *gin.Context) {
	resp := &proto.GetAllProductionStationSignupResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductionStationSignups()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductionStationSignupsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductionStationSignupDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 工站登录记录管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductionStationSignupDetailResponse
// @Router /api/mom/production/productionstationsignup/detail [get]
func GetProductionStationSignupDetail(c *gin.Context) {
	resp := &proto.GetProductionStationSignupDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductionStationSignupByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionStationSignupToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductionStationSignup godoc
// @Summary 删除
// @Description 删除
// @Tags 工站登录记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductionStationSignup"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionstationsignup/delete [delete]
func DeleteProductionStationSignup(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除工站登录记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductionStationSignup(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductionStationSignupRouter(r *gin.Engine) {
	g := r.Group("/api/mom/production/productionstationsignup")

	g.POST("add", AddProductionStationSignup)
	g.PUT("update", UpdateProductionStationSignup)
	g.GET("query", QueryProductionStationSignup)
	g.DELETE("delete", DeleteProductionStationSignup)
	g.GET("all", GetAllProductionStationSignup)
	g.GET("detail", GetProductionStationSignupDetail)
}
