package clients

import (
	"os"

	"dubbo.apache.org/dubbo-go/v3/config"
	ucprovider "github.com/CloudSilk/usercenter/provider"
)

func Init(serviceMode string) {
	if serviceMode == "ALL" {
		userProvider := new(ucprovider.UserProvider)
		UserClient.Add = userProvider.Add
		UserClient.Delete = userProvider.Delete
		UserClient.Export = userProvider.Export
		UserClient.GetDetail = userProvider.GetDetail
		UserClient.Query = userProvider.Query
		UserClient.Update = userProvider.Update
		UserClient.LoginByStaffNo = userProvider.LoginByStaffNo
		UserClient.LogoutByUserName = userProvider.LogoutByUserName
	} else {
		if os.Getenv("MES_DISABLE_AUTH") == "true" {
			config.SetConsumerService(UserClient)
		}
	}
}
