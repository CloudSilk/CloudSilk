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

// AddQualityInspectionOrder godoc
// @Summary 新增检验单
// @Description 按检验类型下的启用标准生成检验单（含GB/T 2828抽样数量）
// @Tags 检验单管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.QualityInspectionOrderInfo true "Add QualityInspectionOrder"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/quality/qualityinspectionorder/add [post]
func AddQualityInspectionOrder(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.QualityInspectionOrderInfo{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新增检验单请求参数无效:%v", transID, err)
		return
	}
	id, err := logic.CreateQualityInspectionOrder(model.PBToQualityInspectionOrder(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateQualityInspectionOrder godoc
// @Summary 更新
// @Description 更新（仅待检验状态）
// @Tags 检验单管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.QualityInspectionOrderInfo true "Update QualityInspectionOrder"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/quality/qualityinspectionorder/update [put]
func UpdateQualityInspectionOrder(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.QualityInspectionOrderInfo{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新检验单请求参数无效:%v", transID, err)
		return
	}
	err = logic.UpdateQualityInspectionOrder(model.PBToQualityInspectionOrder(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryQualityInspectionOrder godoc
// @Summary 查询
// @Description 查询
// @Tags 检验单管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.QueryQualityInspectionOrderResponse
// @Router /api/mom/quality/qualityinspectionorder/query [get]
func QueryQualityInspectionOrder(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.QueryQualityInspectionOrderRequest{}
	resp := &proto.QueryQualityInspectionOrderResponse{}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,查询检验单请求参数无效:%v", transID, err)
		return
	}
	logic.QueryQualityInspectionOrder(req, resp, true)
	c.JSON(http.StatusOK, resp)
}

// GetQualityInspectionOrderDetail godoc
// @Summary 获取详情
// @Description 获取详情（含明细）
// @Tags 检验单管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param id query string true "ID"
// @Success 200 {object} proto.GetQualityInspectionOrderDetailResponse
// @Router /api/mom/quality/qualityinspectionorder/detail [get]
func GetQualityInspectionOrderDetail(c *gin.Context) {
	resp := &proto.GetQualityInspectionOrderDetailResponse{}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		resp.Message = "id不能为空"
		c.JSON(http.StatusOK, resp)
		return
	}
	m, err := logic.GetQualityInspectionOrderByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.QualityInspectionOrderToPB(m)
		resp.Code = proto.Code_Success
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteQualityInspectionOrder godoc
// @Summary 删除
// @Description 删除（已完成的检验单不可删除）
// @Tags 检验单管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.GetByIDsRequest true "Delete QualityInspectionOrder"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/quality/qualityinspectionorder/delete [delete]
func DeleteQualityInspectionOrder(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.GetByIDsRequest{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,删除检验单请求参数无效:%v", transID, err)
		return
	}
	if len(req.Ids) == 0 || req.Ids[0] == "" {
		resp.Code = proto.Code_BadRequest
		resp.Message = "id不能为空"
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteQualityInspectionOrder(req.Ids[0])
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// CompleteQualityInspectionOrder godoc
// @Summary 完成检验
// @Description 录入检测结果，按标准判定单项、按AQL判定整单，可选生成返工记录
// @Tags 检验单管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.CompleteQualityInspectionOrderRequest true "Complete QualityInspectionOrder"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/quality/qualityinspectionorder/complete [put]
func CompleteQualityInspectionOrder(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.CompleteQualityInspectionOrderRequest{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,完成检验请求参数无效:%v", transID, err)
		return
	}
	if err := middleware.Validate.Struct(req); err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	req.InspectionUserID = middleware.GetUserID(c)
	err = logic.CompleteQualityInspectionOrder(req)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// CalcSpc godoc
// @Summary SPC过程能力计算
// @Description 计算均值/标准差/CPK与休哈特控制图界限（可关联检验标准取规格限）
// @Tags 质量统计
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.SpcCalcRequest true "SpcCalcRequest"
// @Success 200 {object} proto.SpcCalcResponse
// @Router /api/mom/quality/spc/calc [post]
func CalcSpc(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.SpcCalcRequest{}
	resp := &proto.SpcCalcResponse{}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,SPC计算请求参数无效:%v", transID, err)
		return
	}
	result, err := logic.CalcSpcFromStandard(req.QualityInspectionStandardID, req.Values, req.LowerSpec, req.UpperSpec)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp = logic.SpcResultToPB(result)
	c.JSON(http.StatusOK, resp)
}
