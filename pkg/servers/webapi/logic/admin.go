package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/CloudSilk/CloudSilk/pkg/clients"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	modelcode "github.com/CloudSilk/pkg/model"
	usercenter "github.com/CloudSilk/usercenter/proto"
	"gorm.io/gorm"
)

func TryLogin(req *proto.LoginRequest, resp *proto.ServiceResponse) {
	var user *usercenter.UserProfile
	var token string
	if req.CardNo == "" {
		if req.Name == "" {
			resp.Code = modelcode.BadRequest
			resp.Message = "账号不能为空"
			return
		}
		if req.Password == "" {
			resp.Code = modelcode.BadRequest
			resp.Message = "密码不能为空"
			return
		}

		//根据帐号获取用户
		_user, _ := clients.UserClient.LoginByStaffNo(context.Background(), &usercenter.LoginByStaffNoRequest{UserName: req.Name, Password: req.Password})
		if _user.Code == modelcode.UserIsNotExist || _user.Code == modelcode.UserNameOrPasswordIsWrong {
			resp.Code = modelcode.UserIsNotExist
			resp.Message = "用户或密码错误"
			return
		}
		if _user.Code == modelcode.UserDisabled {
			resp.Code = modelcode.UserDisabled
			resp.Message = "用户已停用"
			return
		}

		user = _user.User
		token = _user.Message
	} else {
		//根据卡号获取用户
		_user, _ := clients.UserClient.LoginByStaffNo(context.Background(), &usercenter.LoginByStaffNoRequest{StaffNo: req.CardNo})
		if _user.Code == modelcode.UserIsNotExist {
			resp.Code = modelcode.UserIsNotExist
			resp.Message = "卡号不存在，请联系产线主管"
			return
		}
		if _user.Code == modelcode.UserDisabled {
			resp.Code = modelcode.UserDisabled
			resp.Message = "用户已停用"
			return
		}

		user = _user.User
		token = _user.Message
	}

	if req.ProductionStation == "" {
		resp.Code = modelcode.BadRequest
		resp.Message = "ProductionStation不能为空"
		return
	} else {
		_productionStation, err := clients.ProductionStationClient.Get(context.Background(), &proto.GetProductionStationRequest{Code: req.ProductionStation})
		if err == gorm.ErrRecordNotFound {
			resp.Code = modelcode.BadRequest
			resp.Message = "无效的工位编号"
			return
		} else if err != nil {
			resp.Code = modelcode.InternalServerError
			resp.Message = err.Error()
			return
		}

		productionStation := _productionStation.Data

		productionStation.CurrentUserID = user.Id
		_productionStationSignup, err := clients.ProductionStationSignupClient.GetByID(context.Background(), &proto.GetProductionStationSignupRequest{
			ProductionStationID: productionStation.Id,
			LoginUserID:         user.Id,
			HasLogoutTime:       false,
		})
		if err != nil {
			resp.Code = modelcode.InternalServerError
			resp.Message = err.Error()
			return
		}

		productionStationSignup := _productionStationSignup.Data
		now := time.Now()
		if productionStationSignup == nil {
			productionStationSignup = &proto.ProductionStationSignupInfo{
				LoginUserID:         user.Id,
				ProductionStationID: productionStation.Id,
				LoginTime:           now.Format("2006-01-02 15:04:05"),
			}
			if _, err := clients.ProductionStationSignupClient.Add(context.Background(), productionStationSignup); err != nil {
				resp.Code = modelcode.InternalServerError
				resp.Message = err.Error()
				return
			}
		}

		loginTime, err := time.Parse("2006-01-02 15:04:05", productionStationSignup.LoginTime)
		if err != nil {
			resp.Code = modelcode.InternalServerError
			resp.Message = err.Error()
			return
		}
		productionStationSignup.LastHeartbeatTime = now.Format("2006-01-02 15:04:05")
		productionStationSignup.Duration = int32(now.Sub(loginTime).Minutes())

		clients.ProductionStationSignupClient.Update(context.Background(), productionStationSignup)
		clients.ProductionStationClient.Update(context.Background(), productionStation)
	}

	resp.Message = token
	resp.Data = &proto.Data{
		Id:          user.Id,
		Name:        user.UserName,
		ChineseName: user.ChineseName,
		EnglishName: user.EnglishName,
		Photo:       user.Mobile,
		StaffNo:     user.StaffNo,
	}
}

func TryLogout(req *proto.LogoutRequest) error {
	if req.Name == "" {
		return fmt.Errorf("Name不能为空")
	}
	if req.ProductionStation == "" {
		return fmt.Errorf("ProductionStation不能为空")
	}

	_productionStation, err := clients.ProductionStationClient.Get(context.Background(), &proto.GetProductionStationRequest{Code: req.ProductionStation})
	if err != nil {
		return err
	}

	productionStation := _productionStation.Data
	if productionStation == nil {
		return fmt.Errorf("无效的工位编号")
	}

	//根据帐号注销用户
	_user, _ := clients.UserClient.LogoutByUserName(context.Background(), &usercenter.LogoutByUserNameRequest{UserName: req.Name})
	if _user.Code != modelcode.Success {
		return fmt.Errorf(_user.Message)
	}

	userID := _user.Message
	if userID == "" {
		return fmt.Errorf("无效的工位或账号")
	}

	if productionStation.CurrentUserID != userID {
		return fmt.Errorf("登录信息错误，用于账号与工位登录当前账号不符")
	}

	productionStation.CurrentUserID = ""

	_productionStationSignup, err := clients.ProductionStationSignupClient.GetByID(context.Background(), &proto.GetProductionStationSignupRequest{
		ProductionStationID: productionStation.Id,
		LoginUserID:         productionStation.CurrentUserID,
		HasLogoutTime:       false,
	})
	if err != nil {
		return err
	}

	productionStationSignup := _productionStationSignup.Data
	if productionStationSignup != nil {
		productionStationSignup.LastHeartbeatTime = time.Now().Format("2006-01-02 15:04:05")
		productionStationSignup.Duration = 0
	}

	clients.ProductionStationClient.Update(context.Background(), productionStation)
	clients.ProductionStationSignupClient.Update(context.Background(), productionStationSignup)

	return nil
}
