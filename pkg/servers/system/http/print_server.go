package http

import (
	"context"
	"net/http"

	model "github.com/CloudSilk/CloudSilk/pkg/model"
	apipb "github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/system/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/pkg/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddPrintServer godoc
// @Summary 新增
// @Description 新增
// @Tags 打印服务器管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.PrintServerInfo true "Add PrintServer"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/mom/printserver/add [post]
func AddPrintServer(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.PrintServerInfo{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建打印服务器管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreatePrintServer(model.PBToPrintServer(req))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdatePrintServer godoc
// @Summary 更新
// @Description 更新
// @Tags 打印服务器管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body apipb.PrintServerInfo true "Update PrintServer"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/mom/printserver/update [put]
func UpdatePrintServer(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &apipb.PrintServerInfo{}
	resp := &apipb.CommonResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新打印服务器管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdatePrintServer(model.PBToPrintServer(req))
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryPrintServer godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 打印服务器管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Success 200 {object} apipb.QueryPrintServerResponse
// @Router /api/mom/printserver/query [get]
func QueryPrintServer(c *gin.Context) {
	req := &apipb.QueryPrintServerRequest{}
	resp := &apipb.QueryPrintServerResponse{
		Code: apipb.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryPrintServer(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllPrintServer godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 打印服务器管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetAllPrintServerResponse
// @Router /api/mom/printserver/all [get]
func GetAllPrintServer(c *gin.Context) {
	resp := &apipb.GetAllPrintServerResponse{
		Code: apipb.Code_Success,
	}
	list, err := logic.GetAllPrintServers()
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.PrintServersToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetAllPrinter godoc
// @Summary 查询所有打印机
// @Description 查询所有打印机
// @Tags 打印服务器管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetAllPrinterResponse
// @Router /api/mom/printserver/all/printer [get]
func GetAllPrinter(c *gin.Context) {
	resp := &apipb.GetAllPrinterResponse{
		Code: apipb.Code_Success,
	}
	list, err := logic.GetAllPrinters()
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.PrintersToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetPrintServerDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 打印服务器管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} apipb.GetPrintServerDetailResponse
// @Router /api/mom/printserver/detail [get]
func GetPrintServerDetail(c *gin.Context) {
	resp := &apipb.GetPrintServerDetailResponse{
		Code: apipb.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = apipb.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetPrintServerByID(id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.PrintServerToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeletePrintServer godoc
// @Summary 删除
// @Description 删除
// @Tags 打印服务器管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body apipb.DelRequest true "Delete PrintServer"
// @Success 200 {object} apipb.CommonResponse
// @Router /api/mom/printserver/delete [delete]
func DeletePrintServer(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除打印服务器管理请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeletePrintServer(req.Id)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterPrintServerRouter(r *gin.Engine) {
	g := r.Group("/api/mom/printserver")

	g.POST("add", AddPrintServer)
	g.PUT("update", UpdatePrintServer)
	g.GET("query", QueryPrintServer)
	g.DELETE("delete", DeletePrintServer)
	g.GET("all", GetAllPrintServer)
	g.GET("all/printer", GetAllPrinter)
	g.GET("detail", GetPrintServerDetail)
}
