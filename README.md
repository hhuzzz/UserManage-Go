# 用户管理系统

一个基于 Go + Gin + GORM + MySQL 的用户管理系统，提供完整的 CRUD 操作和现代化的 Web 界面。

## 技术栈

- **后端框架**: Gin (Go Web Framework)
- **ORM**: GORM
- **数据库**: MySQL
- **前端**: Bootstrap 5 + HTML Template
- **架构**: 分层架构 (Controller-Service-Repository)

## 项目结构

```
hello/
├── config/          # 配置管理
├── models/          # 数据模型
├── database/        # 数据库连接
├── controllers/     # 控制器层
├── repositories/    # 数据访问层
├── services/        # 业务逻辑层
├── routes/          # 路由配置
├── templates/       # HTML 模板
├── static/          # 静态资源
├── main.go          # 程序入口
└── .env            # 环境变量配置
```

## 快速开始

### 1. 安装依赖

```bash
go mod tidy
```

### 2. 配置数据库

复制 `.env.example` 为 `.env` 并修改数据库配置：

```bash
cp .env.example .env
```

编辑 `.env` 文件：

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=user_management
SERVER_PORT=8080
```

### 3. 创建数据库

在 MySQL 中创建数据库：

```sql
CREATE DATABASE user_management CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 4. 运行项目

```bash
go run main.go
```

### 5. 访问应用

打开浏览器访问：http://localhost:8080

## 功能特性

- ✅ 用户列表展示
- ✅ 创建新用户
- ✅ 编辑用户信息
- ✅ 删除用户
- ✅ 邮箱唯一性验证
- ✅ 现代化 UI 界面
- ✅ 响应式设计
- ✅ AJAX 异步操作
- ✅ Toast 消息提示
- ✅ 模态框表单

## API 接口

### 用户管理接口

- `GET /` - 用户列表页面
- `GET /api/users` - 获取所有用户 (JSON)
- `GET /api/users/:id` - 获取单个用户 (JSON)
- `POST /api/users` - 创建用户 (JSON)
- `PUT /api/users/:id` - 更新用户 (JSON)
- `DELETE /api/users/:id` - 删除用户 (JSON)

## 数据库表结构

### users 表

| 字段 | 类型 | 说明 |
|------|------|------|
| id | INT (主键) | 用户ID |
| name | VARCHAR(100) | 姓名 |
| email | VARCHAR(100) | 邮箱 (唯一) |
| phone | VARCHAR(20) | 电话 |
| age | INT | 年龄 |
| status | INT | 状态 (1:活跃, 0:未激活) |
| created_at | DATETIME | 创建时间 |
| updated_at | DATETIME | 更新时间 |

## 开发说明

### 分层架构说明

- **Controller 层**: 处理 HTTP 请求和响应
- **Service 层**: 处理业务逻辑
- **Repository 层**: 处理数据访问

### 添加新功能

1. 在 `models/` 中定义数据模型
2. 在 `repositories/` 中创建数据访问接口
3. 在 `services/` 中实现业务逻辑
4. 在 `controllers/` 中创建控制器
5. 在 `routes/` 中配置路由

## 许可证

MIT License
