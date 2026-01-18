#!/bin/bash

# 切换到项目根目录
cd "$(dirname "$0")/.."

echo "========================================"
echo "重置用户数据"
echo "========================================"
echo ""
echo "警告: 此操作将删除所有现有用户数据并创建新的测试数据!"
echo ""
read -p "确定要继续吗? (y/n): " confirm
if [ "$confirm" != "y" ] && [ "$confirm" != "Y" ]; then
    echo "操作已取消"
    exit 0
fi

echo ""
echo "正在重置数据..."
echo ""
go run cmd/initdata/main.go
