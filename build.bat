@echo off
REM Windows 批处理构建脚本

setlocal enabledelayedexpansion

set VERSION=%1
if "%VERSION%"=="" set VERSION=dev

set OUTPUT_DIR=dist

echo 🚀 开始构建 Clipboard Bridge %VERSION%
echo ================================================

REM 创建输出目录
if not exist %OUTPUT_DIR% mkdir %OUTPUT_DIR%

echo.
echo 🪟 Windows 构建
set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=1
go build -ldflags="-H windowsgui" -o %OUTPUT_DIR%\ClipboardBridge-windows-amd64.exe
if %errorlevel% equ 0 (
    echo ✅ Windows/amd64 构建成功
) else (
    echo ❌ Windows/amd64 构建失败
    exit /b 1
)

echo.
echo 🐧 Linux 构建
set GOOS=linux
set GOARCH=amd64
go build -o %OUTPUT_DIR%\clipboard-bridge-linux-amd64
if %errorlevel% equ 0 (
    echo ✅ Linux/amd64 构建成功
) else (
    echo ❌ Linux/amd64 构建失败
)

set GOARCH=arm64
go build -o %OUTPUT_DIR%\clipboard-bridge-linux-arm64
if %errorlevel% equ 0 (
    echo ✅ Linux/arm64 构建成功
) else (
    echo ❌ Linux/arm64 构建失败
)

echo.
echo 🍎 macOS 构建
set GOOS=darwin
set GOARCH=amd64
go build -o %OUTPUT_DIR%\clipboard-bridge-macos-amd64
if %errorlevel% equ 0 (
    echo ✅ macOS/amd64 构建成功
) else (
    echo ❌ macOS/amd64 构建失败
)

set GOARCH=arm64
go build -o %OUTPUT_DIR%\clipboard-bridge-macos-arm64
if %errorlevel% equ 0 (
    echo ✅ macOS/arm64 构建成功
) else (
    echo ❌ macOS/arm64 构建失败
)

echo.
echo ================================================
echo ✨ 所有平台构建完成！
echo 📁 输出目录: %OUTPUT_DIR%
dir %OUTPUT_DIR%

endlocal

