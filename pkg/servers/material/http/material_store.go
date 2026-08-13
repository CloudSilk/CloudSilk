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

// AddMaterialStore godoc
// @Summary 新增
// @Description 新增
// @Tags 物料仓库管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialStoreInfo true "Add MaterialStore"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialstore/add [post]
func AddMaterialStore(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialStoreInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建物料仓库请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateMaterialStore(model.PBToMaterialStore(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateMaterialStore godoc
// @Summary 更新
// @Description 更新
// @Tags 物料仓库管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialStoreInfo true "Update MaterialStore"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialstore/update [put]
func UpdateMaterialStore(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialStoreInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新物料仓库请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateMaterialStore(model.PBToMaterialStore(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryMaterialStore godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 物料仓库管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Success 200 {object} proto.QueryMaterialStoreResponse
// @Router /api/mom/material/materialstore/query [get]
func QueryMaterialStore(c *gin.Context) {
	req := &proto.QueryMaterialStoreRequest{}
	resp := &proto.QueryMaterialStoreResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryMaterialStore(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllMaterialStore godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 物料仓库管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllMaterialStoreResponse
// @Router /api/mom/material/materialstore/all [get]
func GetAllMaterialStore(c *gin.Context) {
	resp := &proto.GetAllMaterialStoreResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllMaterialStores()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.MaterialStoresToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetMaterialStoreDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 物料仓库管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetMaterialStoreDetailResponse
// @Router /api/mom/material/materialstore/detail [get]
func GetMaterialStoreDetail(c *gin.Context) {
	resp := &proto.GetMaterialStoreDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetMaterialStoreByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialStoreToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteMaterialStore godoc
// @Summary 删除
// @Description 删除
// @Tags 物料仓库管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete MaterialStore"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialstore/delete [delete]
func DeleteMaterialStore(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除物料仓库请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteMaterialStore(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterMaterialStoreRouter(r *gin.Engine) {
	g := r.Group("/api/mom/material/materialstore")

	g.POST("add", AddMaterialStore)
	g.PUT("update", UpdateMaterialStore)
	g.GET("query", QueryMaterialStore)
	g.DELETE("delete", DeleteMaterialStore)
	g.GET("all", GetAllMaterialStore)
	g.GET("detail", GetMaterialStoreDetail)
}
