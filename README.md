# go-gin-starter

基于 Gin 的 Go Web 服务开箱即用模板。适合单体服务、后台 API 和中小型业务系统。

## 特性

- 默认零外部依赖启动，`/health` 与 `/ready` 可直接用于探针
- PostgreSQL / MySQL 可切换，数据库可关闭，支持可选自动建表
- Redis 可选启用，带连接池与超时
- Zap + Lumberjack 日志，支持切割、压缩和双端输出
- 统一响应、分页、文件下载
- JWT 认证、CORS、请求日志、Panic 恢复、请求超时、简单限流
- Handler → Service → Repository 三层结构
- Dockerfile、docker-compose、Makefile、GitHub Template 就绪

## 快速开始

环境要求：Go 1.27.3+

```bash
go mod tidy
go run ./cmd/server -config config/config.yaml
```

默认配置关闭数据库和 Redis，启动后可直接访问：

```bash
curl http://localhost:8080/health
curl http://localhost:8080/ready
```

预期响应：

```json
{"code":0,"message":"success","data":{"status":"ok"}}
```

## 用模板创建新项目

GitHub 上建议把本仓库设为 Template repository，然后：

```bash
gh repo create my-service --template <owner>/go-gin-starter --private --clone
cd my-service
# 将 go.mod 中的 module 改为你的项目路径
go mod tidy
```

非 GitHub 环境：

```bash
git clone --depth 1 https://github.com/<owner>/go-gin-starter.git my-service
cd my-service
rm -rf .git
git init -b main
git add .
git commit -m "init from go-gin-starter"
git remote add origin <your-repo-url>
git push -u origin main
```

创建后记得修改：

- `go.mod` 的 `module` 路径
- `config/config.yaml` 的 `server.app_name`
- `Makefile` 的 `APP_NAME`

## 项目结构

```text
go-gin-starter/
├── cmd/server/              # 入口
├── config/                  # 配置示例
├── internal/
│   ├── config/              # 配置加载
│   ├── database/            # GORM 初始化，可关闭
│   ├── handler/             # HTTP 处理器
│   ├── logger/              # Zap 日志
│   ├── middleware/          # CORS、JWT、日志、恢复、超时、限流
│   ├── model/               # 数据模型
│   ├── redis/               # Redis 封装，可关闭
│   ├── repository/          # 数据访问
│   ├── response/            # 统一响应
│   ├── router/              # 路由注册
│   └── service/             # 业务逻辑
├── pkg/                     # 框架无关工具包
├── Dockerfile
├── docker-compose.yml
└── Makefile
```

## 配置说明

默认 `config/config.yaml` 关闭数据库：

```yaml
database:
  enabled: false
```

需要数据库时改为 `enabled: true`，并选择驱动：

```yaml
database:
  enabled: true
  driver: postgres   # postgres 或 mysql
  auto_migrate: true # 仅建议开发和演示环境开启
```

本地依赖可启动：

```bash
docker compose up -d postgres redis
```

然后按需在配置中打开 `database.enabled` 和 `redis.enabled`。

## API 示例

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| GET | `/health` | 存活检查 |
| GET | `/ready` | 就绪检查（数据库/Redis 已启用时探活） |
| POST | `/api/v1/users` | 创建用户 |
| GET | `/api/v1/users` | 用户列表 |
| GET | `/api/v1/users/:id` | 用户详情 |
| PUT | `/api/v1/users/:id` | 更新用户 |
| DELETE | `/api/v1/users/:id` | 删除用户 |

未启用数据库时，用户接口返回明确的 `database is disabled` 错误，不会导致服务崩溃。

## 统一响应

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

分页：

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [],
    "total": 0,
    "page": 1,
    "page_size": 10
  }
}
```

## 安全与生产注意事项

- 示例用户密码使用 bcrypt 哈希存储，生产环境仍建议结合密码强度策略和登录限流
- `allowed_origins` 默认 `*`，生产环境应改为明确域名
- `auto_migrate` 只适合开发环境，生产建议使用独立迁移工具
- 限流是单进程内存实现，多实例部署时应替换为 Redis 或网关限流
- JWT 示例只校验 token 是否存在，正式项目需接入完整签名校验

## 常用命令

```bash
make run
make test
make vet
make build
make docker
```

## 技术栈

- Gin
- GORM（PostgreSQL / MySQL）
- Redis (go-redis)
- Zap + Lumberjack
- validator
- YAML

## License

MIT
