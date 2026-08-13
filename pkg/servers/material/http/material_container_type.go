package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/material/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddMaterialContainerType godoc
// @Summary 新增
// @Description 新增
// @Tags 物料容器类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialContainerTypeInfo true "Add MaterialContainerType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialcontainertype/add [post]
func AddMaterialContainerType(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialContainerTypeInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建物料容器类型请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateMaterialContainerType(model.PBToMaterialContainerType(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateMaterialContainerType godoc
// @Summary 更新
// @Description 更新
// @Tags 物料容器类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialContainerTypeInfo true "Update MaterialContainerType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialcontainertype/update [put]
func UpdateMaterialContainerType(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialContainerTypeInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新物料容器类型请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateMaterialContainerType(model.PBToMaterialContainerType(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryMaterialContainerType godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 物料容器类型管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Success 200 {object} proto.QueryMaterialContainerTypeResponse
// @Router /api/mom/material/materialcontainertype/query [get]
func QueryMaterialContainerType(c *gin.Context) {
	req := &proto.QueryMaterialContainerTypeRequest{}
	resp := &proto.QueryMaterialContainerTypeResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryMaterialContainerType(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllMaterialContainerType godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 物料容器类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllMaterialContainerTypeResponse
// @Router /api/mom/material/materialcontainertype/all [get]
func GetAllMaterialContainerType(c *gin.Context) {
	resp := &proto.GetAllMaterialContainerTypeResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllMaterialContainerTypes()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.MaterialContainerTypesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetMaterialContainerTypeDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 物料容器类型管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetMaterialContainerTypeDetailResponse
// @Router /api/mom/material/materialcontainertype/detail [get]
func GetMaterialContainerTypeDetail(c *gin.Context) {
	resp := &proto.GetMaterialContainerTypeDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetMaterialContainerTypeByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialContainerTypeToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteMaterialContainerType godoc
// @Summary 删除
// @Description 删除
// @Tags 物料容器类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete MaterialContainerType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialcontainertype/delete [delete]
func DeleteMaterialContainerType(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除物料容器类型请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteMaterialContainerType(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterMaterialContainerTypeRouter(r *gin.Engine) {
	g := r.Group("/api/mom/material/materialcontainertype")

	g.POST("add", AddMaterialContainerType)
	g.PUT("update", UpdateMaterialContainerType)
	g.GET("query", QueryMaterialContainerType)
	g.DELETE("delete", DeleteMaterialContainerType)
	g.GET("all", GetAllMaterialContainerType)
	g.GET("detail", GetMaterialContainerTypeDetail)
}
