# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 项目概述
这是一个基于 Gin + Vue 的全栈开发平台 (gin-vue-admin) 的后端服务，提供完整的管理系统功能包括用户管理、权限管理、代码生成等功能。

## 常用命令

### 开发环境
- `go run main.go` - 启动开发服务器
- `go mod tidy` - 整理依赖包
- `go mod download` - 下载依赖包

### 测试
- `go test ./...` - 运行所有测试
- `go test ./utils/...` - 运行 utils 包的测试
- `go test -v ./service/system/...` - 运行系统服务的详细测试

### 构建
- `go build -o server main.go` - 构建可执行文件
- `go generate` - 生成代码（包含 swagger 文档生成）

### Swagger 文档
- `swag init` - 生成 swagger 文档
- 访问 `http://127.0.0.1:8888/swagger/index.html` 查看 API 文档

## 代码架构

### 分层架构
项目采用经典的 MVC 分层架构：
- **api/v1/** - API 层，处理 HTTP 请求和响应
- **service/** - 业务逻辑层，包含核心业务逻辑
- **model/** - 数据模型层，定义数据结构和数据库映射
- **router/** - 路由层，定义 API 路由
- **middleware/** - 中间件层，处理认证、日志、CORS 等

### 核心模块
- **core/** - 系统核心，包含服务器启动和配置初始化
- **global/** - 全局变量和配置
- **initialize/** - 系统初始化模块（数据库、Redis、路由等）
- **config/** - 配置结构体定义
- **utils/** - 工具函数库

### MCP 工具集成
项目集成了 MCP (Model Context Protocol) 工具系统：
- **mcp/** - MCP 工具实现，包括 API 生成、字典生成、菜单创建等
- MCP 服务端点：`http://127.0.0.1:8888/sse` (SSE) 和 `/message`

### 插件系统
- **plugin/** - 可扩展的插件系统
- 支持邮件、公告等功能插件
- 插件具有独立的 API、服务、路由和模型

## 数据库和存储

### 支持的数据库
- MySQL (默认)
- PostgreSQL
- SQLite
- SQL Server
- Oracle
- MongoDB (可选)

### ORM 和迁移
- 使用 GORM 作为 ORM
- 表结构在 `initialize/register_init.go` 中注册
- 初始数据在 `source/` 目录中定义

### 缓存
- 支持 Redis 单机和集群模式
- 配置位于 `config.yaml` 的 redis 部分

## 配置管理
- 主配置文件：`config.yaml`
- Docker 配置：`config.docker.yaml`
- 使用 Viper 进行配置管理
- 环境变量支持

## 代码生成
项目提供强大的代码生成功能：
- 自动生成 CRUD API
- 自动生成前端表单和表格
- 模板位于 `resource/` 目录
- 支持自定义模板

## 日志和监控
- 使用 Zap 进行结构化日志记录
- 日志文件位于 `log/` 目录
- 支持操作记录中间件进行审计

## 安全特性
- JWT 认证
- Casbin RBAC 权限控制
- API 限流
- CORS 跨域处理
- 图形验证码

## 文件存储
支持多种存储方式：
- 本地存储 (默认)
- 阿里云 OSS
- 腾讯云 COS  
- AWS S3
- MinIO
- 七牛云
- 华为云 OBS

## 小程序模块
项目现已支持微信小程序用户系统：

### 小程序相关接口
- `POST /miniprogram/login` - 小程序用户登录
- `GET /miniprogram/getUserInfo` - 获取小程序用户信息（需鉴权）
- `PUT /miniprogram/updateProfile` - 更新小程序用户资料（需鉴权）

### 小程序配置
在 `config.yaml` 中配置小程序信息：
```yaml
miniprogram:
    app-id: your-miniprogram-appid
    app-secret: your-miniprogram-secret
```

### 数据库表
- `miniprogram_users` - 小程序用户表，包含微信用户基本信息和登录状态

### 鉴权机制
- 小程序用户使用独立的鉴权中间件 `MiniprogramJWTAuth()`
- 小程序用户权限角色ID为999，与管理员用户完全隔离
- 支持可选鉴权中间件 `MiniprogramOptionalAuth()`

## 开发注意事项
- 使用 Go 1.23+ 版本
- 遵循项目的目录结构和命名规范
- 新增 API 需要添加 Swagger 注释
- 数据库变更需要更新表注册
- 权限相关变更需要更新 Casbin 策略
- 小程序功能与管理员功能完全隔离，不会相互影响