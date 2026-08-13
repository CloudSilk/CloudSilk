package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/product/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddProductReworkType godoc
// @Summary 新增
// @Description 新增
// @Tags 产品返工类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductReworkTypeInfo true "Add ProductReworkType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productreworktype/add [post]
func AddProductReworkType(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductReworkTypeInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建产品返工类型请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateProductReworkType(model.PBToProductReworkType(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateProductReworkType godoc
// @Summary 更新
// @Description 更新
// @Tags 产品返工类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.ProductReworkTypeInfo true "Update ProductReworkType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productreworktype/update [put]
func UpdateProductReworkType(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.ProductReworkTypeInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新产品返工类型请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateProductReworkType(model.PBToProductReworkType(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryProductReworkType godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 产品返工类型管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param code query string false "代号或描述"
// @Success 200 {object} proto.QueryProductReworkTypeResponse
// @Router /api/mom/product/productreworktype/query [get]
func QueryProductReworkType(c *gin.Context) {
	req := &proto.QueryProductReworkTypeRequest{}
	resp := &proto.QueryProductReworkTypeResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryProductReworkType(req, resp, false)

	c.JSON(http.StatusOK, resp)
}

// GetAllProductReworkType godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 产品返工类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllProductReworkTypeResponse
// @Router /api/mom/product/productreworktype/all [get]
func GetAllProductReworkType(c *gin.Context) {
	resp := &proto.GetAllProductReworkTypeResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllProductReworkTypes()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.ProductReworkTypesToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetProductReworkTypeDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 产品返工类型管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetProductReworkTypeDetailResponse
// @Router /api/mom/product/productreworktype/detail [get]
func GetProductReworkTypeDetail(c *gin.Context) {
	resp := &proto.GetProductReworkTypeDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetProductReworkTypeByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.ProductReworkTypeToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteProductReworkType godoc
// @Summary 删除
// @Description 删除
// @Tags 产品返工类型管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete ProductReworkType"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/product/productreworktype/delete [delete]
func DeleteProductReworkType(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除产品返工类型请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteProductReworkType(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterProductReworkTypeRouter(r *gin.Engine) {
	g := r.Group("/api/mom/product/productreworktype")

	g.POST("add", AddProductReworkType)
	g.PUT("update", UpdateProductReworkType)
	g.GET("query", QueryProductReworkType)
	g.DELETE("delete", DeleteProductReworkType)
	g.GET("all", GetAllProductReworkType)
	g.GET("detail", GetProductReworkTypeDetail)
}
