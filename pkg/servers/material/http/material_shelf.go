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

// AddMaterialShelf godoc
// @Summary 新增
// @Description 新增
// @Tags 物料货架管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialShelfInfo true "Add MaterialShelf"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialshelf/add [post]
func AddMaterialShelf(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialShelfInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建物料货架请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateMaterialShelf(model.PBToMaterialShelf(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateMaterialShelf godoc
// @Summary 更新
// @Description 更新
// @Tags 物料货架管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialShelfInfo true "Update MaterialShelf"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialshelf/update [put]
func UpdateMaterialShelf(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialShelfInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新物料货架请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateMaterialShelf(model.PBToMaterialShelf(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryMaterialShelf godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 物料货架管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Success 200 {object} proto.QueryMaterialShelfResponse
// @Router /api/mom/material/materialshelf/query [get]
func QueryMaterialShelf(c *gin.Context) {
	req := &proto.QueryMaterialShelfRequest{}
	resp := &proto.QueryMaterialShelfResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryMaterialShelf(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllMaterialShelf godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 物料货架管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllMaterialShelfResponse
// @Router /api/mom/material/materialshelf/all [get]
func GetAllMaterialShelf(c *gin.Context) {
	resp := &proto.GetAllMaterialShelfResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllMaterialShelfs()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.MaterialShelfsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetMaterialShelfDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 物料货架管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetMaterialShelfDetailResponse
// @Router /api/mom/material/materialshelf/detail [get]
func GetMaterialShelfDetail(c *gin.Context) {
	resp := &proto.GetMaterialShelfDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetMaterialShelfByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialShelfToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteMaterialShelf godoc
// @Summary 删除
// @Description 删除
// @Tags 物料货架管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete MaterialShelf"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialshelf/delete [delete]
func DeleteMaterialShelf(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除物料货架请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteMaterialShelf(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterMaterialShelfRouter(r *gin.Engine) {
	g := r.Group("/api/mom/material/materialshelf")

	g.POST("add", AddMaterialShelf)
	g.PUT("update", UpdateMaterialShelf)
	g.GET("query", QueryMaterialShelf)
	g.DELETE("delete", DeleteMaterialShelf)
	g.GET("all", GetAllMaterialShelf)
	g.GET("detail", GetMaterialShelfDetail)
}
