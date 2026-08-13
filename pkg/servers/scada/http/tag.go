package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/scada/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddScadaTag godoc
// @Summary 新增
// @Description 新增
// @Tags 采集点位
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ScadaTagInfo true "Add ScadaTag"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/scada/tag/add [post]
func AddScadaTag(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ScadaTagInfo{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新增采集点位请求参数无效:%v", transID, err)
		return
	}
	id, err := logic.CreateScadaTag(model.PBToScadaTag(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateScadaTag godoc
// @Summary 更新
// @Description 更新
// @Tags 采集点位
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ScadaTagInfo true "Update ScadaTag"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/scada/tag/update [put]
func UpdateScadaTag(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ScadaTagInfo{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新采集点位请求参数无效:%v", transID, err)
		return
	}
	err = logic.UpdateScadaTag(model.PBToScadaTag(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryScadaTag godoc
// @Summary 查询
// @Description 查询
// @Tags 采集点位
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.QueryScadaTagResponse
// @Router /api/mom/scada/tag/query [get]
func QueryScadaTag(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.QueryScadaTagRequest{}
	resp := &proto.QueryScadaTagResponse{}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,查询采集点位请求参数无效:%v", transID, err)
		return
	}
	logic.QueryScadaTag(req, resp, true)
	c.JSON(http.StatusOK, resp)
}

// GetAllScadaTag godoc
// @Summary 获取所有
// @Description 获取所有
// @Tags 采集点位
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllScadaTagResponse
// @Router /api/mom/scada/tag/all [get]
func GetAllScadaTag(c *gin.Context) {
	resp := &proto.GetAllScadaTagResponse{}
	list, err := logic.GetAllScadaTags()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ScadaTagsToPB(list)
		resp.Code = proto.Code_Success
	}
	c.JSON(http.StatusOK, resp)
}

// GetScadaTagDetail godoc
// @Summary 获取详情
// @Description 获取详情
// @Tags 采集点位
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param id query string true "ID"
// @Success 200 {object} proto.GetScadaTagDetailResponse
// @Router /api/mom/scada/tag/detail [get]
func GetScadaTagDetail(c *gin.Context) {
	resp := &proto.GetScadaTagDetailResponse{}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		resp.Message = "id不能为空"
		c.JSON(http.StatusOK, resp)
		return
	}
	m, err := logic.GetScadaTagByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ScadaTagToPB(m)
		resp.Code = proto.Code_Success
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteScadaTag godoc
// @Summary 删除
// @Description 删除
// @Tags 采集点位
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.GetByIDsRequest true "Delete ScadaTag"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/scada/tag/delete [delete]
func DeleteScadaTag(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.GetByIDsRequest{}
	resp := &proto.CommonResponse{Code: proto.Code_Success}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,删除采集点位请求参数无效:%v", transID, err)
		return
	}
	if len(req.Ids) == 0 || req.Ids[0] == "" {
		resp.Code = proto.Code_BadRequest
		resp.Message = "id不能为空"
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteScadaTag(req.Ids[0])
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}
