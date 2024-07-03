package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/clients"
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/material/logic"
	"github.com/CloudSilk/pkg/utils/log"
	usercenter "github.com/CloudSilk/usercenter/proto"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)

// AddMaterialInventory godoc
// @Summary 新增
// @Description 新增
// @Tags 物料库存管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialInventoryInfo true "Add MaterialInventory"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialinventory/add [post]
func AddMaterialInventory(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialInventoryInfo{CreateUserID: middleware.GetUserID(c)}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,新建物料库存请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}

	id, err := logic.CreateMaterialInventory(model.PBToMaterialInventory(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Message = id
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateMaterialInventory godoc
// @Summary 更新
// @Description 更新
// @Tags 物料库存管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.MaterialInventoryInfo true "Update MaterialInventory"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialinventory/update [put]
func UpdateMaterialInventory(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.MaterialInventoryInfo{}
	resp := &proto.CommonResponse{
		Code: proto.Code_Success,
	}
	err := c.BindJSON(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,更新物料库存请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.UpdateMaterialInventory(model.PBToMaterialInventory(req))
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

// QueryMaterialInventory godoc
// @Summary 分页查询
// @Description 分页查询
// @Tags 物料库存管理
// @Accept  json
// @Produce  octet-stream
// @Param authorization header string true "jwt token"
// @Param pageIndex query int false "从1开始"
// @Param pageSize query int false "默认每页10条"
// @Param orderField query string false "排序字段"
// @Param desc query bool false "是否倒序排序"
// @Param createTime0 query string false "创建时间开始"
// @Param createTime1 query string false "创建时间结束"
// @Param materialInfo query string false "物料代号或描述"
// @Param materialStore query string false "仓库代号或描述"
// @Success 200 {object} proto.QueryMaterialInventoryResponse
// @Router /api/mom/material/materialinventory/query [get]
func QueryMaterialInventory(c *gin.Context) {
	req := &proto.QueryMaterialInventoryRequest{}
	resp := &proto.QueryMaterialInventoryResponse{
		Code: proto.Code_Success,
	}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	logic.QueryMaterialInventory(req, resp, false)
	if resp.Code == proto.Code_Success {
		r := &usercenter.QueryUserRequest{}
		for _, u := range resp.Data {
			r.Ids = append(r.Ids, u.CreateUserID)
		}
		r.PageSize = int64(len(r.Ids))
		users, err := clients.UserClient.Query(context.Background(), r)
		if err == nil && users.Code == usercenter.Code_Success {
			for _, u := range resp.Data {
				for _, u2 := range users.Data {
					if u.CreateUserID == u2.Id {
						u.CreateUserName = u2.Nickname
					}
				}
			}
		}
	}

	c.JSON(http.StatusOK, resp)
}

// GetAllMaterialInventory godoc
// @Summary 查询所有
// @Description 查询所有
// @Tags 物料库存管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetAllMaterialInventoryResponse
// @Router /api/mom/material/materialinventory/all [get]
func GetAllMaterialInventory(c *gin.Context) {
	resp := &proto.GetAllMaterialInventoryResponse{
		Code: proto.Code_Success,
	}
	list, err := logic.GetAllMaterialInventorys()
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.Data = model.MaterialInventorysToPB(list)
	c.JSON(http.StatusOK, resp)
}

// GetMaterialInventoryDetail godoc
// @Summary 查询明细
// @Description 查询明细
// @Tags 物料库存管理
// @Accept  json
// @Produce  json
// @Param id query string true "ID"
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.GetMaterialInventoryDetailResponse
// @Router /api/mom/material/materialinventory/detail [get]
func GetMaterialInventoryDetail(c *gin.Context) {
	resp := &proto.GetMaterialInventoryDetailResponse{
		Code: proto.Code_Success,
	}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		c.JSON(http.StatusOK, resp)
		return
	}
	var err error

	data, err := logic.GetMaterialInventoryByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialInventoryToPB(data)
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteMaterialInventory godoc
// @Summary 删除
// @Description 删除
// @Tags 物料库存管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param data body proto.DelRequest true "Delete MaterialInventory"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/material/materialinventory/delete [delete]
func DeleteMaterialInventory(c *gin.Context) {
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
		log.Warnf(context.Background(), "TransID:%s,删除物料库存请求参数无效:%v", transID, err)
		return
	}
	err = middleware.Validate.Struct(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	err = logic.DeleteMaterialInventory(req.Id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func RegisterMaterialInventoryRouter(r *gin.Engine) {
	g := r.Group("/api/mom/material/materialinventory")

	g.POST("add", AddMaterialInventory)
	g.PUT("update", UpdateMaterialInventory)
	g.GET("query", QueryMaterialInventory)
	g.DELETE("delete", DeleteMaterialInventory)
	g.GET("all", GetAllMaterialInventory)
	g.GET("detail", GetMaterialInventoryDetail)
}
