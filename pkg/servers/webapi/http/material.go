package http

// BindMaterialTray godoc
// @Summary 绑定物料载具
// @Description 绑定物料载具
// @Tags WebAPI
// @Accept  json
// @Produce  json
// @Param authorization header string true "jwt token"
// @Param account body proto.BindMaterialTrayRequest true "BindMaterialTrayRequest"
// @Success 200 {object} proto.CommonResponse
// @Router /api/mom/webapi/material/bindmaterialtray [post]
// func BindMaterialTray(c *gin.Context) {
// 	transID := ucmiddleware.GetTransID(c)
// 	req := &proto.BindMaterialTrayRequest{}
// 	resp := &proto.CommonResponse{Code: 20000}

// 	var err error
// 	if err = c.BindJSON(req); err != nil {
// 		resp.Code = 400
// 		resp.Message = err.Error()
// 		c.JSON(http.StatusOK, resp)
// 		log.Warnf(context.Background(), "TransID:%s,请求绑定物料载具接口参数无效:%v", transID, err)
// 		return
// 	}

// 	if err = ucmiddleware.Validate.Struct(req); err != nil {
// 		resp.Code = 400
// 		resp.Message = err.Error()
// 		c.JSON(http.StatusOK, resp)
// 		return
// 	}

// 	if err := logic.BindMaterialTray(req); err != nil {
// 		resp.Code = 500
// 		resp.Message = err.Error()
// 	}

// 	c.JSON(http.StatusOK, resp)
// }

// func RegisterMaterialRouter(r *gin.Engine) {
// 	g := r.Group("/api/mom/webapi/material")

// 	g.POST("bindmaterialtray", BindMaterialTray)
// }
