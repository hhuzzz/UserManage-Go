# 用户管理系统

一个基于 Go + Gin + GORM + MySQL 的用户管理系统，提供完整的 CRUD 操作、用户认证和现代化的 Web 界面。

## 📋 目录

- [技术栈](#技术栈)
- [项目结构](#项目结构)
- [快速开始](#快速开始)
- [功能特性](#功能特性)
- [API 文档](#api-文档)
- [数据库结构](#数据库结构)
- [测试账号](#测试账号)
- [开发说明](#开发说明)
- [常用脚本](#常用脚本)

## 技术栈

- **后端框架**: Gin (Go Web Framework)
- **ORM**: GORM
- **数据库**: MySQL
- **认证**: JWT (JSON Web Token)
- **密码加密**: bcrypt
- **前端**: Bootstrap 5 + HTML Template
- **架构**: 分层架构 (Controller-Service-Repository)

## 项目结构

```
hello/
├── auth/              # 认证模块 (JWT、认证服务)
├── cmd/               # 命令行工具
│   ├── initdata/      # 数据初始化工具
│   └── genpassword/   # 密码生成工具
├── config/            # 配置管理
├── controllers/       # 控制器层
├── database/          # 数据库连接
├── middleware/       # 中间件
├── models/           # 数据模型
├── repositories/     # 数据访问层
├── routes/           # 路由配置
├── scripts/          # 脚本文件
│   ├── reset_data.bat # Windows 数据重置脚本
│   ├── reset_data.sh  # Linux/Mac 数据重置脚本
│   ├── test_auth.bat  # Windows 认证测试脚本
│   └── test_auth.sh   # Linux/Mac 认证测试脚本
├── services/         # 业务逻辑层
├── static/           # 静态资源
├── templates/        # HTML 模板
├── utils/            # 工具函数
├── main.go           # 程序入口
├── go.mod            # Go 模块定义
├── go.sum            # 依赖锁定文件
├── .env              # 环境变量配置
├── .env.example      # 环境变量示例
└── insert_test_users.sql # 测试数据 SQL
```

## 快速开始

### 1. 克隆项目

```bash
git clone https://github.com/hhuzzz/UserManage-Go.git
cd UserManage-Go
```

### 2. 安装依赖

```bash
go mod download
```

### 3. 配置数据库

复制 `.env.example` 为 `.env` 并修改数据库配置：

```bash
cp .env.example .env
```

编辑 `.env` 文件：

```env
# 数据库配置
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=user_management

# 服务器配置
SERVER_PORT=8080

# JWT 配置
JWT_SECRET=your-secret-key-change-in-production
JWT_EXPIRATION=86400
```

### 4. 创建数据库

在 MySQL 中创建数据库：

```sql
CREATE DATABASE user_management CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 5. 初始化数据 (可选)

运行数据初始化脚本，创建测试用户：

```bash
# Windows
scripts\reset_data.bat

# Linux/Mac
./scripts/reset_data.sh

# 或直接运行
go run cmd/initdata/main.go
```

### 6. 运行项目

```bash
go run main.go
```

### 7. 访问应用

打开浏览器访问：http://localhost:8080

## 功能特性

### 用户管理
- ✅ 用户列表展示
- ✅ 创建新用户
- ✅ 编辑用户信息
- ✅ 删除用户
- ✅ 邮箱唯一性验证
- ✅ 用户状态管理

- ✅按姓名查询
- ✅分页与排序
- ✅用户详情页

### 认证功能
- ✅ 用户注册
- ✅ 用户登录
- ✅ JWT Token 认证
- ✅ 密码加密 (bcrypt)
- ✅ 修改密码
- ✅ 登出功能
- ✅ 受保护的 API 路由

### 前端体验
- ✅ 现代化 UI 界面
- ✅ 响应式设计
- ✅ AJAX 异步操作
- ✅ Toast 消息提示
- ✅ 模态框表单
- ✅ 标签页切换 (登录/注册)
- ✅ 未登录状态提示

## API 文档

详细的 API 文档请查看 [API.md](./docs/API.md)

### 主要端点概览

#### 认证接口 (公开)
- `POST /api/auth/register` - 用户注册
- `POST /api/auth/login` - 用户登录
- `POST /api/auth/logout` - 用户登出

#### 用户管理接口 (需要认证)
- `GET /api/auth/me` - 获取当前用户信息
- `POST /api/auth/change-password` - 修改密码
- `GET /api/users` - 获取所有用户
- `GET /api/users/search` - 按姓名查询用户（分页/排序）
- `GET /api/users/:id` - 获取单个用户
- `POST /api/users` - 创建用户
- `PUT /api/users/:id` - 更新用户
- `DELETE /api/users/:id` - 删除用户

## 数据库结构

### users 表

| 字段 | 类型 | 约束 | 说明 |
|------|------|--------|------|
| id | INT | PRIMARY KEY | 用户ID |
| name | VARCHAR(100) | NOT NULL | 姓名 |
| email | VARCHAR(100) | UNIQUE | 邮箱 |
| password | VARCHAR(255) | NOT NULL | 密码 (bcrypt 加密) |
| phone | VARCHAR(20) | | 电话 |
| age | INT | | 年龄 |
| status | INT | DEFAULT 1 | 状态 (1:活跃, 0:未激活) |
| created_at | DATETIME | | 创建时间 |
| updated_at | DATETIME | | 更新时间 |

## 测试账号

### 管理员账号
- 邮箱: `admin@example.com`
- 密码: `admin123`

### 普通用户 (密码都是: password123)
- 张三 - zhangsan@example.com
- 李四 - lisi@example.com
- 王五 - wangwu@example.com
- 赵六 - zhaoliu@example.com
- 钱七 - qianqi@example.com
- 孙八 - sunba@example.com
- 周九 - zhoujiu@example.com
- 吴十 - wushi@example.com
- 郑十一 - zhengshiyi@example.com
- 王十二 - wangshier@example.com

## 开发说明

### 分层架构说明

- **Controller 层**: 处理 HTTP 请求和响应
- **Service 层**: 处理业务逻辑
- **Repository 层**: 处理数据访问
- **Middleware 层**: 请求拦截和处理

### 添加新功能

1. 在 `models/` 中定义数据模型
2. 在 `repositories/` 中创建数据访问接口
3. 在 `services/` 中实现业务逻辑
4. 在 `controllers/` 中创建控制器
5. 在 `routes/` 中配置路由

## 常用脚本

### 重置数据

```bash
# Windows
scripts\reset_data.bat

# Linux/Mac
./scripts/reset_data.sh
```

### 测试认证功能

```bash
# Windows
scripts\test_auth.bat

# Linux/Mac
./scripts/test_auth.sh
```

### 生成密码哈希

```bash
go run cmd/genpassword/main.go
```

## 安全建议

1. **生产环境必须修改 JWT_SECRET**
2. **建议使用 HTTPS**
3. **Token 存储在 localStorage 中,注意 XSS 防护**
4. **密码使用 bcrypt 加密,强度足够**
5. **所有 API 接口都有认证保护**

## 许可证

MIT License

## 贡献

欢迎提交 Issue 和 Pull Request！
