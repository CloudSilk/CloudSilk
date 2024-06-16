package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/material/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/pkg/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddMaterialTrayBindingRecord godoc
// @Summary 新增
// @Description 新增
// @Tags 物料托盘绑定记录
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialTrayBindingRecordInfo true "Add MaterialTrayBindingRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialtraybindingrecord/add [post]
func AddMaterialTrayBindingRecord(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialTrayBindingRecordInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建物料托盘绑定记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateMaterialTrayBindingRecord(model.PBToMaterialTrayBindingRecord(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateMaterialTrayBindingRecord godoc
// @Summary 更新
// @Description 更新
// @Tags 物料托盘绑定记录
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialTrayBindingRecordInfo true "Update MaterialTrayBindingRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialtraybindingrecord/update [put]
func UpdateMaterialTrayBindingRecord(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialTrayBindingRecordInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新物料托盘绑定记录请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateMaterialTrayBindingRecord(model.PBToMaterialTrayBindingRecord(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryMaterialTrayBindingRecord godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 物料托盘绑定记录
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productionLineID query string false "产线ID"
// @Param createTime0 query string false "发料时间开始"
// @Param createTime1 query string false "发料时间结束"
// @Success 200 {object} proto.QueryMaterialTrayBindingRecordResponse
// @Router /api/mom/material/materialtraybindingrecord/query [get]
func QueryMaterialTrayBindingRecord(c *gin.Context) {
	req := &proto.QueryMaterialTrayBindingRecordRequest{}
	resp := &proto.QueryMaterialTrayBindingRecordResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryMaterialTrayBindingRecord(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllMaterialTrayBindingRecord godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 物料托盘绑定记录
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllMaterialTrayBindingRecordResponse
// @Router /api/mom/material/materialtraybindingrecord/all [get]
func GetAllMaterialTrayBindingRecord(c *gin.Context) {
	resp := &proto.GetAllMaterialTrayBindingRecordResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllMaterialTrayBindingRecords()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.MaterialTrayBindingRecordsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetMaterialTrayBindingRecordDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 物料托盘绑定记录
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetMaterialTrayBindingRecordDetailResponse
// @Router /api/mom/material/materialtraybindingrecord/detail [get]
func GetMaterialTrayBindingRecordDetail(c *gin.Context) {
	resp := &proto.GetMaterialTrayBindingRecordDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetMaterialTrayBindingRecordByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialTrayBindingRecordToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteMaterialTrayBindingRecord godoc
// @Summary 删除
// @Description 删除
// @Tags 物料托盘绑定记录
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete MaterialTrayBindingRecord"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialtraybindingrecord/delete [delete]
func DeleteMaterialTrayBindingRecord(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除物料托盘请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteMaterialTrayBindingRecord(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterMaterialTrayBindingRecordRouter(r *gin.Engine) {
	g := r.Group("/api/mom/material/materialtraybindingrecord")

	g.POST("add", AddMaterialTrayBindingRecord)
	g.PUT("update", UpdateMaterialTrayBindingRecord)
	g.GET("query", QueryMaterialTrayBindingRecord)
	g.DELETE("delete", DeleteMaterialTrayBindingRecord)
	g.GET("all", GetAllMaterialTrayBindingRecord)
	g.GET("detail", GetMaterialTrayBindingRecordDetail)
}
