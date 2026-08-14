# 新模块接入指南（QM / EM / WMS / APS / SCADA / 监控）

v0.2.0 新增了 6 个后端模块。本文说明各模块的 API 入口，以及如何将它们接入前端。

## 前端页面

| 功能 | 路由 | 说明 |
|------|------|------|
| 生产监控大屏 | `/monitoring` | 厂级总览指标卡 + 未处理告警实时表（5 秒轮询） |
| APS 排程 | `/aps` | 生成排程（换型/日窗口参数）、计划列表、CSS 甘特图、下发 |

两个页面均为手写页面（`web/src/pages/monitoring`、`web/src/pages/aps`），直接调用后端 REST API，已注册路由。

## CRUD 管理页（低代码方式）

质量检验单、设备台账/维保、SCADA 设备/点位等主数据维护页面**推荐用平台低代码能力**接入，无需写代码：

1. 以管理员登录，进入低代码页面管理（curd 元数据）
2. 新建页面，`path` 填后端路由前缀，例如：
   - 质量检验类型：`quality/qualityinspectiontype`
   - 质量检验标准：`quality/qualityinspectionstandard`
   - 质量检验单：`quality/qualityinspectionorder`
   - 设备台账：`equipment/equipment`
   - 维保计划：`equipment/maintenplan`
   - 采集设备：`scada/device`、点位：`scada/tag`
3. 各实体均遵循平台 CRUD 约定（`add/update/query/all/detail/delete`），表单字段与 Swagger 模型一一对应

完整字段与参数见 Swagger 文档：`/swagger/mom/index.html`（或 `docs/swagger.json`）。

## API 一览（摘要）

| 模块 | 前缀 | 关键端点 |
|------|------|---------|
| 质量管理 | `/api/mom/quality` | 检验类型/标准/检验单 CRUD；`qualityinspectionorder/complete`（完成检验：判定+AQL+可选返工/让步接收）；`spc/calc`（CPK/控制限） |
| 设备管理 | `/api/mom/equipment` | 设备 CRUD + `changestate` 状态机；维保计划 `maintenplan/execute`（执行维保滚动周期）；维保记录查询；`equipment/oee`（OEE 计算） |
| WMS 作业 | `/api/mom/material/wms` | `receive` 入库、`pickbill` 工单拣货（锁定库存）、`pickbill/complete`（出库+工单联动+可选 AGV）、`pickbill/cancel`、`stocktake` 盘点、`transaction/query` 流水、`alert` 预警 |
| APS | `/api/mom/aps/schedule` | `generate`（约束排程）、`insert`（插单重排）、`release`（下发回写）、`query/detail`（甘特数据）、`void` |
| SCADA | `/api/mom/scada` | `device/*`（含 `test` 连通性测试）、`tag/*`（含 `values` 实时值、`history` 历史）；支持 modbus-tcp 与 opcua 双协议 |
| 监控 | `/api/mom/monitoring` | `overview` 厂级总览、`line` 产线监控（加权 OEE）、`alarms` 告警、`stream` SSE 推送 |


## 配置项（系统参数表 SystemParamsConfig）

### 敏感操作授权（code=`permission`）

| key | 控制的操作 | 值 |
|-----|-----------|-----|
| `aps.manage` | 排程生成/插单/下发/作废/删除 | 角色ID逗号分隔 |
| `quality.judge` | 完成检验（判定/让步/返工） | 同上 |
| `wms.operate` | 入库/拣货/取消/盘点 | 同上 |
| `equipment.manage` | 设备状态流转/维保执行 | 同上 |
| `scada.manage` | 采集设备/点位增删改 | 同上 |

超管角色（`config.yaml superAdminRoleID`）始终放行；**参数未配置时默认拒绝**（安全缺省）；参数变更约 1 分钟内生效。

### SCADA MQTT 转发（code=`scada`）

| key | 说明 | 默认 |
|-----|------|------|
| `mqtt.broker` | MQTT 地址（如 `tcp://127.0.0.1:1883`），**空=禁用转发** | 空 |
| `mqtt.username` / `mqtt.password` | 认证（可选） | 空 |
| `mqtt.topicPrefix` | 主题前缀 | `cloudsilk/scada` |

启用后每个采集成功的点位按 `{prefix}/{设备代号}/{点位名}` 发布 JSON：`{"value":"...","quality":"Good","collectTime":"..."}`。SmartFlow（RuleGo）等规则引擎订阅该主题即可实现"采集→规则→联动"闭环。
