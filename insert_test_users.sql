-- 清理并初始化用户数据
-- 注意: 此脚本会删除所有现有用户,请谨慎使用!

-- 1. 清空现有用户数据
TRUNCATE TABLE users;

-- 2. 插入管理员用户
-- 密码: admin123
INSERT INTO users (name, email, password, phone, age, status, created_at, updated_at)
VALUES (
    '系统管理员',
    'admin@example.com',
    '$2a$10$dFW4KzYxKkpq29a7lgnoUOcKYHsdIdXv4xIzLjlB/wwarQjZuZQhm',
    '13800138000',
    35,
    1,
    NOW(),
    NOW()
);

-- 3. 插入普通用户 (密码都是: password123)
INSERT INTO users (name, email, password, phone, age, status, created_at, updated_at) VALUES
('张三', 'zhangsan@example.com', '$2a$10$dFW4KzYxKkpq29a7lgnoUOcKYHsdIdXv4xIzLjlB/wwarQjZuZQhm', '13800138001', 28, 1, NOW(), NOW()),
('李四', 'lisi@example.com', '$2a$10$dFW4KzYxKkpq29a7lgnoUOcKYHsdIdXv4xIzLjlB/wwarQjZuZQhm', '13800138002', 32, 1, NOW(), NOW()),
('王五', 'wangwu@example.com', '$2a$10$dFW4KzYxKkpq29a7lgnoUOcKYHsdIdXv4xIzLjlB/wwarQjZuZQhm', '13800138003', 25, 1, NOW(), NOW()),
('赵六', 'zhaoliu@example.com', '$2a$10$dFW4KzYxKkpq29a7lgnoUOcKYHsdIdXv4xIzLjlB/wwarQjZuZQhm', '13800138004', 30, 1, NOW(), NOW()),
('钱七', 'qianqi@example.com', '$2a$10$dFW4KzYxKkpq29a7lgnoUOcKYHsdIdXv4xIzLjlB/wwarQjZuZQhm', '13800138005', 27, 1, NOW(), NOW()),
('孙八', 'sunba@example.com', '$2a$10$dFW4KzYxKkpq29a7lgnoUOcKYHsdIdXv4xIzLjlB/wwarQjZuZQhm', '13800138006', 29, 1, NOW(), NOW()),
('周九', 'zhoujiu@example.com', '$2a$10$dFW4KzYxKkpq29a7lgnoUOcKYHsdIdXv4xIzLjlB/wwarQjZuZQhm', '13800138007', 31, 1, NOW(), NOW()),
('吴十', 'wushi@example.com', '$2a$10$dFW4KzYxKkpq29a7lgnoUOcKYHsdIdXv4xIzLjlB/wwarQjZuZQhm', '13800138008', 26, 1, NOW(), NOW()),
('郑十一', 'zhengshiyi@example.com', '$2a$10$dFW4KzYxKkpq29a7lgnoUOcKYHsdIdXv4xIzLjlB/wwarQjZuZQhm', '13800138009', 33, 1, NOW(), NOW()),
('王十二', 'wangshier@example.com', '$2a$10$dFW4KzYxKkpq29a7lgnoUOcKYHsdIdXv4xIzLjlB/wwarQjZuZQhm', '13800138010', 24, 1, NOW(), NOW());
