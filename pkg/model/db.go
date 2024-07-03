package model

import (
	"time"

	"github.com/CloudSilk/pkg/db"
	"github.com/CloudSilk/pkg/db/mysql"
	"github.com/CloudSilk/pkg/db/sqlite"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var DB db.DBClientInterface

// Init Init
func Init(connStr string, debug bool) {
	DB = mysql.NewMysql(connStr, debug)
	AutoMigrate()
}

func InitSqlite(database string, debug bool) {
	DB = sqlite.NewSqlite2("", "", database, "", debug)
	if debug {
		AutoMigrate()
	}
}

func InitDB(client db.DBClientInterface, debug bool) {
	DB = client
	if debug {
		AutoMigrate()
	}
}

// AutoMigrate 自动生成表
func AutoMigrate() {
	DB.DB().AutoMigrate(
		&LabelTemplate{},
		&LabelParameter{},
		&LabelType{},
		&LabelPrintQueue{},
		&LabelPrintQueueParameter{},
		&LabelPrintQueueExecution{},
		&LabelPrintTask{},
		&LabelAdaptationRule{},

		&ProductCategory{},
		&ProductAttribute{},
		&ProductAttributeIdentifier{},
		&ProductAttributeIdentifierAvailableCategory{},
		&ProductBrand{},
		&ProductCategoryAttribute{},
		&ProductCategoryAttributeValue{},
		&ProductModelBom{},
		&ProductModel{},
		&ProductModelAttributeValue{},

		&MaterialCategory{},
		&MaterialInfo{},
		&MaterialSupplier{},
		&AvailableMaterial{},
		&MaterialTray{},
		&MaterialTrayBindingRecord{},
		&MaterialChannelLayer{},
		&MaterialChannel{},
		&MaterialReturnRequestForm{},
		&MaterialReturnType{},
		&MaterialReturnCause{},
		&MaterialReturnCauseAvailableCategory{},
		&MaterialReturnCauseAvailableType{},
		&MaterialReturnSolution{},
		&MaterialReturnSolutionAvailableCause{},
		&MaterialContainer{},
		&MaterialContainerType{},
		&MaterialShelf{},
		&MaterialShelfBin{},
		&MaterialStore{},
		&MaterialStoreAvailableLine{},
		&MaterialInventory{},
		&MaterialStoreFeedRule{},
		&AGVTaskType{},
		&AGVTaskQueue{},
		&WMSBillQueue{},

		&ProductInfo{},
		&ProductOrderAttribute{},
		&ProductOrderPackage{},
		&ProductOrderProcessStep{},
		&ProductOrderProcessStepAttachment{},
		&ProductOrderProcessStepTypeParameter{},
		&ProductOrderProcess{},
		&ProductOrder{},
		&ProductOrderBom{},
		&ProductPackageType{},
		&ProductPackage{},
		&ProductPackageStackRule{},
		&ProductPackageMatchRule{},
		&ProductProcessRoute{},
		&ProductionProcessSop{},
		&ProductionProcessStep{},
		&AvailableProcess{},
		&ProductionRhythm{},
		&ProcessStepType{},
		&ProcessStepTypeParameter{},
		&ProductReworkType{},
		&ProductReworkSolution{},
		&ProductReworkCauseAvailableSolution{},
		&ProductReworkCause{},
		&ProductReworkTypePossibleCause{},
		&ProductReleaseStrategy{},
		&ProductReleaseStrategyComparableAttribute{},
		&ProductAttributeValuateRule{},
		&ProductOrderPriorityRule{},
		&ProductOrderReleaseRule{},
		&ProductOrderAttachment{},
		&ProductOrderLabel{},
		&ProductOrderLabelParameter{},
		&ProductOrderPallet{},
		&ProductIssueRecord{},
		&ProductReleaseRecord{},
		&ProductProcessRecord{},
		&ProductRhythmRecord{},
		&ProductWorkRecord{},
		&ProductTestRecord{},
		&ProductReworkRecord{},
		&ProductPackageRecord{},
		&ProcessStepParameter{},
		&ProcessStepParameterValue{},
		&ProductReworkOperation{},
		&ProductReworkRoute{},
		&ProductReworkProcess{},
		&ProductReworkProcessAvailableStation{},
		&ProductReworkProcessAvailableProcess{},

		&ProductionProcess{},
		&ProductionStationAlarm{},
		&ProductionStationBreakdown{},
		&ProductionStationOutput{},
		&ProductionStationSignup{},
		&ProductionStationStartup{},
		&ProductionCrossway{},
		&ProductionCrosswayLeftTurnStation{},
		&ProductionCrosswayRightTurnStation{},
		&ProductionCrosswayStraightStation{},
		&ProductionFactory{},
		&ProductionLine{},
		&ProductionLineSupportableCategory{},
		&ProductionStation{},
		&ProductionProcessAvailableStation{},

		&RemoteServiceTaskQueue{},
		&RemoteServiceTask{},
		&RemoteServiceTaskParameter{},
		&RemoteService{},
		&SerialNumber{},
		&SystemEventTrigger{},
		&SystemEvent{},
		&SystemEventParameter{},
		&SystemEventSubscription{},
		&SystemParamsConfig{},
		&TaskQueueExecution{},
		&TaskQueue{},
		&TaskQueueParameter{},
		&InvocationAuthentication{},
		&DataMapping{},
		&CodingGeneration{},
		&CodingSerial{},
		&CodingTemplate{},
		&CodingElement{},
		&CodingElementValue{},
		&PrintServer{},
		&Printer{},

		&OperationTrace{},
		&InvocationTrace{},
		&ExceptionTrace{},

		&PersonnelQualification{},
		&PersonnelQualificationType{},
		&PersonnelQualificationTypeAvailableModel{},

		&AttributeExpression{},
		&PropertyExpression{},

		&ProductionStationEfficiency{},
		&ProductionEfficiency{},
	)
}

type ModelID struct {
	ID        string    `json:"id" gorm:"primarykey;size:36"`
	CreatedAt time.Time `json:"createdAt" copier:"-"`
}

func (u *ModelID) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.NewString()
	}
	return
}
