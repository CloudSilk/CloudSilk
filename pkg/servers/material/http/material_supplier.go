package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/material/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/pkg/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddMaterialSupplier godoc
// @Summary 新增
// @Description 新增
// @Tags 物料供应商管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialSupplierInfo true "Add MaterialSupplier"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialsupplier/add [post]
func AddMaterialSupplier(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialSupplierInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建物料载具请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateMaterialSupplier(model.PBToMaterialSupplier(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateMaterialSupplier godoc
// @Summary 更新
// @Description 更新
// @Tags 物料供应商管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialSupplierInfo true "Update MaterialSupplier"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialsupplier/update [put]
func UpdateMaterialSupplier(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialSupplierInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新物料载具请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateMaterialSupplier(model.PBToMaterialSupplier(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryMaterialSupplier godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 物料供应商管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Success 200 {object} proto.QueryMaterialSupplierResponse
// @Router /api/mom/material/materialsupplier/query [get]
func QueryMaterialSupplier(c *gin.Context) {
	req := &proto.QueryMaterialSupplierRequest{}
	resp := &proto.QueryMaterialSupplierResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryMaterialSupplier(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllMaterialSupplier godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 物料供应商管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllMaterialSupplierResponse
// @Router /api/mom/material/materialsupplier/all [get]
func GetAllMaterialSupplier(c *gin.Context) {
	resp := &proto.GetAllMaterialSupplierResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllMaterialSuppliers()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.MaterialSuppliersToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetMaterialSupplierDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 物料供应商管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetMaterialSupplierDetailResponse
// @Router /api/mom/material/materialsupplier/detail [get]
func GetMaterialSupplierDetail(c *gin.Context) {
	resp := &proto.GetMaterialSupplierDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetMaterialSupplierByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialSupplierToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteMaterialSupplier godoc
// @Summary 删除
// @Description 删除
// @Tags 物料供应商管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete MaterialSupplier"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialsupplier/delete [delete]
func DeleteMaterialSupplier(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除物料载具请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteMaterialSupplier(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterMaterialSupplierRouter(r *gin.Engine) {
	g := r.Group("/api/mom/material/materialsupplier")

	g.POST("add", AddMaterialSupplier)
	g.PUT("update", UpdateMaterialSupplier)
	g.GET("query", QueryMaterialSupplier)
	g.DELETE("delete", DeleteMaterialSupplier)
	g.GET("all", GetAllMaterialSupplier)
	g.GET("detail", GetMaterialSupplierDetail)
}
