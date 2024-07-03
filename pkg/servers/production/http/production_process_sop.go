package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/production/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddProductionProcessSop godoc
// @Summary 新增
// @Description 新增
// @Tags 作业手册管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionProcessSopInfo true "Add ProductionProcessSop"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionprocesssop/add [post]
func AddProductionProcessSop(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionProcessSopInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建作业手册请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductionProcessSop(model.PBToProductionProcessSop(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductionProcessSop godoc
// @Summary 更新
// @Description 更新
// @Tags 作业手册管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductionProcessSopInfo true "Update ProductionProcessSop"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionprocesssop/update [put]
func UpdateProductionProcessSop(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductionProcessSopInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新作业手册请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductionProcessSop(model.PBToProductionProcessSop(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductionProcessSop godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 作业手册管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param productionLineID query string false "生产产线ID"
// @Success 200 {object} proto.QueryProductionProcessSopResponse
// @Router /api/mom/production/productionprocesssop/query [get]
func QueryProductionProcessSop(c *gin.Context) {
	req := &proto.QueryProductionProcessSopRequest{}
	resp := &proto.QueryProductionProcessSopResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductionProcessSop(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductionProcessSop godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 作业手册管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductionProcessSopResponse
// @Router /api/mom/production/productionprocesssop/all [get]
func GetAllProductionProcessSop(c *gin.Context) {
	resp := &proto.GetAllProductionProcessSopResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductionProcessSops()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductionProcessSopsToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductionProcessSopDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 作业手册管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductionProcessSopDetailResponse
// @Router /api/mom/production/productionprocesssop/detail [get]
func GetProductionProcessSopDetail(c *gin.Context) {
	resp := &proto.GetProductionProcessSopDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductionProcessSopByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductionProcessSopToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductionProcessSop godoc
// @Summary 删除
// @Description 删除
// @Tags 作业手册管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductionProcessSop"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/production/productionprocesssop/delete [delete]
func DeleteProductionProcessSop(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除作业手册请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductionProcessSop(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductionProcessSopRouter(r *gin.Engine) {
	g := r.Group("/api/mom/production/productionprocesssop")

	g.POST("add", AddProductionProcessSop)
	g.PUT("update", UpdateProductionProcessSop)
	g.GET("query", QueryProductionProcessSop)
	g.DELETE("delete", DeleteProductionProcessSop)
	g.GET("all", GetAllProductionProcessSop)
	g.GET("detail", GetProductionProcessSopDetail)
}
