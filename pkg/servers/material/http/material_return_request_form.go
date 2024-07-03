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

// AddMaterialReturnRequestForm godoc
// @Summary 新增
// @Description 新增
// @Tags 材料退货申请表管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialReturnRequestFormInfo true "Add MaterialReturnRequestForm"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialreturnrequestform/add [post]
func AddMaterialReturnRequestForm(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialReturnRequestFormInfo{CreateUserID: middleware.GetUserID(c)}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建材料退货申请表请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateMaterialReturnRequestForm(model.PBToMaterialReturnRequestForm(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateMaterialReturnRequestForm godoc
// @Summary 更新
// @Description 更新
// @Tags 材料退货申请表管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialReturnRequestFormInfo true "Update MaterialReturnRequestForm"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialreturnrequestform/update [put]
func UpdateMaterialReturnRequestForm(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialReturnRequestFormInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新材料退货申请表请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateMaterialReturnRequestForm(model.PBToMaterialReturnRequestForm(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryMaterialReturnRequestForm godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 材料退货申请表管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Success 200 {object} proto.QueryMaterialReturnRequestFormResponse
// @Router /api/mom/material/materialreturnrequestform/query [get]
func QueryMaterialReturnRequestForm(c *gin.Context) {
	req := &proto.QueryMaterialReturnRequestFormRequest{}
	resp := &proto.QueryMaterialReturnRequestFormResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryMaterialReturnRequestForm(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllMaterialReturnRequestForm godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 材料退货申请表管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllMaterialReturnRequestFormResponse
// @Router /api/mom/material/materialreturnrequestform/all [get]
func GetAllMaterialReturnRequestForm(c *gin.Context) {
	resp := &proto.GetAllMaterialReturnRequestFormResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllMaterialReturnRequestForms()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.MaterialReturnRequestFormsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetMaterialReturnRequestFormDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 材料退货申请表管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetMaterialReturnRequestFormDetailResponse
// @Router /api/mom/material/materialreturnrequestform/detail [get]
func GetMaterialReturnRequestFormDetail(c *gin.Context) {
	resp := &proto.GetMaterialReturnRequestFormDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetMaterialReturnRequestFormByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialReturnRequestFormToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteMaterialReturnRequestForm godoc
// @Summary 删除
// @Description 删除
// @Tags 材料退货申请表管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete MaterialReturnRequestForm"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialreturnrequestform/delete [delete]
func DeleteMaterialReturnRequestForm(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除材料退货申请表请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteMaterialReturnRequestForm(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterMaterialReturnRequestFormRouter(r *gin.Engine) {
	g := r.Group("/api/mom/material/materialreturnrequestform")

	g.POST("add", AddMaterialReturnRequestForm)
	g.PUT("update", UpdateMaterialReturnRequestForm)
	g.GET("query", QueryMaterialReturnRequestForm)
	g.DELETE("delete", DeleteMaterialReturnRequestForm)
	g.GET("all", GetAllMaterialReturnRequestForm)
	g.GET("detail", GetMaterialReturnRequestFormDetail)
}
