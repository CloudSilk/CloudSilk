package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm/clause"
)

func CreateMaterialReturnRequestForm(m *model.MaterialReturnRequestForm) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdateMaterialReturnRequestForm(m *model.MaterialReturnRequestForm) error {
	return model.DB.DB().Omit("created_at", "create_time").Save(m).Error
}

func QueryMaterialReturnRequestForm(req *proto.QueryMaterialReturnRequestFormRequest, resp *proto.QueryMaterialReturnRequestFormResponse, preload bool) {
	db := model.DB.DB().Model(&model.MaterialReturnRequestForm{})

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "created_at desc")
	if err != nil {
		resp.Code = proto.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.MaterialReturnRequestForm
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = proto.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.MaterialReturnRequestFormsToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllMaterialReturnRequestForms() (list []*model.MaterialReturnRequestForm, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetMaterialReturnRequestFormByID(id string) (*model.MaterialReturnRequestForm, error) {
	m := &model.MaterialReturnRequestForm{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetMaterialReturnRequestFormByIDs(ids []string) ([]*model.MaterialReturnRequestForm, error) {
	var m []*model.MaterialReturnRequestForm
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeleteMaterialReturnRequestForm(id string) (err error) {
	return model.DB.DB().Delete(&model.MaterialReturnRequestForm{}, "`id` = ?", id).Error
}

// func ApplyMaterialReturnRequestForm(ids []string) (err error) {
// 	var materialReturnRequestForms []*model.MaterialReturnRequestForm
// 	if err = model.DB.DB().Preload("MaterialReturnType").Preload("MaterialReturnCause").Preload("MaterialInfo").Preload("MaterialSupplier").Preload("ProductionLine").Where("id in (?)", ids).Find(&m).Error; err != nil {
// 		return
// 	}

// 	productionLineID := ""
// 	for i, v := range materialReturnRequestForms {
// 		if v.CheckTime.Valid {
// 			return fmt.Errorf("请勿选择已复核的记录，只能选择没有复核时间和复核人员的记录发送退料申请邮件")
// 		}

// 		if i == 0 {
// 			productionLineID = v.ProductionLineID
// 		} else if v.ProductionLineID != productionLineID {
// 			return fmt.Errorf("请勿跨产线选择退料申请记录，一次批量申请只能选择同一个产线的记录")
// 		}

// 		if v.CurrentState == types.MaterialReturnRequestFormStateApplied || v.CurrentState == types.MaterialReturnRequestFormStateNoRequired {
// 			return fmt.Errorf("该申请单已申请或不需申请,请核对")
// 		}
// 	}

// 	if productionLineID == "" {
// 		return fmt.Errorf("产线或工位不能为空，请核对！")
// 	}

// 	configCode := "10005"
// 	materialReturnRemindConfig := &model.SystemParamsConfig{}
// 	if err = model.DB.DB().First(materialReturnRemindConfig, "`code` = ?", configCode).Error; err == gorm.ErrRecordNotFound {
// 		return fmt.Errorf("未能找到代号为%s的系统配置，请添加此配置项！", configCode)
// 	} else if err != nil {
// 		return
// 	}

// 	spConfigCode := "SP001"
// 	spConfig := &model.SystemParamsConfig{}
// 	if err = model.DB.DB().First(spConfig, "`code` = ?", spConfigCode).Error; err == gorm.ErrRecordNotFound {
// 		return fmt.Errorf("未能找到代号为%s的系统配置，请添加此配置项！", spConfigCode)
// 	} else if err != nil {
// 		return
// 	}

// 	messageSendTaskCodes := []string{}
// 	for _, v := range strings.Split(materialReturnRemindConfig.Value, ",") {
// 		if v != "" {
// 			messageSendTaskCodes = append(messageSendTaskCodes, v)
// 		}
// 	}

// 	productionLine := materialReturnRequestForms[0].ProductionLine

// 	return nil
// }
