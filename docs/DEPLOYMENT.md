# CloudSilk 部署文档

本文档介绍 CloudSilk MOM 系统的部署方法。

## 系统要求

- **操作系统:** Linux (推荐 Ubuntu 20.04+) / Windows 10+
- **内存:** 最少 4GB，推荐 8GB+
- **磁盘:** 最少 10GB 可用空间
- **Go 版本:** 1.20+ (仅开发环境需要)

## 快速开始

### 方式一：Docker 部署（推荐）

#### 1. 构建镜像

```bash
docker build -t cloudsilk/mom:latest .
```

#### 2. 运行容器

```bash
docker run -d \
  --name cloudsilk \
  -p 20000:20000 \
  -v $(pwd)/data:/workspace/data \
  -v $(pwd)/config.yaml:/workspace/config.yaml \
  cloudsilk/mom:latest
```

#### 3. 查看日志

```bash
docker logs -f cloudsilk
```

#### 4. 停止服务

```bash
docker stop cloudsilk
docker rm cloudsilk
```

### 方式二：源码编译

#### 1. 克隆仓库

```bash
git clone https://github.com/CloudSilk/CloudSilk.git
cd CloudSilk
```

#### 2. 安装依赖

```bash
go mod download
```

#### 3. 编译

```bash
# Linux/macOS
go build -o CloudSilk main.go

# Windows
go build -o CloudSilk.exe main.go
```

#### 4. 运行

```bash
# Linux/macOS
./CloudSilk

# Windows
CloudSilk.exe
```

### 方式三：Windows 快速启动

项目提供了 Windows 启动脚本：

```bash
start.bat
```

## 配置说明

配置文件：`config.yaml`

### 数据库配置

```yaml
dbConfigs:
  # Usercenter 用户中心数据库
  Usercenter:
    fileName: "./CloudSilk.s3db"  # SQLite 文件路径
    dbType: "sqlite"
  
  # Curd 配置数据库
  Curd:
    fileName: "./CloudSilk.s3db"
    dbType: "sqlite"
  
  # CloudSilk 主数据库（MySQL）
  CloudSilk:
    connectionStr: "root:password@(127.0.0.1:3306)/mom?charset=utf8mb4&parseTime=True&loc=Local"
    dbType: "mysql"
  
  # 默认数据库
  all:
    fileName: "./CloudSilk.s3db"
    dbType: "sqlite"
```

### 其他配置

```yaml
debug: true  # 调试模式

# Token 配置
token:
  key: "CloudSilk"           # Token 密钥
  redisAddr: ""              # Redis 地址（可选）
  redisUserName: ""          # Redis 用户名
  redisPwd: ""               # Redis 密码
  expired: 86400             # Token 过期时间（秒）

# 角色配置
superAdminRoleID: "1"                              # 超级管理员角色 ID
platformTenantID: "2012ae3c-6f5f-4d10-9a59-ca0dd7c058a5"  # 平台租户 ID
defaultRoleID: 7434dc3b-fabc-4ada-9332-6ddf30e0356a       # 默认角色 ID
defaultPwd: CloudSilk    # 默认密码

# 小程序配置
miniApp:
  - id: ""
    name: ""
    secret: ""
    tenantID: ""
```

## 默认登录信息

- **默认密码:** `CloudSilk`
- **超级管理员角色 ID:** `1`

## 端口说明

| 端口 | 用途 |
|------|------|
| 20000 | HTTP API 服务端口 |

## 常见问题

### 1. 无法登录

**问题：** 部署后无法登录系统

**解决方案：**
1. 检查默认密码是否为 `CloudSilk`
2. 确认数据库连接正常
3. 查看日志文件排查错误

### 2. 数据库连接失败

**问题：** MySQL 连接失败

**解决方案：**
1. 检查 MySQL 服务是否运行
2. 确认 `config.yaml` 中的连接字符串正确
3. 检查防火墙设置

### 3. 端口被占用

**问题：** 20000 端口被占用

**解决方案：**
1. 修改 `config.yaml` 中的端口配置
2. 或者停止占用端口的服务

## 开发环境

### 安装 Dubbo-go

```bash
go get github.com/apache/dubbo-go@v3
```

### 生成 Swagger 文档

```bash
swag init -g main.go -o ./docs
```

### 运行测试

```bash
go test ./...
```

## 生产环境建议

1. **修改默认密码** - 部署后立即修改 `defaultPwd`
2. **关闭调试模式** - 设置 `debug: false`
3. **使用 MySQL** - 生产环境建议使用 MySQL 而非 SQLite
4. **配置 Redis** - 启用 Redis 缓存提升性能
5. **启用 HTTPS** - 使用反向代理（如 Nginx）配置 HTTPS
6. **定期备份** - 定期备份数据库文件

## 相关文档

- [API 文档](./swagger.yaml)
- [用户中心文档](https://github.com/CloudSilk/usercenter)
- [低代码平台](https://github.com/CloudSilk/SwiftEase)

## 支持与反馈

- **GitHub Issues:** https://github.com/CloudSilk/CloudSilk/issues
- **Discord:** https://discord.gg/AXgZhNPv
- **邮箱:** guoxf@swtsoft.com

---

_最后更新：2026-03-08_
