#!/bin/bash
# 多平台构建脚本

set -e

VERSION=${1:-dev}
OUTPUT_DIR="dist"

echo "🚀 开始构建 Clipboard Bridge $VERSION"
echo "================================================"

# 创建输出目录
mkdir -p $OUTPUT_DIR

# 构建函数
build() {
    local OS=$1
    local ARCH=$2
    local OUTPUT=$3
    local FLAGS=$4
    
    echo "📦 构建 $OS/$ARCH..."
    
    GOOS=$OS GOARCH=$ARCH CGO_ENABLED=1 go build $FLAGS -o $OUTPUT_DIR/$OUTPUT
    
    if [ $? -eq 0 ]; then
        echo "✅ $OS/$ARCH 构建成功"
    else
        echo "❌ $OS/$ARCH 构建失败"
        exit 1
    fi
}

# Windows
echo ""
echo "🪟 Windows 构建"
build windows amd64 "ClipboardBridge-windows-amd64.exe" "-ldflags=-H windowsgui"

# Linux
echo ""
echo "🐧 Linux 构建"
build linux amd64 "clipboard-bridge-linux-amd64" ""
build linux arm64 "clipboard-bridge-linux-arm64" ""

# macOS
echo ""
echo "🍎 macOS 构建"
build darwin amd64 "clipboard-bridge-macos-amd64" ""
build darwin arm64 "clipboard-bridge-macos-arm64" ""

echo ""
echo "================================================"
echo "✨ 所有平台构建完成！"
echo "📁 输出目录: $OUTPUT_DIR"
ls -lh $OUTPUT_DIR

