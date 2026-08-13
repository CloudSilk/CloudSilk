package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/quality/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddQualityInspectionType godoc
// @Summary 新增
// @Description 新增
// @Tags 检验类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.QualityInspectionTypeInfo true "Add QualityInspectionType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/quality/qualityinspectiontype/add [post]
func AddQualityInspectionType(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.QualityInspectionTypeInfo{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新增检验类型请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	id, err := logic.CreateQualityInspectionType(model.PBToQualityInspectionType(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateQualityInspectionType godoc
// @Summary 更新
// @Description 更新
// @Tags 检验类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.QualityInspectionTypeInfo true "Update QualityInspectionType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/quality/qualityinspectiontype/update [put]
func UpdateQualityInspectionType(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.QualityInspectionTypeInfo{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新检验类型请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateQualityInspectionType(model.PBToQualityInspectionType(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryQualityInspectionType godoc
// @Summary 查询
// @Description 查询
// @Tags 检验类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.QueryQualityInspectionTypeResponse
// @Router /api/mom/quality/qualityinspectiontype/query [get]
func QueryQualityInspectionType(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.QueryQualityInspectionTypeRequest{}
	resp := &proto.QueryQualityInspectionTypeResponse{}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,查询检验类型请求参数无效:%v", transID, err)
		return
	}
	logic.QueryQualityInspectionType(req, resp, true)
	c.JSON(http.StatusOK, resp)
}

// GetAllQualityInspectionType godoc
// @Summary 获取所有
// @Description 获取所有
// @Tags 检验类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllQualityInspectionTypeResponse
// @Router /api/mom/quality/qualityinspectiontype/all [get]
func GetAllQualityInspectionType(c *gin.Context) {
	resp := &proto.GetAllQualityInspectionTypeResponse{}
	list, err := logic.GetAllQualityInspectionTypes()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.QualityInspectionTypesToPB(list)
		resp.Code = proto.Code_Success
	}
	c.JSON(http.StatusOK, resp)
}

// GetQualityInspectionTypeDetail godoc
// @Summary 获取详情
// @Description 获取详情
// @Tags 检验类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param id query string true "ID"
// @Success 200 {object} proto.GetQualityInspectionTypeDetailResponse
// @Router /api/mom/quality/qualityinspectiontype/detail [get]
func GetQualityInspectionTypeDetail(c *gin.Context) {
	resp := &proto.GetQualityInspectionTypeDetailResponse{}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		resp.Message = "id不能为空"
		c.JSON(http.StatusOK, resp)
		return
	}
	m, err := logic.GetQualityInspectionTypeByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.QualityInspectionTypeToPB(m)
		resp.Code = proto.Code_Success
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteQualityInspectionType godoc
// @Summary 删除
// @Description 删除
// @Tags 检验类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.GetByIDsRequest true "Delete QualityInspectionType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/quality/qualityinspectiontype/delete [delete]
func DeleteQualityInspectionType(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.GetByIDsRequest{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,删除检验类型请求参数无效:%v", transID, err)
		return
	}
	if len(req.Ids) == 0 || req.Ids[0] == "" {
		resp.Code = proto.Code_BadRequest
		resp.Message = "id不能为空"
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteQualityInspectionType(req.Ids[0])
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}
