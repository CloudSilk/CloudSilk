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

// AddMaterialContainer godoc
// @Summary 新增
// @Description 新增
// @Tags 物料容器管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialContainerInfo true "Add MaterialContainer"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialcontainer/add [post]
func AddMaterialContainer(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialContainerInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建物料容器请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateMaterialContainer(model.PBToMaterialContainer(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateMaterialContainer godoc
// @Summary 更新
// @Description 更新
// @Tags 物料容器管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialContainerInfo true "Update MaterialContainer"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialcontainer/update [put]
func UpdateMaterialContainer(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialContainerInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新物料容器请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateMaterialContainer(model.PBToMaterialContainer(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryMaterialContainer godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 物料容器管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "编号或描述"
// @Param currentState query string false "当前状态"
// @Param materialContainerTypeID query string false "容器类型ID"
// @Success 200 {object} proto.QueryMaterialContainerResponse
// @Router /api/mom/material/materialcontainer/query [get]
func QueryMaterialContainer(c *gin.Context) {
	req := &proto.QueryMaterialContainerRequest{}
	resp := &proto.QueryMaterialContainerResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryMaterialContainer(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllMaterialContainer godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 物料容器管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllMaterialContainerResponse
// @Router /api/mom/material/materialcontainer/all [get]
func GetAllMaterialContainer(c *gin.Context) {
	resp := &proto.GetAllMaterialContainerResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllMaterialContainers()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.MaterialContainersToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetMaterialContainerDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 物料容器管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetMaterialContainerDetailResponse
// @Router /api/mom/material/materialcontainer/detail [get]
func GetMaterialContainerDetail(c *gin.Context) {
	resp := &proto.GetMaterialContainerDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetMaterialContainerByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialContainerToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteMaterialContainer godoc
// @Summary 删除
// @Description 删除
// @Tags 物料容器管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete MaterialContainer"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialcontainer/delete [delete]
func DeleteMaterialContainer(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除物料容器请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteMaterialContainer(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterMaterialContainerRouter(r *gin.Engine) {
	g := r.Group("/api/mom/material/materialcontainer")

	g.POST("add", AddMaterialContainer)
	g.PUT("update", UpdateMaterialContainer)
	g.GET("query", QueryMaterialContainer)
	g.DELETE("delete", DeleteMaterialContainer)
	g.GET("all", GetAllMaterialContainer)
	g.GET("detail", GetMaterialContainerDetail)
}
