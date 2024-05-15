package logic

import (
	"github.com/CloudSilk/CloudSilk/pkg/model"
	apipb "github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreatePrintServer(m *model.PrintServer) (string, error) {
	err := model.DB.DB().Create(m).Error
	return m.ID, err
}

func UpdatePrintServer(m *model.PrintServer) error {
	return model.DB.DB().Omit("created_at").Save(m).Error
}

func QueryPrintServer(req *apipb.QueryPrintServerRequest, resp *apipb.QueryPrintServerResponse, preload bool) {
	db := model.DB.DB().Model(&model.PrintServer{})

	orderStr, err := utils.GenerateOrderString(req.SortConfig, "id")
	if err != nil {
		resp.Code = apipb.Code_BadRequest
		resp.Message = err.Error()
		return
	}

	var list []*model.PrintServer
	resp.Records, resp.Pages, err = model.DB.PageQuery(db, req.PageSize, req.PageIndex, orderStr, &list)
	if err != nil {
		resp.Code = apipb.Code_InternalServerError
		resp.Message = err.Error()
	} else {
		resp.Data = model.PrintServersToPB(list)
	}
	resp.Total = resp.Records
}

func GetAllPrintServers() (list []*model.PrintServer, err error) {
	err = model.DB.DB().Find(&list).Error
	return
}

func GetPrintServerByID(id string) (*model.PrintServer, error) {
	m := &model.PrintServer{}
	err := model.DB.DB().Preload(clause.Associations).Where("id = ?", id).First(m).Error
	return m, err
}

func GetPrintServerByIDs(ids []string) ([]*model.PrintServer, error) {
	var m []*model.PrintServer
	err := model.DB.DB().Preload(clause.Associations).Where("id in (?)", ids).Find(&m).Error
	return m, err
}

func DeletePrintServer(id string) (err error) {
	return model.DB.DB().Transaction(func(tx *gorm.DB) error {
		var printers []model.Printer
		if err := tx.Where("print_server_id = ?", id).Find(&printers).Error; err != nil {
			return err
		}

		for _, printer := range printers {
			if err := tx.Delete(&printer).Error; err != nil {
				return err
			}
		}

		if err := tx.Delete(&model.PrintServer{}, "id=?", id).Error; err != nil {
			return err
		}

		return nil
	})
}
