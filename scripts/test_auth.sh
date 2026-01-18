#!/bin/bash

# 认证功能测试脚本

echo "======================================"
echo "认证功能测试"
echo "======================================"
echo ""

# 服务器地址
SERVER="http://localhost:8080"

echo "1. 测试登录接口..."
TOKEN_RESPONSE=$(curl -s -X POST $SERVER/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "password123"
  }')

echo "$TOKEN_RESPONSE" | grep -q "token"

if [ $? -eq 0 ]; then
  echo "✓ 登录成功"
  TOKEN=$(echo $TOKEN_RESPONSE | grep -o '"token":"[^"]*' | cut -d'"' -f4)
  echo "  Token: ${TOKEN:0:50}..."
else
  echo "✗ 登录失败"
  echo "  响应: $TOKEN_RESPONSE"
  exit 1
fi

echo ""
echo "2. 测试获取当前用户信息..."
USER_RESPONSE=$(curl -s -X GET $SERVER/api/auth/me \
  -H "Authorization: Bearer $TOKEN")

echo "$USER_RESPONSE" | grep -q "email"

if [ $? -eq 0 ]; then
  echo "✓ 获取用户信息成功"
  echo "  响应: $USER_RESPONSE"
else
  echo "✗ 获取用户信息失败"
  echo "  响应: $USER_RESPONSE"
fi

echo ""
echo "3. 测试修改密码..."
PWD_RESPONSE=$(curl -s -X POST $SERVER/api/auth/change-password \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "old_password": "password123",
    "new_password": "newpassword123"
  }')

echo "$PWD_RESPONSE" | grep -q "success\|changed"

if [ $? -eq 0 ]; then
  echo "✓ 修改密码成功"
else
  echo "✗ 修改密码失败"
  echo "  响应: $PWD_RESPONSE"
fi

echo ""
echo "4. 测试受保护的用户列表接口..."
USERS_RESPONSE=$(curl -s -X GET $SERVER/api/users \
  -H "Authorization: Bearer $TOKEN")

echo "$USERS_RESPONSE" | grep -q "\[\]"

if [ $? -eq 0 ]; then
  echo "✓ 获取用户列表成功"
else
  echo "✗ 获取用户列表失败"
  echo "  响应: $USERS_RESPONSE"
fi

echo ""
echo "5. 测试未认证访问 (应该失败)..."
FAIL_RESPONSE=$(curl -s -X GET $SERVER/api/users)

echo "$FAIL_RESPONSE" | grep -q "Unauthorized\|authorization"

if [ $? -eq 0 ]; then
  echo "✓ 未认证访问被正确拒绝"
else
  echo "✗ 安全检查失败"
  echo "  响应: $FAIL_RESPONSE"
fi

echo ""
echo "======================================"
echo "测试完成"
echo "======================================"
