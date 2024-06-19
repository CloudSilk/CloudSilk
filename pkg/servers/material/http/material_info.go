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

// AddMaterialInfo godoc
// @Summary 新增
// @Description 新增
// @Tags 物料信息管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialInfoInfo true "Add MaterialInfo"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialinfo/add [post]
func AddMaterialInfo(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialInfoInfo{}
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

	id, err := logic.CreateMaterialInfo(model.PBToMaterialInfo(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateMaterialInfo godoc
// @Summary 更新
// @Description 更新
// @Tags 物料信息管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialInfoInfo true "Update MaterialInfo"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialinfo/update [put]
func UpdateMaterialInfo(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialInfoInfo{}
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
	err = logic.UpdateMaterialInfo(model.PBToMaterialInfo(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryMaterialInfo godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 物料信息管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Success 200 {object} proto.QueryMaterialInfoResponse
// @Router /api/mom/material/materialinfo/query [get]
func QueryMaterialInfo(c *gin.Context) {
	req := &proto.QueryMaterialInfoRequest{}
	resp := &proto.QueryMaterialInfoResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryMaterialInfo(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllMaterialInfo godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 物料信息管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllMaterialInfoResponse
// @Router /api/mom/material/materialinfo/all [get]
func GetAllMaterialInfo(c *gin.Context) {
	resp := &proto.GetAllMaterialInfoResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllMaterialInfos()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.MaterialInfosToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetMaterialInfoDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 物料信息管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetMaterialInfoDetailResponse
// @Router /api/mom/material/materialinfo/detail [get]
func GetMaterialInfoDetail(c *gin.Context) {
	resp := &proto.GetMaterialInfoDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetMaterialInfoByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialInfoToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteMaterialInfo godoc
// @Summary 删除
// @Description 删除
// @Tags 物料信息管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete MaterialInfo"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialinfo/delete [delete]
func DeleteMaterialInfo(c *gin.Context) {
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
	err = logic.DeleteMaterialInfo(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterMaterialInfoRouter(r *gin.Engine) {
	g := r.Group("/api/mom/material/materialinfo")

	g.POST("add", AddMaterialInfo)
	g.PUT("update", UpdateMaterialInfo)
	g.GET("query", QueryMaterialInfo)
	g.DELETE("delete", DeleteMaterialInfo)
	g.GET("all", GetAllMaterialInfo)
	g.GET("detail", GetMaterialInfoDetail)
}
