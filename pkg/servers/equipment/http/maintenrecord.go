package http

import (
	"context"
	"net/http"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/servers/equipment/logic"
	"github.com/CloudSilk/pkg/utils/log"
	"github.com/CloudSilk/usercenter/utils/middleware"
	"github.com/gin-gonic/gin"
)



// QueryEquipmentMaintenanceRecord godoc
// @Summary 查询
// @Description 查询
// @Tags 设备维保记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Success 200 {object} proto.QueryEquipmentMaintenanceRecordResponse
// @Router /api/mom/equipment/maintenrecord/query [get]
func QueryEquipmentMaintenanceRecord(c *gin.Context) {
	transID := middleware.GetTransID(c)
	req := &proto.QueryEquipmentMaintenanceRecordRequest{}
	resp := &proto.QueryEquipmentMaintenanceRecordResponse{}
	err := c.BindQuery(req)
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		c.JSON(http.StatusOK, resp)
		log.Warnf(context.Background(), "TransID:%s,查询设备维保记录请求参数无效:%v", transID, err)
		return
	}
	logic.QueryEquipmentMaintenanceRecord(req, resp, true)
	c.JSON(http.StatusOK, resp)
}


// GetEquipmentMaintenanceRecordDetail godoc
// @Summary 获取详情
// @Description 获取详情
// @Tags 设备维保记录管理
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param id query string true "ID"
// @Success 200 {object} proto.GetEquipmentMaintenanceRecordDetailResponse
// @Router /api/mom/equipment/maintenrecord/detail [get]
func GetEquipmentMaintenanceRecordDetail(c *gin.Context) {
	resp := &proto.GetEquipmentMaintenanceRecordDetailResponse{}
	id := c.Query("id")
	if id == "" {
		resp.Code = proto.Code_BadRequest
		resp.Message = "id不能为空"
		c.JSON(http.StatusOK, resp)
		return
	}
	m, err := logic.GetEquipmentMaintenanceRecordByID(id)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.EquipmentMaintenanceRecordToPB(m)
		resp.Code = proto.Code_Success
	}
	c.JSON(http.StatusOK, resp)
}

