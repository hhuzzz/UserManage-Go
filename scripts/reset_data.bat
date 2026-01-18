@echo off
cd /d %~dp0..
echo ========================================
echo 重置用户数据
echo ========================================
echo.
echo 警告: 此操作将删除所有现有用户数据并创建新的测试数据!
echo.
set /p confirm="确定要继续吗? (y/n): "
if /i not "%confirm%"=="y" (
    echo 操作已取消
    pause
    exit /b 0
)

echo.
echo 正在重置数据...
echo.
go run cmd/initdata/main.go

echo.
pause
