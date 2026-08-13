package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/system/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddSerialNumber godoc
// @Summary 新增
// @Description 新增
// @Tags 序列号管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.SerialNumberInfo true "Add SerialNumber"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/serialnumber/add [post]
func AddSerialNumber(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.SerialNumberInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建序列号管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateSerialNumber(model.PBToSerialNumber(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateSerialNumber godoc
// @Summary 更新
// @Description 更新
// @Tags 序列号管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.SerialNumberInfo true "Update SerialNumber"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/serialnumber/update [put]
func UpdateSerialNumber(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.SerialNumberInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新序列号管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateSerialNumber(model.PBToSerialNumber(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QuerySerialNumber godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 序列号管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param name query string false "名字或描述或前缀"
// @Success 200 {object} proto.QuerySerialNumberResponse
// @Router /api/mom/system/serialnumber/query [get]
func QuerySerialNumber(c *gin.Context) {
	req := &proto.QuerySerialNumberRequest{}
	resp := &proto.QuerySerialNumberResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QuerySerialNumber(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllSerialNumber godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 序列号管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllSerialNumberResponse
// @Router /api/mom/system/serialnumber/all [get]
func GetAllSerialNumber(c *gin.Context) {
	resp := &proto.GetAllSerialNumberResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllSerialNumbers()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.SerialNumbersToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetSerialNumberDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 序列号管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetSerialNumberDetailResponse
// @Router /api/mom/system/serialnumber/detail [get]
func GetSerialNumberDetail(c *gin.Context) {
	resp := &proto.GetSerialNumberDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetSerialNumberByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.SerialNumberToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteSerialNumber godoc
// @Summary 删除
// @Description 删除
// @Tags 序列号管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete SerialNumber"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/system/serialnumber/delete [delete]
func DeleteSerialNumber(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除序列号管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteSerialNumber(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterSerialNumberRouter(r *gin.Engine) {
	g := r.Group("/api/mom/system/serialnumber")

	g.POST("add", AddSerialNumber)
	g.PUT("update", UpdateSerialNumber)
	g.GET("query", QuerySerialNumber)
	g.DELETE("delete", DeleteSerialNumber)
	g.GET("all", GetAllSerialNumber)
	g.GET("detail", GetSerialNumberDetail)
}
