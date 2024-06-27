package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/production/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/pkg/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddPersonnelQualificationType godoc
// @Summary 新增
// @Description 新增
// @Tags 人员资质类型
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.PersonnelQualificationTypeInfo true "Add PersonnelQualificationType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/personnelqualificationtype/add [post]
func AddPersonnelQualificationType(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.PersonnelQualificationTypeInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建人员资质请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreatePersonnelQualificationType(model.PBToPersonnelQualificationType(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdatePersonnelQualificationType godoc
// @Summary 更新
// @Description 更新
// @Tags 人员资质类型
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.PersonnelQualificationTypeInfo true "Update PersonnelQualificationType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/personnelqualificationtype/update [put]
func UpdatePersonnelQualificationType(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.PersonnelQualificationTypeInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新人员资质请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdatePersonnelQualificationType(model.PBToPersonnelQualificationType(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryPersonnelQualificationType godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 人员资质类型
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productionLineID query string false "生产产线ID"
// @Param name query string false "认证人员信息"
// @Success 200 {object} proto.QueryPersonnelQualificationTypeResponse
// @Router /api/mom/production/personnelqualificationtype/query [get]
func QueryPersonnelQualificationType(c *gin.Context) {
	req := &proto.QueryPersonnelQualificationTypeRequest{}
	resp := &proto.QueryPersonnelQualificationTypeResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryPersonnelQualificationType(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllPersonnelQualificationType godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 人员资质类型
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllPersonnelQualificationTypeResponse
// @Router /api/mom/production/personnelqualificationtype/all [get]
func GetAllPersonnelQualificationType(c *gin.Context) {
	resp := &proto.GetAllPersonnelQualificationTypeResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllPersonnelQualificationTypes()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.PersonnelQualificationTypesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetPersonnelQualificationTypeDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 人员资质类型
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetPersonnelQualificationTypeDetailResponse
// @Router /api/mom/production/personnelqualificationtype/detail [get]
func GetPersonnelQualificationTypeDetail(c *gin.Context) {
	resp := &proto.GetPersonnelQualificationTypeDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetPersonnelQualificationTypeByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.PersonnelQualificationTypeToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeletePersonnelQualificationType godoc
// @Summary 删除
// @Description 删除
// @Tags 人员资质类型
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete PersonnelQualificationType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/personnelqualificationtype/delete [delete]
func DeletePersonnelQualificationType(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除人员资质请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeletePersonnelQualificationType(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterPersonnelQualificationTypeRouter(r *gin.Engine) {
	g := r.Group("/api/mom/production/personnelqualificationtype")

	g.POST("add", AddPersonnelQualificationType)
	g.PUT("update", UpdatePersonnelQualificationType)
	g.GET("query", QueryPersonnelQualificationType)
	g.DELETE("delete", DeletePersonnelQualificationType)
	g.GET("all", GetAllPersonnelQualificationType)
	g.GET("detail", GetPersonnelQualificationTypeDetail)
}
