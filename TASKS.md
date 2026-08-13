# 云梭 MOM 生产运营管理系统 — 未完成功能任务清单

> 生成日期：2026-08-13（2026-08-14 更新：合并远程 main 的 PR #17-#20 后修订）
> 分析范围：后端 `pkg/`（432 个 Go 文件）、前端 `web/src`（206 个 TS/TSX 文件）、`csharp/PrintService`
> 分析方法：TODO/FIXME 标记扫描、被注释代码块排查、README 承诺功能与代码实现对照、前端路由/页面完整性检查

## 总体结论

系统实际完成度：**MES 生产执行核心流程基本可用**（上线装配、进/出站、报工、返工、工艺路线校验），外加 MSS 物料/仓储基础功能（仓库/货架/库位/库存/容器/AGV 任务/退料，来自 PR #17-#20）、低代码 CRUD 平台与 BarTender 标签打印服务。

但 README 宣传的六大模块中，**APS 排程、SCADA 网关、设备管理三块尚无代码实现**，质量管理仅有数据表与基础测试记录。此外存在若干被整段注释禁用的接口。

## 已完成（本轮维护）

- ~~TASK-003 infrastructure.go 字段拼写错误~~ — 已由 PR #17-#20 修复（`station_type`、`id`、`code` 均已正确）
- 合并远程 main 丢失的 PR #17-#20（webapi 事务处理、MSS/WMS 模块）
- 修复 PR #20 引入的编译错误（`personnel_qualification.go` 使用了已删除的 `req.Name` 字段）
- Dockerfile 修复：补齐 config.yaml/web/dist 拷贝、前端多阶段构建、端口对齐（见提交记录）

---

## 优先级说明

| 级别 | 含义 |
|------|------|
| P0 | 核心生产流程缺失或线上 Bug，直接影响业务正确性 |
| P1 | 已有骨架但被禁用/不完整的功能，恢复或补齐即可用 |
| P2 | README 承诺但未实现的模块，需要立项开发 |
| P3 | 前端问题与体验缺陷 |
| P4 | 代码质量、死代码清理、国际化 |

---

# 一、P0 — 核心流程缺失与线上 Bug（后端）

### TASK-001 ✅ 工单"签派"环节已实现
- **位置**：`pkg/servers/product/logic/product_order.go`、`pkg/servers/product/http/product_order.go`
- **完成内容**：
  - `DispatchProductOrder(tx, id, dispatchUserID, productionTeam)`：已核验 → 已签派，指派生产班组并写入 OperationTrace 签派轨迹
  - `BatchDispatchProductOrder(ids, ...)`：批量签派，事务包裹
  - HTTP `PUT /api/mom/product/productorder/dispatch?productionTeam=xxx`
  - 发放前置条件兼容"已核验/已签派"两种状态
  - 工单创建自动流转链插入签派步骤：接单 → 核验 → **签派** → 发放

### TASK-002 ✅ 重复进站报工去重已修复
- **位置**：`pkg/servers/webapi/logic/production.go`
- **根因**：进站查询未结束节拍记录时，`product_info_id` 占位符误传了 `productionStation.ID`，永远查不到已有记录，每次进站都新建节拍记录。修正参数后重复进站复用首次进站时间，不重复报工。

### TASK-003 ~~webapi 工位查询接口字段名拼写错误~~ ✅ 已修复
- **位置**：`pkg/servers/webapi/logic/infrastructure.go`
- **结果**：已由合并进来的 PR #17-#20 修复（`station_type` 过滤、返回字段 `id`/`code` 均已正确），无需再处理。

### TASK-004 ✅ 产品型号特性表达式正则解析 panic 防护已完成
- **位置**：`pkg/servers/product_base/logic/product_model.go`
- **完成内容**：`regexp.MustCompile` 改为 `regexp.Compile` + 明确错误返回（指出 RE2 不支持 .NET 环视/反向引用语法），保留 `(?<` → `(?P<` 具名分组转换。遇不兼容表达式时返回可读错误而非 panic。

### TASK-005 ✅ 更新接口空值覆盖外键已修复
- **位置**：`pkg/servers/product/logic/product_package_type.go`、`pkg/servers/product_base/logic/product_model_bom.go`
- **完成内容**：恢复条件 omit 逻辑——`LabelTypeID`/`SystemEventID`（`*string` 判 nil 及空串）与 `ProductModelID` 为空时加入 GORM Omit 列表，更新不再误清外键。

---

# 二、P1 — 被注释禁用的功能（恢复/补齐链路）

### TASK-006 ✅ 质量测试项获取接口已恢复
- **位置**：`pkg/servers/webapi/logic/quality.go`（重写）、`pkg/servers/webapi/http/quality.go`（新增）
- **完成内容**：`GetTestProjectWithParameter` 按当前直连 DB 风格重写——托盘号/序列号定位产品 → 定位测试工序（工位代号）→ 加载工单特性 → 按特性表达式匹配测试步骤 → 汇总测试项目与输入/输出参数。路由 `POST /api/mom/webapi/quality/gettestprojectwithparameter` 已注册。

### TASK-007 ✅ 工位生产看板接口已恢复
- **位置**：`pkg/servers/webapi/logic/production.go`、`pkg/servers/webapi/http/production.go`
- **完成内容**：`GetProductionStationExhibition` 按当前直连 DB 风格重写——工站 → 归属工序 → 产线当前型号的在制工艺路线 → 工单（含型号/类别/BOM）→ SOP 链接 → 在制/已制统计 → 返回聚合 map。路由 `GET /api/mom/webapi/production/getproductionstationexhibition` 已恢复。

### TASK-008 ✅ 测试记录上报接口已恢复
- **位置**：`pkg/servers/webapi/logic/quality.go`、`pkg/servers/webapi/http/production.go`
- **完成内容**：`CreateProductTestRecord` 编排层实现——序列号定位产品、工位代号定位工站与测试工序、兼容两种时间格式并计算耗时，落库 `ProductTestRecord`。路由 `POST /api/mom/webapi/production/createproducttestrecord` 已恢复。

### TASK-009 ✅ 工单核验失败轨迹记录已恢复
- **位置**：`pkg/servers/product/logic/product_order.go`
- **完成内容**：核验失败时用独立连接写入 `TaskQueueExecution`（Success=false、失败原因、数据跟踪索引），不随事务回滚丢失。

### TASK-010 ✅ 工艺路线统一解析策略已完成
- **位置**：`pkg/servers/webapi/logic/production.go`
- **完成内容**：抽取 `resolveProductProcessRoute(tx, productInfo, afterRouteIndex, lastProcessID, order)` 统一策略函数，上线装配与工序流转两处复用；兼容"直接创建模式"（预建路线直接返回）与"工单工艺动态创建模式"（按 ProductOrderProcess 顺序生成）。新增 `production_test.go` 覆盖两种模式（内存 SQLite）。

### TASK-011 ✅ 工序步骤匹配算法已重构
- **位置**：`pkg/model/attribute_expression_match.go`（新增）
- **完成内容**：抽取 `MatchAnyAttributeExpressions` 统一匹配函数（表达式间 OR、同表达式内特性值满足比较运算符即命中、空表达式集返回 InitialValue），四处重复嵌套循环全部替换（作业步骤匹配、测试步骤匹配、发放规则/节拍/工序匹配），并移除残留的 `fmt.Println` 调试输出。新增 9 组边界用例单测。

---

# 三、P2 — README 承诺但未实现的模块（需立项开发）

> 以下模块经全量代码搜索（模型/逻辑/proto/前端）确认无实现，按业务依赖建议排序。

### TASK-012 ✅ 质量管理系统（QM）基础版已实现
- **新增模块**：`pkg/servers/quality/`（logic/http/start）+ `pkg/model/quality_inspection_*.go` + `pkg/proto/quality_inspection.proto`（protoc 正规生成）
- **完成内容**：
  - 检验类型/检验标准（规格限/单位/比较方式/AQL/检验水平）主数据 CRUD
  - 检验单：按检验类型自动生成明细、GB/T 2828 一般检验水平II 简化抽样表确定样本量、单号自动流水
  - 完成检验：实测值按标准判定（复用 MathOperator，支持范围内/等于/大于等）、AQL 接收数判定整单结论、不合格可选生成返工记录并联动产品状态为检查中
  - SPC：均值/标准差/Cpu/Cpl/Cpk/休哈特控制限计算接口 + 单测（抽样表、接收数、Cpk 均有用例）
  - 路由：`/api/mom/quality/{qualityinspectiontype|qualityinspectionstandard|qualityinspectionorder|spc}` 
- **后续增强**（不阻塞）：让步接收工作流、质量看板实时推送、检验标准版本管理

### TASK-013 设备管理系统（EM）— 从零开发
- **现状**：无设备台账/状态机/维护模型。仅 `production_station_breakdown.go` 中有 `EquipmentID`、故障类型/原因/方案等**故障记录**字段。
- **缺失清单**：
  - [ ] 设备台账（主数据、资产、位置、责任班组）
  - [ ] 设备状态实时跟踪（与工站 `CurrentState` 联动）
  - [ ] 预防性维护计划与点检保养
  - [ ] 维修工单流程（复用故障记录数据）
  - [ ] 设备绩效分析（OEE：时间稼动率×性能稼动率×良品率）
- **工作量**：特大

### TASK-014 WMS 仓库管理系统 — 基础已有，作业流程待补齐
- **现状**（2026-08-14 更新）：PR #20 已带来 MSS 物料/仓储基础：`material_store`（仓库）、`material_shelf`/`material_shelf_bin`（货架/库位）、`material_inventory`（库存）、`material_container`（容器）、`agv_task_queue`（AGV 任务）、`wms_bill_queue`（单据队列）、退料单系列（`material_return_*`），模型与 CRUD logic 齐全。
- **仍缺失清单**：
  - [ ] 收货/上架/拣货/补料作业流程（状态机驱动，目前主要是主数据 CRUD）
  - [ ] 出入库单据与工单领料联动（对接 `product_order_bom`、`material_store_feed_rule`）
  - [ ] 盘点与库存调整流程
  - [ ] 库存周转/预警看板
  - [ ] AGV 任务下发的实际对接（驱动接口）
- **工作量**：大（基础比原评估好很多，从"从零开发"降级为"补作业流程"）

### TASK-015 APS 高级计划与排程 — 从零开发
- **现状**：全库搜索 `schedule/scheduling/排程/APS` 零命中。仅有静态的 `product_order_priority_rule.go`（优先级规则）和 `product_order_release_rule.go`（发放规则），非排程算法。
- **缺失清单**：
  - [ ] 资源与产能建模（产线、工位、班组、日历）
  - [ ] 排程引擎（约束满足或启发式算法：交期/优先级/换型时间）
  - [ ] 甘特图展示与人工干预调整
  - [ ] 排程结果下发生成/调整生产工单
  - [ ] 与 MES 实际进度回滚重排（插单、缺料响应）
- **建议**：可作为独立服务开发，复用现有 `SmartFlow` 规则引擎做规则编排
- **工作量**：特大

### TASK-016 SCADA 网关（设备数据采集）— 从零开发
- **现状**：全仓库（Go/C#/TS）搜索 `modbus/opcua/plc/s7/采集/scada` 零业务命中。`csharp/PrintService` 的 MQTT 仅用于接收打印任务。
- **缺失清单**：
  - [ ] 工业协议驱动（建议优先 Modbus，社区库已有 `github.com/CloudSilk/pkg/modbus` 基础；再扩展 OPC UA）
  - [ ] 点位（Tag）配置管理与采集任务调度
  - [ ] 设备连接管理（断线重连、状态上报，对接 TASK-013 设备管理）
  - [ ] 采集数据清洗后进入规则引擎（可复用 `SmartFlow`/RuleGo）与 MES 报工流程
  - [ ] 边缘网关部署形态（Agent 模式 + 中心配置下发）
- **工作量**：特大

### TASK-017 生产监控大屏后端聚合接口
- **现状**：MES 无实时监控数据聚合接口（工位看板接口已在 TASK-007 恢复，但厂级/产线级大屏无数据源）。
- **缺失清单**：
  - [ ] 产线级 OEE/节拍/产量达成率聚合接口
  - [ ] 厂级生产进度、异常告警（复用 `production_station_alarm`）汇总接口
  - [ ] WebSocket/SSE 实时推送
- **前置依赖**：TASK-007
- **工作量**：大

---

# 四、P3 — 前端未完成功能（web/）

### TASK-018 首页 Home 页面导入路径断裂（坏页面）
- **位置**：`web/src/pages/Home/index.tsx` — `import Editor from '@/form/components/common/AtaliEditor/editor'`，但 `src/form` 目录不存在（只有 `src/pages/form`）。路由 `/home` 已注册，运行时必报错。
- **验收标准**：修正导入或移除该页面与路由。
- **工作量**：小

### TASK-019 点击菜单时多页签（Tab）不刷新
- **位置**：`web/src/app.tsx:167` — `// TODO 如何在这边点击的时候更新tab`
- **验收标准**：点击菜单项时已打开的 Tab 正确刷新/激活。
- **工作量**：中

### TASK-020 权限控制形同虚设
- **位置**：`web/src/access.ts` 定义了 `canSeeAdmin`，但所有路由均无 `access` 字段，权限从未生效。
- **验收标准**：路由/菜单接入 usercenter RBAC 权限（菜单权限由后端 curd 元数据下发）。
- **工作量**：中

### TASK-021 服务器地址硬编码
- **位置**：
  - `web/src/pages/bpm/DesignerPage/index.tsx:212`：`fileUrlPrefix='http://101.132.37.232'`
  - `web/src/pages/cell/edit/index.tsx:107`：`cellCache.init('','http://101.132.37.232')`
  - `web/config/proxy.ts:12`：prod target 同 IP
- **验收标准**：全部改为环境变量配置。
- **工作量**：小

### TASK-022 修改密码失败提示类型错误（Bug）
- **位置**：`web/src/components/RightContent/AvatarDropdown.tsx:106` — 失败时调用 `message.success('密码修改失败!')`，应为 `message.error`。
- **工作量**：小

### TASK-023 "个人中心"菜单无点击处理
- **位置**：`web/src/components/RightContent/AvatarDropdown.tsx` — `onMenuClick` 无 `key === "center"` 分支，点击无响应。
- **验收标准**：实现跳转个人中心页（或移除该菜单项）。
- **工作量**：小

### TASK-024 BPM 流程设计器右侧菜单为占位样例
- **位置**：`web/src/pages/bpm/DesignerPage/index.tsx:25-52` — antd 样例菜单（"Navigation One/Two"、"Option 1~12"），onClick 仅 console.log。
- **验收标准**：替换为真实的流程组件面板（节点/网关/属性配置）。
- **工作量**：大

### TASK-025 AI 对话面板使用假数据
- **位置**：`web/src/pages/ResizablePanel/index.tsx:26-30` — `ProChat` request 返回写死的模拟回复，未对接真实 AI 接口（usercenter 已有 OpenAI 兼容网关 `/v1/chat/completions` 可直接对接）。
- **工作量**：中

### TASK-026 多语言切换入口被注释禁用
- **位置**：`web/src/components/RightContent/index.tsx:25` — `<SelectLang/>` 被注释，用户无法切换语言。
- **验收标准**：恢复入口并确认 locale 文件完整（见 TASK-032）。
- **工作量**：小

### TASK-027 登录页国际化失效
- **位置**：`web/src/pages/user/Login/index.tsx:28-29` — `formatMessage` 直接 `return defaultMessage`，真实 `intl.formatMessage` 被注释。
- **工作量**：小

### TASK-028 Dashboard 兜底 formID 硬编码
- **位置**：`web/src/pages/dashboard/index.tsx` — URL 无 `formID` 时硬编码 `formID="7877b188-2593-4c1c-bb1e-7ca7eb9dc0f5"`，跨环境指向错误表单。
- **验收标准**：改为按环境配置或后端默认表单接口获取。
- **工作量**：小

---

# 五、P4 — 代码质量、死代码与国际化

### TASK-029 清理后端陈旧 TODO 标记（5 处）
- **位置**：`pkg/servers/webapi/logic/production.go` 行 106、130、167、201、1660 — TODO 注释下方代码已实现所描述动作，属标记未清理。
- **工作量**：小

### TASK-030 恢复被注释的必填校验与状态分支
- `pkg/servers/product_base/logic/product_model.go:82`：`ProductCategoryID` 非空校验被注释（创建型号时类别可不填）；
- `pkg/servers/product/logic/product_order.go:114-116`：`if false {...}` 死代码，工单状态固定为 `Uploaded`，确认原意后恢复或删除。
- **工作量**：小

### TASK-031 前端死代码清理
- **清单**：
  - `web/src/services/demo/`（OneAPI 生成的样例，无引用）
  - `web/src/components/Footer/`（版权年份过期 2023、外链指向其它系统、已被注释禁用）
  - `web/src/components/Guide/`、`HeaderSearch/`、`HeaderContent/`（无引用）
  - `web/src/pages/editor/index.tsx`（无路由，内容为无关的测试数据）
  - `web/src/pages/dashboard/data.ts`（mock 聊天数据，无引用）
  - `web/src/pages/block/index.tsx`（Blocksuite POC，直接 appendChild 到 body，脱离 React 体系——移除或正式集成）
  - 9 处 `console.log` 调试残留（`app.tsx:58`、`bpm/DesignerPage:115` 等）
- **工作量**：小

### TASK-032 国际化补全与裁剪
- **清单**：
  - zh-TW 相对 zh-CN 缺失约 35 个 key（整份 `menu.*`、`app.settings.*`、`app.setting.*`、`component.globalHeader.*`、`app.pwa.*` 等）
  - en-US `pages.layouts.userLayout.title` 仍是 antd 脚手架原文，zh-CN 已改为"智能工厂"
  - 所有 locale 均为 ant-design-pro 样板文案，与本项目实际页面不符；MOM 业务术语（生产/物料/标签/追溯）完全未做国际化
  - `web/src/pages/404.tsx` 文案硬编码英文未走国际化
- **工作量**：中

### TASK-033 接口调用未纳入 service 层
- **位置**：`web/src/app.tsx:26-38`、`AvatarDropdown.tsx:15-36` — 裸用 `umiRequest` 硬写 URL（`/api/core/auth/user/profile` 等）。
- **验收标准**：统一收敛到 service 层。
- **工作量**：小

---

# 六、社区需求跟踪（对应 GitHub Issues）

| Issue | 需求 | 状态 | 对应任务/说明 |
|-------|------|------|---------------|
| [#12](https://github.com/CloudSilk/CloudSilk/issues/12) | 产品类别支持添加多个产品特性 | 🔄 PR #23 review 中 | 模型层关联字段已有，缺 proto/接口/前端配置入口 |
| [#8](https://github.com/CloudSilk/CloudSilk/issues/8) | 产品特性新建 = 基础信息 + 映射关系组合 | 📋 待排期 | 特性维护页拆分页签，映射对接上层服务描述 |
| [#6](https://github.com/CloudSilk/CloudSilk/issues/6) | 特性配置增加 Group 概念（or/and 嵌套分组） | 📋 待排期 | 需扩展特性表达式解析器与前端配置器 |
| [#5](https://github.com/CloudSilk/CloudSilk/issues/5) | MQTT 推送生产过程数据给 BI/质量等系统 | 📋 待排期 | 与 TASK-017 生产监控/事件推送合并考虑 |

# 七、任务统计与建议实施顺序

| 优先级 | 任务数 | 说明 |
|--------|--------|------|
| P0 | 4 (TASK-001/002/004/005) | 直接影响生产数据正确性，建议立即排期（TASK-003 已修复） |
| P1 | 6 (TASK-006~011) | 恢复/补齐被禁用链路，投入产出比最高 |
| P2 | 6 (TASK-012~017) | 大模块立项，建议按 质量→设备→WMS作业流程→SCADA→APS 顺序分期 |
| P3 | 11 (TASK-018~028) | 前端问题，其中 018/021/022 为快速修复项 |
| P4 | 5 (TASK-029~033) | 技术债清理，可穿插进行 |

**建议第一阶段（快速见效，约 1~2 周）**：TASK-003、005、008、009、018、021、022、023、029、030、031 — 全部为小工作量修复。

**建议第二阶段（核心补全）**：TASK-001、002、004、006、007、010、011、017、019、020、024、025。

**第三阶段（大模块分期立项）**：TASK-012 质量管理 → TASK-013 设备管理 → TASK-014 WMS → TASK-016 SCADA → TASK-015 APS。

**文档修正**：在上述模块落地前，建议同步修订 README.md 中对 APS/WMS/SCADA/设备管理/质量管理的描述，标注"规划中"，避免宣传与实现不符。
