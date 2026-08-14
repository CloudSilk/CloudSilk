# CloudSilk 架构约束与演进指南

## 模块划分（业务域）

| 域 | 代码位置 | 职责 | 可依赖的域 |
|----|---------|------|-----------|
| product | `pkg/servers/product` + `product_base` | 工单、产品、BOM、返工、包装 | system |
| material | `pkg/servers/material` | 物料主数据、线边、WMS 作业（入库/拣货/盘点/流水/预警） | product, system |
| quality | `pkg/servers/quality` | 检验类型/标准/检验单、AQL 判定、SPC | product, system |
| equipment | `pkg/servers/equipment` | 设备台账、状态机、维保、OEE | system |
| aps | `pkg/servers/aps` | 排程（Scheduler 算法插件 + 计划编排） | product, material, system |
| scada | `pkg/servers/scada` | 采集设备/点位、Modbus/OPC UA 采集、MQTT 转发 | system |
| monitoring | `pkg/servers/monitoring` | 大屏聚合（只读） | 各域 logic（只读） |
| label / trace / user / system | 对应目录 | 标签打印、追溯查询、用户、系统配置 | — |

依赖方向必须无环：`aps/quality/material → product → system`。

## 跨域访问规则

1. **禁止跨域直连数据表**：其他模块不得直接读写非本域的 `model.Xxx` 表。
   一律经由属主域导出的 logic 层函数（各域 `crossdomain.go` 为出口清单）：
   - product：`GetProductOrderForPick` / `GetProductInfoByIDTx` / `AddProductOrderIssued` / `MarkProductChecking` / `CreateProductReworkRecordTx` / `GetSchedulableOrders`
   - material：`LatestCompletedPickTime`
2. **跨域函数接受 `*gorm.DB`（可为事务）**：保证跨域流程的单事务语义；函数内不提交事务。
3. **只读报表例外**：`monitoring`、`statistic` 属 CQRS 读侧，允许跨域只读查询（不写任何表），
   但不得包含业务规则；规则一律下沉到属主域。
4. **新增跨域需求时**：先在属主域 `crossdomain.go` 增加语义化出口函数，再调用；
   review 时以此为检查项。

## 微服务拆分路径（演进）

当前为模块化单体（`--service_mode=ALL`）。按上述边界拆分为 dubbogo 服务的步骤：

1. 各域 `logic` 包整体迁移为独立服务，`http` 层留在网关或随服务走；
2. `crossdomain.go` 出口函数逐一替换为 RPC 接口（dubbogo triple），
   事务内调用改为 saga/本地消息表——**出口函数粒度已按此设计**（单聚合、单语义）；
3. `model.DB` 直连替换为各域独立数据源（`config.yaml dbConfigs` 已支持分库）。

## 关键扩展点

- **APS 排程算法**：`aps/logic.Scheduler` 接口 + `SetScheduler()` 注册。
  内置 `GreedyScheduler`（前向贪心：优先级→交期→型号聚簇，换型/齐套/日窗口约束）。
  引入 OR-Tools CP-SAT 等求解器时实现该接口注册即可，编排层零改动。
- **SCADA 采集转发**：采集值变更后发布 MQTT（topic `cloudsilk/scada/{device}/{tag}`），
  SmartFlow（RuleGo）等规则引擎订阅即可接入联动；未配置 broker 时自动禁用。
- **权限**：`pkg/middleware.RequirePermissionParam(key)`，
  角色白名单存于系统参数 `permission/<key>`（逗号分隔角色ID），
  超管角色（`config.yaml superAdminRoleID`）始终放行，未配置默认拒绝。

## 已知例外与技术债

- 老模块（product/production）间仍存在 `clients` RPC 与直连库混用（历史遗留，随拆分统一）；
- dubbogo v3.0.5 依赖链漏洞待大版本升级（见 TASKS.md 遗留）；
- APS 人工拖拽调整需要前端画布交互支持，后端接口已具备（insert 重排）。
