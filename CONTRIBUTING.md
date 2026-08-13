# 贡献指南 | Contributing Guide

感谢你对云梭 CloudSilk 的关注！欢迎通过 Issue 和 Pull Request 参与共建。Issues 和 PR 都会得到维护者的认真处理。

Thank you for your interest in CloudSilk! We welcome issues and pull requests, and both are actively triaged by maintainers.

## 目录

- [报告问题](#报告问题)
- [开发环境](#开发环境)
- [如何提交 PR](#如何提交-pr)
- [代码规范](#代码规范)
- [Proto 变更说明](#proto-变更说明)
- [许可证](#许可证)

## 报告问题

提交 Issue 前请先搜索[已有 Issue](https://github.com/CloudSilk/CloudSilk/issues) 避免重复。一个高质量的 Issue 应包含：

1. **环境信息**：版本/commit、部署方式（Docker / 源码 / Windows 脚本）、数据库类型
2. **复现步骤**：最小化的操作路径
3. **预期行为与实际行为**
4. **日志或截图**

功能建议（enhancement）请描述使用场景，而不仅是解决方案——我们会结合 [TASKS.md](./TASKS.md) 路线图评估优先级。

## 开发环境

```bash
# 要求：Go 1.20+；前端构建另需 Node.js 18+

# 后端（单进程 + SQLite，开箱即用）
go build -o CloudSilk main.go
./CloudSilk --ui ./web/dist --service_mode=ALL --single_db=true --port=48900

# 前端
cd web
WEB_BASE=/web yarn install
WEB_BASE=/web yarn start        # 开发调试
WEB_BASE=/web yarn build        # 构建产物输出到 web/dist

# Swagger 文档重新生成（修改了 HTTP handler 注释后）
swag init --parseDependency --parseInternal --parseDepth 1
```

完整部署说明见 [docs/DEPLOYMENT.md](./docs/DEPLOYMENT.md)。

## 如何提交 PR

1. Fork 仓库，从最新 `main` 创建特性分支：
   ```bash
   git checkout -b fix/short-description
   ```
2. 完成改动并**确保本地可编译、可运行**：
   ```bash
   go build ./...
   ```
   > ⚠️ 提交前必须通过 `go build ./...`。历史上出现过合并后主干无法编译的情况，现在这是硬性要求。
3. 提交信息使用简洁的前缀式描述（中英文均可）：
   - `fix: 修复进站重复报工`
   - `feat: 支持库位批量导入`
   - `docs: 补充部署文档`
4. 在 PR 描述中说明：改了什么、为什么改、如何验证（最好附测试步骤或截图）。
5. 关联相关 Issue（如 `Fixes #12`）。

PR 会在合理时间内得到 review；需要修改时会留下具体意见，更新分支即可。

## 代码规范

- **Go**：`gofmt` 格式化；遵循周边代码风格（错误处理、命名、注释密度保持一致）
- **新增业务模块**遵循现有结构：`pkg/model/`（数据模型）→ `pkg/proto/`（接口定义）→ `pkg/servers/<module>/logic|http|provider/`（逻辑与路由）
- **前端**：业务页面优先通过低代码元数据（curd 引擎）配置而非硬编码页面；仅平台能力改动才需要修改 `web/src`
- 注释与文档中文为主，接口命名用英文

## Proto 变更说明

修改 `.proto` 文件后，**必须用 protoc 重新生成对应的 `.pb.go`**，并将两者一并提交：

```bash
protoc --proto_path=./pkg/proto --go_out=./pkg/proto your_file.proto
```

> ⚠️ 不要手工编辑 `.pb.go`。protobuf 的序列化依赖文件中的 raw descriptor，手改结构体字段而不重新生成，字段在 RPC 序列化时会被静默丢弃，HTTP JSON 场景又能正常工作，问题非常隐蔽。

## 许可证

提交即表示你同意贡献内容以 [Apache 2.0](./LICENSE) 许可证开源。
