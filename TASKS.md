# 云梭 MOM 生产运营管理系统 — 任务清单（✅ 全部完成）

> 生成日期：2026-08-13；完成日期：2026-08-14（33/33 项任务全部完成并推送）
> 分析范围：后端 `pkg/`（432+ Go 文件）、前端 `web/src`（206 个 TS/TSX 文件）、`csharp/PrintService`
> 分析方法：TODO/FIXME 标记扫描、被注释代码块排查、README 承诺功能与代码实现对照、前端路由/页面完整性检查

## 总体结论（终态）

**33 项任务全部完成**：MES 核心流程补全（签派/报工去重/看板/测试上报）、五大新模块落地（质量 QM、设备 EM、WMS 作业流程、APS 排程、SCADA 采集、生产监控）、前端 11 项修复与技术债清零。系统从"MES 执行 + 低代码平台"升级为覆盖 **MES/QM/EM/WMS/APS/SCADA/监控** 的完整 MOM 基础版。

各任务详情中的"后续增强"为迭代方向，不阻塞当前功能可用性。

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
- **增强迭代（2026-08-14 已落地）**：
  - ✅ 让步接收：CompleteQualityInspectionOrderRequest.concession，判定不合格时可特采放行（结论"让步接收"，不生成返工、不改产品状态）
  - 待迭代：质量看板实时推送（可复用监控 SSE）、检验标准版本管理

### TASK-013 ✅ 设备管理系统（EM）基础版已实现
- **新增模块**：`pkg/servers/equipment/` + `pkg/model/equipment.go` + `pkg/proto/equipment.proto`
- **完成内容**：
  - 设备台账 CRUD：代号/型号/厂家/位置/责任班组/投运日期/关联生产工站，空值不覆盖工站外键
  - 设备状态机：在用/停用/维修中/报废 流转接口
  - 预防性维保计划：周期滚动（下次执行日期自动计算）、到期筛选（dueOnly）、执行维保生成记录并滚动周期，异常联动设备状态为维修中
  - 维保记录：按设备/日期区间查询
  - **OEE 计算**：时间稼动率（故障记录停机时长）× 性能稼动率（数量×平均标准节拍/运行时间）× 良品率（测试记录合格率，无测试数据时按返工折算）
- **后续增强**（不阻塞）：点检任务自动生成、备件库存联动

### TASK-014 ✅ WMS 作业流程已实现
- **新增**：`pkg/servers/material/logic/material_operation.go`、`pkg/servers/material/http/material_operation.go`、`pkg/model/material_inventory_transaction.go`、`pkg/proto/material_operation.proto`
- **完成内容**：
  - **收货入库** `POST /wms/receive`：增加账面库存（自动建账）+ 事务流水
  - **工单拣货** `POST /wms/pickbill`：按工单 BOM（需求数量×工单量）逐行锁定库存生成拣货单，缺料行跳过并提示；完成 `PUT /wms/pickbill/complete` 扣减库存、解除锁定、**联动工单发料数量与最新发料时间**；取消解锁回滚
  - **库存盘点** `PUT /wms/stocktake`：实盘调整账面、记录盘盈盘亏差异流水
  - **库存事务流水** `GET /wms/transaction/query`：入库/出库/锁定/解锁/盘点调整全量追溯（含变更前后数量）
  - **库存预警** `GET /wms/alert`：可用量低于补料规则最低库存量或已为负
  - 拣货单状态机：待拣货 → 已完成/已取消
- **增强迭代（2026-08-14 已落地）**：
  - ✅ AGV 任务联动：完成拣货勾选 createAGVTask 时自动生成待签派 AGV 搬运任务（复用 AGV 任务队列与流水号）
  - 待迭代：库位级上架指引、波次拣货

### TASK-015 ✅ APS 排程引擎基础版已实现
- **新增模块**：`pkg/servers/aps/` + `pkg/model/production_schedule.go` + `pkg/proto/production_schedule.proto`
- **完成内容**：
  - **前向贪心排程算法**（纯函数 `ForwardSchedule`，5 组单测覆盖）：优先级降序 → 交期升序 → 稳定原始顺序；工时=数量×标准节拍；支持按日工作窗口跨天顺延
  - 生成排程：按产线取已发放/已签派/生产中（按剩余量）工单，产出计划+明细时段（流水批次号）
  - **下发排程**：回写各工单预计开工/完工时间（EstimateStartTime/FinishTime，甘特基准）
  - 计划状态机：已生成 → 已下发 / 已作废
  - 查询接口含明细时段，可直接渲染甘特图
  - 路由：`/api/mom/aps/schedule/{generate|release|void|query|detail|delete}`
- **增强迭代（2026-08-14 已落地）**：
  - ✅ 约束求解器 `ConstrainedSchedule`：换型时间约束（相邻工单型号切换插入换型间隔）+ 同优先级型号聚簇减少换型 + 物料齐套约束（齐套时间取自最近已完成拣货单）+ 7 组新单测
  - ✅ 插单重排 `POST /aps/schedule/insert`：保持既有工单顺序，插入到严格更高优先级之后，重排落库
  - 待迭代：人工拖拽调整（需前端）

### TASK-016 ✅ SCADA 网关（Modbus TCP）基础版已实现
- **新增模块**：`pkg/servers/scada/` + `pkg/model/scada.go` + `pkg/proto/scada.proto`
- **完成内容**：
  - 采集设备管理：协议/地址/从站号/采集间隔/连接状态/最近错误，连通性测试接口（读保持寄存器探测）
  - 点位（Tag）管理：功能码（线圈/离散输入/保持/输入寄存器）、数据类型（bool/uint16/int16/float32）、缩放系数、单位、历史开关
  - **后台采集器**：随服务启动（sync.Once 保护），每设备独立 goroutine 按配置间隔轮询，实时值 upsert + 历史落库 + 连接状态回写（全部失败置异常并记录首错）
  - 复用 `github.com/CloudSilk/pkg/modbus` 驱动
  - 查询接口：点位实时值（按设备过滤）、点位历史值（时间区间分页）
  - 路由：`/api/mom/scada/{device|tag}/*`
- **增强迭代（2026-08-14 已落地）**：
  - ✅ **OPC UA 驱动**（gopcua）：设备协议 `opcua`（opc.tcp://endpoint），点位按节点ID（ns=2;s=...）采集，单连接顺序读取、数值/布尔/整型格式化、缩放系数；连通性测试读取 ServerStatus 节点；Modbus 与 OPC UA 双协议共存
  - ✅ 断线告警事件：连接失败翻转为异常时写 ExceptionTrace
  - 待迭代：采集数据转发 MQTT/规则引擎

### TASK-017 ✅ 生产监控大屏后端聚合接口已实现
- **新增模块**：`pkg/servers/monitoring/` + `pkg/proto/monitoring.proto`
- **完成内容**：
  - **厂级总览** `GET /monitoring/overview`：产线/工站规模、进行中工单与数量进度（下单/开工/完工合计）、未处理与今日告警数、工站状态分布（待机/作业/故障）
  - **产线监控** `GET /monitoring/line`：OEE 四指标（复用设备模块计算）、周期产量（节拍记录数）、工单达成率、在制数、未处理告警
  - **告警列表** `GET /monitoring/alarms`：按状态过滤、时间倒序、限量
  - **SSE 实时推送** `GET /monitoring/stream`：每 5 秒推送 overview 事件（text/event-stream，EventSource 兼容）
- **增强迭代（2026-08-14 已落地）**：
  - ✅ 产线级 OEE 加权平均：遍历产线全部工站按周期产量加权（无产量工站等权退化）
  - 待迭代：WebSocket 双向推送、前端大屏页面

---

# 四、P3 — 前端未完成功能（web/）

### TASK-018 ✅ 首页断裂导入 — 经核验已不存在
- **结论**：当前代码 `src/pages/Home/index.tsx` 全部为 npm 包导入（`@swiftease/*`），不存在 `@/form` 断裂路径（早期分析基于旧代码状态）。页面为完整的表单设计器，路由正常。

### TASK-019 ✅ 点击菜单多页签不刷新已修复
- **位置**：`web/src/app.tsx`
- **根因**：`tabs.push(...)` 原地修改数组后把同一引用传给 `setTabs`，React 引用相等跳过渲染。
- **修复**：改用 `setTabs([...tabs, {...item, path}])` 新数组引用；清空分支 `setTabs([])`。

### TASK-020 ✅ 权限控制已接入路由
- **位置**：`web/.umirc.ts`、`web/src/access.ts`
- **完成**：block 与 resizable 演示路由挂 `access: "canSeeAdmin"`，access.ts 权限真正生效（未授权用户 403）。

### TASK-021 ✅ 服务器地址硬编码已消除
- **位置**：`bpm/DesignerPage`、`cell/edit`、`config/proxy.ts`
- **完成**：`fileUrlPrefix`/`cellCache.init` 改用 `window.location.origin`（同源部署）；开发代理目标走 `PROXY_TARGET` 环境变量（默认 127.0.0.1:48089）。

### TASK-022 ✅ 密码修改失败提示类型已修正
- **位置**：`web/src/components/RightContent/AvatarDropdown.tsx`
- **完成**：两处失败分支 `message.success` → `message.error`。

### TASK-023 ✅ 个人中心菜单已实现跳转
- **位置**：`web/src/components/RightContent/AvatarDropdown.tsx`
- **完成**：`key === "center"` 分支跳转 `/home`。

### TASK-024 ✅ BPM 设计器右侧菜单已替换为真实操作
- **位置**：`web/src/pages/bpm/DesignerPage/index.tsx`
- **完成**：antd 样例菜单替换为画布操作菜单（撤销/重做/放大/缩小/适应画布/清空），基于 createMenu 回调注入的 X6 Graph 实例实现。

### TASK-025 ✅ AI 对话面板已对接真实接口
- **位置**：`web/src/pages/ResizablePanel/index.tsx`
- **完成**：移除模拟回复，改为调用 OpenAI 兼容端点 `/v1/chat/completions`（usercenter AI 网关），携带 token，失败时给出明确提示。

### TASK-026 ✅ 多语言切换入口已恢复
- **位置**：`web/src/components/RightContent/index.tsx`
- **完成**：取消 `<SelectLang/>` 注释，从 `@umijs/max` 导入。

### TASK-027 ✅ 登录页国际化已接通
- **位置**：`web/src/pages/user/Login/index.tsx`
- **完成**：`formatMessage` 恢复走 `intl.formatMessage`（原直接返回 defaultMessage）。

### TASK-028 ✅ Dashboard 兜底 formID 已改为环境变量配置
- **位置**：`web/.umirc.ts`、`web/src/pages/dashboard/index.tsx`
- **完成**：define 注入 `process.env.DEFAULT_FORM_ID`（默认保持原 ID，可用环境变量覆盖），消除跨环境硬编码。


# 五、P4 — 代码质量、死代码与国际化

### TASK-029 ✅ 后端陈旧 TODO 标记已清理
- **位置**：`pkg/servers/webapi/logic/production.go`
- **完成**：删除"触发事件"与"计算预计下线时间"两处陈旧 TODO（下方代码均已实现相应功能）。全后端已无 TODO 标记。

### TASK-030 ✅ 必填校验与状态分支已恢复
- **完成**：`product_model.go` 创建型号恢复"产品类别不能为空"校验、更新时空类别不覆盖外键（omit）；`product_order.go` 移除 `if false` 死代码并注释说明状态起点。

### TASK-031 ✅ 前端死代码已清理
- **删除**：`src/services/demo/`（OneAPI 样例）、`components/{Footer,Guide,HeaderSearch,HeaderContent}`、`pages/editor`（无路由）、`pages/block`（POC，含路由注册）、dashboard/data.ts 由 grep 确认无引用后移除。
- **清理**：`app.tsx`/`bpm/DesignerPage` 的 `console.log` 调试残留。TS 检查无新增错误。

### TASK-032 ✅ 国际化补全已完成
- **完成**：新增 `zh-TW/pages.ts`（65 个 key 繁体翻译）并注册到 zh-TW.ts；en-US `pages.layouts.userLayout.title` 修正为 Smart Factory；404 页文案中文化。

### TASK-033 ✅ 接口调用已收敛 service 层
- **完成**：新建 `src/services/index.ts`（queryCurrentUser/logout/changePwd，统一 umi-request + token 头），`app.tsx` 与 `AvatarDropdown` 的裸调用全部替换为 service 导入。


# 六、社区需求跟踪（对应 GitHub Issues）

| Issue | 需求 | 状态 | 对应任务/说明 |
|-------|------|------|---------------|
| [#12](https://github.com/CloudSilk/CloudSilk/issues/12) | 产品类别支持添加多个产品特性 | 🔄 PR #23 review 中 | 模型层关联字段已有，缺 proto/接口/前端配置入口 |
| [#8](https://github.com/CloudSilk/CloudSilk/issues/8) | 产品特性新建 = 基础信息 + 映射关系组合 | 📋 待排期 | 特性维护页拆分页签，映射对接上层服务描述 |
| [#6](https://github.com/CloudSilk/CloudSilk/issues/6) | 特性配置增加 Group 概念（or/and 嵌套分组） | 📋 待排期 | 需扩展特性表达式解析器与前端配置器 |
| [#5](https://github.com/CloudSilk/CloudSilk/issues/5) | MQTT 推送生产过程数据给 BI/质量等系统 | 📋 待排期 | 与 TASK-017 生产监控/事件推送合并考虑 |

# 七、任务统计（终态）

| 优先级 | 任务数 | 结果 |
|--------|--------|------|
| P0（TASK-001~005） | 5 | ✅ 全部完成（003 由 PR#17-20 修复，其余本轮实现） |
| P1（TASK-006~011） | 6 | ✅ 全部完成（含 quality.go/看板/测试上报恢复、统一策略重构 + 单测） |
| P2（TASK-012~017） | 6 | ✅ 全部完成（QM/EM/WMS/APS/SCADA/监控六大模块基础版） |
| P3（TASK-018~028） | 11 | ✅ 全部完成（前端修复与功能补全） |
| P4（TASK-029~033） | 5 | ✅ 全部完成（清理/国际化/service 收敛） |

**质量证据**：`go build ./...` 通过；4 个测试包全绿（模型匹配、SPC/AQL、排程算法、工艺路线策略）；前端 tsc 无新增错误；全部提交已推送至 GitHub main（每个任务独立提交）。
