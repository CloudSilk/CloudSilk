package logic

import (
	"testing"

	"github.com/CloudSilk/CloudSilk/pkg/model"
	"github.com/CloudSilk/CloudSilk/pkg/proto"
	"github.com/CloudSilk/CloudSilk/pkg/testutil"
	"gorm.io/gorm"
)

func setupQMDB(t *testing.T) *gorm.DB {
	t.Helper()
	gdb, err := testutil.SetupTestDB()
	if err != nil {
		t.Fatalf("初始化测试数据库失败: %v", err)
	}
	return gdb
}

// 造一套检验体系：类型 + 一个范围内判定的标准 + 已生成明细的检验单
func seedInspectionOrder(t *testing.T, gdb *gorm.DB, lotQTY int32) *model.QualityInspectionOrder {
	t.Helper()
	it := &model.QualityInspectionType{Code: "IQC", Description: "来料检验", Enable: true}
	if err := gdb.Create(it).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}
	standard := &model.QualityInspectionStandard{
		Code:            "STD-LEN",
		Description:     "长度",
		Unit:            "mm",
		CriterionMethod: CriterionWithinRange,
		LowerLimit:      "9.5",
		UpperLimit:      "10.5",
		Aql:             1.0,
		SampleLevel:     "II",
		QualityInspectionTypeID: &it.ID,
		Enable:          true,
	}
	if err := gdb.Create(standard).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}

	order := &model.QualityInspectionOrder{
		QualityInspectionTypeID: &it.ID,
		LotQTY:                  lotQTY,
		SampleQTY:               SampleSizeByLot(lotQTY, "II"),
		CurrentState:            QualityInspectionStateWait,
	}
	item := &model.QualityInspectionOrderItem{
		QualityInspectionStandardID: standard.ID,
	}
	order.QualityInspectionOrderItems = append(order.QualityInspectionOrderItems, item)
	if err := gdb.Create(order).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}
	if err := gdb.Model(&model.QualityInspectionOrderItem{}).Where("`id` = ?", item.ID).
		Update("quality_inspection_order_id", order.ID).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}
	return order
}

// 合格路径：实测值在规格内 → 判定合格，不产生返工
func TestCompleteInspection_Qualified(t *testing.T) {
	gdb := setupQMDB(t)
	order := seedInspectionOrder(t, gdb, 100)
	standardID := order.QualityInspectionOrderItems[0].QualityInspectionStandardID

	err := CompleteQualityInspectionOrder(&proto.CompleteQualityInspectionOrderRequest{
		Id: order.ID,
		Items: []*proto.QualityInspectionOrderItemInfo{
			{QualityInspectionStandardID: standardID, MeasuredValue: "10.0"},
		},
	})
	if err != nil {
		t.Fatalf("完成检验失败: %v", err)
	}

	after := &model.QualityInspectionOrder{}
	gdb.Preload("QualityInspectionOrderItems").First(after, "`id` = ?", order.ID)
	if after.Conclusion != QualityConclusionQualified {
		t.Fatalf("结论应为合格，实际%s", after.Conclusion)
	}
	if !after.ActualInspectionTime.Valid {
		t.Fatalf("应记录实际检验时间")
	}
	item := after.QualityInspectionOrderItems[0]
	if !item.IsQualified || item.MeasuredValue != "10.0" {
		t.Fatalf("明细应判定合格且记录实测值")
	}
}

// 不合格路径：超上限 + 勾选返工 → 返工记录 + 状态联动
func TestCompleteInspection_UnqualifiedWithRework(t *testing.T) {
	gdb := setupQMDB(t)
	product := &model.ProductInfo{ProductSerialNo: "SN-Q-001", CurrentState: "已发放"}
	if err := gdb.Create(product).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}
	order := seedInspectionOrder(t, gdb, 100)
	order.ProductInfoID = &product.ID
	gdb.Save(order)
	standardID := order.QualityInspectionOrderItems[0].QualityInspectionStandardID

	err := CompleteQualityInspectionOrder(&proto.CompleteQualityInspectionOrderRequest{
		Id:          order.ID,
		Items:       []*proto.QualityInspectionOrderItemInfo{
			{QualityInspectionStandardID: standardID, MeasuredValue: "11.0"},
		},
		CreateRework: true,
	})
	if err != nil {
		t.Fatalf("完成检验失败: %v", err)
	}

	after := &model.QualityInspectionOrder{}
	gdb.First(after, "`id` = ?", order.ID)
	if after.Conclusion != QualityConclusionUnqualified {
		t.Fatalf("结论应为不合格，实际%s", after.Conclusion)
	}
	if !after.ReworkCreated {
		t.Fatalf("应已生成返工记录")
	}

	var reworks int64
	gdb.Model(&model.ProductReworkRecord{}).Where("`product_info_id` = ?", product.ID).Count(&reworks)
	if reworks != 1 {
		t.Fatalf("应产生1条返工记录，实际%d", reworks)
	}

	pi := &model.ProductInfo{}
	gdb.First(pi, "`id` = ?", product.ID)
	if pi.CurrentState != "检查中" {
		t.Fatalf("产品状态应联动为检查中，实际%s", pi.CurrentState)
	}
}

// 让步接收：判定不合格但特采放行 → 不生成返工、不改产品状态
func TestCompleteInspection_Concession(t *testing.T) {
	gdb := setupQMDB(t)
	product := &model.ProductInfo{ProductSerialNo: "SN-C-001", CurrentState: "已发放"}
	if err := gdb.Create(product).Error; err != nil {
		t.Fatalf("造数失败: %v", err)
	}
	order := seedInspectionOrder(t, gdb, 100)
	order.ProductInfoID = &product.ID
	gdb.Save(order)
	standardID := order.QualityInspectionOrderItems[0].QualityInspectionStandardID

	err := CompleteQualityInspectionOrder(&proto.CompleteQualityInspectionOrderRequest{
		Id:          order.ID,
		Items:       []*proto.QualityInspectionOrderItemInfo{
			{QualityInspectionStandardID: standardID, MeasuredValue: "99"},
		},
		Concession: true,
	})
	if err != nil {
		t.Fatalf("完成检验失败: %v", err)
	}

	after := &model.QualityInspectionOrder{}
	gdb.First(after, "`id` = ?", order.ID)
	if after.Conclusion != QualityConclusionConcession {
		t.Fatalf("结论应为让步接收，实际%s", after.Conclusion)
	}
	if after.ReworkCreated {
		t.Fatalf("让步接收不应生成返工")
	}

	var reworks int64
	gdb.Model(&model.ProductReworkRecord{}).Count(&reworks)
	if reworks != 0 {
		t.Fatalf("不应产生返工记录，实际%d", reworks)
	}
	pi := &model.ProductInfo{}
	gdb.First(pi, "`id` = ?", product.ID)
	if pi.CurrentState != "已发放" {
		t.Fatalf("让步接收不应改产品状态，实际%s", pi.CurrentState)
	}
}

// 重复提交与缺检测结果应拒绝
func TestCompleteInspection_Invalid(t *testing.T) {
	gdb := setupQMDB(t)
	order := seedInspectionOrder(t, gdb, 100)
	standardID := order.QualityInspectionOrderItems[0].QualityInspectionStandardID

	// 缺检测结果
	if err := CompleteQualityInspectionOrder(&proto.CompleteQualityInspectionOrderRequest{
		Id:    order.ID,
		Items: []*proto.QualityInspectionOrderItemInfo{},
	}); err == nil {
		t.Fatalf("缺少检测结果应拒绝")
	}

	// 正常完成一次
	if err := CompleteQualityInspectionOrder(&proto.CompleteQualityInspectionOrderRequest{
		Id:    order.ID,
		Items: []*proto.QualityInspectionOrderItemInfo{
			{QualityInspectionStandardID: standardID, MeasuredValue: "10"},
		},
	}); err != nil {
		t.Fatalf("完成检验失败: %v", err)
	}

	// 重复提交
	if err := CompleteQualityInspectionOrder(&proto.CompleteQualityInspectionOrderRequest{
		Id:    order.ID,
		Items: []*proto.QualityInspectionOrderItemInfo{
			{QualityInspectionStandardID: standardID, MeasuredValue: "10"},
		},
	}); err == nil {
		t.Fatalf("已完成检验单不能重复提交")
	}
}

// 草稿保存：部分填报不出结论、状态推进到检验中；随后补齐可完成
func TestCompleteInspection_DraftThenFinish(t *testing.T) {
	gdb := setupQMDB(t)
	order := seedInspectionOrder(t, gdb, 100)
	standardID := order.QualityInspectionOrderItems[0].QualityInspectionStandardID

	// 草稿：saveOnly
	if err := CompleteQualityInspectionOrder(&proto.CompleteQualityInspectionOrderRequest{
		Id:        order.ID,
		SaveOnly:  true,
		Items:     []*proto.QualityInspectionOrderItemInfo{
			{QualityInspectionStandardID: standardID, MeasuredValue: "9.8"},
		},
	}); err != nil {
		t.Fatalf("草稿保存失败: %v", err)
	}
	after := &model.QualityInspectionOrder{}
	gdb.Preload("QualityInspectionOrderItems").First(after, "`id` = ?", order.ID)
	if after.Conclusion != "" || after.CurrentState != QualityInspectionStateDoing {
		t.Fatalf("草稿不应出结论，状态应为检验中，实际 结论=%s 状态=%s", after.Conclusion, after.CurrentState)
	}
	if !after.QualityInspectionOrderItems[0].IsQualified || after.QualityInspectionOrderItems[0].MeasuredValue != "9.8" {
		t.Fatalf("草稿应落库单项判定与实测值")
	}

	// 正式完成（复用同一数值）
	if err := CompleteQualityInspectionOrder(&proto.CompleteQualityInspectionOrderRequest{
		Id:    order.ID,
		Items: []*proto.QualityInspectionOrderItemInfo{
			{QualityInspectionStandardID: standardID, MeasuredValue: "9.8"},
		},
	}); err != nil {
		t.Fatalf("完成失败: %v", err)
	}
	gdb.First(after, "`id` = ?", order.ID)
	if after.Conclusion != QualityConclusionQualified {
		t.Fatalf("最终结论应为合格，实际%s", after.Conclusion)
	}
}

// 抽检模式：按该项标准 AQL 接收数判定（Ac=0 时 1 个不合格即整项不合格）
func TestCompleteInspection_SampleAQL(t *testing.T) {
	gdb := setupQMDB(t)
	order := seedInspectionOrder(t, gdb, 5) // 样本量5 → Ac(5,1.0)=0
	standardID := order.QualityInspectionOrderItems[0].QualityInspectionStandardID

	// 1个不良 > Ac=0 → 不合格
	if err := CompleteQualityInspectionOrder(&proto.CompleteQualityInspectionOrderRequest{
		Id:    order.ID,
		Items: []*proto.QualityInspectionOrderItemInfo{
			{QualityInspectionStandardID: standardID, BySample: true, SampleDefectives: 1},
		},
	}); err != nil {
		t.Fatalf("完成失败: %v", err)
	}
	after := &model.QualityInspectionOrder{}
	gdb.Preload("QualityInspectionOrderItems").First(after, "`id` = ?", order.ID)
	if after.Conclusion != QualityConclusionUnqualified {
		t.Fatalf("1不良>Ac0 应不合格，实际%s", after.Conclusion)
	}
	item := after.QualityInspectionOrderItems[0]
	if !item.BySample || item.SampleDefectives != 1 || item.IsQualified {
		t.Fatalf("抽检字段应落库且单项不合格")
	}
}
