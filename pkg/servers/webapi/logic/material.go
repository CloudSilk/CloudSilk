package logic

// 绑定物料载具，写入产品信息
// func BindMaterialTray(req *proto.BindMaterialTrayRequest) error {
// 	if req.MaterialTray == "" {
// 		return fmt.Errorf("MaterialTray不能为空")
// 	}
// 	if req.ProductSerialNo == "" {
// 		return fmt.Errorf("ProductSerialNo不能为空")
// 	}

// 	_materialTray, _ := clients.MaterialTrayClient.Get(context.Background(), &proto.GetMaterialTrayRequest{Identifier: req.MaterialTray})
// 	if _materialTray.Message == gorm.ErrRecordNotFound.Error() {
// 		return fmt.Errorf("无效的物料载具识别码")
// 	}
// 	if _materialTray.Code != modelcode.Success {
// 		return fmt.Errorf(_materialTray.Message)
// 	}
// 	materialTray := _materialTray.Data

// 	if !materialTray.Enable {
// 		return fmt.Errorf("载具已禁用")
// 	}

// 	_productInfo, _ := clients.ProductInfoClient.Get(context.Background(), &proto.GetProductInfoRequest{ProductSerialNo: req.ProductSerialNo})
// 	if _productInfo.Message == gorm.ErrRecordNotFound.Error() {
// 		return fmt.Errorf("无效的产品序列号")
// 	}
// 	if _productInfo.Code != modelcode.Success {
// 		return fmt.Errorf(_productInfo.Message)
// 	}
// 	productInfo := _productInfo.Data

// 	if materialTray.ProductInfoID != "" && materialTray.ProductInfoID != productInfo.Id {
// 		return fmt.Errorf("非法操作，此载具已绑定其他产品信息")
// 	}

// 	_MaterialTrayBindingRecord, _ := clients.MaterialTrayBindingRecordClient.Get(context.Background(), &proto.GetMaterialTrayBindingRecordRequest{
// 		ProductInfoID:  productInfo.Id,
// 		MaterialTrayID: materialTray.Id,
// 	})
// 	if _MaterialTrayBindingRecord.Code != modelcode.Success && _MaterialTrayBindingRecord.Message != gorm.ErrRecordNotFound.Error() {
// 		return fmt.Errorf(_MaterialTrayBindingRecord.Message)
// 	}

// 	if _MaterialTrayBindingRecord.Data == nil {
// 		if _materialTrayBindingRecord, _ := clients.MaterialTrayBindingRecordClient.Add(context.Background(), &proto.MaterialTrayBindingRecordInfo{
// 			MaterialTrayID: materialTray.Id,
// 			ProductInfoID:  productInfo.Id,
// 			CurrentState:   types.MaterialTrayBindingRecordStateEffected,
// 		}); _materialTrayBindingRecord.Code != modelcode.Success {
// 			return fmt.Errorf(_materialTrayBindingRecord.Message)
// 		}
// 	}
// 	materialTray.ProductInfoID = productInfo.Id
// 	materialTray.CurrentState = types.MaterialTrayStateFilled
// 	if _materialTray, _ := clients.MaterialTrayClient.Update(context.Background(), materialTray); _materialTray.Code != modelcode.Success {
// 		return fmt.Errorf(_materialTray.Message)
// 	}

// 	return nil
// }
