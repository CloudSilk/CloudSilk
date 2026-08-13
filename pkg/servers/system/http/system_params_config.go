package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/system/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddSystemParamsConfig godoc
// @Summary 新增
// @Description 新增
// @Tags 系统参数管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.SystemParamsConfigInfo true "Add SystemParamsConfig"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/systemparamsconfig/add [post]
func AddSystemParamsConfig(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.SystemParamsConfigInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建系统参数管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateSystemParamsConfig(model.PBToSystemParamsConfig(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateSystemParamsConfig godoc
// @Summary 更新
// @Description 更新
// @Tags 系统参数管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.SystemParamsConfigInfo true "Update SystemParamsConfig"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/systemparamsconfig/update [put]
func UpdateSystemParamsConfig(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.SystemParamsConfigInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新系统参数管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateSystemParamsConfig(model.PBToSystemParamsConfig(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QuerySystemParamsConfig godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 系统参数管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param key query string false "项"
// @Success 200 {object} proto.QuerySystemParamsConfigResponse
// @Router /api/mom/system/systemparamsconfig/query [get]
func QuerySystemParamsConfig(c *gin.Context) {
	req := &proto.QuerySystemParamsConfigRequest{}
	resp := &proto.QuerySystemParamsConfigResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QuerySystemParamsConfig(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllSystemParamsConfig godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 系统参数管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllSystemParamsConfigResponse
// @Router /api/mom/system/systemparamsconfig/all [get]
func GetAllSystemParamsConfig(c *gin.Context) {
	resp := &proto.GetAllSystemParamsConfigResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllSystemParamsConfigs()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.SystemParamsConfigsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetSystemParamsConfigDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 系统参数管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetSystemParamsConfigDetailResponse
// @Router /api/mom/system/systemparamsconfig/detail [get]
func GetSystemParamsConfigDetail(c *gin.Context) {
	resp := &proto.GetSystemParamsConfigDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetSystemParamsConfigByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.SystemParamsConfigToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteSystemParamsConfig godoc
// @Summary 删除
// @Description 删除
// @Tags 系统参数管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete SystemParamsConfig"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/systemparamsconfig/delete [delete]
func DeleteSystemParamsConfig(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除系统参数管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteSystemParamsConfig(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterSystemParamsConfigRouter(r *gin.Engine) {
	g := r.Group("/api/mom/system/systemparamsconfig")

	g.POST("add", AddSystemParamsConfig)
	g.PUT("update", UpdateSystemParamsConfig)
	g.GET("query", QuerySystemParamsConfig)
	g.DELETE("delete", DeleteSystemParamsConfig)
	g.GET("all", GetAllSystemParamsConfig)
	g.GET("detail", GetSystemParamsConfigDetail)
}
