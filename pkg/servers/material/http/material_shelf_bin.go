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

// AddMaterialShelfBin godoc
// @Summary 新增
// @Description 新增
// @Tags 物料货架库位管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialShelfBinInfo true "Add MaterialShelfBin"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialshelfbin/add [post]
func AddMaterialShelfBin(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialShelfBinInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建物料货架库位请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateMaterialShelfBin(model.PBToMaterialShelfBin(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateMaterialShelfBin godoc
// @Summary 更新
// @Description 更新
// @Tags 物料货架库位管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialShelfBinInfo true "Update MaterialShelfBin"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialshelfbin/update [put]
func UpdateMaterialShelfBin(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialShelfBinInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新物料货架库位请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateMaterialShelfBin(model.PBToMaterialShelfBin(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryMaterialShelfBin godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 物料货架库位管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "编号"
// @Param currentState query string false "当前状态"
// @Success 200 {object} proto.QueryMaterialShelfBinResponse
// @Router /api/mom/material/materialshelfbin/query [get]
func QueryMaterialShelfBin(c *gin.Context) {
	req := &proto.QueryMaterialShelfBinRequest{}
	resp := &proto.QueryMaterialShelfBinResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryMaterialShelfBin(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllMaterialShelfBin godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 物料货架库位管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllMaterialShelfBinResponse
// @Router /api/mom/material/materialshelfbin/all [get]
func GetAllMaterialShelfBin(c *gin.Context) {
	resp := &proto.GetAllMaterialShelfBinResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllMaterialShelfBins()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.MaterialShelfBinsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetMaterialShelfBinDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 物料货架库位管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetMaterialShelfBinDetailResponse
// @Router /api/mom/material/materialshelfbin/detail [get]
func GetMaterialShelfBinDetail(c *gin.Context) {
	resp := &proto.GetMaterialShelfBinDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetMaterialShelfBinByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialShelfBinToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteMaterialShelfBin godoc
// @Summary 删除
// @Description 删除
// @Tags 物料货架库位管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete MaterialShelfBin"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialshelfbin/delete [delete]
func DeleteMaterialShelfBin(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除物料货架库位请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteMaterialShelfBin(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterMaterialShelfBinRouter(r *gin.Engine) {
	g := r.Group("/api/mom/material/materialshelfbin")

	g.POST("add", AddMaterialShelfBin)
	g.PUT("update", UpdateMaterialShelfBin)
	g.GET("query", QueryMaterialShelfBin)
	g.DELETE("delete", DeleteMaterialShelfBin)
	g.GET("all", GetAllMaterialShelfBin)
	g.GET("detail", GetMaterialShelfBinDetail)
}
