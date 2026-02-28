# Go Basic Framework Template

基于 Gin 框架的企业级 Go 项目基础模板，开箱即用。

## 特性

- **数据库适配** - 支持 PostgreSQL（优选）/ MySQL，通过配置文件一键切换
- **Redis 缓存** - 可选启用，支持连接池配置
- **日志系统** - Zap + Lumberjack，支持文件切割、压缩、控制台输出
- **统一响应** - JSON 格式 + 分页数据 + 文件流下载
- **中间件** - CORS 跨域、请求日志、Panic 恢复、JWT 认证
- **分层架构** - Handler → Service → Repository 标准三层
- **工具包** - 丰富的通用工具函数库

## 项目结构

```
go_basic_framework_template/
├── cmd/
│   └── server/
│       └── main.go              # 主入口文件
├── config/
│   ├── config.yaml              # PostgreSQL 配置示例
│   └── config.mysql.yaml        # MySQL 配置示例
├── internal/                    # 私有应用代码
│   ├── config/                  # 配置管理
│   ├── database/                # 数据库连接
│   ├── handler/                 # HTTP 处理器
│   ├── logger/                  # 日志模块
│   ├── middleware/              # 中间件
│   ├── model/                   # 数据模型
│   ├── redis/                   # Redis 缓存
│   ├── repository/              # 数据访问层
│   ├── response/                # 统一响应封装
│   ├── router/                  # 路由配置
│   └── service/                 # 业务逻辑层
├── pkg/                         # 公共工具包
│   ├── converter/               # 类型转换工具
│   ├── hash/                    # 哈希加密工具
│   ├── jwt/                     # JWT 认证工具
│   ├── slice/                   # 切片工具
│   ├── str/                     # 字符串工具
│   ├── utils/                   # 通用工具
│   └── validator/               # 验证工具
├── logs/                        # 日志文件目录
├── bin/                         # 编译输出
├── Makefile                     # 构建脚本
├── go.mod
└── go.sum
```

## 快速开始

### 环境要求

- Go 1.22+
- PostgreSQL 12+ / MySQL 8+
- Redis 6+ (可选)

### 安装依赖

```bash
go mod tidy
```

### 配置文件

编辑 `config/config.yaml`，配置数据库连接信息：

```yaml
server:
  port: 8080
  mode: debug
  app_name: go_basic_framework_template

database:
  driver: postgres # postgres 或 mysql
  host: localhost
  port: 5432
  username: postgres
  password: postgres
  database: myapp
  ssl_mode: disable
  time_zone: Asia/Shanghai
  max_idle_conns: 10
  max_open_conns: 100

redis:
  enabled: false # 是否启用 Redis
  host: localhost
  port: 6379
  password: ''
  db: 0
  pool_size: 10
  min_idle_conns: 5

log:
  level: debug
  filename: logs/app.log
  max_size: 100 # MB
  max_backups: 10
  max_age: 30 # days
  compress: true
  console: true
```

### 运行项目

```bash
# 直接运行
go run ./cmd/server

# 或编译后运行
go build -o bin/server ./cmd/server
./bin/server

# 使用 Makefile
make run
```

### 构建命令

```bash
make build        # 编译项目
make run          # 运行项目
make clean        # 清理构建产物
make test         # 运行测试
make lint         # 运行代码检查
make fmt          # 格式化代码
```

## API 接口

### 健康检查

| 方法 | 路径      | 说明     |
| ---- | --------- | -------- |
| GET  | `/health` | 健康检查 |
| GET  | `/ready`  | 就绪检查 |

### 用户管理（示例）

| 方法   | 路径                | 说明     |
| ------ | ------------------- | -------- |
| POST   | `/api/v1/users`     | 创建用户 |
| GET    | `/api/v1/users`     | 用户列表 |
| GET    | `/api/v1/users/:id` | 用户详情 |
| PUT    | `/api/v1/users/:id` | 更新用户 |
| DELETE | `/api/v1/users/:id` | 删除用户 |

### 统一响应格式

**成功响应**

```json
{
  "code": 0,
  "message": "success",
  "data": { ... }
}
```

**分页响应**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [ ... ],
    "total": 100,
    "page": 1,
    "page_size": 10
  }
}
```

**错误响应**

```json
{
  "code": 400,
  "message": "bad request"
}
```

## 工具包使用

### 字符串工具 (pkg/str)

```go
import "go_basic_framework_template/pkg/str"

str.IsEmpty("")                    // true
str.IsBlank("  ")                  // true
str.CamelToSnake("UserName")       // "user_name"
str.SnakeToCamel("user_name")      // "userName"
str.MaskPhone("13812345678")       // "138****5678"
str.MaskEmail("test@example.com")  // "t***@example.com"
```

### 切片工具 (pkg/slice)

```go
import "go_basic_framework_template/pkg/slice"

slice.Contains([]int{1, 2, 3}, 2)           // true
slice.Unique([]int{1, 2, 2, 3})             // [1, 2, 3]
slice.Chunk([]int{1, 2, 3, 4}, 2)           // [[1, 2], [3, 4]]
slice.Map([]int{1, 2, 3}, func(i int) int { return i * 2 })  // [2, 4, 6]
slice.Filter([]int{1, 2, 3}, func(i int) bool { return i > 1 })  // [2, 3]
```

### 类型转换 (pkg/converter)

```go
import "go_basic_framework_template/pkg/converter"

converter.ToInt("123")           // 123, nil
converter.ToInt64(123)           // 123, nil
converter.ToFloat64("3.14")      // 3.14, nil
converter.ToBool("true")         // true, nil
converter.ToJSON(map[string]int{"a": 1})  // "{\"a\":1}", nil
```

### 验证工具 (pkg/validator)

```go
import "go_basic_framework_template/pkg/validator"

validator.IsEmail("test@example.com")    // true
validator.IsPhone("13812345678")         // true
validator.IsIDCard("110101199001011234") // true
validator.IsURL("https://example.com")   // true
validator.IsIP("192.168.1.1")            // true
validator.IsStrongPassword("Abc123!@")   // true
```

### 哈希工具 (pkg/hash)

```go
import "go_basic_framework_template/pkg/hash"

hash.MD5String("hello")                    // "5d41402abc4b2a76b9719d911017c592"
hash.SHA256String("hello")                 // "2cf24dba5..."
hash.HashPassword("mypassword")           // "$2a$10$..." (bcrypt)
hash.CheckPassword("mypassword", hashStr) // true
```

### JWT 工具 (pkg/jwt)

```go
import "go_basic_framework_template/pkg/jwt"
import "time"

j := jwt.NewJWT("secret-key", "myapp", time.Hour, time.Hour*24*7)

// 生成 Token
token, _ := j.GenerateToken(1, "username", "email@example.com")

// 解析 Token
claims, _ := j.ParseToken(token)
// claims.UserID, claims.Username, claims.Email

// 刷新 Token
newToken, _ := j.RefreshToken(token)
```

### 通用工具 (pkg/utils)

```go
import "go_basic_framework_template/pkg/utils"

utils.RandomString(16)           // "aBc123XyZ..."
utils.RandomNumber(6)            // "123456"
utils.Paginate(2, 10)            // offset=10, limit=10
utils.If(true, "yes", "no")      // "yes"
utils.Default(0, 100)            // 100
utils.Default(50, 100)           // 50
```

## 中间件

### CORS 跨域

已内置 CORS 中间件，默认允许所有来源。可在 `internal/middleware/cors.go` 中修改配置。

### 请求日志

自动记录每个请求的方法、路径、状态码、耗时等信息。

### Panic 恢复

自动捕获 panic 并返回 500 错误，防止服务崩溃。

### JWT 认证

```go
import "go_basic_framework_template/internal/middleware"

// 需要认证的路由
api.GET("/profile", middleware.AuthRequired(), handler.Profile)
```

## 扩展指南

### 添加新的业务模块

1. 在 `internal/model/` 创建数据模型
2. 在 `internal/repository/` 创建数据访问层
3. 在 `internal/service/` 创建业务逻辑层
4. 在 `internal/handler/` 创建 HTTP 处理器
5. 在 `internal/router/router.go` 注册路由

### 添加新的中间件

在 `internal/middleware/` 创建新的中间件文件，并在 `router.go` 中注册。

### 添加新的工具函数

在 `pkg/` 对应的工具包中添加函数，并添加注释说明。

## 技术栈

- [Gin](https://github.com/gin-gonic/gin) - Web 框架
- [GORM](https://gorm.io/) - ORM 框架
- [Zap](https://github.com/uber-go/zap) - 日志库
- [Lumberjack](https://github.com/natefinch/lumberjack) - 日志切割
- [go-redis](https://github.com/redis/go-redis) - Redis 客户端
- [jwt-go](https://github.com/golang-jwt/jwt) - JWT 认证
- [validator](https://github.com/go-playground/validator) - 参数验证
- [yaml](https://github.com/go-yaml/yaml) - YAML 解析

## License

MIT License
