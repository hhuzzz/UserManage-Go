# 部署文档

## 开发环境部署

### 前置要求

- Go 1.25.4 或更高版本
- MySQL 5.7 或更高版本
- Git

### 快速启动

1. **克隆项目**
```bash
git clone https://github.com/hhuzzz/UserManage-Go.git
cd UserManage-Go
```

2. **安装依赖**
```bash
go mod download
```

3. **配置环境变量**
```bash
cp .env.example .env
# 编辑 .env 文件，修改数据库配置
```

4. **创建数据库**
```sql
CREATE DATABASE user_management CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

5. **初始化数据 (可选)**
```bash
go run cmd/initdata/main.go
```

6. **运行服务**
```bash
go run main.go
```

7. **访问应用**
```
http://localhost:8080
```

---

## 生产环境部署

### 1. 编译项目

```bash
# Linux/Mac
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o usermanage

# Windows
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -o usermanage.exe

# 其他平台
GOOS=darwin GOARCH=amd64 go build -o usermanage-mac
```

### 2. 配置生产环境变量

创建生产环境的 `.env` 文件：

```env
# 数据库配置
DB_HOST=your-production-db-host
DB_PORT=3306
DB_USER=production_user
DB_PASSWORD=strong_password_here
DB_NAME=user_management

# 服务器配置
SERVER_PORT=8080

# JWT 配置 (必须修改为强密钥)
JWT_SECRET=very-secure-secret-key-change-in-production-min-32-chars
JWT_EXPIRATION=86400
```

### 3. 使用 Systemd 部署 (Linux)

创建服务文件 `/etc/systemd/system/usermanage.service`:

```ini
[Unit]
Description=User Management System
After=network.target mysql.service

[Service]
Type=simple
User=www-data
WorkingDirectory=/opt/usermanage
ExecStart=/opt/usermanage/usermanage
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
```

启动服务：

```bash
# 复制文件
sudo cp usermanage /opt/usermanage/
sudo cp .env /opt/usermanage/

# 设置权限
sudo chown -R www-data:www-data /opt/usermanage
sudo chmod +x /opt/usermanage/usermanage

# 重新加载 systemd
sudo systemctl daemon-reload

# 启动服务
sudo systemctl start usermanage

# 设置开机自启
sudo systemctl enable usermanage

# 查看状态
sudo systemctl status usermanage

# 查看日志
sudo journalctl -u usermanage -f
```

### 4. 使用 Supervisor 部署

安装 Supervisor:

```bash
sudo apt-get install supervisor
# 或
sudo yum install supervisor
```

创建配置文件 `/etc/supervisor/conf.d/usermanage.conf`:

```ini
[program:usermanage]
command=/opt/usermanage/usermanage
directory=/opt/usermanage
autostart=true
autorestart=true
user=www-data
redirect_stderr=true
stdout_logfile=/var/log/usermanage.log
environment=GO_ENV="production"
```

启动服务：

```bash
sudo supervisorctl reread
sudo supervisorctl update
sudo supervisorctl start usermanage
```

### 5. 使用 Docker 部署

创建 `Dockerfile`:

```dockerfile
FROM golang:1.25.4-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o usermanage

FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /root/
COPY --from=builder /app/usermanage .
COPY --from=builder /app/.env .

EXPOSE 8080
CMD ["./usermanage"]
```

创建 `docker-compose.yml`:

```yaml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      - DB_HOST=mysql
      - DB_PORT=3306
      - DB_USER=root
      - DB_PASSWORD=rootpassword
      - DB_NAME=user_management
      - SERVER_PORT=8080
      - JWT_SECRET=your-jwt-secret-key
    depends_on:
      - mysql
    restart: always

  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: rootpassword
      MYSQL_DATABASE: user_management
    volumes:
      - mysql_data:/var/lib/mysql
    ports:
      - "3306:3306"
    restart: always

volumes:
  mysql_data:
```

启动容器：

```bash
docker-compose up -d
```

---

## 使用 Nginx 反向代理

### 安装 Nginx

```bash
sudo apt-get install nginx
# 或
sudo yum install nginx
```

### 配置 Nginx

创建配置文件 `/etc/nginx/sites-available/usermanage`:

```nginx
server {
    listen 80;
    server_name your-domain.com;

    # 重定向到 HTTPS
    return 301 https://$server_name$request_uri;
}

server {
    listen 443 ssl http2;
    server_name your-domain.com;

    # SSL 证书配置
    ssl_certificate /path/to/cert.pem;
    ssl_certificate_key /path/to/key.pem;

    # 安全头
    add_header X-Frame-Options "SAMEORIGIN";
    add_header X-Content-Type-Options "nosniff";
    add_header X-XSS-Protection "1; mode=block";

    location / {
        proxy_pass http://127.0.0.1:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

启用配置：

```bash
sudo ln -s /etc/nginx/sites-available/usermanage /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx
```

---

## 安全加固

### 1. 数据库安全

```bash
# 创建专用数据库用户
CREATE USER 'usermanage'@'localhost' IDENTIFIED BY 'strong_password';
GRANT ALL PRIVILEGES ON user_management.* TO 'usermanage'@'localhost';
FLUSH PRIVILEGES;

# 删除测试数据
DELETE FROM users WHERE email LIKE '%test%' OR email LIKE '%example%';
```

### 2. 文件权限

```bash
# 设置最小权限
chmod 600 /opt/usermanage/.env
chown root:root /opt/usermanage/.env
```

### 3. 防火墙配置

```bash
# UFW (Ubuntu)
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
sudo ufw enable

# firewalld (CentOS)
sudo firewall-cmd --permanent --add-service=http
sudo firewall-cmd --permanent --add-service=https
sudo firewall-cmd --reload
```

### 4. JWT 密钥

使用强随机密钥：

```bash
# 生成随机密钥
openssl rand -hex 32
```

将生成的密钥设置到 `JWT_SECRET` 环境变量。

---

## 监控和日志

### 应用日志

```bash
# Systemd
sudo journalctl -u usermanage -f

# Supervisor
sudo tail -f /var/log/usermanage.log
```

### 性能监控

使用 Prometheus + Grafana 监控应用性能。

---

## 备份策略

### 数据库备份

创建备份脚本 `/opt/backup/backup.sh`:

```bash
#!/bin/bash
BACKUP_DIR="/opt/backup/mysql"
DATE=$(date +%Y%m%d_%H%M%S)
mysqldump -u root -p user_management | gzip > $BACKUP_DIR/user_management_$DATE.sql.gz

# 保留最近7天的备份
find $BACKUP_DIR -name "*.sql.gz" -mtime +7 -delete
```

设置定时任务：

```bash
# 每天凌晨2点备份
crontab -e
0 2 * * * /opt/backup/backup.sh
```

---

## 故障排查

### 常见问题

1. **服务无法启动**
   - 检查端口是否被占用: `netstat -tuln | grep 8080`
   - 检查数据库连接配置
   - 查看服务日志

2. **数据库连接失败**
   - 检查 MySQL 服务状态
   - 验证数据库用户权限
   - 检查防火墙规则

3. **Token 无效**
   - 检查 JWT_SECRET 是否一致
   - 确认 Token 未过期
   - 验证时间同步

### 日志查看

```bash
# 查看最近的错误
journalctl -u usermanage --priority=err

# 查看特定时间段的日志
journalctl -u usermanage --since "1 hour ago"
```

---

## 更新部署

### 版本更新

```bash
# 1. 备份数据库
mysqldump -u root -p user_management > backup.sql

# 2. 停止服务
sudo systemctl stop usermanage

# 3. 拉取最新代码
cd /opt/usermanage
git pull origin main

# 4. 重新编译
go build -o usermanage

# 5. 启动服务
sudo systemctl start usermanage
```

### 数据库迁移

如果数据库结构有变更，需要手动执行迁移：

```bash
# 运行应用自动迁移
./usermanage
```

GORM 会自动处理表结构的更新。
