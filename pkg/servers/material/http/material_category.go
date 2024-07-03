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

// AddMaterialCategory godoc
// @Summary 新增
// @Description 新增
// @Tags 物料类别管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialCategoryInfo true "Add MaterialCategory"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialcategory/add [post]
func AddMaterialCategory(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialCategoryInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建物料类别请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateMaterialCategory(model.PBToMaterialCategory(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateMaterialCategory godoc
// @Summary 更新
// @Description 更新
// @Tags 物料类别管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialCategoryInfo true "Update MaterialCategory"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialcategory/update [put]
func UpdateMaterialCategory(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialCategoryInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新物料类别请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateMaterialCategory(model.PBToMaterialCategory(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryMaterialCategory godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 物料类别管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Success 200 {object} proto.QueryMaterialCategoryResponse
// @Router /api/mom/material/materialcategory/query [get]
func QueryMaterialCategory(c *gin.Context) {
	req := &proto.QueryMaterialCategoryRequest{}
	resp := &proto.QueryMaterialCategoryResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryMaterialCategory(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllMaterialCategory godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 物料类别管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllMaterialCategoryResponse
// @Router /api/mom/material/materialcategory/all [get]
func GetAllMaterialCategory(c *gin.Context) {
	resp := &proto.GetAllMaterialCategoryResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllMaterialCategorys()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.MaterialCategorysToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetMaterialCategoryDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 物料类别管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetMaterialCategoryDetailResponse
// @Router /api/mom/material/materialcategory/detail [get]
func GetMaterialCategoryDetail(c *gin.Context) {
	resp := &proto.GetMaterialCategoryDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetMaterialCategoryByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialCategoryToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteMaterialCategory godoc
// @Summary 删除
// @Description 删除
// @Tags 物料类别管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete MaterialCategory"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialcategory/delete [delete]
func DeleteMaterialCategory(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除物料类别请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteMaterialCategory(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterMaterialCategoryRouter(r *gin.Engine) {
	g := r.Group("/api/mom/material/materialcategory")

	g.POST("add", AddMaterialCategory)
	g.PUT("update", UpdateMaterialCategory)
	g.GET("query", QueryMaterialCategory)
	g.DELETE("delete", DeleteMaterialCategory)
	g.GET("all", GetAllMaterialCategory)
	g.GET("detail", GetMaterialCategoryDetail)
}
