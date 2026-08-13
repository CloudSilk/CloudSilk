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

// AddMaterialTray godoc
// @Summary 新增
// @Description 新增
// @Tags 物料载具管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialTrayInfo true "Add MaterialTray"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialtray/add [post]
func AddMaterialTray(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialTrayInfo{}
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

	id, err := logic.CreateMaterialTray(model.PBToMaterialTray(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateMaterialTray godoc
// @Summary 更新
// @Description 更新
// @Tags 物料载具管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialTrayInfo true "Update MaterialTray"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialtray/update [put]
func UpdateMaterialTray(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialTrayInfo{}
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
	err = logic.UpdateMaterialTray(model.PBToMaterialTray(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryMaterialTray godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 物料载具管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Success 200 {object} proto.QueryMaterialTrayResponse
// @Router /api/mom/material/materialtray/query [get]
func QueryMaterialTray(c *gin.Context) {
	req := &proto.QueryMaterialTrayRequest{}
	resp := &proto.QueryMaterialTrayResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryMaterialTray(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllMaterialTray godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 物料载具管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllMaterialTrayResponse
// @Router /api/mom/material/materialtray/all [get]
func GetAllMaterialTray(c *gin.Context) {
	resp := &proto.GetAllMaterialTrayResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllMaterialTrays()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.MaterialTraysToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetMaterialTrayDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 物料载具管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetMaterialTrayDetailResponse
// @Router /api/mom/material/materialtray/detail [get]
func GetMaterialTrayDetail(c *gin.Context) {
	resp := &proto.GetMaterialTrayDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetMaterialTrayByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialTrayToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteMaterialTray godoc
// @Summary 删除
// @Description 删除
// @Tags 物料载具管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete MaterialTray"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialtray/delete [delete]
func DeleteMaterialTray(c *gin.Context) {
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
	err = logic.DeleteMaterialTray(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterMaterialTrayRouter(r *gin.Engine) {
	g := r.Group("/api/mom/material/materialtray")

	g.POST("add", AddMaterialTray)
	g.PUT("update", UpdateMaterialTray)
	g.GET("query", QueryMaterialTray)
	g.DELETE("delete", DeleteMaterialTray)
	g.GET("all", GetAllMaterialTray)
	g.GET("detail", GetMaterialTrayDetail)
}
