package clients

import (
	"os"

	"dubbo.apache.org/dubbo-go/v3/config"
	mpp "github.com/CloudSilk/CloudSilk/pkg/servers/material/provider"
	ppp "github.com/CloudSilk/CloudSilk/pkg/servers/product/provider"
	pbpp "github.com/CloudSilk/CloudSilk/pkg/servers/product_base/provider"
	ptpp "github.com/CloudSilk/CloudSilk/pkg/servers/production/provider"
	ptbpp "github.com/CloudSilk/CloudSilk/pkg/servers/production_base/provider"
	spp "github.com/CloudSilk/CloudSilk/pkg/servers/system/provider"
	upp "github.com/CloudSilk/CloudSilk/pkg/servers/user/provider"
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

		productProvider := new(ppp.ProductInfoProvider)
		ProductInfoClient.Get = productProvider.Get
		ProductInfoClient.Query = productProvider.Query

		productionStationProvider := new(ptbpp.ProductionStationProvider)
		ProductionStationClient.Get = productionStationProvider.Get
		ProductionStationClient.Add = productionStationProvider.Add
		ProductionStationClient.Update = productionStationProvider.Update
		ProductionStationClient.Query = productionStationProvider.Query
		ProductionStationClient.Delete = productionStationProvider.Delete
		ProductionStationClient.GetAll = productionStationProvider.GetAll
		ProductionStationClient.GetDetail = productionStationProvider.GetDetail

		productionStationSignupProvider := new(ptpp.ProductionStationSignupProvider)
		ProductionStationSignupClient.Add = productionStationSignupProvider.Add
		ProductionStationSignupClient.Update = productionStationSignupProvider.Update
		ProductionStationSignupClient.Query = productionStationSignupProvider.Query
		ProductionStationSignupClient.Delete = productionStationSignupProvider.Delete
		ProductionStationSignupClient.GetAll = productionStationSignupProvider.GetAll
		ProductionStationSignupClient.GetDetail = productionStationSignupProvider.GetDetail

		productionLineSignupProvider := new(ptbpp.ProductionLineProvider)
		ProductionLineClient.Add = productionLineSignupProvider.Add
		ProductionLineClient.Update = productionLineSignupProvider.Update
		ProductionLineClient.Query = productionLineSignupProvider.Query
		ProductionLineClient.Delete = productionLineSignupProvider.Delete
		ProductionLineClient.GetAll = productionLineSignupProvider.GetAll
		ProductionLineClient.GetDetail = productionLineSignupProvider.GetDetail

		productAttributeProvider := new(pbpp.ProductAttributeProvider)
		ProductAttributeClient.Add = productAttributeProvider.Add
		ProductAttributeClient.Update = productAttributeProvider.Update
		ProductAttributeClient.Query = productAttributeProvider.Query
		ProductAttributeClient.Delete = productAttributeProvider.Delete
		ProductAttributeClient.GetAll = productAttributeProvider.GetAll
		ProductAttributeClient.GetDetail = productAttributeProvider.GetDetail

		productionCrosswayProvider := new(ptbpp.ProductionCrosswayProvider)
		ProductionCrosswayClient.Add = productionCrosswayProvider.Add
		ProductionCrosswayClient.Update = productionCrosswayProvider.Update
		ProductionCrosswayClient.Query = productionCrosswayProvider.Query
		ProductionCrosswayClient.Delete = productionCrosswayProvider.Delete
		ProductionCrosswayClient.GetAll = productionCrosswayProvider.GetAll
		ProductionCrosswayClient.GetDetail = productionCrosswayProvider.GetDetail

		materialTrayProvider := new(mpp.MaterialTrayProvider)
		MaterialTrayClient.Get = materialTrayProvider.Get

		productPackageRecordProvider := new(ppp.ProductPackageRecordProvider)
		ProductPackageRecordClient.Get = productPackageRecordProvider.Get

		productInfoProvider := new(ppp.ProductInfoProvider)
		ProductInfoClient.Get = productInfoProvider.Get
		ProductInfoClient.Query = productInfoProvider.Query

		productOrderProvider := new(ppp.ProductOrderProvider)
		ProductOrderClient.GetDetail = productOrderProvider.GetDetail

		productRhythmRecordProvider := new(ppp.ProductRhythmRecordProvider)
		ProductRhythmRecordClient.Add = productRhythmRecordProvider.Add
		ProductRhythmRecordClient.Get = productRhythmRecordProvider.Get

		productProcessRouteProvider := new(ppp.ProductProcessRouteProvider)
		ProductProcessRouteClient.Add = productProcessRouteProvider.Add
		// ProductProcessRouteClient.Get = productProcessRouteProvider.Get
		ProductProcessRouteClient.Query = productProcessRouteProvider.Query

		productionProcessSopProvider := new(ptpp.ProductionProcessSopProvider)
		ProductionProcessSopClient.Get = productionProcessSopProvider.Get

		productModelProvider := new(pbpp.ProductModelProvider)
		ProductModelClient.GetDetail = productModelProvider.GetDetail

		personnelQualificationProvider := new(upp.PersonnelQualificationProvider)
		PersonnelQualificationClient.Get = personnelQualificationProvider.Get
		PersonnelQualificationClient.Query = personnelQualificationProvider.Query

		systemEventProvider := new(spp.SystemEventProvider)
		SystemEventClient.Get = systemEventProvider.Get

		systemEventTriggerProvider := new(spp.SystemEventTriggerProvider)
		SystemEventTriggerClient.Add = systemEventTriggerProvider.Add

		systemEventTriggerParameterProvider := new(spp.SystemEventTriggerParameterProvider)
		SystemEventTriggerParameterClient.Add = systemEventTriggerParameterProvider.Add

		productReworkRecordProvider := new(ppp.ProductReworkRecordProvider)
		ProductReworkRecordClient.Add = productReworkRecordProvider.Add

		productOrderProcessProvider := new(ppp.ProductOrderProcessProvider)
		ProductOrderProcessClient.Query = productOrderProcessProvider.Query

		productionStationOutputProvider := new(ptpp.ProductionStationOutputProvider)
		ProductionStationOutputClient.Add = productionStationOutputProvider.Add
		ProductionStationOutputClient.Get = productionStationOutputProvider.Get

		productOrderAttributeProvider := new(ppp.ProductOrderAttributeProvider)
		ProductOrderAttributeClient.Query = productOrderAttributeProvider.Query

		processStepParameterProvider := new(ptpp.ProcessStepParameterProvider)
		ProcessStepParameterClient.Query = processStepParameterProvider.Query

		productionProcessStepProvider := new(ptpp.ProductionProcessStepProvider)
		ProductionProcessStepClient.Get = productionProcessStepProvider.Get
		ProductionProcessStepClient.Query = productionProcessStepProvider.Query

		productTestRecordProvider := new(ppp.ProductTestRecordProvider)
		ProductTestRecordClient.Add = productTestRecordProvider.Add
	} else {
		if os.Getenv("MES_DISABLE_AUTH") != "true" {
			config.SetConsumerService(UserClient)
		}
		config.SetConsumerService(ProductionStationClient)
		config.SetConsumerService(ProductionStationSignupClient)
		config.SetConsumerService(ProductionLineClient)
		config.SetConsumerService(ProductAttributeClient)
		config.SetConsumerService(ProductionCrosswayClient)
		config.SetConsumerService(MaterialTrayClient)
		config.SetConsumerService(ProductPackageRecordClient)
		config.SetConsumerService(ProductInfoClient)
		config.SetConsumerService(ProductOrderClient)
		config.SetConsumerService(ProductRhythmRecordClient)
		config.SetConsumerService(ProductProcessRouteClient)
		config.SetConsumerService(ProductionProcessClient)
		config.SetConsumerService(ProductionProcessSopClient)
		config.SetConsumerService(ProductModelClient)
		config.SetConsumerService(PersonnelQualificationClient)
		config.SetConsumerService(SystemEventClient)
		config.SetConsumerService(SystemEventTriggerClient)
		config.SetConsumerService(SystemEventTriggerParameterClient)
		config.SetConsumerService(ProductReworkRecordClient)
		config.SetConsumerService(ProductOrderProcessClient)
		config.SetConsumerService(ProductionStationOutputClient)
		config.SetConsumerService(ProductOrderAttributeClient)
		config.SetConsumerService(ProcessStepParameterClient)
		config.SetConsumerService(ProductionProcessStepClient)
		config.SetConsumerService(ProductTestRecordClient)
	}
}
