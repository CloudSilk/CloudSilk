package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/clients"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/production/logic"
	"github.com/CloudSilk/pkg/utils/log"
	usercenter "github.com/CloudSilk/usercenter/proto"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddPersonnelQualification godoc
// @Summary 新增
// @Description 新增
// @Tags 人员资质管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.PersonnelQualificationInfo true "Add PersonnelQualification"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/personnelqualification/add [post]
func AddPersonnelQualification(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.PersonnelQualificationInfo{
		AuthorizedUserID: middleware.GetUserID(c),
	}
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

	id, err := logic.CreatePersonnelQualification(model.PBToPersonnelQualification(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdatePersonnelQualification godoc
// @Summary 更新
// @Description 更新
// @Tags 人员资质管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.PersonnelQualificationInfo true "Update PersonnelQualification"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/personnelqualification/update [put]
func UpdatePersonnelQualification(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.PersonnelQualificationInfo{}
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
	err = logic.UpdatePersonnelQualification(model.PBToPersonnelQualification(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryPersonnelQualification godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 人员资质管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productionLineID query string false "生产产线ID"
// @Param name query string false "认证人员信息"
// @Success 200 {object} proto.QueryPersonnelQualificationResponse
// @Router /api/mom/production/personnelqualification/query [get]
func QueryPersonnelQualification(c *gin.Context) {
	req := &proto.QueryPersonnelQualificationRequest{}
	resp := &proto.QueryPersonnelQualificationResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryPersonnelQualification(req, resp, false)
	if resp.Code == proto.Code_Success {
		r := &usercenter.QueryUserRequest{}
		for _, u := range resp.Data {
			r.Ids = append(r.Ids, u.CertifiedUserID, u.AuthorizedUserID)
		}
		r.PageSize = int64(len(r.Ids))
		users, err := clients.UserClient.Query(context.Background(), r)
		if err == nil && users.Code == usercenter.Code_Success {
			for _, u := range resp.Data {
				for _, u2 := range users.Data {
					if u.CertifiedUserID == u2.Id {
						u.CertifiedUserName = u2.Nickname
					}
					if u.AuthorizedUserID == u2.Id {
						u.AuthorizedUserName = u2.Nickname
					}
				}
			}
		}
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllPersonnelQualification godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 人员资质管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllPersonnelQualificationResponse
// @Router /api/mom/production/personnelqualification/all [get]
func GetAllPersonnelQualification(c *gin.Context) {
	resp := &proto.GetAllPersonnelQualificationResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllPersonnelQualifications()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.PersonnelQualificationsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetPersonnelQualificationDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 人员资质管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetPersonnelQualificationDetailResponse
// @Router /api/mom/production/personnelqualification/detail [get]
func GetPersonnelQualificationDetail(c *gin.Context) {
	resp := &proto.GetPersonnelQualificationDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetPersonnelQualificationByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.PersonnelQualificationToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeletePersonnelQualification godoc
// @Summary 删除
// @Description 删除
// @Tags 人员资质管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete PersonnelQualification"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/personnelqualification/delete [delete]
func DeletePersonnelQualification(c *gin.Context) {
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
	err = logic.DeletePersonnelQualification(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterPersonnelQualificationRouter(r *gin.Engine) {
	g := r.Group("/api/mom/production/personnelqualification")

	g.POST("add", AddPersonnelQualification)
	g.PUT("update", UpdatePersonnelQualification)
	g.GET("query", QueryPersonnelQualification)
	g.DELETE("delete", DeletePersonnelQualification)
	g.GET("all", GetAllPersonnelQualification)
	g.GET("detail", GetPersonnelQualificationDetail)
}
