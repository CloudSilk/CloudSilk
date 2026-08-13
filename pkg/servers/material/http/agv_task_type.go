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

// AddAGVTaskType godoc
// @Summary 新增
// @Description 新增
// @Tags AGV任务类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.AGVTaskTypeInfo true "Add AGVTaskType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/agvtasktype/add [post]
func AddAGVTaskType(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.AGVTaskTypeInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建AGV任务类型请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateAGVTaskType(model.PBToAGVTaskType(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateAGVTaskType godoc
// @Summary 更新
// @Description 更新
// @Tags AGV任务类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.AGVTaskTypeInfo true "Update AGVTaskType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/agvtasktype/update [put]
func UpdateAGVTaskType(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.AGVTaskTypeInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新AGV任务类型请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateAGVTaskType(model.PBToAGVTaskType(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryAGVTaskType godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags AGV任务类型管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param shelfType query int false "货架类型"
// @Param sceneType query int false "场景类型"
// @Param materialContainerTypeID query string false "容器类型ID"
// @Success 200 {object} proto.QueryAGVTaskTypeResponse
// @Router /api/mom/material/agvtasktype/query [get]
func QueryAGVTaskType(c *gin.Context) {
	req := &proto.QueryAGVTaskTypeRequest{}
	resp := &proto.QueryAGVTaskTypeResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryAGVTaskType(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllAGVTaskType godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags AGV任务类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllAGVTaskTypeResponse
// @Router /api/mom/material/agvtasktype/all [get]
func GetAllAGVTaskType(c *gin.Context) {
	resp := &proto.GetAllAGVTaskTypeResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllAGVTaskTypes()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.AGVTaskTypesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetAGVTaskTypeDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags AGV任务类型管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAGVTaskTypeDetailResponse
// @Router /api/mom/material/agvtasktype/detail [get]
func GetAGVTaskTypeDetail(c *gin.Context) {
	resp := &proto.GetAGVTaskTypeDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetAGVTaskTypeByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.AGVTaskTypeToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteAGVTaskType godoc
// @Summary 删除
// @Description 删除
// @Tags AGV任务类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete AGVTaskType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/agvtasktype/delete [delete]
func DeleteAGVTaskType(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除AGV任务类型请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteAGVTaskType(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterAGVTaskTypeRouter(r *gin.Engine) {
	g := r.Group("/api/mom/material/agvtasktype")

	g.POST("add", AddAGVTaskType)
	g.PUT("update", UpdateAGVTaskType)
	g.GET("query", QueryAGVTaskType)
	g.DELETE("delete", DeleteAGVTaskType)
	g.GET("all", GetAllAGVTaskType)
	g.GET("detail", GetAGVTaskTypeDetail)
}
