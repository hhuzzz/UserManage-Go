# 开发指南

## 开发环境设置

### 1. 安装 Go

下载并安装 Go: https://golang.org/dl/

验证安装：

```bash
go version
```

### 2. 配置 Go 环境

设置 Go 模块代理 (国内开发者)：

```bash
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOSUMDB=sum.golang.google.cn
```

### 3. IDE 推荐

- **GoLand** (JetBrains)
- **VS Code** + Go 插件
- **Vim** + vim-go

---

## 项目架构

### 分层架构

```
┌─────────────────────────────────────┐
│         HTTP Request            │
└──────────────┬──────────────────┘
               │
               ▼
┌─────────────────────────────────────┐
│    Middleware Layer              │
│  (Auth, Logger, etc.)         │
└──────────────┬──────────────────┘
               │
               ▼
┌─────────────────────────────────────┐
│    Controller Layer             │
│  (Request/Response handling)    │
└──────────────┬──────────────────┘
               │
               ▼
┌─────────────────────────────────────┐
│      Service Layer              │
│     (Business logic)           │
└──────────────┬──────────────────┘
               │
               ▼
┌─────────────────────────────────────┐
│   Repository Layer              │
│     (Data access)              │
└──────────────┬──────────────────┘
               │
               ▼
┌─────────────────────────────────────┐
│         Database                │
└─────────────────────────────────────┘
```

### 目录职责

| 目录 | 职责 |
|------|--------|
| `config/` | 配置管理和加载 |
| `models/` | 数据模型定义 |
| `database/` | 数据库连接和初始化 |
| `controllers/` | HTTP 请求处理 |
| `services/` | 业务逻辑实现 |
| `repositories/` | 数据库操作接口 |
| `middleware/` | 请求中间件 |
| `routes/` | 路由配置 |
| `templates/` | HTML 模板 |
| `static/` | 静态资源 |
| `utils/` | 工具函数 |
| `auth/` | 认证相关逻辑 |
| `cmd/` | 命令行工具 |

---

## 开发流程

### 添加新功能

#### 示例：添加用户头像功能

1. **定义数据模型**

在 `models/user.go` 中添加字段：

```go
type User struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Name      string    `json:"name" gorm:"type:varchar(100);not null"`
    Email     string    `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
    Password  string    `json:"-" gorm:"type:varchar(255);not null"`
    Phone     string    `json:"phone" gorm:"type:varchar(20)"`
    Age       int       `json:"age" gorm:"type:int"`
    Status    int       `json:"status" gorm:"type:int;default:1"`
    Avatar    string    `json:"avatar" gorm:"type:varchar(255)"` // 新增
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

2. **更新 Repository 接口**

在 `repositories/user_repository.go` 中添加方法：

```go
type UserRepository interface {
    Create(user *models.User) error
    FindAll() ([]models.User, error)
    FindByID(id uint) (*models.User, error)
    Update(user *models.User) error
    Delete(id uint) error
    FindByEmail(email string) (*models.User, error)
    UpdateAvatar(id uint, avatar string) error // 新增
}
```

实现方法：

```go
func (r *userRepository) UpdateAvatar(id uint, avatar string) error {
    return r.db.Model(&models.User{}).Where("id = ?", id).Update("avatar", avatar).Error
}
```

3. **更新 Service 接口**

在 `services/user_service.go` 中添加方法：

```go
type UserService interface {
    CreateUser(req *models.CreateUserRequest) (*models.User, error)
    GetAllUsers() ([]models.User, error)
    GetUserByID(id uint) (*models.User, error)
    UpdateUser(id uint, req *models.UpdateUserRequest) (*models.User, error)
    DeleteUser(id uint) error
    UpdateUserAvatar(id uint, avatar string) error // 新增
}
```

实现方法：

```go
func (s *userService) UpdateUserAvatar(id uint, avatar string) error {
    return s.repo.UpdateAvatar(id, avatar)
}
```

4. **添加 Controller 方法**

在 `controllers/user_controller.go` 中添加：

```go
func (c *UserController) UpdateAvatar(ctx *gin.Context) {
    var req struct {
        Avatar string `json:"avatar" binding:"required"`
    }

    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := c.userService.UpdateUserAvatar(req.Avatar); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "Avatar updated successfully"})
}
```

5. **配置路由**

在 `routes/routes.go` 中添加路由：

```go
func SetupRoutes(r *gin.Engine, userController *controllers.UserController, ...) {
    // ...
    api := r.Group("/api")
    {
        protected := api.Group("")
        protected.Use(middleware.AuthMiddleware(jwtManager))
        {
            // ...
            protected.POST("/users/:id/avatar", userController.UpdateAvatar)
        }
    }
}
```

6. **更新前端**

在 `templates/index.html` 中添加头像显示和上传功能。

7. **测试**

```bash
go build
go run main.go
```

---

## 代码规范

### 命名规范

- **包名**: 小写单数词 (如 `models`, `controllers`)
- **接口名**: 以 `I` 开头 (如 `IUserRepository`)
- **结构体**: 驼峰命名 (如 `UserController`)
- **常量**: 全大写下划线分隔 (如 `MAX_RETRY`)
- **私有方法**: 小写开头
- **公开方法**: 大写开头

### 注释规范

```go
// Login 验证用户凭据并返回 JWT token
// 参数: email - 用户邮箱, password - 用户密码
// 返回: user - 用户信息, token - JWT token, err - 错误信息
func (s *AuthService) Login(email, password string) (*models.User, string, error) {
    // ...
}
```

### 错误处理

```go
// 返回 JSON 错误
ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

// 返回自定义错误
ctx.JSON(http.StatusUnauthorized, gin.H{
    "error": "User not authenticated",
    "code": "UNAUTHORIZED"
})
```

---

## 测试

### 运行测试

```bash
# 运行所有测试
go test ./...

# 运行指定包的测试
go test ./services

# 显示测试覆盖率
go test -cover ./...
```

### 编写测试示例

```go
// services/user_service_test.go
package services_test

import (
    "testing"
    "hello/models"
    // ...
)

func TestCreateUser(t *testing.T) {
    // 准备测试数据
    req := &models.CreateUserRequest{
        Name:     "Test User",
        Email:    "test@example.com",
        Password: "password123",
    }

    // 执行测试
    user, err := userService.CreateUser(req)

    // 验证结果
    if err != nil {
        t.Errorf("Failed to create user: %v", err)
    }
    if user.ID == 0 {
        t.Error("User ID should not be zero")
    }
}
```

---

## 调试

### 使用 Delve 调试器

安装：

```bash
go install github.com/go-delve/delve/cmd/dlv@latest
```

调试：

```bash
# 调试主程序
dlv debug main.go

# 测试调试
dlv test github.com/hello/services
```

### 常用命令

```go
// 打印日志
log.Printf("Debug info: %+v", data)

// 格式化输出
fmt.Printf("User data: %+v\n", user)

// 检查错误
if err != nil {
    log.Printf("Error occurred: %v", err)
}
```

---

## 性能优化

### 数据库查询优化

```go
// 使用索引
db.Where("email = ?", email).First(&user)

// 限制结果数量
db.Limit(100).Find(&users)

// 预加载关联
db.Preload("Posts").Find(&users)
```

### 缓存策略

使用 Redis 缓存热点数据：

```go
// 伪代码
func (s *userService) GetUserByID(id uint) (*models.User, error) {
    // 先查缓存
    if user, found := cache.Get(id); found {
        return user, nil
    }

    // 查数据库
    user, err := s.repo.FindByID(id)
    if err != nil {
        return nil, err
    }

    // 写缓存
    cache.Set(id, user, time.Hour)
    return user, nil
}
```

---

## 常见任务

### 添加新的 API 接口

1. 在 Controller 中添加处理函数
2. 在 Service 中实现业务逻辑
3. 在 Repository 中添加数据访问方法（如需要）
4. 在 Routes 中注册路由
5. 更新 API 文档

### 修改数据库结构

1. 修改 `models/` 中的结构体定义
2. 运行应用，GORM 自动迁移
3. 备份现有数据（生产环境）
4. 测试新字段功能

### 添加中间件

```go
// middleware/logger.go
package middleware

import (
    "log"
    "time"
    "github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        start := time.Now()

        ctx.Next()

        duration := time.Since(start)
        log.Printf("%s %s - %v",
            ctx.Request.Method,
            ctx.Request.URL.Path,
            duration,
        )
    }
}
```

在路由中使用：

```go
r.Use(middleware.Logger())
```

---

## 工具和命令

### 代码格式化

```bash
# 格式化所有代码
go fmt ./...

# 检查代码格式
go fmt -d ./
```

### 代码检查

```bash
# 静态分析
go vet ./...

# 使用 golangci-lint
golangci-lint run
```

### 依赖管理

```bash
# 整理依赖
go mod tidy

# 查看依赖树
go mod graph

# 验证依赖
go mod verify
```

### 生成文档

```bash
# 生成 godoc
godoc -http=:6060

# 访问: http://localhost:6060
```

---

## Git 工作流

### 分支策略

- `main` - 主分支，稳定版本
- `develop` - 开发分支
- `feature/*` - 功能分支
- `bugfix/*` - 修复分支

### Commit 规范

```
<type>(<scope>): <subject>

<body>

<footer>
```

Type 类型：
- `feat`: 新功能
- `fix`: Bug 修复
- `docs`: 文档更新
- `style`: 代码格式调整
- `refactor`: 代码重构
- `test`: 测试相关
- `chore`: 构建/工具相关

示例：

```
feat(auth): add user registration feature

- Implement register API endpoint
- Add registration form to login page
- Update user creation flow
```

---

## 有用的资源

- [Go 官方文档](https://golang.org/doc/)
- [Gin 文档](https://gin-gonic.com/docs/)
- [GORM 文档](https://gorm.io/docs/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
