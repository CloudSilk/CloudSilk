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

// AddQualityInspectionStandard godoc
// @Summary 新增
// @Description 新增
// @Tags 检验标准管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.QualityInspectionStandardInfo true "Add QualityInspectionStandard"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/quality/qualityinspectionstandard/add [post]
func AddQualityInspectionStandard(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.QualityInspectionStandardInfo{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新增检验标准请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	id, err := logic.CreateQualityInspectionStandard(model.PBToQualityInspectionStandard(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateQualityInspectionStandard godoc
// @Summary 更新
// @Description 更新
// @Tags 检验标准管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.QualityInspectionStandardInfo true "Update QualityInspectionStandard"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/quality/qualityinspectionstandard/update [put]
func UpdateQualityInspectionStandard(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.QualityInspectionStandardInfo{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新检验标准请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateQualityInspectionStandard(model.PBToQualityInspectionStandard(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryQualityInspectionStandard godoc
// @Summary 查询
// @Description 查询
// @Tags 检验标准管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.QueryQualityInspectionStandardResponse
// @Router /api/mom/quality/qualityinspectionstandard/query [get]
func QueryQualityInspectionStandard(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.QueryQualityInspectionStandardRequest{}
	resp := &proto.QueryQualityInspectionStandardResponse{}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,查询检验标准请求参数无效:%v", transID, err)
		return
	}
	logic.QueryQualityInspectionStandard(req, resp, true)
	c.JSON(http.StatusOK, resp)
}

// GetAllQualityInspectionStandard godoc
// @Summary 获取所有
// @Description 获取所有
// @Tags 检验标准管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllQualityInspectionStandardResponse
// @Router /api/mom/quality/qualityinspectionstandard/all [get]
func GetAllQualityInspectionStandard(c *gin.Context) {
	resp := &proto.GetAllQualityInspectionStandardResponse{}
	list, err := logic.GetAllQualityInspectionStandards()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.QualityInspectionStandardsToPB(list)
		resp.Code = proto.Code_Success
	}
	c.JSON(http.StatusOK, resp)
}

// GetQualityInspectionStandardDetail godoc
// @Summary 获取详情
// @Description 获取详情
// @Tags 检验标准管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param id query string true "ID"
// @Success 200 {object} proto.GetQualityInspectionStandardDetailResponse
// @Router /api/mom/quality/qualityinspectionstandard/detail [get]
func GetQualityInspectionStandardDetail(c *gin.Context) {
	resp := &proto.GetQualityInspectionStandardDetailResponse{}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		resp.Message = "id不能为空"
		c.JSON(http.StatusOK, resp)
		return
	}
	m, err := logic.GetQualityInspectionStandardByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.QualityInspectionStandardToPB(m)
		resp.Code = proto.Code_Success
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteQualityInspectionStandard godoc
// @Summary 删除
// @Description 删除
// @Tags 检验标准管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.GetByIDsRequest true "Delete QualityInspectionStandard"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/quality/qualityinspectionstandard/delete [delete]
func DeleteQualityInspectionStandard(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.GetByIDsRequest{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,删除检验标准请求参数无效:%v", transID, err)
		return
	}
	if len(req.Ids) == 0 || req.Ids[0] == "" {
		resp.Code = proto.Code_BadRequest
		resp.Message = "id不能为空"
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteQualityInspectionStandard(req.Ids[0])
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}
