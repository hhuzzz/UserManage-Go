# API 文档

## 基础信息

- **Base URL**: `http://localhost:8080`
- **Content-Type**: `application/json`
- **认证方式**: Bearer Token (JWT)

## 认证说明

所有受保护的 API 都需要在请求头中携带 Token：

```http
Authorization: Bearer <your_token>
```

Token 通过登录接口获取，有效期 24 小时。

---

## 认证接口

### 1. 用户注册

**接口**: `POST /api/auth/register`

**说明**: 注册新用户账号

**请求头**:
```http
Content-Type: application/json
```

**请求体**:
```json
{
  "name": "张三",
  "email": "zhangsan@example.com",
  "password": "password123",
  "phone": "13800138001",
  "age": 25,
  "status": 1
}
```

**参数说明**:

| 参数 | 类型 | 必填 | 说明 |
|------|------|--------|------|
| name | string | 是 | 用户姓名 |
| email | string | 是 | 用户邮箱 (唯一) |
| password | string | 是 | 密码 (至少6位) |
| phone | string | 否 | 电话号码 |
| age | int | 否 | 年龄 |
| status | int | 否 | 状态 (1:活跃, 0:未激活), 默认 1 |

**响应示例**:

成功 (201):
```json
{
  "message": "注册成功",
  "user": {
    "id": 12,
    "name": "张三",
    "email": "zhangsan@example.com"
  }
}
```

失败 (400):
```json
{
  "error": "email already exists"
}
```

---

### 2. 用户登录

**接口**: `POST /api/auth/login`

**说明**: 用户登录获取 Token

**请求头**:
```http
Content-Type: application/json
```

**请求体**:
```json
{
  "email": "admin@example.com",
  "password": "admin123"
}
```

**参数说明**:

| 参数 | 类型 | 必填 | 说明 |
|------|------|--------|------|
| email | string | 是 | 用户邮箱 |
| password | string | 是 | 用户密码 |

**响应示例**:

成功 (200):
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "name": "系统管理员",
    "email": "admin@example.com"
  }
}
```

失败 (401):
```json
{
  "error": "invalid email or password"
}
```

---

### 3. 用户登出

**接口**: `POST /api/auth/logout`

**说明**: 用户登出 (客户端需清除 Token)

**请求头**:
```http
Authorization: Bearer <your_token>
```

**响应示例**:

成功 (200):
```json
{
  "message": "Logged out successfully"
}
```

---

### 4. 获取当前用户信息

**接口**: `GET /api/auth/me`

**说明**: 获取当前登录用户的详细信息

**请求头**:
```http
Authorization: Bearer <your_token>
```

**响应示例**:

成功 (200):
```json
{
  "id": 1,
  "name": "系统管理员",
  "email": "admin@example.com",
  "phone": "13800138000",
  "age": 35,
  "status": 1,
  "created_at": "2026-01-18T14:33:03+08:00"
}
```

失败 (401):
```json
{
  "error": "User not authenticated"
}
```

---

### 5. 修改密码

**接口**: `POST /api/auth/change-password`

**说明**: 修改当前用户密码

**请求头**:
```http
Authorization: Bearer <your_token>
Content-Type: application/json
```

**请求体**:
```json
{
  "old_password": "password123",
  "new_password": "newpassword456"
}
```

**参数说明**:

| 参数 | 类型 | 必填 | 说明 |
|------|------|--------|------|
| old_password | string | 是 | 原密码 |
| new_password | string | 是 | 新密码 (至少6位) |

**响应示例**:

成功 (200):
```json
{
  "message": "Password changed successfully"
}
```

失败 (400):
```json
{
  "error": "old password is incorrect"
}
```

---

## 用户管理接口

以下接口都需要认证，需在请求头中携带 Token。

### 1. 获取用户列表

**接口**: `GET /api/users`

**说明**: 获取所有用户列表

**请求头**:
```http
Authorization: Bearer <your_token>
```

**查询参数** (可选):

| 参数 | 类型 | 说明 |
|------|------|------|
| page | int | 页码 |
| limit | int | 每页数量 |

**响应示例**:

成功 (200):
```json
[
  {
    "id": 1,
    "name": "系统管理员",
    "email": "admin@example.com",
    "phone": "13800138000",
    "age": 35,
    "status": 1,
    "created_at": "2026-01-18T14:33:03+08:00",
    "updated_at": "2026-01-18T14:33:03+08:00"
  },
  {
    "id": 2,
    "name": "张三",
    "email": "zhangsan@example.com",
    "phone": "13800138001",
    "age": 28,
    "status": 1,
    "created_at": "2026-01-18T14:33:03+08:00",
    "updated_at": "2026-01-18T14:33:03+08:00"
  }
]
```

失败 (401):
```json
{
  "error": "Authorization header is required"
}
```

---

### 2. 获取单个用户

**接口**: `GET /api/users/:id`

**说明**: 根据 ID 获取用户详情

**请求头**:
```http
Authorization: Bearer <your_token>
```

**路径参数**:

| 参数 | 类型 | 说明 |
|------|------|------|
| id | int | 用户 ID |

**响应示例**:

成功 (200):
```json
{
  "id": 1,
  "name": "系统管理员",
  "email": "admin@example.com",
  "phone": "13800138000",
  "age": 35,
  "status": 1,
  "created_at": "2026-01-18T14:33:03+08:00",
  "updated_at": "2026-01-18T14:33:03+08:00"
}
```

失败 (404):
```json
{
  "error": "record not found"
}
```

---

### 3. 创建用户

**接口**: `POST /api/users`

**说明**: 创建新用户

**请求头**:
```http
Authorization: Bearer <your_token>
Content-Type: application/json
```

**请求体**:
```json
{
  "name": "新用户",
  "email": "newuser@example.com",
  "password": "password123",
  "phone": "13900000000",
  "age": 30,
  "status": 1
}
```

**参数说明**:

| 参数 | 类型 | 必填 | 说明 |
|------|------|--------|------|
| name | string | 是 | 用户姓名 |
| email | string | 是 | 用户邮箱 (唯一) |
| password | string | 是 | 密码 (至少6位) |
| phone | string | 否 | 电话号码 |
| age | int | 否 | 年龄 |
| status | int | 否 | 状态 (1:活跃, 0:未激活) |

**响应示例**:

成功 (200):
```json
{
  "id": 13,
  "name": "新用户",
  "email": "newuser@example.com",
  "phone": "13900000000",
  "age": 30,
  "status": 1,
  "created_at": "2026-01-18T15:00:00+08:00",
  "updated_at": "2026-01-18T15:00:00+08:00"
}
```

失败 (400):
```json
{
  "error": "email already exists"
}
```

---

### 4. 更新用户

**接口**: `PUT /api/users/:id`

**说明**: 更新用户信息

**请求头**:
```http
Authorization: Bearer <your_token>
Content-Type: application/json
```

**路径参数**:

| 参数 | 类型 | 说明 |
|------|------|------|
| id | int | 用户 ID |

**请求体**:
```json
{
  "name": "更新后的姓名",
  "email": "updated@example.com",
  "phone": "13900000001",
  "age": 31,
  "status": 1
}
```

**参数说明**:

| 参数 | 类型 | 必填 | 说明 |
|------|------|--------|------|
| name | string | 否 | 用户姓名 |
| email | string | 否 | 用户邮箱 (唯一) |
| phone | string | 否 | 电话号码 |
| age | int | 否 | 年龄 |
| status | int | 否 | 状态 (1:活跃, 0:未激活) |

**响应示例**:

成功 (200):
```json
{
  "id": 13,
  "name": "更新后的姓名",
  "email": "updated@example.com",
  "phone": "13900000001",
  "age": 31,
  "status": 1,
  "created_at": "2026-01-18T15:00:00+08:00",
  "updated_at": "2026-01-18T15:30:00+08:00"
}
```

失败 (404):
```json
{
  "error": "record not found"
}
```

---

### 5. 删除用户

**接口**: `DELETE /api/users/:id`

**说明**: 删除指定用户

**请求头**:
```http
Authorization: Bearer <your_token>
```

**路径参数**:

| 参数 | 类型 | 说明 |
|------|------|------|
| id | int | 用户 ID |

**响应示例**:

成功 (200):
```json
{
  "message": "User deleted successfully"
}
```

失败 (404):
```json
{
  "error": "record not found"
}
```

---

## 错误码说明

| HTTP 状态码 | 说明 |
|-------------|------|
| 200 | 请求成功 |
| 201 | 创建成功 |
| 400 | 请求参数错误 |
| 401 | 未认证或 Token 无效 |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |

## 使用示例

### cURL 示例

#### 注册
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "测试用户",
    "email": "test@example.com",
    "password": "password123"
  }'
```

#### 登录
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "admin123"
  }'
```

#### 获取用户列表 (需要 Token)
```bash
TOKEN="your_token_here"
curl -X GET http://localhost:8080/api/users \
  -H "Authorization: Bearer $TOKEN"
```

#### 创建用户 (需要 Token)
```bash
TOKEN="your_token_here"
curl -X POST http://localhost:8080/api/users \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "新用户",
    "email": "new@example.com",
    "password": "password123"
  }'
```

### JavaScript/Fetch 示例

#### 登录
```javascript
async function login(email, password) {
  const response = await fetch('/api/auth/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password })
  });

  const data = await response.json();
  if (response.ok) {
    // 存储 Token
    localStorage.setItem('token', data.token);
    return data;
  } else {
    throw new Error(data.error);
  }
}
```

#### 获取用户列表
```javascript
async function getUsers() {
  const token = localStorage.getItem('token');
  const response = await fetch('/api/users', {
    headers: { 'Authorization': `Bearer ${token}` }
  });
  return await response.json();
}
```

#### 创建用户
```javascript
async function createUser(userData) {
  const token = localStorage.getItem('token');
  const response = await fetch('/api/users', {
    method: 'POST',
    headers: {
      'Authorization': `Bearer ${token}`,
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(userData)
  });
  return await response.json();
}
```

---

## 注意事项

1. **Token 有效期**: 24 小时，过期后需重新登录
2. **密码安全**: 所有密码使用 bcrypt 加密存储
3. **邮箱唯一**: 系统中每个邮箱只能注册一次
4. **认证保护**: 除了登录和注册接口，其他所有接口都需要认证
5. **生产环境**: 请修改 JWT_SECRET 为强密钥
