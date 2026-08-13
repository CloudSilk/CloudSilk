package http

import (
	"context"
	"net/http"

	model "github.com/CloudSilk/CloudSilk/pkg/model"
	apipb "github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/system/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddCodingSerial godoc
// @Summary 新增
// @Description 新增
// @Tags 编码序列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.CodingSerialInfo true "Add CodingSerial"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/mom/codingserial/add [post]
func AddCodingSerial(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.CodingSerialInfo{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建编码序列管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateCodingSerial(model.PBToCodingSerial(req))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateCodingSerial godoc
// @Summary 更新
// @Description 更新
// @Tags 编码序列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.CodingSerialInfo true "Update CodingSerial"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/mom/codingserial/update [put]
func UpdateCodingSerial(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.CodingSerialInfo{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新编码序列管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateCodingSerial(model.PBToCodingSerial(req))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryCodingSerial godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 编码序列管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Success 200 {object} apipb.QueryCodingSerialResponse
// @Router /api/mom/codingserial/query [get]
func QueryCodingSerial(c *gin.Context) {
	req := &apipb.QueryCodingSerialRequest{}
	resp := &apipb.QueryCodingSerialResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryCodingSerial(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllCodingSerial godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 编码序列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetAllCodingSerialResponse
// @Router /api/mom/codingserial/all [get]
func GetAllCodingSerial(c *gin.Context) {
	resp := &apipb.GetAllCodingSerialResponse{
		Code: apipb.Code_Success,
	}
	list, err := logic.GetAllCodingSerials()
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.CodingSerialsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetCodingSerialDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 编码序列管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetCodingSerialDetailResponse
// @Router /api/mom/codingserial/detail [get]
func GetCodingSerialDetail(c *gin.Context) {
	resp := &apipb.GetCodingSerialDetailResponse{
		Code: apipb.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = apipb.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetCodingSerialByID(id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.CodingSerialToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteCodingSerial godoc
// @Summary 删除
// @Description 删除
// @Tags 编码序列管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body apipb.DelRequest true "Delete CodingSerial"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/mom/codingserial/delete [delete]
func DeleteCodingSerial(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.DelRequest{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,删除编码序列管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteCodingSerial(req.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterCodingSerialRouter(r *gin.Engine) {
	g := r.Group("/api/mom/codingserial")

	g.POST("add", AddCodingSerial)
	g.PUT("update", UpdateCodingSerial)
	g.GET("query", QueryCodingSerial)
	g.DELETE("delete", DeleteCodingSerial)
	g.GET("all", GetAllCodingSerial)
	g.GET("detail", GetCodingSerialDetail)
}
