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

// AddMaterialReturnType godoc
// @Summary 新增
// @Description 新增
// @Tags 物料退料类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialReturnTypeInfo true "Add MaterialReturnType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialreturntype/add [post]
func AddMaterialReturnType(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialReturnTypeInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建物料退料类型请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateMaterialReturnType(model.PBToMaterialReturnType(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateMaterialReturnType godoc
// @Summary 更新
// @Description 更新
// @Tags 物料退料类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialReturnTypeInfo true "Update MaterialReturnType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialreturntype/update [put]
func UpdateMaterialReturnType(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialReturnTypeInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新物料退料类型请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateMaterialReturnType(model.PBToMaterialReturnType(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryMaterialReturnType godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 物料退料类型管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Success 200 {object} proto.QueryMaterialReturnTypeResponse
// @Router /api/mom/material/materialreturntype/query [get]
func QueryMaterialReturnType(c *gin.Context) {
	req := &proto.QueryMaterialReturnTypeRequest{}
	resp := &proto.QueryMaterialReturnTypeResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryMaterialReturnType(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllMaterialReturnType godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 物料退料类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllMaterialReturnTypeResponse
// @Router /api/mom/material/materialreturntype/all [get]
func GetAllMaterialReturnType(c *gin.Context) {
	resp := &proto.GetAllMaterialReturnTypeResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllMaterialReturnTypes()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.MaterialReturnTypesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetMaterialReturnTypeDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 物料退料类型管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetMaterialReturnTypeDetailResponse
// @Router /api/mom/material/materialreturntype/detail [get]
func GetMaterialReturnTypeDetail(c *gin.Context) {
	resp := &proto.GetMaterialReturnTypeDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetMaterialReturnTypeByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialReturnTypeToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteMaterialReturnType godoc
// @Summary 删除
// @Description 删除
// @Tags 物料退料类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete MaterialReturnType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialreturntype/delete [delete]
func DeleteMaterialReturnType(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除物料退料类型请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteMaterialReturnType(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterMaterialReturnTypeRouter(r *gin.Engine) {
	g := r.Group("/api/mom/material/materialreturntype")

	g.POST("add", AddMaterialReturnType)
	g.PUT("update", UpdateMaterialReturnType)
	g.GET("query", QueryMaterialReturnType)
	g.DELETE("delete", DeleteMaterialReturnType)
	g.GET("all", GetAllMaterialReturnType)
	g.GET("detail", GetMaterialReturnTypeDetail)
}
