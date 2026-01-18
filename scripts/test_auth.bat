@echo off
REM 认证功能测试脚本 (Windows)

echo ======================================
echo 认证功能测试
echo ======================================
echo.

REM 服务器地址
set SERVER=http://localhost:8080

echo 1. 测试登录接口...
curl -s -X POST %SERVER%/api/auth/login -H "Content-Type: application/json" -d "{\"email\": \"admin@example.com\", \"password\": \"password123\"}" > login_response.json

findstr /C:"token" login_response.json >nul
if %errorlevel% equ 0 (
  echo [OK] 登录成功
  for /f "tokens=2 delims=:," %%a in (login_response.json) do (
    set TOKEN=%%a
    goto :found_token
  )
  :found_token
  set TOKEN=%TOKEN:"=%
  echo Token: %TOKEN:~0,50%...
) else (
  echo [FAIL] 登录失败
  type login_response.json
  pause
  exit /b 1
)

echo.
echo 2. 测试获取当前用户信息...
curl -s -X GET %SERVER%/api/auth/me -H "Authorization: Bearer %TOKEN%" > user_response.json

findstr /C:"email" user_response.json >nul
if %errorlevel% equ 0 (
  echo [OK] 获取用户信息成功
  type user_response.json
) else (
  echo [FAIL] 获取用户信息失败
  type user_response.json
)

echo.
echo 3. 测试修改密码...
curl -s -X POST %SERVER%/api/auth/change-password -H "Content-Type: application/json" -H "Authorization: Bearer %TOKEN%" -d "{\"old_password\": \"password123\", \"new_password\": \"newpassword123\"}" > pwd_response.json

findstr /C:"success\|changed" pwd_response.json >nul
if %errorlevel% equ 0 (
  echo [OK] 修改密码成功
) else (
  echo [FAIL] 修改密码失败
  type pwd_response.json
)

echo.
echo 4. 测试受保护的用户列表接口...
curl -s -X GET %SERVER%/api/users -H "Authorization: Bearer %TOKEN%" > users_response.json

findstr /C:"\[\]" users_response.json >nul
if %errorlevel% equ 0 (
  echo [OK] 获取用户列表成功
) else (
  echo [FAIL] 获取用户列表失败
  type users_response.json
)

echo.
echo 5. 测试未认证访问 (应该失败)...
curl -s -X GET %SERVER%/api/users > fail_response.json

findstr /C:"Unauthorized\|authorization" fail_response.json >nul
if %errorlevel% equ 0 (
  echo [OK] 未认证访问被正确拒绝
) else (
  echo [FAIL] 安全检查失败
  type fail_response.json
)

echo.
echo ======================================
echo 测试完成
echo ======================================

REM 清理临时文件
del login_response.json user_response.json pwd_response.json users_response.json fail_response.json 2>nul

pause
