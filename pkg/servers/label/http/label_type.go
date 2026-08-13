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

// AddLabelType godoc
// @Summary 新增
// @Description 新增
// @Tags 标签类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.LabelTypeInfo true "Add LabelType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/label/labeltype/add [post]
func AddLabelType(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.LabelTypeInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建标签类型管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateLabelType(model.PBToLabelType(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateLabelType godoc
// @Summary 更新
// @Description 更新
// @Tags 标签类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.LabelTypeInfo true "Update LabelType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/label/labeltype/update [put]
func UpdateLabelType(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.LabelTypeInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新标签类型管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateLabelType(model.PBToLabelType(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryLabelType godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 标签类型管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "标签类型或描述"
// @Success 200 {object} proto.QueryLabelTypeResponse
// @Router /api/mom/label/labeltype/query [get]
func QueryLabelType(c *gin.Context) {
	req := &proto.QueryLabelTypeRequest{}
	resp := &proto.QueryLabelTypeResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryLabelType(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllLabelType godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 标签类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllLabelTypeResponse
// @Router /api/mom/label/labeltype/all [get]
func GetAllLabelType(c *gin.Context) {
	resp := &proto.GetAllLabelTypeResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllLabelTypes()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.LabelTypesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetLabelTypeDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 标签类型管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetLabelTypeDetailResponse
// @Router /api/mom/label/labeltype/detail [get]
func GetLabelTypeDetail(c *gin.Context) {
	resp := &proto.GetLabelTypeDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetLabelTypeByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.LabelTypeToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteLabelType godoc
// @Summary 删除
// @Description 删除
// @Tags 标签类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete LabelType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/label/labeltype/delete [delete]
func DeleteLabelType(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除标签类型管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteLabelType(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterLabelTypeRouter(r *gin.Engine) {
	g := r.Group("/api/mom/label/labeltype")

	g.POST("add", AddLabelType)
	g.PUT("update", UpdateLabelType)
	g.GET("query", QueryLabelType)
	g.DELETE("delete", DeleteLabelType)
	g.GET("all", GetAllLabelType)
	g.GET("detail", GetLabelTypeDetail)
}
