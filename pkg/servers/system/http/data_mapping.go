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

// AddDataMapping godoc
// @Summary 新增
// @Description 新增
// @Tags 数据词条管理管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.DataMappingInfo true "Add DataMapping"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/datamapping/add [post]
func AddDataMapping(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.DataMappingInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建数据词条管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateDataMapping(model.PBToDataMapping(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateDataMapping godoc
// @Summary 更新
// @Description 更新
// @Tags 数据词条管理管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.DataMappingInfo true "Update DataMapping"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/datamapping/update [put]
func UpdateDataMapping(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.DataMappingInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新数据词条管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateDataMapping(model.PBToDataMapping(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryDataMapping godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 数据词条管理管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param group query string false "分组"
// @Param code query string false "代号或描述"
// @Success 200 {object} proto.QueryDataMappingResponse
// @Router /api/mom/system/datamapping/query [get]
func QueryDataMapping(c *gin.Context) {
	req := &proto.QueryDataMappingRequest{}
	resp := &proto.QueryDataMappingResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryDataMapping(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllDataMapping godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 数据词条管理管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllDataMappingResponse
// @Router /api/mom/system/datamapping/all [get]
func GetAllDataMapping(c *gin.Context) {
	resp := &proto.GetAllDataMappingResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllDataMappings()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.DataMappingsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetDataMappingDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 数据词条管理管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetDataMappingDetailResponse
// @Router /api/mom/system/datamapping/detail [get]
func GetDataMappingDetail(c *gin.Context) {
	resp := &proto.GetDataMappingDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetDataMappingByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.DataMappingToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteDataMapping godoc
// @Summary 删除
// @Description 删除
// @Tags 数据词条管理管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete DataMapping"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/datamapping/delete [delete]
func DeleteDataMapping(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除数据词条管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteDataMapping(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterDataMappingRouter(r *gin.Engine) {
	g := r.Group("/api/mom/system/datamapping")

	g.POST("add", AddDataMapping)
	g.PUT("update", UpdateDataMapping)
	g.GET("query", QueryDataMapping)
	g.DELETE("delete", DeleteDataMapping)
	g.GET("all", GetAllDataMapping)
	g.GET("detail", GetDataMappingDetail)
}
