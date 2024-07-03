package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/label/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddLabelPrintTask godoc
// @Summary 新增
// @Description 新增
// @Tags 标签打印任务管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.LabelPrintTaskInfo true "Add LabelPrintTask"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/label/labelprinttask/add [post]
func AddLabelPrintTask(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.LabelPrintTaskInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建标签打印任务请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateLabelPrintTask(model.PBToLabelPrintTask(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateLabelPrintTask godoc
// @Summary 更新
// @Description 更新
// @Tags 标签打印任务管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.LabelPrintTaskInfo true "Update LabelPrintTask"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/label/labelprinttask/update [put]
func UpdateLabelPrintTask(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.LabelPrintTaskInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新标签打印任务请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateLabelPrintTask(model.PBToLabelPrintTask(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryLabelPrintTask godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 标签打印任务管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Param productionLineID query string false "生产产线ID"
// @Success 200 {object} proto.QueryLabelPrintTaskResponse
// @Router /api/mom/label/labelprinttask/query [get]
func QueryLabelPrintTask(c *gin.Context) {
	req := &proto.QueryLabelPrintTaskRequest{}
	resp := &proto.QueryLabelPrintTaskResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryLabelPrintTask(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllLabelPrintTask godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 标签打印任务管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllLabelPrintTaskResponse
// @Router /api/mom/label/labelprinttask/all [get]
func GetAllLabelPrintTask(c *gin.Context) {
	resp := &proto.GetAllLabelPrintTaskResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllLabelPrintTasks()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.LabelPrintTasksToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetLabelPrintTaskDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 标签打印任务管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetLabelPrintTaskDetailResponse
// @Router /api/mom/label/labelprinttask/detail [get]
func GetLabelPrintTaskDetail(c *gin.Context) {
	resp := &proto.GetLabelPrintTaskDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetLabelPrintTaskByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.LabelPrintTaskToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteLabelPrintTask godoc
// @Summary 删除
// @Description 删除
// @Tags 标签打印任务管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete LabelPrintTask"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/label/labelprinttask/delete [delete]
func DeleteLabelPrintTask(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除标签打印任务管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteLabelPrintTask(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterLabelPrintTaskRouter(r *gin.Engine) {
	g := r.Group("/api/mom/label/labelprinttask")

	g.POST("add", AddLabelPrintTask)
	g.PUT("update", UpdateLabelPrintTask)
	g.GET("query", QueryLabelPrintTask)
	g.DELETE("delete", DeleteLabelPrintTask)
	g.GET("all", GetAllLabelPrintTask)
	g.GET("detail", GetLabelPrintTaskDetail)
}
